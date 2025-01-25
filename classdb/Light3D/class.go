// Package Light3D provides methods for working with Light3D object instances.
package Light3D

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
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Color"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
Light3D is the [i]abstract[/i] base class for light nodes. As it can't be instantiated, it shouldn't be used directly. Other types of light nodes inherit from it. Light3D contains the common variables and parameters used for lighting.
*/
type Instance [1]gdclass.Light3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLight3D() Instance
}

/*
Returns the [Color] of an idealized blackbody at the given [member light_temperature]. This value is calculated internally based on the [member light_temperature]. This [Color] is multiplied by [member light_color] before being sent to the [RenderingServer].
*/
func (self Instance) GetCorrelatedColor() Color.RGBA { //gd:Light3D.get_correlated_color
	return Color.RGBA(class(self).GetCorrelatedColor())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Light3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Light3D"))
	casted := Instance{*(*gdclass.Light3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) LightIntensityLumens() Float.X {
	return Float.X(Float.X(class(self).GetParam(20)))
}

func (self Instance) SetLightIntensityLumens(value Float.X) {
	class(self).SetParam(20, gd.Float(value))
}

func (self Instance) LightIntensityLux() Float.X {
	return Float.X(Float.X(class(self).GetParam(20)))
}

func (self Instance) SetLightIntensityLux(value Float.X) {
	class(self).SetParam(20, gd.Float(value))
}

func (self Instance) LightTemperature() Float.X {
	return Float.X(Float.X(class(self).GetTemperature()))
}

func (self Instance) SetLightTemperature(value Float.X) {
	class(self).SetTemperature(gd.Float(value))
}

func (self Instance) LightColor() Color.RGBA {
	return Color.RGBA(class(self).GetColor())
}

func (self Instance) SetLightColor(value Color.RGBA) {
	class(self).SetColor(gd.Color(value))
}

func (self Instance) LightEnergy() Float.X {
	return Float.X(Float.X(class(self).GetParam(0)))
}

func (self Instance) SetLightEnergy(value Float.X) {
	class(self).SetParam(0, gd.Float(value))
}

func (self Instance) LightIndirectEnergy() Float.X {
	return Float.X(Float.X(class(self).GetParam(1)))
}

func (self Instance) SetLightIndirectEnergy(value Float.X) {
	class(self).SetParam(1, gd.Float(value))
}

func (self Instance) LightVolumetricFogEnergy() Float.X {
	return Float.X(Float.X(class(self).GetParam(2)))
}

func (self Instance) SetLightVolumetricFogEnergy(value Float.X) {
	class(self).SetParam(2, gd.Float(value))
}

func (self Instance) LightProjector() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetProjector())
}

func (self Instance) SetLightProjector(value [1]gdclass.Texture2D) {
	class(self).SetProjector(value)
}

func (self Instance) LightSize() Float.X {
	return Float.X(Float.X(class(self).GetParam(5)))
}

func (self Instance) SetLightSize(value Float.X) {
	class(self).SetParam(5, gd.Float(value))
}

func (self Instance) LightAngularDistance() Float.X {
	return Float.X(Float.X(class(self).GetParam(5)))
}

func (self Instance) SetLightAngularDistance(value Float.X) {
	class(self).SetParam(5, gd.Float(value))
}

func (self Instance) LightNegative() bool {
	return bool(class(self).IsNegative())
}

func (self Instance) SetLightNegative(value bool) {
	class(self).SetNegative(value)
}

func (self Instance) LightSpecular() Float.X {
	return Float.X(Float.X(class(self).GetParam(3)))
}

func (self Instance) SetLightSpecular(value Float.X) {
	class(self).SetParam(3, gd.Float(value))
}

func (self Instance) LightBakeMode() gdclass.Light3DBakeMode {
	return gdclass.Light3DBakeMode(class(self).GetBakeMode())
}

func (self Instance) SetLightBakeMode(value gdclass.Light3DBakeMode) {
	class(self).SetBakeMode(value)
}

func (self Instance) LightCullMask() int {
	return int(int(class(self).GetCullMask()))
}

func (self Instance) SetLightCullMask(value int) {
	class(self).SetCullMask(gd.Int(value))
}

func (self Instance) ShadowEnabled() bool {
	return bool(class(self).HasShadow())
}

func (self Instance) SetShadowEnabled(value bool) {
	class(self).SetShadow(value)
}

func (self Instance) ShadowBias() Float.X {
	return Float.X(Float.X(class(self).GetParam(15)))
}

func (self Instance) SetShadowBias(value Float.X) {
	class(self).SetParam(15, gd.Float(value))
}

func (self Instance) ShadowNormalBias() Float.X {
	return Float.X(Float.X(class(self).GetParam(14)))
}

func (self Instance) SetShadowNormalBias(value Float.X) {
	class(self).SetParam(14, gd.Float(value))
}

func (self Instance) ShadowReverseCullFace() bool {
	return bool(class(self).GetShadowReverseCullFace())
}

