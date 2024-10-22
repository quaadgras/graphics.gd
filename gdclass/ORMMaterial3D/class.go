package ORMMaterial3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/BaseMaterial3D"
import "grow.graphics/gd/gdclass/Material"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
ORMMaterial3D's properties are inherited from [BaseMaterial3D]. Unlike [StandardMaterial3D], ORMMaterial3D uses a single texture for ambient occlusion, roughness and metallic maps, known as an ORM texture.

*/
type Go [1]classdb.ORMMaterial3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ORMMaterial3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ORMMaterial3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsORMMaterial3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsORMMaterial3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsBaseMaterial3D() BaseMaterial3D.GD { return *((*BaseMaterial3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsBaseMaterial3D() BaseMaterial3D.Go { return *((*BaseMaterial3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.GD { return *((*Material.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMaterial() Material.Go { return *((*Material.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsBaseMaterial3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsBaseMaterial3D(), name)
	}
}
func init() {classdb.Register("ORMMaterial3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
