package InputEventMIDI

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEvent"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
InputEventMIDI stores information about messages from [url=https://en.wikipedia.org/wiki/MIDI]MIDI[/url] (Musical Instrument Digital Interface) devices. These may include musical keyboards, synthesizers, and drum machines.
MIDI messages can be received over a 5-pin MIDI connector or over USB. If your device supports both be sure to check the settings in the device to see which output it is using.
By default, Godot does not detect MIDI devices. You need to call [method OS.open_midi_inputs], first. You can check which devices are detected with [method OS.get_connected_midi_inputs], and close the connection with [method OS.close_midi_inputs].
[codeblocks]
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
[/codeblocks]
[b]Note:[/b] Godot does not support MIDI output, so there is no way to emit MIDI messages from Godot. Only MIDI input is supported.
*/
type Instance [1]classdb.InputEventMIDI

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.InputEventMIDI

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventMIDI"))
	return Instance{classdb.InputEventMIDI(object)}
}

func (self Instance) Channel() int {
	return int(int(class(self).GetChannel()))
}

func (self Instance) SetChannel(value int) {
	class(self).SetChannel(gd.Int(value))
}

func (self Instance) Message() gdconst.MIDIMessage {
	return gdconst.MIDIMessage(class(self).GetMessage())
}

func (self Instance) SetMessage(value gdconst.MIDIMessage) {
	class(self).SetMessage(value)
}

func (self Instance) Pitch() int {
	return int(int(class(self).GetPitch()))
}

func (self Instance) SetPitch(value int) {
	class(self).SetPitch(gd.Int(value))
}

func (self Instance) Velocity() int {
	return int(int(class(self).GetVelocity()))
}

func (self Instance) SetVelocity(value int) {
	class(self).SetVelocity(gd.Int(value))
}

func (self Instance) Instrument() int {
	return int(int(class(self).GetInstrument()))
}

func (self Instance) SetInstrument(value int) {
	class(self).SetInstrument(gd.Int(value))
}

func (self Instance) Pressure() int {
	return int(int(class(self).GetPressure()))
}

func (self Instance) SetPressure(value int) {
	class(self).SetPressure(gd.Int(value))
}

func (self Instance) ControllerNumber() int {
	return int(int(class(self).GetControllerNumber()))
}

func (self Instance) SetControllerNumber(value int) {
	class(self).SetControllerNumber(gd.Int(value))
}

func (self Instance) ControllerValue() int {
	return int(int(class(self).GetControllerValue()))
}

func (self Instance) SetControllerValue(value int) {
	class(self).SetControllerValue(gd.Int(value))
}

//go:nosplit
func (self class) SetChannel(channel gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetChannel() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMessage(message gdconst.MIDIMessage) {
	var frame = callframe.New()
	callframe.Arg(frame, message)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMessage() gdconst.MIDIMessage {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdconst.MIDIMessage](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPitch(pitch gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pitch)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_pitch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPitch() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_pitch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVelocity(velocity gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVelocity() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInstrument(instrument gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, instrument)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_instrument, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInstrument() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_instrument, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressure(pressure gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pressure)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPressure() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetControllerNumber(controller_number gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, controller_number)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_controller_number, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetControllerNumber() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_controller_number, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetControllerValue(controller_value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, controller_value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_controller_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetControllerValue() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_controller_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventMIDI() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventMIDI() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsInputEvent() InputEvent.Advanced {
	return *((*InputEvent.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEvent() InputEvent.Instance {
	return *((*InputEvent.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEvent(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEvent(), name)
	}
}
func init() {
	classdb.Register("InputEventMIDI", func(ptr gd.Object) any { return classdb.InputEventMIDI(ptr) })
}
