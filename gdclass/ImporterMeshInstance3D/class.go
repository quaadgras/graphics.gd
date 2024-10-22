package ImporterMeshInstance3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Go [1]classdb.ImporterMeshInstance3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ImporterMeshInstance3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ImporterMeshInstance3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Mesh() gdclass.ImporterMesh {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.ImporterMesh(class(self).GetMesh(gc))
}

func (self Go) SetMesh(value gdclass.ImporterMesh) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMesh(value)
}

func (self Go) Skin() gdclass.Skin {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Skin(class(self).GetSkin(gc))
}

func (self Go) SetSkin(value gdclass.Skin) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkin(value)
}

func (self Go) SkeletonPath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetSkeletonPath(gc).String())
}

func (self Go) SetSkeletonPath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkeletonPath(gc.String(value).NodePath(gc))
}

func (self Go) LayerMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLayerMask()))
}

func (self Go) SetLayerMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLayerMask(gd.Int(value))
}

func (self Go) CastShadow() classdb.GeometryInstance3DShadowCastingSetting {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GeometryInstance3DShadowCastingSetting(class(self).GetCastShadowsSetting())
}

func (self Go) SetCastShadow(value classdb.GeometryInstance3DShadowCastingSetting) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCastShadowsSetting(value)
}

func (self Go) VisibilityRangeBegin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVisibilityRangeBegin()))
}

func (self Go) SetVisibilityRangeBegin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeBegin(gd.Float(value))
}

func (self Go) VisibilityRangeBeginMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVisibilityRangeBeginMargin()))
}

func (self Go) SetVisibilityRangeBeginMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeBeginMargin(gd.Float(value))
}

func (self Go) VisibilityRangeEnd() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVisibilityRangeEnd()))
}

func (self Go) SetVisibilityRangeEnd(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeEnd(gd.Float(value))
}

func (self Go) VisibilityRangeEndMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVisibilityRangeEndMargin()))
}

func (self Go) SetVisibilityRangeEndMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeEndMargin(gd.Float(value))
}

func (self Go) VisibilityRangeFadeMode() classdb.GeometryInstance3DVisibilityRangeFadeMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GeometryInstance3DVisibilityRangeFadeMode(class(self).GetVisibilityRangeFadeMode())
}

func (self Go) SetVisibilityRangeFadeMode(value classdb.GeometryInstance3DVisibilityRangeFadeMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityRangeFadeMode(value)
}

//go:nosplit
func (self class) SetMesh(mesh gdclass.ImporterMesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMesh(ctx gd.Lifetime) gdclass.ImporterMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.ImporterMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkin(skin gdclass.Skin)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(skin[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkin(ctx gd.Lifetime) gdclass.Skin {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Skin
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkeletonPath(skeleton_path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(skeleton_path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkeletonPath(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLayerMask(layer_mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_layer_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLayerMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_layer_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCastShadowsSetting(shadow_casting_setting classdb.GeometryInstance3DShadowCastingSetting)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shadow_casting_setting)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCastShadowsSetting() classdb.GeometryInstance3DShadowCastingSetting {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GeometryInstance3DShadowCastingSetting](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeEndMargin(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeEndMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeEnd(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeEnd() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeBeginMargin(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeBeginMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeBegin(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeBegin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityRangeFadeMode(mode classdb.GeometryInstance3DVisibilityRangeFadeMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_set_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityRangeFadeMode() classdb.GeometryInstance3DVisibilityRangeFadeMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GeometryInstance3DVisibilityRangeFadeMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImporterMeshInstance3D.Bind_get_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsImporterMeshInstance3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsImporterMeshInstance3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("ImporterMeshInstance3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
