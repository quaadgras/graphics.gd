package main

import (
	"fmt"
	"math"
	"strconv"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioStream"
	"graphics.gd/classdb/AudioStreamPlayer"
	"graphics.gd/classdb/ColorRect"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/HBoxContainer"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMIDI"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Timer"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

// A standard piano with 88 keys has keys from 21 to 108.
// To get a different set of keys, modify these numbers.
// A maximally extended 108-key piano goes from 12 to 119.
// A 76-key piano goes from 23 to 98, 61-key from 36 to 96,
// 49-key from 36 to 84, 37-key from 41 to 77, and 25-key
// from 48 to 72. Middle C is pitch number 60, A440 is 69.
const (
	StartKey = 21
	EndKey   = 108
)

var (
	WhiteKeyScene = Resource.Load[PackedScene.Is[*PianoKey]]("res://piano_keys/white_piano_key.tscn")
	BlackKeyScene = Resource.Load[PackedScene.Is[*PianoKey]]("res://piano_keys/black_piano_key.tscn")
)

type Piano struct {
	ColorRect.Extension[Piano] `gd:"Piano"`

	WhiteKeys HBoxContainer.Instance
	BlackKeys HBoxContainer.Instance

	pianoKeys map[int]*PianoKey
}

func (p *Piano) Ready() {
	if isNoteIndexSharp(pitchIndexToNoteIndex(StartKey)) {
		panic("The start key can't be a sharp note (limitation of this piano-generating algorithm). Try 21.")
	}
	p.pianoKeys = map[int]*PianoKey{}
	for i := StartKey; i <= EndKey; i++ {
		p.pianoKeys[i] = p.createPianoKey(i)
	}
	if p.WhiteKeys.AsNode().GetChildCount() != p.BlackKeys.AsNode().GetChildCount() {
		p.addPlaceholderKey(p.BlackKeys)
	}

	OS.OpenMidiInputs()

	if inputs := OS.GetConnectedMidiInputs(); len(inputs) > 0 {
		fmt.Println(inputs)
	}
}

func (p *Piano) Input(event InputEvent.Instance) {
	midi, ok := Object.As[InputEventMIDI.Instance](event)
	if !ok {
		return
	}
	if midi.Pitch() < StartKey || midi.Pitch() > EndKey {
		// The given pitch isn't on the on-screen keyboard, so return.
		return
	}
	printMidiInfo(midi)
	key := p.pianoKeys[midi.Pitch()]
	if midi.Message() == InputEventMIDI.MidiMessageNoteOn {
		key.Activate()
	} else {
		key.Deactivate()
	}
}

func (p *Piano) addPlaceholderKey(container HBoxContainer.Instance) {
	placeholder := Control.New()
	placeholder.SetSizeFlagsHorizontal(Control.SizeExpandFill)
	placeholder.SetMouseFilter(Control.MouseFilterIgnore)
	placeholder.AsNode().SetName("Placeholder")
	container.AsNode().AddChild(placeholder.AsNode())
}

func (p *Piano) createPianoKey(pitchIndex int) *PianoKey {
	noteIndex := pitchIndexToNoteIndex(pitchIndex)
	var pianoKey *PianoKey
	if isNoteIndexSharp(noteIndex) {
		pianoKey = BlackKeyScene.Instantiate()
		p.BlackKeys.AsNode().AddChild(pianoKey.AsNode())
	} else {
		pianoKey = WhiteKeyScene.Instantiate()
		p.WhiteKeys.AsNode().AddChild(pianoKey.AsNode())
		if isNoteIndexLackingSharp(noteIndex) {
			p.addPlaceholderKey(p.BlackKeys)
		}
	}
	pianoKey.Setup(pitchIndex)
	return pianoKey
}

// B and E, because no B# or E#.
func isNoteIndexLackingSharp(noteIndex int) bool {
	return noteIndex == 2 || noteIndex == 7
}

// A#, C#, D#, F#, and G#.
func isNoteIndexSharp(noteIndex int) bool {
	switch noteIndex {
	case 1, 4, 6, 9, 11:
		return true
	}
	return false
}

func pitchIndexToNoteIndex(pitch int) int {
	return (pitch + 3) % 12
}

func printMidiInfo(midi InputEventMIDI.Instance) {
	fmt.Println("Channel:", midi.Channel())
	fmt.Println("Message:", midi.Message())
	fmt.Println("Pitch:", midi.Pitch())
	fmt.Println("Velocity:", midi.Velocity())
	fmt.Println("Instrument:", midi.Instrument())
	fmt.Println("Pressure:", midi.Pressure())
	fmt.Println("Controller number:", midi.ControllerNumber())
	fmt.Println("Controller value:", midi.ControllerValue())
}

var KeySound = Resource.Load[AudioStream.Instance]("res://piano_keys/A440.wav")

type PianoKey struct {
	Control.Extension[PianoKey] `gd:"PianoKey"`

	Key        *PianoKeyColor
	ColorTimer Timer.Instance

	pitchScale Float.X
	startColor Color.RGBA
}

func (k *PianoKey) Ready() {
	k.startColor = k.Key.AsColorRect().Color()
}

func (k *PianoKey) Setup(pitchIndex int) {
	k.AsNode().SetName("PianoKey" + strconv.Itoa(pitchIndex))
	exponent := (float64(pitchIndex) - 69.0) / 12.0
	k.pitchScale = Float.X(math.Pow(2, exponent))
}

func (k *PianoKey) Activate() {
	k.Key.AsColorRect().SetColor(Color.MulX(Color.Add(Color.RGBA{1, 1, 0, 1}, k.startColor), 0.5))
	audio := AudioStreamPlayer.New()
	k.AsNode().AddChild(audio.AsNode())
	audio.SetStream(KeySound)
	audio.SetPitchScale(k.pitchScale)
	audio.Play()
	k.ColorTimer.Start()
	SceneTree.Get(k.AsNode()).CreateTimer(8).OnTimeout(func() {
		audio.AsNode().QueueFree()
	})
}

func (k *PianoKey) Deactivate() {
	k.Key.AsColorRect().SetColor(k.startColor)
}

// PianoKeyColor exists just to forward clicks on the key to the parent PianoKey.
type PianoKeyColor struct {
	ColorRect.Extension[PianoKeyColor] `gd:"PianoKeyColor"`
}

func (c *PianoKeyColor) GuiInput(event InputEvent.Instance) {
	if button, ok := Object.As[InputEventMouseButton.Instance](event); ok && button.AsInputEvent().IsPressed() {
		if parent, ok := Object.As[*PianoKey](c.AsNode().GetParent()); ok {
			parent.Activate()
		}
	}
}

func main() {
	classdb.Register[Piano]()
	classdb.Register[PianoKey]()
	classdb.Register[PianoKeyColor]()
	startup.Scene()
}
