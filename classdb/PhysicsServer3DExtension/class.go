// Package PhysicsServer3DExtension provides methods for working with PhysicsServer3DExtension object instances.
package PhysicsServer3DExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/AABB"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
This class extends [PhysicsServer3D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsServer3D].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=PhysicsServer3DExtension)
*/
type Instance [1]gdclass.PhysicsServer3DExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsServer3DExtension() Instance
}
type Interface interface {
	WorldBoundaryShapeCreate() Resource.ID
	SeparationRayShapeCreate() Resource.ID
	SphereShapeCreate() Resource.ID
	BoxShapeCreate() Resource.ID
	CapsuleShapeCreate() Resource.ID
	CylinderShapeCreate() Resource.ID
	ConvexPolygonShapeCreate() Resource.ID
	ConcavePolygonShapeCreate() Resource.ID
	HeightmapShapeCreate() Resource.ID
	CustomShapeCreate() Resource.ID
	ShapeSetData(shape Resource.ID, data any)
	ShapeSetCustomSolverBias(shape Resource.ID, bias Float.X)
	ShapeSetMargin(shape Resource.ID, margin Float.X)
	ShapeGetMargin(shape Resource.ID) Float.X
	ShapeGetType(shape Resource.ID) gdclass.PhysicsServer3DShapeType
	ShapeGetData(shape Resource.ID) any
	ShapeGetCustomSolverBias(shape Resource.ID) Float.X
	SpaceCreate() Resource.ID
	SpaceSetActive(space Resource.ID, active bool)
	SpaceIsActive(space Resource.ID) bool
	SpaceSetParam(space Resource.ID, param gdclass.PhysicsServer3DSpaceParameter, value Float.X)
	SpaceGetParam(space Resource.ID, param gdclass.PhysicsServer3DSpaceParameter) Float.X
	SpaceGetDirectState(space Resource.ID) [1]gdclass.PhysicsDirectSpaceState3D
	SpaceSetDebugContacts(space Resource.ID, max_contacts int)
	SpaceGetContacts(space Resource.ID) []Vector3.XYZ
	SpaceGetContactCount(space Resource.ID) int
	AreaCreate() Resource.ID
	AreaSetSpace(area Resource.ID, space Resource.ID)
	AreaGetSpace(area Resource.ID) Resource.ID
	AreaAddShape(area Resource.ID, shape Resource.ID, transform Transform3D.BasisOrigin, disabled bool)
	AreaSetShape(area Resource.ID, shape_idx int, shape Resource.ID)
	AreaSetShapeTransform(area Resource.ID, shape_idx int, transform Transform3D.BasisOrigin)
	AreaSetShapeDisabled(area Resource.ID, shape_idx int, disabled bool)
	AreaGetShapeCount(area Resource.ID) int
	AreaGetShape(area Resource.ID, shape_idx int) Resource.ID
	AreaGetShapeTransform(area Resource.ID, shape_idx int) Transform3D.BasisOrigin
	AreaRemoveShape(area Resource.ID, shape_idx int)
	AreaClearShapes(area Resource.ID)
	AreaAttachObjectInstanceId(area Resource.ID, id int)
	AreaGetObjectInstanceId(area Resource.ID) int
	AreaSetParam(area Resource.ID, param gdclass.PhysicsServer3DAreaParameter, value any)
	AreaSetTransform(area Resource.ID, transform Transform3D.BasisOrigin)
	AreaGetParam(area Resource.ID, param gdclass.PhysicsServer3DAreaParameter) any
	AreaGetTransform(area Resource.ID) Transform3D.BasisOrigin
	AreaSetCollisionLayer(area Resource.ID, layer int)
	AreaGetCollisionLayer(area Resource.ID) int
	AreaSetCollisionMask(area Resource.ID, mask int)
	AreaGetCollisionMask(area Resource.ID) int
	AreaSetMonitorable(area Resource.ID, monitorable bool)
	AreaSetRayPickable(area Resource.ID, enable bool)
	AreaSetMonitorCallback(area Resource.ID, callback Callable.Any)
	AreaSetAreaMonitorCallback(area Resource.ID, callback Callable.Any)
	BodyCreate() Resource.ID
	BodySetSpace(body Resource.ID, space Resource.ID)
	BodyGetSpace(body Resource.ID) Resource.ID
	BodySetMode(body Resource.ID, mode gdclass.PhysicsServer3DBodyMode)
	BodyGetMode(body Resource.ID) gdclass.PhysicsServer3DBodyMode
	BodyAddShape(body Resource.ID, shape Resource.ID, transform Transform3D.BasisOrigin, disabled bool)
	BodySetShape(body Resource.ID, shape_idx int, shape Resource.ID)
	BodySetShapeTransform(body Resource.ID, shape_idx int, transform Transform3D.BasisOrigin)
	BodySetShapeDisabled(body Resource.ID, shape_idx int, disabled bool)
	BodyGetShapeCount(body Resource.ID) int
	BodyGetShape(body Resource.ID, shape_idx int) Resource.ID
	BodyGetShapeTransform(body Resource.ID, shape_idx int) Transform3D.BasisOrigin
	BodyRemoveShape(body Resource.ID, shape_idx int)
	BodyClearShapes(body Resource.ID)
	BodyAttachObjectInstanceId(body Resource.ID, id int)
	BodyGetObjectInstanceId(body Resource.ID) int
	BodySetEnableContinuousCollisionDetection(body Resource.ID, enable bool)
	BodyIsContinuousCollisionDetectionEnabled(body Resource.ID) bool
	BodySetCollisionLayer(body Resource.ID, layer int)
	BodyGetCollisionLayer(body Resource.ID) int
	BodySetCollisionMask(body Resource.ID, mask int)
	BodyGetCollisionMask(body Resource.ID) int
	BodySetCollisionPriority(body Resource.ID, priority Float.X)
	BodyGetCollisionPriority(body Resource.ID) Float.X
	BodySetUserFlags(body Resource.ID, flags int)
	BodyGetUserFlags(body Resource.ID) int
	BodySetParam(body Resource.ID, param gdclass.PhysicsServer3DBodyParameter, value any)
	BodyGetParam(body Resource.ID, param gdclass.PhysicsServer3DBodyParameter) any
	BodyResetMassProperties(body Resource.ID)
	BodySetState(body Resource.ID, state gdclass.PhysicsServer3DBodyState, value any)
	BodyGetState(body Resource.ID, state gdclass.PhysicsServer3DBodyState) any
	BodyApplyCentralImpulse(body Resource.ID, impulse Vector3.XYZ)
	BodyApplyImpulse(body Resource.ID, impulse Vector3.XYZ, position Vector3.XYZ)
	BodyApplyTorqueImpulse(body Resource.ID, impulse Vector3.XYZ)
	BodyApplyCentralForce(body Resource.ID, force Vector3.XYZ)
	BodyApplyForce(body Resource.ID, force Vector3.XYZ, position Vector3.XYZ)
	BodyApplyTorque(body Resource.ID, torque Vector3.XYZ)
	BodyAddConstantCentralForce(body Resource.ID, force Vector3.XYZ)
	BodyAddConstantForce(body Resource.ID, force Vector3.XYZ, position Vector3.XYZ)
	BodyAddConstantTorque(body Resource.ID, torque Vector3.XYZ)
	BodySetConstantForce(body Resource.ID, force Vector3.XYZ)
	BodyGetConstantForce(body Resource.ID) Vector3.XYZ
	BodySetConstantTorque(body Resource.ID, torque Vector3.XYZ)
	BodyGetConstantTorque(body Resource.ID) Vector3.XYZ
	BodySetAxisVelocity(body Resource.ID, axis_velocity Vector3.XYZ)
	BodySetAxisLock(body Resource.ID, axis gdclass.PhysicsServer3DBodyAxis, lock bool)
	BodyIsAxisLocked(body Resource.ID, axis gdclass.PhysicsServer3DBodyAxis) bool
	BodyAddCollisionException(body Resource.ID, excepted_body Resource.ID)
	BodyRemoveCollisionException(body Resource.ID, excepted_body Resource.ID)
	BodyGetCollisionExceptions(body Resource.ID) gd.Array
	BodySetMaxContactsReported(body Resource.ID, amount int)
	BodyGetMaxContactsReported(body Resource.ID) int
	BodySetContactsReportedDepthThreshold(body Resource.ID, threshold Float.X)
	BodyGetContactsReportedDepthThreshold(body Resource.ID) Float.X
	BodySetOmitForceIntegration(body Resource.ID, enable bool)
	BodyIsOmittingForceIntegration(body Resource.ID) bool
	BodySetStateSyncCallback(body Resource.ID, callable Callable.Any)
	BodySetForceIntegrationCallback(body Resource.ID, callable Callable.Any, userdata any)
	BodySetRayPickable(body Resource.ID, enable bool)
	BodyTestMotion(body Resource.ID, from Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, max_collisions int, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) bool
	BodyGetDirectState(body Resource.ID) [1]gdclass.PhysicsDirectBodyState3D
	SoftBodyCreate() Resource.ID
	SoftBodyUpdateRenderingServer(body Resource.ID, rendering_server_handler [1]gdclass.PhysicsServer3DRenderingServerHandler)
	SoftBodySetSpace(body Resource.ID, space Resource.ID)
	SoftBodyGetSpace(body Resource.ID) Resource.ID
	SoftBodySetRayPickable(body Resource.ID, enable bool)
	SoftBodySetCollisionLayer(body Resource.ID, layer int)
	SoftBodyGetCollisionLayer(body Resource.ID) int
	SoftBodySetCollisionMask(body Resource.ID, mask int)
	SoftBodyGetCollisionMask(body Resource.ID) int
	SoftBodyAddCollisionException(body Resource.ID, body_b Resource.ID)
	SoftBodyRemoveCollisionException(body Resource.ID, body_b Resource.ID)
	SoftBodyGetCollisionExceptions(body Resource.ID) gd.Array
	SoftBodySetState(body Resource.ID, state gdclass.PhysicsServer3DBodyState, variant any)
	SoftBodyGetState(body Resource.ID, state gdclass.PhysicsServer3DBodyState) any
	SoftBodySetTransform(body Resource.ID, transform Transform3D.BasisOrigin)
	SoftBodySetSimulationPrecision(body Resource.ID, simulation_precision int)
	SoftBodyGetSimulationPrecision(body Resource.ID) int
	SoftBodySetTotalMass(body Resource.ID, total_mass Float.X)
	SoftBodyGetTotalMass(body Resource.ID) Float.X
	SoftBodySetLinearStiffness(body Resource.ID, linear_stiffness Float.X)
	SoftBodyGetLinearStiffness(body Resource.ID) Float.X
	SoftBodySetPressureCoefficient(body Resource.ID, pressure_coefficient Float.X)
	SoftBodyGetPressureCoefficient(body Resource.ID) Float.X
	SoftBodySetDampingCoefficient(body Resource.ID, damping_coefficient Float.X)
	SoftBodyGetDampingCoefficient(body Resource.ID) Float.X
	SoftBodySetDragCoefficient(body Resource.ID, drag_coefficient Float.X)
	SoftBodyGetDragCoefficient(body Resource.ID) Float.X
	SoftBodySetMesh(body Resource.ID, mesh Resource.ID)
	SoftBodyGetBounds(body Resource.ID) AABB.PositionSize
	SoftBodyMovePoint(body Resource.ID, point_index int, global_position Vector3.XYZ)
	SoftBodyGetPointGlobalPosition(body Resource.ID, point_index int) Vector3.XYZ
	SoftBodyRemoveAllPinnedPoints(body Resource.ID)
	SoftBodyPinPoint(body Resource.ID, point_index int, pin bool)
	SoftBodyIsPointPinned(body Resource.ID, point_index int) bool
	JointCreate() Resource.ID
	JointClear(joint Resource.ID)
	JointMakePin(joint Resource.ID, body_A Resource.ID, local_A Vector3.XYZ, body_B Resource.ID, local_B Vector3.XYZ)
	PinJointSetParam(joint Resource.ID, param gdclass.PhysicsServer3DPinJointParam, value Float.X)
	PinJointGetParam(joint Resource.ID, param gdclass.PhysicsServer3DPinJointParam) Float.X
	PinJointSetLocalA(joint Resource.ID, local_A Vector3.XYZ)
	PinJointGetLocalA(joint Resource.ID) Vector3.XYZ
	PinJointSetLocalB(joint Resource.ID, local_B Vector3.XYZ)
	PinJointGetLocalB(joint Resource.ID) Vector3.XYZ
	JointMakeHinge(joint Resource.ID, body_A Resource.ID, hinge_A Transform3D.BasisOrigin, body_B Resource.ID, hinge_B Transform3D.BasisOrigin)
	JointMakeHingeSimple(joint Resource.ID, body_A Resource.ID, pivot_A Vector3.XYZ, axis_A Vector3.XYZ, body_B Resource.ID, pivot_B Vector3.XYZ, axis_B Vector3.XYZ)
	HingeJointSetParam(joint Resource.ID, param gdclass.PhysicsServer3DHingeJointParam, value Float.X)
	HingeJointGetParam(joint Resource.ID, param gdclass.PhysicsServer3DHingeJointParam) Float.X
	HingeJointSetFlag(joint Resource.ID, flag gdclass.PhysicsServer3DHingeJointFlag, enabled bool)
	HingeJointGetFlag(joint Resource.ID, flag gdclass.PhysicsServer3DHingeJointFlag) bool
	JointMakeSlider(joint Resource.ID, body_A Resource.ID, local_ref_A Transform3D.BasisOrigin, body_B Resource.ID, local_ref_B Transform3D.BasisOrigin)
	SliderJointSetParam(joint Resource.ID, param gdclass.PhysicsServer3DSliderJointParam, value Float.X)
	SliderJointGetParam(joint Resource.ID, param gdclass.PhysicsServer3DSliderJointParam) Float.X
	JointMakeConeTwist(joint Resource.ID, body_A Resource.ID, local_ref_A Transform3D.BasisOrigin, body_B Resource.ID, local_ref_B Transform3D.BasisOrigin)
	ConeTwistJointSetParam(joint Resource.ID, param gdclass.PhysicsServer3DConeTwistJointParam, value Float.X)
	ConeTwistJointGetParam(joint Resource.ID, param gdclass.PhysicsServer3DConeTwistJointParam) Float.X
	JointMakeGeneric6dof(joint Resource.ID, body_A Resource.ID, local_ref_A Transform3D.BasisOrigin, body_B Resource.ID, local_ref_B Transform3D.BasisOrigin)
	Generic6dofJointSetParam(joint Resource.ID, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam, value Float.X)
	Generic6dofJointGetParam(joint Resource.ID, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam) Float.X
	Generic6dofJointSetFlag(joint Resource.ID, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag, enable bool)
	Generic6dofJointGetFlag(joint Resource.ID, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag) bool
	JointGetType(joint Resource.ID) gdclass.PhysicsServer3DJointType
	JointSetSolverPriority(joint Resource.ID, priority int)
	JointGetSolverPriority(joint Resource.ID) int
	JointDisableCollisionsBetweenBodies(joint Resource.ID, disable bool)
	JointIsDisabledCollisionsBetweenBodies(joint Resource.ID) bool
	FreeRid(rid Resource.ID)
	SetActive(active bool)
	Init()
	Step(step Float.X)
	Sync()
	FlushQueries()
	EndSync()
	Finish()
	IsFlushingQueries() bool
	GetProcessInfo(process_info gdclass.PhysicsServer3DProcessInfo) int
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) WorldBoundaryShapeCreate() (_ Resource.ID)                { return }
func (self Implementation) SeparationRayShapeCreate() (_ Resource.ID)                { return }
func (self Implementation) SphereShapeCreate() (_ Resource.ID)                       { return }
func (self Implementation) BoxShapeCreate() (_ Resource.ID)                          { return }
func (self Implementation) CapsuleShapeCreate() (_ Resource.ID)                      { return }
func (self Implementation) CylinderShapeCreate() (_ Resource.ID)                     { return }
func (self Implementation) ConvexPolygonShapeCreate() (_ Resource.ID)                { return }
func (self Implementation) ConcavePolygonShapeCreate() (_ Resource.ID)               { return }
func (self Implementation) HeightmapShapeCreate() (_ Resource.ID)                    { return }
func (self Implementation) CustomShapeCreate() (_ Resource.ID)                       { return }
func (self Implementation) ShapeSetData(shape Resource.ID, data any)                 { return }
func (self Implementation) ShapeSetCustomSolverBias(shape Resource.ID, bias Float.X) { return }
func (self Implementation) ShapeSetMargin(shape Resource.ID, margin Float.X)         { return }
func (self Implementation) ShapeGetMargin(shape Resource.ID) (_ Float.X)             { return }
func (self Implementation) ShapeGetType(shape Resource.ID) (_ gdclass.PhysicsServer3DShapeType) {
	return
}
func (self Implementation) ShapeGetData(shape Resource.ID) (_ any)                 { return }
func (self Implementation) ShapeGetCustomSolverBias(shape Resource.ID) (_ Float.X) { return }
func (self Implementation) SpaceCreate() (_ Resource.ID)                           { return }
func (self Implementation) SpaceSetActive(space Resource.ID, active bool)          { return }
func (self Implementation) SpaceIsActive(space Resource.ID) (_ bool)               { return }
func (self Implementation) SpaceSetParam(space Resource.ID, param gdclass.PhysicsServer3DSpaceParameter, value Float.X) {
	return
}
func (self Implementation) SpaceGetParam(space Resource.ID, param gdclass.PhysicsServer3DSpaceParameter) (_ Float.X) {
	return
}
func (self Implementation) SpaceGetDirectState(space Resource.ID) (_ [1]gdclass.PhysicsDirectSpaceState3D) {
	return
}
func (self Implementation) SpaceSetDebugContacts(space Resource.ID, max_contacts int) { return }
func (self Implementation) SpaceGetContacts(space Resource.ID) (_ []Vector3.XYZ)      { return }
func (self Implementation) SpaceGetContactCount(space Resource.ID) (_ int)            { return }
func (self Implementation) AreaCreate() (_ Resource.ID)                               { return }
func (self Implementation) AreaSetSpace(area Resource.ID, space Resource.ID)          { return }
func (self Implementation) AreaGetSpace(area Resource.ID) (_ Resource.ID)             { return }
func (self Implementation) AreaAddShape(area Resource.ID, shape Resource.ID, transform Transform3D.BasisOrigin, disabled bool) {
	return
}
func (self Implementation) AreaSetShape(area Resource.ID, shape_idx int, shape Resource.ID) { return }
func (self Implementation) AreaSetShapeTransform(area Resource.ID, shape_idx int, transform Transform3D.BasisOrigin) {
	return
}
func (self Implementation) AreaSetShapeDisabled(area Resource.ID, shape_idx int, disabled bool) {
	return
}
func (self Implementation) AreaGetShapeCount(area Resource.ID) (_ int)                   { return }
func (self Implementation) AreaGetShape(area Resource.ID, shape_idx int) (_ Resource.ID) { return }
func (self Implementation) AreaGetShapeTransform(area Resource.ID, shape_idx int) (_ Transform3D.BasisOrigin) {
	return
}
func (self Implementation) AreaRemoveShape(area Resource.ID, shape_idx int)     { return }
func (self Implementation) AreaClearShapes(area Resource.ID)                    { return }
func (self Implementation) AreaAttachObjectInstanceId(area Resource.ID, id int) { return }
func (self Implementation) AreaGetObjectInstanceId(area Resource.ID) (_ int)    { return }
func (self Implementation) AreaSetParam(area Resource.ID, param gdclass.PhysicsServer3DAreaParameter, value any) {
	return
}
func (self Implementation) AreaSetTransform(area Resource.ID, transform Transform3D.BasisOrigin) {
	return
}
func (self Implementation) AreaGetParam(area Resource.ID, param gdclass.PhysicsServer3DAreaParameter) (_ any) {
	return
}
func (self Implementation) AreaGetTransform(area Resource.ID) (_ Transform3D.BasisOrigin)  { return }
func (self Implementation) AreaSetCollisionLayer(area Resource.ID, layer int)              { return }
func (self Implementation) AreaGetCollisionLayer(area Resource.ID) (_ int)                 { return }
func (self Implementation) AreaSetCollisionMask(area Resource.ID, mask int)                { return }
func (self Implementation) AreaGetCollisionMask(area Resource.ID) (_ int)                  { return }
func (self Implementation) AreaSetMonitorable(area Resource.ID, monitorable bool)          { return }
func (self Implementation) AreaSetRayPickable(area Resource.ID, enable bool)               { return }
func (self Implementation) AreaSetMonitorCallback(area Resource.ID, callback Callable.Any) { return }
func (self Implementation) AreaSetAreaMonitorCallback(area Resource.ID, callback Callable.Any) {
	return
}
func (self Implementation) BodyCreate() (_ Resource.ID)                      { return }
func (self Implementation) BodySetSpace(body Resource.ID, space Resource.ID) { return }
func (self Implementation) BodyGetSpace(body Resource.ID) (_ Resource.ID)    { return }
func (self Implementation) BodySetMode(body Resource.ID, mode gdclass.PhysicsServer3DBodyMode) {
	return
}
func (self Implementation) BodyGetMode(body Resource.ID) (_ gdclass.PhysicsServer3DBodyMode) { return }
func (self Implementation) BodyAddShape(body Resource.ID, shape Resource.ID, transform Transform3D.BasisOrigin, disabled bool) {
	return
}
func (self Implementation) BodySetShape(body Resource.ID, shape_idx int, shape Resource.ID) { return }
func (self Implementation) BodySetShapeTransform(body Resource.ID, shape_idx int, transform Transform3D.BasisOrigin) {
	return
}
func (self Implementation) BodySetShapeDisabled(body Resource.ID, shape_idx int, disabled bool) {
	return
}
func (self Implementation) BodyGetShapeCount(body Resource.ID) (_ int)                   { return }
func (self Implementation) BodyGetShape(body Resource.ID, shape_idx int) (_ Resource.ID) { return }
func (self Implementation) BodyGetShapeTransform(body Resource.ID, shape_idx int) (_ Transform3D.BasisOrigin) {
	return
}
func (self Implementation) BodyRemoveShape(body Resource.ID, shape_idx int)     { return }
func (self Implementation) BodyClearShapes(body Resource.ID)                    { return }
func (self Implementation) BodyAttachObjectInstanceId(body Resource.ID, id int) { return }
func (self Implementation) BodyGetObjectInstanceId(body Resource.ID) (_ int)    { return }
func (self Implementation) BodySetEnableContinuousCollisionDetection(body Resource.ID, enable bool) {
	return
}
func (self Implementation) BodyIsContinuousCollisionDetectionEnabled(body Resource.ID) (_ bool) {
	return
}
func (self Implementation) BodySetCollisionLayer(body Resource.ID, layer int)           { return }
func (self Implementation) BodyGetCollisionLayer(body Resource.ID) (_ int)              { return }
func (self Implementation) BodySetCollisionMask(body Resource.ID, mask int)             { return }
func (self Implementation) BodyGetCollisionMask(body Resource.ID) (_ int)               { return }
func (self Implementation) BodySetCollisionPriority(body Resource.ID, priority Float.X) { return }
func (self Implementation) BodyGetCollisionPriority(body Resource.ID) (_ Float.X)       { return }
func (self Implementation) BodySetUserFlags(body Resource.ID, flags int)                { return }
func (self Implementation) BodyGetUserFlags(body Resource.ID) (_ int)                   { return }
func (self Implementation) BodySetParam(body Resource.ID, param gdclass.PhysicsServer3DBodyParameter, value any) {
	return
}
func (self Implementation) BodyGetParam(body Resource.ID, param gdclass.PhysicsServer3DBodyParameter) (_ any) {
	return
}
func (self Implementation) BodyResetMassProperties(body Resource.ID) { return }
func (self Implementation) BodySetState(body Resource.ID, state gdclass.PhysicsServer3DBodyState, value any) {
	return
}
func (self Implementation) BodyGetState(body Resource.ID, state gdclass.PhysicsServer3DBodyState) (_ any) {
	return
}
func (self Implementation) BodyApplyCentralImpulse(body Resource.ID, impulse Vector3.XYZ) { return }
func (self Implementation) BodyApplyImpulse(body Resource.ID, impulse Vector3.XYZ, position Vector3.XYZ) {
	return
}
func (self Implementation) BodyApplyTorqueImpulse(body Resource.ID, impulse Vector3.XYZ) { return }
func (self Implementation) BodyApplyCentralForce(body Resource.ID, force Vector3.XYZ)    { return }
func (self Implementation) BodyApplyForce(body Resource.ID, force Vector3.XYZ, position Vector3.XYZ) {
	return
}
func (self Implementation) BodyApplyTorque(body Resource.ID, torque Vector3.XYZ)            { return }
func (self Implementation) BodyAddConstantCentralForce(body Resource.ID, force Vector3.XYZ) { return }
func (self Implementation) BodyAddConstantForce(body Resource.ID, force Vector3.XYZ, position Vector3.XYZ) {
	return
}
func (self Implementation) BodyAddConstantTorque(body Resource.ID, torque Vector3.XYZ)      { return }
func (self Implementation) BodySetConstantForce(body Resource.ID, force Vector3.XYZ)        { return }
func (self Implementation) BodyGetConstantForce(body Resource.ID) (_ Vector3.XYZ)           { return }
func (self Implementation) BodySetConstantTorque(body Resource.ID, torque Vector3.XYZ)      { return }
func (self Implementation) BodyGetConstantTorque(body Resource.ID) (_ Vector3.XYZ)          { return }
func (self Implementation) BodySetAxisVelocity(body Resource.ID, axis_velocity Vector3.XYZ) { return }
func (self Implementation) BodySetAxisLock(body Resource.ID, axis gdclass.PhysicsServer3DBodyAxis, lock bool) {
	return
}
func (self Implementation) BodyIsAxisLocked(body Resource.ID, axis gdclass.PhysicsServer3DBodyAxis) (_ bool) {
	return
}
func (self Implementation) BodyAddCollisionException(body Resource.ID, excepted_body Resource.ID) {
	return
}
func (self Implementation) BodyRemoveCollisionException(body Resource.ID, excepted_body Resource.ID) {
	return
}
func (self Implementation) BodyGetCollisionExceptions(body Resource.ID) (_ gd.Array) { return }
func (self Implementation) BodySetMaxContactsReported(body Resource.ID, amount int)  { return }
func (self Implementation) BodyGetMaxContactsReported(body Resource.ID) (_ int)      { return }
func (self Implementation) BodySetContactsReportedDepthThreshold(body Resource.ID, threshold Float.X) {
	return
}
func (self Implementation) BodyGetContactsReportedDepthThreshold(body Resource.ID) (_ Float.X) {
	return
}
func (self Implementation) BodySetOmitForceIntegration(body Resource.ID, enable bool)        { return }
func (self Implementation) BodyIsOmittingForceIntegration(body Resource.ID) (_ bool)         { return }
func (self Implementation) BodySetStateSyncCallback(body Resource.ID, callable Callable.Any) { return }
func (self Implementation) BodySetForceIntegrationCallback(body Resource.ID, callable Callable.Any, userdata any) {
	return
}
func (self Implementation) BodySetRayPickable(body Resource.ID, enable bool) { return }
func (self Implementation) BodyTestMotion(body Resource.ID, from Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, max_collisions int, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) (_ bool) {
	return
}
func (self Implementation) BodyGetDirectState(body Resource.ID) (_ [1]gdclass.PhysicsDirectBodyState3D) {
	return
}
func (self Implementation) SoftBodyCreate() (_ Resource.ID) { return }
func (self Implementation) SoftBodyUpdateRenderingServer(body Resource.ID, rendering_server_handler [1]gdclass.PhysicsServer3DRenderingServerHandler) {
	return
}
func (self Implementation) SoftBodySetSpace(body Resource.ID, space Resource.ID)  { return }
func (self Implementation) SoftBodyGetSpace(body Resource.ID) (_ Resource.ID)     { return }
func (self Implementation) SoftBodySetRayPickable(body Resource.ID, enable bool)  { return }
func (self Implementation) SoftBodySetCollisionLayer(body Resource.ID, layer int) { return }
func (self Implementation) SoftBodyGetCollisionLayer(body Resource.ID) (_ int)    { return }
func (self Implementation) SoftBodySetCollisionMask(body Resource.ID, mask int)   { return }
func (self Implementation) SoftBodyGetCollisionMask(body Resource.ID) (_ int)     { return }
func (self Implementation) SoftBodyAddCollisionException(body Resource.ID, body_b Resource.ID) {
	return
}
func (self Implementation) SoftBodyRemoveCollisionException(body Resource.ID, body_b Resource.ID) {
	return
}
func (self Implementation) SoftBodyGetCollisionExceptions(body Resource.ID) (_ gd.Array) { return }
func (self Implementation) SoftBodySetState(body Resource.ID, state gdclass.PhysicsServer3DBodyState, variant any) {
	return
}
func (self Implementation) SoftBodyGetState(body Resource.ID, state gdclass.PhysicsServer3DBodyState) (_ any) {
	return
}
func (self Implementation) SoftBodySetTransform(body Resource.ID, transform Transform3D.BasisOrigin) {
	return
}
func (self Implementation) SoftBodySetSimulationPrecision(body Resource.ID, simulation_precision int) {
	return
}
func (self Implementation) SoftBodyGetSimulationPrecision(body Resource.ID) (_ int)   { return }
func (self Implementation) SoftBodySetTotalMass(body Resource.ID, total_mass Float.X) { return }
func (self Implementation) SoftBodyGetTotalMass(body Resource.ID) (_ Float.X)         { return }
func (self Implementation) SoftBodySetLinearStiffness(body Resource.ID, linear_stiffness Float.X) {
	return
}
func (self Implementation) SoftBodyGetLinearStiffness(body Resource.ID) (_ Float.X) { return }
func (self Implementation) SoftBodySetPressureCoefficient(body Resource.ID, pressure_coefficient Float.X) {
	return
}
func (self Implementation) SoftBodyGetPressureCoefficient(body Resource.ID) (_ Float.X) { return }
func (self Implementation) SoftBodySetDampingCoefficient(body Resource.ID, damping_coefficient Float.X) {
	return
}
func (self Implementation) SoftBodyGetDampingCoefficient(body Resource.ID) (_ Float.X) { return }
func (self Implementation) SoftBodySetDragCoefficient(body Resource.ID, drag_coefficient Float.X) {
	return
}
func (self Implementation) SoftBodyGetDragCoefficient(body Resource.ID) (_ Float.X)  { return }
func (self Implementation) SoftBodySetMesh(body Resource.ID, mesh Resource.ID)       { return }
func (self Implementation) SoftBodyGetBounds(body Resource.ID) (_ AABB.PositionSize) { return }
func (self Implementation) SoftBodyMovePoint(body Resource.ID, point_index int, global_position Vector3.XYZ) {
	return
}
func (self Implementation) SoftBodyGetPointGlobalPosition(body Resource.ID, point_index int) (_ Vector3.XYZ) {
	return
}
func (self Implementation) SoftBodyRemoveAllPinnedPoints(body Resource.ID)                   { return }
func (self Implementation) SoftBodyPinPoint(body Resource.ID, point_index int, pin bool)     { return }
func (self Implementation) SoftBodyIsPointPinned(body Resource.ID, point_index int) (_ bool) { return }
func (self Implementation) JointCreate() (_ Resource.ID)                                     { return }
func (self Implementation) JointClear(joint Resource.ID)                                     { return }
func (self Implementation) JointMakePin(joint Resource.ID, body_A Resource.ID, local_A Vector3.XYZ, body_B Resource.ID, local_B Vector3.XYZ) {
	return
}
func (self Implementation) PinJointSetParam(joint Resource.ID, param gdclass.PhysicsServer3DPinJointParam, value Float.X) {
	return
}
func (self Implementation) PinJointGetParam(joint Resource.ID, param gdclass.PhysicsServer3DPinJointParam) (_ Float.X) {
	return
}
func (self Implementation) PinJointSetLocalA(joint Resource.ID, local_A Vector3.XYZ) { return }
func (self Implementation) PinJointGetLocalA(joint Resource.ID) (_ Vector3.XYZ)      { return }
func (self Implementation) PinJointSetLocalB(joint Resource.ID, local_B Vector3.XYZ) { return }
func (self Implementation) PinJointGetLocalB(joint Resource.ID) (_ Vector3.XYZ)      { return }
func (self Implementation) JointMakeHinge(joint Resource.ID, body_A Resource.ID, hinge_A Transform3D.BasisOrigin, body_B Resource.ID, hinge_B Transform3D.BasisOrigin) {
	return
}
func (self Implementation) JointMakeHingeSimple(joint Resource.ID, body_A Resource.ID, pivot_A Vector3.XYZ, axis_A Vector3.XYZ, body_B Resource.ID, pivot_B Vector3.XYZ, axis_B Vector3.XYZ) {
	return
}
func (self Implementation) HingeJointSetParam(joint Resource.ID, param gdclass.PhysicsServer3DHingeJointParam, value Float.X) {
	return
}
func (self Implementation) HingeJointGetParam(joint Resource.ID, param gdclass.PhysicsServer3DHingeJointParam) (_ Float.X) {
	return
}
func (self Implementation) HingeJointSetFlag(joint Resource.ID, flag gdclass.PhysicsServer3DHingeJointFlag, enabled bool) {
	return
}
func (self Implementation) HingeJointGetFlag(joint Resource.ID, flag gdclass.PhysicsServer3DHingeJointFlag) (_ bool) {
	return
}
func (self Implementation) JointMakeSlider(joint Resource.ID, body_A Resource.ID, local_ref_A Transform3D.BasisOrigin, body_B Resource.ID, local_ref_B Transform3D.BasisOrigin) {
	return
}
func (self Implementation) SliderJointSetParam(joint Resource.ID, param gdclass.PhysicsServer3DSliderJointParam, value Float.X) {
	return
}
func (self Implementation) SliderJointGetParam(joint Resource.ID, param gdclass.PhysicsServer3DSliderJointParam) (_ Float.X) {
	return
}
func (self Implementation) JointMakeConeTwist(joint Resource.ID, body_A Resource.ID, local_ref_A Transform3D.BasisOrigin, body_B Resource.ID, local_ref_B Transform3D.BasisOrigin) {
	return
}
func (self Implementation) ConeTwistJointSetParam(joint Resource.ID, param gdclass.PhysicsServer3DConeTwistJointParam, value Float.X) {
	return
}
func (self Implementation) ConeTwistJointGetParam(joint Resource.ID, param gdclass.PhysicsServer3DConeTwistJointParam) (_ Float.X) {
	return
}
func (self Implementation) JointMakeGeneric6dof(joint Resource.ID, body_A Resource.ID, local_ref_A Transform3D.BasisOrigin, body_B Resource.ID, local_ref_B Transform3D.BasisOrigin) {
	return
}
func (self Implementation) Generic6dofJointSetParam(joint Resource.ID, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam, value Float.X) {
	return
}
func (self Implementation) Generic6dofJointGetParam(joint Resource.ID, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam) (_ Float.X) {
	return
}
func (self Implementation) Generic6dofJointSetFlag(joint Resource.ID, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag, enable bool) {
	return
}
func (self Implementation) Generic6dofJointGetFlag(joint Resource.ID, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag) (_ bool) {
	return
}
func (self Implementation) JointGetType(joint Resource.ID) (_ gdclass.PhysicsServer3DJointType) {
	return
}
func (self Implementation) JointSetSolverPriority(joint Resource.ID, priority int) { return }
func (self Implementation) JointGetSolverPriority(joint Resource.ID) (_ int)       { return }
func (self Implementation) JointDisableCollisionsBetweenBodies(joint Resource.ID, disable bool) {
	return
}
func (self Implementation) JointIsDisabledCollisionsBetweenBodies(joint Resource.ID) (_ bool) { return }
func (self Implementation) FreeRid(rid Resource.ID)                                           { return }
func (self Implementation) SetActive(active bool)                                             { return }
func (self Implementation) Init()                                                             { return }
func (self Implementation) Step(step Float.X)                                                 { return }
func (self Implementation) Sync()                                                             { return }
func (self Implementation) FlushQueries()                                                     { return }
func (self Implementation) EndSync()                                                          { return }
func (self Implementation) Finish()                                                           { return }
func (self Implementation) IsFlushingQueries() (_ bool)                                       { return }
func (self Implementation) GetProcessInfo(process_info gdclass.PhysicsServer3DProcessInfo) (_ int) {
	return
}
func (Instance) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _sphere_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _box_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _capsule_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _cylinder_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _heightmap_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _custom_shape_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _shape_set_data(impl func(ptr unsafe.Pointer, shape Resource.ID, data any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var data = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 1))
		defer pointers.End(data)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, data.Interface())
	}
}
func (Instance) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape Resource.ID, bias Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var bias = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, Float.X(bias))
	}
}
func (Instance) _shape_set_margin(impl func(ptr unsafe.Pointer, shape Resource.ID, margin Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var margin = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, Float.X(margin))
	}
}
func (Instance) _shape_get_margin(impl func(ptr unsafe.Pointer, shape Resource.ID) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _shape_get_type(impl func(ptr unsafe.Pointer, shape Resource.ID) gdclass.PhysicsServer3DShapeType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _shape_get_data(impl func(ptr unsafe.Pointer, shape Resource.ID) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape Resource.ID) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _space_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _space_set_active(impl func(ptr unsafe.Pointer, space Resource.ID, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var active = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, active)
	}
}
func (Instance) _space_is_active(impl func(ptr unsafe.Pointer, space Resource.ID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _space_set_param(impl func(ptr unsafe.Pointer, space Resource.ID, param gdclass.PhysicsServer3DSpaceParameter, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSpaceParameter](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, param, Float.X(value))
	}
}
func (Instance) _space_get_param(impl func(ptr unsafe.Pointer, space Resource.ID, param gdclass.PhysicsServer3DSpaceParameter) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSpaceParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _space_get_direct_state(impl func(ptr unsafe.Pointer, space Resource.ID) [1]gdclass.PhysicsDirectSpaceState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space Resource.ID, max_contacts int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var max_contacts = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, int(max_contacts))
	}
}
func (Instance) _space_get_contacts(impl func(ptr unsafe.Pointer, space Resource.ID) []Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&ret))))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _space_get_contact_count(impl func(ptr unsafe.Pointer, space Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _area_set_space(impl func(ptr unsafe.Pointer, area Resource.ID, space Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, space)
	}
}
func (Instance) _area_get_space(impl func(ptr unsafe.Pointer, area Resource.ID) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _area_add_shape(impl func(ptr unsafe.Pointer, area Resource.ID, shape Resource.ID, transform Transform3D.BasisOrigin, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape = gd.UnsafeGet[gd.RID](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var disabled = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape, transform, disabled)
	}
}
func (Instance) _area_set_shape(impl func(ptr unsafe.Pointer, area Resource.ID, shape_idx int, shape Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var shape = gd.UnsafeGet[gd.RID](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), shape)
	}
}
func (Instance) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area Resource.ID, shape_idx int, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), transform)
	}
}
func (Instance) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area Resource.ID, shape_idx int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var disabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), disabled)
	}
}
func (Instance) _area_get_shape_count(impl func(ptr unsafe.Pointer, area Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_get_shape(impl func(ptr unsafe.Pointer, area Resource.ID, shape_idx int) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area Resource.ID, shape_idx int) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, gd.Transform3D(ret))
	}
}
func (Instance) _area_remove_shape(impl func(ptr unsafe.Pointer, area Resource.ID, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx))
	}
}
func (Instance) _area_clear_shapes(impl func(ptr unsafe.Pointer, area Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area)
	}
}
func (Instance) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area Resource.ID, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(id))
	}
}
func (Instance) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_set_param(impl func(ptr unsafe.Pointer, area Resource.ID, param gdclass.PhysicsServer3DAreaParameter, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DAreaParameter](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, param, value.Interface())
	}
}
func (Instance) _area_set_transform(impl func(ptr unsafe.Pointer, area Resource.ID, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, transform)
	}
}
func (Instance) _area_get_param(impl func(ptr unsafe.Pointer, area Resource.ID, param gdclass.PhysicsServer3DAreaParameter) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DAreaParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _area_get_transform(impl func(ptr unsafe.Pointer, area Resource.ID) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Transform3D(ret))
	}
}
func (Instance) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area Resource.ID, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(layer))
	}
}
func (Instance) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area Resource.ID, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(mask))
	}
}
func (Instance) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_set_monitorable(impl func(ptr unsafe.Pointer, area Resource.ID, monitorable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var monitorable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, monitorable)
	}
}
func (Instance) _area_set_ray_pickable(impl func(ptr unsafe.Pointer, area Resource.ID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, enable)
	}
}
func (Instance) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area Resource.ID, callback Callable.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var callback = pointers.New[gd.Callable](gd.UnsafeGet[callablePointers](p_args, 1))
		defer pointers.End(callback)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}
