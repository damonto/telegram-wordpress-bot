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

// AttachMenuBot represents TL type `attachMenuBot#d90d8dfe`.
// Represents a bot web app that can be launched from the attachment menu »¹
//
// Links:
//  1. https://core.telegram.org/api/bots/attach
//
// See https://core.telegram.org/constructor/attachMenuBot for reference.
type AttachMenuBot struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this bot attachment menu entry should be shown in the attachment menu (toggle
	// using messages.toggleBotInAttachMenu¹)
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.toggleBotInAttachMenu
	Inactive bool
	// True, if the bot supports the "settings_button_pressed" event »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/webapps#settings-button-pressed
	HasSettings bool
	// Whether the bot would like to send messages to the user.
	RequestWriteAccess bool
	// ShowInAttachMenu field of AttachMenuBot.
	ShowInAttachMenu bool
	// ShowInSideMenu field of AttachMenuBot.
	ShowInSideMenu bool
	// SideMenuDisclaimerNeeded field of AttachMenuBot.
	SideMenuDisclaimerNeeded bool
	// Bot ID
	BotID int64
	// Attachment menu item name
	ShortName string
	// List of dialog types where this attachment menu entry should be shown
	//
	// Use SetPeerTypes and GetPeerTypes helpers.
	PeerTypes []AttachMenuPeerTypeClass
	// List of platform-specific static icons and animations to use for the attachment menu
	// button
	Icons []AttachMenuBotIcon
}

// AttachMenuBotTypeID is TL type id of AttachMenuBot.
const AttachMenuBotTypeID = 0xd90d8dfe

// Ensuring interfaces in compile-time for AttachMenuBot.
var (
	_ bin.Encoder     = &AttachMenuBot{}
	_ bin.Decoder     = &AttachMenuBot{}
	_ bin.BareEncoder = &AttachMenuBot{}
	_ bin.BareDecoder = &AttachMenuBot{}
)

func (a *AttachMenuBot) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Inactive == false) {
		return false
	}
	if !(a.HasSettings == false) {
		return false
	}
	if !(a.RequestWriteAccess == false) {
		return false
	}
	if !(a.ShowInAttachMenu == false) {
		return false
	}
	if !(a.ShowInSideMenu == false) {
		return false
	}
	if !(a.SideMenuDisclaimerNeeded == false) {
		return false
	}
	if !(a.BotID == 0) {
		return false
	}
	if !(a.ShortName == "") {
		return false
	}
	if !(a.PeerTypes == nil) {
		return false
	}
	if !(a.Icons == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuBot) String() string {
	if a == nil {
		return "AttachMenuBot(nil)"
	}
	type Alias AttachMenuBot
	return fmt.Sprintf("AttachMenuBot%+v", Alias(*a))
}

