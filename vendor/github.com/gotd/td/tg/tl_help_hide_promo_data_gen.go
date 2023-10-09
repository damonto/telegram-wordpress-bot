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

// HelpHidePromoDataRequest represents TL type `help.hidePromoData#1e251c95`.
// Hide MTProxy/Public Service Announcement information
//
// See https://core.telegram.org/method/help.hidePromoData for reference.
type HelpHidePromoDataRequest struct {
	// Peer to hide
	Peer InputPeerClass
}

// HelpHidePromoDataRequestTypeID is TL type id of HelpHidePromoDataRequest.
const HelpHidePromoDataRequestTypeID = 0x1e251c95

// Ensuring interfaces in compile-time for HelpHidePromoDataRequest.
var (
	_ bin.Encoder     = &HelpHidePromoDataRequest{}
	_ bin.Decoder     = &HelpHidePromoDataRequest{}
	_ bin.BareEncoder = &HelpHidePromoDataRequest{}
	_ bin.BareDecoder = &HelpHidePromoDataRequest{}
)

func (h *HelpHidePromoDataRequest) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *HelpHidePromoDataRequest) String() string {
	if h == nil {
		return "HelpHidePromoDataRequest(nil)"
	}
	type Alias HelpHidePromoDataRequest
	return fmt.Sprintf("HelpHidePromoDataRequest%+v", Alias(*h))
}

// FillFrom fills HelpHidePromoDataRequest from given interface.
func (h *HelpHidePromoDataRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	h.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpHidePromoDataRequest) TypeID() uint32 {
	return HelpHidePromoDataRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpHidePromoDataRequest) TypeName() string {
	return "help.hidePromoData"
}

// TypeInfo returns info about TL type.
func (h *HelpHidePromoDataRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.hidePromoData",
		ID:   HelpHidePromoDataRequestTypeID,
	}
	if h == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (h *HelpHidePromoDataRequest) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode help.hidePromoData#1e251c95 as nil")
	}
	b.PutID(HelpHidePromoDataRequestTypeID)
	return h.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (h *HelpHidePromoDataRequest) EncodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode help.hidePromoData#1e251c95 as nil")
	}
	if h.Peer == nil {
		return fmt.Errorf("unable to encode help.hidePromoData#1e251c95: field peer is nil")
	}
	if err := h.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.hidePromoData#1e251c95: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (h *HelpHidePromoDataRequest) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode help.hidePromoData#1e251c95 to nil")
	}
	if err := b.ConsumeID(HelpHidePromoDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.hidePromoData#1e251c95: %w", err)
	}
	return h.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (h *HelpHidePromoDataRequest) DecodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode help.hidePromoData#1e251c95 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.hidePromoData#1e251c95: field peer: %w", err)
		}
		h.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (h *HelpHidePromoDataRequest) GetPeer() (value InputPeerClass) {
	if h == nil {
		return
	}
	return h.Peer
}

// HelpHidePromoData invokes method help.hidePromoData#1e251c95 returning error if any.
// Hide MTProxy/Public Service Announcement information
//
// See https://core.telegram.org/method/help.hidePromoData for reference.
func (c *Client) HelpHidePromoData(ctx context.Context, peer InputPeerClass) (bool, error) {
	var result BoolBox

	request := &HelpHidePromoDataRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
