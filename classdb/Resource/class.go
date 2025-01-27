// Package Resource provides methods for working with Resource object instances.
package Resource

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
Resource is the base class for all Godot-specific resource types, serving primarily as data containers. Since they inherit from [RefCounted], resources are reference-counted and freed when no longer in use. They can also be nested within other resources, and saved on disk. [PackedScene], one of the most common [Object]s in a Godot project, is also a resource, uniquely capable of storing and instantiating the [Node]s it contains as many times as desired.
In GDScript, resources can loaded from disk by their [member resource_path] using [method @GDScript.load] or [method @GDScript.preload].
The engine keeps a global cache of all loaded resources, referenced by paths (see [method ResourceLoader.has_cached]). A resource will be cached when loaded for the first time and removed from cache once all references are released. When a resource is cached, subsequent loads using its path will return the cached reference.
[b]Note:[/b] In C#, resources will not be freed instantly after they are no longer in use. Instead, garbage collection will run periodically and will free resources that are no longer in use. This means that unused resources will remain in memory for a while before being removed.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=Resource)
*/
type Instance [1]gdclass.Resource

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsResource() Instance
}
type Interface interface {
	//Override this method to customize the newly duplicated resource created from [method PackedScene.instantiate], if the original's [member resource_local_to_scene] is set to [code]true[/code].
	//[b]Example:[/b] Set a random [code]damage[/code] value to every local resource from an instantiated scene.
	//[codeblock]
	//extends Resource
	//
	//var damage = 0
	//
	//func _setup_local_to_scene():
	//    damage = randi_range(10, 40)
	//[/codeblock]
	SetupLocalToScene()
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) SetupLocalToScene() { return }

