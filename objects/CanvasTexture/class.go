package CanvasTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Texture2D"
import "grow.graphics/gd/objects/Texture"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Color"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[CanvasTexture] is an alternative to [ImageTexture] for 2D rendering. It allows using normal maps and specular maps in any node that inherits from [CanvasItem]. [CanvasTexture] also allows overriding the texture's filter and repeat mode independently of the node's properties (or the project settings).
[b]Note:[/b] [CanvasTexture] cannot be used in 3D. It will not display correctly when applied to any [VisualInstance3D], such as [Sprite3D] or [Decal]. For physically-based materials in 3D, use [BaseMaterial3D] instead.
*/
type Instance [1]classdb.CanvasTexture

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CanvasTexture

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CanvasTexture"))
	return Instance{classdb.CanvasTexture(object)}
}

func (self Instance) DiffuseTexture() objects.Texture2D {
	return objects.Texture2D(class(self).GetDiffuseTexture())
}

func (self Instance) SetDiffuseTexture(value objects.Texture2D) {
	class(self).SetDiffuseTexture(value)
}

func (self Instance) NormalTexture() objects.Texture2D {
	return objects.Texture2D(class(self).GetNormalTexture())
}

func (self Instance) SetNormalTexture(value objects.Texture2D) {
	class(self).SetNormalTexture(value)
}

func (self Instance) SpecularTexture() objects.Texture2D {
	return objects.Texture2D(class(self).GetSpecularTexture())
}

func (self Instance) SetSpecularTexture(value objects.Texture2D) {
	class(self).SetSpecularTexture(value)
}

func (self Instance) SpecularColor() Color.RGBA {
	return Color.RGBA(class(self).GetSpecularColor())
}

func (self Instance) SetSpecularColor(value Color.RGBA) {
	class(self).SetSpecularColor(gd.Color(value))
}

func (self Instance) SpecularShininess() Float.X {
	return Float.X(Float.X(class(self).GetSpecularShininess()))
}

func (self Instance) SetSpecularShininess(value Float.X) {
	class(self).SetSpecularShininess(gd.Float(value))
}

func (self Instance) TextureFilter() classdb.CanvasItemTextureFilter {
	return classdb.CanvasItemTextureFilter(class(self).GetTextureFilter())
}

func (self Instance) SetTextureFilter(value classdb.CanvasItemTextureFilter) {
	class(self).SetTextureFilter(value)
}

func (self Instance) TextureRepeat() classdb.CanvasItemTextureRepeat {
	return classdb.CanvasItemTextureRepeat(class(self).GetTextureRepeat())
}

func (self Instance) SetTextureRepeat(value classdb.CanvasItemTextureRepeat) {
	class(self).SetTextureRepeat(value)
}

//go:nosplit
func (self class) SetDiffuseTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_diffuse_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDiffuseTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_diffuse_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_normal_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNormalTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_normal_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecularTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_specular_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecularTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_specular_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecularColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_specular_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecularColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_specular_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecularShininess(shininess gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, shininess)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_specular_shininess, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecularShininess() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_specular_shininess, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureFilter(filter classdb.CanvasItemTextureFilter) {
	var frame = callframe.New()
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureFilter() classdb.CanvasItemTextureFilter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CanvasItemTextureFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureRepeat(repeat classdb.CanvasItemTextureRepeat) {
	var frame = callframe.New()
	callframe.Arg(frame, repeat)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRepeat() classdb.CanvasItemTextureRepeat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CanvasItemTextureRepeat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCanvasTexture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCanvasTexture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {
	classdb.Register("CanvasTexture", func(ptr gd.Object) any { return classdb.CanvasTexture(ptr) })
}
