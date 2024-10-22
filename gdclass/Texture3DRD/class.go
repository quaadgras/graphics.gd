package Texture3DRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture3D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This texture class allows you to use a 3D texture created directly on the [RenderingDevice] as a texture for materials, meshes, etc.

*/
type Go [1]classdb.Texture3DRD
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Texture3DRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Texture3DRD"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) TextureRdRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
		return gd.RID(class(self).GetTextureRdRid())
}

func (self Go) SetTextureRdRid(value gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureRdRid(value)
}

//go:nosplit
func (self class) SetTextureRdRid(texture_rd_rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture_rd_rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3DRD.Bind_set_texture_rd_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureRdRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3DRD.Bind_get_texture_rd_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTexture3DRD() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture3DRD() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture3D() Texture3D.GD { return *((*Texture3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture3D() Texture3D.Go { return *((*Texture3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture3D(), name)
	}
}
func init() {classdb.Register("Texture3DRD", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
