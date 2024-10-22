package GLTFLight

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Represents a light as defined by the [code]KHR_lights_punctual[/code] GLTF extension.

*/
type Go [1]classdb.GLTFLight

/*
Create a new GLTFLight instance from the given Godot [Light3D] node.
*/
func (self Go) FromNode(light_node gdclass.Light3D) gdclass.GLTFLight {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.GLTFLight(class(self).FromNode(gc, light_node))
}

/*
Converts this GLTFLight instance into a Godot [Light3D] node.
*/
func (self Go) ToNode() gdclass.Light3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Light3D(class(self).ToNode(gc))
}

/*
Creates a new GLTFLight instance by parsing the given [Dictionary].
*/
func (self Go) FromDictionary(dictionary gd.Dictionary) gdclass.GLTFLight {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.GLTFLight(class(self).FromDictionary(gc, dictionary))
}

/*
Serializes this GLTFLight instance into a [Dictionary].
*/
func (self Go) ToDictionary() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).ToDictionary(gc))
}
func (self Go) GetAdditionalData(extension_name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetAdditionalData(gc, gc.StringName(extension_name)))
}
func (self Go) SetAdditionalData(extension_name string, additional_data gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdditionalData(gc.StringName(extension_name), additional_data)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFLight
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFLight"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Color() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetColor())
}

func (self Go) SetColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColor(value)
}

func (self Go) Intensity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetIntensity()))
}

func (self Go) SetIntensity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIntensity(gd.Float(value))
}

func (self Go) LightType() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetLightType(gc).String())
}

func (self Go) SetLightType(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLightType(gc.String(value))
}

func (self Go) Range() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRange()))
}

func (self Go) SetRange(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRange(gd.Float(value))
}

func (self Go) InnerConeAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetInnerConeAngle()))
}

func (self Go) SetInnerConeAngle(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInnerConeAngle(gd.Float(value))
}

func (self Go) OuterConeAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetOuterConeAngle()))
}

func (self Go) SetOuterConeAngle(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOuterConeAngle(gd.Float(value))
}

/*
Create a new GLTFLight instance from the given Godot [Light3D] node.
*/
//go:nosplit
func (self class) FromNode(ctx gd.Lifetime, light_node gdclass.Light3D) gdclass.GLTFLight {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(light_node[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFLight.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.GLTFLight
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Converts this GLTFLight instance into a Godot [Light3D] node.
*/
//go:nosplit
func (self class) ToNode(ctx gd.Lifetime) gdclass.Light3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Light3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new GLTFLight instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(ctx gd.Lifetime, dictionary gd.Dictionary) gdclass.GLTFLight {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dictionary))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.GLTFLight.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.GLTFLight
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Serializes this GLTFLight instance into a [Dictionary].
*/
//go:nosplit
func (self class) ToDictionary(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIntensity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_get_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIntensity(intensity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, intensity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_set_intensity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightType(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_get_light_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLightType(light_type gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(light_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_set_light_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRange() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRange(arange gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, arange)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInnerConeAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_get_inner_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInnerConeAngle(inner_cone_angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, inner_cone_angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_set_inner_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOuterConeAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_get_outer_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOuterConeAngle(outer_cone_angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, outer_cone_angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_set_outer_cone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdditionalData(ctx gd.Lifetime, extension_name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_get_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdditionalData(extension_name gd.StringName, additional_data gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension_name))
	callframe.Arg(frame, mmm.Get(additional_data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFLight.Bind_set_additional_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFLight() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFLight() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("GLTFLight", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
