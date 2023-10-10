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

// InputGameID represents TL type `inputGameID#32c3e77`.
// Indicates an already sent game
//
// See https://core.telegram.org/constructor/inputGameID for reference.
type InputGameID struct {
	// game ID from Game¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/type/Game
	ID int64
	// access hash from Game¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/type/Game
	AccessHash int64
}

// InputGameIDTypeID is TL type id of InputGameID.
const InputGameIDTypeID = 0x32c3e77

// construct implements constructor of InputGameClass.
func (i InputGameID) construct() InputGameClass { return &i }

// Ensuring interfaces in compile-time for InputGameID.
var (
	_ bin.Encoder     = &InputGameID{}
	_ bin.Decoder     = &InputGameID{}
	_ bin.BareEncoder = &InputGameID{}
	_ bin.BareDecoder = &InputGameID{}

	_ InputGameClass = &InputGameID{}
)

func (i *InputGameID) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputGameID) String() string {
	if i == nil {
		return "InputGameID(nil)"
	}
	type Alias InputGameID
	return fmt.Sprintf("InputGameID%+v", Alias(*i))
}

// FillFrom fills InputGameID from given interface.
func (i *InputGameID) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputGameID) TypeID() uint32 {
	return InputGameIDTypeID
}

// TypeName returns name of type in TL schema.
func (*InputGameID) TypeName() string {
	return "inputGameID"
}

// TypeInfo returns info about TL type.
func (i *InputGameID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputGameID",
		ID:   InputGameIDTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputGameID) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGameID#32c3e77 as nil")
	}
	b.PutID(InputGameIDTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputGameID) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGameID#32c3e77 as nil")
	}
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputGameID) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGameID#32c3e77 to nil")
	}
	if err := b.ConsumeID(InputGameIDTypeID); err != nil {
		return fmt.Errorf("unable to decode inputGameID#32c3e77: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputGameID) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGameID#32c3e77 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputGameID#32c3e77: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputGameID#32c3e77: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputGameID) GetID() (value int64) {
	if i == nil {
		return
	}
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputGameID) GetAccessHash() (value int64) {
	if i == nil {
		return
	}
	return i.AccessHash
}

// InputGameShortName represents TL type `inputGameShortName#c331e80a`.
// Game by short name
//
// See https://core.telegram.org/constructor/inputGameShortName for reference.
type InputGameShortName struct {
	// The bot that provides the game
	BotID InputUserClass
	// The game's short name, usually obtained from a game link »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/links#game-links
	ShortName string
}

// InputGameShortNameTypeID is TL type id of InputGameShortName.
const InputGameShortNameTypeID = 0xc331e80a

// construct implements constructor of InputGameClass.
func (i InputGameShortName) construct() InputGameClass { return &i }

// Ensuring interfaces in compile-time for InputGameShortName.
var (
	_ bin.Encoder     = &InputGameShortName{}
	_ bin.Decoder     = &InputGameShortName{}
	_ bin.BareEncoder = &InputGameShortName{}
	_ bin.BareDecoder = &InputGameShortName{}

	_ InputGameClass = &InputGameShortName{}
)

func (i *InputGameShortName) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.BotID == nil) {
		return false
	}
	if !(i.ShortName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputGameShortName) String() string {
	if i == nil {
		return "InputGameShortName(nil)"
	}
	type Alias InputGameShortName
	return fmt.Sprintf("InputGameShortName%+v", Alias(*i))
}

// FillFrom fills InputGameShortName from given interface.
func (i *InputGameShortName) FillFrom(from interface {
	GetBotID() (value InputUserClass)
	GetShortName() (value string)
}) {
	i.BotID = from.GetBotID()
	i.ShortName = from.GetShortName()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputGameShortName) TypeID() uint32 {
	return InputGameShortNameTypeID
}

// TypeName returns name of type in TL schema.
func (*InputGameShortName) TypeName() string {
	return "inputGameShortName"
}

// TypeInfo returns info about TL type.
func (i *InputGameShortName) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputGameShortName",
		ID:   InputGameShortNameTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotID",
			SchemaName: "bot_id",
		},
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputGameShortName) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGameShortName#c331e80a as nil")
	}
	b.PutID(InputGameShortNameTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputGameShortName) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGameShortName#c331e80a as nil")
	}
	if i.BotID == nil {
		return fmt.Errorf("unable to encode inputGameShortName#c331e80a: field bot_id is nil")
	}
	if err := i.BotID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputGameShortName#c331e80a: field bot_id: %w", err)
	}
	b.PutString(i.ShortName)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputGameShortName) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGameShortName#c331e80a to nil")
	}
	if err := b.ConsumeID(InputGameShortNameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputGameShortName#c331e80a: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputGameShortName) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGameShortName#c331e80a to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputGameShortName#c331e80a: field bot_id: %w", err)
		}
		i.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputGameShortName#c331e80a: field short_name: %w", err)
		}
		i.ShortName = value
	}
	return nil
}

// GetBotID returns value of BotID field.
func (i *InputGameShortName) GetBotID() (value InputUserClass) {
	if i == nil {
		return
	}
	return i.BotID
}

// GetShortName returns value of ShortName field.
func (i *InputGameShortName) GetShortName() (value string) {
	if i == nil {
		return
	}
	return i.ShortName
}

// InputGameClassName is schema name of InputGameClass.
const InputGameClassName = "InputGame"

// InputGameClass represents InputGame generic type.
//
// See https://core.telegram.org/type/InputGame for reference.
//
// Example:
//
//	g, err := tg.DecodeInputGame(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputGameID: // inputGameID#32c3e77
//	case *tg.InputGameShortName: // inputGameShortName#c331e80a
//	default: panic(v)
//	}
type InputGameClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputGameClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputGame implements binary de-serialization for InputGameClass.
func DecodeInputGame(buf *bin.Buffer) (InputGameClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputGameIDTypeID:
		// Decoding inputGameID#32c3e77.
		v := InputGameID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputGameClass: %w", err)
		}
		return &v, nil
	case InputGameShortNameTypeID:
		// Decoding inputGameShortName#c331e80a.
		v := InputGameShortName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputGameClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputGameClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputGame boxes the InputGameClass providing a helper.
type InputGameBox struct {
	InputGame InputGameClass
}

// Decode implements bin.Decoder for InputGameBox.
func (b *InputGameBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputGameBox to nil")
	}
	v, err := DecodeInputGame(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputGame = v
	return nil
}

// Encode implements bin.Encode for InputGameBox.
func (b *InputGameBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputGame == nil {
		return fmt.Errorf("unable to encode InputGameClass as nil")
	}
	return b.InputGame.Encode(buf)
}