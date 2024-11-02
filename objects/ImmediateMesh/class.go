package ImmediateMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Mesh"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A mesh type optimized for creating geometry manually, similar to OpenGL 1.x immediate mode.
Here's a sample on how to generate a triangular face:
[codeblocks]
[gdscript]
var mesh = ImmediateMesh.new()
mesh.surface_begin(Mesh.PRIMITIVE_TRIANGLES)
mesh.surface_add_vertex(Vector3.LEFT)
mesh.surface_add_vertex(Vector3.FORWARD)
mesh.surface_add_vertex(Vector3.ZERO)
mesh.surface_end()
[/gdscript]
[csharp]
var mesh = new ImmediateMesh();
mesh.SurfaceBegin(Mesh.PrimitiveType.Triangles);
mesh.SurfaceAddVertex(Vector3.Left);
mesh.SurfaceAddVertex(Vector3.Forward);
mesh.SurfaceAddVertex(Vector3.Zero);
mesh.SurfaceEnd();
[/csharp]
[/codeblocks]
[b]Note:[/b] Generating complex geometries with [ImmediateMesh] is highly inefficient. Instead, it is designed to generate simple geometry that changes often.
*/
type Instance [1]classdb.ImmediateMesh

/*
Begin a new surface.
*/
func (self Instance) SurfaceBegin(primitive classdb.MeshPrimitiveType) {
	class(self).SurfaceBegin(primitive, ([1]objects.Material{}[0]))
}

/*
Set the color attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetColor(color gd.Color) {
	class(self).SurfaceSetColor(color)
}

/*
Set the normal attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetNormal(normal gd.Vector3) {
	class(self).SurfaceSetNormal(normal)
}

/*
Set the tangent attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetTangent(tangent gd.Plane) {
	class(self).SurfaceSetTangent(tangent)
}

/*
Set the UV attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetUv(uv gd.Vector2) {
	class(self).SurfaceSetUv(uv)
}

/*
Set the UV2 attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetUv2(uv2 gd.Vector2) {
	class(self).SurfaceSetUv2(uv2)
}

/*
Add a 3D vertex using the current attributes previously set.
*/
func (self Instance) SurfaceAddVertex(vertex gd.Vector3) {
	class(self).SurfaceAddVertex(vertex)
}

/*
Add a 2D vertex using the current attributes previously set.
*/
func (self Instance) SurfaceAddVertex2d(vertex gd.Vector2) {
	class(self).SurfaceAddVertex2d(vertex)
}

/*
End and commit current surface. Note that surface being created will not be visible until this function is called.
*/
func (self Instance) SurfaceEnd() {
	class(self).SurfaceEnd()
}

/*
Clear all surfaces.
*/
func (self Instance) ClearSurfaces() {
	class(self).ClearSurfaces()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ImmediateMesh

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImmediateMesh"))
	return Instance{classdb.ImmediateMesh(object)}
}

/*
Begin a new surface.
*/
//go:nosplit
func (self class) SurfaceBegin(primitive classdb.MeshPrimitiveType, material objects.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, primitive)
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Set the color attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Set the normal attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetNormal(normal gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, normal)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Set the tangent attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetTangent(tangent gd.Plane) {
	var frame = callframe.New()
	callframe.Arg(frame, tangent)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Set the UV attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetUv(uv gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, uv)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_uv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Set the UV2 attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetUv2(uv2 gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, uv2)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_uv2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Add a 3D vertex using the current attributes previously set.
*/
//go:nosplit
func (self class) SurfaceAddVertex(vertex gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, vertex)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_add_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Add a 2D vertex using the current attributes previously set.
*/
//go:nosplit
func (self class) SurfaceAddVertex2d(vertex gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, vertex)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_add_vertex_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
End and commit current surface. Note that surface being created will not be visible until this function is called.
*/
//go:nosplit
func (self class) SurfaceEnd() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clear all surfaces.
*/
//go:nosplit
func (self class) ClearSurfaces() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_clear_surfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsImmediateMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImmediateMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMesh() Mesh.Advanced        { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Mesh.Instance     { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsMesh(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMesh(), name)
	}
}
func init() {
	classdb.Register("ImmediateMesh", func(ptr gd.Object) any { return classdb.ImmediateMesh(ptr) })
}
