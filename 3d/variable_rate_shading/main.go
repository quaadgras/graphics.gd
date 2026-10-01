package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/MobileVRInterface"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/OptionButton"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/classdb/TextureRect"
	"graphics.gd/classdb/VBoxContainer"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/classdb/XRInterface"
	"graphics.gd/classdb/XRServer"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

// VRS lets the user switch between the viewport's variable rate shading modes.
type VRS struct {
	Node3D.Extension[VRS] `gd:"VRS"`

	Texture Texture2D.Instance

	OptionButton OptionButton.Instance `gd:"CanvasLayer/VBoxContainer/HBoxContainer/OptionButton"`
	TextureRect  TextureRect.Instance  `gd:"CanvasLayer/VBoxContainer/TextureRect"`
	Camera       Camera3D.Instance     `gd:"Camera3D"`
	XROrigin     Node3D.Instance       `gd:"XROrigin3D"`
	XRCamera     Camera3D.Instance     `gd:"XROrigin3D/XRCamera3D"`

	xrInterface XRInterface.Instance
}

func (v *VRS) setXrMode() {
	viewport := Viewport.Get(v.AsNode())
	if viewport.VrsMode() == Viewport.VrsXr {
		v.xrInterface = XRServer.FindInterface("Native mobile")
		if v.xrInterface != XRInterface.Nil && v.xrInterface.Initialize() {
			// Disable a lot of VR-specific stuff like lens distortion.
			if mobile, ok := Object.As[MobileVRInterface.Instance](v.xrInterface); ok {
				mobile.SetEyeHeight(0.0)
				mobile.SetK1(0.0)
				mobile.SetK2(0.0)
				mobile.SetOversample(1.0)
			}

			viewport.SetUseXr(true)
			v.XRCamera.SetCurrent(true)

			// Reposition our origin point to work around an engine bug.
			v.XROrigin.SetGlobalTransform(v.Camera.AsNode3D().GlobalTransform())
		}
	} else {
		if v.xrInterface != XRInterface.Nil {
			v.xrInterface.Uninitialize()
		}

		viewport.SetUseXr(false)
		v.Camera.SetCurrent(true)
	}
}

func (v *VRS) updateTexture() {
	viewport := Viewport.Get(v.AsNode())
	switch viewport.VrsMode() {
	case Viewport.VrsDisabled:
		v.TextureRect.AsCanvasItem().SetVisible(false)
	case Viewport.VrsTexture:
		viewport.SetVrsTexture(v.Texture)
		v.TextureRect.SetTexture(v.Texture)
		v.TextureRect.AsCanvasItem().SetVisible(true)
	case Viewport.VrsXr:
		// Doesn't seem to be supported yet. This should be exposed in a future engine version.
		v.TextureRect.AsCanvasItem().SetVisible(false)
	}
}

func (v *VRS) Ready() {
	v.OptionButton.Select(int(Viewport.Get(v.AsNode()).VrsMode()))
	v.updateTexture()
}

func (v *VRS) OnOptionButtonItemSelected(index int) {
	Viewport.Get(v.AsNode()).SetVrsMode(Viewport.VRSMode(index))
	v.setXrMode()
	v.updateTexture()
}

// Information displays the viewport resolution and frames per second.
type Information struct {
	VBoxContainer.Extension[Information] `gd:"Information"`

	Resolution Label.Instance
	FPS        Label.Instance
}

func (i *Information) Ready() {
	Viewport.Get(i.AsNode()).OnSizeChanged(i.onViewportSizeChanged)
	i.onViewportSizeChanged()
}

func (i *Information) Process(_ Float.X) {
	fps := Engine.GetFramesPerSecond()
	i.FPS.SetText(fmt.Sprintf("%d FPS (%.2f mspf)", int(fps), 1000.0/fps))
}

func (i *Information) onViewportSizeChanged() {
	size := Viewport.Get(i.AsNode()).GetVisibleRect().Size
	i.Resolution.SetText(fmt.Sprintf("%d × %d", int(size.X), int(size.Y)))
}

func main() {
	classdb.Register[VRS]()
	classdb.Register[Information]()
	startup.Scene()
}
