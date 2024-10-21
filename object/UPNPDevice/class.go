package UPNPDevice

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Universal Plug and Play (UPnP) device. See [UPNP] for UPnP discovery and utility functions. Provides low-level access to UPNP control commands. Allows to manage port mappings (port forwarding) and to query network information of the device (like local and external IP address and status). Note that methods on this class are synchronous and block the calling thread.

*/
type Simple [1]classdb.UPNPDevice
func (self Simple) IsValidGateway() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsValidGateway())
}
func (self Simple) QueryExternalAddress() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).QueryExternalAddress(gc).String())
}
func (self Simple) AddPortMapping(port int, port_internal int, desc string, proto string, duration int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddPortMapping(gd.Int(port), gd.Int(port_internal), gc.String(desc), gc.String(proto), gd.Int(duration))))
}
func (self Simple) DeletePortMapping(port int, proto string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).DeletePortMapping(gd.Int(port), gc.String(proto))))
}
func (self Simple) SetDescriptionUrl(url string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDescriptionUrl(gc.String(url))
}
func (self Simple) GetDescriptionUrl() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetDescriptionUrl(gc).String())
}
func (self Simple) SetServiceType(atype string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetServiceType(gc.String(atype))
}
func (self Simple) GetServiceType() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetServiceType(gc).String())
}
func (self Simple) SetIgdControlUrl(url string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIgdControlUrl(gc.String(url))
}
func (self Simple) GetIgdControlUrl() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetIgdControlUrl(gc).String())
}
func (self Simple) SetIgdServiceType(atype string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIgdServiceType(gc.String(atype))
}
func (self Simple) GetIgdServiceType() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetIgdServiceType(gc).String())
}
func (self Simple) SetIgdOurAddr(addr string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIgdOurAddr(gc.String(addr))
}
func (self Simple) GetIgdOurAddr() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetIgdOurAddr(gc).String())
}
func (self Simple) SetIgdStatus(status classdb.UPNPDeviceIGDStatus) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIgdStatus(status)
}
func (self Simple) GetIgdStatus() classdb.UPNPDeviceIGDStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.UPNPDeviceIGDStatus(Expert(self).GetIgdStatus())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.UPNPDevice
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsUPNPDevice() Expert { return self[0].AsUPNPDevice() }


//go:nosplit
func (self Simple) AsUPNPDevice() Simple { return self[0].AsUPNPDevice() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("UPNPDevice", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
