package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/CanvasLayer"
	"graphics.gd/classdb/CheckBox"
	"graphics.gd/classdb/HSlider"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/MethodTweener"
	"graphics.gd/classdb/OptionButton"
	"graphics.gd/classdb/Path2D"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/SpinBox"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/classdb/TextureProgressBar"
	"graphics.gd/classdb/Tween"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

type Main struct {
	CanvasLayer.Extension[Main] `gd:"Main"`

	Icon           Sprite2D.Instance           `gd:"%Icon"`
	CountdownLabel Label.Instance              `gd:"%CountdownLabel"`
	Path           Path2D.Instance             `gd:"Path2D"`
	Progress       TextureProgressBar.Instance `gd:"%Progress"`
	SpeedSlider    HSlider.Instance            `gd:"%SpeedSlider"`
	SpeedLabel     Label.Instance              `gd:"%SpeedLabel"`
	Infinite       CheckBox.Instance           `gd:"%Infinite"`
	Loops          SpinBox.Instance            `gd:"%Loops"`
	Reset          CheckBox.Instance           `gd:"%Reset"`
	Ease1          OptionButton.Instance       `gd:"%Ease1"`
	Trans1         OptionButton.Instance       `gd:"%Trans1"`
	Ease3          OptionButton.Instance       `gd:"%Ease3"`
	Trans3         OptionButton.Instance       `gd:"%Trans3"`
	Ease7          OptionButton.Instance       `gd:"%Ease7"`
	Trans7         OptionButton.Instance       `gd:"%Trans7"`

	tween, subTween   Tween.Instance
	iconStartPosition Vector2.XY
}

func (m *Main) Ready() {
	m.iconStartPosition = m.Icon.AsNode2D().Position()
}

func (m *Main) Process(_ Float.X) {
	if m.tween == Tween.Nil || !m.tween.IsRunning() {
		return
	}
	m.Progress.AsRange().SetValue(m.tween.GetTotalElapsedTime())
}

func (m *Main) iconObject() Object.Instance { return Object.Instance(m.Icon.AsObject()) }

func (m *Main) StartAnimation() {
	// Reset the icon to original state.
	m.reset(false)
	// Create the Tween. Also sets the initial animation speed.
	// All methods that modify Tween will return the Tween, so you can chain them.
	m.tween = m.AsNode().CreateTween().SetSpeedScale(m.SpeedSlider.AsRange().Value())
	tween := m.tween
	icon := m.iconObject()

	// Sets the amount of loops. 1 loop = 1 animation cycle, so e.g. 2 loops will play animation twice.
	if m.Infinite.AsBaseButton().ButtonPressed() {
		tween.SetLoops() // Called without arguments, the Tween will loop infinitely.
	} else {
		tween.MoreArgs().SetLoops(int(m.Loops.AsRange().Value()))
	}

	// Step 1

	if m.isStepEnabled("MoveTo", 1.0) {
		// PropertyTweener.Make returns a Tweener object. Its methods can also be chained, but
		// it's stored in a variable here for readability (chained lines tend to be long).
		tweener := PropertyTweener.Make(tween, icon, "position", Vector2.New(400, 250), 1.0)
		tweener.SetEase(Tween.EaseType(m.Ease1.Selected()))
		tweener.SetTrans(Tween.TransitionType(m.Trans1.Selected()))
	}

	// Step 2

	if m.isStepEnabled("ColorRed", 1.0) {
		PropertyTweener.Make(tween, icon, "self_modulate", Color.W3C.Red, 1.0)
	}

	// Step 3

	if m.isStepEnabled("MoveRight", 1.0) {
		// AsRelative() makes the value relative, so in this case it moves the icon
		// 200 pixels from the previous position.
		tweener := PropertyTweener.Make(tween, icon, "position:x", 200.0, 1.0).AsRelative()
		tweener.SetEase(Tween.EaseType(m.Ease3.Selected()))
		tweener.SetTrans(Tween.TransitionType(m.Trans3.Selected()))
	}
	if m.isStepEnabled("Roll", 0.0) {
		// Parallel() makes the Tweener run in parallel to the previous one.
		tweener := PropertyTweener.Make(tween.Parallel(), icon, "rotation", Float.X(Angle.Tau), 1.0)
		tweener.SetEase(Tween.EaseType(m.Ease3.Selected()))
		tweener.SetTrans(Tween.TransitionType(m.Trans3.Selected()))
	}

	// Step 4

	if m.isStepEnabled("MoveLeft", 1.0) {
		PropertyTweener.Make(tween, icon, "position", Vector2.MulX(Vector2.Left, 200), 1.0).AsRelative()
	}
	if m.isStepEnabled("Jump", 0.0) {
		// Jump has 2 substeps, so to make it properly parallel, it can be done in a sub-Tween.
		// Here we are calling a closure that creates a sub-Tween.
		// Any number of Tweens can animate a single object in the same time.
		tween.Parallel().TweenCallback(func() {
			// Note that transition is set on Tween, but ease is set on Tweener.
			// Values set on Tween will affect all Tweeners (as defaults) and values
			// on Tweeners can override them.
			m.subTween = m.AsNode().CreateTween().SetSpeedScale(m.SpeedSlider.AsRange().Value()).SetTrans(Tween.TransSine)
			PropertyTweener.Make(m.subTween, icon, "position:y", -150.0, 0.5).AsRelative().SetEase(Tween.EaseOut)
			PropertyTweener.Make(m.subTween, icon, "position:y", 150.0, 0.5).AsRelative().SetEase(Tween.EaseIn)
		})
	}

	// Step 5

	if m.isStepEnabled("Blink", 2.0) {
		// Loops are handy when creating some animations.
		for range 10 {
			tween.TweenCallback(m.Icon.AsCanvasItem().Hide).SetDelay(0.1)
			tween.TweenCallback(m.Icon.AsCanvasItem().Show).SetDelay(0.1)
		}
	}

	// Step 6

	if m.isStepEnabled("Teleport", 0.5) {
		// Tweening a value with 0 duration makes it change instantly.
		PropertyTweener.Make(tween, icon, "position", Vector2.New(325, 325), 0)
		tween.TweenInterval(0.5)
		// Closures can be used for advanced callbacks.
		tween.TweenCallback(func() { m.Icon.AsNode2D().SetPosition(Vector2.New(680, 215)) })
	}

	// Step 7

	if m.isStepEnabled("Curve", 3.5) {
		// Method tweening is useful for animating values that can't be directly interpolated.
		// It can be used for remapping and some very advanced animations.
		// Here it's used for moving sprite along a path, using an inline closure.
		curve := m.Path.Curve()
		tweener := MethodTweener.Make(tween, Callable.New(func(v Float.X) {
			m.Icon.AsNode2D().SetPosition(Vector2.Add(m.Path.AsNode2D().Position(), curve.MoreArgs().SampleBaked(v, false)))
		}), 0.0, curve.GetBakedLength(), 3.0).SetDelay(0.5)
		tweener.SetEase(Tween.EaseType(m.Ease7.Selected()))
		tweener.SetTrans(Tween.TransitionType(m.Trans7.Selected()))
	}

	// Step 8

	if m.isStepEnabled("Wait", 2.0) {
		// ...
		tween.TweenInterval(2)
	}

	// Step 9

	if m.isStepEnabled("Countdown", 3.0) {
		tween.TweenCallback(m.CountdownLabel.AsCanvasItem().Show)
		MethodTweener.Make(tween, Callable.New(m.doCountdown), 4, 1, 3)
		tween.TweenCallback(m.CountdownLabel.AsCanvasItem().Hide)
	}

	// Step 10

	if m.isStepEnabled("Enlarge", 0.0) {
		PropertyTweener.Make(tween, icon, "scale", Vector2.MulX(Vector2.One, 5), 0.5).SetTrans(Tween.TransElastic).SetEase(Tween.EaseOut)
	}
	if m.isStepEnabled("Vanish", 1.0) {
		PropertyTweener.Make(tween.Parallel(), icon, "self_modulate:a", 0.0, 1.0)
	}

	if m.Loops.AsRange().Value() > 1 || m.Infinite.AsBaseButton().ButtonPressed() {
		tween.TweenCallback(m.Icon.AsCanvasItem().Show)
		tween.TweenCallback(func() { m.Icon.AsCanvasItem().SetSelfModulate(Color.W3C.White) })
	}

	// RESET step

	if m.Reset.AsBaseButton().ButtonPressed() {
		tween.TweenCallback(func() { m.reset(true) })
	}
}

