package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/HBoxContainer"
	"graphics.gd/classdb/HSlider"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/OptionButton"
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/ShaderMaterial"
	"graphics.gd/classdb/Sky"
	"graphics.gd/classdb/SphereMesh"
	"graphics.gd/classdb/StandardMaterial3D"
	"graphics.gd/classdb/WorldEnvironment"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

const MouseSensitivity = 0.001

// SkyShaders demonstrates a volumetric cloud sky shader with a day/night cycle.
type SkyShaders struct {
	Node3D.Extension[SkyShaders] `gd:"SkyShaders"`

	AnimationPlayer  AnimationPlayer.Instance
	Panel            Panel.Instance
	Help             Label.Instance
	Spheres          *Spheres
	WorldEnvironment WorldEnvironment.Instance
	YawCamera        Node3D.Instance
	Camera           Camera3D.Instance `gd:"YawCamera/Camera3D"`

	TimeOfDaySlider HSlider.Instance       `gd:"Panel/MarginContainer/VBoxContainer/TimeOfDay/HSlider"`
	TimeOfDayValue  Label.Instance         `gd:"Panel/MarginContainer/VBoxContainer/TimeOfDay/Value"`
	CurrentSpeed    Label.Instance         `gd:"Panel/MarginContainer/VBoxContainer/TimeOfDay/CurrentSpeed"`
	CoverageValue   Label.Instance         `gd:"Panel/MarginContainer/VBoxContainer/Clouds/CoverageValue"`
	DensityValue    Label.Instance         `gd:"Panel/MarginContainer/VBoxContainer/Clouds/DensityValue"`
	RadianceSizeRow HBoxContainer.Instance `gd:"Panel/MarginContainer/VBoxContainer/RadianceSize"`
	RadianceSize    OptionButton.Instance  `gd:"Panel/MarginContainer/VBoxContainer/RadianceSize/OptionButton"`

	// The camera field of view to smoothly interpolate to.
	desiredFov Float.X
}

func (s *SkyShaders) sky() Sky.Instance {
	return s.WorldEnvironment.Environment().Sky()
}

func (s *SkyShaders) skyMaterial() ShaderMaterial.Instance {
	material, _ := Object.As[ShaderMaterial.Instance](s.sky().SkyMaterial())
	return material
}

func (s *SkyShaders) Ready() {
	s.desiredFov = s.Camera.Fov()
	Input.SetMouseMode(Input.MouseModeCaptured)
}

func (s *SkyShaders) Process(delta Float.X) {
	// Make the slider follow the day/night cycle.
	s.TimeOfDaySlider.AsRange().SetValue(s.AnimationPlayer.CurrentAnimationPosition())
	s.skyMaterial().SetShaderParameter("cloud_time_offset", s.AnimationPlayer.CurrentAnimationPosition())

	s.Camera.SetFov(Float.Lerp(s.Camera.Fov(), s.desiredFov, 1.0-Float.Exp(-delta*10.0)))
}

func (s *SkyShaders) Input(inputEvent InputEvent.Instance) {
	if inputEvent.IsActionPressed("toggle_gui") {
		s.Panel.AsCanvasItem().SetVisible(!s.Panel.AsCanvasItem().Visible())
		s.Help.AsCanvasItem().SetVisible(!s.Help.AsCanvasItem().Visible())
	}

	if inputEvent.IsActionPressed("toggle_spheres") {
		s.Spheres.AsNode3D().SetVisible(!s.Spheres.AsNode3D().Visible())
	}

	if inputEvent.IsActionPressed("toggle_mouse_capture") {
		if Input.MouseMode() == Input.MouseModeCaptured {
			Input.SetMouseMode(Input.MouseModeVisible)
		} else {
			Input.SetMouseMode(Input.MouseModeCaptured)
		}
	}

	if motion, ok := Object.As[InputEventMouseMotion.Instance](inputEvent); ok && Input.MouseMode() == Input.MouseModeCaptured {
		// Mouselook.
		relativeMotion := motion.ScreenRelative()
		rotation := s.YawCamera.Rotation()
		rotation.X = Float.Clamp(rotation.X-Angle.Radians(relativeMotion.Y*MouseSensitivity), -Angle.Tau*0.25, Angle.Tau*0.25)
		rotation.Y -= Angle.Radians(relativeMotion.X * MouseSensitivity)
		s.YawCamera.SetRotation(rotation)
	}

	// Mouse wheel currently doesn't work in input actions. Hardcode mouse wheel as a workaround.
	if inputEvent.IsActionPressed("increase_camera_fov") || Input.IsMouseButtonPressed(Input.MouseButtonWheelDown) {
		s.desiredFov = Float.Clamp(s.desiredFov+5.0, 20.0, 120.0)
	}
	if inputEvent.IsActionPressed("decrease_camera_fov") || Input.IsMouseButtonPressed(Input.MouseButtonWheelUp) {
		s.desiredFov = Float.Clamp(s.desiredFov-5.0, 20.0, 120.0)
	}
}

