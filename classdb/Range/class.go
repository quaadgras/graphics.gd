// Package Range provides methods for working with Range object instances.
package Range

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Range is an abstract base class for controls that represent a number within a range, using a configured [member step] and [member page] size. See e.g. [ScrollBar] and [Slider] for examples of higher-level nodes using Range.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=Range)
*/
type Instance [1]gdclass.Range

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRange() Instance
}
type Interface interface {
	//Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
	ValueChanged(new_value Float.X)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) ValueChanged(new_value Float.X) { return }

/*
Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
*/
func (Instance) _value_changed(impl func(ptr unsafe.Pointer, new_value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var new_value = gd.UnsafeGet[float64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(new_value))
	}
}

/*
Sets the [Range]'s current value to the specified [param value], without emitting the [signal value_changed] signal.
*/
func (self Instance) SetValueNoSignal(value Float.X) { //gd:Range.set_value_no_signal
	Advanced(self).SetValueNoSignal(float64(value))
}

/*
Binds two [Range]s together along with any ranges previously grouped with either of them. When any of range's member variables change, it will share the new value with all other ranges in its group.
*/
func (self Instance) Share(with [1]gdclass.Node) { //gd:Range.share
	Advanced(self).Share(with)
}

/*
Stops the [Range] from sharing its member variables with any other.
*/
func (self Instance) Unshare() { //gd:Range.unshare
	Advanced(self).Unshare()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Range

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Range"))
	casted := Instance{*(*gdclass.Range)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) MinValue() Float.X {
	return Float.X(Float.X(class(self).GetMin()))
}

func (self Instance) SetMinValue(value Float.X) {
	class(self).SetMin(float64(value))
}

func (self Instance) MaxValue() Float.X {
	return Float.X(Float.X(class(self).GetMax()))
}

func (self Instance) SetMaxValue(value Float.X) {
	class(self).SetMax(float64(value))
}

func (self Instance) Step() Float.X {
	return Float.X(Float.X(class(self).GetStep()))
}

func (self Instance) SetStep(value Float.X) {
	class(self).SetStep(float64(value))
}

func (self Instance) Page() Float.X {
	return Float.X(Float.X(class(self).GetPage()))
}

func (self Instance) SetPage(value Float.X) {
	class(self).SetPage(float64(value))
}

func (self Instance) Value() Float.X {
	return Float.X(Float.X(class(self).GetValue()))
}

func (self Instance) SetValue(value Float.X) {
	class(self).SetValue(float64(value))
}

func (self Instance) Ratio() Float.X {
	return Float.X(Float.X(class(self).GetAsRatio()))
}

func (self Instance) SetRatio(value Float.X) {
	class(self).SetAsRatio(float64(value))
}

func (self Instance) ExpEdit() bool {
	return bool(class(self).IsRatioExp())
}

func (self Instance) SetExpEdit(value bool) {
	class(self).SetExpRatio(value)
}

func (self Instance) Rounded() bool {
	return bool(class(self).IsUsingRoundedValues())
}

func (self Instance) SetRounded(value bool) {
	class(self).SetUseRoundedValues(value)
}

func (self Instance) AllowGreater() bool {
	return bool(class(self).IsGreaterAllowed())
}

func (self Instance) SetAllowGreater(value bool) {
	class(self).SetAllowGreater(value)
}

func (self Instance) AllowLesser() bool {
	return bool(class(self).IsLesserAllowed())
}

func (self Instance) SetAllowLesser(value bool) {
	class(self).SetAllowLesser(value)
}

/*
Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
*/
func (class) _value_changed(impl func(ptr unsafe.Pointer, new_value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var new_value = gd.UnsafeGet[float64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, new_value)
	}
}

//go:nosplit
func (self class) GetValue() float64 { //gd:Range.get_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMin() float64 { //gd:Range.get_min
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMax() float64 { //gd:Range.get_max
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetStep() float64 { //gd:Range.get_step
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPage() float64 { //gd:Range.get_page
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_page, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetAsRatio() float64 { //gd:Range.get_as_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_as_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetValue(value float64) { //gd:Range.set_value
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [Range]'s current value to the specified [param value], without emitting the [signal value_changed] signal.
*/
//go:nosplit
func (self class) SetValueNoSignal(value float64) { //gd:Range.set_value_no_signal
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_value_no_signal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetMin(minimum float64) { //gd:Range.set_min
	var frame = callframe.New()
	callframe.Arg(frame, minimum)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetMax(maximum float64) { //gd:Range.set_max
	var frame = callframe.New()
	callframe.Arg(frame, maximum)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetStep(step float64) { //gd:Range.set_step
	var frame = callframe.New()
	callframe.Arg(frame, step)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetPage(pagesize float64) { //gd:Range.set_page
	var frame = callframe.New()
	callframe.Arg(frame, pagesize)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_page, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAsRatio(value float64) { //gd:Range.set_as_ratio
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_as_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseRoundedValues(enabled bool) { //gd:Range.set_use_rounded_values
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_use_rounded_values, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingRoundedValues() bool { //gd:Range.is_using_rounded_values
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_is_using_rounded_values, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExpRatio(enabled bool) { //gd:Range.set_exp_ratio
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_exp_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRatioExp() bool { //gd:Range.is_ratio_exp
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_is_ratio_exp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowGreater(allow bool) { //gd:Range.set_allow_greater
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_allow_greater, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsGreaterAllowed() bool { //gd:Range.is_greater_allowed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_is_greater_allowed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowLesser(allow bool) { //gd:Range.set_allow_lesser
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_allow_lesser, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsLesserAllowed() bool { //gd:Range.is_lesser_allowed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_is_lesser_allowed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Binds two [Range]s together along with any ranges previously grouped with either of them. When any of range's member variables change, it will share the new value with all other ranges in its group.
*/
//go:nosplit
func (self class) Share(with [1]gdclass.Node) { //gd:Range.share
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(with[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_share, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops the [Range] from sharing its member variables with any other.
*/
//go:nosplit
func (self class) Unshare() { //gd:Range.unshare
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_unshare, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnValueChanged(cb func(value Float.X)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("value_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
}

func (self class) AsRange() Advanced           { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRange() Instance        { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_value_changed":
		return reflect.ValueOf(self._value_changed)
	default:
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_value_changed":
		return reflect.ValueOf(self._value_changed)
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("Range", func(ptr gd.Object) any { return [1]gdclass.Range{*(*gdclass.Range)(unsafe.Pointer(&ptr))} })
}
