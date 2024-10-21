package PhysicsServer2DExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PhysicsServer2D"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class extends [PhysicsServer2D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsServer2D].
	// PhysicsServer2DExtension methods that can be overridden by a [Class] that extends it.
	type PhysicsServer2DExtension interface {
		//Overridable version of [method PhysicsServer2D.world_boundary_shape_create].
		WorldBoundaryShapeCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.separation_ray_shape_create].
		SeparationRayShapeCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.segment_shape_create].
		SegmentShapeCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.circle_shape_create].
		CircleShapeCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.rectangle_shape_create].
		RectangleShapeCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.capsule_shape_create].
		CapsuleShapeCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.convex_polygon_shape_create].
		ConvexPolygonShapeCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.concave_polygon_shape_create].
		ConcavePolygonShapeCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.shape_set_data].
		ShapeSetData(shape gd.RID, data gd.Variant) 
		//Should set the custom solver bias for the given [param shape]. It defines how much bodies are forced to separate on contact.
		//Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
		ShapeSetCustomSolverBias(shape gd.RID, bias gd.Float) 
		//Overridable version of [method PhysicsServer2D.shape_get_type].
		ShapeGetType(shape gd.RID) classdb.PhysicsServer2DShapeType
		//Overridable version of [method PhysicsServer2D.shape_get_data].
		ShapeGetData(shape gd.RID) gd.Variant
		//Should return the custom solver bias of the given [param shape], which defines how much bodies are forced to separate on contact when this shape is involved.
		//Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
		ShapeGetCustomSolverBias(shape gd.RID) gd.Float
		//Given two shapes and their parameters, should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
		//Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
		ShapeCollide(shape_A gd.RID, xform_A gd.Transform2D, motion_A gd.Vector2, shape_B gd.RID, xform_B gd.Transform2D, motion_B gd.Vector2, results unsafe.Pointer, result_max gd.Int, result_count *int32) bool
		//Overridable version of [method PhysicsServer2D.space_create].
		SpaceCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.space_set_active].
		SpaceSetActive(space gd.RID, active bool) 
		//Overridable version of [method PhysicsServer2D.space_is_active].
		SpaceIsActive(space gd.RID) bool
		//Overridable version of [method PhysicsServer2D.space_set_param].
		SpaceSetParam(space gd.RID, param classdb.PhysicsServer2DSpaceParameter, value gd.Float) 
		//Overridable version of [method PhysicsServer2D.space_get_param].
		SpaceGetParam(space gd.RID, param classdb.PhysicsServer2DSpaceParameter) gd.Float
		//Overridable version of [method PhysicsServer2D.space_get_direct_state].
		SpaceGetDirectState(space gd.RID) object.PhysicsDirectSpaceState2D
		//Used internally to allow the given [param space] to store contact points, up to [param max_contacts]. This is automatically set for the main [World2D]'s space when [member SceneTree.debug_collisions_hint] is [code]true[/code], or by checking "Visible Collision Shapes" in the editor. Only works in debug builds.
		//Overridable version of [PhysicsServer2D]'s internal [code]space_set_debug_contacts[/code] method.
		SpaceSetDebugContacts(space gd.RID, max_contacts gd.Int) 
		//Should return the positions of all contacts that have occurred during the last physics step in the given [param space]. See also [method _space_get_contact_count] and [method _space_set_debug_contacts].
		//Overridable version of [PhysicsServer2D]'s internal [code]space_get_contacts[/code] method.
		SpaceGetContacts(space gd.RID) gd.PackedVector2Array
		//Should return how many contacts have occurred during the last physics step in the given [param space]. See also [method _space_get_contacts] and [method _space_set_debug_contacts].
		//Overridable version of [PhysicsServer2D]'s internal [code]space_get_contact_count[/code] method.
		SpaceGetContactCount(space gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.area_create].
		AreaCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.area_set_space].
		AreaSetSpace(area gd.RID, space gd.RID) 
		//Overridable version of [method PhysicsServer2D.area_get_space].
		AreaGetSpace(area gd.RID) gd.RID
		//Overridable version of [method PhysicsServer2D.area_add_shape].
		AreaAddShape(area gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool) 
		//Overridable version of [method PhysicsServer2D.area_set_shape].
		AreaSetShape(area gd.RID, shape_idx gd.Int, shape gd.RID) 
		//Overridable version of [method PhysicsServer2D.area_set_shape_transform].
		AreaSetShapeTransform(area gd.RID, shape_idx gd.Int, transform gd.Transform2D) 
		//Overridable version of [method PhysicsServer2D.area_set_shape_disabled].
		AreaSetShapeDisabled(area gd.RID, shape_idx gd.Int, disabled bool) 
		//Overridable version of [method PhysicsServer2D.area_get_shape_count].
		AreaGetShapeCount(area gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.area_get_shape].
		AreaGetShape(area gd.RID, shape_idx gd.Int) gd.RID
		//Overridable version of [method PhysicsServer2D.area_get_shape_transform].
		AreaGetShapeTransform(area gd.RID, shape_idx gd.Int) gd.Transform2D
		//Overridable version of [method PhysicsServer2D.area_remove_shape].
		AreaRemoveShape(area gd.RID, shape_idx gd.Int) 
		//Overridable version of [method PhysicsServer2D.area_clear_shapes].
		AreaClearShapes(area gd.RID) 
		//Overridable version of [method PhysicsServer2D.area_attach_object_instance_id].
		AreaAttachObjectInstanceId(area gd.RID, id gd.Int) 
		//Overridable version of [method PhysicsServer2D.area_get_object_instance_id].
		AreaGetObjectInstanceId(area gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.area_attach_canvas_instance_id].
		AreaAttachCanvasInstanceId(area gd.RID, id gd.Int) 
		//Overridable version of [method PhysicsServer2D.area_get_canvas_instance_id].
		AreaGetCanvasInstanceId(area gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.area_set_param].
		AreaSetParam(area gd.RID, param classdb.PhysicsServer2DAreaParameter, value gd.Variant) 
		//Overridable version of [method PhysicsServer2D.area_set_transform].
		AreaSetTransform(area gd.RID, transform gd.Transform2D) 
		//Overridable version of [method PhysicsServer2D.area_get_param].
		AreaGetParam(area gd.RID, param classdb.PhysicsServer2DAreaParameter) gd.Variant
		//Overridable version of [method PhysicsServer2D.area_get_transform].
		AreaGetTransform(area gd.RID) gd.Transform2D
		//Overridable version of [method PhysicsServer2D.area_set_collision_layer].
		AreaSetCollisionLayer(area gd.RID, layer gd.Int) 
		//Overridable version of [method PhysicsServer2D.area_get_collision_layer].
		AreaGetCollisionLayer(area gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.area_set_collision_mask].
		AreaSetCollisionMask(area gd.RID, mask gd.Int) 
		//Overridable version of [method PhysicsServer2D.area_get_collision_mask].
		AreaGetCollisionMask(area gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.area_set_monitorable].
		AreaSetMonitorable(area gd.RID, monitorable bool) 
		//If set to [code]true[/code], allows the area with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
		//Overridable version of [PhysicsServer2D]'s internal [code]area_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
		AreaSetPickable(area gd.RID, pickable bool) 
		//Overridable version of [method PhysicsServer2D.area_set_monitor_callback].
		AreaSetMonitorCallback(area gd.RID, callback gd.Callable) 
		//Overridable version of [method PhysicsServer2D.area_set_area_monitor_callback].
		AreaSetAreaMonitorCallback(area gd.RID, callback gd.Callable) 
		//Overridable version of [method PhysicsServer2D.body_create].
		BodyCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.body_set_space].
		BodySetSpace(body gd.RID, space gd.RID) 
		//Overridable version of [method PhysicsServer2D.body_get_space].
		BodyGetSpace(body gd.RID) gd.RID
		//Overridable version of [method PhysicsServer2D.body_set_mode].
		BodySetMode(body gd.RID, mode classdb.PhysicsServer2DBodyMode) 
		//Overridable version of [method PhysicsServer2D.body_get_mode].
		BodyGetMode(body gd.RID) classdb.PhysicsServer2DBodyMode
		//Overridable version of [method PhysicsServer2D.body_add_shape].
		BodyAddShape(body gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool) 
		//Overridable version of [method PhysicsServer2D.body_set_shape].
		BodySetShape(body gd.RID, shape_idx gd.Int, shape gd.RID) 
		//Overridable version of [method PhysicsServer2D.body_set_shape_transform].
		BodySetShapeTransform(body gd.RID, shape_idx gd.Int, transform gd.Transform2D) 
		//Overridable version of [method PhysicsServer2D.body_get_shape_count].
		BodyGetShapeCount(body gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.body_get_shape].
		BodyGetShape(body gd.RID, shape_idx gd.Int) gd.RID
		//Overridable version of [method PhysicsServer2D.body_get_shape_transform].
		BodyGetShapeTransform(body gd.RID, shape_idx gd.Int) gd.Transform2D
		//Overridable version of [method PhysicsServer2D.body_set_shape_disabled].
		BodySetShapeDisabled(body gd.RID, shape_idx gd.Int, disabled bool) 
		//Overridable version of [method PhysicsServer2D.body_set_shape_as_one_way_collision].
		BodySetShapeAsOneWayCollision(body gd.RID, shape_idx gd.Int, enable bool, margin gd.Float) 
		//Overridable version of [method PhysicsServer2D.body_remove_shape].
		BodyRemoveShape(body gd.RID, shape_idx gd.Int) 
		//Overridable version of [method PhysicsServer2D.body_clear_shapes].
		BodyClearShapes(body gd.RID) 
		//Overridable version of [method PhysicsServer2D.body_attach_object_instance_id].
		BodyAttachObjectInstanceId(body gd.RID, id gd.Int) 
		//Overridable version of [method PhysicsServer2D.body_get_object_instance_id].
		BodyGetObjectInstanceId(body gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.body_attach_canvas_instance_id].
		BodyAttachCanvasInstanceId(body gd.RID, id gd.Int) 
		//Overridable version of [method PhysicsServer2D.body_get_canvas_instance_id].
		BodyGetCanvasInstanceId(body gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.body_set_continuous_collision_detection_mode].
		BodySetContinuousCollisionDetectionMode(body gd.RID, mode classdb.PhysicsServer2DCCDMode) 
		//Overridable version of [method PhysicsServer2D.body_get_continuous_collision_detection_mode].
		BodyGetContinuousCollisionDetectionMode(body gd.RID) classdb.PhysicsServer2DCCDMode
		//Overridable version of [method PhysicsServer2D.body_set_collision_layer].
		BodySetCollisionLayer(body gd.RID, layer gd.Int) 
		//Overridable version of [method PhysicsServer2D.body_get_collision_layer].
		BodyGetCollisionLayer(body gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.body_set_collision_mask].
		BodySetCollisionMask(body gd.RID, mask gd.Int) 
		//Overridable version of [method PhysicsServer2D.body_get_collision_mask].
		BodyGetCollisionMask(body gd.RID) gd.Int
		//Overridable version of [method PhysicsServer2D.body_set_collision_priority].
		BodySetCollisionPriority(body gd.RID, priority gd.Float) 
		//Overridable version of [method PhysicsServer2D.body_get_collision_priority].
		BodyGetCollisionPriority(body gd.RID) gd.Float
		//Overridable version of [method PhysicsServer2D.body_set_param].
		BodySetParam(body gd.RID, param classdb.PhysicsServer2DBodyParameter, value gd.Variant) 
		//Overridable version of [method PhysicsServer2D.body_get_param].
		BodyGetParam(body gd.RID, param classdb.PhysicsServer2DBodyParameter) gd.Variant
		//Overridable version of [method PhysicsServer2D.body_reset_mass_properties].
		BodyResetMassProperties(body gd.RID) 
		//Overridable version of [method PhysicsServer2D.body_set_state].
		BodySetState(body gd.RID, state classdb.PhysicsServer2DBodyState, value gd.Variant) 
		//Overridable version of [method PhysicsServer2D.body_get_state].
		BodyGetState(body gd.RID, state classdb.PhysicsServer2DBodyState) gd.Variant
		//Overridable version of [method PhysicsServer2D.body_apply_central_impulse].
		BodyApplyCentralImpulse(body gd.RID, impulse gd.Vector2) 
		//Overridable version of [method PhysicsServer2D.body_apply_torque_impulse].
		BodyApplyTorqueImpulse(body gd.RID, impulse gd.Float) 
		//Overridable version of [method PhysicsServer2D.body_apply_impulse].
		BodyApplyImpulse(body gd.RID, impulse gd.Vector2, position gd.Vector2) 
		//Overridable version of [method PhysicsServer2D.body_apply_central_force].
		BodyApplyCentralForce(body gd.RID, force gd.Vector2) 
		//Overridable version of [method PhysicsServer2D.body_apply_force].
		BodyApplyForce(body gd.RID, force gd.Vector2, position gd.Vector2) 
		//Overridable version of [method PhysicsServer2D.body_apply_torque].
		BodyApplyTorque(body gd.RID, torque gd.Float) 
		//Overridable version of [method PhysicsServer2D.body_add_constant_central_force].
		BodyAddConstantCentralForce(body gd.RID, force gd.Vector2) 
		//Overridable version of [method PhysicsServer2D.body_add_constant_force].
		BodyAddConstantForce(body gd.RID, force gd.Vector2, position gd.Vector2) 
		//Overridable version of [method PhysicsServer2D.body_add_constant_torque].
		BodyAddConstantTorque(body gd.RID, torque gd.Float) 
		//Overridable version of [method PhysicsServer2D.body_set_constant_force].
		BodySetConstantForce(body gd.RID, force gd.Vector2) 
		//Overridable version of [method PhysicsServer2D.body_get_constant_force].
		BodyGetConstantForce(body gd.RID) gd.Vector2
		//Overridable version of [method PhysicsServer2D.body_set_constant_torque].
		BodySetConstantTorque(body gd.RID, torque gd.Float) 
		//Overridable version of [method PhysicsServer2D.body_get_constant_torque].
		BodyGetConstantTorque(body gd.RID) gd.Float
		//Overridable version of [method PhysicsServer2D.body_set_axis_velocity].
		BodySetAxisVelocity(body gd.RID, axis_velocity gd.Vector2) 
		//Overridable version of [method PhysicsServer2D.body_add_collision_exception].
		BodyAddCollisionException(body gd.RID, excepted_body gd.RID) 
		//Overridable version of [method PhysicsServer2D.body_remove_collision_exception].
		BodyRemoveCollisionException(body gd.RID, excepted_body gd.RID) 
		//Returns the [RID]s of all bodies added as collision exceptions for the given [param body]. See also [method _body_add_collision_exception] and [method _body_remove_collision_exception].
		//Overridable version of [PhysicsServer2D]'s internal [code]body_get_collision_exceptions[/code] method. Corresponds to [method PhysicsBody2D.get_collision_exceptions].
		BodyGetCollisionExceptions(body gd.RID) gd.ArrayOf[gd.RID]
		//Overridable version of [method PhysicsServer2D.body_set_max_contacts_reported].
		BodySetMaxContactsReported(body gd.RID, amount gd.Int) 
		//Overridable version of [method PhysicsServer2D.body_get_max_contacts_reported].
		BodyGetMaxContactsReported(body gd.RID) gd.Int
		//Overridable version of [PhysicsServer2D]'s internal [code]body_set_contacts_reported_depth_threshold[/code] method.
		//[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
		BodySetContactsReportedDepthThreshold(body gd.RID, threshold gd.Float) 
		//Overridable version of [PhysicsServer2D]'s internal [code]body_get_contacts_reported_depth_threshold[/code] method.
		//[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
		BodyGetContactsReportedDepthThreshold(body gd.RID) gd.Float
		//Overridable version of [method PhysicsServer2D.body_set_omit_force_integration].
		BodySetOmitForceIntegration(body gd.RID, enable bool) 
		//Overridable version of [method PhysicsServer2D.body_is_omitting_force_integration].
		BodyIsOmittingForceIntegration(body gd.RID) bool
		//Assigns the [param body] to call the given [param callable] during the synchronization phase of the loop, before [method _step] is called. See also [method _sync].
		//Overridable version of [method PhysicsServer2D.body_set_state_sync_callback].
		BodySetStateSyncCallback(body gd.RID, callable gd.Callable) 
		//Overridable version of [method PhysicsServer2D.body_set_force_integration_callback].
		BodySetForceIntegrationCallback(body gd.RID, callable gd.Callable, userdata gd.Variant) 
		//Given a [param body], a [param shape], and their respective parameters, this method should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
		//Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
		BodyCollideShape(body gd.RID, body_shape gd.Int, shape gd.RID, shape_xform gd.Transform2D, motion gd.Vector2, results unsafe.Pointer, result_max gd.Int, result_count *int32) bool
		//If set to [code]true[/code], allows the body with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
		//Overridable version of [PhysicsServer2D]'s internal [code]body_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
		BodySetPickable(body gd.RID, pickable bool) 
		//Overridable version of [method PhysicsServer2D.body_get_direct_state].
		BodyGetDirectState(body gd.RID) object.PhysicsDirectBodyState2D
		//Overridable version of [method PhysicsServer2D.body_test_motion]. Unlike the exposed implementation, this method does not receive all of the arguments inside a [PhysicsTestMotionParameters2D].
		BodyTestMotion(body gd.RID, from gd.Transform2D, motion gd.Vector2, margin gd.Float, collide_separation_ray bool, recovery_as_collision bool, result *classdb.PhysicsServer2DExtensionMotionResult) bool
		//Overridable version of [method PhysicsServer2D.joint_create].
		JointCreate() gd.RID
		//Overridable version of [method PhysicsServer2D.joint_clear].
		JointClear(joint gd.RID) 
		//Overridable version of [method PhysicsServer2D.joint_set_param].
		JointSetParam(joint gd.RID, param classdb.PhysicsServer2DJointParam, value gd.Float) 
		//Overridable version of [method PhysicsServer2D.joint_get_param].
		JointGetParam(joint gd.RID, param classdb.PhysicsServer2DJointParam) gd.Float
		//Overridable version of [method PhysicsServer2D.joint_disable_collisions_between_bodies].
		JointDisableCollisionsBetweenBodies(joint gd.RID, disable bool) 
		//Overridable version of [method PhysicsServer2D.joint_is_disabled_collisions_between_bodies].
		JointIsDisabledCollisionsBetweenBodies(joint gd.RID) bool
		//Overridable version of [method PhysicsServer2D.joint_make_pin].
		JointMakePin(joint gd.RID, anchor gd.Vector2, body_a gd.RID, body_b gd.RID) 
		//Overridable version of [method PhysicsServer2D.joint_make_groove].
		JointMakeGroove(joint gd.RID, a_groove1 gd.Vector2, a_groove2 gd.Vector2, b_anchor gd.Vector2, body_a gd.RID, body_b gd.RID) 
		//Overridable version of [method PhysicsServer2D.joint_make_damped_spring].
		JointMakeDampedSpring(joint gd.RID, anchor_a gd.Vector2, anchor_b gd.Vector2, body_a gd.RID, body_b gd.RID) 
		//Overridable version of [method PhysicsServer2D.pin_joint_set_flag].
		PinJointSetFlag(joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag, enabled bool) 
		//Overridable version of [method PhysicsServer2D.pin_joint_get_flag].
		PinJointGetFlag(joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag) bool
		//Overridable version of [method PhysicsServer2D.pin_joint_set_param].
		PinJointSetParam(joint gd.RID, param classdb.PhysicsServer2DPinJointParam, value gd.Float) 
		//Overridable version of [method PhysicsServer2D.pin_joint_get_param].
		PinJointGetParam(joint gd.RID, param classdb.PhysicsServer2DPinJointParam) gd.Float
		//Overridable version of [method PhysicsServer2D.damped_spring_joint_set_param].
		DampedSpringJointSetParam(joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam, value gd.Float) 
		//Overridable version of [method PhysicsServer2D.damped_spring_joint_get_param].
		DampedSpringJointGetParam(joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam) gd.Float
		//Overridable version of [method PhysicsServer2D.joint_get_type].
		JointGetType(joint gd.RID) classdb.PhysicsServer2DJointType
		//Overridable version of [method PhysicsServer2D.free_rid].
		FreeRid(rid gd.RID) 
		//Overridable version of [method PhysicsServer2D.set_active].
		SetActive(active bool) 
		//Called when the main loop is initialized and creates a new instance of this physics server. See also [method MainLoop._initialize] and [method _finish].
		//Overridable version of [PhysicsServer2D]'s internal [code]init[/code] method.
		Init() 
		//Called every physics step to process the physics simulation. [param step] is the time elapsed since the last physics step, in seconds. It is usually the same as [method Node.get_physics_process_delta_time].
		//Overridable version of [PhysicsServer2D]'s internal [code skip-lint]step[/code] method.
		Step(step gd.Float) 
		//Called to indicate that the physics server is synchronizing and cannot access physics states if running on a separate thread. See also [method _end_sync].
		//Overridable version of [PhysicsServer2D]'s internal [code]sync[/code] method.
		Sync() 
		//Called every physics step before [method _step] to process all remaining queries.
		//Overridable version of [PhysicsServer2D]'s internal [code]flush_queries[/code] method.
		FlushQueries() 
		//Called to indicate that the physics server has stopped synchronizing. It is in the loop's iteration/physics phase, and can access physics objects even if running on a separate thread. See also [method _sync].
		//Overridable version of [PhysicsServer2D]'s internal [code]end_sync[/code] method.
		EndSync() 
		//Called when the main loop finalizes to shut down the physics server. See also [method MainLoop._finalize] and [method _init].
		//Overridable version of [PhysicsServer2D]'s internal [code]finish[/code] method.
		Finish() 
		//Overridable method that should return [code]true[/code] when the physics server is processing queries. See also [method _flush_queries].
		//Overridable version of [PhysicsServer2D]'s internal [code]is_flushing_queries[/code] method.
		IsFlushingQueries() bool
		//Overridable version of [method PhysicsServer2D.get_process_info].
		GetProcessInfo(process_info classdb.PhysicsServer2DProcessInfo) gd.Int
	}

