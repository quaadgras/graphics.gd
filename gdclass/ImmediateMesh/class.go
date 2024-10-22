package ImmediateMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Mesh"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.ImmediateMesh

/*
Begin a new surface.
*/
func (self Go) SurfaceBegin(primitive classdb.MeshPrimitiveType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SurfaceBegin(primitive, ([1]gdclass.Material{}[0]))
}

/*
Set the color attribute that will be pushed with the next vertex.
*/
func (self Go) SurfaceSetColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SurfaceSetColor(color)
}

/*
Set the normal attribute that will be pushed with the next vertex.
*/
func (self Go) SurfaceSetNormal(normal gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SurfaceSetNormal(normal)
}

/*
Set the tangent attribute that will be pushed with the next vertex.
*/
func (self Go) SurfaceSetTangent(tangent gd.Plane) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SurfaceSetTangent(tangent)
}

/*
Set the UV attribute that will be pushed with the next vertex.
*/
func (self Go) SurfaceSetUv(uv gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SurfaceSetUv(uv)
}

/*
Set the UV2 attribute that will be pushed with the next vertex.
*/
func (self Go) SurfaceSetUv2(uv2 gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SurfaceSetUv2(uv2)
}

/*
Add a 3D vertex using the current attributes previously set.
*/
func (self Go) SurfaceAddVertex(vertex gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SurfaceAddVertex(vertex)
}

/*
Add a 2D vertex using the current attributes previously set.
*/
func (self Go) SurfaceAddVertex2d(vertex gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SurfaceAddVertex2d(vertex)
}

/*
End and commit current surface. Note that surface being created will not be visible until this function is called.
*/
func (self Go) SurfaceEnd() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SurfaceEnd()
}

/*
Clear all surfaces.
*/
func (self Go) ClearSurfaces() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearSurfaces()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ImmediateMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ImmediateMesh"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Begin a new surface.
*/
//go:nosplit
func (self class) SurfaceBegin(primitive classdb.MeshPrimitiveType, material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, primitive)
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_surface_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set the color attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_surface_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set the normal attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetNormal(normal gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, normal)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_surface_set_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set the tangent attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetTangent(tangent gd.Plane)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tangent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_surface_set_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set the UV attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetUv(uv gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, uv)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_surface_set_uv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set the UV2 attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetUv2(uv2 gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, uv2)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_surface_set_uv2, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Add a 3D vertex using the current attributes previously set.
*/
//go:nosplit
func (self class) SurfaceAddVertex(vertex gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vertex)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_surface_add_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Add a 2D vertex using the current attributes previously set.
*/
//go:nosplit
func (self class) SurfaceAddVertex2d(vertex gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, vertex)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_surface_add_vertex_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
End and commit current surface. Note that surface being created will not be visible until this function is called.
*/
//go:nosplit
func (self class) SurfaceEnd()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_surface_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clear all surfaces.
*/
//go:nosplit
func (self class) ClearSurfaces()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImmediateMesh.Bind_clear_surfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsImmediateMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsImmediateMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMesh() Mesh.GD { return *((*Mesh.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMesh() Mesh.Go { return *((*Mesh.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMesh(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMesh(), name)
	}
}
func init() {classdb.Register("ImmediateMesh", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
