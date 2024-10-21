package EditorUndoRedoManager

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[EditorUndoRedoManager] is a manager for [UndoRedo] objects associated with edited scenes. Each scene has its own undo history and [EditorUndoRedoManager] ensures that each action performed in the editor gets associated with a proper scene. For actions not related to scenes ([ProjectSettings] edits, external resources, etc.), a separate global history is used.
The usage is mostly the same as [UndoRedo]. You create and commit actions and the manager automatically decides under-the-hood what scenes it belongs to. The scene is deduced based on the first operation in an action, using the object from the operation. The rules are as follows:
- If the object is a [Node], use the currently edited scene;
- If the object is a built-in resource, use the scene from its path;
- If the object is external resource or anything else, use global history.
This guessing can sometimes yield false results, so you can provide a custom context object when creating an action.
[EditorUndoRedoManager] is intended to be used by Godot editor plugins. You can obtain it using [method EditorPlugin.get_undo_redo]. For non-editor uses or plugins that don't need to integrate with the editor's undo history, use [UndoRedo] instead.
The manager's API is mostly the same as in [UndoRedo], so you can refer to its documentation for more examples. The main difference is that [EditorUndoRedoManager] uses object + method name for actions, instead of [Callable].

*/
type Simple [1]classdb.EditorUndoRedoManager
func (self Simple) CreateAction(name string, merge_mode classdb.UndoRedoMergeMode, custom_context gd.Object, backward_undo_ops bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateAction(gc.String(name), merge_mode, custom_context, backward_undo_ops)
}
func (self Simple) CommitAction(execute bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CommitAction(execute)
}
func (self Simple) IsCommittingAction() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCommittingAction())
}
func (self Simple) ForceFixedHistory() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceFixedHistory()
}
func (self Simple) AddDoProperty(obj gd.Object, property string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddDoProperty(obj, gc.StringName(property), value)
}
func (self Simple) AddUndoProperty(obj gd.Object, property string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddUndoProperty(obj, gc.StringName(property), value)
}
func (self Simple) AddDoReference(obj gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddDoReference(obj)
}
func (self Simple) AddUndoReference(obj gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddUndoReference(obj)
}
func (self Simple) GetObjectHistoryId(obj gd.Object) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetObjectHistoryId(obj)))
}
func (self Simple) GetHistoryUndoRedo(id int) [1]classdb.UndoRedo {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.UndoRedo(Expert(self).GetHistoryUndoRedo(gc, gd.Int(id)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorUndoRedoManager
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Create a new action. After this is called, do all your calls to [method add_do_method], [method add_undo_method], [method add_do_property], and [method add_undo_property], then commit the action with [method commit_action].
The way actions are merged is dictated by the [param merge_mode] argument. See [enum UndoRedo.MergeMode] for details.
If [param custom_context] object is provided, it will be used for deducing target history (instead of using the first operation).
The way undo operation are ordered in actions is dictated by [param backward_undo_ops]. When [param backward_undo_ops] is [code]false[/code] undo option are ordered in the same order they were added. Which means the first operation to be added will be the first to be undone.
*/
//go:nosplit
func (self class) CreateAction(name gd.String, merge_mode classdb.UndoRedoMergeMode, custom_context gd.Object, backward_undo_ops bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, merge_mode)
	callframe.Arg(frame, mmm.End(custom_context.AsPointer())[0])
	callframe.Arg(frame, backward_undo_ops)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_create_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Commit the action. If [param execute] is true (default), all "do" methods/properties are called/set when this function is called.
*/
//go:nosplit
func (self class) CommitAction(execute bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, execute)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_commit_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the [EditorUndoRedoManager] is currently committing the action, i.e. running its "do" method or property change (see [method commit_action]).
*/
//go:nosplit
func (self class) IsCommittingAction() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_is_committing_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Forces the next operation (e.g. [method add_do_method]) to use the action's history rather than guessing it from the object. This is sometimes needed when a history can't be correctly determined, like for a nested resource that doesn't have a path yet.
This method should only be used when absolutely necessary, otherwise it might cause invalid history state. For most of complex cases, the [code]custom_context[/code] parameter of [method create_action] is sufficient.
*/
//go:nosplit
func (self class) ForceFixedHistory()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_force_fixed_history, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Register a property value change for "do".
If this is the first operation, the [param object] will be used to deduce target undo history.
*/
//go:nosplit
func (self class) AddDoProperty(obj gd.Object, property gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_add_do_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Register a property value change for "undo".
If this is the first operation, the [param object] will be used to deduce target undo history.
*/
//go:nosplit
func (self class) AddUndoProperty(obj gd.Object, property gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_add_undo_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Register a reference for "do" that will be erased if the "do" history is lost. This is useful mostly for new nodes created for the "do" call. Do not use for resources.
*/
//go:nosplit
func (self class) AddDoReference(obj gd.Object)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_add_do_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Register a reference for "undo" that will be erased if the "undo" history is lost. This is useful mostly for nodes removed with the "do" call (not the "undo" call!).
*/
//go:nosplit
func (self class) AddUndoReference(obj gd.Object)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_add_undo_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the history ID deduced from the given [param object]. It can be used with [method get_history_undo_redo].
*/
//go:nosplit
func (self class) GetObjectHistoryId(obj gd.Object) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_get_object_history_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [UndoRedo] object associated with the given history [param id].
[param id] above [code]0[/code] are mapped to the opened scene tabs (but it doesn't match their order). [param id] of [code]0[/code] or lower have special meaning (see [enum SpecialHistory]).
Best used with [method get_object_history_id]. This method is only provided in case you need some more advanced methods of [UndoRedo] (but keep in mind that directly operating on the [UndoRedo] object might affect editor's stability).
*/
//go:nosplit
func (self class) GetHistoryUndoRedo(ctx gd.Lifetime, id gd.Int) object.UndoRedo {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorUndoRedoManager.Bind_get_history_undo_redo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.UndoRedo
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorUndoRedoManager() Expert { return self[0].AsEditorUndoRedoManager() }


//go:nosplit
func (self Simple) AsEditorUndoRedoManager() Simple { return self[0].AsEditorUndoRedoManager() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorUndoRedoManager", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type SpecialHistory = classdb.EditorUndoRedoManagerSpecialHistory

const (
/*Global history not associated with any scene, but with external resources etc.*/
	GlobalHistory SpecialHistory = 0
/*History associated with remote inspector. Used when live editing a running project.*/
	RemoteHistory SpecialHistory = -9
/*Invalid "null" history. It's a special value, not associated with any object.*/
	InvalidHistory SpecialHistory = -99
)
