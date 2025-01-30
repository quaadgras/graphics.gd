// Package MeshLibrary provides methods for working with MeshLibrary object instances.
package MeshLibrary

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
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
import "graphics.gd/variant/Transform3D"

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
A library of meshes. Contains a list of [Mesh] resources, each with a name and ID. Each item can also include collision and navigation shapes. This resource is used in [GridMap].
*/
type Instance [1]gdclass.MeshLibrary

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMeshLibrary() Instance
}

/*
Creates a new item in the library with the given ID.
You can get an unused ID from [method get_last_unused_item_id].
*/
func (self Instance) CreateItem(id int) { //gd:MeshLibrary.create_item
	class(self).CreateItem(int64(id))
}

/*
Sets the item's name.
This name is shown in the editor. It can also be used to look up the item later using [method find_item_by_name].
*/
func (self Instance) SetItemName(id int, name string) { //gd:MeshLibrary.set_item_name
	class(self).SetItemName(int64(id), String.New(name))
}

/*
Sets the item's mesh.
*/
func (self Instance) SetItemMesh(id int, mesh [1]gdclass.Mesh) { //gd:MeshLibrary.set_item_mesh
	class(self).SetItemMesh(int64(id), mesh)
}

/*
Sets the transform to apply to the item's mesh.
*/
func (self Instance) SetItemMeshTransform(id int, mesh_transform Transform3D.BasisOrigin) { //gd:MeshLibrary.set_item_mesh_transform
	class(self).SetItemMeshTransform(int64(id), Transform3D.BasisOrigin(mesh_transform))
}

/*
Sets the item's navigation mesh.
*/
func (self Instance) SetItemNavigationMesh(id int, navigation_mesh [1]gdclass.NavigationMesh) { //gd:MeshLibrary.set_item_navigation_mesh
	class(self).SetItemNavigationMesh(int64(id), navigation_mesh)
}

/*
Sets the transform to apply to the item's navigation mesh.
*/
func (self Instance) SetItemNavigationMeshTransform(id int, navigation_mesh Transform3D.BasisOrigin) { //gd:MeshLibrary.set_item_navigation_mesh_transform
	class(self).SetItemNavigationMeshTransform(int64(id), Transform3D.BasisOrigin(navigation_mesh))
}

/*
Sets the item's navigation layers bitmask.
*/
func (self Instance) SetItemNavigationLayers(id int, navigation_layers int) { //gd:MeshLibrary.set_item_navigation_layers
	class(self).SetItemNavigationLayers(int64(id), int64(navigation_layers))
}

/*
Sets an item's collision shapes.
The array should consist of [Shape3D] objects, each followed by a [Transform3D] that will be applied to it. For shapes that should not have a transform, use [constant Transform3D.IDENTITY].
*/
func (self Instance) SetItemShapes(id int, shapes []any) { //gd:MeshLibrary.set_item_shapes
	class(self).SetItemShapes(int64(id), gd.EngineArrayFromSlice(shapes))
}

/*
Sets a texture to use as the item's preview icon in the editor.
*/
func (self Instance) SetItemPreview(id int, texture [1]gdclass.Texture2D) { //gd:MeshLibrary.set_item_preview
	class(self).SetItemPreview(int64(id), texture)
}

/*
Returns the item's name.
*/
func (self Instance) GetItemName(id int) string { //gd:MeshLibrary.get_item_name
	return string(class(self).GetItemName(int64(id)).String())
}

/*
Returns the item's mesh.
*/
func (self Instance) GetItemMesh(id int) [1]gdclass.Mesh { //gd:MeshLibrary.get_item_mesh
	return [1]gdclass.Mesh(class(self).GetItemMesh(int64(id)))
}

/*
Returns the transform applied to the item's mesh.
*/
func (self Instance) GetItemMeshTransform(id int) Transform3D.BasisOrigin { //gd:MeshLibrary.get_item_mesh_transform
	return Transform3D.BasisOrigin(class(self).GetItemMeshTransform(int64(id)))
}

