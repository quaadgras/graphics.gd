package ImmediateMesh

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Plane"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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
type Instance [1]gdclass.ImmediateMesh
type Any interface {
	gd.IsClass
	AsImmediateMesh() Instance
}

/*
Begin a new surface.
*/
func (self Instance) SurfaceBegin(primitive gdclass.MeshPrimitiveType) {
	class(self).SurfaceBegin(primitive, [1][1]gdclass.Material{}[0])
}

/*
Set the color attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetColor(color Color.RGBA) {
	class(self).SurfaceSetColor(gd.Color(color))
}

/*
Set the normal attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetNormal(normal Vector3.XYZ) {
	class(self).SurfaceSetNormal(gd.Vector3(normal))
}

/*
Set the tangent attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetTangent(tangent Plane.NormalD) {
	class(self).SurfaceSetTangent(gd.Plane(tangent))
}

/*
Set the UV attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetUv(uv Vector2.XY) {
	class(self).SurfaceSetUv(gd.Vector2(uv))
}

/*
Set the UV2 attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetUv2(uv2 Vector2.XY) {
	class(self).SurfaceSetUv2(gd.Vector2(uv2))
}

/*
Add a 3D vertex using the current attributes previously set.
*/
func (self Instance) SurfaceAddVertex(vertex Vector3.XYZ) {
	class(self).SurfaceAddVertex(gd.Vector3(vertex))
}

/*
Add a 2D vertex using the current attributes previously set.
*/
func (self Instance) SurfaceAddVertex2d(vertex Vector2.XY) {
	class(self).SurfaceAddVertex2d(gd.Vector2(vertex))
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
type class [1]gdclass.ImmediateMesh

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImmediateMesh"))
	return Instance{*(*gdclass.ImmediateMesh)(unsafe.Pointer(&object))}
}

/*
Begin a new surface.
*/
//go:nosplit
func (self class) SurfaceBegin(primitive gdclass.MeshPrimitiveType, material [1]gdclass.Material) {
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
	gdclass.Register("ImmediateMesh", func(ptr gd.Object) any {
		return [1]gdclass.ImmediateMesh{*(*gdclass.ImmediateMesh)(unsafe.Pointer(&ptr))}
	})
}
