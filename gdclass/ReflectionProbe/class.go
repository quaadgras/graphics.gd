package ReflectionProbe

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Captures its surroundings as a cubemap, and stores versions of it with increasing levels of blur to simulate different material roughnesses.
The [ReflectionProbe] is used to create high-quality reflections at a low performance cost (when [member update_mode] is [constant UPDATE_ONCE]). [ReflectionProbe]s can be blended together and with the rest of the scene smoothly. [ReflectionProbe]s can also be combined with [VoxelGI], SDFGI ([member Environment.sdfgi_enabled]) and screen-space reflections ([member Environment.ssr_enabled]) to get more accurate reflections in specific areas. [ReflectionProbe]s render all objects within their [member cull_mask], so updating them can be quite expensive. It is best to update them once with the important static objects and then leave them as-is.
[b]Note:[/b] Unlike [VoxelGI] and SDFGI, [ReflectionProbe]s only source their environment from a [WorldEnvironment] node. If you specify an [Environment] resource within a [Camera3D] node, it will be ignored by the [ReflectionProbe]. This can lead to incorrect lighting within the [ReflectionProbe].
[b]Note:[/b] Reflection probes are only supported in the Forward+ and Mobile rendering methods, not Compatibility. When using the Mobile rendering method, only 8 reflection probes can be displayed on each mesh resource. Attempting to display more than 8 reflection probes on a single mesh resource will result in reflection probes flickering in and out as the camera moves.
[b]Note:[/b] When using the Mobile rendering method, reflection probes will only correctly affect meshes whose visibility AABB intersects with the reflection probe's AABB. If using a shader to deform the mesh in a way that makes it go outside its AABB, [member GeometryInstance3D.extra_cull_margin] must be increased on the mesh. Otherwise, the reflection probe may not be visible on the mesh.

*/
type Go [1]classdb.ReflectionProbe
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ReflectionProbe
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ReflectionProbe"))
	return Go{classdb.ReflectionProbe(object)}
}

func (self Go) UpdateMode() classdb.ReflectionProbeUpdateMode {
		return classdb.ReflectionProbeUpdateMode(class(self).GetUpdateMode())
}

func (self Go) SetUpdateMode(value classdb.ReflectionProbeUpdateMode) {
	class(self).SetUpdateMode(value)
}

func (self Go) Intensity() float64 {
		return float64(float64(class(self).GetIntensity()))
}

func (self Go) SetIntensity(value float64) {
	class(self).SetIntensity(gd.Float(value))
}

func (self Go) MaxDistance() float64 {
		return float64(float64(class(self).GetMaxDistance()))
}

func (self Go) SetMaxDistance(value float64) {
	class(self).SetMaxDistance(gd.Float(value))
}

func (self Go) Size() gd.Vector3 {
		return gd.Vector3(class(self).GetSize())
}

func (self Go) SetSize(value gd.Vector3) {
	class(self).SetSize(value)
}

func (self Go) OriginOffset() gd.Vector3 {
		return gd.Vector3(class(self).GetOriginOffset())
}

func (self Go) SetOriginOffset(value gd.Vector3) {
	class(self).SetOriginOffset(value)
}

func (self Go) BoxProjection() bool {
		return bool(class(self).IsBoxProjectionEnabled())
}

func (self Go) SetBoxProjection(value bool) {
	class(self).SetEnableBoxProjection(value)
}

func (self Go) Interior() bool {
		return bool(class(self).IsSetAsInterior())
}

func (self Go) SetInterior(value bool) {
	class(self).SetAsInterior(value)
}

func (self Go) EnableShadows() bool {
		return bool(class(self).AreShadowsEnabled())
}

func (self Go) SetEnableShadows(value bool) {
	class(self).SetEnableShadows(value)
}

func (self Go) CullMask() int {
		return int(int(class(self).GetCullMask()))
}

func (self Go) SetCullMask(value int) {
	class(self).SetCullMask(gd.Int(value))
}

func (self Go) ReflectionMask() int {
		return int(int(class(self).GetReflectionMask()))
}

