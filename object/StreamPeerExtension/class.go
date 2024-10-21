package StreamPeerExtension

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/StreamPeer"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Simple [1]classdb.StreamPeerExtension
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.StreamPeerExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func (class) _get_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes gd.Int, r_received *int32) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args,1)
		var r_received = gd.UnsafeGet[*int32](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_bytes, r_received)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_partial_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes gd.Int, r_received *int32) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args,1)
		var r_received = gd.UnsafeGet[*int32](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_bytes, r_received)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _put_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes gd.Int, r_sent *int32) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args,1)
		var r_sent = gd.UnsafeGet[*int32](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, p_bytes, r_sent)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _put_partial_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes gd.Int, r_sent *int32) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args,1)
		var r_sent = gd.UnsafeGet[*int32](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, p_bytes, r_sent)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_available_bytes(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}


//go:nosplit
func (self class) AsStreamPeerExtension() Expert { return self[0].AsStreamPeerExtension() }


//go:nosplit
func (self Simple) AsStreamPeerExtension() Simple { return self[0].AsStreamPeerExtension() }


//go:nosplit
func (self class) AsStreamPeer() StreamPeer.Expert { return self[0].AsStreamPeer() }


//go:nosplit
func (self Simple) AsStreamPeer() StreamPeer.Simple { return self[0].AsStreamPeer() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_data": return reflect.ValueOf(self._get_data);
	case "_get_partial_data": return reflect.ValueOf(self._get_partial_data);
	case "_put_data": return reflect.ValueOf(self._put_data);
	case "_put_partial_data": return reflect.ValueOf(self._put_partial_data);
	case "_get_available_bytes": return reflect.ValueOf(self._get_available_bytes);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("StreamPeerExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
