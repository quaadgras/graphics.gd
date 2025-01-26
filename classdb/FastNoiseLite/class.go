// Package FastNoiseLite provides methods for working with FastNoiseLite object instances.
package FastNoiseLite

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/Noise"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any

/*
This class generates noise using the FastNoiseLite library, which is a collection of several noise algorithms including Cellular, Perlin, Value, and more.
Most generated noise values are in the range of [code][-1, 1][/code], but not always. Some of the cellular noise algorithms return results above [code]1[/code].
*/
type Instance [1]gdclass.FastNoiseLite

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFastNoiseLite() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FastNoiseLite

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FastNoiseLite"))
	casted := Instance{*(*gdclass.FastNoiseLite)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) NoiseType() gdclass.FastNoiseLiteNoiseType {
	return gdclass.FastNoiseLiteNoiseType(class(self).GetNoiseType())
}

func (self Instance) SetNoiseType(value gdclass.FastNoiseLiteNoiseType) {
	class(self).SetNoiseType(value)
}

func (self Instance) Seed() int {
	return int(int(class(self).GetSeed()))
}

func (self Instance) SetSeed(value int) {
	class(self).SetSeed(gd.Int(value))
}

func (self Instance) Frequency() Float.X {
	return Float.X(Float.X(class(self).GetFrequency()))
}

func (self Instance) SetFrequency(value Float.X) {
	class(self).SetFrequency(gd.Float(value))
}

func (self Instance) Offset() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetOffset())
}

func (self Instance) SetOffset(value Vector3.XYZ) {
	class(self).SetOffset(gd.Vector3(value))
}

func (self Instance) FractalType() gdclass.FastNoiseLiteFractalType {
	return gdclass.FastNoiseLiteFractalType(class(self).GetFractalType())
}

func (self Instance) SetFractalType(value gdclass.FastNoiseLiteFractalType) {
	class(self).SetFractalType(value)
}

func (self Instance) FractalOctaves() int {
	return int(int(class(self).GetFractalOctaves()))
}

func (self Instance) SetFractalOctaves(value int) {
	class(self).SetFractalOctaves(gd.Int(value))
}

func (self Instance) FractalLacunarity() Float.X {
	return Float.X(Float.X(class(self).GetFractalLacunarity()))
}

func (self Instance) SetFractalLacunarity(value Float.X) {
	class(self).SetFractalLacunarity(gd.Float(value))
}

func (self Instance) FractalGain() Float.X {
	return Float.X(Float.X(class(self).GetFractalGain()))
}

func (self Instance) SetFractalGain(value Float.X) {
	class(self).SetFractalGain(gd.Float(value))
}

func (self Instance) FractalWeightedStrength() Float.X {
	return Float.X(Float.X(class(self).GetFractalWeightedStrength()))
}

func (self Instance) SetFractalWeightedStrength(value Float.X) {
	class(self).SetFractalWeightedStrength(gd.Float(value))
}

func (self Instance) FractalPingPongStrength() Float.X {
	return Float.X(Float.X(class(self).GetFractalPingPongStrength()))
}

func (self Instance) SetFractalPingPongStrength(value Float.X) {
	class(self).SetFractalPingPongStrength(gd.Float(value))
}

func (self Instance) CellularDistanceFunction() gdclass.FastNoiseLiteCellularDistanceFunction {
	return gdclass.FastNoiseLiteCellularDistanceFunction(class(self).GetCellularDistanceFunction())
}

func (self Instance) SetCellularDistanceFunction(value gdclass.FastNoiseLiteCellularDistanceFunction) {
	class(self).SetCellularDistanceFunction(value)
}

func (self Instance) CellularJitter() Float.X {
	return Float.X(Float.X(class(self).GetCellularJitter()))
}

func (self Instance) SetCellularJitter(value Float.X) {
	class(self).SetCellularJitter(gd.Float(value))
}

func (self Instance) CellularReturnType() gdclass.FastNoiseLiteCellularReturnType {
	return gdclass.FastNoiseLiteCellularReturnType(class(self).GetCellularReturnType())
}

