package OccluderPolygon2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Editor facility that helps you draw a 2D polygon used as resource for [LightOccluder2D].
*/
type Instance [1]classdb.OccluderPolygon2D
type Any interface {
	gd.IsClass
	AsOccluderPolygon2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OccluderPolygon2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OccluderPolygon2D"))
	return Instance{classdb.OccluderPolygon2D(object)}
}

func (self Instance) Closed() bool {
	return bool(class(self).IsClosed())
}

func (self Instance) SetClosed(value bool) {
	class(self).SetClosed(value)
}

func (self Instance) CullMode() classdb.OccluderPolygon2DCullMode {
	return classdb.OccluderPolygon2DCullMode(class(self).GetCullMode())
}

func (self Instance) SetCullMode(value classdb.OccluderPolygon2DCullMode) {
	class(self).SetCullMode(value)
}

func (self Instance) Polygon() []Vector2.XY {
	return []Vector2.XY(class(self).GetPolygon().AsSlice())
}

func (self Instance) SetPolygon(value []Vector2.XY) {
	class(self).SetPolygon(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

//go:nosplit
func (self class) SetClosed(closed bool) {
	var frame = callframe.New()
	callframe.Arg(frame, closed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_set_closed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsClosed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_is_closed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMode(cull_mode classdb.OccluderPolygon2DCullMode) {
	var frame = callframe.New()
	callframe.Arg(frame, cull_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_set_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMode() classdb.OccluderPolygon2DCullMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.OccluderPolygon2DCullMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_get_cull_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolygon() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("OccluderPolygon2D", func(ptr gd.Object) any { return classdb.OccluderPolygon2D(ptr) })
}

type CullMode = classdb.OccluderPolygon2DCullMode

const (
	/*Culling is disabled. See [member cull_mode].*/
	CullDisabled CullMode = 0
	/*Culling is performed in the clockwise direction. See [member cull_mode].*/
	CullClockwise CullMode = 1
	/*Culling is performed in the counterclockwise direction. See [member cull_mode].*/
	CullCounterClockwise CullMode = 2
)
