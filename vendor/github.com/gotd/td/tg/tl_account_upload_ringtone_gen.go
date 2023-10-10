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

// AccountUploadRingtoneRequest represents TL type `account.uploadRingtone#831a83a2`.
// Upload notification sound, use account.saveRingtone¹ to convert it and add it to the
// list of saved notification sounds.
//
// Links:
//  1. https://core.telegram.org/method/account.saveRingtone
//
// See https://core.telegram.org/method/account.uploadRingtone for reference.
type AccountUploadRingtoneRequest struct {
	// Notification sound
	File InputFileClass
	// File name
	FileName string
	// MIME type of file
	MimeType string
}

// AccountUploadRingtoneRequestTypeID is TL type id of AccountUploadRingtoneRequest.
const AccountUploadRingtoneRequestTypeID = 0x831a83a2

// Ensuring interfaces in compile-time for AccountUploadRingtoneRequest.
var (
	_ bin.Encoder     = &AccountUploadRingtoneRequest{}
	_ bin.Decoder     = &AccountUploadRingtoneRequest{}
	_ bin.BareEncoder = &AccountUploadRingtoneRequest{}
	_ bin.BareDecoder = &AccountUploadRingtoneRequest{}
)

func (u *AccountUploadRingtoneRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.File == nil) {
		return false
	}
	if !(u.FileName == "") {
		return false
	}
	if !(u.MimeType == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUploadRingtoneRequest) String() string {
	if u == nil {
		return "AccountUploadRingtoneRequest(nil)"
	}
	type Alias AccountUploadRingtoneRequest
	return fmt.Sprintf("AccountUploadRingtoneRequest%+v", Alias(*u))
}

// FillFrom fills AccountUploadRingtoneRequest from given interface.
func (u *AccountUploadRingtoneRequest) FillFrom(from interface {
	GetFile() (value InputFileClass)
	GetFileName() (value string)
	GetMimeType() (value string)
}) {
	u.File = from.GetFile()
	u.FileName = from.GetFileName()
	u.MimeType = from.GetMimeType()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUploadRingtoneRequest) TypeID() uint32 {
	return AccountUploadRingtoneRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUploadRingtoneRequest) TypeName() string {
	return "account.uploadRingtone"
}

// TypeInfo returns info about TL type.
func (u *AccountUploadRingtoneRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.uploadRingtone",
		ID:   AccountUploadRingtoneRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "File",
			SchemaName: "file",
		},
		{
			Name:       "FileName",
			SchemaName: "file_name",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUploadRingtoneRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.uploadRingtone#831a83a2 as nil")
	}
	b.PutID(AccountUploadRingtoneRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUploadRingtoneRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.uploadRingtone#831a83a2 as nil")
	}
	if u.File == nil {
		return fmt.Errorf("unable to encode account.uploadRingtone#831a83a2: field file is nil")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadRingtone#831a83a2: field file: %w", err)
	}
	b.PutString(u.FileName)
	b.PutString(u.MimeType)
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUploadRingtoneRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.uploadRingtone#831a83a2 to nil")
	}
	if err := b.ConsumeID(AccountUploadRingtoneRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.uploadRingtone#831a83a2: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUploadRingtoneRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.uploadRingtone#831a83a2 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadRingtone#831a83a2: field file: %w", err)
		}
		u.File = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadRingtone#831a83a2: field file_name: %w", err)
		}
		u.FileName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadRingtone#831a83a2: field mime_type: %w", err)
		}
		u.MimeType = value
	}
	return nil
}

// GetFile returns value of File field.
func (u *AccountUploadRingtoneRequest) GetFile() (value InputFileClass) {
	if u == nil {
		return
	}
	return u.File
}

// GetFileName returns value of FileName field.
func (u *AccountUploadRingtoneRequest) GetFileName() (value string) {
	if u == nil {
		return
	}
	return u.FileName
}

// GetMimeType returns value of MimeType field.
func (u *AccountUploadRingtoneRequest) GetMimeType() (value string) {
	if u == nil {
		return
	}
	return u.MimeType
}

// AccountUploadRingtone invokes method account.uploadRingtone#831a83a2 returning error if any.
// Upload notification sound, use account.saveRingtone¹ to convert it and add it to the
// list of saved notification sounds.
//
// Links:
//  1. https://core.telegram.org/method/account.saveRingtone
//
// See https://core.telegram.org/method/account.uploadRingtone for reference.
func (c *Client) AccountUploadRingtone(ctx context.Context, request *AccountUploadRingtoneRequest) (DocumentClass, error) {
	var result DocumentBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Document, nil
}