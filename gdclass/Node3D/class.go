package Node3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Most basic 3D game object, with a [Transform3D] and visibility settings. All other 3D game objects inherit from [Node3D]. Use [Node3D] as a parent node to move, scale, rotate and show/hide children in a 3D project.
Affine operations (rotate, scale, translate) happen in parent's local coordinate system, unless the [Node3D] object is set as top-level. Affine operations in this coordinate system correspond to direct affine operations on the [Node3D]'s transform. The word local below refers to this coordinate system. The coordinate system that is attached to the [Node3D] object itself is referred to as object-local coordinate system.
[b]Note:[/b] Unless otherwise specified, all methods that have angle parameters must have angles specified as [i]radians[/i]. To convert degrees to radians, use [method @GlobalScope.deg_to_rad].
[b]Note:[/b] Be aware that "Spatial" nodes are now called "Node3D" starting with Godot 4. Any Godot 3.x references to "Spatial" nodes refer to "Node3D" in Godot 4.

*/
type Go [1]classdb.Node3D

/*
Returns the parent [Node3D], or [code]null[/code] if no parent exists, the parent is not of type [Node3D], or [member top_level] is [code]true[/code].
[b]Note:[/b] Calling this method is not equivalent to [code]get_parent() as Node3D[/code], which does not take [member top_level] into account.
*/
func (self Go) GetParentNode3d() gdclass.Node3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Node3D(class(self).GetParentNode3d(gc))
}

/*
Sets whether the node ignores notification that its transformation (global or local) changed.
*/
func (self Go) SetIgnoreTransformNotification(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIgnoreTransformNotification(enabled)
}

/*
Sets whether the node uses a scale of [code](1, 1, 1)[/code] or its local transformation scale. Changes to the local transformation scale are preserved.
*/
func (self Go) SetDisableScale(disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisableScale(disable)
}

/*
Returns whether this node uses a scale of [code](1, 1, 1)[/code] or its local transformation scale.
*/
func (self Go) IsScaleDisabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsScaleDisabled())
}

/*
Returns the current [World3D] resource this [Node3D] node is registered to.
*/
func (self Go) GetWorld3d() gdclass.World3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.World3D(class(self).GetWorld3d(gc))
}

/*
Forces the transform to update. Transform changes in physics are not instant for performance reasons. Transforms are accumulated and then set. Use this if you need an up-to-date transform when doing physics operations.
*/
func (self Go) ForceUpdateTransform() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ForceUpdateTransform()
}

/*
Updates all the [Node3D] gizmos attached to this node.
*/
func (self Go) UpdateGizmos() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UpdateGizmos()
}

/*
Attach an editor gizmo to this [Node3D].
[b]Note:[/b] The gizmo object would typically be an instance of [EditorNode3DGizmo], but the argument type is kept generic to avoid creating a dependency on editor classes in [Node3D].
*/
func (self Go) AddGizmo(gizmo gdclass.Node3DGizmo) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddGizmo(gizmo)
}

/*
Returns all the gizmos attached to this [Node3D].
*/
func (self Go) GetGizmos() gd.ArrayOf[gdclass.Node3DGizmo] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Node3DGizmo](class(self).GetGizmos(gc))
}

/*
Clear all gizmos attached to this [Node3D].
*/
func (self Go) ClearGizmos() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearGizmos()
}

/*
Set subgizmo selection for this node in the editor.
[b]Note:[/b] The gizmo object would typically be an instance of [EditorNode3DGizmo], but the argument type is kept generic to avoid creating a dependency on editor classes in [Node3D].
*/
func (self Go) SetSubgizmoSelection(gizmo gdclass.Node3DGizmo, id int, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSubgizmoSelection(gizmo, gd.Int(id), transform)
}

/*
Clears subgizmo selection for this node in the editor. Useful when subgizmo IDs become invalid after a property change.
*/
func (self Go) ClearSubgizmoSelection() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearSubgizmoSelection()
}

/*
Returns [code]true[/code] if the node is present in the [SceneTree], its [member visible] property is [code]true[/code] and all its ancestors are also visible. If any ancestor is hidden, this node will not be visible in the scene tree.
*/
func (self Go) IsVisibleInTree() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsVisibleInTree())
}

/*
Enables rendering of this node. Changes [member visible] to [code]true[/code].
*/
func (self Go) Show() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Show()
}