// FillFrom fills AttachMenuBot from given interface.
func (a *AttachMenuBot) FillFrom(from interface {
	GetInactive() (value bool)
	GetHasSettings() (value bool)
	GetRequestWriteAccess() (value bool)
	GetShowInAttachMenu() (value bool)
	GetShowInSideMenu() (value bool)
	GetSideMenuDisclaimerNeeded() (value bool)
	GetBotID() (value int64)
	GetShortName() (value string)
	GetPeerTypes() (value []AttachMenuPeerTypeClass, ok bool)
	GetIcons() (value []AttachMenuBotIcon)
}) {
	a.Inactive = from.GetInactive()
	a.HasSettings = from.GetHasSettings()
	a.RequestWriteAccess = from.GetRequestWriteAccess()
	a.ShowInAttachMenu = from.GetShowInAttachMenu()
	a.ShowInSideMenu = from.GetShowInSideMenu()
	a.SideMenuDisclaimerNeeded = from.GetSideMenuDisclaimerNeeded()
	a.BotID = from.GetBotID()
	a.ShortName = from.GetShortName()
	if val, ok := from.GetPeerTypes(); ok {
		a.PeerTypes = val
	}

	a.Icons = from.GetIcons()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuBot) TypeID() uint32 {
	return AttachMenuBotTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuBot) TypeName() string {
	return "attachMenuBot"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuBot) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuBot",
		ID:   AttachMenuBotTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Inactive",
			SchemaName: "inactive",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "HasSettings",
			SchemaName: "has_settings",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "RequestWriteAccess",
			SchemaName: "request_write_access",
			Null:       !a.Flags.Has(2),
		},
		{
			Name:       "ShowInAttachMenu",
			SchemaName: "show_in_attach_menu",
			Null:       !a.Flags.Has(3),
		},
		{
			Name:       "ShowInSideMenu",
			SchemaName: "show_in_side_menu",
			Null:       !a.Flags.Has(4),
		},
		{
			Name:       "SideMenuDisclaimerNeeded",
			SchemaName: "side_menu_disclaimer_needed",
			Null:       !a.Flags.Has(5),
		},
		{
			Name:       "BotID",
			SchemaName: "bot_id",
		},
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
		{
			Name:       "PeerTypes",
			SchemaName: "peer_types",
			Null:       !a.Flags.Has(3),
		},
		{
			Name:       "Icons",
			SchemaName: "icons",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *AttachMenuBot) SetFlags() {
	if !(a.Inactive == false) {
		a.Flags.Set(0)
	}
	if !(a.HasSettings == false) {
		a.Flags.Set(1)
	}
	if !(a.RequestWriteAccess == false) {
		a.Flags.Set(2)
	}
	if !(a.ShowInAttachMenu == false) {
		a.Flags.Set(3)
	}
	if !(a.ShowInSideMenu == false) {
		a.Flags.Set(4)
	}
	if !(a.SideMenuDisclaimerNeeded == false) {
		a.Flags.Set(5)
	}
	if !(a.PeerTypes == nil) {
		a.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (a *AttachMenuBot) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBot#d90d8dfe as nil")
	}
	b.PutID(AttachMenuBotTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuBot) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBot#d90d8dfe as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode attachMenuBot#d90d8dfe: field flags: %w", err)
	}
	b.PutLong(a.BotID)
	b.PutString(a.ShortName)
	if a.Flags.Has(3) {
		b.PutVectorHeader(len(a.PeerTypes))
		for idx, v := range a.PeerTypes {
			if v == nil {
				return fmt.Errorf("unable to encode attachMenuBot#d90d8dfe: field peer_types element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode attachMenuBot#d90d8dfe: field peer_types element with index %d: %w", idx, err)
			}
		}
	}
	b.PutVectorHeader(len(a.Icons))
	for idx, v := range a.Icons {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode attachMenuBot#d90d8dfe: field icons element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuBot) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBot#d90d8dfe to nil")
	}
	if err := b.ConsumeID(AttachMenuBotTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuBot#d90d8dfe: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuBot) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBot#d90d8dfe to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode attachMenuBot#d90d8dfe: field flags: %w", err)
		}
	}
	a.Inactive = a.Flags.Has(0)
	a.HasSettings = a.Flags.Has(1)
	a.RequestWriteAccess = a.Flags.Has(2)
	a.ShowInAttachMenu = a.Flags.Has(3)
	a.ShowInSideMenu = a.Flags.Has(4)
	a.SideMenuDisclaimerNeeded = a.Flags.Has(5)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBot#d90d8dfe: field bot_id: %w", err)
		}
		a.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBot#d90d8dfe: field short_name: %w", err)
		}
		a.ShortName = value
	}
	if a.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBot#d90d8dfe: field peer_types: %w", err)
		}

		if headerLen > 0 {
			a.PeerTypes = make([]AttachMenuPeerTypeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeAttachMenuPeerType(b)
			if err != nil {
				return fmt.Errorf("unable to decode attachMenuBot#d90d8dfe: field peer_types: %w", err)
			}
			a.PeerTypes = append(a.PeerTypes, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBot#d90d8dfe: field icons: %w", err)
		}

		if headerLen > 0 {
			a.Icons = make([]AttachMenuBotIcon, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value AttachMenuBotIcon
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode attachMenuBot#d90d8dfe: field icons: %w", err)
			}
			a.Icons = append(a.Icons, value)
		}
	}
	return nil
}

