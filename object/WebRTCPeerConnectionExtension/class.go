package WebRTCPeerConnectionExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/WebRTCPeerConnection"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Simple [1]classdb.WebRTCPeerConnectionExtension
func (Simple) _get_connection_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionConnectionState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_gathering_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionGatheringState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_signaling_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionSignalingState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _initialize(impl func(ptr unsafe.Pointer, p_config gd.Dictionary) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var p_config = mmm.Let[gd.Dictionary](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_config)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _create_data_channel(impl func(ptr unsafe.Pointer, p_label string, p_config gd.Dictionary) [1]classdb.WebRTCDataChannel, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var p_label = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var p_config = mmm.Let[gd.Dictionary](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label.String(), p_config)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _create_offer(impl func(ptr unsafe.Pointer) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _set_remote_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var p_type = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var p_sdp = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _set_local_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var p_type = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var p_sdp = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name string, p_sdp_mline_index int, p_sdp_name string) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var p_sdp_mid_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var p_sdp_mline_index = gd.UnsafeGet[gd.Int](p_args,1)
		var p_sdp_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name.String(), int(p_sdp_mline_index), p_sdp_name.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _poll(impl func(ptr unsafe.Pointer) gd.Error, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _close(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.WebRTCPeerConnectionExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func (class) _get_connection_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionConnectionState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_gathering_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionGatheringState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_signaling_state(impl func(ptr unsafe.Pointer) classdb.WebRTCPeerConnectionSignalingState, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _initialize(impl func(ptr unsafe.Pointer, p_config gd.Dictionary) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var p_config = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_config)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _create_data_channel(impl func(ptr unsafe.Pointer, p_label gd.String, p_config gd.Dictionary) object.WebRTCDataChannel, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var p_label = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var p_config = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label, p_config)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

func (class) _create_offer(impl func(ptr unsafe.Pointer) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_remote_description(impl func(ptr unsafe.Pointer, p_type gd.String, p_sdp gd.String) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var p_type = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var p_sdp = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_local_description(impl func(ptr unsafe.Pointer, p_type gd.String, p_sdp gd.String) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var p_type = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var p_sdp = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name gd.String, p_sdp_mline_index gd.Int, p_sdp_name gd.String) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var p_sdp_mid_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var p_sdp_mline_index = gd.UnsafeGet[gd.Int](p_args,1)
		var p_sdp_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name, p_sdp_mline_index, p_sdp_name)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _poll(impl func(ptr unsafe.Pointer) int64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _close(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}


//go:nosplit
func (self class) AsWebRTCPeerConnectionExtension() Expert { return self[0].AsWebRTCPeerConnectionExtension() }


//go:nosplit
func (self Simple) AsWebRTCPeerConnectionExtension() Simple { return self[0].AsWebRTCPeerConnectionExtension() }


//go:nosplit
func (self class) AsWebRTCPeerConnection() WebRTCPeerConnection.Expert { return self[0].AsWebRTCPeerConnection() }


//go:nosplit
func (self Simple) AsWebRTCPeerConnection() WebRTCPeerConnection.Simple { return self[0].AsWebRTCPeerConnection() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("WebRTCPeerConnectionExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
