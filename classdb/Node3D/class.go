// Package Node3D provides methods for working with Node3D object instances.
package Node3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Quaternion"
import "graphics.gd/variant/Basis"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Most basic 3D game object, with a [Transform3D] and visibility settings. All other 3D game objects inherit from [Node3D]. Use [Node3D] as a parent node to move, scale, rotate and show/hide children in a 3D project.
Affine operations (rotate, scale, translate) happen in parent's local coordinate system, unless the [Node3D] object is set as top-level. Affine operations in this coordinate system correspond to direct affine operations on the [Node3D]'s transform. The word local below refers to this coordinate system. The coordinate system that is attached to the [Node3D] object itself is referred to as object-local coordinate system.
[b]Note:[/b] Unless otherwise specified, all methods that have angle parameters must have angles specified as [i]radians[/i]. To convert degrees to radians, use [method @GlobalScope.deg_to_rad].
[b]Note:[/b] Be aware that "Spatial" nodes are now called "Node3D" starting with Godot 4. Any Godot 3.x references to "Spatial" nodes refer to "Node3D" in Godot 4.
*/
type Instance [1]gdclass.Node3D
type Any interface {
	gd.IsClass
	AsNode3D() Instance
}

/*
Returns the parent [Node3D], or [code]null[/code] if no parent exists, the parent is not of type [Node3D], or [member top_level] is [code]true[/code].
[b]Note:[/b] Calling this method is not equivalent to [code]get_parent() as Node3D[/code], which does not take [member top_level] into account.
*/
func (self Instance) GetParentNode3d() [1]gdclass.Node3D {
	return [1]gdclass.Node3D(class(self).GetParentNode3d())
}

/*
Sets whether the node ignores notification that its transformation (global or local) changed.
*/
func (self Instance) SetIgnoreTransformNotification(enabled bool) {
	class(self).SetIgnoreTransformNotification(enabled)
}

/*
Sets whether the node uses a scale of [code](1, 1, 1)[/code] or its local transformation scale. Changes to the local transformation scale are preserved.
*/
func (self Instance) SetDisableScale(disable bool) {
	class(self).SetDisableScale(disable)
}

/*
Returns whether this node uses a scale of [code](1, 1, 1)[/code] or its local transformation scale.
*/
func (self Instance) IsScaleDisabled() bool {
	return bool(class(self).IsScaleDisabled())
}

/*
Returns the current [World3D] resource this [Node3D] node is registered to.
*/
func (self Instance) GetWorld3d() [1]gdclass.World3D {
	return [1]gdclass.World3D(class(self).GetWorld3d())
}

/*
Forces the transform to update. Transform changes in physics are not instant for performance reasons. Transforms are accumulated and then set. Use this if you need an up-to-date transform when doing physics operations.
*/
func (self Instance) ForceUpdateTransform() {
	class(self).ForceUpdateTransform()
}

/*
Updates all the [Node3D] gizmos attached to this node.
*/
func (self Instance) UpdateGizmos() {
	class(self).UpdateGizmos()
}

/*
Attach an editor gizmo to this [Node3D].
[b]Note:[/b] The gizmo object would typically be an instance of [EditorNode3DGizmo], but the argument type is kept generic to avoid creating a dependency on editor classes in [Node3D].
*/
func (self Instance) AddGizmo(gizmo [1]gdclass.Node3DGizmo) {
	class(self).AddGizmo(gizmo)
}

/*
Returns all the gizmos attached to this [Node3D].
*/
func (self Instance) GetGizmos() gd.Array {
	return gd.Array(class(self).GetGizmos())
}

/*
Clear all gizmos attached to this [Node3D].
*/
func (self Instance) ClearGizmos() {
	class(self).ClearGizmos()
}

/*
Set subgizmo selection for this node in the editor.
[b]Note:[/b] The gizmo object would typically be an instance of [EditorNode3DGizmo], but the argument type is kept generic to avoid creating a dependency on editor classes in [Node3D].
*/
func (self Instance) SetSubgizmoSelection(gizmo [1]gdclass.Node3DGizmo, id int, transform Transform3D.BasisOrigin) {
	class(self).SetSubgizmoSelection(gizmo, gd.Int(id), gd.Transform3D(transform))
}

