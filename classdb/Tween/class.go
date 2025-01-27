// Package Tween provides methods for working with Tween object instances.
package Tween

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ = slices.Delete[[]struct{}, struct{}]

/*
Tweens are mostly useful for animations requiring a numerical property to be interpolated over a range of values. The name [i]tween[/i] comes from [i]in-betweening[/i], an animation technique where you specify [i]keyframes[/i] and the computer interpolates the frames that appear between them. Animating something with a [Tween] is called tweening.
[Tween] is more suited than [AnimationPlayer] for animations where you don't know the final values in advance. For example, interpolating a dynamically-chosen camera zoom value is best done with a [Tween]; it would be difficult to do the same thing with an [AnimationPlayer] node. Tweens are also more light-weight than [AnimationPlayer], so they are very much suited for simple animations or general tasks that don't require visual tweaking provided by the editor. They can be used in a "fire-and-forget" manner for some logic that normally would be done by code. You can e.g. make something shoot periodically by using a looped [CallbackTweener] with a delay.
A [Tween] can be created by using either [method SceneTree.create_tween] or [method Node.create_tween]. [Tween]s created manually (i.e. by using [code]Tween.new()[/code]) are invalid and can't be used for tweening values.
A tween animation is created by adding [Tweener]s to the [Tween] object, using [method tween_property], [method tween_interval], [method tween_callback] or [method tween_method]:
[codeblocks]
[gdscript]
var tween = get_tree().create_tween()
tween.tween_property($Sprite, "modulate", Color.RED, 1)
tween.tween_property($Sprite, "scale", Vector2(), 1)
tween.tween_callback($Sprite.queue_free)
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween();
tween.TweenProperty(GetNode("Sprite"), "modulate", Colors.Red, 1.0f);
tween.TweenProperty(GetNode("Sprite"), "scale", Vector2.Zero, 1.0f);
tween.TweenCallback(Callable.From(GetNode("Sprite").QueueFree));
[/csharp]
[/codeblocks]
This sequence will make the [code]$Sprite[/code] node turn red, then shrink, before finally calling [method Node.queue_free] to free the sprite. [Tweener]s are executed one after another by default. This behavior can be changed using [method parallel] and [method set_parallel].
When a [Tweener] is created with one of the [code]tween_*[/code] methods, a chained method call can be used to tweak the properties of this [Tweener]. For example, if you want to set a different transition type in the above example, you can use [method set_trans]:
[codeblocks]
[gdscript]
var tween = get_tree().create_tween()
tween.tween_property($Sprite, "modulate", Color.RED, 1).set_trans(Tween.TRANS_SINE)
tween.tween_property($Sprite, "scale", Vector2(), 1).set_trans(Tween.TRANS_BOUNCE)
tween.tween_callback($Sprite.queue_free)
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween();
tween.TweenProperty(GetNode("Sprite"), "modulate", Colors.Red, 1.0f).SetTrans(Tween.TransitionType.Sine);
tween.TweenProperty(GetNode("Sprite"), "scale", Vector2.Zero, 1.0f).SetTrans(Tween.TransitionType.Bounce);
tween.TweenCallback(Callable.From(GetNode("Sprite").QueueFree));
[/csharp]
[/codeblocks]
Most of the [Tween] methods can be chained this way too. In the following example the [Tween] is bound to the running script's node and a default transition is set for its [Tweener]s:
[codeblocks]
[gdscript]
var tween = get_tree().create_tween().bind_node(self).set_trans(Tween.TRANS_ELASTIC)
tween.tween_property($Sprite, "modulate", Color.RED, 1)
tween.tween_property($Sprite, "scale", Vector2(), 1)
tween.tween_callback($Sprite.queue_free)
[/gdscript]
[csharp]
var tween = GetTree().CreateTween().BindNode(this).SetTrans(Tween.TransitionType.Elastic);
tween.TweenProperty(GetNode("Sprite"), "modulate", Colors.Red, 1.0f);
tween.TweenProperty(GetNode("Sprite"), "scale", Vector2.Zero, 1.0f);
tween.TweenCallback(Callable.From(GetNode("Sprite").QueueFree));
[/csharp]
[/codeblocks]
Another interesting use for [Tween]s is animating arbitrary sets of objects:
[codeblocks]
[gdscript]
var tween = create_tween()
for sprite in get_children():

	tween.tween_property(sprite, "position", Vector2(0, 0), 1)

[/gdscript]
[csharp]
Tween tween = CreateTween();
foreach (Node sprite in GetChildren())

	tween.TweenProperty(sprite, "position", Vector2.Zero, 1.0f);

[/csharp]
[/codeblocks]
In the example above, all children of a node are moved one after another to position (0, 0).
You should avoid using more than one [Tween] per object's property. If two or more tweens animate one property at the same time, the last one created will take priority and assign the final value. If you want to interrupt and restart an animation, consider assigning the [Tween] to a variable:
[codeblocks]
[gdscript]
var tween
func animate():

	if tween:
	    tween.kill() # Abort the previous animation.
	tween = create_tween()

[/gdscript]
[csharp]
private Tween _tween;

public void Animate()

	{
	    if (_tween != null)
	        _tween.Kill(); // Abort the previous animation
	    _tween = CreateTween();
	}

[/csharp]
[/codeblocks]
Some [Tweener]s use transitions and eases. The first accepts a [enum TransitionType] constant, and refers to the way the timing of the animation is handled (see [url=https://easings.net/]easings.net[/url] for some examples). The second accepts an [enum EaseType] constant, and controls where the [code]trans_type[/code] is applied to the interpolation (in the beginning, the end, or both). If you don't know which transition and easing to pick, you can try different [enum TransitionType] constants with [constant EASE_IN_OUT], and use the one that looks best.
[url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/tween_cheatsheet.webp]Tween easing and transition types cheatsheet[/url]
[b]Note:[/b] Tweens are not designed to be re-used and trying to do so results in an undefined behavior. Create a new Tween for each animation and every time you replay an animation from start. Keep in mind that Tweens start immediately, so only create a Tween when you want to start animating.
[b]Note:[/b] The tween is processed after all of the nodes in the current frame, i.e. node's [method Node._process] method would be called before the tween (or [method Node._physics_process] depending on the value passed to [method set_process_mode]).
*/
type Instance [1]gdclass.Tween

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTween() Instance
}

