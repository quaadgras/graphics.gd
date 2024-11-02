package BackBufferCopy

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"
import "grow.graphics/gd/variant/Rect2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Node for back-buffering the currently-displayed screen. The region defined in the [BackBufferCopy] node is buffered with the content of the screen it covers, or the entire screen according to the [member copy_mode]. It can be accessed in shader scripts using the screen texture (i.e. a uniform sampler with [code]hint_screen_texture[/code]).
[b]Note:[/b] Since this node inherits from [Node2D] (and not [Control]), anchors and margins won't apply to child [Control]-derived nodes. This can be problematic when resizing the window. To avoid this, add [Control]-derived nodes as [i]siblings[/i] to the [BackBufferCopy] node instead of adding them as children.
*/
type Instance [1]classdb.BackBufferCopy

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.BackBufferCopy

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BackBufferCopy"))
	return Instance{classdb.BackBufferCopy(object)}
}

func (self Instance) CopyMode() classdb.BackBufferCopyCopyMode {
	return classdb.BackBufferCopyCopyMode(class(self).GetCopyMode())
}

func (self Instance) SetCopyMode(value classdb.BackBufferCopyCopyMode) {
	class(self).SetCopyMode(value)
}

func (self Instance) Rect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetRect())
}

func (self Instance) SetRect(value Rect2.PositionSize) {
	class(self).SetRect(gd.Rect2(value))
}

//go:nosplit
func (self class) SetRect(rect gd.Rect2) {
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
func (self class) SetCopyMode(copy_mode classdb.BackBufferCopyCopyMode) {
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
func (self class) AsBackBufferCopy() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBackBufferCopy() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced     { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance  { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("BackBufferCopy", func(ptr gd.Object) any { return classdb.BackBufferCopy(ptr) })
}

type CopyMode = classdb.BackBufferCopyCopyMode

const (
	/*Disables the buffering mode. This means the [BackBufferCopy] node will directly use the portion of screen it covers.*/
	CopyModeDisabled CopyMode = 0
	/*[BackBufferCopy] buffers a rectangular region.*/
	CopyModeRect CopyMode = 1
	/*[BackBufferCopy] buffers the entire screen.*/
	CopyModeViewport CopyMode = 2
)