/*
Clears subgizmo selection for this node in the editor. Useful when subgizmo IDs become invalid after a property change.
*/
func (self Instance) ClearSubgizmoSelection() {
	class(self).ClearSubgizmoSelection()
}

/*
Returns [code]true[/code] if the node is present in the [SceneTree], its [member visible] property is [code]true[/code] and all its ancestors are also visible. If any ancestor is hidden, this node will not be visible in the scene tree.
*/
func (self Instance) IsVisibleInTree() bool {
	return bool(class(self).IsVisibleInTree())
}

/*
Enables rendering of this node. Changes [member visible] to [code]true[/code].
*/
func (self Instance) Show() {
	class(self).Show()
}

/*
Disables rendering of this node. Changes [member visible] to [code]false[/code].
*/
func (self Instance) Hide() {
	class(self).Hide()
}

/*
Sets whether the node notifies about its local transformation changes. [Node3D] will not propagate this by default.
*/
func (self Instance) SetNotifyLocalTransform(enable bool) {
	class(self).SetNotifyLocalTransform(enable)
}

/*
Returns whether node notifies about its local transformation changes. [Node3D] will not propagate this by default.
*/
func (self Instance) IsLocalTransformNotificationEnabled() bool {
	return bool(class(self).IsLocalTransformNotificationEnabled())
}

/*
Sets whether the node notifies about its global and local transformation changes. [Node3D] will not propagate this by default, unless it is in the editor context and it has a valid gizmo.
*/
func (self Instance) SetNotifyTransform(enable bool) {
	class(self).SetNotifyTransform(enable)
}

/*
Returns whether the node notifies about its global and local transformation changes. [Node3D] will not propagate this by default.
*/
func (self Instance) IsTransformNotificationEnabled() bool {
	return bool(class(self).IsTransformNotificationEnabled())
}

/*
Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians.
*/
func (self Instance) Rotate(axis Vector3.XYZ, angle Float.X) {
	class(self).Rotate(gd.Vector3(axis), gd.Float(angle))
}

/*
Rotates the global (world) transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in global coordinate system.
*/
func (self Instance) GlobalRotate(axis Vector3.XYZ, angle Float.X) {
	class(self).GlobalRotate(gd.Vector3(axis), gd.Float(angle))
}

/*
Scales the global (world) transformation by the given [Vector3] scale factors.
*/
func (self Instance) GlobalScale(scale Vector3.XYZ) {
	class(self).GlobalScale(gd.Vector3(scale))
}

/*
Moves the global (world) transformation by [Vector3] offset. The offset is in global coordinate system.
*/
func (self Instance) GlobalTranslate(offset Vector3.XYZ) {
	class(self).GlobalTranslate(gd.Vector3(offset))
}

/*
Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in object-local coordinate system.
*/
func (self Instance) RotateObjectLocal(axis Vector3.XYZ, angle Float.X) {
	class(self).RotateObjectLocal(gd.Vector3(axis), gd.Float(angle))
}

/*
Scales the local transformation by given 3D scale factors in object-local coordinate system.
*/
func (self Instance) ScaleObjectLocal(scale Vector3.XYZ) {
	class(self).ScaleObjectLocal(gd.Vector3(scale))
}

/*
Changes the node's position by the given offset [Vector3] in local space.
*/
func (self Instance) TranslateObjectLocal(offset Vector3.XYZ) {
	class(self).TranslateObjectLocal(gd.Vector3(offset))
}

/*
Rotates the local transformation around the X axis by angle in radians.
*/
func (self Instance) RotateX(angle Float.X) {
	class(self).RotateX(gd.Float(angle))
}

/*
Rotates the local transformation around the Y axis by angle in radians.
*/
func (self Instance) RotateY(angle Float.X) {
	class(self).RotateY(gd.Float(angle))
}

/*
Rotates the local transformation around the Z axis by angle in radians.
*/
func (self Instance) RotateZ(angle Float.X) {
	class(self).RotateZ(gd.Float(angle))
}

/*
Changes the node's position by the given offset [Vector3].
Note that the translation [param offset] is affected by the node's scale, so if scaled by e.g. [code](10, 1, 1)[/code], a translation by an offset of [code](2, 0, 0)[/code] would actually add 20 ([code]2 * 10[/code]) to the X coordinate.
*/
func (self Instance) Translate(offset Vector3.XYZ) {
	class(self).Translate(gd.Vector3(offset))
}