/*
Returns the item's navigation mesh.
*/
func (self Instance) GetItemNavigationMesh(id int) [1]gdclass.NavigationMesh { //gd:MeshLibrary.get_item_navigation_mesh
	return [1]gdclass.NavigationMesh(class(self).GetItemNavigationMesh(int64(id)))
}

/*
Returns the transform applied to the item's navigation mesh.
*/
func (self Instance) GetItemNavigationMeshTransform(id int) Transform3D.BasisOrigin { //gd:MeshLibrary.get_item_navigation_mesh_transform
	return Transform3D.BasisOrigin(class(self).GetItemNavigationMeshTransform(int64(id)))
}

/*
Returns the item's navigation layers bitmask.
*/
func (self Instance) GetItemNavigationLayers(id int) int { //gd:MeshLibrary.get_item_navigation_layers
	return int(int(class(self).GetItemNavigationLayers(int64(id))))
}

/*
Returns an item's collision shapes.
The array consists of each [Shape3D] followed by its [Transform3D].
*/
func (self Instance) GetItemShapes(id int) []any { //gd:MeshLibrary.get_item_shapes
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetItemShapes(int64(id)))))
}

/*
When running in the editor, returns a generated item preview (a 3D rendering in isometric perspective). When used in a running project, returns the manually-defined item preview which can be set using [method set_item_preview]. Returns an empty [Texture2D] if no preview was manually set in a running project.
*/
func (self Instance) GetItemPreview(id int) [1]gdclass.Texture2D { //gd:MeshLibrary.get_item_preview
	return [1]gdclass.Texture2D(class(self).GetItemPreview(int64(id)))
}

/*
Removes the item.
*/
func (self Instance) RemoveItem(id int) { //gd:MeshLibrary.remove_item
	class(self).RemoveItem(int64(id))
}

/*
Returns the first item with the given name, or [code]-1[/code] if no item is found.
*/
func (self Instance) FindItemByName(name string) int { //gd:MeshLibrary.find_item_by_name
	return int(int(class(self).FindItemByName(String.New(name))))
}

/*
Clears the library.
*/
func (self Instance) Clear() { //gd:MeshLibrary.clear
	class(self).Clear()
}

/*
Returns the list of item IDs in use.
*/
func (self Instance) GetItemList() []int32 { //gd:MeshLibrary.get_item_list
	return []int32(slices.Collect(class(self).GetItemList().Values()))
}

