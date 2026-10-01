package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioServer"
	"graphics.gd/classdb/AudioStreamPlayer"
	"graphics.gd/classdb/CenterContainer"
	"graphics.gd/classdb/Node"
	"graphics.gd/startup"
)

// AudioEffects toggles the various audio effects on the Master bus while
// playing sound effects and background music.
type AudioEffects struct {
	CenterContainer.Extension[AudioEffects] `gd:"AudioEffects"`

	SoundEffects struct {
		Node.Instance

		Music    AudioStreamPlayer.Instance
		Ding     AudioStreamPlayer.Instance
		Glass    AudioStreamPlayer.Instance
		Meow     AudioStreamPlayer.Instance
		Beeps    AudioStreamPlayer.Instance
		Trombone AudioStreamPlayer.Instance
		Static   AudioStreamPlayer.Instance
		Whistle  AudioStreamPlayer.Instance
	}
}

func (a *AudioEffects) OnToggleMusicToggled(buttonPressed bool) {
	if buttonPressed {
		a.SoundEffects.Music.Play()
	} else {
		a.SoundEffects.Music.Stop()
	}
}

func (a *AudioEffects) OnDingButtonPressed()     { a.SoundEffects.Ding.Play() }
func (a *AudioEffects) OnGlassButtonPressed()    { a.SoundEffects.Glass.Play() }
func (a *AudioEffects) OnMeowButtonPressed()     { a.SoundEffects.Meow.Play() }
func (a *AudioEffects) OnBeepsButtonPressed()    { a.SoundEffects.Beeps.Play() }
func (a *AudioEffects) OnTromboneButtonPressed() { a.SoundEffects.Trombone.Play() }
func (a *AudioEffects) OnStaticButtonPressed()   { a.SoundEffects.Static.Play() }
func (a *AudioEffects) OnWhistleButtonPressed()  { a.SoundEffects.Whistle.Play() }

// setEffect toggles the effect at the given index on the Master bus.
func setEffect(idx int, enabled bool) {
	AudioServer.SetBusEffectEnabled(0, AudioServer.Effect(idx), enabled)
}

func (a *AudioEffects) OnToggleAmplifyToggled(on bool)        { setEffect(0, on) }
func (a *AudioEffects) OnToggleBandLimiterToggled(on bool)    { setEffect(1, on) }
func (a *AudioEffects) OnToggleBandPassFilterToggled(on bool) { setEffect(2, on) }
func (a *AudioEffects) OnToggleChorusToggled(on bool)         { setEffect(3, on) }
func (a *AudioEffects) OnToggleCompressorToggled(on bool)     { setEffect(4, on) }
func (a *AudioEffects) OnToggleDelayToggled(on bool)          { setEffect(5, on) }
func (a *AudioEffects) OnToggleDistortionToggled(on bool)     { setEffect(6, on) }
func (a *AudioEffects) OnToggleEq6Toggled(on bool)            { setEffect(7, on) }
func (a *AudioEffects) OnToggleEq10Toggled(on bool)           { setEffect(8, on) }
func (a *AudioEffects) OnToggleEq21Toggled(on bool)           { setEffect(9, on) }
func (a *AudioEffects) OnToggleHighPassFilterToggled(on bool) { setEffect(10, on) }
func (a *AudioEffects) OnToggleLowShelfFilterToggled(on bool) { setEffect(11, on) }
func (a *AudioEffects) OnToggleNotchFilterToggled(on bool)    { setEffect(12, on) }
func (a *AudioEffects) OnTogglePannerToggled(on bool)         { setEffect(13, on) }
func (a *AudioEffects) OnTogglePhaserToggled(on bool)         { setEffect(14, on) }
func (a *AudioEffects) OnTogglePitchShiftToggled(on bool)     { setEffect(15, on) }
func (a *AudioEffects) OnToggleReverbToggled(on bool)         { setEffect(16, on) }
func (a *AudioEffects) OnToggleStereoEnhanceToggled(on bool)  { setEffect(17, on) }

func main() {
	classdb.Register[AudioEffects]()
	startup.Scene()
}
