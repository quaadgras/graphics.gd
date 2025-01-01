package SubViewportContainer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Container"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A container that displays the contents of underlying [SubViewport] child nodes. It uses the combined size of the [SubViewport]s as minimum size, unless [member stretch] is enabled.
[b]Note:[/b] Changing a [SubViewportContainer]'s [member Control.scale] will cause its contents to appear distorted. To change its visual size without causing distortion, adjust the node's margins instead (if it's not already in a container).
[b]Note:[/b] The [SubViewportContainer] forwards mouse-enter and mouse-exit notifications to its sub-viewports.

	// SubViewportContainer methods that can be overridden by a [Class] that extends it.
	type SubViewportContainer interface {
		//Virtual method to be implemented by the user. If it returns [code]true[/code], the [param event] is propagated to [SubViewport] children. Propagation doesn't happen if it returns [code]false[/code]. If the function is not implemented, all events are propagated to SubViewports.
		PropagateInputEvent(event objects.InputEvent) bool
	}
*/
type Instance [1]classdb.SubViewportContainer
type Any interface {
	gd.IsClass
	AsSubViewportContainer() Instance
}

/*
Virtual method to be implemented by the user. If it returns [code]true[/code], the [param event] is propagated to [SubViewport] children. Propagation doesn't happen if it returns [code]false[/code]. If the function is not implemented, all events are propagated to SubViewports.
*/
func (Instance) _propagate_input_event(impl func(ptr unsafe.Pointer, event objects.InputEvent) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var event = objects.InputEvent{pointers.New[classdb.InputEvent]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SubViewportContainer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SubViewportContainer"))
	return Instance{classdb.SubViewportContainer(object)}
}

func (self Instance) Stretch() bool {
	return bool(class(self).IsStretchEnabled())
}

func (self Instance) SetStretch(value bool) {
	class(self).SetStretch(value)
}

func (self Instance) StretchShrink() int {
	return int(int(class(self).GetStretchShrink()))
}

func (self Instance) SetStretchShrink(value int) {
	class(self).SetStretchShrink(gd.Int(value))
}

/*
Virtual method to be implemented by the user. If it returns [code]true[/code], the [param event] is propagated to [SubViewport] children. Propagation doesn't happen if it returns [code]false[/code]. If the function is not implemented, all events are propagated to SubViewports.
*/
func (class) _propagate_input_event(impl func(ptr unsafe.Pointer, event objects.InputEvent) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var event = objects.InputEvent{pointers.New[classdb.InputEvent]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) SetStretch(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewportContainer.Bind_set_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsStretchEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewportContainer.Bind_is_stretch_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStretchShrink(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewportContainer.Bind_set_stretch_shrink, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStretchShrink() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SubViewportContainer.Bind_get_stretch_shrink, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSubViewportContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSubViewportContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
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
	case "_propagate_input_event":
		return reflect.ValueOf(self._propagate_input_event)
	default:
		return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_propagate_input_event":
		return reflect.ValueOf(self._propagate_input_event)
	default:
		return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {
	classdb.Register("SubViewportContainer", func(ptr gd.Object) any { return [1]classdb.SubViewportContainer{classdb.SubViewportContainer(ptr)} })
}
