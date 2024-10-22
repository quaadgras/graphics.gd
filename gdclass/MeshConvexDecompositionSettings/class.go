package MeshConvexDecompositionSettings

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Parameters to be used with a [Mesh] convex decomposition operation.

*/
type Go [1]classdb.MeshConvexDecompositionSettings
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MeshConvexDecompositionSettings
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("MeshConvexDecompositionSettings"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) MaxConcavity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMaxConcavity()))
}

func (self Go) SetMaxConcavity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxConcavity(gd.Float(value))
}

func (self Go) SymmetryPlanesClippingBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSymmetryPlanesClippingBias()))
}

func (self Go) SetSymmetryPlanesClippingBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSymmetryPlanesClippingBias(gd.Float(value))
}

func (self Go) RevolutionAxesClippingBias() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRevolutionAxesClippingBias()))
}

func (self Go) SetRevolutionAxesClippingBias(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRevolutionAxesClippingBias(gd.Float(value))
}

func (self Go) MinVolumePerConvexHull() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMinVolumePerConvexHull()))
}

func (self Go) SetMinVolumePerConvexHull(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinVolumePerConvexHull(gd.Float(value))
}

func (self Go) Resolution() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetResolution()))
}

func (self Go) SetResolution(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetResolution(gd.Int(value))
}

func (self Go) MaxNumVerticesPerConvexHull() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxNumVerticesPerConvexHull()))
}

func (self Go) SetMaxNumVerticesPerConvexHull(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxNumVerticesPerConvexHull(gd.Int(value))
}

func (self Go) PlaneDownsampling() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPlaneDownsampling()))
}

func (self Go) SetPlaneDownsampling(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPlaneDownsampling(gd.Int(value))
}

func (self Go) ConvexHullDownsampling() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetConvexHullDownsampling()))
}

func (self Go) SetConvexHullDownsampling(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConvexHullDownsampling(gd.Int(value))
}

func (self Go) NormalizeMesh() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetNormalizeMesh())
}

func (self Go) SetNormalizeMesh(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNormalizeMesh(value)
}

func (self Go) Mode() classdb.MeshConvexDecompositionSettingsMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.MeshConvexDecompositionSettingsMode(class(self).GetMode())
}

func (self Go) SetMode(value classdb.MeshConvexDecompositionSettingsMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMode(value)
}

func (self Go) ConvexHullApproximation() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetConvexHullApproximation())
}

func (self Go) SetConvexHullApproximation(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConvexHullApproximation(value)
}

func (self Go) MaxConvexHulls() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxConvexHulls()))
}

func (self Go) SetMaxConvexHulls(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxConvexHulls(gd.Int(value))
}

func (self Go) ProjectHullVertices() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetProjectHullVertices())
}

func (self Go) SetProjectHullVertices(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProjectHullVertices(value)
}

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
func (self class) AsMeshConvexDecompositionSettings() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMeshConvexDecompositionSettings() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("MeshConvexDecompositionSettings", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Mode = classdb.MeshConvexDecompositionSettingsMode

const (
/*Constant for voxel-based approximate convex decomposition.*/
	ConvexDecompositionModeVoxel Mode = 0
/*Constant for tetrahedron-based approximate convex decomposition.*/
	ConvexDecompositionModeTetrahedron Mode = 1
)
