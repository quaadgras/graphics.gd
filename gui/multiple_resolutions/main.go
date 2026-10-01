// The root Control node ("Main") and AspectRatioContainer nodes are the most important
// pieces of this demo.
// Both nodes have their Layout set to Full Rect
// (with their rect spread across the whole viewport, and Anchor set to Full Rect).
package main

import (
	"fmt"
	"math"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AspectRatioContainer"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/HSlider"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/OptionButton"
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Window"
	"graphics.gd/startup"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
)

const options = "Panel/AspectRatioContainer/Panel/CenterContainer/Options/"

type Main struct {
	Control.Extension[Main] `gd:"Main"`

	Panel Panel.Instance
	ARC   AspectRatioContainer.Instance `gd:"Panel/AspectRatioContainer"`

	BaseSizeOption      OptionButton.Instance `gd:"Panel/AspectRatioContainer/Panel/CenterContainer/Options/WindowBaseSize/OptionButton"`
	StretchAspectOption OptionButton.Instance `gd:"Panel/AspectRatioContainer/Panel/CenterContainer/Options/WindowStretchAspect/OptionButton"`
	ScaleFactorSlider   HSlider.Instance      `gd:"Panel/AspectRatioContainer/Panel/CenterContainer/Options/WindowScaleFactor/HSlider"`
	ScaleFactorValue    Label.Instance        `gd:"Panel/AspectRatioContainer/Panel/CenterContainer/Options/WindowScaleFactor/Value"`
	GUIMarginSlider     HSlider.Instance      `gd:"Panel/AspectRatioContainer/Panel/CenterContainer/Options/GUIMargin/HSlider"`
	GUIMarginValue      Label.Instance        `gd:"Panel/AspectRatioContainer/Panel/CenterContainer/Options/GUIMargin/Value"`

	baseWindowSize Vector2i.XY
	stretchMode    Window.ContentScaleMode
	stretchAspect  Window.ContentScaleAspect
	scaleFactor    Float.X
	guiAspectRatio Float.X
	guiMargin      Float.X
}

func NewMain() *Main {
	return &Main{
		// These defaults match this demo's project settings. Adjust as needed if adapting this
		// in your own project.
		baseWindowSize: Vector2i.XY{
			X: int32(settingInt("display/window/size/viewport_width")),
			Y: int32(settingInt("display/window/size/viewport_height")),
		},
		stretchMode:    Window.ContentScaleModeCanvasItems,
		stretchAspect:  Window.ContentScaleAspectExpand,
		scaleFactor:    1.0,
		guiAspectRatio: -1.0,
	}
}

func settingInt(name string) int {
	switch v := ProjectSettings.GetSetting(name, 0).(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	}
	return 0
}

func (m *Main) window() Window.Instance { return SceneTree.Get(m.AsNode()).Root() }

func (m *Main) Ready() {
	// The `resized` signal will be emitted when the window size changes, as the root Control node
	// is resized whenever the window size changes. This is because the root Control node
	// uses a Full Rect anchor, so its size will always be equal to the window size.
	m.AsControl().OnResized(m.OnResized)
	Callable.Defer(Callable.New(m.updateContainer))
}

func (m *Main) updateContainer() {
	// The code within this function needs to be run deferred to work around an issue with containers
	// having a 1-frame delay with updates.
	// Otherwise, `panel.size` returns a value of the previous frame, which results in incorrect
	// sizing of the inner AspectRatioContainer when using the Fit to Window setting.
	for range 2 {
		if math.Abs(float64(m.guiAspectRatio+1.0)) < 1e-5 {
			// Fit to Window. Tell the AspectRatioContainer to use the same aspect ratio as the window,
			// making the AspectRatioContainer not have any visible effect.
			m.ARC.SetRatio(Vector2.Aspect(m.Panel.AsControl().Size()))
			// Apply GUI offset on the AspectRatioContainer's parent (Panel).
			// This also makes the GUI offset apply on controls located outside the AspectRatioContainer
			// (such as the inner side label in this demo).
			m.Panel.AsControl().SetOffsetTop(m.guiMargin)
			m.Panel.AsControl().SetOffsetBottom(-m.guiMargin)
		} else {
			// Constrained aspect ratio.
			m.ARC.SetRatio(Float.Min(Vector2.Aspect(m.Panel.AsControl().Size()), m.guiAspectRatio))
			// Adjust top and bottom offsets relative to the aspect ratio when it's constrained.
			// This ensures that GUI offset settings behave exactly as if the window had the
			// original aspect ratio size.
			m.Panel.AsControl().SetOffsetTop(m.guiMargin / m.guiAspectRatio)
			m.Panel.AsControl().SetOffsetBottom(-m.guiMargin / m.guiAspectRatio)
		}
		m.Panel.AsControl().SetOffsetLeft(m.guiMargin)
		m.Panel.AsControl().SetOffsetRight(-m.guiMargin)
	}
}

