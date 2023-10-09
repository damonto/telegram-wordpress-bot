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

// MediaAreaVenue represents TL type `mediaAreaVenue#be82db9c`.
//
// See https://core.telegram.org/constructor/mediaAreaVenue for reference.
type MediaAreaVenue struct {
	// Coordinates field of MediaAreaVenue.
	Coordinates MediaAreaCoordinates
	// Geo field of MediaAreaVenue.
	Geo GeoPointClass
	// Title field of MediaAreaVenue.
	Title string
	// Address field of MediaAreaVenue.
	Address string
	// Provider field of MediaAreaVenue.
	Provider string
	// VenueID field of MediaAreaVenue.
	VenueID string
	// VenueType field of MediaAreaVenue.
	VenueType string
}

// MediaAreaVenueTypeID is TL type id of MediaAreaVenue.
const MediaAreaVenueTypeID = 0xbe82db9c

// construct implements constructor of MediaAreaClass.
func (m MediaAreaVenue) construct() MediaAreaClass { return &m }

// Ensuring interfaces in compile-time for MediaAreaVenue.
var (
	_ bin.Encoder     = &MediaAreaVenue{}
	_ bin.Decoder     = &MediaAreaVenue{}
	_ bin.BareEncoder = &MediaAreaVenue{}
	_ bin.BareDecoder = &MediaAreaVenue{}

	_ MediaAreaClass = &MediaAreaVenue{}
)

func (m *MediaAreaVenue) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Coordinates.Zero()) {
		return false
	}
	if !(m.Geo == nil) {
		return false
	}
	if !(m.Title == "") {
		return false
	}
	if !(m.Address == "") {
		return false
	}
	if !(m.Provider == "") {
		return false
	}
	if !(m.VenueID == "") {
		return false
	}
	if !(m.VenueType == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MediaAreaVenue) String() string {
	if m == nil {
		return "MediaAreaVenue(nil)"
	}
	type Alias MediaAreaVenue
	return fmt.Sprintf("MediaAreaVenue%+v", Alias(*m))
}

// FillFrom fills MediaAreaVenue from given interface.
func (m *MediaAreaVenue) FillFrom(from interface {
	GetCoordinates() (value MediaAreaCoordinates)
	GetGeo() (value GeoPointClass)
	GetTitle() (value string)
	GetAddress() (value string)
	GetProvider() (value string)
	GetVenueID() (value string)
	GetVenueType() (value string)
}) {
	m.Coordinates = from.GetCoordinates()
	m.Geo = from.GetGeo()
	m.Title = from.GetTitle()
	m.Address = from.GetAddress()
	m.Provider = from.GetProvider()
	m.VenueID = from.GetVenueID()
	m.VenueType = from.GetVenueType()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MediaAreaVenue) TypeID() uint32 {
	return MediaAreaVenueTypeID
}

// TypeName returns name of type in TL schema.
func (*MediaAreaVenue) TypeName() string {
	return "mediaAreaVenue"
}

// TypeInfo returns info about TL type.
func (m *MediaAreaVenue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "mediaAreaVenue",
		ID:   MediaAreaVenueTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Coordinates",
			SchemaName: "coordinates",
		},
		{
			Name:       "Geo",
			SchemaName: "geo",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Address",
			SchemaName: "address",
		},
		{
			Name:       "Provider",
			SchemaName: "provider",
		},
		{
			Name:       "VenueID",
			SchemaName: "venue_id",
		},
		{
			Name:       "VenueType",
			SchemaName: "venue_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MediaAreaVenue) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode mediaAreaVenue#be82db9c as nil")
	}
	b.PutID(MediaAreaVenueTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MediaAreaVenue) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode mediaAreaVenue#be82db9c as nil")
	}
	if err := m.Coordinates.Encode(b); err != nil {
		return fmt.Errorf("unable to encode mediaAreaVenue#be82db9c: field coordinates: %w", err)
	}
	if m.Geo == nil {
		return fmt.Errorf("unable to encode mediaAreaVenue#be82db9c: field geo is nil")
	}
	if err := m.Geo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode mediaAreaVenue#be82db9c: field geo: %w", err)
	}
	b.PutString(m.Title)
	b.PutString(m.Address)
	b.PutString(m.Provider)
	b.PutString(m.VenueID)
	b.PutString(m.VenueType)
	return nil
}

