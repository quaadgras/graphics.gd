// Package FBXState provides methods for working with FBXState object instances.
package FBXState

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/GLTFState"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The FBXState handles the state data imported from FBX files.
*/
type Instance [1]gdclass.FBXState
type Any interface {
	gd.IsClass
	AsFBXState() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FBXState

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FBXState"))
	casted := Instance{*(*gdclass.FBXState)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GLTFState.Advanced(self.AsGLTFState()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GLTFState.Instance(self.AsGLTFState()), name)
	}
}
func init() {
	gdclass.Register("FBXState", func(ptr gd.Object) any { return [1]gdclass.FBXState{*(*gdclass.FBXState)(unsafe.Pointer(&ptr))} })
}
