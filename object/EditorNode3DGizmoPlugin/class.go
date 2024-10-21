package EditorNode3DGizmoPlugin

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[EditorNode3DGizmoPlugin] allows you to define a new type of Gizmo. There are two main ways to do so: extending [EditorNode3DGizmoPlugin] for the simpler gizmos, or creating a new [EditorNode3DGizmo] type. See the tutorial in the documentation for more info.
To use [EditorNode3DGizmoPlugin], register it using the [method EditorPlugin.add_node_3d_gizmo_plugin] method first.
	// EditorNode3DGizmoPlugin methods that can be overridden by a [Class] that extends it.
	type EditorNode3DGizmoPlugin interface {
		//Override this method to define which Node3D nodes have a gizmo from this plugin. Whenever a [Node3D] node is added to a scene this method is called, if it returns [code]true[/code] the node gets a generic [EditorNode3DGizmo] assigned and is added to this plugin's list of active gizmos.
		HasGizmo(for_node_3d object.Node3D) bool
		//Override this method to return a custom [EditorNode3DGizmo] for the spatial nodes of your choice, return [code]null[/code] for the rest of nodes. See also [method _has_gizmo].
		CreateGizmo(for_node_3d object.Node3D) object.EditorNode3DGizmo
		//Override this method to provide the name that will appear in the gizmo visibility menu.
		GetGizmoName() gd.String
		//Override this method to set the gizmo's priority. Gizmos with higher priority will have precedence when processing inputs like handles or subgizmos selection.
		//All built-in editor gizmos return a priority of [code]-1[/code]. If not overridden, this method will return [code]0[/code], which means custom gizmos will automatically get higher priority than built-in gizmos.
		GetPriority() gd.Int
		//Override this method to define whether the gizmos handled by this plugin can be hidden or not. Returns [code]true[/code] if not overridden.
		CanBeHidden() bool
		//Override this method to define whether Node3D with this gizmo should be selectable even when the gizmo is hidden.
		IsSelectableWhenHidden() bool
		//Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method EditorNode3DGizmo.clear] at the beginning of this method and then add visual elements depending on the node's properties.
		Redraw(gizmo object.EditorNode3DGizmo) 
		//Override this method to provide gizmo's handle names. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
		GetHandleName(gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool) gd.String
		//Override this method to return [code]true[/code] whenever to given handle should be highlighted in the editor. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
		IsHandleHighlighted(gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool) bool
		//Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
		//The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
		//Called for this plugin's active gizmos.
		GetHandleValue(gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool) gd.Variant
		BeginHandleAction(gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool) 
		//Override this method to update the node's properties when the user drags a gizmo handle (previously added with [method EditorNode3DGizmo.add_handles]). The provided [param screen_pos] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
		//The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
		//Called for this plugin's active gizmos.
		SetHandle(gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool, camera object.Camera3D, screen_pos gd.Vector2) 
		//Override this method to commit a handle being edited (handles must have been previously added by [method EditorNode3DGizmo.add_handles] during [method _redraw]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
		//If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
		//The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
		//Called for this plugin's active gizmos.
		CommitHandle(gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool, restore gd.Variant, cancel bool) 
		//Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param screen_pos] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
		SubgizmosIntersectRay(gizmo object.EditorNode3DGizmo, camera object.Camera3D, screen_pos gd.Vector2) gd.Int
		//Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and [param frustum_planes], this method should return which subgizmos are contained within the frustums. The [param frustum_planes] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, these identifiers can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
		SubgizmosIntersectFrustum(gizmo object.EditorNode3DGizmo, camera object.Camera3D, frustum_planes gd.ArrayOf[gd.Plane]) gd.PackedInt32Array
		//Override this method to return the current transform of a subgizmo. As with all subgizmo methods, the transform should be in local space respect to the gizmo's Node3D. This transform will be requested at the start of an edit and used in the [code]restore[/code] argument in [method _commit_subgizmos]. Called for this plugin's active gizmos.
		GetSubgizmoTransform(gizmo object.EditorNode3DGizmo, subgizmo_id gd.Int) gd.Transform3D
		//Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the Node3D's local coordinate system. Called for this plugin's active gizmos.
		SetSubgizmoTransform(gizmo object.EditorNode3DGizmo, subgizmo_id gd.Int, transform gd.Transform3D) 
		//Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
		//If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action. As with all subgizmo methods, transforms are given in local space respect to the gizmo's Node3D. Called for this plugin's active gizmos.
		CommitSubgizmos(gizmo object.EditorNode3DGizmo, ids gd.PackedInt32Array, restores gd.ArrayOf[gd.Transform3D], cancel bool) 
	}