// Decode implements bin.Decoder.
func (m *MediaAreaVenue) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode mediaAreaVenue#be82db9c to nil")
	}
	if err := b.ConsumeID(MediaAreaVenueTypeID); err != nil {
		return fmt.Errorf("unable to decode mediaAreaVenue#be82db9c: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MediaAreaVenue) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode mediaAreaVenue#be82db9c to nil")
	}
	{
		if err := m.Coordinates.Decode(b); err != nil {
			return fmt.Errorf("unable to decode mediaAreaVenue#be82db9c: field coordinates: %w", err)
		}
	}
	{
		value, err := DecodeGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode mediaAreaVenue#be82db9c: field geo: %w", err)
		}
		m.Geo = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaAreaVenue#be82db9c: field title: %w", err)
		}
		m.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaAreaVenue#be82db9c: field address: %w", err)
		}
		m.Address = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaAreaVenue#be82db9c: field provider: %w", err)
		}
		m.Provider = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaAreaVenue#be82db9c: field venue_id: %w", err)
		}
		m.VenueID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode mediaAreaVenue#be82db9c: field venue_type: %w", err)
		}
		m.VenueType = value
	}
	return nil
}

// GetCoordinates returns value of Coordinates field.
func (m *MediaAreaVenue) GetCoordinates() (value MediaAreaCoordinates) {
	if m == nil {
		return
	}
	return m.Coordinates
}

// GetGeo returns value of Geo field.
func (m *MediaAreaVenue) GetGeo() (value GeoPointClass) {
	if m == nil {
		return
	}
	return m.Geo
}

// GetTitle returns value of Title field.
func (m *MediaAreaVenue) GetTitle() (value string) {
	if m == nil {
		return
	}
	return m.Title
}

// GetAddress returns value of Address field.
func (m *MediaAreaVenue) GetAddress() (value string) {
	if m == nil {
		return
	}
	return m.Address
}

// GetProvider returns value of Provider field.
func (m *MediaAreaVenue) GetProvider() (value string) {
	if m == nil {
		return
	}
	return m.Provider
}

// GetVenueID returns value of VenueID field.
func (m *MediaAreaVenue) GetVenueID() (value string) {
	if m == nil {
		return
	}
	return m.VenueID
}

// GetVenueType returns value of VenueType field.
func (m *MediaAreaVenue) GetVenueType() (value string) {
	if m == nil {
		return
	}
	return m.VenueType
}

// InputMediaAreaVenue represents TL type `inputMediaAreaVenue#b282217f`.
//
// See https://core.telegram.org/constructor/inputMediaAreaVenue for reference.
type InputMediaAreaVenue struct {
	// Coordinates field of InputMediaAreaVenue.
	Coordinates MediaAreaCoordinates
	// QueryID field of InputMediaAreaVenue.
	QueryID int64
	// ResultID field of InputMediaAreaVenue.
	ResultID string
}

// InputMediaAreaVenueTypeID is TL type id of InputMediaAreaVenue.
const InputMediaAreaVenueTypeID = 0xb282217f

// construct implements constructor of MediaAreaClass.
func (i InputMediaAreaVenue) construct() MediaAreaClass { return &i }

// Ensuring interfaces in compile-time for InputMediaAreaVenue.
var (
	_ bin.Encoder     = &InputMediaAreaVenue{}
	_ bin.Decoder     = &InputMediaAreaVenue{}
	_ bin.BareEncoder = &InputMediaAreaVenue{}
	_ bin.BareDecoder = &InputMediaAreaVenue{}

	_ MediaAreaClass = &InputMediaAreaVenue{}
)

