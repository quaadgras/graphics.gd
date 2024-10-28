package BoxShape3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Shape3D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A 3D box shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape3D].
[b]Performance:[/b] [BoxShape3D] is fast to check collisions against. It is faster than [CapsuleShape3D] and [CylinderShape3D], but slower than [SphereShape3D].

*/
type Go [1]classdb.BoxShape3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.BoxShape3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BoxShape3D"))
	return Go{classdb.BoxShape3D(object)}
}

func (self Go) Size() gd.Vector3 {
		return gd.Vector3(class(self).GetSize())
}

func (self Go) SetSize(value gd.Vector3) {
	class(self).SetSize(value)
}

//go:nosplit
func (self class) SetSize(size gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxShape3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxShape3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsBoxShape3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsBoxShape3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsShape3D() Shape3D.GD { return *((*Shape3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsShape3D() Shape3D.Go { return *((*Shape3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape3D(), name)
	}
}
func init() {classdb.Register("BoxShape3D", func(ptr gd.Object) any { return classdb.BoxShape3D(ptr) })}
