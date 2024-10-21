package PhysicsServer3D

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
PhysicsServer3D is the server responsible for all 3D physics. It can directly create and manipulate all physics objects:
- A [i]space[/i] is a self-contained world for a physics simulation. It contains bodies, areas, and joints. Its state can be queried for collision and intersection information, and several parameters of the simulation can be modified.
- A [i]shape[/i] is a geometric shape such as a sphere, a box, a cylinder, or a polygon. It can be used for collision detection by adding it to a body/area, possibly with an extra transformation relative to the body/area's origin. Bodies/areas can have multiple (transformed) shapes added to them, and a single shape can be added to bodies/areas multiple times with different local transformations.
- A [i]body[/i] is a physical object which can be in static, kinematic, or rigid mode. Its state (such as position and velocity) can be queried and updated. A force integration callback can be set to customize the body's physics.
- An [i]area[/i] is a region in space which can be used to detect bodies and areas entering and exiting it. A body monitoring callback can be set to report entering/exiting body shapes, and similarly an area monitoring callback can be set. Gravity and damping can be overridden within the area by setting area parameters.
- A [i]joint[/i] is a constraint, either between two bodies or on one body relative to a point. Parameters such as the joint bias and the rest length of a spring joint can be adjusted.
Physics objects in [PhysicsServer3D] may be created and manipulated independently; they do not have to be tied to nodes in the scene tree.
[b]Note:[/b] All the 3D physics nodes use the physics server internally. Adding a physics node to the scene tree will cause a corresponding physics object to be created in the physics server. A rigid body node registers a callback that updates the node's transform with the transform of the respective body object in the physics server (every physics update). An area node registers a callback to inform the area node about overlaps with the respective area object in the physics server. The raycast node queries the direct state of the relevant space in the physics server.