func (Instance) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area Resource.ID, callback Callable.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var callback = pointers.New[gd.Callable](gd.UnsafeGet[callablePointers](p_args, 1))
		defer pointers.End(callback)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}
func (Instance) _body_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_space(impl func(ptr unsafe.Pointer, body Resource.ID, space Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}
func (Instance) _body_get_space(impl func(ptr unsafe.Pointer, body Resource.ID) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_mode(impl func(ptr unsafe.Pointer, body Resource.ID, mode gdclass.PhysicsServer3DBodyMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mode = gd.UnsafeGet[gdclass.PhysicsServer3DBodyMode](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}
func (Instance) _body_get_mode(impl func(ptr unsafe.Pointer, body Resource.ID) gdclass.PhysicsServer3DBodyMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_add_shape(impl func(ptr unsafe.Pointer, body Resource.ID, shape Resource.ID, transform Transform3D.BasisOrigin, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape = gd.UnsafeGet[gd.RID](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var disabled = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape, transform, disabled)
	}
}
func (Instance) _body_set_shape(impl func(ptr unsafe.Pointer, body Resource.ID, shape_idx int, shape Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var shape = gd.UnsafeGet[gd.RID](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), shape)
	}
}
func (Instance) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body Resource.ID, shape_idx int, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), transform)
	}
}
func (Instance) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body Resource.ID, shape_idx int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var disabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), disabled)
	}
}
func (Instance) _body_get_shape_count(impl func(ptr unsafe.Pointer, body Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_get_shape(impl func(ptr unsafe.Pointer, body Resource.ID, shape_idx int) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body Resource.ID, shape_idx int) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, gd.Transform3D(ret))
	}
}
func (Instance) _body_remove_shape(impl func(ptr unsafe.Pointer, body Resource.ID, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx))
	}
}
func (Instance) _body_clear_shapes(impl func(ptr unsafe.Pointer, body Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}
func (Instance) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body Resource.ID, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(id))
	}
}
func (Instance) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_enable_continuous_collision_detection(impl func(ptr unsafe.Pointer, body Resource.ID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _body_is_continuous_collision_detection_enabled(impl func(ptr unsafe.Pointer, body Resource.ID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body Resource.ID, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(layer))
	}
}
func (Instance) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body Resource.ID, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(mask))
	}
}
func (Instance) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body Resource.ID, priority Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var priority = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(priority))
	}
}
func (Instance) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body Resource.ID) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _body_set_user_flags(impl func(ptr unsafe.Pointer, body Resource.ID, flags int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var flags = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(flags))
	}
}
func (Instance) _body_get_user_flags(impl func(ptr unsafe.Pointer, body Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_param(impl func(ptr unsafe.Pointer, body Resource.ID, param gdclass.PhysicsServer3DBodyParameter, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DBodyParameter](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, param, value.Interface())
	}
}
func (Instance) _body_get_param(impl func(ptr unsafe.Pointer, body Resource.ID, param gdclass.PhysicsServer3DBodyParameter) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DBodyParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}
func (Instance) _body_set_state(impl func(ptr unsafe.Pointer, body Resource.ID, state gdclass.PhysicsServer3DBodyState, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, value.Interface())
	}
}
func (Instance) _body_get_state(impl func(ptr unsafe.Pointer, body Resource.ID, state gdclass.PhysicsServer3DBodyState) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body Resource.ID, impulse Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}
func (Instance) _body_apply_impulse(impl func(ptr unsafe.Pointer, body Resource.ID, impulse Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse, position)
	}
}
func (Instance) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body Resource.ID, impulse Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}
func (Instance) _body_apply_central_force(impl func(ptr unsafe.Pointer, body Resource.ID, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}
func (Instance) _body_apply_force(impl func(ptr unsafe.Pointer, body Resource.ID, force Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}
func (Instance) _body_apply_torque(impl func(ptr unsafe.Pointer, body Resource.ID, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}
func (Instance) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body Resource.ID, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}
func (Instance) _body_add_constant_force(impl func(ptr unsafe.Pointer, body Resource.ID, force Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}
func (Instance) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body Resource.ID, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}
func (Instance) _body_set_constant_force(impl func(ptr unsafe.Pointer, body Resource.ID, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}
func (Instance) _body_get_constant_force(impl func(ptr unsafe.Pointer, body Resource.ID) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body Resource.ID, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}
func (Instance) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body Resource.ID) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body Resource.ID, axis_velocity Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis_velocity = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis_velocity)
	}
}
func (Instance) _body_set_axis_lock(impl func(ptr unsafe.Pointer, body Resource.ID, axis gdclass.PhysicsServer3DBodyAxis, lock bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gdclass.PhysicsServer3DBodyAxis](p_args, 1)
		var lock = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis, lock)
	}
}
func (Instance) _body_is_axis_locked(impl func(ptr unsafe.Pointer, body Resource.ID, axis gdclass.PhysicsServer3DBodyAxis) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gdclass.PhysicsServer3DBodyAxis](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, axis)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body Resource.ID, excepted_body Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}
func (Instance) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body Resource.ID, excepted_body Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}
func (Instance) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body Resource.ID) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body Resource.ID, amount int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var amount = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(amount))
	}
}
func (Instance) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body Resource.ID, threshold Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var threshold = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(threshold))
	}
}
func (Instance) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body Resource.ID) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body Resource.ID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body Resource.ID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body Resource.ID, callable Callable.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var callable = pointers.New[gd.Callable](gd.UnsafeGet[callablePointers](p_args, 1))
		defer pointers.End(callable)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable)
	}
}
func (Instance) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body Resource.ID, callable Callable.Any, userdata any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var callable = pointers.New[gd.Callable](gd.UnsafeGet[callablePointers](p_args, 1))
		defer pointers.End(callable)
		var userdata = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		defer pointers.End(userdata)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable, userdata.Interface())
	}
}
func (Instance) _body_set_ray_pickable(impl func(ptr unsafe.Pointer, body Resource.ID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _body_test_motion(impl func(ptr unsafe.Pointer, body Resource.ID, from Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, max_collisions int, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var from = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var max_collisions = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_separation_ray = gd.UnsafeGet[bool](p_args, 5)
		var recovery_as_collision = gd.UnsafeGet[bool](p_args, 6)
		var result = gd.UnsafeGet[*MotionResult](p_args, 7)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, Float.X(margin), int(max_collisions), collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_get_direct_state(impl func(ptr unsafe.Pointer, body Resource.ID) [1]gdclass.PhysicsDirectBodyState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _soft_body_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _soft_body_update_rendering_server(impl func(ptr unsafe.Pointer, body Resource.ID, rendering_server_handler [1]gdclass.PhysicsServer3DRenderingServerHandler)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var rendering_server_handler = [1]gdclass.PhysicsServer3DRenderingServerHandler{pointers.New[gdclass.PhysicsServer3DRenderingServerHandler]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(rendering_server_handler[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, rendering_server_handler)
	}
}
func (Instance) _soft_body_set_space(impl func(ptr unsafe.Pointer, body Resource.ID, space Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}
func (Instance) _soft_body_get_space(impl func(ptr unsafe.Pointer, body Resource.ID) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _soft_body_set_ray_pickable(impl func(ptr unsafe.Pointer, body Resource.ID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _soft_body_set_collision_layer(impl func(ptr unsafe.Pointer, body Resource.ID, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(layer))
	}
}
func (Instance) _soft_body_get_collision_layer(impl func(ptr unsafe.Pointer, body Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _soft_body_set_collision_mask(impl func(ptr unsafe.Pointer, body Resource.ID, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(mask))
	}
}
func (Instance) _soft_body_get_collision_mask(impl func(ptr unsafe.Pointer, body Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _soft_body_add_collision_exception(impl func(ptr unsafe.Pointer, body Resource.ID, body_b Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_b = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}
func (Instance) _soft_body_remove_collision_exception(impl func(ptr unsafe.Pointer, body Resource.ID, body_b Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_b = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}
func (Instance) _soft_body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body Resource.ID) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _soft_body_set_state(impl func(ptr unsafe.Pointer, body Resource.ID, state gdclass.PhysicsServer3DBodyState, variant any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)
		var variant = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		defer pointers.End(variant)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, variant.Interface())
	}
}
func (Instance) _soft_body_get_state(impl func(ptr unsafe.Pointer, body Resource.ID, state gdclass.PhysicsServer3DBodyState) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _soft_body_set_transform(impl func(ptr unsafe.Pointer, body Resource.ID, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, transform)
	}
}
func (Instance) _soft_body_set_simulation_precision(impl func(ptr unsafe.Pointer, body Resource.ID, simulation_precision int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var simulation_precision = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(simulation_precision))
	}
}
func (Instance) _soft_body_get_simulation_precision(impl func(ptr unsafe.Pointer, body Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _soft_body_set_total_mass(impl func(ptr unsafe.Pointer, body Resource.ID, total_mass Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var total_mass = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(total_mass))
	}
}
func (Instance) _soft_body_get_total_mass(impl func(ptr unsafe.Pointer, body Resource.ID) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_linear_stiffness(impl func(ptr unsafe.Pointer, body Resource.ID, linear_stiffness Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var linear_stiffness = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(linear_stiffness))
	}
}
func (Instance) _soft_body_get_linear_stiffness(impl func(ptr unsafe.Pointer, body Resource.ID) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_pressure_coefficient(impl func(ptr unsafe.Pointer, body Resource.ID, pressure_coefficient Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var pressure_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(pressure_coefficient))
	}
}
func (Instance) _soft_body_get_pressure_coefficient(impl func(ptr unsafe.Pointer, body Resource.ID) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_damping_coefficient(impl func(ptr unsafe.Pointer, body Resource.ID, damping_coefficient Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var damping_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(damping_coefficient))
	}
}
func (Instance) _soft_body_get_damping_coefficient(impl func(ptr unsafe.Pointer, body Resource.ID) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_drag_coefficient(impl func(ptr unsafe.Pointer, body Resource.ID, drag_coefficient Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var drag_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(drag_coefficient))
	}
}
func (Instance) _soft_body_get_drag_coefficient(impl func(ptr unsafe.Pointer, body Resource.ID) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_mesh(impl func(ptr unsafe.Pointer, body Resource.ID, mesh Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mesh = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mesh)
	}
}
func (Instance) _soft_body_get_bounds(impl func(ptr unsafe.Pointer, body Resource.ID) AABB.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.AABB(ret))
	}
}
func (Instance) _soft_body_move_point(impl func(ptr unsafe.Pointer, body Resource.ID, point_index int, global_position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var global_position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(point_index), global_position)
	}
}
func (Instance) _soft_body_get_point_global_position(impl func(ptr unsafe.Pointer, body Resource.ID, point_index int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(point_index))
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _soft_body_remove_all_pinned_points(impl func(ptr unsafe.Pointer, body Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}
func (Instance) _soft_body_pin_point(impl func(ptr unsafe.Pointer, body Resource.ID, point_index int, pin bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var pin = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(point_index), pin)
	}
}
func (Instance) _soft_body_is_point_pinned(impl func(ptr unsafe.Pointer, body Resource.ID, point_index int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(point_index))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_create(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_clear(impl func(ptr unsafe.Pointer, joint Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint)
	}
}
func (Instance) _joint_make_pin(impl func(ptr unsafe.Pointer, joint Resource.ID, body_A Resource.ID, local_A Vector3.XYZ, body_B Resource.ID, local_B Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_A = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_B = gd.UnsafeGet[gd.Vector3](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_A, body_B, local_B)
	}
}
func (Instance) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint Resource.ID, param gdclass.PhysicsServer3DPinJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DPinJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}
func (Instance) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint Resource.ID, param gdclass.PhysicsServer3DPinJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DPinJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _pin_joint_set_local_a(impl func(ptr unsafe.Pointer, joint Resource.ID, local_A Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var local_A = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_A)
	}
}
func (Instance) _pin_joint_get_local_a(impl func(ptr unsafe.Pointer, joint Resource.ID) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _pin_joint_set_local_b(impl func(ptr unsafe.Pointer, joint Resource.ID, local_B Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var local_B = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_B)
	}
}
func (Instance) _pin_joint_get_local_b(impl func(ptr unsafe.Pointer, joint Resource.ID) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _joint_make_hinge(impl func(ptr unsafe.Pointer, joint Resource.ID, body_A Resource.ID, hinge_A Transform3D.BasisOrigin, body_B Resource.ID, hinge_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var hinge_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var hinge_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, hinge_A, body_B, hinge_B)
	}
}
func (Instance) _joint_make_hinge_simple(impl func(ptr unsafe.Pointer, joint Resource.ID, body_A Resource.ID, pivot_A Vector3.XYZ, axis_A Vector3.XYZ, body_B Resource.ID, pivot_B Vector3.XYZ, axis_B Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var pivot_A = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var axis_A = gd.UnsafeGet[gd.Vector3](p_args, 3)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 4)
		var pivot_B = gd.UnsafeGet[gd.Vector3](p_args, 5)
		var axis_B = gd.UnsafeGet[gd.Vector3](p_args, 6)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, pivot_A, axis_A, body_B, pivot_B, axis_B)
	}
}
func (Instance) _hinge_joint_set_param(impl func(ptr unsafe.Pointer, joint Resource.ID, param gdclass.PhysicsServer3DHingeJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}
func (Instance) _hinge_joint_get_param(impl func(ptr unsafe.Pointer, joint Resource.ID, param gdclass.PhysicsServer3DHingeJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _hinge_joint_set_flag(impl func(ptr unsafe.Pointer, joint Resource.ID, flag gdclass.PhysicsServer3DHingeJointFlag, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointFlag](p_args, 1)
		var enabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, flag, enabled)
	}
}
func (Instance) _hinge_joint_get_flag(impl func(ptr unsafe.Pointer, joint Resource.ID, flag gdclass.PhysicsServer3DHingeJointFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointFlag](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_make_slider(impl func(ptr unsafe.Pointer, joint Resource.ID, body_A Resource.ID, local_ref_A Transform3D.BasisOrigin, body_B Resource.ID, local_ref_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}
func (Instance) _slider_joint_set_param(impl func(ptr unsafe.Pointer, joint Resource.ID, param gdclass.PhysicsServer3DSliderJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSliderJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}
func (Instance) _slider_joint_get_param(impl func(ptr unsafe.Pointer, joint Resource.ID, param gdclass.PhysicsServer3DSliderJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSliderJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _joint_make_cone_twist(impl func(ptr unsafe.Pointer, joint Resource.ID, body_A Resource.ID, local_ref_A Transform3D.BasisOrigin, body_B Resource.ID, local_ref_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}
func (Instance) _cone_twist_joint_set_param(impl func(ptr unsafe.Pointer, joint Resource.ID, param gdclass.PhysicsServer3DConeTwistJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DConeTwistJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}
func (Instance) _cone_twist_joint_get_param(impl func(ptr unsafe.Pointer, joint Resource.ID, param gdclass.PhysicsServer3DConeTwistJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DConeTwistJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _joint_make_generic_6dof(impl func(ptr unsafe.Pointer, joint Resource.ID, body_A Resource.ID, local_ref_A Transform3D.BasisOrigin, body_B Resource.ID, local_ref_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}
func (Instance) _generic_6dof_joint_set_param(impl func(ptr unsafe.Pointer, joint Resource.ID, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)
		var value = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, param, Float.X(value))
	}
}
func (Instance) _generic_6dof_joint_get_param(impl func(ptr unsafe.Pointer, joint Resource.ID, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _generic_6dof_joint_set_flag(impl func(ptr unsafe.Pointer, joint Resource.ID, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)
		var enable = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, flag, enable)
	}
}
func (Instance) _generic_6dof_joint_get_flag(impl func(ptr unsafe.Pointer, joint Resource.ID, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, flag)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_get_type(impl func(ptr unsafe.Pointer, joint Resource.ID) gdclass.PhysicsServer3DJointType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_set_solver_priority(impl func(ptr unsafe.Pointer, joint Resource.ID, priority int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var priority = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, int(priority))
	}
}
func (Instance) _joint_get_solver_priority(impl func(ptr unsafe.Pointer, joint Resource.ID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint Resource.ID, disable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var disable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, disable)
	}
}
func (Instance) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint Resource.ID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _free_rid(impl func(ptr unsafe.Pointer, rid Resource.ID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}
func (Instance) _set_active(impl func(ptr unsafe.Pointer, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var active = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, active)
	}
}
func (Instance) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _step(impl func(ptr unsafe.Pointer, step Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var step = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(step))
	}
}
func (Instance) _sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _flush_queries(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _end_sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_process_info(impl func(ptr unsafe.Pointer, process_info gdclass.PhysicsServer3DProcessInfo) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var process_info = gd.UnsafeGet[gdclass.PhysicsServer3DProcessInfo](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (self Instance) BodyTestMotionIsExcludingBody(body Resource.ID) bool {
	return bool(class(self).BodyTestMotionIsExcludingBody(body))
}
func (self Instance) BodyTestMotionIsExcludingObject(obj int) bool {
	return bool(class(self).BodyTestMotionIsExcludingObject(gd.Int(obj)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsServer3DExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsServer3DExtension"))
	casted := Instance{*(*gdclass.PhysicsServer3DExtension)(unsafe.Pointer(&object))}
	return casted
}

func (class) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _sphere_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _box_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _capsule_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _cylinder_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _heightmap_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _custom_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _shape_set_data(impl func(ptr unsafe.Pointer, shape gd.RID, data gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var data = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, data)
	}
}

func (class) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID, bias gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var bias = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, bias)
	}
}

func (class) _shape_set_margin(impl func(ptr unsafe.Pointer, shape gd.RID, margin gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var margin = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, margin)
	}
}

func (class) _shape_get_margin(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _shape_get_type(impl func(ptr unsafe.Pointer, shape gd.RID) gdclass.PhysicsServer3DShapeType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _shape_get_data(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_set_active(impl func(ptr unsafe.Pointer, space gd.RID, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var active = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, active)
	}
}

func (class) _space_is_active(impl func(ptr unsafe.Pointer, space gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_set_param(impl func(ptr unsafe.Pointer, space gd.RID, param gdclass.PhysicsServer3DSpaceParameter, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSpaceParameter](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, param, value)
	}
}

func (class) _space_get_param(impl func(ptr unsafe.Pointer, space gd.RID, param gdclass.PhysicsServer3DSpaceParameter) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSpaceParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_get_direct_state(impl func(ptr unsafe.Pointer, space gd.RID) [1]gdclass.PhysicsDirectSpaceState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space gd.RID, max_contacts gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var max_contacts = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, max_contacts)
	}
}

func (class) _space_get_contacts(impl func(ptr unsafe.Pointer, space gd.RID) gd.PackedVector3Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _space_get_contact_count(impl func(ptr unsafe.Pointer, space gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_space(impl func(ptr unsafe.Pointer, area gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, space)
	}
}

func (class) _area_get_space(impl func(ptr unsafe.Pointer, area gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_add_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape = gd.UnsafeGet[gd.RID](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var disabled = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape, transform, disabled)
	}
}

func (class) _area_set_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, shape gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var shape = gd.UnsafeGet[gd.RID](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, shape)
	}
}

func (class) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, transform)
	}
}

func (class) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var disabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, disabled)
	}
}

func (class) _area_get_shape_count(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_get_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_remove_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx)
	}
}

func (class) _area_clear_shapes(impl func(ptr unsafe.Pointer, area gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area)
	}
}

func (class) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID, id gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, id)
	}
}

func (class) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_param(impl func(ptr unsafe.Pointer, area gd.RID, param gdclass.PhysicsServer3DAreaParameter, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DAreaParameter](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, param, value)
	}
}

func (class) _area_set_transform(impl func(ptr unsafe.Pointer, area gd.RID, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, transform)
	}
}

func (class) _area_get_param(impl func(ptr unsafe.Pointer, area gd.RID, param gdclass.PhysicsServer3DAreaParameter) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DAreaParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _area_get_transform(impl func(ptr unsafe.Pointer, area gd.RID) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID, layer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, layer)
	}
}

func (class) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID, mask gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, mask)
	}
}

