package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/Area3D"
	"graphics.gd/classdb/AudioStreamPlayer3D"
	"graphics.gd/classdb/CharacterBody3D"
	"graphics.gd/classdb/CollisionShape3D"
	"graphics.gd/classdb/Curve"
	"graphics.gd/classdb/KinematicCollision3D"
	"graphics.gd/classdb/Marker3D"
	"graphics.gd/classdb/Mesh"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/PhysicsServer3D"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/ShapeCast3D"
	"graphics.gd/classdb/SurfaceTool"
	"graphics.gd/classdb/Timer"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
)

type MeleeAttackArea struct {
	Area3D.Extension[MeleeAttackArea] `gd:"MeleeAttackArea"`

	CollisionShape CollisionShape3D.Instance `gd:"CollisionShape3d"`
}

func (area MeleeAttackArea) Activate() {
	Callable.Defer(Callable.New(func() {
		area.CollisionShape.SetDisabled(false)
	}))
}

func (area MeleeAttackArea) Deactivate() {
	Callable.Defer(Callable.New(func() {
		area.CollisionShape.SetDisabled(true)
	}))
}

func (area MeleeAttackArea) Ready() {
	area.AsArea3D().OnBodyEntered(func(body Node3D.Instance) {
		if damagable, ok := Object.As[Damageable](body); ok {
			var impact_point = Vector3.Sub(area.AsNode3D().GlobalPosition(), body.GlobalPosition())
			var force = Vector3.Neg(impact_point)
			damagable.Damage(impact_point, force)
		}
	})
}

type Bullet struct {
	Node3D.Extension[Bullet] `gd:"Bullet"`

	ScaleDecay    Curve.Instance
	DistanceLimit Float.X

	Velocity Vector3.XYZ
	Shooter  Node.ID

	Area3D  Area3D.Instance `gd:"Area3D"`
	Visuals Node3D.Instance `gd:"Bullet"`

	ProjectileSound AudioStreamPlayer3D.Instance `gd:"ProjectileSound"`

	TimeAlive  Float.X
	AliveLimit Float.X
}

func NewBullet() *Bullet {
	return &Bullet{
		DistanceLimit: 5,
	}
}

func (b *Bullet) Ready() {
	b.Area3D.OnBodyEntered(func(body Node3D.Instance) {
		if body.ID() == Node3D.ID(b.Shooter) {
			return
		}
		if body.AsNode().IsInGroup("damageables") {
			var impact_point = Vector3.Sub(b.Visuals.GlobalPosition(), body.GlobalPosition())
			Object.Call(body, "damage", impact_point, b.Velocity)
		}
		body.AsNode().QueueFree()
	})
	b.AsNode3D().LookAt(Vector3.Add(b.AsNode3D().GlobalPosition(), b.Velocity))
	b.AliveLimit = b.DistanceLimit / Vector3.Length(b.Velocity)
	b.ProjectileSound.SetPitchScale(Float.RandomlyDistributed(1.0, 0.1))
	b.ProjectileSound.Play()
}

func (b *Bullet) Process(delta Float.X) {
	b.AsNode3D().SetGlobalPosition(Vector3.Add(b.AsNode3D().GlobalPosition(), Vector3.MulX(b.Velocity, delta)))
	b.TimeAlive += delta
	b.Visuals.SetScale(Vector3.MulX(Vector3.One, b.ScaleDecay.Sample(b.TimeAlive)))
	if b.TimeAlive > b.AliveLimit {
		b.AsNode().QueueFree()
	}
}

var ExplosionScene = Resource.Load[PackedScene.Is[Node3D.Instance]]("res://Player/ExplosionVisuals/explosion_scene.tscn")

type Grenade struct {
	CharacterBody3D.Extension[Grenade] `gd:"Grenade"`

	Gravity Float.X

	velocity Vector3.XYZ

	ExplosionArea       Area3D.Instance              `gd:"ExplosionArea"`
	ExplosionSound      AudioStreamPlayer3D.Instance `gd:"ExplosionSound"`
	ExplosionStartTimer Timer.Instance               `gd:"ExplosionStartTimer"`

	Visuals struct {
		Node3D.Instance

		AnimationPlayer AnimationPlayer.Instance
	} `gd:"grenade"`
}

func NewGrenade() *Grenade {
	return &Grenade{
		Gravity: DefaultGravity,
	}
}

func (g *Grenade) Throw(throw_velocity Vector3.XYZ) {
	g.velocity = throw_velocity
}

func (g *Grenade) Process(delta Float.X) {
	g.Visuals.RotateObjectLocal(Vector3.Right, 10*Angle.Radians(delta))
}

