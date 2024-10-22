package RDPipelineDepthStencilState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[RDPipelineDepthStencilState] controls the way depth and stencil comparisons are performed when sampling those values using [RenderingDevice].

*/
type Go [1]classdb.RDPipelineDepthStencilState
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDPipelineDepthStencilState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RDPipelineDepthStencilState"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) EnableDepthTest() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableDepthTest())
}

func (self Go) SetEnableDepthTest(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableDepthTest(value)
}

func (self Go) EnableDepthWrite() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableDepthWrite())
}

func (self Go) SetEnableDepthWrite(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableDepthWrite(value)
}

func (self Go) DepthCompareOperator() classdb.RenderingDeviceCompareOperator {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceCompareOperator(class(self).GetDepthCompareOperator())
}

func (self Go) SetDepthCompareOperator(value classdb.RenderingDeviceCompareOperator) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepthCompareOperator(value)
}

func (self Go) EnableDepthRange() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableDepthRange())
}

func (self Go) SetEnableDepthRange(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableDepthRange(value)
}

func (self Go) DepthRangeMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDepthRangeMin()))
}

func (self Go) SetDepthRangeMin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepthRangeMin(gd.Float(value))
}

func (self Go) DepthRangeMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDepthRangeMax()))
}

func (self Go) SetDepthRangeMax(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDepthRangeMax(gd.Float(value))
}

func (self Go) EnableStencil() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetEnableStencil())
}

func (self Go) SetEnableStencil(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableStencil(value)
}

func (self Go) FrontOpFail() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceStencilOperation(class(self).GetFrontOpFail())
}

func (self Go) SetFrontOpFail(value classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrontOpFail(value)
}

func (self Go) FrontOpPass() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceStencilOperation(class(self).GetFrontOpPass())
}

func (self Go) SetFrontOpPass(value classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrontOpPass(value)
}

func (self Go) FrontOpDepthFail() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceStencilOperation(class(self).GetFrontOpDepthFail())
}

func (self Go) SetFrontOpDepthFail(value classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrontOpDepthFail(value)
}

func (self Go) FrontOpCompare() classdb.RenderingDeviceCompareOperator {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceCompareOperator(class(self).GetFrontOpCompare())
}

func (self Go) SetFrontOpCompare(value classdb.RenderingDeviceCompareOperator) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrontOpCompare(value)
}

func (self Go) FrontOpCompareMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFrontOpCompareMask()))
}

func (self Go) SetFrontOpCompareMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrontOpCompareMask(gd.Int(value))
}

func (self Go) FrontOpWriteMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFrontOpWriteMask()))
}

func (self Go) SetFrontOpWriteMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrontOpWriteMask(gd.Int(value))
}

func (self Go) FrontOpReference() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFrontOpReference()))
}

func (self Go) SetFrontOpReference(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrontOpReference(gd.Int(value))
}

func (self Go) BackOpFail() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceStencilOperation(class(self).GetBackOpFail())
}

func (self Go) SetBackOpFail(value classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBackOpFail(value)
}

func (self Go) BackOpPass() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceStencilOperation(class(self).GetBackOpPass())
}

func (self Go) SetBackOpPass(value classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBackOpPass(value)
}

func (self Go) BackOpDepthFail() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceStencilOperation(class(self).GetBackOpDepthFail())
}

func (self Go) SetBackOpDepthFail(value classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBackOpDepthFail(value)
}

func (self Go) BackOpCompare() classdb.RenderingDeviceCompareOperator {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RenderingDeviceCompareOperator(class(self).GetBackOpCompare())
}

func (self Go) SetBackOpCompare(value classdb.RenderingDeviceCompareOperator) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBackOpCompare(value)
}

func (self Go) BackOpCompareMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBackOpCompareMask()))
}

func (self Go) SetBackOpCompareMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBackOpCompareMask(gd.Int(value))
}

func (self Go) BackOpWriteMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBackOpWriteMask()))
}

func (self Go) SetBackOpWriteMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBackOpWriteMask(gd.Int(value))
}

func (self Go) BackOpReference() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBackOpReference()))
}

func (self Go) SetBackOpReference(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBackOpReference(gd.Int(value))
}

//go:nosplit
func (self class) SetEnableDepthTest(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_enable_depth_test, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableDepthTest() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_enable_depth_test, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableDepthWrite(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_enable_depth_write, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableDepthWrite() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_enable_depth_write, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthCompareOperator(p_member classdb.RenderingDeviceCompareOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_depth_compare_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthCompareOperator() classdb.RenderingDeviceCompareOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceCompareOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_depth_compare_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableDepthRange(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_enable_depth_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableDepthRange() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_enable_depth_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthRangeMin(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_depth_range_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthRangeMin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_depth_range_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthRangeMax(p_member gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_depth_range_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthRangeMax() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_depth_range_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableStencil(p_member bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_enable_stencil, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnableStencil() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_enable_stencil, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrontOpFail(p_member classdb.RenderingDeviceStencilOperation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_front_op_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrontOpFail() classdb.RenderingDeviceStencilOperation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_front_op_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrontOpPass(p_member classdb.RenderingDeviceStencilOperation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_front_op_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrontOpPass() classdb.RenderingDeviceStencilOperation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_front_op_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrontOpDepthFail(p_member classdb.RenderingDeviceStencilOperation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_front_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrontOpDepthFail() classdb.RenderingDeviceStencilOperation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_front_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrontOpCompare(p_member classdb.RenderingDeviceCompareOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_front_op_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrontOpCompare() classdb.RenderingDeviceCompareOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceCompareOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_front_op_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrontOpCompareMask(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_front_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrontOpCompareMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_front_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrontOpWriteMask(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_front_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrontOpWriteMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_front_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrontOpReference(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_front_op_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrontOpReference() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_front_op_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBackOpFail(p_member classdb.RenderingDeviceStencilOperation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_back_op_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBackOpFail() classdb.RenderingDeviceStencilOperation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_back_op_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBackOpPass(p_member classdb.RenderingDeviceStencilOperation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_back_op_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBackOpPass() classdb.RenderingDeviceStencilOperation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_back_op_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBackOpDepthFail(p_member classdb.RenderingDeviceStencilOperation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_back_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBackOpDepthFail() classdb.RenderingDeviceStencilOperation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_back_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBackOpCompare(p_member classdb.RenderingDeviceCompareOperator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_back_op_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBackOpCompare() classdb.RenderingDeviceCompareOperator {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceCompareOperator](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_back_op_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBackOpCompareMask(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_back_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBackOpCompareMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_back_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBackOpWriteMask(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_back_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBackOpWriteMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_back_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBackOpReference(p_member gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_set_back_op_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBackOpReference() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RDPipelineDepthStencilState.Bind_get_back_op_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDPipelineDepthStencilState() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDPipelineDepthStencilState() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("RDPipelineDepthStencilState", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
