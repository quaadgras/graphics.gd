package PhysicsServer3DManager

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
[PhysicsServer3DManager] is the API for registering [PhysicsServer3D] implementations and for setting the default implementation.
[b]Note:[/b] It is not possible to switch physics servers at runtime. This class is only used on startup at the server initialization level, by Godot itself and possibly by GDExtensions.

*/
type Simple [1]classdb.PhysicsServer3DManager
func (self Simple) RegisterServer(name string, create_callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegisterServer(gc.String(name), create_callback)
}
func (self Simple) SetDefaultServer(name string, priority int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultServer(gc.String(name), gd.Int(priority))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicsServer3DManager
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Register a [PhysicsServer3D] implementation by passing a [param name] and a [Callable] that returns a [PhysicsServer3D] object.
*/
//go:nosplit
func (self class) RegisterServer(name gd.String, create_callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(create_callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3DManager.Bind_register_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set the default [PhysicsServer3D] implementation to the one identified by [param name], if [param priority] is greater than the priority of the current default implementation.
*/
//go:nosplit
func (self class) SetDefaultServer(name gd.String, priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3DManager.Bind_set_default_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsPhysicsServer3DManager() Expert { return self[0].AsPhysicsServer3DManager() }


//go:nosplit
func (self Simple) AsPhysicsServer3DManager() Simple { return self[0].AsPhysicsServer3DManager() }

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
func init() {classdb.Register("PhysicsServer3DManager", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
