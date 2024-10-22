package AnimationTree

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AnimationMixer"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A node used for advanced animation transitions in an [AnimationPlayer].
[b]Note:[/b] When linked with an [AnimationPlayer], several properties and methods of the corresponding [AnimationPlayer] will not function as expected. Playback and transitions should be handled using only the [AnimationTree] and its constituent [AnimationNode](s). The [AnimationPlayer] node should be used solely for adding, deleting, and editing animations.

*/
type Go [1]classdb.AnimationTree

/*
Sets the process notification in which to update animations.
*/
func (self Go) SetProcessCallback(mode classdb.AnimationTreeAnimationProcessCallback) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProcessCallback(mode)
}

/*
Returns the process notification in which to update animations.
*/
func (self Go) GetProcessCallback() classdb.AnimationTreeAnimationProcessCallback {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationTreeAnimationProcessCallback(class(self).GetProcessCallback())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationTree
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimationTree"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) TreeRoot() gdclass.AnimationRootNode {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AnimationRootNode(class(self).GetTreeRoot(gc))
}

func (self Go) SetTreeRoot(value gdclass.AnimationRootNode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTreeRoot(value)
}

func (self Go) AdvanceExpressionBaseNode() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetAdvanceExpressionBaseNode(gc).String())
}

func (self Go) SetAdvanceExpressionBaseNode(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdvanceExpressionBaseNode(gc.String(value).NodePath(gc))
}

func (self Go) AnimPlayer() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetAnimationPlayer(gc).String())
}

func (self Go) SetAnimPlayer(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAnimationPlayer(gc.String(value).NodePath(gc))
}

//go:nosplit
func (self class) SetTreeRoot(animation_node gdclass.AnimationRootNode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation_node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_set_tree_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTreeRoot(ctx gd.Lifetime) gdclass.AnimationRootNode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_get_tree_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AnimationRootNode
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdvanceExpressionBaseNode(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_set_advance_expression_base_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdvanceExpressionBaseNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_get_advance_expression_base_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAnimationPlayer(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_set_animation_player, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAnimationPlayer(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_get_animation_player, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the process notification in which to update animations.
*/
//go:nosplit
func (self class) SetProcessCallback(mode classdb.AnimationTreeAnimationProcessCallback)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_set_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the process notification in which to update animations.
*/
//go:nosplit
func (self class) GetProcessCallback() classdb.AnimationTreeAnimationProcessCallback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationTreeAnimationProcessCallback](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_get_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnAnimationPlayerChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("animation_player_changed"), gc.Callable(cb), 0)
}


func (self class) AsAnimationTree() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationTree() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationMixer() AnimationMixer.GD { return *((*AnimationMixer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationMixer() AnimationMixer.Go { return *((*AnimationMixer.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationMixer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationMixer(), name)
	}
}
func init() {classdb.Register("AnimationTree", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type AnimationProcessCallback = classdb.AnimationTreeAnimationProcessCallback

const (
	AnimationProcessPhysics AnimationProcessCallback = 0
	AnimationProcessIdle AnimationProcessCallback = 1
	AnimationProcessManual AnimationProcessCallback = 2
)
