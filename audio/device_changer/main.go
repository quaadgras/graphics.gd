package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioServer"
	"graphics.gd/classdb/AudioStreamPlayer"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/ItemList"
	"graphics.gd/classdb/Label"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
)

// DeviceChanger lists the available audio output devices and lets the
// user switch between them while a sound is playing.
type DeviceChanger struct {
	Control.Extension[DeviceChanger] `gd:"DeviceChanger"`

	ItemList          ItemList.Instance
	DeviceInfo        Label.Instance
	PlayAudio         Button.Instance
	AudioStreamPlayer AudioStreamPlayer.Instance
}

func (d *DeviceChanger) Ready() {
	for _, item := range AudioServer.GetOutputDeviceList() {
		d.ItemList.AddItem(item)
	}
	device := AudioServer.OutputDevice()
	for i := range d.ItemList.ItemCount() {
		if device == d.ItemList.GetItemText(i) {
			d.ItemList.Select(i)
			break
		}
	}
}

func (d *DeviceChanger) Process(_ Float.X) {
	speakerModeText := "Stereo"
	switch AudioServer.GetSpeakerMode() {
	case AudioServer.SpeakerSurround31:
		speakerModeText = "Surround 3.1"
	case AudioServer.SpeakerSurround51:
		speakerModeText = "Surround 5.1"
	case AudioServer.SpeakerSurround71:
		speakerModeText = "Surround 7.1"
	}
	d.DeviceInfo.SetText("Current Device: " + AudioServer.OutputDevice() + "\n" +
		"Speaker Mode: " + speakerModeText)
}

func (d *DeviceChanger) OnButtonButtonDown() {
	for _, item := range d.ItemList.GetSelectedItems() {
		device := d.ItemList.GetItemText(int(item))
		AudioServer.SetOutputDevice(device)
	}
}

func (d *DeviceChanger) OnPlayAudioButtonDown() {
	if d.AudioStreamPlayer.Playing() {
		d.AudioStreamPlayer.Stop()
		d.PlayAudio.SetText("Play Audio")
	} else {
		d.AudioStreamPlayer.Play()
		d.PlayAudio.SetText("Stop Audio")
	}
}

func main() {
	classdb.Register[DeviceChanger]()
	startup.Scene()
}
