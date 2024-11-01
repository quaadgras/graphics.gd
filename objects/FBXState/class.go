package FBXState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/GLTFState"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
The FBXState handles the state data imported from FBX files.
*/
type Instance [1]classdb.FBXState

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.FBXState

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FBXState"))
	return Instance{classdb.FBXState(object)}
}

func (self Instance) AllowGeometryHelperNodes() bool {
	return bool(class(self).GetAllowGeometryHelperNodes())
}

func (self Instance) SetAllowGeometryHelperNodes(value bool) {
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
func (self class) SetAllowGeometryHelperNodes(allow bool) {
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FBXState.Bind_set_allow_geometry_helper_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsFBXState() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFBXState() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGLTFState() GLTFState.Advanced {
	return *((*GLTFState.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGLTFState() GLTFState.Instance {
	return *((*GLTFState.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsGLTFState(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsGLTFState(), name)
	}
}
func init() { classdb.Register("FBXState", func(ptr gd.Object) any { return classdb.FBXState(ptr) }) }
