package LightmapGI

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [LightmapGI] node is used to compute and store baked lightmaps. Lightmaps are used to provide high-quality indirect lighting with very little light leaking. [LightmapGI] can also provide rough reflections using spherical harmonics if [member directional] is enabled. Dynamic objects can receive indirect lighting thanks to [i]light probes[/i], which can be automatically placed by setting [member generate_probes_subdiv] to a value other than [constant GENERATE_PROBES_DISABLED]. Additional lightmap probes can also be added by creating [LightmapProbe] nodes. The downside is that lightmaps are fully static and cannot be baked in an exported project. Baking a [LightmapGI] node is also slower compared to [VoxelGI].
[b]Procedural generation:[/b] Lightmap baking functionality is only available in the editor. This means [LightmapGI] is not suited to procedurally generated or user-built levels. For procedurally generated or user-built levels, use [VoxelGI] or SDFGI instead (see [member Environment.sdfgi_enabled]).
[b]Performance:[/b] [LightmapGI] provides the best possible run-time performance for global illumination. It is suitable for low-end hardware including integrated graphics and mobile devices.
[b]Note:[/b] Due to how lightmaps work, most properties only have a visible effect once lightmaps are baked again.
[b]Note:[/b] Lightmap baking on [CSGShape3D]s and [PrimitiveMesh]es is not supported, as these cannot store UV2 data required for baking.
[b]Note:[/b] If no custom lightmappers are installed, [LightmapGI] can only be baked from devices that support the Forward+ or Mobile rendering backends.

