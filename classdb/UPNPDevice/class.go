// Package UPNPDevice provides methods for working with UPNPDevice object instances.
package UPNPDevice

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Universal Plug and Play (UPnP) device. See [UPNP] for UPnP discovery and utility functions. Provides low-level access to UPNP control commands. Allows to manage port mappings (port forwarding) and to query network information of the device (like local and external IP address and status). Note that methods on this class are synchronous and block the calling thread.
*/
type Instance [1]gdclass.UPNPDevice

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsUPNPDevice() Instance
}

/*
Returns [code]true[/code] if this is a valid IGD (InternetGatewayDevice) which potentially supports port forwarding.
*/
func (self Instance) IsValidGateway() bool {
	return bool(class(self).IsValidGateway())
}

/*
Returns the external IP address of this [UPNPDevice] or an empty string.
*/
func (self Instance) QueryExternalAddress() string {
	return string(class(self).QueryExternalAddress().String())
}

/*
Adds a port mapping to forward the given external port on this [UPNPDevice] for the given protocol to the local machine. See [method UPNP.add_port_mapping].
*/
func (self Instance) AddPortMapping(port int) int {
	return int(int(class(self).AddPortMapping(gd.Int(port), gd.Int(0), gd.NewString(""), gd.NewString("UDP"), gd.Int(0))))
}

/*
Deletes the port mapping identified by the given port and protocol combination on this device. See [method UPNP.delete_port_mapping].
*/
func (self Instance) DeletePortMapping(port int) int {
	return int(int(class(self).DeletePortMapping(gd.Int(port), gd.NewString("UDP"))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.UPNPDevice

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("UPNPDevice"))
	casted := Instance{*(*gdclass.UPNPDevice)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) DescriptionUrl() string {
	return string(class(self).GetDescriptionUrl().String())
}

func (self Instance) SetDescriptionUrl(value string) {
	class(self).SetDescriptionUrl(gd.NewString(value))
}

func (self Instance) ServiceType() string {
	return string(class(self).GetServiceType().String())
}

func (self Instance) SetServiceType(value string) {
	class(self).SetServiceType(gd.NewString(value))
}

func (self Instance) IgdControlUrl() string {
	return string(class(self).GetIgdControlUrl().String())
}

func (self Instance) SetIgdControlUrl(value string) {
	class(self).SetIgdControlUrl(gd.NewString(value))
}

func (self Instance) IgdServiceType() string {
	return string(class(self).GetIgdServiceType().String())
}

func (self Instance) SetIgdServiceType(value string) {
	class(self).SetIgdServiceType(gd.NewString(value))
}

func (self Instance) IgdOurAddr() string {
	return string(class(self).GetIgdOurAddr().String())
}

func (self Instance) SetIgdOurAddr(value string) {
	class(self).SetIgdOurAddr(gd.NewString(value))
}

func (self Instance) IgdStatus() gdclass.UPNPDeviceIGDStatus {
	return gdclass.UPNPDeviceIGDStatus(class(self).GetIgdStatus())
}

func (self Instance) SetIgdStatus(value gdclass.UPNPDeviceIGDStatus) {
	class(self).SetIgdStatus(value)
}

/*
Returns [code]true[/code] if this is a valid IGD (InternetGatewayDevice) which potentially supports port forwarding.
*/
//go:nosplit
func (self class) IsValidGateway() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_is_valid_gateway, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_query_external_address, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
	callframe.Arg(frame, pointers.Get(desc))
	callframe.Arg(frame, pointers.Get(proto))
	callframe.Arg(frame, duration)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_add_port_mapping, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	callframe.Arg(frame, pointers.Get(proto))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_delete_port_mapping, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDescriptionUrl(url gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(url))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_description_url, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDescriptionUrl() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_description_url, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetServiceType(atype gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atype))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_service_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetServiceType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_service_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIgdControlUrl(url gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(url))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_igd_control_url, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIgdControlUrl() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_igd_control_url, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIgdServiceType(atype gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atype))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_igd_service_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIgdServiceType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_igd_service_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIgdOurAddr(addr gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(addr))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_igd_our_addr, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIgdOurAddr() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_igd_our_addr, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIgdStatus(status gdclass.UPNPDeviceIGDStatus) {
	var frame = callframe.New()
	callframe.Arg(frame, status)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_set_igd_status, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIgdStatus() gdclass.UPNPDeviceIGDStatus {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.UPNPDeviceIGDStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNPDevice.Bind_get_igd_status, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsUPNPDevice() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsUPNPDevice() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("UPNPDevice", func(ptr gd.Object) any { return [1]gdclass.UPNPDevice{*(*gdclass.UPNPDevice)(unsafe.Pointer(&ptr))} })
}

type IGDStatus = gdclass.UPNPDeviceIGDStatus

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