/*
Override this method to customize the newly duplicated resource created from [method PackedScene.instantiate], if the original's [member resource_local_to_scene] is set to [code]true[/code].
[b]Example:[/b] Set a random [code]damage[/code] value to every local resource from an instantiated scene.
[codeblock]
extends Resource

var damage = 0

func _setup_local_to_scene():

	damage = randi_range(10, 40)

[/codeblock]
*/
func (Instance) _setup_local_to_scene(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Sets the [member resource_path] to [param path], potentially overriding an existing cache entry for this path. Further attempts to load an overridden resource by path will instead return this resource.
*/
func (self Instance) TakeOverPath(path string) { //gd:Resource.take_over_path
	class(self).TakeOverPath(String.New(path))
}

/*
Returns the [RID] of this resource (or an empty RID). Many resources (such as [Texture2D], [Mesh], and so on) are high-level abstractions of resources stored in a specialized server ([DisplayServer], [RenderingServer], etc.), so this function will return the original [RID].
*/
func (self Instance) GetRid() ID { //gd:Resource.get_rid
	return ID(class(self).GetRid())
}

/*
If [member resource_local_to_scene] is set to [code]true[/code] and the resource has been loaded from a [PackedScene] instantiation, returns the root [Node] of the scene where this resource is used. Otherwise, returns [code]null[/code].
*/
func (self Instance) GetLocalScene() [1]gdclass.Node { //gd:Resource.get_local_scene
	return [1]gdclass.Node(class(self).GetLocalScene())
}

/*
Calls [method _setup_local_to_scene]. If [member resource_local_to_scene] is set to [code]true[/code], this method is automatically called from [method PackedScene.instantiate] by the newly duplicated resource within the scene instance.
*/
func (self Instance) SetupLocalToScene() { //gd:Resource.setup_local_to_scene
	class(self).SetupLocalToScene()
}

/*
Generates a unique identifier for a resource to be contained inside a [PackedScene], based on the current date, time, and a random value. The returned string is only composed of letters ([code]a[/code] to [code]y[/code]) and numbers ([code]0[/code] to [code]8[/code]). See also [member resource_scene_unique_id].
*/
func GenerateSceneUniqueId() string { //gd:Resource.generate_scene_unique_id
	self := Instance{}
	return string(class(self).GenerateSceneUniqueId().String())
}

/*
Emits the [signal changed] signal. This method is called automatically for some built-in resources.
[b]Note:[/b] For custom resources, it's recommended to call this method whenever a meaningful change occurs, such as a modified property. This ensures that custom [Object]s depending on the resource are properly updated.
[codeblock]
var damage:

	set(new_value):
	    if damage != new_value:
	        damage = new_value
	        emit_changed()

[/codeblock]
*/
func (self Instance) EmitChanged() { //gd:Resource.emit_changed
	class(self).EmitChanged()
}

/*
Duplicates this resource, returning a new resource with its [code]export[/code]ed or [constant PROPERTY_USAGE_STORAGE] properties copied from the original.
If [param subresources] is [code]false[/code], a shallow copy is returned; nested resources within subresources are not duplicated and are shared with the original resource (with one exception; see below). If [param subresources] is [code]true[/code], a deep copy is returned; nested subresources will be duplicated and are not shared (with two exceptions; see below).
[param subresources] is usually respected, with the following exceptions:
- Subresource properties with the [constant PROPERTY_USAGE_ALWAYS_DUPLICATE] flag are always duplicated.
- Subresource properties with the [constant PROPERTY_USAGE_NEVER_DUPLICATE] flag are never duplicated.
- Subresources inside [Array] and [Dictionary] properties are never duplicated.
[b]Note:[/b] For custom resources, this method will fail if [method Object._init] has been defined with required parameters.
*/
func (self Instance) Duplicate() [1]gdclass.Resource { //gd:Resource.duplicate
	return [1]gdclass.Resource(class(self).Duplicate(false))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Resource

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Resource"))
	casted := Instance{*(*gdclass.Resource)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) ResourceLocalToScene() bool {
	return bool(class(self).IsLocalToScene())
}

func (self Instance) SetResourceLocalToScene(value bool) {
	class(self).SetLocalToScene(value)
}

func (self Instance) ResourcePath() string {
	return string(class(self).GetPath().String())
}

func (self Instance) SetResourcePath(value string) {
	class(self).SetPath(String.New(value))
}

func (self Instance) ResourceName() string {
	return string(class(self).GetName().String())
}

func (self Instance) SetResourceName(value string) {
	class(self).SetName(String.New(value))
}

func (self Instance) ResourceSceneUniqueId() string {
	return string(class(self).GetSceneUniqueId().String())
}

func (self Instance) SetResourceSceneUniqueId(value string) {
	class(self).SetSceneUniqueId(String.New(value))
}

/*
Override this method to customize the newly duplicated resource created from [method PackedScene.instantiate], if the original's [member resource_local_to_scene] is set to [code]true[/code].
[b]Example:[/b] Set a random [code]damage[/code] value to every local resource from an instantiated scene.
[codeblock]
extends Resource

var damage = 0

func _setup_local_to_scene():

	damage = randi_range(10, 40)

[/codeblock]
*/
func (class) _setup_local_to_scene(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

//go:nosplit
func (self class) SetPath(path String.Readable) { //gd:Resource.set_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_set_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [member resource_path] to [param path], potentially overriding an existing cache entry for this path. Further attempts to load an overridden resource by path will instead return this resource.
*/
//go:nosplit
func (self class) TakeOverPath(path String.Readable) { //gd:Resource.take_over_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_take_over_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPath() String.Readable { //gd:Resource.get_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetName(name String.Readable) { //gd:Resource.set_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_set_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetName() String.Readable { //gd:Resource.get_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the [RID] of this resource (or an empty RID). Many resources (such as [Texture2D], [Mesh], and so on) are high-level abstractions of resources stored in a specialized server ([DisplayServer], [RenderingServer], etc.), so this function will return the original [RID].
*/
//go:nosplit
func (self class) GetRid() gd.RID { //gd:Resource.get_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLocalToScene(enable bool) { //gd:Resource.set_local_to_scene
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_set_local_to_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsLocalToScene() bool { //gd:Resource.is_local_to_scene
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_is_local_to_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [member resource_local_to_scene] is set to [code]true[/code] and the resource has been loaded from a [PackedScene] instantiation, returns the root [Node] of the scene where this resource is used. Otherwise, returns [code]null[/code].
*/
//go:nosplit
func (self class) GetLocalScene() [1]gdclass.Node { //gd:Resource.get_local_scene
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_get_local_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Calls [method _setup_local_to_scene]. If [member resource_local_to_scene] is set to [code]true[/code], this method is automatically called from [method PackedScene.instantiate] by the newly duplicated resource within the scene instance.
*/
//go:nosplit
func (self class) SetupLocalToScene() { //gd:Resource.setup_local_to_scene
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_setup_local_to_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Generates a unique identifier for a resource to be contained inside a [PackedScene], based on the current date, time, and a random value. The returned string is only composed of letters ([code]a[/code] to [code]y[/code]) and numbers ([code]0[/code] to [code]8[/code]). See also [member resource_scene_unique_id].
*/
//go:nosplit
func (self class) GenerateSceneUniqueId() String.Readable { //gd:Resource.generate_scene_unique_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_generate_scene_unique_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSceneUniqueId(id String.Readable) { //gd:Resource.set_scene_unique_id
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(id)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_set_scene_unique_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSceneUniqueId() String.Readable { //gd:Resource.get_scene_unique_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_get_scene_unique_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Emits the [signal changed] signal. This method is called automatically for some built-in resources.
[b]Note:[/b] For custom resources, it's recommended to call this method whenever a meaningful change occurs, such as a modified property. This ensures that custom [Object]s depending on the resource are properly updated.
[codeblock]
var damage:
    set(new_value):
        if damage != new_value:
            damage = new_value
            emit_changed()
[/codeblock]
*/
//go:nosplit
func (self class) EmitChanged() { //gd:Resource.emit_changed
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_emit_changed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Duplicates this resource, returning a new resource with its [code]export[/code]ed or [constant PROPERTY_USAGE_STORAGE] properties copied from the original.
If [param subresources] is [code]false[/code], a shallow copy is returned; nested resources within subresources are not duplicated and are shared with the original resource (with one exception; see below). If [param subresources] is [code]true[/code], a deep copy is returned; nested subresources will be duplicated and are not shared (with two exceptions; see below).
[param subresources] is usually respected, with the following exceptions:
- Subresource properties with the [constant PROPERTY_USAGE_ALWAYS_DUPLICATE] flag are always duplicated.
- Subresource properties with the [constant PROPERTY_USAGE_NEVER_DUPLICATE] flag are never duplicated.
- Subresources inside [Array] and [Dictionary] properties are never duplicated.
[b]Note:[/b] For custom resources, this method will fail if [method Object._init] has been defined with required parameters.
*/
//go:nosplit
func (self class) Duplicate(subresources bool) [1]gdclass.Resource { //gd:Resource.duplicate
	var frame = callframe.New()
	callframe.Arg(frame, subresources)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Resource.Bind_duplicate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}
func (self Instance) OnChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSetupLocalToSceneRequested(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("setup_local_to_scene_requested"), gd.NewCallable(cb), 0)
}

func (self class) AsResource() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResource() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_setup_local_to_scene":
		return reflect.ValueOf(self._setup_local_to_scene)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_setup_local_to_scene":
		return reflect.ValueOf(self._setup_local_to_scene)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("Resource", func(ptr gd.Object) any { return [1]gdclass.Resource{*(*gdclass.Resource)(unsafe.Pointer(&ptr))} })
}
