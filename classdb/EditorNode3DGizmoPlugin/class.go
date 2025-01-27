// Package EditorNode3DGizmoPlugin provides methods for working with EditorNode3DGizmoPlugin object instances.
package EditorNode3DGizmoPlugin

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Plane"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Color"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable

/*
[EditorNode3DGizmoPlugin] allows you to define a new type of Gizmo. There are two main ways to do so: extending [EditorNode3DGizmoPlugin] for the simpler gizmos, or creating a new [EditorNode3DGizmo] type. See the tutorial in the documentation for more info.
To use [EditorNode3DGizmoPlugin], register it using the [method EditorPlugin.add_node_3d_gizmo_plugin] method first.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorNode3DGizmoPlugin)
*/
type Instance [1]gdclass.EditorNode3DGizmoPlugin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorNode3DGizmoPlugin() Instance
}
type Interface interface {
	//Override this method to define which Node3D nodes have a gizmo from this plugin. Whenever a [Node3D] node is added to a scene this method is called, if it returns [code]true[/code] the node gets a generic [EditorNode3DGizmo] assigned and is added to this plugin's list of active gizmos.
	HasGizmo(for_node_3d [1]gdclass.Node3D) bool
	//Override this method to return a custom [EditorNode3DGizmo] for the spatial nodes of your choice, return [code]null[/code] for the rest of nodes. See also [method _has_gizmo].
	CreateGizmo(for_node_3d [1]gdclass.Node3D) [1]gdclass.EditorNode3DGizmo
	//Override this method to provide the name that will appear in the gizmo visibility menu.
	GetGizmoName() string
	//Override this method to set the gizmo's priority. Gizmos with higher priority will have precedence when processing inputs like handles or subgizmos selection.
	//All built-in editor gizmos return a priority of [code]-1[/code]. If not overridden, this method will return [code]0[/code], which means custom gizmos will automatically get higher priority than built-in gizmos.
	GetPriority() int
	//Override this method to define whether the gizmos handled by this plugin can be hidden or not. Returns [code]true[/code] if not overridden.
	CanBeHidden() bool
	//Override this method to define whether Node3D with this gizmo should be selectable even when the gizmo is hidden.
	IsSelectableWhenHidden() bool
	//Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method EditorNode3DGizmo.clear] at the beginning of this method and then add visual elements depending on the node's properties.
	Redraw(gizmo [1]gdclass.EditorNode3DGizmo)
	//Override this method to provide gizmo's handle names. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
	GetHandleName(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) string
	//Override this method to return [code]true[/code] whenever to given handle should be highlighted in the editor. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
	IsHandleHighlighted(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) bool
	//Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
	//The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
	//Called for this plugin's active gizmos.
	GetHandleValue(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) any
	BeginHandleAction(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool)
	//Override this method to update the node's properties when the user drags a gizmo handle (previously added with [method EditorNode3DGizmo.add_handles]). The provided [param screen_pos] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
	//The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
	//Called for this plugin's active gizmos.
	SetHandle(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool, camera [1]gdclass.Camera3D, screen_pos Vector2.XY)
	//Override this method to commit a handle being edited (handles must have been previously added by [method EditorNode3DGizmo.add_handles] during [method _redraw]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
	//If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
	//The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
	//Called for this plugin's active gizmos.
	CommitHandle(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool, restore any, cancel bool)
	//Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param screen_pos] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
	SubgizmosIntersectRay(gizmo [1]gdclass.EditorNode3DGizmo, camera [1]gdclass.Camera3D, screen_pos Vector2.XY) int
	//Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and [param frustum_planes], this method should return which subgizmos are contained within the frustums. The [param frustum_planes] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, these identifiers can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
	SubgizmosIntersectFrustum(gizmo [1]gdclass.EditorNode3DGizmo, camera [1]gdclass.Camera3D, frustum_planes []Plane.NormalD) []int32
	//Override this method to return the current transform of a subgizmo. As with all subgizmo methods, the transform should be in local space respect to the gizmo's Node3D. This transform will be requested at the start of an edit and used in the [code]restore[/code] argument in [method _commit_subgizmos]. Called for this plugin's active gizmos.
	GetSubgizmoTransform(gizmo [1]gdclass.EditorNode3DGizmo, subgizmo_id int) Transform3D.BasisOrigin
	//Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the Node3D's local coordinate system. Called for this plugin's active gizmos.
	SetSubgizmoTransform(gizmo [1]gdclass.EditorNode3DGizmo, subgizmo_id int, transform Transform3D.BasisOrigin)
	//Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
	//If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action. As with all subgizmo methods, transforms are given in local space respect to the gizmo's Node3D. Called for this plugin's active gizmos.
	CommitSubgizmos(gizmo [1]gdclass.EditorNode3DGizmo, ids []int32, restores []Transform3D.BasisOrigin, cancel bool)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) HasGizmo(for_node_3d [1]gdclass.Node3D) (_ bool) { return }
