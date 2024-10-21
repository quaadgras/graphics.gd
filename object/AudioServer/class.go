package AudioServer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[AudioServer] is a low-level server interface for audio access. It is in charge of creating sample data (playable audio) as well as its playback via a voice interface.

*/
type Simple [1]classdb.AudioServer
func (self Simple) SetBusCount(amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBusCount(gd.Int(amount))
}
func (self Simple) GetBusCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBusCount()))
}
func (self Simple) RemoveBus(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveBus(gd.Int(index))
}
func (self Simple) AddBus(at_position int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddBus(gd.Int(at_position))
}
func (self Simple) MoveBus(index int, to_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveBus(gd.Int(index), gd.Int(to_index))
}
func (self Simple) SetBusName(bus_idx int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBusName(gd.Int(bus_idx), gc.String(name))
}
func (self Simple) GetBusName(bus_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBusName(gc, gd.Int(bus_idx)).String())
}
func (self Simple) GetBusIndex(bus_name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBusIndex(gc.StringName(bus_name))))
}
func (self Simple) GetBusChannels(bus_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBusChannels(gd.Int(bus_idx))))
}
func (self Simple) SetBusVolumeDb(bus_idx int, volume_db float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBusVolumeDb(gd.Int(bus_idx), gd.Float(volume_db))
}
func (self Simple) GetBusVolumeDb(bus_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBusVolumeDb(gd.Int(bus_idx))))
}
func (self Simple) SetBusSend(bus_idx int, send string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBusSend(gd.Int(bus_idx), gc.StringName(send))
}
func (self Simple) GetBusSend(bus_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBusSend(gc, gd.Int(bus_idx)).String())
}
func (self Simple) SetBusSolo(bus_idx int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBusSolo(gd.Int(bus_idx), enable)
}
func (self Simple) IsBusSolo(bus_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBusSolo(gd.Int(bus_idx)))
}
func (self Simple) SetBusMute(bus_idx int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBusMute(gd.Int(bus_idx), enable)
}
func (self Simple) IsBusMute(bus_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBusMute(gd.Int(bus_idx)))
}
func (self Simple) SetBusBypassEffects(bus_idx int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBusBypassEffects(gd.Int(bus_idx), enable)
}
func (self Simple) IsBusBypassingEffects(bus_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBusBypassingEffects(gd.Int(bus_idx)))
}
func (self Simple) AddBusEffect(bus_idx int, effect [1]classdb.AudioEffect, at_position int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddBusEffect(gd.Int(bus_idx), effect, gd.Int(at_position))
}
func (self Simple) RemoveBusEffect(bus_idx int, effect_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveBusEffect(gd.Int(bus_idx), gd.Int(effect_idx))
}
func (self Simple) GetBusEffectCount(bus_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBusEffectCount(gd.Int(bus_idx))))
}
func (self Simple) GetBusEffect(bus_idx int, effect_idx int) [1]classdb.AudioEffect {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioEffect(Expert(self).GetBusEffect(gc, gd.Int(bus_idx), gd.Int(effect_idx)))
}
func (self Simple) GetBusEffectInstance(bus_idx int, effect_idx int, channel int) [1]classdb.AudioEffectInstance {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioEffectInstance(Expert(self).GetBusEffectInstance(gc, gd.Int(bus_idx), gd.Int(effect_idx), gd.Int(channel)))
}
func (self Simple) SwapBusEffects(bus_idx int, effect_idx int, by_effect_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SwapBusEffects(gd.Int(bus_idx), gd.Int(effect_idx), gd.Int(by_effect_idx))
}
func (self Simple) SetBusEffectEnabled(bus_idx int, effect_idx int, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBusEffectEnabled(gd.Int(bus_idx), gd.Int(effect_idx), enabled)
}
func (self Simple) IsBusEffectEnabled(bus_idx int, effect_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBusEffectEnabled(gd.Int(bus_idx), gd.Int(effect_idx)))
}
func (self Simple) GetBusPeakVolumeLeftDb(bus_idx int, channel int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBusPeakVolumeLeftDb(gd.Int(bus_idx), gd.Int(channel))))
}
func (self Simple) GetBusPeakVolumeRightDb(bus_idx int, channel int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBusPeakVolumeRightDb(gd.Int(bus_idx), gd.Int(channel))))
}
func (self Simple) SetPlaybackSpeedScale(scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPlaybackSpeedScale(gd.Float(scale))
}
func (self Simple) GetPlaybackSpeedScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPlaybackSpeedScale()))
}
func (self Simple) Lock() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Lock()
}
func (self Simple) Unlock() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Unlock()
}
func (self Simple) GetSpeakerMode() classdb.AudioServerSpeakerMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AudioServerSpeakerMode(Expert(self).GetSpeakerMode())
}
func (self Simple) GetMixRate() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMixRate()))
}
func (self Simple) GetOutputDeviceList() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetOutputDeviceList(gc))
}
func (self Simple) GetOutputDevice() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetOutputDevice(gc).String())
}
func (self Simple) SetOutputDevice(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOutputDevice(gc.String(name))
}
func (self Simple) GetTimeToNextMix() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTimeToNextMix()))
}
func (self Simple) GetTimeSinceLastMix() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTimeSinceLastMix()))
}
func (self Simple) GetOutputLatency() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetOutputLatency()))
}
func (self Simple) GetInputDeviceList() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetInputDeviceList(gc))
}
func (self Simple) GetInputDevice() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetInputDevice(gc).String())
}
func (self Simple) SetInputDevice(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputDevice(gc.String(name))
}
func (self Simple) SetBusLayout(bus_layout [1]classdb.AudioBusLayout) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBusLayout(bus_layout)
}
func (self Simple) GenerateBusLayout() [1]classdb.AudioBusLayout {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioBusLayout(Expert(self).GenerateBusLayout(gc))
}
func (self Simple) SetEnableTaggingUsedAudioStreams(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableTaggingUsedAudioStreams(enable)
}
func (self Simple) IsStreamRegisteredAsSample(stream [1]classdb.AudioStream) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsStreamRegisteredAsSample(stream))
}
func (self Simple) RegisterStreamAsSample(stream [1]classdb.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegisterStreamAsSample(stream)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetBusCount(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_bus_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBusCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the bus at index [param index].
*/
//go:nosplit
func (self class) RemoveBus(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_remove_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a bus at [param at_position].
*/
//go:nosplit
func (self class) AddBus(at_position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, at_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_add_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the bus from index [param index] to index [param to_index].
*/
//go:nosplit
func (self class) MoveBus(index gd.Int, to_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, to_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_move_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the name of the bus at index [param bus_idx] to [param name].
*/
//go:nosplit
func (self class) SetBusName(bus_idx gd.Int, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the bus with the index [param bus_idx].
*/
//go:nosplit
func (self class) GetBusName(ctx gd.Lifetime, bus_idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the index of the bus with the name [param bus_name]. Returns [code]-1[/code] if no bus with the specified name exist.
*/
//go:nosplit
func (self class) GetBusIndex(bus_name gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bus_name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of channels of the bus at index [param bus_idx].
*/
//go:nosplit
func (self class) GetBusChannels(bus_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_channels, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the volume of the bus at index [param bus_idx] to [param volume_db].
*/
//go:nosplit
func (self class) SetBusVolumeDb(bus_idx gd.Int, volume_db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, volume_db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_bus_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the volume of the bus at index [param bus_idx] in dB.
*/
//go:nosplit
func (self class) GetBusVolumeDb(bus_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Connects the output of the bus at [param bus_idx] to the bus named [param send].
*/
//go:nosplit
func (self class) SetBusSend(bus_idx gd.Int, send gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, mmm.Get(send))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_bus_send, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the bus that the bus at index [param bus_idx] sends to.
*/
//go:nosplit
func (self class) GetBusSend(ctx gd.Lifetime, bus_idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_send, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
If [code]true[/code], the bus at index [param bus_idx] is in solo mode.
*/
//go:nosplit
func (self class) SetBusSolo(bus_idx gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_bus_solo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the bus at index [param bus_idx] is in solo mode.
*/
//go:nosplit
func (self class) IsBusSolo(bus_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_is_bus_solo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the bus at index [param bus_idx] is muted.
*/
//go:nosplit
func (self class) SetBusMute(bus_idx gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_bus_mute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the bus at index [param bus_idx] is muted.
*/
//go:nosplit
func (self class) IsBusMute(bus_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_is_bus_mute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the bus at index [param bus_idx] is bypassing effects.
*/
//go:nosplit
func (self class) SetBusBypassEffects(bus_idx gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_bus_bypass_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the bus at index [param bus_idx] is bypassing effects.
*/
//go:nosplit
func (self class) IsBusBypassingEffects(bus_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_is_bus_bypassing_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an [AudioEffect] effect to the bus [param bus_idx] at [param at_position].
*/
//go:nosplit
func (self class) AddBusEffect(bus_idx gd.Int, effect object.AudioEffect, at_position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, mmm.Get(effect[0].AsPointer())[0])
	callframe.Arg(frame, at_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_add_bus_effect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the effect at index [param effect_idx] from the bus at index [param bus_idx].
*/
//go:nosplit
func (self class) RemoveBusEffect(bus_idx gd.Int, effect_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, effect_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_remove_bus_effect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of effects on the bus at [param bus_idx].
*/
//go:nosplit
func (self class) GetBusEffectCount(bus_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_effect_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [AudioEffect] at position [param effect_idx] in bus [param bus_idx].
*/
//go:nosplit
func (self class) GetBusEffect(ctx gd.Lifetime, bus_idx gd.Int, effect_idx gd.Int) object.AudioEffect {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, effect_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_effect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioEffect
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [AudioEffectInstance] assigned to the given bus and effect indices (and optionally channel).
*/
//go:nosplit
func (self class) GetBusEffectInstance(ctx gd.Lifetime, bus_idx gd.Int, effect_idx gd.Int, channel gd.Int) object.AudioEffectInstance {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, effect_idx)
	callframe.Arg(frame, channel)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_effect_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioEffectInstance
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Swaps the position of two effects in bus [param bus_idx].
*/
//go:nosplit
func (self class) SwapBusEffects(bus_idx gd.Int, effect_idx gd.Int, by_effect_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, effect_idx)
	callframe.Arg(frame, by_effect_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_swap_bus_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the effect at index [param effect_idx] on the bus at index [param bus_idx] is enabled.
*/
//go:nosplit
func (self class) SetBusEffectEnabled(bus_idx gd.Int, effect_idx gd.Int, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, effect_idx)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_bus_effect_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the effect at index [param effect_idx] on the bus at index [param bus_idx] is enabled.
*/
//go:nosplit
func (self class) IsBusEffectEnabled(bus_idx gd.Int, effect_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, effect_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_is_bus_effect_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the peak volume of the left speaker at bus index [param bus_idx] and channel index [param channel].
*/
//go:nosplit
func (self class) GetBusPeakVolumeLeftDb(bus_idx gd.Int, channel gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, channel)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_peak_volume_left_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the peak volume of the right speaker at bus index [param bus_idx] and channel index [param channel].
*/
//go:nosplit
func (self class) GetBusPeakVolumeRightDb(bus_idx gd.Int, channel gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bus_idx)
	callframe.Arg(frame, channel)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_bus_peak_volume_right_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlaybackSpeedScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_playback_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlaybackSpeedScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_playback_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Locks the audio driver's main loop.
[b]Note:[/b] Remember to unlock it afterwards.
*/
//go:nosplit
func (self class) Lock()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_lock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unlocks the audio driver's main loop. (After locking it, you should always unlock it.)
*/
//go:nosplit
func (self class) Unlock()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_unlock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the speaker configuration.
*/
//go:nosplit
func (self class) GetSpeakerMode() classdb.AudioServerSpeakerMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioServerSpeakerMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_speaker_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the sample rate at the output of the [AudioServer].
*/
//go:nosplit
func (self class) GetMixRate() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_mix_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the names of all audio output devices detected on the system.
*/
//go:nosplit
func (self class) GetOutputDeviceList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_output_device_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetOutputDevice(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_output_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOutputDevice(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_output_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the relative time until the next mix occurs.
*/
//go:nosplit
func (self class) GetTimeToNextMix() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_time_to_next_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the relative time since the last mix occurred.
*/
//go:nosplit
func (self class) GetTimeSinceLastMix() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_time_since_last_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the audio driver's effective output latency. This is based on [member ProjectSettings.audio/driver/output_latency], but the exact returned value will differ depending on the operating system and audio driver.
[b]Note:[/b] This can be expensive; it is not recommended to call [method get_output_latency] every frame.
*/
//go:nosplit
func (self class) GetOutputLatency() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_output_latency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the names of all audio input devices detected on the system.
[b]Note:[/b] [member ProjectSettings.audio/driver/enable_input] must be [code]true[/code] for audio input to work. See also that setting's description for caveats related to permissions and operating system privacy settings.
*/
//go:nosplit
func (self class) GetInputDeviceList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_input_device_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetInputDevice(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_get_input_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInputDevice(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_input_device, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Overwrites the currently used [AudioBusLayout].
*/
//go:nosplit
func (self class) SetBusLayout(bus_layout object.AudioBusLayout)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bus_layout[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_bus_layout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Generates an [AudioBusLayout] using the available buses and effects.
*/
//go:nosplit
func (self class) GenerateBusLayout(ctx gd.Lifetime) object.AudioBusLayout {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_generate_bus_layout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioBusLayout
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
If set to [code]true[/code], all instances of [AudioStreamPlayback] will call [method AudioStreamPlayback._tag_used_streams] every mix step.
[b]Note:[/b] This is enabled by default in the editor, as it is used by editor plugins for the audio stream previews.
*/
//go:nosplit
func (self class) SetEnableTaggingUsedAudioStreams(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_set_enable_tagging_used_audio_streams, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the stream is registered as a sample. The engine will not have to register it before playing the sample.
If [code]false[/code], the stream will have to be registered before playing it. To prevent lag spikes, register the stream as sample with [method register_stream_as_sample].
*/
//go:nosplit
func (self class) IsStreamRegisteredAsSample(stream object.AudioStream) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_is_stream_registered_as_sample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Forces the registration of a stream as a sample.
[b]Note:[/b] Lag spikes may occur when calling this method, especially on single-threaded builds. It is suggested to call this method while loading assets, where the lag spike could be masked, instead of registering the sample right before it needs to be played.
*/
//go:nosplit
func (self class) RegisterStreamAsSample(stream object.AudioStream)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioServer.Bind_register_stream_as_sample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsAudioServer() Expert { return self[0].AsAudioServer() }


//go:nosplit
func (self Simple) AsAudioServer() Simple { return self[0].AsAudioServer() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioServer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type SpeakerMode = classdb.AudioServerSpeakerMode

const (
/*Two or fewer speakers were detected.*/
	SpeakerModeStereo SpeakerMode = 0
/*A 3.1 channel surround setup was detected.*/
	SpeakerSurround31 SpeakerMode = 1
/*A 5.1 channel surround setup was detected.*/
	SpeakerSurround51 SpeakerMode = 2
/*A 7.1 channel surround setup was detected.*/
	SpeakerSurround71 SpeakerMode = 3
)
type PlaybackType = classdb.AudioServerPlaybackType

const (
/*The playback will be considered of the type declared at [member ProjectSettings.audio/general/default_playback_type].*/
	PlaybackTypeDefault PlaybackType = 0
/*Force the playback to be considered as a stream.*/
	PlaybackTypeStream PlaybackType = 1
/*Force the playback to be considered as a sample. This can provide lower latency and more stable playback (with less risk of audio crackling), at the cost of having less flexibility.
[b]Note:[/b] Only currently supported on the web platform.
[b]Note:[/b] [AudioEffect]s are not supported when playback is considered as a sample.*/
	PlaybackTypeSample PlaybackType = 2
/*Represents the size of the [enum PlaybackType] enum.*/
	PlaybackTypeMax PlaybackType = 3
)