func (m *Main) OnGuiAspectRatioItemSelected(index int) {
	switch index {
	case 0: // Fit to Window
		m.guiAspectRatio = -1.0
	case 1: // 5:4
		m.guiAspectRatio = 5.0 / 4.0
	case 2: // 4:3
		m.guiAspectRatio = 4.0 / 3.0
	case 3: // 3:2
		m.guiAspectRatio = 3.0 / 2.0
	case 4: // 16:10
		m.guiAspectRatio = 16.0 / 10.0
	case 5: // 16:9
		m.guiAspectRatio = 16.0 / 9.0
	case 6: // 21:9
		m.guiAspectRatio = 21.0 / 9.0
	}
	Callable.Defer(Callable.New(m.updateContainer))
}

func (m *Main) OnResized() {
	Callable.Defer(Callable.New(m.updateContainer))
}

func (m *Main) OnGuiMarginDragEnded(_ bool) {
	m.guiMargin = m.GUIMarginSlider.AsRange().Value()
	m.GUIMarginValue.SetText(fmt.Sprint(m.guiMargin))
	Callable.Defer(Callable.New(m.updateContainer))
}

func (m *Main) OnWindowBaseSizeItemSelected(index int) {
	switch index {
	case 0: // 648×648 (1:1)
		m.baseWindowSize = Vector2i.XY{X: 648, Y: 648}
	case 1: // 640×480 (4:3)
		m.baseWindowSize = Vector2i.XY{X: 640, Y: 480}
	case 2: // 720×480 (3:2)
		m.baseWindowSize = Vector2i.XY{X: 720, Y: 480}
	case 3: // 800×600 (4:3)
		m.baseWindowSize = Vector2i.XY{X: 800, Y: 600}
	case 4: // 1152×648 (16:9)
		m.baseWindowSize = Vector2i.XY{X: 1152, Y: 648}
	case 5: // 1280×720 (16:9)
		m.baseWindowSize = Vector2i.XY{X: 1280, Y: 720}
	case 6: // 1280×800 (16:10)
		m.baseWindowSize = Vector2i.XY{X: 1280, Y: 800}
	case 7: // 1680×720 (21:9)
		m.baseWindowSize = Vector2i.XY{X: 1680, Y: 720}
	}
	m.window().SetContentScaleSize(m.baseWindowSize)
	Callable.Defer(Callable.New(m.updateContainer))
}

func (m *Main) OnWindowStretchModeItemSelected(index int) {
	m.stretchMode = Window.ContentScaleMode(index)
	m.window().SetContentScaleMode(m.stretchMode)

	// Disable irrelevant options when the stretch mode is Disabled.
	disabled := m.stretchMode == Window.ContentScaleModeDisabled
	m.BaseSizeOption.AsBaseButton().SetDisabled(disabled)
	m.StretchAspectOption.AsBaseButton().SetDisabled(disabled)
}

func (m *Main) OnWindowStretchAspectItemSelected(index int) {
	m.stretchAspect = Window.ContentScaleAspect(index)
	m.window().SetContentScaleAspect(m.stretchAspect)
}

func (m *Main) OnWindowScaleFactorDragEnded(_ bool) {
	m.scaleFactor = m.ScaleFactorSlider.AsRange().Value()
	m.ScaleFactorValue.SetText(fmt.Sprintf("%d%%", int(m.scaleFactor*100)))
	m.window().SetContentScaleFactor(m.scaleFactor)
}

func (m *Main) OnWindowStretchScaleModeItemSelected(index int) {
	m.window().SetContentScaleStretch(Window.ContentScaleStretch(index))
}

func main() {
	classdb.Register[Main](NewMain)
	startup.Scene()
}
