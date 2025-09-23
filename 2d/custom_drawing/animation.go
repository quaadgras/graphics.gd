package main

import (
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Panel"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

type Animation struct {
	Panel.Extension[Animation] `gd:"CustomAnimation"`

	antialiasable

	time Float.X
}

func (a *Animation) Process(delta Float.X) {
	a.time += delta // Increment a counter variable that we use in [Animation.Draw].
	// Force redrawing on every processed frame, so that the animation can visibly progress.
	// Only do this when the node is visible in tree, so that we don't force continuous redrawing
	// when not needed (low-processor usage mode is enabled in this demo).
	if a.AsCanvasItem().IsVisibleInTree() {
		a.AsCanvasItem().QueueRedraw()
	}
}

func (a *Animation) Draw() {
	var (
		margin = Vector2.New(240, 70)
		offset = Vector2.Zero
	)
	// Draw an animated arc to simulate a circular progress bar.
	// The start angle is set so the arc starts from the top.
	const PointCount = 48
	var progress = Angle.Radians(Float.Wrap(a.time, 0, 1))
	a.AsCanvasItem().MoreArgs().DrawArc(
		Vector2.Add(margin, offset),
		50,
		0.75*Angle.Tau,
		(0.75+progress)*Angle.Tau,
		PointCount,
		Color.W3C.MediumAquamarine,
		a.LineWidth(),
		a.UseAntialiasing,
	)
}

type AnimationSlice struct {
	Control.Extension[AnimationSlice] `gd:"CustomAnimationSlice"`

	antialiasable
}

func (a *AnimationSlice) Draw() {
	var (
		margin = Vector2.New(240, 70)
		offset = Vector2.New(0, 150)
	)
	// This is an example of using draw commands to create animations.
	// For "continuous" animations, you can use a timer within `Draw()` and call `AsCanvasItem().QueueRedraw()`
	// in `Process(Float.X)` to redraw every frame.
	// Animation length in seconds. The animation will loop after the specified duration.
	const (
		AnimationLength = 2.0
		AnimationFrames = 5 //per second
	)
	// Declare an animation frame with randomized rotation and color for each frame.
	// `DrawAnimationSlice()` makes it so the following draw commands are only visible
	// on screen when the current time is within the animation slice.
	// NOTE: Pause does not affect animations drawn by `DrawAnimationSlice()`
	// (they will keep playing).
	canvas := a.AsCanvasItem()
	for frame := range AnimationFrames {
		// `Remap()` is useful to determine the time slice in which a frame is visible.
		// For example, on the 2nd frame, `slice_begin` is `0.2` and `slice_end` is `0.4`.
		var slice_begin = Float.Remap(Float.X(frame), 0, AnimationFrames, 0, AnimationLength)
		var slice_end = Float.Remap(Float.X(frame+1), 0, AnimationFrames, 0, AnimationLength)
		canvas.DrawAnimationSlice(AnimationLength, slice_begin, slice_end)
		canvas.MoreArgs().DrawSetTransform(Vector2.Add(margin, offset), Angle.InRadians(Angle.Degrees(Float.RandomBetween(-5, 5))), Vector2.One)
		canvas.MoreArgs().DrawRect(
			Rect2.PositionSize{Size: Vector2.New(100, 50)},
			Color.HSV(Float.Random(), 0.4, 1.0),
			true,
			-1,
			a.UseAntialiasing,
		)
	}
}
