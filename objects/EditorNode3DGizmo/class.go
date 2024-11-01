package EditorNode3DGizmo

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node3DGizmo"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Gizmo that is used for providing custom visualization and editing (handles and subgizmos) for [Node3D] objects. Can be overridden to create custom gizmos, but for simple gizmos creating a [EditorNode3DGizmoPlugin] is usually recommended.

	// EditorNode3DGizmo methods that can be overridden by a [Class] that extends it.
	type EditorNode3DGizmo interface {
		//Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method clear] at the beginning of this method and then add visual elements depending on the node's properties.
		Redraw()
		//Override this method to return the name of an edited handle (handles must have been previously added by [method add_handles]). Handles can be named for reference to the user when editing.
		//The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
		GetHandleName(id int, secondary bool) string
		//Override this method to return [code]true[/code] whenever the given handle should be highlighted in the editor.
		//The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
		IsHandleHighlighted(id int, secondary bool) bool
		//Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
		//The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
		GetHandleValue(id int, secondary bool) gd.Variant
		BeginHandleAction(id int, secondary bool)
		//Override this method to update the node properties when the user drags a gizmo handle (previously added with [method add_handles]). The provided [param point] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
		//The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method add_handles] for more information).
		SetHandle(id int, secondary bool, camera objects.Camera3D, point gd.Vector2)
		//Override this method to commit a handle being edited (handles must have been previously added by [method add_handles]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
		//If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
		//The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method add_handles] for more information).
		CommitHandle(id int, secondary bool, restore gd.Variant, cancel bool)
		//Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param point] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
		SubgizmosIntersectRay(camera objects.Camera3D, point gd.Vector2) int
		//Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and a [param frustum], this method should return which subgizmos are contained within the frustum. The [param frustum] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
		SubgizmosIntersectFrustum(camera objects.Camera3D, frustum gd.Array) []int32
		//Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the [Node3D]'s local coordinate system.
		SetSubgizmoTransform(id int, transform gd.Transform3D)
		//Override this method to return the current transform of a subgizmo. This transform will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_subgizmos].
		GetSubgizmoTransform(id int) gd.Transform3D
		//Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
		//If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action.
		CommitSubgizmos(ids []int32, restores gd.Array, cancel bool)
	}
*/
type Instance [1]classdb.EditorNode3DGizmo

/*
Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method clear] at the beginning of this method and then add visual elements depending on the node's properties.
*/
func (Instance) _redraw(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to return the name of an edited handle (handles must have been previously added by [method add_handles]). Handles can be named for reference to the user when editing.
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
*/
func (Instance) _get_handle_name(impl func(ptr unsafe.Pointer, id int, secondary bool) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id), secondary)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to return [code]true[/code] whenever the given handle should be highlighted in the editor.
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
*/
func (Instance) _is_handle_highlighted(impl func(ptr unsafe.Pointer, id int, secondary bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id), secondary)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
*/
func (Instance) _get_handle_value(impl func(ptr unsafe.Pointer, id int, secondary bool) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id), secondary)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _begin_handle_action(impl func(ptr unsafe.Pointer, id int, secondary bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(id), secondary)
	}
}

