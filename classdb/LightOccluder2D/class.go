// Package LightOccluder2D provides methods for working with LightOccluder2D object instances.
package LightOccluder2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Occludes light cast by a Light2D, casting shadows. The LightOccluder2D must be provided with an [OccluderPolygon2D] in order for the shadow to be computed.
*/
type Instance [1]gdclass.LightOccluder2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLightOccluder2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.LightOccluder2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("LightOccluder2D"))
	casted := Instance{*(*gdclass.LightOccluder2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Occluder() [1]gdclass.OccluderPolygon2D {
	return [1]gdclass.OccluderPolygon2D(class(self).GetOccluderPolygon())
}

func (self Instance) SetOccluder(value [1]gdclass.OccluderPolygon2D) {
	class(self).SetOccluderPolygon(value)
}

func (self Instance) SdfCollision() bool {
	return bool(class(self).IsSetAsSdfCollision())
}

func (self Instance) SetSdfCollision(value bool) {
	class(self).SetAsSdfCollision(value)
}

func (self Instance) OccluderLightMask() int {
	return int(int(class(self).GetOccluderLightMask()))
}

func (self Instance) SetOccluderLightMask(value int) {
	class(self).SetOccluderLightMask(gd.Int(value))
}

//go:nosplit
func (self class) SetOccluderPolygon(polygon [1]gdclass.OccluderPolygon2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightOccluder2D.Bind_set_occluder_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOccluderPolygon() [1]gdclass.OccluderPolygon2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightOccluder2D.Bind_get_occluder_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OccluderPolygon2D{gd.PointerWithOwnershipTransferredToGo[gdclass.OccluderPolygon2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOccluderLightMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightOccluder2D.Bind_set_occluder_light_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOccluderLightMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightOccluder2D.Bind_get_occluder_light_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAsSdfCollision(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightOccluder2D.Bind_set_as_sdf_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSetAsSdfCollision() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightOccluder2D.Bind_is_set_as_sdf_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLightOccluder2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLightOccluder2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced      { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance   { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("LightOccluder2D", func(ptr gd.Object) any {
		return [1]gdclass.LightOccluder2D{*(*gdclass.LightOccluder2D)(unsafe.Pointer(&ptr))}
	})
}