/*
Creates and appends a [PropertyTweener]. This method tweens a [param property] of an [param object] between an initial value and [param final_val] in a span of time equal to [param duration], in seconds. The initial value by default is the property's value at the time the tweening of the [PropertyTweener] starts.
[b]Example:[/b]
[codeblocks]
[gdscript]
var tween = create_tween()
tween.tween_property($Sprite, "position", Vector2(100, 200), 1)
tween.tween_property($Sprite, "position", Vector2(200, 300), 1)
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(GetNode("Sprite"), "position", new Vector2(100.0f, 200.0f), 1.0f);
tween.TweenProperty(GetNode("Sprite"), "position", new Vector2(200.0f, 300.0f), 1.0f);
[/csharp]
[/codeblocks]
will move the sprite to position (100, 200) and then to (200, 300). If you use [method PropertyTweener.from] or [method PropertyTweener.from_current], the starting position will be overwritten by the given value instead. See other methods in [PropertyTweener] to see how the tweening can be tweaked further.
[b]Note:[/b] You can find the correct property name by hovering over the property in the Inspector. You can also provide the components of a property directly by using [code]"property:component"[/code] (eg. [code]position:x[/code]), where it would only apply to that particular component.
[b]Example:[/b] Moving an object twice from the same position, with different transition types:
[codeblocks]
[gdscript]
var tween = create_tween()
tween.tween_property($Sprite, "position", Vector2.RIGHT * 300, 1).as_relative().set_trans(Tween.TRANS_SINE)
tween.tween_property($Sprite, "position", Vector2.RIGHT * 300, 1).as_relative().from_current().set_trans(Tween.TRANS_EXPO)
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(GetNode("Sprite"), "position", Vector2.Right * 300.0f, 1.0f).AsRelative().SetTrans(Tween.TransitionType.Sine);
tween.TweenProperty(GetNode("Sprite"), "position", Vector2.Right * 300.0f, 1.0f).AsRelative().FromCurrent().SetTrans(Tween.TransitionType.Expo);
[/csharp]
[/codeblocks]
*/
func (self Instance) TweenProperty(obj Object.Instance, property string, final_val any, duration Float.X) [1]gdclass.PropertyTweener { //gd:Tween.tween_property
	return [1]gdclass.PropertyTweener(class(self).TweenProperty(obj, Path.ToNode(String.New(property)), gd.NewVariant(final_val), gd.Float(duration)))
}

/*
Creates and appends an [IntervalTweener]. This method can be used to create delays in the tween animation, as an alternative to using the delay in other [Tweener]s, or when there's no animation (in which case the [Tween] acts as a timer). [param time] is the length of the interval, in seconds.
[b]Example:[/b] Creating an interval in code execution:
[codeblocks]
[gdscript]
# ... some code
await create_tween().tween_interval(2).finished
# ... more code
[/gdscript]
[csharp]
// ... some code
await ToSignal(CreateTween().TweenInterval(2.0f), Tween.SignalName.Finished);
// ... more code
[/csharp]
[/codeblocks]
[b]Example:[/b] Creating an object that moves back and forth and jumps every few seconds:
[codeblocks]
[gdscript]
var tween = create_tween().set_loops()
tween.tween_property($Sprite, "position:x", 200.0, 1).as_relative()
tween.tween_callback(jump)
tween.tween_interval(2)
tween.tween_property($Sprite, "position:x", -200.0, 1).as_relative()
tween.tween_callback(jump)
tween.tween_interval(2)
[/gdscript]
[csharp]
Tween tween = CreateTween().SetLoops();
tween.TweenProperty(GetNode("Sprite"), "position:x", 200.0f, 1.0f).AsRelative();
tween.TweenCallback(Callable.From(Jump));
tween.TweenInterval(2.0f);
tween.TweenProperty(GetNode("Sprite"), "position:x", -200.0f, 1.0f).AsRelative();
tween.TweenCallback(Callable.From(Jump));
tween.TweenInterval(2.0f);
[/csharp]
[/codeblocks]
*/
func (self Instance) TweenInterval(time Float.X) [1]gdclass.IntervalTweener { //gd:Tween.tween_interval
	return [1]gdclass.IntervalTweener(class(self).TweenInterval(gd.Float(time)))
}

