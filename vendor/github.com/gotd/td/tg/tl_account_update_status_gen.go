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

// AccountUpdateStatusRequest represents TL type `account.updateStatus#6628562c`.
// Updates online user status.
//
// See https://core.telegram.org/method/account.updateStatus for reference.
type AccountUpdateStatusRequest struct {
	// If (boolTrue)¹ is transmitted, user status will change to (userStatusOffline)².
	//
	// Links:
	//  1) https://core.telegram.org/constructor/boolTrue
	//  2) https://core.telegram.org/constructor/userStatusOffline
	Offline bool
}

// AccountUpdateStatusRequestTypeID is TL type id of AccountUpdateStatusRequest.
const AccountUpdateStatusRequestTypeID = 0x6628562c

// Ensuring interfaces in compile-time for AccountUpdateStatusRequest.
var (
	_ bin.Encoder     = &AccountUpdateStatusRequest{}
	_ bin.Decoder     = &AccountUpdateStatusRequest{}
	_ bin.BareEncoder = &AccountUpdateStatusRequest{}
	_ bin.BareDecoder = &AccountUpdateStatusRequest{}
)

func (u *AccountUpdateStatusRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Offline == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateStatusRequest) String() string {
	if u == nil {
		return "AccountUpdateStatusRequest(nil)"
	}
	type Alias AccountUpdateStatusRequest
	return fmt.Sprintf("AccountUpdateStatusRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateStatusRequest from given interface.
func (u *AccountUpdateStatusRequest) FillFrom(from interface {
	GetOffline() (value bool)
}) {
	u.Offline = from.GetOffline()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateStatusRequest) TypeID() uint32 {
	return AccountUpdateStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateStatusRequest) TypeName() string {
	return "account.updateStatus"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateStatus",
		ID:   AccountUpdateStatusRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Offline",
			SchemaName: "offline",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdateStatusRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateStatus#6628562c as nil")
	}
	b.PutID(AccountUpdateStatusRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateStatusRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateStatus#6628562c as nil")
	}
	b.PutBool(u.Offline)
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateStatusRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateStatus#6628562c to nil")
	}
	if err := b.ConsumeID(AccountUpdateStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateStatus#6628562c: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateStatusRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateStatus#6628562c to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateStatus#6628562c: field offline: %w", err)
		}
		u.Offline = value
	}
	return nil
}

// GetOffline returns value of Offline field.
func (u *AccountUpdateStatusRequest) GetOffline() (value bool) {
	if u == nil {
		return
	}
	return u.Offline
}

// AccountUpdateStatus invokes method account.updateStatus#6628562c returning error if any.
// Updates online user status.
//
// Possible errors:
//
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//
// See https://core.telegram.org/method/account.updateStatus for reference.
func (c *Client) AccountUpdateStatus(ctx context.Context, offline bool) (bool, error) {
	var result BoolBox

	request := &AccountUpdateStatusRequest{
		Offline: offline,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}