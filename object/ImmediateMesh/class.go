package ImmediateMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Mesh"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
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
type Simple [1]classdb.ImmediateMesh
func (self Simple) SurfaceBegin(primitive classdb.MeshPrimitiveType, material [1]classdb.Material) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SurfaceBegin(primitive, material)
}
func (self Simple) SurfaceSetColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SurfaceSetColor(color)
}
func (self Simple) SurfaceSetNormal(normal gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SurfaceSetNormal(normal)
}
func (self Simple) SurfaceSetTangent(tangent gd.Plane) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SurfaceSetTangent(tangent)
}
func (self Simple) SurfaceSetUv(uv gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SurfaceSetUv(uv)
}
func (self Simple) SurfaceSetUv2(uv2 gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SurfaceSetUv2(uv2)
}
func (self Simple) SurfaceAddVertex(vertex gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SurfaceAddVertex(vertex)
}
func (self Simple) SurfaceAddVertex2d(vertex gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SurfaceAddVertex2d(vertex)
}
func (self Simple) SurfaceEnd() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SurfaceEnd()
}
func (self Simple) ClearSurfaces() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearSurfaces()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ImmediateMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Begin a new surface.
*/
//go:nosplit
func (self class) SurfaceBegin(primitive classdb.MeshPrimitiveType, material object.Material)  {
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

//go:nosplit
func (self class) AsImmediateMesh() Expert { return self[0].AsImmediateMesh() }


//go:nosplit
func (self Simple) AsImmediateMesh() Simple { return self[0].AsImmediateMesh() }


//go:nosplit
func (self class) AsMesh() Mesh.Expert { return self[0].AsMesh() }


//go:nosplit
func (self Simple) AsMesh() Mesh.Simple { return self[0].AsMesh() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

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
func init() {classdb.Register("ImmediateMesh", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