func (i *InputMediaAreaVenue) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Coordinates.Zero()) {
		return false
	}
	if !(i.QueryID == 0) {
		return false
	}
	if !(i.ResultID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputMediaAreaVenue) String() string {
	if i == nil {
		return "InputMediaAreaVenue(nil)"
	}
	type Alias InputMediaAreaVenue
	return fmt.Sprintf("InputMediaAreaVenue%+v", Alias(*i))
}

// FillFrom fills InputMediaAreaVenue from given interface.
func (i *InputMediaAreaVenue) FillFrom(from interface {
	GetCoordinates() (value MediaAreaCoordinates)
	GetQueryID() (value int64)
	GetResultID() (value string)
}) {
	i.Coordinates = from.GetCoordinates()
	i.QueryID = from.GetQueryID()
	i.ResultID = from.GetResultID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputMediaAreaVenue) TypeID() uint32 {
	return InputMediaAreaVenueTypeID
}

// TypeName returns name of type in TL schema.
func (*InputMediaAreaVenue) TypeName() string {
	return "inputMediaAreaVenue"
}

// TypeInfo returns info about TL type.
func (i *InputMediaAreaVenue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputMediaAreaVenue",
		ID:   InputMediaAreaVenueTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Coordinates",
			SchemaName: "coordinates",
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "ResultID",
			SchemaName: "result_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputMediaAreaVenue) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMediaAreaVenue#b282217f as nil")
	}
	b.PutID(InputMediaAreaVenueTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputMediaAreaVenue) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMediaAreaVenue#b282217f as nil")
	}
	if err := i.Coordinates.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputMediaAreaVenue#b282217f: field coordinates: %w", err)
	}
	b.PutLong(i.QueryID)
	b.PutString(i.ResultID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMediaAreaVenue) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMediaAreaVenue#b282217f to nil")
	}
	if err := b.ConsumeID(InputMediaAreaVenueTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMediaAreaVenue#b282217f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputMediaAreaVenue) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMediaAreaVenue#b282217f to nil")
	}
	{
		if err := i.Coordinates.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputMediaAreaVenue#b282217f: field coordinates: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputMediaAreaVenue#b282217f: field query_id: %w", err)
		}
		i.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputMediaAreaVenue#b282217f: field result_id: %w", err)
		}
		i.ResultID = value
	}
	return nil
}

// GetCoordinates returns value of Coordinates field.
func (i *InputMediaAreaVenue) GetCoordinates() (value MediaAreaCoordinates) {
	if i == nil {
		return
	}
	return i.Coordinates
}

// GetQueryID returns value of QueryID field.
func (i *InputMediaAreaVenue) GetQueryID() (value int64) {
	if i == nil {
		return
	}
	return i.QueryID
}

// GetResultID returns value of ResultID field.
func (i *InputMediaAreaVenue) GetResultID() (value string) {
	if i == nil {
		return
	}
	return i.ResultID
}

// MediaAreaGeoPoint represents TL type `mediaAreaGeoPoint#df8b3b22`.
//
// See https://core.telegram.org/constructor/mediaAreaGeoPoint for reference.
type MediaAreaGeoPoint struct {
	// Coordinates field of MediaAreaGeoPoint.
	Coordinates MediaAreaCoordinates
	// Geo field of MediaAreaGeoPoint.
	Geo GeoPointClass
}

// MediaAreaGeoPointTypeID is TL type id of MediaAreaGeoPoint.
const MediaAreaGeoPointTypeID = 0xdf8b3b22

// construct implements constructor of MediaAreaClass.
func (m MediaAreaGeoPoint) construct() MediaAreaClass { return &m }

// Ensuring interfaces in compile-time for MediaAreaGeoPoint.
var (
	_ bin.Encoder     = &MediaAreaGeoPoint{}
	_ bin.Decoder     = &MediaAreaGeoPoint{}
	_ bin.BareEncoder = &MediaAreaGeoPoint{}
	_ bin.BareDecoder = &MediaAreaGeoPoint{}

	_ MediaAreaClass = &MediaAreaGeoPoint{}
)

