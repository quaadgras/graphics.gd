/*
[gdscript]
func _ready():
    OS.open_midi_inputs()
    print(OS.get_connected_midi_inputs())

func _input(input_event):
    if input_event is InputEventMIDI:
        _print_midi_info(input_event)

func _print_midi_info(midi_event):
    print(midi_event)
    print("Channel ", midi_event.channel)
    print("Message ", midi_event.message)
    print("Pitch ", midi_event.pitch)
    print("Velocity ", midi_event.velocity)
    print("Instrument ", midi_event.instrument)
    print("Pressure ", midi_event.pressure)
    print("Controller number: ", midi_event.controller_number)
    print("Controller value: ", midi_event.controller_value)
[/gdscript]
[csharp]
public override void _Ready()
{
    OS.OpenMidiInputs();
    GD.Print(OS.GetConnectedMidiInputs());
}

public override void _Input(InputEvent inputEvent)
{
    if (inputEvent is InputEventMidi midiEvent)
    {
        PrintMIDIInfo(midiEvent);
    }
}

private void PrintMIDIInfo(InputEventMidi midiEvent)
{
    GD.Print(midiEvent);
    GD.Print($"Channel {midiEvent.Channel}");
    GD.Print($"Message {midiEvent.Message}");
    GD.Print($"Pitch {midiEvent.Pitch}");
    GD.Print($"Velocity {midiEvent.Velocity}");
    GD.Print($"Instrument {midiEvent.Instrument}");
    GD.Print($"Pressure {midiEvent.Pressure}");
    GD.Print($"Controller number: {midiEvent.ControllerNumber}");
    GD.Print($"Controller value: {midiEvent.ControllerValue}");
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMIDI"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/OS"
	"graphics.gd/variant/Object"
)

type MyMIDI struct {
	Node.Extension[MyMIDI]
}

func (n MyMIDI) Ready() {
	OS.OpenMidiInputs()
	fmt.Println(OS.GetConnectedMidiInputs())
}

func (n MyMIDI) Input(event InputEvent.Instance) {
	if event, ok := Object.As[InputEventMIDI.Instance](event); ok {
		fmt.Println(event)
		fmt.Println("Channel ", event.Channel())
		fmt.Println("Message ", event.Message())
		fmt.Println("Pitch ", event.Pitch())
		fmt.Println("Velocity ", event.Velocity())
		fmt.Println("Instrument ", event.Instrument())
		fmt.Println("Pressure ", event.Pressure())
		fmt.Println("Controller number: ", event.ControllerNumber())
		fmt.Println("Controller value: ", event.ControllerValue())
	}
}
