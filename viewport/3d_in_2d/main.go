package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AnimatedSprite2D"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/classdb/SubViewport"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
)

// ThreeDInTwoD renders a 3D scene into a SubViewport and shows it as a sprite.
type ThreeDInTwoD struct {
	Node2D.Extension[ThreeDInTwoD] `gd:"ThreeDInTwoD"`

	SubViewport      SubViewport.Instance
	ViewportSprite   Sprite2D.Instance
	AnimatedSprite2D AnimatedSprite2D.Instance

	viewportInitialSize Vector2i.XY
}

func (t *ThreeDInTwoD) Ready() {
	t.viewportInitialSize = t.SubViewport.Size()
	t.AnimatedSprite2D.Play()
	Viewport.Get(t.AsNode()).OnSizeChanged(t.rootViewportSizeChanged)
}

// rootViewportSizeChanged is called when the root's viewport size changes (i.e. when the window is resized).
// This is done to handle multiple resolutions without losing quality.
func (t *ThreeDInTwoD) rootViewportSizeChanged() {
	// The viewport is resized depending on the window height.
	// To compensate for the larger resolution, the viewport sprite is scaled down.
	height := Viewport.Get(t.AsNode()).GetVisibleRect().Size.Y
	t.SubViewport.SetSize(Vector2i.XY{X: int32(height), Y: int32(height)})
	t.ViewportSprite.AsNode2D().SetScale(Vector2.MulX(Vector2.One, Float.X(t.viewportInitialSize.Y)/height))
}

// Robot3D is a simple class that rotates the model.
type Robot3D struct {
	Node3D.Extension[Robot3D] `gd:"Robot3D"`

	Model Node3D.Instance
}

func (r *Robot3D) Process(delta Float.X) {
	rotation := r.Model.Rotation()
	rotation.Y += Angle.Radians(delta * 0.7)
	r.Model.SetRotation(rotation)
}

func main() {
	classdb.Register[ThreeDInTwoD]()
	classdb.Register[Robot3D]()
	startup.Scene()
}
