package RandomNumberGenerator

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
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
type Go [1]classdb.RandomNumberGenerator

/*
Returns a pseudo-random 32-bit unsigned integer between [code]0[/code] and [code]4294967295[/code] (inclusive).
*/
func (self Go) Randi() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).Randi()))
}

/*
Returns a pseudo-random float between [code]0.0[/code] and [code]1.0[/code] (inclusive).
*/
func (self Go) Randf() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).Randf()))
}

/*
Returns a [url=https://en.wikipedia.org/wiki/Normal_distribution]normally-distributed[/url], pseudo-random floating-point number from the specified [param mean] and a standard [param deviation]. This is also known as a Gaussian distribution.
[b]Note:[/b] This method uses the [url=https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform]Box-Muller transform[/url] algorithm.
*/
func (self Go) Randfn() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).Randfn(gd.Float(0.0), gd.Float(1.0))))
}

/*
Returns a pseudo-random float between [param from] and [param to] (inclusive).
*/
func (self Go) RandfRange(from float64, to float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).RandfRange(gd.Float(from), gd.Float(to))))
}

/*
Returns a pseudo-random 32-bit signed integer between [param from] and [param to] (inclusive).
*/
func (self Go) RandiRange(from int, to int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).RandiRange(gd.Int(from), gd.Int(to))))
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
func (self Go) RandWeighted(weights []float32) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).RandWeighted(gc.PackedFloat32Slice(weights))))
}

/*
Sets up a time-based seed for this [RandomNumberGenerator] instance. Unlike the [@GlobalScope] random number generation functions, different [RandomNumberGenerator] instances can use different seeds.
*/
func (self Go) Randomize() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Randomize()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RandomNumberGenerator
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RandomNumberGenerator"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Seed() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSeed()))
}

func (self Go) SetSeed(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSeed(gd.Int(value))
}

func (self Go) State() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetState()))
}

func (self Go) SetState(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetState(gd.Int(value))
}

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
func (self class) AsRandomNumberGenerator() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRandomNumberGenerator() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("RandomNumberGenerator", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
