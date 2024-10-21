package FBXState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/GLTFState"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The FBXState handles the state data imported from FBX files.

*/
type Simple [1]classdb.FBXState
func (self Simple) GetAllowGeometryHelperNodes() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAllowGeometryHelperNodes())
}
func (self Simple) SetAllowGeometryHelperNodes(allow bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowGeometryHelperNodes(allow)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.FBXState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) GetAllowGeometryHelperNodes() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FBXState.Bind_get_allow_geometry_helper_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowGeometryHelperNodes(allow bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FBXState.Bind_set_allow_geometry_helper_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsFBXState() Expert { return self[0].AsFBXState() }


//go:nosplit
func (self Simple) AsFBXState() Simple { return self[0].AsFBXState() }


//go:nosplit
func (self class) AsGLTFState() GLTFState.Expert { return self[0].AsGLTFState() }


//go:nosplit
func (self Simple) AsGLTFState() GLTFState.Simple { return self[0].AsGLTFState() }


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
func init() {classdb.Register("FBXState", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
