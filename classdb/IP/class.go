// Package IP provides methods for working with IP object instances.
package IP

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "net/netip"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

/*
IP contains support functions for the Internet Protocol (IP). TCP/IP support is in different classes (see [StreamPeerTCP] and [TCPServer]). IP provides DNS hostname resolution support, both blocking and threaded.
*/
var self [1]gdclass.IP
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.IP)
	self = *(*[1]gdclass.IP)(unsafe.Pointer(&obj))
}

/*
Returns a given hostname's IPv4 or IPv6 address when resolved (blocking-type method). The address type returned depends on the [enum Type] constant given as [param ip_type].
*/
func ResolveHostname(host string) string { //gd:IP.resolve_hostname
	once.Do(singleton)
	return string(class(self).ResolveHostname(gd.NewString(host), 3).String())
}

/*
Resolves a given hostname in a blocking way. Addresses are returned as an [Array] of IPv4 or IPv6 addresses depending on [param ip_type].
*/
func ResolveHostnameAddresses(host string) []string { //gd:IP.resolve_hostname_addresses
	once.Do(singleton)
	return []string(class(self).ResolveHostnameAddresses(gd.NewString(host), 3).Strings())
}

/*
Creates a queue item to resolve a hostname to an IPv4 or IPv6 address depending on the [enum Type] constant given as [param ip_type]. Returns the queue ID if successful, or [constant RESOLVER_INVALID_ID] on error.
*/
func ResolveHostnameQueueItem(host string) int { //gd:IP.resolve_hostname_queue_item
	once.Do(singleton)
	return int(int(class(self).ResolveHostnameQueueItem(gd.NewString(host), 3)))
}

/*
Returns a queued hostname's status as a [enum ResolverStatus] constant, given its queue [param id].
*/
func GetResolveItemStatus(id int) gdclass.IPResolverStatus { //gd:IP.get_resolve_item_status
	once.Do(singleton)
	return gdclass.IPResolverStatus(class(self).GetResolveItemStatus(gd.Int(id)))
}

/*
Returns a queued hostname's IP address, given its queue [param id]. Returns an empty string on error or if resolution hasn't happened yet (see [method get_resolve_item_status]).
*/
func GetResolveItemAddress(id int) string { //gd:IP.get_resolve_item_address
	once.Do(singleton)
	return string(class(self).GetResolveItemAddress(gd.Int(id)).String())
}

/*
Returns resolved addresses, or an empty array if an error happened or resolution didn't happen yet (see [method get_resolve_item_status]).
*/
func GetResolveItemAddresses(id int) []any { //gd:IP.get_resolve_item_addresses
	once.Do(singleton)
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetResolveItemAddresses(gd.Int(id)))))
}

/*
Removes a given item [param id] from the queue. This should be used to free a queue after it has completed to enable more queries to happen.
*/
func EraseResolveItem(id int) { //gd:IP.erase_resolve_item
	once.Do(singleton)
	class(self).EraseResolveItem(gd.Int(id))
}

/*
Returns all the user's current IPv4 and IPv6 addresses as an array.
*/
func GetLocalAddresses() []string { //gd:IP.get_local_addresses
	once.Do(singleton)
	return []string(class(self).GetLocalAddresses().Strings())
}

/*
Returns all network adapters as an array.
Each adapter is a dictionary of the form:
[codeblock]

	{
	    "index": "1", # Interface index.
	    "name": "eth0", # Interface name.
	    "friendly": "Ethernet One", # A friendly name (might be empty).
	    "addresses": ["192.168.1.101"], # An array of IP addresses associated to this interface.
	}

[/codeblock]
*/
func GetLocalInterfaces() []LocalInterface { //gd:IP.get_local_interfaces
	once.Do(singleton)
	return []LocalInterface(gd.ArrayAs[[]LocalInterface](gd.InternalArray(class(self).GetLocalInterfaces())))
}

