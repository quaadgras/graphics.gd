package StatusIndicator

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

type Instance [1]classdb.StatusIndicator

/*
Returns the status indicator rectangle in screen coordinates. If this status indicator is not visible, returns an empty [Rect2].
*/
func (self Instance) GetRect() gd.Rect2 {
	return gd.Rect2(class(self).GetRect())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StatusIndicator

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StatusIndicator"))
	return Instance{classdb.StatusIndicator(object)}
}

func (self Instance) Tooltip() string {
	return string(class(self).GetTooltip().String())
}

func (self Instance) SetTooltip(value string) {
	class(self).SetTooltip(gd.NewString(value))
}

func (self Instance) Icon() objects.Texture2D {
	return objects.Texture2D(class(self).GetIcon())
}

func (self Instance) SetIcon(value objects.Texture2D) {
	class(self).SetIcon(value)
}

func (self Instance) Menu() string {
	return string(class(self).GetMenu().String())
}

func (self Instance) SetMenu(value string) {
	class(self).SetMenu(gd.NewString(value).NodePath())
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisible(value)
}

//go:nosplit
func (self class) SetTooltip(tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_set_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTooltip() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_get_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIcon(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIcon() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisible(visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMenu(menu gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_set_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMenu() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the status indicator rectangle in screen coordinates. If this status indicator is not visible, returns an empty [Rect2].
*/
//go:nosplit
func (self class) GetRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnPressed(cb func(mouse_button int, mouse_position gd.Vector2i)) {
	self[0].AsObject().Connect(gd.NewStringName("pressed"), gd.NewCallable(cb), 0)
}

func (self class) AsStatusIndicator() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStatusIndicator() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced          { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance       { return *((*Node.Instance)(unsafe.Pointer(&self))) }

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
	classdb.Register("StatusIndicator", func(ptr gd.Object) any { return classdb.StatusIndicator(ptr) })
}