func (m *MediaAreaGeoPoint) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Coordinates.Zero()) {
		return false
	}
	if !(m.Geo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MediaAreaGeoPoint) String() string {
	if m == nil {
		return "MediaAreaGeoPoint(nil)"
	}
	type Alias MediaAreaGeoPoint
	return fmt.Sprintf("MediaAreaGeoPoint%+v", Alias(*m))
}

// FillFrom fills MediaAreaGeoPoint from given interface.
func (m *MediaAreaGeoPoint) FillFrom(from interface {
	GetCoordinates() (value MediaAreaCoordinates)
	GetGeo() (value GeoPointClass)
}) {
	m.Coordinates = from.GetCoordinates()
	m.Geo = from.GetGeo()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MediaAreaGeoPoint) TypeID() uint32 {
	return MediaAreaGeoPointTypeID
}

// TypeName returns name of type in TL schema.
func (*MediaAreaGeoPoint) TypeName() string {
	return "mediaAreaGeoPoint"
}

// TypeInfo returns info about TL type.
func (m *MediaAreaGeoPoint) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "mediaAreaGeoPoint",
		ID:   MediaAreaGeoPointTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Coordinates",
			SchemaName: "coordinates",
		},
		{
			Name:       "Geo",
			SchemaName: "geo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MediaAreaGeoPoint) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode mediaAreaGeoPoint#df8b3b22 as nil")
	}
	b.PutID(MediaAreaGeoPointTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MediaAreaGeoPoint) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode mediaAreaGeoPoint#df8b3b22 as nil")
	}
	if err := m.Coordinates.Encode(b); err != nil {
		return fmt.Errorf("unable to encode mediaAreaGeoPoint#df8b3b22: field coordinates: %w", err)
	}
	if m.Geo == nil {
		return fmt.Errorf("unable to encode mediaAreaGeoPoint#df8b3b22: field geo is nil")
	}
	if err := m.Geo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode mediaAreaGeoPoint#df8b3b22: field geo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MediaAreaGeoPoint) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode mediaAreaGeoPoint#df8b3b22 to nil")
	}
	if err := b.ConsumeID(MediaAreaGeoPointTypeID); err != nil {
		return fmt.Errorf("unable to decode mediaAreaGeoPoint#df8b3b22: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MediaAreaGeoPoint) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode mediaAreaGeoPoint#df8b3b22 to nil")
	}
	{
		if err := m.Coordinates.Decode(b); err != nil {
			return fmt.Errorf("unable to decode mediaAreaGeoPoint#df8b3b22: field coordinates: %w", err)
		}
	}
	{
		value, err := DecodeGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode mediaAreaGeoPoint#df8b3b22: field geo: %w", err)
		}
		m.Geo = value
	}
	return nil
}

// GetCoordinates returns value of Coordinates field.
func (m *MediaAreaGeoPoint) GetCoordinates() (value MediaAreaCoordinates) {
	if m == nil {
		return
	}
	return m.Coordinates
}

// GetGeo returns value of Geo field.
func (m *MediaAreaGeoPoint) GetGeo() (value GeoPointClass) {
	if m == nil {
		return
	}
	return m.Geo
}

// MediaAreaSuggestedReaction represents TL type `mediaAreaSuggestedReaction#14455871`.
//
// See https://core.telegram.org/constructor/mediaAreaSuggestedReaction for reference.
type MediaAreaSuggestedReaction struct {
	// Flags field of MediaAreaSuggestedReaction.
	Flags bin.Fields
	// Dark field of MediaAreaSuggestedReaction.
	Dark bool
	// Flipped field of MediaAreaSuggestedReaction.
	Flipped bool
	// Coordinates field of MediaAreaSuggestedReaction.
	Coordinates MediaAreaCoordinates
	// Reaction field of MediaAreaSuggestedReaction.
	Reaction ReactionClass
}

// MediaAreaSuggestedReactionTypeID is TL type id of MediaAreaSuggestedReaction.
const MediaAreaSuggestedReactionTypeID = 0x14455871

// construct implements constructor of MediaAreaClass.
func (m MediaAreaSuggestedReaction) construct() MediaAreaClass { return &m }

