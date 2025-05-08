// Package Particle provides a particle shader pipeline used for shading 2D and 3D particles.
package Particle

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/internal"
	"graphics.gd/shaders/internal/gpu"
	"graphics.gd/shaders/mat4"
	"graphics.gd/shaders/uint"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

/*
Shader that runs before the object is drawn. They are used for calculating material properties such as color, position, and rotation.
They are drawn with any regular material for CanvasItem or Spatial, depending on whether they are 2D or 3D.

Particle shaders are unique because they are not used to draw the object itself; they are used to calculate particle properties,
which are then used by a CanvasItem or Spatial shader. They contain two processor functions: start() and process().

Unlike other shader types, particle shaders keep the data that was output the previous frame. Therefore, particle shaders can be used for complex effects that take place over multiple frames.
*/
type Shader struct {
	shader
}

type shader = internal.Shader

func (Shader) ShaderType() string       { return "particle" }
func (Shader) RenderMode() []RenderMode { return nil }
func (Shader) Pipeline() [3]string {
	return [3]string{"start", "process", ""}
}

func (Shader) Fragment(state Startup) Process { return Process{} }
func (Shader) Material(state Process) State   { return State{} }
func (Shader) Lighting(Process) struct{}      { return struct{}{} }

type RenderMode string

const (
	KeepData          RenderMode = "keep_data"           // Do not clear previous data on restart.
	DisableForce      RenderMode = "disable_force"       // Disable attractor force.
	DisableVelocity   RenderMode = "disable_velocity"    // Ignore VELOCITY value.
	CollisionUseScale RenderMode = "collision_use_scale" // Scale the particle's size for collisions.
)

type State struct {
	Active    bool.X           `gd:"ACTIVE"`    // true when the particle is active, can be set false.
	Color     vec4.RGBA        `gd:"COLOR"`     // Particle color, can be written to and accessed in mesh's vertex function.
	Velocity  vec3.XYZ         `gd:"VELOCITY"`  // Particle velocity, can be modified.
	Transform mat4.ColumnMajor `gd:"TRANSFORM"` // Particle transform.
	Custom    vec4.XYZW        `gd:"CUSTOM"`    // Custom particle data. Accessible from shader of mesh as INSTANCE_CUSTOM.
	Mass      float.X          `gd:"MASS"`      // Particle mass, intended to be used with attractors. Equals 1.0 by default.

	EmitSubParticle func(xform mat4.ColumnMajor, velocity vec3.XYZ, color vec4.RGBA, custom vec4.XYZW, flags uint.X) bool.X `gd:"emit_subparticle"`
}

// Lifetime of the particle
func (State) Lifetime() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("LIFETIME")))
}

// Delta process time.
func (State) Delta() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("DELTA")))
}

// Number is a unique number since emission start.
func (State) Number() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("NUMBER")))
}

// Index (from total particles).
func (State) Index() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("INDEX")))
}

// EmissionTransform (used for non-local systems).
func (State) EmissionTransform() mat4.ColumnMajor {
	return gpu.NewMat4Expression(gpu.New(gpu.Identifier("EMISSION_TRANSFORM")))
}

// RandomSeed used as base for random.
func (State) RandomSeed() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("RANDOM_SEED")))
}

// UserData is user data passed to the shader.
func (State) UserData() [6]vec4.XYZW {
	return [6]vec4.XYZW{
		gpu.NewVec4Expression(gpu.New(gpu.Identifier("USERDATA1"))),
		gpu.NewVec4Expression(gpu.New(gpu.Identifier("USERDATA2"))),
		gpu.NewVec4Expression(gpu.New(gpu.Identifier("USERDATA3"))),
		gpu.NewVec4Expression(gpu.New(gpu.Identifier("USERDATA4"))),
		gpu.NewVec4Expression(gpu.New(gpu.Identifier("USERDATA5"))),
		gpu.NewVec4Expression(gpu.New(gpu.Identifier("USERDATA6"))),
	}
}

