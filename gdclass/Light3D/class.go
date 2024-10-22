package Light3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Light3D is the [i]abstract[/i] base class for light nodes. As it can't be instantiated, it shouldn't be used directly. Other types of light nodes inherit from it. Light3D contains the common variables and parameters used for lighting.

*/
type Go [1]classdb.Light3D

/*
Returns the [Color] of an idealized blackbody at the given [member light_temperature]. This value is calculated internally based on the [member light_temperature]. This [Color] is multiplied by [member light_color] before being sent to the [RenderingServer].
*/
func (self Go) GetCorrelatedColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetCorrelatedColor())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Light3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Light3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) LightIntensityLumens() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(20)))
}

func (self Go) SetLightIntensityLumens(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(20, gd.Float(value))
}

func (self Go) LightIntensityLux() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(20)))
}

func (self Go) SetLightIntensityLux(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(20, gd.Float(value))
}

func (self Go) LightTemperature() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTemperature()))
}

func (self Go) SetLightTemperature(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTemperature(gd.Float(value))
}

func (self Go) LightColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetColor())
}

func (self Go) SetLightColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColor(value)
}

func (self Go) LightEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(0)))
}

func (self Go) SetLightEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(0, gd.Float(value))
}

func (self Go) LightIndirectEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(1)))
}

func (self Go) SetLightIndirectEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(1, gd.Float(value))
}

func (self Go) LightVolumetricFogEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(2)))
}

func (self Go) SetLightVolumetricFogEnergy(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(2, gd.Float(value))
}

func (self Go) LightProjector() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetProjector(gc))
}

func (self Go) SetLightProjector(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProjector(value)
}

func (self Go) LightSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(5)))
}

func (self Go) SetLightSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(5, gd.Float(value))
}

func (self Go) LightAngularDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(5)))
}

func (self Go) SetLightAngularDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(5, gd.Float(value))
}

func (self Go) LightNegative() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsNegative())
}

func (self Go) SetLightNegative(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNegative(value)
}

func (self Go) LightSpecular() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(3)))
}

func (self Go) SetLightSpecular(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(3, gd.Float(value))
}

func (self Go) LightBakeMode() classdb.Light3DBakeMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Light3DBakeMode(class(self).GetBakeMode())
}

func (self Go) SetLightBakeMode(value classdb.Light3DBakeMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBakeMode(value)
}

func (self Go) LightCullMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCullMask()))
}

func (self Go) SetLightCullMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCullMask(gd.Int(value))
}

func (self Go) ShadowEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).HasShadow())
}

func (self Go) SetShadowEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShadow(value)
}

func (self Go) ShadowBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(15)))
}

func (self Go) SetShadowBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(15, gd.Float(value))
}

func (self Go) ShadowNormalBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(14)))
}

func (self Go) SetShadowNormalBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(14, gd.Float(value))
}

func (self Go) ShadowReverseCullFace() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetShadowReverseCullFace())
}

func (self Go) SetShadowReverseCullFace(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShadowReverseCullFace(value)
}

func (self Go) ShadowTransmittanceBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(19)))
}

func (self Go) SetShadowTransmittanceBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(19, gd.Float(value))
}

func (self Go) ShadowOpacity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(17)))
}

func (self Go) SetShadowOpacity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(17, gd.Float(value))
}

func (self Go) ShadowBlur() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetParam(18)))
}

func (self Go) SetShadowBlur(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetParam(18, gd.Float(value))
}

func (self Go) DistanceFadeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDistanceFadeEnabled())
}

func (self Go) SetDistanceFadeEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnableDistanceFade(value)
}

func (self Go) DistanceFadeBegin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDistanceFadeBegin()))
}

func (self Go) SetDistanceFadeBegin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDistanceFadeBegin(gd.Float(value))
}

func (self Go) DistanceFadeShadow() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDistanceFadeShadow()))
}

func (self Go) SetDistanceFadeShadow(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDistanceFadeShadow(gd.Float(value))
}

func (self Go) DistanceFadeLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDistanceFadeLength()))
}

func (self Go) SetDistanceFadeLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDistanceFadeLength(gd.Float(value))
}

func (self Go) EditorOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEditorOnly())
}

func (self Go) SetEditorOnly(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditorOnly(value)
}

//go:nosplit
func (self class) SetEditorOnly(editor_only bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, editor_only)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditorOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_is_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of the specified [enum Light3D.Param] parameter.
*/
//go:nosplit
func (self class) SetParam(param classdb.Light3DParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified [enum Light3D.Param] parameter.
*/
//go:nosplit
func (self class) GetParam(param classdb.Light3DParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadow(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasShadow() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_has_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNegative(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_negative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsNegative() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_is_negative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCullMask(cull_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cull_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCullMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableDistanceFade(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_enable_distance_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDistanceFadeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_is_distance_fade_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDistanceFadeBegin(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDistanceFadeBegin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDistanceFadeShadow(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_distance_fade_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDistanceFadeShadow() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_distance_fade_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDistanceFadeLength(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDistanceFadeLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadowReverseCullFace(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_shadow_reverse_cull_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowReverseCullFace() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_shadow_reverse_cull_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBakeMode(bake_mode classdb.Light3DBakeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bake_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_bake_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakeMode() classdb.Light3DBakeMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Light3DBakeMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_bake_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProjector(projector gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(projector[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_projector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProjector(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_projector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTemperature(temperature gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, temperature)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_set_temperature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTemperature() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_temperature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Color] of an idealized blackbody at the given [member light_temperature]. This value is calculated internally based on the [member light_temperature]. This [Color] is multiplied by [member light_color] before being sent to the [RenderingServer].
*/
//go:nosplit
func (self class) GetCorrelatedColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Light3D.Bind_get_correlated_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLight3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsLight3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}
func init() {classdb.Register("Light3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Param = classdb.Light3DParam

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
type BakeMode = classdb.Light3DBakeMode

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