/*
Override this method to update the node properties when the user drags a gizmo handle (previously added with [method add_handles]). The provided [param point] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method add_handles] for more information).
*/
func (Instance) _set_handle(impl func(ptr unsafe.Pointer, id int, secondary bool, camera objects.Camera3D, point gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		var camera = objects.Camera3D{pointers.New[classdb.Camera3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
		defer pointers.End(camera[0])
		var point = gd.UnsafeGet[gd.Vector2](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(id), secondary, camera, point)
	}
}

/*
Override this method to commit a handle being edited (handles must have been previously added by [method add_handles]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method add_handles] for more information).
*/
func (Instance) _commit_handle(impl func(ptr unsafe.Pointer, id int, secondary bool, restore gd.Variant, cancel bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		var restore = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		defer pointers.End(restore)
		var cancel = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(id), secondary, restore, cancel)
	}
}

/*
Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param point] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
*/
func (Instance) _subgizmos_intersect_ray(impl func(ptr unsafe.Pointer, camera objects.Camera3D, point gd.Vector2) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var camera = objects.Camera3D{pointers.New[classdb.Camera3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(camera[0])
		var point = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, camera, point)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and a [param frustum], this method should return which subgizmos are contained within the frustum. The [param frustum] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
*/
func (Instance) _subgizmos_intersect_frustum(impl func(ptr unsafe.Pointer, camera objects.Camera3D, frustum gd.Array) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var camera = objects.Camera3D{pointers.New[classdb.Camera3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(camera[0])
		var frustum = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(frustum)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, camera, frustum)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the [Node3D]'s local coordinate system.
*/
func (Instance) _set_subgizmo_transform(impl func(ptr unsafe.Pointer, id int, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(id), transform)
	}
}

/*
Override this method to return the current transform of a subgizmo. This transform will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_subgizmos].
*/
func (Instance) _get_subgizmo_transform(impl func(ptr unsafe.Pointer, id int) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action.
*/
func (Instance) _commit_subgizmos(impl func(ptr unsafe.Pointer, ids []int32, restores gd.Array, cancel bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var ids = pointers.New[gd.PackedInt32Array](gd.UnsafeGet[[2]uintptr](p_args, 0))
		defer pointers.End(ids)
		var restores = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(restores)
		var cancel = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, ids.AsSlice(), restores, cancel)
	}
}

/*
Adds lines to the gizmo (as sets of 2 points), with a given material. The lines are used for visualizing the gizmo. Call this method during [method _redraw].
*/
func (self Instance) AddLines(lines []gd.Vector3, material objects.Material) {
	class(self).AddLines(gd.NewPackedVector3Slice(lines), material, false, gd.Color{1, 1, 1, 1})
}

/*
Adds a mesh to the gizmo with the specified [param material], local [param transform] and [param skeleton]. Call this method during [method _redraw].
*/
func (self Instance) AddMesh(mesh objects.Mesh) {
	class(self).AddMesh(mesh, ([1]objects.Material{}[0]), gd.NewTransform3D(1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0), ([1]objects.SkinReference{}[0]))
}

/*
Adds the specified [param segments] to the gizmo's collision shape for picking. Call this method during [method _redraw].
*/
func (self Instance) AddCollisionSegments(segments []gd.Vector3) {
	class(self).AddCollisionSegments(gd.NewPackedVector3Slice(segments))
}

/*
Adds collision triangles to the gizmo for picking. A [TriangleMesh] can be generated from a regular [Mesh] too. Call this method during [method _redraw].
*/
func (self Instance) AddCollisionTriangles(triangles objects.TriangleMesh) {
	class(self).AddCollisionTriangles(triangles)
}

/*
Adds an unscaled billboard for visualization and selection. Call this method during [method _redraw].
*/
func (self Instance) AddUnscaledBillboard(material objects.Material) {
	class(self).AddUnscaledBillboard(material, gd.Float(1), gd.Color{1, 1, 1, 1})
}

/*
Adds a list of handles (points) which can be used to edit the properties of the gizmo's [Node3D]. The [param ids] argument can be used to specify a custom identifier for each handle, if an empty array is passed, the ids will be assigned automatically from the [param handles] argument order.
The [param secondary] argument marks the added handles as secondary, meaning they will normally have lower selection priority than regular handles. When the user is holding the shift key secondary handles will switch to have higher priority than regular handles. This change in priority can be used to place multiple handles at the same point while still giving the user control on their selection.
There are virtual methods which will be called upon editing of these handles. Call this method during [method _redraw].
*/
func (self Instance) AddHandles(handles []gd.Vector3, material objects.Material, ids []int32) {
	class(self).AddHandles(gd.NewPackedVector3Slice(handles), material, gd.NewPackedInt32Slice(ids), false, false)
}

/*
Sets the reference [Node3D] node for the gizmo. [param node] must inherit from [Node3D].
*/
func (self Instance) SetNode3d(node objects.Node) {
	class(self).SetNode3d(node)
}

/*
Returns the [Node3D] node associated with this gizmo.
*/
func (self Instance) GetNode3d() objects.Node3D {
	return objects.Node3D(class(self).GetNode3d())
}

/*
Returns the [EditorNode3DGizmoPlugin] that owns this gizmo. It's useful to retrieve materials using [method EditorNode3DGizmoPlugin.get_material].
*/
func (self Instance) GetPlugin() objects.EditorNode3DGizmoPlugin {
	return objects.EditorNode3DGizmoPlugin(class(self).GetPlugin())
}

/*
Removes everything in the gizmo including meshes, collisions and handles.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Sets the gizmo's hidden state. If [code]true[/code], the gizmo will be hidden. If [code]false[/code], it will be shown.
*/
func (self Instance) SetHidden(hidden bool) {
	class(self).SetHidden(hidden)
}

/*
Returns [code]true[/code] if the given subgizmo is currently selected. Can be used to highlight selected elements during [method _redraw].
*/
func (self Instance) IsSubgizmoSelected(id int) bool {
	return bool(class(self).IsSubgizmoSelected(gd.Int(id)))
}

/*
Returns a list of the currently selected subgizmos. Can be used to highlight selected elements during [method _redraw].
*/
func (self Instance) GetSubgizmoSelection() []int32 {
	return []int32(class(self).GetSubgizmoSelection().AsSlice())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorNode3DGizmo

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorNode3DGizmo"))
	return Instance{classdb.EditorNode3DGizmo(object)}
}

/*
Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method clear] at the beginning of this method and then add visual elements depending on the node's properties.
*/
func (class) _redraw(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to return the name of an edited handle (handles must have been previously added by [method add_handles]). Handles can be named for reference to the user when editing.
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
*/
func (class) _get_handle_name(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id, secondary)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to return [code]true[/code] whenever the given handle should be highlighted in the editor.
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
*/
func (class) _is_handle_highlighted(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id, secondary)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
*/
func (class) _get_handle_value(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id, secondary)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _begin_handle_action(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, id, secondary)
	}
}

/*
Override this method to update the node properties when the user drags a gizmo handle (previously added with [method add_handles]). The provided [param point] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method add_handles] for more information).
*/
func (class) _set_handle(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool, camera objects.Camera3D, point gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		var camera = objects.Camera3D{pointers.New[classdb.Camera3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 2)})}
		defer pointers.End(camera[0])
		var point = gd.UnsafeGet[gd.Vector2](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, id, secondary, camera, point)
	}
}

