package FBXState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GLTFState"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The FBXState handles the state data imported from FBX files.

*/
type Go [1]classdb.FBXState
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.FBXState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FBXState"))
	return Go{classdb.FBXState(object)}
}

func (self Go) AllowGeometryHelperNodes() bool {
		return bool(class(self).GetAllowGeometryHelperNodes())
}

func (self Go) SetAllowGeometryHelperNodes(value bool) {
	class(self).SetAllowGeometryHelperNodes(value)
}

//go:nosplit
func (self class) GetAllowGeometryHelperNodes() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FBXState.Bind_get_allow_geometry_helper_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowGeometryHelperNodes(allow bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FBXState.Bind_set_allow_geometry_helper_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsFBXState() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsFBXState() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGLTFState() GLTFState.GD { return *((*GLTFState.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFState() GLTFState.Go { return *((*GLTFState.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGLTFState(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGLTFState(), name)
	}
}
func init() {classdb.Register("FBXState", func(ptr gd.Object) any { return classdb.FBXState(ptr) })}
