package VisualShaderNodeTexture2DArray

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeSample3D"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Translated to [code]uniform sampler2DArray[/code] in the shader language.

*/
type Simple [1]classdb.VisualShaderNodeTexture2DArray
func (self Simple) SetTextureArray(value [1]classdb.Texture2DArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureArray(value)
}
func (self Simple) GetTextureArray() [1]classdb.Texture2DArray {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2DArray(Expert(self).GetTextureArray(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeTexture2DArray
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTextureArray(value object.Texture2DArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(value[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeTexture2DArray.Bind_set_texture_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureArray(ctx gd.Lifetime) object.Texture2DArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeTexture2DArray.Bind_get_texture_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2DArray
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeTexture2DArray() Expert { return self[0].AsVisualShaderNodeTexture2DArray() }


//go:nosplit
func (self Simple) AsVisualShaderNodeTexture2DArray() Simple { return self[0].AsVisualShaderNodeTexture2DArray() }


//go:nosplit
func (self class) AsVisualShaderNodeSample3D() VisualShaderNodeSample3D.Expert { return self[0].AsVisualShaderNodeSample3D() }


//go:nosplit
func (self Simple) AsVisualShaderNodeSample3D() VisualShaderNodeSample3D.Simple { return self[0].AsVisualShaderNodeSample3D() }


//go:nosplit
func (self class) AsVisualShaderNode() VisualShaderNode.Expert { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self Simple) AsVisualShaderNode() VisualShaderNode.Simple { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VisualShaderNodeTexture2DArray", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