/*
Creates and appends a [CallbackTweener]. This method can be used to call an arbitrary method in any object. Use [method Callable.bind] to bind additional arguments for the call.
[b]Example:[/b] Object that keeps shooting every 1 second:
[codeblocks]
[gdscript]
var tween = get_tree().create_tween().set_loops()
tween.tween_callback(shoot).set_delay(1)
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween().SetLoops();
tween.TweenCallback(Callable.From(Shoot)).SetDelay(1.0f);
[/csharp]
[/codeblocks]
[b]Example:[/b] Turning a sprite red and then blue, with 2 second delay:
[codeblocks]
[gdscript]
var tween = get_tree().create_tween()
tween.tween_callback($Sprite.set_modulate.bind(Color.RED)).set_delay(2)
tween.tween_callback($Sprite.set_modulate.bind(Color.BLUE)).set_delay(2)
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween();
Sprite2D sprite = GetNode<Sprite2D>("Sprite");
tween.TweenCallback(Callable.From(() => sprite.Modulate = Colors.Red)).SetDelay(2.0f);
tween.TweenCallback(Callable.From(() => sprite.Modulate = Colors.Blue)).SetDelay(2.0f);
[/csharp]
[/codeblocks]
*/
func (self Instance) TweenCallback(callback func()) [1]gdclass.CallbackTweener { //gd:Tween.tween_callback
	return [1]gdclass.CallbackTweener(class(self).TweenCallback(Callable.New(callback)))
}

/*
Creates and appends a [MethodTweener]. This method is similar to a combination of [method tween_callback] and [method tween_property]. It calls a method over time with a tweened value provided as an argument. The value is tweened between [param from] and [param to] over the time specified by [param duration], in seconds. Use [method Callable.bind] to bind additional arguments for the call. You can use [method MethodTweener.set_ease] and [method MethodTweener.set_trans] to tweak the easing and transition of the value or [method MethodTweener.set_delay] to delay the tweening.
[b]Example:[/b] Making a 3D object look from one point to another point:
[codeblocks]
[gdscript]
var tween = create_tween()
tween.tween_method(look_at.bind(Vector3.UP), Vector3(-1, 0, -1), Vector3(1, 0, -1), 1) # The look_at() method takes up vector as second argument.
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenMethod(Callable.From((Vector3 target) => LookAt(target, Vector3.Up)), new Vector3(-1.0f, 0.0f, -1.0f), new Vector3(1.0f, 0.0f, -1.0f), 1.0f); // Use lambdas to bind additional arguments for the call.
[/csharp]
[/codeblocks]
[b]Example:[/b] Setting the text of a [Label], using an intermediate method and after a delay:
[codeblocks]
[gdscript]
func _ready():

	var tween = create_tween()
	tween.tween_method(set_label_text, 0, 10, 1).set_delay(1)

func set_label_text(value: int):

	$Label.text = "Counting " + str(value)

[/gdscript]
[csharp]
public override void _Ready()

	{
	    base._Ready();

	    Tween tween = CreateTween();
	    tween.TweenMethod(Callable.From<int>(SetLabelText), 0.0f, 10.0f, 1.0f).SetDelay(1.0f);
	}

private void SetLabelText(int value)

	{
	    GetNode<Label>("Label").Text = $"Counting {value}";
	}

[/csharp]
[/codeblocks]
*/
func (self Instance) TweenMethod(method func(value any), from any, to any, duration Float.X) [1]gdclass.MethodTweener { //gd:Tween.tween_method
	return [1]gdclass.MethodTweener(class(self).TweenMethod(Callable.New(method), gd.NewVariant(from), gd.NewVariant(to), gd.Float(duration)))
}

/*
Processes the [Tween] by the given [param delta] value, in seconds. This is mostly useful for manual control when the [Tween] is paused. It can also be used to end the [Tween] animation immediately, by setting [param delta] longer than the whole duration of the [Tween] animation.
Returns [code]true[/code] if the [Tween] still has [Tweener]s that haven't finished.
*/
func (self Instance) CustomStep(delta Float.X) bool { //gd:Tween.custom_step
	return bool(class(self).CustomStep(gd.Float(delta)))
}

/*
Stops the tweening and resets the [Tween] to its initial state. This will not remove any appended [Tweener]s.
[b]Note:[/b] If a Tween is stopped and not bound to any node, it will exist indefinitely until manually started or invalidated. If you lose a reference to such Tween, you can retrieve it using [method SceneTree.get_processed_tweens].
*/
func (self Instance) Stop() { //gd:Tween.stop
	class(self).Stop()
}

/*
Pauses the tweening. The animation can be resumed by using [method play].
[b]Note:[/b] If a Tween is paused and not bound to any node, it will exist indefinitely until manually started or invalidated. If you lose a reference to such Tween, you can retrieve it using [method SceneTree.get_processed_tweens].
*/
func (self Instance) Pause() { //gd:Tween.pause
	class(self).Pause()
}

/*
Resumes a paused or stopped [Tween].
*/
func (self Instance) Play() { //gd:Tween.play
	class(self).Play()
}

/*
Aborts all tweening operations and invalidates the [Tween].
*/
func (self Instance) Kill() { //gd:Tween.kill
	class(self).Kill()
}

/*
Returns the total time in seconds the [Tween] has been animating (i.e. the time since it started, not counting pauses etc.). The time is affected by [method set_speed_scale], and [method stop] will reset it to [code]0[/code].
[b]Note:[/b] As it results from accumulating frame deltas, the time returned after the [Tween] has finished animating will be slightly greater than the actual [Tween] duration.
*/
func (self Instance) GetTotalElapsedTime() Float.X { //gd:Tween.get_total_elapsed_time
	return Float.X(Float.X(class(self).GetTotalElapsedTime()))
}

/*
Returns whether the [Tween] is currently running, i.e. it wasn't paused and it's not finished.
*/
func (self Instance) IsRunning() bool { //gd:Tween.is_running
	return bool(class(self).IsRunning())
}

