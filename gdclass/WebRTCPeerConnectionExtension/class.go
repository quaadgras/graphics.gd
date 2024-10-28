package WebRTCPeerConnectionExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/WebRTCPeerConnection"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

type Go [1]classdb.WebRTCPeerConnectionExtension
func (Go) _get_connection_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionConnectionState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Go) _get_gathering_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionGatheringState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Go) _get_signaling_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionSignalingState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Go) _initialize(impl func(ptr unsafe.Pointer, p_config gd.Dictionary) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_config = discreet.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(p_config)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_config)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Go) _create_data_channel(impl func(ptr unsafe.Pointer, p_label string, p_config gd.Dictionary) gdclass.WebRTCDataChannel, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_label = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(p_label)
		var p_config = discreet.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(p_config)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label.String(), p_config)
ptr, ok := discreet.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Go) _create_offer(impl func(ptr unsafe.Pointer) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Go) _set_remote_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_type = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(p_type)
		var p_sdp = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(p_sdp)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Go) _set_local_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_type = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(p_type)
		var p_sdp = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(p_sdp)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Go) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name string, p_sdp_mline_index int, p_sdp_name string) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_sdp_mid_name = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(p_sdp_mid_name)
		var p_sdp_mline_index = gd.UnsafeGet[gd.Int](p_args,1)
		var p_sdp_name = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		defer discreet.End(p_sdp_name)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name.String(), int(p_sdp_mline_index), p_sdp_name.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Go) _poll(impl func(ptr unsafe.Pointer) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Go) _close(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.WebRTCPeerConnectionExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCPeerConnectionExtension"))
	return Go{classdb.WebRTCPeerConnectionExtension(object)}
}

func (class) _get_connection_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionConnectionState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_gathering_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionGatheringState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_signaling_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionSignalingState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _initialize(impl func(ptr unsafe.Pointer, p_config gd.Dictionary) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_config = discreet.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_config)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _create_data_channel(impl func(ptr unsafe.Pointer, p_label gd.String, p_config gd.Dictionary) gdclass.WebRTCDataChannel, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_label = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		var p_config = discreet.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label, p_config)
ptr, ok := discreet.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _create_offer(impl func(ptr unsafe.Pointer) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_remote_description(impl func(ptr unsafe.Pointer, p_type gd.String, p_sdp gd.String) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_type = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		var p_sdp = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_local_description(impl func(ptr unsafe.Pointer, p_type gd.String, p_sdp gd.String) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_type = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		var p_sdp = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name gd.String, p_sdp_mline_index gd.Int, p_sdp_name gd.String) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_sdp_mid_name = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		var p_sdp_mline_index = gd.UnsafeGet[gd.Int](p_args,1)
		var p_sdp_name = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name, p_sdp_mline_index, p_sdp_name)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _poll(impl func(ptr unsafe.Pointer) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _close(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
	}
}

func (self class) AsWebRTCPeerConnectionExtension() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsWebRTCPeerConnectionExtension() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsWebRTCPeerConnection() WebRTCPeerConnection.GD { return *((*WebRTCPeerConnection.GD)(unsafe.Pointer(&self))) }
func (self Go) AsWebRTCPeerConnection() WebRTCPeerConnection.Go { return *((*WebRTCPeerConnection.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_connection_state": return reflect.ValueOf(self._get_connection_state);
	case "_get_gathering_state": return reflect.ValueOf(self._get_gathering_state);
	case "_get_signaling_state": return reflect.ValueOf(self._get_signaling_state);
	case "_initialize": return reflect.ValueOf(self._initialize);
	case "_create_data_channel": return reflect.ValueOf(self._create_data_channel);
	case "_create_offer": return reflect.ValueOf(self._create_offer);
	case "_set_remote_description": return reflect.ValueOf(self._set_remote_description);
	case "_set_local_description": return reflect.ValueOf(self._set_local_description);
	case "_add_ice_candidate": return reflect.ValueOf(self._add_ice_candidate);
	case "_poll": return reflect.ValueOf(self._poll);
	case "_close": return reflect.ValueOf(self._close);
	default: return gd.VirtualByName(self.AsWebRTCPeerConnection(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_connection_state": return reflect.ValueOf(self._get_connection_state);
	case "_get_gathering_state": return reflect.ValueOf(self._get_gathering_state);
	case "_get_signaling_state": return reflect.ValueOf(self._get_signaling_state);
	case "_initialize": return reflect.ValueOf(self._initialize);
	case "_create_data_channel": return reflect.ValueOf(self._create_data_channel);
	case "_create_offer": return reflect.ValueOf(self._create_offer);
	case "_set_remote_description": return reflect.ValueOf(self._set_remote_description);
	case "_set_local_description": return reflect.ValueOf(self._set_local_description);
	case "_add_ice_candidate": return reflect.ValueOf(self._add_ice_candidate);
	case "_poll": return reflect.ValueOf(self._poll);
	case "_close": return reflect.ValueOf(self._close);
	default: return gd.VirtualByName(self.AsWebRTCPeerConnection(), name)
	}
}
func init() {classdb.Register("WebRTCPeerConnectionExtension", func(ptr gd.Object) any { return classdb.WebRTCPeerConnectionExtension(ptr) })}