*/
type Simple [1]classdb.PhysicsServer2DExtension
func (Simple) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _segment_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _circle_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _rectangle_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _capsule_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shape_set_data(impl func(ptr unsafe.Pointer, shape gd.RID, data gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		var data = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shape, data)
		gc.End()
	}
}
func (Simple) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID, bias float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		var bias = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shape, float64(bias))
		gc.End()
	}
}
func (Simple) _shape_get_type(impl func(ptr unsafe.Pointer, shape gd.RID) classdb.PhysicsServer2DShapeType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shape_get_data(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _shape_collide(impl func(ptr unsafe.Pointer, shape_A gd.RID, xform_A gd.Transform2D, motion_A gd.Vector2, shape_B gd.RID, xform_B gd.Transform2D, motion_B gd.Vector2, results unsafe.Pointer, result_max int, result_count *int32) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shape_A = gd.UnsafeGet[gd.RID](p_args,0)
		var xform_A = gd.UnsafeGet[gd.Transform2D](p_args,1)
		var motion_A = gd.UnsafeGet[gd.Vector2](p_args,2)
		var shape_B = gd.UnsafeGet[gd.RID](p_args,3)
		var xform_B = gd.UnsafeGet[gd.Transform2D](p_args,4)
		var motion_B = gd.UnsafeGet[gd.Vector2](p_args,5)
		var results = gd.UnsafeGet[unsafe.Pointer](p_args,6)
		var result_max = gd.UnsafeGet[gd.Int](p_args,7)
		var result_count = gd.UnsafeGet[*int32](p_args,8)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_A, xform_A, motion_A, shape_B, xform_B, motion_B, results, int(result_max), result_count)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _space_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _space_set_active(impl func(ptr unsafe.Pointer, space gd.RID, active bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		var active = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, space, active)
		gc.End()
	}
}
func (Simple) _space_is_active(impl func(ptr unsafe.Pointer, space gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _space_set_param(impl func(ptr unsafe.Pointer, space gd.RID, param classdb.PhysicsServer2DSpaceParameter, value float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DSpaceParameter](p_args,1)
		var value = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, space, param, float64(value))
		gc.End()
	}
}
func (Simple) _space_get_param(impl func(ptr unsafe.Pointer, space gd.RID, param classdb.PhysicsServer2DSpaceParameter) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DSpaceParameter](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _space_get_direct_state(impl func(ptr unsafe.Pointer, space gd.RID) [1]classdb.PhysicsDirectSpaceState2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space gd.RID, max_contacts int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		var max_contacts = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, space, int(max_contacts))
		gc.End()
	}
}
func (Simple) _space_get_contacts(impl func(ptr unsafe.Pointer, space gd.RID) gd.PackedVector2Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _space_get_contact_count(impl func(ptr unsafe.Pointer, space gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _area_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _area_set_space(impl func(ptr unsafe.Pointer, area gd.RID, space gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var space = gd.UnsafeGet[gd.RID](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, space)
		gc.End()
	}
}
func (Simple) _area_get_space(impl func(ptr unsafe.Pointer, area gd.RID) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _area_add_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape = gd.UnsafeGet[gd.RID](p_args,1)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,2)
		var disabled = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, shape, transform, disabled)
		gc.End()
	}
}
func (Simple) _area_set_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int, shape gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var shape = gd.UnsafeGet[gd.RID](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, int(shape_idx), shape)
		gc.End()
	}
}
func (Simple) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, int(shape_idx), transform)
		gc.End()
	}
}
func (Simple) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int, disabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var disabled = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, int(shape_idx), disabled)
		gc.End()
	}
}
func (Simple) _area_get_shape_count(impl func(ptr unsafe.Pointer, area gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _area_get_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _area_remove_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, int(shape_idx))
		gc.End()
	}
}
func (Simple) _area_clear_shapes(impl func(ptr unsafe.Pointer, area gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area)
		gc.End()
	}
}
func (Simple) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID, id int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, int(id))
		gc.End()
	}
}
func (Simple) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _area_attach_canvas_instance_id(impl func(ptr unsafe.Pointer, area gd.RID, id int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, int(id))
		gc.End()
	}
}
func (Simple) _area_get_canvas_instance_id(impl func(ptr unsafe.Pointer, area gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _area_set_param(impl func(ptr unsafe.Pointer, area gd.RID, param classdb.PhysicsServer2DAreaParameter, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DAreaParameter](p_args,1)
		var value = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, param, value)
		gc.End()
	}
}
func (Simple) _area_set_transform(impl func(ptr unsafe.Pointer, area gd.RID, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, transform)
		gc.End()
	}
}
func (Simple) _area_get_param(impl func(ptr unsafe.Pointer, area gd.RID, param classdb.PhysicsServer2DAreaParameter) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DAreaParameter](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _area_get_transform(impl func(ptr unsafe.Pointer, area gd.RID) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID, layer int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var layer = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, int(layer))
		gc.End()
	}
}
func (Simple) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID, mask int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var mask = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, int(mask))
		gc.End()
	}
}
func (Simple) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _area_set_monitorable(impl func(ptr unsafe.Pointer, area gd.RID, monitorable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var monitorable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, monitorable)
		gc.End()
	}
}
func (Simple) _area_set_pickable(impl func(ptr unsafe.Pointer, area gd.RID, pickable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var pickable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, pickable)
		gc.End()
	}
}
func (Simple) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var callback = mmm.Let[gd.Callable](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, callback)
		gc.End()
	}
}
func (Simple) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var callback = mmm.Let[gd.Callable](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, callback)
		gc.End()
	}
}
func (Simple) _body_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _body_set_space(impl func(ptr unsafe.Pointer, body gd.RID, space gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var space = gd.UnsafeGet[gd.RID](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, space)
		gc.End()
	}
}
func (Simple) _body_get_space(impl func(ptr unsafe.Pointer, body gd.RID) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _body_set_mode(impl func(ptr unsafe.Pointer, body gd.RID, mode classdb.PhysicsServer2DBodyMode) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var mode = gd.UnsafeGet[classdb.PhysicsServer2DBodyMode](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, mode)
		gc.End()
	}
}
func (Simple) _body_get_mode(impl func(ptr unsafe.Pointer, body gd.RID) classdb.PhysicsServer2DBodyMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _body_add_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape = gd.UnsafeGet[gd.RID](p_args,1)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,2)
		var disabled = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, shape, transform, disabled)
		gc.End()
	}
}
func (Simple) _body_set_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int, shape gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var shape = gd.UnsafeGet[gd.RID](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(shape_idx), shape)
		gc.End()
	}
}
func (Simple) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(shape_idx), transform)
		gc.End()
	}
}
func (Simple) _body_get_shape_count(impl func(ptr unsafe.Pointer, body gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _body_get_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int, disabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var disabled = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(shape_idx), disabled)
		gc.End()
	}
}
func (Simple) _body_set_shape_as_one_way_collision(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int, enable bool, margin float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var enable = gd.UnsafeGet[bool](p_args,2)
		var margin = gd.UnsafeGet[gd.Float](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(shape_idx), enable, float64(margin))
		gc.End()
	}
}
func (Simple) _body_remove_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(shape_idx))
		gc.End()
	}
}
func (Simple) _body_clear_shapes(impl func(ptr unsafe.Pointer, body gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body)
		gc.End()
	}
}
func (Simple) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID, id int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(id))
		gc.End()
	}
}
func (Simple) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _body_attach_canvas_instance_id(impl func(ptr unsafe.Pointer, body gd.RID, id int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(id))
		gc.End()
	}
}
func (Simple) _body_get_canvas_instance_id(impl func(ptr unsafe.Pointer, body gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _body_set_continuous_collision_detection_mode(impl func(ptr unsafe.Pointer, body gd.RID, mode classdb.PhysicsServer2DCCDMode) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var mode = gd.UnsafeGet[classdb.PhysicsServer2DCCDMode](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, mode)
		gc.End()
	}
}
func (Simple) _body_get_continuous_collision_detection_mode(impl func(ptr unsafe.Pointer, body gd.RID) classdb.PhysicsServer2DCCDMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID, layer int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var layer = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(layer))
		gc.End()
	}
}
func (Simple) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID, mask int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var mask = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(mask))
		gc.End()
	}
}
func (Simple) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID, priority float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var priority = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, float64(priority))
		gc.End()
	}
}
func (Simple) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _body_set_param(impl func(ptr unsafe.Pointer, body gd.RID, param classdb.PhysicsServer2DBodyParameter, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DBodyParameter](p_args,1)
		var value = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, param, value)
		gc.End()
	}
}
func (Simple) _body_get_param(impl func(ptr unsafe.Pointer, body gd.RID, param classdb.PhysicsServer2DBodyParameter) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DBodyParameter](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body)
		gc.End()
	}
}
func (Simple) _body_set_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer2DBodyState, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var state = gd.UnsafeGet[classdb.PhysicsServer2DBodyState](p_args,1)
		var value = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, state, value)
		gc.End()
	}
}
func (Simple) _body_get_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer2DBodyState) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var state = gd.UnsafeGet[classdb.PhysicsServer2DBodyState](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, impulse)
		gc.End()
	}
}
func (Simple) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var impulse = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, float64(impulse))
		gc.End()
	}
}
func (Simple) _body_apply_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,1)
		var position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, impulse, position)
		gc.End()
	}
}
func (Simple) _body_apply_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force)
		gc.End()
	}
}
func (Simple) _body_apply_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		var position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force, position)
		gc.End()
	}
}
func (Simple) _body_apply_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var torque = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, float64(torque))
		gc.End()
	}
}
func (Simple) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force)
		gc.End()
	}
}
func (Simple) _body_add_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		var position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force, position)
		gc.End()
	}
}
func (Simple) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var torque = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, float64(torque))
		gc.End()
	}
}
func (Simple) _body_set_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force)
		gc.End()
	}
}
func (Simple) _body_get_constant_force(impl func(ptr unsafe.Pointer, body gd.RID) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var torque = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, float64(torque))
		gc.End()
	}
}
func (Simple) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body gd.RID, axis_velocity gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var axis_velocity = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, axis_velocity)
		gc.End()
	}
}
func (Simple) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, excepted_body)
		gc.End()
	}
}
func (Simple) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, excepted_body)
		gc.End()
	}
}
func (Simple) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body gd.RID) gd.ArrayOf[gd.RID], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Simple) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID, amount int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var amount = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, int(amount))
		gc.End()
	}
}
func (Simple) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID, threshold float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var threshold = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, float64(threshold))
		gc.End()
	}
}
func (Simple) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body gd.RID, enable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var enable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, enable)
		gc.End()
	}
}
func (Simple) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var callable = mmm.Let[gd.Callable](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, callable)
		gc.End()
	}
}
func (Simple) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable, userdata gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var callable = mmm.Let[gd.Callable](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		var userdata = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, callable, userdata)
		gc.End()
	}
}
func (Simple) _body_collide_shape(impl func(ptr unsafe.Pointer, body gd.RID, body_shape int, shape gd.RID, shape_xform gd.Transform2D, motion gd.Vector2, results unsafe.Pointer, result_max int, result_count *int32) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var body_shape = gd.UnsafeGet[gd.Int](p_args,1)
		var shape = gd.UnsafeGet[gd.RID](p_args,2)
		var shape_xform = gd.UnsafeGet[gd.Transform2D](p_args,3)
		var motion = gd.UnsafeGet[gd.Vector2](p_args,4)
		var results = gd.UnsafeGet[unsafe.Pointer](p_args,5)
		var result_max = gd.UnsafeGet[gd.Int](p_args,6)
		var result_count = gd.UnsafeGet[*int32](p_args,7)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(body_shape), shape, shape_xform, motion, results, int(result_max), result_count)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _body_set_pickable(impl func(ptr unsafe.Pointer, body gd.RID, pickable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var pickable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, pickable)
		gc.End()
	}
}
func (Simple) _body_get_direct_state(impl func(ptr unsafe.Pointer, body gd.RID) [1]classdb.PhysicsDirectBodyState2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _body_test_motion(impl func(ptr unsafe.Pointer, body gd.RID, from gd.Transform2D, motion gd.Vector2, margin float64, collide_separation_ray bool, recovery_as_collision bool, result *classdb.PhysicsServer2DExtensionMotionResult) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var from = gd.UnsafeGet[gd.Transform2D](p_args,1)
		var motion = gd.UnsafeGet[gd.Vector2](p_args,2)
		var margin = gd.UnsafeGet[gd.Float](p_args,3)
		var collide_separation_ray = gd.UnsafeGet[bool](p_args,4)
		var recovery_as_collision = gd.UnsafeGet[bool](p_args,5)
		var result = gd.UnsafeGet[*classdb.PhysicsServer2DExtensionMotionResult](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, float64(margin), collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _joint_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _joint_clear(impl func(ptr unsafe.Pointer, joint gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint)
		gc.End()
	}
}
func (Simple) _joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DJointParam, value float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DJointParam](p_args,1)
		var value = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, param, float64(value))
		gc.End()
	}
}
func (Simple) _joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DJointParam) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DJointParam](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID, disable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var disable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, disable)
		gc.End()
	}
}
func (Simple) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _joint_make_pin(impl func(ptr unsafe.Pointer, joint gd.RID, anchor gd.Vector2, body_a gd.RID, body_b gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var anchor = gd.UnsafeGet[gd.Vector2](p_args,1)
		var body_a = gd.UnsafeGet[gd.RID](p_args,2)
		var body_b = gd.UnsafeGet[gd.RID](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, anchor, body_a, body_b)
		gc.End()
	}
}
func (Simple) _joint_make_groove(impl func(ptr unsafe.Pointer, joint gd.RID, a_groove1 gd.Vector2, a_groove2 gd.Vector2, b_anchor gd.Vector2, body_a gd.RID, body_b gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var a_groove1 = gd.UnsafeGet[gd.Vector2](p_args,1)
		var a_groove2 = gd.UnsafeGet[gd.Vector2](p_args,2)
		var b_anchor = gd.UnsafeGet[gd.Vector2](p_args,3)
		var body_a = gd.UnsafeGet[gd.RID](p_args,4)
		var body_b = gd.UnsafeGet[gd.RID](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, a_groove1, a_groove2, b_anchor, body_a, body_b)
		gc.End()
	}
}
func (Simple) _joint_make_damped_spring(impl func(ptr unsafe.Pointer, joint gd.RID, anchor_a gd.Vector2, anchor_b gd.Vector2, body_a gd.RID, body_b gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var anchor_a = gd.UnsafeGet[gd.Vector2](p_args,1)
		var anchor_b = gd.UnsafeGet[gd.Vector2](p_args,2)
		var body_a = gd.UnsafeGet[gd.RID](p_args,3)
		var body_b = gd.UnsafeGet[gd.RID](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, anchor_a, anchor_b, body_a, body_b)
		gc.End()
	}
}
func (Simple) _pin_joint_set_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var flag = gd.UnsafeGet[classdb.PhysicsServer2DPinJointFlag](p_args,1)
		var enabled = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, flag, enabled)
		gc.End()
	}
}
func (Simple) _pin_joint_get_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var flag = gd.UnsafeGet[classdb.PhysicsServer2DPinJointFlag](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DPinJointParam, value float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DPinJointParam](p_args,1)
		var value = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, param, float64(value))
		gc.End()
	}
}
func (Simple) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DPinJointParam) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DPinJointParam](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _damped_spring_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam, value float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DDampedSpringParam](p_args,1)
		var value = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, param, float64(value))
		gc.End()
	}
}
func (Simple) _damped_spring_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DDampedSpringParam](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _joint_get_type(impl func(ptr unsafe.Pointer, joint gd.RID) classdb.PhysicsServer2DJointType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _free_rid(impl func(ptr unsafe.Pointer, rid gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, rid)
		gc.End()
	}
}
func (Simple) _set_active(impl func(ptr unsafe.Pointer, active bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var active = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, active)
		gc.End()
	}
}
func (Simple) _init(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _step(impl func(ptr unsafe.Pointer, step float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var step = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(step))
		gc.End()
	}
}
func (Simple) _sync(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _flush_queries(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _end_sync(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _finish(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_process_info(impl func(ptr unsafe.Pointer, process_info classdb.PhysicsServer2DProcessInfo) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var process_info = gd.UnsafeGet[classdb.PhysicsServer2DProcessInfo](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (self Simple) BodyTestMotionIsExcludingBody(body gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).BodyTestMotionIsExcludingBody(body))
}
func (self Simple) BodyTestMotionIsExcludingObject(obj int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).BodyTestMotionIsExcludingObject(gd.Int(obj)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicsServer2DExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Overridable version of [method PhysicsServer2D.world_boundary_shape_create].
*/
func (class) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.separation_ray_shape_create].
*/
func (class) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.segment_shape_create].
*/
func (class) _segment_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.circle_shape_create].
*/
func (class) _circle_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.rectangle_shape_create].
*/
func (class) _rectangle_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.capsule_shape_create].
*/
func (class) _capsule_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.convex_polygon_shape_create].
*/
func (class) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.concave_polygon_shape_create].
*/
func (class) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.shape_set_data].
*/
func (class) _shape_set_data(impl func(ptr unsafe.Pointer, shape gd.RID, data gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		var data = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shape, data)
		ctx.End()
	}
}

