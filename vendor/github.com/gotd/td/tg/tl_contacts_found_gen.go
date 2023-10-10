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

// ContactsFound represents TL type `contacts.found#b3134d9d`.
// Users found by name substring and auxiliary data.
//
// See https://core.telegram.org/constructor/contacts.found for reference.
type ContactsFound struct {
	// Personalized results
	MyResults []PeerClass
	// List of found user identifiers
	Results []PeerClass
	// Found chats
	Chats []ChatClass
	// List of users
	Users []UserClass
}

// ContactsFoundTypeID is TL type id of ContactsFound.
const ContactsFoundTypeID = 0xb3134d9d

// Ensuring interfaces in compile-time for ContactsFound.
var (
	_ bin.Encoder     = &ContactsFound{}
	_ bin.Decoder     = &ContactsFound{}
	_ bin.BareEncoder = &ContactsFound{}
	_ bin.BareDecoder = &ContactsFound{}
)

func (f *ContactsFound) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.MyResults == nil) {
		return false
	}
	if !(f.Results == nil) {
		return false
	}
	if !(f.Chats == nil) {
		return false
	}
	if !(f.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *ContactsFound) String() string {
	if f == nil {
		return "ContactsFound(nil)"
	}
	type Alias ContactsFound
	return fmt.Sprintf("ContactsFound%+v", Alias(*f))
}

// FillFrom fills ContactsFound from given interface.
func (f *ContactsFound) FillFrom(from interface {
	GetMyResults() (value []PeerClass)
	GetResults() (value []PeerClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	f.MyResults = from.GetMyResults()
	f.Results = from.GetResults()
	f.Chats = from.GetChats()
	f.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsFound) TypeID() uint32 {
	return ContactsFoundTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsFound) TypeName() string {
	return "contacts.found"
}

// TypeInfo returns info about TL type.
func (f *ContactsFound) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.found",
		ID:   ContactsFoundTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MyResults",
			SchemaName: "my_results",
		},
		{
			Name:       "Results",
			SchemaName: "results",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *ContactsFound) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode contacts.found#b3134d9d as nil")
	}
	b.PutID(ContactsFoundTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *ContactsFound) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode contacts.found#b3134d9d as nil")
	}
	b.PutVectorHeader(len(f.MyResults))
	for idx, v := range f.MyResults {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field my_results element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field my_results element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Results))
	for idx, v := range f.Results {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field results element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field results element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Chats))
	for idx, v := range f.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Users))
	for idx, v := range f.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *ContactsFound) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode contacts.found#b3134d9d to nil")
	}
	if err := b.ConsumeID(ContactsFoundTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.found#b3134d9d: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *ContactsFound) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode contacts.found#b3134d9d to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.found#b3134d9d: field my_results: %w", err)
		}

		if headerLen > 0 {
			f.MyResults = make([]PeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.found#b3134d9d: field my_results: %w", err)
			}
			f.MyResults = append(f.MyResults, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.found#b3134d9d: field results: %w", err)
		}

		if headerLen > 0 {
			f.Results = make([]PeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.found#b3134d9d: field results: %w", err)
			}
			f.Results = append(f.Results, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.found#b3134d9d: field chats: %w", err)
		}

		if headerLen > 0 {
			f.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.found#b3134d9d: field chats: %w", err)
			}
			f.Chats = append(f.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.found#b3134d9d: field users: %w", err)
		}

		if headerLen > 0 {
			f.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.found#b3134d9d: field users: %w", err)
			}
			f.Users = append(f.Users, value)
		}
	}
	return nil
}

// GetMyResults returns value of MyResults field.
func (f *ContactsFound) GetMyResults() (value []PeerClass) {
	if f == nil {
		return
	}
	return f.MyResults
}

// GetResults returns value of Results field.
func (f *ContactsFound) GetResults() (value []PeerClass) {
	if f == nil {
		return
	}
	return f.Results
}

// GetChats returns value of Chats field.
func (f *ContactsFound) GetChats() (value []ChatClass) {
	if f == nil {
		return
	}
	return f.Chats
}

// GetUsers returns value of Users field.
func (f *ContactsFound) GetUsers() (value []UserClass) {
	if f == nil {
		return
	}
	return f.Users
}

// MapMyResults returns field MyResults wrapped in PeerClassArray helper.
func (f *ContactsFound) MapMyResults() (value PeerClassArray) {
	return PeerClassArray(f.MyResults)
}

// MapResults returns field Results wrapped in PeerClassArray helper.
func (f *ContactsFound) MapResults() (value PeerClassArray) {
	return PeerClassArray(f.Results)
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (f *ContactsFound) MapChats() (value ChatClassArray) {
	return ChatClassArray(f.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (f *ContactsFound) MapUsers() (value UserClassArray) {
	return UserClassArray(f.Users)
}