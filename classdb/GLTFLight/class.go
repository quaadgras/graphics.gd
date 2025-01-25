// Package GLTFLight provides methods for working with GLTFLight object instances.
package GLTFLight

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Represents a light as defined by the [code]KHR_lights_punctual[/code] GLTF extension.
*/
type Instance [1]gdclass.GLTFLight

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFLight() Instance
}

/*
Create a new GLTFLight instance from the given Godot [Light3D] node.
*/
func FromNode(light_node [1]gdclass.Light3D) [1]gdclass.GLTFLight {
	self := Instance{}
	return [1]gdclass.GLTFLight(class(self).FromNode(light_node))
}

/*
Converts this GLTFLight instance into a Godot [Light3D] node.
*/
func (self Instance) ToNode() [1]gdclass.Light3D {
	return [1]gdclass.Light3D(class(self).ToNode())
}

/*
Creates a new GLTFLight instance by parsing the given [Dictionary].
*/
func FromDictionary(dictionary map[any]any) [1]gdclass.GLTFLight {
	self := Instance{}
	return [1]gdclass.GLTFLight(class(self).FromDictionary(gd.NewVariant(dictionary).Interface().(gd.Dictionary)))
}

/*
Serializes this GLTFLight instance into a [Dictionary].
*/
func (self Instance) ToDictionary() map[any]any {
	return map[any]any(gd.DictionaryAs[any, any](class(self).ToDictionary()))
}
func (self Instance) GetAdditionalData(extension_name string) any {
	return any(class(self).GetAdditionalData(gd.NewStringName(extension_name)).Interface())
}
func (self Instance) SetAdditionalData(extension_name string, additional_data any) {
	class(self).SetAdditionalData(gd.NewStringName(extension_name), gd.NewVariant(additional_data))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFLight

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFLight"))
	casted := Instance{*(*gdclass.GLTFLight)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Color() Color.RGBA {
	return Color.RGBA(class(self).GetColor())
}

func (self Instance) SetColor(value Color.RGBA) {
	class(self).SetColor(gd.Color(value))
}

func (self Instance) Intensity() Float.X {
	return Float.X(Float.X(class(self).GetIntensity()))
}

func (self Instance) SetIntensity(value Float.X) {
	class(self).SetIntensity(gd.Float(value))
}

func (self Instance) LightType() string {
	return string(class(self).GetLightType().String())
}

func (self Instance) SetLightType(value string) {
	class(self).SetLightType(gd.NewString(value))
}

func (self Instance) Range() Float.X {
	return Float.X(Float.X(class(self).GetRange()))
}

func (self Instance) SetRange(value Float.X) {
	class(self).SetRange(gd.Float(value))
}

func (self Instance) InnerConeAngle() Float.X {
	return Float.X(Float.X(class(self).GetInnerConeAngle()))
}

func (self Instance) SetInnerConeAngle(value Float.X) {
	class(self).SetInnerConeAngle(gd.Float(value))
}

func (self Instance) OuterConeAngle() Float.X {
	return Float.X(Float.X(class(self).GetOuterConeAngle()))
}

func (self Instance) SetOuterConeAngle(value Float.X) {
	class(self).SetOuterConeAngle(gd.Float(value))
}

/*
Create a new GLTFLight instance from the given Godot [Light3D] node.
*/
//go:nosplit
func (self class) FromNode(light_node [1]gdclass.Light3D) [1]gdclass.GLTFLight {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(light_node[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GLTFLight{gd.PointerWithOwnershipTransferredToGo[gdclass.GLTFLight](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Converts this GLTFLight instance into a Godot [Light3D] node.
*/
//go:nosplit
func (self class) ToNode() [1]gdclass.Light3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Light3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Light3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new GLTFLight instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(dictionary gd.Dictionary) [1]gdclass.GLTFLight {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dictionary))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GLTFLight{gd.PointerWithOwnershipTransferredToGo[gdclass.GLTFLight](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Serializes this GLTFLight instance into a [Dictionary].
*/
//go:nosplit
func (self class) ToDictionary() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIntensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIntensity(intensity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_intensity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLightType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_light_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLightType(light_type gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(light_type))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_light_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRange() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRange(arange gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, arange)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInnerConeAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_inner_cone_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInnerConeAngle(inner_cone_angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, inner_cone_angle)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_inner_cone_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOuterConeAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_outer_cone_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOuterConeAngle(outer_cone_angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, outer_cone_angle)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_outer_cone_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdditionalData(extension_name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	callframe.Arg(frame, pointers.Get(additional_data))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsGLTFLight() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFLight() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GLTFLight", func(ptr gd.Object) any { return [1]gdclass.GLTFLight{*(*gdclass.GLTFLight)(unsafe.Pointer(&ptr))} })
}