func (class) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_monitorable(impl func(ptr unsafe.Pointer, area gd.RID, monitorable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var monitorable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, monitorable)
	}
}

func (class) _area_set_ray_pickable(impl func(ptr unsafe.Pointer, area gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, enable)
	}
}

func (class) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var callback = pointers.New[gd.Callable](gd.UnsafeGet[callablePointers](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

func (class) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var callback = pointers.New[gd.Callable](gd.UnsafeGet[callablePointers](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

func (class) _body_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_space(impl func(ptr unsafe.Pointer, body gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}

func (class) _body_get_space(impl func(ptr unsafe.Pointer, body gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_mode(impl func(ptr unsafe.Pointer, body gd.RID, mode gdclass.PhysicsServer3DBodyMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mode = gd.UnsafeGet[gdclass.PhysicsServer3DBodyMode](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}

func (class) _body_get_mode(impl func(ptr unsafe.Pointer, body gd.RID) gdclass.PhysicsServer3DBodyMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_add_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape = gd.UnsafeGet[gd.RID](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var disabled = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape, transform, disabled)
	}
}

func (class) _body_set_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, shape gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var shape = gd.UnsafeGet[gd.RID](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, shape)
	}
}

func (class) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, transform)
	}
}

func (class) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var disabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, disabled)
	}
}