/*
Resets this node's transformations (like scale, skew and taper) preserving its rotation and translation by performing Gram-Schmidt orthonormalization on this node's [Transform3D].
*/
func (self Instance) Orthonormalize() {
	class(self).Orthonormalize()
}

/*
Reset all transformations for this node (sets its [Transform3D] to the identity matrix).
*/
func (self Instance) SetIdentity() {
	class(self).SetIdentity()
}

/*
Rotates the node so that the local forward axis (-Z, [constant Vector3.FORWARD]) points toward the [param target] position.
The local up axis (+Y) points as close to the [param up] vector as possible while staying perpendicular to the local forward axis. The resulting transform is orthogonal, and the scale is preserved. Non-uniform scaling may not work correctly.
The [param target] position cannot be the same as the node's position, the [param up] vector cannot be zero, and the direction from the node's position to the [param target] vector cannot be parallel to the [param up] vector.
Operations take place in global space, which means that the node must be in the scene tree.
If [param use_model_front] is [code]true[/code], the +Z axis (asset front) is treated as forward (implies +X is left) and points toward the [param target] position. By default, the -Z axis (camera forward) is treated as forward (implies +X is right).
*/
func (self Instance) LookAt(target Vector3.XYZ) {
	class(self).LookAt(gd.Vector3(target), gd.Vector3(gd.Vector3{0, 1, 0}), false)
}

/*
Moves the node to the specified [param position], and then rotates the node to point toward the [param target] as per [method look_at]. Operations take place in global space.
*/
func (self Instance) LookAtFromPosition(position Vector3.XYZ, target Vector3.XYZ) {
	class(self).LookAtFromPosition(gd.Vector3(position), gd.Vector3(target), gd.Vector3(gd.Vector3{0, 1, 0}), false)
}

/*
Transforms [param global_point] from world space to this node's local space.
*/
func (self Instance) ToLocal(global_point Vector3.XYZ) Vector3.XYZ {
	return Vector3.XYZ(class(self).ToLocal(gd.Vector3(global_point)))
}

/*
Transforms [param local_point] from this node's local space to world space.
*/
func (self Instance) ToGlobal(local_point Vector3.XYZ) Vector3.XYZ {
	return Vector3.XYZ(class(self).ToGlobal(gd.Vector3(local_point)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Node3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Node3D"))
	return Instance{*(*gdclass.Node3D)(unsafe.Pointer(&object))}
}

func (self Instance) Transform() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetTransform())
}

func (self Instance) SetTransform(value Transform3D.BasisOrigin) {
	class(self).SetTransform(gd.Transform3D(value))
}

func (self Instance) GlobalTransform() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetGlobalTransform())
}

func (self Instance) SetGlobalTransform(value Transform3D.BasisOrigin) {
	class(self).SetGlobalTransform(gd.Transform3D(value))
}

func (self Instance) Position() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetPosition())
}

func (self Instance) SetPosition(value Vector3.XYZ) {
	class(self).SetPosition(gd.Vector3(value))
}

func (self Instance) Rotation() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetRotation())
}

func (self Instance) SetRotation(value Vector3.XYZ) {
	class(self).SetRotation(gd.Vector3(value))
}

func (self Instance) RotationDegrees() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetRotationDegrees())
}

func (self Instance) SetRotationDegrees(value Vector3.XYZ) {
	class(self).SetRotationDegrees(gd.Vector3(value))
}

func (self Instance) Quaternion() Quaternion.IJKX {
	return Quaternion.IJKX(class(self).GetQuaternion())
}

func (self Instance) SetQuaternion(value Quaternion.IJKX) {
	class(self).SetQuaternion(gd.Quaternion(value))
}

func (self Instance) Basis() Basis.XYZ {
	return Basis.XYZ(class(self).GetBasis())
}

func (self Instance) SetBasis(value Basis.XYZ) {
	class(self).SetBasis(gd.Basis(value))
}

func (self Instance) Scale() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetScale())
}

func (self Instance) SetScale(value Vector3.XYZ) {
	class(self).SetScale(gd.Vector3(value))
}

func (self Instance) RotationEditMode() gdclass.Node3DRotationEditMode {
	return gdclass.Node3DRotationEditMode(class(self).GetRotationEditMode())
}

func (self Instance) SetRotationEditMode(value gdclass.Node3DRotationEditMode) {
	class(self).SetRotationEditMode(value)
}

