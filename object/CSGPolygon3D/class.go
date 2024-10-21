package CSGPolygon3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/CSGPrimitive3D"
import "grow.graphics/gd/object/CSGShape3D"
import "grow.graphics/gd/object/GeometryInstance3D"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An array of 2D points is extruded to quickly and easily create a variety of 3D meshes. See also [CSGMesh3D] for using 3D meshes as CSG nodes.
[b]Note:[/b] CSG nodes are intended to be used for level prototyping. Creating CSG nodes has a significant CPU cost compared to creating a [MeshInstance3D] with a [PrimitiveMesh]. Moving a CSG node within another CSG node also has a significant CPU cost, so it should be avoided during gameplay.

*/
type Simple [1]classdb.CSGPolygon3D
func (self Simple) SetPolygon(polygon gd.PackedVector2Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPolygon(polygon)
}
func (self Simple) GetPolygon() gd.PackedVector2Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector2Array(Expert(self).GetPolygon(gc))
}
func (self Simple) SetMode(mode classdb.CSGPolygon3DMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMode(mode)
}
func (self Simple) GetMode() classdb.CSGPolygon3DMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CSGPolygon3DMode(Expert(self).GetMode())
}
func (self Simple) SetDepth(depth float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepth(gd.Float(depth))
}
func (self Simple) GetDepth() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDepth()))
}
func (self Simple) SetSpinDegrees(degrees float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSpinDegrees(gd.Float(degrees))
}
func (self Simple) GetSpinDegrees() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSpinDegrees()))
}
func (self Simple) SetSpinSides(spin_sides int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSpinSides(gd.Int(spin_sides))
}
func (self Simple) GetSpinSides() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSpinSides()))
}
func (self Simple) SetPathNode(path gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathNode(path)
}
func (self Simple) GetPathNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetPathNode(gc))
}
func (self Simple) SetPathIntervalType(interval_type classdb.CSGPolygon3DPathIntervalType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathIntervalType(interval_type)
}
func (self Simple) GetPathIntervalType() classdb.CSGPolygon3DPathIntervalType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CSGPolygon3DPathIntervalType(Expert(self).GetPathIntervalType())
}
func (self Simple) SetPathInterval(interval float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathInterval(gd.Float(interval))
}
func (self Simple) GetPathInterval() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPathInterval()))
}
func (self Simple) SetPathSimplifyAngle(degrees float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathSimplifyAngle(gd.Float(degrees))
}
func (self Simple) GetPathSimplifyAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPathSimplifyAngle()))
}
func (self Simple) SetPathRotation(path_rotation classdb.CSGPolygon3DPathRotation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathRotation(path_rotation)
}
func (self Simple) GetPathRotation() classdb.CSGPolygon3DPathRotation {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CSGPolygon3DPathRotation(Expert(self).GetPathRotation())
}
func (self Simple) SetPathLocal(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathLocal(enable)
}
func (self Simple) IsPathLocal() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPathLocal())
}
func (self Simple) SetPathContinuousU(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathContinuousU(enable)
}
func (self Simple) IsPathContinuousU() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPathContinuousU())
}
func (self Simple) SetPathUDistance(distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathUDistance(gd.Float(distance))
}
func (self Simple) GetPathUDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPathUDistance()))
}
func (self Simple) SetPathJoined(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathJoined(enable)
}
func (self Simple) IsPathJoined() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPathJoined())
}
func (self Simple) SetMaterial(material [1]classdb.Material) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaterial(material)
}
func (self Simple) GetMaterial() [1]classdb.Material {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Material(Expert(self).GetMaterial(gc))
}
func (self Simple) SetSmoothFaces(smooth_faces bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSmoothFaces(smooth_faces)
}
func (self Simple) GetSmoothFaces() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetSmoothFaces())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CSGPolygon3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(polygon))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPolygon(ctx gd.Lifetime) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMode(mode classdb.CSGPolygon3DMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMode() classdb.CSGPolygon3DMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CSGPolygon3DMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepth(depth gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepth() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpinDegrees(degrees gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_spin_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpinDegrees() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_spin_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpinSides(spin_sides gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, spin_sides)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_spin_sides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpinSides() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_spin_sides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathNode(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathIntervalType(interval_type classdb.CSGPolygon3DPathIntervalType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interval_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_interval_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathIntervalType() classdb.CSGPolygon3DPathIntervalType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CSGPolygon3DPathIntervalType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_interval_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathInterval(interval gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interval)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathInterval() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathSimplifyAngle(degrees gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_simplify_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathSimplifyAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_simplify_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathRotation(path_rotation classdb.CSGPolygon3DPathRotation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, path_rotation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathRotation() classdb.CSGPolygon3DPathRotation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CSGPolygon3DPathRotation](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathLocal(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPathLocal() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_is_path_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathContinuousU(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_continuous_u, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPathContinuousU() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_is_path_continuous_u, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathUDistance(distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_u_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathUDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_path_u_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathJoined(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_path_joined, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPathJoined() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_is_path_joined, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) object.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSmoothFaces(smooth_faces bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, smooth_faces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_set_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSmoothFaces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPolygon3D.Bind_get_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCSGPolygon3D() Expert { return self[0].AsCSGPolygon3D() }


//go:nosplit
func (self Simple) AsCSGPolygon3D() Simple { return self[0].AsCSGPolygon3D() }


//go:nosplit
func (self class) AsCSGPrimitive3D() CSGPrimitive3D.Expert { return self[0].AsCSGPrimitive3D() }


//go:nosplit
func (self Simple) AsCSGPrimitive3D() CSGPrimitive3D.Simple { return self[0].AsCSGPrimitive3D() }


//go:nosplit
func (self class) AsCSGShape3D() CSGShape3D.Expert { return self[0].AsCSGShape3D() }


//go:nosplit
func (self Simple) AsCSGShape3D() CSGShape3D.Simple { return self[0].AsCSGShape3D() }


//go:nosplit
func (self class) AsGeometryInstance3D() GeometryInstance3D.Expert { return self[0].AsGeometryInstance3D() }


//go:nosplit
func (self Simple) AsGeometryInstance3D() GeometryInstance3D.Simple { return self[0].AsGeometryInstance3D() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("CSGPolygon3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Mode = classdb.CSGPolygon3DMode

const (
/*The [member polygon] shape is extruded along the negative Z axis.*/
	ModeDepth Mode = 0
/*The [member polygon] shape is extruded by rotating it around the Y axis.*/
	ModeSpin Mode = 1
/*The [member polygon] shape is extruded along the [Path3D] specified in [member path_node].*/
	ModePath Mode = 2
)
type PathRotation = classdb.CSGPolygon3DPathRotation

const (
/*The [member polygon] shape is not rotated.
[b]Note:[/b] Requires the path Z coordinates to continually decrease to ensure viable shapes.*/
	PathRotationPolygon PathRotation = 0
/*The [member polygon] shape is rotated along the path, but it is not rotated around the path axis.
[b]Note:[/b] Requires the path Z coordinates to continually decrease to ensure viable shapes.*/
	PathRotationPath PathRotation = 1
/*The [member polygon] shape follows the path and its rotations around the path axis.*/
	PathRotationPathFollow PathRotation = 2
)
type PathIntervalType = classdb.CSGPolygon3DPathIntervalType

const (
/*When [member mode] is set to [constant MODE_PATH], [member path_interval] will determine the distance, in meters, each interval of the path will extrude.*/
	PathIntervalDistance PathIntervalType = 0
/*When [member mode] is set to [constant MODE_PATH], [member path_interval] will subdivide the polygons along the path.*/
	PathIntervalSubdivide PathIntervalType = 1
)