/*
Returns whether the [Tween] is valid. A valid [Tween] is a [Tween] contained by the scene tree (i.e. the array from [method SceneTree.get_processed_tweens] will contain this [Tween]). A [Tween] might become invalid when it has finished tweening, is killed, or when created with [code]Tween.new()[/code]. Invalid [Tween]s can't have [Tweener]s appended.
*/
func (self Instance) IsValid() bool { //gd:Tween.is_valid
	return bool(class(self).IsValid())
}

/*
Binds this [Tween] with the given [param node]. [Tween]s are processed directly by the [SceneTree], so they run independently of the animated nodes. When you bind a [Node] with the [Tween], the [Tween] will halt the animation when the object is not inside tree and the [Tween] will be automatically killed when the bound object is freed. Also [constant TWEEN_PAUSE_BOUND] will make the pausing behavior dependent on the bound node.
For a shorter way to create and bind a [Tween], you can use [method Node.create_tween].
*/
func (self Instance) BindNode(node [1]gdclass.Node) [1]gdclass.Tween { //gd:Tween.bind_node
	return [1]gdclass.Tween(class(self).BindNode(node))
}

/*
Determines whether the [Tween] should run after process frames (see [method Node._process]) or physics frames (see [method Node._physics_process]).
Default value is [constant TWEEN_PROCESS_IDLE].
*/
func (self Instance) SetProcessMode(mode gdclass.TweenTweenProcessMode) [1]gdclass.Tween { //gd:Tween.set_process_mode
	return [1]gdclass.Tween(class(self).SetProcessMode(mode))
}

/*
Determines the behavior of the [Tween] when the [SceneTree] is paused. Check [enum TweenPauseMode] for options.
Default value is [constant TWEEN_PAUSE_BOUND].
*/
func (self Instance) SetPauseMode(mode gdclass.TweenTweenPauseMode) [1]gdclass.Tween { //gd:Tween.set_pause_mode
	return [1]gdclass.Tween(class(self).SetPauseMode(mode))
}

/*
If [param parallel] is [code]true[/code], the [Tweener]s appended after this method will by default run simultaneously, as opposed to sequentially.
[b]Note:[/b] Just like with [method parallel], the tweener added right before this method will also be part of the parallel step.
[codeblock]
tween.tween_property(self, "position", Vector2(300, 0), 0.5)
tween.set_parallel()
tween.tween_property(self, "modulate", Color.GREEN, 0.5) # Runs together with the position tweener.
[/codeblock]
*/
func (self Instance) SetParallel() [1]gdclass.Tween { //gd:Tween.set_parallel
	return [1]gdclass.Tween(class(self).SetParallel(true))
}

/*
Sets the number of times the tweening sequence will be repeated, i.e. [code]set_loops(2)[/code] will run the animation twice.
Calling this method without arguments will make the [Tween] run infinitely, until either it is killed with [method kill], the [Tween]'s bound node is freed, or all the animated objects have been freed (which makes further animation impossible).
[b]Warning:[/b] Make sure to always add some duration/delay when using infinite loops. To prevent the game freezing, 0-duration looped animations (e.g. a single [CallbackTweener] with no delay) are stopped after a small number of loops, which may produce unexpected results. If a [Tween]'s lifetime depends on some node, always use [method bind_node].
*/
func (self Instance) SetLoops() [1]gdclass.Tween { //gd:Tween.set_loops
	return [1]gdclass.Tween(class(self).SetLoops(gd.Int(0)))
}

/*
Returns the number of remaining loops for this [Tween] (see [method set_loops]). A return value of [code]-1[/code] indicates an infinitely looping [Tween], and a return value of [code]0[/code] indicates that the [Tween] has already finished.
*/
func (self Instance) GetLoopsLeft() int { //gd:Tween.get_loops_left
	return int(int(class(self).GetLoopsLeft()))
}

/*
Scales the speed of tweening. This affects all [Tweener]s and their delays.
*/
func (self Instance) SetSpeedScale(speed Float.X) [1]gdclass.Tween { //gd:Tween.set_speed_scale
	return [1]gdclass.Tween(class(self).SetSpeedScale(gd.Float(speed)))
}

/*
Sets the default transition type for [PropertyTweener]s and [MethodTweener]s animated by this [Tween].
If not specified, the default value is [constant TRANS_LINEAR].
*/
func (self Instance) SetTrans(trans gdclass.TweenTransitionType) [1]gdclass.Tween { //gd:Tween.set_trans
	return [1]gdclass.Tween(class(self).SetTrans(trans))
}

/*
Sets the default ease type for [PropertyTweener]s and [MethodTweener]s animated by this [Tween].
If not specified, the default value is [constant EASE_IN_OUT].
*/
func (self Instance) SetEase(ease gdclass.TweenEaseType) [1]gdclass.Tween { //gd:Tween.set_ease
	return [1]gdclass.Tween(class(self).SetEase(ease))
}

/*
Makes the next [Tweener] run parallelly to the previous one.
[b]Example:[/b]
[codeblocks]
[gdscript]
var tween = create_tween()
tween.tween_property(...)
tween.parallel().tween_property(...)
tween.parallel().tween_property(...)
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(...);
tween.Parallel().TweenProperty(...);
tween.Parallel().TweenProperty(...);
[/csharp]
[/codeblocks]
All [Tweener]s in the example will run at the same time.
You can make the [Tween] parallel by default by using [method set_parallel].
*/
func (self Instance) Parallel() [1]gdclass.Tween { //gd:Tween.parallel
	return [1]gdclass.Tween(class(self).Parallel())
}