func (class) _body_get_shape_count(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_get_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_remove_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx)
	}
}

func (class) _body_clear_shapes(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

func (class) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID, id gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, id)
	}
}

func (class) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_enable_continuous_collision_detection(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _body_is_continuous_collision_detection_enabled(impl func(ptr unsafe.Pointer, body gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID, layer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, layer)
	}
}

func (class) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID, mask gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mask)
	}
}

func (class) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID, priority gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var priority = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, priority)
	}
}

func (class) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_user_flags(impl func(ptr unsafe.Pointer, body gd.RID, flags gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var flags = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, flags)
	}
}

func (class) _body_get_user_flags(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_param(impl func(ptr unsafe.Pointer, body gd.RID, param gdclass.PhysicsServer3DBodyParameter, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DBodyParameter](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, param, value)
	}
}

func (class) _body_get_param(impl func(ptr unsafe.Pointer, body gd.RID, param gdclass.PhysicsServer3DBodyParameter) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DBodyParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

func (class) _body_set_state(impl func(ptr unsafe.Pointer, body gd.RID, state gdclass.PhysicsServer3DBodyState, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, value)
	}
}

func (class) _body_get_state(impl func(ptr unsafe.Pointer, body gd.RID, state gdclass.PhysicsServer3DBodyState) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}

