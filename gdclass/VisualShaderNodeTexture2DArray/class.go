package VisualShaderNodeTexture2DArray

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeSample3D"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Translated to [code]uniform sampler2DArray[/code] in the shader language.

*/
type Go [1]classdb.VisualShaderNodeTexture2DArray
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeTexture2DArray
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeTexture2DArray"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) TextureArray() gdclass.Texture2DArray {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2DArray(class(self).GetTextureArray(gc))
}

func (self Go) SetTextureArray(value gdclass.Texture2DArray) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureArray(value)
}

//go:nosplit
func (self class) SetTextureArray(value gdclass.Texture2DArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(value[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeTexture2DArray.Bind_set_texture_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureArray(ctx gd.Lifetime) gdclass.Texture2DArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeTexture2DArray.Bind_get_texture_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2DArray
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTexture2DArray() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeTexture2DArray() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeSample3D() VisualShaderNodeSample3D.GD { return *((*VisualShaderNodeSample3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeSample3D() VisualShaderNodeSample3D.Go { return *((*VisualShaderNodeSample3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeSample3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeSample3D(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeTexture2DArray", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