func (self implementation) CreateGizmo(for_node_3d [1]gdclass.Node3D) (_ [1]gdclass.EditorNode3DGizmo) {
	return
}
func (self implementation) GetGizmoName() (_ string)                  { return }
func (self implementation) GetPriority() (_ int)                      { return }
func (self implementation) CanBeHidden() (_ bool)                     { return }
func (self implementation) IsSelectableWhenHidden() (_ bool)          { return }
func (self implementation) Redraw(gizmo [1]gdclass.EditorNode3DGizmo) { return }
func (self implementation) GetHandleName(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) (_ string) {
	return
}
func (self implementation) IsHandleHighlighted(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) (_ bool) {
	return
}
func (self implementation) GetHandleValue(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) (_ any) {
	return
}
func (self implementation) BeginHandleAction(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) {
	return
}
func (self implementation) SetHandle(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool, camera [1]gdclass.Camera3D, screen_pos Vector2.XY) {
	return
}
func (self implementation) CommitHandle(gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool, restore any, cancel bool) {
	return
}
func (self implementation) SubgizmosIntersectRay(gizmo [1]gdclass.EditorNode3DGizmo, camera [1]gdclass.Camera3D, screen_pos Vector2.XY) (_ int) {
	return
}
func (self implementation) SubgizmosIntersectFrustum(gizmo [1]gdclass.EditorNode3DGizmo, camera [1]gdclass.Camera3D, frustum_planes []Plane.NormalD) (_ []int32) {
	return
}
func (self implementation) GetSubgizmoTransform(gizmo [1]gdclass.EditorNode3DGizmo, subgizmo_id int) (_ Transform3D.BasisOrigin) {
	return
}
func (self implementation) SetSubgizmoTransform(gizmo [1]gdclass.EditorNode3DGizmo, subgizmo_id int, transform Transform3D.BasisOrigin) {
	return
}
func (self implementation) CommitSubgizmos(gizmo [1]gdclass.EditorNode3DGizmo, ids []int32, restores []Transform3D.BasisOrigin, cancel bool) {
	return
}