/*
Should set the custom solver bias for the given [param shape]. It defines how much bodies are forced to separate on contact.
Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
*/
func (class) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID, bias gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		var bias = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shape, bias)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.shape_get_type].
*/
func (class) _shape_get_type(impl func(ptr unsafe.Pointer, shape gd.RID) classdb.PhysicsServer2DShapeType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.shape_get_data].
*/
func (class) _shape_get_data(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Should return the custom solver bias of the given [param shape], which defines how much bodies are forced to separate on contact when this shape is involved.
Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
*/
func (class) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Given two shapes and their parameters, should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
*/
func (class) _shape_collide(impl func(ptr unsafe.Pointer, shape_A gd.RID, xform_A gd.Transform2D, motion_A gd.Vector2, shape_B gd.RID, xform_B gd.Transform2D, motion_B gd.Vector2, results unsafe.Pointer, result_max gd.Int, result_count *int32) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shape_A = gd.UnsafeGet[gd.RID](p_args,0)
		var xform_A = gd.UnsafeGet[gd.Transform2D](p_args,1)
		var motion_A = gd.UnsafeGet[gd.Vector2](p_args,2)
		var shape_B = gd.UnsafeGet[gd.RID](p_args,3)
		var xform_B = gd.UnsafeGet[gd.Transform2D](p_args,4)
		var motion_B = gd.UnsafeGet[gd.Vector2](p_args,5)
		var results = gd.UnsafeGet[unsafe.Pointer](p_args,6)
		var result_max = gd.UnsafeGet[gd.Int](p_args,7)
		var result_count = gd.UnsafeGet[*int32](p_args,8)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_A, xform_A, motion_A, shape_B, xform_B, motion_B, results, result_max, result_count)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.space_create].