func (self Instance) SetShadowReverseCullFace(value bool) {
	class(self).SetShadowReverseCullFace(value)
}

func (self Instance) ShadowTransmittanceBias() Float.X {
	return Float.X(Float.X(class(self).GetParam(19)))
}

func (self Instance) SetShadowTransmittanceBias(value Float.X) {
	class(self).SetParam(19, gd.Float(value))
}

func (self Instance) ShadowOpacity() Float.X {
	return Float.X(Float.X(class(self).GetParam(17)))
}

func (self Instance) SetShadowOpacity(value Float.X) {
	class(self).SetParam(17, gd.Float(value))
}

func (self Instance) ShadowBlur() Float.X {
	return Float.X(Float.X(class(self).GetParam(18)))
}

func (self Instance) SetShadowBlur(value Float.X) {
	class(self).SetParam(18, gd.Float(value))
}

func (self Instance) DistanceFadeEnabled() bool {
	return bool(class(self).IsDistanceFadeEnabled())
}

func (self Instance) SetDistanceFadeEnabled(value bool) {
	class(self).SetEnableDistanceFade(value)
}

func (self Instance) DistanceFadeBegin() Float.X {
	return Float.X(Float.X(class(self).GetDistanceFadeBegin()))
}

func (self Instance) SetDistanceFadeBegin(value Float.X) {
	class(self).SetDistanceFadeBegin(gd.Float(value))
}

func (self Instance) DistanceFadeShadow() Float.X {
	return Float.X(Float.X(class(self).GetDistanceFadeShadow()))
}

func (self Instance) SetDistanceFadeShadow(value Float.X) {
	class(self).SetDistanceFadeShadow(gd.Float(value))
}

func (self Instance) DistanceFadeLength() Float.X {
	return Float.X(Float.X(class(self).GetDistanceFadeLength()))
}

func (self Instance) SetDistanceFadeLength(value Float.X) {
	class(self).SetDistanceFadeLength(gd.Float(value))
}

func (self Instance) EditorOnly() bool {
	return bool(class(self).IsEditorOnly())
}

func (self Instance) SetEditorOnly(value bool) {
	class(self).SetEditorOnly(value)
}