*/
type Simple [1]classdb.EditorNode3DGizmoPlugin
func (self Simple) CreateMaterial(name string, color gd.Color, billboard bool, on_top bool, use_vertex_color bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateMaterial(gc.String(name), color, billboard, on_top, use_vertex_color)
}
func (self Simple) CreateIconMaterial(name string, texture [1]classdb.Texture2D, on_top bool, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateIconMaterial(gc.String(name), texture, on_top, color)
}
func (self Simple) CreateHandleMaterial(name string, billboard bool, texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateHandleMaterial(gc.String(name), billboard, texture)
}
func (self Simple) AddMaterial(name string, material [1]classdb.StandardMaterial3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddMaterial(gc.String(name), material)
}
func (self Simple) GetMaterial(name string, gizmo [1]classdb.EditorNode3DGizmo) [1]classdb.StandardMaterial3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.StandardMaterial3D(Expert(self).GetMaterial(gc, gc.String(name), gizmo))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorNode3DGizmoPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Override this method to define which Node3D nodes have a gizmo from this plugin. Whenever a [Node3D] node is added to a scene this method is called, if it returns [code]true[/code] the node gets a generic [EditorNode3DGizmo] assigned and is added to this plugin's list of active gizmos.
*/
func (class) _has_gizmo(impl func(ptr unsafe.Pointer, for_node_3d object.Node3D) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var for_node_3d object.Node3D
		for_node_3d[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_node_3d)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to return a custom [EditorNode3DGizmo] for the spatial nodes of your choice, return [code]null[/code] for the rest of nodes. See also [method _has_gizmo].
*/
func (class) _create_gizmo(impl func(ptr unsafe.Pointer, for_node_3d object.Node3D) object.EditorNode3DGizmo, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var for_node_3d object.Node3D
		for_node_3d[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_node_3d)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Override this method to provide the name that will appear in the gizmo visibility menu.
*/
func (class) _get_gizmo_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to set the gizmo's priority. Gizmos with higher priority will have precedence when processing inputs like handles or subgizmos selection.
All built-in editor gizmos return a priority of [code]-1[/code]. If not overridden, this method will return [code]0[/code], which means custom gizmos will automatically get higher priority than built-in gizmos.
*/
func (class) _get_priority(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to define whether the gizmos handled by this plugin can be hidden or not. Returns [code]true[/code] if not overridden.
*/
func (class) _can_be_hidden(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to define whether Node3D with this gizmo should be selectable even when the gizmo is hidden.
*/
func (class) _is_selectable_when_hidden(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method EditorNode3DGizmo.clear] at the beginning of this method and then add visual elements depending on the node's properties.
*/
func (class) _redraw(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, gizmo)
		ctx.End()
	}
}

/*
Override this method to provide gizmo's handle names. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
*/
func (class) _get_handle_name(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var handle_id = gd.UnsafeGet[gd.Int](p_args,1)
		var secondary = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, handle_id, secondary)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to return [code]true[/code] whenever to given handle should be highlighted in the editor. The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information). Called for this plugin's active gizmos.
*/
func (class) _is_handle_highlighted(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var handle_id = gd.UnsafeGet[gd.Int](p_args,1)
		var secondary = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, handle_id, secondary)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
Called for this plugin's active gizmos.
*/
func (class) _get_handle_value(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var handle_id = gd.UnsafeGet[gd.Int](p_args,1)
		var secondary = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, handle_id, secondary)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _begin_handle_action(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var handle_id = gd.UnsafeGet[gd.Int](p_args,1)
		var secondary = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, gizmo, handle_id, secondary)
		ctx.End()
	}
}

/*
Override this method to update the node's properties when the user drags a gizmo handle (previously added with [method EditorNode3DGizmo.add_handles]). The provided [param screen_pos] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
Called for this plugin's active gizmos.
*/
func (class) _set_handle(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool, camera object.Camera3D, screen_pos gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var handle_id = gd.UnsafeGet[gd.Int](p_args,1)
		var secondary = gd.UnsafeGet[bool](p_args,2)
		var camera object.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,3)}))
		var screen_pos = gd.UnsafeGet[gd.Vector2](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, gizmo, handle_id, secondary, camera, screen_pos)
		ctx.End()
	}
}

/*
Override this method to commit a handle being edited (handles must have been previously added by [method EditorNode3DGizmo.add_handles] during [method _redraw]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method EditorNode3DGizmo.add_handles] for more information).
Called for this plugin's active gizmos.
*/
func (class) _commit_handle(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, handle_id gd.Int, secondary bool, restore gd.Variant, cancel bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var handle_id = gd.UnsafeGet[gd.Int](p_args,1)
		var secondary = gd.UnsafeGet[bool](p_args,2)
		var restore = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,3))
		var cancel = gd.UnsafeGet[bool](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, gizmo, handle_id, secondary, restore, cancel)
		ctx.End()
	}
}

