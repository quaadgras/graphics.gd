package Joint3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Abstract base class for all joints in 3D physics. 3D joints bind together two physics bodies ([member node_a] and [member node_b]) and apply a constraint. If only one body is defined, it is attached to a fixed [StaticBody3D] without collision shapes.
*/
type Instance [1]classdb.Joint3D

/*
Returns the joint's internal [RID] from the [PhysicsServer3D].
*/
func (self Instance) GetRid() gd.RID {
	return gd.RID(class(self).GetRid())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Joint3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Joint3D"))
	return Instance{classdb.Joint3D(object)}
}

func (self Instance) NodeA() string {
	return string(class(self).GetNodeA().String())
}

func (self Instance) SetNodeA(value string) {
	class(self).SetNodeA(gd.NewString(value).NodePath())
}

func (self Instance) NodeB() string {
	return string(class(self).GetNodeB().String())
}

func (self Instance) SetNodeB(value string) {
	class(self).SetNodeB(gd.NewString(value).NodePath())
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
func (self class) SetNodeA(node gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_set_node_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNodeA() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_node_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNodeB(node gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_set_node_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNodeB() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_node_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSolverPriority(priority gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_set_solver_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSolverPriority() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_solver_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeNodesFromCollision(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_set_exclude_nodes_from_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeNodesFromCollision() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_exclude_nodes_from_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the joint's internal [RID] from the [PhysicsServer3D].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint3D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() { classdb.Register("Joint3D", func(ptr gd.Object) any { return classdb.Joint3D(ptr) }) }
