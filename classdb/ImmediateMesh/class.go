// Package ImmediateMesh provides methods for working with ImmediateMesh object instances.
package ImmediateMesh

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Plane"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

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

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsImmediateMesh() Instance
}

/*
Begin a new surface.
*/
func (self Instance) SurfaceBegin(primitive gdclass.MeshPrimitiveType) { //gd:ImmediateMesh.surface_begin
	class(self).SurfaceBegin(primitive, [1][1]gdclass.Material{}[0])
}

/*
Set the color attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetColor(color Color.RGBA) { //gd:ImmediateMesh.surface_set_color
	class(self).SurfaceSetColor(Color.RGBA(color))
}

/*
Set the normal attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetNormal(normal Vector3.XYZ) { //gd:ImmediateMesh.surface_set_normal
	class(self).SurfaceSetNormal(Vector3.XYZ(normal))
}

/*
Set the tangent attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetTangent(tangent Plane.NormalD) { //gd:ImmediateMesh.surface_set_tangent
	class(self).SurfaceSetTangent(Plane.NormalD(tangent))
}

/*
Set the UV attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetUv(uv Vector2.XY) { //gd:ImmediateMesh.surface_set_uv
	class(self).SurfaceSetUv(Vector2.XY(uv))
}

/*
Set the UV2 attribute that will be pushed with the next vertex.
*/
func (self Instance) SurfaceSetUv2(uv2 Vector2.XY) { //gd:ImmediateMesh.surface_set_uv2
	class(self).SurfaceSetUv2(Vector2.XY(uv2))
}

/*
Add a 3D vertex using the current attributes previously set.
*/
func (self Instance) SurfaceAddVertex(vertex Vector3.XYZ) { //gd:ImmediateMesh.surface_add_vertex
	class(self).SurfaceAddVertex(Vector3.XYZ(vertex))
}

/*
Add a 2D vertex using the current attributes previously set.
*/
func (self Instance) SurfaceAddVertex2d(vertex Vector2.XY) { //gd:ImmediateMesh.surface_add_vertex_2d
	class(self).SurfaceAddVertex2d(Vector2.XY(vertex))
}

/*
End and commit current surface. Note that surface being created will not be visible until this function is called.
*/
func (self Instance) SurfaceEnd() { //gd:ImmediateMesh.surface_end
	class(self).SurfaceEnd()
}

/*
Clear all surfaces.
*/
func (self Instance) ClearSurfaces() { //gd:ImmediateMesh.clear_surfaces
	class(self).ClearSurfaces()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ImmediateMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImmediateMesh"))
	casted := Instance{*(*gdclass.ImmediateMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Begin a new surface.
*/
//go:nosplit
func (self class) SurfaceBegin(primitive gdclass.MeshPrimitiveType, material [1]gdclass.Material) { //gd:ImmediateMesh.surface_begin
	var frame = callframe.New()
	callframe.Arg(frame, primitive)
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Set the color attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetColor(color Color.RGBA) { //gd:ImmediateMesh.surface_set_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Set the normal attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetNormal(normal Vector3.XYZ) { //gd:ImmediateMesh.surface_set_normal
	var frame = callframe.New()
	callframe.Arg(frame, normal)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Set the tangent attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetTangent(tangent Plane.NormalD) { //gd:ImmediateMesh.surface_set_tangent
	var frame = callframe.New()
	callframe.Arg(frame, tangent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Set the UV attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetUv(uv Vector2.XY) { //gd:ImmediateMesh.surface_set_uv
	var frame = callframe.New()
	callframe.Arg(frame, uv)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_uv, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Set the UV2 attribute that will be pushed with the next vertex.
*/
//go:nosplit
func (self class) SurfaceSetUv2(uv2 Vector2.XY) { //gd:ImmediateMesh.surface_set_uv2
	var frame = callframe.New()
	callframe.Arg(frame, uv2)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_set_uv2, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Add a 3D vertex using the current attributes previously set.
*/
//go:nosplit
func (self class) SurfaceAddVertex(vertex Vector3.XYZ) { //gd:ImmediateMesh.surface_add_vertex
	var frame = callframe.New()
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_add_vertex, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Add a 2D vertex using the current attributes previously set.
*/
//go:nosplit
func (self class) SurfaceAddVertex2d(vertex Vector2.XY) { //gd:ImmediateMesh.surface_add_vertex_2d
	var frame = callframe.New()
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_add_vertex_2d, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
End and commit current surface. Note that surface being created will not be visible until this function is called.
*/
//go:nosplit
func (self class) SurfaceEnd() { //gd:ImmediateMesh.surface_end
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_surface_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clear all surfaces.
*/
//go:nosplit
func (self class) ClearSurfaces() { //gd:ImmediateMesh.clear_surfaces
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImmediateMesh.Bind_clear_surfaces, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Mesh.Advanced(self.AsMesh()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Mesh.Instance(self.AsMesh()), name)
	}
}
func init() {
	gdclass.Register("ImmediateMesh", func(ptr gd.Object) any {
		return [1]gdclass.ImmediateMesh{*(*gdclass.ImmediateMesh)(unsafe.Pointer(&ptr))}
	})
}