/*
Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param screen_pos] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
*/
func (class) _subgizmos_intersect_ray(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, camera object.Camera3D, screen_pos gd.Vector2) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var camera object.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var screen_pos = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, camera, screen_pos)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and [param frustum_planes], this method should return which subgizmos are contained within the frustums. The [param frustum_planes] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, these identifiers can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos]. Called for this plugin's active gizmos.
*/
func (class) _subgizmos_intersect_frustum(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, camera object.Camera3D, frustum_planes gd.ArrayOf[gd.Plane]) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var camera object.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var frustum_planes = gd.TypedArray[gd.Plane](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2)))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, camera, frustum_planes)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to return the current transform of a subgizmo. As with all subgizmo methods, the transform should be in local space respect to the gizmo's Node3D. This transform will be requested at the start of an edit and used in the [code]restore[/code] argument in [method _commit_subgizmos]. Called for this plugin's active gizmos.
*/
func (class) _get_subgizmo_transform(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, subgizmo_id gd.Int) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var subgizmo_id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gizmo, subgizmo_id)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the Node3D's local coordinate system. Called for this plugin's active gizmos.
*/
func (class) _set_subgizmo_transform(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, subgizmo_id gd.Int, transform gd.Transform3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var subgizmo_id = gd.UnsafeGet[gd.Int](p_args,1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, gizmo, subgizmo_id, transform)
		ctx.End()
	}
}

/*
Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action. As with all subgizmo methods, transforms are given in local space respect to the gizmo's Node3D. Called for this plugin's active gizmos.
*/
func (class) _commit_subgizmos(impl func(ptr unsafe.Pointer, gizmo object.EditorNode3DGizmo, ids gd.PackedInt32Array, restores gd.ArrayOf[gd.Transform3D], cancel bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var gizmo object.EditorNode3DGizmo
		gizmo[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var ids = mmm.Let[gd.PackedInt32Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		var restores = gd.TypedArray[gd.Transform3D](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2)))
		var cancel = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, gizmo, ids, restores, cancel)
		ctx.End()
	}
}

/*
Creates an unshaded material with its variants (selected and/or editable) and adds them to the internal material list. They can then be accessed with [method get_material] and used in [method EditorNode3DGizmo.add_mesh] and [method EditorNode3DGizmo.add_lines]. Should not be overridden.
*/
//go:nosplit
func (self class) CreateMaterial(name gd.String, color gd.Color, billboard bool, on_top bool, use_vertex_color bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, color)
	callframe.Arg(frame, billboard)
	callframe.Arg(frame, on_top)
	callframe.Arg(frame, use_vertex_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmoPlugin.Bind_create_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates an icon material with its variants (selected and/or editable) and adds them to the internal material list. They can then be accessed with [method get_material] and used in [method EditorNode3DGizmo.add_unscaled_billboard]. Should not be overridden.
*/
//go:nosplit
func (self class) CreateIconMaterial(name gd.String, texture object.Texture2D, on_top bool, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, on_top)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmoPlugin.Bind_create_icon_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a handle material with its variants (selected and/or editable) and adds them to the internal material list. They can then be accessed with [method get_material] and used in [method EditorNode3DGizmo.add_handles]. Should not be overridden.
You can optionally provide a texture to use instead of the default icon.
*/
//go:nosplit
func (self class) CreateHandleMaterial(name gd.String, billboard bool, texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, billboard)
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmoPlugin.Bind_create_handle_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new material to the internal material list for the plugin. It can then be accessed with [method get_material]. Should not be overridden.
*/
//go:nosplit
func (self class) AddMaterial(name gd.String, material object.StandardMaterial3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmoPlugin.Bind_add_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets material from the internal list of materials. If an [EditorNode3DGizmo] is provided, it will try to get the corresponding variant (selected and/or editable).
*/
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime, name gd.String, gizmo object.EditorNode3DGizmo) object.StandardMaterial3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(gizmo[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmoPlugin.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.StandardMaterial3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorNode3DGizmoPlugin() Expert { return self[0].AsEditorNode3DGizmoPlugin() }


//go:nosplit
func (self Simple) AsEditorNode3DGizmoPlugin() Simple { return self[0].AsEditorNode3DGizmoPlugin() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_has_gizmo": return reflect.ValueOf(self._has_gizmo);
	case "_create_gizmo": return reflect.ValueOf(self._create_gizmo);
	case "_get_gizmo_name": return reflect.ValueOf(self._get_gizmo_name);
	case "_get_priority": return reflect.ValueOf(self._get_priority);
	case "_can_be_hidden": return reflect.ValueOf(self._can_be_hidden);
	case "_is_selectable_when_hidden": return reflect.ValueOf(self._is_selectable_when_hidden);
	case "_redraw": return reflect.ValueOf(self._redraw);
	case "_get_handle_name": return reflect.ValueOf(self._get_handle_name);
	case "_is_handle_highlighted": return reflect.ValueOf(self._is_handle_highlighted);
	case "_get_handle_value": return reflect.ValueOf(self._get_handle_value);
	case "_begin_handle_action": return reflect.ValueOf(self._begin_handle_action);
	case "_set_handle": return reflect.ValueOf(self._set_handle);
	case "_commit_handle": return reflect.ValueOf(self._commit_handle);
	case "_subgizmos_intersect_ray": return reflect.ValueOf(self._subgizmos_intersect_ray);
	case "_subgizmos_intersect_frustum": return reflect.ValueOf(self._subgizmos_intersect_frustum);
	case "_get_subgizmo_transform": return reflect.ValueOf(self._get_subgizmo_transform);
	case "_set_subgizmo_transform": return reflect.ValueOf(self._set_subgizmo_transform);
	case "_commit_subgizmos": return reflect.ValueOf(self._commit_subgizmos);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorNode3DGizmoPlugin", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