*/
type Simple [1]classdb.LightmapGI
func (self Simple) SetLightData(data [1]classdb.LightmapGIData) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLightData(data)
}
func (self Simple) GetLightData() [1]classdb.LightmapGIData {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.LightmapGIData(Expert(self).GetLightData(gc))
}
func (self Simple) SetBakeQuality(bake_quality classdb.LightmapGIBakeQuality) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBakeQuality(bake_quality)
}
func (self Simple) GetBakeQuality() classdb.LightmapGIBakeQuality {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.LightmapGIBakeQuality(Expert(self).GetBakeQuality())
}
func (self Simple) SetBounces(bounces int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBounces(gd.Int(bounces))
}
func (self Simple) GetBounces() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBounces()))
}
func (self Simple) SetBounceIndirectEnergy(bounce_indirect_energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBounceIndirectEnergy(gd.Float(bounce_indirect_energy))
}
func (self Simple) GetBounceIndirectEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBounceIndirectEnergy()))
}
func (self Simple) SetGenerateProbes(subdivision classdb.LightmapGIGenerateProbes) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGenerateProbes(subdivision)
}
func (self Simple) GetGenerateProbes() classdb.LightmapGIGenerateProbes {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.LightmapGIGenerateProbes(Expert(self).GetGenerateProbes())
}
func (self Simple) SetBias(bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBias(gd.Float(bias))
}
func (self Simple) GetBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBias()))
}
func (self Simple) SetEnvironmentMode(mode classdb.LightmapGIEnvironmentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnvironmentMode(mode)
}
func (self Simple) GetEnvironmentMode() classdb.LightmapGIEnvironmentMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.LightmapGIEnvironmentMode(Expert(self).GetEnvironmentMode())
}
func (self Simple) SetEnvironmentCustomSky(sky [1]classdb.Sky) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnvironmentCustomSky(sky)
}
func (self Simple) GetEnvironmentCustomSky() [1]classdb.Sky {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Sky(Expert(self).GetEnvironmentCustomSky(gc))
}
func (self Simple) SetEnvironmentCustomColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnvironmentCustomColor(color)
}
func (self Simple) GetEnvironmentCustomColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetEnvironmentCustomColor())
}
func (self Simple) SetEnvironmentCustomEnergy(energy float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnvironmentCustomEnergy(gd.Float(energy))
}
func (self Simple) GetEnvironmentCustomEnergy() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEnvironmentCustomEnergy()))
}
func (self Simple) SetTexelScale(texel_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTexelScale(gd.Float(texel_scale))
}
func (self Simple) GetTexelScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTexelScale()))
}
func (self Simple) SetMaxTextureSize(max_texture_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxTextureSize(gd.Int(max_texture_size))
}
func (self Simple) GetMaxTextureSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxTextureSize()))
}
func (self Simple) SetUseDenoiser(use_denoiser bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseDenoiser(use_denoiser)
}
func (self Simple) IsUsingDenoiser() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingDenoiser())
}
func (self Simple) SetDenoiserStrength(denoiser_strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDenoiserStrength(gd.Float(denoiser_strength))
}
func (self Simple) GetDenoiserStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDenoiserStrength()))
}
func (self Simple) SetDenoiserRange(denoiser_range int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDenoiserRange(gd.Int(denoiser_range))
}
func (self Simple) GetDenoiserRange() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDenoiserRange()))
}
func (self Simple) SetInterior(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInterior(enable)
}
func (self Simple) IsInterior() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInterior())
}
func (self Simple) SetDirectional(directional bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDirectional(directional)
}
func (self Simple) IsDirectional() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDirectional())
}
func (self Simple) SetUseTextureForBounces(use_texture_for_bounces bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseTextureForBounces(use_texture_for_bounces)
}
func (self Simple) IsUsingTextureForBounces() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingTextureForBounces())
}
func (self Simple) SetCameraAttributes(camera_attributes [1]classdb.CameraAttributes) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCameraAttributes(camera_attributes)
}
func (self Simple) GetCameraAttributes() [1]classdb.CameraAttributes {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.CameraAttributes(Expert(self).GetCameraAttributes(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.LightmapGI
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetLightData(data object.LightmapGIData)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_light_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightData(ctx gd.Lifetime) object.LightmapGIData {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_light_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.LightmapGIData
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBakeQuality(bake_quality classdb.LightmapGIBakeQuality)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bake_quality)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_bake_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakeQuality() classdb.LightmapGIBakeQuality {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.LightmapGIBakeQuality](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_bake_quality, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBounces(bounces gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bounces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBounces() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBounceIndirectEnergy(bounce_indirect_energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bounce_indirect_energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_bounce_indirect_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBounceIndirectEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_bounce_indirect_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGenerateProbes(subdivision classdb.LightmapGIGenerateProbes)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, subdivision)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_generate_probes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGenerateProbes() classdb.LightmapGIGenerateProbes {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.LightmapGIGenerateProbes](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_generate_probes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBias(bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironmentMode(mode classdb.LightmapGIEnvironmentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_environment_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironmentMode() classdb.LightmapGIEnvironmentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.LightmapGIEnvironmentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_environment_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironmentCustomSky(sky object.Sky)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(sky[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_environment_custom_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironmentCustomSky(ctx gd.Lifetime) object.Sky {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_environment_custom_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Sky
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironmentCustomColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_environment_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironmentCustomColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_environment_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironmentCustomEnergy(energy gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_environment_custom_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironmentCustomEnergy() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_environment_custom_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTexelScale(texel_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texel_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_texel_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexelScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_texel_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxTextureSize(max_texture_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_texture_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_max_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxTextureSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_max_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseDenoiser(use_denoiser bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_denoiser)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_use_denoiser, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingDenoiser() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_is_using_denoiser, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDenoiserStrength(denoiser_strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, denoiser_strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_denoiser_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDenoiserStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_denoiser_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDenoiserRange(denoiser_range gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, denoiser_range)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_denoiser_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDenoiserRange() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_denoiser_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInterior(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsInterior() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_is_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDirectional(directional bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, directional)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_directional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDirectional() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_is_directional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseTextureForBounces(use_texture_for_bounces bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_texture_for_bounces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_use_texture_for_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingTextureForBounces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_is_using_texture_for_bounces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCameraAttributes(camera_attributes object.CameraAttributes)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(camera_attributes[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_set_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCameraAttributes(ctx gd.Lifetime) object.CameraAttributes {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LightmapGI.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.CameraAttributes
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsLightmapGI() Expert { return self[0].AsLightmapGI() }


//go:nosplit
func (self Simple) AsLightmapGI() Simple { return self[0].AsLightmapGI() }


//go:nosplit
func (self class) AsVisualInstance3D() VisualInstance3D.Expert { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self Simple) AsVisualInstance3D() VisualInstance3D.Simple { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("LightmapGI", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type BakeQuality = classdb.LightmapGIBakeQuality

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
type GenerateProbes = classdb.LightmapGIGenerateProbes

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
type BakeError = classdb.LightmapGIBakeError

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
type EnvironmentMode = classdb.LightmapGIEnvironmentMode

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