/*
Override this method to define which Node3D nodes have a gizmo from this plugin. Whenever a [Node3D] node is added to a scene this method is called, if it returns [code]true[/code] the node gets a generic [EditorNode3DGizmo] assigned and is added to this plugin's list of active gizmos.
*/
func (Instance) _has_gizmo(impl func(ptr unsafe.Pointer, for_node_3d [1]gdclass.Node3D) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_node_3d = [1]gdclass.Node3D{pointers.New[gdclass.Node3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(for_node_3d[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_node_3d)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to return a custom [EditorNode3DGizmo] for the spatial nodes of your choice, return [code]null[/code] for the rest of nodes. See also [method _has_gizmo].
*/
func (Instance) _create_gizmo(impl func(ptr unsafe.Pointer, for_node_3d [1]gdclass.Node3D) [1]gdclass.EditorNode3DGizmo) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_node_3d = [1]gdclass.Node3D{pointers.New[gdclass.Node3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(for_node_3d[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_node_3d)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to provide the name that will appear in the gizmo visibility menu.
*/
func (Instance) _get_gizmo_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to set the gizmo's priority. Gizmos with higher priority will have precedence when processing inputs like handles or subgizmos selection.
All built-in editor gizmos return a priority of [code]-1[/code]. If not overridden, this method will return [code]0[/code], which means custom gizmos will automatically get higher priority than built-in gizmos.
*/
func (Instance) _get_priority(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Override this method to define whether the gizmos handled by this plugin can be hidden or not. Returns [code]true[/code] if not overridden.
*/
func (Instance) _can_be_hidden(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define whether Node3D with this gizmo should be selectable even when the gizmo is hidden.
*/
func (Instance) _is_selectable_when_hidden(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method EditorNode3DGizmo.clear] at the beginning of this method and then add visual elements depending on the node's properties.
*/
func (Instance) _redraw(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo)
	}
}

/*
Override this method to provide gizmo's handle names. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
*/
func (Instance) _get_handle_name(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, int(handle_id), secondary)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to return [code]true[/code] whenever to given handle should be highlighted in the editor. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
*/
func (Instance) _is_handle_highlighted(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, int(handle_id), secondary)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
Called for this plugin's active gizmos.
*/
func (Instance) _get_handle_value(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, int(handle_id), secondary)
		ptr, ok := pointers.End(gd.NewVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _begin_handle_action(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, int(handle_id), secondary)
	}
}

/*
Override this method to update the node's properties when the user drags a gizmo handle (previously added with [method EditorNode3DGizmo.add_handles]). The provided [param screen_pos] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
Called for this plugin's active gizmos.
*/
func (Instance) _set_handle(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool, camera [1]gdclass.Camera3D, screen_pos Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		var camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}

		defer pointers.End(camera[0])
		var screen_pos = gd.UnsafeGet[gd.Vector2](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, int(handle_id), secondary, camera, screen_pos)
	}
}

/*
Override this method to commit a handle being edited (handles must have been previously added by [method EditorNode3DGizmo.add_handles] during [method _redraw]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
Called for this plugin's active gizmos.
*/
func (Instance) _commit_handle(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id int, secondary bool, restore any, cancel bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		var restore = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 3))
		defer pointers.End(restore)
		var cancel = gd.UnsafeGet[bool](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, int(handle_id), secondary, restore.Interface(), cancel)
	}
}

/*
Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param screen_pos] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
*/
func (Instance) _subgizmos_intersect_ray(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, camera [1]gdclass.Camera3D, screen_pos Vector2.XY) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(camera[0])
		var screen_pos = gd.UnsafeGet[gd.Vector2](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, camera, screen_pos)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and [param frustum_planes], this method should return which subgizmos are contained within the frustums. The [param frustum_planes] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, these identifiers can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
*/
func (Instance) _subgizmos_intersect_frustum(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, camera [1]gdclass.Camera3D, frustum_planes []Plane.NormalD) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(camera[0])
		var frustum_planes = Array.Through(gd.ArrayProxy[gd.Plane]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalArray(frustum_planes))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, camera, gd.ArrayAs[[]Plane.NormalD](gd.InternalArray(frustum_planes)))
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to return the current transform of a subgizmo. As with all subgizmo methods, the transform should be in local space respect to the gizmo's Node3D. This transform will be requested at the start of an edit and used in the [code]restore[/code] argument in [method _commit_subgizmos]. Called for this plugin's active gizmos.
*/
func (Instance) _get_subgizmo_transform(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, subgizmo_id int) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var subgizmo_id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, int(subgizmo_id))
		gd.UnsafeSet(p_back, gd.Transform3D(ret))
	}
}

/*
Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the Node3D's local coordinate system. Called for this plugin's active gizmos.
*/
func (Instance) _set_subgizmo_transform(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, subgizmo_id int, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var subgizmo_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, int(subgizmo_id), transform)
	}
}

/*
Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action. As with all subgizmo methods, transforms are given in local space respect to the gizmo's Node3D. Called for this plugin's active gizmos.
*/
func (Instance) _commit_subgizmos(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, ids []int32, restores []Transform3D.BasisOrigin, cancel bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var ids = pointers.New[gd.PackedInt32Array](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(ids)
		var restores = Array.Through(gd.ArrayProxy[gd.Transform3D]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalArray(restores))
		var cancel = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, ids.AsSlice(), gd.ArrayAs[[]Transform3D.BasisOrigin](gd.InternalArray(restores)), cancel)
	}
}

