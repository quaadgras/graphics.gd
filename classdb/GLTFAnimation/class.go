// Package GLTFAnimation provides methods for working with GLTFAnimation object instances.
package GLTFAnimation

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
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

type Instance [1]gdclass.GLTFAnimation

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFAnimation() Instance
}

/*
Gets additional arbitrary data in this [GLTFAnimation] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
func (self Instance) GetAdditionalData(extension_name string) any {
	return any(class(self).GetAdditionalData(gd.NewStringName(extension_name)).Interface())
}

/*
Sets additional arbitrary data in this [GLTFAnimation] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
func (self Instance) SetAdditionalData(extension_name string, additional_data any) {
	class(self).SetAdditionalData(gd.NewStringName(extension_name), gd.NewVariant(additional_data))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFAnimation

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFAnimation"))
	casted := Instance{*(*gdclass.GLTFAnimation)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) OriginalName() string {
	return string(class(self).GetOriginalName().String())
}

func (self Instance) SetOriginalName(value string) {
	class(self).SetOriginalName(gd.NewString(value))
}

func (self Instance) Loop() bool {
	return bool(class(self).GetLoop())
}

func (self Instance) SetLoop(value bool) {
	class(self).SetLoop(value)
}

//go:nosplit
func (self class) GetOriginalName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAnimation.Bind_get_original_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOriginalName(original_name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(original_name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAnimation.Bind_set_original_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLoop() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAnimation.Bind_get_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoop(loop bool) {
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAnimation.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets additional arbitrary data in this [GLTFAnimation] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the return value can be anything you set. If nothing was set, the return value is null.
*/
//go:nosplit
func (self class) GetAdditionalData(extension_name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAnimation.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets additional arbitrary data in this [GLTFAnimation] instance. This can be used to keep per-node state data in [GLTFDocumentExtension] classes, which is important because they are stateless.
The first argument should be the [GLTFDocumentExtension] name (does not have to match the extension name in the GLTF file), and the second argument can be anything you want.
*/
//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	callframe.Arg(frame, pointers.Get(additional_data))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAnimation.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsGLTFAnimation() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFAnimation() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GLTFAnimation", func(ptr gd.Object) any {
		return [1]gdclass.GLTFAnimation{*(*gdclass.GLTFAnimation)(unsafe.Pointer(&ptr))}
	})
}
