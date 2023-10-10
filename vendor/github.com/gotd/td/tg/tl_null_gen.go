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

// Null represents TL type `null#56730bcc`.
// Corresponds to an arbitrary empty object.
//
// See https://core.telegram.org/constructor/null for reference.
type Null struct {
}

// NullTypeID is TL type id of Null.
const NullTypeID = 0x56730bcc

// Ensuring interfaces in compile-time for Null.
var (
	_ bin.Encoder     = &Null{}
	_ bin.Decoder     = &Null{}
	_ bin.BareEncoder = &Null{}
	_ bin.BareDecoder = &Null{}
)

func (n *Null) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *Null) String() string {
	if n == nil {
		return "Null(nil)"
	}
	type Alias Null
	return fmt.Sprintf("Null%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Null) TypeID() uint32 {
	return NullTypeID
}

// TypeName returns name of type in TL schema.
func (*Null) TypeName() string {
	return "null"
}

// TypeInfo returns info about TL type.
func (n *Null) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "null",
		ID:   NullTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *Null) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode null#56730bcc as nil")
	}
	b.PutID(NullTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *Null) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode null#56730bcc as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *Null) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode null#56730bcc to nil")
	}
	if err := b.ConsumeID(NullTypeID); err != nil {
		return fmt.Errorf("unable to decode null#56730bcc: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *Null) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode null#56730bcc to nil")
	}
	return nil
}