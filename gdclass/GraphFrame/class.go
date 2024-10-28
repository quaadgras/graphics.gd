package GraphFrame

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GraphElement"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
GraphFrame is a special [GraphElement] to which other [GraphElement]s can be attached. It can be configured to automatically resize to enclose all attached [GraphElement]s. If the frame is moved, all the attached [GraphElement]s inside it will be moved as well.
A GraphFrame is always kept behind the connection layer and other [GraphElement]s inside a [GraphEdit].

*/
type Go [1]classdb.GraphFrame

/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default.
This can be used to add custom controls to the title bar such as option or close buttons.
*/
func (self Go) GetTitlebarHbox() gdclass.HBoxContainer {
	return gdclass.HBoxContainer(class(self).GetTitlebarHbox())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GraphFrame
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphFrame"))
	return Go{classdb.GraphFrame(object)}
}

func (self Go) Title() string {
		return string(class(self).GetTitle().String())
}

func (self Go) SetTitle(value string) {
	class(self).SetTitle(gd.NewString(value))
}

func (self Go) AutoshrinkEnabled() bool {
		return bool(class(self).IsAutoshrinkEnabled())
}

func (self Go) SetAutoshrinkEnabled(value bool) {
	class(self).SetAutoshrinkEnabled(value)
}

func (self Go) AutoshrinkMargin() int {
		return int(int(class(self).GetAutoshrinkMargin()))
}

func (self Go) SetAutoshrinkMargin(value int) {
	class(self).SetAutoshrinkMargin(gd.Int(value))
}

func (self Go) DragMargin() int {
		return int(int(class(self).GetDragMargin()))
}

func (self Go) SetDragMargin(value int) {
	class(self).SetDragMargin(gd.Int(value))
}

func (self Go) TintColorEnabled() bool {
		return bool(class(self).IsTintColorEnabled())
}

func (self Go) SetTintColorEnabled(value bool) {
	class(self).SetTintColorEnabled(value)
}

func (self Go) TintColor() gd.Color {
		return gd.Color(class(self).GetTintColor())
}

func (self Go) SetTintColor(value gd.Color) {
	class(self).SetTintColor(value)
}

//go:nosplit
func (self class) SetTitle(title gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(title))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTitle() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default.
This can be used to add custom controls to the title bar such as option or close buttons.
*/
//go:nosplit
func (self class) GetTitlebarHbox() gdclass.HBoxContainer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_get_titlebar_hbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.HBoxContainer{classdb.HBoxContainer(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoshrinkEnabled(shrink bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, shrink)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_set_autoshrink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoshrinkEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_is_autoshrink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoshrinkMargin(autoshrink_margin gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, autoshrink_margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_set_autoshrink_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoshrinkMargin() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_get_autoshrink_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragMargin(drag_margin gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, drag_margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_set_drag_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDragMargin() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_get_drag_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTintColorEnabled(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_set_tint_color_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsTintColorEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_is_tint_color_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTintColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_set_tint_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTintColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_get_tint_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnAutoshrinkChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("autoshrink_changed"), gd.NewCallable(cb), 0)
}


func (self class) AsGraphFrame() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGraphFrame() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGraphElement() GraphElement.GD { return *((*GraphElement.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGraphElement() GraphElement.Go { return *((*GraphElement.Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsGraphElement(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGraphElement(), name)
	}
}
func init() {classdb.Register("GraphFrame", func(ptr gd.Object) any { return classdb.GraphFrame(ptr) })}