func (s *SkyShaders) OnTimeOfDayValueChanged(value Float.X) {
	// Update time of day.
	s.AnimationPlayer.SeekTo(value)
	s.TimeOfDayValue.SetText(fmt.Sprintf("%.2f", value))
}

func (s *SkyShaders) OnSpeedMinusPressed() {
	// Allow minimum value to be zero so that the special case can be reached for pausing.
	s.AnimationPlayer.SetSpeedScale(Float.Clamp(s.AnimationPlayer.SpeedScale()*0.5, 0.0, 12.8))
	if s.AnimationPlayer.SpeedScale() < 0.0499 {
		// Going below 0.5× speed; pause.
		s.AnimationPlayer.SetSpeedScale(0.0)
	}
	s.updateSpeedLabel()
}

func (s *SkyShaders) OnSpeedPlusPressed() {
	s.AnimationPlayer.SetSpeedScale(Float.Clamp(s.AnimationPlayer.SpeedScale()*2.0, 0.05, 12.8))
	if Float.IsApproximatelyZero(s.AnimationPlayer.SpeedScale()) {
		// Currently paused; resume playback.
		s.AnimationPlayer.SetSpeedScale(0.1)
	}
	s.updateSpeedLabel()
}

func (s *SkyShaders) updateSpeedLabel() {
	// The default speed scale for this AnimationPlayer is internally 0.1, so multiply the displayed value by 10.
	if Float.IsApproximatelyZero(s.AnimationPlayer.SpeedScale()) {
		s.CurrentSpeed.SetText("Pause")
	} else {
		s.CurrentSpeed.SetText(fmt.Sprintf("%.2f×", s.AnimationPlayer.SpeedScale()*10))
	}
}

func (s *SkyShaders) OnCloudCoverageValueChanged(value Float.X) {
	s.skyMaterial().SetShaderParameter("cloud_coverage", value)
	s.CoverageValue.SetText(fmt.Sprintf("%d%%", int(value*100)))
}

func (s *SkyShaders) OnCloudDensityValueChanged(value Float.X) {
	s.skyMaterial().SetShaderParameter("cloud_density", value)
	s.DensityValue.SetText(fmt.Sprintf("%d%%", int(value*100)))
}

func (s *SkyShaders) OnProcessModeItemSelected(index int) {
	switch index {
	case 0:
		s.sky().SetProcessMode(Sky.ProcessModeQuality)
		s.RadianceSizeRow.AsCanvasItem().SetVisible(true)
		// Reset radiance size as the engine forces radiance size to 256 after switching to the Real-Time process mode.
		s.OnRadianceSizeItemSelected(s.RadianceSize.Selected())
	case 1:
		s.sky().SetProcessMode(Sky.ProcessModeIncremental)
		s.RadianceSizeRow.AsCanvasItem().SetVisible(true)
		// Reset radiance size as the engine forces radiance size to 256 after switching to the Real-Time process mode.
		s.OnRadianceSizeItemSelected(s.RadianceSize.Selected())
	case 2:
		s.sky().SetProcessMode(Sky.ProcessModeRealtime)
		// Radiance size is forced to 256 by the engine when using Real-Time process mode.
		s.RadianceSizeRow.AsCanvasItem().SetVisible(false)
	}
}

func (s *SkyShaders) OnRadianceSizeItemSelected(index int) {
	switch index {
	case 0:
		s.sky().SetRadianceSize(Sky.RadianceSize32)
	case 1:
		s.sky().SetRadianceSize(Sky.RadianceSize64)
	case 2:
		s.sky().SetRadianceSize(Sky.RadianceSize128)
	case 3:
		s.sky().SetRadianceSize(Sky.RadianceSize256)
	case 4:
		s.sky().SetRadianceSize(Sky.RadianceSize512)
	}
}

// Spheres creates spheres representing various levels of material roughness and metallic.
type Spheres struct {
	Node3D.Extension[Spheres] `gd:"Spheres"`
	classdb.Tool
}

func (s *Spheres) Ready() {
	// Create spheres to represent various levels of material roughness and metallic.
	for roughness := 0; roughness < 11; roughness++ {
		for metallic := 0; metallic < 11; metallic++ {
			sphere := MeshInstance3D.New()
			sphere.SetMesh(SphereMesh.New().AsMesh())
			// Center the spheres around the node origin.
			sphere.AsNode3D().SetPosition(Vector3.XYZ{Float.X(roughness) - 5, 0, Float.X(metallic) - 5})

			material := StandardMaterial3D.New()
			material.AsBaseMaterial3D().SetAlbedoColor(Color.RGBA{R: 0.5, G: 0.5, B: 0.5, A: 1})
			material.AsBaseMaterial3D().SetRoughness(Float.X(roughness) * 0.1)
			material.AsBaseMaterial3D().SetMetallic(Float.X(metallic) * 0.1)
			sphere.AsGeometryInstance3D().SetMaterialOverride(material.AsMaterial())

			s.AsNode().AddChild(sphere.AsNode())
		}
	}
}

func main() {
	classdb.Register[SkyShaders]()
	classdb.Register[Spheres]()
	classdb.Register[VolumetricClouds](NewVolumetricClouds)
	startup.Scene()
}