/*
Removes all of a [param hostname]'s cached references. If no [param hostname] is given, all cached IP addresses are removed.
*/
func ClearCache() { //gd:IP.clear_cache
	once.Do(singleton)
	class(self).ClearCache(gd.NewString(""))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.IP

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns a given hostname's IPv4 or IPv6 address when resolved (blocking-type method). The address type returned depends on the [enum Type] constant given as [param ip_type].
*/
//go:nosplit
func (self class) ResolveHostname(host gd.String, ip_type gdclass.IPType) gd.String { //gd:IP.resolve_hostname
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(host))
	callframe.Arg(frame, ip_type)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_resolve_hostname, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Resolves a given hostname in a blocking way. Addresses are returned as an [Array] of IPv4 or IPv6 addresses depending on [param ip_type].
*/
//go:nosplit
func (self class) ResolveHostnameAddresses(host gd.String, ip_type gdclass.IPType) gd.PackedStringArray { //gd:IP.resolve_hostname_addresses
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(host))
	callframe.Arg(frame, ip_type)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_resolve_hostname_addresses, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates a queue item to resolve a hostname to an IPv4 or IPv6 address depending on the [enum Type] constant given as [param ip_type]. Returns the queue ID if successful, or [constant RESOLVER_INVALID_ID] on error.
*/
//go:nosplit
func (self class) ResolveHostnameQueueItem(host gd.String, ip_type gdclass.IPType) gd.Int { //gd:IP.resolve_hostname_queue_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(host))
	callframe.Arg(frame, ip_type)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_resolve_hostname_queue_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a queued hostname's status as a [enum ResolverStatus] constant, given its queue [param id].
*/
//go:nosplit
func (self class) GetResolveItemStatus(id gd.Int) gdclass.IPResolverStatus { //gd:IP.get_resolve_item_status
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gdclass.IPResolverStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_get_resolve_item_status, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a queued hostname's IP address, given its queue [param id]. Returns an empty string on error or if resolution hasn't happened yet (see [method get_resolve_item_status]).
*/
//go:nosplit
func (self class) GetResolveItemAddress(id gd.Int) gd.String { //gd:IP.get_resolve_item_address
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_get_resolve_item_address, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns resolved addresses, or an empty array if an error happened or resolution didn't happen yet (see [method get_resolve_item_status]).
*/
//go:nosplit
func (self class) GetResolveItemAddresses(id gd.Int) Array.Any { //gd:IP.get_resolve_item_addresses
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_get_resolve_item_addresses, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes a given item [param id] from the queue. This should be used to free a queue after it has completed to enable more queries to happen.
*/
//go:nosplit
func (self class) EraseResolveItem(id gd.Int) { //gd:IP.erase_resolve_item
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_erase_resolve_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns all the user's current IPv4 and IPv6 addresses as an array.
*/
//go:nosplit
func (self class) GetLocalAddresses() gd.PackedStringArray { //gd:IP.get_local_addresses
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_get_local_addresses, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns all network adapters as an array.
Each adapter is a dictionary of the form:
[codeblock]
{
    "index": "1", # Interface index.
    "name": "eth0", # Interface name.
    "friendly": "Ethernet One", # A friendly name (might be empty).
    "addresses": ["192.168.1.101"], # An array of IP addresses associated to this interface.
}
[/codeblock]
*/
//go:nosplit
func (self class) GetLocalInterfaces() Array.Contains[Dictionary.Any] { //gd:IP.get_local_interfaces
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_get_local_interfaces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes all of a [param hostname]'s cached references. If no [param hostname] is given, all cached IP addresses are removed.
*/
//go:nosplit
func (self class) ClearCache(hostname gd.String) { //gd:IP.clear_cache
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(hostname))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.IP.Bind_clear_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("IP", func(ptr gd.Object) any { return [1]gdclass.IP{*(*gdclass.IP)(unsafe.Pointer(&ptr))} })
}

type ResolverStatus = gdclass.IPResolverStatus //gd:IP.ResolverStatus

const (
	/*DNS hostname resolver status: No status.*/
	ResolverStatusNone ResolverStatus = 0
	/*DNS hostname resolver status: Waiting.*/
	ResolverStatusWaiting ResolverStatus = 1
	/*DNS hostname resolver status: Done.*/
	ResolverStatusDone ResolverStatus = 2
	/*DNS hostname resolver status: Error.*/
	ResolverStatusError ResolverStatus = 3
)

type Type = gdclass.IPType //gd:IP.Type

const (
	/*Address type: None.*/
	TypeNone Type = 0
	/*Address type: Internet protocol version 4 (IPv4).*/
	TypeIpv4 Type = 1
	/*Address type: Internet protocol version 6 (IPv6).*/
	TypeIpv6 Type = 2
	/*Address type: Any.*/
	TypeAny Type = 3
)

type LocalInterface struct {
	Index     string       `gd:"index"`
	Name      string       `gd:"name"`
	Friendly  string       `gd:"friendly"`
	Addresses []netip.Addr `gd:"addresses"`
}