/*
Creates an unshaded material with its variants (selected and/or editable) and adds them to the internal material list. They can then be accessed with [method get_material] and used in [method EditorNode3DGizmo.add_mesh] and [method EditorNode3DGizmo.add_lines]. Should not be overridden.
*/
func (self Instance) CreateMaterial(name string, color Color.RGBA) { //gd:EditorNode3DGizmoPlugin.create_material
	class(self).CreateMaterial(String.New(name), gd.Color(color), false, false, false)
}

/*
Creates an icon material with its variants (selected and/or editable) and adds them to the internal material list. They can then be accessed with [method get_material] and used in [method EditorNode3DGizmo.add_unscaled_billboard]. Should not be overridden.
*/
func (self Instance) CreateIconMaterial(name string, texture [1]gdclass.Texture2D) { //gd:EditorNode3DGizmoPlugin.create_icon_material
	class(self).CreateIconMaterial(String.New(name), texture, false, gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Creates a handle material with its variants (selected and/or editable) and adds them to the internal material list. They can then be accessed with [method get_material] and used in [method EditorNode3DGizmo.add_handles]. Should not be overridden.
You can optionally provide a texture to use instead of the default icon.
*/
func (self Instance) CreateHandleMaterial(name string) { //gd:EditorNode3DGizmoPlugin.create_handle_material
	class(self).CreateHandleMaterial(String.New(name), false, [1][1]gdclass.Texture2D{}[0])
}

/*
Adds a new material to the internal material list for the plugin. It can then be accessed with [method get_material]. Should not be overridden.
*/
func (self Instance) AddMaterial(name string, material [1]gdclass.StandardMaterial3D) { //gd:EditorNode3DGizmoPlugin.add_material
	class(self).AddMaterial(String.New(name), material)
}

/*
Gets material from the internal list of materials. If an [EditorNode3DGizmo] is provided, it will try to get the corresponding variant (selected and/or editable).
*/
func (self Instance) GetMaterial(name string) [1]gdclass.StandardMaterial3D { //gd:EditorNode3DGizmoPlugin.get_material
	return [1]gdclass.StandardMaterial3D(class(self).GetMaterial(String.New(name), [1][1]gdclass.EditorNode3DGizmo{}[0]))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorNode3DGizmoPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorNode3DGizmoPlugin"))
	casted := Instance{*(*gdclass.EditorNode3DGizmoPlugin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Override this method to define which Node3D nodes have a gizmo from this plugin. Whenever a [Node3D] node is added to a scene this method is called, if it returns [code]true[/code] the node gets a generic [EditorNode3DGizmo] assigned and is added to this plugin's list of active gizmos.
*/
func (class) _has_gizmo(impl func(ptr unsafe.Pointer, for_node_3d [1]gdclass.Node3D) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_node_3d = [1]gdclass.Node3D{pointers.New[gdclass.Node3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(for_node_3d[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_node_3d)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to return a custom [EditorNode3DGizmo] for the spatial nodes of your choice, return [code]null[/code] for the rest of nodes. See also [method _has_gizmo].
*/
func (class) _create_gizmo(impl func(ptr unsafe.Pointer, for_node_3d [1]gdclass.Node3D) [1]gdclass.EditorNode3DGizmo) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_node_3d = [1]gdclass.Node3D{pointers.New[gdclass.Node3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(for_node_3d[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_node_3d)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to provide the name that will appear in the gizmo visibility menu.
*/
func (class) _get_gizmo_name(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to set the gizmo's priority. Gizmos with higher priority will have precedence when processing inputs like handles or subgizmos selection.
All built-in editor gizmos return a priority of [code]-1[/code]. If not overridden, this method will return [code]0[/code], which means custom gizmos will automatically get higher priority than built-in gizmos.
*/
func (class) _get_priority(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define whether the gizmos handled by this plugin can be hidden or not. Returns [code]true[/code] if not overridden.
*/
func (class) _can_be_hidden(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define whether Node3D with this gizmo should be selectable even when the gizmo is hidden.
*/
func (class) _is_selectable_when_hidden(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method EditorNode3DGizmo.clear] at the beginning of this method and then add visual elements depending on the node's properties.
*/
func (class) _redraw(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo)
	}
}

/*
Override this method to provide gizmo's handle names. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
*/
func (class) _get_handle_name(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id gd.Int, secondary bool) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, handle_id, secondary)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to return [code]true[/code] whenever to given handle should be highlighted in the editor. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
*/
func (class) _is_handle_highlighted(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id gd.Int, secondary bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, handle_id, secondary)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
Called for this plugin's active gizmos.
*/
func (class) _get_handle_value(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id gd.Int, secondary bool) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, handle_id, secondary)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _begin_handle_action(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id gd.Int, secondary bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, handle_id, secondary)
	}
}

/*
Override this method to update the node's properties when the user drags a gizmo handle (previously added with [method EditorNode3DGizmo.add_handles]). The provided [param screen_pos] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
Called for this plugin's active gizmos.
*/
func (class) _set_handle(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id gd.Int, secondary bool, camera [1]gdclass.Camera3D, screen_pos gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		var camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}

		defer pointers.End(camera[0])
		var screen_pos = gd.UnsafeGet[gd.Vector2](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, handle_id, secondary, camera, screen_pos)
	}
}

/*
Override this method to commit a handle being edited (handles must have been previously added by [method EditorNode3DGizmo.add_handles] during [method _redraw]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
Called for this plugin's active gizmos.
*/
func (class) _commit_handle(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, handle_id gd.Int, secondary bool, restore gd.Variant, cancel bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var handle_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var secondary = gd.UnsafeGet[bool](p_args, 2)

		var restore = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 3))
		defer pointers.End(restore)
		var cancel = gd.UnsafeGet[bool](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, handle_id, secondary, restore, cancel)
	}
}

/*
Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param screen_pos] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
*/
func (class) _subgizmos_intersect_ray(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, camera [1]gdclass.Camera3D, screen_pos gd.Vector2) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(camera[0])
		var screen_pos = gd.UnsafeGet[gd.Vector2](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, camera, screen_pos)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and [param frustum_planes], this method should return which subgizmos are contained within the frustums. The [param frustum_planes] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, these identifiers can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
*/
func (class) _subgizmos_intersect_frustum(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, camera [1]gdclass.Camera3D, frustum_planes Array.Contains[gd.Plane]) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var camera = [1]gdclass.Camera3D{pointers.New[gdclass.Camera3D]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(camera[0])
		var frustum_planes = Array.Through(gd.ArrayProxy[gd.Plane]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalArray(frustum_planes))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, camera, frustum_planes)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to return the current transform of a subgizmo. As with all subgizmo methods, the transform should be in local space respect to the gizmo's Node3D. This transform will be requested at the start of an edit and used in the [code]restore[/code] argument in [method _commit_subgizmos]. Called for this plugin's active gizmos.
*/
func (class) _get_subgizmo_transform(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, subgizmo_id gd.Int) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var subgizmo_id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, subgizmo_id)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the Node3D's local coordinate system. Called for this plugin's active gizmos.
*/
func (class) _set_subgizmo_transform(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, subgizmo_id gd.Int, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var subgizmo_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, subgizmo_id, transform)
	}
}

/*
Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action. As with all subgizmo methods, transforms are given in local space respect to the gizmo's Node3D. Called for this plugin's active gizmos.
*/
func (class) _commit_subgizmos(impl func(ptr unsafe.Pointer, gizmo [1]gdclass.EditorNode3DGizmo, ids gd.PackedInt32Array, restores Array.Contains[gd.Transform3D], cancel bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var gizmo = [1]gdclass.EditorNode3DGizmo{pointers.New[gdclass.EditorNode3DGizmo]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(gizmo[0])
		var ids = pointers.New[gd.PackedInt32Array](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(ids)
		var restores = Array.Through(gd.ArrayProxy[gd.Transform3D]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalArray(restores))
		var cancel = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, gizmo, ids, restores, cancel)
	}
}

/*
Creates an unshaded material with its variants (selected and/or editable) and adds them to the internal material list. They can then be accessed with [method get_material] and used in [method EditorNode3DGizmo.add_mesh] and [method EditorNode3DGizmo.add_lines]. Should not be overridden.
*/
//go:nosplit
func (self class) CreateMaterial(name String.Readable, color gd.Color, billboard bool, on_top bool, use_vertex_color bool) { //gd:EditorNode3DGizmoPlugin.create_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, color)
	callframe.Arg(frame, billboard)
	callframe.Arg(frame, on_top)
	callframe.Arg(frame, use_vertex_color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmoPlugin.Bind_create_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates an icon material with its variants (selected and/or editable) and adds them to the internal material list. They can then be accessed with [method get_material] and used in [method EditorNode3DGizmo.add_unscaled_billboard]. Should not be overridden.
*/
//go:nosplit
func (self class) CreateIconMaterial(name String.Readable, texture [1]gdclass.Texture2D, on_top bool, color gd.Color) { //gd:EditorNode3DGizmoPlugin.create_icon_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, on_top)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmoPlugin.Bind_create_icon_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a handle material with its variants (selected and/or editable) and adds them to the internal material list. They can then be accessed with [method get_material] and used in [method EditorNode3DGizmo.add_handles]. Should not be overridden.
You can optionally provide a texture to use instead of the default icon.
*/
//go:nosplit
func (self class) CreateHandleMaterial(name String.Readable, billboard bool, texture [1]gdclass.Texture2D) { //gd:EditorNode3DGizmoPlugin.create_handle_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, billboard)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmoPlugin.Bind_create_handle_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a new material to the internal material list for the plugin. It can then be accessed with [method get_material]. Should not be overridden.
*/
//go:nosplit
func (self class) AddMaterial(name String.Readable, material [1]gdclass.StandardMaterial3D) { //gd:EditorNode3DGizmoPlugin.add_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmoPlugin.Bind_add_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets material from the internal list of materials. If an [EditorNode3DGizmo] is provided, it will try to get the corresponding variant (selected and/or editable).
*/
//go:nosplit
func (self class) GetMaterial(name String.Readable, gizmo [1]gdclass.EditorNode3DGizmo) [1]gdclass.StandardMaterial3D { //gd:EditorNode3DGizmoPlugin.get_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(gizmo[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorNode3DGizmoPlugin.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.StandardMaterial3D{gd.PointerWithOwnershipTransferredToGo[gdclass.StandardMaterial3D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsEditorNode3DGizmoPlugin() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorNode3DGizmoPlugin() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_has_gizmo":
		return reflect.ValueOf(self._has_gizmo)
	case "_create_gizmo":
		return reflect.ValueOf(self._create_gizmo)
	case "_get_gizmo_name":
		return reflect.ValueOf(self._get_gizmo_name)
	case "_get_priority":
		return reflect.ValueOf(self._get_priority)
	case "_can_be_hidden":
		return reflect.ValueOf(self._can_be_hidden)
	case "_is_selectable_when_hidden":
		return reflect.ValueOf(self._is_selectable_when_hidden)
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
	case "_get_subgizmo_transform":
		return reflect.ValueOf(self._get_subgizmo_transform)
	case "_set_subgizmo_transform":
		return reflect.ValueOf(self._set_subgizmo_transform)
	case "_commit_subgizmos":
		return reflect.ValueOf(self._commit_subgizmos)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_has_gizmo":
		return reflect.ValueOf(self._has_gizmo)
	case "_create_gizmo":
		return reflect.ValueOf(self._create_gizmo)
	case "_get_gizmo_name":
		return reflect.ValueOf(self._get_gizmo_name)
	case "_get_priority":
		return reflect.ValueOf(self._get_priority)
	case "_can_be_hidden":
		return reflect.ValueOf(self._can_be_hidden)
	case "_is_selectable_when_hidden":
		return reflect.ValueOf(self._is_selectable_when_hidden)
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
	case "_get_subgizmo_transform":
		return reflect.ValueOf(self._get_subgizmo_transform)
	case "_set_subgizmo_transform":
		return reflect.ValueOf(self._set_subgizmo_transform)
	case "_commit_subgizmos":
		return reflect.ValueOf(self._commit_subgizmos)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("EditorNode3DGizmoPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorNode3DGizmoPlugin{*(*gdclass.EditorNode3DGizmoPlugin)(unsafe.Pointer(&ptr))}
	})
}
