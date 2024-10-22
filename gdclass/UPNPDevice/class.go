package UPNPDevice

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Universal Plug and Play (UPnP) device. See [UPNP] for UPnP discovery and utility functions. Provides low-level access to UPNP control commands. Allows to manage port mappings (port forwarding) and to query network information of the device (like local and external IP address and status). Note that methods on this class are synchronous and block the calling thread.

*/
type Go [1]classdb.UPNPDevice

/*
Returns [code]true[/code] if this is a valid IGD (InternetGatewayDevice) which potentially supports port forwarding.
*/
func (self Go) IsValidGateway() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsValidGateway())
}

/*
Returns the external IP address of this [UPNPDevice] or an empty string.
*/
func (self Go) QueryExternalAddress() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).QueryExternalAddress(gc).String())
}

/*
Adds a port mapping to forward the given external port on this [UPNPDevice] for the given protocol to the local machine. See [method UPNP.add_port_mapping].
*/
func (self Go) AddPortMapping(port int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).AddPortMapping(gd.Int(port), gd.Int(0), gc.String(""), gc.String("UDP"), gd.Int(0))))
}

/*
Deletes the port mapping identified by the given port and protocol combination on this device. See [method UPNP.delete_port_mapping].
*/
func (self Go) DeletePortMapping(port int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).DeletePortMapping(gd.Int(port), gc.String("UDP"))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.UPNPDevice
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("UPNPDevice"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) DescriptionUrl() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetDescriptionUrl(gc).String())
}

func (self Go) SetDescriptionUrl(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDescriptionUrl(gc.String(value))
}

func (self Go) ServiceType() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetServiceType(gc).String())
}

func (self Go) SetServiceType(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetServiceType(gc.String(value))
}

func (self Go) IgdControlUrl() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetIgdControlUrl(gc).String())
}

func (self Go) SetIgdControlUrl(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIgdControlUrl(gc.String(value))
}

func (self Go) IgdServiceType() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetIgdServiceType(gc).String())
}

func (self Go) SetIgdServiceType(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIgdServiceType(gc.String(value))
}

func (self Go) IgdOurAddr() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetIgdOurAddr(gc).String())
}

func (self Go) SetIgdOurAddr(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIgdOurAddr(gc.String(value))
}

func (self Go) IgdStatus() classdb.UPNPDeviceIGDStatus {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.UPNPDeviceIGDStatus(class(self).GetIgdStatus())
}

func (self Go) SetIgdStatus(value classdb.UPNPDeviceIGDStatus) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIgdStatus(value)
}

/*
Returns [code]true[/code] if this is a valid IGD (InternetGatewayDevice) which potentially supports port forwarding.
*/
//go:nosplit
func (self class) IsValidGateway() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_is_valid_gateway, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the external IP address of this [UPNPDevice] or an empty string.
*/
//go:nosplit
func (self class) QueryExternalAddress(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_query_external_address, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds a port mapping to forward the given external port on this [UPNPDevice] for the given protocol to the local machine. See [method UPNP.add_port_mapping].
*/
//go:nosplit
func (self class) AddPortMapping(port gd.Int, port_internal gd.Int, desc gd.String, proto gd.String, duration gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, port_internal)
	callframe.Arg(frame, mmm.Get(desc))
	callframe.Arg(frame, mmm.Get(proto))
	callframe.Arg(frame, duration)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_add_port_mapping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Deletes the port mapping identified by the given port and protocol combination on this device. See [method UPNP.delete_port_mapping].
*/
//go:nosplit
func (self class) DeletePortMapping(port gd.Int, proto gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, mmm.Get(proto))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_delete_port_mapping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDescriptionUrl(url gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(url))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_set_description_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDescriptionUrl(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_get_description_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetServiceType(atype gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(atype))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_set_service_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetServiceType(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_get_service_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgdControlUrl(url gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(url))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_set_igd_control_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIgdControlUrl(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_get_igd_control_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgdServiceType(atype gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(atype))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_set_igd_service_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIgdServiceType(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_get_igd_service_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgdOurAddr(addr gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(addr))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_set_igd_our_addr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIgdOurAddr(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_get_igd_our_addr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgdStatus(status classdb.UPNPDeviceIGDStatus)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, status)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_set_igd_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIgdStatus() classdb.UPNPDeviceIGDStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.UPNPDeviceIGDStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UPNPDevice.Bind_get_igd_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsUPNPDevice() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsUPNPDevice() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("UPNPDevice", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type IGDStatus = classdb.UPNPDeviceIGDStatus

const (
/*OK.*/
	IgdStatusOk IGDStatus = 0
/*HTTP error.*/
	IgdStatusHttpError IGDStatus = 1
/*Empty HTTP response.*/
	IgdStatusHttpEmpty IGDStatus = 2
/*Returned response contained no URLs.*/
	IgdStatusNoUrls IGDStatus = 3
/*Not a valid IGD.*/
	IgdStatusNoIgd IGDStatus = 4
/*Disconnected.*/
	IgdStatusDisconnected IGDStatus = 5
/*Unknown device.*/
	IgdStatusUnknownDevice IGDStatus = 6
/*Invalid control.*/
	IgdStatusInvalidControl IGDStatus = 7
/*Memory allocation error.*/
	IgdStatusMallocError IGDStatus = 8
/*Unknown error.*/
	IgdStatusUnknownError IGDStatus = 9
)
