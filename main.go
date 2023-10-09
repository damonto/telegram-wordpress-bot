package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/celestix/gotgproto"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/dispatcher/handlers/filters"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/sessionMaker"
	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/tg"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"
)

var (
	telegramApiId        int
	telegramApiHash      string
	telegramPhoneNumber  string
	wordpressEndpoint    string
	wordpressAppUsername string
	wordpressAppPassword string
)

func init() {
	flag.IntVar(&telegramApiId, "tg-api-id", 0, "Your telegram api id")
	flag.StringVar(&telegramApiHash, "tg-api-hash", "", "Your telegram api hash")
	flag.StringVar(&telegramPhoneNumber, "tg-phone-number", "", "Your telegram phone number in international format. (+18883334444)")
	flag.StringVar(&wordpressEndpoint, "wp-endpoint", "", "Your wordpress url")
	flag.StringVar(&wordpressAppUsername, "wp-app-username", "", "Your wordpress application username")
	flag.StringVar(&wordpressAppPassword, "wp-app-password", "", "Your wordpress application password")
	flag.Parse()
}

func main() {
	client, err := gotgproto.NewClient(
		telegramApiId,
		telegramApiHash,
		gotgproto.ClientType{Phone: telegramPhoneNumber},
		&gotgproto.ClientOpts{
			Session: sessionMaker.NewSession("tg-wp-bot", sessionMaker.Session),
		},
	)
	if err != nil {
		panic(err)
	}

	wpClient := NewWordpressClient(wordpressEndpoint, wordpressAppUsername, wordpressAppPassword)
	slog.Info("client has been started", "username", client.Self.Username)

	client.Dispatcher.AddHandler(handlers.NewMessage(filters.Message.All, func(ctx *ext.Context, u *ext.Update) error {
		userClass := u.EffectiveChat().GetInputUser()
		if userClass == nil {
			return nil
		}

		if userClass.(*tg.InputUser).GetUserID() != client.Self.ID || !u.EffectiveUser().GetSelf() {
			return nil
		}

		messages, err := getMessages(client.API(), ctx, u.EffectiveMessage.Text)
		if err != nil {
			slog.Error("failed to get messages", "error", err)
			ctx.Reply(u, err.Error(), nil)
			return nil
		}

		downloaded, err := download(client.API(), ctx, messages)
		if err != nil {
			slog.Error("failed to download media", "error", err)
			ctx.Reply(u, err.Error(), nil)
			return nil
		}
		message := messages[0]
		title := message.Message
		if message.Message == "" {
			title = time.Now().Format(time.DateTime)
		}
		postUrl, err := wpClient.CreatePost(title, message.Message, downloaded)
		if err != nil {
			ctx.Reply(u, err.Error(), nil)
			return nil
		}

		ctx.Reply(u, "Post created: "+postUrl, nil)
		return nil
	}))

	client.Idle()
}

func download(client *tg.Client, ctx context.Context, messages []*tg.Message) ([]string, error) {
	if err := os.MkdirAll("./downloads", os.ModePerm); err != nil {
		return []string{}, err
	}

	mediaDocuments := sync.Map{}
	errgrp := errgroup.Group{}
	for _, message := range messages {
		media, ok := message.GetMedia()
		if !ok {
			continue
		}
		switch media.TypeName() {
		case "messageMediaPhoto":
			func(message *tg.Message) {
				errgrp.Go(func() error {
					mediaPhoto := media.(*tg.MessageMediaPhoto)
					path, err := downloadPhoto(ctx, client, mediaPhoto.Photo)
					if err != nil {
						return err
					}
					mediaDocuments.Store(mediaPhoto.Photo.GetID(), path)
					return nil
				})
			}(message)
		case "messageMediaDocument":
			func(message *tg.Message) {
				errgrp.Go(func() error {
					mediaDocument := media.(*tg.MessageMediaDocument)
					path, err := downloadMediaDocument(ctx, client, mediaDocument.Document)
					if err != nil {
						return err
					}
					mediaDocuments.Store(mediaDocument.Document.GetID(), path)
					return nil
				})
			}(message)
		default:
			return []string{}, fmt.Errorf("unsupported media type %s", media.TypeName())
		}
	}

	if err := errgrp.Wait(); err != nil {
		return []string{}, err
	}

	downloadedMedia := []string{}
	mediaDocuments.Range(func(key, value any) bool {
		downloadedMedia = append(downloadedMedia, value.(string))
		return true
	})

	return downloadedMedia, nil
}

func downloadMediaDocument(ctx context.Context, client *tg.Client, mediaDocument tg.DocumentClass) (string, error) {
	media, ok := mediaDocument.AsNotEmpty()
	if !ok {
		return "", errors.New("media is empty")
	}

	mimeType := strings.Split(media.GetMimeType(), "/")
	filePath := fmt.Sprintf("./downloads/%d.%s", media.GetID(), mimeType[len(mimeType)-1])
	slog.Info("downloading photo", "id", media.GetID(), "path", filePath)
	downloader := downloader.NewDownloader()
	_, err := downloader.Download(client, media.AsInputDocumentFileLocation()).ToPath(ctx, filePath)
	return filePath, err
}

