package Texture3DRD

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Texture3D"
import "graphics.gd/objects/Texture"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This texture class allows you to use a 3D texture created directly on the [RenderingDevice] as a texture for materials, meshes, etc.
*/
type Instance [1]classdb.Texture3DRD
type Any interface {
	gd.IsClass
	AsTexture3DRD() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Texture3DRD

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Texture3DRD"))
	return Instance{classdb.Texture3DRD(object)}
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
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3DRD.Bind_set_texture_rd_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRdRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Texture3DRD.Bind_get_texture_rd_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture3D(), name)
	}
}
func init() {
	classdb.Register("Texture3DRD", func(ptr gd.Object) any { return classdb.Texture3DRD(ptr) })
}