*/
func (class) _space_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.space_set_active].
*/
func (class) _space_set_active(impl func(ptr unsafe.Pointer, space gd.RID, active bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		var active = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, space, active)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.space_is_active].
*/
func (class) _space_is_active(impl func(ptr unsafe.Pointer, space gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.space_set_param].
*/
func (class) _space_set_param(impl func(ptr unsafe.Pointer, space gd.RID, param classdb.PhysicsServer2DSpaceParameter, value gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DSpaceParameter](p_args,1)
		var value = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, space, param, value)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.space_get_param].
*/
func (class) _space_get_param(impl func(ptr unsafe.Pointer, space gd.RID, param classdb.PhysicsServer2DSpaceParameter) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DSpaceParameter](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.space_get_direct_state].
*/
func (class) _space_get_direct_state(impl func(ptr unsafe.Pointer, space gd.RID) object.PhysicsDirectSpaceState2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Used internally to allow the given [param space] to store contact points, up to [param max_contacts]. This is automatically set for the main [World2D]'s space when [member SceneTree.debug_collisions_hint] is [code]true[/code], or by checking "Visible Collision Shapes" in the editor. Only works in debug builds.
Overridable version of [PhysicsServer2D]'s internal [code]space_set_debug_contacts[/code] method.
*/
func (class) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space gd.RID, max_contacts gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		var max_contacts = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, space, max_contacts)
		ctx.End()
	}
}

/*
Should return the positions of all contacts that have occurred during the last physics step in the given [param space]. See also [method _space_get_contact_count] and [method _space_set_debug_contacts].
Overridable version of [PhysicsServer2D]'s internal [code]space_get_contacts[/code] method.
*/
func (class) _space_get_contacts(impl func(ptr unsafe.Pointer, space gd.RID) gd.PackedVector2Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Should return how many contacts have occurred during the last physics step in the given [param space]. See also [method _space_get_contacts] and [method _space_set_debug_contacts].
Overridable version of [PhysicsServer2D]'s internal [code]space_get_contact_count[/code] method.
*/
func (class) _space_get_contact_count(impl func(ptr unsafe.Pointer, space gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var space = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_create].
*/
func (class) _area_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.area_set_space].
*/
func (class) _area_set_space(impl func(ptr unsafe.Pointer, area gd.RID, space gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var space = gd.UnsafeGet[gd.RID](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, space)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_space].
*/
func (class) _area_get_space(impl func(ptr unsafe.Pointer, area gd.RID) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_add_shape].
*/
func (class) _area_add_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape = gd.UnsafeGet[gd.RID](p_args,1)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,2)
		var disabled = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, shape, transform, disabled)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_shape].
