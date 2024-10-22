package Sky

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [Sky] class uses a [Material] to render a 3D environment's background and the light it emits by updating the reflection/radiance cubemaps.

*/
type Go [1]classdb.Sky
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Sky
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Sky"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) SkyMaterial() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Material(class(self).GetMaterial(gc))
}

func (self Go) SetSkyMaterial(value gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaterial(value)
}

func (self Go) ProcessMode() classdb.SkyProcessMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.SkyProcessMode(class(self).GetProcessMode())
}

func (self Go) SetProcessMode(value classdb.SkyProcessMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProcessMode(value)
}

func (self Go) RadianceSize() classdb.SkyRadianceSize {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.SkyRadianceSize(class(self).GetRadianceSize())
}

func (self Go) SetRadianceSize(value classdb.SkyRadianceSize) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRadianceSize(value)
}

//go:nosplit
func (self class) SetRadianceSize(size classdb.SkyRadianceSize)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sky.Bind_set_radiance_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadianceSize() classdb.SkyRadianceSize {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SkyRadianceSize](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sky.Bind_get_radiance_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProcessMode(mode classdb.SkyProcessMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sky.Bind_set_process_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProcessMode() classdb.SkyProcessMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SkyProcessMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sky.Bind_get_process_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaterial(material gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sky.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sky.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsSky() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSky() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("Sky", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type RadianceSize = classdb.SkyRadianceSize

const (
/*Radiance texture size is 32×32 pixels.*/
	RadianceSize32 RadianceSize = 0
/*Radiance texture size is 64×64 pixels.*/
	RadianceSize64 RadianceSize = 1
/*Radiance texture size is 128×128 pixels.*/
	RadianceSize128 RadianceSize = 2
/*Radiance texture size is 256×256 pixels.*/
	RadianceSize256 RadianceSize = 3
/*Radiance texture size is 512×512 pixels.*/
	RadianceSize512 RadianceSize = 4
/*Radiance texture size is 1024×1024 pixels.*/
	RadianceSize1024 RadianceSize = 5
/*Radiance texture size is 2048×2048 pixels.*/
	RadianceSize2048 RadianceSize = 6
/*Represents the size of the [enum RadianceSize] enum.*/
	RadianceSizeMax RadianceSize = 7
)
type ProcessMode = classdb.SkyProcessMode

const (
/*Automatically selects the appropriate process mode based on your sky shader. If your shader uses [code]TIME[/code] or [code]POSITION[/code], this will use [constant PROCESS_MODE_REALTIME]. If your shader uses any of the [code]LIGHT_*[/code] variables or any custom uniforms, this uses [constant PROCESS_MODE_INCREMENTAL]. Otherwise, this defaults to [constant PROCESS_MODE_QUALITY].*/
	ProcessModeAutomatic ProcessMode = 0
/*Uses high quality importance sampling to process the radiance map. In general, this results in much higher quality than [constant PROCESS_MODE_REALTIME] but takes much longer to generate. This should not be used if you plan on changing the sky at runtime. If you are finding that the reflection is not blurry enough and is showing sparkles or fireflies, try increasing [member ProjectSettings.rendering/reflections/sky_reflections/ggx_samples].*/
	ProcessModeQuality ProcessMode = 1
/*Uses the same high quality importance sampling to process the radiance map as [constant PROCESS_MODE_QUALITY], but updates over several frames. The number of frames is determined by [member ProjectSettings.rendering/reflections/sky_reflections/roughness_layers]. Use this when you need highest quality radiance maps, but have a sky that updates slowly.*/
	ProcessModeIncremental ProcessMode = 2
/*Uses the fast filtering algorithm to process the radiance map. In general this results in lower quality, but substantially faster run times. If you need better quality, but still need to update the sky every frame, consider turning on [member ProjectSettings.rendering/reflections/sky_reflections/fast_filter_high_quality].
[b]Note:[/b] The fast filtering algorithm is limited to 256×256 cubemaps, so [member radiance_size] must be set to [constant RADIANCE_SIZE_256]. Otherwise, a warning is printed and the overridden radiance size is ignored.*/
	ProcessModeRealtime ProcessMode = 3
)
