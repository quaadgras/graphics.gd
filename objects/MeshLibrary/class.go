package MeshLibrary

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Transform3D"
import "grow.graphics/gd/variant/Array"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A library of meshes. Contains a list of [Mesh] resources, each with a name and ID. Each item can also include collision and navigation shapes. This resource is used in [GridMap].
*/
type Instance [1]classdb.MeshLibrary

/*
Creates a new item in the library with the given ID.
You can get an unused ID from [method get_last_unused_item_id].
*/
func (self Instance) CreateItem(id int) {
	class(self).CreateItem(gd.Int(id))
}

/*
Sets the item's name.
This name is shown in the editor. It can also be used to look up the item later using [method find_item_by_name].
*/
func (self Instance) SetItemName(id int, name string) {
	class(self).SetItemName(gd.Int(id), gd.NewString(name))
}

/*
Sets the item's mesh.
*/
func (self Instance) SetItemMesh(id int, mesh objects.Mesh) {
	class(self).SetItemMesh(gd.Int(id), mesh)
}

/*
Sets the transform to apply to the item's mesh.
*/
func (self Instance) SetItemMeshTransform(id int, mesh_transform Transform3D.BasisOrigin) {
	class(self).SetItemMeshTransform(gd.Int(id), gd.Transform3D(mesh_transform))
}

/*
Sets the item's navigation mesh.
*/
func (self Instance) SetItemNavigationMesh(id int, navigation_mesh objects.NavigationMesh) {
	class(self).SetItemNavigationMesh(gd.Int(id), navigation_mesh)
}

/*
Sets the transform to apply to the item's navigation mesh.
*/
func (self Instance) SetItemNavigationMeshTransform(id int, navigation_mesh Transform3D.BasisOrigin) {
	class(self).SetItemNavigationMeshTransform(gd.Int(id), gd.Transform3D(navigation_mesh))
}

/*
Sets the item's navigation layers bitmask.
*/
func (self Instance) SetItemNavigationLayers(id int, navigation_layers int) {
	class(self).SetItemNavigationLayers(gd.Int(id), gd.Int(navigation_layers))
}

/*
Sets an item's collision shapes.
The array should consist of [Shape3D] objects, each followed by a [Transform3D] that will be applied to it. For shapes that should not have a transform, use [constant Transform3D.IDENTITY].
*/
func (self Instance) SetItemShapes(id int, shapes Array.Any) {
	class(self).SetItemShapes(gd.Int(id), shapes)
}

/*
Sets a texture to use as the item's preview icon in the editor.
*/
func (self Instance) SetItemPreview(id int, texture objects.Texture2D) {
	class(self).SetItemPreview(gd.Int(id), texture)
}

/*
Returns the item's name.
*/
func (self Instance) GetItemName(id int) string {
	return string(class(self).GetItemName(gd.Int(id)).String())
}

/*
Returns the item's mesh.
*/
func (self Instance) GetItemMesh(id int) objects.Mesh {
	return objects.Mesh(class(self).GetItemMesh(gd.Int(id)))
}

/*
Returns the transform applied to the item's mesh.
*/
func (self Instance) GetItemMeshTransform(id int) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetItemMeshTransform(gd.Int(id)))
}

/*
Returns the item's navigation mesh.
*/
func (self Instance) GetItemNavigationMesh(id int) objects.NavigationMesh {
	return objects.NavigationMesh(class(self).GetItemNavigationMesh(gd.Int(id)))
}

/*
Returns the transform applied to the item's navigation mesh.
*/
func (self Instance) GetItemNavigationMeshTransform(id int) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetItemNavigationMeshTransform(gd.Int(id)))
}

/*
Returns the item's navigation layers bitmask.
*/
func (self Instance) GetItemNavigationLayers(id int) int {
	return int(int(class(self).GetItemNavigationLayers(gd.Int(id))))
}

/*
Returns an item's collision shapes.
The array consists of each [Shape3D] followed by its [Transform3D].
*/
func (self Instance) GetItemShapes(id int) Array.Any {
	return Array.Any(class(self).GetItemShapes(gd.Int(id)))
}

/*
When running in the editor, returns a generated item preview (a 3D rendering in isometric perspective). When used in a running project, returns the manually-defined item preview which can be set using [method set_item_preview]. Returns an empty [Texture2D] if no preview was manually set in a running project.
*/
func (self Instance) GetItemPreview(id int) objects.Texture2D {
	return objects.Texture2D(class(self).GetItemPreview(gd.Int(id)))
}

