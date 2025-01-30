// Package OpenXRInteractionProfile provides methods for working with OpenXRInteractionProfile object instances.
package OpenXRInteractionProfile

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
This object stores suggested bindings for an interaction profile. Interaction profiles define the metadata for a tracked XR device such as an XR controller.
For more information see the [url=https://www.khronos.org/registry/OpenXR/specs/1.0/html/xrspec.html#semantic-path-interaction-profiles]interaction profiles info in the OpenXR specification[/url].
*/
type Instance [1]gdclass.OpenXRInteractionProfile

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRInteractionProfile() Instance
}

/*
Get the number of bindings in this interaction profile.
*/
func (self Instance) GetBindingCount() int { //gd:OpenXRInteractionProfile.get_binding_count
	return int(int(class(self).GetBindingCount()))
}

/*
Retrieve the binding at this index.
*/
func (self Instance) GetBinding(index int) [1]gdclass.OpenXRIPBinding { //gd:OpenXRInteractionProfile.get_binding
	return [1]gdclass.OpenXRIPBinding(class(self).GetBinding(int64(index)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRInteractionProfile

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRInteractionProfile"))
	casted := Instance{*(*gdclass.OpenXRInteractionProfile)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) InteractionProfilePath() string {
	return string(class(self).GetInteractionProfilePath().String())
}

func (self Instance) SetInteractionProfilePath(value string) {
	class(self).SetInteractionProfilePath(String.New(value))
}

func (self Instance) Bindings() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetBindings())))
}

func (self Instance) SetBindings(value []any) {
	class(self).SetBindings(gd.EngineArrayFromSlice(value))
}

//go:nosplit
func (self class) SetInteractionProfilePath(interaction_profile_path String.Readable) { //gd:OpenXRInteractionProfile.set_interaction_profile_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(interaction_profile_path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_set_interaction_profile_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInteractionProfilePath() String.Readable { //gd:OpenXRInteractionProfile.get_interaction_profile_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_get_interaction_profile_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Get the number of bindings in this interaction profile.
*/
//go:nosplit
func (self class) GetBindingCount() int64 { //gd:OpenXRInteractionProfile.get_binding_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_get_binding_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Retrieve the binding at this index.
*/
//go:nosplit
func (self class) GetBinding(index int64) [1]gdclass.OpenXRIPBinding { //gd:OpenXRInteractionProfile.get_binding
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_get_binding, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRIPBinding{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRIPBinding](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBindings(bindings Array.Any) { //gd:OpenXRInteractionProfile.set_bindings
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(bindings)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_set_bindings, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindings() Array.Any { //gd:OpenXRInteractionProfile.get_bindings
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_get_bindings, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsOpenXRInteractionProfile() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRInteractionProfile() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("OpenXRInteractionProfile", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRInteractionProfile{*(*gdclass.OpenXRInteractionProfile)(unsafe.Pointer(&ptr))}
	})
}
