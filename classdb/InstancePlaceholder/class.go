// Package InstancePlaceholder provides methods for working with InstancePlaceholder object instances.
package InstancePlaceholder

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Turning on the option [b]Load As Placeholder[/b] for an instantiated scene in the editor causes it to be replaced by an [InstancePlaceholder] when running the game, this will not replace the node in the editor. This makes it possible to delay actually loading the scene until calling [method create_instance]. This is useful to avoid loading large scenes all at once by loading parts of it selectively.
The [InstancePlaceholder] does not have a transform. This causes any child nodes to be positioned relatively to the [Viewport] from point (0,0), rather than their parent as displayed in the editor. Replacing the placeholder with a scene with a transform will transform children relatively to their parent again.
*/
type Instance [1]gdclass.InstancePlaceholder
type Any interface {
	gd.IsClass
	AsInstancePlaceholder() Instance
}

/*
Returns the list of properties that will be applied to the node when [method create_instance] is called.
If [param with_order] is [code]true[/code], a key named [code].order[/code] (note the leading period) is added to the dictionary. This [code].order[/code] key is an [Array] of [String] property names specifying the order in which properties will be applied (with index 0 being the first).
*/
func (self Instance) GetStoredValues() Dictionary.Any {
	return Dictionary.Any(class(self).GetStoredValues(false))
}

/*
Call this method to actually load in the node. The created node will be placed as a sibling [i]above[/i] the [InstancePlaceholder] in the scene tree. The [Node]'s reference is also returned for convenience.
[b]Note:[/b] [method create_instance] is not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
func (self Instance) CreateInstance() [1]gdclass.Node {
	return [1]gdclass.Node(class(self).CreateInstance(false, [1][1]gdclass.PackedScene{}[0]))
}

/*
Gets the path to the [PackedScene] resource file that is loaded by default when calling [method create_instance]. Not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
func (self Instance) GetInstancePath() string {
	return string(class(self).GetInstancePath().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InstancePlaceholder

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InstancePlaceholder"))
	return Instance{*(*gdclass.InstancePlaceholder)(unsafe.Pointer(&object))}
}

/*
Returns the list of properties that will be applied to the node when [method create_instance] is called.
If [param with_order] is [code]true[/code], a key named [code].order[/code] (note the leading period) is added to the dictionary. This [code].order[/code] key is an [Array] of [String] property names specifying the order in which properties will be applied (with index 0 being the first).
*/
//go:nosplit
func (self class) GetStoredValues(with_order bool) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, with_order)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InstancePlaceholder.Bind_get_stored_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Call this method to actually load in the node. The created node will be placed as a sibling [i]above[/i] the [InstancePlaceholder] in the scene tree. The [Node]'s reference is also returned for convenience.
[b]Note:[/b] [method create_instance] is not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
//go:nosplit
func (self class) CreateInstance(replace bool, custom_scene [1]gdclass.PackedScene) [1]gdclass.Node {
	var frame = callframe.New()
	callframe.Arg(frame, replace)
	callframe.Arg(frame, pointers.Get(custom_scene[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InstancePlaceholder.Bind_create_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Node{gd.PointerWithOwnershipTransferredToGo[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Gets the path to the [PackedScene] resource file that is loaded by default when calling [method create_instance]. Not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
//go:nosplit
func (self class) GetInstancePath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InstancePlaceholder.Bind_get_instance_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsInstancePlaceholder() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInstancePlaceholder() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced              { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance           { return *((*Node.Instance)(unsafe.Pointer(&self))) }

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
	gdclass.Register("InstancePlaceholder", func(ptr gd.Object) any {
		return [1]gdclass.InstancePlaceholder{*(*gdclass.InstancePlaceholder)(unsafe.Pointer(&ptr))}
	})
}