func (class) _body_apply_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse, position)
	}
}

func (class) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}

func (class) _body_apply_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

func (class) _body_apply_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

func (class) _body_apply_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

func (class) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

func (class) _body_add_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

func (class) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

func (class) _body_set_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

func (class) _body_get_constant_force(impl func(ptr unsafe.Pointer, body gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

func (class) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body gd.RID, axis_velocity gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis_velocity = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis_velocity)
	}
}

func (class) _body_set_axis_lock(impl func(ptr unsafe.Pointer, body gd.RID, axis gdclass.PhysicsServer3DBodyAxis, lock bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gdclass.PhysicsServer3DBodyAxis](p_args, 1)
		var lock = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis, lock)
	}
}

func (class) _body_is_axis_locked(impl func(ptr unsafe.Pointer, body gd.RID, axis gdclass.PhysicsServer3DBodyAxis) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gdclass.PhysicsServer3DBodyAxis](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, axis)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

func (class) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

func (class) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body gd.RID) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID, amount gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var amount = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, amount)
	}
}

func (class) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID, threshold gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var threshold = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, threshold)
	}
}

func (class) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var callable = pointers.New[gd.Callable](gd.UnsafeGet[callablePointers](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable)
	}
}

func (class) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable, userdata gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var callable = pointers.New[gd.Callable](gd.UnsafeGet[callablePointers](p_args, 1))
		var userdata = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable, userdata)
	}
}

func (class) _body_set_ray_pickable(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _body_test_motion(impl func(ptr unsafe.Pointer, body gd.RID, from gd.Transform3D, motion gd.Vector3, margin gd.Float, max_collisions gd.Int, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var from = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var max_collisions = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_separation_ray = gd.UnsafeGet[bool](p_args, 5)
		var recovery_as_collision = gd.UnsafeGet[bool](p_args, 6)
		var result = gd.UnsafeGet[*MotionResult](p_args, 7)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, margin, max_collisions, collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_get_direct_state(impl func(ptr unsafe.Pointer, body gd.RID) [1]gdclass.PhysicsDirectBodyState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _soft_body_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_update_rendering_server(impl func(ptr unsafe.Pointer, body gd.RID, rendering_server_handler [1]gdclass.PhysicsServer3DRenderingServerHandler)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var rendering_server_handler = [1]gdclass.PhysicsServer3DRenderingServerHandler{pointers.New[gdclass.PhysicsServer3DRenderingServerHandler]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(rendering_server_handler[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, rendering_server_handler)
	}
}

func (class) _soft_body_set_space(impl func(ptr unsafe.Pointer, body gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}

func (class) _soft_body_get_space(impl func(ptr unsafe.Pointer, body gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_ray_pickable(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _soft_body_set_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID, layer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, layer)
	}
}

func (class) _soft_body_get_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID, mask gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mask)
	}
}

func (class) _soft_body_get_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_add_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, body_b gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_b = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}

func (class) _soft_body_remove_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, body_b gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_b = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}

func (class) _soft_body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body gd.RID) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _soft_body_set_state(impl func(ptr unsafe.Pointer, body gd.RID, state gdclass.PhysicsServer3DBodyState, variant gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)
		var variant = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, variant)
	}
}

func (class) _soft_body_get_state(impl func(ptr unsafe.Pointer, body gd.RID, state gdclass.PhysicsServer3DBodyState) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _soft_body_set_transform(impl func(ptr unsafe.Pointer, body gd.RID, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, transform)
	}
}

func (class) _soft_body_set_simulation_precision(impl func(ptr unsafe.Pointer, body gd.RID, simulation_precision gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var simulation_precision = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, simulation_precision)
	}
}

func (class) _soft_body_get_simulation_precision(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_total_mass(impl func(ptr unsafe.Pointer, body gd.RID, total_mass gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var total_mass = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, total_mass)
	}
}

func (class) _soft_body_get_total_mass(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_linear_stiffness(impl func(ptr unsafe.Pointer, body gd.RID, linear_stiffness gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var linear_stiffness = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, linear_stiffness)
	}
}

func (class) _soft_body_get_linear_stiffness(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_pressure_coefficient(impl func(ptr unsafe.Pointer, body gd.RID, pressure_coefficient gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var pressure_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, pressure_coefficient)
	}
}

func (class) _soft_body_get_pressure_coefficient(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_damping_coefficient(impl func(ptr unsafe.Pointer, body gd.RID, damping_coefficient gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var damping_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, damping_coefficient)
	}
}

func (class) _soft_body_get_damping_coefficient(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_drag_coefficient(impl func(ptr unsafe.Pointer, body gd.RID, drag_coefficient gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var drag_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, drag_coefficient)
	}
}

func (class) _soft_body_get_drag_coefficient(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_mesh(impl func(ptr unsafe.Pointer, body gd.RID, mesh gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mesh = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mesh)
	}
}

func (class) _soft_body_get_bounds(impl func(ptr unsafe.Pointer, body gd.RID) gd.AABB) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_move_point(impl func(ptr unsafe.Pointer, body gd.RID, point_index gd.Int, global_position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var global_position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, point_index, global_position)
	}
}

