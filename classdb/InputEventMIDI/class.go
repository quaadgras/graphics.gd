// Package InputEventMIDI provides methods for working with InputEventMIDI object instances.
package InputEventMIDI

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/InputEvent"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

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
type Instance [1]gdclass.InputEventMIDI

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsInputEventMIDI() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InputEventMIDI

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventMIDI"))
	casted := Instance{*(*gdclass.InputEventMIDI)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Channel() int {
	return int(int(class(self).GetChannel()))
}

func (self Instance) SetChannel(value int) {
	class(self).SetChannel(gd.Int(value))
}

func (self Instance) Message() MIDIMessage {
	return MIDIMessage(class(self).GetMessage())
}

func (self Instance) SetMessage(value MIDIMessage) {
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetChannel() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMessage(message MIDIMessage) {
	var frame = callframe.New()
	callframe.Arg(frame, message)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMessage() MIDIMessage {
	var frame = callframe.New()
	var r_ret = callframe.Ret[MIDIMessage](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPitch(pitch gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pitch)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_pitch, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPitch() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_pitch, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVelocity(velocity gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVelocity() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInstrument(instrument gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, instrument)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_instrument, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInstrument() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_instrument, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressure(pressure gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pressure)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_pressure, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPressure() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_pressure, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetControllerNumber(controller_number gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, controller_number)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_controller_number, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetControllerNumber() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_controller_number, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetControllerValue(controller_value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, controller_value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_set_controller_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetControllerValue() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMIDI.Bind_get_controller_value, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEvent.Advanced(self.AsInputEvent()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEvent.Instance(self.AsInputEvent()), name)
	}
}
func init() {
	gdclass.Register("InputEventMIDI", func(ptr gd.Object) any {
		return [1]gdclass.InputEventMIDI{*(*gdclass.InputEventMIDI)(unsafe.Pointer(&ptr))}
	})
}

type MIDIMessage int

const (
	/*Does not correspond to any MIDI message. This is the default value of [member InputEventMIDI.message].*/
	MidiMessageNone MIDIMessage = 0
	/*MIDI message sent when a note is released.
	  [b]Note:[/b] Not all MIDI devices send this message; some may send [constant MIDI_MESSAGE_NOTE_ON] with [member InputEventMIDI.velocity] set to [code]0[/code].*/
	MidiMessageNoteOff MIDIMessage = 8
	/*MIDI message sent when a note is pressed.*/
	MidiMessageNoteOn MIDIMessage = 9
	/*MIDI message sent to indicate a change in pressure while a note is being pressed down, also called aftertouch.*/
	MidiMessageAftertouch MIDIMessage = 10
	/*MIDI message sent when a controller value changes. In a MIDI device, a controller is any input that doesn't play notes. These may include sliders for volume, balance, and panning, as well as switches and pedals. See the [url=https://en.wikipedia.org/wiki/General_MIDI#Controller_events]General MIDI specification[/url] for a small list.*/
	MidiMessageControlChange MIDIMessage = 11
	/*MIDI message sent when the MIDI device changes its current instrument (also called [i]program[/i] or [i]preset[/i]).*/
	MidiMessageProgramChange MIDIMessage = 12
	/*MIDI message sent to indicate a change in pressure for the whole channel. Some MIDI devices may send this instead of [constant MIDI_MESSAGE_AFTERTOUCH].*/
	MidiMessageChannelPressure MIDIMessage = 13
	/*MIDI message sent when the value of the pitch bender changes, usually a wheel on the MIDI device.*/
	MidiMessagePitchBend MIDIMessage = 14
	/*MIDI system exclusive (SysEx) message. This type of message is not standardized and it's highly dependent on the MIDI device sending it.
	  [b]Note:[/b] Getting this message's data from [InputEventMIDI] is not implemented.*/
	MidiMessageSystemExclusive MIDIMessage = 240
	/*MIDI message sent every quarter frame to keep connected MIDI devices synchronized. Related to [constant MIDI_MESSAGE_TIMING_CLOCK].
	  [b]Note:[/b] Getting this message's data from [InputEventMIDI] is not implemented.*/
	MidiMessageQuarterFrame MIDIMessage = 241
	/*MIDI message sent to jump onto a new position in the current sequence or song.
	  [b]Note:[/b] Getting this message's data from [InputEventMIDI] is not implemented.*/
	MidiMessageSongPositionPointer MIDIMessage = 242
	/*MIDI message sent to select a sequence or song to play.
	  [b]Note:[/b] Getting this message's data from [InputEventMIDI] is not implemented.*/
	MidiMessageSongSelect MIDIMessage = 243
	/*MIDI message sent to request a tuning calibration. Used on analog synthesizers. Most modern MIDI devices do not need this message.*/
	MidiMessageTuneRequest MIDIMessage = 246
	/*MIDI message sent 24 times after [constant MIDI_MESSAGE_QUARTER_FRAME], to keep connected MIDI devices synchronized.*/
	MidiMessageTimingClock MIDIMessage = 248
	/*MIDI message sent to start the current sequence or song from the beginning.*/
	MidiMessageStart MIDIMessage = 250
	/*MIDI message sent to resume from the point the current sequence or song was paused.*/
	MidiMessageContinue MIDIMessage = 251
	/*MIDI message sent to pause the current sequence or song.*/
	MidiMessageStop MIDIMessage = 252
	/*MIDI message sent repeatedly while the MIDI device is idle, to tell the receiver that the connection is alive. Most MIDI devices do not send this message.*/
	MidiMessageActiveSensing MIDIMessage = 254
	/*MIDI message sent to reset a MIDI device to its default state, as if it was just turned on. It should not be sent when the MIDI device is being turned on.*/
	MidiMessageSystemReset MIDIMessage = 255
)
