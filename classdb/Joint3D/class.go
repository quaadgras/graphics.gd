// Package Joint3D provides methods for working with Joint3D object instances.
package Joint3D

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
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/RID"

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
Abstract base class for all joints in 3D physics. 3D joints bind together two physics bodies ([member node_a] and [member node_b]) and apply a constraint. If only one body is defined, it is attached to a fixed [StaticBody3D] without collision shapes.
*/
type Instance [1]gdclass.Joint3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsJoint3D() Instance
}

/*
Returns the joint's internal [RID] from the [PhysicsServer3D].
*/
func (self Instance) GetRid() RID.Joint3D { //gd:Joint3D.get_rid
	return RID.Joint3D(class(self).GetRid())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Joint3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Joint3D"))
	casted := Instance{*(*gdclass.Joint3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) NodeA() NodePath.String {
	return NodePath.String(class(self).GetNodeA().String())
}

func (self Instance) SetNodeA(value NodePath.String) {
	class(self).SetNodeA(gd.NewString(string(value)).NodePath())
}

func (self Instance) NodeB() NodePath.String {
	return NodePath.String(class(self).GetNodeB().String())
}

func (self Instance) SetNodeB(value NodePath.String) {
	class(self).SetNodeB(gd.NewString(string(value)).NodePath())
}

func (self Instance) SolverPriority() int {
	return int(int(class(self).GetSolverPriority()))
}

func (self Instance) SetSolverPriority(value int) {
	class(self).SetSolverPriority(gd.Int(value))
}

func (self Instance) ExcludeNodesFromCollision() bool {
	return bool(class(self).GetExcludeNodesFromCollision())
}

func (self Instance) SetExcludeNodesFromCollision(value bool) {
	class(self).SetExcludeNodesFromCollision(value)
}

//go:nosplit
func (self class) SetNodeA(node gd.NodePath) { //gd:Joint3D.set_node_a
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_set_node_a, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNodeA() gd.NodePath { //gd:Joint3D.get_node_a
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_node_a, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNodeB(node gd.NodePath) { //gd:Joint3D.set_node_b
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_set_node_b, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNodeB() gd.NodePath { //gd:Joint3D.get_node_b
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_node_b, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSolverPriority(priority gd.Int) { //gd:Joint3D.set_solver_priority
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_set_solver_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSolverPriority() gd.Int { //gd:Joint3D.get_solver_priority
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_solver_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeNodesFromCollision(enable bool) { //gd:Joint3D.set_exclude_nodes_from_collision
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_set_exclude_nodes_from_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeNodesFromCollision() bool { //gd:Joint3D.get_exclude_nodes_from_collision
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_exclude_nodes_from_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the joint's internal [RID] from the [PhysicsServer3D].
*/
//go:nosplit
func (self class) GetRid() gd.RID { //gd:Joint3D.get_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsJoint3D() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJoint3D() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

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
	gdclass.Register("Joint3D", func(ptr gd.Object) any { return [1]gdclass.Joint3D{*(*gdclass.Joint3D)(unsafe.Pointer(&ptr))} })
}
