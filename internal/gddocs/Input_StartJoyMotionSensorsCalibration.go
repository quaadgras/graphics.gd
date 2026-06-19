/*
[gdscript]
const GYRO_SENSITIVITY = 10.0

func _ready():
	# In this example we only use the first connected joypad (id 0).
	if 0 not in Input.get_connected_joypads():
		return

	if not Input.has_joy_motion_sensors(0):
		return

	# We must enable the motion sensors before using them.
	Input.set_joy_motion_sensors_enabled(0, true)

	# (Tell the users here that they need to put their joypads on a flat surface and wait for confirmation.)

	# Start the calibration process.
	calibrate_motion()

func _process(delta):
	# Only move the object if the joypad motion sensors are calibrated.
	if Input.is_joy_motion_sensors_calibrated(0):
		move_object(delta)

func calibrate_motion():
	Input.start_joy_motion_sensors_calibration(0)

	# Wait for some time.
	await get_tree().create_timer(1.0).timeout

	Input.stop_joy_motion_sensors_calibration(0)
	# The joypad is now calibrated.

func move_object(delta):
	var node: Node3D = ... # Put your node here.

	var gyro := Input.get_joy_gyroscope(0)
	node.rotation.x -= -gyro.y * GYRO_SENSITIVITY * delta # Use rotation around the Y axis (yaw) here.
	node.rotation.y += -gyro.x * GYRO_SENSITIVITY * delta # Use rotation around the X axis (pitch) here.
[/gdscript]
[csharp]
private const float GyroSensitivity = 10.0;

public override void _Ready()
{
	// In this example we only use the first connected joypad (id 0).
	if (!Input.GetConnectedJoypads().Contains(0))
	{
		return;
	}

	if (!Input.HasJoyMotionSensors(0))
	{
		return;
	}

	// We must enable the accelerometer and the gyroscope before using them.
	Input.SetJoyMotionSensorsEnabled(0, true);

	// (Tell the users here that they need to put their joypads on a flat surface and wait for confirmation.)

	// Start the calibration process.
	CalibrateMotion();
}

public override void _Process(double delta)
{
	// Only move the object if the joypad motion sensors are calibrated.
	if (Input.IsJoyMotionSensorsCalibrated(0))
	{
		MoveObject(delta);
	}
}

private async Task CalibrateMotion()
{
	Input.StartJoyMotionSensorsCalibration(0);

	// Wait for some time.
	await ToSignal(GetTree().CreateTimer(1.0), SceneTreeTimer.SignalName.Timeout);

	Input.StopJoyMotionSensorsCalibration(0);
	// The joypad is now calibrated.
}

private void MoveObject(double delta)
{
	Node3D node = ... ; // Put your object here.
	Vector3 gyro = Input.GetJoyGyroscope(0);
	Vector3 rotation = node.Rotation;
	rotation.X -= -gyro.Y * GyroSensitivity * (float)delta; // Use rotation around the Y axis (yaw) here.
	rotation.Y += -gyro.X * GyroSensitivity * (float)delta; // Use rotation around the X axis (pitch) here.
	node.Rotation = rotation;
}
[/csharp]
*/

package main

import (
	"slices"

	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
)

const gyroSensitivity = 10.0

type joyMotionExample struct {
	Node.Extension[joyMotionExample]

	Object Node3D.Instance // Put your node here.
}

func (n joyMotionExample) Ready() {
	// In this example we only use the first connected joypad (id 0).
	if !slices.Contains(Input.GetConnectedJoypads(), 0) {
		return
	}
	if !Input.HasJoyMotionSensors(0) {
		return
	}
	// We must enable the motion sensors before using them.
	Input.SetJoyMotionSensorsEnabled(0, true)
	// (Tell the users to put their joypads on a flat surface and wait for confirmation.)
	// Start the calibration process.
	n.calibrateMotion()
}

func (n joyMotionExample) Process(delta Float.X) {
	// Only move the object if the joypad motion sensors are calibrated.
	if Input.IsJoyMotionSensorsCalibrated(0) {
		n.moveObject(delta)
	}
}

func (n joyMotionExample) calibrateMotion() {
	Input.StartJoyMotionSensorsCalibration(0)
	// Wait for some time: await get_tree().create_timer(1.0).timeout
	Input.StopJoyMotionSensorsCalibration(0)
	// The joypad is now calibrated.
}

func (n joyMotionExample) moveObject(delta Float.X) {
	var gyro = Input.GetJoyGyroscope(0)
	var rotation = n.Object.Rotation()
	rotation.X -= Angle.Radians(-gyro.Y * gyroSensitivity * delta) // Use rotation around the Y axis (yaw).
	rotation.Y += Angle.Radians(-gyro.X * gyroSensitivity * delta) // Use rotation around the X axis (pitch).
	n.Object.SetRotation(rotation)
}
