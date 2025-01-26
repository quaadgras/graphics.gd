// Package RDPipelineDepthStencilState provides methods for working with RDPipelineDepthStencilState object instances.
package RDPipelineDepthStencilState

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
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
var _ Dictionary.Any
var _ RID.Any

/*
[RDPipelineDepthStencilState] controls the way depth and stencil comparisons are performed when sampling those values using [RenderingDevice].
*/
type Instance [1]gdclass.RDPipelineDepthStencilState

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDPipelineDepthStencilState() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDPipelineDepthStencilState

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineDepthStencilState"))
	casted := Instance{*(*gdclass.RDPipelineDepthStencilState)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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

func (self Instance) DepthCompareOperator() gdclass.RenderingDeviceCompareOperator {
	return gdclass.RenderingDeviceCompareOperator(class(self).GetDepthCompareOperator())
}

func (self Instance) SetDepthCompareOperator(value gdclass.RenderingDeviceCompareOperator) {
	class(self).SetDepthCompareOperator(value)
}

func (self Instance) EnableDepthRange() bool {
	return bool(class(self).GetEnableDepthRange())
}

func (self Instance) SetEnableDepthRange(value bool) {
	class(self).SetEnableDepthRange(value)
}

func (self Instance) DepthRangeMin() Float.X {
	return Float.X(Float.X(class(self).GetDepthRangeMin()))
}

func (self Instance) SetDepthRangeMin(value Float.X) {
	class(self).SetDepthRangeMin(gd.Float(value))
}

func (self Instance) DepthRangeMax() Float.X {
	return Float.X(Float.X(class(self).GetDepthRangeMax()))
}

func (self Instance) SetDepthRangeMax(value Float.X) {
	class(self).SetDepthRangeMax(gd.Float(value))
}

func (self Instance) EnableStencil() bool {
	return bool(class(self).GetEnableStencil())
}

func (self Instance) SetEnableStencil(value bool) {
	class(self).SetEnableStencil(value)
}

func (self Instance) FrontOpFail() gdclass.RenderingDeviceStencilOperation {
	return gdclass.RenderingDeviceStencilOperation(class(self).GetFrontOpFail())
}

func (self Instance) SetFrontOpFail(value gdclass.RenderingDeviceStencilOperation) {
	class(self).SetFrontOpFail(value)
}

func (self Instance) FrontOpPass() gdclass.RenderingDeviceStencilOperation {
	return gdclass.RenderingDeviceStencilOperation(class(self).GetFrontOpPass())
}

func (self Instance) SetFrontOpPass(value gdclass.RenderingDeviceStencilOperation) {
	class(self).SetFrontOpPass(value)
}

func (self Instance) FrontOpDepthFail() gdclass.RenderingDeviceStencilOperation {
	return gdclass.RenderingDeviceStencilOperation(class(self).GetFrontOpDepthFail())
}

func (self Instance) SetFrontOpDepthFail(value gdclass.RenderingDeviceStencilOperation) {
	class(self).SetFrontOpDepthFail(value)
}

func (self Instance) FrontOpCompare() gdclass.RenderingDeviceCompareOperator {
	return gdclass.RenderingDeviceCompareOperator(class(self).GetFrontOpCompare())
}

func (self Instance) SetFrontOpCompare(value gdclass.RenderingDeviceCompareOperator) {
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

func (self Instance) BackOpFail() gdclass.RenderingDeviceStencilOperation {
	return gdclass.RenderingDeviceStencilOperation(class(self).GetBackOpFail())
}

func (self Instance) SetBackOpFail(value gdclass.RenderingDeviceStencilOperation) {
	class(self).SetBackOpFail(value)
}

func (self Instance) BackOpPass() gdclass.RenderingDeviceStencilOperation {
	return gdclass.RenderingDeviceStencilOperation(class(self).GetBackOpPass())
}

func (self Instance) SetBackOpPass(value gdclass.RenderingDeviceStencilOperation) {
	class(self).SetBackOpPass(value)
}

func (self Instance) BackOpDepthFail() gdclass.RenderingDeviceStencilOperation {
	return gdclass.RenderingDeviceStencilOperation(class(self).GetBackOpDepthFail())
}

func (self Instance) SetBackOpDepthFail(value gdclass.RenderingDeviceStencilOperation) {
	class(self).SetBackOpDepthFail(value)
}

func (self Instance) BackOpCompare() gdclass.RenderingDeviceCompareOperator {
	return gdclass.RenderingDeviceCompareOperator(class(self).GetBackOpCompare())
}

func (self Instance) SetBackOpCompare(value gdclass.RenderingDeviceCompareOperator) {
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
func (self class) SetEnableDepthTest(p_member bool) { //gd:RDPipelineDepthStencilState.set_enable_depth_test
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_enable_depth_test, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableDepthTest() bool { //gd:RDPipelineDepthStencilState.get_enable_depth_test
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_enable_depth_test, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableDepthWrite(p_member bool) { //gd:RDPipelineDepthStencilState.set_enable_depth_write
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_enable_depth_write, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableDepthWrite() bool { //gd:RDPipelineDepthStencilState.get_enable_depth_write
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_enable_depth_write, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthCompareOperator(p_member gdclass.RenderingDeviceCompareOperator) { //gd:RDPipelineDepthStencilState.set_depth_compare_operator
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_depth_compare_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthCompareOperator() gdclass.RenderingDeviceCompareOperator { //gd:RDPipelineDepthStencilState.get_depth_compare_operator
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceCompareOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_depth_compare_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableDepthRange(p_member bool) { //gd:RDPipelineDepthStencilState.set_enable_depth_range
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_enable_depth_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableDepthRange() bool { //gd:RDPipelineDepthStencilState.get_enable_depth_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_enable_depth_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthRangeMin(p_member gd.Float) { //gd:RDPipelineDepthStencilState.set_depth_range_min
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_depth_range_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthRangeMin() gd.Float { //gd:RDPipelineDepthStencilState.get_depth_range_min
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_depth_range_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthRangeMax(p_member gd.Float) { //gd:RDPipelineDepthStencilState.set_depth_range_max
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_depth_range_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthRangeMax() gd.Float { //gd:RDPipelineDepthStencilState.get_depth_range_max
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_depth_range_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableStencil(p_member bool) { //gd:RDPipelineDepthStencilState.set_enable_stencil
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_enable_stencil, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableStencil() bool { //gd:RDPipelineDepthStencilState.get_enable_stencil
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_enable_stencil, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpFail(p_member gdclass.RenderingDeviceStencilOperation) { //gd:RDPipelineDepthStencilState.set_front_op_fail
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_fail, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpFail() gdclass.RenderingDeviceStencilOperation { //gd:RDPipelineDepthStencilState.get_front_op_fail
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_fail, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpPass(p_member gdclass.RenderingDeviceStencilOperation) { //gd:RDPipelineDepthStencilState.set_front_op_pass
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpPass() gdclass.RenderingDeviceStencilOperation { //gd:RDPipelineDepthStencilState.get_front_op_pass
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpDepthFail(p_member gdclass.RenderingDeviceStencilOperation) { //gd:RDPipelineDepthStencilState.set_front_op_depth_fail
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpDepthFail() gdclass.RenderingDeviceStencilOperation { //gd:RDPipelineDepthStencilState.get_front_op_depth_fail
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpCompare(p_member gdclass.RenderingDeviceCompareOperator) { //gd:RDPipelineDepthStencilState.set_front_op_compare
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_compare, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpCompare() gdclass.RenderingDeviceCompareOperator { //gd:RDPipelineDepthStencilState.get_front_op_compare
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceCompareOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_compare, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpCompareMask(p_member gd.Int) { //gd:RDPipelineDepthStencilState.set_front_op_compare_mask
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpCompareMask() gd.Int { //gd:RDPipelineDepthStencilState.get_front_op_compare_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpWriteMask(p_member gd.Int) { //gd:RDPipelineDepthStencilState.set_front_op_write_mask
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpWriteMask() gd.Int { //gd:RDPipelineDepthStencilState.get_front_op_write_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontOpReference(p_member gd.Int) { //gd:RDPipelineDepthStencilState.set_front_op_reference
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_front_op_reference, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontOpReference() gd.Int { //gd:RDPipelineDepthStencilState.get_front_op_reference
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_front_op_reference, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpFail(p_member gdclass.RenderingDeviceStencilOperation) { //gd:RDPipelineDepthStencilState.set_back_op_fail
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_fail, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpFail() gdclass.RenderingDeviceStencilOperation { //gd:RDPipelineDepthStencilState.get_back_op_fail
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_fail, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpPass(p_member gdclass.RenderingDeviceStencilOperation) { //gd:RDPipelineDepthStencilState.set_back_op_pass
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpPass() gdclass.RenderingDeviceStencilOperation { //gd:RDPipelineDepthStencilState.get_back_op_pass
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpDepthFail(p_member gdclass.RenderingDeviceStencilOperation) { //gd:RDPipelineDepthStencilState.set_back_op_depth_fail
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpDepthFail() gdclass.RenderingDeviceStencilOperation { //gd:RDPipelineDepthStencilState.get_back_op_depth_fail
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceStencilOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_depth_fail, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpCompare(p_member gdclass.RenderingDeviceCompareOperator) { //gd:RDPipelineDepthStencilState.set_back_op_compare
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_compare, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpCompare() gdclass.RenderingDeviceCompareOperator { //gd:RDPipelineDepthStencilState.get_back_op_compare
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceCompareOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_compare, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpCompareMask(p_member gd.Int) { //gd:RDPipelineDepthStencilState.set_back_op_compare_mask
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpCompareMask() gd.Int { //gd:RDPipelineDepthStencilState.get_back_op_compare_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_compare_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpWriteMask(p_member gd.Int) { //gd:RDPipelineDepthStencilState.set_back_op_write_mask
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpWriteMask() gd.Int { //gd:RDPipelineDepthStencilState.get_back_op_write_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_write_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBackOpReference(p_member gd.Int) { //gd:RDPipelineDepthStencilState.set_back_op_reference
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_set_back_op_reference, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBackOpReference() gd.Int { //gd:RDPipelineDepthStencilState.get_back_op_reference
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineDepthStencilState.Bind_get_back_op_reference, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RDPipelineDepthStencilState", func(ptr gd.Object) any {
		return [1]gdclass.RDPipelineDepthStencilState{*(*gdclass.RDPipelineDepthStencilState)(unsafe.Pointer(&ptr))}
	})
}
