package RDPipelineDepthStencilState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[RDPipelineDepthStencilState] controls the way depth and stencil comparisons are performed when sampling those values using [RenderingDevice].
*/
type Instance [1]classdb.RDPipelineDepthStencilState

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RDPipelineDepthStencilState

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineDepthStencilState"))
	return Instance{classdb.RDPipelineDepthStencilState(object)}
}

func (self Instance) EnableDepthTest() bool {
	return bool(class(self).GetEnableDepthTest())
}

func (self Instance) SetEnableDepthTest(value bool) {
	class(self).SetEnableDepthTest(value)
}

func (self Instance) EnableDepthWrite() bool {
	return bool(class(self).GetEnableDepthWrite())
}

func (self Instance) SetEnableDepthWrite(value bool) {
	class(self).SetEnableDepthWrite(value)
}

func (self Instance) DepthCompareOperator() classdb.RenderingDeviceCompareOperator {
	return classdb.RenderingDeviceCompareOperator(class(self).GetDepthCompareOperator())
}

func (self Instance) SetDepthCompareOperator(value classdb.RenderingDeviceCompareOperator) {
	class(self).SetDepthCompareOperator(value)
}

func (self Instance) EnableDepthRange() bool {
	return bool(class(self).GetEnableDepthRange())
}

func (self Instance) SetEnableDepthRange(value bool) {
	class(self).SetEnableDepthRange(value)
}

func (self Instance) DepthRangeMin() float64 {
	return float64(float64(class(self).GetDepthRangeMin()))
}

func (self Instance) SetDepthRangeMin(value float64) {
	class(self).SetDepthRangeMin(gd.Float(value))
}

func (self Instance) DepthRangeMax() float64 {
	return float64(float64(class(self).GetDepthRangeMax()))
}

func (self Instance) SetDepthRangeMax(value float64) {
	class(self).SetDepthRangeMax(gd.Float(value))
}

func (self Instance) EnableStencil() bool {
	return bool(class(self).GetEnableStencil())
}

func (self Instance) SetEnableStencil(value bool) {
	class(self).SetEnableStencil(value)
}

func (self Instance) FrontOpFail() classdb.RenderingDeviceStencilOperation {
	return classdb.RenderingDeviceStencilOperation(class(self).GetFrontOpFail())
}

func (self Instance) SetFrontOpFail(value classdb.RenderingDeviceStencilOperation) {
	class(self).SetFrontOpFail(value)
}

func (self Instance) FrontOpPass() classdb.RenderingDeviceStencilOperation {
	return classdb.RenderingDeviceStencilOperation(class(self).GetFrontOpPass())
}

func (self Instance) SetFrontOpPass(value classdb.RenderingDeviceStencilOperation) {
	class(self).SetFrontOpPass(value)
}

func (self Instance) FrontOpDepthFail() classdb.RenderingDeviceStencilOperation {
	return classdb.RenderingDeviceStencilOperation(class(self).GetFrontOpDepthFail())
}

func (self Instance) SetFrontOpDepthFail(value classdb.RenderingDeviceStencilOperation) {
	class(self).SetFrontOpDepthFail(value)
}

func (self Instance) FrontOpCompare() classdb.RenderingDeviceCompareOperator {
	return classdb.RenderingDeviceCompareOperator(class(self).GetFrontOpCompare())
}

func (self Instance) SetFrontOpCompare(value classdb.RenderingDeviceCompareOperator) {
	class(self).SetFrontOpCompare(value)
}

func (self Instance) FrontOpCompareMask() int {
	return int(int(class(self).GetFrontOpCompareMask()))
}

func (self Instance) SetFrontOpCompareMask(value int) {
	class(self).SetFrontOpCompareMask(gd.Int(value))
}

func (self Instance) FrontOpWriteMask() int {
	return int(int(class(self).GetFrontOpWriteMask()))
}

func (self Instance) SetFrontOpWriteMask(value int) {
	class(self).SetFrontOpWriteMask(gd.Int(value))
}

func (self Instance) FrontOpReference() int {
	return int(int(class(self).GetFrontOpReference()))
}

func (self Instance) SetFrontOpReference(value int) {
	class(self).SetFrontOpReference(gd.Int(value))
}

