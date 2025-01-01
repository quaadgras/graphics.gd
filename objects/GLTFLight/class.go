package GLTFLight

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Represents a light as defined by the [code]KHR_lights_punctual[/code] GLTF extension.
*/
type Instance [1]classdb.GLTFLight
type Any interface {
	gd.IsClass
	AsGLTFLight() Instance
}

/*
Create a new GLTFLight instance from the given Godot [Light3D] node.
*/
func FromNode(light_node objects.Light3D) objects.GLTFLight {
	self := Instance{}
	return objects.GLTFLight(class(self).FromNode(light_node))
}

/*
Converts this GLTFLight instance into a Godot [Light3D] node.
*/
func (self Instance) ToNode() objects.Light3D {
	return objects.Light3D(class(self).ToNode())
}

/*
Creates a new GLTFLight instance by parsing the given [Dictionary].
*/
func FromDictionary(dictionary Dictionary.Any) objects.GLTFLight {
	self := Instance{}
	return objects.GLTFLight(class(self).FromDictionary(dictionary))
}

/*
Serializes this GLTFLight instance into a [Dictionary].
*/
func (self Instance) ToDictionary() Dictionary.Any {
	return Dictionary.Any(class(self).ToDictionary())
}
func (self Instance) GetAdditionalData(extension_name string) any {
	return any(class(self).GetAdditionalData(gd.NewStringName(extension_name)).Interface())
}
func (self Instance) SetAdditionalData(extension_name string, additional_data any) {
	class(self).SetAdditionalData(gd.NewStringName(extension_name), gd.NewVariant(additional_data))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GLTFLight

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFLight"))
	return Instance{classdb.GLTFLight(object)}
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
func (self class) FromNode(light_node objects.Light3D) objects.GLTFLight {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(light_node[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.GLTFLight{classdb.GLTFLight(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Converts this GLTFLight instance into a Godot [Light3D] node.
*/
//go:nosplit
func (self class) ToNode() objects.Light3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Light3D{classdb.Light3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Creates a new GLTFLight instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(dictionary gd.Dictionary) objects.GLTFLight {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(dictionary))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.GLTFLight{classdb.GLTFLight(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Serializes this GLTFLight instance into a [Dictionary].
*/
//go:nosplit
func (self class) ToDictionary() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIntensity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIntensity(intensity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLightType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_light_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLightType(light_type gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(light_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_light_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRange() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRange(arange gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, arange)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInnerConeAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_inner_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInnerConeAngle(inner_cone_angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, inner_cone_angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_inner_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOuterConeAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_outer_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOuterConeAngle(outer_cone_angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, outer_cone_angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_outer_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdditionalData(extension_name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension_name))
	callframe.Arg(frame, pointers.Get(additional_data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFLight.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("GLTFLight", func(ptr gd.Object) any { return [1]classdb.GLTFLight{classdb.GLTFLight(ptr)} })
}
