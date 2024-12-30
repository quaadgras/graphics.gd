package GraphFrame

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/GraphElement"
import "graphics.gd/objects/Container"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Color"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
GraphFrame is a special [GraphElement] to which other [GraphElement]s can be attached. It can be configured to automatically resize to enclose all attached [GraphElement]s. If the frame is moved, all the attached [GraphElement]s inside it will be moved as well.
A GraphFrame is always kept behind the connection layer and other [GraphElement]s inside a [GraphEdit].
*/
type Instance [1]classdb.GraphFrame
type Any interface {
	gd.IsClass
	AsGraphFrame() Instance
}

/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default.
This can be used to add custom controls to the title bar such as option or close buttons.
*/
func (self Instance) GetTitlebarHbox() objects.HBoxContainer {
	return objects.HBoxContainer(class(self).GetTitlebarHbox())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GraphFrame

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphFrame"))
	return Instance{classdb.GraphFrame(object)}
}

func (self Instance) Title() string {
	return string(class(self).GetTitle().String())
}

func (self Instance) SetTitle(value string) {
	class(self).SetTitle(gd.NewString(value))
}

func (self Instance) AutoshrinkEnabled() bool {
	return bool(class(self).IsAutoshrinkEnabled())
}

func (self Instance) SetAutoshrinkEnabled(value bool) {
	class(self).SetAutoshrinkEnabled(value)
}

func (self Instance) AutoshrinkMargin() int {
	return int(int(class(self).GetAutoshrinkMargin()))
}

func (self Instance) SetAutoshrinkMargin(value int) {
	class(self).SetAutoshrinkMargin(gd.Int(value))
}

func (self Instance) DragMargin() int {
	return int(int(class(self).GetDragMargin()))
}

func (self Instance) SetDragMargin(value int) {
	class(self).SetDragMargin(gd.Int(value))
}

func (self Instance) TintColorEnabled() bool {
	return bool(class(self).IsTintColorEnabled())
}

func (self Instance) SetTintColorEnabled(value bool) {
	class(self).SetTintColorEnabled(value)
}

func (self Instance) TintColor() Color.RGBA {
	return Color.RGBA(class(self).GetTintColor())
}

func (self Instance) SetTintColor(value Color.RGBA) {
	class(self).SetTintColor(gd.Color(value))
}

//go:nosplit
func (self class) SetTitle(title gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTitle() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [HBoxContainer] used for the title bar, only containing a [Label] for displaying the title by default.
This can be used to add custom controls to the title bar such as option or close buttons.
*/
//go:nosplit
func (self class) GetTitlebarHbox() objects.HBoxContainer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphFrame.Bind_get_titlebar_hbox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.HBoxContainer{classdb.HBoxContainer(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoshrinkEnabled(shrink bool) {
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
func (self class) SetAutoshrinkMargin(autoshrink_margin gd.Int) {
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
func (self class) SetDragMargin(drag_margin gd.Int) {
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
func (self class) SetTintColorEnabled(enable bool) {
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
func (self class) SetTintColor(color gd.Color) {
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
func (self Instance) OnAutoshrinkChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("autoshrink_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsGraphFrame() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGraphFrame() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGraphElement() GraphElement.Advanced {
	return *((*GraphElement.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGraphElement() GraphElement.Instance {
	return *((*GraphElement.Instance)(unsafe.Pointer(&self)))
}
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
	default:
		return gd.VirtualByName(self.AsGraphElement(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsGraphElement(), name)
	}
}
func init() {
	classdb.Register("GraphFrame", func(ptr gd.Object) any { return classdb.GraphFrame(ptr) })
}
