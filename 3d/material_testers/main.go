package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Environment"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/OptionButton"
	"graphics.gd/classdb/PanoramaSkyMaterial"
	"graphics.gd/classdb/ReflectionProbe"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/classdb/WorldEnvironment"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

const (
	InterpSpeed = 2
	RotSpeed    = 0.003
	ZoomSpeed   = 0.1
	MainButtons = Input.MouseButtonMaskLeft | Input.MouseButtonMaskMiddle | Input.MouseButtonMaskRight
)

type background struct {
	path string
	name string
}

var backgrounds = []background{
	{path: "res://backgrounds/schelde.hdr", name: "Riverside"},
	{path: "res://backgrounds/lobby.hdr", name: "Lobby"},
	{path: "res://backgrounds/park.hdr", name: "Park"},
	{path: "res://backgrounds/night.hdr", name: "Night"},
	{path: "res://backgrounds/experiment.hdr", name: "Experiment"},
}

// MaterialTester lets the user inspect a series of material test beds.
type MaterialTester struct {
	Node3D.Extension[MaterialTester] `gd:"MaterialTester"`

	Testers      Node3D.Instance
	MaterialName Label.Instance `gd:"UI/MaterialName"`

	CameraHolder Node3D.Instance   // Has a position and rotates on Y.
	RotationX    Node3D.Instance   `gd:"CameraHolder/RotationX"`
	Camera       Camera3D.Instance `gd:"CameraHolder/RotationX/Camera"`

	WorldEnvironment WorldEnvironment.Instance
	Background       OptionButton.Instance `gd:"UI/Background"`
	Previous         Button.Instance       `gd:"UI/Previous"`
	Next             Button.Instance       `gd:"UI/Next"`

	testerIndex int
	rotX        Float.X
	rotY        Float.X
	zoom        Float.X
}

func NewMaterialTester() *MaterialTester {
	return &MaterialTester{
		rotX: -0.5, // This must be kept in sync with RotationX.
		rotY: -0.5, // This must be kept in sync with CameraHolder.
		zoom: 5.0,
	}
}

func (m *MaterialTester) Ready() {
	if RenderingServer.GetCurrentRenderingMethod() == "gl_compatibility" {
		// Tweak scene brightness to better match Forward+/Mobile.
		environment := m.WorldEnvironment.Environment()
		environment.SetTonemapMode(Environment.ToneMapperFilmic)
		environment.SetBackgroundEnergyMultiplier(2.0)
	}

	for _, background := range backgrounds {
		m.Background.AddItem(background.name)
	}

	m.UpdateGui()
}

func (m *MaterialTester) UnhandledInput(inputEvent InputEvent.Instance) {
	if inputEvent.IsActionPressed("ui_left") {
		m.OnPreviousPressed()
	}
	if inputEvent.IsActionPressed("ui_right") {
		m.OnNextPressed()
	}

	if button, ok := Object.As[InputEventMouseButton.Instance](inputEvent); ok {
		if button.ButtonIndex() == Input.MouseButtonWheelUp {
			m.zoom -= ZoomSpeed
		}
		if button.ButtonIndex() == Input.MouseButtonWheelDown {
			m.zoom += ZoomSpeed
		}
		m.zoom = Float.Clamp(m.zoom, 2, 8)
		position := m.Camera.AsNode3D().Position()
		position.Z = m.zoom
		m.Camera.AsNode3D().SetPosition(position)
	}

	if motion, ok := Object.As[InputEventMouseMotion.Instance](inputEvent); ok && motion.AsInputEventMouse().ButtonMask()&MainButtons != 0 {
		// Use `screen_relative` to make mouse sensitivity independent of viewport resolution.
		relativeMotion := motion.ScreenRelative()
		m.rotY -= relativeMotion.X * RotSpeed
		m.rotY = Float.Clamp(m.rotY, -1.95, 1.95)
		m.rotX -= relativeMotion.Y * RotSpeed
		m.rotX = Float.Clamp(m.rotX, -1.4, 0.45)
		holder := m.CameraHolder.Transform()
		holder.Basis = Basis.FromEuler(Euler.Radians{Y: Angle.Radians(m.rotY)}, Angle.OrderYXZ)
		m.CameraHolder.SetTransform(holder)
		rotationX := m.RotationX.Transform()
		rotationX.Basis = Basis.FromEuler(Euler.Radians{X: Angle.Radians(m.rotX)}, Angle.OrderYXZ)
		m.RotationX.SetTransform(rotationX)
	}
}

