// Package ResourcePreloader provides methods for working with ResourcePreloader object instances.
package ResourcePreloader

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/Node"

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
var _ = slices.Delete[[]struct{}, struct{}]

/*
This node is used to preload sub-resources inside a scene, so when the scene is loaded, all the resources are ready to use and can be retrieved from the preloader. You can add the resources using the ResourcePreloader tab when the node is selected.
GDScript has a simplified [method @GDScript.preload] built-in method which can be used in most situations, leaving the use of [ResourcePreloader] for more advanced scenarios.
*/
type Instance [1]gdclass.ResourcePreloader

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsResourcePreloader() Instance
}

/*
Adds a resource to the preloader with the given [param name]. If a resource with the given [param name] already exists, the new resource will be renamed to "[param name] N" where N is an incrementing number starting from 2.
*/
func (self Instance) AddResource(name string, resource [1]gdclass.Resource) { //gd:ResourcePreloader.add_resource
	class(self).AddResource(String.Name(String.New(name)), resource)
}

/*
Removes the resource associated to [param name] from the preloader.
*/
func (self Instance) RemoveResource(name string) { //gd:ResourcePreloader.remove_resource
	class(self).RemoveResource(String.Name(String.New(name)))
}

/*
Renames a resource inside the preloader from [param name] to [param newname].
*/
func (self Instance) RenameResource(name string, newname string) { //gd:ResourcePreloader.rename_resource
	class(self).RenameResource(String.Name(String.New(name)), String.Name(String.New(newname)))
}

/*
Returns [code]true[/code] if the preloader contains a resource associated to [param name].
*/
func (self Instance) HasResource(name string) bool { //gd:ResourcePreloader.has_resource
	return bool(class(self).HasResource(String.Name(String.New(name))))
}

/*
Returns the resource associated to [param name].
*/
func (self Instance) GetResource(name string) [1]gdclass.Resource { //gd:ResourcePreloader.get_resource
	return [1]gdclass.Resource(class(self).GetResource(String.Name(String.New(name))))
}

/*
Returns the list of resources inside the preloader.
*/
func (self Instance) GetResourceList() []string { //gd:ResourcePreloader.get_resource_list
	return []string(class(self).GetResourceList().Strings())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ResourcePreloader

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourcePreloader"))
	casted := Instance{*(*gdclass.ResourcePreloader)(unsafe.Pointer(&object))}
	return casted
}

/*
Adds a resource to the preloader with the given [param name]. If a resource with the given [param name] already exists, the new resource will be renamed to "[param name] N" where N is an incrementing number starting from 2.
*/
//go:nosplit
func (self class) AddResource(name String.Name, resource [1]gdclass.Resource) { //gd:ResourcePreloader.add_resource
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(resource[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_add_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the resource associated to [param name] from the preloader.
*/
//go:nosplit
func (self class) RemoveResource(name String.Name) { //gd:ResourcePreloader.remove_resource
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_remove_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Renames a resource inside the preloader from [param name] to [param newname].
*/
//go:nosplit
func (self class) RenameResource(name String.Name, newname String.Name) { //gd:ResourcePreloader.rename_resource
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(newname)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_rename_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the preloader contains a resource associated to [param name].
*/
//go:nosplit
func (self class) HasResource(name String.Name) bool { //gd:ResourcePreloader.has_resource
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_has_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the resource associated to [param name].
*/
//go:nosplit
func (self class) GetResource(name String.Name) [1]gdclass.Resource { //gd:ResourcePreloader.get_resource
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_get_resource, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the list of resources inside the preloader.
*/
//go:nosplit
func (self class) GetResourceList() Packed.Strings { //gd:ResourcePreloader.get_resource_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourcePreloader.Bind_get_resource_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
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
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("ResourcePreloader", func(ptr gd.Object) any {
		return [1]gdclass.ResourcePreloader{*(*gdclass.ResourcePreloader)(unsafe.Pointer(&ptr))}
	})
}
