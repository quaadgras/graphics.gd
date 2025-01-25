// Package BoxContainer provides methods for working with BoxContainer object instances.
package BoxContainer

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
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
A container that arranges its child controls horizontally or vertically, rearranging them automatically when their minimum size changes.
*/
type Instance [1]gdclass.BoxContainer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsBoxContainer() Instance
}

/*
Adds a [Control] node to the box as a spacer. If [param begin] is [code]true[/code], it will insert the [Control] node in front of all other children.
*/
func (self Instance) AddSpacer(begin bool) [1]gdclass.Control {
	return [1]gdclass.Control(class(self).AddSpacer(begin))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.BoxContainer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BoxContainer"))
	casted := Instance{*(*gdclass.BoxContainer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Alignment() gdclass.BoxContainerAlignmentMode {
	return gdclass.BoxContainerAlignmentMode(class(self).GetAlignment())
}

func (self Instance) SetAlignment(value gdclass.BoxContainerAlignmentMode) {
	class(self).SetAlignment(value)
}

func (self Instance) Vertical() bool {
	return bool(class(self).IsVertical())
}

func (self Instance) SetVertical(value bool) {
	class(self).SetVertical(value)
}

/*
Adds a [Control] node to the box as a spacer. If [param begin] is [code]true[/code], it will insert the [Control] node in front of all other children.
*/
//go:nosplit
func (self class) AddSpacer(begin bool) [1]gdclass.Control {
	var frame = callframe.New()
	callframe.Arg(frame, begin)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_add_spacer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Control{gd.PointerLifetimeBoundTo[gdclass.Control](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlignment(alignment gdclass.BoxContainerAlignmentMode) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_set_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlignment() gdclass.BoxContainerAlignmentMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.BoxContainerAlignmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_get_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVertical(vertical bool) {
	var frame = callframe.New()
	callframe.Arg(frame, vertical)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVertical() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxContainer.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsBoxContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBoxContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Container.Advanced(self.AsContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Container.Instance(self.AsContainer()), name)
	}
}
func init() {
	gdclass.Register("BoxContainer", func(ptr gd.Object) any {
		return [1]gdclass.BoxContainer{*(*gdclass.BoxContainer)(unsafe.Pointer(&ptr))}
	})
}

type AlignmentMode = gdclass.BoxContainerAlignmentMode

const (
	/*The child controls will be arranged at the beginning of the container, i.e. top if orientation is vertical, left if orientation is horizontal (right for RTL layout).*/
	AlignmentBegin AlignmentMode = 0
	/*The child controls will be centered in the container.*/
	AlignmentCenter AlignmentMode = 1
	/*The child controls will be arranged at the end of the container, i.e. bottom if orientation is vertical, right if orientation is horizontal (left for RTL layout).*/
	AlignmentEnd AlignmentMode = 2
)
