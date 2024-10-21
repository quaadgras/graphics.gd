package ProceduralSkyMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Material"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[ProceduralSkyMaterial] provides a way to create an effective background quickly by defining procedural parameters for the sun, the sky and the ground. The sky and ground are defined by a main color, a color at the horizon, and an easing curve to interpolate between them. Suns are described by a position in the sky, a color, and a max angle from the sun at which the easing curve ends. The max angle therefore defines the size of the sun in the sky.
[ProceduralSkyMaterial] supports up to 4 suns, using the color, and energy, direction, and angular distance of the first four [DirectionalLight3D] nodes in the scene. This means that the suns are defined individually by the properties of their corresponding [DirectionalLight3D]s and globally by [member sun_angle_max] and [member sun_curve].
[ProceduralSkyMaterial] uses a lightweight shader to draw the sky and is therefore suited for real-time updates. This makes it a great option for a sky that is simple and computationally cheap, but unrealistic. If you need a more realistic procedural option, use [PhysicalSkyMaterial].

*/
type Simple [1]classdb.ProceduralSkyMaterial
func (self Simple) SetSkyTopColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkyTopColor(color)
}
func (self Simple) GetSkyTopColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetSkyTopColor())
}
func (self Simple) SetSkyHorizonColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkyHorizonColor(color)
}
func (self Simple) GetSkyHorizonColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetSkyHorizonColor())
}
func (self Simple) SetSkyCurve(curve float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkyCurve(gd.Float(curve))
}
func (self Simple) GetSkyCurve() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSkyCurve()))
}
func (self Simple) SetSkyEnergyMultiplier(multiplier float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkyEnergyMultiplier(gd.Float(multiplier))
}
func (self Simple) GetSkyEnergyMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSkyEnergyMultiplier()))
}
func (self Simple) SetSkyCover(sky_cover [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkyCover(sky_cover)
}
func (self Simple) GetSkyCover() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetSkyCover(gc))
}
func (self Simple) SetSkyCoverModulate(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSkyCoverModulate(color)
}
func (self Simple) GetSkyCoverModulate() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetSkyCoverModulate())
}
func (self Simple) SetGroundBottomColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGroundBottomColor(color)
}
func (self Simple) GetGroundBottomColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetGroundBottomColor())
}
func (self Simple) SetGroundHorizonColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGroundHorizonColor(color)
}
func (self Simple) GetGroundHorizonColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetGroundHorizonColor())
}
func (self Simple) SetGroundCurve(curve float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGroundCurve(gd.Float(curve))
}
func (self Simple) GetGroundCurve() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGroundCurve()))
}
func (self Simple) SetGroundEnergyMultiplier(energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGroundEnergyMultiplier(gd.Float(energy))
}
func (self Simple) GetGroundEnergyMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetGroundEnergyMultiplier()))
}
func (self Simple) SetSunAngleMax(degrees float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSunAngleMax(gd.Float(degrees))
}
func (self Simple) GetSunAngleMax() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSunAngleMax()))
}
func (self Simple) SetSunCurve(curve float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSunCurve(gd.Float(curve))
}
func (self Simple) GetSunCurve() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSunCurve()))
}
func (self Simple) SetUseDebanding(use_debanding bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseDebanding(use_debanding)
}
func (self Simple) GetUseDebanding() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseDebanding())
}
func (self Simple) SetEnergyMultiplier(multiplier float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnergyMultiplier(gd.Float(multiplier))
}
func (self Simple) GetEnergyMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEnergyMultiplier()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ProceduralSkyMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetSkyTopColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_sky_top_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyTopColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_sky_top_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyHorizonColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_sky_horizon_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyHorizonColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_sky_horizon_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyCurve(curve gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_sky_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyCurve() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_sky_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyEnergyMultiplier(multiplier gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_sky_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyEnergyMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_sky_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyCover(sky_cover object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(sky_cover[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_sky_cover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyCover(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_sky_cover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyCoverModulate(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_sky_cover_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyCoverModulate() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_sky_cover_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroundBottomColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_ground_bottom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroundBottomColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_ground_bottom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroundHorizonColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_ground_horizon_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroundHorizonColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_ground_horizon_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroundCurve(curve gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_ground_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroundCurve() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_ground_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroundEnergyMultiplier(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_ground_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroundEnergyMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_ground_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSunAngleMax(degrees gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_sun_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSunAngleMax() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_sun_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSunCurve(curve gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_sun_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSunCurve() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_sun_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseDebanding(use_debanding bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_debanding)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseDebanding() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnergyMultiplier(multiplier gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_set_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnergyMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ProceduralSkyMaterial.Bind_get_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsProceduralSkyMaterial() Expert { return self[0].AsProceduralSkyMaterial() }


//go:nosplit
func (self Simple) AsProceduralSkyMaterial() Simple { return self[0].AsProceduralSkyMaterial() }


//go:nosplit
func (self class) AsMaterial() Material.Expert { return self[0].AsMaterial() }


//go:nosplit
func (self Simple) AsMaterial() Material.Simple { return self[0].AsMaterial() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


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
func init() {classdb.Register("ProceduralSkyMaterial", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
