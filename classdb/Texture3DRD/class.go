// Package Texture3DRD provides methods for working with Texture3DRD object instances.
package Texture3DRD

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Texture3D"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This texture class allows you to use a 3D texture created directly on the [RenderingDevice] as a texture for materials, meshes, etc.
*/
type Instance [1]gdclass.Texture3DRD

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTexture3DRD() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Texture3DRD

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Texture3DRD"))
	casted := Instance{*(*gdclass.Texture3DRD)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) TextureRdRid() Resource.ID {
	return Resource.ID(class(self).GetTextureRdRid())
}

func (self Instance) SetTextureRdRid(value Resource.ID) {
	class(self).SetTextureRdRid(value)
}

//go:nosplit
func (self class) SetTextureRdRid(texture_rd_rid gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, texture_rd_rid)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3DRD.Bind_set_texture_rd_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRdRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3DRD.Bind_get_texture_rd_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTexture3DRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture3DRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture3D() Texture3D.Advanced {
	return *((*Texture3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture3D() Texture3D.Instance {
	return *((*Texture3D.Instance)(unsafe.Pointer(&self)))
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture3D.Advanced(self.AsTexture3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture3D.Instance(self.AsTexture3D()), name)
	}
}
func init() {
	gdclass.Register("Texture3DRD", func(ptr gd.Object) any { return [1]gdclass.Texture3DRD{*(*gdclass.Texture3DRD)(unsafe.Pointer(&ptr))} })
}
