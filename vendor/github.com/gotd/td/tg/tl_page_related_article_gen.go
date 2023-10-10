// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// PageRelatedArticle represents TL type `pageRelatedArticle#b390dc08`.
// Related article
//
// See https://core.telegram.org/constructor/pageRelatedArticle for reference.
type PageRelatedArticle struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// URL of article
	URL string
	// Webpage ID of generated IV preview
	WebpageID int64
	// Title
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Description
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// ID of preview photo
	//
	// Use SetPhotoID and GetPhotoID helpers.
	PhotoID int64
	// Author name
	//
	// Use SetAuthor and GetAuthor helpers.
	Author string
	// Date of publication
	//
	// Use SetPublishedDate and GetPublishedDate helpers.
	PublishedDate int
}

// PageRelatedArticleTypeID is TL type id of PageRelatedArticle.
const PageRelatedArticleTypeID = 0xb390dc08

// Ensuring interfaces in compile-time for PageRelatedArticle.
var (
	_ bin.Encoder     = &PageRelatedArticle{}
	_ bin.Decoder     = &PageRelatedArticle{}
	_ bin.BareEncoder = &PageRelatedArticle{}
	_ bin.BareDecoder = &PageRelatedArticle{}
)

func (p *PageRelatedArticle) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.URL == "") {
		return false
	}
	if !(p.WebpageID == 0) {
		return false
	}
	if !(p.Title == "") {
		return false
	}
	if !(p.Description == "") {
		return false
	}
	if !(p.PhotoID == 0) {
		return false
	}
	if !(p.Author == "") {
		return false
	}
	if !(p.PublishedDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageRelatedArticle) String() string {
	if p == nil {
		return "PageRelatedArticle(nil)"
	}
	type Alias PageRelatedArticle
	return fmt.Sprintf("PageRelatedArticle%+v", Alias(*p))
}

// FillFrom fills PageRelatedArticle from given interface.
func (p *PageRelatedArticle) FillFrom(from interface {
	GetURL() (value string)
	GetWebpageID() (value int64)
	GetTitle() (value string, ok bool)
	GetDescription() (value string, ok bool)
	GetPhotoID() (value int64, ok bool)
	GetAuthor() (value string, ok bool)
	GetPublishedDate() (value int, ok bool)
}) {
	p.URL = from.GetURL()
	p.WebpageID = from.GetWebpageID()
	if val, ok := from.GetTitle(); ok {
		p.Title = val
	}

	if val, ok := from.GetDescription(); ok {
		p.Description = val
	}

	if val, ok := from.GetPhotoID(); ok {
		p.PhotoID = val
	}

	if val, ok := from.GetAuthor(); ok {
		p.Author = val
	}

	if val, ok := from.GetPublishedDate(); ok {
		p.PublishedDate = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageRelatedArticle) TypeID() uint32 {
	return PageRelatedArticleTypeID
}

// TypeName returns name of type in TL schema.
func (*PageRelatedArticle) TypeName() string {
	return "pageRelatedArticle"
}

// TypeInfo returns info about TL type.
func (p *PageRelatedArticle) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageRelatedArticle",
		ID:   PageRelatedArticleTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "WebpageID",
			SchemaName: "webpage_id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "Description",
			SchemaName: "description",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "PhotoID",
			SchemaName: "photo_id",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "Author",
			SchemaName: "author",
			Null:       !p.Flags.Has(3),
		},
		{
			Name:       "PublishedDate",
			SchemaName: "published_date",
			Null:       !p.Flags.Has(4),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *PageRelatedArticle) SetFlags() {
	if !(p.Title == "") {
		p.Flags.Set(0)
	}
	if !(p.Description == "") {
		p.Flags.Set(1)
	}
	if !(p.PhotoID == 0) {
		p.Flags.Set(2)
	}
	if !(p.Author == "") {
		p.Flags.Set(3)
	}
	if !(p.PublishedDate == 0) {
		p.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (p *PageRelatedArticle) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageRelatedArticle#b390dc08 as nil")
	}
	b.PutID(PageRelatedArticleTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageRelatedArticle) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageRelatedArticle#b390dc08 as nil")
	}
	p.SetFlags()
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageRelatedArticle#b390dc08: field flags: %w", err)
	}
	b.PutString(p.URL)
	b.PutLong(p.WebpageID)
	if p.Flags.Has(0) {
		b.PutString(p.Title)
	}
	if p.Flags.Has(1) {
		b.PutString(p.Description)
	}
	if p.Flags.Has(2) {
		b.PutLong(p.PhotoID)
	}
	if p.Flags.Has(3) {
		b.PutString(p.Author)
	}
	if p.Flags.Has(4) {
		b.PutInt(p.PublishedDate)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageRelatedArticle) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageRelatedArticle#b390dc08 to nil")
	}
	if err := b.ConsumeID(PageRelatedArticleTypeID); err != nil {
		return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageRelatedArticle) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageRelatedArticle#b390dc08 to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field url: %w", err)
		}
		p.URL = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field webpage_id: %w", err)
		}
		p.WebpageID = value
	}
	if p.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field title: %w", err)
		}
		p.Title = value
	}
	if p.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field description: %w", err)
		}
		p.Description = value
	}
	if p.Flags.Has(2) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field photo_id: %w", err)
		}
		p.PhotoID = value
	}
	if p.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field author: %w", err)
		}
		p.Author = value
	}
	if p.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field published_date: %w", err)
		}
		p.PublishedDate = value
	}
	return nil
}

// GetURL returns value of URL field.
func (p *PageRelatedArticle) GetURL() (value string) {
	if p == nil {
		return
	}
	return p.URL
}

// GetWebpageID returns value of WebpageID field.
func (p *PageRelatedArticle) GetWebpageID() (value int64) {
	if p == nil {
		return
	}
	return p.WebpageID
}

// SetTitle sets value of Title conditional field.
func (p *PageRelatedArticle) SetTitle(value string) {
	p.Flags.Set(0)
	p.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetTitle() (value string, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.Title, true
}

// SetDescription sets value of Description conditional field.
func (p *PageRelatedArticle) SetDescription(value string) {
	p.Flags.Set(1)
	p.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetDescription() (value string, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Description, true
}

// SetPhotoID sets value of PhotoID conditional field.
func (p *PageRelatedArticle) SetPhotoID(value int64) {
	p.Flags.Set(2)
	p.PhotoID = value
}

// GetPhotoID returns value of PhotoID conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetPhotoID() (value int64, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.PhotoID, true
}

// SetAuthor sets value of Author conditional field.
func (p *PageRelatedArticle) SetAuthor(value string) {
	p.Flags.Set(3)
	p.Author = value
}

// GetAuthor returns value of Author conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetAuthor() (value string, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(3) {
		return value, false
	}
	return p.Author, true
}

// SetPublishedDate sets value of PublishedDate conditional field.
func (p *PageRelatedArticle) SetPublishedDate(value int) {
	p.Flags.Set(4)
	p.PublishedDate = value
}

// GetPublishedDate returns value of PublishedDate conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetPublishedDate() (value int, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.PublishedDate, true
}