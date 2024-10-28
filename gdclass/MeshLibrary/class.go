package MeshLibrary

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A library of meshes. Contains a list of [Mesh] resources, each with a name and ID. Each item can also include collision and navigation shapes. This resource is used in [GridMap].

*/
type Go [1]classdb.MeshLibrary

/*
Creates a new item in the library with the given ID.
You can get an unused ID from [method get_last_unused_item_id].
*/
func (self Go) CreateItem(id int) {
	class(self).CreateItem(gd.Int(id))
}

/*
Sets the item's name.
This name is shown in the editor. It can also be used to look up the item later using [method find_item_by_name].
*/
func (self Go) SetItemName(id int, name string) {
	class(self).SetItemName(gd.Int(id), gd.NewString(name))
}

/*
Sets the item's mesh.
*/
func (self Go) SetItemMesh(id int, mesh gdclass.Mesh) {
	class(self).SetItemMesh(gd.Int(id), mesh)
}

/*
Sets the transform to apply to the item's mesh.
*/
func (self Go) SetItemMeshTransform(id int, mesh_transform gd.Transform3D) {
	class(self).SetItemMeshTransform(gd.Int(id), mesh_transform)
}

/*
Sets the item's navigation mesh.
*/
func (self Go) SetItemNavigationMesh(id int, navigation_mesh gdclass.NavigationMesh) {
	class(self).SetItemNavigationMesh(gd.Int(id), navigation_mesh)
}

/*
Sets the transform to apply to the item's navigation mesh.
*/
func (self Go) SetItemNavigationMeshTransform(id int, navigation_mesh gd.Transform3D) {
	class(self).SetItemNavigationMeshTransform(gd.Int(id), navigation_mesh)
}

/*
Sets the item's navigation layers bitmask.
*/
func (self Go) SetItemNavigationLayers(id int, navigation_layers int) {
	class(self).SetItemNavigationLayers(gd.Int(id), gd.Int(navigation_layers))
}

/*
Sets an item's collision shapes.
The array should consist of [Shape3D] objects, each followed by a [Transform3D] that will be applied to it. For shapes that should not have a transform, use [constant Transform3D.IDENTITY].
*/
func (self Go) SetItemShapes(id int, shapes gd.Array) {
	class(self).SetItemShapes(gd.Int(id), shapes)
}

/*
Sets a texture to use as the item's preview icon in the editor.
*/
func (self Go) SetItemPreview(id int, texture gdclass.Texture2D) {
	class(self).SetItemPreview(gd.Int(id), texture)
}

/*
Returns the item's name.
*/
func (self Go) GetItemName(id int) string {
	return string(class(self).GetItemName(gd.Int(id)).String())
}

/*
Returns the item's mesh.
*/
func (self Go) GetItemMesh(id int) gdclass.Mesh {
	return gdclass.Mesh(class(self).GetItemMesh(gd.Int(id)))
}

/*
Returns the transform applied to the item's mesh.
*/
func (self Go) GetItemMeshTransform(id int) gd.Transform3D {
	return gd.Transform3D(class(self).GetItemMeshTransform(gd.Int(id)))
}

/*
Returns the item's navigation mesh.
*/
func (self Go) GetItemNavigationMesh(id int) gdclass.NavigationMesh {
	return gdclass.NavigationMesh(class(self).GetItemNavigationMesh(gd.Int(id)))
}

/*
Returns the transform applied to the item's navigation mesh.
*/
func (self Go) GetItemNavigationMeshTransform(id int) gd.Transform3D {
	return gd.Transform3D(class(self).GetItemNavigationMeshTransform(gd.Int(id)))
}

/*
Returns the item's navigation layers bitmask.
*/
func (self Go) GetItemNavigationLayers(id int) int {
	return int(int(class(self).GetItemNavigationLayers(gd.Int(id))))
}

/*
Returns an item's collision shapes.
The array consists of each [Shape3D] followed by its [Transform3D].
*/
func (self Go) GetItemShapes(id int) gd.Array {
	return gd.Array(class(self).GetItemShapes(gd.Int(id)))
}

/*
When running in the editor, returns a generated item preview (a 3D rendering in isometric perspective). When used in a running project, returns the manually-defined item preview which can be set using [method set_item_preview]. Returns an empty [Texture2D] if no preview was manually set in a running project.
*/
func (self Go) GetItemPreview(id int) gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetItemPreview(gd.Int(id)))
}

/*
Removes the item.
*/
func (self Go) RemoveItem(id int) {
	class(self).RemoveItem(gd.Int(id))
}

/*
Returns the first item with the given name, or [code]-1[/code] if no item is found.
*/
func (self Go) FindItemByName(name string) int {
	return int(int(class(self).FindItemByName(gd.NewString(name))))
}

/*
Clears the library.
*/
func (self Go) Clear() {
	class(self).Clear()
}

/*
Returns the list of item IDs in use.
*/
func (self Go) GetItemList() []int32 {
	return []int32(class(self).GetItemList().AsSlice())
}

/*
Gets an unused ID for a new item.
*/
func (self Go) GetLastUnusedItemId() int {
	return int(int(class(self).GetLastUnusedItemId()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MeshLibrary
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MeshLibrary"))
	return Go{classdb.MeshLibrary(object)}
}

/*
Creates a new item in the library with the given ID.
You can get an unused ID from [method get_last_unused_item_id].
*/
//go:nosplit
func (self class) CreateItem(id gd.Int)  {
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
func (self class) SetItemName(id gd.Int, name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the item's mesh.
*/
//go:nosplit
func (self class) SetItemMesh(id gd.Int, mesh gdclass.Mesh)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, discreet.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transform to apply to the item's mesh.
*/
//go:nosplit
func (self class) SetItemMeshTransform(id gd.Int, mesh_transform gd.Transform3D)  {
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
func (self class) SetItemNavigationMesh(id gd.Int, navigation_mesh gdclass.NavigationMesh)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, discreet.Get(navigation_mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transform to apply to the item's navigation mesh.
*/
//go:nosplit
func (self class) SetItemNavigationMeshTransform(id gd.Int, navigation_mesh gd.Transform3D)  {
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
func (self class) SetItemNavigationLayers(id gd.Int, navigation_layers gd.Int)  {
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
func (self class) SetItemShapes(id gd.Int, shapes gd.Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, discreet.Get(shapes))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_set_item_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a texture to use as the item's preview icon in the editor.
*/
//go:nosplit
func (self class) SetItemPreview(id gd.Int, texture gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, discreet.Get(texture[0])[0])
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
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the item's mesh.
*/
//go:nosplit
func (self class) GetItemMesh(id gd.Int) gdclass.Mesh {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Mesh{classdb.Mesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
func (self class) GetItemNavigationMesh(id gd.Int) gdclass.NavigationMesh {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.NavigationMesh{classdb.NavigationMesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
When running in the editor, returns a generated item preview (a 3D rendering in isometric perspective). When used in a running project, returns the manually-defined item preview which can be set using [method set_item_preview]. Returns an empty [Texture2D] if no preview was manually set in a running project.
*/
//go:nosplit
func (self class) GetItemPreview(id gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshLibrary.Bind_get_item_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Removes the item.
*/
//go:nosplit
func (self class) RemoveItem(id gd.Int)  {
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
	callframe.Arg(frame, discreet.Get(name))
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
func (self class) Clear()  {
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
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
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
func (self class) AsMeshLibrary() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMeshLibrary() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("MeshLibrary", func(ptr gd.Object) any { return classdb.MeshLibrary(ptr) })}
