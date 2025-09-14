/*
@onready var audio_player = $AudioStreamPlayer

func _ready():
    get_window().files_dropped.connect(_on_files_dropped)

func _on_files_dropped(files):
    if files[0].get_extension() == "wav":
        audio_player.stream = AudioStreamWAV.load_from_file(files[0], {
                "force/max_rate": true,
                "force/max_rate_hz": 11025
            })
        audio_player.play()
*/

package main

import (
	"path/filepath"

	"graphics.gd/classdb/AudioStreamPlayer"
	"graphics.gd/classdb/AudioStreamWAV"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Window"
)

type MyAudioLoader struct {
	Node.Extension[MyAudioLoader]

	AudioPlayer AudioStreamPlayer.Instance
}

func (m *MyAudioLoader) Ready() {
	Window.Get(m.AsNode()).OnFilesDropped(func(files []string) {
		if filepath.Ext(files[0]) == "wav" {
			m.AudioPlayer.SetStream(AudioStreamWAV.LoadFromFile(files[0], AudioStreamWAV.Options{
				ForceMaxRate:   true,
				ForceMaxRateHz: 11025,
			}).AsAudioStream())
			m.AudioPlayer.Play()
		}
	})
}
