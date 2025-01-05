// Package ORMMaterial3D provides methods for working with ORMMaterial3D object instances.
package ORMMaterial3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/BaseMaterial3D"
import "graphics.gd/classdb/Material"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
ORMMaterial3D's properties are inherited from [BaseMaterial3D]. Unlike [StandardMaterial3D], ORMMaterial3D uses a single texture for ambient occlusion, roughness and metallic maps, known as an ORM texture.
*/
type Instance [1]gdclass.ORMMaterial3D
type Any interface {
	gd.IsClass
	AsORMMaterial3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ORMMaterial3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ORMMaterial3D"))
	return Instance{*(*gdclass.ORMMaterial3D)(unsafe.Pointer(&object))}
}

func (self class) AsORMMaterial3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsORMMaterial3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsBaseMaterial3D() BaseMaterial3D.Advanced {
	return *((*BaseMaterial3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBaseMaterial3D() BaseMaterial3D.Instance {
	return *((*BaseMaterial3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMaterial() Material.Advanced {
	return *((*Material.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMaterial() Material.Instance {
	return *((*Material.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(BaseMaterial3D.Advanced(self.AsBaseMaterial3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(BaseMaterial3D.Instance(self.AsBaseMaterial3D()), name)
	}
}
func init() {
	gdclass.Register("ORMMaterial3D", func(ptr gd.Object) any {
		return [1]gdclass.ORMMaterial3D{*(*gdclass.ORMMaterial3D)(unsafe.Pointer(&ptr))}
	})
}
