// Package FlowContainer provides methods for working with FlowContainer object instances.
package FlowContainer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
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

/*
A container that arranges its child controls horizontally or vertically and wraps them around at the borders. This is similar to how text in a book wraps around when no more words can fit on a line.
*/
type Instance [1]gdclass.FlowContainer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFlowContainer() Instance
}

/*
Returns the current line count.
*/
func (self Instance) GetLineCount() int {
	return int(int(class(self).GetLineCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FlowContainer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FlowContainer"))
	casted := Instance{*(*gdclass.FlowContainer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Alignment() gdclass.FlowContainerAlignmentMode {
	return gdclass.FlowContainerAlignmentMode(class(self).GetAlignment())
}

func (self Instance) SetAlignment(value gdclass.FlowContainerAlignmentMode) {
	class(self).SetAlignment(value)
}

func (self Instance) LastWrapAlignment() gdclass.FlowContainerLastWrapAlignmentMode {
	return gdclass.FlowContainerLastWrapAlignmentMode(class(self).GetLastWrapAlignment())
}

func (self Instance) SetLastWrapAlignment(value gdclass.FlowContainerLastWrapAlignmentMode) {
	class(self).SetLastWrapAlignment(value)
}

func (self Instance) Vertical() bool {
	return bool(class(self).IsVertical())
}

func (self Instance) SetVertical(value bool) {
	class(self).SetVertical(value)
}

func (self Instance) ReverseFill() bool {
	return bool(class(self).IsReverseFill())
}

func (self Instance) SetReverseFill(value bool) {
	class(self).SetReverseFill(value)
}

/*
Returns the current line count.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FlowContainer.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlignment(alignment gdclass.FlowContainerAlignmentMode) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FlowContainer.Bind_set_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlignment() gdclass.FlowContainerAlignmentMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FlowContainerAlignmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FlowContainer.Bind_get_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLastWrapAlignment(last_wrap_alignment gdclass.FlowContainerLastWrapAlignmentMode) {
	var frame = callframe.New()
	callframe.Arg(frame, last_wrap_alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FlowContainer.Bind_set_last_wrap_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLastWrapAlignment() gdclass.FlowContainerLastWrapAlignmentMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FlowContainerLastWrapAlignmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FlowContainer.Bind_get_last_wrap_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVertical(vertical bool) {
	var frame = callframe.New()
	callframe.Arg(frame, vertical)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FlowContainer.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVertical() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FlowContainer.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReverseFill(reverse_fill bool) {
	var frame = callframe.New()
	callframe.Arg(frame, reverse_fill)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FlowContainer.Bind_set_reverse_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsReverseFill() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FlowContainer.Bind_is_reverse_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFlowContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFlowContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("FlowContainer", func(ptr gd.Object) any {
		return [1]gdclass.FlowContainer{*(*gdclass.FlowContainer)(unsafe.Pointer(&ptr))}
	})
}

type AlignmentMode = gdclass.FlowContainerAlignmentMode

const (
	/*The child controls will be arranged at the beginning of the container, i.e. top if orientation is vertical, left if orientation is horizontal (right for RTL layout).*/
	AlignmentBegin AlignmentMode = 0
	/*The child controls will be centered in the container.*/
	AlignmentCenter AlignmentMode = 1
	/*The child controls will be arranged at the end of the container, i.e. bottom if orientation is vertical, right if orientation is horizontal (left for RTL layout).*/
	AlignmentEnd AlignmentMode = 2
)

type LastWrapAlignmentMode = gdclass.FlowContainerLastWrapAlignmentMode

const (
	/*The last partially filled row or column will wrap aligned to the previous row or column in accordance with [member alignment].*/
	LastWrapAlignmentInherit LastWrapAlignmentMode = 0
	/*The last partially filled row or column will wrap aligned to the beginning of the previous row or column.*/
	LastWrapAlignmentBegin LastWrapAlignmentMode = 1
	/*The last partially filled row or column will wrap aligned to the center of the previous row or column.*/
	LastWrapAlignmentCenter LastWrapAlignmentMode = 2
	/*The last partially filled row or column will wrap aligned to the end of the previous row or column.*/
	LastWrapAlignmentEnd LastWrapAlignmentMode = 3
)
