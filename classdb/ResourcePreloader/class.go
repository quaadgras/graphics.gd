// Package ResourcePreloader provides methods for working with ResourcePreloader object instances.
package ResourcePreloader

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This node is used to preload sub-resources inside a scene, so when the scene is loaded, all the resources are ready to use and can be retrieved from the preloader. You can add the resources using the ResourcePreloader tab when the node is selected.
GDScript has a simplified [method @GDScript.preload] built-in method which can be used in most situations, leaving the use of [ResourcePreloader] for more advanced scenarios.
*/
type Instance [1]gdclass.ResourcePreloader
type Any interface {
	gd.IsClass
	AsResourcePreloader() Instance
}

/*
Adds a resource to the preloader with the given [param name]. If a resource with the given [param name] already exists, the new resource will be renamed to "[param name] N" where N is an incrementing number starting from 2.
*/
func (self Instance) AddResource(name string, resource [1]gdclass.Resource) {
	class(self).AddResource(gd.NewStringName(name), resource)
}

/*
Removes the resource associated to [param name] from the preloader.
*/
func (self Instance) RemoveResource(name string) {
	class(self).RemoveResource(gd.NewStringName(name))
}

/*
Renames a resource inside the preloader from [param name] to [param newname].
*/
func (self Instance) RenameResource(name string, newname string) {
	class(self).RenameResource(gd.NewStringName(name), gd.NewStringName(newname))
}

/*
Returns [code]true[/code] if the preloader contains a resource associated to [param name].
*/
func (self Instance) HasResource(name string) bool {
	return bool(class(self).HasResource(gd.NewStringName(name)))
}

/*
Returns the resource associated to [param name].
*/
func (self Instance) GetResource(name string) [1]gdclass.Resource {
	return [1]gdclass.Resource(class(self).GetResource(gd.NewStringName(name)))
}

/*
Returns the list of resources inside the preloader.
*/
func (self Instance) GetResourceList() []string {
	return []string(class(self).GetResourceList().Strings())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ResourcePreloader

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourcePreloader"))
	return Instance{*(*gdclass.ResourcePreloader)(unsafe.Pointer(&object))}
}

/*
Adds a resource to the preloader with the given [param name]. If a resource with the given [param name] already exists, the new resource will be renamed to "[param name] N" where N is an incrementing number starting from 2.
*/
//go:nosplit
func (self class) AddResource(name gd.StringName, resource [1]gdclass.Resource) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(resource[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_add_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the resource associated to [param name] from the preloader.
*/
//go:nosplit
func (self class) RemoveResource(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_remove_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Renames a resource inside the preloader from [param name] to [param newname].
*/
//go:nosplit
func (self class) RenameResource(name gd.StringName, newname gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(newname))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_rename_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the preloader contains a resource associated to [param name].
*/
//go:nosplit
func (self class) HasResource(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_has_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the resource associated to [param name].
*/
//go:nosplit
func (self class) GetResource(name gd.StringName) [1]gdclass.Resource {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_get_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the list of resources inside the preloader.
*/
//go:nosplit
func (self class) GetResourceList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_get_resource_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsResourcePreloader() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResourcePreloader() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced            { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance         { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {
	gdclass.Register("ResourcePreloader", func(ptr gd.Object) any {
		return [1]gdclass.ResourcePreloader{*(*gdclass.ResourcePreloader)(unsafe.Pointer(&ptr))}
	})
}
