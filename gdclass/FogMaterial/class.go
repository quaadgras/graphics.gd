package FogMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

/*
A [Material] resource that can be used by [FogVolume]s to draw volumetric effects.
If you need more advanced effects, use a custom [url=$DOCS_URL/tutorials/shaders/shader_reference/fog_shader.html]fog shader[/url].

*/
type Go [1]classdb.FogMaterial
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.FogMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("FogMaterial"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Density() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDensity()))
}

func (self Go) SetDensity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDensity(gd.Float(value))
}

func (self Go) Albedo() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetAlbedo())
}

func (self Go) SetAlbedo(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAlbedo(value)
}

func (self Go) Emission() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetEmission())
}

func (self Go) SetEmission(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEmission(value)
}

func (self Go) HeightFalloff() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetHeightFalloff()))
}

func (self Go) SetHeightFalloff(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHeightFalloff(gd.Float(value))
}

func (self Go) EdgeFade() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEdgeFade()))
}

func (self Go) SetEdgeFade(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEdgeFade(gd.Float(value))
}

func (self Go) DensityTexture() gdclass.Texture3D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture3D(class(self).GetDensityTexture(gc))
}

func (self Go) SetDensityTexture(value gdclass.Texture3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDensityTexture(value)
}

//go:nosplit
func (self class) SetDensity(density gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_set_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_get_density, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlbedo(albedo gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, albedo)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_set_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlbedo() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_get_albedo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEmission(emission gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, emission)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_set_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEmission() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_get_emission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeightFalloff(height_falloff gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height_falloff)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_set_height_falloff, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeightFalloff() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_get_height_falloff, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEdgeFade(edge_fade gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, edge_fade)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_set_edge_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEdgeFade() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_get_edge_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDensityTexture(density_texture gdclass.Texture3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(density_texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_set_density_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDensityTexture(ctx gd.Lifetime) gdclass.Texture3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FogMaterial.Bind_get_density_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsFogMaterial() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsFogMaterial() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("FogMaterial", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
