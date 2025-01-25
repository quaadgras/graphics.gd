// Package RDPipelineMultisampleState provides methods for working with RDPipelineMultisampleState object instances.
package RDPipelineMultisampleState

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[RDPipelineMultisampleState] is used to control how multisample or supersample antialiasing is being performed when rendering using [RenderingDevice].
*/
type Instance [1]gdclass.RDPipelineMultisampleState

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDPipelineMultisampleState() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDPipelineMultisampleState

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineMultisampleState"))
	casted := Instance{*(*gdclass.RDPipelineMultisampleState)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SampleCount() gdclass.RenderingDeviceTextureSamples {
	return gdclass.RenderingDeviceTextureSamples(class(self).GetSampleCount())
}

func (self Instance) SetSampleCount(value gdclass.RenderingDeviceTextureSamples) {
	class(self).SetSampleCount(value)
}

func (self Instance) EnableSampleShading() bool {
	return bool(class(self).GetEnableSampleShading())
}

func (self Instance) SetEnableSampleShading(value bool) {
	class(self).SetEnableSampleShading(value)
}

func (self Instance) MinSampleShading() Float.X {
	return Float.X(Float.X(class(self).GetMinSampleShading()))
}

func (self Instance) SetMinSampleShading(value Float.X) {
	class(self).SetMinSampleShading(gd.Float(value))
}

func (self Instance) EnableAlphaToCoverage() bool {
	return bool(class(self).GetEnableAlphaToCoverage())
}

func (self Instance) SetEnableAlphaToCoverage(value bool) {
	class(self).SetEnableAlphaToCoverage(value)
}

func (self Instance) EnableAlphaToOne() bool {
	return bool(class(self).GetEnableAlphaToOne())
}

func (self Instance) SetEnableAlphaToOne(value bool) {
	class(self).SetEnableAlphaToOne(value)
}

func (self Instance) SampleMasks() []int {
	return []int(gd.ArrayAs[[]int](class(self).GetSampleMasks()))
}

func (self Instance) SetSampleMasks(value []int) {
	class(self).SetSampleMasks(gd.NewVariant(value).Interface().(gd.Array))
}

//go:nosplit
func (self class) SetSampleCount(p_member gdclass.RenderingDeviceTextureSamples) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_sample_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSampleCount() gdclass.RenderingDeviceTextureSamples {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceTextureSamples](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_sample_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableSampleShading(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_enable_sample_shading, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableSampleShading() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_enable_sample_shading, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinSampleShading(p_member gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_min_sample_shading, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinSampleShading() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_min_sample_shading, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableAlphaToCoverage(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_enable_alpha_to_coverage, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableAlphaToCoverage() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_enable_alpha_to_coverage, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableAlphaToOne(p_member bool) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_enable_alpha_to_one, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableAlphaToOne() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_enable_alpha_to_one, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSampleMasks(masks gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(masks))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_set_sample_masks, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSampleMasks() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineMultisampleState.Bind_get_sample_masks, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsRDPipelineMultisampleState() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineMultisampleState() Instance {
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
	gdclass.Register("RDPipelineMultisampleState", func(ptr gd.Object) any {
		return [1]gdclass.RDPipelineMultisampleState{*(*gdclass.RDPipelineMultisampleState)(unsafe.Pointer(&ptr))}
	})
}
