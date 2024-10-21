package FastNoiseLite

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Noise"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class generates noise using the FastNoiseLite library, which is a collection of several noise algorithms including Cellular, Perlin, Value, and more.
Most generated noise values are in the range of [code][-1, 1][/code], but not always. Some of the cellular noise algorithms return results above [code]1[/code].

*/
type Simple [1]classdb.FastNoiseLite
func (self Simple) SetNoiseType(atype classdb.FastNoiseLiteNoiseType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNoiseType(atype)
}
func (self Simple) GetNoiseType() classdb.FastNoiseLiteNoiseType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.FastNoiseLiteNoiseType(Expert(self).GetNoiseType())
}
func (self Simple) SetSeed(seed int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSeed(gd.Int(seed))
}
func (self Simple) GetSeed() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSeed()))
}
func (self Simple) SetFrequency(freq float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFrequency(gd.Float(freq))
}
func (self Simple) GetFrequency() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFrequency()))
}
func (self Simple) SetOffset(offset gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOffset(offset)
}
func (self Simple) GetOffset() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetOffset())
}
func (self Simple) SetFractalType(atype classdb.FastNoiseLiteFractalType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFractalType(atype)
}
func (self Simple) GetFractalType() classdb.FastNoiseLiteFractalType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.FastNoiseLiteFractalType(Expert(self).GetFractalType())
}
func (self Simple) SetFractalOctaves(octave_count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFractalOctaves(gd.Int(octave_count))
}
func (self Simple) GetFractalOctaves() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFractalOctaves()))
}
func (self Simple) SetFractalLacunarity(lacunarity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFractalLacunarity(gd.Float(lacunarity))
}
func (self Simple) GetFractalLacunarity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFractalLacunarity()))
}
func (self Simple) SetFractalGain(gain float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFractalGain(gd.Float(gain))
}
func (self Simple) GetFractalGain() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFractalGain()))
}
func (self Simple) SetFractalWeightedStrength(weighted_strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFractalWeightedStrength(gd.Float(weighted_strength))
}
func (self Simple) GetFractalWeightedStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFractalWeightedStrength()))
}
func (self Simple) SetFractalPingPongStrength(ping_pong_strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFractalPingPongStrength(gd.Float(ping_pong_strength))
}
func (self Simple) GetFractalPingPongStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFractalPingPongStrength()))
}
func (self Simple) SetCellularDistanceFunction(fn classdb.FastNoiseLiteCellularDistanceFunction) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellularDistanceFunction(fn)
}
func (self Simple) GetCellularDistanceFunction() classdb.FastNoiseLiteCellularDistanceFunction {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.FastNoiseLiteCellularDistanceFunction(Expert(self).GetCellularDistanceFunction())
}
func (self Simple) SetCellularJitter(jitter float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellularJitter(gd.Float(jitter))
}
func (self Simple) GetCellularJitter() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCellularJitter()))
}
func (self Simple) SetCellularReturnType(ret classdb.FastNoiseLiteCellularReturnType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellularReturnType(ret)
}
func (self Simple) GetCellularReturnType() classdb.FastNoiseLiteCellularReturnType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.FastNoiseLiteCellularReturnType(Expert(self).GetCellularReturnType())
}
func (self Simple) SetDomainWarpEnabled(domain_warp_enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDomainWarpEnabled(domain_warp_enabled)
}
func (self Simple) IsDomainWarpEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDomainWarpEnabled())
}
func (self Simple) SetDomainWarpType(domain_warp_type classdb.FastNoiseLiteDomainWarpType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDomainWarpType(domain_warp_type)
}
func (self Simple) GetDomainWarpType() classdb.FastNoiseLiteDomainWarpType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.FastNoiseLiteDomainWarpType(Expert(self).GetDomainWarpType())
}
func (self Simple) SetDomainWarpAmplitude(domain_warp_amplitude float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDomainWarpAmplitude(gd.Float(domain_warp_amplitude))
}
func (self Simple) GetDomainWarpAmplitude() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDomainWarpAmplitude()))
}
func (self Simple) SetDomainWarpFrequency(domain_warp_frequency float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDomainWarpFrequency(gd.Float(domain_warp_frequency))
}
func (self Simple) GetDomainWarpFrequency() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDomainWarpFrequency()))
}
func (self Simple) SetDomainWarpFractalType(domain_warp_fractal_type classdb.FastNoiseLiteDomainWarpFractalType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDomainWarpFractalType(domain_warp_fractal_type)
}
func (self Simple) GetDomainWarpFractalType() classdb.FastNoiseLiteDomainWarpFractalType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.FastNoiseLiteDomainWarpFractalType(Expert(self).GetDomainWarpFractalType())
}
func (self Simple) SetDomainWarpFractalOctaves(domain_warp_octave_count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDomainWarpFractalOctaves(gd.Int(domain_warp_octave_count))
}
func (self Simple) GetDomainWarpFractalOctaves() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDomainWarpFractalOctaves()))
}
func (self Simple) SetDomainWarpFractalLacunarity(domain_warp_lacunarity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDomainWarpFractalLacunarity(gd.Float(domain_warp_lacunarity))
}
func (self Simple) GetDomainWarpFractalLacunarity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDomainWarpFractalLacunarity()))
}
func (self Simple) SetDomainWarpFractalGain(domain_warp_gain float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDomainWarpFractalGain(gd.Float(domain_warp_gain))
}
func (self Simple) GetDomainWarpFractalGain() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDomainWarpFractalGain()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.FastNoiseLite
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetNoiseType(atype classdb.FastNoiseLiteNoiseType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_noise_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNoiseType() classdb.FastNoiseLiteNoiseType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FastNoiseLiteNoiseType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_noise_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSeed(seed gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, seed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_seed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSeed() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_seed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrequency(freq gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, freq)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrequency() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOffset(offset gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOffset() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFractalType(atype classdb.FastNoiseLiteFractalType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_fractal_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFractalType() classdb.FastNoiseLiteFractalType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FastNoiseLiteFractalType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_fractal_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFractalOctaves(octave_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, octave_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_fractal_octaves, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFractalOctaves() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_fractal_octaves, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFractalLacunarity(lacunarity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lacunarity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_fractal_lacunarity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFractalLacunarity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_fractal_lacunarity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFractalGain(gain gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gain)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_fractal_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFractalGain() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_fractal_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFractalWeightedStrength(weighted_strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, weighted_strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_fractal_weighted_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFractalWeightedStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_fractal_weighted_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFractalPingPongStrength(ping_pong_strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ping_pong_strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_fractal_ping_pong_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFractalPingPongStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_fractal_ping_pong_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCellularDistanceFunction(fn classdb.FastNoiseLiteCellularDistanceFunction)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_cellular_distance_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellularDistanceFunction() classdb.FastNoiseLiteCellularDistanceFunction {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FastNoiseLiteCellularDistanceFunction](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_cellular_distance_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCellularJitter(jitter gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, jitter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_cellular_jitter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellularJitter() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_cellular_jitter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCellularReturnType(ret classdb.FastNoiseLiteCellularReturnType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ret)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_cellular_return_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCellularReturnType() classdb.FastNoiseLiteCellularReturnType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FastNoiseLiteCellularReturnType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_cellular_return_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDomainWarpEnabled(domain_warp_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_domain_warp_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDomainWarpEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_is_domain_warp_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDomainWarpType(domain_warp_type classdb.FastNoiseLiteDomainWarpType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_domain_warp_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDomainWarpType() classdb.FastNoiseLiteDomainWarpType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FastNoiseLiteDomainWarpType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_domain_warp_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDomainWarpAmplitude(domain_warp_amplitude gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_amplitude)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_domain_warp_amplitude, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDomainWarpAmplitude() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_domain_warp_amplitude, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDomainWarpFrequency(domain_warp_frequency gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_frequency)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_domain_warp_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDomainWarpFrequency() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_domain_warp_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDomainWarpFractalType(domain_warp_fractal_type classdb.FastNoiseLiteDomainWarpFractalType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_fractal_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_domain_warp_fractal_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDomainWarpFractalType() classdb.FastNoiseLiteDomainWarpFractalType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.FastNoiseLiteDomainWarpFractalType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_domain_warp_fractal_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDomainWarpFractalOctaves(domain_warp_octave_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_octave_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_domain_warp_fractal_octaves, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDomainWarpFractalOctaves() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_domain_warp_fractal_octaves, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDomainWarpFractalLacunarity(domain_warp_lacunarity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_lacunarity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_domain_warp_fractal_lacunarity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDomainWarpFractalLacunarity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_domain_warp_fractal_lacunarity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDomainWarpFractalGain(domain_warp_gain gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_gain)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_set_domain_warp_fractal_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDomainWarpFractalGain() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FastNoiseLite.Bind_get_domain_warp_fractal_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsFastNoiseLite() Expert { return self[0].AsFastNoiseLite() }


//go:nosplit
func (self Simple) AsFastNoiseLite() Simple { return self[0].AsFastNoiseLite() }


//go:nosplit
func (self class) AsNoise() Noise.Expert { return self[0].AsNoise() }


//go:nosplit
func (self Simple) AsNoise() Noise.Simple { return self[0].AsNoise() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

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
func init() {classdb.Register("FastNoiseLite", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type NoiseType = classdb.FastNoiseLiteNoiseType

const (
/*A lattice of points are assigned random values then interpolated based on neighboring values.*/
	TypeValue NoiseType = 5
/*Similar to Value noise, but slower. Has more variance in peaks and valleys.
Cubic noise can be used to avoid certain artifacts when using value noise to create a bumpmap. In general, you should always use this mode if the value noise is being used for a heightmap or bumpmap.*/
	TypeValueCubic NoiseType = 4
/*A lattice of random gradients. Their dot products are interpolated to obtain values in between the lattices.*/
	TypePerlin NoiseType = 3
/*Cellular includes both Worley noise and Voronoi diagrams which creates various regions of the same value.*/
	TypeCellular NoiseType = 2
/*As opposed to [constant TYPE_PERLIN], gradients exist in a simplex lattice rather than a grid lattice, avoiding directional artifacts.*/
	TypeSimplex NoiseType = 0
/*Modified, higher quality version of [constant TYPE_SIMPLEX], but slower.*/
	TypeSimplexSmooth NoiseType = 1
)
type FractalType = classdb.FastNoiseLiteFractalType

const (
/*No fractal noise.*/
	FractalNone FractalType = 0
/*Method using Fractional Brownian Motion to combine octaves into a fractal.*/
	FractalFbm FractalType = 1
/*Method of combining octaves into a fractal resulting in a "ridged" look.*/
	FractalRidged FractalType = 2
/*Method of combining octaves into a fractal with a ping pong effect.*/
	FractalPingPong FractalType = 3
)
type CellularDistanceFunction = classdb.FastNoiseLiteCellularDistanceFunction

const (
/*Euclidean distance to the nearest point.*/
	DistanceEuclidean CellularDistanceFunction = 0
/*Squared Euclidean distance to the nearest point.*/
	DistanceEuclideanSquared CellularDistanceFunction = 1
/*Manhattan distance (taxicab metric) to the nearest point.*/
	DistanceManhattan CellularDistanceFunction = 2
/*Blend of [constant DISTANCE_EUCLIDEAN] and [constant DISTANCE_MANHATTAN] to give curved cell boundaries*/
	DistanceHybrid CellularDistanceFunction = 3
)
type CellularReturnType = classdb.FastNoiseLiteCellularReturnType

const (
/*The cellular distance function will return the same value for all points within a cell.*/
	ReturnCellValue CellularReturnType = 0
/*The cellular distance function will return a value determined by the distance to the nearest point.*/
	ReturnDistance CellularReturnType = 1
/*The cellular distance function returns the distance to the second-nearest point.*/
	ReturnDistance2 CellularReturnType = 2
/*The distance to the nearest point is added to the distance to the second-nearest point.*/
	ReturnDistance2Add CellularReturnType = 3
/*The distance to the nearest point is subtracted from the distance to the second-nearest point.*/
	ReturnDistance2Sub CellularReturnType = 4
/*The distance to the nearest point is multiplied with the distance to the second-nearest point.*/
	ReturnDistance2Mul CellularReturnType = 5
/*The distance to the nearest point is divided by the distance to the second-nearest point.*/
	ReturnDistance2Div CellularReturnType = 6
)
type DomainWarpType = classdb.FastNoiseLiteDomainWarpType

const (
/*The domain is warped using the simplex noise algorithm.*/
	DomainWarpSimplex DomainWarpType = 0
/*The domain is warped using a simplified version of the simplex noise algorithm.*/
	DomainWarpSimplexReduced DomainWarpType = 1
/*The domain is warped using a simple noise grid (not as smooth as the other methods, but more performant).*/
	DomainWarpBasicGrid DomainWarpType = 2
)
type DomainWarpFractalType = classdb.FastNoiseLiteDomainWarpFractalType

const (
/*No fractal noise for warping the space.*/
	DomainWarpFractalNone DomainWarpFractalType = 0
/*Warping the space progressively, octave for octave, resulting in a more "liquified" distortion.*/
	DomainWarpFractalProgressive DomainWarpFractalType = 1
/*Warping the space independently for each octave, resulting in a more chaotic distortion.*/
	DomainWarpFractalIndependent DomainWarpFractalType = 2
)