func downloadPhoto(ctx context.Context, client *tg.Client, photoClass tg.PhotoClass) (string, error) {
	photo, ok := photoClass.AsNotEmpty()
	if !ok {
		return "", errors.New("photo is empty")
	}

	downloader := downloader.NewDownloader()
	filePath := fmt.Sprintf("./downloads/%d.png", photo.GetID())
	slog.Info("downloading photo", "id", photo.GetID(), "path", filePath)
	_, err := downloader.Download(client, &tg.InputPhotoFileLocation{
		ID:            photo.GetID(),
		AccessHash:    photo.GetAccessHash(),
		FileReference: photo.GetFileReference(),
		ThumbSize:     "y",
	}).ToPath(ctx, filePath)
	return filePath, err
}

func getMessages(client *tg.Client, ctx context.Context, messageLink string) ([]*tg.Message, error) {
	var err error
	inputChannel, messageId, err := parseMessageLink(client, messageLink)
	if err != nil {
		return nil, err
	}

	var messageClass tg.MessagesMessagesClass
	messageClass, err = client.ChannelsGetMessages(ctx, &tg.ChannelsGetMessagesRequest{
		Channel: inputChannel,
		ID:      []tg.InputMessageClass{&tg.InputMessageID{ID: int(messageId)}},
	})
	if err != nil {
		return nil, err
	}

	message := messageClass.(*tg.MessagesChannelMessages).GetMessages()[0].(*tg.Message)
	// media group.
	if message.GroupedID != 0 {
		messageClass, err = client.ChannelsGetMessages(ctx, &tg.ChannelsGetMessagesRequest{
			Channel: inputChannel,
			ID:      messageIdRange(messageId-9, messageId+10),
		})

		if err != nil {
			slog.Error("failed to get messages", "messageId", messageId, "groupedId", message.GroupedID, "error", err)
			return nil, err
		}
	}

	messages := make([]*tg.Message, 0)
	for _, channelMessageClass := range messageClass.(*tg.MessagesChannelMessages).GetMessages() {
		if channelMessageClass.TypeName() == "message" {
			channelMessage := channelMessageClass.(*tg.Message)
			if channelMessage.GroupedID == message.GroupedID {
				messages = append(messages, channelMessage)
			}
		}
	}
	return messages, nil
}

func messageIdRange(start int, end int) []tg.InputMessageClass {
	numbers := []tg.InputMessageClass{}
	for i := start; i <= end; i++ {
		numbers = append(numbers, &tg.InputMessageID{ID: i})
	}
	return numbers
}

func parseMessageLink(client *tg.Client, messageLink string) (*tg.InputChannel, int, error) {
	if !strings.Contains(messageLink, "https://t.me") {
		return nil, 0, errors.New("incorrect message link")
	}

	factor := strings.Split(strings.Replace(messageLink, "https://t.me/", "", 1), "/")
	messageId, _ := strconv.ParseInt(factor[len(factor)-1], 10, 0)
	// 群组都是 https://t.me/c/channelId/messageId
	if len(factor) == 3 && factor[0] == "c" {
		channelId, _ := strconv.ParseInt(factor[1], 10, 0)
		inputChannel, err := resolveChannelByChannelId(client, channelId)
		return inputChannel, int(messageId), err
	}

	inputChannel, err := resolveChannelByUsername(client, factor[0])
	return inputChannel, int(messageId), err
}

func resolveChannelByChannelId(client *tg.Client, channelId int64) (*tg.InputChannel, error) {
	messagesChatsClass, err := client.ChannelsGetChannels(context.Background(), []tg.InputChannelClass{&tg.InputChannel{ChannelID: channelId}})
	if err != nil {
		return nil, err
	}

	messagesChats := messagesChatsClass.(*tg.MessagesChats)
	channel := messagesChats.GetChats()[0].(*tg.Channel)
	accessHash, ok := channel.GetAccessHash()
	if !ok {
		return nil, errors.New("failed to get channel access hash")
	}
	return &tg.InputChannel{
		ChannelID:  channel.GetID(),
		AccessHash: accessHash,
	}, err
}

func resolveChannelByUsername(client *tg.Client, username string) (*tg.InputChannel, error) {
	resovler := peer.DefaultResolver(client)
	resolvedClass, err := resovler.ResolveDomain(context.Background(), username)
	if err != nil {
		return nil, err
	}

	if resolvedClass.TypeName() != "inputPeerChannel" {
		return nil, fmt.Errorf("incorrect username %s", username)
	}

	inputPeerChannel := resolvedClass.(*tg.InputPeerChannel)
	return &tg.InputChannel{
		ChannelID:  inputPeerChannel.GetChannelID(),
		AccessHash: inputPeerChannel.GetAccessHash(),
	}, nil
}