func (g *Grenade) PhysicsProcess(delta Float.X) {
	g.velocity = Vector3.Add(g.velocity, Vector3.MulX(Vector3.Down, g.Gravity*delta))
	var collision = g.AsPhysicsBody3D().MoveAndCollide(Vector3.MulX(g.velocity, delta))
	if collision != KinematicCollision3D.Nil {
		g.velocity = Vector3.MulX(Vector3.Bounce(g.velocity, collision.GetNormal()), 0.7)
		if g.ExplosionStartTimer.IsStopped() {
			g.ExplosionStartTimer.Start()
		}
	}
}

func (g *Grenade) Ready() {
	g.Visuals.AnimationPlayer.PlayNamed("wave")
	g.ExplosionStartTimer.OnTimeout(func() {
		g.AsNode().SetPhysicsProcess(false)
		g.ExplosionSound.SetPitchScale(Float.RandomlyDistributed(2, 0.1))
		g.ExplosionSound.Play()
		var bodies = g.ExplosionArea.GetOverlappingBodies()
		for _, body := range bodies {
			if damageable, ok := Object.As[Damageable](body); ok {
				// add some variance to the impact point
				var impact_point = Vector3.Normalized(Vector3.Sub(g.AsNode3D().GlobalPosition(), body.GlobalPosition()))
				impact_point = Vector3.Normalized(Vector3.Add(impact_point, Vector3.Down))
				var force = Vector3.MulX(Vector3.Neg(impact_point), 10)
				damageable.Damage(impact_point, force)
			}
		}
		var explosion = ExplosionScene.Instantiate()
		g.AsNode().GetParent().AddChild(explosion.AsNode())
		explosion.SetGlobalPosition(g.AsNode3D().GlobalPosition())
		g.AsNode3D().Hide()
		g.ExplosionSound.OnFinished(func() {
			g.AsNode().QueueFree()
		})
	})
}

var GrenadeScene = Resource.Load[PackedScene.Is[*Grenade]]("res://Player/Grenade.tscn")

type GrenadeLauncher struct {
	Node3D.Extension[GrenadeLauncher] `gd:"GrenadeLauncher"`

	MinThrowDistance Float.X
	MaxThrowDistance Float.X
	Gravity          Float.X

	FromLookPosition Vector3.XYZ
	ThrowDirection   Vector3.XYZ

	SnapMesh          Node3D.Instance         `gd:"%SnapMesh"`
	Raycast           ShapeCast3D.Instance    `gd:"%ShapeCast3D"`
	LaunchPoint       Marker3D.Instance       `gd:"%LaunchPoint"`
	TrailMeshInstance MeshInstance3D.Instance `gd:"%TrailMeshInstance"`

	throwVelocity Vector3.XYZ
	timeToLand    Float.X
}

func NewGrenadeLauncher() *GrenadeLauncher {
	return &GrenadeLauncher{
		MinThrowDistance: 7,
		MaxThrowDistance: 16,
		Gravity:          DefaultGravity,
	}
}

func (weapon *GrenadeLauncher) PhysicsProcess(delta Float.X) {
	if weapon.AsNode3D().Visible() {
		weapon.update_throw_velocity()
		weapon.draw_throw_path()
	}
}

func (weapon *GrenadeLauncher) ThrowGrenade() bool {
	if !weapon.AsNode3D().Visible() {
		return false
	}
	var grenade = GrenadeScene.Instantiate()
	weapon.AsNode().GetParent().AddChild(grenade.AsNode())
	grenade.AsNode3D().SetGlobalPosition(weapon.LaunchPoint.AsNode3D().GlobalPosition())
	grenade.Throw(weapon.throwVelocity)
	parent := Object.To[CharacterBody3D.Instance](grenade.AsNode().GetParent())
	PhysicsServer3D.BodyAddCollisionException(parent.AsCollisionObject3D().GetRid(), grenade.AsCollisionObject3D().GetRid())
	return true
}