*/
type Simple [1]classdb.PhysicsServer3D
func (self Simple) WorldBoundaryShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).WorldBoundaryShapeCreate())
}
func (self Simple) SeparationRayShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SeparationRayShapeCreate())
}
func (self Simple) SphereShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SphereShapeCreate())
}
func (self Simple) BoxShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).BoxShapeCreate())
}
func (self Simple) CapsuleShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CapsuleShapeCreate())
}
func (self Simple) CylinderShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CylinderShapeCreate())
}
func (self Simple) ConvexPolygonShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ConvexPolygonShapeCreate())
}
func (self Simple) ConcavePolygonShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ConcavePolygonShapeCreate())
}
func (self Simple) HeightmapShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).HeightmapShapeCreate())
}
func (self Simple) CustomShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CustomShapeCreate())
}
func (self Simple) ShapeSetData(shape gd.RID, data gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ShapeSetData(shape, data)
}
func (self Simple) ShapeSetMargin(shape gd.RID, margin float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ShapeSetMargin(shape, gd.Float(margin))
}
func (self Simple) ShapeGetType(shape gd.RID) classdb.PhysicsServer3DShapeType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicsServer3DShapeType(Expert(self).ShapeGetType(shape))
}
func (self Simple) ShapeGetData(shape gd.RID) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).ShapeGetData(gc, shape))
}
func (self Simple) ShapeGetMargin(shape gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ShapeGetMargin(shape)))
}
func (self Simple) SpaceCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SpaceCreate())
}
func (self Simple) SpaceSetActive(space gd.RID, active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SpaceSetActive(space, active)
}
func (self Simple) SpaceIsActive(space gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).SpaceIsActive(space))
}
func (self Simple) SpaceSetParam(space gd.RID, param classdb.PhysicsServer3DSpaceParameter, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SpaceSetParam(space, param, gd.Float(value))
}
func (self Simple) SpaceGetParam(space gd.RID, param classdb.PhysicsServer3DSpaceParameter) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).SpaceGetParam(space, param)))
}
func (self Simple) SpaceGetDirectState(space gd.RID) [1]classdb.PhysicsDirectSpaceState3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PhysicsDirectSpaceState3D(Expert(self).SpaceGetDirectState(gc, space))
}
func (self Simple) AreaCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).AreaCreate())
}
func (self Simple) AreaSetSpace(area gd.RID, space gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetSpace(area, space)
}
func (self Simple) AreaGetSpace(area gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).AreaGetSpace(area))
}
func (self Simple) AreaAddShape(area gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaAddShape(area, shape, transform, disabled)
}
func (self Simple) AreaSetShape(area gd.RID, shape_idx int, shape gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetShape(area, gd.Int(shape_idx), shape)
}
func (self Simple) AreaSetShapeTransform(area gd.RID, shape_idx int, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetShapeTransform(area, gd.Int(shape_idx), transform)
}
func (self Simple) AreaSetShapeDisabled(area gd.RID, shape_idx int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetShapeDisabled(area, gd.Int(shape_idx), disabled)
}
func (self Simple) AreaGetShapeCount(area gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AreaGetShapeCount(area)))
}
func (self Simple) AreaGetShape(area gd.RID, shape_idx int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).AreaGetShape(area, gd.Int(shape_idx)))
}
func (self Simple) AreaGetShapeTransform(area gd.RID, shape_idx int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).AreaGetShapeTransform(area, gd.Int(shape_idx)))
}
func (self Simple) AreaRemoveShape(area gd.RID, shape_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaRemoveShape(area, gd.Int(shape_idx))
}
func (self Simple) AreaClearShapes(area gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaClearShapes(area)
}
func (self Simple) AreaSetCollisionLayer(area gd.RID, layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetCollisionLayer(area, gd.Int(layer))
}
func (self Simple) AreaGetCollisionLayer(area gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AreaGetCollisionLayer(area)))
}
func (self Simple) AreaSetCollisionMask(area gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetCollisionMask(area, gd.Int(mask))
}
func (self Simple) AreaGetCollisionMask(area gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AreaGetCollisionMask(area)))
}
func (self Simple) AreaSetParam(area gd.RID, param classdb.PhysicsServer3DAreaParameter, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetParam(area, param, value)
}
func (self Simple) AreaSetTransform(area gd.RID, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetTransform(area, transform)
}
func (self Simple) AreaGetParam(area gd.RID, param classdb.PhysicsServer3DAreaParameter) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).AreaGetParam(gc, area, param))
}
func (self Simple) AreaGetTransform(area gd.RID) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).AreaGetTransform(area))
}
func (self Simple) AreaAttachObjectInstanceId(area gd.RID, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaAttachObjectInstanceId(area, gd.Int(id))
}
func (self Simple) AreaGetObjectInstanceId(area gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AreaGetObjectInstanceId(area)))
}
func (self Simple) AreaSetMonitorCallback(area gd.RID, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetMonitorCallback(area, callback)
}
func (self Simple) AreaSetAreaMonitorCallback(area gd.RID, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetAreaMonitorCallback(area, callback)
}
func (self Simple) AreaSetMonitorable(area gd.RID, monitorable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetMonitorable(area, monitorable)
}
func (self Simple) AreaSetRayPickable(area gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetRayPickable(area, enable)
}
func (self Simple) BodyCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).BodyCreate())
}
func (self Simple) BodySetSpace(body gd.RID, space gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetSpace(body, space)
}
func (self Simple) BodyGetSpace(body gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).BodyGetSpace(body))
}
func (self Simple) BodySetMode(body gd.RID, mode classdb.PhysicsServer3DBodyMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetMode(body, mode)
}
func (self Simple) BodyGetMode(body gd.RID) classdb.PhysicsServer3DBodyMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicsServer3DBodyMode(Expert(self).BodyGetMode(body))
}
func (self Simple) BodySetCollisionLayer(body gd.RID, layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetCollisionLayer(body, gd.Int(layer))
}
func (self Simple) BodyGetCollisionLayer(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).BodyGetCollisionLayer(body)))
}
func (self Simple) BodySetCollisionMask(body gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetCollisionMask(body, gd.Int(mask))
}
func (self Simple) BodyGetCollisionMask(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).BodyGetCollisionMask(body)))
}
func (self Simple) BodySetCollisionPriority(body gd.RID, priority float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetCollisionPriority(body, gd.Float(priority))
}
func (self Simple) BodyGetCollisionPriority(body gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).BodyGetCollisionPriority(body)))
}
func (self Simple) BodyAddShape(body gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAddShape(body, shape, transform, disabled)
}
func (self Simple) BodySetShape(body gd.RID, shape_idx int, shape gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetShape(body, gd.Int(shape_idx), shape)
}
func (self Simple) BodySetShapeTransform(body gd.RID, shape_idx int, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetShapeTransform(body, gd.Int(shape_idx), transform)
}
func (self Simple) BodySetShapeDisabled(body gd.RID, shape_idx int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetShapeDisabled(body, gd.Int(shape_idx), disabled)
}
func (self Simple) BodyGetShapeCount(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).BodyGetShapeCount(body)))
}
func (self Simple) BodyGetShape(body gd.RID, shape_idx int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).BodyGetShape(body, gd.Int(shape_idx)))
}
func (self Simple) BodyGetShapeTransform(body gd.RID, shape_idx int) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).BodyGetShapeTransform(body, gd.Int(shape_idx)))
}
func (self Simple) BodyRemoveShape(body gd.RID, shape_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyRemoveShape(body, gd.Int(shape_idx))
}
func (self Simple) BodyClearShapes(body gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyClearShapes(body)
}
func (self Simple) BodyAttachObjectInstanceId(body gd.RID, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAttachObjectInstanceId(body, gd.Int(id))
}
func (self Simple) BodyGetObjectInstanceId(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).BodyGetObjectInstanceId(body)))
}
func (self Simple) BodySetEnableContinuousCollisionDetection(body gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetEnableContinuousCollisionDetection(body, enable)
}
func (self Simple) BodyIsContinuousCollisionDetectionEnabled(body gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).BodyIsContinuousCollisionDetectionEnabled(body))
}
func (self Simple) BodySetParam(body gd.RID, param classdb.PhysicsServer3DBodyParameter, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetParam(body, param, value)
}
func (self Simple) BodyGetParam(body gd.RID, param classdb.PhysicsServer3DBodyParameter) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).BodyGetParam(gc, body, param))
}
func (self Simple) BodyResetMassProperties(body gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyResetMassProperties(body)
}
func (self Simple) BodySetState(body gd.RID, state classdb.PhysicsServer3DBodyState, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetState(body, state, value)
}
func (self Simple) BodyGetState(body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).BodyGetState(gc, body, state))
}
func (self Simple) BodyApplyCentralImpulse(body gd.RID, impulse gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyCentralImpulse(body, impulse)
}
func (self Simple) BodyApplyImpulse(body gd.RID, impulse gd.Vector3, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyImpulse(body, impulse, position)
}
func (self Simple) BodyApplyTorqueImpulse(body gd.RID, impulse gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyTorqueImpulse(body, impulse)
}
func (self Simple) BodyApplyCentralForce(body gd.RID, force gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyCentralForce(body, force)
}
func (self Simple) BodyApplyForce(body gd.RID, force gd.Vector3, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyForce(body, force, position)
}
func (self Simple) BodyApplyTorque(body gd.RID, torque gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyTorque(body, torque)
}
func (self Simple) BodyAddConstantCentralForce(body gd.RID, force gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAddConstantCentralForce(body, force)
}
func (self Simple) BodyAddConstantForce(body gd.RID, force gd.Vector3, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAddConstantForce(body, force, position)
}
func (self Simple) BodyAddConstantTorque(body gd.RID, torque gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAddConstantTorque(body, torque)
}
func (self Simple) BodySetConstantForce(body gd.RID, force gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetConstantForce(body, force)
}
func (self Simple) BodyGetConstantForce(body gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).BodyGetConstantForce(body))
}
func (self Simple) BodySetConstantTorque(body gd.RID, torque gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetConstantTorque(body, torque)
}
func (self Simple) BodyGetConstantTorque(body gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).BodyGetConstantTorque(body))
}
func (self Simple) BodySetAxisVelocity(body gd.RID, axis_velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetAxisVelocity(body, axis_velocity)
}
func (self Simple) BodySetAxisLock(body gd.RID, axis classdb.PhysicsServer3DBodyAxis, lock bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetAxisLock(body, axis, lock)
}
func (self Simple) BodyIsAxisLocked(body gd.RID, axis classdb.PhysicsServer3DBodyAxis) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).BodyIsAxisLocked(body, axis))
}
func (self Simple) BodyAddCollisionException(body gd.RID, excepted_body gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAddCollisionException(body, excepted_body)
}
func (self Simple) BodyRemoveCollisionException(body gd.RID, excepted_body gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyRemoveCollisionException(body, excepted_body)
}
func (self Simple) BodySetMaxContactsReported(body gd.RID, amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetMaxContactsReported(body, gd.Int(amount))
}
func (self Simple) BodyGetMaxContactsReported(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).BodyGetMaxContactsReported(body)))
}
func (self Simple) BodySetOmitForceIntegration(body gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetOmitForceIntegration(body, enable)
}
func (self Simple) BodyIsOmittingForceIntegration(body gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).BodyIsOmittingForceIntegration(body))
}
func (self Simple) BodySetStateSyncCallback(body gd.RID, callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetStateSyncCallback(body, callable)
}
func (self Simple) BodySetForceIntegrationCallback(body gd.RID, callable gd.Callable, userdata gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetForceIntegrationCallback(body, callable, userdata)
}
func (self Simple) BodySetRayPickable(body gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetRayPickable(body, enable)
}
func (self Simple) BodyTestMotion(body gd.RID, parameters [1]classdb.PhysicsTestMotionParameters3D, result [1]classdb.PhysicsTestMotionResult3D) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).BodyTestMotion(body, parameters, result))
}
func (self Simple) BodyGetDirectState(body gd.RID) [1]classdb.PhysicsDirectBodyState3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PhysicsDirectBodyState3D(Expert(self).BodyGetDirectState(gc, body))
}
func (self Simple) SoftBodyCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SoftBodyCreate())
}
func (self Simple) SoftBodyUpdateRenderingServer(body gd.RID, rendering_server_handler [1]classdb.PhysicsServer3DRenderingServerHandler) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodyUpdateRenderingServer(body, rendering_server_handler)
}
func (self Simple) SoftBodySetSpace(body gd.RID, space gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetSpace(body, space)
}
func (self Simple) SoftBodyGetSpace(body gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SoftBodyGetSpace(body))
}
func (self Simple) SoftBodySetMesh(body gd.RID, mesh gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetMesh(body, mesh)
}
func (self Simple) SoftBodyGetBounds(body gd.RID) gd.AABB {
	gc := gd.GarbageCollector(); _ = gc
	return gd.AABB(Expert(self).SoftBodyGetBounds(body))
}
func (self Simple) SoftBodySetCollisionLayer(body gd.RID, layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetCollisionLayer(body, gd.Int(layer))
}
func (self Simple) SoftBodyGetCollisionLayer(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).SoftBodyGetCollisionLayer(body)))
}
func (self Simple) SoftBodySetCollisionMask(body gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetCollisionMask(body, gd.Int(mask))
}
func (self Simple) SoftBodyGetCollisionMask(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).SoftBodyGetCollisionMask(body)))
}
func (self Simple) SoftBodyAddCollisionException(body gd.RID, body_b gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodyAddCollisionException(body, body_b)
}
func (self Simple) SoftBodyRemoveCollisionException(body gd.RID, body_b gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodyRemoveCollisionException(body, body_b)
}
func (self Simple) SoftBodySetState(body gd.RID, state classdb.PhysicsServer3DBodyState, variant gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetState(body, state, variant)
}
func (self Simple) SoftBodyGetState(body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).SoftBodyGetState(gc, body, state))
}
func (self Simple) SoftBodySetTransform(body gd.RID, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetTransform(body, transform)
}
func (self Simple) SoftBodySetRayPickable(body gd.RID, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetRayPickable(body, enable)
}
func (self Simple) SoftBodySetSimulationPrecision(body gd.RID, simulation_precision int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetSimulationPrecision(body, gd.Int(simulation_precision))
}
func (self Simple) SoftBodyGetSimulationPrecision(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).SoftBodyGetSimulationPrecision(body)))
}
func (self Simple) SoftBodySetTotalMass(body gd.RID, total_mass float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetTotalMass(body, gd.Float(total_mass))
}
func (self Simple) SoftBodyGetTotalMass(body gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).SoftBodyGetTotalMass(body)))
}
func (self Simple) SoftBodySetLinearStiffness(body gd.RID, stiffness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetLinearStiffness(body, gd.Float(stiffness))
}
func (self Simple) SoftBodyGetLinearStiffness(body gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).SoftBodyGetLinearStiffness(body)))
}
func (self Simple) SoftBodySetPressureCoefficient(body gd.RID, pressure_coefficient float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetPressureCoefficient(body, gd.Float(pressure_coefficient))
}
func (self Simple) SoftBodyGetPressureCoefficient(body gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).SoftBodyGetPressureCoefficient(body)))
}
func (self Simple) SoftBodySetDampingCoefficient(body gd.RID, damping_coefficient float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetDampingCoefficient(body, gd.Float(damping_coefficient))
}
func (self Simple) SoftBodyGetDampingCoefficient(body gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).SoftBodyGetDampingCoefficient(body)))
}
func (self Simple) SoftBodySetDragCoefficient(body gd.RID, drag_coefficient float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodySetDragCoefficient(body, gd.Float(drag_coefficient))
}
func (self Simple) SoftBodyGetDragCoefficient(body gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).SoftBodyGetDragCoefficient(body)))
}
func (self Simple) SoftBodyMovePoint(body gd.RID, point_index int, global_position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodyMovePoint(body, gd.Int(point_index), global_position)
}
func (self Simple) SoftBodyGetPointGlobalPosition(body gd.RID, point_index int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).SoftBodyGetPointGlobalPosition(body, gd.Int(point_index)))
}
func (self Simple) SoftBodyRemoveAllPinnedPoints(body gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodyRemoveAllPinnedPoints(body)
}
func (self Simple) SoftBodyPinPoint(body gd.RID, point_index int, pin bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SoftBodyPinPoint(body, gd.Int(point_index), pin)
}
func (self Simple) SoftBodyIsPointPinned(body gd.RID, point_index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).SoftBodyIsPointPinned(body, gd.Int(point_index)))
}
func (self Simple) JointCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).JointCreate())
}
func (self Simple) JointClear(joint gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointClear(joint)
}
func (self Simple) JointMakePin(joint gd.RID, body_A gd.RID, local_A gd.Vector3, body_B gd.RID, local_B gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointMakePin(joint, body_A, local_A, body_B, local_B)
}
func (self Simple) PinJointSetParam(joint gd.RID, param classdb.PhysicsServer3DPinJointParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PinJointSetParam(joint, param, gd.Float(value))
}
func (self Simple) PinJointGetParam(joint gd.RID, param classdb.PhysicsServer3DPinJointParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).PinJointGetParam(joint, param)))
}
func (self Simple) PinJointSetLocalA(joint gd.RID, local_A gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PinJointSetLocalA(joint, local_A)
}
func (self Simple) PinJointGetLocalA(joint gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).PinJointGetLocalA(joint))
}
func (self Simple) PinJointSetLocalB(joint gd.RID, local_B gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PinJointSetLocalB(joint, local_B)
}
func (self Simple) PinJointGetLocalB(joint gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).PinJointGetLocalB(joint))
}
func (self Simple) JointMakeHinge(joint gd.RID, body_A gd.RID, hinge_A gd.Transform3D, body_B gd.RID, hinge_B gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointMakeHinge(joint, body_A, hinge_A, body_B, hinge_B)
}
func (self Simple) HingeJointSetParam(joint gd.RID, param classdb.PhysicsServer3DHingeJointParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).HingeJointSetParam(joint, param, gd.Float(value))
}
func (self Simple) HingeJointGetParam(joint gd.RID, param classdb.PhysicsServer3DHingeJointParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).HingeJointGetParam(joint, param)))
}
func (self Simple) HingeJointSetFlag(joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).HingeJointSetFlag(joint, flag, enabled)
}
func (self Simple) HingeJointGetFlag(joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HingeJointGetFlag(joint, flag))
}
func (self Simple) JointMakeSlider(joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointMakeSlider(joint, body_A, local_ref_A, body_B, local_ref_B)
}
func (self Simple) SliderJointSetParam(joint gd.RID, param classdb.PhysicsServer3DSliderJointParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SliderJointSetParam(joint, param, gd.Float(value))
}
func (self Simple) SliderJointGetParam(joint gd.RID, param classdb.PhysicsServer3DSliderJointParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).SliderJointGetParam(joint, param)))
}
func (self Simple) JointMakeConeTwist(joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointMakeConeTwist(joint, body_A, local_ref_A, body_B, local_ref_B)
}
func (self Simple) ConeTwistJointSetParam(joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ConeTwistJointSetParam(joint, param, gd.Float(value))
}
func (self Simple) ConeTwistJointGetParam(joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ConeTwistJointGetParam(joint, param)))
}
func (self Simple) JointGetType(joint gd.RID) classdb.PhysicsServer3DJointType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicsServer3DJointType(Expert(self).JointGetType(joint))
}
func (self Simple) JointSetSolverPriority(joint gd.RID, priority int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointSetSolverPriority(joint, gd.Int(priority))
}
func (self Simple) JointGetSolverPriority(joint gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).JointGetSolverPriority(joint)))
}
func (self Simple) JointDisableCollisionsBetweenBodies(joint gd.RID, disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointDisableCollisionsBetweenBodies(joint, disable)
}
func (self Simple) JointIsDisabledCollisionsBetweenBodies(joint gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).JointIsDisabledCollisionsBetweenBodies(joint))
}
func (self Simple) JointMakeGeneric6dof(joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointMakeGeneric6dof(joint, body_A, local_ref_A, body_B, local_ref_B)
}
func (self Simple) Generic6dofJointSetParam(joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Generic6dofJointSetParam(joint, axis, param, gd.Float(value))
}
func (self Simple) Generic6dofJointGetParam(joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).Generic6dofJointGetParam(joint, axis, param)))
}
func (self Simple) Generic6dofJointSetFlag(joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Generic6dofJointSetFlag(joint, axis, flag, enable)
}
func (self Simple) Generic6dofJointGetFlag(joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).Generic6dofJointGetFlag(joint, axis, flag))
}
func (self Simple) FreeRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FreeRid(rid)
}
func (self Simple) SetActive(active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetActive(active)
}
func (self Simple) GetProcessInfo(process_info classdb.PhysicsServer3DProcessInfo) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetProcessInfo(process_info)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicsServer3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) WorldBoundaryShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_world_boundary_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SeparationRayShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_separation_ray_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SphereShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_sphere_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) BoxShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_box_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) CapsuleShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_capsule_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) CylinderShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_cylinder_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) ConvexPolygonShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_convex_polygon_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) ConcavePolygonShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_concave_polygon_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) HeightmapShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_heightmap_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) CustomShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_custom_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the shape data that defines its shape and size. The data to be passed depends on the kind of shape created [method shape_get_type].
*/
//go:nosplit
func (self class) ShapeSetData(shape gd.RID, data gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_shape_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the collision margin for the shape.
[b]Note:[/b] This is not used in Godot Physics.
*/
//go:nosplit
func (self class) ShapeSetMargin(shape gd.RID, margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_shape_set_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the type of shape (see [enum ShapeType] constants).
*/
//go:nosplit
func (self class) ShapeGetType(shape gd.RID) classdb.PhysicsServer3DShapeType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret = callframe.Ret[classdb.PhysicsServer3DShapeType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_shape_get_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the shape data.
*/
//go:nosplit
func (self class) ShapeGetData(ctx gd.Lifetime, shape gd.RID) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_shape_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the collision margin for the shape.
[b]Note:[/b] This is not used in Godot Physics, so will always return [code]0[/code].
*/
//go:nosplit
func (self class) ShapeGetMargin(shape gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_shape_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a space. A space is a collection of parameters for the physics engine that can be assigned to an area or a body. It can be assigned to an area with [method area_set_space], or to a body with [method body_set_space].
*/
//go:nosplit
func (self class) SpaceCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_space_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Marks a space as active. It will not have an effect, unless it is assigned to an area or body.
*/
//go:nosplit
func (self class) SpaceSetActive(space gd.RID, active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_space_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the space is active.
*/
//go:nosplit
func (self class) SpaceIsActive(space gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_space_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value for a space parameter. A list of available parameters is on the [enum SpaceParameter] constants.
*/
//go:nosplit
func (self class) SpaceSetParam(space gd.RID, param classdb.PhysicsServer3DSpaceParameter, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_space_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of a space parameter.
*/
//go:nosplit
func (self class) SpaceGetParam(space gd.RID, param classdb.PhysicsServer3DSpaceParameter) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_space_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the state of a space, a [PhysicsDirectSpaceState3D]. This object can be used to make collision/intersection queries.
*/
//go:nosplit
func (self class) SpaceGetDirectState(ctx gd.Lifetime, space gd.RID) object.PhysicsDirectSpaceState3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_space_get_direct_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PhysicsDirectSpaceState3D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a 3D area object in the physics server, and returns the [RID] that identifies it. The default settings for the created area include a collision layer and mask set to [code]1[/code], and [code]monitorable[/code] set to [code]false[/code].
Use [method area_add_shape] to add shapes to it, use [method area_set_transform] to set its transform, and use [method area_set_space] to add the area to a space. If you want the area to be detectable use [method area_set_monitorable].
*/
//go:nosplit
func (self class) AreaCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Assigns a space to the area.
*/
//go:nosplit
func (self class) AreaSetSpace(area gd.RID, space gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the space assigned to the area.
*/
//go:nosplit
func (self class) AreaGetSpace(area gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_get_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a shape to the area, along with a transform matrix. Shapes are usually referenced by their index, so you should track which shape has a given index.
*/
//go:nosplit
func (self class) AreaAddShape(area gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape)
	callframe.Arg(frame, transform)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_add_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Substitutes a given area shape by another. The old shape is selected by its index, the new one by its [RID].
*/
//go:nosplit
func (self class) AreaSetShape(area gd.RID, shape_idx gd.Int, shape gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transform matrix for an area shape.
*/
//go:nosplit
func (self class) AreaSetShapeTransform(area gd.RID, shape_idx gd.Int, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_shape_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreaSetShapeDisabled(area gd.RID, shape_idx gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_shape_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of shapes assigned to an area.
*/
//go:nosplit
func (self class) AreaGetShapeCount(area gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_get_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [RID] of the nth shape of an area.
*/
//go:nosplit
func (self class) AreaGetShape(area gd.RID, shape_idx gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform matrix of a shape within an area.
*/
//go:nosplit
func (self class) AreaGetShapeTransform(area gd.RID, shape_idx gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_get_shape_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes a shape from an area. It does not delete the shape, so it can be reassigned later.
*/
//go:nosplit
func (self class) AreaRemoveShape(area gd.RID, shape_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_remove_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all shapes from an area. It does not delete the shapes, so they can be reassigned later.
*/
//go:nosplit
func (self class) AreaClearShapes(area gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_clear_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Assigns the area to one or many physics layers.
*/
//go:nosplit
func (self class) AreaSetCollisionLayer(area gd.RID, layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers an area belongs to.
*/
//go:nosplit
func (self class) AreaGetCollisionLayer(area gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets which physics layers the area will monitor.
*/
//go:nosplit
func (self class) AreaSetCollisionMask(area gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers an area can contact with.
*/
//go:nosplit
func (self class) AreaGetCollisionMask(area gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value for an area parameter. A list of available parameters is on the [enum AreaParameter] constants.
*/
//go:nosplit
func (self class) AreaSetParam(area gd.RID, param classdb.PhysicsServer3DAreaParameter, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, param)
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transform matrix for an area.
*/
//go:nosplit
func (self class) AreaSetTransform(area gd.RID, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an area parameter value. A list of available parameters is on the [enum AreaParameter] constants.
*/
//go:nosplit
func (self class) AreaGetParam(ctx gd.Lifetime, area gd.RID, param classdb.PhysicsServer3DAreaParameter) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the transform matrix for an area.
*/
//go:nosplit
func (self class) AreaGetTransform(area gd.RID) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Assigns the area to a descendant of [Object], so it can exist in the node tree.
*/
//go:nosplit
func (self class) AreaAttachObjectInstanceId(area gd.RID, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_attach_object_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets the instance ID of the object the area is assigned to.
*/
//go:nosplit
func (self class) AreaGetObjectInstanceId(area gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_get_object_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the area's body monitor callback. This callback will be called when any other (shape of a) body enters or exits (a shape of) the given area, and must take the following five parameters:
1. an integer [code]status[/code]: either [constant AREA_BODY_ADDED] or [constant AREA_BODY_REMOVED] depending on whether the other body shape entered or exited the area,
2. an [RID] [code]body_rid[/code]: the [RID] of the body that entered or exited the area,
3. an integer [code]instance_id[/code]: the [code]ObjectID[/code] attached to the body,
4. an integer [code]body_shape_idx[/code]: the index of the shape of the body that entered or exited the area,
5. an integer [code]self_shape_idx[/code]: the index of the shape of the area where the body entered or exited.
By counting (or keeping track of) the shapes that enter and exit, it can be determined if a body (with all its shapes) is entering for the first time or exiting for the last time.
*/
//go:nosplit
func (self class) AreaSetMonitorCallback(area gd.RID, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_monitor_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the area's area monitor callback. This callback will be called when any other (shape of an) area enters or exits (a shape of) the given area, and must take the following five parameters:
1. an integer [code]status[/code]: either [constant AREA_BODY_ADDED] or [constant AREA_BODY_REMOVED] depending on whether the other area's shape entered or exited the area,
2. an [RID] [code]area_rid[/code]: the [RID] of the other area that entered or exited the area,
3. an integer [code]instance_id[/code]: the [code]ObjectID[/code] attached to the other area,
4. an integer [code]area_shape_idx[/code]: the index of the shape of the other area that entered or exited the area,
5. an integer [code]self_shape_idx[/code]: the index of the shape of the area where the other area entered or exited.
By counting (or keeping track of) the shapes that enter and exit, it can be determined if an area (with all its shapes) is entering for the first time or exiting for the last time.
*/
//go:nosplit
func (self class) AreaSetAreaMonitorCallback(area gd.RID, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_area_monitor_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreaSetMonitorable(area gd.RID, monitorable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, monitorable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_monitorable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets object pickable with rays.
*/
//go:nosplit
func (self class) AreaSetRayPickable(area gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_area_set_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a 3D body object in the physics server, and returns the [RID] that identifies it. The default settings for the created area include a collision layer and mask set to [code]1[/code], and body mode set to [constant BODY_MODE_RIGID].
Use [method body_add_shape] to add shapes to it, use [method body_set_state] to set its transform, and use [method body_set_space] to add the body to a space.
*/
//go:nosplit
func (self class) BodyCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Assigns a space to the body (see [method space_create]).
*/
//go:nosplit
func (self class) BodySetSpace(body gd.RID, space gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [RID] of the space assigned to a body.
*/
//go:nosplit
func (self class) BodyGetSpace(body gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the body mode, from one of the [enum BodyMode] constants.
*/
//go:nosplit
func (self class) BodySetMode(body gd.RID, mode classdb.PhysicsServer3DBodyMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the body mode.
*/
//go:nosplit
func (self class) BodyGetMode(body gd.RID) classdb.PhysicsServer3DBodyMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[classdb.PhysicsServer3DBodyMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the physics layer or layers a body belongs to.
*/
//go:nosplit
func (self class) BodySetCollisionLayer(body gd.RID, layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers a body belongs to.
*/
//go:nosplit
func (self class) BodyGetCollisionLayer(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the physics layer or layers a body can collide with.
*/
//go:nosplit
func (self class) BodySetCollisionMask(body gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers a body can collide with.
*/
//go:nosplit
func (self class) BodyGetCollisionMask(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the body's collision priority.
*/
//go:nosplit
func (self class) BodySetCollisionPriority(body gd.RID, priority gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the body's collision priority.
*/
//go:nosplit
func (self class) BodyGetCollisionPriority(body gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a shape to the body, along with a transform matrix. Shapes are usually referenced by their index, so you should track which shape has a given index.
*/
//go:nosplit
func (self class) BodyAddShape(body gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape)
	callframe.Arg(frame, transform)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_add_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Substitutes a given body shape by another. The old shape is selected by its index, the new one by its [RID].
*/
//go:nosplit
func (self class) BodySetShape(body gd.RID, shape_idx gd.Int, shape gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transform matrix for a body shape.
*/
//go:nosplit
func (self class) BodySetShapeTransform(body gd.RID, shape_idx gd.Int, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_shape_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) BodySetShapeDisabled(body gd.RID, shape_idx gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_shape_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of shapes assigned to a body.
*/
//go:nosplit
func (self class) BodyGetShapeCount(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [RID] of the nth shape of a body.
*/
//go:nosplit
func (self class) BodyGetShape(body gd.RID, shape_idx gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform matrix of a body shape.
*/
//go:nosplit
func (self class) BodyGetShapeTransform(body gd.RID, shape_idx gd.Int) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_shape_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes a shape from a body. The shape is not deleted, so it can be reused afterwards.
*/
//go:nosplit
func (self class) BodyRemoveShape(body gd.RID, shape_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_remove_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all shapes from a body.
*/
//go:nosplit
func (self class) BodyClearShapes(body gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_clear_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Assigns the area to a descendant of [Object], so it can exist in the node tree.
*/
//go:nosplit
func (self class) BodyAttachObjectInstanceId(body gd.RID, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_attach_object_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets the instance ID of the object the area is assigned to.
*/
//go:nosplit
func (self class) BodyGetObjectInstanceId(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_object_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the continuous collision detection mode is enabled.
Continuous collision detection tries to predict where a moving body will collide, instead of moving it and correcting its movement if it collided.
*/
//go:nosplit
func (self class) BodySetEnableContinuousCollisionDetection(body gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_enable_continuous_collision_detection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the continuous collision detection mode is enabled.
*/
//go:nosplit
func (self class) BodyIsContinuousCollisionDetectionEnabled(body gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_is_continuous_collision_detection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a body parameter. A list of available parameters is on the [enum BodyParameter] constants.
*/
//go:nosplit
func (self class) BodySetParam(body gd.RID, param classdb.PhysicsServer3DBodyParameter, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, param)
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of a body parameter. A list of available parameters is on the [enum BodyParameter] constants.
*/
//go:nosplit
func (self class) BodyGetParam(ctx gd.Lifetime, body gd.RID, param classdb.PhysicsServer3DBodyParameter) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Restores the default inertia and center of mass based on shapes to cancel any custom values previously set using [method body_set_param].
*/
//go:nosplit
func (self class) BodyResetMassProperties(body gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_reset_mass_properties, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a body state (see [enum BodyState] constants).
*/
//go:nosplit
func (self class) BodySetState(body gd.RID, state classdb.PhysicsServer3DBodyState, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, state)
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a body state.
*/
//go:nosplit
func (self class) BodyGetState(ctx gd.Lifetime, body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, state)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method body_apply_impulse] at the body's center of mass.
*/
//go:nosplit
func (self class) BodyApplyCentralImpulse(body gd.RID, impulse gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) BodyApplyImpulse(body gd.RID, impulse gd.Vector3, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
*/
//go:nosplit
func (self class) BodyApplyTorqueImpulse(body gd.RID, impulse gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_apply_torque_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method body_apply_force] at the body's center of mass.
*/
//go:nosplit
func (self class) BodyApplyCentralForce(body gd.RID, force gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_apply_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) BodyApplyForce(body gd.RID, force gd.Vector3, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_apply_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
*/
//go:nosplit
func (self class) BodyApplyTorque(body gd.RID, torque gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_apply_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]body_set_constant_force(body, Vector3(0, 0, 0))[/code].
This is equivalent to using [method body_add_constant_force] at the body's center of mass.
*/
//go:nosplit
func (self class) BodyAddConstantCentralForce(body gd.RID, force gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_add_constant_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]body_set_constant_force(body, Vector3(0, 0, 0))[/code].
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) BodyAddConstantForce(body gd.RID, force gd.Vector3, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_add_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]body_set_constant_torque(body, Vector3(0, 0, 0))[/code].
*/
//go:nosplit
func (self class) BodyAddConstantTorque(body gd.RID, torque gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_add_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the body's total constant positional forces applied during each physics update.
See [method body_add_constant_force] and [method body_add_constant_central_force].
*/
//go:nosplit
func (self class) BodySetConstantForce(body gd.RID, force gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the body's total constant positional forces applied during each physics update.
See [method body_add_constant_force] and [method body_add_constant_central_force].
*/
//go:nosplit
func (self class) BodyGetConstantForce(body gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the body's total constant rotational forces applied during each physics update.
See [method body_add_constant_torque].
*/
//go:nosplit
func (self class) BodySetConstantTorque(body gd.RID, torque gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the body's total constant rotational forces applied during each physics update.
See [method body_add_constant_torque].
*/
//go:nosplit
func (self class) BodyGetConstantTorque(body gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets an axis velocity. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
//go:nosplit
func (self class) BodySetAxisVelocity(body gd.RID, axis_velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, axis_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_axis_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) BodySetAxisLock(body gd.RID, axis classdb.PhysicsServer3DBodyAxis, lock bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, axis)
	callframe.Arg(frame, lock)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_axis_lock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) BodyIsAxisLocked(body gd.RID, axis classdb.PhysicsServer3DBodyAxis) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, axis)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_is_axis_locked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a body to the list of bodies exempt from collisions.
*/
//go:nosplit
func (self class) BodyAddCollisionException(body gd.RID, excepted_body gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, excepted_body)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_add_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a body from the list of bodies exempt from collisions.
Continuous collision detection tries to predict where a moving body will collide, instead of moving it and correcting its movement if it collided.
*/
//go:nosplit
func (self class) BodyRemoveCollisionException(body gd.RID, excepted_body gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, excepted_body)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_remove_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the maximum contacts to report. Bodies can keep a log of the contacts with other bodies. This is enabled by setting the maximum number of contacts reported to a number greater than 0.
*/
//go:nosplit
func (self class) BodySetMaxContactsReported(body gd.RID, amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum contacts that can be reported. See [method body_set_max_contacts_reported].
*/
//go:nosplit
func (self class) BodyGetMaxContactsReported(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the body omits the standard force integration. If [param enable] is [code]true[/code], the body will not automatically use applied forces, torques, and damping to update the body's linear and angular velocity. In this case, [method body_set_force_integration_callback] can be used to manually update the linear and angular velocity instead.
This method is called when the property [member RigidBody3D.custom_integrator] is set.
*/
//go:nosplit
func (self class) BodySetOmitForceIntegration(body gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_omit_force_integration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the body is omitting the standard force integration. See [method body_set_omit_force_integration].
*/
//go:nosplit
func (self class) BodyIsOmittingForceIntegration(body gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_is_omitting_force_integration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the body's state synchronization callback function to [param callable]. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the callback.
The function [param callable] will be called every physics frame, assuming that the body was active during the previous physics tick, and can be used to fetch the latest state from the physics server.
The function [param callable] must take the following parameters:
1. [code]state[/code]: a [PhysicsDirectBodyState3D], used to retrieve the body's state.
*/
//go:nosplit
func (self class) BodySetStateSyncCallback(body gd.RID, callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mmm.Get(callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_state_sync_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the body's custom force integration callback function to [param callable]. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the custom callback.
The function [param callable] will be called every physics tick, before the standard force integration (see [method body_set_omit_force_integration]). It can be used for example to update the body's linear and angular velocity based on contact with other bodies.
If [param userdata] is not [code]null[/code], the function [param callable] must take the following two parameters:
1. [code]state[/code]: a [PhysicsDirectBodyState3D], used to retrieve and modify the body's state,
2. [code skip-lint]userdata[/code]: a [Variant]; its value will be the [param userdata] passed into this method.
If [param userdata] is [code]null[/code], then [param callable] must take only the [code]state[/code] parameter.
*/
//go:nosplit
func (self class) BodySetForceIntegrationCallback(body gd.RID, callable gd.Callable, userdata gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mmm.Get(callable))
	callframe.Arg(frame, mmm.Get(userdata))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_force_integration_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the body pickable with rays if [param enable] is set.
*/
//go:nosplit
func (self class) BodySetRayPickable(body gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_set_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a collision would result from moving along a motion vector from a given point in space. [PhysicsTestMotionParameters3D] is passed to set motion parameters. [PhysicsTestMotionResult3D] can be passed to return additional information.
*/
//go:nosplit
func (self class) BodyTestMotion(body gd.RID, parameters object.PhysicsTestMotionParameters3D, result object.PhysicsTestMotionResult3D) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mmm.Get(parameters[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(result[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_test_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [PhysicsDirectBodyState3D] of the body. Returns [code]null[/code] if the body is destroyed or removed from the physics space.
*/
//go:nosplit
func (self class) BodyGetDirectState(ctx gd.Lifetime, body gd.RID) object.PhysicsDirectBodyState3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_body_get_direct_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PhysicsDirectBodyState3D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new soft body and returns its internal [RID].
*/
//go:nosplit
func (self class) SoftBodyCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Requests that the physics server updates the rendering server with the latest positions of the given soft body's points through the [param rendering_server_handler] interface.
*/
//go:nosplit
func (self class) SoftBodyUpdateRenderingServer(body gd.RID, rendering_server_handler object.PhysicsServer3DRenderingServerHandler)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mmm.Get(rendering_server_handler[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_update_rendering_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Assigns a space to the given soft body (see [method space_create]).
*/
//go:nosplit
func (self class) SoftBodySetSpace(body gd.RID, space gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [RID] of the space assigned to the given soft body.
*/
//go:nosplit
func (self class) SoftBodyGetSpace(body gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the mesh of the given soft body.
*/
//go:nosplit
func (self class) SoftBodySetMesh(body gd.RID, mesh gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mesh)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the bounds of the given soft body in global coordinates.
*/
//go:nosplit
func (self class) SoftBodyGetBounds(body gd.RID) gd.AABB {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.AABB](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_bounds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the physics layer or layers the given soft body belongs to.
*/
//go:nosplit
func (self class) SoftBodySetCollisionLayer(body gd.RID, layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers that the given soft body belongs to.
*/
//go:nosplit
func (self class) SoftBodyGetCollisionLayer(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the physics layer or layers the given soft body can collide with.
*/
//go:nosplit
func (self class) SoftBodySetCollisionMask(body gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers that the given soft body can collide with.
*/
//go:nosplit
func (self class) SoftBodyGetCollisionMask(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds the given body to the list of bodies exempt from collisions.
*/
//go:nosplit
func (self class) SoftBodyAddCollisionException(body gd.RID, body_b gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, body_b)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_add_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given body from the list of bodies exempt from collisions.
*/
//go:nosplit
func (self class) SoftBodyRemoveCollisionException(body gd.RID, body_b gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, body_b)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_remove_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given body state for the given body (see [enum BodyState] constants).
[b]Note:[/b] Godot's default physics implementation does not support [constant BODY_STATE_LINEAR_VELOCITY], [constant BODY_STATE_ANGULAR_VELOCITY], [constant BODY_STATE_SLEEPING], or [constant BODY_STATE_CAN_SLEEP].
*/
//go:nosplit
func (self class) SoftBodySetState(body gd.RID, state classdb.PhysicsServer3DBodyState, variant gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, state)
	callframe.Arg(frame, mmm.Get(variant))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given soft body state (see [enum BodyState] constants).
[b]Note:[/b] Godot's default physics implementation does not support [constant BODY_STATE_LINEAR_VELOCITY], [constant BODY_STATE_ANGULAR_VELOCITY], [constant BODY_STATE_SLEEPING], or [constant BODY_STATE_CAN_SLEEP].
*/
//go:nosplit
func (self class) SoftBodyGetState(ctx gd.Lifetime, body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, state)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the global transform of the given soft body.
*/
//go:nosplit
func (self class) SoftBodySetTransform(body gd.RID, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether the given soft body will be pickable when using object picking.
*/
//go:nosplit
func (self class) SoftBodySetRayPickable(body gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the simulation precision of the given soft body. Increasing this value will improve the resulting simulation, but can affect performance. Use with care.
*/
//go:nosplit
func (self class) SoftBodySetSimulationPrecision(body gd.RID, simulation_precision gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, simulation_precision)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_simulation_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the simulation precision of the given soft body.
*/
//go:nosplit
func (self class) SoftBodyGetSimulationPrecision(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_simulation_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the total mass for the given soft body.
*/
//go:nosplit
func (self class) SoftBodySetTotalMass(body gd.RID, total_mass gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, total_mass)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_total_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the total mass assigned to the given soft body.
*/
//go:nosplit
func (self class) SoftBodyGetTotalMass(body gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_total_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the linear stiffness of the given soft body. Higher values will result in a stiffer body, while lower values will increase the body's ability to bend. The value can be between [code]0.0[/code] and [code]1.0[/code] (inclusive).
*/
//go:nosplit
func (self class) SoftBodySetLinearStiffness(body gd.RID, stiffness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, stiffness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_linear_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the linear stiffness of the given soft body.
*/
//go:nosplit
func (self class) SoftBodyGetLinearStiffness(body gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_linear_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the pressure coefficient of the given soft body. Simulates pressure build-up from inside this body. Higher values increase the strength of this effect.
*/
//go:nosplit
func (self class) SoftBodySetPressureCoefficient(body gd.RID, pressure_coefficient gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, pressure_coefficient)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_pressure_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the pressure coefficient of the given soft body.
*/
//go:nosplit
func (self class) SoftBodyGetPressureCoefficient(body gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_pressure_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the damping coefficient of the given soft body. Higher values will slow down the body more noticeably when forces are applied.
*/
//go:nosplit
func (self class) SoftBodySetDampingCoefficient(body gd.RID, damping_coefficient gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, damping_coefficient)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_damping_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the damping coefficient of the given soft body.
*/
//go:nosplit
func (self class) SoftBodyGetDampingCoefficient(body gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_damping_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the drag coefficient of the given soft body. Higher values increase this body's air resistance.
[b]Note:[/b] This value is currently unused by Godot's default physics implementation.
*/
//go:nosplit
func (self class) SoftBodySetDragCoefficient(body gd.RID, drag_coefficient gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, drag_coefficient)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_set_drag_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the drag coefficient of the given soft body.
*/
//go:nosplit
func (self class) SoftBodyGetDragCoefficient(body gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_drag_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves the given soft body point to a position in global coordinates.
*/
//go:nosplit
func (self class) SoftBodyMovePoint(body gd.RID, point_index gd.Int, global_position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, point_index)
	callframe.Arg(frame, global_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_move_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current position of the given soft body point in global coordinates.
*/
//go:nosplit
func (self class) SoftBodyGetPointGlobalPosition(body gd.RID, point_index gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, point_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_get_point_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Unpins all points of the given soft body.
*/
//go:nosplit
func (self class) SoftBodyRemoveAllPinnedPoints(body gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_remove_all_pinned_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Pins or unpins the given soft body point based on the value of [param pin].
[b]Note:[/b] Pinning a point effectively makes it kinematic, preventing it from being affected by forces, but you can still move it using [method soft_body_move_point].
*/
//go:nosplit
func (self class) SoftBodyPinPoint(body gd.RID, point_index gd.Int, pin bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, point_index)
	callframe.Arg(frame, pin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_pin_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the given soft body point is pinned.
*/
//go:nosplit
func (self class) SoftBodyIsPointPinned(body gd.RID, point_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, point_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_soft_body_is_point_pinned, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) JointCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) JointClear(joint gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) JointMakePin(joint gd.RID, body_A gd.RID, local_A gd.Vector3, body_B gd.RID, local_B gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, body_A)
	callframe.Arg(frame, local_A)
	callframe.Arg(frame, body_B)
	callframe.Arg(frame, local_B)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_make_pin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a pin_joint parameter (see [enum PinJointParam] constants).
*/
//go:nosplit
func (self class) PinJointSetParam(joint gd.RID, param classdb.PhysicsServer3DPinJointParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_pin_joint_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a pin_joint parameter (see [enum PinJointParam] constants).
*/
//go:nosplit
func (self class) PinJointGetParam(joint gd.RID, param classdb.PhysicsServer3DPinJointParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_pin_joint_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets position of the joint in the local space of body a of the joint.
*/
//go:nosplit
func (self class) PinJointSetLocalA(joint gd.RID, local_A gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, local_A)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_pin_joint_set_local_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns position of the joint in the local space of body a of the joint.
*/
//go:nosplit
func (self class) PinJointGetLocalA(joint gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_pin_joint_get_local_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets position of the joint in the local space of body b of the joint.
*/
//go:nosplit
func (self class) PinJointSetLocalB(joint gd.RID, local_B gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, local_B)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_pin_joint_set_local_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns position of the joint in the local space of body b of the joint.
*/
//go:nosplit
func (self class) PinJointGetLocalB(joint gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_pin_joint_get_local_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) JointMakeHinge(joint gd.RID, body_A gd.RID, hinge_A gd.Transform3D, body_B gd.RID, hinge_B gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, body_A)
	callframe.Arg(frame, hinge_A)
	callframe.Arg(frame, body_B)
	callframe.Arg(frame, hinge_B)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_make_hinge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a hinge_joint parameter (see [enum HingeJointParam] constants).
*/
//go:nosplit
func (self class) HingeJointSetParam(joint gd.RID, param classdb.PhysicsServer3DHingeJointParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_hinge_joint_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a hinge_joint parameter (see [enum HingeJointParam]).
*/
//go:nosplit
func (self class) HingeJointGetParam(joint gd.RID, param classdb.PhysicsServer3DHingeJointParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_hinge_joint_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a hinge_joint flag (see [enum HingeJointFlag] constants).
*/
//go:nosplit
func (self class) HingeJointSetFlag(joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_hinge_joint_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a hinge_joint flag (see [enum HingeJointFlag] constants).
*/
//go:nosplit
func (self class) HingeJointGetFlag(joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_hinge_joint_get_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) JointMakeSlider(joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, body_A)
	callframe.Arg(frame, local_ref_A)
	callframe.Arg(frame, body_B)
	callframe.Arg(frame, local_ref_B)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_make_slider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a slider_joint parameter (see [enum SliderJointParam] constants).
*/
//go:nosplit
func (self class) SliderJointSetParam(joint gd.RID, param classdb.PhysicsServer3DSliderJointParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_slider_joint_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a slider_joint parameter (see [enum SliderJointParam] constants).
*/
//go:nosplit
func (self class) SliderJointGetParam(joint gd.RID, param classdb.PhysicsServer3DSliderJointParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_slider_joint_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) JointMakeConeTwist(joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, body_A)
	callframe.Arg(frame, local_ref_A)
	callframe.Arg(frame, body_B)
	callframe.Arg(frame, local_ref_B)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_make_cone_twist, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a cone_twist_joint parameter (see [enum ConeTwistJointParam] constants).
*/
//go:nosplit
func (self class) ConeTwistJointSetParam(joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_cone_twist_joint_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a cone_twist_joint parameter (see [enum ConeTwistJointParam] constants).
*/
//go:nosplit
func (self class) ConeTwistJointGetParam(joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_cone_twist_joint_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the type of the Joint3D.
*/
//go:nosplit
func (self class) JointGetType(joint gd.RID) classdb.PhysicsServer3DJointType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[classdb.PhysicsServer3DJointType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_get_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the priority value of the Joint3D.
*/
//go:nosplit
func (self class) JointSetSolverPriority(joint gd.RID, priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_set_solver_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets the priority value of the Joint3D.
*/
//go:nosplit
func (self class) JointGetSolverPriority(joint gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_get_solver_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the bodies attached to the [Joint3D] will collide with each other.
*/
//go:nosplit
func (self class) JointDisableCollisionsBetweenBodies(joint gd.RID, disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_disable_collisions_between_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the bodies attached to the [Joint3D] will collide with each other.
*/
//go:nosplit
func (self class) JointIsDisabledCollisionsBetweenBodies(joint gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_is_disabled_collisions_between_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Make the joint a generic six degrees of freedom (6DOF) joint. Use [method generic_6dof_joint_set_flag] and [method generic_6dof_joint_set_param] to set the joint's flags and parameters respectively.
*/
//go:nosplit
func (self class) JointMakeGeneric6dof(joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, body_A)
	callframe.Arg(frame, local_ref_A)
	callframe.Arg(frame, body_B)
	callframe.Arg(frame, local_ref_B)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_joint_make_generic_6dof, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the value of a given generic 6DOF joint parameter. See [enum G6DOFJointAxisParam] for the list of available parameters.
*/
//go:nosplit
func (self class) Generic6dofJointSetParam(joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, axis)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_generic_6dof_joint_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of a generic 6DOF joint parameter. See [enum G6DOFJointAxisParam] for the list of available parameters.
*/
//go:nosplit
func (self class) Generic6dofJointGetParam(joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, axis)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_generic_6dof_joint_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of a given generic 6DOF joint flag. See [enum G6DOFJointAxisFlag] for the list of available flags.
*/
//go:nosplit
func (self class) Generic6dofJointSetFlag(joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, axis)
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_generic_6dof_joint_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of a generic 6DOF joint flag. See [enum G6DOFJointAxisFlag] for the list of available flags.
*/
//go:nosplit
func (self class) Generic6dofJointGetFlag(joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, axis)
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_generic_6dof_joint_get_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Destroys any of the objects created by PhysicsServer3D. If the [RID] passed is not one of the objects that can be created by PhysicsServer3D, an error will be sent to the console.
*/
//go:nosplit
func (self class) FreeRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Activates or deactivates the 3D physics engine.
*/
//go:nosplit
func (self class) SetActive(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns information about the current state of the 3D physics engine. See [enum ProcessInfo] for a list of available states.
*/
//go:nosplit
func (self class) GetProcessInfo(process_info classdb.PhysicsServer3DProcessInfo) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, process_info)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer3D.Bind_get_process_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPhysicsServer3D() Expert { return self[0].AsPhysicsServer3D() }


//go:nosplit
func (self Simple) AsPhysicsServer3D() Simple { return self[0].AsPhysicsServer3D() }

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
func init() {classdb.Register("PhysicsServer3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type JointType = classdb.PhysicsServer3DJointType

const (
/*The [Joint3D] is a [PinJoint3D].*/
	JointTypePin JointType = 0
/*The [Joint3D] is a [HingeJoint3D].*/
	JointTypeHinge JointType = 1
/*The [Joint3D] is a [SliderJoint3D].*/
	JointTypeSlider JointType = 2
/*The [Joint3D] is a [ConeTwistJoint3D].*/
	JointTypeConeTwist JointType = 3
/*The [Joint3D] is a [Generic6DOFJoint3D].*/
	JointType6dof JointType = 4
/*Represents the size of the [enum JointType] enum.*/
	JointTypeMax JointType = 5
)
type PinJointParam = classdb.PhysicsServer3DPinJointParam

const (
/*The strength with which the pinned objects try to stay in positional relation to each other.
The higher, the stronger.*/
	PinJointBias PinJointParam = 0
/*The strength with which the pinned objects try to stay in velocity relation to each other.
The higher, the stronger.*/
	PinJointDamping PinJointParam = 1
/*If above 0, this value is the maximum value for an impulse that this Joint3D puts on its ends.*/
	PinJointImpulseClamp PinJointParam = 2
)
type HingeJointParam = classdb.PhysicsServer3DHingeJointParam

const (
/*The speed with which the two bodies get pulled together when they move in different directions.*/
	HingeJointBias HingeJointParam = 0
/*The maximum rotation across the Hinge.*/
	HingeJointLimitUpper HingeJointParam = 1
/*The minimum rotation across the Hinge.*/
	HingeJointLimitLower HingeJointParam = 2
/*The speed with which the rotation across the axis perpendicular to the hinge gets corrected.*/
	HingeJointLimitBias HingeJointParam = 3
	HingeJointLimitSoftness HingeJointParam = 4
/*The lower this value, the more the rotation gets slowed down.*/
	HingeJointLimitRelaxation HingeJointParam = 5
/*Target speed for the motor.*/
	HingeJointMotorTargetVelocity HingeJointParam = 6
/*Maximum acceleration for the motor.*/
	HingeJointMotorMaxImpulse HingeJointParam = 7
)
type HingeJointFlag = classdb.PhysicsServer3DHingeJointFlag

const (
/*If [code]true[/code], the Hinge has a maximum and a minimum rotation.*/
	HingeJointFlagUseLimit HingeJointFlag = 0
/*If [code]true[/code], a motor turns the Hinge.*/
	HingeJointFlagEnableMotor HingeJointFlag = 1
)
type SliderJointParam = classdb.PhysicsServer3DSliderJointParam

const (
/*The maximum difference between the pivot points on their X axis before damping happens.*/
	SliderJointLinearLimitUpper SliderJointParam = 0
/*The minimum difference between the pivot points on their X axis before damping happens.*/
	SliderJointLinearLimitLower SliderJointParam = 1
/*A factor applied to the movement across the slider axis once the limits get surpassed. The lower, the slower the movement.*/
	SliderJointLinearLimitSoftness SliderJointParam = 2
/*The amount of restitution once the limits are surpassed. The lower, the more velocity-energy gets lost.*/
	SliderJointLinearLimitRestitution SliderJointParam = 3
/*The amount of damping once the slider limits are surpassed.*/
	SliderJointLinearLimitDamping SliderJointParam = 4
/*A factor applied to the movement across the slider axis as long as the slider is in the limits. The lower, the slower the movement.*/
	SliderJointLinearMotionSoftness SliderJointParam = 5
/*The amount of restitution inside the slider limits.*/
	SliderJointLinearMotionRestitution SliderJointParam = 6
/*The amount of damping inside the slider limits.*/
	SliderJointLinearMotionDamping SliderJointParam = 7
/*A factor applied to the movement across axes orthogonal to the slider.*/
	SliderJointLinearOrthogonalSoftness SliderJointParam = 8
/*The amount of restitution when movement is across axes orthogonal to the slider.*/
	SliderJointLinearOrthogonalRestitution SliderJointParam = 9
/*The amount of damping when movement is across axes orthogonal to the slider.*/
	SliderJointLinearOrthogonalDamping SliderJointParam = 10
/*The upper limit of rotation in the slider.*/
	SliderJointAngularLimitUpper SliderJointParam = 11
/*The lower limit of rotation in the slider.*/
	SliderJointAngularLimitLower SliderJointParam = 12
/*A factor applied to the all rotation once the limit is surpassed.*/
	SliderJointAngularLimitSoftness SliderJointParam = 13
/*The amount of restitution of the rotation when the limit is surpassed.*/
	SliderJointAngularLimitRestitution SliderJointParam = 14
/*The amount of damping of the rotation when the limit is surpassed.*/
	SliderJointAngularLimitDamping SliderJointParam = 15
/*A factor that gets applied to the all rotation in the limits.*/
	SliderJointAngularMotionSoftness SliderJointParam = 16
/*The amount of restitution of the rotation in the limits.*/
	SliderJointAngularMotionRestitution SliderJointParam = 17
/*The amount of damping of the rotation in the limits.*/
	SliderJointAngularMotionDamping SliderJointParam = 18
/*A factor that gets applied to the all rotation across axes orthogonal to the slider.*/
	SliderJointAngularOrthogonalSoftness SliderJointParam = 19
/*The amount of restitution of the rotation across axes orthogonal to the slider.*/
	SliderJointAngularOrthogonalRestitution SliderJointParam = 20
/*The amount of damping of the rotation across axes orthogonal to the slider.*/
	SliderJointAngularOrthogonalDamping SliderJointParam = 21
/*Represents the size of the [enum SliderJointParam] enum.*/
	SliderJointMax SliderJointParam = 22
)
type ConeTwistJointParam = classdb.PhysicsServer3DConeTwistJointParam

const (
/*Swing is rotation from side to side, around the axis perpendicular to the twist axis.
The swing span defines, how much rotation will not get corrected along the swing axis.
Could be defined as looseness in the [ConeTwistJoint3D].
If below 0.05, this behavior is locked.*/
	ConeTwistJointSwingSpan ConeTwistJointParam = 0
/*Twist is the rotation around the twist axis, this value defined how far the joint can twist.
Twist is locked if below 0.05.*/
	ConeTwistJointTwistSpan ConeTwistJointParam = 1
/*The speed with which the swing or twist will take place.
The higher, the faster.*/
	ConeTwistJointBias ConeTwistJointParam = 2
/*The ease with which the Joint3D twists, if it's too low, it takes more force to twist the joint.*/
	ConeTwistJointSoftness ConeTwistJointParam = 3
/*Defines, how fast the swing- and twist-speed-difference on both sides gets synced.*/
	ConeTwistJointRelaxation ConeTwistJointParam = 4
)
type G6DOFJointAxisParam = classdb.PhysicsServer3DG6DOFJointAxisParam

const (
/*The minimum difference between the pivot points' axes.*/
	G6dofJointLinearLowerLimit G6DOFJointAxisParam = 0
/*The maximum difference between the pivot points' axes.*/
	G6dofJointLinearUpperLimit G6DOFJointAxisParam = 1
/*A factor that gets applied to the movement across the axes. The lower, the slower the movement.*/
	G6dofJointLinearLimitSoftness G6DOFJointAxisParam = 2
/*The amount of restitution on the axes movement. The lower, the more velocity-energy gets lost.*/
	G6dofJointLinearRestitution G6DOFJointAxisParam = 3
/*The amount of damping that happens at the linear motion across the axes.*/
	G6dofJointLinearDamping G6DOFJointAxisParam = 4
/*The velocity that the joint's linear motor will attempt to reach.*/
	G6dofJointLinearMotorTargetVelocity G6DOFJointAxisParam = 5
/*The maximum force that the linear motor can apply while trying to reach the target velocity.*/
	G6dofJointLinearMotorForceLimit G6DOFJointAxisParam = 6
	G6dofJointLinearSpringStiffness G6DOFJointAxisParam = 7
	G6dofJointLinearSpringDamping G6DOFJointAxisParam = 8
	G6dofJointLinearSpringEquilibriumPoint G6DOFJointAxisParam = 9
/*The minimum rotation in negative direction to break loose and rotate around the axes.*/
	G6dofJointAngularLowerLimit G6DOFJointAxisParam = 10
/*The minimum rotation in positive direction to break loose and rotate around the axes.*/
	G6dofJointAngularUpperLimit G6DOFJointAxisParam = 11
/*A factor that gets multiplied onto all rotations across the axes.*/
	G6dofJointAngularLimitSoftness G6DOFJointAxisParam = 12
/*The amount of rotational damping across the axes. The lower, the more damping occurs.*/
	G6dofJointAngularDamping G6DOFJointAxisParam = 13
/*The amount of rotational restitution across the axes. The lower, the more restitution occurs.*/
	G6dofJointAngularRestitution G6DOFJointAxisParam = 14
/*The maximum amount of force that can occur, when rotating around the axes.*/
	G6dofJointAngularForceLimit G6DOFJointAxisParam = 15
/*When correcting the crossing of limits in rotation across the axes, this error tolerance factor defines how much the correction gets slowed down. The lower, the slower.*/
	G6dofJointAngularErp G6DOFJointAxisParam = 16
/*Target speed for the motor at the axes.*/
	G6dofJointAngularMotorTargetVelocity G6DOFJointAxisParam = 17
/*Maximum acceleration for the motor at the axes.*/
	G6dofJointAngularMotorForceLimit G6DOFJointAxisParam = 18
	G6dofJointAngularSpringStiffness G6DOFJointAxisParam = 19
	G6dofJointAngularSpringDamping G6DOFJointAxisParam = 20
	G6dofJointAngularSpringEquilibriumPoint G6DOFJointAxisParam = 21
/*Represents the size of the [enum G6DOFJointAxisParam] enum.*/
	G6dofJointMax G6DOFJointAxisParam = 22
)
type G6DOFJointAxisFlag = classdb.PhysicsServer3DG6DOFJointAxisFlag

const (
/*If set, linear motion is possible within the given limits.*/
	G6dofJointFlagEnableLinearLimit G6DOFJointAxisFlag = 0
/*If set, rotational motion is possible.*/
	G6dofJointFlagEnableAngularLimit G6DOFJointAxisFlag = 1
	G6dofJointFlagEnableAngularSpring G6DOFJointAxisFlag = 2
	G6dofJointFlagEnableLinearSpring G6DOFJointAxisFlag = 3
/*If set, there is a rotational motor across these axes.*/
	G6dofJointFlagEnableMotor G6DOFJointAxisFlag = 4
/*If set, there is a linear motor on this axis that targets a specific velocity.*/
	G6dofJointFlagEnableLinearMotor G6DOFJointAxisFlag = 5
/*Represents the size of the [enum G6DOFJointAxisFlag] enum.*/
	G6dofJointFlagMax G6DOFJointAxisFlag = 6
)
type ShapeType = classdb.PhysicsServer3DShapeType

const (
/*The [Shape3D] is a [WorldBoundaryShape3D].*/
	ShapeWorldBoundary ShapeType = 0
/*The [Shape3D] is a [SeparationRayShape3D].*/
	ShapeSeparationRay ShapeType = 1
/*The [Shape3D] is a [SphereShape3D].*/
	ShapeSphere ShapeType = 2
/*The [Shape3D] is a [BoxShape3D].*/
	ShapeBox ShapeType = 3
/*The [Shape3D] is a [CapsuleShape3D].*/
	ShapeCapsule ShapeType = 4
/*The [Shape3D] is a [CylinderShape3D].*/
	ShapeCylinder ShapeType = 5
/*The [Shape3D] is a [ConvexPolygonShape3D].*/
	ShapeConvexPolygon ShapeType = 6
/*The [Shape3D] is a [ConcavePolygonShape3D].*/
	ShapeConcavePolygon ShapeType = 7
/*The [Shape3D] is a [HeightMapShape3D].*/
	ShapeHeightmap ShapeType = 8
/*The [Shape3D] is used internally for a soft body. Any attempt to create this kind of shape results in an error.*/
	ShapeSoftBody ShapeType = 9
/*This constant is used internally by the engine. Any attempt to create this kind of shape results in an error.*/
	ShapeCustom ShapeType = 10
)
type AreaParameter = classdb.PhysicsServer3DAreaParameter

const (
/*Constant to set/get gravity override mode in an area. See [enum AreaSpaceOverrideMode] for possible values.*/
	AreaParamGravityOverrideMode AreaParameter = 0
/*Constant to set/get gravity strength in an area.*/
	AreaParamGravity AreaParameter = 1
/*Constant to set/get gravity vector/center in an area.*/
	AreaParamGravityVector AreaParameter = 2
/*Constant to set/get whether the gravity vector of an area is a direction, or a center point.*/
	AreaParamGravityIsPoint AreaParameter = 3
/*Constant to set/get the distance at which the gravity strength is equal to the gravity controlled by [constant AREA_PARAM_GRAVITY]. For example, on a planet 100 meters in radius with a surface gravity of 4.0 m/s², set the gravity to 4.0 and the unit distance to 100.0. The gravity will have falloff according to the inverse square law, so in the example, at 200 meters from the center the gravity will be 1.0 m/s² (twice the distance, 1/4th the gravity), at 50 meters it will be 16.0 m/s² (half the distance, 4x the gravity), and so on.
The above is true only when the unit distance is a positive number. When this is set to 0.0, the gravity will be constant regardless of distance.*/
	AreaParamGravityPointUnitDistance AreaParameter = 4
/*Constant to set/get linear damping override mode in an area. See [enum AreaSpaceOverrideMode] for possible values.*/
	AreaParamLinearDampOverrideMode AreaParameter = 5
/*Constant to set/get the linear damping factor of an area.*/
	AreaParamLinearDamp AreaParameter = 6
/*Constant to set/get angular damping override mode in an area. See [enum AreaSpaceOverrideMode] for possible values.*/
	AreaParamAngularDampOverrideMode AreaParameter = 7
/*Constant to set/get the angular damping factor of an area.*/
	AreaParamAngularDamp AreaParameter = 8
/*Constant to set/get the priority (order of processing) of an area.*/
	AreaParamPriority AreaParameter = 9
/*Constant to set/get the magnitude of area-specific wind force. This wind force only applies to [SoftBody3D] nodes. Other physics bodies are currently not affected by wind.*/
	AreaParamWindForceMagnitude AreaParameter = 10
/*Constant to set/get the 3D vector that specifies the origin from which an area-specific wind blows.*/
	AreaParamWindSource AreaParameter = 11
/*Constant to set/get the 3D vector that specifies the direction in which an area-specific wind blows.*/
	AreaParamWindDirection AreaParameter = 12
/*Constant to set/get the exponential rate at which wind force decreases with distance from its origin.*/
	AreaParamWindAttenuationFactor AreaParameter = 13
)
type AreaSpaceOverrideMode = classdb.PhysicsServer3DAreaSpaceOverrideMode

const (
/*This area does not affect gravity/damp. These are generally areas that exist only to detect collisions, and objects entering or exiting them.*/
	AreaSpaceOverrideDisabled AreaSpaceOverrideMode = 0
/*This area adds its gravity/damp values to whatever has been calculated so far. This way, many overlapping areas can combine their physics to make interesting effects.*/
	AreaSpaceOverrideCombine AreaSpaceOverrideMode = 1
/*This area adds its gravity/damp values to whatever has been calculated so far. Then stops taking into account the rest of the areas, even the default one.*/
	AreaSpaceOverrideCombineReplace AreaSpaceOverrideMode = 2
/*This area replaces any gravity/damp, even the default one, and stops taking into account the rest of the areas.*/
	AreaSpaceOverrideReplace AreaSpaceOverrideMode = 3
/*This area replaces any gravity/damp calculated so far, but keeps calculating the rest of the areas, down to the default one.*/
	AreaSpaceOverrideReplaceCombine AreaSpaceOverrideMode = 4
)
type BodyMode = classdb.PhysicsServer3DBodyMode

const (
/*Constant for static bodies. In this mode, a body can be only moved by user code and doesn't collide with other bodies along its path when moved.*/
	BodyModeStatic BodyMode = 0
/*Constant for kinematic bodies. In this mode, a body can be only moved by user code and collides with other bodies along its path.*/
	BodyModeKinematic BodyMode = 1
/*Constant for rigid bodies. In this mode, a body can be pushed by other bodies and has forces applied.*/
	BodyModeRigid BodyMode = 2
/*Constant for linear rigid bodies. In this mode, a body can not rotate, and only its linear velocity is affected by external forces.*/
	BodyModeRigidLinear BodyMode = 3
)
type BodyParameter = classdb.PhysicsServer3DBodyParameter

const (
/*Constant to set/get a body's bounce factor.*/
	BodyParamBounce BodyParameter = 0
/*Constant to set/get a body's friction.*/
	BodyParamFriction BodyParameter = 1
/*Constant to set/get a body's mass.*/
	BodyParamMass BodyParameter = 2
/*Constant to set/get a body's inertia.*/
	BodyParamInertia BodyParameter = 3
/*Constant to set/get a body's center of mass position in the body's local coordinate system.*/
	BodyParamCenterOfMass BodyParameter = 4
/*Constant to set/get a body's gravity multiplier.*/
	BodyParamGravityScale BodyParameter = 5
/*Constant to set/get a body's linear damping mode. See [enum BodyDampMode] for possible values.*/
	BodyParamLinearDampMode BodyParameter = 6
/*Constant to set/get a body's angular damping mode. See [enum BodyDampMode] for possible values.*/
	BodyParamAngularDampMode BodyParameter = 7
/*Constant to set/get a body's linear damping factor.*/
	BodyParamLinearDamp BodyParameter = 8
/*Constant to set/get a body's angular damping factor.*/
	BodyParamAngularDamp BodyParameter = 9
/*Represents the size of the [enum BodyParameter] enum.*/
	BodyParamMax BodyParameter = 10
)
type BodyDampMode = classdb.PhysicsServer3DBodyDampMode

const (
/*The body's damping value is added to any value set in areas or the default value.*/
	BodyDampModeCombine BodyDampMode = 0
/*The body's damping value replaces any value set in areas or the default value.*/
	BodyDampModeReplace BodyDampMode = 1
)
type BodyState = classdb.PhysicsServer3DBodyState

const (
/*Constant to set/get the current transform matrix of the body.*/
	BodyStateTransform BodyState = 0
/*Constant to set/get the current linear velocity of the body.*/
	BodyStateLinearVelocity BodyState = 1
/*Constant to set/get the current angular velocity of the body.*/
	BodyStateAngularVelocity BodyState = 2
/*Constant to sleep/wake up a body, or to get whether it is sleeping.*/
	BodyStateSleeping BodyState = 3
/*Constant to set/get whether the body can sleep.*/
	BodyStateCanSleep BodyState = 4
)
type AreaBodyStatus = classdb.PhysicsServer3DAreaBodyStatus

const (
/*The value of the first parameter and area callback function receives, when an object enters one of its shapes.*/
	AreaBodyAdded AreaBodyStatus = 0
/*The value of the first parameter and area callback function receives, when an object exits one of its shapes.*/
	AreaBodyRemoved AreaBodyStatus = 1
)
type ProcessInfo = classdb.PhysicsServer3DProcessInfo

const (
/*Constant to get the number of objects that are not sleeping.*/
	InfoActiveObjects ProcessInfo = 0
/*Constant to get the number of possible collisions.*/
	InfoCollisionPairs ProcessInfo = 1
/*Constant to get the number of space regions where a collision could occur.*/
	InfoIslandCount ProcessInfo = 2
)
type SpaceParameter = classdb.PhysicsServer3DSpaceParameter

const (
/*Constant to set/get the maximum distance a pair of bodies has to move before their collision status has to be recalculated.*/
	SpaceParamContactRecycleRadius SpaceParameter = 0
/*Constant to set/get the maximum distance a shape can be from another before they are considered separated and the contact is discarded.*/
	SpaceParamContactMaxSeparation SpaceParameter = 1
/*Constant to set/get the maximum distance a shape can penetrate another shape before it is considered a collision.*/
	SpaceParamContactMaxAllowedPenetration SpaceParameter = 2
/*Constant to set/get the default solver bias for all physics contacts. A solver bias is a factor controlling how much two objects "rebound", after overlapping, to avoid leaving them in that state because of numerical imprecision.*/
	SpaceParamContactDefaultBias SpaceParameter = 3
/*Constant to set/get the threshold linear velocity of activity. A body marked as potentially inactive for both linear and angular velocity will be put to sleep after the time given.*/
	SpaceParamBodyLinearVelocitySleepThreshold SpaceParameter = 4
/*Constant to set/get the threshold angular velocity of activity. A body marked as potentially inactive for both linear and angular velocity will be put to sleep after the time given.*/
	SpaceParamBodyAngularVelocitySleepThreshold SpaceParameter = 5
/*Constant to set/get the maximum time of activity. A body marked as potentially inactive for both linear and angular velocity will be put to sleep after this time.*/
	SpaceParamBodyTimeToSleep SpaceParameter = 6
/*Constant to set/get the number of solver iterations for contacts and constraints. The greater the number of iterations, the more accurate the collisions and constraints will be. However, a greater number of iterations requires more CPU power, which can decrease performance.*/
	SpaceParamSolverIterations SpaceParameter = 7
)
type BodyAxis = classdb.PhysicsServer3DBodyAxis

const (
	BodyAxisLinearX BodyAxis = 1
	BodyAxisLinearY BodyAxis = 2
	BodyAxisLinearZ BodyAxis = 4
	BodyAxisAngularX BodyAxis = 8
	BodyAxisAngularY BodyAxis = 16
	BodyAxisAngularZ BodyAxis = 32
)
