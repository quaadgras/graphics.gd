package TextServerManager

import "unsafe"
import "sync"
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
[TextServerManager] is the API backend for loading, enumerating, and switching [TextServer]s.
[b]Note:[/b] Switching text server at runtime is possible, but will invalidate all fonts and text buffers. Make sure to unload all controls, fonts, and themes before doing so.

*/
var self gdclass.TextServerManager
var once sync.Once
func singleton() {
	gc := gd.Static
	obj := gc.API.Object.GetSingleton(gc, gc.API.Singletons.TextServerManager)
	self = *(*gdclass.TextServerManager)(unsafe.Pointer(&obj))
}

/*
Registers a [TextServer] interface.
*/
func AddInterface(intf gdclass.TextServer) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).AddInterface(intf)
}

/*
Returns the number of interfaces currently registered.
*/
func GetInterfaceCount() int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).GetInterfaceCount()))
}

/*
Removes an interface. All fonts and shaped text caches should be freed before removing an interface.
*/
func RemoveInterface(intf gdclass.TextServer) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).RemoveInterface(intf)
}

/*
Returns the interface registered at a given index.
*/
func GetInterface(idx int) gdclass.TextServer {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.TextServer(class(self).GetInterface(gc, gd.Int(idx)))
}

/*
Returns a list of available interfaces, with the index and name of each interface.
*/
func GetInterfaces() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.ArrayOf[gd.Dictionary](class(self).GetInterfaces(gc))
}

/*
Finds an interface by its [param name].
*/
func FindInterface(name string) gdclass.TextServer {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.TextServer(class(self).FindInterface(gc, gc.String(name)))
}

/*
Sets the primary [TextServer] interface.
*/
func SetPrimaryInterface(index gdclass.TextServer) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).SetPrimaryInterface(index)
}

/*
Returns the primary [TextServer] interface currently in use.
*/
func GetPrimaryInterface() gdclass.TextServer {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.TextServer(class(self).GetPrimaryInterface(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.TextServerManager
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Registers a [TextServer] interface.
*/
//go:nosplit
func (self class) AddInterface(intf gdclass.TextServer)  {
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
func (self class) RemoveInterface(intf gdclass.TextServer)  {
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
func (self class) GetInterface(ctx gd.Lifetime, idx gd.Int) gdclass.TextServer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_get_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TextServer
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
func (self class) FindInterface(ctx gd.Lifetime, name gd.String) gdclass.TextServer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_find_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TextServer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the primary [TextServer] interface.
*/
//go:nosplit
func (self class) SetPrimaryInterface(index gdclass.TextServer)  {
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
func (self class) GetPrimaryInterface(ctx gd.Lifetime) gdclass.TextServer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextServerManager.Bind_get_primary_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TextServer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func OnInterfaceAdded(cb func(interface_name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("interface_added"), gc.Callable(cb), 0)
}


func OnInterfaceRemoved(cb func(interface_name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("interface_removed"), gc.Callable(cb), 0)
}


func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("TextServerManager", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