// FlagEmitPosition for using on the last argument of emit_subparticle() function to assign a
// position to a new particle's transform.
func (State) FlagEmitPosition() uint.X {
	return gpu.NewUintExpression(gpu.New(gpu.Identifier("FLAG_EMIT_POSITION")))
}

// FlagEmitRotScale for using on the last argument of emit_subparticle() function to assign a
// rotation and scale to a new particle's transform.
func (State) FlagEmitRotScale() uint.X {
	return gpu.NewUintExpression(gpu.New(gpu.Identifier("FLAG_EMIT_ROT_SCALE")))
}

// FlagEmitVelocity for using on the last argument of emit_subparticle() function to assign a
// velocity to a new particle.
func (State) FlagEmitVelocity() uint.X {
	return gpu.NewUintExpression(gpu.New(gpu.Identifier("FLAG_EMIT_VELOCITY")))
}

// FlagEmitColor for using on the last argument of emit_subparticle() function to assign a
// color to a new particle.
func (State) FlagEmitColor() uint.X {
	return gpu.NewUintExpression(gpu.New(gpu.Identifier("FLAG_EMIT_COLOR")))
}

// FlagEmitCustom for using on the last argument of emit_subparticle() function to assign a
// custom value to a new particle.
func (State) FlagEmitCustom() uint.X {
	return gpu.NewUintExpression(gpu.New(gpu.Identifier("FLAG_EMIT_CUSTOM")))
}

// EmitterVelocity is the Velocity of the Particles2D (3D) node.
func (State) EmitterVelocity() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("EMITTER_VELOCITY")))
}

// InterpolateToEnd Value of interp_to_end (3D) property of Particles node.
func (State) InterpolateToEnd() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("INTERPOLATE_TO_END")))
}

// AmountRatio value of amount_ratio (3D) property of Particles node.
func (State) AmountRatio() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("AMOUNT_RATIO")))
}

type Startup State

// RestartPosition is true if particle is restarted, or emitted without a custom position (i.e. this particle was created by
// emit_subparticle() without the FLAG_EMIT_POSITION flag).
func (state Startup) RestartPosition() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("RESTART_POSITION")))
}

// RestartRotScale is true if particle is restarted, or emitted without a custom rotation or scale (i.e. this particle was
// created by emit_subparticle() without the FLAG_EMIT_ROT_SCALE flag).
func (state Startup) RestartRotScale() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("RESTART_ROT_SCALE")))
}

// RestartVelocity is true if particle is restarted, or emitted without a custom velocity (i.e. this particle was created by
// emit_subparticle() without the FLAG_EMIT_VELOCITY flag).
func (state Startup) RestartVelocity() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("RESTART_VELOCITY")))
}

// RestartColor is true if particle is restarted, or emitted without a custom color (i.e. this particle was created by
// emit_subparticle() without the FLAG_EMIT_COLOR flag).
func (state Startup) RestartColor() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("RESTART_COLOR")))
}

// RestartCustom is true if particle is restarted, or emitted without a custom property (i.e. this particle was created by
// emit_subparticle() without the FLAG_EMIT_CUSTOM flag).
func (state Startup) RestartCustom() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("RESTART_CUSTOM")))
}

type Process State

// Restart is true if the current process frame is first for the particle.
func (state Process) Restart() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("RESTART")))
}

// Collided is true when the particle has collided with a particle collider.
func (state Process) Collided() bool.X {
	return gpu.NewBoolExpression(gpu.New(gpu.Identifier("COLLIDED")))
}

// CollisionNormal of the last collision. If there is no collision detected it is equal to (0.0, 0.0, 0.0).
func (state Process) CollisionNormal() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("COLLISION_NORMAL")))
}

// CollisionDepth is the length of the normal of the last collision. If there is no collision detected it is
// equal to 0.0.
func (state Process) CollisionDepth() float.X {
	return gpu.NewFloatExpression(gpu.New(gpu.Identifier("COLLISION_DEPTH")))
}

// AttractorForce is the combined force of the attractors at the moment on that particle.
func (state Process) AttractorForce() vec3.XYZ {
	return gpu.NewVec3Expression(gpu.New(gpu.Identifier("ATTRACTOR_FORCE")))
}
