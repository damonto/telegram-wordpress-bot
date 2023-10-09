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

// MessagesGetPollVotesRequest represents TL type `messages.getPollVotes#b86e380e`.
// Get poll results for non-anonymous polls
//
// See https://core.telegram.org/method/messages.getPollVotes for reference.
type MessagesGetPollVotesRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Chat where the poll was sent
	Peer InputPeerClass
	// Message ID
	ID int
	// Get only results for the specified poll option
	//
	// Use SetOption and GetOption helpers.
	Option []byte
	// Offset for results, taken from the next_offset field of messages.votesList¹,
	// initially an empty string. Note: if no more results are available, the method call
	// will return an empty next_offset; thus, avoid providing the next_offset returned in
	// messages.votesList² if it is empty, to avoid an infinite loop.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messages.votesList
	//  2) https://core.telegram.org/constructor/messages.votesList
	//
	// Use SetOffset and GetOffset helpers.
	Offset string
	// Number of results to return
	Limit int
}

// MessagesGetPollVotesRequestTypeID is TL type id of MessagesGetPollVotesRequest.
const MessagesGetPollVotesRequestTypeID = 0xb86e380e

// Ensuring interfaces in compile-time for MessagesGetPollVotesRequest.
var (
	_ bin.Encoder     = &MessagesGetPollVotesRequest{}
	_ bin.Decoder     = &MessagesGetPollVotesRequest{}
	_ bin.BareEncoder = &MessagesGetPollVotesRequest{}
	_ bin.BareDecoder = &MessagesGetPollVotesRequest{}
)

func (g *MessagesGetPollVotesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == 0) {
		return false
	}
	if !(g.Option == nil) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetPollVotesRequest) String() string {
	if g == nil {
		return "MessagesGetPollVotesRequest(nil)"
	}
	type Alias MessagesGetPollVotesRequest
	return fmt.Sprintf("MessagesGetPollVotesRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetPollVotesRequest from given interface.
func (g *MessagesGetPollVotesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value int)
	GetOption() (value []byte, ok bool)
	GetOffset() (value string, ok bool)
	GetLimit() (value int)
}) {
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
	if val, ok := from.GetOption(); ok {
		g.Option = val
	}

	if val, ok := from.GetOffset(); ok {
		g.Offset = val
	}

	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetPollVotesRequest) TypeID() uint32 {
	return MessagesGetPollVotesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetPollVotesRequest) TypeName() string {
	return "messages.getPollVotes"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetPollVotesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getPollVotes",
		ID:   MessagesGetPollVotesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Option",
			SchemaName: "option",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *MessagesGetPollVotesRequest) SetFlags() {
	if !(g.Option == nil) {
		g.Flags.Set(0)
	}
	if !(g.Offset == "") {
		g.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (g *MessagesGetPollVotesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPollVotes#b86e380e as nil")
	}
	b.PutID(MessagesGetPollVotesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetPollVotesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPollVotes#b86e380e as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getPollVotes#b86e380e: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getPollVotes#b86e380e: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getPollVotes#b86e380e: field peer: %w", err)
	}
	b.PutInt(g.ID)
	if g.Flags.Has(0) {
		b.PutBytes(g.Option)
	}
	if g.Flags.Has(1) {
		b.PutString(g.Offset)
	}
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetPollVotesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPollVotes#b86e380e to nil")
	}
	if err := b.ConsumeID(MessagesGetPollVotesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPollVotes#b86e380e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetPollVotesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPollVotes#b86e380e to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getPollVotes#b86e380e: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPollVotes#b86e380e: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPollVotes#b86e380e: field id: %w", err)
		}
		g.ID = value
	}
	if g.Flags.Has(0) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPollVotes#b86e380e: field option: %w", err)
		}
		g.Option = value
	}
	if g.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPollVotes#b86e380e: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPollVotes#b86e380e: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetPollVotesRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *MessagesGetPollVotesRequest) GetID() (value int) {
	if g == nil {
		return
	}
	return g.ID
}

// SetOption sets value of Option conditional field.
func (g *MessagesGetPollVotesRequest) SetOption(value []byte) {
	g.Flags.Set(0)
	g.Option = value
}

// GetOption returns value of Option conditional field and
// boolean which is true if field was set.
func (g *MessagesGetPollVotesRequest) GetOption() (value []byte, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Option, true
}

// SetOffset sets value of Offset conditional field.
func (g *MessagesGetPollVotesRequest) SetOffset(value string) {
	g.Flags.Set(1)
	g.Offset = value
}

// GetOffset returns value of Offset conditional field and
// boolean which is true if field was set.
func (g *MessagesGetPollVotesRequest) GetOffset() (value string, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(1) {
		return value, false
	}
	return g.Offset, true
}

// GetLimit returns value of Limit field.
func (g *MessagesGetPollVotesRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// MessagesGetPollVotes invokes method messages.getPollVotes#b86e380e returning error if any.
// Get poll results for non-anonymous polls
//
// Possible errors:
//
//	403 BROADCAST_FORBIDDEN: Participants of polls in channels should stay anonymous.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	403 POLL_VOTE_REQUIRED: Cast a vote in the poll before calling this method.
//
// See https://core.telegram.org/method/messages.getPollVotes for reference.
func (c *Client) MessagesGetPollVotes(ctx context.Context, request *MessagesGetPollVotesRequest) (*MessagesVotesList, error) {
	var result MessagesVotesList

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
