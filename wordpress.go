package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-resty/resty/v2"
	"golang.org/x/sync/errgroup"
)

type WordpressClient interface {
	CreatePost(title, content string, media []string) (string, error)
}

type wordpressClient struct {
	Endpoint    string
	restyClient *resty.Client
}

func NewWordpressClient(endpoint, username, password string) WordpressClient {
	restyClient := resty.New()

	restyClient.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		c.SetBasicAuth(username, password)
		return nil
	})
	return &wordpressClient{
		Endpoint:    endpoint,
		restyClient: restyClient,
	}
}

func (w wordpressClient) uri(path string) string {
	return w.Endpoint + "/wp-json/wp/v2/" + path
}

type mediaResponse struct {
	Id        int    `json:"id"`
	SourceUrl string `json:"source_url"`
}

func (w wordpressClient) uploadMedia(path string) (int, string, error) {
	mediaResp := mediaResponse{}
	resp, err := w.restyClient.
		R().
		SetFile("file", path).
		SetResult(&mediaResp).
		SetHeader("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(path))).
		Post(w.uri("media"))

	if resp.IsError() {
		return 0, "", errors.New(string(resp.Body()))
	}
	return mediaResp.Id, mediaResp.SourceUrl, err
}

type postResponse struct {
	Id   int    `json:"id"`
	Link string `json:"link"`
}

func (w wordpressClient) CreatePost(title, content string, media []string) (string, error) {
	errgrp := errgroup.Group{}
	uploadedMedia := sync.Map{}
	for _, m := range media {
		func(localPath string) {
			errgrp.Go(func() error {
				mediaId, remoteUrl, err := w.uploadMedia(localPath)
				if err != nil {
					return err
				}
				uploadedMedia.Store(mediaId, remoteUrl)
				return nil
			})
		}(m)
	}

	if err := errgrp.Wait(); err != nil {
		return "", err
	}

	var featuredMediaId int
	uploadedMedia.Range(func(key, value any) bool {
		content += w.GenerateHtml(value.(string))

		if strings.HasSuffix(value.(string), "jpg") && featuredMediaId == 0 {
			featuredMediaId = key.(int)
		}
		return true
	})

	payload := map[string]interface{}{
		"status":         "publish",
		"title":          title,
		"content":        content,
		"featured_media": featuredMediaId,
	}

	postResp := postResponse{}
	resp, err := w.restyClient.R().SetBody(payload).SetResult(&postResp).Post(w.uri("posts"))
	if resp.IsError() {
		return "", errors.New(string(resp.Body()))
	}
	return postResp.Link, err
}

func (w wordpressClient) GenerateHtml(value string) string {
	if strings.HasSuffix(value, "jpg") {
		return fmt.Sprintf("<img width='100%%' src='%s' /> <br />", value)
	}
	return fmt.Sprintf("<video width='100%%' controls><source src='%s' type='video/mp4'>Your browser does not support the video tag.</video> <br />", value)
}