func (self Instance) RotationOrder() EulerOrder {
	return EulerOrder(class(self).GetRotationOrder())
}

func (self Instance) SetRotationOrder(value EulerOrder) {
	class(self).SetRotationOrder(value)
}

func (self Instance) TopLevel() bool {
	return bool(class(self).IsSetAsTopLevel())
}

func (self Instance) SetTopLevel(value bool) {
	class(self).SetAsTopLevel(value)
}

func (self Instance) GlobalPosition() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetGlobalPosition())
}

func (self Instance) SetGlobalPosition(value Vector3.XYZ) {
	class(self).SetGlobalPosition(gd.Vector3(value))
}

func (self Instance) GlobalBasis() Basis.XYZ {
	return Basis.XYZ(class(self).GetGlobalBasis())
}

func (self Instance) SetGlobalBasis(value Basis.XYZ) {
	class(self).SetGlobalBasis(gd.Basis(value))
}

func (self Instance) GlobalRotation() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetGlobalRotation())
}

func (self Instance) SetGlobalRotation(value Vector3.XYZ) {
	class(self).SetGlobalRotation(gd.Vector3(value))
}

func (self Instance) GlobalRotationDegrees() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetGlobalRotationDegrees())
}

func (self Instance) SetGlobalRotationDegrees(value Vector3.XYZ) {
	class(self).SetGlobalRotationDegrees(gd.Vector3(value))
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisible(value)
}

func (self Instance) VisibilityParent() NodePath.String {
	return NodePath.String(class(self).GetVisibilityParent().String())
}

func (self Instance) SetVisibilityParent(value NodePath.String) {
	class(self).SetVisibilityParent(gd.NewString(string(value)).NodePath())
}

