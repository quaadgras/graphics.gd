package OccluderInstance3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

/*
Occlusion culling can improve rendering performance in closed/semi-open areas by hiding geometry that is occluded by other objects.
The occlusion culling system is mostly static. [OccluderInstance3D]s can be moved or hidden at run-time, but doing so will trigger a background recomputation that can take several frames. It is recommended to only move [OccluderInstance3D]s sporadically (e.g. for procedural generation purposes), rather than doing so every frame.
The occlusion culling system works by rendering the occluders on the CPU in parallel using [url=https://www.embree.org/]Embree[/url], drawing the result to a low-resolution buffer then using this to cull 3D nodes individually. In the 3D editor, you can preview the occlusion culling buffer by choosing [b]Perspective > Debug Advanced... > Occlusion Culling Buffer[/b] in the top-left corner of the 3D viewport. The occlusion culling buffer quality can be adjusted in the Project Settings.
[b]Baking:[/b] Select an [OccluderInstance3D] node, then use the [b]Bake Occluders[/b] button at the top of the 3D editor. Only opaque materials will be taken into account; transparent materials (alpha-blended or alpha-tested) will be ignored by the occluder generation.
[b]Note:[/b] Occlusion culling is only effective if [member ProjectSettings.rendering/occlusion_culling/use_occlusion_culling] is [code]true[/code]. Enabling occlusion culling has a cost on the CPU. Only enable occlusion culling if you actually plan to use it. Large open scenes with few or no objects blocking the view will generally not benefit much from occlusion culling. Large open scenes generally benefit more from mesh LOD and visibility ranges ([member GeometryInstance3D.visibility_range_begin] and [member GeometryInstance3D.visibility_range_end]) compared to occlusion culling.
[b]Note:[/b] Due to memory constraints, occlusion culling is not supported by default in Web export templates. It can be enabled by compiling custom Web export templates with [code]module_raycast_enabled=yes[/code].

*/
type Go [1]classdb.OccluderInstance3D

/*
Based on [param value], enables or disables the specified layer in the [member bake_mask], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetBakeMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBakeMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member bake_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetBakeMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetBakeMaskValue(gd.Int(layer_number)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OccluderInstance3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OccluderInstance3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Occluder() gdclass.Occluder3D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Occluder3D(class(self).GetOccluder(gc))
}

func (self Go) SetOccluder(value gdclass.Occluder3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOccluder(value)
}

func (self Go) BakeMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBakeMask()))
}

func (self Go) SetBakeMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBakeMask(gd.Int(value))
}

func (self Go) BakeSimplificationDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBakeSimplificationDistance()))
}

func (self Go) SetBakeSimplificationDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBakeSimplificationDistance(gd.Float(value))
}

//go:nosplit
func (self class) SetBakeMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderInstance3D.Bind_set_bake_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakeMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderInstance3D.Bind_get_bake_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member bake_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetBakeMaskValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderInstance3D.Bind_set_bake_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member bake_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetBakeMaskValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderInstance3D.Bind_get_bake_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBakeSimplificationDistance(simplification_distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, simplification_distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderInstance3D.Bind_set_bake_simplification_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakeSimplificationDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderInstance3D.Bind_get_bake_simplification_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOccluder(occluder gdclass.Occluder3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(occluder[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderInstance3D.Bind_set_occluder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOccluder(ctx gd.Lifetime) gdclass.Occluder3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OccluderInstance3D.Bind_get_occluder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Occluder3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsOccluderInstance3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOccluderInstance3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("OccluderInstance3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
