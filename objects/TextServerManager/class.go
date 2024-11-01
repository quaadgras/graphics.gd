package TextServerManager

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[TextServerManager] is the API backend for loading, enumerating, and switching [TextServer]s.
[b]Note:[/b] Switching text server at runtime is possible, but will invalidate all fonts and text buffers. Make sure to unload all controls, fonts, and themes before doing so.
*/
var self objects.TextServerManager
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.TextServerManager)
	self = *(*objects.TextServerManager)(unsafe.Pointer(&obj))
}

/*
Registers a [TextServer] interface.
*/
func AddInterface(intf objects.TextServer) {
	once.Do(singleton)
	class(self).AddInterface(intf)
}

/*
Returns the number of interfaces currently registered.
*/
func GetInterfaceCount() int {
	once.Do(singleton)
	return int(int(class(self).GetInterfaceCount()))
}

/*
Removes an interface. All fonts and shaped text caches should be freed before removing an interface.
*/
func RemoveInterface(intf objects.TextServer) {
	once.Do(singleton)
	class(self).RemoveInterface(intf)
}

/*
Returns the interface registered at a given index.
*/
func GetInterface(idx int) objects.TextServer {
	once.Do(singleton)
	return objects.TextServer(class(self).GetInterface(gd.Int(idx)))
}

/*
Returns a list of available interfaces, with the index and name of each interface.
*/
func GetInterfaces() gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).GetInterfaces())
}

/*
Finds an interface by its [param name].
*/
func FindInterface(name string) objects.TextServer {
	once.Do(singleton)
	return objects.TextServer(class(self).FindInterface(gd.NewString(name)))
}

/*
Sets the primary [TextServer] interface.
*/
func SetPrimaryInterface(index objects.TextServer) {
	once.Do(singleton)
	class(self).SetPrimaryInterface(index)
}

/*
Returns the primary [TextServer] interface currently in use.
*/
func GetPrimaryInterface() objects.TextServer {
	once.Do(singleton)
	return objects.TextServer(class(self).GetPrimaryInterface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.TextServerManager

func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Registers a [TextServer] interface.
*/
//go:nosplit
func (self class) AddInterface(intf objects.TextServer) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(intf[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServerManager.Bind_add_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of interfaces currently registered.
*/
//go:nosplit
func (self class) GetInterfaceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServerManager.Bind_get_interface_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes an interface. All fonts and shaped text caches should be freed before removing an interface.
*/
//go:nosplit
func (self class) RemoveInterface(intf objects.TextServer) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(intf[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServerManager.Bind_remove_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the interface registered at a given index.
*/
//go:nosplit
func (self class) GetInterface(idx gd.Int) objects.TextServer {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServerManager.Bind_get_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.TextServer{classdb.TextServer(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns a list of available interfaces, with the index and name of each interface.
*/
//go:nosplit
func (self class) GetInterfaces() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServerManager.Bind_get_interfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Finds an interface by its [param name].
*/
//go:nosplit
func (self class) FindInterface(name gd.String) objects.TextServer {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServerManager.Bind_find_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.TextServer{classdb.TextServer(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets the primary [TextServer] interface.
*/
//go:nosplit
func (self class) SetPrimaryInterface(index objects.TextServer) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(index[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServerManager.Bind_set_primary_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the primary [TextServer] interface currently in use.
*/
//go:nosplit
func (self class) GetPrimaryInterface() objects.TextServer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServerManager.Bind_get_primary_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.TextServer{classdb.TextServer(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func OnInterfaceAdded(cb func(interface_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("interface_added"), gd.NewCallable(cb), 0)
}

func OnInterfaceRemoved(cb func(interface_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("interface_removed"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("TextServerManager", func(ptr gd.Object) any { return classdb.TextServerManager(ptr) })
}