// Ensuring interfaces in compile-time for MediaAreaSuggestedReaction.
var (
	_ bin.Encoder     = &MediaAreaSuggestedReaction{}
	_ bin.Decoder     = &MediaAreaSuggestedReaction{}
	_ bin.BareEncoder = &MediaAreaSuggestedReaction{}
	_ bin.BareDecoder = &MediaAreaSuggestedReaction{}

	_ MediaAreaClass = &MediaAreaSuggestedReaction{}
)

func (m *MediaAreaSuggestedReaction) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Dark == false) {
		return false
	}
	if !(m.Flipped == false) {
		return false
	}
	if !(m.Coordinates.Zero()) {
		return false
	}
	if !(m.Reaction == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MediaAreaSuggestedReaction) String() string {
	if m == nil {
		return "MediaAreaSuggestedReaction(nil)"
	}
	type Alias MediaAreaSuggestedReaction
	return fmt.Sprintf("MediaAreaSuggestedReaction%+v", Alias(*m))
}

// FillFrom fills MediaAreaSuggestedReaction from given interface.
func (m *MediaAreaSuggestedReaction) FillFrom(from interface {
	GetDark() (value bool)
	GetFlipped() (value bool)
	GetCoordinates() (value MediaAreaCoordinates)
	GetReaction() (value ReactionClass)
}) {
	m.Dark = from.GetDark()
	m.Flipped = from.GetFlipped()
	m.Coordinates = from.GetCoordinates()
	m.Reaction = from.GetReaction()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MediaAreaSuggestedReaction) TypeID() uint32 {
	return MediaAreaSuggestedReactionTypeID
}

// TypeName returns name of type in TL schema.
func (*MediaAreaSuggestedReaction) TypeName() string {
	return "mediaAreaSuggestedReaction"
}

// TypeInfo returns info about TL type.
func (m *MediaAreaSuggestedReaction) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "mediaAreaSuggestedReaction",
		ID:   MediaAreaSuggestedReactionTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dark",
			SchemaName: "dark",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "Flipped",
			SchemaName: "flipped",
			Null:       !m.Flags.Has(1),
		},
		{
			Name:       "Coordinates",
			SchemaName: "coordinates",
		},
		{
			Name:       "Reaction",
			SchemaName: "reaction",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (m *MediaAreaSuggestedReaction) SetFlags() {
	if !(m.Dark == false) {
		m.Flags.Set(0)
	}
	if !(m.Flipped == false) {
		m.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (m *MediaAreaSuggestedReaction) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode mediaAreaSuggestedReaction#14455871 as nil")
	}
	b.PutID(MediaAreaSuggestedReactionTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MediaAreaSuggestedReaction) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode mediaAreaSuggestedReaction#14455871 as nil")
	}
	m.SetFlags()
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode mediaAreaSuggestedReaction#14455871: field flags: %w", err)
	}
	if err := m.Coordinates.Encode(b); err != nil {
		return fmt.Errorf("unable to encode mediaAreaSuggestedReaction#14455871: field coordinates: %w", err)
	}
	if m.Reaction == nil {
		return fmt.Errorf("unable to encode mediaAreaSuggestedReaction#14455871: field reaction is nil")
	}
	if err := m.Reaction.Encode(b); err != nil {
		return fmt.Errorf("unable to encode mediaAreaSuggestedReaction#14455871: field reaction: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MediaAreaSuggestedReaction) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode mediaAreaSuggestedReaction#14455871 to nil")
	}
	if err := b.ConsumeID(MediaAreaSuggestedReactionTypeID); err != nil {
		return fmt.Errorf("unable to decode mediaAreaSuggestedReaction#14455871: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MediaAreaSuggestedReaction) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode mediaAreaSuggestedReaction#14455871 to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode mediaAreaSuggestedReaction#14455871: field flags: %w", err)
		}
	}
	m.Dark = m.Flags.Has(0)
	m.Flipped = m.Flags.Has(1)
	{
		if err := m.Coordinates.Decode(b); err != nil {
			return fmt.Errorf("unable to decode mediaAreaSuggestedReaction#14455871: field coordinates: %w", err)
		}
	}
	{
		value, err := DecodeReaction(b)
		if err != nil {
			return fmt.Errorf("unable to decode mediaAreaSuggestedReaction#14455871: field reaction: %w", err)
		}
		m.Reaction = value
	}
	return nil
}

