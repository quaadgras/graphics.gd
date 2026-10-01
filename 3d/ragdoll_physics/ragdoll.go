package main

import (
	"graphics.gd/classdb/AudioStreamPlayer3D"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PhysicalBone3D"
	"graphics.gd/classdb/PhysicalBoneSimulator3D"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

const (
	ImpactSoundSpeedSmall = 0.3
	ImpactSoundSpeedBig   = 1.0
)

// MannequinyRagdoll wraps the mannequiny model (instanced as the "Model" child)
// with its ragdoll simulation and impact sounds.
type MannequinyRagdoll struct {
	Node3D.Extension[MannequinyRagdoll]

	// InitialVelocity is the velocity to apply on the first physics frame.
	InitialVelocity Vector3.XYZ

	Simulator PhysicalBoneSimulator3D.Instance `gd:"Model/root/root_001/Skeleton3D/PhysicalBoneSimulator3D"`
	// Used to play an impact sound on sudden velocity changes.
	// We use the pelvis bone as it's close to the center of mass of the character model.
	// For more detailed impact sounds, you could place multiple AudioStreamPlayer nodes as a child
	// of each limb.
	Pelvis           PhysicalBone3D.Instance `gd:"Model/root/root_001/Skeleton3D/PhysicalBoneSimulator3D/Physical Bone pelvis"`
	ImpactSoundBig   AudioStreamPlayer3D.Instance
	ImpactSoundSmall AudioStreamPlayer3D.Instance

	previousPelvisSpeed Float.X
}

func (m *MannequinyRagdoll) Ready() {
	m.Simulator.PhysicalBonesStartSimulation()
	if !Vector3.IsApproximatelyZero(m.InitialVelocity) {
		for _, child := range m.Simulator.AsNode().GetChildren() {
			// Give the ragdoll an initial motion by applying velocity on all its bones upon being spawned.
			if bone, ok := Object.As[PhysicalBone3D.Instance](child); ok {
				bone.ApplyCentralImpulse(m.InitialVelocity)
			}
		}
	}
}

func (m *MannequinyRagdoll) PhysicsProcess(_ Float.X) {
	var pelvisSpeed = Vector3.Length(m.Pelvis.LinearVelocity())
	// Ensure the speed used to determine the threshold doesn't change with time scale.
	var impactSpeed = (m.previousPelvisSpeed - pelvisSpeed) / Engine.TimeScale()
	if impactSpeed > ImpactSoundSpeedBig {
		m.ImpactSoundBig.Play()
	} else if impactSpeed > ImpactSoundSpeedSmall {
		m.ImpactSoundSmall.Play()
	}
	m.previousPelvisSpeed = pelvisSpeed
}
