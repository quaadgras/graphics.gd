package PanoramaSkyMaterial

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
A resource referenced in a [Sky] that is used to draw a background. [PanoramaSkyMaterial] functions similar to skyboxes in other engines, except it uses an equirectangular sky map instead of a [Cubemap].
Using an HDR panorama is strongly recommended for accurate, high-quality reflections. Godot supports the Radiance HDR ([code].hdr[/code]) and OpenEXR ([code].exr[/code]) image formats for this purpose.
You can use [url=https://danilw.github.io/GLSL-howto/cubemap_to_panorama_js/cubemap_to_panorama.html]this tool[/url] to convert a cubemap to an equirectangular sky map.

*/
type Go [1]classdb.PanoramaSkyMaterial
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PanoramaSkyMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PanoramaSkyMaterial"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Panorama() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetPanorama(gc))
}

func (self Go) SetPanorama(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPanorama(value)
}

func (self Go) Filter() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFilteringEnabled())
}

func (self Go) SetFilter(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFilteringEnabled(value)
}

func (self Go) EnergyMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEnergyMultiplier()))
}

func (self Go) SetEnergyMultiplier(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnergyMultiplier(gd.Float(value))
}

//go:nosplit
func (self class) SetPanorama(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PanoramaSkyMaterial.Bind_set_panorama, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPanorama(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PanoramaSkyMaterial.Bind_get_panorama, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFilteringEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PanoramaSkyMaterial.Bind_set_filtering_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFilteringEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PanoramaSkyMaterial.Bind_is_filtering_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PanoramaSkyMaterial.Bind_set_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnergyMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PanoramaSkyMaterial.Bind_get_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPanoramaSkyMaterial() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPanoramaSkyMaterial() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("PanoramaSkyMaterial", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
