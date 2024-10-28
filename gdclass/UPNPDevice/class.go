package UPNPDevice

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Universal Plug and Play (UPnP) device. See [UPNP] for UPnP discovery and utility functions. Provides low-level access to UPNP control commands. Allows to manage port mappings (port forwarding) and to query network information of the device (like local and external IP address and status). Note that methods on this class are synchronous and block the calling thread.

*/
type Go [1]classdb.UPNPDevice

/*
Returns [code]true[/code] if this is a valid IGD (InternetGatewayDevice) which potentially supports port forwarding.
*/
func (self Go) IsValidGateway() bool {
	return bool(class(self).IsValidGateway())
}

/*
Returns the external IP address of this [UPNPDevice] or an empty string.
*/
func (self Go) QueryExternalAddress() string {
	return string(class(self).QueryExternalAddress().String())
}

/*
Adds a port mapping to forward the given external port on this [UPNPDevice] for the given protocol to the local machine. See [method UPNP.add_port_mapping].
*/
func (self Go) AddPortMapping(port int) int {
	return int(int(class(self).AddPortMapping(gd.Int(port), gd.Int(0), gd.NewString(""), gd.NewString("UDP"), gd.Int(0))))
}

/*
Deletes the port mapping identified by the given port and protocol combination on this device. See [method UPNP.delete_port_mapping].
*/
func (self Go) DeletePortMapping(port int) int {
	return int(int(class(self).DeletePortMapping(gd.Int(port), gd.NewString("UDP"))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.UPNPDevice
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("UPNPDevice"))
	return Go{classdb.UPNPDevice(object)}
}

func (self Go) DescriptionUrl() string {
		return string(class(self).GetDescriptionUrl().String())
}

func (self Go) SetDescriptionUrl(value string) {
	class(self).SetDescriptionUrl(gd.NewString(value))
}

func (self Go) ServiceType() string {
		return string(class(self).GetServiceType().String())
}

func (self Go) SetServiceType(value string) {
	class(self).SetServiceType(gd.NewString(value))
}

func (self Go) IgdControlUrl() string {
		return string(class(self).GetIgdControlUrl().String())
}

func (self Go) SetIgdControlUrl(value string) {
	class(self).SetIgdControlUrl(gd.NewString(value))
}

func (self Go) IgdServiceType() string {
		return string(class(self).GetIgdServiceType().String())
}

func (self Go) SetIgdServiceType(value string) {
	class(self).SetIgdServiceType(gd.NewString(value))
}

func (self Go) IgdOurAddr() string {
		return string(class(self).GetIgdOurAddr().String())
}

func (self Go) SetIgdOurAddr(value string) {
	class(self).SetIgdOurAddr(gd.NewString(value))
}

func (self Go) IgdStatus() classdb.UPNPDeviceIGDStatus {
		return classdb.UPNPDeviceIGDStatus(class(self).GetIgdStatus())
}

func (self Go) SetIgdStatus(value classdb.UPNPDeviceIGDStatus) {
	class(self).SetIgdStatus(value)
}

/*
Returns [code]true[/code] if this is a valid IGD (InternetGatewayDevice) which potentially supports port forwarding.
*/
//go:nosplit
func (self class) IsValidGateway() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_is_valid_gateway, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the external IP address of this [UPNPDevice] or an empty string.
*/
//go:nosplit
func (self class) QueryExternalAddress() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_query_external_address, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds a port mapping to forward the given external port on this [UPNPDevice] for the given protocol to the local machine. See [method UPNP.add_port_mapping].
*/
//go:nosplit
func (self class) AddPortMapping(port gd.Int, port_internal gd.Int, desc gd.String, proto gd.String, duration gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, port_internal)
	callframe.Arg(frame, discreet.Get(desc))
	callframe.Arg(frame, discreet.Get(proto))
	callframe.Arg(frame, duration)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_add_port_mapping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Deletes the port mapping identified by the given port and protocol combination on this device. See [method UPNP.delete_port_mapping].
*/
//go:nosplit
func (self class) DeletePortMapping(port gd.Int, proto gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, discreet.Get(proto))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_delete_port_mapping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDescriptionUrl(url gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(url))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_description_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDescriptionUrl() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_description_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetServiceType(atype gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(atype))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_service_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetServiceType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_service_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgdControlUrl(url gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(url))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_igd_control_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIgdControlUrl() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_igd_control_url, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgdServiceType(atype gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(atype))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_igd_service_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIgdServiceType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_igd_service_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgdOurAddr(addr gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(addr))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_igd_our_addr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIgdOurAddr() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_igd_our_addr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgdStatus(status classdb.UPNPDeviceIGDStatus)  {
	var frame = callframe.New()
	callframe.Arg(frame, status)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_igd_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIgdStatus() classdb.UPNPDeviceIGDStatus {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.UPNPDeviceIGDStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_igd_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func init() {classdb.Register("UPNPDevice", func(ptr gd.Object) any { return classdb.UPNPDevice(ptr) })}
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
