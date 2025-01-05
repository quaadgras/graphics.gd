// Package Container provides methods for working with Container object instances.
package Container

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Rect2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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
type Instance [1]gdclass.Container
type Any interface {
	gd.IsClass
	AsContainer() Instance
}

/*
Implement to return a list of allowed horizontal [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
*/
func (Instance) _get_allowed_size_flags_horizontal(impl func(ptr unsafe.Pointer) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Implement to return a list of allowed vertical [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
*/
func (Instance) _get_allowed_size_flags_vertical(impl func(ptr unsafe.Pointer) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Queue resort of the contained children. This is called automatically anyway, but can be called upon request.
*/
func (self Instance) QueueSort() {
	class(self).QueueSort()
}

/*
Fit a child control in a given rect. This is mainly a helper for creating custom container classes.
*/
func (self Instance) FitChildInRect(child [1]gdclass.Control, rect Rect2.PositionSize) {
	class(self).FitChildInRect(child, gd.Rect2(rect))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Container

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Container"))
	return Instance{*(*gdclass.Container)(unsafe.Pointer(&object))}
}

/*
Implement to return a list of allowed horizontal [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
*/
func (class) _get_allowed_size_flags_horizontal(impl func(ptr unsafe.Pointer) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Implement to return a list of allowed vertical [enum Control.SizeFlags] for child nodes. This doesn't technically prevent the usages of any other size flags, if your implementation requires that. This only limits the options available to the user in the Inspector dock.
[b]Note:[/b] Having no size flags is equal to having [constant Control.SIZE_SHRINK_BEGIN]. As such, this value is always implicitly allowed.
*/
func (class) _get_allowed_size_flags_vertical(impl func(ptr unsafe.Pointer) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Queue resort of the contained children. This is called automatically anyway, but can be called upon request.
*/
//go:nosplit
func (self class) QueueSort() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Container.Bind_queue_sort, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Fit a child control in a given rect. This is mainly a helper for creating custom container classes.
*/
//go:nosplit
func (self class) FitChildInRect(child [1]gdclass.Control, rect gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(child[0])[0])
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Container.Bind_fit_child_in_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnPreSortChildren(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pre_sort_children"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSortChildren(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("sort_children"), gd.NewCallable(cb), 0)
}

func (self class) AsContainer() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsContainer() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_allowed_size_flags_horizontal":
		return reflect.ValueOf(self._get_allowed_size_flags_horizontal)
	case "_get_allowed_size_flags_vertical":
		return reflect.ValueOf(self._get_allowed_size_flags_vertical)
	default:
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_allowed_size_flags_horizontal":
		return reflect.ValueOf(self._get_allowed_size_flags_horizontal)
	case "_get_allowed_size_flags_vertical":
		return reflect.ValueOf(self._get_allowed_size_flags_vertical)
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("Container", func(ptr gd.Object) any { return [1]gdclass.Container{*(*gdclass.Container)(unsafe.Pointer(&ptr))} })
}
