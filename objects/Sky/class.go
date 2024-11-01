package Sky

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
The [Sky] class uses a [Material] to render a 3D environment's background and the light it emits by updating the reflection/radiance cubemaps.
*/
type Instance [1]classdb.Sky

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Sky

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Sky"))
	return Instance{classdb.Sky(object)}
}

func (self Instance) SkyMaterial() objects.Material {
	return objects.Material(class(self).GetMaterial())
}

func (self Instance) SetSkyMaterial(value objects.Material) {
	class(self).SetMaterial(value)
}

func (self Instance) ProcessMode() classdb.SkyProcessMode {
	return classdb.SkyProcessMode(class(self).GetProcessMode())
}

func (self Instance) SetProcessMode(value classdb.SkyProcessMode) {
	class(self).SetProcessMode(value)
}

func (self Instance) RadianceSize() classdb.SkyRadianceSize {
	return classdb.SkyRadianceSize(class(self).GetRadianceSize())
}

func (self Instance) SetRadianceSize(value classdb.SkyRadianceSize) {
	class(self).SetRadianceSize(value)
}

//go:nosplit
func (self class) SetRadianceSize(size classdb.SkyRadianceSize) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sky.Bind_set_radiance_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadianceSize() classdb.SkyRadianceSize {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SkyRadianceSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sky.Bind_get_radiance_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProcessMode(mode classdb.SkyProcessMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sky.Bind_set_process_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetProcessMode() classdb.SkyProcessMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SkyProcessMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sky.Bind_get_process_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaterial(material objects.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sky.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaterial() objects.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Sky.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Material{classdb.Material(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsSky() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSky() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("Sky", func(ptr gd.Object) any { return classdb.Sky(ptr) }) }

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
