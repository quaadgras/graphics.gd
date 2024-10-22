package Container

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for all GUI containers. A [Container] automatically arranges its child controls in a certain way. This class can be inherited to make custom container types.
	// Container methods that can be overridden by a [Class] that extends it.
	type Container interface {
		//Implement to return a list of allowed horizontal [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
		//[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
		GetAllowedSizeFlagsHorizontal() []int32
		//Implement to return a list of allowed vertical [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
		//[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
		GetAllowedSizeFlagsVertical() []int32
	}

*/
type Go [1]classdb.Container

/*
Implement to return a list of allowed horizontal [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
*/
func (Go) _get_allowed_size_flags_horizontal(impl func(ptr unsafe.Pointer) []int32, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedInt32Slice(ret)))
		gc.End()
	}
}

/*
Implement to return a list of allowed vertical [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
*/
func (Go) _get_allowed_size_flags_vertical(impl func(ptr unsafe.Pointer) []int32, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedInt32Slice(ret)))
		gc.End()
	}
}

/*
Queue resort of the contained children. This is called automatically anyway, but can be called upon request.
*/
func (self Go) QueueSort() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).QueueSort()
}

/*
Fit a child control in a given rect. This is mainly a helper for creating custom container classes.
*/
func (self Go) FitChildInRect(child gdclass.Control, rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FitChildInRect(child, rect)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Container
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Container"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Implement to return a list of allowed horizontal [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
*/
func (class) _get_allowed_size_flags_horizontal(impl func(ptr unsafe.Pointer) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Implement to return a list of allowed vertical [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
*/
func (class) _get_allowed_size_flags_vertical(impl func(ptr unsafe.Pointer) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Queue resort of the contained children. This is called automatically anyway, but can be called upon request.
*/
//go:nosplit
func (self class) QueueSort()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Container.Bind_queue_sort, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Fit a child control in a given rect. This is mainly a helper for creating custom container classes.
*/
//go:nosplit
func (self class) FitChildInRect(child gdclass.Control, rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(child[0].AsPointer())[0])
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Container.Bind_fit_child_in_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnPreSortChildren(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("pre_sort_children"), gc.Callable(cb), 0)
}


func (self Go) OnSortChildren(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("sort_children"), gc.Callable(cb), 0)
}


func (self class) AsContainer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_allowed_size_flags_horizontal": return reflect.ValueOf(self._get_allowed_size_flags_horizontal);
	case "_get_allowed_size_flags_vertical": return reflect.ValueOf(self._get_allowed_size_flags_vertical);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_allowed_size_flags_horizontal": return reflect.ValueOf(self._get_allowed_size_flags_horizontal);
	case "_get_allowed_size_flags_vertical": return reflect.ValueOf(self._get_allowed_size_flags_vertical);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("Container", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