func (class) _soft_body_get_point_global_position(impl func(ptr unsafe.Pointer, body gd.RID, point_index gd.Int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, point_index)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_remove_all_pinned_points(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

func (class) _soft_body_pin_point(impl func(ptr unsafe.Pointer, body gd.RID, point_index gd.Int, pin bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var pin = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, point_index, pin)
	}
}

func (class) _soft_body_is_point_pinned(impl func(ptr unsafe.Pointer, body gd.RID, point_index gd.Int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, point_index)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_clear(impl func(ptr unsafe.Pointer, joint gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint)
	}
}

func (class) _joint_make_pin(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_A gd.Vector3, body_B gd.RID, local_B gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_A = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_B = gd.UnsafeGet[gd.Vector3](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_A, body_B, local_B)
	}
}

func (class) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer3DPinJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DPinJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer3DPinJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DPinJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _pin_joint_set_local_a(impl func(ptr unsafe.Pointer, joint gd.RID, local_A gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var local_A = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_A)
	}
}

func (class) _pin_joint_get_local_a(impl func(ptr unsafe.Pointer, joint gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _pin_joint_set_local_b(impl func(ptr unsafe.Pointer, joint gd.RID, local_B gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var local_B = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_B)
	}
}

func (class) _pin_joint_get_local_b(impl func(ptr unsafe.Pointer, joint gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_hinge(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, hinge_A gd.Transform3D, body_B gd.RID, hinge_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var hinge_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var hinge_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, hinge_A, body_B, hinge_B)
	}
}

func (class) _joint_make_hinge_simple(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, pivot_A gd.Vector3, axis_A gd.Vector3, body_B gd.RID, pivot_B gd.Vector3, axis_B gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var pivot_A = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var axis_A = gd.UnsafeGet[gd.Vector3](p_args, 3)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 4)
		var pivot_B = gd.UnsafeGet[gd.Vector3](p_args, 5)
		var axis_B = gd.UnsafeGet[gd.Vector3](p_args, 6)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, pivot_A, axis_A, body_B, pivot_B, axis_B)
	}
}

func (class) _hinge_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer3DHingeJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _hinge_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer3DHingeJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _hinge_joint_set_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag gdclass.PhysicsServer3DHingeJointFlag, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointFlag](p_args, 1)
		var enabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, flag, enabled)
	}
}

func (class) _hinge_joint_get_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag gdclass.PhysicsServer3DHingeJointFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointFlag](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_slider(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}

func (class) _slider_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer3DSliderJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSliderJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _slider_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer3DSliderJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSliderJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_cone_twist(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}

func (class) _cone_twist_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer3DConeTwistJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DConeTwistJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _cone_twist_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer3DConeTwistJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DConeTwistJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_generic_6dof(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}

func (class) _generic_6dof_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)
		var value = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, param, value)
	}
}

func (class) _generic_6dof_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var param = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _generic_6dof_joint_set_flag(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)
		var enable = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, flag, enable)
	}
}

func (class) _generic_6dof_joint_get_flag(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, flag)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_get_type(impl func(ptr unsafe.Pointer, joint gd.RID) gdclass.PhysicsServer3DJointType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_set_solver_priority(impl func(ptr unsafe.Pointer, joint gd.RID, priority gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var priority = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, priority)
	}
}

func (class) _joint_get_solver_priority(impl func(ptr unsafe.Pointer, joint gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID, disable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var disable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, disable)
	}
}

func (class) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _free_rid(impl func(ptr unsafe.Pointer, rid gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}

func (class) _set_active(impl func(ptr unsafe.Pointer, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var active = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, active)
	}
}

func (class) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _step(impl func(ptr unsafe.Pointer, step gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var step = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, step)
	}
}

