package main

import (
	"strconv"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/CharacterBody3D"
	"graphics.gd/classdb/GeometryInstance3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/ShaderMaterial"
	"graphics.gd/classdb/StandardMaterial3D"
	"graphics.gd/classdb/SubViewport"
	"graphics.gd/classdb/TextureRect"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
)

// SplitPlayer moves the player.
type SplitPlayer struct {
	CharacterBody3D.Extension[SplitPlayer] `gd:"SplitPlayer"`

	PlayerId  int `range:"1,2"`
	WalkSpeed Float.X
}

func NewSplitPlayer() *SplitPlayer {
	return &SplitPlayer{PlayerId: 1, WalkSpeed: 2.0}
}

func (p *SplitPlayer) PhysicsProcess(_ Float.X) {
	body := p.AsCharacterBody3D()
	id := strconv.Itoa(p.PlayerId)
	moveDirection := Input.GetVector(
		"move_left_player"+id,
		"move_right_player"+id,
		"move_up_player"+id,
		"move_down_player"+id,
	)
	velocity := body.Velocity()
	velocity.X += moveDirection.X * p.WalkSpeed
	velocity.Z += moveDirection.Y * p.WalkSpeed

	// Apply friction.
	velocity = Vector3.MulX(velocity, 0.9)
	body.SetVelocity(velocity)

	body.MoveAndSlide()
}

// WallColoring sets a random color to all objects in the "walls" group.
type WallColoring struct {
	Node3D.Extension[WallColoring] `gd:"WallColoring"`
	classdb.Tool
}

func (w *WallColoring) Ready() {
	for _, wall := range SceneTree.Get(w.AsNode()).GetNodesInGroup("walls") {
		geometry, ok := Object.As[GeometryInstance3D.Instance](wall)
		if !ok {
			continue
		}
		material := StandardMaterial3D.New()
		material.AsBaseMaterial3D().SetAlbedoColor(Color.RGBA{R: Float.Random(), G: Float.Random(), B: Float.Random(), A: 1})
		geometry.SetMaterialOverride(material.AsMaterial())
	}
}

// CameraController handles the motion of both player cameras as well as communication with the
// SplitScreen shader to achieve the dynamic split screen effect.
//
// Cameras are placed on the segment joining the two players, either in the middle
// if players are close enough or at a fixed distance if they are not.
// In the first case, both cameras being at the same location, only the view of
// the first one is used for the entire screen thus allowing the players to play
// on an unsplit screen.
// In the second case, the screen is split in two with a line perpendicular to the
// segment joining the two players.
type CameraController struct {
	Node3D.Extension[CameraController] `gd:"CameraController"`

	MaxSeparation              Float.X    // The distance between players at which the view starts to split.
	SplitLineThickness         Float.X    // The thickness of the split line in pixels.
	SplitLineColor             Color.RGBA // Color of the split line.
	AdaptiveSplitLineThickness bool       // If true, the split line thickness will vary depending on the distance between players.

	Player1   *SplitPlayer         `gd:"../Player1"`
	Player2   *SplitPlayer         `gd:"../Player2"`
	View      TextureRect.Instance `gd:"View"`
	Viewport1 SubViewport.Instance `gd:"Viewport1"`
	Viewport2 SubViewport.Instance `gd:"Viewport2"`
	Camera1   Camera3D.Instance    `gd:"Viewport1/Camera1"`
	Camera2   Camera3D.Instance    `gd:"Viewport2/Camera2"`

	material ShaderMaterial.Instance
}

func NewCameraController() *CameraController {
	return &CameraController{
		MaxSeparation:              20.0,
		SplitLineThickness:         3.0,
		SplitLineColor:             Color.X11.Black,
		AdaptiveSplitLineThickness: true,
	}
}

