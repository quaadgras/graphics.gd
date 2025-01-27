// Package AudioStreamGenerator provides methods for working with AudioStreamGenerator object instances.
package AudioStreamGenerator

import "unsafe"
import "reflect"
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
import "graphics.gd/classdb/AudioStream"
import "graphics.gd/classdb/Resource"
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

/*
[AudioStreamGenerator] is a type of audio stream that does not play back sounds on its own; instead, it expects a script to generate audio data for it. See also [AudioStreamGeneratorPlayback].
Here's a sample on how to use it to generate a sine wave:
[codeblocks]
[gdscript]
var playback # Will hold the AudioStreamGeneratorPlayback.
@onready var sample_hz = $AudioStreamPlayer.stream.mix_rate
var pulse_hz = 440.0 # The frequency of the sound wave.

func _ready():

	$AudioStreamPlayer.play()
	playback = $AudioStreamPlayer.get_stream_playback()
	fill_buffer()

func fill_buffer():

	var phase = 0.0
	var increment = pulse_hz / sample_hz
	var frames_available = playback.get_frames_available()

	for i in range(frames_available):
	    playback.push_frame(Vector2.ONE * sin(phase * TAU))
	    phase = fmod(phase + increment, 1.0)

[/gdscript]
[csharp]
[Export] public AudioStreamPlayer Player { get; set; }

private AudioStreamGeneratorPlayback _playback; // Will hold the AudioStreamGeneratorPlayback.
private float _sampleHz;
private float _pulseHz = 440.0f; // The frequency of the sound wave.

public override void _Ready()

	{
	    if (Player.Stream is AudioStreamGenerator generator) // Type as a generator to access MixRate.
	    {
	        _sampleHz = generator.MixRate;
	        Player.Play();
	        _playback = (AudioStreamGeneratorPlayback)Player.GetStreamPlayback();
	        FillBuffer();
	    }
	}

public void FillBuffer()

	{
	    double phase = 0.0;
	    float increment = _pulseHz / _sampleHz;
	    int framesAvailable = _playback.GetFramesAvailable();

	    for (int i = 0; i < framesAvailable; i++)
	    {
	        _playback.PushFrame(Vector2.One * (float)Mathf.Sin(phase * Mathf.Tau));
	        phase = Mathf.PosMod(phase + increment, 1.0);
	    }
	}

[/csharp]
[/codeblocks]
In the example above, the "AudioStreamPlayer" node must use an [AudioStreamGenerator] as its stream. The [code]fill_buffer[/code] function provides audio data for approximating a sine wave.
See also [AudioEffectSpectrumAnalyzer] for performing real-time audio spectrum analysis.
[b]Note:[/b] Due to performance constraints, this class is best used from C# or from a compiled language via GDExtension. If you still want to use this class from GDScript, consider using a lower [member mix_rate] such as 11,025 Hz or 22,050 Hz.
*/
type Instance [1]gdclass.AudioStreamGenerator

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamGenerator() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamGenerator

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamGenerator"))
	casted := Instance{*(*gdclass.AudioStreamGenerator)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) MixRate() Float.X {
	return Float.X(Float.X(class(self).GetMixRate()))
}

func (self Instance) SetMixRate(value Float.X) {
	class(self).SetMixRate(gd.Float(value))
}

func (self Instance) BufferLength() Float.X {
	return Float.X(Float.X(class(self).GetBufferLength()))
}

func (self Instance) SetBufferLength(value Float.X) {
	class(self).SetBufferLength(gd.Float(value))
}

//go:nosplit
func (self class) SetMixRate(hz gd.Float) { //gd:AudioStreamGenerator.set_mix_rate
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGenerator.Bind_set_mix_rate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMixRate() gd.Float { //gd:AudioStreamGenerator.get_mix_rate
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGenerator.Bind_get_mix_rate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBufferLength(seconds gd.Float) { //gd:AudioStreamGenerator.set_buffer_length
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGenerator.Bind_set_buffer_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBufferLength() gd.Float { //gd:AudioStreamGenerator.get_buffer_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGenerator.Bind_get_buffer_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamGenerator() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamGenerator() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAudioStream() AudioStream.Advanced {
	return *((*AudioStream.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStream() AudioStream.Instance {
	return *((*AudioStream.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Advanced(self.AsAudioStream()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Instance(self.AsAudioStream()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamGenerator", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamGenerator{*(*gdclass.AudioStreamGenerator)(unsafe.Pointer(&ptr))}
	})
}
