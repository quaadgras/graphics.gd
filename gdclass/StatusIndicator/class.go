package StatusIndicator

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

type Go [1]classdb.StatusIndicator

/*
Returns the status indicator rectangle in screen coordinates. If this status indicator is not visible, returns an empty [Rect2].
*/
func (self Go) GetRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).GetRect())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.StatusIndicator
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("StatusIndicator"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Tooltip() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetTooltip(gc).String())
}

func (self Go) SetTooltip(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTooltip(gc.String(value))
}

func (self Go) Icon() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetIcon(gc))
}

func (self Go) SetIcon(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIcon(value)
}

func (self Go) Menu() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetMenu(gc).String())
}

func (self Go) SetMenu(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMenu(gc.String(value).NodePath(gc))
}

func (self Go) Visible() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVisible())
}

func (self Go) SetVisible(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisible(value)
}

//go:nosplit
func (self class) SetTooltip(tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StatusIndicator.Bind_set_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTooltip(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StatusIndicator.Bind_get_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIcon(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StatusIndicator.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIcon(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StatusIndicator.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StatusIndicator.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StatusIndicator.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMenu(menu gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(menu))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StatusIndicator.Bind_set_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMenu(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StatusIndicator.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the status indicator rectangle in screen coordinates. If this status indicator is not visible, returns an empty [Rect2].
*/
//go:nosplit
func (self class) GetRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StatusIndicator.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnPressed(cb func(mouse_button int, mouse_position gd.Vector2i)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("pressed"), gc.Callable(cb), 0)
}


func (self class) AsStatusIndicator() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsStatusIndicator() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("StatusIndicator", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
