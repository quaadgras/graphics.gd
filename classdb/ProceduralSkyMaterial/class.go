// Package ProceduralSkyMaterial provides methods for working with ProceduralSkyMaterial object instances.
package ProceduralSkyMaterial

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Material"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[ProceduralSkyMaterial] provides a way to create an effective background quickly by defining procedural parameters for the sun, the sky and the ground. The sky and ground are defined by a main color, a color at the horizon, and an easing curve to interpolate between them. Suns are described by a position in the sky, a color, and a max angle from the sun at which the easing curve ends. The max angle therefore defines the size of the sun in the sky.
[ProceduralSkyMaterial] supports up to 4 suns, using the color, and energy, direction, and angular distance of the first four [DirectionalLight3D] nodes in the scene. This means that the suns are defined individually by the properties of their corresponding [DirectionalLight3D]s and globally by [member sun_angle_max] and [member sun_curve].
[ProceduralSkyMaterial] uses a lightweight shader to draw the sky and is therefore suited for real-time updates. This makes it a great option for a sky that is simple and computationally cheap, but unrealistic. If you need a more realistic procedural option, use [PhysicalSkyMaterial].
*/
type Instance [1]gdclass.ProceduralSkyMaterial
type Any interface {
	gd.IsClass
	AsProceduralSkyMaterial() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ProceduralSkyMaterial

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ProceduralSkyMaterial"))
	return Instance{*(*gdclass.ProceduralSkyMaterial)(unsafe.Pointer(&object))}
}

func (self Instance) SkyTopColor() Color.RGBA {
	return Color.RGBA(class(self).GetSkyTopColor())
}

func (self Instance) SetSkyTopColor(value Color.RGBA) {
	class(self).SetSkyTopColor(gd.Color(value))
}

func (self Instance) SkyHorizonColor() Color.RGBA {
	return Color.RGBA(class(self).GetSkyHorizonColor())
}

func (self Instance) SetSkyHorizonColor(value Color.RGBA) {
	class(self).SetSkyHorizonColor(gd.Color(value))
}

func (self Instance) SkyCurve() Float.X {
	return Float.X(Float.X(class(self).GetSkyCurve()))
}

func (self Instance) SetSkyCurve(value Float.X) {
	class(self).SetSkyCurve(gd.Float(value))
}

func (self Instance) SkyEnergyMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetSkyEnergyMultiplier()))
}

func (self Instance) SetSkyEnergyMultiplier(value Float.X) {
	class(self).SetSkyEnergyMultiplier(gd.Float(value))
}

func (self Instance) SkyCover() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetSkyCover())
}

func (self Instance) SetSkyCover(value [1]gdclass.Texture2D) {
	class(self).SetSkyCover(value)
}

func (self Instance) SkyCoverModulate() Color.RGBA {
	return Color.RGBA(class(self).GetSkyCoverModulate())
}

func (self Instance) SetSkyCoverModulate(value Color.RGBA) {
	class(self).SetSkyCoverModulate(gd.Color(value))
}

func (self Instance) GroundBottomColor() Color.RGBA {
	return Color.RGBA(class(self).GetGroundBottomColor())
}

func (self Instance) SetGroundBottomColor(value Color.RGBA) {
	class(self).SetGroundBottomColor(gd.Color(value))
}

func (self Instance) GroundHorizonColor() Color.RGBA {
	return Color.RGBA(class(self).GetGroundHorizonColor())
}

func (self Instance) SetGroundHorizonColor(value Color.RGBA) {
	class(self).SetGroundHorizonColor(gd.Color(value))
}

func (self Instance) GroundCurve() Float.X {
	return Float.X(Float.X(class(self).GetGroundCurve()))
}

func (self Instance) SetGroundCurve(value Float.X) {
	class(self).SetGroundCurve(gd.Float(value))
}

func (self Instance) GroundEnergyMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetGroundEnergyMultiplier()))
}

func (self Instance) SetGroundEnergyMultiplier(value Float.X) {
	class(self).SetGroundEnergyMultiplier(gd.Float(value))
}

func (self Instance) SunAngleMax() Float.X {
	return Float.X(Float.X(class(self).GetSunAngleMax()))
}

func (self Instance) SetSunAngleMax(value Float.X) {
	class(self).SetSunAngleMax(gd.Float(value))
}

func (self Instance) SunCurve() Float.X {
	return Float.X(Float.X(class(self).GetSunCurve()))
}

func (self Instance) SetSunCurve(value Float.X) {
	class(self).SetSunCurve(gd.Float(value))
}

func (self Instance) UseDebanding() bool {
	return bool(class(self).GetUseDebanding())
}

func (self Instance) SetUseDebanding(value bool) {
	class(self).SetUseDebanding(value)
}

func (self Instance) EnergyMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetEnergyMultiplier()))
}

func (self Instance) SetEnergyMultiplier(value Float.X) {
	class(self).SetEnergyMultiplier(gd.Float(value))
}

//go:nosplit
func (self class) SetSkyTopColor(color gd.Color) {
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
func (self class) SetSkyHorizonColor(color gd.Color) {
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
func (self class) SetSkyCurve(curve gd.Float) {
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
func (self class) SetSkyEnergyMultiplier(multiplier gd.Float) {
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
func (self class) SetSkyCover(sky_cover [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(sky_cover[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_set_sky_cover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyCover() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProceduralSkyMaterial.Bind_get_sky_cover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyCoverModulate(color gd.Color) {
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
func (self class) SetGroundBottomColor(color gd.Color) {
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
func (self class) SetGroundHorizonColor(color gd.Color) {
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
func (self class) SetGroundCurve(curve gd.Float) {
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
func (self class) SetGroundEnergyMultiplier(energy gd.Float) {
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
func (self class) SetSunAngleMax(degrees gd.Float) {
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
func (self class) SetSunCurve(curve gd.Float) {
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
func (self class) SetUseDebanding(use_debanding bool) {
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
func (self class) SetEnergyMultiplier(multiplier gd.Float) {
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
func (self class) AsProceduralSkyMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsProceduralSkyMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.Advanced {
	return *((*Material.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMaterial() Material.Instance {
	return *((*Material.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsMaterial(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMaterial(), name)
	}
}
func init() {
	gdclass.Register("ProceduralSkyMaterial", func(ptr gd.Object) any {
		return [1]gdclass.ProceduralSkyMaterial{*(*gdclass.ProceduralSkyMaterial)(unsafe.Pointer(&ptr))}
	})
}