func (class) _sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _flush_queries(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _end_sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_process_info(impl func(ptr unsafe.Pointer, process_info gdclass.PhysicsServer3DProcessInfo) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var process_info = gd.UnsafeGet[gdclass.PhysicsServer3DProcessInfo](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) BodyTestMotionIsExcludingBody(body gd.RID) bool {
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DExtension.Bind_body_test_motion_is_excluding_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) BodyTestMotionIsExcludingObject(obj gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, obj)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DExtension.Bind_body_test_motion_is_excluding_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsServer3DExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicsServer3DExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_world_boundary_shape_create":
		return reflect.ValueOf(self._world_boundary_shape_create)
	case "_separation_ray_shape_create":
		return reflect.ValueOf(self._separation_ray_shape_create)
	case "_sphere_shape_create":
		return reflect.ValueOf(self._sphere_shape_create)
	case "_box_shape_create":
		return reflect.ValueOf(self._box_shape_create)
	case "_capsule_shape_create":
		return reflect.ValueOf(self._capsule_shape_create)
	case "_cylinder_shape_create":
		return reflect.ValueOf(self._cylinder_shape_create)
	case "_convex_polygon_shape_create":
		return reflect.ValueOf(self._convex_polygon_shape_create)
	case "_concave_polygon_shape_create":
		return reflect.ValueOf(self._concave_polygon_shape_create)
	case "_heightmap_shape_create":
		return reflect.ValueOf(self._heightmap_shape_create)
	case "_custom_shape_create":
		return reflect.ValueOf(self._custom_shape_create)
	case "_shape_set_data":
		return reflect.ValueOf(self._shape_set_data)
	case "_shape_set_custom_solver_bias":
		return reflect.ValueOf(self._shape_set_custom_solver_bias)
	case "_shape_set_margin":
		return reflect.ValueOf(self._shape_set_margin)
	case "_shape_get_margin":
		return reflect.ValueOf(self._shape_get_margin)
	case "_shape_get_type":
		return reflect.ValueOf(self._shape_get_type)
	case "_shape_get_data":
		return reflect.ValueOf(self._shape_get_data)
	case "_shape_get_custom_solver_bias":
		return reflect.ValueOf(self._shape_get_custom_solver_bias)
	case "_space_create":
		return reflect.ValueOf(self._space_create)
	case "_space_set_active":
		return reflect.ValueOf(self._space_set_active)
	case "_space_is_active":
		return reflect.ValueOf(self._space_is_active)
	case "_space_set_param":
		return reflect.ValueOf(self._space_set_param)
	case "_space_get_param":
		return reflect.ValueOf(self._space_get_param)
	case "_space_get_direct_state":
		return reflect.ValueOf(self._space_get_direct_state)
	case "_space_set_debug_contacts":
		return reflect.ValueOf(self._space_set_debug_contacts)
	case "_space_get_contacts":
		return reflect.ValueOf(self._space_get_contacts)
	case "_space_get_contact_count":
		return reflect.ValueOf(self._space_get_contact_count)
	case "_area_create":
		return reflect.ValueOf(self._area_create)
	case "_area_set_space":
		return reflect.ValueOf(self._area_set_space)
	case "_area_get_space":
		return reflect.ValueOf(self._area_get_space)
	case "_area_add_shape":
		return reflect.ValueOf(self._area_add_shape)
	case "_area_set_shape":
		return reflect.ValueOf(self._area_set_shape)
	case "_area_set_shape_transform":
		return reflect.ValueOf(self._area_set_shape_transform)
	case "_area_set_shape_disabled":
		return reflect.ValueOf(self._area_set_shape_disabled)
	case "_area_get_shape_count":
		return reflect.ValueOf(self._area_get_shape_count)
	case "_area_get_shape":
		return reflect.ValueOf(self._area_get_shape)
	case "_area_get_shape_transform":
		return reflect.ValueOf(self._area_get_shape_transform)
	case "_area_remove_shape":
		return reflect.ValueOf(self._area_remove_shape)
	case "_area_clear_shapes":
		return reflect.ValueOf(self._area_clear_shapes)
	case "_area_attach_object_instance_id":
		return reflect.ValueOf(self._area_attach_object_instance_id)
	case "_area_get_object_instance_id":
		return reflect.ValueOf(self._area_get_object_instance_id)
	case "_area_set_param":
		return reflect.ValueOf(self._area_set_param)
	case "_area_set_transform":
		return reflect.ValueOf(self._area_set_transform)
	case "_area_get_param":
		return reflect.ValueOf(self._area_get_param)
	case "_area_get_transform":
		return reflect.ValueOf(self._area_get_transform)
	case "_area_set_collision_layer":
		return reflect.ValueOf(self._area_set_collision_layer)
	case "_area_get_collision_layer":
		return reflect.ValueOf(self._area_get_collision_layer)
	case "_area_set_collision_mask":
		return reflect.ValueOf(self._area_set_collision_mask)
	case "_area_get_collision_mask":
		return reflect.ValueOf(self._area_get_collision_mask)
	case "_area_set_monitorable":
		return reflect.ValueOf(self._area_set_monitorable)
	case "_area_set_ray_pickable":
		return reflect.ValueOf(self._area_set_ray_pickable)
	case "_area_set_monitor_callback":
		return reflect.ValueOf(self._area_set_monitor_callback)
	case "_area_set_area_monitor_callback":
		return reflect.ValueOf(self._area_set_area_monitor_callback)
	case "_body_create":
		return reflect.ValueOf(self._body_create)
	case "_body_set_space":
		return reflect.ValueOf(self._body_set_space)
	case "_body_get_space":
		return reflect.ValueOf(self._body_get_space)
	case "_body_set_mode":
		return reflect.ValueOf(self._body_set_mode)
	case "_body_get_mode":
		return reflect.ValueOf(self._body_get_mode)
	case "_body_add_shape":
		return reflect.ValueOf(self._body_add_shape)
	case "_body_set_shape":
		return reflect.ValueOf(self._body_set_shape)
	case "_body_set_shape_transform":
		return reflect.ValueOf(self._body_set_shape_transform)
	case "_body_set_shape_disabled":
		return reflect.ValueOf(self._body_set_shape_disabled)
	case "_body_get_shape_count":
		return reflect.ValueOf(self._body_get_shape_count)
	case "_body_get_shape":
		return reflect.ValueOf(self._body_get_shape)
	case "_body_get_shape_transform":
		return reflect.ValueOf(self._body_get_shape_transform)
	case "_body_remove_shape":
		return reflect.ValueOf(self._body_remove_shape)
	case "_body_clear_shapes":
		return reflect.ValueOf(self._body_clear_shapes)
	case "_body_attach_object_instance_id":
		return reflect.ValueOf(self._body_attach_object_instance_id)
	case "_body_get_object_instance_id":
		return reflect.ValueOf(self._body_get_object_instance_id)
	case "_body_set_enable_continuous_collision_detection":
		return reflect.ValueOf(self._body_set_enable_continuous_collision_detection)
	case "_body_is_continuous_collision_detection_enabled":
		return reflect.ValueOf(self._body_is_continuous_collision_detection_enabled)
	case "_body_set_collision_layer":
		return reflect.ValueOf(self._body_set_collision_layer)
	case "_body_get_collision_layer":
		return reflect.ValueOf(self._body_get_collision_layer)
	case "_body_set_collision_mask":
		return reflect.ValueOf(self._body_set_collision_mask)
	case "_body_get_collision_mask":
		return reflect.ValueOf(self._body_get_collision_mask)
	case "_body_set_collision_priority":
		return reflect.ValueOf(self._body_set_collision_priority)
	case "_body_get_collision_priority":
		return reflect.ValueOf(self._body_get_collision_priority)
	case "_body_set_user_flags":
		return reflect.ValueOf(self._body_set_user_flags)
	case "_body_get_user_flags":
		return reflect.ValueOf(self._body_get_user_flags)
	case "_body_set_param":
		return reflect.ValueOf(self._body_set_param)
	case "_body_get_param":
		return reflect.ValueOf(self._body_get_param)
	case "_body_reset_mass_properties":
		return reflect.ValueOf(self._body_reset_mass_properties)
	case "_body_set_state":
		return reflect.ValueOf(self._body_set_state)
	case "_body_get_state":
		return reflect.ValueOf(self._body_get_state)
	case "_body_apply_central_impulse":
		return reflect.ValueOf(self._body_apply_central_impulse)
	case "_body_apply_impulse":
		return reflect.ValueOf(self._body_apply_impulse)
	case "_body_apply_torque_impulse":
		return reflect.ValueOf(self._body_apply_torque_impulse)
	case "_body_apply_central_force":
		return reflect.ValueOf(self._body_apply_central_force)
	case "_body_apply_force":
		return reflect.ValueOf(self._body_apply_force)
	case "_body_apply_torque":
		return reflect.ValueOf(self._body_apply_torque)
	case "_body_add_constant_central_force":
		return reflect.ValueOf(self._body_add_constant_central_force)
	case "_body_add_constant_force":
		return reflect.ValueOf(self._body_add_constant_force)
	case "_body_add_constant_torque":
		return reflect.ValueOf(self._body_add_constant_torque)
	case "_body_set_constant_force":
		return reflect.ValueOf(self._body_set_constant_force)
	case "_body_get_constant_force":
		return reflect.ValueOf(self._body_get_constant_force)
	case "_body_set_constant_torque":
		return reflect.ValueOf(self._body_set_constant_torque)
	case "_body_get_constant_torque":
		return reflect.ValueOf(self._body_get_constant_torque)
	case "_body_set_axis_velocity":
		return reflect.ValueOf(self._body_set_axis_velocity)
	case "_body_set_axis_lock":
		return reflect.ValueOf(self._body_set_axis_lock)
	case "_body_is_axis_locked":
		return reflect.ValueOf(self._body_is_axis_locked)
	case "_body_add_collision_exception":
		return reflect.ValueOf(self._body_add_collision_exception)
	case "_body_remove_collision_exception":
		return reflect.ValueOf(self._body_remove_collision_exception)
	case "_body_get_collision_exceptions":
		return reflect.ValueOf(self._body_get_collision_exceptions)
	case "_body_set_max_contacts_reported":
		return reflect.ValueOf(self._body_set_max_contacts_reported)
	case "_body_get_max_contacts_reported":
		return reflect.ValueOf(self._body_get_max_contacts_reported)
	case "_body_set_contacts_reported_depth_threshold":
		return reflect.ValueOf(self._body_set_contacts_reported_depth_threshold)
	case "_body_get_contacts_reported_depth_threshold":
		return reflect.ValueOf(self._body_get_contacts_reported_depth_threshold)
	case "_body_set_omit_force_integration":
		return reflect.ValueOf(self._body_set_omit_force_integration)
	case "_body_is_omitting_force_integration":
		return reflect.ValueOf(self._body_is_omitting_force_integration)
	case "_body_set_state_sync_callback":
		return reflect.ValueOf(self._body_set_state_sync_callback)
	case "_body_set_force_integration_callback":
		return reflect.ValueOf(self._body_set_force_integration_callback)
	case "_body_set_ray_pickable":
		return reflect.ValueOf(self._body_set_ray_pickable)
	case "_body_test_motion":
		return reflect.ValueOf(self._body_test_motion)
	case "_body_get_direct_state":
		return reflect.ValueOf(self._body_get_direct_state)
	case "_soft_body_create":
		return reflect.ValueOf(self._soft_body_create)
	case "_soft_body_update_rendering_server":
		return reflect.ValueOf(self._soft_body_update_rendering_server)
	case "_soft_body_set_space":
		return reflect.ValueOf(self._soft_body_set_space)
	case "_soft_body_get_space":
		return reflect.ValueOf(self._soft_body_get_space)
	case "_soft_body_set_ray_pickable":
		return reflect.ValueOf(self._soft_body_set_ray_pickable)
	case "_soft_body_set_collision_layer":
		return reflect.ValueOf(self._soft_body_set_collision_layer)
	case "_soft_body_get_collision_layer":
		return reflect.ValueOf(self._soft_body_get_collision_layer)
	case "_soft_body_set_collision_mask":
		return reflect.ValueOf(self._soft_body_set_collision_mask)
	case "_soft_body_get_collision_mask":
		return reflect.ValueOf(self._soft_body_get_collision_mask)
	case "_soft_body_add_collision_exception":
		return reflect.ValueOf(self._soft_body_add_collision_exception)
	case "_soft_body_remove_collision_exception":
		return reflect.ValueOf(self._soft_body_remove_collision_exception)
	case "_soft_body_get_collision_exceptions":
		return reflect.ValueOf(self._soft_body_get_collision_exceptions)
	case "_soft_body_set_state":
		return reflect.ValueOf(self._soft_body_set_state)
	case "_soft_body_get_state":
		return reflect.ValueOf(self._soft_body_get_state)
	case "_soft_body_set_transform":
		return reflect.ValueOf(self._soft_body_set_transform)
	case "_soft_body_set_simulation_precision":
		return reflect.ValueOf(self._soft_body_set_simulation_precision)
	case "_soft_body_get_simulation_precision":
		return reflect.ValueOf(self._soft_body_get_simulation_precision)
	case "_soft_body_set_total_mass":
		return reflect.ValueOf(self._soft_body_set_total_mass)
	case "_soft_body_get_total_mass":
		return reflect.ValueOf(self._soft_body_get_total_mass)
	case "_soft_body_set_linear_stiffness":
		return reflect.ValueOf(self._soft_body_set_linear_stiffness)
	case "_soft_body_get_linear_stiffness":
		return reflect.ValueOf(self._soft_body_get_linear_stiffness)
	case "_soft_body_set_pressure_coefficient":
		return reflect.ValueOf(self._soft_body_set_pressure_coefficient)
	case "_soft_body_get_pressure_coefficient":
		return reflect.ValueOf(self._soft_body_get_pressure_coefficient)
	case "_soft_body_set_damping_coefficient":
		return reflect.ValueOf(self._soft_body_set_damping_coefficient)
	case "_soft_body_get_damping_coefficient":
		return reflect.ValueOf(self._soft_body_get_damping_coefficient)
	case "_soft_body_set_drag_coefficient":
		return reflect.ValueOf(self._soft_body_set_drag_coefficient)
	case "_soft_body_get_drag_coefficient":
		return reflect.ValueOf(self._soft_body_get_drag_coefficient)
	case "_soft_body_set_mesh":
		return reflect.ValueOf(self._soft_body_set_mesh)
	case "_soft_body_get_bounds":
		return reflect.ValueOf(self._soft_body_get_bounds)
	case "_soft_body_move_point":
		return reflect.ValueOf(self._soft_body_move_point)
	case "_soft_body_get_point_global_position":
		return reflect.ValueOf(self._soft_body_get_point_global_position)
	case "_soft_body_remove_all_pinned_points":
		return reflect.ValueOf(self._soft_body_remove_all_pinned_points)
	case "_soft_body_pin_point":
		return reflect.ValueOf(self._soft_body_pin_point)
	case "_soft_body_is_point_pinned":
		return reflect.ValueOf(self._soft_body_is_point_pinned)
	case "_joint_create":
		return reflect.ValueOf(self._joint_create)
	case "_joint_clear":
		return reflect.ValueOf(self._joint_clear)
	case "_joint_make_pin":
		return reflect.ValueOf(self._joint_make_pin)
	case "_pin_joint_set_param":
		return reflect.ValueOf(self._pin_joint_set_param)
	case "_pin_joint_get_param":
		return reflect.ValueOf(self._pin_joint_get_param)
	case "_pin_joint_set_local_a":
		return reflect.ValueOf(self._pin_joint_set_local_a)
	case "_pin_joint_get_local_a":
		return reflect.ValueOf(self._pin_joint_get_local_a)
	case "_pin_joint_set_local_b":
		return reflect.ValueOf(self._pin_joint_set_local_b)
	case "_pin_joint_get_local_b":
		return reflect.ValueOf(self._pin_joint_get_local_b)
	case "_joint_make_hinge":
		return reflect.ValueOf(self._joint_make_hinge)
	case "_joint_make_hinge_simple":
		return reflect.ValueOf(self._joint_make_hinge_simple)
	case "_hinge_joint_set_param":
		return reflect.ValueOf(self._hinge_joint_set_param)
	case "_hinge_joint_get_param":
		return reflect.ValueOf(self._hinge_joint_get_param)
	case "_hinge_joint_set_flag":
		return reflect.ValueOf(self._hinge_joint_set_flag)
	case "_hinge_joint_get_flag":
		return reflect.ValueOf(self._hinge_joint_get_flag)
	case "_joint_make_slider":
		return reflect.ValueOf(self._joint_make_slider)
	case "_slider_joint_set_param":
		return reflect.ValueOf(self._slider_joint_set_param)
	case "_slider_joint_get_param":
		return reflect.ValueOf(self._slider_joint_get_param)
	case "_joint_make_cone_twist":
		return reflect.ValueOf(self._joint_make_cone_twist)
	case "_cone_twist_joint_set_param":
		return reflect.ValueOf(self._cone_twist_joint_set_param)
	case "_cone_twist_joint_get_param":
		return reflect.ValueOf(self._cone_twist_joint_get_param)
	case "_joint_make_generic_6dof":
		return reflect.ValueOf(self._joint_make_generic_6dof)
	case "_generic_6dof_joint_set_param":
		return reflect.ValueOf(self._generic_6dof_joint_set_param)
	case "_generic_6dof_joint_get_param":
		return reflect.ValueOf(self._generic_6dof_joint_get_param)
	case "_generic_6dof_joint_set_flag":
		return reflect.ValueOf(self._generic_6dof_joint_set_flag)
	case "_generic_6dof_joint_get_flag":
		return reflect.ValueOf(self._generic_6dof_joint_get_flag)
	case "_joint_get_type":
		return reflect.ValueOf(self._joint_get_type)
	case "_joint_set_solver_priority":
		return reflect.ValueOf(self._joint_set_solver_priority)
	case "_joint_get_solver_priority":
		return reflect.ValueOf(self._joint_get_solver_priority)
	case "_joint_disable_collisions_between_bodies":
		return reflect.ValueOf(self._joint_disable_collisions_between_bodies)
	case "_joint_is_disabled_collisions_between_bodies":
		return reflect.ValueOf(self._joint_is_disabled_collisions_between_bodies)
	case "_free_rid":
		return reflect.ValueOf(self._free_rid)
	case "_set_active":
		return reflect.ValueOf(self._set_active)
	case "_init":
		return reflect.ValueOf(self._init)
	case "_step":
		return reflect.ValueOf(self._step)
	case "_sync":
		return reflect.ValueOf(self._sync)
	case "_flush_queries":
		return reflect.ValueOf(self._flush_queries)
	case "_end_sync":
		return reflect.ValueOf(self._end_sync)
	case "_finish":
		return reflect.ValueOf(self._finish)
	case "_is_flushing_queries":
		return reflect.ValueOf(self._is_flushing_queries)
	case "_get_process_info":
		return reflect.ValueOf(self._get_process_info)
	default:
		return reflect.Value{}
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_world_boundary_shape_create":
		return reflect.ValueOf(self._world_boundary_shape_create)
	case "_separation_ray_shape_create":
		return reflect.ValueOf(self._separation_ray_shape_create)
	case "_sphere_shape_create":
		return reflect.ValueOf(self._sphere_shape_create)
	case "_box_shape_create":
		return reflect.ValueOf(self._box_shape_create)
	case "_capsule_shape_create":
		return reflect.ValueOf(self._capsule_shape_create)
	case "_cylinder_shape_create":
		return reflect.ValueOf(self._cylinder_shape_create)
	case "_convex_polygon_shape_create":
		return reflect.ValueOf(self._convex_polygon_shape_create)
	case "_concave_polygon_shape_create":
		return reflect.ValueOf(self._concave_polygon_shape_create)
	case "_heightmap_shape_create":
		return reflect.ValueOf(self._heightmap_shape_create)
	case "_custom_shape_create":
		return reflect.ValueOf(self._custom_shape_create)
	case "_shape_set_data":
		return reflect.ValueOf(self._shape_set_data)
	case "_shape_set_custom_solver_bias":
		return reflect.ValueOf(self._shape_set_custom_solver_bias)
	case "_shape_set_margin":
		return reflect.ValueOf(self._shape_set_margin)
	case "_shape_get_margin":
		return reflect.ValueOf(self._shape_get_margin)
	case "_shape_get_type":
		return reflect.ValueOf(self._shape_get_type)
	case "_shape_get_data":
		return reflect.ValueOf(self._shape_get_data)
	case "_shape_get_custom_solver_bias":
		return reflect.ValueOf(self._shape_get_custom_solver_bias)
	case "_space_create":
		return reflect.ValueOf(self._space_create)
	case "_space_set_active":
		return reflect.ValueOf(self._space_set_active)
	case "_space_is_active":
		return reflect.ValueOf(self._space_is_active)
	case "_space_set_param":
		return reflect.ValueOf(self._space_set_param)
	case "_space_get_param":
		return reflect.ValueOf(self._space_get_param)
	case "_space_get_direct_state":
		return reflect.ValueOf(self._space_get_direct_state)
	case "_space_set_debug_contacts":
		return reflect.ValueOf(self._space_set_debug_contacts)
	case "_space_get_contacts":
		return reflect.ValueOf(self._space_get_contacts)
	case "_space_get_contact_count":
		return reflect.ValueOf(self._space_get_contact_count)
	case "_area_create":
		return reflect.ValueOf(self._area_create)
	case "_area_set_space":
		return reflect.ValueOf(self._area_set_space)
	case "_area_get_space":
		return reflect.ValueOf(self._area_get_space)
	case "_area_add_shape":
		return reflect.ValueOf(self._area_add_shape)
	case "_area_set_shape":
		return reflect.ValueOf(self._area_set_shape)
	case "_area_set_shape_transform":
		return reflect.ValueOf(self._area_set_shape_transform)
	case "_area_set_shape_disabled":
		return reflect.ValueOf(self._area_set_shape_disabled)
	case "_area_get_shape_count":
		return reflect.ValueOf(self._area_get_shape_count)
	case "_area_get_shape":
		return reflect.ValueOf(self._area_get_shape)
	case "_area_get_shape_transform":
		return reflect.ValueOf(self._area_get_shape_transform)
	case "_area_remove_shape":
		return reflect.ValueOf(self._area_remove_shape)
	case "_area_clear_shapes":
		return reflect.ValueOf(self._area_clear_shapes)
	case "_area_attach_object_instance_id":
		return reflect.ValueOf(self._area_attach_object_instance_id)
	case "_area_get_object_instance_id":
		return reflect.ValueOf(self._area_get_object_instance_id)
	case "_area_set_param":
		return reflect.ValueOf(self._area_set_param)
	case "_area_set_transform":
		return reflect.ValueOf(self._area_set_transform)
	case "_area_get_param":
		return reflect.ValueOf(self._area_get_param)
	case "_area_get_transform":
		return reflect.ValueOf(self._area_get_transform)
	case "_area_set_collision_layer":
		return reflect.ValueOf(self._area_set_collision_layer)
	case "_area_get_collision_layer":
		return reflect.ValueOf(self._area_get_collision_layer)
	case "_area_set_collision_mask":
		return reflect.ValueOf(self._area_set_collision_mask)
	case "_area_get_collision_mask":
		return reflect.ValueOf(self._area_get_collision_mask)
	case "_area_set_monitorable":
		return reflect.ValueOf(self._area_set_monitorable)
	case "_area_set_ray_pickable":
		return reflect.ValueOf(self._area_set_ray_pickable)
	case "_area_set_monitor_callback":
		return reflect.ValueOf(self._area_set_monitor_callback)
	case "_area_set_area_monitor_callback":
		return reflect.ValueOf(self._area_set_area_monitor_callback)
	case "_body_create":
		return reflect.ValueOf(self._body_create)
	case "_body_set_space":
		return reflect.ValueOf(self._body_set_space)
	case "_body_get_space":
		return reflect.ValueOf(self._body_get_space)
	case "_body_set_mode":
		return reflect.ValueOf(self._body_set_mode)
	case "_body_get_mode":
		return reflect.ValueOf(self._body_get_mode)
	case "_body_add_shape":
		return reflect.ValueOf(self._body_add_shape)
	case "_body_set_shape":
		return reflect.ValueOf(self._body_set_shape)
	case "_body_set_shape_transform":
		return reflect.ValueOf(self._body_set_shape_transform)
	case "_body_set_shape_disabled":
		return reflect.ValueOf(self._body_set_shape_disabled)
	case "_body_get_shape_count":
		return reflect.ValueOf(self._body_get_shape_count)
	case "_body_get_shape":
		return reflect.ValueOf(self._body_get_shape)
	case "_body_get_shape_transform":
		return reflect.ValueOf(self._body_get_shape_transform)
	case "_body_remove_shape":
		return reflect.ValueOf(self._body_remove_shape)
	case "_body_clear_shapes":
		return reflect.ValueOf(self._body_clear_shapes)
	case "_body_attach_object_instance_id":
		return reflect.ValueOf(self._body_attach_object_instance_id)
	case "_body_get_object_instance_id":
		return reflect.ValueOf(self._body_get_object_instance_id)
	case "_body_set_enable_continuous_collision_detection":
		return reflect.ValueOf(self._body_set_enable_continuous_collision_detection)
	case "_body_is_continuous_collision_detection_enabled":
		return reflect.ValueOf(self._body_is_continuous_collision_detection_enabled)
	case "_body_set_collision_layer":
		return reflect.ValueOf(self._body_set_collision_layer)
	case "_body_get_collision_layer":
		return reflect.ValueOf(self._body_get_collision_layer)
	case "_body_set_collision_mask":
		return reflect.ValueOf(self._body_set_collision_mask)
	case "_body_get_collision_mask":
		return reflect.ValueOf(self._body_get_collision_mask)
	case "_body_set_collision_priority":
		return reflect.ValueOf(self._body_set_collision_priority)
	case "_body_get_collision_priority":
		return reflect.ValueOf(self._body_get_collision_priority)
	case "_body_set_user_flags":
		return reflect.ValueOf(self._body_set_user_flags)
	case "_body_get_user_flags":
		return reflect.ValueOf(self._body_get_user_flags)
	case "_body_set_param":
		return reflect.ValueOf(self._body_set_param)
	case "_body_get_param":
		return reflect.ValueOf(self._body_get_param)
	case "_body_reset_mass_properties":
		return reflect.ValueOf(self._body_reset_mass_properties)
	case "_body_set_state":
		return reflect.ValueOf(self._body_set_state)
	case "_body_get_state":
		return reflect.ValueOf(self._body_get_state)
	case "_body_apply_central_impulse":
		return reflect.ValueOf(self._body_apply_central_impulse)
	case "_body_apply_impulse":
		return reflect.ValueOf(self._body_apply_impulse)
	case "_body_apply_torque_impulse":
		return reflect.ValueOf(self._body_apply_torque_impulse)
	case "_body_apply_central_force":
		return reflect.ValueOf(self._body_apply_central_force)
	case "_body_apply_force":
		return reflect.ValueOf(self._body_apply_force)
	case "_body_apply_torque":
		return reflect.ValueOf(self._body_apply_torque)
	case "_body_add_constant_central_force":
		return reflect.ValueOf(self._body_add_constant_central_force)
	case "_body_add_constant_force":
		return reflect.ValueOf(self._body_add_constant_force)
	case "_body_add_constant_torque":
		return reflect.ValueOf(self._body_add_constant_torque)
	case "_body_set_constant_force":
		return reflect.ValueOf(self._body_set_constant_force)
	case "_body_get_constant_force":
		return reflect.ValueOf(self._body_get_constant_force)
	case "_body_set_constant_torque":
		return reflect.ValueOf(self._body_set_constant_torque)
	case "_body_get_constant_torque":
		return reflect.ValueOf(self._body_get_constant_torque)
	case "_body_set_axis_velocity":
		return reflect.ValueOf(self._body_set_axis_velocity)
	case "_body_set_axis_lock":
		return reflect.ValueOf(self._body_set_axis_lock)
	case "_body_is_axis_locked":
		return reflect.ValueOf(self._body_is_axis_locked)
	case "_body_add_collision_exception":
		return reflect.ValueOf(self._body_add_collision_exception)
	case "_body_remove_collision_exception":
		return reflect.ValueOf(self._body_remove_collision_exception)
	case "_body_get_collision_exceptions":
		return reflect.ValueOf(self._body_get_collision_exceptions)
	case "_body_set_max_contacts_reported":
		return reflect.ValueOf(self._body_set_max_contacts_reported)
	case "_body_get_max_contacts_reported":
		return reflect.ValueOf(self._body_get_max_contacts_reported)
	case "_body_set_contacts_reported_depth_threshold":
		return reflect.ValueOf(self._body_set_contacts_reported_depth_threshold)
	case "_body_get_contacts_reported_depth_threshold":
		return reflect.ValueOf(self._body_get_contacts_reported_depth_threshold)
	case "_body_set_omit_force_integration":
		return reflect.ValueOf(self._body_set_omit_force_integration)
	case "_body_is_omitting_force_integration":
		return reflect.ValueOf(self._body_is_omitting_force_integration)
	case "_body_set_state_sync_callback":
		return reflect.ValueOf(self._body_set_state_sync_callback)
	case "_body_set_force_integration_callback":
		return reflect.ValueOf(self._body_set_force_integration_callback)
	case "_body_set_ray_pickable":
		return reflect.ValueOf(self._body_set_ray_pickable)
	case "_body_test_motion":
		return reflect.ValueOf(self._body_test_motion)
	case "_body_get_direct_state":
		return reflect.ValueOf(self._body_get_direct_state)
	case "_soft_body_create":
		return reflect.ValueOf(self._soft_body_create)
	case "_soft_body_update_rendering_server":
		return reflect.ValueOf(self._soft_body_update_rendering_server)
	case "_soft_body_set_space":
		return reflect.ValueOf(self._soft_body_set_space)
	case "_soft_body_get_space":
		return reflect.ValueOf(self._soft_body_get_space)
	case "_soft_body_set_ray_pickable":
		return reflect.ValueOf(self._soft_body_set_ray_pickable)
	case "_soft_body_set_collision_layer":
		return reflect.ValueOf(self._soft_body_set_collision_layer)
	case "_soft_body_get_collision_layer":
		return reflect.ValueOf(self._soft_body_get_collision_layer)
	case "_soft_body_set_collision_mask":
		return reflect.ValueOf(self._soft_body_set_collision_mask)
	case "_soft_body_get_collision_mask":
		return reflect.ValueOf(self._soft_body_get_collision_mask)
	case "_soft_body_add_collision_exception":
		return reflect.ValueOf(self._soft_body_add_collision_exception)
	case "_soft_body_remove_collision_exception":
		return reflect.ValueOf(self._soft_body_remove_collision_exception)
	case "_soft_body_get_collision_exceptions":
		return reflect.ValueOf(self._soft_body_get_collision_exceptions)
	case "_soft_body_set_state":
		return reflect.ValueOf(self._soft_body_set_state)
	case "_soft_body_get_state":
		return reflect.ValueOf(self._soft_body_get_state)
	case "_soft_body_set_transform":
		return reflect.ValueOf(self._soft_body_set_transform)
	case "_soft_body_set_simulation_precision":
		return reflect.ValueOf(self._soft_body_set_simulation_precision)
	case "_soft_body_get_simulation_precision":
		return reflect.ValueOf(self._soft_body_get_simulation_precision)
	case "_soft_body_set_total_mass":
		return reflect.ValueOf(self._soft_body_set_total_mass)
	case "_soft_body_get_total_mass":
		return reflect.ValueOf(self._soft_body_get_total_mass)
	case "_soft_body_set_linear_stiffness":
		return reflect.ValueOf(self._soft_body_set_linear_stiffness)
	case "_soft_body_get_linear_stiffness":
		return reflect.ValueOf(self._soft_body_get_linear_stiffness)
	case "_soft_body_set_pressure_coefficient":
		return reflect.ValueOf(self._soft_body_set_pressure_coefficient)
	case "_soft_body_get_pressure_coefficient":
		return reflect.ValueOf(self._soft_body_get_pressure_coefficient)
	case "_soft_body_set_damping_coefficient":
		return reflect.ValueOf(self._soft_body_set_damping_coefficient)
	case "_soft_body_get_damping_coefficient":
		return reflect.ValueOf(self._soft_body_get_damping_coefficient)
	case "_soft_body_set_drag_coefficient":
		return reflect.ValueOf(self._soft_body_set_drag_coefficient)
	case "_soft_body_get_drag_coefficient":
		return reflect.ValueOf(self._soft_body_get_drag_coefficient)
	case "_soft_body_set_mesh":
		return reflect.ValueOf(self._soft_body_set_mesh)
	case "_soft_body_get_bounds":
		return reflect.ValueOf(self._soft_body_get_bounds)
	case "_soft_body_move_point":
		return reflect.ValueOf(self._soft_body_move_point)
	case "_soft_body_get_point_global_position":
		return reflect.ValueOf(self._soft_body_get_point_global_position)
	case "_soft_body_remove_all_pinned_points":
		return reflect.ValueOf(self._soft_body_remove_all_pinned_points)
	case "_soft_body_pin_point":
		return reflect.ValueOf(self._soft_body_pin_point)
	case "_soft_body_is_point_pinned":
		return reflect.ValueOf(self._soft_body_is_point_pinned)
	case "_joint_create":
		return reflect.ValueOf(self._joint_create)
	case "_joint_clear":
		return reflect.ValueOf(self._joint_clear)
	case "_joint_make_pin":
		return reflect.ValueOf(self._joint_make_pin)
	case "_pin_joint_set_param":
		return reflect.ValueOf(self._pin_joint_set_param)
	case "_pin_joint_get_param":
		return reflect.ValueOf(self._pin_joint_get_param)
	case "_pin_joint_set_local_a":
		return reflect.ValueOf(self._pin_joint_set_local_a)
	case "_pin_joint_get_local_a":
		return reflect.ValueOf(self._pin_joint_get_local_a)
	case "_pin_joint_set_local_b":
		return reflect.ValueOf(self._pin_joint_set_local_b)
	case "_pin_joint_get_local_b":
		return reflect.ValueOf(self._pin_joint_get_local_b)
	case "_joint_make_hinge":
		return reflect.ValueOf(self._joint_make_hinge)
	case "_joint_make_hinge_simple":
		return reflect.ValueOf(self._joint_make_hinge_simple)
	case "_hinge_joint_set_param":
		return reflect.ValueOf(self._hinge_joint_set_param)
	case "_hinge_joint_get_param":
		return reflect.ValueOf(self._hinge_joint_get_param)
	case "_hinge_joint_set_flag":
		return reflect.ValueOf(self._hinge_joint_set_flag)
	case "_hinge_joint_get_flag":
		return reflect.ValueOf(self._hinge_joint_get_flag)
	case "_joint_make_slider":
		return reflect.ValueOf(self._joint_make_slider)
	case "_slider_joint_set_param":
		return reflect.ValueOf(self._slider_joint_set_param)
	case "_slider_joint_get_param":
		return reflect.ValueOf(self._slider_joint_get_param)
	case "_joint_make_cone_twist":
		return reflect.ValueOf(self._joint_make_cone_twist)
	case "_cone_twist_joint_set_param":
		return reflect.ValueOf(self._cone_twist_joint_set_param)
	case "_cone_twist_joint_get_param":
		return reflect.ValueOf(self._cone_twist_joint_get_param)
	case "_joint_make_generic_6dof":
		return reflect.ValueOf(self._joint_make_generic_6dof)
	case "_generic_6dof_joint_set_param":
		return reflect.ValueOf(self._generic_6dof_joint_set_param)
	case "_generic_6dof_joint_get_param":
		return reflect.ValueOf(self._generic_6dof_joint_get_param)
	case "_generic_6dof_joint_set_flag":
		return reflect.ValueOf(self._generic_6dof_joint_set_flag)
	case "_generic_6dof_joint_get_flag":
		return reflect.ValueOf(self._generic_6dof_joint_get_flag)
	case "_joint_get_type":
		return reflect.ValueOf(self._joint_get_type)
	case "_joint_set_solver_priority":
		return reflect.ValueOf(self._joint_set_solver_priority)
	case "_joint_get_solver_priority":
		return reflect.ValueOf(self._joint_get_solver_priority)
	case "_joint_disable_collisions_between_bodies":
		return reflect.ValueOf(self._joint_disable_collisions_between_bodies)
	case "_joint_is_disabled_collisions_between_bodies":
		return reflect.ValueOf(self._joint_is_disabled_collisions_between_bodies)
	case "_free_rid":
		return reflect.ValueOf(self._free_rid)
	case "_set_active":
		return reflect.ValueOf(self._set_active)
	case "_init":
		return reflect.ValueOf(self._init)
	case "_step":
		return reflect.ValueOf(self._step)
	case "_sync":
		return reflect.ValueOf(self._sync)
	case "_flush_queries":
		return reflect.ValueOf(self._flush_queries)
	case "_end_sync":
		return reflect.ValueOf(self._end_sync)
	case "_finish":
		return reflect.ValueOf(self._finish)
	case "_is_flushing_queries":
		return reflect.ValueOf(self._is_flushing_queries)
	case "_get_process_info":
		return reflect.ValueOf(self._get_process_info)
	default:
		return reflect.Value{}
	}
}
func init() {
	gdclass.Register("PhysicsServer3DExtension", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsServer3DExtension{*(*gdclass.PhysicsServer3DExtension)(unsafe.Pointer(&ptr))}
	})
}