/*
Used to chain two [Tweener]s after [method set_parallel] is called with [code]true[/code].
[codeblocks]
[gdscript]
var tween = create_tween().set_parallel(true)
tween.tween_property(...)
tween.tween_property(...) # Will run parallelly with above.
tween.chain().tween_property(...) # Will run after two above are finished.
[/gdscript]
[csharp]
Tween tween = CreateTween().SetParallel(true);
tween.TweenProperty(...);
tween.TweenProperty(...); // Will run parallelly with above.
tween.Chain().TweenProperty(...); // Will run after two above are finished.
[/csharp]
[/codeblocks]
*/
func (self Instance) Chain() [1]gdclass.Tween { //gd:Tween.chain
	return [1]gdclass.Tween(class(self).Chain())
}

/*
This method can be used for manual interpolation of a value, when you don't want [Tween] to do animating for you. It's similar to [method @GlobalScope.lerp], but with support for custom transition and easing.
[param initial_value] is the starting value of the interpolation.
[param delta_value] is the change of the value in the interpolation, i.e. it's equal to [code]final_value - initial_value[/code].
[param elapsed_time] is the time in seconds that passed after the interpolation started and it's used to control the position of the interpolation. E.g. when it's equal to half of the [param duration], the interpolated value will be halfway between initial and final values. This value can also be greater than [param duration] or lower than 0, which will extrapolate the value.
[param duration] is the total time of the interpolation.
[b]Note:[/b] If [param duration] is equal to [code]0[/code], the method will always return the final value, regardless of [param elapsed_time] provided.
*/
func InterpolateValue(initial_value any, delta_value any, elapsed_time Float.X, duration Float.X, trans_type gdclass.TweenTransitionType, ease_type gdclass.TweenEaseType) any { //gd:Tween.interpolate_value
	self := Instance{}
	return any(class(self).InterpolateValue(gd.NewVariant(initial_value), gd.NewVariant(delta_value), gd.Float(elapsed_time), gd.Float(duration), trans_type, ease_type).Interface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Tween

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Tween"))
	casted := Instance{*(*gdclass.Tween)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates and appends a [PropertyTweener]. This method tweens a [param property] of an [param object] between an initial value and [param final_val] in a span of time equal to [param duration], in seconds. The initial value by default is the property's value at the time the tweening of the [PropertyTweener] starts.
[b]Example:[/b]
[codeblocks]
[gdscript]
var tween = create_tween()
tween.tween_property($Sprite, "position", Vector2(100, 200), 1)
tween.tween_property($Sprite, "position", Vector2(200, 300), 1)
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(GetNode("Sprite"), "position", new Vector2(100.0f, 200.0f), 1.0f);
tween.TweenProperty(GetNode("Sprite"), "position", new Vector2(200.0f, 300.0f), 1.0f);
[/csharp]
[/codeblocks]
will move the sprite to position (100, 200) and then to (200, 300). If you use [method PropertyTweener.from] or [method PropertyTweener.from_current], the starting position will be overwritten by the given value instead. See other methods in [PropertyTweener] to see how the tweening can be tweaked further.
[b]Note:[/b] You can find the correct property name by hovering over the property in the Inspector. You can also provide the components of a property directly by using [code]"property:component"[/code] (eg. [code]position:x[/code]), where it would only apply to that particular component.
[b]Example:[/b] Moving an object twice from the same position, with different transition types:
[codeblocks]
[gdscript]
var tween = create_tween()
tween.tween_property($Sprite, "position", Vector2.RIGHT * 300, 1).as_relative().set_trans(Tween.TRANS_SINE)
tween.tween_property($Sprite, "position", Vector2.RIGHT * 300, 1).as_relative().from_current().set_trans(Tween.TRANS_EXPO)
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(GetNode("Sprite"), "position", Vector2.Right * 300.0f, 1.0f).AsRelative().SetTrans(Tween.TransitionType.Sine);
tween.TweenProperty(GetNode("Sprite"), "position", Vector2.Right * 300.0f, 1.0f).AsRelative().FromCurrent().SetTrans(Tween.TransitionType.Expo);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) TweenProperty(obj [1]gd.Object, property Path.ToNode, final_val gd.Variant, duration gd.Float) [1]gdclass.PropertyTweener { //gd:Tween.tween_property
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(property)))
	callframe.Arg(frame, pointers.Get(final_val))
	callframe.Arg(frame, duration)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_tween_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PropertyTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.PropertyTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates and appends an [IntervalTweener]. This method can be used to create delays in the tween animation, as an alternative to using the delay in other [Tweener]s, or when there's no animation (in which case the [Tween] acts as a timer). [param time] is the length of the interval, in seconds.
[b]Example:[/b] Creating an interval in code execution:
[codeblocks]
[gdscript]
# ... some code
await create_tween().tween_interval(2).finished
# ... more code
[/gdscript]
[csharp]
// ... some code
await ToSignal(CreateTween().TweenInterval(2.0f), Tween.SignalName.Finished);
// ... more code
[/csharp]
[/codeblocks]
[b]Example:[/b] Creating an object that moves back and forth and jumps every few seconds:
[codeblocks]
[gdscript]
var tween = create_tween().set_loops()
tween.tween_property($Sprite, "position:x", 200.0, 1).as_relative()
tween.tween_callback(jump)
tween.tween_interval(2)
tween.tween_property($Sprite, "position:x", -200.0, 1).as_relative()
tween.tween_callback(jump)
tween.tween_interval(2)
[/gdscript]
[csharp]
Tween tween = CreateTween().SetLoops();
tween.TweenProperty(GetNode("Sprite"), "position:x", 200.0f, 1.0f).AsRelative();
tween.TweenCallback(Callable.From(Jump));
tween.TweenInterval(2.0f);
tween.TweenProperty(GetNode("Sprite"), "position:x", -200.0f, 1.0f).AsRelative();
tween.TweenCallback(Callable.From(Jump));
tween.TweenInterval(2.0f);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) TweenInterval(time gd.Float) [1]gdclass.IntervalTweener { //gd:Tween.tween_interval
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_tween_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.IntervalTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.IntervalTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates and appends a [CallbackTweener]. This method can be used to call an arbitrary method in any object. Use [method Callable.bind] to bind additional arguments for the call.
[b]Example:[/b] Object that keeps shooting every 1 second:
[codeblocks]
[gdscript]
var tween = get_tree().create_tween().set_loops()
tween.tween_callback(shoot).set_delay(1)
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween().SetLoops();
tween.TweenCallback(Callable.From(Shoot)).SetDelay(1.0f);
[/csharp]
[/codeblocks]
[b]Example:[/b] Turning a sprite red and then blue, with 2 second delay:
[codeblocks]
[gdscript]
var tween = get_tree().create_tween()
tween.tween_callback($Sprite.set_modulate.bind(Color.RED)).set_delay(2)
tween.tween_callback($Sprite.set_modulate.bind(Color.BLUE)).set_delay(2)
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween();
Sprite2D sprite = GetNode<Sprite2D>("Sprite");
tween.TweenCallback(Callable.From(() => sprite.Modulate = Colors.Red)).SetDelay(2.0f);
tween.TweenCallback(Callable.From(() => sprite.Modulate = Colors.Blue)).SetDelay(2.0f);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) TweenCallback(callback Callable.Function) [1]gdclass.CallbackTweener { //gd:Tween.tween_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_tween_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CallbackTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.CallbackTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates and appends a [MethodTweener]. This method is similar to a combination of [method tween_callback] and [method tween_property]. It calls a method over time with a tweened value provided as an argument. The value is tweened between [param from] and [param to] over the time specified by [param duration], in seconds. Use [method Callable.bind] to bind additional arguments for the call. You can use [method MethodTweener.set_ease] and [method MethodTweener.set_trans] to tweak the easing and transition of the value or [method MethodTweener.set_delay] to delay the tweening.
[b]Example:[/b] Making a 3D object look from one point to another point:
[codeblocks]
[gdscript]
var tween = create_tween()
tween.tween_method(look_at.bind(Vector3.UP), Vector3(-1, 0, -1), Vector3(1, 0, -1), 1) # The look_at() method takes up vector as second argument.
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenMethod(Callable.From((Vector3 target) => LookAt(target, Vector3.Up)), new Vector3(-1.0f, 0.0f, -1.0f), new Vector3(1.0f, 0.0f, -1.0f), 1.0f); // Use lambdas to bind additional arguments for the call.
[/csharp]
[/codeblocks]
[b]Example:[/b] Setting the text of a [Label], using an intermediate method and after a delay:
[codeblocks]
[gdscript]
func _ready():
    var tween = create_tween()
    tween.tween_method(set_label_text, 0, 10, 1).set_delay(1)

func set_label_text(value: int):
    $Label.text = "Counting " + str(value)
[/gdscript]
[csharp]
public override void _Ready()
{
    base._Ready();

    Tween tween = CreateTween();
    tween.TweenMethod(Callable.From<int>(SetLabelText), 0.0f, 10.0f, 1.0f).SetDelay(1.0f);
}

private void SetLabelText(int value)
{
    GetNode<Label>("Label").Text = $"Counting {value}";
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) TweenMethod(method Callable.Function, from gd.Variant, to gd.Variant, duration gd.Float) [1]gdclass.MethodTweener { //gd:Tween.tween_method
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(method)))
	callframe.Arg(frame, pointers.Get(from))
	callframe.Arg(frame, pointers.Get(to))
	callframe.Arg(frame, duration)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_tween_method, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.MethodTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.MethodTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Processes the [Tween] by the given [param delta] value, in seconds. This is mostly useful for manual control when the [Tween] is paused. It can also be used to end the [Tween] animation immediately, by setting [param delta] longer than the whole duration of the [Tween] animation.
