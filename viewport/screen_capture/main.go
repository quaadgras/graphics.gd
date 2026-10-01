package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/TextureRect"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

// ScreenCapture captures the viewport into a texture when the button is pressed.
type ScreenCapture struct {
	Control.Extension[ScreenCapture] `gd:"ScreenCapture"`

	CapturedImage TextureRect.Instance
	CaptureButton Button.Instance
}

func (s *ScreenCapture) Ready() {
	// Focus button for keyboard/gamepad-friendly navigation.
	s.CaptureButton.AsControl().GrabFocus()
}

func (s *ScreenCapture) OnCaptureButtonPressed() {
	// Retrieve the captured image.
	img := Viewport.Get(s.AsNode()).GetTexture().AsTexture2D().GetImage()

	// Create a texture for it.
	tex := ImageTexture.CreateFromImage(img)

	// Set the texture to the captured image node.
	s.CapturedImage.SetTexture(tex.AsTexture2D())

	// Colorize the button with a random color, so you can see which button belongs to which capture.
	s.CaptureButton.AsCanvasItem().SetModulate(Color.HSV(Float.Random(), Float.RandomBetween(0.2, 0.8), 1.0))
}

func main() {
	classdb.Register[ScreenCapture]()
	startup.Scene()
}
