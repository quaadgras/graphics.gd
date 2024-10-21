package MeshConvexDecompositionSettings

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Parameters to be used with a [Mesh] convex decomposition operation.

*/
type Simple [1]classdb.MeshConvexDecompositionSettings
func (self Simple) SetMaxConcavity(max_concavity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxConcavity(gd.Float(max_concavity))
}
func (self Simple) GetMaxConcavity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMaxConcavity()))
}
func (self Simple) SetSymmetryPlanesClippingBias(symmetry_planes_clipping_bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSymmetryPlanesClippingBias(gd.Float(symmetry_planes_clipping_bias))
}
func (self Simple) GetSymmetryPlanesClippingBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSymmetryPlanesClippingBias()))
}
func (self Simple) SetRevolutionAxesClippingBias(revolution_axes_clipping_bias float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRevolutionAxesClippingBias(gd.Float(revolution_axes_clipping_bias))
}
func (self Simple) GetRevolutionAxesClippingBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRevolutionAxesClippingBias()))
}
func (self Simple) SetMinVolumePerConvexHull(min_volume_per_convex_hull float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinVolumePerConvexHull(gd.Float(min_volume_per_convex_hull))
}
func (self Simple) GetMinVolumePerConvexHull() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMinVolumePerConvexHull()))
}
func (self Simple) SetResolution(min_volume_per_convex_hull int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetResolution(gd.Int(min_volume_per_convex_hull))
}
func (self Simple) GetResolution() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetResolution()))
}
func (self Simple) SetMaxNumVerticesPerConvexHull(max_num_vertices_per_convex_hull int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxNumVerticesPerConvexHull(gd.Int(max_num_vertices_per_convex_hull))
}
func (self Simple) GetMaxNumVerticesPerConvexHull() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxNumVerticesPerConvexHull()))
}
func (self Simple) SetPlaneDownsampling(plane_downsampling int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPlaneDownsampling(gd.Int(plane_downsampling))
}
func (self Simple) GetPlaneDownsampling() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPlaneDownsampling()))
}
func (self Simple) SetConvexHullDownsampling(convex_hull_downsampling int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConvexHullDownsampling(gd.Int(convex_hull_downsampling))
}
func (self Simple) GetConvexHullDownsampling() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetConvexHullDownsampling()))
}
func (self Simple) SetNormalizeMesh(normalize_mesh bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNormalizeMesh(normalize_mesh)
}
func (self Simple) GetNormalizeMesh() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetNormalizeMesh())
}
func (self Simple) SetMode(mode classdb.MeshConvexDecompositionSettingsMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMode(mode)
}
func (self Simple) GetMode() classdb.MeshConvexDecompositionSettingsMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.MeshConvexDecompositionSettingsMode(Expert(self).GetMode())
}
func (self Simple) SetConvexHullApproximation(convex_hull_approximation bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetConvexHullApproximation(convex_hull_approximation)
}
func (self Simple) GetConvexHullApproximation() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetConvexHullApproximation())
}
func (self Simple) SetMaxConvexHulls(max_convex_hulls int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxConvexHulls(gd.Int(max_convex_hulls))
}
func (self Simple) GetMaxConvexHulls() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxConvexHulls()))
}
func (self Simple) SetProjectHullVertices(project_hull_vertices bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProjectHullVertices(project_hull_vertices)
}
func (self Simple) GetProjectHullVertices() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetProjectHullVertices())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.MeshConvexDecompositionSettings
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetMaxConcavity(max_concavity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_concavity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_max_concavity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxConcavity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_max_concavity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSymmetryPlanesClippingBias(symmetry_planes_clipping_bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, symmetry_planes_clipping_bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_symmetry_planes_clipping_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSymmetryPlanesClippingBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_symmetry_planes_clipping_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRevolutionAxesClippingBias(revolution_axes_clipping_bias gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, revolution_axes_clipping_bias)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_revolution_axes_clipping_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRevolutionAxesClippingBias() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_revolution_axes_clipping_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinVolumePerConvexHull(min_volume_per_convex_hull gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, min_volume_per_convex_hull)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_min_volume_per_convex_hull, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinVolumePerConvexHull() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_min_volume_per_convex_hull, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetResolution(min_volume_per_convex_hull gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, min_volume_per_convex_hull)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetResolution() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxNumVerticesPerConvexHull(max_num_vertices_per_convex_hull gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_num_vertices_per_convex_hull)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_max_num_vertices_per_convex_hull, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxNumVerticesPerConvexHull() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_max_num_vertices_per_convex_hull, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlaneDownsampling(plane_downsampling gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, plane_downsampling)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_plane_downsampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlaneDownsampling() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_plane_downsampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConvexHullDownsampling(convex_hull_downsampling gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, convex_hull_downsampling)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_convex_hull_downsampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConvexHullDownsampling() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_convex_hull_downsampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNormalizeMesh(normalize_mesh bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, normalize_mesh)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_normalize_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNormalizeMesh() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_normalize_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMode(mode classdb.MeshConvexDecompositionSettingsMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMode() classdb.MeshConvexDecompositionSettingsMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MeshConvexDecompositionSettingsMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConvexHullApproximation(convex_hull_approximation bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, convex_hull_approximation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_convex_hull_approximation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConvexHullApproximation() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_convex_hull_approximation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxConvexHulls(max_convex_hulls gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_convex_hulls)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_max_convex_hulls, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxConvexHulls() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_max_convex_hulls, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProjectHullVertices(project_hull_vertices bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, project_hull_vertices)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_set_project_hull_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProjectHullVertices() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MeshConvexDecompositionSettings.Bind_get_project_hull_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsMeshConvexDecompositionSettings() Expert { return self[0].AsMeshConvexDecompositionSettings() }


//go:nosplit
func (self Simple) AsMeshConvexDecompositionSettings() Simple { return self[0].AsMeshConvexDecompositionSettings() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("MeshConvexDecompositionSettings", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Mode = classdb.MeshConvexDecompositionSettingsMode

const (
/*Constant for voxel-based approximate convex decomposition.*/
	ConvexDecompositionModeVoxel Mode = 0
/*Constant for tetrahedron-based approximate convex decomposition.*/
	ConvexDecompositionModeTetrahedron Mode = 1
)
