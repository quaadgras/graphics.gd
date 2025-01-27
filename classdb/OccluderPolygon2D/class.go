// Package OccluderPolygon2D provides methods for working with OccluderPolygon2D object instances.
package OccluderPolygon2D

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
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"

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
var _ Path.ToNode

/*
Editor facility that helps you draw a 2D polygon used as resource for [LightOccluder2D].
*/
type Instance [1]gdclass.OccluderPolygon2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOccluderPolygon2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OccluderPolygon2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OccluderPolygon2D"))
	casted := Instance{*(*gdclass.OccluderPolygon2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Closed() bool {
	return bool(class(self).IsClosed())
}

func (self Instance) SetClosed(value bool) {
	class(self).SetClosed(value)
}

func (self Instance) CullMode() gdclass.OccluderPolygon2DCullMode {
	return gdclass.OccluderPolygon2DCullMode(class(self).GetCullMode())
}

func (self Instance) SetCullMode(value gdclass.OccluderPolygon2DCullMode) {
	class(self).SetCullMode(value)
}

func (self Instance) Polygon() []Vector2.XY {
	return []Vector2.XY(class(self).GetPolygon().AsSlice())
}

func (self Instance) SetPolygon(value []Vector2.XY) {
	class(self).SetPolygon(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

//go:nosplit
func (self class) SetClosed(closed bool) { //gd:OccluderPolygon2D.set_closed
	var frame = callframe.New()
	callframe.Arg(frame, closed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_set_closed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsClosed() bool { //gd:OccluderPolygon2D.is_closed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_is_closed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMode(cull_mode gdclass.OccluderPolygon2DCullMode) { //gd:OccluderPolygon2D.set_cull_mode
	var frame = callframe.New()
	callframe.Arg(frame, cull_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_set_cull_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMode() gdclass.OccluderPolygon2DCullMode { //gd:OccluderPolygon2D.get_cull_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.OccluderPolygon2DCullMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_get_cull_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array) { //gd:OccluderPolygon2D.set_polygon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolygon() gd.PackedVector2Array { //gd:OccluderPolygon2D.get_polygon
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsOccluderPolygon2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOccluderPolygon2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("OccluderPolygon2D", func(ptr gd.Object) any {
		return [1]gdclass.OccluderPolygon2D{*(*gdclass.OccluderPolygon2D)(unsafe.Pointer(&ptr))}
	})
}

type CullMode = gdclass.OccluderPolygon2DCullMode //gd:OccluderPolygon2D.CullMode

const (
	/*Culling is disabled. See [member cull_mode].*/
	CullDisabled CullMode = 0
	/*Culling is performed in the clockwise direction. See [member cull_mode].*/
	CullClockwise CullMode = 1
	/*Culling is performed in the counterclockwise direction. See [member cull_mode].*/
	CullCounterClockwise CullMode = 2
)