/*
Disables rendering of this node. Changes [member visible] to [code]false[/code].
*/
func (self Go) Hide() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Hide()
}

/*
Sets whether the node notifies about its local transformation changes. [Node3D] will not propagate this by default.
*/
func (self Go) SetNotifyLocalTransform(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNotifyLocalTransform(enable)
}

/*
Returns whether node notifies about its local transformation changes. [Node3D] will not propagate this by default.
*/
func (self Go) IsLocalTransformNotificationEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLocalTransformNotificationEnabled())
}

/*
Sets whether the node notifies about its global and local transformation changes. [Node3D] will not propagate this by default, unless it is in the editor context and it has a valid gizmo.
*/
func (self Go) SetNotifyTransform(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNotifyTransform(enable)
}

/*
Returns whether the node notifies about its global and local transformation changes. [Node3D] will not propagate this by default.
*/
func (self Go) IsTransformNotificationEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsTransformNotificationEnabled())
}

/*
Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians.
*/
func (self Go) Rotate(axis gd.Vector3, angle float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Rotate(axis, gd.Float(angle))
}

/*
Rotates the global (world) transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in global coordinate system.
*/
func (self Go) GlobalRotate(axis gd.Vector3, angle float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).GlobalRotate(axis, gd.Float(angle))
}

/*
Scales the global (world) transformation by the given [Vector3] scale factors.
*/
func (self Go) GlobalScale(scale gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).GlobalScale(scale)
}

/*
Moves the global (world) transformation by [Vector3] offset. The offset is in global coordinate system.
*/
func (self Go) GlobalTranslate(offset gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).GlobalTranslate(offset)
}

/*
Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in object-local coordinate system.
*/
func (self Go) RotateObjectLocal(axis gd.Vector3, angle float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RotateObjectLocal(axis, gd.Float(angle))
}

/*
Scales the local transformation by given 3D scale factors in object-local coordinate system.
*/
func (self Go) ScaleObjectLocal(scale gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ScaleObjectLocal(scale)
}

/*
Changes the node's position by the given offset [Vector3] in local space.
*/
func (self Go) TranslateObjectLocal(offset gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TranslateObjectLocal(offset)
}

/*
Rotates the local transformation around the X axis by angle in radians.
*/
func (self Go) RotateX(angle float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RotateX(gd.Float(angle))
}

/*
Rotates the local transformation around the Y axis by angle in radians.
*/
func (self Go) RotateY(angle float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RotateY(gd.Float(angle))
}

/*
Rotates the local transformation around the Z axis by angle in radians.
*/
func (self Go) RotateZ(angle float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RotateZ(gd.Float(angle))
}

/*
Changes the node's position by the given offset [Vector3].
Note that the translation [param offset] is affected by the node's scale, so if scaled by e.g. [code](10, 1, 1)[/code], a translation by an offset of [code](2, 0, 0)[/code] would actually add 20 ([code]2 * 10[/code]) to the X coordinate.
*/
func (self Go) Translate(offset gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Translate(offset)
}

/*
Resets this node's transformations (like scale, skew and taper) preserving its rotation and translation by performing Gram-Schmidt orthonormalization on this node's [Transform3D].
*/
func (self Go) Orthonormalize() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Orthonormalize()
}

/*
Reset all transformations for this node (sets its [Transform3D] to the identity matrix).
*/
func (self Go) SetIdentity() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIdentity()
}

/*
Rotates the node so that the local forward axis (-Z, [constant Vector3.FORWARD]) points toward the [param target] position.
The local up axis (+Y) points as close to the [param up] vector as possible while staying perpendicular to the local forward axis. The resulting transform is orthogonal, and the scale is preserved. Non-uniform scaling may not work correctly.
The [param target] position cannot be the same as the node's position, the [param up] vector cannot be zero, and the direction from the node's position to the [param target] vector cannot be parallel to the [param up] vector.
Operations take place in global space, which means that the node must be in the scene tree.
If [param use_model_front] is [code]true[/code], the +Z axis (asset front) is treated as forward (implies +X is left) and points toward the [param target] position. By default, the -Z axis (camera forward) is treated as forward (implies +X is right).
*/
func (self Go) LookAt(target gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).LookAt(target, gd.Vector3{0, 1, 0}, false)
}

/*
Moves the node to the specified [param position], and then rotates the node to point toward the [param target] as per [method look_at]. Operations take place in global space.
*/
func (self Go) LookAtFromPosition(position gd.Vector3, target gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).LookAtFromPosition(position, target, gd.Vector3{0, 1, 0}, false)
}

