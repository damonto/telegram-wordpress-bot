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

// ReceivedNotifyMessage represents TL type `receivedNotifyMessage#a384b779`.
// Message ID, for which PUSH-notifications were cancelled.
//
// See https://core.telegram.org/constructor/receivedNotifyMessage for reference.
type ReceivedNotifyMessage struct {
	// Message ID, for which PUSH-notifications were canceled
	ID int
	// Reserved for future use
	Flags int
}

// ReceivedNotifyMessageTypeID is TL type id of ReceivedNotifyMessage.
const ReceivedNotifyMessageTypeID = 0xa384b779

// Ensuring interfaces in compile-time for ReceivedNotifyMessage.
var (
	_ bin.Encoder     = &ReceivedNotifyMessage{}
	_ bin.Decoder     = &ReceivedNotifyMessage{}
	_ bin.BareEncoder = &ReceivedNotifyMessage{}
	_ bin.BareDecoder = &ReceivedNotifyMessage{}
)

func (r *ReceivedNotifyMessage) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ID == 0) {
		return false
	}
	if !(r.Flags == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReceivedNotifyMessage) String() string {
	if r == nil {
		return "ReceivedNotifyMessage(nil)"
	}
	type Alias ReceivedNotifyMessage
	return fmt.Sprintf("ReceivedNotifyMessage%+v", Alias(*r))
}

// FillFrom fills ReceivedNotifyMessage from given interface.
func (r *ReceivedNotifyMessage) FillFrom(from interface {
	GetID() (value int)
	GetFlags() (value int)
}) {
	r.ID = from.GetID()
	r.Flags = from.GetFlags()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReceivedNotifyMessage) TypeID() uint32 {
	return ReceivedNotifyMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*ReceivedNotifyMessage) TypeName() string {
	return "receivedNotifyMessage"
}

// TypeInfo returns info about TL type.
func (r *ReceivedNotifyMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "receivedNotifyMessage",
		ID:   ReceivedNotifyMessageTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReceivedNotifyMessage) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode receivedNotifyMessage#a384b779 as nil")
	}
	b.PutID(ReceivedNotifyMessageTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReceivedNotifyMessage) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode receivedNotifyMessage#a384b779 as nil")
	}
	b.PutInt(r.ID)
	b.PutInt(r.Flags)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReceivedNotifyMessage) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode receivedNotifyMessage#a384b779 to nil")
	}
	if err := b.ConsumeID(ReceivedNotifyMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode receivedNotifyMessage#a384b779: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReceivedNotifyMessage) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode receivedNotifyMessage#a384b779 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode receivedNotifyMessage#a384b779: field id: %w", err)
		}
		r.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode receivedNotifyMessage#a384b779: field flags: %w", err)
		}
		r.Flags = value
	}
	return nil
}

// GetID returns value of ID field.
func (r *ReceivedNotifyMessage) GetID() (value int) {
	if r == nil {
		return
	}
	return r.ID
}

// GetFlags returns value of Flags field.
func (r *ReceivedNotifyMessage) GetFlags() (value int) {
	if r == nil {
		return
	}
	return r.Flags
}