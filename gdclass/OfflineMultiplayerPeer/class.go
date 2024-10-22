package OfflineMultiplayerPeer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/MultiplayerPeer"
import "grow.graphics/gd/gdclass/PacketPeer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This is the default [member MultiplayerAPI.multiplayer_peer] for the [member Node.multiplayer]. It mimics the behavior of a server with no peers connected.
This means that the [SceneTree] will act as the multiplayer authority by default. Calls to [method MultiplayerAPI.is_server] will return [code]true[/code], and calls to [method MultiplayerAPI.get_unique_id] will return [constant MultiplayerPeer.TARGET_PEER_SERVER].

*/
type Go [1]classdb.OfflineMultiplayerPeer
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OfflineMultiplayerPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OfflineMultiplayerPeer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsOfflineMultiplayerPeer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOfflineMultiplayerPeer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMultiplayerPeer() MultiplayerPeer.GD { return *((*MultiplayerPeer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMultiplayerPeer() MultiplayerPeer.Go { return *((*MultiplayerPeer.Go)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.GD { return *((*PacketPeer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPacketPeer() PacketPeer.Go { return *((*PacketPeer.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMultiplayerPeer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMultiplayerPeer(), name)
	}
}
func init() {classdb.Register("OfflineMultiplayerPeer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
