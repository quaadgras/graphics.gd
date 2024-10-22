package InstancePlaceholder

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
Turning on the option [b]Load As Placeholder[/b] for an instantiated scene in the editor causes it to be replaced by an [InstancePlaceholder] when running the game, this will not replace the node in the editor. This makes it possible to delay actually loading the scene until calling [method create_instance]. This is useful to avoid loading large scenes all at once by loading parts of it selectively.
The [InstancePlaceholder] does not have a transform. This causes any child nodes to be positioned relatively to the [Viewport] from point (0,0), rather than their parent as displayed in the editor. Replacing the placeholder with a scene with a transform will transform children relatively to their parent again.

*/
type Go [1]classdb.InstancePlaceholder

/*
Returns the list of properties that will be applied to the node when [method create_instance] is called.
If [param with_order] is [code]true[/code], a key named [code].order[/code] (note the leading period) is added to the dictionary. This [code].order[/code] key is an [Array] of [String] property names specifying the order in which properties will be applied (with index 0 being the first).
*/
func (self Go) GetStoredValues() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetStoredValues(gc, false))
}

/*
Call this method to actually load in the node. The created node will be placed as a sibling [i]above[/i] the [InstancePlaceholder] in the scene tree. The [Node]'s reference is also returned for convenience.
[b]Note:[/b] [method create_instance] is not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
func (self Go) CreateInstance() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Node(class(self).CreateInstance(gc, false, ([1]gdclass.PackedScene{}[0])))
}

/*
Gets the path to the [PackedScene] resource file that is loaded by default when calling [method create_instance]. Not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
func (self Go) GetInstancePath() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetInstancePath(gc).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InstancePlaceholder
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("InstancePlaceholder"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the list of properties that will be applied to the node when [method create_instance] is called.
If [param with_order] is [code]true[/code], a key named [code].order[/code] (note the leading period) is added to the dictionary. This [code].order[/code] key is an [Array] of [String] property names specifying the order in which properties will be applied (with index 0 being the first).
*/
//go:nosplit
func (self class) GetStoredValues(ctx gd.Lifetime, with_order bool) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, with_order)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InstancePlaceholder.Bind_get_stored_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Call this method to actually load in the node. The created node will be placed as a sibling [i]above[/i] the [InstancePlaceholder] in the scene tree. The [Node]'s reference is also returned for convenience.
[b]Note:[/b] [method create_instance] is not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
//go:nosplit
func (self class) CreateInstance(ctx gd.Lifetime, replace bool, custom_scene gdclass.PackedScene) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, replace)
	callframe.Arg(frame, mmm.Get(custom_scene[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InstancePlaceholder.Bind_create_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Gets the path to the [PackedScene] resource file that is loaded by default when calling [method create_instance]. Not thread-safe. Use [method Object.call_deferred] if calling from a thread.
*/
//go:nosplit
func (self class) GetInstancePath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InstancePlaceholder.Bind_get_instance_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsInstancePlaceholder() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInstancePlaceholder() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("InstancePlaceholder", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
