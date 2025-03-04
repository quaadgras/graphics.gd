// Package AnimationNodeExtension provides methods for working with AnimationNodeExtension object instances.
package AnimationNodeExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AnimationNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[AnimationNodeExtension] exposes the APIs of [AnimationRootNode] to allow users to extend it from GDScript, C#, or C++. This class is not meant to be used directly, but to be extended by other classes. It is used to create custom nodes for the [AnimationTree] system.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=AnimationNodeExtension)
*/
type Instance [1]gdclass.AnimationNodeExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationNodeExtension() Instance
}
type Interface interface {
	//A version of the [method AnimationNode._process] method that is meant to be overridden by custom nodes. It returns a [PackedFloat32Array] with the processed animation data.
	//The [PackedFloat64Array] parameter contains the playback information, containing the following values encoded as floating point numbers (in order): playback time and delta, start and end times, whether a seek was requested (encoded as a float greater than [code]0[/code]), whether the seek request was externally requested (encoded as a float greater than [code]0[/code]), the current [enum Animation.LoopedFlag] (encoded as a float), and the current blend weight.
	//The function must return a [PackedFloat32Array] of the node's time info, containing the following values (in order): animation length, time position, delta, [enum Animation.LoopMode] (encoded as a float), whether the animation is about to end (encoded as a float greater than [code]0[/code]) and whether the animation is infinite (encoded as a float greater than [code]0[/code]). All values must be included in the returned array.
	ProcessAnimationNode(playback_info []float64, test_only bool) []float32
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) ProcessAnimationNode(playback_info []float64, test_only bool) (_ []float32) {
	return
}

/*
A version of the [method AnimationNode._process] method that is meant to be overridden by custom nodes. It returns a [PackedFloat32Array] with the processed animation data.
The [PackedFloat64Array] parameter contains the playback information, containing the following values encoded as floating point numbers (in order): playback time and delta, start and end times, whether a seek was requested (encoded as a float greater than [code]0[/code]), whether the seek request was externally requested (encoded as a float greater than [code]0[/code]), the current [enum Animation.LoopedFlag] (encoded as a float), and the current blend weight.
The function must return a [PackedFloat32Array] of the node's time info, containing the following values (in order): animation length, time position, delta, [enum Animation.LoopMode] (encoded as a float), whether the animation is about to end (encoded as a float greater than [code]0[/code]) and whether the animation is infinite (encoded as a float greater than [code]0[/code]). All values must be included in the returned array.
*/
func (Instance) _process_animation_node(impl func(ptr unsafe.Pointer, playback_info []float64, test_only bool) []float32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var playback_info = Packed.Array[float64](Array.Through(gd.PackedProxy[gd.PackedFloat64Array, float64]{}, pointers.Pack(pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 0)))))
		defer pointers.End(gd.InternalPacked[gd.PackedFloat64Array, float64](playback_info))
		var test_only = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, slices.Collect(playback_info.Values()), test_only)
		ptr, ok := pointers.End(gd.InternalPacked[gd.PackedFloat32Array, float32](Packed.New(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if the animation for the given [param node_info] is looping.
*/
func IsLooping(node_info []float32) bool { //gd:AnimationNodeExtension.is_looping
	self := Instance{}
	return bool(class(self).IsLooping(Packed.New(node_info...)))
}

/*
Returns the animation's remaining time for the given node info. For looping animations, it will only return the remaining time if [param break_loop] is [code]true[/code], a large integer value will be returned otherwise.
*/
func GetRemainingTime(node_info []float32, break_loop bool) Float.X { //gd:AnimationNodeExtension.get_remaining_time
	self := Instance{}
	return Float.X(Float.X(class(self).GetRemainingTime(Packed.New(node_info...), break_loop)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationNodeExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeExtension"))
	casted := Instance{*(*gdclass.AnimationNodeExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
A version of the [method AnimationNode._process] method that is meant to be overridden by custom nodes. It returns a [PackedFloat32Array] with the processed animation data.
The [PackedFloat64Array] parameter contains the playback information, containing the following values encoded as floating point numbers (in order): playback time and delta, start and end times, whether a seek was requested (encoded as a float greater than [code]0[/code]), whether the seek request was externally requested (encoded as a float greater than [code]0[/code]), the current [enum Animation.LoopedFlag] (encoded as a float), and the current blend weight.
The function must return a [PackedFloat32Array] of the node's time info, containing the following values (in order): animation length, time position, delta, [enum Animation.LoopMode] (encoded as a float), whether the animation is about to end (encoded as a float greater than [code]0[/code]) and whether the animation is infinite (encoded as a float greater than [code]0[/code]). All values must be included in the returned array.
*/
func (class) _process_animation_node(impl func(ptr unsafe.Pointer, playback_info Packed.Array[float64], test_only bool) Packed.Array[float32]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var playback_info = Packed.Array[float64](Array.Through(gd.PackedProxy[gd.PackedFloat64Array, float64]{}, pointers.Pack(pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 0)))))
		defer pointers.End(gd.InternalPacked[gd.PackedFloat64Array, float64](playback_info))
		var test_only = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, playback_info, test_only)
		ptr, ok := pointers.End(gd.InternalPacked[gd.PackedFloat32Array, float32](ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if the animation for the given [param node_info] is looping.
*/
//go:nosplit
func (self class) IsLooping(node_info Packed.Array[float32]) bool { //gd:AnimationNodeExtension.is_looping
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](node_info)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeExtension.Bind_is_looping, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the animation's remaining time for the given node info. For looping animations, it will only return the remaining time if [param break_loop] is [code]true[/code], a large integer value will be returned otherwise.
*/
//go:nosplit
func (self class) GetRemainingTime(node_info Packed.Array[float32], break_loop bool) float64 { //gd:AnimationNodeExtension.get_remaining_time
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](node_info)))
	callframe.Arg(frame, break_loop)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeExtension.Bind_get_remaining_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationNode() AnimationNode.Advanced {
	return *((*AnimationNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNode() AnimationNode.Instance {
	return *((*AnimationNode.Instance)(unsafe.Pointer(&self)))
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
	case "_process_animation_node":
		return reflect.ValueOf(self._process_animation_node)
	default:
		return gd.VirtualByName(AnimationNode.Advanced(self.AsAnimationNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_process_animation_node":
		return reflect.ValueOf(self._process_animation_node)
	default:
		return gd.VirtualByName(AnimationNode.Instance(self.AsAnimationNode()), name)
	}
}
func init() {
	gdclass.Register("AnimationNodeExtension", func(ptr gd.Object) any {
		return [1]gdclass.AnimationNodeExtension{*(*gdclass.AnimationNodeExtension)(unsafe.Pointer(&ptr))}
	})
}
