package AudioStreamGenerator

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioStream"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.AudioStreamGenerator
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamGenerator
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStreamGenerator"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) MixRate() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMixRate()))
}

func (self Go) SetMixRate(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMixRate(gd.Float(value))
}

func (self Go) BufferLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBufferLength()))
}

func (self Go) SetBufferLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBufferLength(gd.Float(value))
}

//go:nosplit
func (self class) SetMixRate(hz gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGenerator.Bind_set_mix_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMixRate() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGenerator.Bind_get_mix_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBufferLength(seconds gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGenerator.Bind_set_buffer_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBufferLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGenerator.Bind_get_buffer_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamGenerator() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamGenerator() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioStream() AudioStream.GD { return *((*AudioStream.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStream() AudioStream.Go { return *((*AudioStream.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStream(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStream(), name)
	}
}
func init() {classdb.Register("AudioStreamGenerator", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
