package RandomNumberGenerator

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
RandomNumberGenerator is a class for generating pseudo-random numbers. It currently uses [url=https://www.pcg-random.org/]PCG32[/url].
[b]Note:[/b] The underlying algorithm is an implementation detail and should not be depended upon.
To generate a random float number (within a given range) based on a time-dependent seed:
[codeblock]
var rng = RandomNumberGenerator.new()
func _ready():
    var my_random_number = rng.randf_range(-10.0, 10.0)
[/codeblock]

*/
type Simple [1]classdb.RandomNumberGenerator
func (self Simple) SetSeed(seed int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSeed(gd.Int(seed))
}
func (self Simple) GetSeed() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSeed()))
}
func (self Simple) SetState(state int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetState(gd.Int(state))
}
func (self Simple) GetState() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetState()))
}
func (self Simple) Randi() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Randi()))
}
func (self Simple) Randf() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).Randf()))
}
func (self Simple) Randfn(mean float64, deviation float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).Randfn(gd.Float(mean), gd.Float(deviation))))
}
func (self Simple) RandfRange(from float64, to float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).RandfRange(gd.Float(from), gd.Float(to))))
}
func (self Simple) RandiRange(from int, to int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).RandiRange(gd.Int(from), gd.Int(to))))
}
func (self Simple) RandWeighted(weights gd.PackedFloat32Array) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).RandWeighted(weights)))
}
func (self Simple) Randomize() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Randomize()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RandomNumberGenerator
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetSeed(seed gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, seed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_set_seed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSeed() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_get_seed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetState(state gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, state)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_set_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetState() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_get_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a pseudo-random 32-bit unsigned integer between [code]0[/code] and [code]4294967295[/code] (inclusive).
*/
//go:nosplit
func (self class) Randi() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_randi, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a pseudo-random float between [code]0.0[/code] and [code]1.0[/code] (inclusive).
*/
//go:nosplit
func (self class) Randf() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_randf, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [url=https://en.wikipedia.org/wiki/Normal_distribution]normally-distributed[/url], pseudo-random floating-point number from the specified [param mean] and a standard [param deviation]. This is also known as a Gaussian distribution.
[b]Note:[/b] This method uses the [url=https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform]Box-Muller transform[/url] algorithm.
*/
//go:nosplit
func (self class) Randfn(mean gd.Float, deviation gd.Float) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mean)
	callframe.Arg(frame, deviation)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_randfn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a pseudo-random float between [param from] and [param to] (inclusive).
*/
//go:nosplit
func (self class) RandfRange(from gd.Float, to gd.Float) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_randf_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a pseudo-random 32-bit signed integer between [param from] and [param to] (inclusive).
*/
//go:nosplit
func (self class) RandiRange(from gd.Int, to gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_randi_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a random index with non-uniform weights. Prints an error and returns [code]-1[/code] if the array is empty.
[codeblocks]
[gdscript]
var rng = RandomNumberGenerator.new()

var my_array = ["one", "two", "three", "four"]
var weights = PackedFloat32Array([0.5, 1, 1, 2])

# Prints one of the four elements in `my_array`.
# It is more likely to print "four", and less likely to print "one".
print(my_array[rng.rand_weighted(weights)])
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) RandWeighted(weights gd.PackedFloat32Array) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(weights))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_rand_weighted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets up a time-based seed for this [RandomNumberGenerator] instance. Unlike the [@GlobalScope] random number generation functions, different [RandomNumberGenerator] instances can use different seeds.
*/
//go:nosplit
func (self class) Randomize()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RandomNumberGenerator.Bind_randomize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsRandomNumberGenerator() Expert { return self[0].AsRandomNumberGenerator() }


//go:nosplit
func (self Simple) AsRandomNumberGenerator() Simple { return self[0].AsRandomNumberGenerator() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RandomNumberGenerator", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
