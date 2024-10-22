package InputEventMIDI

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEvent"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.InputEventMIDI
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InputEventMIDI
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("InputEventMIDI"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Channel() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetChannel()))
}

func (self Go) SetChannel(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetChannel(gd.Int(value))
}

func (self Go) Message() gd.MIDIMessage {
	gc := gd.GarbageCollector(); _ = gc
		return gd.MIDIMessage(class(self).GetMessage())
}

func (self Go) SetMessage(value gd.MIDIMessage) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMessage(value)
}

func (self Go) Pitch() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPitch()))
}

func (self Go) SetPitch(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPitch(gd.Int(value))
}

func (self Go) Velocity() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetVelocity()))
}

func (self Go) SetVelocity(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVelocity(gd.Int(value))
}

func (self Go) Instrument() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetInstrument()))
}

func (self Go) SetInstrument(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInstrument(gd.Int(value))
}

func (self Go) Pressure() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPressure()))
}

func (self Go) SetPressure(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPressure(gd.Int(value))
}

func (self Go) ControllerNumber() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetControllerNumber()))
}

func (self Go) SetControllerNumber(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetControllerNumber(gd.Int(value))
}

func (self Go) ControllerValue() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetControllerValue()))
}

func (self Go) SetControllerValue(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetControllerValue(gd.Int(value))
}

//go:nosplit
func (self class) SetChannel(channel gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_set_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetChannel() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_get_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMessage(message gd.MIDIMessage)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, message)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_set_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMessage() gd.MIDIMessage {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.MIDIMessage](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_get_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPitch(pitch gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pitch)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_set_pitch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPitch() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_get_pitch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVelocity(velocity gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVelocity() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInstrument(instrument gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, instrument)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_set_instrument, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInstrument() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_get_instrument, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPressure(pressure gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressure)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_set_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPressure() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_get_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetControllerNumber(controller_number gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, controller_number)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_set_controller_number, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetControllerNumber() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_get_controller_number, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetControllerValue(controller_value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, controller_value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_set_controller_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetControllerValue() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventMIDI.Bind_get_controller_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventMIDI() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventMIDI() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEvent() InputEvent.GD { return *((*InputEvent.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEvent() InputEvent.Go { return *((*InputEvent.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEvent(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEvent(), name)
	}
}
func init() {classdb.Register("InputEventMIDI", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
