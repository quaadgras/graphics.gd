package StandardMaterial3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/BaseMaterial3D"
import "graphics.gd/objects/Material"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[StandardMaterial3D]'s properties are inherited from [BaseMaterial3D]. [StandardMaterial3D] uses separate textures for ambient occlusion, roughness and metallic maps. To use a single ORM map for all 3 textures, use an [ORMMaterial3D] instead.
*/
type Instance [1]classdb.StandardMaterial3D
type Any interface {
	gd.IsClass
	AsStandardMaterial3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StandardMaterial3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StandardMaterial3D"))
	return Instance{classdb.StandardMaterial3D(object)}
}

func (self class) AsStandardMaterial3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStandardMaterial3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsBaseMaterial3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsBaseMaterial3D(), name)
	}
}
func init() {
	classdb.Register("StandardMaterial3D", func(ptr gd.Object) any { return [1]classdb.StandardMaterial3D{classdb.StandardMaterial3D(ptr)} })
}