func (c *CameraController) Ready() {
	c.material, _ = Object.As[ShaderMaterial.Instance](c.View.AsCanvasItem().Material())

	c.onSizeChanged()
	c.updateSplitscreen()

	Viewport.Get(c.AsNode()).OnSizeChanged(c.onSizeChanged)

	c.material.SetShaderParameter("viewport1", c.Viewport1.AsViewport().GetTexture())
	c.material.SetShaderParameter("viewport2", c.Viewport2.AsViewport().GetTexture())
}

func (c *CameraController) Process(_ Float.X) {
	c.moveCameras()
	c.updateSplitscreen()
}

func (c *CameraController) moveCameras() {
	positionDifference := c.positionDifferenceInWorld()

	distance := Float.Clamp(horizontalLength(positionDifference), 0, c.MaxSeparation)

	positionDifference = Vector3.MulX(Vector3.Normalized(positionDifference), distance)

	player1Position := c.Player1.AsNode3D().Position()
	player2Position := c.Player2.AsNode3D().Position()

	camera1Position := c.Camera1.AsNode3D().Position()
	camera1Position.X = player1Position.X + positionDifference.X/2.0
	camera1Position.Z = player1Position.Z + positionDifference.Z/2.0
	c.Camera1.AsNode3D().SetPosition(camera1Position)

	camera2Position := c.Camera2.AsNode3D().Position()
	camera2Position.X = player2Position.X - positionDifference.X/2.0
	camera2Position.Z = player2Position.Z - positionDifference.Z/2.0
	c.Camera2.AsNode3D().SetPosition(camera2Position)
}

func (c *CameraController) updateSplitscreen() {
	screenSize := Viewport.Get(c.AsNode()).GetVisibleRect().Size
	player1Position := Vector2.Div(c.Camera1.UnprojectPosition(c.Player1.AsNode3D().Position()), screenSize)
	player2Position := Vector2.Div(c.Camera2.UnprojectPosition(c.Player2.AsNode3D().Position()), screenSize)

	var thickness Float.X
	if c.AdaptiveSplitLineThickness {
		positionDifference := c.positionDifferenceInWorld()
		distance := horizontalLength(positionDifference)
		thickness = Float.Lerp(0, c.SplitLineThickness, (distance-c.MaxSeparation)/c.MaxSeparation)
		thickness = Float.Clamp(thickness, 0, c.SplitLineThickness)
	} else {
		thickness = c.SplitLineThickness
	}

	c.material.SetShaderParameter("split_active", c.isSplitState())
	c.material.SetShaderParameter("player1_position", player1Position)
	c.material.SetShaderParameter("player2_position", player2Position)
	c.material.SetShaderParameter("split_line_thickness", thickness)
	c.material.SetShaderParameter("split_line_color", c.SplitLineColor)
}

// isSplitState returns true if split screen is active (which occurs when players are
// too far apart from each other), false otherwise.
// Only the horizontal components (x, z) are used for distance computation.
func (c *CameraController) isSplitState() bool {
	positionDifference := c.positionDifferenceInWorld()
	separationDistance := horizontalLength(positionDifference)
	return separationDistance > c.MaxSeparation
}

func (c *CameraController) onSizeChanged() {
	screenSize := Viewport.Get(c.AsNode()).GetVisibleRect().Size

	size := Vector2i.XY{X: int32(screenSize.X), Y: int32(screenSize.Y)}
	c.Viewport1.SetSize(size)
	c.Viewport2.SetSize(size)

	c.material.SetShaderParameter("viewport_size", screenSize)
}

func (c *CameraController) positionDifferenceInWorld() Vector3.XYZ {
	return Vector3.Sub(c.Player2.AsNode3D().Position(), c.Player1.AsNode3D().Position())
}

func horizontalLength(v Vector3.XYZ) Float.X {
	return Vector2.Length(Vector2.XY{X: v.X, Y: v.Z})
}

func main() {
	classdb.Register[SplitPlayer](NewSplitPlayer)
	classdb.Register[WallColoring]()
	classdb.Register[CameraController](NewCameraController)
	startup.Scene()
}