func (self Instance) SetCellularReturnType(value gdclass.FastNoiseLiteCellularReturnType) {
	class(self).SetCellularReturnType(value)
}

func (self Instance) DomainWarpEnabled() bool {
	return bool(class(self).IsDomainWarpEnabled())
}

func (self Instance) SetDomainWarpEnabled(value bool) {
	class(self).SetDomainWarpEnabled(value)
}

func (self Instance) DomainWarpType() gdclass.FastNoiseLiteDomainWarpType {
	return gdclass.FastNoiseLiteDomainWarpType(class(self).GetDomainWarpType())
}

func (self Instance) SetDomainWarpType(value gdclass.FastNoiseLiteDomainWarpType) {
	class(self).SetDomainWarpType(value)
}

func (self Instance) DomainWarpAmplitude() Float.X {
	return Float.X(Float.X(class(self).GetDomainWarpAmplitude()))
}

func (self Instance) SetDomainWarpAmplitude(value Float.X) {
	class(self).SetDomainWarpAmplitude(gd.Float(value))
}

func (self Instance) DomainWarpFrequency() Float.X {
	return Float.X(Float.X(class(self).GetDomainWarpFrequency()))
}

func (self Instance) SetDomainWarpFrequency(value Float.X) {
	class(self).SetDomainWarpFrequency(gd.Float(value))
}

func (self Instance) DomainWarpFractalType() gdclass.FastNoiseLiteDomainWarpFractalType {
	return gdclass.FastNoiseLiteDomainWarpFractalType(class(self).GetDomainWarpFractalType())
}

func (self Instance) SetDomainWarpFractalType(value gdclass.FastNoiseLiteDomainWarpFractalType) {
	class(self).SetDomainWarpFractalType(value)
}

func (self Instance) DomainWarpFractalOctaves() int {
	return int(int(class(self).GetDomainWarpFractalOctaves()))
}

func (self Instance) SetDomainWarpFractalOctaves(value int) {
	class(self).SetDomainWarpFractalOctaves(gd.Int(value))
}

func (self Instance) DomainWarpFractalLacunarity() Float.X {
	return Float.X(Float.X(class(self).GetDomainWarpFractalLacunarity()))
}

func (self Instance) SetDomainWarpFractalLacunarity(value Float.X) {
	class(self).SetDomainWarpFractalLacunarity(gd.Float(value))
}

func (self Instance) DomainWarpFractalGain() Float.X {
	return Float.X(Float.X(class(self).GetDomainWarpFractalGain()))
}

func (self Instance) SetDomainWarpFractalGain(value Float.X) {
	class(self).SetDomainWarpFractalGain(gd.Float(value))
}

