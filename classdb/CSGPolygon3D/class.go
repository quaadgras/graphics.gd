// Package CSGPolygon3D provides methods for working with CSGPolygon3D object instances.
package CSGPolygon3D

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
import "graphics.gd/classdb/CSGPrimitive3D"
import "graphics.gd/classdb/CSGShape3D"
import "graphics.gd/classdb/GeometryInstance3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/NodePath"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
An array of 2D points is extruded to quickly and easily create a variety of 3D meshes. See also [CSGMesh3D] for using 3D meshes as CSG nodes.
[b]Note:[/b] CSG nodes are intended to be used for level prototyping. Creating CSG nodes has a significant CPU cost compared to creating a [MeshInstance3D] with a [PrimitiveMesh]. Moving a CSG node within another CSG node also has a significant CPU cost, so it should be avoided during gameplay.
*/
type Instance [1]gdclass.CSGPolygon3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCSGPolygon3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CSGPolygon3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CSGPolygon3D"))
	casted := Instance{*(*gdclass.CSGPolygon3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Polygon() []Vector2.XY {
	return []Vector2.XY(class(self).GetPolygon().AsSlice())
}

func (self Instance) SetPolygon(value []Vector2.XY) {
	class(self).SetPolygon(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

func (self Instance) Mode() gdclass.CSGPolygon3DMode {
	return gdclass.CSGPolygon3DMode(class(self).GetMode())
}

func (self Instance) SetMode(value gdclass.CSGPolygon3DMode) {
	class(self).SetMode(value)
}

func (self Instance) Depth() Float.X {
	return Float.X(Float.X(class(self).GetDepth()))
}

func (self Instance) SetDepth(value Float.X) {
	class(self).SetDepth(gd.Float(value))
}

func (self Instance) SpinDegrees() Float.X {
	return Float.X(Float.X(class(self).GetSpinDegrees()))
}

func (self Instance) SetSpinDegrees(value Float.X) {
	class(self).SetSpinDegrees(gd.Float(value))
}

func (self Instance) SpinSides() int {
	return int(int(class(self).GetSpinSides()))
}

func (self Instance) SetSpinSides(value int) {
	class(self).SetSpinSides(gd.Int(value))
}

func (self Instance) PathNode() NodePath.String {
	return NodePath.String(class(self).GetPathNode().String())
}

func (self Instance) SetPathNode(value NodePath.String) {
	class(self).SetPathNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) PathIntervalType() gdclass.CSGPolygon3DPathIntervalType {
	return gdclass.CSGPolygon3DPathIntervalType(class(self).GetPathIntervalType())
}

func (self Instance) SetPathIntervalType(value gdclass.CSGPolygon3DPathIntervalType) {
	class(self).SetPathIntervalType(value)
}

func (self Instance) PathInterval() Float.X {
	return Float.X(Float.X(class(self).GetPathInterval()))
}

func (self Instance) SetPathInterval(value Float.X) {
	class(self).SetPathInterval(gd.Float(value))
}

func (self Instance) PathSimplifyAngle() Float.X {
	return Float.X(Float.X(class(self).GetPathSimplifyAngle()))
}

func (self Instance) SetPathSimplifyAngle(value Float.X) {
	class(self).SetPathSimplifyAngle(gd.Float(value))
}

func (self Instance) PathRotation() gdclass.CSGPolygon3DPathRotation {
	return gdclass.CSGPolygon3DPathRotation(class(self).GetPathRotation())
}

func (self Instance) SetPathRotation(value gdclass.CSGPolygon3DPathRotation) {
	class(self).SetPathRotation(value)
}

func (self Instance) PathLocal() bool {
	return bool(class(self).IsPathLocal())
}

func (self Instance) SetPathLocal(value bool) {
	class(self).SetPathLocal(value)
}

func (self Instance) PathContinuousU() bool {
	return bool(class(self).IsPathContinuousU())
}

func (self Instance) SetPathContinuousU(value bool) {
	class(self).SetPathContinuousU(value)
}

func (self Instance) PathUDistance() Float.X {
	return Float.X(Float.X(class(self).GetPathUDistance()))
}

func (self Instance) SetPathUDistance(value Float.X) {
	class(self).SetPathUDistance(gd.Float(value))
}

func (self Instance) PathJoined() bool {
	return bool(class(self).IsPathJoined())
}

func (self Instance) SetPathJoined(value bool) {
	class(self).SetPathJoined(value)
}

func (self Instance) SmoothFaces() bool {
	return bool(class(self).GetSmoothFaces())
}

func (self Instance) SetSmoothFaces(value bool) {
	class(self).SetSmoothFaces(value)
}

func (self Instance) Material() [1]gdclass.Material {
	return [1]gdclass.Material(class(self).GetMaterial())
}

func (self Instance) SetMaterial(value [1]gdclass.Material) {
	class(self).SetMaterial(value)
}

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array) { //gd:CSGPolygon3D.set_polygon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolygon() gd.PackedVector2Array { //gd:CSGPolygon3D.get_polygon
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMode(mode gdclass.CSGPolygon3DMode) { //gd:CSGPolygon3D.set_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMode() gdclass.CSGPolygon3DMode { //gd:CSGPolygon3D.get_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CSGPolygon3DMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepth(depth gd.Float) { //gd:CSGPolygon3D.set_depth
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepth() gd.Float { //gd:CSGPolygon3D.get_depth
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpinDegrees(degrees gd.Float) { //gd:CSGPolygon3D.set_spin_degrees
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_spin_degrees, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpinDegrees() gd.Float { //gd:CSGPolygon3D.get_spin_degrees
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_spin_degrees, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpinSides(spin_sides gd.Int) { //gd:CSGPolygon3D.set_spin_sides
	var frame = callframe.New()
	callframe.Arg(frame, spin_sides)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_spin_sides, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpinSides() gd.Int { //gd:CSGPolygon3D.get_spin_sides
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_spin_sides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathNode(path gd.NodePath) { //gd:CSGPolygon3D.set_path_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_path_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathNode() gd.NodePath { //gd:CSGPolygon3D.get_path_node
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_path_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathIntervalType(interval_type gdclass.CSGPolygon3DPathIntervalType) { //gd:CSGPolygon3D.set_path_interval_type
	var frame = callframe.New()
	callframe.Arg(frame, interval_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_path_interval_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathIntervalType() gdclass.CSGPolygon3DPathIntervalType { //gd:CSGPolygon3D.get_path_interval_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CSGPolygon3DPathIntervalType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_path_interval_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathInterval(interval gd.Float) { //gd:CSGPolygon3D.set_path_interval
	var frame = callframe.New()
	callframe.Arg(frame, interval)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_path_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathInterval() gd.Float { //gd:CSGPolygon3D.get_path_interval
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_path_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathSimplifyAngle(degrees gd.Float) { //gd:CSGPolygon3D.set_path_simplify_angle
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_path_simplify_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathSimplifyAngle() gd.Float { //gd:CSGPolygon3D.get_path_simplify_angle
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_path_simplify_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathRotation(path_rotation gdclass.CSGPolygon3DPathRotation) { //gd:CSGPolygon3D.set_path_rotation
	var frame = callframe.New()
	callframe.Arg(frame, path_rotation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_path_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathRotation() gdclass.CSGPolygon3DPathRotation { //gd:CSGPolygon3D.get_path_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CSGPolygon3DPathRotation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_path_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathLocal(enable bool) { //gd:CSGPolygon3D.set_path_local
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_path_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPathLocal() bool { //gd:CSGPolygon3D.is_path_local
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_is_path_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathContinuousU(enable bool) { //gd:CSGPolygon3D.set_path_continuous_u
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_path_continuous_u, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPathContinuousU() bool { //gd:CSGPolygon3D.is_path_continuous_u
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_is_path_continuous_u, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathUDistance(distance gd.Float) { //gd:CSGPolygon3D.set_path_u_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_path_u_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathUDistance() gd.Float { //gd:CSGPolygon3D.get_path_u_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_path_u_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathJoined(enable bool) { //gd:CSGPolygon3D.set_path_joined
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_path_joined, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPathJoined() bool { //gd:CSGPolygon3D.is_path_joined
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_is_path_joined, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaterial(material [1]gdclass.Material) { //gd:CSGPolygon3D.set_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaterial() [1]gdclass.Material { //gd:CSGPolygon3D.get_material
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSmoothFaces(smooth_faces bool) { //gd:CSGPolygon3D.set_smooth_faces
	var frame = callframe.New()
	callframe.Arg(frame, smooth_faces)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_set_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSmoothFaces() bool { //gd:CSGPolygon3D.get_smooth_faces
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGPolygon3D.Bind_get_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCSGPolygon3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCSGPolygon3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsCSGPrimitive3D() CSGPrimitive3D.Advanced {
	return *((*CSGPrimitive3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCSGPrimitive3D() CSGPrimitive3D.Instance {
	return *((*CSGPrimitive3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCSGShape3D() CSGShape3D.Advanced {
	return *((*CSGShape3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCSGShape3D() CSGShape3D.Instance {
	return *((*CSGShape3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(CSGPrimitive3D.Advanced(self.AsCSGPrimitive3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CSGPrimitive3D.Instance(self.AsCSGPrimitive3D()), name)
	}
}
func init() {
	gdclass.Register("CSGPolygon3D", func(ptr gd.Object) any {
		return [1]gdclass.CSGPolygon3D{*(*gdclass.CSGPolygon3D)(unsafe.Pointer(&ptr))}
	})
}

type Mode = gdclass.CSGPolygon3DMode //gd:CSGPolygon3D.Mode

const (
	/*The [member polygon] shape is extruded along the negative Z axis.*/
	ModeDepth Mode = 0
	/*The [member polygon] shape is extruded by rotating it around the Y axis.*/
	ModeSpin Mode = 1
	/*The [member polygon] shape is extruded along the [Path3D] specified in [member path_node].*/
	ModePath Mode = 2
)

type PathRotation = gdclass.CSGPolygon3DPathRotation //gd:CSGPolygon3D.PathRotation

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

type PathIntervalType = gdclass.CSGPolygon3DPathIntervalType //gd:CSGPolygon3D.PathIntervalType

const (
	/*When [member mode] is set to [constant MODE_PATH], [member path_interval] will determine the distance, in meters, each interval of the path will extrude.*/
	PathIntervalDistance PathIntervalType = 0
	/*When [member mode] is set to [constant MODE_PATH], [member path_interval] will subdivide the polygons along the path.*/
	PathIntervalSubdivide PathIntervalType = 1
)
