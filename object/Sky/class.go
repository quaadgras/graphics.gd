package Sky

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [Sky] class uses a [Material] to render a 3D environment's background and the light it emits by updating the reflection/radiance cubemaps.

*/
type Simple [1]classdb.Sky
func (self Simple) SetRadianceSize(size classdb.SkyRadianceSize) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRadianceSize(size)
}
func (self Simple) GetRadianceSize() classdb.SkyRadianceSize {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.SkyRadianceSize(Expert(self).GetRadianceSize())
}
func (self Simple) SetProcessMode(mode classdb.SkyProcessMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProcessMode(mode)
}
func (self Simple) GetProcessMode() classdb.SkyProcessMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.SkyProcessMode(Expert(self).GetProcessMode())
}
func (self Simple) SetMaterial(material [1]classdb.Material) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaterial(material)
}
func (self Simple) GetMaterial() [1]classdb.Material {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Material(Expert(self).GetMaterial(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Sky
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) SetMaterial(material object.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sky.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) object.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Sky.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSky() Expert { return self[0].AsSky() }


//go:nosplit
func (self Simple) AsSky() Simple { return self[0].AsSky() }


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
func init() {classdb.Register("Sky", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