// SetInactive sets value of Inactive conditional field.
func (a *AttachMenuBot) SetInactive(value bool) {
	if value {
		a.Flags.Set(0)
		a.Inactive = true
	} else {
		a.Flags.Unset(0)
		a.Inactive = false
	}
}

// GetInactive returns value of Inactive conditional field.
func (a *AttachMenuBot) GetInactive() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// SetHasSettings sets value of HasSettings conditional field.
func (a *AttachMenuBot) SetHasSettings(value bool) {
	if value {
		a.Flags.Set(1)
		a.HasSettings = true
	} else {
		a.Flags.Unset(1)
		a.HasSettings = false
	}
}

// GetHasSettings returns value of HasSettings conditional field.
func (a *AttachMenuBot) GetHasSettings() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(1)
}

// SetRequestWriteAccess sets value of RequestWriteAccess conditional field.
func (a *AttachMenuBot) SetRequestWriteAccess(value bool) {
	if value {
		a.Flags.Set(2)
		a.RequestWriteAccess = true
	} else {
		a.Flags.Unset(2)
		a.RequestWriteAccess = false
	}
}

// GetRequestWriteAccess returns value of RequestWriteAccess conditional field.
func (a *AttachMenuBot) GetRequestWriteAccess() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(2)
}

// SetShowInAttachMenu sets value of ShowInAttachMenu conditional field.
func (a *AttachMenuBot) SetShowInAttachMenu(value bool) {
	if value {
		a.Flags.Set(3)
		a.ShowInAttachMenu = true
	} else {
		a.Flags.Unset(3)
		a.ShowInAttachMenu = false
	}
}

// GetShowInAttachMenu returns value of ShowInAttachMenu conditional field.
func (a *AttachMenuBot) GetShowInAttachMenu() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(3)
}

// SetShowInSideMenu sets value of ShowInSideMenu conditional field.
func (a *AttachMenuBot) SetShowInSideMenu(value bool) {
	if value {
		a.Flags.Set(4)
		a.ShowInSideMenu = true
	} else {
		a.Flags.Unset(4)
		a.ShowInSideMenu = false
	}
}

// GetShowInSideMenu returns value of ShowInSideMenu conditional field.
func (a *AttachMenuBot) GetShowInSideMenu() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(4)
}

// SetSideMenuDisclaimerNeeded sets value of SideMenuDisclaimerNeeded conditional field.
func (a *AttachMenuBot) SetSideMenuDisclaimerNeeded(value bool) {
	if value {
		a.Flags.Set(5)
		a.SideMenuDisclaimerNeeded = true
	} else {
		a.Flags.Unset(5)
		a.SideMenuDisclaimerNeeded = false
	}
}

// GetSideMenuDisclaimerNeeded returns value of SideMenuDisclaimerNeeded conditional field.
func (a *AttachMenuBot) GetSideMenuDisclaimerNeeded() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(5)
}

// GetBotID returns value of BotID field.
func (a *AttachMenuBot) GetBotID() (value int64) {
	if a == nil {
		return
	}
	return a.BotID
}

// GetShortName returns value of ShortName field.
func (a *AttachMenuBot) GetShortName() (value string) {
	if a == nil {
		return
	}
	return a.ShortName
}

// SetPeerTypes sets value of PeerTypes conditional field.
func (a *AttachMenuBot) SetPeerTypes(value []AttachMenuPeerTypeClass) {
	a.Flags.Set(3)
	a.PeerTypes = value
}

// GetPeerTypes returns value of PeerTypes conditional field and
// boolean which is true if field was set.
func (a *AttachMenuBot) GetPeerTypes() (value []AttachMenuPeerTypeClass, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(3) {
		return value, false
	}
	return a.PeerTypes, true
}

// GetIcons returns value of Icons field.
func (a *AttachMenuBot) GetIcons() (value []AttachMenuBotIcon) {
	if a == nil {
		return
	}
	return a.Icons
}

// MapPeerTypes returns field PeerTypes wrapped in AttachMenuPeerTypeClassArray helper.
func (a *AttachMenuBot) MapPeerTypes() (value AttachMenuPeerTypeClassArray, ok bool) {
	if !a.Flags.Has(3) {
		return value, false
	}
	return AttachMenuPeerTypeClassArray(a.PeerTypes), true
}