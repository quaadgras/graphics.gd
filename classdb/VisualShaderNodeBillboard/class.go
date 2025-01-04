// Package VisualShaderNodeBillboard provides methods for working with VisualShaderNodeBillboard object instances.
package VisualShaderNodeBillboard

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The output port of this node needs to be connected to [code]Model View Matrix[/code] port of [VisualShaderNodeOutput].
*/
type Instance [1]gdclass.VisualShaderNodeBillboard
type Any interface {
	gd.IsClass
	AsVisualShaderNodeBillboard() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeBillboard

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeBillboard"))
	return Instance{*(*gdclass.VisualShaderNodeBillboard)(unsafe.Pointer(&object))}
}

func (self Instance) BillboardType() gdclass.VisualShaderNodeBillboardBillboardType {
	return gdclass.VisualShaderNodeBillboardBillboardType(class(self).GetBillboardType())
}

func (self Instance) SetBillboardType(value gdclass.VisualShaderNodeBillboardBillboardType) {
	class(self).SetBillboardType(value)
}

func (self Instance) KeepScale() bool {
	return bool(class(self).IsKeepScaleEnabled())
}

func (self Instance) SetKeepScale(value bool) {
	class(self).SetKeepScaleEnabled(value)
}

//go:nosplit
func (self class) SetBillboardType(billboard_type gdclass.VisualShaderNodeBillboardBillboardType) {
	var frame = callframe.New()
	callframe.Arg(frame, billboard_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBillboard.Bind_set_billboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBillboardType() gdclass.VisualShaderNodeBillboardBillboardType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeBillboardBillboardType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBillboard.Bind_get_billboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetKeepScaleEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBillboard.Bind_set_keep_scale_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsKeepScaleEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBillboard.Bind_is_keep_scale_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeBillboard() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeBillboard() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeBillboard", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeBillboard{*(*gdclass.VisualShaderNodeBillboard)(unsafe.Pointer(&ptr))}
	})
}

type BillboardType = gdclass.VisualShaderNodeBillboardBillboardType

const (
	/*Billboarding is disabled and the node does nothing.*/
	BillboardTypeDisabled BillboardType = 0
	/*A standard billboarding algorithm is enabled.*/
	BillboardTypeEnabled BillboardType = 1
	/*A billboarding algorithm to rotate around Y-axis is enabled.*/
	BillboardTypeFixedY BillboardType = 2
	/*A billboarding algorithm designed to use on particles is enabled.*/
	BillboardTypeParticles BillboardType = 3
	/*Represents the size of the [enum BillboardType] enum.*/
	BillboardTypeMax BillboardType = 4
)
