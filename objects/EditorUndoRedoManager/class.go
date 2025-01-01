package EditorUndoRedoManager

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.EditorUndoRedoManager
type Any interface {
	gd.IsClass
	AsEditorUndoRedoManager() Instance
}

/*
Create a new action. After this is called, do all your calls to [method add_do_method], [method add_undo_method], [method add_do_property], and [method add_undo_property], then commit the action with [method commit_action].
The way actions are merged is dictated by the [param merge_mode] argument. See [enum UndoRedo.MergeMode] for details.
If [param custom_context] object is provided, it will be used for deducing target history (instead of using the first operation).
The way undo operation are ordered in actions is dictated by [param backward_undo_ops]. When [param backward_undo_ops] is [code]false[/code] undo option are ordered in the same order they were added. Which means the first operation to be added will be the first to be undone.
*/
func (self Instance) CreateAction(name string) {
	class(self).CreateAction(gd.NewString(name), 0, [1]gd.Object{}[0], false)
}

/*
Commit the action. If [param execute] is true (default), all "do" methods/properties are called/set when this function is called.
*/
func (self Instance) CommitAction() {
	class(self).CommitAction(true)
}

/*
Returns [code]true[/code] if the [EditorUndoRedoManager] is currently committing the action, i.e. running its "do" method or property change (see [method commit_action]).
*/
func (self Instance) IsCommittingAction() bool {
	return bool(class(self).IsCommittingAction())
}

/*
Forces the next operation (e.g. [method add_do_method]) to use the action's history rather than guessing it from the object. This is sometimes needed when a history can't be correctly determined, like for a nested resource that doesn't have a path yet.
This method should only be used when absolutely necessary, otherwise it might cause invalid history state. For most of complex cases, the [code]custom_context[/code] parameter of [method create_action] is sufficient.
*/
func (self Instance) ForceFixedHistory() {
	class(self).ForceFixedHistory()
}

/*
Register a property value change for "do".
If this is the first operation, the [param object] will be used to deduce target undo history.
*/
func (self Instance) AddDoProperty(obj gd.Object, property string, value any) {
	class(self).AddDoProperty(obj, gd.NewStringName(property), gd.NewVariant(value))
}

/*
Register a property value change for "undo".
If this is the first operation, the [param object] will be used to deduce target undo history.
*/
func (self Instance) AddUndoProperty(obj gd.Object, property string, value any) {
	class(self).AddUndoProperty(obj, gd.NewStringName(property), gd.NewVariant(value))
}

/*
Register a reference for "do" that will be erased if the "do" history is lost. This is useful mostly for new nodes created for the "do" call. Do not use for resources.
*/
func (self Instance) AddDoReference(obj gd.Object) {
	class(self).AddDoReference(obj)
}

/*
Register a reference for "undo" that will be erased if the "undo" history is lost. This is useful mostly for nodes removed with the "do" call (not the "undo" call!).
*/
func (self Instance) AddUndoReference(obj gd.Object) {
	class(self).AddUndoReference(obj)
}

/*
Returns the history ID deduced from the given [param object]. It can be used with [method get_history_undo_redo].
*/
func (self Instance) GetObjectHistoryId(obj gd.Object) int {
	return int(int(class(self).GetObjectHistoryId(obj)))
}

