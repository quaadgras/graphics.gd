package CSGPrimitive3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/CSGShape3D"
import "grow.graphics/gd/object/GeometryInstance3D"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Parent class for various CSG primitives. It contains code and functionality that is common between them. It cannot be used directly. Instead use one of the various classes that inherit from it.
[b]Note:[/b] CSG nodes are intended to be used for level prototyping. Creating CSG nodes has a significant CPU cost compared to creating a [MeshInstance3D] with a [PrimitiveMesh]. Moving a CSG node within another CSG node also has a significant CPU cost, so it should be avoided during gameplay.

*/
type Simple [1]classdb.CSGPrimitive3D
func (self Simple) SetFlipFaces(flip_faces bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlipFaces(flip_faces)
}
func (self Simple) GetFlipFaces() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlipFaces())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CSGPrimitive3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetFlipFaces(flip_faces bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flip_faces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPrimitive3D.Bind_set_flip_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlipFaces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGPrimitive3D.Bind_get_flip_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCSGPrimitive3D() Expert { return self[0].AsCSGPrimitive3D() }


//go:nosplit
func (self Simple) AsCSGPrimitive3D() Simple { return self[0].AsCSGPrimitive3D() }


//go:nosplit
func (self class) AsCSGShape3D() CSGShape3D.Expert { return self[0].AsCSGShape3D() }


//go:nosplit
func (self Simple) AsCSGShape3D() CSGShape3D.Simple { return self[0].AsCSGShape3D() }


//go:nosplit
func (self class) AsGeometryInstance3D() GeometryInstance3D.Expert { return self[0].AsGeometryInstance3D() }


//go:nosplit
func (self Simple) AsGeometryInstance3D() GeometryInstance3D.Simple { return self[0].AsGeometryInstance3D() }


//go:nosplit
func (self class) AsVisualInstance3D() VisualInstance3D.Expert { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self Simple) AsVisualInstance3D() VisualInstance3D.Simple { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

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
func init() {classdb.Register("CSGPrimitive3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