func (self Instance) BackOpFail() classdb.RenderingDeviceStencilOperation {
	return classdb.RenderingDeviceStencilOperation(class(self).GetBackOpFail())
}

func (self Instance) SetBackOpFail(value classdb.RenderingDeviceStencilOperation) {
	class(self).SetBackOpFail(value)
}

func (self Instance) BackOpPass() classdb.RenderingDeviceStencilOperation {
	return classdb.RenderingDeviceStencilOperation(class(self).GetBackOpPass())
}

func (self Instance) SetBackOpPass(value classdb.RenderingDeviceStencilOperation) {
	class(self).SetBackOpPass(value)
}

func (self Instance) BackOpDepthFail() classdb.RenderingDeviceStencilOperation {
	return classdb.RenderingDeviceStencilOperation(class(self).GetBackOpDepthFail())
}

func (self Instance) SetBackOpDepthFail(value classdb.RenderingDeviceStencilOperation) {
	class(self).SetBackOpDepthFail(value)
}

func (self Instance) BackOpCompare() classdb.RenderingDeviceCompareOperator {
	return classdb.RenderingDeviceCompareOperator(class(self).GetBackOpCompare())
}

func (self Instance) SetBackOpCompare(value classdb.RenderingDeviceCompareOperator) {
	class(self).SetBackOpCompare(value)
}

func (self Instance) BackOpCompareMask() int {
	return int(int(class(self).GetBackOpCompareMask()))
}

func (self Instance) SetBackOpCompareMask(value int) {
	class(self).SetBackOpCompareMask(gd.Int(value))
}

func (self Instance) BackOpWriteMask() int {
	return int(int(class(self).GetBackOpWriteMask()))
}

func (self Instance) SetBackOpWriteMask(value int) {
	class(self).SetBackOpWriteMask(gd.Int(value))
}

func (self Instance) BackOpReference() int {
	return int(int(class(self).GetBackOpReference()))
}

func (self Instance) SetBackOpReference(value int) {
	class(self).SetBackOpReference(gd.Int(value))
}

//go:nosplit
func (self class) SetEnableDepthTest(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_enable_depth_test, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableDepthTest() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_enable_depth_test, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableDepthWrite(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_enable_depth_write, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableDepthWrite() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_enable_depth_write, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthCompareOperator(p_member classdb.RenderingDeviceCompareOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_depth_compare_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthCompareOperator() classdb.RenderingDeviceCompareOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceCompareOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_depth_compare_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableDepthRange(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_enable_depth_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableDepthRange() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_enable_depth_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthRangeMin(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_depth_range_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthRangeMin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_depth_range_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthRangeMax(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_depth_range_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthRangeMax() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_depth_range_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableStencil(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_enable_stencil, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableStencil() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_enable_stencil, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpFail(p_member classdb.RenderingDeviceStencilOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpFail() classdb.RenderingDeviceStencilOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpPass(p_member classdb.RenderingDeviceStencilOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpPass() classdb.RenderingDeviceStencilOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpDepthFail(p_member classdb.RenderingDeviceStencilOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpDepthFail() classdb.RenderingDeviceStencilOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpCompare(p_member classdb.RenderingDeviceCompareOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpCompare() classdb.RenderingDeviceCompareOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceCompareOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpCompareMask(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpCompareMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpWriteMask(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpWriteMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpReference(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpReference() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpFail(p_member classdb.RenderingDeviceStencilOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpFail() classdb.RenderingDeviceStencilOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpPass(p_member classdb.RenderingDeviceStencilOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpPass() classdb.RenderingDeviceStencilOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpDepthFail(p_member classdb.RenderingDeviceStencilOperation) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpDepthFail() classdb.RenderingDeviceStencilOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpCompare(p_member classdb.RenderingDeviceCompareOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpCompare() classdb.RenderingDeviceCompareOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceCompareOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_compare, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpCompareMask(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpCompareMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpWriteMask(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpWriteMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpReference(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpReference() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDPipelineDepthStencilState() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineDepthStencilState() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("RDPipelineDepthStencilState", func(ptr gd.Object) any { return classdb.RDPipelineDepthStencilState(ptr) })
}