//go:nosplit
func (self class) SetNoiseType(atype gdclass.FastNoiseLiteNoiseType) { //gd:FastNoiseLite.set_noise_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_noise_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNoiseType() gdclass.FastNoiseLiteNoiseType { //gd:FastNoiseLite.get_noise_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FastNoiseLiteNoiseType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_noise_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSeed(seed gd.Int) { //gd:FastNoiseLite.set_seed
	var frame = callframe.New()
	callframe.Arg(frame, seed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_seed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSeed() gd.Int { //gd:FastNoiseLite.get_seed
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_seed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrequency(freq gd.Float) { //gd:FastNoiseLite.set_frequency
	var frame = callframe.New()
	callframe.Arg(frame, freq)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_frequency, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrequency() gd.Float { //gd:FastNoiseLite.get_frequency
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_frequency, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset gd.Vector3) { //gd:FastNoiseLite.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() gd.Vector3 { //gd:FastNoiseLite.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFractalType(atype gdclass.FastNoiseLiteFractalType) { //gd:FastNoiseLite.set_fractal_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_fractal_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFractalType() gdclass.FastNoiseLiteFractalType { //gd:FastNoiseLite.get_fractal_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FastNoiseLiteFractalType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_fractal_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFractalOctaves(octave_count gd.Int) { //gd:FastNoiseLite.set_fractal_octaves
	var frame = callframe.New()
	callframe.Arg(frame, octave_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_fractal_octaves, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFractalOctaves() gd.Int { //gd:FastNoiseLite.get_fractal_octaves
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_fractal_octaves, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFractalLacunarity(lacunarity gd.Float) { //gd:FastNoiseLite.set_fractal_lacunarity
	var frame = callframe.New()
	callframe.Arg(frame, lacunarity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_fractal_lacunarity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFractalLacunarity() gd.Float { //gd:FastNoiseLite.get_fractal_lacunarity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_fractal_lacunarity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFractalGain(gain gd.Float) { //gd:FastNoiseLite.set_fractal_gain
	var frame = callframe.New()
	callframe.Arg(frame, gain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_fractal_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFractalGain() gd.Float { //gd:FastNoiseLite.get_fractal_gain
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_fractal_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFractalWeightedStrength(weighted_strength gd.Float) { //gd:FastNoiseLite.set_fractal_weighted_strength
	var frame = callframe.New()
	callframe.Arg(frame, weighted_strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_fractal_weighted_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFractalWeightedStrength() gd.Float { //gd:FastNoiseLite.get_fractal_weighted_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_fractal_weighted_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFractalPingPongStrength(ping_pong_strength gd.Float) { //gd:FastNoiseLite.set_fractal_ping_pong_strength
	var frame = callframe.New()
	callframe.Arg(frame, ping_pong_strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_fractal_ping_pong_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFractalPingPongStrength() gd.Float { //gd:FastNoiseLite.get_fractal_ping_pong_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_fractal_ping_pong_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCellularDistanceFunction(fn gdclass.FastNoiseLiteCellularDistanceFunction) { //gd:FastNoiseLite.set_cellular_distance_function
	var frame = callframe.New()
	callframe.Arg(frame, fn)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_cellular_distance_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellularDistanceFunction() gdclass.FastNoiseLiteCellularDistanceFunction { //gd:FastNoiseLite.get_cellular_distance_function
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FastNoiseLiteCellularDistanceFunction](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_cellular_distance_function, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCellularJitter(jitter gd.Float) { //gd:FastNoiseLite.set_cellular_jitter
	var frame = callframe.New()
	callframe.Arg(frame, jitter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_cellular_jitter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellularJitter() gd.Float { //gd:FastNoiseLite.get_cellular_jitter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_cellular_jitter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCellularReturnType(ret gdclass.FastNoiseLiteCellularReturnType) { //gd:FastNoiseLite.set_cellular_return_type
	var frame = callframe.New()
	callframe.Arg(frame, ret)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_cellular_return_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellularReturnType() gdclass.FastNoiseLiteCellularReturnType { //gd:FastNoiseLite.get_cellular_return_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FastNoiseLiteCellularReturnType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_cellular_return_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDomainWarpEnabled(domain_warp_enabled bool) { //gd:FastNoiseLite.set_domain_warp_enabled
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_domain_warp_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDomainWarpEnabled() bool { //gd:FastNoiseLite.is_domain_warp_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_is_domain_warp_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDomainWarpType(domain_warp_type gdclass.FastNoiseLiteDomainWarpType) { //gd:FastNoiseLite.set_domain_warp_type
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_domain_warp_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDomainWarpType() gdclass.FastNoiseLiteDomainWarpType { //gd:FastNoiseLite.get_domain_warp_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FastNoiseLiteDomainWarpType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_domain_warp_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDomainWarpAmplitude(domain_warp_amplitude gd.Float) { //gd:FastNoiseLite.set_domain_warp_amplitude
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_amplitude)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_domain_warp_amplitude, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDomainWarpAmplitude() gd.Float { //gd:FastNoiseLite.get_domain_warp_amplitude
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_domain_warp_amplitude, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDomainWarpFrequency(domain_warp_frequency gd.Float) { //gd:FastNoiseLite.set_domain_warp_frequency
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_frequency)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_domain_warp_frequency, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDomainWarpFrequency() gd.Float { //gd:FastNoiseLite.get_domain_warp_frequency
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_domain_warp_frequency, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDomainWarpFractalType(domain_warp_fractal_type gdclass.FastNoiseLiteDomainWarpFractalType) { //gd:FastNoiseLite.set_domain_warp_fractal_type
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_fractal_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_domain_warp_fractal_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDomainWarpFractalType() gdclass.FastNoiseLiteDomainWarpFractalType { //gd:FastNoiseLite.get_domain_warp_fractal_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.FastNoiseLiteDomainWarpFractalType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_domain_warp_fractal_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDomainWarpFractalOctaves(domain_warp_octave_count gd.Int) { //gd:FastNoiseLite.set_domain_warp_fractal_octaves
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_octave_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_domain_warp_fractal_octaves, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDomainWarpFractalOctaves() gd.Int { //gd:FastNoiseLite.get_domain_warp_fractal_octaves
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_domain_warp_fractal_octaves, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDomainWarpFractalLacunarity(domain_warp_lacunarity gd.Float) { //gd:FastNoiseLite.set_domain_warp_fractal_lacunarity
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_lacunarity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_domain_warp_fractal_lacunarity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDomainWarpFractalLacunarity() gd.Float { //gd:FastNoiseLite.get_domain_warp_fractal_lacunarity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_domain_warp_fractal_lacunarity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDomainWarpFractalGain(domain_warp_gain gd.Float) { //gd:FastNoiseLite.set_domain_warp_fractal_gain
	var frame = callframe.New()
	callframe.Arg(frame, domain_warp_gain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_set_domain_warp_fractal_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDomainWarpFractalGain() gd.Float { //gd:FastNoiseLite.get_domain_warp_fractal_gain
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FastNoiseLite.Bind_get_domain_warp_fractal_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFastNoiseLite() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFastNoiseLite() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNoise() Noise.Advanced      { return *((*Noise.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNoise() Noise.Instance   { return *((*Noise.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Noise.Advanced(self.AsNoise()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Noise.Instance(self.AsNoise()), name)
	}
}
func init() {
	gdclass.Register("FastNoiseLite", func(ptr gd.Object) any {
		return [1]gdclass.FastNoiseLite{*(*gdclass.FastNoiseLite)(unsafe.Pointer(&ptr))}
	})
}

type NoiseType = gdclass.FastNoiseLiteNoiseType //gd:FastNoiseLite.NoiseType

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

type FractalType = gdclass.FastNoiseLiteFractalType //gd:FastNoiseLite.FractalType

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

type CellularDistanceFunction = gdclass.FastNoiseLiteCellularDistanceFunction //gd:FastNoiseLite.CellularDistanceFunction

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

type CellularReturnType = gdclass.FastNoiseLiteCellularReturnType //gd:FastNoiseLite.CellularReturnType

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

type DomainWarpType = gdclass.FastNoiseLiteDomainWarpType //gd:FastNoiseLite.DomainWarpType

const (
	/*The domain is warped using the simplex noise algorithm.*/
	DomainWarpSimplex DomainWarpType = 0
	/*The domain is warped using a simplified version of the simplex noise algorithm.*/
	DomainWarpSimplexReduced DomainWarpType = 1
	/*The domain is warped using a simple noise grid (not as smooth as the other methods, but more performant).*/
	DomainWarpBasicGrid DomainWarpType = 2
)

type DomainWarpFractalType = gdclass.FastNoiseLiteDomainWarpFractalType //gd:FastNoiseLite.DomainWarpFractalType

const (
	/*No fractal noise for warping the space.*/
	DomainWarpFractalNone DomainWarpFractalType = 0
	/*Warping the space progressively, octave for octave, resulting in a more "liquified" distortion.*/
	DomainWarpFractalProgressive DomainWarpFractalType = 1
	/*Warping the space independently for each octave, resulting in a more chaotic distortion.*/
	DomainWarpFractalIndependent DomainWarpFractalType = 2
)
