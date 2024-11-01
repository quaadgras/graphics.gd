package Range

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Range is an abstract base class for controls that represent a number within a range, using a configured [member step] and [member page] size. See e.g. [ScrollBar] and [Slider] for examples of higher-level nodes using Range.

	// Range methods that can be overridden by a [Class] that extends it.
	type Range interface {
		//Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
		ValueChanged(new_value float64)
	}
*/
type Instance [1]classdb.Range

/*
Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
*/
func (Instance) _value_changed(impl func(ptr unsafe.Pointer, new_value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var new_value = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, float64(new_value))
	}
}

/*
Sets the [Range]'s current value to the specified [param value], without emitting the [signal value_changed] signal.
*/
func (self Instance) SetValueNoSignal(value float64) {
	class(self).SetValueNoSignal(gd.Float(value))
}

/*
Binds two [Range]s together along with any ranges previously grouped with either of them. When any of range's member variables change, it will share the new value with all other ranges in its group.
*/
func (self Instance) Share(with objects.Node) {
	class(self).Share(with)
}

/*
Stops the [Range] from sharing its member variables with any other.
*/
func (self Instance) Unshare() {
	class(self).Unshare()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Range

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Range"))
	return Instance{classdb.Range(object)}
}

func (self Instance) MinValue() float64 {
	return float64(float64(class(self).GetMin()))
}

func (self Instance) SetMinValue(value float64) {
	class(self).SetMin(gd.Float(value))
}

func (self Instance) MaxValue() float64 {
	return float64(float64(class(self).GetMax()))
}

func (self Instance) SetMaxValue(value float64) {
	class(self).SetMax(gd.Float(value))
}

func (self Instance) Step() float64 {
	return float64(float64(class(self).GetStep()))
}

func (self Instance) SetStep(value float64) {
	class(self).SetStep(gd.Float(value))
}

func (self Instance) Page() float64 {
	return float64(float64(class(self).GetPage()))
}

func (self Instance) SetPage(value float64) {
	class(self).SetPage(gd.Float(value))
}

func (self Instance) Value() float64 {
	return float64(float64(class(self).GetValue()))
}

func (self Instance) SetValue(value float64) {
	class(self).SetValue(gd.Float(value))
}

func (self Instance) Ratio() float64 {
	return float64(float64(class(self).GetAsRatio()))
}

func (self Instance) SetRatio(value float64) {
	class(self).SetAsRatio(gd.Float(value))
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
func (class) _value_changed(impl func(ptr unsafe.Pointer, new_value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var new_value = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, new_value)
	}
}

//go:nosplit
func (self class) GetValue() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMax() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetStep() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPage() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_page, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetAsRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_get_as_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetValue(value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the [Range]'s current value to the specified [param value], without emitting the [signal value_changed] signal.
*/
//go:nosplit
func (self class) SetValueNoSignal(value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_value_no_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetMin(minimum gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, minimum)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetMax(maximum gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, maximum)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetStep(step gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, step)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetPage(pagesize gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pagesize)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_page, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetAsRatio(value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_as_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseRoundedValues(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_use_rounded_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingRoundedValues() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_is_using_rounded_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExpRatio(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_exp_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsRatioExp() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_is_ratio_exp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowGreater(allow bool) {
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_allow_greater, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsGreaterAllowed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_is_greater_allowed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowLesser(allow bool) {
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_allow_lesser, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsLesserAllowed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_is_lesser_allowed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Binds two [Range]s together along with any ranges previously grouped with either of them. When any of range's member variables change, it will share the new value with all other ranges in its group.
*/
//go:nosplit
func (self class) Share(with objects.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(with[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_share, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Stops the [Range] from sharing its member variables with any other.
*/
//go:nosplit
func (self class) Unshare() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_unshare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnValueChanged(cb func(value float64)) {
	self[0].AsObject().Connect(gd.NewStringName("value_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
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
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_value_changed":
		return reflect.ValueOf(self._value_changed)
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() { classdb.Register("Range", func(ptr gd.Object) any { return classdb.Range(ptr) }) }
