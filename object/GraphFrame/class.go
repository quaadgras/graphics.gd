package GraphFrame

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/GraphElement"
import "grow.graphics/gd/object/Container"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
GraphFrame is a special [GraphElement] to which other [GraphElement]s can be attached. It can be configured to automatically resize to enclose all attached [GraphElement]s. If the frame is moved, all the attached [GraphElement]s inside it will be moved as well.
A GraphFrame is always kept behind the connection layer and other [GraphElement]s inside a [GraphEdit].

*/
type Simple [1]classdb.GraphFrame
func (self Simple) SetTitle(title string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTitle(gc.String(title))
}
func (self Simple) GetTitle() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTitle(gc).String())
}
func (self Simple) GetTitlebarHbox() [1]classdb.HBoxContainer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.HBoxContainer(Expert(self).GetTitlebarHbox(gc))
}
func (self Simple) SetAutoshrinkEnabled(shrink bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoshrinkEnabled(shrink)
}
func (self Simple) IsAutoshrinkEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAutoshrinkEnabled())
}
func (self Simple) SetAutoshrinkMargin(autoshrink_margin int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoshrinkMargin(gd.Int(autoshrink_margin))
}
func (self Simple) GetAutoshrinkMargin() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetAutoshrinkMargin()))
}
func (self Simple) SetDragMargin(drag_margin int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDragMargin(gd.Int(drag_margin))
}
func (self Simple) GetDragMargin() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDragMargin()))
}
func (self Simple) SetTintColorEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTintColorEnabled(enable)
}
func (self Simple) IsTintColorEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsTintColorEnabled())
}
func (self Simple) SetTintColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTintColor(color)
}
func (self Simple) GetTintColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetTintColor())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GraphFrame
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTitle(title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTitle(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default.
This can be used to add custom controls to the title bar such as option or close buttons.
*/
//go:nosplit
func (self class) GetTitlebarHbox(ctx gd.Lifetime) object.HBoxContainer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_get_titlebar_hbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.HBoxContainer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoshrinkEnabled(shrink bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shrink)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_set_autoshrink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoshrinkEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_is_autoshrink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoshrinkMargin(autoshrink_margin gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, autoshrink_margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_set_autoshrink_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoshrinkMargin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_get_autoshrink_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragMargin(drag_margin gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, drag_margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_set_drag_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDragMargin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_get_drag_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTintColorEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_set_tint_color_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsTintColorEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_is_tint_color_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTintColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_set_tint_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTintColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphFrame.Bind_get_tint_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsGraphFrame() Expert { return self[0].AsGraphFrame() }


//go:nosplit
func (self Simple) AsGraphFrame() Simple { return self[0].AsGraphFrame() }


//go:nosplit
func (self class) AsGraphElement() GraphElement.Expert { return self[0].AsGraphElement() }


//go:nosplit
func (self Simple) AsGraphElement() GraphElement.Simple { return self[0].AsGraphElement() }


//go:nosplit
func (self class) AsContainer() Container.Expert { return self[0].AsContainer() }


//go:nosplit
func (self Simple) AsContainer() Container.Simple { return self[0].AsContainer() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("GraphFrame", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
