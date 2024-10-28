package Range

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Range is an abstract base class for controls that represent a number within a range, using a configured [member step] and [member page] size. See e.g. [ScrollBar] and [Slider] for examples of higher-level nodes using Range.
	// Range methods that can be overridden by a [Class] that extends it.
	type Range interface {
		//Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
		ValueChanged(new_value float64) 
	}

*/
type Go [1]classdb.Range

/*
Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
*/
func (Go) _value_changed(impl func(ptr unsafe.Pointer, new_value float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var new_value = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(new_value))
	}
}

/*
Sets the [Range]'s current value to the specified [param value], without emitting the [signal value_changed] signal.
*/
func (self Go) SetValueNoSignal(value float64) {
	class(self).SetValueNoSignal(gd.Float(value))
}

/*
Binds two [Range]s together along with any ranges previously grouped with either of them. When any of range's member variables change, it will share the new value with all other ranges in its group.
*/
func (self Go) Share(with gdclass.Node) {
	class(self).Share(with)
}

/*
Stops the [Range] from sharing its member variables with any other.
*/
func (self Go) Unshare() {
	class(self).Unshare()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Range
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Range"))
	return Go{classdb.Range(object)}
}

func (self Go) MinValue() float64 {
		return float64(float64(class(self).GetMin()))
}

func (self Go) SetMinValue(value float64) {
	class(self).SetMin(gd.Float(value))
}

func (self Go) MaxValue() float64 {
		return float64(float64(class(self).GetMax()))
}

func (self Go) SetMaxValue(value float64) {
	class(self).SetMax(gd.Float(value))
}

func (self Go) Step() float64 {
		return float64(float64(class(self).GetStep()))
}

func (self Go) SetStep(value float64) {
	class(self).SetStep(gd.Float(value))
}

func (self Go) Page() float64 {
		return float64(float64(class(self).GetPage()))
}

func (self Go) SetPage(value float64) {
	class(self).SetPage(gd.Float(value))
}

func (self Go) Value() float64 {
		return float64(float64(class(self).GetValue()))
}

func (self Go) SetValue(value float64) {
	class(self).SetValue(gd.Float(value))
}

func (self Go) Ratio() float64 {
		return float64(float64(class(self).GetAsRatio()))
}

func (self Go) SetRatio(value float64) {
	class(self).SetAsRatio(gd.Float(value))
}

func (self Go) ExpEdit() bool {
		return bool(class(self).IsRatioExp())
}

func (self Go) SetExpEdit(value bool) {
	class(self).SetExpRatio(value)
}

func (self Go) Rounded() bool {
		return bool(class(self).IsUsingRoundedValues())
}

func (self Go) SetRounded(value bool) {
	class(self).SetUseRoundedValues(value)
}

func (self Go) AllowGreater() bool {
		return bool(class(self).IsGreaterAllowed())
}

func (self Go) SetAllowGreater(value bool) {
	class(self).SetAllowGreater(value)
}

func (self Go) AllowLesser() bool {
		return bool(class(self).IsLesserAllowed())
}

func (self Go) SetAllowLesser(value bool) {
	class(self).SetAllowLesser(value)
}

/*
Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
*/
func (class) _value_changed(impl func(ptr unsafe.Pointer, new_value gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var new_value = gd.UnsafeGet[gd.Float](p_args,0)
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
func (self class) SetValue(value gd.Float)  {
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
func (self class) SetValueNoSignal(value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_value_no_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMin(minimum gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, minimum)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMax(maximum gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, maximum)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetStep(step gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, step)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetPage(pagesize gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, pagesize)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_page, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAsRatio(value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_set_as_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetUseRoundedValues(enabled bool)  {
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
func (self class) SetExpRatio(enabled bool)  {
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
func (self class) SetAllowGreater(allow bool)  {
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
func (self class) SetAllowLesser(allow bool)  {
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
func (self class) Share(with gdclass.Node)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(with[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_share, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops the [Range] from sharing its member variables with any other.
*/
//go:nosplit
func (self class) Unshare()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Range.Bind_unshare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnValueChanged(cb func(value float64)) {
	self[0].AsObject().Connect(gd.NewStringName("value_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
}


func (self class) AsRange() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRange() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_value_changed": return reflect.ValueOf(self._value_changed);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_value_changed": return reflect.ValueOf(self._value_changed);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("Range", func(ptr gd.Object) any { return classdb.Range(ptr) })}
