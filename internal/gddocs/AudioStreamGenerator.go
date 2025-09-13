/*
[gdscript]
var playback # Will hold the AudioStreamGeneratorPlayback.
@onready var sample_hz = $AudioStreamPlayer.stream.mix_rate
var pulse_hz = 440.0 # The frequency of the sound wave.
var phase = 0.0

func _ready():
    $AudioStreamPlayer.play()
    playback = $AudioStreamPlayer.get_stream_playback()
    fill_buffer()

func fill_buffer():
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
private double phase = 0.0;

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
    float increment = _pulseHz / _sampleHz;
    int framesAvailable = _playback.GetFramesAvailable();

    for (int i = 0; i < framesAvailable; i++)
    {
        _playback.PushFrame(Vector2.One * (float)Mathf.Sin(phase * Mathf.Tau));
        phase = Mathf.PosMod(phase + increment, 1.0);
    }
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AudioStreamGenerator"
	"graphics.gd/classdb/AudioStreamGeneratorPlayback"
	"graphics.gd/classdb/AudioStreamPlayer"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

type ExampleAudioStreamGenerator struct {
	Node.Extension[ExampleAudioStreamGenerator]

	AudioStreamPlayer AudioStreamPlayer.Instance

	playback AudioStreamGeneratorPlayback.Instance

	sample_hz, pulse_hz Float.X
	phase               Angle.Radians
}

func (eg *ExampleAudioStreamGenerator) Ready() {
	eg.sample_hz = Object.To[AudioStreamGenerator.Instance](eg.AudioStreamPlayer.Stream()).AsAudioStreamGenerator().MixRate()
	eg.AudioStreamPlayer.Play()
	eg.playback = Object.To[AudioStreamGeneratorPlayback.Instance](eg.AudioStreamPlayer.GetStreamPlayback())
	eg.FillBuffer()
}

func (eg *ExampleAudioStreamGenerator) FillBuffer() {
	increment := Angle.Radians(eg.pulse_hz / eg.sample_hz)
	frames_available := eg.playback.GetFramesAvailable()
	for range frames_available {
		eg.playback.PushFrame(Vector2.AddX(Vector2.One, Angle.Sin(eg.phase*Angle.Tau)))
		eg.phase = Float.Mod(eg.phase+increment, 1.0)
	}
}
