package StreamPeerExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/StreamPeer"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

type Instance [1]classdb.StreamPeerExtension

func (Instance) _get_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes int, r_received *int32) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_received = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, int(r_bytes), r_received)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_partial_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes int, r_received *int32) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_received = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, int(r_bytes), r_received)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _put_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes int, r_sent *int32) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_sent = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, int(p_bytes), r_sent)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _put_partial_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes int, r_sent *int32) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_sent = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, int(p_bytes), r_sent)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_available_bytes(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StreamPeerExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeerExtension"))
	return Instance{classdb.StreamPeerExtension(object)}
}

func (class) _get_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes gd.Int, r_received *int32) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_received = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_bytes, r_received)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_partial_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes gd.Int, r_received *int32) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_received = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_bytes, r_received)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _put_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes gd.Int, r_sent *int32) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_sent = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, p_bytes, r_sent)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _put_partial_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes gd.Int, r_sent *int32) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_sent = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, p_bytes, r_sent)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_available_bytes(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsStreamPeerExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStreamPeerExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStreamPeer() StreamPeer.Advanced {
	return *((*StreamPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStreamPeer() StreamPeer.Instance {
	return *((*StreamPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_data":
		return reflect.ValueOf(self._get_data)
	case "_get_partial_data":
		return reflect.ValueOf(self._get_partial_data)
	case "_put_data":
		return reflect.ValueOf(self._put_data)
	case "_put_partial_data":
		return reflect.ValueOf(self._put_partial_data)
	case "_get_available_bytes":
		return reflect.ValueOf(self._get_available_bytes)
	default:
		return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_data":
		return reflect.ValueOf(self._get_data)
	case "_get_partial_data":
		return reflect.ValueOf(self._get_partial_data)
	case "_put_data":
		return reflect.ValueOf(self._put_data)
	case "_put_partial_data":
		return reflect.ValueOf(self._put_partial_data)
	case "_get_available_bytes":
		return reflect.ValueOf(self._get_available_bytes)
	default:
		return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}
func init() {
	classdb.Register("StreamPeerExtension", func(ptr gd.Object) any { return classdb.StreamPeerExtension(ptr) })
}
