package CollisionShape3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A node that provides a [Shape3D] to a [CollisionObject3D] parent and allows to edit it. This can give a detection shape to an [Area3D] or turn a [PhysicsBody3D] into a solid object.
[b]Warning:[/b] A non-uniformly scaled [CollisionShape3D] will likely not behave as expected. Make sure to keep its scale the same on all axes and adjust its [member shape] resource instead.

*/
type Go [1]classdb.CollisionShape3D

/*
This method does nothing.
*/
func (self Go) ResourceChanged(resource gdclass.Resource) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ResourceChanged(resource)
}

/*
Sets the collision shape's shape to the addition of all its convexed [MeshInstance3D] siblings geometry.
*/
func (self Go) MakeConvexFromSiblings() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MakeConvexFromSiblings()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CollisionShape3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CollisionShape3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Shape() gdclass.Shape3D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Shape3D(class(self).GetShape(gc))
}

func (self Go) SetShape(value gdclass.Shape3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShape(value)
}

func (self Go) Disabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDisabled())
}

func (self Go) SetDisabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisabled(value)
}

/*
This method does nothing.
*/
//go:nosplit
func (self class) ResourceChanged(resource gdclass.Resource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(resource[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionShape3D.Bind_resource_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetShape(shape gdclass.Shape3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shape[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionShape3D.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShape(ctx gd.Lifetime) gdclass.Shape3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionShape3D.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Shape3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionShape3D.Bind_set_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDisabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionShape3D.Bind_is_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the collision shape's shape to the addition of all its convexed [MeshInstance3D] siblings geometry.
*/
//go:nosplit
func (self class) MakeConvexFromSiblings()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CollisionShape3D.Bind_make_convex_from_siblings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsCollisionShape3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCollisionShape3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("CollisionShape3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