/*
Override this method to commit a handle being edited (handles must have been previously added by [method add_handles]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method add_handles] for more information).
*/
func (class) _commit_handle(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool, restore gd.Variant, cancel bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var secondary = gd.UnsafeGet[bool](p_args, 1)
		var restore = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		var cancel = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, id, secondary, restore, cancel)
	}
}

/*
Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param point] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
*/
func (class) _subgizmos_intersect_ray(impl func(ptr unsafe.Pointer, camera objects.Camera3D, point gd.Vector2) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var camera = objects.Camera3D{pointers.New[classdb.Camera3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(camera[0])
		var point = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, camera, point)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and a [param frustum], this method should return which subgizmos are contained within the frustum. The [param frustum] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
*/
func (class) _subgizmos_intersect_frustum(impl func(ptr unsafe.Pointer, camera objects.Camera3D, frustum gd.Array) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var camera = objects.Camera3D{pointers.New[classdb.Camera3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(camera[0])
		var frustum = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, camera, frustum)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the [Node3D]'s local coordinate system.
*/
func (class) _set_subgizmo_transform(impl func(ptr unsafe.Pointer, id gd.Int, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, id, transform)
	}
}

/*
Override this method to return the current transform of a subgizmo. This transform will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_subgizmos].
*/
func (class) _get_subgizmo_transform(impl func(ptr unsafe.Pointer, id gd.Int) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action.
*/
func (class) _commit_subgizmos(impl func(ptr unsafe.Pointer, ids gd.PackedInt32Array, restores gd.Array, cancel bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var ids = pointers.New[gd.PackedInt32Array](gd.UnsafeGet[[2]uintptr](p_args, 0))
		var restores = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var cancel = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, ids, restores, cancel)
	}
}

/*
Adds lines to the gizmo (as sets of 2 points), with a given material. The lines are used for visualizing the gizmo. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddLines(lines gd.PackedVector3Array, material objects.Material, billboard bool, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(lines))
	callframe.Arg(frame, pointers.Get(material[0])[0])
	callframe.Arg(frame, billboard)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_add_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a mesh to the gizmo with the specified [param material], local [param transform] and [param skeleton]. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddMesh(mesh objects.Mesh, material objects.Material, transform gd.Transform3D, skeleton objects.SkinReference) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	callframe.Arg(frame, pointers.Get(material[0])[0])
	callframe.Arg(frame, transform)
	callframe.Arg(frame, pointers.Get(skeleton[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_add_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds the specified [param segments] to the gizmo's collision shape for picking. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddCollisionSegments(segments gd.PackedVector3Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(segments))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_add_collision_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds collision triangles to the gizmo for picking. A [TriangleMesh] can be generated from a regular [Mesh] too. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddCollisionTriangles(triangles objects.TriangleMesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(triangles[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_add_collision_triangles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds an unscaled billboard for visualization and selection. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddUnscaledBillboard(material objects.Material, default_scale gd.Float, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	callframe.Arg(frame, default_scale)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_add_unscaled_billboard, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a list of handles (points) which can be used to edit the properties of the gizmo's [Node3D]. The [param ids] argument can be used to specify a custom identifier for each handle, if an empty array is passed, the ids will be assigned automatically from the [param handles] argument order.
The [param secondary] argument marks the added handles as secondary, meaning they will normally have lower selection priority than regular handles. When the user is holding the shift key secondary handles will switch to have higher priority than regular handles. This change in priority can be used to place multiple handles at the same point while still giving the user control on their selection.
There are virtual methods which will be called upon editing of these handles. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddHandles(handles gd.PackedVector3Array, material objects.Material, ids gd.PackedInt32Array, billboard bool, secondary bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(handles))
	callframe.Arg(frame, pointers.Get(material[0])[0])
	callframe.Arg(frame, pointers.Get(ids))
	callframe.Arg(frame, billboard)
	callframe.Arg(frame, secondary)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_add_handles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the reference [Node3D] node for the gizmo. [param node] must inherit from [Node3D].
*/
//go:nosplit
func (self class) SetNode3d(node objects.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(node[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_set_node_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Node3D] node associated with this gizmo.
*/
//go:nosplit
func (self class) GetNode3d() objects.Node3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_get_node_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Node3D{classdb.Node3D(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the [EditorNode3DGizmoPlugin] that owns this gizmo. It's useful to retrieve materials using [method EditorNode3DGizmoPlugin.get_material].
*/
//go:nosplit
func (self class) GetPlugin() objects.EditorNode3DGizmoPlugin {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_get_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.EditorNode3DGizmoPlugin{classdb.EditorNode3DGizmoPlugin(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Removes everything in the gizmo including meshes, collisions and handles.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the gizmo's hidden state. If [code]true[/code], the gizmo will be hidden. If [code]false[/code], it will be shown.
*/
//go:nosplit
func (self class) SetHidden(hidden bool) {
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_set_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given subgizmo is currently selected. Can be used to highlight selected elements during [method _redraw].
*/
//go:nosplit
func (self class) IsSubgizmoSelected(id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_is_subgizmo_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a list of the currently selected subgizmos. Can be used to highlight selected elements during [method _redraw].
*/
//go:nosplit
func (self class) GetSubgizmoSelection() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmo.Bind_get_subgizmo_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorNode3DGizmo() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorNode3DGizmo() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3DGizmo() Node3DGizmo.Advanced {
	return *((*Node3DGizmo.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsNode3DGizmo() Node3DGizmo.Instance {
	return *((*Node3DGizmo.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_redraw":
		return reflect.ValueOf(self._redraw)
	case "_get_handle_name":
		return reflect.ValueOf(self._get_handle_name)
	case "_is_handle_highlighted":
		return reflect.ValueOf(self._is_handle_highlighted)
	case "_get_handle_value":
		return reflect.ValueOf(self._get_handle_value)
	case "_begin_handle_action":
		return reflect.ValueOf(self._begin_handle_action)
	case "_set_handle":
		return reflect.ValueOf(self._set_handle)
	case "_commit_handle":
		return reflect.ValueOf(self._commit_handle)
	case "_subgizmos_intersect_ray":
		return reflect.ValueOf(self._subgizmos_intersect_ray)
	case "_subgizmos_intersect_frustum":
		return reflect.ValueOf(self._subgizmos_intersect_frustum)
	case "_set_subgizmo_transform":
		return reflect.ValueOf(self._set_subgizmo_transform)
	case "_get_subgizmo_transform":
		return reflect.ValueOf(self._get_subgizmo_transform)
	case "_commit_subgizmos":
		return reflect.ValueOf(self._commit_subgizmos)
	default:
		return gd.VirtualByName(self.AsNode3DGizmo(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_redraw":
		return reflect.ValueOf(self._redraw)
	case "_get_handle_name":
		return reflect.ValueOf(self._get_handle_name)
	case "_is_handle_highlighted":
		return reflect.ValueOf(self._is_handle_highlighted)
	case "_get_handle_value":
		return reflect.ValueOf(self._get_handle_value)
	case "_begin_handle_action":
		return reflect.ValueOf(self._begin_handle_action)
	case "_set_handle":
		return reflect.ValueOf(self._set_handle)
	case "_commit_handle":
		return reflect.ValueOf(self._commit_handle)
	case "_subgizmos_intersect_ray":
		return reflect.ValueOf(self._subgizmos_intersect_ray)
	case "_subgizmos_intersect_frustum":
		return reflect.ValueOf(self._subgizmos_intersect_frustum)
	case "_set_subgizmo_transform":
		return reflect.ValueOf(self._set_subgizmo_transform)
	case "_get_subgizmo_transform":
		return reflect.ValueOf(self._get_subgizmo_transform)
	case "_commit_subgizmos":
		return reflect.ValueOf(self._commit_subgizmos)
	default:
		return gd.VirtualByName(self.AsNode3DGizmo(), name)
	}
}
func init() {
	classdb.Register("EditorNode3DGizmo", func(ptr gd.Object) any { return classdb.EditorNode3DGizmo(ptr) })
}
