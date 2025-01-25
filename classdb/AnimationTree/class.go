// Package AnimationTree provides methods for working with AnimationTree object instances.
package AnimationTree

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
import "graphics.gd/classdb/AnimationMixer"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/NodePath"

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
A node used for advanced animation transitions in an [AnimationPlayer].
[b]Note:[/b] When linked with an [AnimationPlayer], several properties and methods of the corresponding [AnimationPlayer] will not function as expected. Playback and transitions should be handled using only the [AnimationTree] and its constituent [AnimationNode](s). The [AnimationPlayer] node should be used solely for adding, deleting, and editing animations.
*/
type Instance [1]gdclass.AnimationTree

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationTree() Instance
}

/*
Sets the process notification in which to update animations.
*/
func (self Instance) SetProcessCallback(mode gdclass.AnimationTreeAnimationProcessCallback) {
	class(self).SetProcessCallback(mode)
}

/*
Returns the process notification in which to update animations.
*/
func (self Instance) GetProcessCallback() gdclass.AnimationTreeAnimationProcessCallback {
	return gdclass.AnimationTreeAnimationProcessCallback(class(self).GetProcessCallback())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationTree

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationTree"))
	casted := Instance{*(*gdclass.AnimationTree)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) TreeRoot() [1]gdclass.AnimationRootNode {
	return [1]gdclass.AnimationRootNode(class(self).GetTreeRoot())
}

func (self Instance) SetTreeRoot(value [1]gdclass.AnimationRootNode) {
	class(self).SetTreeRoot(value)
}

func (self Instance) AdvanceExpressionBaseNode() NodePath.String {
	return NodePath.String(class(self).GetAdvanceExpressionBaseNode().String())
}

func (self Instance) SetAdvanceExpressionBaseNode(value NodePath.String) {
	class(self).SetAdvanceExpressionBaseNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) AnimPlayer() NodePath.String {
	return NodePath.String(class(self).GetAnimationPlayer().String())
}

func (self Instance) SetAnimPlayer(value NodePath.String) {
	class(self).SetAnimationPlayer(gd.NewString(string(value)).NodePath())
}

//go:nosplit
func (self class) SetTreeRoot(animation_node [1]gdclass.AnimationRootNode) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation_node[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_set_tree_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTreeRoot() [1]gdclass.AnimationRootNode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_get_tree_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AnimationRootNode{gd.PointerWithOwnershipTransferredToGo[gdclass.AnimationRootNode](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdvanceExpressionBaseNode(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_set_advance_expression_base_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdvanceExpressionBaseNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_get_advance_expression_base_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnimationPlayer(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_set_animation_player, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnimationPlayer() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_get_animation_player, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the process notification in which to update animations.
*/
//go:nosplit
func (self class) SetProcessCallback(mode gdclass.AnimationTreeAnimationProcessCallback) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_set_process_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the process notification in which to update animations.
*/
//go:nosplit
func (self class) GetProcessCallback() gdclass.AnimationTreeAnimationProcessCallback {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationTreeAnimationProcessCallback](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_get_process_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnAnimationPlayerChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_player_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationTree() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationTree() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAnimationMixer() AnimationMixer.Advanced {
	return *((*AnimationMixer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationMixer() AnimationMixer.Instance {
	return *((*AnimationMixer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationMixer.Advanced(self.AsAnimationMixer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationMixer.Instance(self.AsAnimationMixer()), name)
	}
}
func init() {
	gdclass.Register("AnimationTree", func(ptr gd.Object) any {
		return [1]gdclass.AnimationTree{*(*gdclass.AnimationTree)(unsafe.Pointer(&ptr))}
	})
}

type AnimationProcessCallback = gdclass.AnimationTreeAnimationProcessCallback

const (
	AnimationProcessPhysics AnimationProcessCallback = 0
	AnimationProcessIdle    AnimationProcessCallback = 1
	AnimationProcessManual  AnimationProcessCallback = 2
)
