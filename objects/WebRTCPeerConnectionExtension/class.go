package WebRTCPeerConnectionExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/WebRTCPeerConnection"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

type Instance [1]classdb.WebRTCPeerConnectionExtension

func (Instance) _get_connection_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionConnectionState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_gathering_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionGatheringState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_signaling_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionSignalingState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _initialize(impl func(ptr unsafe.Pointer, p_config gd.Dictionary) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_config = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(p_config)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_config)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _create_data_channel(impl func(ptr unsafe.Pointer, p_label string, p_config gd.Dictionary) objects.WebRTCDataChannel) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_label = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(p_label)
		var p_config = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(p_config)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label.String(), p_config)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _create_offer(impl func(ptr unsafe.Pointer) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _set_remote_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_type = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(p_type)
		var p_sdp = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(p_sdp)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _set_local_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_type = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(p_type)
		var p_sdp = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(p_sdp)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name string, p_sdp_mline_index int, p_sdp_name string) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_sdp_mid_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(p_sdp_mid_name)
		var p_sdp_mline_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var p_sdp_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(p_sdp_name)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name.String(), int(p_sdp_mline_index), p_sdp_name.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _poll(impl func(ptr unsafe.Pointer) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.WebRTCPeerConnectionExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCPeerConnectionExtension"))
	return Instance{classdb.WebRTCPeerConnectionExtension(object)}
}

func (class) _get_connection_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionConnectionState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_gathering_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionGatheringState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_signaling_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionSignalingState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _initialize(impl func(ptr unsafe.Pointer, p_config gd.Dictionary) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_config = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_config)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _create_data_channel(impl func(ptr unsafe.Pointer, p_label gd.String, p_config gd.Dictionary) objects.WebRTCDataChannel) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_label = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var p_config = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label, p_config)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _create_offer(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_remote_description(impl func(ptr unsafe.Pointer, p_type gd.String, p_sdp gd.String) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_type = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var p_sdp = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_local_description(impl func(ptr unsafe.Pointer, p_type gd.String, p_sdp gd.String) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_type = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var p_sdp = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name gd.String, p_sdp_mline_index gd.Int, p_sdp_name gd.String) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_sdp_mid_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var p_sdp_mline_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var p_sdp_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name, p_sdp_mline_index, p_sdp_name)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _poll(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (self class) AsWebRTCPeerConnectionExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsWebRTCPeerConnectionExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsWebRTCPeerConnection() WebRTCPeerConnection.Advanced {
	return *((*WebRTCPeerConnection.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsWebRTCPeerConnection() WebRTCPeerConnection.Instance {
	return *((*WebRTCPeerConnection.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_connection_state":
		return reflect.ValueOf(self._get_connection_state)
	case "_get_gathering_state":
		return reflect.ValueOf(self._get_gathering_state)
	case "_get_signaling_state":
		return reflect.ValueOf(self._get_signaling_state)
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_create_data_channel":
		return reflect.ValueOf(self._create_data_channel)
	case "_create_offer":
		return reflect.ValueOf(self._create_offer)
	case "_set_remote_description":
		return reflect.ValueOf(self._set_remote_description)
	case "_set_local_description":
		return reflect.ValueOf(self._set_local_description)
	case "_add_ice_candidate":
		return reflect.ValueOf(self._add_ice_candidate)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	default:
		return gd.VirtualByName(self.AsWebRTCPeerConnection(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_connection_state":
		return reflect.ValueOf(self._get_connection_state)
	case "_get_gathering_state":
		return reflect.ValueOf(self._get_gathering_state)
	case "_get_signaling_state":
		return reflect.ValueOf(self._get_signaling_state)
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_create_data_channel":
		return reflect.ValueOf(self._create_data_channel)
	case "_create_offer":
		return reflect.ValueOf(self._create_offer)
	case "_set_remote_description":
		return reflect.ValueOf(self._set_remote_description)
	case "_set_local_description":
		return reflect.ValueOf(self._set_local_description)
	case "_add_ice_candidate":
		return reflect.ValueOf(self._add_ice_candidate)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	default:
		return gd.VirtualByName(self.AsWebRTCPeerConnection(), name)
	}
}
func init() {
	classdb.Register("WebRTCPeerConnectionExtension", func(ptr gd.Object) any { return classdb.WebRTCPeerConnectionExtension(ptr) })
}