//go:nosplit
func (self class) SetEditorOnly(editor_only bool) { //gd:Light3D.set_editor_only
	var frame = callframe.New()
	callframe.Arg(frame, editor_only)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_editor_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditorOnly() bool { //gd:Light3D.is_editor_only
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_is_editor_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the value of the specified [enum Light3D.Param] parameter.
*/
//go:nosplit
func (self class) SetParam(param gdclass.Light3DParam, value gd.Float) { //gd:Light3D.set_param
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the value of the specified [enum Light3D.Param] parameter.
*/
//go:nosplit
func (self class) GetParam(param gdclass.Light3DParam) gd.Float { //gd:Light3D.get_param
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadow(enabled bool) { //gd:Light3D.set_shadow
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_shadow, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) HasShadow() bool { //gd:Light3D.has_shadow
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_has_shadow, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNegative(enabled bool) { //gd:Light3D.set_negative
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_negative, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsNegative() bool { //gd:Light3D.is_negative
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_is_negative, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMask(cull_mask gd.Int) { //gd:Light3D.set_cull_mask
	var frame = callframe.New()
	callframe.Arg(frame, cull_mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMask() gd.Int { //gd:Light3D.get_cull_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableDistanceFade(enable bool) { //gd:Light3D.set_enable_distance_fade
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_enable_distance_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDistanceFadeEnabled() bool { //gd:Light3D.is_distance_fade_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_is_distance_fade_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeBegin(distance gd.Float) { //gd:Light3D.set_distance_fade_begin
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeBegin() gd.Float { //gd:Light3D.get_distance_fade_begin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeShadow(distance gd.Float) { //gd:Light3D.set_distance_fade_shadow
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_distance_fade_shadow, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeShadow() gd.Float { //gd:Light3D.get_distance_fade_shadow
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_distance_fade_shadow, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeLength(distance gd.Float) { //gd:Light3D.set_distance_fade_length
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeLength() gd.Float { //gd:Light3D.get_distance_fade_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) { //gd:Light3D.set_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color { //gd:Light3D.get_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowReverseCullFace(enable bool) { //gd:Light3D.set_shadow_reverse_cull_face
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_shadow_reverse_cull_face, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowReverseCullFace() bool { //gd:Light3D.get_shadow_reverse_cull_face
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_shadow_reverse_cull_face, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeMode(bake_mode gdclass.Light3DBakeMode) { //gd:Light3D.set_bake_mode
	var frame = callframe.New()
	callframe.Arg(frame, bake_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_bake_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeMode() gdclass.Light3DBakeMode { //gd:Light3D.get_bake_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Light3DBakeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_bake_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProjector(projector [1]gdclass.Texture2D) { //gd:Light3D.set_projector
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(projector[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_projector, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProjector() [1]gdclass.Texture2D { //gd:Light3D.get_projector
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_projector, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTemperature(temperature gd.Float) { //gd:Light3D.set_temperature
	var frame = callframe.New()
	callframe.Arg(frame, temperature)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_temperature, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTemperature() gd.Float { //gd:Light3D.get_temperature
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_temperature, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Color] of an idealized blackbody at the given [member light_temperature]. This value is calculated internally based on the [member light_temperature]. This [Color] is multiplied by [member light_color] before being sent to the [RenderingServer].
*/
//go:nosplit
func (self class) GetCorrelatedColor() gd.Color { //gd:Light3D.get_correlated_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_correlated_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLight3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLight3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Advanced(self.AsVisualInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Instance(self.AsVisualInstance3D()), name)
	}
}
func init() {
	gdclass.Register("Light3D", func(ptr gd.Object) any { return [1]gdclass.Light3D{*(*gdclass.Light3D)(unsafe.Pointer(&ptr))} })
}

type Param = gdclass.Light3DParam //gd:Light3D.Param

const (
	/*Constant for accessing [member light_energy].*/
	ParamEnergy Param = 0
	/*Constant for accessing [member light_indirect_energy].*/
	ParamIndirectEnergy Param = 1
	/*Constant for accessing [member light_volumetric_fog_energy].*/
	ParamVolumetricFogEnergy Param = 2
	/*Constant for accessing [member light_specular].*/
	ParamSpecular Param = 3
	/*Constant for accessing [member OmniLight3D.omni_range] or [member SpotLight3D.spot_range].*/
	ParamRange Param = 4
	/*Constant for accessing [member light_size].*/
	ParamSize Param = 5
	/*Constant for accessing [member OmniLight3D.omni_attenuation] or [member SpotLight3D.spot_attenuation].*/
	ParamAttenuation Param = 6
	/*Constant for accessing [member SpotLight3D.spot_angle].*/
	ParamSpotAngle Param = 7
	/*Constant for accessing [member SpotLight3D.spot_angle_attenuation].*/
	ParamSpotAttenuation Param = 8
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_max_distance].*/
	ParamShadowMaxDistance Param = 9
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_split_1].*/
	ParamShadowSplit1Offset Param = 10
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_split_2].*/
	ParamShadowSplit2Offset Param = 11
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_split_3].*/
	ParamShadowSplit3Offset Param = 12
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_fade_start].*/
	ParamShadowFadeStart Param = 13
	/*Constant for accessing [member shadow_normal_bias].*/
	ParamShadowNormalBias Param = 14
	/*Constant for accessing [member shadow_bias].*/
	ParamShadowBias Param = 15
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_pancake_size].*/
	ParamShadowPancakeSize Param = 16
	/*Constant for accessing [member shadow_opacity].*/
	ParamShadowOpacity Param = 17
	/*Constant for accessing [member shadow_blur].*/
	ParamShadowBlur Param = 18
	/*Constant for accessing [member shadow_transmittance_bias].*/
	ParamTransmittanceBias Param = 19
	/*Constant for accessing [member light_intensity_lumens] and [member light_intensity_lux]. Only used when [member ProjectSettings.rendering/lights_and_shadows/use_physical_light_units] is [code]true[/code].*/
	ParamIntensity Param = 20
	/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 21
)

type BakeMode = gdclass.Light3DBakeMode //gd:Light3D.BakeMode

const (
	/*Light is ignored when baking. This is the fastest mode, but the light will be taken into account when baking global illumination. This mode should generally be used for dynamic lights that change quickly, as the effect of global illumination is less noticeable on those lights.
	  [b]Note:[/b] Hiding a light does [i]not[/i] affect baking [LightmapGI]. Hiding a light will still affect baking [VoxelGI] and SDFGI (see [member Environment.sdfgi_enabled]).*/
	BakeDisabled BakeMode = 0
	/*Light is taken into account in static baking ([VoxelGI], [LightmapGI], SDFGI ([member Environment.sdfgi_enabled])). The light can be moved around or modified, but its global illumination will not update in real-time. This is suitable for subtle changes (such as flickering torches), but generally not large changes such as toggling a light on and off.
	  [b]Note:[/b] The light is not baked in [LightmapGI] if [member editor_only] is [code]true[/code].*/
	BakeStatic BakeMode = 1
	/*Light is taken into account in dynamic baking ([VoxelGI] and SDFGI ([member Environment.sdfgi_enabled]) only). The light can be moved around or modified with global illumination updating in real-time. The light's global illumination appearance will be slightly different compared to [constant BAKE_STATIC]. This has a greater performance cost compared to [constant BAKE_STATIC]. When using SDFGI, the update speed of dynamic lights is affected by [member ProjectSettings.rendering/global_illumination/sdfgi/frames_to_update_lights].*/
	BakeDynamic BakeMode = 2
)
