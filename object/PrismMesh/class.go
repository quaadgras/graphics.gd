package PrismMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PrimitiveMesh"
import "grow.graphics/gd/object/Mesh"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Class representing a prism-shaped [PrimitiveMesh].

*/
type Simple [1]classdb.PrismMesh
func (self Simple) SetLeftToRight(left_to_right float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLeftToRight(gd.Float(left_to_right))
}
func (self Simple) GetLeftToRight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetLeftToRight()))
}
func (self Simple) SetSize(size gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSize(size)
}
func (self Simple) GetSize() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetSize())
}
func (self Simple) SetSubdivideWidth(segments int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubdivideWidth(gd.Int(segments))
}
func (self Simple) GetSubdivideWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSubdivideWidth()))
}
func (self Simple) SetSubdivideHeight(segments int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubdivideHeight(gd.Int(segments))
}
func (self Simple) GetSubdivideHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSubdivideHeight()))
}
func (self Simple) SetSubdivideDepth(segments int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubdivideDepth(gd.Int(segments))
}
func (self Simple) GetSubdivideDepth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSubdivideDepth()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PrismMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetLeftToRight(left_to_right gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, left_to_right)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_set_left_to_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLeftToRight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_get_left_to_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSize(size gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubdivideWidth(segments gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_set_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubdivideWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_get_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubdivideHeight(segments gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_set_subdivide_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubdivideHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_get_subdivide_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubdivideDepth(segments gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_set_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubdivideDepth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PrismMesh.Bind_get_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPrismMesh() Expert { return self[0].AsPrismMesh() }


//go:nosplit
func (self Simple) AsPrismMesh() Simple { return self[0].AsPrismMesh() }


//go:nosplit
func (self class) AsPrimitiveMesh() PrimitiveMesh.Expert { return self[0].AsPrimitiveMesh() }


//go:nosplit
func (self Simple) AsPrimitiveMesh() PrimitiveMesh.Simple { return self[0].AsPrimitiveMesh() }


//go:nosplit
func (self class) AsMesh() Mesh.Expert { return self[0].AsMesh() }


//go:nosplit
func (self Simple) AsMesh() Mesh.Simple { return self[0].AsMesh() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PrismMesh", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
