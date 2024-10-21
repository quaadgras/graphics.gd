package PhysicsServer2D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
PhysicsServer2D is the server responsible for all 2D physics. It can directly create and manipulate all physics objects:
- A [i]space[/i] is a self-contained world for a physics simulation. It contains bodies, areas, and joints. Its state can be queried for collision and intersection information, and several parameters of the simulation can be modified.
- A [i]shape[/i] is a geometric shape such as a circle, a rectangle, a capsule, or a polygon. It can be used for collision detection by adding it to a body/area, possibly with an extra transformation relative to the body/area's origin. Bodies/areas can have multiple (transformed) shapes added to them, and a single shape can be added to bodies/areas multiple times with different local transformations.
- A [i]body[/i] is a physical object which can be in static, kinematic, or rigid mode. Its state (such as position and velocity) can be queried and updated. A force integration callback can be set to customize the body's physics.
- An [i]area[/i] is a region in space which can be used to detect bodies and areas entering and exiting it. A body monitoring callback can be set to report entering/exiting body shapes, and similarly an area monitoring callback can be set. Gravity and damping can be overridden within the area by setting area parameters.
- A [i]joint[/i] is a constraint, either between two bodies or on one body relative to a point. Parameters such as the joint bias and the rest length of a spring joint can be adjusted.
Physics objects in [PhysicsServer2D] may be created and manipulated independently; they do not have to be tied to nodes in the scene tree.
[b]Note:[/b] All the 2D physics nodes use the physics server internally. Adding a physics node to the scene tree will cause a corresponding physics object to be created in the physics server. A rigid body node registers a callback that updates the node's transform with the transform of the respective body object in the physics server (every physics update). An area node registers a callback to inform the area node about overlaps with the respective area object in the physics server. The raycast node queries the direct state of the relevant space in the physics server.

