// Package FogMaterial provides methods for working with FogMaterial object instances.
package FogMaterial

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
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Material"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Color"

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

/*
A [Material] resource that can be used by [FogVolume]s to draw volumetric effects.
If you need more advanced effects, use a custom [url=$DOCS_URL/tutorials/shaders/shader_reference/fog_shader.html]fog shader[/url].
*/
type Instance [1]gdclass.FogMaterial

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFogMaterial() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FogMaterial

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FogMaterial"))
	casted := Instance{*(*gdclass.FogMaterial)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Density() Float.X {
	return Float.X(Float.X(class(self).GetDensity()))
}

func (self Instance) SetDensity(value Float.X) {
	class(self).SetDensity(gd.Float(value))
}

func (self Instance) Albedo() Color.RGBA {
	return Color.RGBA(class(self).GetAlbedo())
}

func (self Instance) SetAlbedo(value Color.RGBA) {
	class(self).SetAlbedo(gd.Color(value))
}

func (self Instance) Emission() Color.RGBA {
	return Color.RGBA(class(self).GetEmission())
}

func (self Instance) SetEmission(value Color.RGBA) {
	class(self).SetEmission(gd.Color(value))
}

func (self Instance) HeightFalloff() Float.X {
	return Float.X(Float.X(class(self).GetHeightFalloff()))
}

func (self Instance) SetHeightFalloff(value Float.X) {
	class(self).SetHeightFalloff(gd.Float(value))
}

func (self Instance) EdgeFade() Float.X {
	return Float.X(Float.X(class(self).GetEdgeFade()))
}

func (self Instance) SetEdgeFade(value Float.X) {
	class(self).SetEdgeFade(gd.Float(value))
}

func (self Instance) DensityTexture() [1]gdclass.Texture3D {
	return [1]gdclass.Texture3D(class(self).GetDensityTexture())
}

func (self Instance) SetDensityTexture(value [1]gdclass.Texture3D) {
	class(self).SetDensityTexture(value)
}

//go:nosplit
func (self class) SetDensity(density gd.Float) { //gd:FogMaterial.set_density
	var frame = callframe.New()
	callframe.Arg(frame, density)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_density, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDensity() gd.Float { //gd:FogMaterial.get_density
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_density, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlbedo(albedo gd.Color) { //gd:FogMaterial.set_albedo
	var frame = callframe.New()
	callframe.Arg(frame, albedo)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_albedo, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlbedo() gd.Color { //gd:FogMaterial.get_albedo
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_albedo, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmission(emission gd.Color) { //gd:FogMaterial.set_emission
	var frame = callframe.New()
	callframe.Arg(frame, emission)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_emission, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmission() gd.Color { //gd:FogMaterial.get_emission
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_emission, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeightFalloff(height_falloff gd.Float) { //gd:FogMaterial.set_height_falloff
	var frame = callframe.New()
	callframe.Arg(frame, height_falloff)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_height_falloff, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeightFalloff() gd.Float { //gd:FogMaterial.get_height_falloff
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_height_falloff, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEdgeFade(edge_fade gd.Float) { //gd:FogMaterial.set_edge_fade
	var frame = callframe.New()
	callframe.Arg(frame, edge_fade)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_edge_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEdgeFade() gd.Float { //gd:FogMaterial.get_edge_fade
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_edge_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDensityTexture(density_texture [1]gdclass.Texture3D) { //gd:FogMaterial.set_density_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(density_texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_set_density_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDensityTexture() [1]gdclass.Texture3D { //gd:FogMaterial.get_density_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FogMaterial.Bind_get_density_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture3D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsFogMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFogMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.Advanced {
	return *((*Material.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMaterial() Material.Instance {
	return *((*Material.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Material.Advanced(self.AsMaterial()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Material.Instance(self.AsMaterial()), name)
	}
}
func init() {
	gdclass.Register("FogMaterial", func(ptr gd.Object) any { return [1]gdclass.FogMaterial{*(*gdclass.FogMaterial)(unsafe.Pointer(&ptr))} })
}