/*
Removes the item.
*/
func (self Instance) RemoveItem(id int) {
	class(self).RemoveItem(gd.Int(id))
}

/*
Returns the first item with the given name, or [code]-1[/code] if no item is found.
*/
func (self Instance) FindItemByName(name string) int {
	return int(int(class(self).FindItemByName(gd.NewString(name))))
}

/*
Clears the library.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Returns the list of item IDs in use.
*/
func (self Instance) GetItemList() []int32 {
	return []int32(class(self).GetItemList().AsSlice())
}

/*
Gets an unused ID for a new item.
*/
func (self Instance) GetLastUnusedItemId() int {
	return int(int(class(self).GetLastUnusedItemId()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MeshLibrary

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MeshLibrary"))
	return Instance{classdb.MeshLibrary(object)}
}

/*
Creates a new item in the library with the given ID.
You can get an unused ID from [method get_last_unused_item_id].
*/
//go:nosplit
func (self class) CreateItem(id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_create_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the item's name.
This name is shown in the editor. It can also be used to look up the item later using [method find_item_by_name].
*/
//go:nosplit
func (self class) SetItemName(id gd.Int, name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the item's mesh.
*/
//go:nosplit
func (self class) SetItemMesh(id gd.Int, mesh objects.Mesh) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the transform to apply to the item's mesh.
*/
//go:nosplit
func (self class) SetItemMeshTransform(id gd.Int, mesh_transform gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mesh_transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_mesh_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the item's navigation mesh.
*/
//go:nosplit
func (self class) SetItemNavigationMesh(id gd.Int, navigation_mesh objects.NavigationMesh) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the transform to apply to the item's navigation mesh.
*/
//go:nosplit
func (self class) SetItemNavigationMeshTransform(id gd.Int, navigation_mesh gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, navigation_mesh)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_navigation_mesh_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the item's navigation layers bitmask.
*/
//go:nosplit
func (self class) SetItemNavigationLayers(id gd.Int, navigation_layers gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets an item's collision shapes.
The array should consist of [Shape3D] objects, each followed by a [Transform3D] that will be applied to it. For shapes that should not have a transform, use [constant Transform3D.IDENTITY].
*/
//go:nosplit
func (self class) SetItemShapes(id gd.Int, shapes gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(shapes))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets a texture to use as the item's preview icon in the editor.
*/
//go:nosplit
func (self class) SetItemPreview(id gd.Int, texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the item's name.
*/
//go:nosplit
func (self class) GetItemName(id gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the item's mesh.
*/
//go:nosplit
func (self class) GetItemMesh(id gd.Int) objects.Mesh {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Mesh{classdb.Mesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the transform applied to the item's mesh.
*/
//go:nosplit
func (self class) GetItemMeshTransform(id gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_mesh_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the item's navigation mesh.
*/
//go:nosplit
func (self class) GetItemNavigationMesh(id gd.Int) objects.NavigationMesh {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.NavigationMesh{classdb.NavigationMesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the transform applied to the item's navigation mesh.
*/
//go:nosplit
func (self class) GetItemNavigationMeshTransform(id gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_navigation_mesh_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the item's navigation layers bitmask.
*/
//go:nosplit
func (self class) GetItemNavigationLayers(id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an item's collision shapes.
The array consists of each [Shape3D] followed by its [Transform3D].
*/
//go:nosplit
func (self class) GetItemShapes(id gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
When running in the editor, returns a generated item preview (a 3D rendering in isometric perspective). When used in a running project, returns the manually-defined item preview which can be set using [method set_item_preview]. Returns an empty [Texture2D] if no preview was manually set in a running project.
*/
//go:nosplit
func (self class) GetItemPreview(id gd.Int) objects.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Removes the item.
*/
//go:nosplit
func (self class) RemoveItem(id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the first item with the given name, or [code]-1[/code] if no item is found.
*/
//go:nosplit
func (self class) FindItemByName(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_find_item_by_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears the library.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the list of item IDs in use.
*/
//go:nosplit
func (self class) GetItemList() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Gets an unused ID for a new item.
*/
//go:nosplit
func (self class) GetLastUnusedItemId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_last_unused_item_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("MeshLibrary", func(ptr gd.Object) any { return classdb.MeshLibrary(ptr) })
}
