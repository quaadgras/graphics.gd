// Package BackBufferCopy provides methods for working with BackBufferCopy object instances.
package BackBufferCopy

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Rect2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode

/*
Node for back-buffering the currently-displayed screen. The region defined in the [BackBufferCopy] node is buffered with the content of the screen it covers, or the entire screen according to the [member copy_mode]. It can be accessed in shader scripts using the screen texture (i.e. a uniform sampler with [code]hint_screen_texture[/code]).
[b]Note:[/b] Since this node inherits from [Node2D] (and not [Control]), anchors and margins won't apply to child [Control]-derived nodes. This can be problematic when resizing the window. To avoid this, add [Control]-derived nodes as [i]siblings[/i] to the [BackBufferCopy] node instead of adding them as children.
*/
type Instance [1]gdclass.BackBufferCopy

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsBackBufferCopy() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.BackBufferCopy

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BackBufferCopy"))
	casted := Instance{*(*gdclass.BackBufferCopy)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) CopyMode() gdclass.BackBufferCopyCopyMode {
	return gdclass.BackBufferCopyCopyMode(class(self).GetCopyMode())
}

func (self Instance) SetCopyMode(value gdclass.BackBufferCopyCopyMode) {
	class(self).SetCopyMode(value)
}

func (self Instance) Rect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetRect())
}

func (self Instance) SetRect(value Rect2.PositionSize) {
	class(self).SetRect(gd.Rect2(value))
}

//go:nosplit
func (self class) SetRect(rect gd.Rect2) { //gd:BackBufferCopy.set_rect
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BackBufferCopy.Bind_set_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRect() gd.Rect2 { //gd:BackBufferCopy.get_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BackBufferCopy.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCopyMode(copy_mode gdclass.BackBufferCopyCopyMode) { //gd:BackBufferCopy.set_copy_mode
	var frame = callframe.New()
	callframe.Arg(frame, copy_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BackBufferCopy.Bind_set_copy_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCopyMode() gdclass.BackBufferCopyCopyMode { //gd:BackBufferCopy.get_copy_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BackBufferCopyCopyMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BackBufferCopy.Bind_get_copy_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
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
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("BackBufferCopy", func(ptr gd.Object) any {
		return [1]gdclass.BackBufferCopy{*(*gdclass.BackBufferCopy)(unsafe.Pointer(&ptr))}
	})
}

type CopyMode = gdclass.BackBufferCopyCopyMode //gd:BackBufferCopy.CopyMode

const (
	/*Disables the buffering mode. This means the [BackBufferCopy] node will directly use the portion of screen it covers.*/
	CopyModeDisabled CopyMode = 0
	/*[BackBufferCopy] buffers a rectangular region.*/
	CopyModeRect CopyMode = 1
	/*[BackBufferCopy] buffers the entire screen.*/
	CopyModeViewport CopyMode = 2
)
