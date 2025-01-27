// Package PlaneMesh provides methods for working with PlaneMesh object instances.
package PlaneMesh

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
import "graphics.gd/classdb/PrimitiveMesh"
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector3"

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

/*
Class representing a planar [PrimitiveMesh]. This flat mesh does not have a thickness. By default, this mesh is aligned on the X and Z axes; this default rotation isn't suited for use with billboarded materials. For billboarded materials, change [member orientation] to [constant FACE_Z].
[b]Note:[/b] When using a large textured [PlaneMesh] (e.g. as a floor), you may stumble upon UV jittering issues depending on the camera angle. To solve this, increase [member subdivide_depth] and [member subdivide_width] until you no longer notice UV jittering.
*/
type Instance [1]gdclass.PlaneMesh

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPlaneMesh() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PlaneMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PlaneMesh"))
	casted := Instance{*(*gdclass.PlaneMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Size() Vector2.XY {
	return Vector2.XY(class(self).GetSize())
}

func (self Instance) SetSize(value Vector2.XY) {
	class(self).SetSize(gd.Vector2(value))
}

func (self Instance) SubdivideWidth() int {
	return int(int(class(self).GetSubdivideWidth()))
}

func (self Instance) SetSubdivideWidth(value int) {
	class(self).SetSubdivideWidth(gd.Int(value))
}

func (self Instance) SubdivideDepth() int {
	return int(int(class(self).GetSubdivideDepth()))
}

func (self Instance) SetSubdivideDepth(value int) {
	class(self).SetSubdivideDepth(gd.Int(value))
}

func (self Instance) CenterOffset() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetCenterOffset())
}

func (self Instance) SetCenterOffset(value Vector3.XYZ) {
	class(self).SetCenterOffset(gd.Vector3(value))
}

func (self Instance) Orientation() gdclass.PlaneMeshOrientation {
	return gdclass.PlaneMeshOrientation(class(self).GetOrientation())
}

func (self Instance) SetOrientation(value gdclass.PlaneMeshOrientation) {
	class(self).SetOrientation(value)
}

//go:nosplit
func (self class) SetSize(size gd.Vector2) { //gd:PlaneMesh.set_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector2 { //gd:PlaneMesh.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubdivideWidth(subdivide gd.Int) { //gd:PlaneMesh.set_subdivide_width
	var frame = callframe.New()
	callframe.Arg(frame, subdivide)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_set_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubdivideWidth() gd.Int { //gd:PlaneMesh.get_subdivide_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_get_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubdivideDepth(subdivide gd.Int) { //gd:PlaneMesh.set_subdivide_depth
	var frame = callframe.New()
	callframe.Arg(frame, subdivide)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_set_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubdivideDepth() gd.Int { //gd:PlaneMesh.get_subdivide_depth
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_get_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCenterOffset(offset gd.Vector3) { //gd:PlaneMesh.set_center_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_set_center_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterOffset() gd.Vector3 { //gd:PlaneMesh.get_center_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_get_center_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOrientation(orientation gdclass.PlaneMeshOrientation) { //gd:PlaneMesh.set_orientation
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_set_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOrientation() gdclass.PlaneMeshOrientation { //gd:PlaneMesh.get_orientation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.PlaneMeshOrientation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaneMesh.Bind_get_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPlaneMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPlaneMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPrimitiveMesh() PrimitiveMesh.Advanced {
	return *((*PrimitiveMesh.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPrimitiveMesh() PrimitiveMesh.Instance {
	return *((*PrimitiveMesh.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMesh() Mesh.Advanced    { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Mesh.Instance { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PrimitiveMesh.Advanced(self.AsPrimitiveMesh()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PrimitiveMesh.Instance(self.AsPrimitiveMesh()), name)
	}
}
func init() {
	gdclass.Register("PlaneMesh", func(ptr gd.Object) any { return [1]gdclass.PlaneMesh{*(*gdclass.PlaneMesh)(unsafe.Pointer(&ptr))} })
}

type Orientation = gdclass.PlaneMeshOrientation //gd:PlaneMesh.Orientation

const (
	/*[PlaneMesh] will face the positive X-axis.*/
	FaceX Orientation = 0
	/*[PlaneMesh] will face the positive Y-axis. This matches the behavior of the [PlaneMesh] in Godot 3.x.*/
	FaceY Orientation = 1
	/*[PlaneMesh] will face the positive Z-axis. This matches the behavior of the QuadMesh in Godot 3.x.*/
	FaceZ Orientation = 2
)
