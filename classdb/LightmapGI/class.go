// Package LightmapGI provides methods for working with LightmapGI object instances.
package LightmapGI

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Color"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

/*
The [LightmapGI] node is used to compute and store baked lightmaps. Lightmaps are used to provide high-quality indirect lighting with very little light leaking. [LightmapGI] can also provide rough reflections using spherical harmonics if [member directional] is enabled. Dynamic objects can receive indirect lighting thanks to [i]light probes[/i], which can be automatically placed by setting [member generate_probes_subdiv] to a value other than [constant GENERATE_PROBES_DISABLED]. Additional lightmap probes can also be added by creating [LightmapProbe] nodes. The downside is that lightmaps are fully static and cannot be baked in an exported project. Baking a [LightmapGI] node is also slower compared to [VoxelGI].
[b]Procedural generation:[/b] Lightmap baking functionality is only available in the editor. This means [LightmapGI] is not suited to procedurally generated or user-built levels. For procedurally generated or user-built levels, use [VoxelGI] or SDFGI instead (see [member Environment.sdfgi_enabled]).
[b]Performance:[/b] [LightmapGI] provides the best possible run-time performance for global illumination. It is suitable for low-end hardware including integrated graphics and mobile devices.
[b]Note:[/b] Due to how lightmaps work, most properties only have a visible effect once lightmaps are baked again.
[b]Note:[/b] Lightmap baking on [CSGShape3D]s and [PrimitiveMesh]es is not supported, as these cannot store UV2 data required for baking.
[b]Note:[/b] If no custom lightmappers are installed, [LightmapGI] can only be baked from devices that support the Forward+ or Mobile rendering backends.
*/
type Instance [1]gdclass.LightmapGI

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLightmapGI() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.LightmapGI

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("LightmapGI"))
	casted := Instance{*(*gdclass.LightmapGI)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Quality() gdclass.LightmapGIBakeQuality {
	return gdclass.LightmapGIBakeQuality(class(self).GetBakeQuality())
}

func (self Instance) SetQuality(value gdclass.LightmapGIBakeQuality) {
	class(self).SetBakeQuality(value)
}

func (self Instance) Bounces() int {
	return int(int(class(self).GetBounces()))
}

func (self Instance) SetBounces(value int) {
	class(self).SetBounces(gd.Int(value))
}

func (self Instance) BounceIndirectEnergy() Float.X {
	return Float.X(Float.X(class(self).GetBounceIndirectEnergy()))
}

func (self Instance) SetBounceIndirectEnergy(value Float.X) {
	class(self).SetBounceIndirectEnergy(gd.Float(value))
}

func (self Instance) Directional() bool {
	return bool(class(self).IsDirectional())
}

func (self Instance) SetDirectional(value bool) {
	class(self).SetDirectional(value)
}

func (self Instance) UseTextureForBounces() bool {
	return bool(class(self).IsUsingTextureForBounces())
}

func (self Instance) SetUseTextureForBounces(value bool) {
	class(self).SetUseTextureForBounces(value)
}

func (self Instance) Interior() bool {
	return bool(class(self).IsInterior())
}

func (self Instance) SetInterior(value bool) {
	class(self).SetInterior(value)
}

func (self Instance) UseDenoiser() bool {
	return bool(class(self).IsUsingDenoiser())
}

func (self Instance) SetUseDenoiser(value bool) {
	class(self).SetUseDenoiser(value)
}

func (self Instance) DenoiserStrength() Float.X {
	return Float.X(Float.X(class(self).GetDenoiserStrength()))
}

func (self Instance) SetDenoiserStrength(value Float.X) {
	class(self).SetDenoiserStrength(gd.Float(value))
}

func (self Instance) DenoiserRange() int {
	return int(int(class(self).GetDenoiserRange()))
}

func (self Instance) SetDenoiserRange(value int) {
	class(self).SetDenoiserRange(gd.Int(value))
}

func (self Instance) Bias() Float.X {
	return Float.X(Float.X(class(self).GetBias()))
}

func (self Instance) SetBias(value Float.X) {
	class(self).SetBias(gd.Float(value))
}

func (self Instance) TexelScale() Float.X {
	return Float.X(Float.X(class(self).GetTexelScale()))
}

func (self Instance) SetTexelScale(value Float.X) {
	class(self).SetTexelScale(gd.Float(value))
}

func (self Instance) MaxTextureSize() int {
	return int(int(class(self).GetMaxTextureSize()))
}

func (self Instance) SetMaxTextureSize(value int) {
	class(self).SetMaxTextureSize(gd.Int(value))
}

func (self Instance) EnvironmentMode() gdclass.LightmapGIEnvironmentMode {
	return gdclass.LightmapGIEnvironmentMode(class(self).GetEnvironmentMode())
}

func (self Instance) SetEnvironmentMode(value gdclass.LightmapGIEnvironmentMode) {
	class(self).SetEnvironmentMode(value)
}

func (self Instance) EnvironmentCustomSky() [1]gdclass.Sky {
	return [1]gdclass.Sky(class(self).GetEnvironmentCustomSky())
}

func (self Instance) SetEnvironmentCustomSky(value [1]gdclass.Sky) {
	class(self).SetEnvironmentCustomSky(value)
}

func (self Instance) EnvironmentCustomColor() Color.RGBA {
	return Color.RGBA(class(self).GetEnvironmentCustomColor())
}

func (self Instance) SetEnvironmentCustomColor(value Color.RGBA) {
	class(self).SetEnvironmentCustomColor(gd.Color(value))
}

func (self Instance) EnvironmentCustomEnergy() Float.X {
	return Float.X(Float.X(class(self).GetEnvironmentCustomEnergy()))
}

func (self Instance) SetEnvironmentCustomEnergy(value Float.X) {
	class(self).SetEnvironmentCustomEnergy(gd.Float(value))
}

func (self Instance) CameraAttributes() [1]gdclass.CameraAttributes {
	return [1]gdclass.CameraAttributes(class(self).GetCameraAttributes())
}

func (self Instance) SetCameraAttributes(value [1]gdclass.CameraAttributes) {
	class(self).SetCameraAttributes(value)
}

func (self Instance) GenerateProbesSubdiv() gdclass.LightmapGIGenerateProbes {
	return gdclass.LightmapGIGenerateProbes(class(self).GetGenerateProbes())
}

func (self Instance) SetGenerateProbesSubdiv(value gdclass.LightmapGIGenerateProbes) {
	class(self).SetGenerateProbes(value)
}

func (self Instance) LightData() [1]gdclass.LightmapGIData {
	return [1]gdclass.LightmapGIData(class(self).GetLightData())
}

func (self Instance) SetLightData(value [1]gdclass.LightmapGIData) {
	class(self).SetLightData(value)
}

//go:nosplit
func (self class) SetLightData(data [1]gdclass.LightmapGIData) { //gd:LightmapGI.set_light_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_light_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLightData() [1]gdclass.LightmapGIData { //gd:LightmapGI.get_light_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_light_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.LightmapGIData{gd.PointerWithOwnershipTransferredToGo[gdclass.LightmapGIData](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeQuality(bake_quality gdclass.LightmapGIBakeQuality) { //gd:LightmapGI.set_bake_quality
	var frame = callframe.New()
	callframe.Arg(frame, bake_quality)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_bake_quality, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeQuality() gdclass.LightmapGIBakeQuality { //gd:LightmapGI.get_bake_quality
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.LightmapGIBakeQuality](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_bake_quality, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBounces(bounces gd.Int) { //gd:LightmapGI.set_bounces
	var frame = callframe.New()
	callframe.Arg(frame, bounces)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_bounces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBounces() gd.Int { //gd:LightmapGI.get_bounces
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_bounces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBounceIndirectEnergy(bounce_indirect_energy gd.Float) { //gd:LightmapGI.set_bounce_indirect_energy
	var frame = callframe.New()
	callframe.Arg(frame, bounce_indirect_energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_bounce_indirect_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBounceIndirectEnergy() gd.Float { //gd:LightmapGI.get_bounce_indirect_energy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_bounce_indirect_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGenerateProbes(subdivision gdclass.LightmapGIGenerateProbes) { //gd:LightmapGI.set_generate_probes
	var frame = callframe.New()
	callframe.Arg(frame, subdivision)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_generate_probes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGenerateProbes() gdclass.LightmapGIGenerateProbes { //gd:LightmapGI.get_generate_probes
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.LightmapGIGenerateProbes](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_generate_probes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBias(bias gd.Float) { //gd:LightmapGI.set_bias
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBias() gd.Float { //gd:LightmapGI.get_bias
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnvironmentMode(mode gdclass.LightmapGIEnvironmentMode) { //gd:LightmapGI.set_environment_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_environment_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironmentMode() gdclass.LightmapGIEnvironmentMode { //gd:LightmapGI.get_environment_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.LightmapGIEnvironmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_environment_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnvironmentCustomSky(sky [1]gdclass.Sky) { //gd:LightmapGI.set_environment_custom_sky
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(sky[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_environment_custom_sky, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironmentCustomSky() [1]gdclass.Sky { //gd:LightmapGI.get_environment_custom_sky
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_environment_custom_sky, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Sky{gd.PointerWithOwnershipTransferredToGo[gdclass.Sky](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnvironmentCustomColor(color gd.Color) { //gd:LightmapGI.set_environment_custom_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_environment_custom_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironmentCustomColor() gd.Color { //gd:LightmapGI.get_environment_custom_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_environment_custom_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnvironmentCustomEnergy(energy gd.Float) { //gd:LightmapGI.set_environment_custom_energy
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_environment_custom_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironmentCustomEnergy() gd.Float { //gd:LightmapGI.get_environment_custom_energy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_environment_custom_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexelScale(texel_scale gd.Float) { //gd:LightmapGI.set_texel_scale
	var frame = callframe.New()
	callframe.Arg(frame, texel_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_texel_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexelScale() gd.Float { //gd:LightmapGI.get_texel_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_texel_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxTextureSize(max_texture_size gd.Int) { //gd:LightmapGI.set_max_texture_size
	var frame = callframe.New()
	callframe.Arg(frame, max_texture_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_max_texture_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxTextureSize() gd.Int { //gd:LightmapGI.get_max_texture_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_max_texture_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseDenoiser(use_denoiser bool) { //gd:LightmapGI.set_use_denoiser
	var frame = callframe.New()
	callframe.Arg(frame, use_denoiser)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_use_denoiser, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingDenoiser() bool { //gd:LightmapGI.is_using_denoiser
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_is_using_denoiser, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDenoiserStrength(denoiser_strength gd.Float) { //gd:LightmapGI.set_denoiser_strength
	var frame = callframe.New()
	callframe.Arg(frame, denoiser_strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_denoiser_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDenoiserStrength() gd.Float { //gd:LightmapGI.get_denoiser_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_denoiser_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDenoiserRange(denoiser_range gd.Int) { //gd:LightmapGI.set_denoiser_range
	var frame = callframe.New()
	callframe.Arg(frame, denoiser_range)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_denoiser_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDenoiserRange() gd.Int { //gd:LightmapGI.get_denoiser_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_denoiser_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInterior(enable bool) { //gd:LightmapGI.set_interior
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_interior, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsInterior() bool { //gd:LightmapGI.is_interior
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_is_interior, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDirectional(directional bool) { //gd:LightmapGI.set_directional
	var frame = callframe.New()
	callframe.Arg(frame, directional)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_directional, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDirectional() bool { //gd:LightmapGI.is_directional
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_is_directional, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseTextureForBounces(use_texture_for_bounces bool) { //gd:LightmapGI.set_use_texture_for_bounces
	var frame = callframe.New()
	callframe.Arg(frame, use_texture_for_bounces)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_use_texture_for_bounces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingTextureForBounces() bool { //gd:LightmapGI.is_using_texture_for_bounces
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_is_using_texture_for_bounces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCameraAttributes(camera_attributes [1]gdclass.CameraAttributes) { //gd:LightmapGI.set_camera_attributes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(camera_attributes[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCameraAttributes() [1]gdclass.CameraAttributes { //gd:LightmapGI.get_camera_attributes
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGI.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CameraAttributes{gd.PointerWithOwnershipTransferredToGo[gdclass.CameraAttributes](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsLightmapGI() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLightmapGI() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Advanced(self.AsVisualInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Instance(self.AsVisualInstance3D()), name)
	}
}
func init() {
	gdclass.Register("LightmapGI", func(ptr gd.Object) any { return [1]gdclass.LightmapGI{*(*gdclass.LightmapGI)(unsafe.Pointer(&ptr))} })
}

type BakeQuality = gdclass.LightmapGIBakeQuality //gd:LightmapGI.BakeQuality

const (
	/*Low bake quality (fastest bake times). The quality of this preset can be adjusted by changing [member ProjectSettings.rendering/lightmapping/bake_quality/low_quality_ray_count] and [member ProjectSettings.rendering/lightmapping/bake_quality/low_quality_probe_ray_count].*/
	BakeQualityLow BakeQuality = 0
	/*Medium bake quality (fast bake times). The quality of this preset can be adjusted by changing [member ProjectSettings.rendering/lightmapping/bake_quality/medium_quality_ray_count] and [member ProjectSettings.rendering/lightmapping/bake_quality/medium_quality_probe_ray_count].*/
	BakeQualityMedium BakeQuality = 1
	/*High bake quality (slow bake times). The quality of this preset can be adjusted by changing [member ProjectSettings.rendering/lightmapping/bake_quality/high_quality_ray_count] and [member ProjectSettings.rendering/lightmapping/bake_quality/high_quality_probe_ray_count].*/
	BakeQualityHigh BakeQuality = 2
	/*Highest bake quality (slowest bake times). The quality of this preset can be adjusted by changing [member ProjectSettings.rendering/lightmapping/bake_quality/ultra_quality_ray_count] and [member ProjectSettings.rendering/lightmapping/bake_quality/ultra_quality_probe_ray_count].*/
	BakeQualityUltra BakeQuality = 3
)

type GenerateProbes = gdclass.LightmapGIGenerateProbes //gd:LightmapGI.GenerateProbes

const (
	/*Don't generate lightmap probes for lighting dynamic objects.*/
	GenerateProbesDisabled GenerateProbes = 0
	/*Lowest level of subdivision (fastest bake times, smallest file sizes).*/
	GenerateProbesSubdiv4 GenerateProbes = 1
	/*Low level of subdivision (fast bake times, small file sizes).*/
	GenerateProbesSubdiv8 GenerateProbes = 2
	/*High level of subdivision (slow bake times, large file sizes).*/
	GenerateProbesSubdiv16 GenerateProbes = 3
	/*Highest level of subdivision (slowest bake times, largest file sizes).*/
	GenerateProbesSubdiv32 GenerateProbes = 4
)

type BakeError = gdclass.LightmapGIBakeError //gd:LightmapGI.BakeError

const (
	/*Lightmap baking was successful.*/
	BakeErrorOk BakeError = 0
	/*Lightmap baking failed because the root node for the edited scene could not be accessed.*/
	BakeErrorNoSceneRoot BakeError = 1
	/*Lightmap baking failed as the lightmap data resource is embedded in a foreign resource.*/
	BakeErrorForeignData BakeError = 2
	/*Lightmap baking failed as there is no lightmapper available in this Godot build.*/
	BakeErrorNoLightmapper BakeError = 3
	/*Lightmap baking failed as the [LightmapGIData] save path isn't configured in the resource.*/
	BakeErrorNoSavePath BakeError = 4
	/*Lightmap baking failed as there are no meshes whose [member GeometryInstance3D.gi_mode] is [constant GeometryInstance3D.GI_MODE_STATIC] and with valid UV2 mapping in the current scene. You may need to select 3D scenes in the Import dock and change their global illumination mode accordingly.*/
	BakeErrorNoMeshes BakeError = 5
	/*Lightmap baking failed as the lightmapper failed to analyze some of the meshes marked as static for baking.*/
	BakeErrorMeshesInvalid BakeError = 6
	/*Lightmap baking failed as the resulting image couldn't be saved or imported by Godot after it was saved.*/
	BakeErrorCantCreateImage BakeError = 7
	/*The user aborted the lightmap baking operation (typically by clicking the [b]Cancel[/b] button in the progress dialog).*/
	BakeErrorUserAborted BakeError = 8
	/*Lightmap baking failed as the maximum texture size is too small to fit some of the meshes marked for baking.*/
	BakeErrorTextureSizeTooSmall BakeError = 9
	/*Lightmap baking failed as the lightmap is too small.*/
	BakeErrorLightmapTooSmall BakeError = 10
	/*Lightmap baking failed as the lightmap was unable to fit into an atlas.*/
	BakeErrorAtlasTooSmall BakeError = 11
)

type EnvironmentMode = gdclass.LightmapGIEnvironmentMode //gd:LightmapGI.EnvironmentMode

const (
	/*Ignore environment lighting when baking lightmaps.*/
	EnvironmentModeDisabled EnvironmentMode = 0
	/*Use the scene's environment lighting when baking lightmaps.
	  [b]Note:[/b] If baking lightmaps in a scene with no [WorldEnvironment] node, this will act like [constant ENVIRONMENT_MODE_DISABLED]. The editor's preview sky and sun is [i]not[/i] taken into account by [LightmapGI] when baking lightmaps.*/
	EnvironmentModeScene EnvironmentMode = 1
	/*Use [member environment_custom_sky] as a source of environment lighting when baking lightmaps.*/
	EnvironmentModeCustomSky EnvironmentMode = 2
	/*Use [member environment_custom_color] multiplied by [member environment_custom_energy] as a constant source of environment lighting when baking lightmaps.*/
	EnvironmentModeCustomColor EnvironmentMode = 3
)
