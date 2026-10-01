package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioEffectRecord"
	"graphics.gd/classdb/AudioServer"
	"graphics.gd/classdb/AudioStreamPlayer"
	"graphics.gd/classdb/AudioStreamWAV"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/startup"
	"graphics.gd/variant/Object"
)

// MicRecord records audio from the microphone through an AudioEffectRecord
// on the "Record" bus and can play it back or save it to a WAV file.
type MicRecord struct {
	Control.Extension[MicRecord] `gd:"MicRecord"`

	Status            Label.Instance
	RecordButton      Button.Instance
	PlayButton        Button.Instance
	PlayMusic         Button.Instance
	SaveButton        Button.Instance
	Filename          LineEdit.Instance `gd:"SaveButton/Filename"`
	AudioStreamPlayer AudioStreamPlayer.Instance
	Music             AudioStreamPlayer.Instance `gd:"AudioStreamPlayer2"`

	effect    AudioEffectRecord.Instance
	recording AudioStreamWAV.Instance

	stereo  bool
	mixRate int
	format  AudioStreamWAV.Format
}

func NewMicRecord() *MicRecord {
	return &MicRecord{
		stereo:  true,
		mixRate: 44100,                       // This is the default mix rate on recordings.
		format:  AudioStreamWAV.Format16Bits, // This is the default format on recordings.
	}
}

func (m *MicRecord) Ready() {
	idx := AudioServer.GetBusIndex("Record")
	m.effect, _ = Object.As[AudioEffectRecord.Instance](AudioServer.GetBusEffect(AudioServer.Bus(idx), 0))
}

func (m *MicRecord) OnRecordButtonPressed() {
	if m.effect.IsRecordingActive() {
		m.recording = m.effect.GetRecording()
		m.PlayButton.AsBaseButton().SetDisabled(false)
		m.SaveButton.AsBaseButton().SetDisabled(false)
		m.effect.SetRecordingActive(false)
		m.recording.SetMixRate(m.mixRate)
		m.recording.SetFormat(m.format)
		m.recording.SetStereo(m.stereo)
		m.RecordButton.SetText("Record")
		m.Status.SetText("")
	} else {
		m.PlayButton.AsBaseButton().SetDisabled(true)
		m.SaveButton.AsBaseButton().SetDisabled(true)
		m.effect.SetRecordingActive(true)
		m.RecordButton.SetText("Stop")
		m.Status.SetText("Status: Recording...")
	}
}

func (m *MicRecord) OnPlayButtonPressed() {
	formatName := "IMA ADPCM compressed"
	switch m.recording.Format() {
	case AudioStreamWAV.Format8Bits:
		formatName = "8-bit uncompressed"
	case AudioStreamWAV.Format16Bits:
		formatName = "16-bit uncompressed"
	}
	stereoText := "No"
	if m.recording.Stereo() {
		stereoText = "Yes"
	}
	fmt.Println("\nPlaying recording:", m.recording)
	fmt.Println("Format:", formatName)
	fmt.Println("Mix rate:", m.recording.MixRate(), "Hz")
	fmt.Println("Stereo:", stereoText)
	fmt.Println("Size:", len(m.recording.Data()), "bytes")
	m.AudioStreamPlayer.SetStream(m.recording.AsAudioStream())
	m.AudioStreamPlayer.Play()
}

func (m *MicRecord) OnPlayMusicPressed() {
	if m.Music.Playing() {
		m.Music.Stop()
		m.PlayMusic.SetText("Play Music")
	} else {
		m.Music.Play()
		m.PlayMusic.SetText("Stop Music")
	}
}

func (m *MicRecord) OnSaveButtonPressed() {
	savePath := m.Filename.Text()
	if err := m.recording.SaveToWav(savePath); err != nil {
		m.Status.SetText("Status: Failed to save WAV file: " + err.Error())
		return
	}
	m.Status.SetText(fmt.Sprintf("Status: Saved WAV file to: %s\n(%s)",
		savePath, ProjectSettings.GlobalizePath(savePath)))
}

func (m *MicRecord) OnMixRateOptionButtonItemSelected(index int) {
	switch index {
	case 0:
		m.mixRate = 11025
	case 1:
		m.mixRate = 16000
	case 2:
		m.mixRate = 22050
	case 3:
		m.mixRate = 32000
	case 4:
		m.mixRate = 44100
	case 5:
		m.mixRate = 48000
	}
	if m.recording != AudioStreamWAV.Nil {
		m.recording.SetMixRate(m.mixRate)
	}
}

func (m *MicRecord) OnFormatOptionButtonItemSelected(index int) {
	switch index {
	case 0:
		m.format = AudioStreamWAV.Format8Bits
	case 1:
		m.format = AudioStreamWAV.Format16Bits
	case 2:
		m.format = AudioStreamWAV.FormatImaAdpcm
	}
	if m.recording != AudioStreamWAV.Nil {
		m.recording.SetFormat(m.format)
	}
}

func (m *MicRecord) OnStereoCheckButtonToggled(buttonPressed bool) {
	m.stereo = buttonPressed
	if m.recording != AudioStreamWAV.Nil {
		m.recording.SetStereo(m.stereo)
	}
}

func (m *MicRecord) OnOpenUserFolderButtonPressed() {
	OS.ShellOpen(ProjectSettings.GlobalizePath("user://"))
}

func main() {
	classdb.Register[MicRecord](NewMicRecord)
	startup.Scene()
}
