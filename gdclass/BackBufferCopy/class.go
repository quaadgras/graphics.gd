package BackBufferCopy

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Node for back-buffering the currently-displayed screen. The region defined in the [BackBufferCopy] node is buffered with the content of the screen it covers, or the entire screen according to the [member copy_mode]. It can be accessed in shader scripts using the screen texture (i.e. a uniform sampler with [code]hint_screen_texture[/code]).
[b]Note:[/b] Since this node inherits from [Node2D] (and not [Control]), anchors and margins won't apply to child [Control]-derived nodes. This can be problematic when resizing the window. To avoid this, add [Control]-derived nodes as [i]siblings[/i] to the [BackBufferCopy] node instead of adding them as children.

*/
type Go [1]classdb.BackBufferCopy
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.BackBufferCopy
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BackBufferCopy"))
	return Go{classdb.BackBufferCopy(object)}
}

func (self Go) CopyMode() classdb.BackBufferCopyCopyMode {
		return classdb.BackBufferCopyCopyMode(class(self).GetCopyMode())
}

func (self Go) SetCopyMode(value classdb.BackBufferCopyCopyMode) {
	class(self).SetCopyMode(value)
}

func (self Go) Rect() gd.Rect2 {
		return gd.Rect2(class(self).GetRect())
}

func (self Go) SetRect(value gd.Rect2) {
	class(self).SetRect(value)
}

//go:nosplit
func (self class) SetRect(rect gd.Rect2)  {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BackBufferCopy.Bind_set_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BackBufferCopy.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCopyMode(copy_mode classdb.BackBufferCopyCopyMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, copy_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BackBufferCopy.Bind_set_copy_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCopyMode() classdb.BackBufferCopyCopyMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BackBufferCopyCopyMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BackBufferCopy.Bind_get_copy_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsBackBufferCopy() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsBackBufferCopy() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("BackBufferCopy", func(ptr gd.Object) any { return classdb.BackBufferCopy(ptr) })}
type CopyMode = classdb.BackBufferCopyCopyMode

const (
/*Disables the buffering mode. This means the [BackBufferCopy] node will directly use the portion of screen it covers.*/
	CopyModeDisabled CopyMode = 0
/*[BackBufferCopy] buffers a rectangular region.*/
	CopyModeRect CopyMode = 1
/*[BackBufferCopy] buffers the entire screen.*/
	CopyModeViewport CopyMode = 2
)
