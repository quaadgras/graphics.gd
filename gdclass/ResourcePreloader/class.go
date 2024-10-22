package ResourcePreloader

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This node is used to preload sub-resources inside a scene, so when the scene is loaded, all the resources are ready to use and can be retrieved from the preloader. You can add the resources using the ResourcePreloader tab when the node is selected.
GDScript has a simplified [method @GDScript.preload] built-in method which can be used in most situations, leaving the use of [ResourcePreloader] for more advanced scenarios.

*/
type Go [1]classdb.ResourcePreloader

/*
Adds a resource to the preloader with the given [param name]. If a resource with the given [param name] already exists, the new resource will be renamed to "[param name] N" where N is an incrementing number starting from 2.
*/
func (self Go) AddResource(name string, resource gdclass.Resource) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddResource(gc.StringName(name), resource)
}

/*
Removes the resource associated to [param name] from the preloader.
*/
func (self Go) RemoveResource(name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveResource(gc.StringName(name))
}

/*
Renames a resource inside the preloader from [param name] to [param newname].
*/
func (self Go) RenameResource(name string, newname string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RenameResource(gc.StringName(name), gc.StringName(newname))
}

/*
Returns [code]true[/code] if the preloader contains a resource associated to [param name].
*/
func (self Go) HasResource(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasResource(gc.StringName(name)))
}

/*
Returns the resource associated to [param name].
*/
func (self Go) GetResource(name string) gdclass.Resource {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Resource(class(self).GetResource(gc, gc.StringName(name)))
}

/*
Returns the list of resources inside the preloader.
*/
func (self Go) GetResourceList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetResourceList(gc).Strings(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ResourcePreloader
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ResourcePreloader"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Adds a resource to the preloader with the given [param name]. If a resource with the given [param name] already exists, the new resource will be renamed to "[param name] N" where N is an incrementing number starting from 2.
*/
//go:nosplit
func (self class) AddResource(name gd.StringName, resource gdclass.Resource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(resource[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourcePreloader.Bind_add_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the resource associated to [param name] from the preloader.
*/
//go:nosplit
func (self class) RemoveResource(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourcePreloader.Bind_remove_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Renames a resource inside the preloader from [param name] to [param newname].
*/
//go:nosplit
func (self class) RenameResource(name gd.StringName, newname gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(newname))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourcePreloader.Bind_rename_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the preloader contains a resource associated to [param name].
*/
//go:nosplit
func (self class) HasResource(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourcePreloader.Bind_has_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the resource associated to [param name].
*/
//go:nosplit
func (self class) GetResource(ctx gd.Lifetime, name gd.StringName) gdclass.Resource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourcePreloader.Bind_get_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Resource
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the list of resources inside the preloader.
*/
//go:nosplit
func (self class) GetResourceList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourcePreloader.Bind_get_resource_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsResourcePreloader() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourcePreloader() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("ResourcePreloader", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
