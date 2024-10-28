package ProceduralSkyMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Material"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[ProceduralSkyMaterial] provides a way to create an effective background quickly by defining procedural parameters for the sun, the sky and the ground. The sky and ground are defined by a main color, a color at the horizon, and an easing curve to interpolate between them. Suns are described by a position in the sky, a color, and a max angle from the sun at which the easing curve ends. The max angle therefore defines the size of the sun in the sky.
[ProceduralSkyMaterial] supports up to 4 suns, using the color, and energy, direction, and angular distance of the first four [DirectionalLight3D] nodes in the scene. This means that the suns are defined individually by the properties of their corresponding [DirectionalLight3D]s and globally by [member sun_angle_max] and [member sun_curve].
[ProceduralSkyMaterial] uses a lightweight shader to draw the sky and is therefore suited for real-time updates. This makes it a great option for a sky that is simple and computationally cheap, but unrealistic. If you need a more realistic procedural option, use [PhysicalSkyMaterial].

*/
type Go [1]classdb.ProceduralSkyMaterial
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ProceduralSkyMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ProceduralSkyMaterial"))
	return Go{classdb.ProceduralSkyMaterial(object)}
}

func (self Go) SkyTopColor() gd.Color {
		return gd.Color(class(self).GetSkyTopColor())
}

func (self Go) SetSkyTopColor(value gd.Color) {
	class(self).SetSkyTopColor(value)
}

func (self Go) SkyHorizonColor() gd.Color {
		return gd.Color(class(self).GetSkyHorizonColor())
}

func (self Go) SetSkyHorizonColor(value gd.Color) {
	class(self).SetSkyHorizonColor(value)
}

func (self Go) SkyCurve() float64 {
		return float64(float64(class(self).GetSkyCurve()))
}

func (self Go) SetSkyCurve(value float64) {
	class(self).SetSkyCurve(gd.Float(value))
}

func (self Go) SkyEnergyMultiplier() float64 {
		return float64(float64(class(self).GetSkyEnergyMultiplier()))
}

func (self Go) SetSkyEnergyMultiplier(value float64) {
	class(self).SetSkyEnergyMultiplier(gd.Float(value))
}

func (self Go) SkyCover() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetSkyCover())
}

func (self Go) SetSkyCover(value gdclass.Texture2D) {
	class(self).SetSkyCover(value)
}

func (self Go) SkyCoverModulate() gd.Color {
		return gd.Color(class(self).GetSkyCoverModulate())
}

func (self Go) SetSkyCoverModulate(value gd.Color) {
	class(self).SetSkyCoverModulate(value)
}

func (self Go) GroundBottomColor() gd.Color {
		return gd.Color(class(self).GetGroundBottomColor())
}

func (self Go) SetGroundBottomColor(value gd.Color) {
	class(self).SetGroundBottomColor(value)
}

func (self Go) GroundHorizonColor() gd.Color {
		return gd.Color(class(self).GetGroundHorizonColor())
}

func (self Go) SetGroundHorizonColor(value gd.Color) {
	class(self).SetGroundHorizonColor(value)
}

func (self Go) GroundCurve() float64 {
		return float64(float64(class(self).GetGroundCurve()))
}

func (self Go) SetGroundCurve(value float64) {
	class(self).SetGroundCurve(gd.Float(value))
}

func (self Go) GroundEnergyMultiplier() float64 {
		return float64(float64(class(self).GetGroundEnergyMultiplier()))
}

func (self Go) SetGroundEnergyMultiplier(value float64) {
	class(self).SetGroundEnergyMultiplier(gd.Float(value))
}

func (self Go) SunAngleMax() float64 {
		return float64(float64(class(self).GetSunAngleMax()))
}

func (self Go) SetSunAngleMax(value float64) {
	class(self).SetSunAngleMax(gd.Float(value))
}

func (self Go) SunCurve() float64 {
		return float64(float64(class(self).GetSunCurve()))
}

func (self Go) SetSunCurve(value float64) {
	class(self).SetSunCurve(gd.Float(value))
}

func (self Go) UseDebanding() bool {
		return bool(class(self).GetUseDebanding())
}

func (self Go) SetUseDebanding(value bool) {
	class(self).SetUseDebanding(value)
}

func (self Go) EnergyMultiplier() float64 {
		return float64(float64(class(self).GetEnergyMultiplier()))
}

func (self Go) SetEnergyMultiplier(value float64) {
	class(self).SetEnergyMultiplier(gd.Float(value))
}

//go:nosplit
func (self class) SetSkyTopColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_top_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyTopColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_top_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyHorizonColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_horizon_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyHorizonColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_horizon_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyCurve(curve gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyCurve() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyEnergyMultiplier(multiplier gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyEnergyMultiplier() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyCover(sky_cover gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(sky_cover[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_cover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyCover() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_cover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkyCoverModulate(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_cover_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkyCoverModulate() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_cover_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroundBottomColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_ground_bottom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroundBottomColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_ground_bottom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroundHorizonColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_ground_horizon_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroundHorizonColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_ground_horizon_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroundCurve(curve gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_ground_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroundCurve() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_ground_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroundEnergyMultiplier(energy gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_ground_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroundEnergyMultiplier() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_ground_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSunAngleMax(degrees gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sun_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSunAngleMax() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sun_angle_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSunCurve(curve gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sun_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSunCurve() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sun_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseDebanding(use_debanding bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, use_debanding)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseDebanding() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnergyMultiplier(multiplier gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnergyMultiplier() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsProceduralSkyMaterial() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsProceduralSkyMaterial() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.GD { return *((*Material.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMaterial() Material.Go { return *((*Material.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMaterial(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMaterial(), name)
	}
}
func init() {classdb.Register("ProceduralSkyMaterial", func(ptr gd.Object) any { return classdb.ProceduralSkyMaterial(ptr) })}
