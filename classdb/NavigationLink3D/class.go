// Package NavigationLink3D provides methods for working with NavigationLink3D object instances.
package NavigationLink3D

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
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"

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
A link between two positions on [NavigationRegion3D]s that agents can be routed through. These positions can be on the same [NavigationRegion3D] or on two different ones. Links are useful to express navigation methods other than traveling along the surface of the navigation mesh, such as ziplines, teleporters, or gaps that can be jumped across.
*/
type Instance [1]gdclass.NavigationLink3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNavigationLink3D() Instance
}

/*
Returns the [RID] of this link on the [NavigationServer3D].
*/
func (self Instance) GetRid() Resource.ID { //gd:NavigationLink3D.get_rid
	return Resource.ID(class(self).GetRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetNavigationLayerValue(layer_number int, value bool) { //gd:NavigationLink3D.set_navigation_layer_value
	class(self).SetNavigationLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetNavigationLayerValue(layer_number int) bool { //gd:NavigationLink3D.get_navigation_layer_value
	return bool(class(self).GetNavigationLayerValue(gd.Int(layer_number)))
}

/*
Sets the [member start_position] that is relative to the link from a global [param position].
*/
func (self Instance) SetGlobalStartPosition(position Vector3.XYZ) { //gd:NavigationLink3D.set_global_start_position
	class(self).SetGlobalStartPosition(gd.Vector3(position))
}

/*
Returns the [member start_position] that is relative to the link as a global position.
*/
func (self Instance) GetGlobalStartPosition() Vector3.XYZ { //gd:NavigationLink3D.get_global_start_position
	return Vector3.XYZ(class(self).GetGlobalStartPosition())
}

/*
Sets the [member end_position] that is relative to the link from a global [param position].
*/
func (self Instance) SetGlobalEndPosition(position Vector3.XYZ) { //gd:NavigationLink3D.set_global_end_position
	class(self).SetGlobalEndPosition(gd.Vector3(position))
}

/*
Returns the [member end_position] that is relative to the link as a global position.
*/
func (self Instance) GetGlobalEndPosition() Vector3.XYZ { //gd:NavigationLink3D.get_global_end_position
	return Vector3.XYZ(class(self).GetGlobalEndPosition())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationLink3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationLink3D"))
	casted := Instance{*(*gdclass.NavigationLink3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Enabled() bool {
	return bool(class(self).IsEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) Bidirectional() bool {
	return bool(class(self).IsBidirectional())
}

func (self Instance) SetBidirectional(value bool) {
	class(self).SetBidirectional(value)
}

func (self Instance) NavigationLayers() int {
	return int(int(class(self).GetNavigationLayers()))
}

func (self Instance) SetNavigationLayers(value int) {
	class(self).SetNavigationLayers(gd.Int(value))
}

func (self Instance) StartPosition() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetStartPosition())
}

func (self Instance) SetStartPosition(value Vector3.XYZ) {
	class(self).SetStartPosition(gd.Vector3(value))
}

func (self Instance) EndPosition() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetEndPosition())
}

func (self Instance) SetEndPosition(value Vector3.XYZ) {
	class(self).SetEndPosition(gd.Vector3(value))
}

func (self Instance) EnterCost() Float.X {
	return Float.X(Float.X(class(self).GetEnterCost()))
}

func (self Instance) SetEnterCost(value Float.X) {
	class(self).SetEnterCost(gd.Float(value))
}

func (self Instance) TravelCost() Float.X {
	return Float.X(Float.X(class(self).GetTravelCost()))
}

func (self Instance) SetTravelCost(value Float.X) {
	class(self).SetTravelCost(gd.Float(value))
}

/*
Returns the [RID] of this link on the [NavigationServer3D].
*/
//go:nosplit
func (self class) GetRid() gd.RID { //gd:NavigationLink3D.get_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnabled(enabled bool) { //gd:NavigationLink3D.set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEnabled() bool { //gd:NavigationLink3D.is_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBidirectional(bidirectional bool) { //gd:NavigationLink3D.set_bidirectional
	var frame = callframe.New()
	callframe.Arg(frame, bidirectional)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_bidirectional, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsBidirectional() bool { //gd:NavigationLink3D.is_bidirectional
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_is_bidirectional, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int) { //gd:NavigationLink3D.set_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationLayers() gd.Int { //gd:NavigationLink3D.get_navigation_layers
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerValue(layer_number gd.Int, value bool) { //gd:NavigationLink3D.set_navigation_layer_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetNavigationLayerValue(layer_number gd.Int) bool { //gd:NavigationLink3D.get_navigation_layer_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_get_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStartPosition(position gd.Vector3) { //gd:NavigationLink3D.set_start_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStartPosition() gd.Vector3 { //gd:NavigationLink3D.get_start_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_get_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEndPosition(position gd.Vector3) { //gd:NavigationLink3D.set_end_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEndPosition() gd.Vector3 { //gd:NavigationLink3D.get_end_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_get_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [member start_position] that is relative to the link from a global [param position].
*/
//go:nosplit
func (self class) SetGlobalStartPosition(position gd.Vector3) { //gd:NavigationLink3D.set_global_start_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_global_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [member start_position] that is relative to the link as a global position.
*/
//go:nosplit
func (self class) GetGlobalStartPosition() gd.Vector3 { //gd:NavigationLink3D.get_global_start_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_get_global_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [member end_position] that is relative to the link from a global [param position].
*/
//go:nosplit
func (self class) SetGlobalEndPosition(position gd.Vector3) { //gd:NavigationLink3D.set_global_end_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_global_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [member end_position] that is relative to the link as a global position.
*/
//go:nosplit
func (self class) GetGlobalEndPosition() gd.Vector3 { //gd:NavigationLink3D.get_global_end_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_get_global_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnterCost(enter_cost gd.Float) { //gd:NavigationLink3D.set_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, enter_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnterCost() gd.Float { //gd:NavigationLink3D.get_enter_cost
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTravelCost(travel_cost gd.Float) { //gd:NavigationLink3D.set_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, travel_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTravelCost() gd.Float { //gd:NavigationLink3D.get_travel_cost
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink3D.Bind_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNavigationLink3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNavigationLink3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced       { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance    { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced           { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance        { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("NavigationLink3D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationLink3D{*(*gdclass.NavigationLink3D)(unsafe.Pointer(&ptr))}
	})
}
