// Package ImporterMeshInstance3D provides methods for working with ImporterMeshInstance3D object instances.
package ImporterMeshInstance3D

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
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Float"

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
var _ String.Readable

type Instance [1]gdclass.ImporterMeshInstance3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsImporterMeshInstance3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ImporterMeshInstance3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImporterMeshInstance3D"))
	casted := Instance{*(*gdclass.ImporterMeshInstance3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Mesh() [1]gdclass.ImporterMesh {
	return [1]gdclass.ImporterMesh(class(self).GetMesh())
}

func (self Instance) SetMesh(value [1]gdclass.ImporterMesh) {
	class(self).SetMesh(value)
}

func (self Instance) Skin() [1]gdclass.Skin {
	return [1]gdclass.Skin(class(self).GetSkin())
}

func (self Instance) SetSkin(value [1]gdclass.Skin) {
	class(self).SetSkin(value)
}

func (self Instance) SkeletonPath() NodePath.String {
	return NodePath.String(class(self).GetSkeletonPath().String())
}

func (self Instance) SetSkeletonPath(value NodePath.String) {
	class(self).SetSkeletonPath(gd.NewString(string(value)).NodePath())
}

func (self Instance) LayerMask() int {
	return int(int(class(self).GetLayerMask()))
}

func (self Instance) SetLayerMask(value int) {
	class(self).SetLayerMask(gd.Int(value))
}

func (self Instance) CastShadow() gdclass.GeometryInstance3DShadowCastingSetting {
	return gdclass.GeometryInstance3DShadowCastingSetting(class(self).GetCastShadowsSetting())
}

func (self Instance) SetCastShadow(value gdclass.GeometryInstance3DShadowCastingSetting) {
	class(self).SetCastShadowsSetting(value)
}

func (self Instance) VisibilityRangeBegin() Float.X {
	return Float.X(Float.X(class(self).GetVisibilityRangeBegin()))
}

func (self Instance) SetVisibilityRangeBegin(value Float.X) {
	class(self).SetVisibilityRangeBegin(gd.Float(value))
}

func (self Instance) VisibilityRangeBeginMargin() Float.X {
	return Float.X(Float.X(class(self).GetVisibilityRangeBeginMargin()))
}

func (self Instance) SetVisibilityRangeBeginMargin(value Float.X) {
	class(self).SetVisibilityRangeBeginMargin(gd.Float(value))
}

func (self Instance) VisibilityRangeEnd() Float.X {
	return Float.X(Float.X(class(self).GetVisibilityRangeEnd()))
}

func (self Instance) SetVisibilityRangeEnd(value Float.X) {
	class(self).SetVisibilityRangeEnd(gd.Float(value))
}

func (self Instance) VisibilityRangeEndMargin() Float.X {
	return Float.X(Float.X(class(self).GetVisibilityRangeEndMargin()))
}

func (self Instance) SetVisibilityRangeEndMargin(value Float.X) {
	class(self).SetVisibilityRangeEndMargin(gd.Float(value))
}

func (self Instance) VisibilityRangeFadeMode() gdclass.GeometryInstance3DVisibilityRangeFadeMode {
	return gdclass.GeometryInstance3DVisibilityRangeFadeMode(class(self).GetVisibilityRangeFadeMode())
}

func (self Instance) SetVisibilityRangeFadeMode(value gdclass.GeometryInstance3DVisibilityRangeFadeMode) {
	class(self).SetVisibilityRangeFadeMode(value)
}

//go:nosplit
func (self class) SetMesh(mesh [1]gdclass.ImporterMesh) { //gd:ImporterMeshInstance3D.set_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() [1]gdclass.ImporterMesh { //gd:ImporterMeshInstance3D.get_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ImporterMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.ImporterMesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkin(skin [1]gdclass.Skin) { //gd:ImporterMeshInstance3D.set_skin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_skin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkin() [1]gdclass.Skin { //gd:ImporterMeshInstance3D.get_skin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skin{gd.PointerWithOwnershipTransferredToGo[gdclass.Skin](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkeletonPath(skeleton_path gd.NodePath) { //gd:ImporterMeshInstance3D.set_skeleton_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skeleton_path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkeletonPath() gd.NodePath { //gd:ImporterMeshInstance3D.get_skeleton_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLayerMask(layer_mask gd.Int) { //gd:ImporterMeshInstance3D.set_layer_mask
	var frame = callframe.New()
	callframe.Arg(frame, layer_mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_layer_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLayerMask() gd.Int { //gd:ImporterMeshInstance3D.get_layer_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_layer_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCastShadowsSetting(shadow_casting_setting gdclass.GeometryInstance3DShadowCastingSetting) { //gd:ImporterMeshInstance3D.set_cast_shadows_setting
	var frame = callframe.New()
	callframe.Arg(frame, shadow_casting_setting)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCastShadowsSetting() gdclass.GeometryInstance3DShadowCastingSetting { //gd:ImporterMeshInstance3D.get_cast_shadows_setting
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GeometryInstance3DShadowCastingSetting](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeEndMargin(distance gd.Float) { //gd:ImporterMeshInstance3D.set_visibility_range_end_margin
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeEndMargin() gd.Float { //gd:ImporterMeshInstance3D.get_visibility_range_end_margin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeEnd(distance gd.Float) { //gd:ImporterMeshInstance3D.set_visibility_range_end
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeEnd() gd.Float { //gd:ImporterMeshInstance3D.get_visibility_range_end
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeBeginMargin(distance gd.Float) { //gd:ImporterMeshInstance3D.set_visibility_range_begin_margin
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeBeginMargin() gd.Float { //gd:ImporterMeshInstance3D.get_visibility_range_begin_margin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeBegin(distance gd.Float) { //gd:ImporterMeshInstance3D.set_visibility_range_begin
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeBegin() gd.Float { //gd:ImporterMeshInstance3D.get_visibility_range_begin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeFadeMode(mode gdclass.GeometryInstance3DVisibilityRangeFadeMode) { //gd:ImporterMeshInstance3D.set_visibility_range_fade_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeFadeMode() gdclass.GeometryInstance3DVisibilityRangeFadeMode { //gd:ImporterMeshInstance3D.get_visibility_range_fade_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GeometryInstance3DVisibilityRangeFadeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsImporterMeshInstance3D() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImporterMeshInstance3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("ImporterMeshInstance3D", func(ptr gd.Object) any {
		return [1]gdclass.ImporterMeshInstance3D{*(*gdclass.ImporterMeshInstance3D)(unsafe.Pointer(&ptr))}
	})
}