func (self Go) SetReflectionMask(value int) {
	class(self).SetReflectionMask(gd.Int(value))
}

func (self Go) MeshLodThreshold() float64 {
		return float64(float64(class(self).GetMeshLodThreshold()))
}

func (self Go) SetMeshLodThreshold(value float64) {
	class(self).SetMeshLodThreshold(gd.Float(value))
}

func (self Go) AmbientMode() classdb.ReflectionProbeAmbientMode {
		return classdb.ReflectionProbeAmbientMode(class(self).GetAmbientMode())
}

func (self Go) SetAmbientMode(value classdb.ReflectionProbeAmbientMode) {
	class(self).SetAmbientMode(value)
}

func (self Go) AmbientColor() gd.Color {
		return gd.Color(class(self).GetAmbientColor())
}

func (self Go) SetAmbientColor(value gd.Color) {
	class(self).SetAmbientColor(value)
}

func (self Go) AmbientColorEnergy() float64 {
		return float64(float64(class(self).GetAmbientColorEnergy()))
}

func (self Go) SetAmbientColorEnergy(value float64) {
	class(self).SetAmbientColorEnergy(gd.Float(value))
}

//go:nosplit
func (self class) SetIntensity(intensity gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIntensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientMode(ambient classdb.ReflectionProbeAmbientMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, ambient)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_ambient_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientMode() classdb.ReflectionProbeAmbientMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ReflectionProbeAmbientMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_ambient_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientColor(ambient gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, ambient)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_ambient_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_ambient_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAmbientColorEnergy(ambient_energy gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, ambient_energy)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_ambient_color_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAmbientColorEnergy() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_ambient_color_energy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxDistance(max_distance gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, max_distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMeshLodThreshold(ratio gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_mesh_lod_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMeshLodThreshold() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_mesh_lod_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSize(size gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOriginOffset(origin_offset gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, origin_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_origin_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOriginOffset() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_origin_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAsInterior(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_as_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSetAsInterior() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_is_set_as_interior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableBoxProjection(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_enable_box_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsBoxProjectionEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_is_box_projection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableShadows(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_enable_shadows, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreShadowsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_are_shadows_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCullMask(layers gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCullMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReflectionMask(layers gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_reflection_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReflectionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_reflection_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateMode(mode classdb.ReflectionProbeUpdateMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_set_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateMode() classdb.ReflectionProbeUpdateMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ReflectionProbeUpdateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ReflectionProbe.Bind_get_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsReflectionProbe() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsReflectionProbe() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}
func init() {classdb.Register("ReflectionProbe", func(ptr gd.Object) any { return classdb.ReflectionProbe(ptr) })}
type UpdateMode = classdb.ReflectionProbeUpdateMode

const (
/*Update the probe once on the next frame (recommended for most objects). The corresponding radiance map will be generated over the following six frames. This takes more time to update than [constant UPDATE_ALWAYS], but it has a lower performance cost and can result in higher-quality reflections. The ReflectionProbe is updated when its transform changes, but not when nearby geometry changes. You can force a [ReflectionProbe] update by moving the [ReflectionProbe] slightly in any direction.*/
	UpdateOnce UpdateMode = 0
/*Update the probe every frame. This provides better results for fast-moving dynamic objects (such as cars). However, it has a significant performance cost. Due to the cost, it's recommended to only use one ReflectionProbe with [constant UPDATE_ALWAYS] at most per scene. For all other use cases, use [constant UPDATE_ONCE].*/
	UpdateAlways UpdateMode = 1
)
type AmbientMode = classdb.ReflectionProbeAmbientMode

const (
/*Do not apply any ambient lighting inside the [ReflectionProbe]'s box defined by its [member size].*/
	AmbientDisabled AmbientMode = 0
/*Apply automatically-sourced environment lighting inside the [ReflectionProbe]'s box defined by its [member size].*/
	AmbientEnvironment AmbientMode = 1
/*Apply custom ambient lighting inside the [ReflectionProbe]'s box defined by its [member size]. See [member ambient_color] and [member ambient_color_energy].*/
	AmbientColor AmbientMode = 2
)