/*
Transforms [param global_point] from world space to this node's local space.
*/
func (self Go) ToLocal(global_point gd.Vector3) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).ToLocal(global_point))
}

/*
Transforms [param local_point] from this node's local space to world space.
*/
func (self Go) ToGlobal(local_point gd.Vector3) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).ToGlobal(local_point))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Node3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Node3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Transform() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform3D(class(self).GetTransform())
}

func (self Go) SetTransform(value gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTransform(value)
}

func (self Go) GlobalTransform() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform3D(class(self).GetGlobalTransform())
}

func (self Go) SetGlobalTransform(value gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlobalTransform(value)
}

func (self Go) Position() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetPosition())
}

func (self Go) SetPosition(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPosition(value)
}

func (self Go) Rotation() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetRotation())
}

func (self Go) SetRotation(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRotation(value)
}

func (self Go) RotationDegrees() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetRotationDegrees())
}

func (self Go) SetRotationDegrees(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRotationDegrees(value)
}

func (self Go) Quaternion() gd.Quaternion {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Quaternion(class(self).GetQuaternion())
}

func (self Go) SetQuaternion(value gd.Quaternion) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetQuaternion(value)
}

func (self Go) Basis() gd.Basis {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Basis(class(self).GetBasis())
}

func (self Go) SetBasis(value gd.Basis) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBasis(value)
}

func (self Go) Scale() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetScale())
}

func (self Go) SetScale(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScale(value)
}

func (self Go) RotationEditMode() classdb.Node3DRotationEditMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Node3DRotationEditMode(class(self).GetRotationEditMode())
}

func (self Go) SetRotationEditMode(value classdb.Node3DRotationEditMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRotationEditMode(value)
}

func (self Go) RotationOrder() gd.EulerOrder {
	gc := gd.GarbageCollector(); _ = gc
		return gd.EulerOrder(class(self).GetRotationOrder())
}

func (self Go) SetRotationOrder(value gd.EulerOrder) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRotationOrder(value)
}

func (self Go) TopLevel() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSetAsTopLevel())
}

func (self Go) SetTopLevel(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAsTopLevel(value)
}

func (self Go) GlobalPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetGlobalPosition())
}

func (self Go) SetGlobalPosition(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlobalPosition(value)
}

func (self Go) GlobalBasis() gd.Basis {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Basis(class(self).GetGlobalBasis())
}

func (self Go) SetGlobalBasis(value gd.Basis) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlobalBasis(value)
}

func (self Go) GlobalRotation() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetGlobalRotation())
}

func (self Go) SetGlobalRotation(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlobalRotation(value)
}

func (self Go) GlobalRotationDegrees() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetGlobalRotationDegrees())
}

func (self Go) SetGlobalRotationDegrees(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGlobalRotationDegrees(value)
}

func (self Go) Visible() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVisible())
}

func (self Go) SetVisible(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisible(value)
}

func (self Go) VisibilityParent() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetVisibilityParent(gc).String())
}

func (self Go) SetVisibilityParent(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityParent(gc.String(value).NodePath(gc))
}