Returns [code]true[/code] if the [Tween] still has [Tweener]s that haven't finished.
*/
//go:nosplit
func (self class) CustomStep(delta gd.Float) bool { //gd:Tween.custom_step
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_custom_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Stops the tweening and resets the [Tween] to its initial state. This will not remove any appended [Tweener]s.
[b]Note:[/b] If a Tween is stopped and not bound to any node, it will exist indefinitely until manually started or invalidated. If you lose a reference to such Tween, you can retrieve it using [method SceneTree.get_processed_tweens].
*/
//go:nosplit
func (self class) Stop() { //gd:Tween.stop
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Pauses the tweening. The animation can be resumed by using [method play].
[b]Note:[/b] If a Tween is paused and not bound to any node, it will exist indefinitely until manually started or invalidated. If you lose a reference to such Tween, you can retrieve it using [method SceneTree.get_processed_tweens].
*/
//go:nosplit
func (self class) Pause() { //gd:Tween.pause
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_pause, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Resumes a paused or stopped [Tween].
*/
//go:nosplit
func (self class) Play() { //gd:Tween.play
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_play, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Aborts all tweening operations and invalidates the [Tween].
*/
//go:nosplit
func (self class) Kill() { //gd:Tween.kill
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_kill, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the total time in seconds the [Tween] has been animating (i.e. the time since it started, not counting pauses etc.). The time is affected by [method set_speed_scale], and [method stop] will reset it to [code]0[/code].
[b]Note:[/b] As it results from accumulating frame deltas, the time returned after the [Tween] has finished animating will be slightly greater than the actual [Tween] duration.
*/
//go:nosplit
func (self class) GetTotalElapsedTime() gd.Float { //gd:Tween.get_total_elapsed_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_get_total_elapsed_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the [Tween] is currently running, i.e. it wasn't paused and it's not finished.
*/
//go:nosplit
func (self class) IsRunning() bool { //gd:Tween.is_running
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_is_running, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the [Tween] is valid. A valid [Tween] is a [Tween] contained by the scene tree (i.e. the array from [method SceneTree.get_processed_tweens] will contain this [Tween]). A [Tween] might become invalid when it has finished tweening, is killed, or when created with [code]Tween.new()[/code]. Invalid [Tween]s can't have [Tweener]s appended.
*/
//go:nosplit
func (self class) IsValid() bool { //gd:Tween.is_valid
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_is_valid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Binds this [Tween] with the given [param node]. [Tween]s are processed directly by the [SceneTree], so they run independently of the animated nodes. When you bind a [Node] with the [Tween], the [Tween] will halt the animation when the object is not inside tree and the [Tween] will be automatically killed when the bound object is freed. Also [constant TWEEN_PAUSE_BOUND] will make the pausing behavior dependent on the bound node.
For a shorter way to create and bind a [Tween], you can use [method Node.create_tween].
*/
//go:nosplit
func (self class) BindNode(node [1]gdclass.Node) [1]gdclass.Tween { //gd:Tween.bind_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_bind_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Determines whether the [Tween] should run after process frames (see [method Node._process]) or physics frames (see [method Node._physics_process]).
Default value is [constant TWEEN_PROCESS_IDLE].
*/
//go:nosplit
func (self class) SetProcessMode(mode gdclass.TweenTweenProcessMode) [1]gdclass.Tween { //gd:Tween.set_process_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_set_process_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Determines the behavior of the [Tween] when the [SceneTree] is paused. Check [enum TweenPauseMode] for options.
Default value is [constant TWEEN_PAUSE_BOUND].
*/
//go:nosplit
func (self class) SetPauseMode(mode gdclass.TweenTweenPauseMode) [1]gdclass.Tween { //gd:Tween.set_pause_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_set_pause_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
If [param parallel] is [code]true[/code], the [Tweener]s appended after this method will by default run simultaneously, as opposed to sequentially.
[b]Note:[/b] Just like with [method parallel], the tweener added right before this method will also be part of the parallel step.
[codeblock]
tween.tween_property(self, "position", Vector2(300, 0), 0.5)
tween.set_parallel()
tween.tween_property(self, "modulate", Color.GREEN, 0.5) # Runs together with the position tweener.
[/codeblock]
*/
//go:nosplit
func (self class) SetParallel(parallel bool) [1]gdclass.Tween { //gd:Tween.set_parallel
	var frame = callframe.New()
	callframe.Arg(frame, parallel)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_set_parallel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the number of times the tweening sequence will be repeated, i.e. [code]set_loops(2)[/code] will run the animation twice.
Calling this method without arguments will make the [Tween] run infinitely, until either it is killed with [method kill], the [Tween]'s bound node is freed, or all the animated objects have been freed (which makes further animation impossible).
[b]Warning:[/b] Make sure to always add some duration/delay when using infinite loops. To prevent the game freezing, 0-duration looped animations (e.g. a single [CallbackTweener] with no delay) are stopped after a small number of loops, which may produce unexpected results. If a [Tween]'s lifetime depends on some node, always use [method bind_node].
*/
//go:nosplit
func (self class) SetLoops(loops gd.Int) [1]gdclass.Tween { //gd:Tween.set_loops
	var frame = callframe.New()
	callframe.Arg(frame, loops)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_set_loops, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the number of remaining loops for this [Tween] (see [method set_loops]). A return value of [code]-1[/code] indicates an infinitely looping [Tween], and a return value of [code]0[/code] indicates that the [Tween] has already finished.
*/
//go:nosplit
func (self class) GetLoopsLeft() gd.Int { //gd:Tween.get_loops_left
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_get_loops_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Scales the speed of tweening. This affects all [Tweener]s and their delays.
*/
//go:nosplit
func (self class) SetSpeedScale(speed gd.Float) [1]gdclass.Tween { //gd:Tween.set_speed_scale
	var frame = callframe.New()
	callframe.Arg(frame, speed)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the default transition type for [PropertyTweener]s and [MethodTweener]s animated by this [Tween].
If not specified, the default value is [constant TRANS_LINEAR].
*/
//go:nosplit
func (self class) SetTrans(trans gdclass.TweenTransitionType) [1]gdclass.Tween { //gd:Tween.set_trans
	var frame = callframe.New()
	callframe.Arg(frame, trans)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_set_trans, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the default ease type for [PropertyTweener]s and [MethodTweener]s animated by this [Tween].
If not specified, the default value is [constant EASE_IN_OUT].
*/
//go:nosplit
func (self class) SetEase(ease gdclass.TweenEaseType) [1]gdclass.Tween { //gd:Tween.set_ease
	var frame = callframe.New()
	callframe.Arg(frame, ease)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_set_ease, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Makes the next [Tweener] run parallelly to the previous one.
[b]Example:[/b]
[codeblocks]
[gdscript]
var tween = create_tween()
tween.tween_property(...)
tween.parallel().tween_property(...)
tween.parallel().tween_property(...)
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(...);
tween.Parallel().TweenProperty(...);
tween.Parallel().TweenProperty(...);
[/csharp]
[/codeblocks]
All [Tweener]s in the example will run at the same time.
You can make the [Tween] parallel by default by using [method set_parallel].
*/
//go:nosplit
func (self class) Parallel() [1]gdclass.Tween { //gd:Tween.parallel
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_parallel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Used to chain two [Tweener]s after [method set_parallel] is called with [code]true[/code].
[codeblocks]
[gdscript]
var tween = create_tween().set_parallel(true)
tween.tween_property(...)
tween.tween_property(...) # Will run parallelly with above.
tween.chain().tween_property(...) # Will run after two above are finished.
[/gdscript]
[csharp]
Tween tween = CreateTween().SetParallel(true);
tween.TweenProperty(...);
tween.TweenProperty(...); // Will run parallelly with above.
tween.Chain().TweenProperty(...); // Will run after two above are finished.
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) Chain() [1]gdclass.Tween { //gd:Tween.chain
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_chain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tween{gd.PointerWithOwnershipTransferredToGo[gdclass.Tween](r_ret.Get())}
	frame.Free()
	return ret
}

/*
This method can be used for manual interpolation of a value, when you don't want [Tween] to do animating for you. It's similar to [method @GlobalScope.lerp], but with support for custom transition and easing.
[param initial_value] is the starting value of the interpolation.
[param delta_value] is the change of the value in the interpolation, i.e. it's equal to [code]final_value - initial_value[/code].
[param elapsed_time] is the time in seconds that passed after the interpolation started and it's used to control the position of the interpolation. E.g. when it's equal to half of the [param duration], the interpolated value will be halfway between initial and final values. This value can also be greater than [param duration] or lower than 0, which will extrapolate the value.
[param duration] is the total time of the interpolation.
[b]Note:[/b] If [param duration] is equal to [code]0[/code], the method will always return the final value, regardless of [param elapsed_time] provided.
*/
//go:nosplit
func (self class) InterpolateValue(initial_value gd.Variant, delta_value gd.Variant, elapsed_time gd.Float, duration gd.Float, trans_type gdclass.TweenTransitionType, ease_type gdclass.TweenEaseType) gd.Variant { //gd:Tween.interpolate_value
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(initial_value))
	callframe.Arg(frame, pointers.Get(delta_value))
	callframe.Arg(frame, elapsed_time)
	callframe.Arg(frame, duration)
	callframe.Arg(frame, trans_type)
	callframe.Arg(frame, ease_type)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tween.Bind_interpolate_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnStepFinished(cb func(idx int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("step_finished"), gd.NewCallable(cb), 0)
}

func (self Instance) OnLoopFinished(cb func(loop_count int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("loop_finished"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsTween() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTween() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("Tween", func(ptr gd.Object) any { return [1]gdclass.Tween{*(*gdclass.Tween)(unsafe.Pointer(&ptr))} })
}

type TweenProcessMode = gdclass.TweenTweenProcessMode //gd:Tween.TweenProcessMode

const (
	/*The [Tween] updates after each physics frame (see [method Node._physics_process]).*/
	TweenProcessPhysics TweenProcessMode = 0
	/*The [Tween] updates after each process frame (see [method Node._process]).*/
	TweenProcessIdle TweenProcessMode = 1
)

type TweenPauseMode = gdclass.TweenTweenPauseMode //gd:Tween.TweenPauseMode

const (
	/*If the [Tween] has a bound node, it will process when that node can process (see [member Node.process_mode]). Otherwise it's the same as [constant TWEEN_PAUSE_STOP].*/
	TweenPauseBound TweenPauseMode = 0
	/*If [SceneTree] is paused, the [Tween] will also pause.*/
	TweenPauseStop TweenPauseMode = 1
	/*The [Tween] will process regardless of whether [SceneTree] is paused.*/
	TweenPauseProcess TweenPauseMode = 2
)

type TransitionType = gdclass.TweenTransitionType //gd:Tween.TransitionType

const (
	/*The animation is interpolated linearly.*/
	TransLinear TransitionType = 0
	/*The animation is interpolated using a sine function.*/
	TransSine TransitionType = 1
	/*The animation is interpolated with a quintic (to the power of 5) function.*/
	TransQuint TransitionType = 2
	/*The animation is interpolated with a quartic (to the power of 4) function.*/
	TransQuart TransitionType = 3
	/*The animation is interpolated with a quadratic (to the power of 2) function.*/
	TransQuad TransitionType = 4
	/*The animation is interpolated with an exponential (to the power of x) function.*/
	TransExpo TransitionType = 5
	/*The animation is interpolated with elasticity, wiggling around the edges.*/
	TransElastic TransitionType = 6
	/*The animation is interpolated with a cubic (to the power of 3) function.*/
	TransCubic TransitionType = 7
	/*The animation is interpolated with a function using square roots.*/
	TransCirc TransitionType = 8
	/*The animation is interpolated by bouncing at the end.*/
	TransBounce TransitionType = 9
	/*The animation is interpolated backing out at ends.*/
	TransBack TransitionType = 10
	/*The animation is interpolated like a spring towards the end.*/
	TransSpring TransitionType = 11
)

type EaseType = gdclass.TweenEaseType //gd:Tween.EaseType

const (
	/*The interpolation starts slowly and speeds up towards the end.*/
	EaseIn EaseType = 0
	/*The interpolation starts quickly and slows down towards the end.*/
	EaseOut EaseType = 1
	/*A combination of [constant EASE_IN] and [constant EASE_OUT]. The interpolation is slowest at both ends.*/
	EaseInOut EaseType = 2
	/*A combination of [constant EASE_IN] and [constant EASE_OUT]. The interpolation is fastest at both ends.*/
	EaseOutIn EaseType = 3
)