*/
func (class) _area_set_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, shape gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var shape = gd.UnsafeGet[gd.RID](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, shape_idx, shape)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_shape_transform].
*/
func (class) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, shape_idx, transform)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_shape_disabled].
*/
func (class) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, disabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var disabled = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, shape_idx, disabled)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_shape_count].
*/
func (class) _area_get_shape_count(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_shape].
*/
func (class) _area_get_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_shape_transform].
*/
func (class) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_remove_shape].
*/
func (class) _area_remove_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, shape_idx)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_clear_shapes].
*/
func (class) _area_clear_shapes(impl func(ptr unsafe.Pointer, area gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_attach_object_instance_id].
*/
func (class) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID, id gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, id)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_object_instance_id].
*/
func (class) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_attach_canvas_instance_id].
*/
func (class) _area_attach_canvas_instance_id(impl func(ptr unsafe.Pointer, area gd.RID, id gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, id)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_canvas_instance_id].
*/
func (class) _area_get_canvas_instance_id(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_param].
*/
func (class) _area_set_param(impl func(ptr unsafe.Pointer, area gd.RID, param classdb.PhysicsServer2DAreaParameter, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DAreaParameter](p_args,1)
		var value = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, param, value)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_transform].
*/
func (class) _area_set_transform(impl func(ptr unsafe.Pointer, area gd.RID, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, transform)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_param].
*/
func (class) _area_get_param(impl func(ptr unsafe.Pointer, area gd.RID, param classdb.PhysicsServer2DAreaParameter) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DAreaParameter](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_transform].
*/
func (class) _area_get_transform(impl func(ptr unsafe.Pointer, area gd.RID) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_collision_layer].
*/
func (class) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID, layer gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var layer = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, layer)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_collision_layer].
*/
func (class) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_collision_mask].
*/
func (class) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID, mask gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var mask = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, mask)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_collision_mask].
*/
func (class) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_monitorable].
*/
func (class) _area_set_monitorable(impl func(ptr unsafe.Pointer, area gd.RID, monitorable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var monitorable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, monitorable)
		ctx.End()
	}
}

/*
If set to [code]true[/code], allows the area with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
Overridable version of [PhysicsServer2D]'s internal [code]area_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
*/
func (class) _area_set_pickable(impl func(ptr unsafe.Pointer, area gd.RID, pickable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var pickable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, pickable)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_monitor_callback].
*/
func (class) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var callback = mmm.Let[gd.Callable](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, callback)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_area_monitor_callback].
*/
func (class) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var area = gd.UnsafeGet[gd.RID](p_args,0)
		var callback = mmm.Let[gd.Callable](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, area, callback)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_create].
*/
func (class) _body_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.body_set_space].
*/
func (class) _body_set_space(impl func(ptr unsafe.Pointer, body gd.RID, space gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var space = gd.UnsafeGet[gd.RID](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, space)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_space].
*/
func (class) _body_get_space(impl func(ptr unsafe.Pointer, body gd.RID) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_mode].
*/
func (class) _body_set_mode(impl func(ptr unsafe.Pointer, body gd.RID, mode classdb.PhysicsServer2DBodyMode) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var mode = gd.UnsafeGet[classdb.PhysicsServer2DBodyMode](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, mode)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_mode].
*/
func (class) _body_get_mode(impl func(ptr unsafe.Pointer, body gd.RID) classdb.PhysicsServer2DBodyMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_shape].
*/
func (class) _body_add_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape = gd.UnsafeGet[gd.RID](p_args,1)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,2)
		var disabled = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, shape, transform, disabled)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape].
*/
func (class) _body_set_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, shape gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var shape = gd.UnsafeGet[gd.RID](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, shape_idx, shape)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape_transform].
*/
func (class) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, shape_idx, transform)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_shape_count].
*/
func (class) _body_get_shape_count(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_shape].
*/
func (class) _body_get_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_shape_transform].
*/
func (class) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape_disabled].
*/
func (class) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, disabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var disabled = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, shape_idx, disabled)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape_as_one_way_collision].
*/
func (class) _body_set_shape_as_one_way_collision(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, enable bool, margin gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		var enable = gd.UnsafeGet[bool](p_args,2)
		var margin = gd.UnsafeGet[gd.Float](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, shape_idx, enable, margin)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_remove_shape].
*/
func (class) _body_remove_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, shape_idx)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_clear_shapes].
*/
func (class) _body_clear_shapes(impl func(ptr unsafe.Pointer, body gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_attach_object_instance_id].
*/
func (class) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID, id gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, id)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_object_instance_id].
*/
func (class) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_attach_canvas_instance_id].
*/
func (class) _body_attach_canvas_instance_id(impl func(ptr unsafe.Pointer, body gd.RID, id gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var id = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, id)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_canvas_instance_id].
*/
func (class) _body_get_canvas_instance_id(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_continuous_collision_detection_mode].
*/
func (class) _body_set_continuous_collision_detection_mode(impl func(ptr unsafe.Pointer, body gd.RID, mode classdb.PhysicsServer2DCCDMode) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var mode = gd.UnsafeGet[classdb.PhysicsServer2DCCDMode](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, mode)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_continuous_collision_detection_mode].
*/
func (class) _body_get_continuous_collision_detection_mode(impl func(ptr unsafe.Pointer, body gd.RID) classdb.PhysicsServer2DCCDMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_collision_layer].
*/
func (class) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID, layer gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var layer = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, layer)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_collision_layer].
*/
func (class) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_collision_mask].
*/
func (class) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID, mask gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var mask = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, mask)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_collision_mask].
*/
func (class) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_collision_priority].
*/
func (class) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID, priority gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var priority = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, priority)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_collision_priority].
*/
func (class) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_param].
*/
func (class) _body_set_param(impl func(ptr unsafe.Pointer, body gd.RID, param classdb.PhysicsServer2DBodyParameter, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DBodyParameter](p_args,1)
		var value = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, param, value)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_param].
*/
func (class) _body_get_param(impl func(ptr unsafe.Pointer, body gd.RID, param classdb.PhysicsServer2DBodyParameter) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DBodyParameter](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_reset_mass_properties].
*/
func (class) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_state].
*/
func (class) _body_set_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer2DBodyState, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var state = gd.UnsafeGet[classdb.PhysicsServer2DBodyState](p_args,1)
		var value = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, state, value)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_state].
*/
func (class) _body_get_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer2DBodyState) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var state = gd.UnsafeGet[classdb.PhysicsServer2DBodyState](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_central_impulse].
*/
func (class) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, impulse)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_torque_impulse].
*/
func (class) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var impulse = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, impulse)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_impulse].
*/
func (class) _body_apply_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,1)
		var position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, impulse, position)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_central_force].
*/
func (class) _body_apply_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_force].
*/
func (class) _body_apply_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		var position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force, position)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_torque].
*/
func (class) _body_apply_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var torque = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, torque)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_constant_central_force].
*/
func (class) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_constant_force].
*/
func (class) _body_add_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		var position = gd.UnsafeGet[gd.Vector2](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force, position)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_constant_torque].
*/
func (class) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var torque = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, torque)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_constant_force].
*/
func (class) _body_set_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var force = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, force)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_constant_force].
*/
func (class) _body_get_constant_force(impl func(ptr unsafe.Pointer, body gd.RID) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_constant_torque].
*/
func (class) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var torque = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, torque)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_constant_torque].
*/
func (class) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_axis_velocity].
*/
func (class) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body gd.RID, axis_velocity gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var axis_velocity = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, axis_velocity)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_collision_exception].
*/
func (class) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, excepted_body)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_remove_collision_exception].
*/
func (class) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, excepted_body)
		ctx.End()
	}
}

/*
Returns the [RID]s of all bodies added as collision exceptions for the given [param body]. See also [method _body_add_collision_exception] and [method _body_remove_collision_exception].
Overridable version of [PhysicsServer2D]'s internal [code]body_get_collision_exceptions[/code] method. Corresponds to [method PhysicsBody2D.get_collision_exceptions].
*/
func (class) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body gd.RID) gd.ArrayOf[gd.RID], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_max_contacts_reported].
*/
func (class) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID, amount gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var amount = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, amount)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_max_contacts_reported].
*/
func (class) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [PhysicsServer2D]'s internal [code]body_set_contacts_reported_depth_threshold[/code] method.
[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
*/
func (class) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID, threshold gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var threshold = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, threshold)
		ctx.End()
	}
}

/*
Overridable version of [PhysicsServer2D]'s internal [code]body_get_contacts_reported_depth_threshold[/code] method.
[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
*/
func (class) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_omit_force_integration].
*/
func (class) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body gd.RID, enable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var enable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, enable)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_is_omitting_force_integration].
*/
func (class) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Assigns the [param body] to call the given [param callable] during the synchronization phase of the loop, before [method _step] is called. See also [method _sync].
Overridable version of [method PhysicsServer2D.body_set_state_sync_callback].
*/
func (class) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var callable = mmm.Let[gd.Callable](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, callable)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_force_integration_callback].
*/
func (class) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable, userdata gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var callable = mmm.Let[gd.Callable](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		var userdata = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, callable, userdata)
		ctx.End()
	}
}

/*
Given a [param body], a [param shape], and their respective parameters, this method should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
*/
func (class) _body_collide_shape(impl func(ptr unsafe.Pointer, body gd.RID, body_shape gd.Int, shape gd.RID, shape_xform gd.Transform2D, motion gd.Vector2, results unsafe.Pointer, result_max gd.Int, result_count *int32) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var body_shape = gd.UnsafeGet[gd.Int](p_args,1)
		var shape = gd.UnsafeGet[gd.RID](p_args,2)
		var shape_xform = gd.UnsafeGet[gd.Transform2D](p_args,3)
		var motion = gd.UnsafeGet[gd.Vector2](p_args,4)
		var results = gd.UnsafeGet[unsafe.Pointer](p_args,5)
		var result_max = gd.UnsafeGet[gd.Int](p_args,6)
		var result_count = gd.UnsafeGet[*int32](p_args,7)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, body_shape, shape, shape_xform, motion, results, result_max, result_count)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
