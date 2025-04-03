// Package Mesh provides methods for working with Mesh object instances.
package Mesh

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/AABB"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2i"
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
Mesh is a type of [Resource] that contains vertex array-based geometry, divided in [i]surfaces[/i]. Each surface contains a completely separate array and a material used to draw it. Design wise, a mesh with multiple surfaces is preferred to a single surface, because objects created in 3D editing software commonly contain multiple materials. The maximum number of surfaces per mesh is [constant RenderingServer.MAX_MESH_SURFACES].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=Mesh)
*/
type Instance [1]gdclass.Mesh
type Expanded [1]gdclass.Mesh

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMesh() Instance
}
type Interface interface {
	//Virtual method to override the surface count for a custom class extending [Mesh].
	GetSurfaceCount() int
	//Virtual method to override the surface array length for a custom class extending [Mesh].
	SurfaceGetArrayLen(index int) int
	//Virtual method to override the surface array index length for a custom class extending [Mesh].
	SurfaceGetArrayIndexLen(index int) int
	//Virtual method to override the surface arrays for a custom class extending [Mesh].
	SurfaceGetArrays(index int) []any
	//Virtual method to override the blend shape arrays for a custom class extending [Mesh].
	SurfaceGetBlendShapeArrays(index int) [][]any
	//Virtual method to override the surface LODs for a custom class extending [Mesh].
	SurfaceGetLods(index int) map[any]any
	//Virtual method to override the surface format for a custom class extending [Mesh].
	SurfaceGetFormat(index int) int
	//Virtual method to override the surface primitive type for a custom class extending [Mesh].
	SurfaceGetPrimitiveType(index int) int
	//Virtual method to override the setting of a [param material] at the given [param index] for a custom class extending [Mesh].
	SurfaceSetMaterial(index int, material [1]gdclass.Material)
	//Virtual method to override the surface material for a custom class extending [Mesh].
	SurfaceGetMaterial(index int) [1]gdclass.Material
	//Virtual method to override the number of blend shapes for a custom class extending [Mesh].
	GetBlendShapeCount() int
	//Virtual method to override the retrieval of blend shape names for a custom class extending [Mesh].
	GetBlendShapeName(index int) string
	//Virtual method to override the names of blend shapes for a custom class extending [Mesh].
	SetBlendShapeName(index int, name string)
	//Virtual method to override the [AABB] for a custom class extending [Mesh].
	GetAabb() AABB.PositionSize
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetSurfaceCount() (_ int)                                   { return }
func (self implementation) SurfaceGetArrayLen(index int) (_ int)                       { return }
func (self implementation) SurfaceGetArrayIndexLen(index int) (_ int)                  { return }
func (self implementation) SurfaceGetArrays(index int) (_ []any)                       { return }
func (self implementation) SurfaceGetBlendShapeArrays(index int) (_ [][]any)           { return }
func (self implementation) SurfaceGetLods(index int) (_ map[any]any)                   { return }
func (self implementation) SurfaceGetFormat(index int) (_ int)                         { return }
func (self implementation) SurfaceGetPrimitiveType(index int) (_ int)                  { return }
func (self implementation) SurfaceSetMaterial(index int, material [1]gdclass.Material) { return }
func (self implementation) SurfaceGetMaterial(index int) (_ [1]gdclass.Material)       { return }
func (self implementation) GetBlendShapeCount() (_ int)                                { return }
func (self implementation) GetBlendShapeName(index int) (_ string)                     { return }
func (self implementation) SetBlendShapeName(index int, name string)                   { return }
func (self implementation) GetAabb() (_ AABB.PositionSize)                             { return }

/*
Virtual method to override the surface count for a custom class extending [Mesh].
*/
func (Instance) _get_surface_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Virtual method to override the surface array length for a custom class extending [Mesh].
*/
func (Instance) _surface_get_array_len(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Virtual method to override the surface array index length for a custom class extending [Mesh].
*/
func (Instance) _surface_get_array_index_len(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Virtual method to override the surface arrays for a custom class extending [Mesh].
*/
func (Instance) _surface_get_arrays(impl func(ptr unsafe.Pointer, index int) []any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(gd.InternalArray(gd.EngineArrayFromSlice(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the blend shape arrays for a custom class extending [Mesh].
*/
func (Instance) _surface_get_blend_shape_arrays(impl func(ptr unsafe.Pointer, index int) [][]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Array.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the surface LODs for a custom class extending [Mesh].
*/
func (Instance) _surface_get_lods(impl func(ptr unsafe.Pointer, index int) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the surface format for a custom class extending [Mesh].
*/
func (Instance) _surface_get_format(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Virtual method to override the surface primitive type for a custom class extending [Mesh].
*/
func (Instance) _surface_get_primitive_type(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Virtual method to override the setting of a [param material] at the given [param index] for a custom class extending [Mesh].
*/
func (Instance) _surface_set_material(impl func(ptr unsafe.Pointer, index int, material [1]gdclass.Material)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		var material = [1]gdclass.Material{pointers.New[gdclass.Material]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(material[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(index), material)
	}
}

/*
Virtual method to override the surface material for a custom class extending [Mesh].
*/
func (Instance) _surface_get_material(impl func(ptr unsafe.Pointer, index int) [1]gdclass.Material) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the number of blend shapes for a custom class extending [Mesh].
*/
func (Instance) _get_blend_shape_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Virtual method to override the retrieval of blend shape names for a custom class extending [Mesh].
*/
func (Instance) _get_blend_shape_name(impl func(ptr unsafe.Pointer, index int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(gd.InternalStringName(String.Name(String.New(ret))))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the names of blend shapes for a custom class extending [Mesh].
*/
func (Instance) _set_blend_shape_name(impl func(ptr unsafe.Pointer, index int, name string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		var name = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1)))))
		defer pointers.End(gd.InternalStringName(name))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(index), name.String())
	}
}

/*
Virtual method to override the [AABB] for a custom class extending [Mesh].
*/
func (Instance) _get_aabb(impl func(ptr unsafe.Pointer) AABB.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, AABB.PositionSize(ret))
	}
}

/*
Returns the smallest [AABB] enclosing this mesh in local space. Not affected by [code]custom_aabb[/code].
[b]Note:[/b] This is only implemented for [ArrayMesh] and [PrimitiveMesh].
*/
func (self Instance) GetAabb() AABB.PositionSize { //gd:Mesh.get_aabb
	return AABB.PositionSize(Advanced(self).GetAabb())
}

/*
Returns all the vertices that make up the faces of the mesh. Each three vertices represent one triangle.
*/
func (self Instance) GetFaces() []Vector3.XYZ { //gd:Mesh.get_faces
	return []Vector3.XYZ(slices.Collect(Advanced(self).GetFaces().Values()))
}

/*
Returns the number of surfaces that the [Mesh] holds. This is equivalent to [method MeshInstance3D.get_surface_override_material_count].
*/
func (self Instance) GetSurfaceCount() int { //gd:Mesh.get_surface_count
	return int(int(Advanced(self).GetSurfaceCount()))
}

/*
Returns the arrays for the vertices, normals, UVs, etc. that make up the requested surface (see [method ArrayMesh.add_surface_from_arrays]).
*/
func (self Instance) SurfaceGetArrays(surf_idx int) []any { //gd:Mesh.surface_get_arrays
	return []any(gd.ArrayAs[[]any](gd.InternalArray(Advanced(self).SurfaceGetArrays(int64(surf_idx)))))
}

/*
Returns the blend shape arrays for the requested surface.
*/
func (self Instance) SurfaceGetBlendShapeArrays(surf_idx int) [][]any { //gd:Mesh.surface_get_blend_shape_arrays
	return [][]any(gd.ArrayAs[[][]any](gd.InternalArray(Advanced(self).SurfaceGetBlendShapeArrays(int64(surf_idx)))))
}

/*
Sets a [Material] for a given surface. Surface will be rendered using this material.
[b]Note:[/b] This assigns the material within the [Mesh] resource, not the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties. To set the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, use [method MeshInstance3D.set_surface_override_material] instead.
*/
func (self Instance) SurfaceSetMaterial(surf_idx int, material [1]gdclass.Material) { //gd:Mesh.surface_set_material
	Advanced(self).SurfaceSetMaterial(int64(surf_idx), material)
}

/*
Returns a [Material] in a given surface. Surface is rendered using this material.
[b]Note:[/b] This returns the material within the [Mesh] resource, not the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties. To get the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, use [method MeshInstance3D.get_surface_override_material] instead.
*/
func (self Instance) SurfaceGetMaterial(surf_idx int) [1]gdclass.Material { //gd:Mesh.surface_get_material
	return [1]gdclass.Material(Advanced(self).SurfaceGetMaterial(int64(surf_idx)))
}

/*
Creates a placeholder version of this resource ([PlaceholderMesh]).
*/
func (self Instance) CreatePlaceholder() [1]gdclass.Resource { //gd:Mesh.create_placeholder
	return [1]gdclass.Resource(Advanced(self).CreatePlaceholder())
}

/*
Calculate a [ConcavePolygonShape3D] from the mesh.
*/
func (self Instance) CreateTrimeshShape() [1]gdclass.ConcavePolygonShape3D { //gd:Mesh.create_trimesh_shape
	return [1]gdclass.ConcavePolygonShape3D(Advanced(self).CreateTrimeshShape())
}

/*
Calculate a [ConvexPolygonShape3D] from the mesh.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
func (self Instance) CreateConvexShape() [1]gdclass.ConvexPolygonShape3D { //gd:Mesh.create_convex_shape
	return [1]gdclass.ConvexPolygonShape3D(Advanced(self).CreateConvexShape(true, false))
}

/*
Calculate a [ConvexPolygonShape3D] from the mesh.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
func (self Expanded) CreateConvexShape(clean bool, simplify bool) [1]gdclass.ConvexPolygonShape3D { //gd:Mesh.create_convex_shape
	return [1]gdclass.ConvexPolygonShape3D(Advanced(self).CreateConvexShape(clean, simplify))
}

/*
Calculate an outline mesh at a defined offset (margin) from the original mesh.
[b]Note:[/b] This method typically returns the vertices in reverse order (e.g. clockwise to counterclockwise).
*/
func (self Instance) CreateOutline(margin Float.X) [1]gdclass.Mesh { //gd:Mesh.create_outline
	return [1]gdclass.Mesh(Advanced(self).CreateOutline(float64(margin)))
}

/*
Generate a [TriangleMesh] from the mesh. Considers only surfaces using one of these primitive types: [constant PRIMITIVE_TRIANGLES], [constant PRIMITIVE_TRIANGLE_STRIP].
*/
func (self Instance) GenerateTriangleMesh() [1]gdclass.TriangleMesh { //gd:Mesh.generate_triangle_mesh
	return [1]gdclass.TriangleMesh(Advanced(self).GenerateTriangleMesh())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Mesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Mesh"))
	casted := Instance{*(*gdclass.Mesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) LightmapSizeHint() Vector2i.XY {
	return Vector2i.XY(class(self).GetLightmapSizeHint())
}

func (self Instance) SetLightmapSizeHint(value Vector2i.XY) {
	class(self).SetLightmapSizeHint(Vector2i.XY(value))
}

/*
Virtual method to override the surface count for a custom class extending [Mesh].
*/
func (class) _get_surface_count(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the surface array length for a custom class extending [Mesh].
*/
func (class) _surface_get_array_len(impl func(ptr unsafe.Pointer, index int64) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the surface array index length for a custom class extending [Mesh].
*/
func (class) _surface_get_array_index_len(impl func(ptr unsafe.Pointer, index int64) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the surface arrays for a custom class extending [Mesh].
*/
func (class) _surface_get_arrays(impl func(ptr unsafe.Pointer, index int64) Array.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the blend shape arrays for a custom class extending [Mesh].
*/
func (class) _surface_get_blend_shape_arrays(impl func(ptr unsafe.Pointer, index int64) Array.Contains[Array.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the surface LODs for a custom class extending [Mesh].
*/
func (class) _surface_get_lods(impl func(ptr unsafe.Pointer, index int64) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the surface format for a custom class extending [Mesh].
*/
func (class) _surface_get_format(impl func(ptr unsafe.Pointer, index int64) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the surface primitive type for a custom class extending [Mesh].
*/
func (class) _surface_get_primitive_type(impl func(ptr unsafe.Pointer, index int64) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the setting of a [param material] at the given [param index] for a custom class extending [Mesh].
*/
func (class) _surface_set_material(impl func(ptr unsafe.Pointer, index int64, material [1]gdclass.Material)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		var material = [1]gdclass.Material{pointers.New[gdclass.Material]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(material[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, index, material)
	}
}

/*
Virtual method to override the surface material for a custom class extending [Mesh].
*/
func (class) _surface_get_material(impl func(ptr unsafe.Pointer, index int64) [1]gdclass.Material) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the number of blend shapes for a custom class extending [Mesh].
*/
func (class) _get_blend_shape_count(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the retrieval of blend shape names for a custom class extending [Mesh].
*/
func (class) _get_blend_shape_name(impl func(ptr unsafe.Pointer, index int64) String.Name) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(gd.InternalStringName(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the names of blend shapes for a custom class extending [Mesh].
*/
func (class) _set_blend_shape_name(impl func(ptr unsafe.Pointer, index int64, name String.Name)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		var name = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1)))))
		defer pointers.End(gd.InternalStringName(name))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, index, name)
	}
}

/*
Virtual method to override the [AABB] for a custom class extending [Mesh].
*/
func (class) _get_aabb(impl func(ptr unsafe.Pointer) AABB.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) SetLightmapSizeHint(size Vector2i.XY) { //gd:Mesh.set_lightmap_size_hint
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_set_lightmap_size_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLightmapSizeHint() Vector2i.XY { //gd:Mesh.get_lightmap_size_hint
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_get_lightmap_size_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the smallest [AABB] enclosing this mesh in local space. Not affected by [code]custom_aabb[/code].
[b]Note:[/b] This is only implemented for [ArrayMesh] and [PrimitiveMesh].
*/
//go:nosplit
func (self class) GetAabb() AABB.PositionSize { //gd:Mesh.get_aabb
	var frame = callframe.New()
	var r_ret = callframe.Ret[AABB.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_get_aabb, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns all the vertices that make up the faces of the mesh. Each three vertices represent one triangle.
*/
//go:nosplit
func (self class) GetFaces() Packed.Array[Vector3.XYZ] { //gd:Mesh.get_faces
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_get_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the number of surfaces that the [Mesh] holds. This is equivalent to [method MeshInstance3D.get_surface_override_material_count].
*/
//go:nosplit
func (self class) GetSurfaceCount() int64 { //gd:Mesh.get_surface_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_get_surface_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the arrays for the vertices, normals, UVs, etc. that make up the requested surface (see [method ArrayMesh.add_surface_from_arrays]).
*/
//go:nosplit
func (self class) SurfaceGetArrays(surf_idx int64) Array.Any { //gd:Mesh.surface_get_arrays
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_surface_get_arrays, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the blend shape arrays for the requested surface.
*/
//go:nosplit
func (self class) SurfaceGetBlendShapeArrays(surf_idx int64) Array.Contains[Array.Any] { //gd:Mesh.surface_get_blend_shape_arrays
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_surface_get_blend_shape_arrays, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Array.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets a [Material] for a given surface. Surface will be rendered using this material.
[b]Note:[/b] This assigns the material within the [Mesh] resource, not the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties. To set the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, use [method MeshInstance3D.set_surface_override_material] instead.
*/
//go:nosplit
func (self class) SurfaceSetMaterial(surf_idx int64, material [1]gdclass.Material) { //gd:Mesh.surface_set_material
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_surface_set_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a [Material] in a given surface. Surface is rendered using this material.
[b]Note:[/b] This returns the material within the [Mesh] resource, not the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties. To get the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, use [method MeshInstance3D.get_surface_override_material] instead.
*/
//go:nosplit
func (self class) SurfaceGetMaterial(surf_idx int64) [1]gdclass.Material { //gd:Mesh.surface_get_material
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_surface_get_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a placeholder version of this resource ([PlaceholderMesh]).
*/
//go:nosplit
func (self class) CreatePlaceholder() [1]gdclass.Resource { //gd:Mesh.create_placeholder
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Calculate a [ConcavePolygonShape3D] from the mesh.
*/
//go:nosplit
func (self class) CreateTrimeshShape() [1]gdclass.ConcavePolygonShape3D { //gd:Mesh.create_trimesh_shape
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_create_trimesh_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ConcavePolygonShape3D{gd.PointerWithOwnershipTransferredToGo[gdclass.ConcavePolygonShape3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Calculate a [ConvexPolygonShape3D] from the mesh.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
//go:nosplit
func (self class) CreateConvexShape(clean bool, simplify bool) [1]gdclass.ConvexPolygonShape3D { //gd:Mesh.create_convex_shape
	var frame = callframe.New()
	callframe.Arg(frame, clean)
	callframe.Arg(frame, simplify)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_create_convex_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ConvexPolygonShape3D{gd.PointerWithOwnershipTransferredToGo[gdclass.ConvexPolygonShape3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Calculate an outline mesh at a defined offset (margin) from the original mesh.
[b]Note:[/b] This method typically returns the vertices in reverse order (e.g. clockwise to counterclockwise).
*/
//go:nosplit
func (self class) CreateOutline(margin float64) [1]gdclass.Mesh { //gd:Mesh.create_outline
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_create_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Mesh{gd.PointerWithOwnershipTransferredToGo[gdclass.Mesh](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Generate a [TriangleMesh] from the mesh. Considers only surfaces using one of these primitive types: [constant PRIMITIVE_TRIANGLES], [constant PRIMITIVE_TRIANGLE_STRIP].
*/
//go:nosplit
func (self class) GenerateTriangleMesh() [1]gdclass.TriangleMesh { //gd:Mesh.generate_triangle_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_generate_triangle_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TriangleMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.TriangleMesh](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_get_surface_count":
		return reflect.ValueOf(self._get_surface_count)
	case "_surface_get_array_len":
		return reflect.ValueOf(self._surface_get_array_len)
	case "_surface_get_array_index_len":
		return reflect.ValueOf(self._surface_get_array_index_len)
	case "_surface_get_arrays":
		return reflect.ValueOf(self._surface_get_arrays)
	case "_surface_get_blend_shape_arrays":
		return reflect.ValueOf(self._surface_get_blend_shape_arrays)
	case "_surface_get_lods":
		return reflect.ValueOf(self._surface_get_lods)
	case "_surface_get_format":
		return reflect.ValueOf(self._surface_get_format)
	case "_surface_get_primitive_type":
		return reflect.ValueOf(self._surface_get_primitive_type)
	case "_surface_set_material":
		return reflect.ValueOf(self._surface_set_material)
	case "_surface_get_material":
		return reflect.ValueOf(self._surface_get_material)
	case "_get_blend_shape_count":
		return reflect.ValueOf(self._get_blend_shape_count)
	case "_get_blend_shape_name":
		return reflect.ValueOf(self._get_blend_shape_name)
	case "_set_blend_shape_name":
		return reflect.ValueOf(self._set_blend_shape_name)
	case "_get_aabb":
		return reflect.ValueOf(self._get_aabb)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_surface_count":
		return reflect.ValueOf(self._get_surface_count)
	case "_surface_get_array_len":
		return reflect.ValueOf(self._surface_get_array_len)
	case "_surface_get_array_index_len":
		return reflect.ValueOf(self._surface_get_array_index_len)
	case "_surface_get_arrays":
		return reflect.ValueOf(self._surface_get_arrays)
	case "_surface_get_blend_shape_arrays":
		return reflect.ValueOf(self._surface_get_blend_shape_arrays)
	case "_surface_get_lods":
		return reflect.ValueOf(self._surface_get_lods)
	case "_surface_get_format":
		return reflect.ValueOf(self._surface_get_format)
	case "_surface_get_primitive_type":
		return reflect.ValueOf(self._surface_get_primitive_type)
	case "_surface_set_material":
		return reflect.ValueOf(self._surface_set_material)
	case "_surface_get_material":
		return reflect.ValueOf(self._surface_get_material)
	case "_get_blend_shape_count":
		return reflect.ValueOf(self._get_blend_shape_count)
	case "_get_blend_shape_name":
		return reflect.ValueOf(self._get_blend_shape_name)
	case "_set_blend_shape_name":
		return reflect.ValueOf(self._set_blend_shape_name)
	case "_get_aabb":
		return reflect.ValueOf(self._get_aabb)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Mesh", func(ptr gd.Object) any { return [1]gdclass.Mesh{*(*gdclass.Mesh)(unsafe.Pointer(&ptr))} })
}

type PrimitiveType = gdclass.MeshPrimitiveType //gd:Mesh.PrimitiveType

const (
	/*Render array as points (one vertex equals one point).*/
	PrimitivePoints PrimitiveType = 0
	/*Render array as lines (every two vertices a line is created).*/
	PrimitiveLines PrimitiveType = 1
	/*Render array as line strip.*/
	PrimitiveLineStrip PrimitiveType = 2
	/*Render array as triangles (every three vertices a triangle is created).*/
	PrimitiveTriangles PrimitiveType = 3
	/*Render array as triangle strips.*/
	PrimitiveTriangleStrip PrimitiveType = 4
)

type ArrayType = gdclass.MeshArrayType //gd:Mesh.ArrayType

const (
	/*[PackedVector3Array], [PackedVector2Array], or [Array] of vertex positions.*/
	ArrayVertex ArrayType = 0
	/*[PackedVector3Array] of vertex normals.
	  [b]Note:[/b] The array has to consist of normal vectors, otherwise they will be normalized by the engine, potentially causing visual discrepancies.*/
	ArrayNormal ArrayType = 1
	/*[PackedFloat32Array] of vertex tangents. Each element in groups of 4 floats, first 3 floats determine the tangent, and the last the binormal direction as -1 or 1.*/
	ArrayTangent ArrayType = 2
	/*[PackedColorArray] of vertex colors.*/
	ArrayColor ArrayType = 3
	/*[PackedVector2Array] for UV coordinates.*/
	ArrayTexUv ArrayType = 4
	/*[PackedVector2Array] for second UV coordinates.*/
	ArrayTexUv2 ArrayType = 5
	/*Contains custom color channel 0. [PackedByteArray] if [code](format >> Mesh.ARRAY_FORMAT_CUSTOM0_SHIFT) & Mesh.ARRAY_FORMAT_CUSTOM_MASK[/code] is [constant ARRAY_CUSTOM_RGBA8_UNORM], [constant ARRAY_CUSTOM_RGBA8_SNORM], [constant ARRAY_CUSTOM_RG_HALF], or [constant ARRAY_CUSTOM_RGBA_HALF]. [PackedFloat32Array] otherwise.*/
	ArrayCustom0 ArrayType = 6
	/*Contains custom color channel 1. [PackedByteArray] if [code](format >> Mesh.ARRAY_FORMAT_CUSTOM1_SHIFT) & Mesh.ARRAY_FORMAT_CUSTOM_MASK[/code] is [constant ARRAY_CUSTOM_RGBA8_UNORM], [constant ARRAY_CUSTOM_RGBA8_SNORM], [constant ARRAY_CUSTOM_RG_HALF], or [constant ARRAY_CUSTOM_RGBA_HALF]. [PackedFloat32Array] otherwise.*/
	ArrayCustom1 ArrayType = 7
	/*Contains custom color channel 2. [PackedByteArray] if [code](format >> Mesh.ARRAY_FORMAT_CUSTOM2_SHIFT) & Mesh.ARRAY_FORMAT_CUSTOM_MASK[/code] is [constant ARRAY_CUSTOM_RGBA8_UNORM], [constant ARRAY_CUSTOM_RGBA8_SNORM], [constant ARRAY_CUSTOM_RG_HALF], or [constant ARRAY_CUSTOM_RGBA_HALF]. [PackedFloat32Array] otherwise.*/
	ArrayCustom2 ArrayType = 8
	/*Contains custom color channel 3. [PackedByteArray] if [code](format >> Mesh.ARRAY_FORMAT_CUSTOM3_SHIFT) & Mesh.ARRAY_FORMAT_CUSTOM_MASK[/code] is [constant ARRAY_CUSTOM_RGBA8_UNORM], [constant ARRAY_CUSTOM_RGBA8_SNORM], [constant ARRAY_CUSTOM_RG_HALF], or [constant ARRAY_CUSTOM_RGBA_HALF]. [PackedFloat32Array] otherwise.*/
	ArrayCustom3 ArrayType = 9
	/*[PackedFloat32Array] or [PackedInt32Array] of bone indices. Contains either 4 or 8 numbers per vertex depending on the presence of the [constant ARRAY_FLAG_USE_8_BONE_WEIGHTS] flag.*/
	ArrayBones ArrayType = 10
	/*[PackedFloat32Array] or [PackedFloat64Array] of bone weights in the range [code]0.0[/code] to [code]1.0[/code] (inclusive). Contains either 4 or 8 numbers per vertex depending on the presence of the [constant ARRAY_FLAG_USE_8_BONE_WEIGHTS] flag.*/
	ArrayWeights ArrayType = 11
	/*[PackedInt32Array] of integers used as indices referencing vertices, colors, normals, tangents, and textures. All of those arrays must have the same number of elements as the vertex array. No index can be beyond the vertex array size. When this index array is present, it puts the function into "index mode," where the index selects the [i]i[/i]'th vertex, normal, tangent, color, UV, etc. This means if you want to have different normals or colors along an edge, you have to duplicate the vertices.
	  For triangles, the index array is interpreted as triples, referring to the vertices of each triangle. For lines, the index array is in pairs indicating the start and end of each line.*/
	ArrayIndex ArrayType = 12
	/*Represents the size of the [enum ArrayType] enum.*/
	ArrayMax ArrayType = 13
)

type ArrayCustomFormat = gdclass.MeshArrayCustomFormat //gd:Mesh.ArrayCustomFormat

const (
	/*Indicates this custom channel contains unsigned normalized byte colors from 0 to 1, encoded as [PackedByteArray].*/
	ArrayCustomRgba8Unorm ArrayCustomFormat = 0
	/*Indicates this custom channel contains signed normalized byte colors from -1 to 1, encoded as [PackedByteArray].*/
	ArrayCustomRgba8Snorm ArrayCustomFormat = 1
	/*Indicates this custom channel contains half precision float colors, encoded as [PackedByteArray]. Only red and green channels are used.*/
	ArrayCustomRgHalf ArrayCustomFormat = 2
	/*Indicates this custom channel contains half precision float colors, encoded as [PackedByteArray].*/
	ArrayCustomRgbaHalf ArrayCustomFormat = 3
	/*Indicates this custom channel contains full float colors, in a [PackedFloat32Array]. Only the red channel is used.*/
	ArrayCustomRFloat ArrayCustomFormat = 4
	/*Indicates this custom channel contains full float colors, in a [PackedFloat32Array]. Only red and green channels are used.*/
	ArrayCustomRgFloat ArrayCustomFormat = 5
	/*Indicates this custom channel contains full float colors, in a [PackedFloat32Array]. Only red, green and blue channels are used.*/
	ArrayCustomRgbFloat ArrayCustomFormat = 6
	/*Indicates this custom channel contains full float colors, in a [PackedFloat32Array].*/
	ArrayCustomRgbaFloat ArrayCustomFormat = 7
	/*Represents the size of the [enum ArrayCustomFormat] enum.*/
	ArrayCustomMax ArrayCustomFormat = 8
)

type ArrayFormat = gdclass.MeshArrayFormat //gd:Mesh.ArrayFormat

const (
	/*Mesh array contains vertices. All meshes require a vertex array so this should always be present.*/
	ArrayFormatVertex ArrayFormat = 1
	/*Mesh array contains normals.*/
	ArrayFormatNormal ArrayFormat = 2
	/*Mesh array contains tangents.*/
	ArrayFormatTangent ArrayFormat = 4
	/*Mesh array contains colors.*/
	ArrayFormatColor ArrayFormat = 8
	/*Mesh array contains UVs.*/
	ArrayFormatTexUv ArrayFormat = 16
	/*Mesh array contains second UV.*/
	ArrayFormatTexUv2 ArrayFormat = 32
	/*Mesh array contains custom channel index 0.*/
	ArrayFormatCustom0 ArrayFormat = 64
	/*Mesh array contains custom channel index 1.*/
	ArrayFormatCustom1 ArrayFormat = 128
	/*Mesh array contains custom channel index 2.*/
	ArrayFormatCustom2 ArrayFormat = 256
	/*Mesh array contains custom channel index 3.*/
	ArrayFormatCustom3 ArrayFormat = 512
	/*Mesh array contains bones.*/
	ArrayFormatBones ArrayFormat = 1024
	/*Mesh array contains bone weights.*/
	ArrayFormatWeights ArrayFormat = 2048
	/*Mesh array uses indices.*/
	ArrayFormatIndex ArrayFormat = 4096
	/*Mask of mesh channels permitted in blend shapes.*/
	ArrayFormatBlendShapeMask ArrayFormat = 7
	/*Shift of first custom channel.*/
	ArrayFormatCustomBase ArrayFormat = 13
	/*Number of format bits per custom channel. See [enum ArrayCustomFormat].*/
	ArrayFormatCustomBits ArrayFormat = 3
	/*Amount to shift [enum ArrayCustomFormat] for custom channel index 0.*/
	ArrayFormatCustom0Shift ArrayFormat = 13
	/*Amount to shift [enum ArrayCustomFormat] for custom channel index 1.*/
	ArrayFormatCustom1Shift ArrayFormat = 16
	/*Amount to shift [enum ArrayCustomFormat] for custom channel index 2.*/
	ArrayFormatCustom2Shift ArrayFormat = 19
	/*Amount to shift [enum ArrayCustomFormat] for custom channel index 3.*/
	ArrayFormatCustom3Shift ArrayFormat = 22
	/*Mask of custom format bits per custom channel. Must be shifted by one of the SHIFT constants. See [enum ArrayCustomFormat].*/
	ArrayFormatCustomMask ArrayFormat = 7
	/*Shift of first compress flag. Compress flags should be passed to [method ArrayMesh.add_surface_from_arrays] and [method SurfaceTool.commit].*/
	ArrayCompressFlagsBase ArrayFormat = 25
	/*Flag used to mark that the array contains 2D vertices.*/
	ArrayFlagUse2dVertices ArrayFormat = 33554432
	/*Flag indices that the mesh data will use [code]GL_DYNAMIC_DRAW[/code] on GLES. Unused on Vulkan.*/
	ArrayFlagUseDynamicUpdate ArrayFormat = 67108864
	/*Flag used to mark that the mesh contains up to 8 bone influences per vertex. This flag indicates that [constant ARRAY_BONES] and [constant ARRAY_WEIGHTS] elements will have double length.*/
	ArrayFlagUse8BoneWeights ArrayFormat = 134217728
	/*Flag used to mark that the mesh intentionally contains no vertex array.*/
	ArrayFlagUsesEmptyVertexArray ArrayFormat = 268435456
	/*Flag used to mark that a mesh is using compressed attributes (vertices, normals, tangents, UVs). When this form of compression is enabled, vertex positions will be packed into an RGBA16UNORM attribute and scaled in the vertex shader. The normal and tangent will be packed into an RG16UNORM representing an axis, and a 16-bit float stored in the A-channel of the vertex. UVs will use 16-bit normalized floats instead of full 32-bit signed floats. When using this compression mode you must use either vertices, normals, and tangents or only vertices. You cannot use normals without tangents. Importers will automatically enable this compression if they can.*/
	ArrayFlagCompressAttributes ArrayFormat = 536870912
)

type BlendShapeMode = gdclass.MeshBlendShapeMode //gd:Mesh.BlendShapeMode

const (
	/*Blend shapes are normalized.*/
	BlendShapeModeNormalized BlendShapeMode = 0
	/*Blend shapes are relative to base weight.*/
	BlendShapeModeRelative BlendShapeMode = 1
)