//go:nosplit
func (self class) SetTransform(local gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, local)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransform() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPosition(position gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPosition() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotation(euler_radians gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, euler_radians)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRotation() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotationDegrees(euler_degrees gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, euler_degrees)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRotationDegrees() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotationOrder(order EulerOrder) {
	var frame = callframe.New()
	callframe.Arg(frame, order)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_rotation_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRotationOrder() EulerOrder {
	var frame = callframe.New()
	var r_ret = callframe.Ret[EulerOrder](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_rotation_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotationEditMode(edit_mode gdclass.Node3DRotationEditMode) {
	var frame = callframe.New()
	callframe.Arg(frame, edit_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_rotation_edit_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRotationEditMode() gdclass.Node3DRotationEditMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Node3DRotationEditMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_rotation_edit_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScale(scale gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScale() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetQuaternion(quaternion gd.Quaternion) {
	var frame = callframe.New()
	callframe.Arg(frame, quaternion)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_quaternion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetQuaternion() gd.Quaternion {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_quaternion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBasis(basis gd.Basis) {
	var frame = callframe.New()
	callframe.Arg(frame, basis)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBasis() gd.Basis {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalTransform(global gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_global_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalTransform() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_global_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalPosition(position gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalPosition() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalBasis(basis gd.Basis) {
	var frame = callframe.New()
	callframe.Arg(frame, basis)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_global_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalBasis() gd.Basis {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_global_basis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalRotation(euler_radians gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, euler_radians)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_global_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalRotation() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_global_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalRotationDegrees(euler_degrees gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, euler_degrees)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_global_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalRotationDegrees() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_global_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the parent [Node3D], or [code]null[/code] if no parent exists, the parent is not of type [Node3D], or [member top_level] is [code]true[/code].
[b]Note:[/b] Calling this method is not equivalent to [code]get_parent() as Node3D[/code], which does not take [member top_level] into account.
*/
//go:nosplit
func (self class) GetParentNode3d() [1]gdclass.Node3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_parent_node_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Node3D{gd.PointerMustAssertInstanceID[gdclass.Node3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets whether the node ignores notification that its transformation (global or local) changed.
*/
//go:nosplit
func (self class) SetIgnoreTransformNotification(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_ignore_transform_notification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetAsTopLevel(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_as_top_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSetAsTopLevel() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_is_set_as_top_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the node uses a scale of [code](1, 1, 1)[/code] or its local transformation scale. Changes to the local transformation scale are preserved.
*/
//go:nosplit
func (self class) SetDisableScale(disable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_disable_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether this node uses a scale of [code](1, 1, 1)[/code] or its local transformation scale.
*/
//go:nosplit
func (self class) IsScaleDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_is_scale_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current [World3D] resource this [Node3D] node is registered to.
*/
//go:nosplit
func (self class) GetWorld3d() [1]gdclass.World3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.World3D{gd.PointerWithOwnershipTransferredToGo[gdclass.World3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Forces the transform to update. Transform changes in physics are not instant for performance reasons. Transforms are accumulated and then set. Use this if you need an up-to-date transform when doing physics operations.
*/
//go:nosplit
func (self class) ForceUpdateTransform() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_force_update_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibilityParent(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_visibility_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityParent() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_visibility_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Updates all the [Node3D] gizmos attached to this node.
*/
//go:nosplit
func (self class) UpdateGizmos() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_update_gizmos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Attach an editor gizmo to this [Node3D].
[b]Note:[/b] The gizmo object would typically be an instance of [EditorNode3DGizmo], but the argument type is kept generic to avoid creating a dependency on editor classes in [Node3D].
*/
//go:nosplit
func (self class) AddGizmo(gizmo [1]gdclass.Node3DGizmo) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gizmo[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_add_gizmo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns all the gizmos attached to this [Node3D].
*/
//go:nosplit
func (self class) GetGizmos() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_get_gizmos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clear all gizmos attached to this [Node3D].
*/
//go:nosplit
func (self class) ClearGizmos() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_clear_gizmos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Set subgizmo selection for this node in the editor.
[b]Note:[/b] The gizmo object would typically be an instance of [EditorNode3DGizmo], but the argument type is kept generic to avoid creating a dependency on editor classes in [Node3D].
*/
//go:nosplit
func (self class) SetSubgizmoSelection(gizmo [1]gdclass.Node3DGizmo, id gd.Int, transform gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gizmo[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_subgizmo_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears subgizmo selection for this node in the editor. Useful when subgizmo IDs become invalid after a property change.
*/
//go:nosplit
func (self class) ClearSubgizmoSelection() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_clear_subgizmo_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisible(visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the node is present in the [SceneTree], its [member visible] property is [code]true[/code] and all its ancestors are also visible. If any ancestor is hidden, this node will not be visible in the scene tree.
*/
//go:nosplit
func (self class) IsVisibleInTree() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_is_visible_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Enables rendering of this node. Changes [member visible] to [code]true[/code].
*/
//go:nosplit
func (self class) Show() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_show, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Disables rendering of this node. Changes [member visible] to [code]false[/code].
*/
//go:nosplit
func (self class) Hide() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_hide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets whether the node notifies about its local transformation changes. [Node3D] will not propagate this by default.
*/
//go:nosplit
func (self class) SetNotifyLocalTransform(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_notify_local_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether node notifies about its local transformation changes. [Node3D] will not propagate this by default.
*/
//go:nosplit
func (self class) IsLocalTransformNotificationEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_is_local_transform_notification_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the node notifies about its global and local transformation changes. [Node3D] will not propagate this by default, unless it is in the editor context and it has a valid gizmo.
*/
//go:nosplit
func (self class) SetNotifyTransform(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_notify_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the node notifies about its global and local transformation changes. [Node3D] will not propagate this by default.
*/
//go:nosplit
func (self class) IsTransformNotificationEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_is_transform_notification_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians.
*/
//go:nosplit
func (self class) Rotate(axis gd.Vector3, angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_rotate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Rotates the global (world) transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in global coordinate system.
*/
//go:nosplit
func (self class) GlobalRotate(axis gd.Vector3, angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_global_rotate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Scales the global (world) transformation by the given [Vector3] scale factors.
*/
//go:nosplit
func (self class) GlobalScale(scale gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_global_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Moves the global (world) transformation by [Vector3] offset. The offset is in global coordinate system.
*/
//go:nosplit
func (self class) GlobalTranslate(offset gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_global_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Rotates the local transformation around axis, a unit [Vector3], by specified angle in radians. The rotation axis is in object-local coordinate system.
*/
//go:nosplit
func (self class) RotateObjectLocal(axis gd.Vector3, angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_rotate_object_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Scales the local transformation by given 3D scale factors in object-local coordinate system.
*/
//go:nosplit
func (self class) ScaleObjectLocal(scale gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_scale_object_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Changes the node's position by the given offset [Vector3] in local space.
*/
//go:nosplit
func (self class) TranslateObjectLocal(offset gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_translate_object_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Rotates the local transformation around the X axis by angle in radians.
*/
//go:nosplit
func (self class) RotateX(angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_rotate_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Rotates the local transformation around the Y axis by angle in radians.
*/
//go:nosplit
func (self class) RotateY(angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_rotate_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Rotates the local transformation around the Z axis by angle in radians.
*/
//go:nosplit
func (self class) RotateZ(angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_rotate_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Changes the node's position by the given offset [Vector3].
Note that the translation [param offset] is affected by the node's scale, so if scaled by e.g. [code](10, 1, 1)[/code], a translation by an offset of [code](2, 0, 0)[/code] would actually add 20 ([code]2 * 10[/code]) to the X coordinate.
*/
//go:nosplit
func (self class) Translate(offset gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Resets this node's transformations (like scale, skew and taper) preserving its rotation and translation by performing Gram-Schmidt orthonormalization on this node's [Transform3D].
*/
//go:nosplit
func (self class) Orthonormalize() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_orthonormalize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Reset all transformations for this node (sets its [Transform3D] to the identity matrix).
*/
//go:nosplit
func (self class) SetIdentity() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_set_identity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) LookAt(target gd.Vector3, up gd.Vector3, use_model_front bool) {
	var frame = callframe.New()
	callframe.Arg(frame, target)
	callframe.Arg(frame, up)
	callframe.Arg(frame, use_model_front)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_look_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Moves the node to the specified [param position], and then rotates the node to point toward the [param target] as per [method look_at]. Operations take place in global space.
*/
//go:nosplit
func (self class) LookAtFromPosition(position gd.Vector3, target gd.Vector3, up gd.Vector3, use_model_front bool) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, target)
	callframe.Arg(frame, up)
	callframe.Arg(frame, use_model_front)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_look_at_from_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Transforms [param global_point] from world space to this node's local space.
*/
//go:nosplit
func (self class) ToLocal(global_point gd.Vector3) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, global_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_to_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Transforms [param local_point] from this node's local space to world space.
*/
//go:nosplit
func (self class) ToGlobal(local_point gd.Vector3) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, local_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node3D.Bind_to_global, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnVisibilityChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("visibility_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsNode3D() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {
	gdclass.Register("Node3D", func(ptr gd.Object) any { return [1]gdclass.Node3D{*(*gdclass.Node3D)(unsafe.Pointer(&ptr))} })
}

type RotationEditMode = gdclass.Node3DRotationEditMode

const (
	/*The rotation is edited using [Vector3] Euler angles.*/
	RotationEditModeEuler RotationEditMode = 0
	/*The rotation is edited using a [Quaternion].*/
	RotationEditModeQuaternion RotationEditMode = 1
	/*The rotation is edited using a [Basis]. In this mode, [member scale] can't be edited separately.*/
	RotationEditModeBasis RotationEditMode = 2
)

type EulerOrder int

const (
	/*Specifies that Euler angles should be in XYZ order. When composing, the order is X, Y, Z. When decomposing, the order is reversed, first Z, then Y, and X last.*/
	EulerOrderXyz EulerOrder = 0
	/*Specifies that Euler angles should be in XZY order. When composing, the order is X, Z, Y. When decomposing, the order is reversed, first Y, then Z, and X last.*/
	EulerOrderXzy EulerOrder = 1
	/*Specifies that Euler angles should be in YXZ order. When composing, the order is Y, X, Z. When decomposing, the order is reversed, first Z, then X, and Y last.*/
	EulerOrderYxz EulerOrder = 2
	/*Specifies that Euler angles should be in YZX order. When composing, the order is Y, Z, X. When decomposing, the order is reversed, first X, then Z, and Y last.*/
	EulerOrderYzx EulerOrder = 3
	/*Specifies that Euler angles should be in ZXY order. When composing, the order is Z, X, Y. When decomposing, the order is reversed, first Y, then X, and Z last.*/
	EulerOrderZxy EulerOrder = 4
	/*Specifies that Euler angles should be in ZYX order. When composing, the order is Z, Y, X. When decomposing, the order is reversed, first X, then Y, and Z last.*/
	EulerOrderZyx EulerOrder = 5
)
