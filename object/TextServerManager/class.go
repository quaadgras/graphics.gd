package TextServerManager

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
[TextServerManager] is the API backend for loading, enumerating, and switching [TextServer]s.
[b]Note:[/b] Switching text server at runtime is possible, but will invalidate all fonts and text buffers. Make sure to unload all controls, fonts, and themes before doing so.

*/
type Simple [1]classdb.TextServerManager
func (self Simple) AddInterface(intf [1]classdb.TextServer) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddInterface(intf)
}
func (self Simple) GetInterfaceCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInterfaceCount()))
}
func (self Simple) RemoveInterface(intf [1]classdb.TextServer) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveInterface(intf)
}
func (self Simple) GetInterface(idx int) [1]classdb.TextServer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TextServer(Expert(self).GetInterface(gc, gd.Int(idx)))
}
func (self Simple) GetInterfaces() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetInterfaces(gc))
}
func (self Simple) FindInterface(name string) [1]classdb.TextServer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TextServer(Expert(self).FindInterface(gc, gc.String(name)))
}
func (self Simple) SetPrimaryInterface(index [1]classdb.TextServer) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPrimaryInterface(index)
}
func (self Simple) GetPrimaryInterface() [1]classdb.TextServer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TextServer(Expert(self).GetPrimaryInterface(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TextServerManager
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Registers a [TextServer] interface.
*/
//go:nosplit
func (self class) AddInterface(intf object.TextServer)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(intf[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_add_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of interfaces currently registered.
*/
//go:nosplit
func (self class) GetInterfaceCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_get_interface_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes an interface. All fonts and shaped text caches should be freed before removing an interface.
*/
//go:nosplit
func (self class) RemoveInterface(intf object.TextServer)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(intf[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_remove_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the interface registered at a given index.
*/
//go:nosplit
func (self class) GetInterface(ctx gd.Lifetime, idx gd.Int) object.TextServer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_get_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TextServer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a list of available interfaces, with the index and name of each interface.
*/
//go:nosplit
func (self class) GetInterfaces(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_get_interfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Finds an interface by its [param name].
*/
//go:nosplit
func (self class) FindInterface(ctx gd.Lifetime, name gd.String) object.TextServer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_find_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TextServer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the primary [TextServer] interface.
*/
//go:nosplit
func (self class) SetPrimaryInterface(index object.TextServer)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(index[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_set_primary_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the primary [TextServer] interface currently in use.
*/
//go:nosplit
func (self class) GetPrimaryInterface(ctx gd.Lifetime) object.TextServer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_get_primary_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TextServer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTextServerManager() Expert { return self[0].AsTextServerManager() }


//go:nosplit
func (self Simple) AsTextServerManager() Simple { return self[0].AsTextServerManager() }

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
func init() {classdb.Register("TextServerManager", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