If set to [code]true[/code], allows the body with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
Overridable version of [PhysicsServer2D]'s internal [code]body_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
*/
func (class) _body_set_pickable(impl func(ptr unsafe.Pointer, body gd.RID, pickable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var pickable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, body, pickable)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_direct_state].
*/
func (class) _body_get_direct_state(impl func(ptr unsafe.Pointer, body gd.RID) object.PhysicsDirectBodyState2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.body_test_motion]. Unlike the exposed implementation, this method does not receive all of the arguments inside a [PhysicsTestMotionParameters2D].
*/
func (class) _body_test_motion(impl func(ptr unsafe.Pointer, body gd.RID, from gd.Transform2D, motion gd.Vector2, margin gd.Float, collide_separation_ray bool, recovery_as_collision bool, result *classdb.PhysicsServer2DExtensionMotionResult) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var body = gd.UnsafeGet[gd.RID](p_args,0)
		var from = gd.UnsafeGet[gd.Transform2D](p_args,1)
		var motion = gd.UnsafeGet[gd.Vector2](p_args,2)
		var margin = gd.UnsafeGet[gd.Float](p_args,3)
		var collide_separation_ray = gd.UnsafeGet[bool](p_args,4)
		var recovery_as_collision = gd.UnsafeGet[bool](p_args,5)
		var result = gd.UnsafeGet[*classdb.PhysicsServer2DExtensionMotionResult](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, margin, collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_create].
*/
func (class) _joint_create(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.joint_clear].
*/
func (class) _joint_clear(impl func(ptr unsafe.Pointer, joint gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_set_param].
*/
func (class) _joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DJointParam, value gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DJointParam](p_args,1)
		var value = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, param, value)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_get_param].
*/
func (class) _joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DJointParam) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DJointParam](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_disable_collisions_between_bodies].
*/
func (class) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID, disable bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var disable = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, disable)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_is_disabled_collisions_between_bodies].
*/
func (class) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_make_pin].
*/
func (class) _joint_make_pin(impl func(ptr unsafe.Pointer, joint gd.RID, anchor gd.Vector2, body_a gd.RID, body_b gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var anchor = gd.UnsafeGet[gd.Vector2](p_args,1)
		var body_a = gd.UnsafeGet[gd.RID](p_args,2)
		var body_b = gd.UnsafeGet[gd.RID](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, anchor, body_a, body_b)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_make_groove].
*/
func (class) _joint_make_groove(impl func(ptr unsafe.Pointer, joint gd.RID, a_groove1 gd.Vector2, a_groove2 gd.Vector2, b_anchor gd.Vector2, body_a gd.RID, body_b gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var a_groove1 = gd.UnsafeGet[gd.Vector2](p_args,1)
		var a_groove2 = gd.UnsafeGet[gd.Vector2](p_args,2)
		var b_anchor = gd.UnsafeGet[gd.Vector2](p_args,3)
		var body_a = gd.UnsafeGet[gd.RID](p_args,4)
		var body_b = gd.UnsafeGet[gd.RID](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, a_groove1, a_groove2, b_anchor, body_a, body_b)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_make_damped_spring].
*/
func (class) _joint_make_damped_spring(impl func(ptr unsafe.Pointer, joint gd.RID, anchor_a gd.Vector2, anchor_b gd.Vector2, body_a gd.RID, body_b gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var anchor_a = gd.UnsafeGet[gd.Vector2](p_args,1)
		var anchor_b = gd.UnsafeGet[gd.Vector2](p_args,2)
		var body_a = gd.UnsafeGet[gd.RID](p_args,3)
		var body_b = gd.UnsafeGet[gd.RID](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, anchor_a, anchor_b, body_a, body_b)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_set_flag].
*/
func (class) _pin_joint_set_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var flag = gd.UnsafeGet[classdb.PhysicsServer2DPinJointFlag](p_args,1)
		var enabled = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, flag, enabled)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_get_flag].
*/
func (class) _pin_joint_get_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag classdb.PhysicsServer2DPinJointFlag) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var flag = gd.UnsafeGet[classdb.PhysicsServer2DPinJointFlag](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_set_param].
*/
func (class) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DPinJointParam, value gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DPinJointParam](p_args,1)
		var value = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, param, value)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_get_param].
*/
func (class) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DPinJointParam) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DPinJointParam](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.damped_spring_joint_set_param].
*/
func (class) _damped_spring_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam, value gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DDampedSpringParam](p_args,1)
		var value = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, joint, param, value)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.damped_spring_joint_get_param].
*/
func (class) _damped_spring_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer2DDampedSpringParam) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		var param = gd.UnsafeGet[classdb.PhysicsServer2DDampedSpringParam](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_get_type].
*/
func (class) _joint_get_type(impl func(ptr unsafe.Pointer, joint gd.RID) classdb.PhysicsServer2DJointType, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var joint = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.free_rid].
*/
func (class) _free_rid(impl func(ptr unsafe.Pointer, rid gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, rid)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsServer2D.set_active].
*/
func (class) _set_active(impl func(ptr unsafe.Pointer, active bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var active = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, active)
		ctx.End()
	}
}

/*
Called when the main loop is initialized and creates a new instance of this physics server. See also [method MainLoop._initialize] and [method _finish].
Overridable version of [PhysicsServer2D]'s internal [code]init[/code] method.
*/
func (class) _init(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called every physics step to process the physics simulation. [param step] is the time elapsed since the last physics step, in seconds. It is usually the same as [method Node.get_physics_process_delta_time].
Overridable version of [PhysicsServer2D]'s internal [code skip-lint]step[/code] method.
*/
func (class) _step(impl func(ptr unsafe.Pointer, step gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var step = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, step)
		ctx.End()
	}
}

