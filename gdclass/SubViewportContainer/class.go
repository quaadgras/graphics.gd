package SubViewportContainer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A container that displays the contents of underlying [SubViewport] child nodes. It uses the combined size of the [SubViewport]s as minimum size, unless [member stretch] is enabled.
[b]Note:[/b] Changing a [SubViewportContainer]'s [member Control.scale] will cause its contents to appear distorted. To change its visual size without causing distortion, adjust the node's margins instead (if it's not already in a container).
[b]Note:[/b] The [SubViewportContainer] forwards mouse-enter and mouse-exit notifications to its sub-viewports.
	// SubViewportContainer methods that can be overridden by a [Class] that extends it.
	type SubViewportContainer interface {
		//Virtual method to be implemented by the user. If it returns [code]true[/code], the [param event] is propagated to [SubViewport] children. Propagation doesn't happen if it returns [code]false[/code]. If the function is not implemented, all events are propagated to SubViewports.
		PropagateInputEvent(event gdclass.InputEvent) bool
	}

*/
type Go [1]classdb.SubViewportContainer

/*
Virtual method to be implemented by the user. If it returns [code]true[/code], the [param event] is propagated to [SubViewport] children. Propagation doesn't happen if it returns [code]false[/code]. If the function is not implemented, all events are propagated to SubViewports.
*/
func (Go) _propagate_input_event(impl func(ptr unsafe.Pointer, event gdclass.InputEvent) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var event gdclass.InputEvent
		event[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SubViewportContainer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SubViewportContainer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Stretch() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsStretchEnabled())
}

func (self Go) SetStretch(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStretch(value)
}

func (self Go) StretchShrink() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetStretchShrink()))
}

func (self Go) SetStretchShrink(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStretchShrink(gd.Int(value))
}

/*
Virtual method to be implemented by the user. If it returns [code]true[/code], the [param event] is propagated to [SubViewport] children. Propagation doesn't happen if it returns [code]false[/code]. If the function is not implemented, all events are propagated to SubViewports.
*/
func (class) _propagate_input_event(impl func(ptr unsafe.Pointer, event gdclass.InputEvent) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var event gdclass.InputEvent
		event[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetStretch(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewportContainer.Bind_set_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsStretchEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewportContainer.Bind_is_stretch_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStretchShrink(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewportContainer.Bind_set_stretch_shrink, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStretchShrink() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SubViewportContainer.Bind_get_stretch_shrink, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSubViewportContainer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSubViewportContainer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_propagate_input_event": return reflect.ValueOf(self._propagate_input_event);
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_propagate_input_event": return reflect.ValueOf(self._propagate_input_event);
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {classdb.Register("SubViewportContainer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
