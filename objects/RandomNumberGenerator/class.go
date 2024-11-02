package RandomNumberGenerator

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.RandomNumberGenerator

/*
Returns a pseudo-random 32-bit unsigned integer between [code]0[/code] and [code]4294967295[/code] (inclusive).
*/
func (self Instance) Randi() int {
	return int(int(class(self).Randi()))
}

/*
Returns a pseudo-random float between [code]0.0[/code] and [code]1.0[/code] (inclusive).
*/
func (self Instance) Randf() float64 {
	return float64(float64(class(self).Randf()))
}

/*
Returns a [url=https://en.wikipedia.org/wiki/Normal_distribution]normally-distributed[/url], pseudo-random floating-point number from the specified [param mean] and a standard [param deviation]. This is also known as a Gaussian distribution.
[b]Note:[/b] This method uses the [url=https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform]Box-Muller transform[/url] algorithm.
*/
func (self Instance) Randfn() float64 {
	return float64(float64(class(self).Randfn(gd.Float(0.0), gd.Float(1.0))))
}

/*
Returns a pseudo-random float between [param from] and [param to] (inclusive).
*/
func (self Instance) RandfRange(from float64, to float64) float64 {
	return float64(float64(class(self).RandfRange(gd.Float(from), gd.Float(to))))
}

/*
Returns a pseudo-random 32-bit signed integer between [param from] and [param to] (inclusive).
*/
func (self Instance) RandiRange(from int, to int) int {
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
func (self Instance) RandWeighted(weights []float32) int {
	return int(int(class(self).RandWeighted(gd.NewPackedFloat32Slice(weights))))
}

/*
Sets up a time-based seed for this [RandomNumberGenerator] instance. Unlike the [@GlobalScope] random number generation functions, different [RandomNumberGenerator] instances can use different seeds.
*/
func (self Instance) Randomize() {
	class(self).Randomize()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RandomNumberGenerator

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RandomNumberGenerator"))
	return Instance{classdb.RandomNumberGenerator(object)}
}

func (self Instance) Seed() int {
	return int(int(class(self).GetSeed()))
}

func (self Instance) SetSeed(value int) {
	class(self).SetSeed(gd.Int(value))
}

func (self Instance) State() int {
	return int(int(class(self).GetState()))
}

func (self Instance) SetState(value int) {
	class(self).SetState(gd.Int(value))
}

//go:nosplit
func (self class) SetSeed(seed gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, seed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_set_seed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSeed() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_get_seed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetState(state gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, state)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_set_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetState() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_get_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a pseudo-random 32-bit unsigned integer between [code]0[/code] and [code]4294967295[/code] (inclusive).
*/
//go:nosplit
func (self class) Randi() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randi, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a pseudo-random float between [code]0.0[/code] and [code]1.0[/code] (inclusive).
*/
//go:nosplit
func (self class) Randf() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randf, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, mean)
	callframe.Arg(frame, deviation)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randfn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a pseudo-random float between [param from] and [param to] (inclusive).
*/
//go:nosplit
func (self class) RandfRange(from gd.Float, to gd.Float) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randf_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a pseudo-random 32-bit signed integer between [param from] and [param to] (inclusive).
*/
//go:nosplit
func (self class) RandiRange(from gd.Int, to gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randi_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(weights))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_rand_weighted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets up a time-based seed for this [RandomNumberGenerator] instance. Unlike the [@GlobalScope] random number generation functions, different [RandomNumberGenerator] instances can use different seeds.
*/
//go:nosplit
func (self class) Randomize() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randomize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsRandomNumberGenerator() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRandomNumberGenerator() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted          { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted       { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("RandomNumberGenerator", func(ptr gd.Object) any { return classdb.RandomNumberGenerator(ptr) })
}
