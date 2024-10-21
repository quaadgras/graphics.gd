package Range

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Range is an abstract base class for controls that represent a number within a range, using a configured [member step] and [member page] size. See e.g. [ScrollBar] and [Slider] for examples of higher-level nodes using Range.
	// Range methods that can be overridden by a [Class] that extends it.
	type Range interface {
		//Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
		ValueChanged(new_value gd.Float) 
	}

*/
type Simple [1]classdb.Range
func (Simple) _value_changed(impl func(ptr unsafe.Pointer, new_value float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var new_value = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(new_value))
		gc.End()
	}
}
func (self Simple) GetValue() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetValue()))
}
func (self Simple) GetMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMin()))
}
func (self Simple) GetMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMax()))
}
func (self Simple) GetStep() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetStep()))
}
func (self Simple) GetPage() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPage()))
}
func (self Simple) GetAsRatio() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAsRatio()))
}
func (self Simple) SetValue(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetValue(gd.Float(value))
}
func (self Simple) SetValueNoSignal(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetValueNoSignal(gd.Float(value))
}
func (self Simple) SetMin(minimum float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMin(gd.Float(minimum))
}
func (self Simple) SetMax(maximum float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMax(gd.Float(maximum))
}
func (self Simple) SetStep(step float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStep(gd.Float(step))
}
func (self Simple) SetPage(pagesize float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPage(gd.Float(pagesize))
}
func (self Simple) SetAsRatio(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAsRatio(gd.Float(value))
}
func (self Simple) SetUseRoundedValues(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseRoundedValues(enabled)
}
func (self Simple) IsUsingRoundedValues() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingRoundedValues())
}
func (self Simple) SetExpRatio(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExpRatio(enabled)
}
func (self Simple) IsRatioExp() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRatioExp())
}
func (self Simple) SetAllowGreater(allow bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowGreater(allow)
}
func (self Simple) IsGreaterAllowed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsGreaterAllowed())
}
func (self Simple) SetAllowLesser(allow bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowLesser(allow)
}
func (self Simple) IsLesserAllowed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLesserAllowed())
}
func (self Simple) Share(with [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Share(with)
}
func (self Simple) Unshare() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Unshare()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Range
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Called when the [Range]'s value is changed (following the same conditions as [signal value_changed]).
*/
func (class) _value_changed(impl func(ptr unsafe.Pointer, new_value gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var new_value = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, new_value)
		ctx.End()
	}
}

//go:nosplit
func (self class) GetValue() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_get_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetMin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_get_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetMax() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_get_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetStep() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_get_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetPage() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_get_page, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetAsRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_get_as_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetValue(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [Range]'s current value to the specified [param value], without emitting the [signal value_changed] signal.
*/
//go:nosplit
func (self class) SetValueNoSignal(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_value_no_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMin(minimum gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, minimum)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMax(maximum gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, maximum)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetStep(step gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, step)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_step, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetPage(pagesize gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pagesize)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_page, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAsRatio(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_as_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetUseRoundedValues(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_use_rounded_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingRoundedValues() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_is_using_rounded_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExpRatio(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_exp_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRatioExp() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_is_ratio_exp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowGreater(allow bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_allow_greater, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsGreaterAllowed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_is_greater_allowed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowLesser(allow bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_set_allow_lesser, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLesserAllowed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_is_lesser_allowed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Binds two [Range]s together along with any ranges previously grouped with either of them. When any of range's member variables change, it will share the new value with all other ranges in its group.
*/
//go:nosplit
func (self class) Share(with object.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(with[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_share, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops the [Range] from sharing its member variables with any other.
*/
//go:nosplit
func (self class) Unshare()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Range.Bind_unshare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsRange() Expert { return self[0].AsRange() }


//go:nosplit
func (self Simple) AsRange() Simple { return self[0].AsRange() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_value_changed": return reflect.ValueOf(self._value_changed);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_value_changed": return reflect.ValueOf(self._value_changed);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Range", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
