// Package SkeletonModifier3D provides methods for working with SkeletonModifier3D object instances.
package SkeletonModifier3D

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
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

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

/*
[SkeletonModifier3D] retrieves a target [Skeleton3D] by having a [Skeleton3D] parent.
If there is [AnimationMixer], modification always performs after playback process of the [AnimationMixer].
This node should be used to implement custom IK solvers, constraints, or skeleton physics.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=SkeletonModifier3D)
*/
type Instance [1]gdclass.SkeletonModifier3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeletonModifier3D() Instance
}
type Interface interface {
	//Override this virtual method to implement a custom skeleton modifier. You should do things like get the [Skeleton3D]'s current pose and apply the pose here.
	//[method _process_modification] must not apply [member influence] to bone poses because the [Skeleton3D] automatically applies influence to all bone poses set by the modifier.
	ProcessModification()
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) ProcessModification() { return }

/*
Override this virtual method to implement a custom skeleton modifier. You should do things like get the [Skeleton3D]'s current pose and apply the pose here.
[method _process_modification] must not apply [member influence] to bone poses because the [Skeleton3D] automatically applies influence to all bone poses set by the modifier.
*/
func (Instance) _process_modification(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Get parent [Skeleton3D] node if found.
*/
func (self Instance) GetSkeleton() [1]gdclass.Skeleton3D { //gd:SkeletonModifier3D.get_skeleton
	return [1]gdclass.Skeleton3D(class(self).GetSkeleton())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonModifier3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModifier3D"))
	casted := Instance{*(*gdclass.SkeletonModifier3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Active() bool {
	return bool(class(self).IsActive())
}

func (self Instance) SetActive(value bool) {
	class(self).SetActive(value)
}

func (self Instance) Influence() Float.X {
	return Float.X(Float.X(class(self).GetInfluence()))
}

func (self Instance) SetInfluence(value Float.X) {
	class(self).SetInfluence(gd.Float(value))
}

/*
Override this virtual method to implement a custom skeleton modifier. You should do things like get the [Skeleton3D]'s current pose and apply the pose here.
[method _process_modification] must not apply [member influence] to bone poses because the [Skeleton3D] automatically applies influence to all bone poses set by the modifier.
*/
func (class) _process_modification(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Get parent [Skeleton3D] node if found.
*/
//go:nosplit
func (self class) GetSkeleton() [1]gdclass.Skeleton3D { //gd:SkeletonModifier3D.get_skeleton
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModifier3D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skeleton3D{gd.PointerMustAssertInstanceID[gdclass.Skeleton3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetActive(active bool) { //gd:SkeletonModifier3D.set_active
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModifier3D.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsActive() bool { //gd:SkeletonModifier3D.is_active
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModifier3D.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInfluence(influence gd.Float) { //gd:SkeletonModifier3D.set_influence
	var frame = callframe.New()
	callframe.Arg(frame, influence)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModifier3D.Bind_set_influence, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInfluence() gd.Float { //gd:SkeletonModifier3D.get_influence
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModifier3D.Bind_get_influence, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnModificationProcessed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("modification_processed"), gd.NewCallable(cb), 0)
}

func (self class) AsSkeletonModifier3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSkeletonModifier3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced         { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance      { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced             { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance          { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process_modification":
		return reflect.ValueOf(self._process_modification)
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_process_modification":
		return reflect.ValueOf(self._process_modification)
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("SkeletonModifier3D", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonModifier3D{*(*gdclass.SkeletonModifier3D)(unsafe.Pointer(&ptr))}
	})
}
