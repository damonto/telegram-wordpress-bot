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

// AccountCreateThemeRequest represents TL type `account.createTheme#652e4400`.
// Create a theme
//
// See https://core.telegram.org/method/account.createTheme for reference.
type AccountCreateThemeRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Unique theme ID used to generate theme deep links¹, can be empty to autogenerate a
	// random ID.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#theme-links
	Slug string
	// Theme name
	Title string
	// Theme file
	//
	// Use SetDocument and GetDocument helpers.
	Document InputDocumentClass
	// Theme settings, multiple values can be provided for the different base themes
	// (day/night mode, etc).
	//
	// Use SetSettings and GetSettings helpers.
	Settings []InputThemeSettings
}

// AccountCreateThemeRequestTypeID is TL type id of AccountCreateThemeRequest.
const AccountCreateThemeRequestTypeID = 0x652e4400

// Ensuring interfaces in compile-time for AccountCreateThemeRequest.
var (
	_ bin.Encoder     = &AccountCreateThemeRequest{}
	_ bin.Decoder     = &AccountCreateThemeRequest{}
	_ bin.BareEncoder = &AccountCreateThemeRequest{}
	_ bin.BareDecoder = &AccountCreateThemeRequest{}
)

func (c *AccountCreateThemeRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Slug == "") {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.Document == nil) {
		return false
	}
	if !(c.Settings == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountCreateThemeRequest) String() string {
	if c == nil {
		return "AccountCreateThemeRequest(nil)"
	}
	type Alias AccountCreateThemeRequest
	return fmt.Sprintf("AccountCreateThemeRequest%+v", Alias(*c))
}

// FillFrom fills AccountCreateThemeRequest from given interface.
func (c *AccountCreateThemeRequest) FillFrom(from interface {
	GetSlug() (value string)
	GetTitle() (value string)
	GetDocument() (value InputDocumentClass, ok bool)
	GetSettings() (value []InputThemeSettings, ok bool)
}) {
	c.Slug = from.GetSlug()
	c.Title = from.GetTitle()
	if val, ok := from.GetDocument(); ok {
		c.Document = val
	}

	if val, ok := from.GetSettings(); ok {
		c.Settings = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountCreateThemeRequest) TypeID() uint32 {
	return AccountCreateThemeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountCreateThemeRequest) TypeName() string {
	return "account.createTheme"
}

// TypeInfo returns info about TL type.
func (c *AccountCreateThemeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.createTheme",
		ID:   AccountCreateThemeRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Document",
			SchemaName: "document",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
			Null:       !c.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *AccountCreateThemeRequest) SetFlags() {
	if !(c.Document == nil) {
		c.Flags.Set(2)
	}
	if !(c.Settings == nil) {
		c.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (c *AccountCreateThemeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.createTheme#652e4400 as nil")
	}
	b.PutID(AccountCreateThemeRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AccountCreateThemeRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.createTheme#652e4400 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.createTheme#652e4400: field flags: %w", err)
	}
	b.PutString(c.Slug)
	b.PutString(c.Title)
	if c.Flags.Has(2) {
		if c.Document == nil {
			return fmt.Errorf("unable to encode account.createTheme#652e4400: field document is nil")
		}
		if err := c.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.createTheme#652e4400: field document: %w", err)
		}
	}
	if c.Flags.Has(3) {
		b.PutVectorHeader(len(c.Settings))
		for idx, v := range c.Settings {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode account.createTheme#652e4400: field settings element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountCreateThemeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.createTheme#652e4400 to nil")
	}
	if err := b.ConsumeID(AccountCreateThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.createTheme#652e4400: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AccountCreateThemeRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.createTheme#652e4400 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.createTheme#652e4400: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.createTheme#652e4400: field slug: %w", err)
		}
		c.Slug = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.createTheme#652e4400: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(2) {
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.createTheme#652e4400: field document: %w", err)
		}
		c.Document = value
	}
	if c.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.createTheme#652e4400: field settings: %w", err)
		}

		if headerLen > 0 {
			c.Settings = make([]InputThemeSettings, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputThemeSettings
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.createTheme#652e4400: field settings: %w", err)
			}
			c.Settings = append(c.Settings, value)
		}
	}
	return nil
}

// GetSlug returns value of Slug field.
func (c *AccountCreateThemeRequest) GetSlug() (value string) {
	if c == nil {
		return
	}
	return c.Slug
}

// GetTitle returns value of Title field.
func (c *AccountCreateThemeRequest) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// SetDocument sets value of Document conditional field.
func (c *AccountCreateThemeRequest) SetDocument(value InputDocumentClass) {
	c.Flags.Set(2)
	c.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (c *AccountCreateThemeRequest) GetDocument() (value InputDocumentClass, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.Document, true
}

// SetSettings sets value of Settings conditional field.
func (c *AccountCreateThemeRequest) SetSettings(value []InputThemeSettings) {
	c.Flags.Set(3)
	c.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (c *AccountCreateThemeRequest) GetSettings() (value []InputThemeSettings, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(3) {
		return value, false
	}
	return c.Settings, true
}

// GetDocumentAsNotEmpty returns mapped value of Document conditional field and
// boolean which is true if field was set.
func (c *AccountCreateThemeRequest) GetDocumentAsNotEmpty() (*InputDocument, bool) {
	if value, ok := c.GetDocument(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// AccountCreateTheme invokes method account.createTheme#652e4400 returning error if any.
// Create a theme
//
// Possible errors:
//
//	400 THEME_MIME_INVALID: The theme's MIME type is invalid.
//	400 THEME_TITLE_INVALID: The specified theme title is invalid.
//
// See https://core.telegram.org/method/account.createTheme for reference.
func (c *Client) AccountCreateTheme(ctx context.Context, request *AccountCreateThemeRequest) (*Theme, error) {
	var result Theme

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}