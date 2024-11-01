package MeshConvexDecompositionSettings

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Parameters to be used with a [Mesh] convex decomposition operation.
*/
type Instance [1]classdb.MeshConvexDecompositionSettings

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MeshConvexDecompositionSettings

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MeshConvexDecompositionSettings"))
	return Instance{classdb.MeshConvexDecompositionSettings(object)}
}

func (self Instance) MaxConcavity() float64 {
	return float64(float64(class(self).GetMaxConcavity()))
}

func (self Instance) SetMaxConcavity(value float64) {
	class(self).SetMaxConcavity(gd.Float(value))
}

func (self Instance) SymmetryPlanesClippingBias() float64 {
	return float64(float64(class(self).GetSymmetryPlanesClippingBias()))
}

func (self Instance) SetSymmetryPlanesClippingBias(value float64) {
	class(self).SetSymmetryPlanesClippingBias(gd.Float(value))
}

func (self Instance) RevolutionAxesClippingBias() float64 {
	return float64(float64(class(self).GetRevolutionAxesClippingBias()))
}

func (self Instance) SetRevolutionAxesClippingBias(value float64) {
	class(self).SetRevolutionAxesClippingBias(gd.Float(value))
}

func (self Instance) MinVolumePerConvexHull() float64 {
	return float64(float64(class(self).GetMinVolumePerConvexHull()))
}

func (self Instance) SetMinVolumePerConvexHull(value float64) {
	class(self).SetMinVolumePerConvexHull(gd.Float(value))
}

func (self Instance) Resolution() int {
	return int(int(class(self).GetResolution()))
}

func (self Instance) SetResolution(value int) {
	class(self).SetResolution(gd.Int(value))
}

func (self Instance) MaxNumVerticesPerConvexHull() int {
	return int(int(class(self).GetMaxNumVerticesPerConvexHull()))
}

func (self Instance) SetMaxNumVerticesPerConvexHull(value int) {
	class(self).SetMaxNumVerticesPerConvexHull(gd.Int(value))
}

func (self Instance) PlaneDownsampling() int {
	return int(int(class(self).GetPlaneDownsampling()))
}

func (self Instance) SetPlaneDownsampling(value int) {
	class(self).SetPlaneDownsampling(gd.Int(value))
}

func (self Instance) ConvexHullDownsampling() int {
	return int(int(class(self).GetConvexHullDownsampling()))
}

func (self Instance) SetConvexHullDownsampling(value int) {
	class(self).SetConvexHullDownsampling(gd.Int(value))
}

func (self Instance) NormalizeMesh() bool {
	return bool(class(self).GetNormalizeMesh())
}

func (self Instance) SetNormalizeMesh(value bool) {
	class(self).SetNormalizeMesh(value)
}

func (self Instance) Mode() classdb.MeshConvexDecompositionSettingsMode {
	return classdb.MeshConvexDecompositionSettingsMode(class(self).GetMode())
}

func (self Instance) SetMode(value classdb.MeshConvexDecompositionSettingsMode) {
	class(self).SetMode(value)
}

func (self Instance) ConvexHullApproximation() bool {
	return bool(class(self).GetConvexHullApproximation())
}

func (self Instance) SetConvexHullApproximation(value bool) {
	class(self).SetConvexHullApproximation(value)
}

func (self Instance) MaxConvexHulls() int {
	return int(int(class(self).GetMaxConvexHulls()))
}

func (self Instance) SetMaxConvexHulls(value int) {
	class(self).SetMaxConvexHulls(gd.Int(value))
}

func (self Instance) ProjectHullVertices() bool {
	return bool(class(self).GetProjectHullVertices())
}

func (self Instance) SetProjectHullVertices(value bool) {
	class(self).SetProjectHullVertices(value)
}

//go:nosplit
func (self class) SetMaxConcavity(max_concavity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, max_concavity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_max_concavity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxConcavity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_max_concavity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSymmetryPlanesClippingBias(symmetry_planes_clipping_bias gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, symmetry_planes_clipping_bias)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_symmetry_planes_clipping_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSymmetryPlanesClippingBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_symmetry_planes_clipping_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRevolutionAxesClippingBias(revolution_axes_clipping_bias gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, revolution_axes_clipping_bias)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_revolution_axes_clipping_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRevolutionAxesClippingBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_revolution_axes_clipping_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinVolumePerConvexHull(min_volume_per_convex_hull gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, min_volume_per_convex_hull)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_min_volume_per_convex_hull, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinVolumePerConvexHull() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_min_volume_per_convex_hull, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetResolution(min_volume_per_convex_hull gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, min_volume_per_convex_hull)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetResolution() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxNumVerticesPerConvexHull(max_num_vertices_per_convex_hull gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_num_vertices_per_convex_hull)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_max_num_vertices_per_convex_hull, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxNumVerticesPerConvexHull() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_max_num_vertices_per_convex_hull, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlaneDownsampling(plane_downsampling gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, plane_downsampling)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_plane_downsampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlaneDownsampling() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_plane_downsampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConvexHullDownsampling(convex_hull_downsampling gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, convex_hull_downsampling)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_convex_hull_downsampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetConvexHullDownsampling() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_convex_hull_downsampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalizeMesh(normalize_mesh bool) {
	var frame = callframe.New()
	callframe.Arg(frame, normalize_mesh)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_normalize_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNormalizeMesh() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_normalize_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMode(mode classdb.MeshConvexDecompositionSettingsMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMode() classdb.MeshConvexDecompositionSettingsMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MeshConvexDecompositionSettingsMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConvexHullApproximation(convex_hull_approximation bool) {
	var frame = callframe.New()
	callframe.Arg(frame, convex_hull_approximation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_convex_hull_approximation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetConvexHullApproximation() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_convex_hull_approximation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxConvexHulls(max_convex_hulls gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_convex_hulls)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_max_convex_hulls, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxConvexHulls() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_max_convex_hulls, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProjectHullVertices(project_hull_vertices bool) {
	var frame = callframe.New()
	callframe.Arg(frame, project_hull_vertices)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_set_project_hull_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetProjectHullVertices() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshConvexDecompositionSettings.Bind_get_project_hull_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsMeshConvexDecompositionSettings() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMeshConvexDecompositionSettings() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("MeshConvexDecompositionSettings", func(ptr gd.Object) any { return classdb.MeshConvexDecompositionSettings(ptr) })
}

type Mode = classdb.MeshConvexDecompositionSettingsMode

const (
	/*Constant for voxel-based approximate convex decomposition.*/
	ConvexDecompositionModeVoxel Mode = 0
	/*Constant for tetrahedron-based approximate convex decomposition.*/
	ConvexDecompositionModeTetrahedron Mode = 1
)
