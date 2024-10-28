package GLTFLight

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Represents a light as defined by the [code]KHR_lights_punctual[/code] GLTF extension.

*/
type Go [1]classdb.GLTFLight

/*
Create a new GLTFLight instance from the given Godot [Light3D] node.
*/
func (self Go) FromNode(light_node gdclass.Light3D) gdclass.GLTFLight {
	return gdclass.GLTFLight(class(self).FromNode(light_node))
}

/*
Converts this GLTFLight instance into a Godot [Light3D] node.
*/
func (self Go) ToNode() gdclass.Light3D {
	return gdclass.Light3D(class(self).ToNode())
}

/*
Creates a new GLTFLight instance by parsing the given [Dictionary].
*/
func (self Go) FromDictionary(dictionary gd.Dictionary) gdclass.GLTFLight {
	return gdclass.GLTFLight(class(self).FromDictionary(dictionary))
}

/*
Serializes this GLTFLight instance into a [Dictionary].
*/
func (self Go) ToDictionary() gd.Dictionary {
	return gd.Dictionary(class(self).ToDictionary())
}
func (self Go) GetAdditionalData(extension_name string) gd.Variant {
	return gd.Variant(class(self).GetAdditionalData(gd.NewStringName(extension_name)))
}
func (self Go) SetAdditionalData(extension_name string, additional_data gd.Variant) {
	class(self).SetAdditionalData(gd.NewStringName(extension_name), additional_data)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFLight
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFLight"))
	return Go{classdb.GLTFLight(object)}
}

func (self Go) Color() gd.Color {
		return gd.Color(class(self).GetColor())
}

func (self Go) SetColor(value gd.Color) {
	class(self).SetColor(value)
}

func (self Go) Intensity() float64 {
		return float64(float64(class(self).GetIntensity()))
}

func (self Go) SetIntensity(value float64) {
	class(self).SetIntensity(gd.Float(value))
}

func (self Go) LightType() string {
		return string(class(self).GetLightType().String())
}

func (self Go) SetLightType(value string) {
	class(self).SetLightType(gd.NewString(value))
}

func (self Go) Range() float64 {
		return float64(float64(class(self).GetRange()))
}

func (self Go) SetRange(value float64) {
	class(self).SetRange(gd.Float(value))
}

func (self Go) InnerConeAngle() float64 {
		return float64(float64(class(self).GetInnerConeAngle()))
}

func (self Go) SetInnerConeAngle(value float64) {
	class(self).SetInnerConeAngle(gd.Float(value))
}

func (self Go) OuterConeAngle() float64 {
		return float64(float64(class(self).GetOuterConeAngle()))
}

func (self Go) SetOuterConeAngle(value float64) {
	class(self).SetOuterConeAngle(gd.Float(value))
}

/*
Create a new GLTFLight instance from the given Godot [Light3D] node.
*/
//go:nosplit
func (self class) FromNode(light_node gdclass.Light3D) gdclass.GLTFLight {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(light_node[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.GLTFLight{classdb.GLTFLight(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Converts this GLTFLight instance into a Godot [Light3D] node.
*/
//go:nosplit
func (self class) ToNode() gdclass.Light3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Light3D{classdb.Light3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Creates a new GLTFLight instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(dictionary gd.Dictionary) gdclass.GLTFLight {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(dictionary))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.GLTFLight{classdb.GLTFLight(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Serializes this GLTFLight instance into a [Dictionary].
*/
//go:nosplit
func (self class) ToDictionary() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIntensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIntensity(intensity gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_light_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLightType(light_type gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(light_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_light_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRange() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRange(arange gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, arange)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInnerConeAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_inner_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInnerConeAngle(inner_cone_angle gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, inner_cone_angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_inner_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOuterConeAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_outer_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOuterConeAngle(outer_cone_angle gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, outer_cone_angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_outer_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdditionalData(extension_name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(extension_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(extension_name))
	callframe.Arg(frame, discreet.Get(additional_data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFLight() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFLight() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("GLTFLight", func(ptr gd.Object) any { return classdb.GLTFLight(ptr) })}