func (m *Main) doCountdown(number int) {
	m.CountdownLabel.SetText(fmt.Sprint(number))
}

func (m *Main) reset(soft bool) {
	m.Icon.AsNode2D().SetPosition(m.iconStartPosition)
	m.Icon.AsCanvasItem().SetSelfModulate(Color.W3C.White)
	m.Icon.AsNode2D().SetRotation(0)
	m.Icon.AsNode2D().SetScale(Vector2.One)
	m.Icon.AsCanvasItem().Show()
	m.CountdownLabel.AsCanvasItem().Hide()

	if soft {
		// Only reset properties.
		return
	}

	if m.tween != Tween.Nil {
		m.tween.Kill()
		m.tween = Tween.Nil
	}

	if m.subTween != Tween.Nil {
		m.subTween.Kill()
		m.subTween = Tween.Nil
	}

	m.Progress.AsRange().SetMaxValue(0)
}

func (m *Main) isStepEnabled(step string, expectedTime Float.X) bool {
	button, _ := Object.As[CheckBox.Instance](m.AsNode().GetNode("%" + step))
	enabled := button.AsBaseButton().ButtonPressed()
	if enabled {
		m.Progress.AsRange().SetMaxValue(m.Progress.AsRange().MaxValue() + expectedTime)
	}
	return enabled
}

func pauseResume(tween Tween.Instance) {
	if tween != Tween.Nil && tween.IsValid() {
		if tween.IsRunning() {
			tween.Pause()
		} else {
			tween.Play()
		}
	}
}

func (m *Main) PauseResume() {
	pauseResume(m.tween)
	pauseResume(m.subTween)
}

func (m *Main) KillTween() {
	if m.tween != Tween.Nil {
		m.tween.Kill()
	}
	if m.subTween != Tween.Nil {
		m.subTween.Kill()
	}
}

func (m *Main) SpeedChanged(value Float.X) {
	if m.tween != Tween.Nil {
		m.tween.SetSpeedScale(value)
	}
	if m.subTween != Tween.Nil {
		m.subTween.SetSpeedScale(value)
	}
	m.SpeedLabel.SetText(fmt.Sprint("x", value))
}

func (m *Main) InfiniteToggled(buttonPressed bool) {
	m.Loops.SetEditable(!buttonPressed)
}

func main() {
	classdb.Register[Main]()
	startup.Scene()
}
