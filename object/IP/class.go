package IP

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
IP contains support functions for the Internet Protocol (IP). TCP/IP support is in different classes (see [StreamPeerTCP] and [TCPServer]). IP provides DNS hostname resolution support, both blocking and threaded.

*/
type Simple [1]classdb.IP
func (self Simple) ResolveHostname(host string, ip_type classdb.IPType) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).ResolveHostname(gc, gc.String(host), ip_type).String())
}
func (self Simple) ResolveHostnameAddresses(host string, ip_type classdb.IPType) gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).ResolveHostnameAddresses(gc, gc.String(host), ip_type))
}
func (self Simple) ResolveHostnameQueueItem(host string, ip_type classdb.IPType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).ResolveHostnameQueueItem(gc.String(host), ip_type)))
}
func (self Simple) GetResolveItemStatus(id int) classdb.IPResolverStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.IPResolverStatus(Expert(self).GetResolveItemStatus(gd.Int(id)))
}
func (self Simple) GetResolveItemAddress(id int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetResolveItemAddress(gc, gd.Int(id)).String())
}
func (self Simple) GetResolveItemAddresses(id int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetResolveItemAddresses(gc, gd.Int(id)))
}
func (self Simple) EraseResolveItem(id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EraseResolveItem(gd.Int(id))
}
func (self Simple) GetLocalAddresses() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetLocalAddresses(gc))
}
func (self Simple) GetLocalInterfaces() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetLocalInterfaces(gc))
}
func (self Simple) ClearCache(hostname string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearCache(gc.String(hostname))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.IP
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns a given hostname's IPv4 or IPv6 address when resolved (blocking-type method). The address type returned depends on the [enum Type] constant given as [param ip_type].
*/
//go:nosplit
func (self class) ResolveHostname(ctx gd.Lifetime, host gd.String, ip_type classdb.IPType) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(host))
	callframe.Arg(frame, ip_type)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_resolve_hostname, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Resolves a given hostname in a blocking way. Addresses are returned as an [Array] of IPv4 or IPv6 addresses depending on [param ip_type].
*/
//go:nosplit
func (self class) ResolveHostnameAddresses(ctx gd.Lifetime, host gd.String, ip_type classdb.IPType) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(host))
	callframe.Arg(frame, ip_type)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_resolve_hostname_addresses, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates a queue item to resolve a hostname to an IPv4 or IPv6 address depending on the [enum Type] constant given as [param ip_type]. Returns the queue ID if successful, or [constant RESOLVER_INVALID_ID] on error.
*/
//go:nosplit
func (self class) ResolveHostnameQueueItem(host gd.String, ip_type classdb.IPType) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(host))
	callframe.Arg(frame, ip_type)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_resolve_hostname_queue_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a queued hostname's status as a [enum ResolverStatus] constant, given its queue [param id].
*/
//go:nosplit
func (self class) GetResolveItemStatus(id gd.Int) classdb.IPResolverStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[classdb.IPResolverStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_get_resolve_item_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a queued hostname's IP address, given its queue [param id]. Returns an empty string on error or if resolution hasn't happened yet (see [method get_resolve_item_status]).
*/
//go:nosplit
func (self class) GetResolveItemAddress(ctx gd.Lifetime, id gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_get_resolve_item_address, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns resolved addresses, or an empty array if an error happened or resolution didn't happen yet (see [method get_resolve_item_status]).
*/
//go:nosplit
func (self class) GetResolveItemAddresses(ctx gd.Lifetime, id gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_get_resolve_item_addresses, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes a given item [param id] from the queue. This should be used to free a queue after it has completed to enable more queries to happen.
*/
//go:nosplit
func (self class) EraseResolveItem(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_erase_resolve_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns all the user's current IPv4 and IPv6 addresses as an array.
*/
//go:nosplit
func (self class) GetLocalAddresses(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_get_local_addresses, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
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
func (self class) GetLocalInterfaces(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_get_local_interfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Removes all of a [param hostname]'s cached references. If no [param hostname] is given, all cached IP addresses are removed.
*/
//go:nosplit
func (self class) ClearCache(hostname gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(hostname))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.IP.Bind_clear_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsIP() Expert { return self[0].AsIP() }


//go:nosplit
func (self Simple) AsIP() Simple { return self[0].AsIP() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("IP", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ResolverStatus = classdb.IPResolverStatus

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
type Type = classdb.IPType

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