//go:nosplit
func (self class) SetTransform(local gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransform() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPosition(position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRotation(euler_radians gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, euler_radians)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRotation() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRotationDegrees(euler_degrees gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, euler_degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRotationDegrees() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRotationOrder(order gd.EulerOrder)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_rotation_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRotationOrder() gd.EulerOrder {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EulerOrder](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_rotation_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRotationEditMode(edit_mode classdb.Node3DRotationEditMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, edit_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_rotation_edit_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRotationEditMode() classdb.Node3DRotationEditMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Node3DRotationEditMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_rotation_edit_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScale(scale gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScale() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetQuaternion(quaternion gd.Quaternion)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, quaternion)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_quaternion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetQuaternion() gd.Quaternion {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_quaternion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBasis(basis gd.Basis)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, basis)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBasis() gd.Basis {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalTransform(global gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_global_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalTransform() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_global_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalPosition(position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalBasis(basis gd.Basis)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, basis)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_global_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalBasis() gd.Basis {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_global_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalRotation(euler_radians gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, euler_radians)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_global_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalRotation() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_global_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalRotationDegrees(euler_degrees gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, euler_degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_global_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalRotationDegrees() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_global_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the parent [Node3D], or [code]null[/code] if no parent exists, the parent is not of type [Node3D], or [member top_level] is [code]true[/code].
[b]Note:[/b] Calling this method is not equivalent to [code]get_parent() as Node3D[/code], which does not take [member top_level] into account.
*/
//go:nosplit
func (self class) GetParentNode3d(ctx gd.Lifetime) gdclass.Node3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_parent_node_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node3D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets whether the node ignores notification that its transformation (global or local) changed.
*/
//go:nosplit
func (self class) SetIgnoreTransformNotification(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_ignore_transform_notification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAsTopLevel(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_as_top_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSetAsTopLevel() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_is_set_as_top_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the node uses a scale of [code](1, 1, 1)[/code] or its local transformation scale. Changes to the local transformation scale are preserved.
*/
//go:nosplit
func (self class) SetDisableScale(disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_disable_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether this node uses a scale of [code](1, 1, 1)[/code] or its local transformation scale.
*/
//go:nosplit
func (self class) IsScaleDisabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_is_scale_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current [World3D] resource this [Node3D] node is registered to.
*/
//go:nosplit
func (self class) GetWorld3d(ctx gd.Lifetime) gdclass.World3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.World3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Forces the transform to update. Transform changes in physics are not instant for performance reasons. Transforms are accumulated and then set. Use this if you need an up-to-date transform when doing physics operations.
*/
//go:nosplit
func (self class) ForceUpdateTransform()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_force_update_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVisibilityParent(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_visibility_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityParent(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_visibility_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Updates all the [Node3D] gizmos attached to this node.
*/
//go:nosplit
func (self class) UpdateGizmos()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_update_gizmos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Attach an editor gizmo to this [Node3D].
[b]Note:[/b] The gizmo object would typically be an instance of [EditorNode3DGizmo], but the argument type is kept generic to avoid creating a dependency on editor classes in [Node3D].
*/
//go:nosplit
func (self class) AddGizmo(gizmo gdclass.Node3DGizmo)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(gizmo[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_add_gizmo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns all the gizmos attached to this [Node3D].
*/
//go:nosplit
func (self class) GetGizmos(ctx gd.Lifetime) gd.ArrayOf[gdclass.Node3DGizmo] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_get_gizmos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Node3DGizmo](ret)
}
/*
Clear all gizmos attached to this [Node3D].
*/
//go:nosplit
func (self class) ClearGizmos()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_clear_gizmos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set subgizmo selection for this node in the editor.
[b]Note:[/b] The gizmo object would typically be an instance of [EditorNode3DGizmo], but the argument type is kept generic to avoid creating a dependency on editor classes in [Node3D].
*/
//go:nosplit
func (self class) SetSubgizmoSelection(gizmo gdclass.Node3DGizmo, id gd.Int, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(gizmo[0].AsPointer())[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_subgizmo_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears subgizmo selection for this node in the editor. Useful when subgizmo IDs become invalid after a property change.
*/
//go:nosplit
func (self class) ClearSubgizmoSelection()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_clear_subgizmo_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the node is present in the [SceneTree], its [member visible] property is [code]true[/code] and all its ancestors are also visible. If any ancestor is hidden, this node will not be visible in the scene tree.
*/
//go:nosplit
func (self class) IsVisibleInTree() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_is_visible_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Enables rendering of this node. Changes [member visible] to [code]true[/code].
*/
//go:nosplit
func (self class) Show()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_show, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disables rendering of this node. Changes [member visible] to [code]false[/code].
*/
//go:nosplit
func (self class) Hide()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_hide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether the node notifies about its local transformation changes. [Node3D] will not propagate this by default.
*/
//go:nosplit
func (self class) SetNotifyLocalTransform(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_notify_local_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether node notifies about its local transformation changes. [Node3D] will not propagate this by default.
*/
//go:nosplit
func (self class) IsLocalTransformNotificationEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_is_local_transform_notification_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the node notifies about its global and local transformation changes. [Node3D] will not propagate this by default, unless it is in the editor context and it has a valid gizmo.
*/
//go:nosplit
func (self class) SetNotifyTransform(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_notify_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the node notifies about its global and local transformation changes. [Node3D] will not propagate this by default.
*/
//go:nosplit
func (self class) IsTransformNotificationEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_is_transform_notification_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians.
*/
//go:nosplit
func (self class) Rotate(axis gd.Vector3, angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_rotate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Rotates the global (world) transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in global coordinate system.
*/
//go:nosplit
func (self class) GlobalRotate(axis gd.Vector3, angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_global_rotate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Scales the global (world) transformation by the given [Vector3] scale factors.
*/
//go:nosplit
func (self class) GlobalScale(scale gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_global_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the global (world) transformation by [Vector3] offset. The offset is in global coordinate system.
*/
//go:nosplit
func (self class) GlobalTranslate(offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_global_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in object-local coordinate system.
*/
//go:nosplit
func (self class) RotateObjectLocal(axis gd.Vector3, angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_rotate_object_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Scales the local transformation by given 3D scale factors in object-local coordinate system.
*/
//go:nosplit
func (self class) ScaleObjectLocal(scale gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_scale_object_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes the node's position by the given offset [Vector3] in local space.
*/
//go:nosplit
func (self class) TranslateObjectLocal(offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_translate_object_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Rotates the local transformation around the X axis by angle in radians.
*/
//go:nosplit
func (self class) RotateX(angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_rotate_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Rotates the local transformation around the Y axis by angle in radians.
*/
//go:nosplit
func (self class) RotateY(angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_rotate_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Rotates the local transformation around the Z axis by angle in radians.
*/
//go:nosplit
func (self class) RotateZ(angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_rotate_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes the node's position by the given offset [Vector3].
Note that the translation [param offset] is affected by the node's scale, so if scaled by e.g. [code](10, 1, 1)[/code], a translation by an offset of [code](2, 0, 0)[/code] would actually add 20 ([code]2 * 10[/code]) to the X coordinate.
*/
//go:nosplit
func (self class) Translate(offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Resets this node's transformations (like scale, skew and taper) preserving its rotation and translation by performing Gram-Schmidt orthonormalization on this node's [Transform3D].
*/
//go:nosplit
func (self class) Orthonormalize()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_orthonormalize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Reset all transformations for this node (sets its [Transform3D] to the identity matrix).
*/
//go:nosplit
func (self class) SetIdentity()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_set_identity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Rotates the node so that the local forward axis (-Z, [constant Vector3.FORWARD]) points toward the [param target] position.
The local up axis (+Y) points as close to the [param up] vector as possible while staying perpendicular to the local forward axis. The resulting transform is orthogonal, and the scale is preserved. Non-uniform scaling may not work correctly.
The [param target] position cannot be the same as the node's position, the [param up] vector cannot be zero, and the direction from the node's position to the [param target] vector cannot be parallel to the [param up] vector.
Operations take place in global space, which means that the node must be in the scene tree.
If [param use_model_front] is [code]true[/code], the +Z axis (asset front) is treated as forward (implies +X is left) and points toward the [param target] position. By default, the -Z axis (camera forward) is treated as forward (implies +X is right).
*/
//go:nosplit
func (self class) LookAt(target gd.Vector3, up gd.Vector3, use_model_front bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, target)
	callframe.Arg(frame, up)
	callframe.Arg(frame, use_model_front)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_look_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the node to the specified [param position], and then rotates the node to point toward the [param target] as per [method look_at]. Operations take place in global space.
*/
//go:nosplit
func (self class) LookAtFromPosition(position gd.Vector3, target gd.Vector3, up gd.Vector3, use_model_front bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, target)
	callframe.Arg(frame, up)
	callframe.Arg(frame, use_model_front)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_look_at_from_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Transforms [param global_point] from world space to this node's local space.
*/
//go:nosplit
func (self class) ToLocal(global_point gd.Vector3) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, global_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_to_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Transforms [param local_point] from this node's local space to world space.
*/
//go:nosplit
func (self class) ToGlobal(local_point gd.Vector3) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Node3D.Bind_to_global, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnVisibilityChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("visibility_changed"), gc.Callable(cb), 0)
}


func (self class) AsNode3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("Node3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type RotationEditMode = classdb.Node3DRotationEditMode

const (
/*The rotation is edited using [Vector3] Euler angles.*/
	RotationEditModeEuler RotationEditMode = 0
/*The rotation is edited using a [Quaternion].*/
	RotationEditModeQuaternion RotationEditMode = 1
/*The rotation is edited using a [Basis]. In this mode, [member scale] can't be edited separately.*/
	RotationEditModeBasis RotationEditMode = 2
)
