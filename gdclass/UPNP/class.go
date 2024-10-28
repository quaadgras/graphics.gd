package UPNP

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
This class can be used to discover compatible [UPNPDevice]s on the local network and execute commands on them, like managing port mappings (for port forwarding/NAT traversal) and querying the local and remote network IP address. Note that methods on this class are synchronous and block the calling thread.
To forward a specific port (here [code]7777[/code], note both [method discover] and [method add_port_mapping] can return errors that should be checked):
[codeblock]
var upnp = UPNP.new()
upnp.discover()
upnp.add_port_mapping(7777)
[/codeblock]
To close a specific port (e.g. after you have finished using it):
[codeblock]
upnp.delete_port_mapping(port)
[/codeblock]
[b]Note:[/b] UPnP discovery blocks the current thread. To perform discovery without blocking the main thread, use [Thread]s like this:
[codeblock]
# Emitted when UPnP port mapping setup is completed (regardless of success or failure).
signal upnp_completed(error)

# Replace this with your own server port number between 1024 and 65535.
const SERVER_PORT = 3928
var thread = null

func _upnp_setup(server_port):
    # UPNP queries take some time.
    var upnp = UPNP.new()
    var err = upnp.discover()

    if err != OK:
        push_error(str(err))
        emit_signal("upnp_completed", err)
        return

    if upnp.get_gateway() and upnp.get_gateway().is_valid_gateway():
        upnp.add_port_mapping(server_port, server_port, ProjectSettings.get_setting("application/config/name"), "UDP")
        upnp.add_port_mapping(server_port, server_port, ProjectSettings.get_setting("application/config/name"), "TCP")
        emit_signal("upnp_completed", OK)

func _ready():
    thread = Thread.new()
    thread.start(_upnp_setup.bind(SERVER_PORT))

func _exit_tree():
    # Wait for thread finish here to handle game exit while the thread is running.
    thread.wait_to_finish()