/*
Gets an unused ID for a new item.
*/
func (self Instance) GetLastUnusedItemId() int { //gd:MeshLibrary.get_last_unused_item_id
	return int(int(class(self).GetLastUnusedItemId()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MeshLibrary

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MeshLibrary"))
	casted := Instance{*(*gdclass.MeshLibrary)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates a new item in the library with the given ID.
You can get an unused ID from [method get_last_unused_item_id].
*/
//go:nosplit
func (self class) CreateItem(id int64) { //gd:MeshLibrary.create_item
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_create_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the item's name.
This name is shown in the editor. It can also be used to look up the item later using [method find_item_by_name].
*/
//go:nosplit
func (self class) SetItemName(id int64, name String.Readable) { //gd:MeshLibrary.set_item_name
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the item's mesh.
*/
//go:nosplit
func (self class) SetItemMesh(id int64, mesh [1]gdclass.Mesh) { //gd:MeshLibrary.set_item_mesh
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the transform to apply to the item's mesh.
*/
//go:nosplit
func (self class) SetItemMeshTransform(id int64, mesh_transform Transform3D.BasisOrigin) { //gd:MeshLibrary.set_item_mesh_transform
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mesh_transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_mesh_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the item's navigation mesh.
*/
//go:nosplit
func (self class) SetItemNavigationMesh(id int64, navigation_mesh [1]gdclass.NavigationMesh) { //gd:MeshLibrary.set_item_navigation_mesh
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the transform to apply to the item's navigation mesh.
*/
//go:nosplit
func (self class) SetItemNavigationMeshTransform(id int64, navigation_mesh Transform3D.BasisOrigin) { //gd:MeshLibrary.set_item_navigation_mesh_transform
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, navigation_mesh)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_navigation_mesh_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the item's navigation layers bitmask.
*/
//go:nosplit
func (self class) SetItemNavigationLayers(id int64, navigation_layers int64) { //gd:MeshLibrary.set_item_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets an item's collision shapes.
The array should consist of [Shape3D] objects, each followed by a [Transform3D] that will be applied to it. For shapes that should not have a transform, use [constant Transform3D.IDENTITY].
*/
//go:nosplit
func (self class) SetItemShapes(id int64, shapes Array.Any) { //gd:MeshLibrary.set_item_shapes
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(shapes)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_shapes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets a texture to use as the item's preview icon in the editor.
*/
//go:nosplit
func (self class) SetItemPreview(id int64, texture [1]gdclass.Texture2D) { //gd:MeshLibrary.set_item_preview
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_preview, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the item's name.
*/
//go:nosplit
func (self class) GetItemName(id int64) String.Readable { //gd:MeshLibrary.get_item_name
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the item's mesh.
*/
//go:nosplit
func (self class) GetItemMesh(id int64) [1]gdclass.Mesh { //gd:MeshLibrary.get_item_mesh
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Mesh{gd.PointerWithOwnershipTransferredToGo[gdclass.Mesh](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the transform applied to the item's mesh.
*/
//go:nosplit
func (self class) GetItemMeshTransform(id int64) Transform3D.BasisOrigin { //gd:MeshLibrary.get_item_mesh_transform
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_mesh_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the item's navigation mesh.
*/
//go:nosplit
func (self class) GetItemNavigationMesh(id int64) [1]gdclass.NavigationMesh { //gd:MeshLibrary.get_item_navigation_mesh
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.NavigationMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.NavigationMesh](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the transform applied to the item's navigation mesh.
*/
//go:nosplit
func (self class) GetItemNavigationMeshTransform(id int64) Transform3D.BasisOrigin { //gd:MeshLibrary.get_item_navigation_mesh_transform
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_navigation_mesh_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the item's navigation layers bitmask.
*/
//go:nosplit
func (self class) GetItemNavigationLayers(id int64) int64 { //gd:MeshLibrary.get_item_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an item's collision shapes.
The array consists of each [Shape3D] followed by its [Transform3D].
*/
//go:nosplit
func (self class) GetItemShapes(id int64) Array.Any { //gd:MeshLibrary.get_item_shapes
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_shapes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
When running in the editor, returns a generated item preview (a 3D rendering in isometric perspective). When used in a running project, returns the manually-defined item preview which can be set using [method set_item_preview]. Returns an empty [Texture2D] if no preview was manually set in a running project.
*/
//go:nosplit
func (self class) GetItemPreview(id int64) [1]gdclass.Texture2D { //gd:MeshLibrary.get_item_preview
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_preview, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Removes the item.
*/
//go:nosplit
func (self class) RemoveItem(id int64) { //gd:MeshLibrary.remove_item
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the first item with the given name, or [code]-1[/code] if no item is found.
*/
//go:nosplit
func (self class) FindItemByName(name String.Readable) int64 { //gd:MeshLibrary.find_item_by_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_find_item_by_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears the library.
*/
//go:nosplit
func (self class) Clear() { //gd:MeshLibrary.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the list of item IDs in use.
*/
//go:nosplit
func (self class) GetItemList() Packed.Array[int32] { //gd:MeshLibrary.get_item_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Gets an unused ID for a new item.
*/
//go:nosplit
func (self class) GetLastUnusedItemId() int64 { //gd:MeshLibrary.get_last_unused_item_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_last_unused_item_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsMeshLibrary() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMeshLibrary() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("MeshLibrary", func(ptr gd.Object) any { return [1]gdclass.MeshLibrary{*(*gdclass.MeshLibrary)(unsafe.Pointer(&ptr))} })
}
