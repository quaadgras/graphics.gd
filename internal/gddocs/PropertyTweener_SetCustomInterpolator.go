/*
[gdscript]
@export var curve: Curve

func _ready():
	var tween = create_tween()
	# Interpolate the value using a custom curve.
	tween.tween_property(self, "position:x", 300, 1).as_relative().set_custom_interpolator(tween_curve)

func tween_curve(v):
	return curve.sample_baked(v)
[/gdscript]
[csharp]
[Export]
public Curve Curve { get; set; }

public override void _Ready()
{
	Tween tween = CreateTween();
	// Interpolate the value using a custom curve.
	Callable tweenCurveCallable = Callable.From<float, float>(TweenCurve);
	tween.TweenProperty(this, "position:x", 300.0f, 1.0f).AsRelative().SetCustomInterpolator(tweenCurveCallable);
}

private float TweenCurve(float value)
{
	return Curve.SampleBaked(value);
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Curve"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/Tween"
)

func PropertyTweener_SetCustomInterpolator() {
	var curve Curve.Instance
	var tween Tween.Instance
	PropertyTweener.Make(tween, self, "position:x", 300, 1).AsRelative().SetCustomInterpolator(func(v float32) float32 {
		return curve.SampleBaked(v)
	})

}
