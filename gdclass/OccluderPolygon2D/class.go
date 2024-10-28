package OccluderPolygon2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Editor facility that helps you draw a 2D polygon used as resource for [LightOccluder2D].

*/
type Go [1]classdb.OccluderPolygon2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OccluderPolygon2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OccluderPolygon2D"))
	return Go{classdb.OccluderPolygon2D(object)}
}

func (self Go) Closed() bool {
		return bool(class(self).IsClosed())
}

func (self Go) SetClosed(value bool) {
	class(self).SetClosed(value)
}

func (self Go) CullMode() classdb.OccluderPolygon2DCullMode {
		return classdb.OccluderPolygon2DCullMode(class(self).GetCullMode())
}

func (self Go) SetCullMode(value classdb.OccluderPolygon2DCullMode) {
	class(self).SetCullMode(value)
}

func (self Go) Polygon() []gd.Vector2 {
		return []gd.Vector2(class(self).GetPolygon().AsSlice())
}

func (self Go) SetPolygon(value []gd.Vector2) {
	class(self).SetPolygon(gd.NewPackedVector2Slice(value))
}

//go:nosplit
func (self class) SetClosed(closed bool)  {
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
func (self class) SetCullMode(cull_mode classdb.OccluderPolygon2DCullMode)  {
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
func (self class) SetPolygon(polygon gd.PackedVector2Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPolygon() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OccluderPolygon2D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsOccluderPolygon2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOccluderPolygon2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("OccluderPolygon2D", func(ptr gd.Object) any { return classdb.OccluderPolygon2D(ptr) })}
type CullMode = classdb.OccluderPolygon2DCullMode

const (
/*Culling is disabled. See [member cull_mode].*/
	CullDisabled CullMode = 0
/*Culling is performed in the clockwise direction. See [member cull_mode].*/
	CullClockwise CullMode = 1
/*Culling is performed in the counterclockwise direction. See [member cull_mode].*/
	CullCounterClockwise CullMode = 2
)
