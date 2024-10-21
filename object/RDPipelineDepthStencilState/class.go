package RDPipelineDepthStencilState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[RDPipelineDepthStencilState] controls the way depth and stencil comparisons are performed when sampling those values using [RenderingDevice].

*/
type Simple [1]classdb.RDPipelineDepthStencilState
func (self Simple) SetEnableDepthTest(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableDepthTest(p_member)
}
func (self Simple) GetEnableDepthTest() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableDepthTest())
}
func (self Simple) SetEnableDepthWrite(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableDepthWrite(p_member)
}
func (self Simple) GetEnableDepthWrite() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableDepthWrite())
}
func (self Simple) SetDepthCompareOperator(p_member classdb.RenderingDeviceCompareOperator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthCompareOperator(p_member)
}
func (self Simple) GetDepthCompareOperator() classdb.RenderingDeviceCompareOperator {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceCompareOperator(Expert(self).GetDepthCompareOperator())
}
func (self Simple) SetEnableDepthRange(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableDepthRange(p_member)
}
func (self Simple) GetEnableDepthRange() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableDepthRange())
}
func (self Simple) SetDepthRangeMin(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthRangeMin(gd.Float(p_member))
}
func (self Simple) GetDepthRangeMin() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDepthRangeMin()))
}
func (self Simple) SetDepthRangeMax(p_member float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepthRangeMax(gd.Float(p_member))
}
func (self Simple) GetDepthRangeMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDepthRangeMax()))
}
func (self Simple) SetEnableStencil(p_member bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableStencil(p_member)
}
func (self Simple) GetEnableStencil() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnableStencil())
}
func (self Simple) SetFrontOpFail(p_member classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrontOpFail(p_member)
}
func (self Simple) GetFrontOpFail() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceStencilOperation(Expert(self).GetFrontOpFail())
}
func (self Simple) SetFrontOpPass(p_member classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrontOpPass(p_member)
}
func (self Simple) GetFrontOpPass() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceStencilOperation(Expert(self).GetFrontOpPass())
}
func (self Simple) SetFrontOpDepthFail(p_member classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrontOpDepthFail(p_member)
}
func (self Simple) GetFrontOpDepthFail() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceStencilOperation(Expert(self).GetFrontOpDepthFail())
}
func (self Simple) SetFrontOpCompare(p_member classdb.RenderingDeviceCompareOperator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrontOpCompare(p_member)
}
func (self Simple) GetFrontOpCompare() classdb.RenderingDeviceCompareOperator {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceCompareOperator(Expert(self).GetFrontOpCompare())
}
func (self Simple) SetFrontOpCompareMask(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrontOpCompareMask(gd.Int(p_member))
}
func (self Simple) GetFrontOpCompareMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFrontOpCompareMask()))
}
func (self Simple) SetFrontOpWriteMask(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrontOpWriteMask(gd.Int(p_member))
}
func (self Simple) GetFrontOpWriteMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFrontOpWriteMask()))
}
func (self Simple) SetFrontOpReference(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrontOpReference(gd.Int(p_member))
}
func (self Simple) GetFrontOpReference() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFrontOpReference()))
}
func (self Simple) SetBackOpFail(p_member classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBackOpFail(p_member)
}
func (self Simple) GetBackOpFail() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceStencilOperation(Expert(self).GetBackOpFail())
}
func (self Simple) SetBackOpPass(p_member classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBackOpPass(p_member)
}
func (self Simple) GetBackOpPass() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceStencilOperation(Expert(self).GetBackOpPass())
}
func (self Simple) SetBackOpDepthFail(p_member classdb.RenderingDeviceStencilOperation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBackOpDepthFail(p_member)
}
func (self Simple) GetBackOpDepthFail() classdb.RenderingDeviceStencilOperation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceStencilOperation(Expert(self).GetBackOpDepthFail())
}
func (self Simple) SetBackOpCompare(p_member classdb.RenderingDeviceCompareOperator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBackOpCompare(p_member)
}
func (self Simple) GetBackOpCompare() classdb.RenderingDeviceCompareOperator {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.RenderingDeviceCompareOperator(Expert(self).GetBackOpCompare())
}
func (self Simple) SetBackOpCompareMask(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBackOpCompareMask(gd.Int(p_member))
}
func (self Simple) GetBackOpCompareMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBackOpCompareMask()))
}
func (self Simple) SetBackOpWriteMask(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBackOpWriteMask(gd.Int(p_member))
}
func (self Simple) GetBackOpWriteMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBackOpWriteMask()))
}
func (self Simple) SetBackOpReference(p_member int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBackOpReference(gd.Int(p_member))
}
func (self Simple) GetBackOpReference() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBackOpReference()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RDPipelineDepthStencilState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsRDPipelineDepthStencilState() Expert { return self[0].AsRDPipelineDepthStencilState() }


//go:nosplit
func (self Simple) AsRDPipelineDepthStencilState() Simple { return self[0].AsRDPipelineDepthStencilState() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RDPipelineDepthStencilState", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
