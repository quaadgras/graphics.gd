package FogMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Material"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A [Material] resource that can be used by [FogVolume]s to draw volumetric effects.
If you need more advanced effects, use a custom [url=$DOCS_URL/tutorials/shaders/shader_reference/fog_shader.html]fog shader[/url].
*/
type Instance [1]classdb.FogMaterial

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.FogMaterial

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FogMaterial"))
	return Instance{classdb.FogMaterial(object)}
}

func (self Instance) Density() float64 {
	return float64(float64(class(self).GetDensity()))
}

func (self Instance) SetDensity(value float64) {
	class(self).SetDensity(gd.Float(value))
}

func (self Instance) Albedo() gd.Color {
	return gd.Color(class(self).GetAlbedo())
}

func (self Instance) SetAlbedo(value gd.Color) {
	class(self).SetAlbedo(value)
}

func (self Instance) Emission() gd.Color {
	return gd.Color(class(self).GetEmission())
}

func (self Instance) SetEmission(value gd.Color) {
	class(self).SetEmission(value)
}

func (self Instance) HeightFalloff() float64 {
	return float64(float64(class(self).GetHeightFalloff()))
}

func (self Instance) SetHeightFalloff(value float64) {
	class(self).SetHeightFalloff(gd.Float(value))
}

func (self Instance) EdgeFade() float64 {
	return float64(float64(class(self).GetEdgeFade()))
}

func (self Instance) SetEdgeFade(value float64) {
	class(self).SetEdgeFade(gd.Float(value))
}

func (self Instance) DensityTexture() gdclass.Texture3D {
	return gdclass.Texture3D(class(self).GetDensityTexture())
}

func (self Instance) SetDensityTexture(value gdclass.Texture3D) {
	class(self).SetDensityTexture(value)
}

//go:nosplit
func (self class) SetDensity(density gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlbedo(albedo gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, albedo)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlbedo() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmission(emission gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, emission)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmission() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeightFalloff(height_falloff gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, height_falloff)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_height_falloff, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeightFalloff() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_height_falloff, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEdgeFade(edge_fade gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, edge_fade)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_edge_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEdgeFade() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_edge_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDensityTexture(density_texture gdclass.Texture3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(density_texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_density_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDensityTexture() gdclass.Texture3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_density_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture3D{classdb.Texture3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsFogMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFogMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("FogMaterial", func(ptr gd.Object) any { return classdb.FogMaterial(ptr) })
}