func (weapon *GrenadeLauncher) update_throw_velocity() {
	var camera = Viewport.Get(weapon.AsNode()).GetCamera3d()
	var up_ratio = Float.Clamp(max(camera.AsNode3D().Rotation().X+0.5, -0.4)*2, 0, 1)
	var base_throw_distance = Float.Lerp(weapon.MinThrowDistance, weapon.MaxThrowDistance, Float.X(up_ratio))
	var throw_distance = base_throw_distance
	var global_camera_look_position = Vector3.Add(weapon.FromLookPosition, Vector3.MulX(weapon.ThrowDirection, throw_distance))
	weapon.Raycast.SetTargetPosition(Vector3.Sub(global_camera_look_position, weapon.Raycast.AsNode3D().GlobalPosition()))
	var to_target = weapon.Raycast.TargetPosition()
	if weapon.Raycast.GetCollisionCount() != 0 {
		if node, ok := Object.As[Node3D.Instance](weapon.Raycast.GetCollider(0)); ok {
			var has_target = node != Node3D.Nil && node.AsNode().IsInGroup("targeteables")
			weapon.SnapMesh.SetVisible(has_target)
			if has_target {
				to_target = Vector3.Sub(node.GlobalPosition(), weapon.LaunchPoint.AsNode3D().GlobalPosition())
				weapon.SnapMesh.SetGlobalPosition(Vector3.Add(weapon.LaunchPoint.AsNode3D().GlobalPosition(), to_target))
				weapon.SnapMesh.LookAt(weapon.LaunchPoint.AsNode3D().GlobalPosition())
			}
		}
	} else {
		weapon.SnapMesh.SetVisible(false)
	}
	var peak_height = max(to_target.Y+0.25, weapon.LaunchPoint.AsNode3D().Position().Y+0.25)
	var motion_up = peak_height
	var time_going_up = Float.Sqrt(2 * motion_up / weapon.Gravity)
	var motion_down = to_target.Y - peak_height
	var time_going_down = Float.Sqrt(-2 * motion_down / weapon.Gravity)
	weapon.timeToLand = time_going_up + time_going_down
	var target_position_xz_plane = Vector3.New(to_target.X, 0, to_target.Z)
	var start_position_xz_plane = Vector3.New(weapon.LaunchPoint.AsNode3D().Position().X, 0, weapon.LaunchPoint.AsNode3D().Position().Z)
	var forward_velocity = Vector3.DivX(Vector3.Sub(target_position_xz_plane, start_position_xz_plane), weapon.timeToLand)
	var velocity_up = Float.Sqrt(2 * weapon.Gravity * motion_up)
	weapon.throwVelocity = Vector3.Add(Vector3.MulX(Vector3.Up, velocity_up), forward_velocity)
}

func (weapon *GrenadeLauncher) draw_throw_path() {
	const TimeStep = 0.05
	const TrailWidth = 0.25
	var forward_direction = Vector3.Normalized(Vector3.New(weapon.ThrowDirection.X, 0, weapon.ThrowDirection.Z))
	var left_direction = Vector3.Cross(Vector3.Up, forward_direction)
	var offset_left = Vector3.MulX(left_direction, TrailWidth/2)
	var offset_right = Vector3.MulX(Vector3.Neg(left_direction), -TrailWidth/2)
	var st = SurfaceTool.New()
	st.Begin(Mesh.PrimitiveTriangles)
	var end_time = weapon.timeToLand + 0.5
	var point_previous Vector3.XYZ
	var time_current Float.X
	for time_current < end_time {
		time_current += TimeStep
		var point_current = Vector3.Add(Vector3.MulX(weapon.throwVelocity, time_current), Vector3.MulX(Vector3.Down, weapon.Gravity*0.5*time_current*time_current))
		var trail_point_left_end = Vector3.Add(point_current, offset_left)
		var trail_point_right_end = Vector3.Add(point_current, offset_right)
		var trail_point_left_start = Vector3.Add(point_previous, offset_left)
		var trail_point_right_start = Vector3.Add(point_previous, offset_right)
		var uv_progress_end = time_current / end_time
		var uv_progress_start = uv_progress_end - (TimeStep / end_time)
		var uv_value_right_start = Vector2.MulX(Vector2.Right, uv_progress_start)
		var uv_value_right_end = Vector2.MulX(Vector2.Right, uv_progress_end)
		var uv_value_left_start = Vector2.Add(Vector2.Down, uv_value_right_start)
		var uv_value_left_end = Vector2.Add(Vector2.Down, uv_value_right_end)
		point_previous = point_current
		st.SetUv(uv_value_right_end)
		st.AddVertex(trail_point_right_end)
		st.SetUv(uv_value_left_start)
		st.AddVertex(trail_point_left_start)
		st.SetUv(uv_value_left_end)
		st.AddVertex(trail_point_left_end)
		st.SetUv(uv_value_right_start)
		st.AddVertex(trail_point_right_start)
		st.SetUv(uv_value_left_start)
		st.AddVertex(trail_point_left_start)
		st.SetUv(uv_value_right_end)
		st.AddVertex(trail_point_right_end)
	}
	st.GenerateNormals()
	weapon.TrailMeshInstance.SetMesh(st.Commit().AsMesh())
}
