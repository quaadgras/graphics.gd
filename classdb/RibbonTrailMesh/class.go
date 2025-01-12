// Package RibbonTrailMesh provides methods for working with RibbonTrailMesh object instances.
package RibbonTrailMesh

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/PrimitiveMesh"
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[RibbonTrailMesh] represents a straight ribbon-shaped mesh with variable width. The ribbon is composed of a number of flat or cross-shaped sections, each with the same [member section_length] and number of [member section_segments]. A [member curve] is sampled along the total length of the ribbon, meaning that the curve determines the size of the ribbon along its length.
This primitive mesh is usually used for particle trails.
*/
type Instance [1]gdclass.RibbonTrailMesh

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRibbonTrailMesh() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RibbonTrailMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RibbonTrailMesh"))
	casted := Instance{*(*gdclass.RibbonTrailMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Shape() gdclass.RibbonTrailMeshShape {
	return gdclass.RibbonTrailMeshShape(class(self).GetShape())
}

func (self Instance) SetShape(value gdclass.RibbonTrailMeshShape) {
	class(self).SetShape(value)
}

func (self Instance) Size() Float.X {
	return Float.X(Float.X(class(self).GetSize()))
}

func (self Instance) SetSize(value Float.X) {
	class(self).SetSize(gd.Float(value))
}

func (self Instance) Sections() int {
	return int(int(class(self).GetSections()))
}

func (self Instance) SetSections(value int) {
	class(self).SetSections(gd.Int(value))
}

func (self Instance) SectionLength() Float.X {
	return Float.X(Float.X(class(self).GetSectionLength()))
}

func (self Instance) SetSectionLength(value Float.X) {
	class(self).SetSectionLength(gd.Float(value))
}

func (self Instance) SectionSegments() int {
	return int(int(class(self).GetSectionSegments()))
}

func (self Instance) SetSectionSegments(value int) {
	class(self).SetSectionSegments(gd.Int(value))
}

func (self Instance) Curve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetCurve())
}

func (self Instance) SetCurve(value [1]gdclass.Curve) {
	class(self).SetCurve(value)
}

//go:nosplit
func (self class) SetSize(size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSections(sections gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, sections)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_set_sections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSections() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_get_sections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSectionLength(section_length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, section_length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_set_section_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSectionLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_get_section_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSectionSegments(section_segments gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, section_segments)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_set_section_segments, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSectionSegments() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_get_section_segments, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurve(curve [1]gdclass.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurve() [1]gdclass.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShape(shape gdclass.RibbonTrailMeshShape) {
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShape() gdclass.RibbonTrailMeshShape {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RibbonTrailMeshShape](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RibbonTrailMesh.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRibbonTrailMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRibbonTrailMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("RibbonTrailMesh", func(ptr gd.Object) any {
		return [1]gdclass.RibbonTrailMesh{*(*gdclass.RibbonTrailMesh)(unsafe.Pointer(&ptr))}
	})
}

type Shape = gdclass.RibbonTrailMeshShape

const (
	/*Gives the mesh a single flat face.*/
	ShapeFlat Shape = 0
	/*Gives the mesh two perpendicular flat faces, making a cross shape.*/
	ShapeCross Shape = 1
)