/*
Called to indicate that the physics server is synchronizing and cannot access physics states if running on a separate thread. See also [method _end_sync].
Overridable version of [PhysicsServer2D]'s internal [code]sync[/code] method.
*/
func (class) _sync(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called every physics step before [method _step] to process all remaining queries.
Overridable version of [PhysicsServer2D]'s internal [code]flush_queries[/code] method.
*/
func (class) _flush_queries(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called to indicate that the physics server has stopped synchronizing. It is in the loop's iteration/physics phase, and can access physics objects even if running on a separate thread. See also [method _sync].
Overridable version of [PhysicsServer2D]'s internal [code]end_sync[/code] method.
*/
func (class) _end_sync(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the main loop finalizes to shut down the physics server. See also [method MainLoop._finalize] and [method _init].
Overridable version of [PhysicsServer2D]'s internal [code]finish[/code] method.
*/
func (class) _finish(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Overridable method that should return [code]true[/code] when the physics server is processing queries. See also [method _flush_queries].
Overridable version of [PhysicsServer2D]'s internal [code]is_flushing_queries[/code] method.
*/
func (class) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable version of [method PhysicsServer2D.get_process_info].
*/
func (class) _get_process_info(impl func(ptr unsafe.Pointer, process_info classdb.PhysicsServer2DProcessInfo) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var process_info = gd.UnsafeGet[classdb.PhysicsServer2DProcessInfo](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns [code]true[/code] if the body with the given [RID] is being excluded from [method _body_test_motion]. See also [method Object.get_instance_id].
*/
//go:nosplit
func (self class) BodyTestMotionIsExcludingBody(body gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2DExtension.Bind_body_test_motion_is_excluding_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the object with the given instance ID is being excluded from [method _body_test_motion]. See also [method Object.get_instance_id].
*/
//go:nosplit
func (self class) BodyTestMotionIsExcludingObject(obj gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obj)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PhysicsServer2DExtension.Bind_body_test_motion_is_excluding_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPhysicsServer2DExtension() Expert { return self[0].AsPhysicsServer2DExtension() }


//go:nosplit
func (self Simple) AsPhysicsServer2DExtension() Simple { return self[0].AsPhysicsServer2DExtension() }


//go:nosplit
func (self class) AsPhysicsServer2D() PhysicsServer2D.Expert { return self[0].AsPhysicsServer2D() }


//go:nosplit
func (self Simple) AsPhysicsServer2D() PhysicsServer2D.Simple { return self[0].AsPhysicsServer2D() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_world_boundary_shape_create": return reflect.ValueOf(self._world_boundary_shape_create);
	case "_separation_ray_shape_create": return reflect.ValueOf(self._separation_ray_shape_create);
	case "_segment_shape_create": return reflect.ValueOf(self._segment_shape_create);
	case "_circle_shape_create": return reflect.ValueOf(self._circle_shape_create);
	case "_rectangle_shape_create": return reflect.ValueOf(self._rectangle_shape_create);
	case "_capsule_shape_create": return reflect.ValueOf(self._capsule_shape_create);
	case "_convex_polygon_shape_create": return reflect.ValueOf(self._convex_polygon_shape_create);
	case "_concave_polygon_shape_create": return reflect.ValueOf(self._concave_polygon_shape_create);
	case "_shape_set_data": return reflect.ValueOf(self._shape_set_data);
	case "_shape_set_custom_solver_bias": return reflect.ValueOf(self._shape_set_custom_solver_bias);
	case "_shape_get_type": return reflect.ValueOf(self._shape_get_type);
	case "_shape_get_data": return reflect.ValueOf(self._shape_get_data);
	case "_shape_get_custom_solver_bias": return reflect.ValueOf(self._shape_get_custom_solver_bias);
	case "_shape_collide": return reflect.ValueOf(self._shape_collide);
	case "_space_create": return reflect.ValueOf(self._space_create);
	case "_space_set_active": return reflect.ValueOf(self._space_set_active);
	case "_space_is_active": return reflect.ValueOf(self._space_is_active);
	case "_space_set_param": return reflect.ValueOf(self._space_set_param);
	case "_space_get_param": return reflect.ValueOf(self._space_get_param);
	case "_space_get_direct_state": return reflect.ValueOf(self._space_get_direct_state);
	case "_space_set_debug_contacts": return reflect.ValueOf(self._space_set_debug_contacts);
	case "_space_get_contacts": return reflect.ValueOf(self._space_get_contacts);
	case "_space_get_contact_count": return reflect.ValueOf(self._space_get_contact_count);
	case "_area_create": return reflect.ValueOf(self._area_create);
	case "_area_set_space": return reflect.ValueOf(self._area_set_space);
	case "_area_get_space": return reflect.ValueOf(self._area_get_space);
	case "_area_add_shape": return reflect.ValueOf(self._area_add_shape);
	case "_area_set_shape": return reflect.ValueOf(self._area_set_shape);
	case "_area_set_shape_transform": return reflect.ValueOf(self._area_set_shape_transform);
	case "_area_set_shape_disabled": return reflect.ValueOf(self._area_set_shape_disabled);
	case "_area_get_shape_count": return reflect.ValueOf(self._area_get_shape_count);
	case "_area_get_shape": return reflect.ValueOf(self._area_get_shape);
	case "_area_get_shape_transform": return reflect.ValueOf(self._area_get_shape_transform);
	case "_area_remove_shape": return reflect.ValueOf(self._area_remove_shape);
	case "_area_clear_shapes": return reflect.ValueOf(self._area_clear_shapes);
	case "_area_attach_object_instance_id": return reflect.ValueOf(self._area_attach_object_instance_id);
	case "_area_get_object_instance_id": return reflect.ValueOf(self._area_get_object_instance_id);
	case "_area_attach_canvas_instance_id": return reflect.ValueOf(self._area_attach_canvas_instance_id);
	case "_area_get_canvas_instance_id": return reflect.ValueOf(self._area_get_canvas_instance_id);
	case "_area_set_param": return reflect.ValueOf(self._area_set_param);
	case "_area_set_transform": return reflect.ValueOf(self._area_set_transform);
	case "_area_get_param": return reflect.ValueOf(self._area_get_param);
	case "_area_get_transform": return reflect.ValueOf(self._area_get_transform);
	case "_area_set_collision_layer": return reflect.ValueOf(self._area_set_collision_layer);
	case "_area_get_collision_layer": return reflect.ValueOf(self._area_get_collision_layer);
	case "_area_set_collision_mask": return reflect.ValueOf(self._area_set_collision_mask);
	case "_area_get_collision_mask": return reflect.ValueOf(self._area_get_collision_mask);
	case "_area_set_monitorable": return reflect.ValueOf(self._area_set_monitorable);
	case "_area_set_pickable": return reflect.ValueOf(self._area_set_pickable);
	case "_area_set_monitor_callback": return reflect.ValueOf(self._area_set_monitor_callback);
	case "_area_set_area_monitor_callback": return reflect.ValueOf(self._area_set_area_monitor_callback);
	case "_body_create": return reflect.ValueOf(self._body_create);
	case "_body_set_space": return reflect.ValueOf(self._body_set_space);
	case "_body_get_space": return reflect.ValueOf(self._body_get_space);
	case "_body_set_mode": return reflect.ValueOf(self._body_set_mode);
	case "_body_get_mode": return reflect.ValueOf(self._body_get_mode);
	case "_body_add_shape": return reflect.ValueOf(self._body_add_shape);
	case "_body_set_shape": return reflect.ValueOf(self._body_set_shape);
	case "_body_set_shape_transform": return reflect.ValueOf(self._body_set_shape_transform);
	case "_body_get_shape_count": return reflect.ValueOf(self._body_get_shape_count);
	case "_body_get_shape": return reflect.ValueOf(self._body_get_shape);
	case "_body_get_shape_transform": return reflect.ValueOf(self._body_get_shape_transform);
	case "_body_set_shape_disabled": return reflect.ValueOf(self._body_set_shape_disabled);
	case "_body_set_shape_as_one_way_collision": return reflect.ValueOf(self._body_set_shape_as_one_way_collision);
	case "_body_remove_shape": return reflect.ValueOf(self._body_remove_shape);
	case "_body_clear_shapes": return reflect.ValueOf(self._body_clear_shapes);
	case "_body_attach_object_instance_id": return reflect.ValueOf(self._body_attach_object_instance_id);
	case "_body_get_object_instance_id": return reflect.ValueOf(self._body_get_object_instance_id);
	case "_body_attach_canvas_instance_id": return reflect.ValueOf(self._body_attach_canvas_instance_id);
	case "_body_get_canvas_instance_id": return reflect.ValueOf(self._body_get_canvas_instance_id);
	case "_body_set_continuous_collision_detection_mode": return reflect.ValueOf(self._body_set_continuous_collision_detection_mode);
	case "_body_get_continuous_collision_detection_mode": return reflect.ValueOf(self._body_get_continuous_collision_detection_mode);
	case "_body_set_collision_layer": return reflect.ValueOf(self._body_set_collision_layer);
	case "_body_get_collision_layer": return reflect.ValueOf(self._body_get_collision_layer);
	case "_body_set_collision_mask": return reflect.ValueOf(self._body_set_collision_mask);
	case "_body_get_collision_mask": return reflect.ValueOf(self._body_get_collision_mask);
	case "_body_set_collision_priority": return reflect.ValueOf(self._body_set_collision_priority);
	case "_body_get_collision_priority": return reflect.ValueOf(self._body_get_collision_priority);
	case "_body_set_param": return reflect.ValueOf(self._body_set_param);
	case "_body_get_param": return reflect.ValueOf(self._body_get_param);
	case "_body_reset_mass_properties": return reflect.ValueOf(self._body_reset_mass_properties);
	case "_body_set_state": return reflect.ValueOf(self._body_set_state);
	case "_body_get_state": return reflect.ValueOf(self._body_get_state);
	case "_body_apply_central_impulse": return reflect.ValueOf(self._body_apply_central_impulse);
	case "_body_apply_torque_impulse": return reflect.ValueOf(self._body_apply_torque_impulse);
	case "_body_apply_impulse": return reflect.ValueOf(self._body_apply_impulse);
	case "_body_apply_central_force": return reflect.ValueOf(self._body_apply_central_force);
	case "_body_apply_force": return reflect.ValueOf(self._body_apply_force);
	case "_body_apply_torque": return reflect.ValueOf(self._body_apply_torque);
	case "_body_add_constant_central_force": return reflect.ValueOf(self._body_add_constant_central_force);
	case "_body_add_constant_force": return reflect.ValueOf(self._body_add_constant_force);
	case "_body_add_constant_torque": return reflect.ValueOf(self._body_add_constant_torque);
	case "_body_set_constant_force": return reflect.ValueOf(self._body_set_constant_force);
	case "_body_get_constant_force": return reflect.ValueOf(self._body_get_constant_force);
	case "_body_set_constant_torque": return reflect.ValueOf(self._body_set_constant_torque);
	case "_body_get_constant_torque": return reflect.ValueOf(self._body_get_constant_torque);
	case "_body_set_axis_velocity": return reflect.ValueOf(self._body_set_axis_velocity);
	case "_body_add_collision_exception": return reflect.ValueOf(self._body_add_collision_exception);
	case "_body_remove_collision_exception": return reflect.ValueOf(self._body_remove_collision_exception);
	case "_body_get_collision_exceptions": return reflect.ValueOf(self._body_get_collision_exceptions);
	case "_body_set_max_contacts_reported": return reflect.ValueOf(self._body_set_max_contacts_reported);
	case "_body_get_max_contacts_reported": return reflect.ValueOf(self._body_get_max_contacts_reported);
	case "_body_set_contacts_reported_depth_threshold": return reflect.ValueOf(self._body_set_contacts_reported_depth_threshold);
	case "_body_get_contacts_reported_depth_threshold": return reflect.ValueOf(self._body_get_contacts_reported_depth_threshold);
	case "_body_set_omit_force_integration": return reflect.ValueOf(self._body_set_omit_force_integration);
	case "_body_is_omitting_force_integration": return reflect.ValueOf(self._body_is_omitting_force_integration);
	case "_body_set_state_sync_callback": return reflect.ValueOf(self._body_set_state_sync_callback);
	case "_body_set_force_integration_callback": return reflect.ValueOf(self._body_set_force_integration_callback);
	case "_body_collide_shape": return reflect.ValueOf(self._body_collide_shape);
	case "_body_set_pickable": return reflect.ValueOf(self._body_set_pickable);
	case "_body_get_direct_state": return reflect.ValueOf(self._body_get_direct_state);
	case "_body_test_motion": return reflect.ValueOf(self._body_test_motion);
	case "_joint_create": return reflect.ValueOf(self._joint_create);
	case "_joint_clear": return reflect.ValueOf(self._joint_clear);
	case "_joint_set_param": return reflect.ValueOf(self._joint_set_param);
	case "_joint_get_param": return reflect.ValueOf(self._joint_get_param);
	case "_joint_disable_collisions_between_bodies": return reflect.ValueOf(self._joint_disable_collisions_between_bodies);
	case "_joint_is_disabled_collisions_between_bodies": return reflect.ValueOf(self._joint_is_disabled_collisions_between_bodies);
	case "_joint_make_pin": return reflect.ValueOf(self._joint_make_pin);
	case "_joint_make_groove": return reflect.ValueOf(self._joint_make_groove);
	case "_joint_make_damped_spring": return reflect.ValueOf(self._joint_make_damped_spring);
	case "_pin_joint_set_flag": return reflect.ValueOf(self._pin_joint_set_flag);
	case "_pin_joint_get_flag": return reflect.ValueOf(self._pin_joint_get_flag);
	case "_pin_joint_set_param": return reflect.ValueOf(self._pin_joint_set_param);
	case "_pin_joint_get_param": return reflect.ValueOf(self._pin_joint_get_param);
	case "_damped_spring_joint_set_param": return reflect.ValueOf(self._damped_spring_joint_set_param);
	case "_damped_spring_joint_get_param": return reflect.ValueOf(self._damped_spring_joint_get_param);
	case "_joint_get_type": return reflect.ValueOf(self._joint_get_type);
	case "_free_rid": return reflect.ValueOf(self._free_rid);
	case "_set_active": return reflect.ValueOf(self._set_active);
	case "_init": return reflect.ValueOf(self._init);
	case "_step": return reflect.ValueOf(self._step);
	case "_sync": return reflect.ValueOf(self._sync);
	case "_flush_queries": return reflect.ValueOf(self._flush_queries);
	case "_end_sync": return reflect.ValueOf(self._end_sync);
	case "_finish": return reflect.ValueOf(self._finish);
	case "_is_flushing_queries": return reflect.ValueOf(self._is_flushing_queries);
	case "_get_process_info": return reflect.ValueOf(self._get_process_info);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_world_boundary_shape_create": return reflect.ValueOf(self._world_boundary_shape_create);
	case "_separation_ray_shape_create": return reflect.ValueOf(self._separation_ray_shape_create);
	case "_segment_shape_create": return reflect.ValueOf(self._segment_shape_create);
	case "_circle_shape_create": return reflect.ValueOf(self._circle_shape_create);
	case "_rectangle_shape_create": return reflect.ValueOf(self._rectangle_shape_create);
	case "_capsule_shape_create": return reflect.ValueOf(self._capsule_shape_create);
	case "_convex_polygon_shape_create": return reflect.ValueOf(self._convex_polygon_shape_create);
	case "_concave_polygon_shape_create": return reflect.ValueOf(self._concave_polygon_shape_create);
	case "_shape_set_data": return reflect.ValueOf(self._shape_set_data);
	case "_shape_set_custom_solver_bias": return reflect.ValueOf(self._shape_set_custom_solver_bias);
	case "_shape_get_type": return reflect.ValueOf(self._shape_get_type);
	case "_shape_get_data": return reflect.ValueOf(self._shape_get_data);
	case "_shape_get_custom_solver_bias": return reflect.ValueOf(self._shape_get_custom_solver_bias);
	case "_shape_collide": return reflect.ValueOf(self._shape_collide);
	case "_space_create": return reflect.ValueOf(self._space_create);
	case "_space_set_active": return reflect.ValueOf(self._space_set_active);
	case "_space_is_active": return reflect.ValueOf(self._space_is_active);
	case "_space_set_param": return reflect.ValueOf(self._space_set_param);
	case "_space_get_param": return reflect.ValueOf(self._space_get_param);
	case "_space_get_direct_state": return reflect.ValueOf(self._space_get_direct_state);
	case "_space_set_debug_contacts": return reflect.ValueOf(self._space_set_debug_contacts);
	case "_space_get_contacts": return reflect.ValueOf(self._space_get_contacts);
	case "_space_get_contact_count": return reflect.ValueOf(self._space_get_contact_count);
	case "_area_create": return reflect.ValueOf(self._area_create);
	case "_area_set_space": return reflect.ValueOf(self._area_set_space);
	case "_area_get_space": return reflect.ValueOf(self._area_get_space);
	case "_area_add_shape": return reflect.ValueOf(self._area_add_shape);
	case "_area_set_shape": return reflect.ValueOf(self._area_set_shape);
	case "_area_set_shape_transform": return reflect.ValueOf(self._area_set_shape_transform);
	case "_area_set_shape_disabled": return reflect.ValueOf(self._area_set_shape_disabled);
	case "_area_get_shape_count": return reflect.ValueOf(self._area_get_shape_count);
	case "_area_get_shape": return reflect.ValueOf(self._area_get_shape);
	case "_area_get_shape_transform": return reflect.ValueOf(self._area_get_shape_transform);
	case "_area_remove_shape": return reflect.ValueOf(self._area_remove_shape);
	case "_area_clear_shapes": return reflect.ValueOf(self._area_clear_shapes);
	case "_area_attach_object_instance_id": return reflect.ValueOf(self._area_attach_object_instance_id);
	case "_area_get_object_instance_id": return reflect.ValueOf(self._area_get_object_instance_id);
	case "_area_attach_canvas_instance_id": return reflect.ValueOf(self._area_attach_canvas_instance_id);
	case "_area_get_canvas_instance_id": return reflect.ValueOf(self._area_get_canvas_instance_id);
	case "_area_set_param": return reflect.ValueOf(self._area_set_param);
	case "_area_set_transform": return reflect.ValueOf(self._area_set_transform);
	case "_area_get_param": return reflect.ValueOf(self._area_get_param);
	case "_area_get_transform": return reflect.ValueOf(self._area_get_transform);
	case "_area_set_collision_layer": return reflect.ValueOf(self._area_set_collision_layer);
	case "_area_get_collision_layer": return reflect.ValueOf(self._area_get_collision_layer);
	case "_area_set_collision_mask": return reflect.ValueOf(self._area_set_collision_mask);
	case "_area_get_collision_mask": return reflect.ValueOf(self._area_get_collision_mask);
	case "_area_set_monitorable": return reflect.ValueOf(self._area_set_monitorable);
	case "_area_set_pickable": return reflect.ValueOf(self._area_set_pickable);
	case "_area_set_monitor_callback": return reflect.ValueOf(self._area_set_monitor_callback);
	case "_area_set_area_monitor_callback": return reflect.ValueOf(self._area_set_area_monitor_callback);
	case "_body_create": return reflect.ValueOf(self._body_create);
	case "_body_set_space": return reflect.ValueOf(self._body_set_space);
	case "_body_get_space": return reflect.ValueOf(self._body_get_space);
	case "_body_set_mode": return reflect.ValueOf(self._body_set_mode);
	case "_body_get_mode": return reflect.ValueOf(self._body_get_mode);
	case "_body_add_shape": return reflect.ValueOf(self._body_add_shape);
	case "_body_set_shape": return reflect.ValueOf(self._body_set_shape);
	case "_body_set_shape_transform": return reflect.ValueOf(self._body_set_shape_transform);
	case "_body_get_shape_count": return reflect.ValueOf(self._body_get_shape_count);
	case "_body_get_shape": return reflect.ValueOf(self._body_get_shape);
	case "_body_get_shape_transform": return reflect.ValueOf(self._body_get_shape_transform);
	case "_body_set_shape_disabled": return reflect.ValueOf(self._body_set_shape_disabled);
	case "_body_set_shape_as_one_way_collision": return reflect.ValueOf(self._body_set_shape_as_one_way_collision);
	case "_body_remove_shape": return reflect.ValueOf(self._body_remove_shape);
	case "_body_clear_shapes": return reflect.ValueOf(self._body_clear_shapes);
	case "_body_attach_object_instance_id": return reflect.ValueOf(self._body_attach_object_instance_id);
	case "_body_get_object_instance_id": return reflect.ValueOf(self._body_get_object_instance_id);
	case "_body_attach_canvas_instance_id": return reflect.ValueOf(self._body_attach_canvas_instance_id);
	case "_body_get_canvas_instance_id": return reflect.ValueOf(self._body_get_canvas_instance_id);
	case "_body_set_continuous_collision_detection_mode": return reflect.ValueOf(self._body_set_continuous_collision_detection_mode);
	case "_body_get_continuous_collision_detection_mode": return reflect.ValueOf(self._body_get_continuous_collision_detection_mode);
	case "_body_set_collision_layer": return reflect.ValueOf(self._body_set_collision_layer);
	case "_body_get_collision_layer": return reflect.ValueOf(self._body_get_collision_layer);
	case "_body_set_collision_mask": return reflect.ValueOf(self._body_set_collision_mask);
	case "_body_get_collision_mask": return reflect.ValueOf(self._body_get_collision_mask);
	case "_body_set_collision_priority": return reflect.ValueOf(self._body_set_collision_priority);
	case "_body_get_collision_priority": return reflect.ValueOf(self._body_get_collision_priority);
	case "_body_set_param": return reflect.ValueOf(self._body_set_param);
	case "_body_get_param": return reflect.ValueOf(self._body_get_param);
	case "_body_reset_mass_properties": return reflect.ValueOf(self._body_reset_mass_properties);
	case "_body_set_state": return reflect.ValueOf(self._body_set_state);
	case "_body_get_state": return reflect.ValueOf(self._body_get_state);
	case "_body_apply_central_impulse": return reflect.ValueOf(self._body_apply_central_impulse);
	case "_body_apply_torque_impulse": return reflect.ValueOf(self._body_apply_torque_impulse);
	case "_body_apply_impulse": return reflect.ValueOf(self._body_apply_impulse);
	case "_body_apply_central_force": return reflect.ValueOf(self._body_apply_central_force);
	case "_body_apply_force": return reflect.ValueOf(self._body_apply_force);
	case "_body_apply_torque": return reflect.ValueOf(self._body_apply_torque);
	case "_body_add_constant_central_force": return reflect.ValueOf(self._body_add_constant_central_force);
	case "_body_add_constant_force": return reflect.ValueOf(self._body_add_constant_force);
	case "_body_add_constant_torque": return reflect.ValueOf(self._body_add_constant_torque);
	case "_body_set_constant_force": return reflect.ValueOf(self._body_set_constant_force);
	case "_body_get_constant_force": return reflect.ValueOf(self._body_get_constant_force);
	case "_body_set_constant_torque": return reflect.ValueOf(self._body_set_constant_torque);
	case "_body_get_constant_torque": return reflect.ValueOf(self._body_get_constant_torque);
	case "_body_set_axis_velocity": return reflect.ValueOf(self._body_set_axis_velocity);
	case "_body_add_collision_exception": return reflect.ValueOf(self._body_add_collision_exception);
	case "_body_remove_collision_exception": return reflect.ValueOf(self._body_remove_collision_exception);
	case "_body_get_collision_exceptions": return reflect.ValueOf(self._body_get_collision_exceptions);
	case "_body_set_max_contacts_reported": return reflect.ValueOf(self._body_set_max_contacts_reported);
	case "_body_get_max_contacts_reported": return reflect.ValueOf(self._body_get_max_contacts_reported);
	case "_body_set_contacts_reported_depth_threshold": return reflect.ValueOf(self._body_set_contacts_reported_depth_threshold);
	case "_body_get_contacts_reported_depth_threshold": return reflect.ValueOf(self._body_get_contacts_reported_depth_threshold);
	case "_body_set_omit_force_integration": return reflect.ValueOf(self._body_set_omit_force_integration);
	case "_body_is_omitting_force_integration": return reflect.ValueOf(self._body_is_omitting_force_integration);
	case "_body_set_state_sync_callback": return reflect.ValueOf(self._body_set_state_sync_callback);
	case "_body_set_force_integration_callback": return reflect.ValueOf(self._body_set_force_integration_callback);
	case "_body_collide_shape": return reflect.ValueOf(self._body_collide_shape);
	case "_body_set_pickable": return reflect.ValueOf(self._body_set_pickable);
	case "_body_get_direct_state": return reflect.ValueOf(self._body_get_direct_state);
	case "_body_test_motion": return reflect.ValueOf(self._body_test_motion);
	case "_joint_create": return reflect.ValueOf(self._joint_create);
	case "_joint_clear": return reflect.ValueOf(self._joint_clear);
	case "_joint_set_param": return reflect.ValueOf(self._joint_set_param);
	case "_joint_get_param": return reflect.ValueOf(self._joint_get_param);
	case "_joint_disable_collisions_between_bodies": return reflect.ValueOf(self._joint_disable_collisions_between_bodies);
	case "_joint_is_disabled_collisions_between_bodies": return reflect.ValueOf(self._joint_is_disabled_collisions_between_bodies);
	case "_joint_make_pin": return reflect.ValueOf(self._joint_make_pin);
	case "_joint_make_groove": return reflect.ValueOf(self._joint_make_groove);
	case "_joint_make_damped_spring": return reflect.ValueOf(self._joint_make_damped_spring);
	case "_pin_joint_set_flag": return reflect.ValueOf(self._pin_joint_set_flag);
	case "_pin_joint_get_flag": return reflect.ValueOf(self._pin_joint_get_flag);
	case "_pin_joint_set_param": return reflect.ValueOf(self._pin_joint_set_param);
	case "_pin_joint_get_param": return reflect.ValueOf(self._pin_joint_get_param);
	case "_damped_spring_joint_set_param": return reflect.ValueOf(self._damped_spring_joint_set_param);
	case "_damped_spring_joint_get_param": return reflect.ValueOf(self._damped_spring_joint_get_param);
	case "_joint_get_type": return reflect.ValueOf(self._joint_get_type);
	case "_free_rid": return reflect.ValueOf(self._free_rid);
	case "_set_active": return reflect.ValueOf(self._set_active);
	case "_init": return reflect.ValueOf(self._init);
	case "_step": return reflect.ValueOf(self._step);
	case "_sync": return reflect.ValueOf(self._sync);
	case "_flush_queries": return reflect.ValueOf(self._flush_queries);
	case "_end_sync": return reflect.ValueOf(self._end_sync);
	case "_finish": return reflect.ValueOf(self._finish);
	case "_is_flushing_queries": return reflect.ValueOf(self._is_flushing_queries);
	case "_get_process_info": return reflect.ValueOf(self._get_process_info);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PhysicsServer2DExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
