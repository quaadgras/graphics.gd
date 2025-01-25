// Package AspectRatioContainer provides methods for working with AspectRatioContainer object instances.
package AspectRatioContainer

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
import "graphics.gd/variant/Float"

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
A container type that arranges its child controls in a way that preserves their proportions automatically when the container is resized. Useful when a container has a dynamic size and the child nodes must adjust their sizes accordingly without losing their aspect ratios.
*/
type Instance [1]gdclass.AspectRatioContainer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAspectRatioContainer() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AspectRatioContainer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AspectRatioContainer"))
	casted := Instance{*(*gdclass.AspectRatioContainer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Ratio() Float.X {
	return Float.X(Float.X(class(self).GetRatio()))
}

func (self Instance) SetRatio(value Float.X) {
	class(self).SetRatio(gd.Float(value))
}

func (self Instance) StretchMode() gdclass.AspectRatioContainerStretchMode {
	return gdclass.AspectRatioContainerStretchMode(class(self).GetStretchMode())
}

func (self Instance) SetStretchMode(value gdclass.AspectRatioContainerStretchMode) {
	class(self).SetStretchMode(value)
}

func (self Instance) AlignmentHorizontal() gdclass.AspectRatioContainerAlignmentMode {
	return gdclass.AspectRatioContainerAlignmentMode(class(self).GetAlignmentHorizontal())
}

func (self Instance) SetAlignmentHorizontal(value gdclass.AspectRatioContainerAlignmentMode) {
	class(self).SetAlignmentHorizontal(value)
}

func (self Instance) AlignmentVertical() gdclass.AspectRatioContainerAlignmentMode {
	return gdclass.AspectRatioContainerAlignmentMode(class(self).GetAlignmentVertical())
}

func (self Instance) SetAlignmentVertical(value gdclass.AspectRatioContainerAlignmentMode) {
	class(self).SetAlignmentVertical(value)
}

//go:nosplit
func (self class) SetRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AspectRatioContainer.Bind_set_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AspectRatioContainer.Bind_get_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStretchMode(stretch_mode gdclass.AspectRatioContainerStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, stretch_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AspectRatioContainer.Bind_set_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStretchMode() gdclass.AspectRatioContainerStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AspectRatioContainerStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AspectRatioContainer.Bind_get_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlignmentHorizontal(alignment_horizontal gdclass.AspectRatioContainerAlignmentMode) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment_horizontal)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AspectRatioContainer.Bind_set_alignment_horizontal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlignmentHorizontal() gdclass.AspectRatioContainerAlignmentMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AspectRatioContainerAlignmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AspectRatioContainer.Bind_get_alignment_horizontal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlignmentVertical(alignment_vertical gdclass.AspectRatioContainerAlignmentMode) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment_vertical)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AspectRatioContainer.Bind_set_alignment_vertical, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlignmentVertical() gdclass.AspectRatioContainerAlignmentMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AspectRatioContainerAlignmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AspectRatioContainer.Bind_get_alignment_vertical, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAspectRatioContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAspectRatioContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AspectRatioContainer", func(ptr gd.Object) any {
		return [1]gdclass.AspectRatioContainer{*(*gdclass.AspectRatioContainer)(unsafe.Pointer(&ptr))}
	})
}

type StretchMode = gdclass.AspectRatioContainerStretchMode

const (
	/*The height of child controls is automatically adjusted based on the width of the container.*/
	StretchWidthControlsHeight StretchMode = 0
	/*The width of child controls is automatically adjusted based on the height of the container.*/
	StretchHeightControlsWidth StretchMode = 1
	/*The bounding rectangle of child controls is automatically adjusted to fit inside the container while keeping the aspect ratio.*/
	StretchFit StretchMode = 2
	/*The width and height of child controls is automatically adjusted to make their bounding rectangle cover the entire area of the container while keeping the aspect ratio.
	  When the bounding rectangle of child controls exceed the container's size and [member Control.clip_contents] is enabled, this allows to show only the container's area restricted by its own bounding rectangle.*/
	StretchCover StretchMode = 3
)

type AlignmentMode = gdclass.AspectRatioContainerAlignmentMode

const (
	/*Aligns child controls with the beginning (left or top) of the container.*/
	AlignmentBegin AlignmentMode = 0
	/*Aligns child controls with the center of the container.*/
	AlignmentCenter AlignmentMode = 1
	/*Aligns child controls with the end (right or bottom) of the container.*/
	AlignmentEnd AlignmentMode = 2
)