/*
Returns the [UndoRedo] object associated with the given history [param id].
[param id] above [code]0[/code] are mapped to the opened scene tabs (but it doesn't match their order). [param id] of [code]0[/code] or lower have special meaning (see [enum SpecialHistory]).
Best used with [method get_object_history_id]. This method is only provided in case you need some more advanced methods of [UndoRedo] (but keep in mind that directly operating on the [UndoRedo] object might affect editor's stability).
*/
func (self Instance) GetHistoryUndoRedo(id int) objects.UndoRedo {
	return objects.UndoRedo(class(self).GetHistoryUndoRedo(gd.Int(id)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorUndoRedoManager

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorUndoRedoManager"))
	return Instance{classdb.EditorUndoRedoManager(object)}
}

/*
Create a new action. After this is called, do all your calls to [method add_do_method], [method add_undo_method], [method add_do_property], and [method add_undo_property], then commit the action with [method commit_action].
The way actions are merged is dictated by the [param merge_mode] argument. See [enum UndoRedo.MergeMode] for details.
If [param custom_context] object is provided, it will be used for deducing target history (instead of using the first operation).
The way undo operation are ordered in actions is dictated by [param backward_undo_ops]. When [param backward_undo_ops] is [code]false[/code] undo option are ordered in the same order they were added. Which means the first operation to be added will be the first to be undone.
*/
//go:nosplit
func (self class) CreateAction(name gd.String, merge_mode classdb.UndoRedoMergeMode, custom_context gd.Object, backward_undo_ops bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, merge_mode)
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(custom_context))
	callframe.Arg(frame, backward_undo_ops)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_create_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Commit the action. If [param execute] is true (default), all "do" methods/properties are called/set when this function is called.
*/
//go:nosplit
func (self class) CommitAction(execute bool) {
	var frame = callframe.New()
	callframe.Arg(frame, execute)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_commit_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the [EditorUndoRedoManager] is currently committing the action, i.e. running its "do" method or property change (see [method commit_action]).
*/
//go:nosplit
func (self class) IsCommittingAction() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_is_committing_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Forces the next operation (e.g. [method add_do_method]) to use the action's history rather than guessing it from the object. This is sometimes needed when a history can't be correctly determined, like for a nested resource that doesn't have a path yet.
This method should only be used when absolutely necessary, otherwise it might cause invalid history state. For most of complex cases, the [code]custom_context[/code] parameter of [method create_action] is sufficient.
*/
//go:nosplit
func (self class) ForceFixedHistory() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_force_fixed_history, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Register a property value change for "do".
If this is the first operation, the [param object] will be used to deduce target undo history.
*/
//go:nosplit
func (self class) AddDoProperty(obj gd.Object, property gd.StringName, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, pointers.Get(property))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_add_do_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Register a property value change for "undo".
If this is the first operation, the [param object] will be used to deduce target undo history.
*/
//go:nosplit
func (self class) AddUndoProperty(obj gd.Object, property gd.StringName, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, pointers.Get(property))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_add_undo_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Register a reference for "do" that will be erased if the "do" history is lost. This is useful mostly for new nodes created for the "do" call. Do not use for resources.
*/
//go:nosplit
func (self class) AddDoReference(obj gd.Object) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_add_do_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Register a reference for "undo" that will be erased if the "undo" history is lost. This is useful mostly for nodes removed with the "do" call (not the "undo" call!).
*/
//go:nosplit
func (self class) AddUndoReference(obj gd.Object) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_add_undo_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the history ID deduced from the given [param object]. It can be used with [method get_history_undo_redo].
*/
//go:nosplit
func (self class) GetObjectHistoryId(obj gd.Object) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_get_object_history_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetHistoryUndoRedo(id gd.Int) objects.UndoRedo {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorUndoRedoManager.Bind_get_history_undo_redo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.UndoRedo{classdb.UndoRedo(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
func (self Instance) OnHistoryChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("history_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnVersionChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("version_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorUndoRedoManager() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorUndoRedoManager() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("EditorUndoRedoManager", func(ptr gd.Object) any { return [1]classdb.EditorUndoRedoManager{classdb.EditorUndoRedoManager(ptr)} })
}

type SpecialHistory = classdb.EditorUndoRedoManagerSpecialHistory

const (
	/*Global history not associated with any scene, but with external resources etc.*/
	GlobalHistory SpecialHistory = 0
	/*History associated with remote inspector. Used when live editing a running project.*/
	RemoteHistory SpecialHistory = -9
	/*Invalid "null" history. It's a special value, not associated with any object.*/
	InvalidHistory SpecialHistory = -99
)