func (m *MaterialTester) Process(delta Float.X) {
	currentTester, _ := Object.As[Node3D.Instance](m.Testers.AsNode().GetChild(m.testerIndex))
	// This code assumes CameraHolder's Y and Z coordinates are already correct.
	targetPosition := currentTester.Transform().Origin.X
	holder := m.CameraHolder.Transform()
	holder.Origin.X = Float.Lerp(holder.Origin.X, targetPosition, InterpSpeed*delta)
	m.CameraHolder.SetTransform(holder)
}

func (m *MaterialTester) OnPreviousPressed() {
	if m.testerIndex > 0 {
		m.testerIndex--
	}
	m.UpdateGui()
}

func (m *MaterialTester) OnNextPressed() {
	if m.testerIndex < m.Testers.AsNode().GetChildCount()-1 {
		m.testerIndex++
	}
	m.UpdateGui()
}

func (m *MaterialTester) UpdateGui() {
	m.MaterialName.SetText(m.Testers.AsNode().GetChild(m.testerIndex).Name())
	m.Previous.AsBaseButton().SetDisabled(m.testerIndex == 0)
	m.Next.AsBaseButton().SetDisabled(m.testerIndex == m.Testers.AsNode().GetChildCount()-1)
}

func (m *MaterialTester) reflectionProbes(yield func(ReflectionProbe.Instance)) {
	for _, node := range SceneTree.Get(m.AsNode()).GetNodesInGroup("reflection_probe") {
		if probe, ok := Object.As[ReflectionProbe.Instance](node); ok {
			yield(probe)
		}
	}
}

func (m *MaterialTester) OnBgItemSelected(index int) {
	skyMaterial, _ := Object.As[PanoramaSkyMaterial.Instance](m.WorldEnvironment.Environment().Sky().SkyMaterial())
	skyMaterial.SetPanorama(Resource.Load[Texture2D.Instance](backgrounds[index].path))

	// Force reflection probes to update by moving them slightly.
	m.reflectionProbes(func(probe ReflectionProbe.Instance) {
		position := probe.AsNode3D().Position()
		position.Y += Float.RandomBetween(-0.0001, 0.0001)
		probe.AsNode3D().SetPosition(position)
	})
}

func (m *MaterialTester) OnReflectionProbesItemSelected(index int) {
	switch index {
	case 0: // No Reflection Probes
		m.reflectionProbes(func(probe ReflectionProbe.Instance) {
			probe.AsNode3D().SetVisible(false)
		})
	case 1: // Reflection Probes (Reflection only)
		m.reflectionProbes(func(probe ReflectionProbe.Instance) {
			probe.AsNode3D().SetVisible(true)
			probe.SetAmbientMode(ReflectionProbe.AmbientDisabled)
		})
	case 2: // Reflection Probes (Reflection + Ambient)
		m.reflectionProbes(func(probe ReflectionProbe.Instance) {
			probe.AsNode3D().SetVisible(true)
			probe.SetAmbientMode(ReflectionProbe.AmbientEnvironment)
		})
	}
}

func (m *MaterialTester) OnQuitPressed() {
	SceneTree.Get(m.AsNode()).Quit()
}

func main() {
	classdb.Register[MaterialTester](NewMaterialTester)
	startup.Scene()
}
