package NavigationLink2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A link between two positions on [NavigationRegion2D]s that agents can be routed through. These positions can be on the same [NavigationRegion2D] or on two different ones. Links are useful to express navigation methods other than traveling along the surface of the navigation polygon, such as ziplines, teleporters, or gaps that can be jumped across.
*/
type Instance [1]classdb.NavigationLink2D

/*
Returns the [RID] of this link on the [NavigationServer2D].
*/
func (self Instance) GetRid() gd.RID {
	return gd.RID(class(self).GetRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetNavigationLayerValue(layer_number int, value bool) {
	class(self).SetNavigationLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetNavigationLayerValue(layer_number int) bool {
	return bool(class(self).GetNavigationLayerValue(gd.Int(layer_number)))
}

/*
Sets the [member start_position] that is relative to the link from a global [param position].
*/
func (self Instance) SetGlobalStartPosition(position gd.Vector2) {
	class(self).SetGlobalStartPosition(position)
}

/*
Returns the [member start_position] that is relative to the link as a global position.
*/
func (self Instance) GetGlobalStartPosition() gd.Vector2 {
	return gd.Vector2(class(self).GetGlobalStartPosition())
}

/*
Sets the [member end_position] that is relative to the link from a global [param position].
*/
func (self Instance) SetGlobalEndPosition(position gd.Vector2) {
	class(self).SetGlobalEndPosition(position)
}

/*
Returns the [member end_position] that is relative to the link as a global position.
*/
func (self Instance) GetGlobalEndPosition() gd.Vector2 {
	return gd.Vector2(class(self).GetGlobalEndPosition())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.NavigationLink2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationLink2D"))
	return Instance{classdb.NavigationLink2D(object)}
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

func (self Instance) StartPosition() gd.Vector2 {
	return gd.Vector2(class(self).GetStartPosition())
}

func (self Instance) SetStartPosition(value gd.Vector2) {
	class(self).SetStartPosition(value)
}

func (self Instance) EndPosition() gd.Vector2 {
	return gd.Vector2(class(self).GetEndPosition())
}

func (self Instance) SetEndPosition(value gd.Vector2) {
	class(self).SetEndPosition(value)
}

func (self Instance) EnterCost() float64 {
	return float64(float64(class(self).GetEnterCost()))
}

func (self Instance) SetEnterCost(value float64) {
	class(self).SetEnterCost(gd.Float(value))
}

func (self Instance) TravelCost() float64 {
	return float64(float64(class(self).GetTravelCost()))
}

func (self Instance) SetTravelCost(value float64) {
	class(self).SetTravelCost(gd.Float(value))
}

/*
Returns the [RID] of this link on the [NavigationServer2D].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBidirectional(bidirectional bool) {
	var frame = callframe.New()
	callframe.Arg(frame, bidirectional)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_bidirectional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsBidirectional() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_is_bidirectional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetNavigationLayerValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStartPosition(position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStartPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEndPosition(position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEndPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [member start_position] that is relative to the link from a global [param position].
*/
//go:nosplit
func (self class) SetGlobalStartPosition(position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_global_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [member start_position] that is relative to the link as a global position.
*/
//go:nosplit
func (self class) GetGlobalStartPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_global_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [member end_position] that is relative to the link from a global [param position].
*/
//go:nosplit
func (self class) SetGlobalEndPosition(position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_global_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [member end_position] that is relative to the link as a global position.
*/
//go:nosplit
func (self class) GetGlobalEndPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_global_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnterCost(enter_cost gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, enter_cost)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnterCost() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTravelCost(travel_cost gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, travel_cost)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTravelCost() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNavigationLink2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNavigationLink2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced       { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance    { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("NavigationLink2D", func(ptr gd.Object) any { return classdb.NavigationLink2D(ptr) })
}
