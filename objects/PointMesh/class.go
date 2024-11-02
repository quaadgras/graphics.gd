package PointMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/PrimitiveMesh"
import "grow.graphics/gd/objects/Mesh"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
The PointMesh is made from a single point. Instead of relying on triangles, points are rendered as a single rectangle on the screen with a constant size. They are intended to be used with Particle systems, but can be used as a cheap way to render constant size billboarded sprites (for example in a point cloud).
PointMeshes, must be used with a material that has a point size. Point size can be accessed in a shader with [code]POINT_SIZE[/code], or in a [BaseMaterial3D] by setting [member BaseMaterial3D.use_point_size] and the variable [member BaseMaterial3D.point_size].
When using PointMeshes, properties that normally alter vertices will be ignored, including billboard mode, grow, and cull face.
*/
type Instance [1]classdb.PointMesh

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PointMesh

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PointMesh"))
	return Instance{classdb.PointMesh(object)}
}

func (self class) AsPointMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPointMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPrimitiveMesh() PrimitiveMesh.Advanced {
	return *((*PrimitiveMesh.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPrimitiveMesh() PrimitiveMesh.Instance {
	return *((*PrimitiveMesh.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMesh() Mesh.Advanced    { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Mesh.Instance { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}
func init() { classdb.Register("PointMesh", func(ptr gd.Object) any { return classdb.PointMesh(ptr) }) }
