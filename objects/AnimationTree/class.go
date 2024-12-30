package AnimationTree

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/AnimationMixer"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Path"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A node used for advanced animation transitions in an [AnimationPlayer].
[b]Note:[/b] When linked with an [AnimationPlayer], several properties and methods of the corresponding [AnimationPlayer] will not function as expected. Playback and transitions should be handled using only the [AnimationTree] and its constituent [AnimationNode](s). The [AnimationPlayer] node should be used solely for adding, deleting, and editing animations.
*/
type Instance [1]classdb.AnimationTree
type Any interface {
	gd.IsClass
	AsAnimationTree() Instance
}

/*
Sets the process notification in which to update animations.
*/
func (self Instance) SetProcessCallback(mode classdb.AnimationTreeAnimationProcessCallback) {
	class(self).SetProcessCallback(mode)
}

/*
Returns the process notification in which to update animations.
*/
func (self Instance) GetProcessCallback() classdb.AnimationTreeAnimationProcessCallback {
	return classdb.AnimationTreeAnimationProcessCallback(class(self).GetProcessCallback())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationTree

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationTree"))
	return Instance{classdb.AnimationTree(object)}
}

func (self Instance) TreeRoot() objects.AnimationRootNode {
	return objects.AnimationRootNode(class(self).GetTreeRoot())
}

func (self Instance) SetTreeRoot(value objects.AnimationRootNode) {
	class(self).SetTreeRoot(value)
}

func (self Instance) AdvanceExpressionBaseNode() Path.String {
	return Path.String(class(self).GetAdvanceExpressionBaseNode().String())
}

func (self Instance) SetAdvanceExpressionBaseNode(value Path.String) {
	class(self).SetAdvanceExpressionBaseNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) AnimPlayer() Path.String {
	return Path.String(class(self).GetAnimationPlayer().String())
}

func (self Instance) SetAnimPlayer(value Path.String) {
	class(self).SetAnimationPlayer(gd.NewString(string(value)).NodePath())
}

//go:nosplit
func (self class) SetTreeRoot(animation_node objects.AnimationRootNode) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation_node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_set_tree_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTreeRoot() objects.AnimationRootNode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_get_tree_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.AnimationRootNode{classdb.AnimationRootNode(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdvanceExpressionBaseNode(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_set_advance_expression_base_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdvanceExpressionBaseNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_get_advance_expression_base_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnimationPlayer(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_set_animation_player, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnimationPlayer() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_get_animation_player, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the process notification in which to update animations.
*/
//go:nosplit
func (self class) SetProcessCallback(mode classdb.AnimationTreeAnimationProcessCallback) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_set_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the process notification in which to update animations.
*/
//go:nosplit
func (self class) GetProcessCallback() classdb.AnimationTreeAnimationProcessCallback {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationTreeAnimationProcessCallback](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationTree.Bind_get_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnAnimationPlayerChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("animation_player_changed"), gd.NewCallable(cb), 0)
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
		return gd.VirtualByName(self.AsAnimationMixer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationMixer(), name)
	}
}
func init() {
	classdb.Register("AnimationTree", func(ptr gd.Object) any { return classdb.AnimationTree(ptr) })
}

type AnimationProcessCallback = classdb.AnimationTreeAnimationProcessCallback

const (
	AnimationProcessPhysics AnimationProcessCallback = 0
	AnimationProcessIdle    AnimationProcessCallback = 1
	AnimationProcessManual  AnimationProcessCallback = 2
)