// SetDark sets value of Dark conditional field.
func (m *MediaAreaSuggestedReaction) SetDark(value bool) {
	if value {
		m.Flags.Set(0)
		m.Dark = true
	} else {
		m.Flags.Unset(0)
		m.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (m *MediaAreaSuggestedReaction) GetDark() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(0)
}

// SetFlipped sets value of Flipped conditional field.
func (m *MediaAreaSuggestedReaction) SetFlipped(value bool) {
	if value {
		m.Flags.Set(1)
		m.Flipped = true
	} else {
		m.Flags.Unset(1)
		m.Flipped = false
	}
}

// GetFlipped returns value of Flipped conditional field.
func (m *MediaAreaSuggestedReaction) GetFlipped() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(1)
}

// GetCoordinates returns value of Coordinates field.
func (m *MediaAreaSuggestedReaction) GetCoordinates() (value MediaAreaCoordinates) {
	if m == nil {
		return
	}
	return m.Coordinates
}

// GetReaction returns value of Reaction field.
func (m *MediaAreaSuggestedReaction) GetReaction() (value ReactionClass) {
	if m == nil {
		return
	}
	return m.Reaction
}

// MediaAreaClassName is schema name of MediaAreaClass.
const MediaAreaClassName = "MediaArea"

// MediaAreaClass represents MediaArea generic type.
//
// See https://core.telegram.org/type/MediaArea for reference.
//
// Example:
//
//	g, err := tg.DecodeMediaArea(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MediaAreaVenue: // mediaAreaVenue#be82db9c
//	case *tg.InputMediaAreaVenue: // inputMediaAreaVenue#b282217f
//	case *tg.MediaAreaGeoPoint: // mediaAreaGeoPoint#df8b3b22
//	case *tg.MediaAreaSuggestedReaction: // mediaAreaSuggestedReaction#14455871
//	default: panic(v)
//	}
type MediaAreaClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MediaAreaClass

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

	// Coordinates field of MediaAreaVenue.
	GetCoordinates() (value MediaAreaCoordinates)
}

// DecodeMediaArea implements binary de-serialization for MediaAreaClass.
func DecodeMediaArea(buf *bin.Buffer) (MediaAreaClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MediaAreaVenueTypeID:
		// Decoding mediaAreaVenue#be82db9c.
		v := MediaAreaVenue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MediaAreaClass: %w", err)
		}
		return &v, nil
	case InputMediaAreaVenueTypeID:
		// Decoding inputMediaAreaVenue#b282217f.
		v := InputMediaAreaVenue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MediaAreaClass: %w", err)
		}
		return &v, nil
	case MediaAreaGeoPointTypeID:
		// Decoding mediaAreaGeoPoint#df8b3b22.
		v := MediaAreaGeoPoint{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MediaAreaClass: %w", err)
		}
		return &v, nil
	case MediaAreaSuggestedReactionTypeID:
		// Decoding mediaAreaSuggestedReaction#14455871.
		v := MediaAreaSuggestedReaction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MediaAreaClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MediaAreaClass: %w", bin.NewUnexpectedID(id))
	}
}

// MediaArea boxes the MediaAreaClass providing a helper.
type MediaAreaBox struct {
	MediaArea MediaAreaClass
}

// Decode implements bin.Decoder for MediaAreaBox.
func (b *MediaAreaBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MediaAreaBox to nil")
	}
	v, err := DecodeMediaArea(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MediaArea = v
	return nil
}

// Encode implements bin.Encode for MediaAreaBox.
func (b *MediaAreaBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MediaArea == nil {
		return fmt.Errorf("unable to encode MediaAreaClass as nil")
	}
	return b.MediaArea.Encode(buf)
}