*/
type Simple [1]classdb.PhysicsServer2D
func (self Simple) WorldBoundaryShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).WorldBoundaryShapeCreate())
}
func (self Simple) SeparationRayShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SeparationRayShapeCreate())
}
func (self Simple) SegmentShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SegmentShapeCreate())
}
func (self Simple) CircleShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CircleShapeCreate())
}
func (self Simple) RectangleShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).RectangleShapeCreate())
}
func (self Simple) CapsuleShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CapsuleShapeCreate())
}
func (self Simple) ConvexPolygonShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ConvexPolygonShapeCreate())
}
func (self Simple) ConcavePolygonShapeCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ConcavePolygonShapeCreate())
}
func (self Simple) ShapeSetData(shape gd.RID, data gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ShapeSetData(shape, data)
}
func (self Simple) ShapeGetType(shape gd.RID) classdb.PhysicsServer2DShapeType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicsServer2DShapeType(Expert(self).ShapeGetType(shape))
}
func (self Simple) ShapeGetData(shape gd.RID) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).ShapeGetData(gc, shape))
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
func (self Simple) SpaceSetParam(space gd.RID, param classdb.PhysicsServer2DSpaceParameter, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SpaceSetParam(space, param, gd.Float(value))
}
func (self Simple) SpaceGetParam(space gd.RID, param classdb.PhysicsServer2DSpaceParameter) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).SpaceGetParam(space, param)))
}
func (self Simple) SpaceGetDirectState(space gd.RID) [1]classdb.PhysicsDirectSpaceState2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PhysicsDirectSpaceState2D(Expert(self).SpaceGetDirectState(gc, space))
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
func (self Simple) AreaAddShape(area gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaAddShape(area, shape, transform, disabled)
}
func (self Simple) AreaSetShape(area gd.RID, shape_idx int, shape gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetShape(area, gd.Int(shape_idx), shape)
}
func (self Simple) AreaSetShapeTransform(area gd.RID, shape_idx int, transform gd.Transform2D) {
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
func (self Simple) AreaGetShapeTransform(area gd.RID, shape_idx int) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).AreaGetShapeTransform(area, gd.Int(shape_idx)))
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
func (self Simple) AreaSetParam(area gd.RID, param classdb.PhysicsServer2DAreaParameter, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetParam(area, param, value)
}
func (self Simple) AreaSetTransform(area gd.RID, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaSetTransform(area, transform)
}
func (self Simple) AreaGetParam(area gd.RID, param classdb.PhysicsServer2DAreaParameter) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).AreaGetParam(gc, area, param))
}
func (self Simple) AreaGetTransform(area gd.RID) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).AreaGetTransform(area))
}
func (self Simple) AreaAttachObjectInstanceId(area gd.RID, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaAttachObjectInstanceId(area, gd.Int(id))
}
func (self Simple) AreaGetObjectInstanceId(area gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AreaGetObjectInstanceId(area)))
}
func (self Simple) AreaAttachCanvasInstanceId(area gd.RID, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AreaAttachCanvasInstanceId(area, gd.Int(id))
}
func (self Simple) AreaGetCanvasInstanceId(area gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AreaGetCanvasInstanceId(area)))
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
func (self Simple) BodySetMode(body gd.RID, mode classdb.PhysicsServer2DBodyMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetMode(body, mode)
}
func (self Simple) BodyGetMode(body gd.RID) classdb.PhysicsServer2DBodyMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicsServer2DBodyMode(Expert(self).BodyGetMode(body))
}
func (self Simple) BodyAddShape(body gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAddShape(body, shape, transform, disabled)
}
func (self Simple) BodySetShape(body gd.RID, shape_idx int, shape gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetShape(body, gd.Int(shape_idx), shape)
}
func (self Simple) BodySetShapeTransform(body gd.RID, shape_idx int, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetShapeTransform(body, gd.Int(shape_idx), transform)
}
func (self Simple) BodyGetShapeCount(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).BodyGetShapeCount(body)))
}
func (self Simple) BodyGetShape(body gd.RID, shape_idx int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).BodyGetShape(body, gd.Int(shape_idx)))
}
func (self Simple) BodyGetShapeTransform(body gd.RID, shape_idx int) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).BodyGetShapeTransform(body, gd.Int(shape_idx)))
}
func (self Simple) BodyRemoveShape(body gd.RID, shape_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyRemoveShape(body, gd.Int(shape_idx))
}
func (self Simple) BodyClearShapes(body gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyClearShapes(body)
}
func (self Simple) BodySetShapeDisabled(body gd.RID, shape_idx int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetShapeDisabled(body, gd.Int(shape_idx), disabled)
}
func (self Simple) BodySetShapeAsOneWayCollision(body gd.RID, shape_idx int, enable bool, margin float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetShapeAsOneWayCollision(body, gd.Int(shape_idx), enable, gd.Float(margin))
}
func (self Simple) BodyAttachObjectInstanceId(body gd.RID, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAttachObjectInstanceId(body, gd.Int(id))
}
func (self Simple) BodyGetObjectInstanceId(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).BodyGetObjectInstanceId(body)))
}
func (self Simple) BodyAttachCanvasInstanceId(body gd.RID, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAttachCanvasInstanceId(body, gd.Int(id))
}
func (self Simple) BodyGetCanvasInstanceId(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).BodyGetCanvasInstanceId(body)))
}
func (self Simple) BodySetContinuousCollisionDetectionMode(body gd.RID, mode classdb.PhysicsServer2DCCDMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetContinuousCollisionDetectionMode(body, mode)
}
func (self Simple) BodyGetContinuousCollisionDetectionMode(body gd.RID) classdb.PhysicsServer2DCCDMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicsServer2DCCDMode(Expert(self).BodyGetContinuousCollisionDetectionMode(body))
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
func (self Simple) BodySetParam(body gd.RID, param classdb.PhysicsServer2DBodyParameter, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetParam(body, param, value)
}
func (self Simple) BodyGetParam(body gd.RID, param classdb.PhysicsServer2DBodyParameter) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).BodyGetParam(gc, body, param))
}
func (self Simple) BodyResetMassProperties(body gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyResetMassProperties(body)
}
func (self Simple) BodySetState(body gd.RID, state classdb.PhysicsServer2DBodyState, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetState(body, state, value)
}
func (self Simple) BodyGetState(body gd.RID, state classdb.PhysicsServer2DBodyState) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).BodyGetState(gc, body, state))
}
func (self Simple) BodyApplyCentralImpulse(body gd.RID, impulse gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyCentralImpulse(body, impulse)
}
func (self Simple) BodyApplyTorqueImpulse(body gd.RID, impulse float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyTorqueImpulse(body, gd.Float(impulse))
}
func (self Simple) BodyApplyImpulse(body gd.RID, impulse gd.Vector2, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyImpulse(body, impulse, position)
}
func (self Simple) BodyApplyCentralForce(body gd.RID, force gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyCentralForce(body, force)
}
func (self Simple) BodyApplyForce(body gd.RID, force gd.Vector2, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyForce(body, force, position)
}
func (self Simple) BodyApplyTorque(body gd.RID, torque float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyApplyTorque(body, gd.Float(torque))
}
func (self Simple) BodyAddConstantCentralForce(body gd.RID, force gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAddConstantCentralForce(body, force)
}
func (self Simple) BodyAddConstantForce(body gd.RID, force gd.Vector2, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAddConstantForce(body, force, position)
}
func (self Simple) BodyAddConstantTorque(body gd.RID, torque float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodyAddConstantTorque(body, gd.Float(torque))
}
func (self Simple) BodySetConstantForce(body gd.RID, force gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetConstantForce(body, force)
}
func (self Simple) BodyGetConstantForce(body gd.RID) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).BodyGetConstantForce(body))
}
func (self Simple) BodySetConstantTorque(body gd.RID, torque float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetConstantTorque(body, gd.Float(torque))
}
func (self Simple) BodyGetConstantTorque(body gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).BodyGetConstantTorque(body)))
}
func (self Simple) BodySetAxisVelocity(body gd.RID, axis_velocity gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BodySetAxisVelocity(body, axis_velocity)
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
func (self Simple) BodyTestMotion(body gd.RID, parameters [1]classdb.PhysicsTestMotionParameters2D, result [1]classdb.PhysicsTestMotionResult2D) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).BodyTestMotion(body, parameters, result))
}
func (self Simple) BodyGetDirectState(body gd.RID) [1]classdb.PhysicsDirectBodyState2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PhysicsDirectBodyState2D(Expert(self).BodyGetDirectState(gc, body))
}
func (self Simple) JointCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).JointCreate())
}
func (self Simple) JointClear(joint gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointClear(joint)
}
func (self Simple) JointSetParam(joint gd.RID, param classdb.PhysicsServer2DJointParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointSetParam(joint, param, gd.Float(value))
}
func (self Simple) JointGetParam(joint gd.RID, param classdb.PhysicsServer2DJointParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).JointGetParam(joint, param)))
}
func (self Simple) JointDisableCollisionsBetweenBodies(joint gd.RID, disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointDisableCollisionsBetweenBodies(joint, disable)
}
func (self Simple) JointIsDisabledCollisionsBetweenBodies(joint gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).JointIsDisabledCollisionsBetweenBodies(joint))
}
func (self Simple) JointMakePin(joint gd.RID, anchor gd.Vector2, body_a gd.RID, body_b gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointMakePin(joint, anchor, body_a, body_b)
}
func (self Simple) JointMakeGroove(joint gd.RID, groove1_a gd.Vector2, groove2_a gd.Vector2, anchor_b gd.Vector2, body_a gd.RID, body_b gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointMakeGroove(joint, groove1_a, groove2_a, anchor_b, body_a, body_b)
}
func (self Simple) JointMakeDampedSpring(joint gd.RID, anchor_a gd.Vector2, anchor_b gd.Vector2, body_a gd.RID, body_b gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).JointMakeDampedSpring(joint, anchor_a, anchor_b, body_a, body_b)
}
func (self Simple) PinJointSetFlag(joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PinJointSetFlag(joint, flag, enabled)
}
func (self Simple) PinJointGetFlag(joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).PinJointGetFlag(joint, flag))
}
func (self Simple) PinJointSetParam(joint gd.RID, param classdb.PhysicsServer2DPinJointParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PinJointSetParam(joint, param, gd.Float(value))
}
func (self Simple) PinJointGetParam(joint gd.RID, param classdb.PhysicsServer2DPinJointParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).PinJointGetParam(joint, param)))
}
func (self Simple) DampedSpringJointSetParam(joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DampedSpringJointSetParam(joint, param, gd.Float(value))
}
func (self Simple) DampedSpringJointGetParam(joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).DampedSpringJointGetParam(joint, param)))
}
func (self Simple) JointGetType(joint gd.RID) classdb.PhysicsServer2DJointType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PhysicsServer2DJointType(Expert(self).JointGetType(joint))
}
func (self Simple) FreeRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FreeRid(rid)
}
func (self Simple) SetActive(active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetActive(active)
}
func (self Simple) GetProcessInfo(process_info classdb.PhysicsServer2DProcessInfo) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetProcessInfo(process_info)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicsServer2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a 2D world boundary shape in the physics server, and returns the [RID] that identifies it. Use [method shape_set_data] to set the shape's normal direction and distance properties.
*/
//go:nosplit
func (self class) WorldBoundaryShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_world_boundary_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a 2D separation ray shape in the physics server, and returns the [RID] that identifies it. Use [method shape_set_data] to set the shape's [code]length[/code] and [code]slide_on_slope[/code] properties.
*/
//go:nosplit
func (self class) SeparationRayShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_separation_ray_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a 2D segment shape in the physics server, and returns the [RID] that identifies it. Use [method shape_set_data] to set the segment's start and end points.
*/
//go:nosplit
func (self class) SegmentShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_segment_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a 2D circle shape in the physics server, and returns the [RID] that identifies it. Use [method shape_set_data] to set the circle's radius.
*/
//go:nosplit
func (self class) CircleShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_circle_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a 2D rectangle shape in the physics server, and returns the [RID] that identifies it. Use [method shape_set_data] to set the rectangle's half-extents.
*/
//go:nosplit
func (self class) RectangleShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_rectangle_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a 2D capsule shape in the physics server, and returns the [RID] that identifies it. Use [method shape_set_data] to set the capsule's height and radius.
*/
//go:nosplit
func (self class) CapsuleShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_capsule_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a 2D convex polygon shape in the physics server, and returns the [RID] that identifies it. Use [method shape_set_data] to set the convex polygon's points.
*/
//go:nosplit
func (self class) ConvexPolygonShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_convex_polygon_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a 2D concave polygon shape in the physics server, and returns the [RID] that identifies it. Use [method shape_set_data] to set the concave polygon's segments.
*/
//go:nosplit
func (self class) ConcavePolygonShapeCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_concave_polygon_shape_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the shape data that defines the configuration of the shape. The [param data] to be passed depends on the shape's type (see [method shape_get_type]):
- [constant SHAPE_WORLD_BOUNDARY]: an array of length two containing a [Vector2] [code]normal[/code] direction and a [float] distance [code]d[/code],
- [constant SHAPE_SEPARATION_RAY]: a dictionary containing the key [code]length[/code] with a [float] value and the key [code]slide_on_slope[/code] with a [bool] value,
- [constant SHAPE_SEGMENT]: a [Rect2] [code]rect[/code] containing the first point of the segment in [code]rect.position[/code] and the second point of the segment in [code]rect.size[/code],
- [constant SHAPE_CIRCLE]: a [float] [code]radius[/code],
- [constant SHAPE_RECTANGLE]: a [Vector2] [code]half_extents[/code],
- [constant SHAPE_CAPSULE]: an array of length two (or a [Vector2]) containing a [float] [code]height[/code] and a [float] [code]radius[/code],
- [constant SHAPE_CONVEX_POLYGON]: either a [PackedVector2Array] of points defining a convex polygon in counterclockwise order (the clockwise outward normal of each segment formed by consecutive points is calculated internally), or a [PackedFloat32Array] of length divisible by four so that every 4-tuple of [float]s contains the coordinates of a point followed by the coordinates of the clockwise outward normal vector to the segment between the current point and the next point,
- [constant SHAPE_CONCAVE_POLYGON]: a [PackedVector2Array] of length divisible by two (each pair of points forms one segment).
[b]Warning:[/b] In the case of [constant SHAPE_CONVEX_POLYGON], this method does not check if the points supplied actually form a convex polygon (unlike the [member CollisionPolygon2D.polygon] property).
*/
//go:nosplit
func (self class) ShapeSetData(shape gd.RID, data gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_shape_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the shape's type (see [enum ShapeType]).
*/
//go:nosplit
func (self class) ShapeGetType(shape gd.RID) classdb.PhysicsServer2DShapeType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret = callframe.Ret[classdb.PhysicsServer2DShapeType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_shape_get_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the shape data that defines the configuration of the shape, such as the half-extents of a rectangle or the segments of a concave shape. See [method shape_set_data] for the precise format of this data in each case.
*/
//go:nosplit
func (self class) ShapeGetData(ctx gd.Lifetime, shape gd.RID) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_shape_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates a 2D space in the physics server, and returns the [RID] that identifies it. A space contains bodies and areas, and controls the stepping of the physics simulation of the objects in it.
*/
//go:nosplit
func (self class) SpaceCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_space_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Activates or deactivates the space. If [param active] is [code]false[/code], then the physics server will not do anything with this space in its physics step.
*/
//go:nosplit
func (self class) SpaceSetActive(space gd.RID, active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_space_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the space is active.
*/
//go:nosplit
func (self class) SpaceIsActive(space gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_space_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of the given space parameter. See [enum SpaceParameter] for the list of available parameters.
*/
//go:nosplit
func (self class) SpaceSetParam(space gd.RID, param classdb.PhysicsServer2DSpaceParameter, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_space_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the given space parameter. See [enum SpaceParameter] for the list of available parameters.
*/
//go:nosplit
func (self class) SpaceGetParam(space gd.RID, param classdb.PhysicsServer2DSpaceParameter) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_space_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the state of a space, a [PhysicsDirectSpaceState2D]. This object can be used for collision/intersection queries.
*/
//go:nosplit
func (self class) SpaceGetDirectState(ctx gd.Lifetime, space gd.RID) object.PhysicsDirectSpaceState2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_space_get_direct_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PhysicsDirectSpaceState2D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a 2D area object in the physics server, and returns the [RID] that identifies it. The default settings for the created area include a collision layer and mask set to [code]1[/code], and [code]monitorable[/code] set to [code]false[/code].
Use [method area_add_shape] to add shapes to it, use [method area_set_transform] to set its transform, and use [method area_set_space] to add the area to a space. If you want the area to be detectable use [method area_set_monitorable].
*/
//go:nosplit
func (self class) AreaCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds the area to the given space, after removing the area from the previously assigned space (if any).
[b]Note:[/b] To remove an area from a space without immediately adding it back elsewhere, use [code]PhysicsServer2D.area_set_space(area, RID())[/code].
*/
//go:nosplit
func (self class) AreaSetSpace(area gd.RID, space gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [RID] of the space assigned to the area. Returns an empty [RID] if no space is assigned.
*/
//go:nosplit
func (self class) AreaGetSpace(area gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a shape to the area, with the given local transform. The shape (together with its [param transform] and [param disabled] properties) is added to an array of shapes, and the shapes of an area are usually referenced by their index in this array.
*/
//go:nosplit
func (self class) AreaAddShape(area gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape)
	callframe.Arg(frame, transform)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_add_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Replaces the area's shape at the given index by another shape, while not affecting the [code]transform[/code] and [code]disabled[/code] properties at the same index.
*/
//go:nosplit
func (self class) AreaSetShape(area gd.RID, shape_idx gd.Int, shape gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the local transform matrix of the area's shape with the given index.
*/
//go:nosplit
func (self class) AreaSetShapeTransform(area gd.RID, shape_idx gd.Int, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_shape_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the disabled property of the area's shape with the given index. If [param disabled] is [code]true[/code], then the shape will not detect any other shapes entering or exiting it.
*/
//go:nosplit
func (self class) AreaSetShapeDisabled(area gd.RID, shape_idx gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_shape_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of shapes added to the area.
*/
//go:nosplit
func (self class) AreaGetShapeCount(area gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [RID] of the shape with the given index in the area's array of shapes.
*/
//go:nosplit
func (self class) AreaGetShape(area gd.RID, shape_idx gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the local transform matrix of the shape with the given index in the area's array of shapes.
*/
//go:nosplit
func (self class) AreaGetShapeTransform(area gd.RID, shape_idx gd.Int) gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_shape_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the shape with the given index from the area's array of shapes. The shape itself is not deleted, so it can continue to be used elsewhere or added back later. As a result of this operation, the area's shapes which used to have indices higher than [param shape_idx] will have their index decreased by one.
*/
//go:nosplit
func (self class) AreaRemoveShape(area gd.RID, shape_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, shape_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_remove_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all shapes from the area. This does not delete the shapes themselves, so they can continue to be used elsewhere or added back later.
*/
//go:nosplit
func (self class) AreaClearShapes(area gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_clear_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Assigns the area to one or many physics layers, via a bitmask.
*/
//go:nosplit
func (self class) AreaSetCollisionLayer(area gd.RID, layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers the area belongs to, as a bitmask.
*/
//go:nosplit
func (self class) AreaGetCollisionLayer(area gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets which physics layers the area will monitor, via a bitmask.
*/
//go:nosplit
func (self class) AreaSetCollisionMask(area gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers the area can contact with, as a bitmask.
*/
//go:nosplit
func (self class) AreaGetCollisionMask(area gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of the given area parameter. See [enum AreaParameter] for the list of available parameters.
*/
//go:nosplit
func (self class) AreaSetParam(area gd.RID, param classdb.PhysicsServer2DAreaParameter, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, param)
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the transform matrix of the area.
*/
//go:nosplit
func (self class) AreaSetTransform(area gd.RID, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the given area parameter. See [enum AreaParameter] for the list of available parameters.
*/
//go:nosplit
func (self class) AreaGetParam(ctx gd.Lifetime, area gd.RID, param classdb.PhysicsServer2DAreaParameter) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the transform matrix of the area.
*/
//go:nosplit
func (self class) AreaGetTransform(area gd.RID) gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Attaches the [code]ObjectID[/code] of an [Object] to the area. Use [method Object.get_instance_id] to get the [code]ObjectID[/code] of a [CollisionObject2D].
*/
//go:nosplit
func (self class) AreaAttachObjectInstanceId(area gd.RID, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_attach_object_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]ObjectID[/code] attached to the area. Use [method @GlobalScope.instance_from_id] to retrieve an [Object] from a nonzero [code]ObjectID[/code].
*/
//go:nosplit
func (self class) AreaGetObjectInstanceId(area gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_object_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Attaches the [code]ObjectID[/code] of a canvas to the area. Use [method Object.get_instance_id] to get the [code]ObjectID[/code] of a [CanvasLayer].
*/
//go:nosplit
func (self class) AreaAttachCanvasInstanceId(area gd.RID, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_attach_canvas_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]ObjectID[/code] of the canvas attached to the area. Use [method @GlobalScope.instance_from_id] to retrieve a [CanvasLayer] from a nonzero [code]ObjectID[/code].
*/
//go:nosplit
func (self class) AreaGetCanvasInstanceId(area gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_get_canvas_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_monitor_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_area_monitor_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether the area is monitorable or not. If [param monitorable] is [code]true[/code], the area monitoring callback of other areas will be called when this area enters or exits them.
*/
//go:nosplit
func (self class) AreaSetMonitorable(area gd.RID, monitorable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, area)
	callframe.Arg(frame, monitorable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_area_set_monitorable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a 2D body object in the physics server, and returns the [RID] that identifies it. The default settings for the created area include a collision layer and mask set to [code]1[/code], and body mode set to [constant BODY_MODE_RIGID].
Use [method body_add_shape] to add shapes to it, use [method body_set_state] to set its transform, and use [method body_set_space] to add the body to a space.
*/
//go:nosplit
func (self class) BodyCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds the body to the given space, after removing the body from the previously assigned space (if any). If the body's mode is set to [constant BODY_MODE_RIGID], then adding the body to a space will have the following additional effects:
- If the parameter [constant BODY_PARAM_CENTER_OF_MASS] has never been set explicitly, then the value of that parameter will be recalculated based on the body's shapes.
- If the parameter [constant BODY_PARAM_INERTIA] is set to a value [code]<= 0.0[/code], then the value of that parameter will be recalculated based on the body's shapes, mass, and center of mass.
[b]Note:[/b] To remove a body from a space without immediately adding it back elsewhere, use [code]PhysicsServer2D.body_set_space(body, RID())[/code].
*/
//go:nosplit
func (self class) BodySetSpace(body gd.RID, space gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [RID] of the space assigned to the body. Returns an empty [RID] if no space is assigned.
*/
//go:nosplit
func (self class) BodyGetSpace(body gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the body's mode. See [enum BodyMode] for the list of available modes.
*/
//go:nosplit
func (self class) BodySetMode(body gd.RID, mode classdb.PhysicsServer2DBodyMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the body's mode (see [enum BodyMode]).
*/
//go:nosplit
func (self class) BodyGetMode(body gd.RID) classdb.PhysicsServer2DBodyMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[classdb.PhysicsServer2DBodyMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a shape to the area, with the given local transform. The shape (together with its [param transform] and [param disabled] properties) is added to an array of shapes, and the shapes of a body are usually referenced by their index in this array.
*/
//go:nosplit
func (self class) BodyAddShape(body gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape)
	callframe.Arg(frame, transform)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_add_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Replaces the body's shape at the given index by another shape, while not affecting the [code]transform[/code], [code]disabled[/code], and one-way collision properties at the same index.
*/
//go:nosplit
func (self class) BodySetShape(body gd.RID, shape_idx gd.Int, shape gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the local transform matrix of the body's shape with the given index.
*/
//go:nosplit
func (self class) BodySetShapeTransform(body gd.RID, shape_idx gd.Int, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_shape_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of shapes added to the body.
*/
//go:nosplit
func (self class) BodyGetShapeCount(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [RID] of the shape with the given index in the body's array of shapes.
*/
//go:nosplit
func (self class) BodyGetShape(body gd.RID, shape_idx gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the local transform matrix of the shape with the given index in the area's array of shapes.
*/
//go:nosplit
func (self class) BodyGetShapeTransform(body gd.RID, shape_idx gd.Int) gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_shape_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the shape with the given index from the body's array of shapes. The shape itself is not deleted, so it can continue to be used elsewhere or added back later. As a result of this operation, the body's shapes which used to have indices higher than [param shape_idx] will have their index decreased by one.
*/
//go:nosplit
func (self class) BodyRemoveShape(body gd.RID, shape_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_remove_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all shapes from the body. This does not delete the shapes themselves, so they can continue to be used elsewhere or added back later.
*/
//go:nosplit
func (self class) BodyClearShapes(body gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_clear_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the disabled property of the body's shape with the given index. If [param disabled] is [code]true[/code], then the shape will be ignored in all collision detection.
*/
//go:nosplit
func (self class) BodySetShapeDisabled(body gd.RID, shape_idx gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_shape_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the one-way collision properties of the body's shape with the given index. If [param enable] is [code]true[/code], the one-way collision direction given by the shape's local upward axis [code]body_get_shape_transform(body, shape_idx).y[/code] will be used to ignore collisions with the shape in the opposite direction, and to ensure depenetration of kinematic bodies happens in this direction.
*/
//go:nosplit
func (self class) BodySetShapeAsOneWayCollision(body gd.RID, shape_idx gd.Int, enable bool, margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, shape_idx)
	callframe.Arg(frame, enable)
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_shape_as_one_way_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Attaches the [code]ObjectID[/code] of an [Object] to the body. Use [method Object.get_instance_id] to get the [code]ObjectID[/code] of a [CollisionObject2D].
*/
//go:nosplit
func (self class) BodyAttachObjectInstanceId(body gd.RID, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_attach_object_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]ObjectID[/code] attached to the body. Use [method @GlobalScope.instance_from_id] to retrieve an [Object] from a nonzero [code]ObjectID[/code].
*/
//go:nosplit
func (self class) BodyGetObjectInstanceId(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_object_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Attaches the [code]ObjectID[/code] of a canvas to the body. Use [method Object.get_instance_id] to get the [code]ObjectID[/code] of a [CanvasLayer].
*/
//go:nosplit
func (self class) BodyAttachCanvasInstanceId(body gd.RID, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_attach_canvas_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]ObjectID[/code] of the canvas attached to the body. Use [method @GlobalScope.instance_from_id] to retrieve a [CanvasLayer] from a nonzero [code]ObjectID[/code].
*/
//go:nosplit
func (self class) BodyGetCanvasInstanceId(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_canvas_instance_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the continuous collision detection mode using one of the [enum CCDMode] constants.
Continuous collision detection tries to predict where a moving body would collide in between physics updates, instead of moving it and correcting its movement if it collided.
*/
//go:nosplit
func (self class) BodySetContinuousCollisionDetectionMode(body gd.RID, mode classdb.PhysicsServer2DCCDMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_continuous_collision_detection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the body's continuous collision detection mode (see [enum CCDMode]).
*/
//go:nosplit
func (self class) BodyGetContinuousCollisionDetectionMode(body gd.RID) classdb.PhysicsServer2DCCDMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[classdb.PhysicsServer2DCCDMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_continuous_collision_detection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the physics layer or layers the body belongs to, via a bitmask.
*/
//go:nosplit
func (self class) BodySetCollisionLayer(body gd.RID, layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers the body belongs to, as a bitmask.
*/
//go:nosplit
func (self class) BodyGetCollisionLayer(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the physics layer or layers the body can collide with, via a bitmask.
*/
//go:nosplit
func (self class) BodySetCollisionMask(body gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics layer or layers the body can collide with, as a bitmask.
*/
//go:nosplit
func (self class) BodyGetCollisionMask(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the body's collision priority. This is used in the depenetration phase of [method body_test_motion]. The higher the priority is, the lower the penetration into the body will be.
*/
//go:nosplit
func (self class) BodySetCollisionPriority(body gd.RID, priority gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the body's collision priority. This is used in the depenetration phase of [method body_test_motion]. The higher the priority is, the lower the penetration into the body will be.
*/
//go:nosplit
func (self class) BodyGetCollisionPriority(body gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of the given body parameter. See [enum BodyParameter] for the list of available parameters.
*/
//go:nosplit
func (self class) BodySetParam(body gd.RID, param classdb.PhysicsServer2DBodyParameter, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, param)
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the given body parameter. See [enum BodyParameter] for the list of available parameters.
*/
//go:nosplit
func (self class) BodyGetParam(ctx gd.Lifetime, body gd.RID, param classdb.PhysicsServer2DBodyParameter) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Restores the default inertia and center of mass of the body based on its shapes. This undoes any custom values previously set using [method body_set_param].
*/
//go:nosplit
func (self class) BodyResetMassProperties(body gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_reset_mass_properties, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the value of a body's state. See [enum BodyState] for the list of available states.
[b]Note:[/b] The state change doesn't take effect immediately. The state will change on the next physics frame.
*/
//go:nosplit
func (self class) BodySetState(body gd.RID, state classdb.PhysicsServer2DBodyState, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, state)
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the given state of the body. See [enum BodyState] for the list of available states.
*/
//go:nosplit
func (self class) BodyGetState(ctx gd.Lifetime, body gd.RID, state classdb.PhysicsServer2DBodyState) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, state)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Applies a directional impulse to the body, at the body's center of mass. The impulse does not affect rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method body_apply_impulse] at the body's center of mass.
*/
//go:nosplit
func (self class) BodyApplyCentralImpulse(body gd.RID, impulse gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a rotational impulse to the body. The impulse does not affect position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
*/
//go:nosplit
func (self class) BodyApplyTorqueImpulse(body gd.RID, impulse gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_apply_torque_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a positioned impulse to the body. The impulse can affect rotation if [param position] is different from the body's center of mass.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) BodyApplyImpulse(body gd.RID, impulse gd.Vector2, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a directional force to the body, at the body's center of mass. The force does not affect rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method body_apply_force] at the body's center of mass.
*/
//go:nosplit
func (self class) BodyApplyCentralForce(body gd.RID, force gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_apply_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a positioned force to the body. The force can affect rotation if [param position] is different from the body's center of mass. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) BodyApplyForce(body gd.RID, force gd.Vector2, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_apply_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a rotational force to the body. The force does not affect position. A force is time dependent and meant to be applied every physics update.
*/
//go:nosplit
func (self class) BodyApplyTorque(body gd.RID, torque gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_apply_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant directional force to the body. The force does not affect rotation. The force remains applied over time until cleared with [code]PhysicsServer2D.body_set_constant_force(body, Vector2(0, 0))[/code].
This is equivalent to using [method body_add_constant_force] at the body's center of mass.
*/
//go:nosplit
func (self class) BodyAddConstantCentralForce(body gd.RID, force gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_add_constant_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant positioned force to the body. The force can affect rotation if [param position] is different from the body's center of mass. The force remains applied over time until cleared with [code]PhysicsServer2D.body_set_constant_force(body, Vector2(0, 0))[/code].
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) BodyAddConstantForce(body gd.RID, force gd.Vector2, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_add_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant rotational force to the body. The force does not affect position. The force remains applied over time until cleared with [code]PhysicsServer2D.body_set_constant_torque(body, 0)[/code].
*/
//go:nosplit
func (self class) BodyAddConstantTorque(body gd.RID, torque gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_add_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the body's total constant positional force applied during each physics update.
See [method body_add_constant_force] and [method body_add_constant_central_force].
*/
//go:nosplit
func (self class) BodySetConstantForce(body gd.RID, force gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the body's total constant positional force applied during each physics update.
See [method body_add_constant_force] and [method body_add_constant_central_force].
*/
//go:nosplit
func (self class) BodyGetConstantForce(body gd.RID) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the body's total constant rotational force applied during each physics update.
See [method body_add_constant_torque].
*/
//go:nosplit
func (self class) BodySetConstantTorque(body gd.RID, torque gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the body's total constant rotational force applied during each physics update.
See [method body_add_constant_torque].
*/
//go:nosplit
func (self class) BodyGetConstantTorque(body gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Modifies the body's linear velocity so that its projection to the axis [code]axis_velocity.normalized()[/code] is exactly [code]axis_velocity.length()[/code]. This is useful for jumping behavior.
*/
//go:nosplit
func (self class) BodySetAxisVelocity(body gd.RID, axis_velocity gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, axis_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_axis_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds [param excepted_body] to the body's list of collision exceptions, so that collisions with it are ignored.
*/
//go:nosplit
func (self class) BodyAddCollisionException(body gd.RID, excepted_body gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, excepted_body)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_add_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes [param excepted_body] from the body's list of collision exceptions, so that collisions with it are no longer ignored.
*/
//go:nosplit
func (self class) BodyRemoveCollisionException(body gd.RID, excepted_body gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, excepted_body)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_remove_collision_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the maximum number of contacts that the body can report. If [param amount] is greater than zero, then the body will keep track of at most this many contacts with other bodies.
*/
//go:nosplit
func (self class) BodySetMaxContactsReported(body gd.RID, amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum number of contacts that the body can report. See [method body_set_max_contacts_reported].
*/
//go:nosplit
func (self class) BodyGetMaxContactsReported(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the body omits the standard force integration. If [param enable] is [code]true[/code], the body will not automatically use applied forces, torques, and damping to update the body's linear and angular velocity. In this case, [method body_set_force_integration_callback] can be used to manually update the linear and angular velocity instead.
This method is called when the property [member RigidBody2D.custom_integrator] is set.
*/
//go:nosplit
func (self class) BodySetOmitForceIntegration(body gd.RID, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_omit_force_integration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_is_omitting_force_integration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the body's state synchronization callback function to [param callable]. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the callback.
The function [param callable] will be called every physics frame, assuming that the body was active during the previous physics tick, and can be used to fetch the latest state from the physics server.
The function [param callable] must take the following parameters:
1. [code]state[/code]: a [PhysicsDirectBodyState2D], used to retrieve the body's state.
*/
//go:nosplit
func (self class) BodySetStateSyncCallback(body gd.RID, callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mmm.Get(callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_state_sync_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the body's custom force integration callback function to [param callable]. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the custom callback.
The function [param callable] will be called every physics tick, before the standard force integration (see [method body_set_omit_force_integration]). It can be used for example to update the body's linear and angular velocity based on contact with other bodies.
If [param userdata] is not [code]null[/code], the function [param callable] must take the following two parameters:
1. [code]state[/code]: a [PhysicsDirectBodyState2D] used to retrieve and modify the body's state,
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_set_force_integration_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a collision would result from moving the body along a motion vector from a given point in space. See [PhysicsTestMotionParameters2D] for the available motion parameters. Optionally a [PhysicsTestMotionResult2D] object can be passed, which will be used to store the information about the resulting collision.
*/
//go:nosplit
func (self class) BodyTestMotion(body gd.RID, parameters object.PhysicsTestMotionParameters2D, result object.PhysicsTestMotionResult2D) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	callframe.Arg(frame, mmm.Get(parameters[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(result[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_test_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [PhysicsDirectBodyState2D] of the body. Returns [code]null[/code] if the body is destroyed or not assigned to a space.
*/
//go:nosplit
func (self class) BodyGetDirectState(ctx gd.Lifetime, body gd.RID) object.PhysicsDirectBodyState2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_body_get_direct_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PhysicsDirectBodyState2D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a 2D joint in the physics server, and returns the [RID] that identifies it. To set the joint type, use [method joint_make_damped_spring], [method joint_make_groove] or [method joint_make_pin]. Use [method joint_set_param] to set generic joint parameters.
*/
//go:nosplit
func (self class) JointCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Destroys the joint with the given [RID], creates a new uninitialized joint, and makes the [RID] refer to this new joint.
*/
//go:nosplit
func (self class) JointClear(joint gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the value of the given joint parameter. See [enum JointParam] for the list of available parameters.
*/
//go:nosplit
func (self class) JointSetParam(joint gd.RID, param classdb.PhysicsServer2DJointParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the given joint parameter. See [enum JointParam] for the list of available parameters.
*/
//go:nosplit
func (self class) JointGetParam(joint gd.RID, param classdb.PhysicsServer2DJointParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the bodies attached to the [Joint2D] will collide with each other.
*/
//go:nosplit
func (self class) JointDisableCollisionsBetweenBodies(joint gd.RID, disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_disable_collisions_between_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the bodies attached to the [Joint2D] will collide with each other.
*/
//go:nosplit
func (self class) JointIsDisabledCollisionsBetweenBodies(joint gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_is_disabled_collisions_between_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Makes the joint a pin joint. If [param body_b] is an empty [RID], then [param body_a] is pinned to the point [param anchor] (given in global coordinates); otherwise, [param body_a] is pinned to [param body_b] at the point [param anchor] (given in global coordinates). To set the parameters which are specific to the pin joint, see [method pin_joint_set_param].
*/
//go:nosplit
func (self class) JointMakePin(joint gd.RID, anchor gd.Vector2, body_a gd.RID, body_b gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, anchor)
	callframe.Arg(frame, body_a)
	callframe.Arg(frame, body_b)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_make_pin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Makes the joint a groove joint.
*/
//go:nosplit
func (self class) JointMakeGroove(joint gd.RID, groove1_a gd.Vector2, groove2_a gd.Vector2, anchor_b gd.Vector2, body_a gd.RID, body_b gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, groove1_a)
	callframe.Arg(frame, groove2_a)
	callframe.Arg(frame, anchor_b)
	callframe.Arg(frame, body_a)
	callframe.Arg(frame, body_b)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_make_groove, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Makes the joint a damped spring joint, attached at the point [param anchor_a] (given in global coordinates) on the body [param body_a] and at the point [param anchor_b] (given in global coordinates) on the body [param body_b]. To set the parameters which are specific to the damped spring, see [method damped_spring_joint_set_param].
*/
//go:nosplit
func (self class) JointMakeDampedSpring(joint gd.RID, anchor_a gd.Vector2, anchor_b gd.Vector2, body_a gd.RID, body_b gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, anchor_a)
	callframe.Arg(frame, anchor_b)
	callframe.Arg(frame, body_a)
	callframe.Arg(frame, body_b)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_make_damped_spring, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a pin joint flag (see [enum PinJointFlag] constants).
*/
//go:nosplit
func (self class) PinJointSetFlag(joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_pin_joint_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a pin joint flag (see [enum PinJointFlag] constants).
*/
//go:nosplit
func (self class) PinJointGetFlag(joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_pin_joint_get_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a pin joint parameter. See [enum PinJointParam] for a list of available parameters.
*/
//go:nosplit
func (self class) PinJointSetParam(joint gd.RID, param classdb.PhysicsServer2DPinJointParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_pin_joint_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of a pin joint parameter. See [enum PinJointParam] for a list of available parameters.
*/
//go:nosplit
func (self class) PinJointGetParam(joint gd.RID, param classdb.PhysicsServer2DPinJointParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_pin_joint_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of the given damped spring joint parameter. See [enum DampedSpringParam] for the list of available parameters.
*/
//go:nosplit
func (self class) DampedSpringJointSetParam(joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_damped_spring_joint_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the given damped spring joint parameter. See [enum DampedSpringParam] for the list of available parameters.
*/
//go:nosplit
func (self class) DampedSpringJointGetParam(joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_damped_spring_joint_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the joint's type (see [enum JointType]).
*/
//go:nosplit
func (self class) JointGetType(joint gd.RID) classdb.PhysicsServer2DJointType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[classdb.PhysicsServer2DJointType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_joint_get_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Destroys any of the objects created by PhysicsServer2D. If the [RID] passed is not one of the objects that can be created by PhysicsServer2D, an error will be printed to the console.
*/
//go:nosplit
func (self class) FreeRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Activates or deactivates the 2D physics server. If [param active] is [code]false[/code], then the physics server will not do anything in its physics step.
*/
//go:nosplit
func (self class) SetActive(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns information about the current state of the 2D physics engine. See [enum ProcessInfo] for the list of available states.
*/
//go:nosplit
func (self class) GetProcessInfo(process_info classdb.PhysicsServer2DProcessInfo) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, process_info)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2D.Bind_get_process_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPhysicsServer2D() Expert { return self[0].AsPhysicsServer2D() }


//go:nosplit
func (self Simple) AsPhysicsServer2D() Simple { return self[0].AsPhysicsServer2D() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PhysicsServer2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type SpaceParameter = classdb.PhysicsServer2DSpaceParameter

const (
/*Constant to set/get the maximum distance a pair of bodies has to move before their collision status has to be recalculated. The default value of this parameter is [member ProjectSettings.physics/2d/solver/contact_recycle_radius].*/
	SpaceParamContactRecycleRadius SpaceParameter = 0
/*Constant to set/get the maximum distance a shape can be from another before they are considered separated and the contact is discarded. The default value of this parameter is [member ProjectSettings.physics/2d/solver/contact_max_separation].*/
	SpaceParamContactMaxSeparation SpaceParameter = 1
/*Constant to set/get the maximum distance a shape can penetrate another shape before it is considered a collision. The default value of this parameter is [member ProjectSettings.physics/2d/solver/contact_max_allowed_penetration].*/
	SpaceParamContactMaxAllowedPenetration SpaceParameter = 2
/*Constant to set/get the default solver bias for all physics contacts. A solver bias is a factor controlling how much two objects "rebound", after overlapping, to avoid leaving them in that state because of numerical imprecision. The default value of this parameter is [member ProjectSettings.physics/2d/solver/default_contact_bias].*/
	SpaceParamContactDefaultBias SpaceParameter = 3
/*Constant to set/get the threshold linear velocity of activity. A body marked as potentially inactive for both linear and angular velocity will be put to sleep after the time given. The default value of this parameter is [member ProjectSettings.physics/2d/sleep_threshold_linear].*/
	SpaceParamBodyLinearVelocitySleepThreshold SpaceParameter = 4
/*Constant to set/get the threshold angular velocity of activity. A body marked as potentially inactive for both linear and angular velocity will be put to sleep after the time given. The default value of this parameter is [member ProjectSettings.physics/2d/sleep_threshold_angular].*/
	SpaceParamBodyAngularVelocitySleepThreshold SpaceParameter = 5
/*Constant to set/get the maximum time of activity. A body marked as potentially inactive for both linear and angular velocity will be put to sleep after this time. The default value of this parameter is [member ProjectSettings.physics/2d/time_before_sleep].*/
	SpaceParamBodyTimeToSleep SpaceParameter = 6
/*Constant to set/get the default solver bias for all physics constraints. A solver bias is a factor controlling how much two objects "rebound", after violating a constraint, to avoid leaving them in that state because of numerical imprecision. The default value of this parameter is [member ProjectSettings.physics/2d/solver/default_constraint_bias].*/
	SpaceParamConstraintDefaultBias SpaceParameter = 7
/*Constant to set/get the number of solver iterations for all contacts and constraints. The greater the number of iterations, the more accurate the collisions will be. However, a greater number of iterations requires more CPU power, which can decrease performance. The default value of this parameter is [member ProjectSettings.physics/2d/solver/solver_iterations].*/
	SpaceParamSolverIterations SpaceParameter = 8
)
type ShapeType = classdb.PhysicsServer2DShapeType

const (
/*This is the constant for creating world boundary shapes. A world boundary shape is an [i]infinite[/i] line with an origin point, and a normal. Thus, it can be used for front/behind checks.*/
	ShapeWorldBoundary ShapeType = 0
/*This is the constant for creating separation ray shapes. A separation ray is defined by a length and separates itself from what is touching its far endpoint. Useful for character controllers.*/
	ShapeSeparationRay ShapeType = 1
/*This is the constant for creating segment shapes. A segment shape is a [i]finite[/i] line from a point A to a point B. It can be checked for intersections.*/
	ShapeSegment ShapeType = 2
/*This is the constant for creating circle shapes. A circle shape only has a radius. It can be used for intersections and inside/outside checks.*/
	ShapeCircle ShapeType = 3
/*This is the constant for creating rectangle shapes. A rectangle shape is defined by a width and a height. It can be used for intersections and inside/outside checks.*/
	ShapeRectangle ShapeType = 4
/*This is the constant for creating capsule shapes. A capsule shape is defined by a radius and a length. It can be used for intersections and inside/outside checks.*/
	ShapeCapsule ShapeType = 5
/*This is the constant for creating convex polygon shapes. A polygon is defined by a list of points. It can be used for intersections and inside/outside checks.*/
	ShapeConvexPolygon ShapeType = 6
/*This is the constant for creating concave polygon shapes. A polygon is defined by a list of points. It can be used for intersections checks, but not for inside/outside checks.*/
	ShapeConcavePolygon ShapeType = 7
/*This constant is used internally by the engine. Any attempt to create this kind of shape results in an error.*/
	ShapeCustom ShapeType = 8
)
type AreaParameter = classdb.PhysicsServer2DAreaParameter

const (
/*Constant to set/get gravity override mode in an area. See [enum AreaSpaceOverrideMode] for possible values. The default value of this parameter is [constant AREA_SPACE_OVERRIDE_DISABLED].*/
	AreaParamGravityOverrideMode AreaParameter = 0
/*Constant to set/get gravity strength in an area. The default value of this parameter is [code]9.80665[/code].*/
	AreaParamGravity AreaParameter = 1
/*Constant to set/get gravity vector/center in an area. The default value of this parameter is [code]Vector2(0, -1)[/code].*/
	AreaParamGravityVector AreaParameter = 2
/*Constant to set/get whether the gravity vector of an area is a direction, or a center point. The default value of this parameter is [code]false[/code].*/
	AreaParamGravityIsPoint AreaParameter = 3
/*Constant to set/get the distance at which the gravity strength is equal to the gravity controlled by [constant AREA_PARAM_GRAVITY]. For example, on a planet 100 pixels in radius with a surface gravity of 4.0 px/s², set the gravity to 4.0 and the unit distance to 100.0. The gravity will have falloff according to the inverse square law, so in the example, at 200 pixels from the center the gravity will be 1.0 px/s² (twice the distance, 1/4th the gravity), at 50 pixels it will be 16.0 px/s² (half the distance, 4x the gravity), and so on.
The above is true only when the unit distance is a positive number. When the unit distance is set to 0.0, the gravity will be constant regardless of distance. The default value of this parameter is [code]0.0[/code].*/
	AreaParamGravityPointUnitDistance AreaParameter = 4
/*Constant to set/get linear damping override mode in an area. See [enum AreaSpaceOverrideMode] for possible values. The default value of this parameter is [constant AREA_SPACE_OVERRIDE_DISABLED].*/
	AreaParamLinearDampOverrideMode AreaParameter = 5
/*Constant to set/get the linear damping factor of an area. The default value of this parameter is [code]0.1[/code].*/
	AreaParamLinearDamp AreaParameter = 6
/*Constant to set/get angular damping override mode in an area. See [enum AreaSpaceOverrideMode] for possible values. The default value of this parameter is [constant AREA_SPACE_OVERRIDE_DISABLED].*/
	AreaParamAngularDampOverrideMode AreaParameter = 7
/*Constant to set/get the angular damping factor of an area. The default value of this parameter is [code]1.0[/code].*/
	AreaParamAngularDamp AreaParameter = 8
/*Constant to set/get the priority (order of processing) of an area. The default value of this parameter is [code]0[/code].*/
	AreaParamPriority AreaParameter = 9
)
type AreaSpaceOverrideMode = classdb.PhysicsServer2DAreaSpaceOverrideMode

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
type BodyMode = classdb.PhysicsServer2DBodyMode

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
type BodyParameter = classdb.PhysicsServer2DBodyParameter

const (
/*Constant to set/get a body's bounce factor. The default value of this parameter is [code]0.0[/code].*/
	BodyParamBounce BodyParameter = 0
/*Constant to set/get a body's friction. The default value of this parameter is [code]1.0[/code].*/
	BodyParamFriction BodyParameter = 1
/*Constant to set/get a body's mass. The default value of this parameter is [code]1.0[/code]. If the body's mode is set to [constant BODY_MODE_RIGID], then setting this parameter will have the following additional effects:
- If the parameter [constant BODY_PARAM_CENTER_OF_MASS] has never been set explicitly, then the value of that parameter will be recalculated based on the body's shapes.
- If the parameter [constant BODY_PARAM_INERTIA] is set to a value [code]<= 0.0[/code], then the value of that parameter will be recalculated based on the body's shapes, mass, and center of mass.*/
	BodyParamMass BodyParameter = 2
/*Constant to set/get a body's inertia. The default value of this parameter is [code]0.0[/code]. If the body's inertia is set to a value [code]<= 0.0[/code], then the inertia will be recalculated based on the body's shapes, mass, and center of mass.*/
	BodyParamInertia BodyParameter = 3
/*Constant to set/get a body's center of mass position in the body's local coordinate system. The default value of this parameter is [code]Vector2(0,0)[/code]. If this parameter is never set explicitly, then it is recalculated based on the body's shapes when setting the parameter [constant BODY_PARAM_MASS] or when calling [method body_set_space].*/
	BodyParamCenterOfMass BodyParameter = 4
/*Constant to set/get a body's gravity multiplier. The default value of this parameter is [code]1.0[/code].*/
	BodyParamGravityScale BodyParameter = 5
/*Constant to set/get a body's linear damping mode. See [enum BodyDampMode] for possible values. The default value of this parameter is [constant BODY_DAMP_MODE_COMBINE].*/
	BodyParamLinearDampMode BodyParameter = 6
/*Constant to set/get a body's angular damping mode. See [enum BodyDampMode] for possible values. The default value of this parameter is [constant BODY_DAMP_MODE_COMBINE].*/
	BodyParamAngularDampMode BodyParameter = 7
/*Constant to set/get a body's linear damping factor. The default value of this parameter is [code]0.0[/code].*/
	BodyParamLinearDamp BodyParameter = 8
/*Constant to set/get a body's angular damping factor. The default value of this parameter is [code]0.0[/code].*/
	BodyParamAngularDamp BodyParameter = 9
/*Represents the size of the [enum BodyParameter] enum.*/
	BodyParamMax BodyParameter = 10
)
type BodyDampMode = classdb.PhysicsServer2DBodyDampMode

const (
/*The body's damping value is added to any value set in areas or the default value.*/
	BodyDampModeCombine BodyDampMode = 0
/*The body's damping value replaces any value set in areas or the default value.*/
	BodyDampModeReplace BodyDampMode = 1
)
type BodyState = classdb.PhysicsServer2DBodyState

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
type JointType = classdb.PhysicsServer2DJointType

const (
/*Constant to create pin joints.*/
	JointTypePin JointType = 0
/*Constant to create groove joints.*/
	JointTypeGroove JointType = 1
/*Constant to create damped spring joints.*/
	JointTypeDampedSpring JointType = 2
/*Represents the size of the [enum JointType] enum.*/
	JointTypeMax JointType = 3
)
type JointParam = classdb.PhysicsServer2DJointParam

const (
/*Constant to set/get how fast the joint pulls the bodies back to satisfy the joint constraint. The lower the value, the more the two bodies can pull on the joint. The default value of this parameter is [code]0.0[/code].
[b]Note:[/b] In Godot Physics, this parameter is only used for pin joints and groove joints.*/
	JointParamBias JointParam = 0
/*Constant to set/get the maximum speed with which the joint can apply corrections. The default value of this parameter is [code]3.40282e+38[/code].
[b]Note:[/b] In Godot Physics, this parameter is only used for groove joints.*/
	JointParamMaxBias JointParam = 1
/*Constant to set/get the maximum force that the joint can use to act on the two bodies. The default value of this parameter is [code]3.40282e+38[/code].
[b]Note:[/b] In Godot Physics, this parameter is only used for groove joints.*/
	JointParamMaxForce JointParam = 2
)
type PinJointParam = classdb.PhysicsServer2DPinJointParam

const (
/*Constant to set/get a how much the bond of the pin joint can flex. The default value of this parameter is [code]0.0[/code].*/
	PinJointSoftness PinJointParam = 0
/*The maximum rotation around the pin.*/
	PinJointLimitUpper PinJointParam = 1
/*The minimum rotation around the pin.*/
	PinJointLimitLower PinJointParam = 2
/*Target speed for the motor. In radians per second.*/
	PinJointMotorTargetVelocity PinJointParam = 3
)
type PinJointFlag = classdb.PhysicsServer2DPinJointFlag

const (
/*If [code]true[/code], the pin has a maximum and a minimum rotation.*/
	PinJointFlagAngularLimitEnabled PinJointFlag = 0
/*If [code]true[/code], a motor turns the pin.*/
	PinJointFlagMotorEnabled PinJointFlag = 1
)
type DampedSpringParam = classdb.PhysicsServer2DDampedSpringParam

const (
/*Sets the resting length of the spring joint. The joint will always try to go to back this length when pulled apart. The default value of this parameter is the distance between the joint's anchor points.*/
	DampedSpringRestLength DampedSpringParam = 0
/*Sets the stiffness of the spring joint. The joint applies a force equal to the stiffness times the distance from its resting length. The default value of this parameter is [code]20.0[/code].*/
	DampedSpringStiffness DampedSpringParam = 1
/*Sets the damping ratio of the spring joint. A value of 0 indicates an undamped spring, while 1 causes the system to reach equilibrium as fast as possible (critical damping). The default value of this parameter is [code]1.5[/code].*/
	DampedSpringDamping DampedSpringParam = 2
)
type CCDMode = classdb.PhysicsServer2DCCDMode

const (
/*Disables continuous collision detection. This is the fastest way to detect body collisions, but it can miss small and/or fast-moving objects.*/
	CcdModeDisabled CCDMode = 0
/*Enables continuous collision detection by raycasting. It is faster than shapecasting, but less precise.*/
	CcdModeCastRay CCDMode = 1
/*Enables continuous collision detection by shapecasting. It is the slowest CCD method, and the most precise.*/
	CcdModeCastShape CCDMode = 2
)
type AreaBodyStatus = classdb.PhysicsServer2DAreaBodyStatus

const (
/*The value of the first parameter and area callback function receives, when an object enters one of its shapes.*/
	AreaBodyAdded AreaBodyStatus = 0
/*The value of the first parameter and area callback function receives, when an object exits one of its shapes.*/
	AreaBodyRemoved AreaBodyStatus = 1
)
type ProcessInfo = classdb.PhysicsServer2DProcessInfo

const (
/*Constant to get the number of objects that are not sleeping.*/
	InfoActiveObjects ProcessInfo = 0
/*Constant to get the number of possible collisions.*/
	InfoCollisionPairs ProcessInfo = 1
/*Constant to get the number of space regions where a collision could occur.*/
	InfoIslandCount ProcessInfo = 2
)