[/codeblock]
[b]Terminology:[/b] In the context of UPnP networking, "gateway" (or "internet gateway device", short IGD) refers to network devices that allow computers in the local network to access the internet ("wide area network", WAN). These gateways are often also called "routers".
[b]Pitfalls:[/b]
- As explained above, these calls are blocking and shouldn't be run on the main thread, especially as they can block for multiple seconds at a time. Use threading!
- Networking is physical and messy. Packets get lost in transit or get filtered, addresses, free ports and assigned mappings change, and devices may leave or join the network at any time. Be mindful of this, be diligent when checking and handling errors, and handle these gracefully if you can: add clear error UI, timeouts and re-try handling.
- Port mappings may change (and be removed) at any time, and the remote/external IP address of the gateway can change likewise. You should consider re-querying the external IP and try to update/refresh the port mapping periodically (for example, every 5 minutes and on networking failures).
- Not all devices support UPnP, and some users disable UPnP support. You need to handle this (e.g. documenting and requiring the user to manually forward ports, or adding alternative methods of NAT traversal, like a relay/mirror server, or NAT hole punching, STUN/TURN, etc.).
- Consider what happens on mapping conflicts. Maybe multiple users on the same network would like to play your game at the same time, or maybe another application uses the same port. Make the port configurable, and optimally choose a port automatically (re-trying with a different port on failure).
[b]Further reading:[/b] If you want to know more about UPnP (and the Internet Gateway Device (IGD) and Port Control Protocol (PCP) specifically), [url=https://en.wikipedia.org/wiki/Universal_Plug_and_Play]Wikipedia[/url] is a good first stop, the specification can be found at the [url=https://openconnectivity.org/developer/specifications/upnp-resources/upnp/]Open Connectivity Foundation[/url] and Godot's implementation is based on the [url=https://github.com/miniupnp/miniupnp]MiniUPnP client[/url].

*/
type Go [1]classdb.UPNP

/*
Returns the number of discovered [UPNPDevice]s.
*/
func (self Go) GetDeviceCount() int {
	return int(int(class(self).GetDeviceCount()))
}

/*
Returns the [UPNPDevice] at the given [param index].
*/
func (self Go) GetDevice(index int) gdclass.UPNPDevice {
	return gdclass.UPNPDevice(class(self).GetDevice(gd.Int(index)))
}

/*
Adds the given [UPNPDevice] to the list of discovered devices.
*/
func (self Go) AddDevice(device gdclass.UPNPDevice) {
	class(self).AddDevice(device)
}

/*
Sets the device at [param index] from the list of discovered devices to [param device].
*/
func (self Go) SetDevice(index int, device gdclass.UPNPDevice) {
	class(self).SetDevice(gd.Int(index), device)
}

/*
Removes the device at [param index] from the list of discovered devices.
*/
func (self Go) RemoveDevice(index int) {
	class(self).RemoveDevice(gd.Int(index))
}

/*
Clears the list of discovered devices.
*/
func (self Go) ClearDevices() {
	class(self).ClearDevices()
}

/*
Returns the default gateway. That is the first discovered [UPNPDevice] that is also a valid IGD (InternetGatewayDevice).
*/
func (self Go) GetGateway() gdclass.UPNPDevice {
	return gdclass.UPNPDevice(class(self).GetGateway())
}

/*
Discovers local [UPNPDevice]s. Clears the list of previously discovered devices.
Filters for IGD (InternetGatewayDevice) type devices by default, as those manage port forwarding. [param timeout] is the time to wait for responses in milliseconds. [param ttl] is the time-to-live; only touch this if you know what you're doing.
See [enum UPNPResult] for possible return values.
*/
func (self Go) Discover() int {
	return int(int(class(self).Discover(gd.Int(2000), gd.Int(2), gd.NewString("InternetGatewayDevice"))))
}

/*
Returns the external [IP] address of the default gateway (see [method get_gateway]) as string. Returns an empty string on error.
*/
func (self Go) QueryExternalAddress() string {
	return string(class(self).QueryExternalAddress().String())
}

/*
Adds a mapping to forward the external [param port] (between 1 and 65535, although recommended to use port 1024 or above) on the default gateway (see [method get_gateway]) to the [param port_internal] on the local machine for the given protocol [param proto] (either [code]"TCP"[/code] or [code]"UDP"[/code], with UDP being the default). If a port mapping for the given port and protocol combination already exists on that gateway device, this method tries to overwrite it. If that is not desired, you can retrieve the gateway manually with [method get_gateway] and call [method add_port_mapping] on it, if any. Note that forwarding a well-known port (below 1024) with UPnP may fail depending on the device.
Depending on the gateway device, if a mapping for that port already exists, it will either be updated or it will refuse this command due to that conflict, especially if the existing mapping for that port wasn't created via UPnP or points to a different network address (or device) than this one.
If [param port_internal] is [code]0[/code] (the default), the same port number is used for both the external and the internal port (the [param port] value).
The description ([param desc]) is shown in some routers management UIs and can be used to point out which application added the mapping.
The mapping's lease [param duration] can be limited by specifying a duration in seconds. The default of [code]0[/code] means no duration, i.e. a permanent lease and notably some devices only support these permanent leases. Note that whether permanent or not, this is only a request and the gateway may still decide at any point to remove the mapping (which usually happens on a reboot of the gateway, when its external IP address changes, or on some models when it detects a port mapping has become inactive, i.e. had no traffic for multiple minutes). If not [code]0[/code] (permanent), the allowed range according to spec is between [code]120[/code] (2 minutes) and [code]86400[/code] seconds (24 hours).
See [enum UPNPResult] for possible return values.
*/
func (self Go) AddPortMapping(port int) int {
	return int(int(class(self).AddPortMapping(gd.Int(port), gd.Int(0), gd.NewString(""), gd.NewString("UDP"), gd.Int(0))))
}

/*
Deletes the port mapping for the given port and protocol combination on the default gateway (see [method get_gateway]) if one exists. [param port] must be a valid port between 1 and 65535, [param proto] can be either [code]"TCP"[/code] or [code]"UDP"[/code]. May be refused for mappings pointing to addresses other than this one, for well-known ports (below 1024), or for mappings not added via UPnP. See [enum UPNPResult] for possible return values.
*/
func (self Go) DeletePortMapping(port int) int {
	return int(int(class(self).DeletePortMapping(gd.Int(port), gd.NewString("UDP"))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.UPNP
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("UPNP"))
	return Go{classdb.UPNP(object)}
}

func (self Go) DiscoverMulticastIf() string {
		return string(class(self).GetDiscoverMulticastIf().String())
}

func (self Go) SetDiscoverMulticastIf(value string) {
	class(self).SetDiscoverMulticastIf(gd.NewString(value))
}

func (self Go) DiscoverLocalPort() int {
		return int(int(class(self).GetDiscoverLocalPort()))
}

func (self Go) SetDiscoverLocalPort(value int) {
	class(self).SetDiscoverLocalPort(gd.Int(value))
}

func (self Go) DiscoverIpv6() bool {
		return bool(class(self).IsDiscoverIpv6())
}

func (self Go) SetDiscoverIpv6(value bool) {
	class(self).SetDiscoverIpv6(value)
}

/*
Returns the number of discovered [UPNPDevice]s.
*/
//go:nosplit
func (self class) GetDeviceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_get_device_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [UPNPDevice] at the given [param index].
*/
//go:nosplit
func (self class) GetDevice(index gd.Int) gdclass.UPNPDevice {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_get_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.UPNPDevice{classdb.UPNPDevice(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Adds the given [UPNPDevice] to the list of discovered devices.
*/
//go:nosplit
func (self class) AddDevice(device gdclass.UPNPDevice)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(device[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_add_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the device at [param index] from the list of discovered devices to [param device].
*/
//go:nosplit
func (self class) SetDevice(index gd.Int, device gdclass.UPNPDevice)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, discreet.Get(device[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_set_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the device at [param index] from the list of discovered devices.
*/
//go:nosplit
func (self class) RemoveDevice(index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_remove_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the list of discovered devices.
*/
//go:nosplit
func (self class) ClearDevices()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_clear_devices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the default gateway. That is the first discovered [UPNPDevice] that is also a valid IGD (InternetGatewayDevice).
*/
//go:nosplit
func (self class) GetGateway() gdclass.UPNPDevice {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_get_gateway, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.UPNPDevice{classdb.UPNPDevice(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Discovers local [UPNPDevice]s. Clears the list of previously discovered devices.
Filters for IGD (InternetGatewayDevice) type devices by default, as those manage port forwarding. [param timeout] is the time to wait for responses in milliseconds. [param ttl] is the time-to-live; only touch this if you know what you're doing.
See [enum UPNPResult] for possible return values.
*/
//go:nosplit
func (self class) Discover(timeout gd.Int, ttl gd.Int, device_filter gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	callframe.Arg(frame, ttl)
	callframe.Arg(frame, discreet.Get(device_filter))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_discover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the external [IP] address of the default gateway (see [method get_gateway]) as string. Returns an empty string on error.
*/
//go:nosplit
func (self class) QueryExternalAddress() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_query_external_address, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds a mapping to forward the external [param port] (between 1 and 65535, although recommended to use port 1024 or above) on the default gateway (see [method get_gateway]) to the [param port_internal] on the local machine for the given protocol [param proto] (either [code]"TCP"[/code] or [code]"UDP"[/code], with UDP being the default). If a port mapping for the given port and protocol combination already exists on that gateway device, this method tries to overwrite it. If that is not desired, you can retrieve the gateway manually with [method get_gateway] and call [method add_port_mapping] on it, if any. Note that forwarding a well-known port (below 1024) with UPnP may fail depending on the device.
Depending on the gateway device, if a mapping for that port already exists, it will either be updated or it will refuse this command due to that conflict, especially if the existing mapping for that port wasn't created via UPnP or points to a different network address (or device) than this one.
If [param port_internal] is [code]0[/code] (the default), the same port number is used for both the external and the internal port (the [param port] value).
The description ([param desc]) is shown in some routers management UIs and can be used to point out which application added the mapping.
The mapping's lease [param duration] can be limited by specifying a duration in seconds. The default of [code]0[/code] means no duration, i.e. a permanent lease and notably some devices only support these permanent leases. Note that whether permanent or not, this is only a request and the gateway may still decide at any point to remove the mapping (which usually happens on a reboot of the gateway, when its external IP address changes, or on some models when it detects a port mapping has become inactive, i.e. had no traffic for multiple minutes). If not [code]0[/code] (permanent), the allowed range according to spec is between [code]120[/code] (2 minutes) and [code]86400[/code] seconds (24 hours).
See [enum UPNPResult] for possible return values.
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_add_port_mapping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Deletes the port mapping for the given port and protocol combination on the default gateway (see [method get_gateway]) if one exists. [param port] must be a valid port between 1 and 65535, [param proto] can be either [code]"TCP"[/code] or [code]"UDP"[/code]. May be refused for mappings pointing to addresses other than this one, for well-known ports (below 1024), or for mappings not added via UPnP. See [enum UPNPResult] for possible return values.
*/
//go:nosplit
func (self class) DeletePortMapping(port gd.Int, proto gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, discreet.Get(proto))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_delete_port_mapping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiscoverMulticastIf(m_if gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(m_if))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_set_discover_multicast_if, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDiscoverMulticastIf() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_get_discover_multicast_if, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiscoverLocalPort(port gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_set_discover_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDiscoverLocalPort() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_get_discover_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDiscoverIpv6(ipv6 bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, ipv6)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_set_discover_ipv6, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDiscoverIpv6() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.UPNP.Bind_is_discover_ipv6, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsUPNP() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsUPNP() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("UPNP", func(ptr gd.Object) any { return classdb.UPNP(ptr) })}
type UPNPResult = classdb.UPNPUPNPResult

const (
/*UPNP command or discovery was successful.*/
	UpnpResultSuccess UPNPResult = 0
/*Not authorized to use the command on the [UPNPDevice]. May be returned when the user disabled UPNP on their router.*/
	UpnpResultNotAuthorized UPNPResult = 1
/*No port mapping was found for the given port, protocol combination on the given [UPNPDevice].*/
	UpnpResultPortMappingNotFound UPNPResult = 2
/*Inconsistent parameters.*/
	UpnpResultInconsistentParameters UPNPResult = 3
/*No such entry in array. May be returned if a given port, protocol combination is not found on an [UPNPDevice].*/
	UpnpResultNoSuchEntryInArray UPNPResult = 4
/*The action failed.*/
	UpnpResultActionFailed UPNPResult = 5
/*The [UPNPDevice] does not allow wildcard values for the source IP address.*/
	UpnpResultSrcIpWildcardNotPermitted UPNPResult = 6
/*The [UPNPDevice] does not allow wildcard values for the external port.*/
	UpnpResultExtPortWildcardNotPermitted UPNPResult = 7
/*The [UPNPDevice] does not allow wildcard values for the internal port.*/
	UpnpResultIntPortWildcardNotPermitted UPNPResult = 8
/*The remote host value must be a wildcard.*/
	UpnpResultRemoteHostMustBeWildcard UPNPResult = 9
/*The external port value must be a wildcard.*/
	UpnpResultExtPortMustBeWildcard UPNPResult = 10
/*No port maps are available. May also be returned if port mapping functionality is not available.*/
	UpnpResultNoPortMapsAvailable UPNPResult = 11
/*Conflict with other mechanism. May be returned instead of [constant UPNP_RESULT_CONFLICT_WITH_OTHER_MAPPING] if a port mapping conflicts with an existing one.*/
	UpnpResultConflictWithOtherMechanism UPNPResult = 12
/*Conflict with an existing port mapping.*/
	UpnpResultConflictWithOtherMapping UPNPResult = 13
/*External and internal port values must be the same.*/
	UpnpResultSamePortValuesRequired UPNPResult = 14
/*Only permanent leases are supported. Do not use the [code]duration[/code] parameter when adding port mappings.*/
	UpnpResultOnlyPermanentLeaseSupported UPNPResult = 15
/*Invalid gateway.*/
	UpnpResultInvalidGateway UPNPResult = 16
/*Invalid port.*/
	UpnpResultInvalidPort UPNPResult = 17
/*Invalid protocol.*/
	UpnpResultInvalidProtocol UPNPResult = 18
/*Invalid duration.*/
	UpnpResultInvalidDuration UPNPResult = 19
/*Invalid arguments.*/
	UpnpResultInvalidArgs UPNPResult = 20
/*Invalid response.*/
	UpnpResultInvalidResponse UPNPResult = 21
/*Invalid parameter.*/
	UpnpResultInvalidParam UPNPResult = 22
/*HTTP error.*/
	UpnpResultHttpError UPNPResult = 23
/*Socket error.*/
	UpnpResultSocketError UPNPResult = 24
/*Error allocating memory.*/
	UpnpResultMemAllocError UPNPResult = 25
/*No gateway available. You may need to call [method discover] first, or discovery didn't detect any valid IGDs (InternetGatewayDevices).*/
	UpnpResultNoGateway UPNPResult = 26
/*No devices available. You may need to call [method discover] first, or discovery didn't detect any valid [UPNPDevice]s.*/
	UpnpResultNoDevices UPNPResult = 27
/*Unknown error.*/
	UpnpResultUnknownError UPNPResult = 28
)
