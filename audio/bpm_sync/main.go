package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioServer"
	"graphics.gd/classdb/AudioStreamPlayer"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/Time"
	"graphics.gd/startup"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Float"
)

type SyncSource Enum.Int[struct {
	SystemClock SyncSource
	SoundClock  SyncSource
}]

var SyncSources = Enum.Values[SyncSource]()

const (
	BPM  = 116
	Bars = 4

	CompensateFrames = 2
	CompensateHz     = 60.0
)

// BPMSync shows the current beat of a song, synced either to the system clock or the sound clock.
type BPMSync struct {
	Panel.Extension[BPMSync] `gd:"BPMSync"`

	Player AudioStreamPlayer.Instance
	Label  Label.Instance

	playing    bool
	syncSource SyncSource

	// Used by system clock.
	timeBegin int
	timeDelay Float.X
}

func (b *BPMSync) Process(_ Float.X) {
	if !b.playing || !b.Player.Playing() {
		return
	}
	var time Float.X
	switch b.syncSource {
	case SyncSources.SystemClock:
		// Obtain from ticks.
		time = Float.X(Time.GetTicksUsec()-b.timeBegin) / 1000000.0
		// Compensate.
		time -= b.timeDelay
	case SyncSources.SoundClock:
		time = b.Player.GetPlaybackPosition() + AudioServer.GetTimeSinceLastMix() - AudioServer.GetOutputLatency() + (1/CompensateHz)*CompensateFrames
	}
	beat := int(time * BPM / 60.0)
	seconds := int(time)
	secondsTotal := int(b.Player.Stream().GetLength())
	b.Label.SetText(fmt.Sprintf("BEAT: %d/%d TIME: %d:%02d / %d:%02d",
		beat%Bars+1, Bars, seconds/60, seconds%60, secondsTotal/60, secondsTotal%60))
}

func (b *BPMSync) OnPlaySystemPressed() {
	b.syncSource = SyncSources.SystemClock
	b.timeBegin = Time.GetTicksUsec()
	b.timeDelay = AudioServer.GetTimeToNextMix() + AudioServer.GetOutputLatency()
	b.playing = true
	b.Player.Play()
}

func (b *BPMSync) OnPlaySoundPressed() {
	b.syncSource = SyncSources.SoundClock
	b.playing = true
	b.Player.Play()
}

func main() {
	classdb.Register[BPMSync]()
	startup.Scene()
}
