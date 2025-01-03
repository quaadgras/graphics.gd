package EncodedObjectAsID

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Utility class which holds a reference to the internal identifier of an [Object] instance, as given by [method Object.get_instance_id]. This ID can then be used to retrieve the object instance with [method @GlobalScope.instance_from_id].
This class is used internally by the editor inspector and script debugger, but can also be used in plugins to pass and display objects as their IDs.
*/
type Instance [1]classdb.EncodedObjectAsID
type Any interface {
	gd.IsClass
	AsEncodedObjectAsID() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EncodedObjectAsID

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EncodedObjectAsID"))
	return Instance{*(*classdb.EncodedObjectAsID)(unsafe.Pointer(&object))}
}

func (self Instance) ObjectId() int {
	return int(int(class(self).GetObjectId()))
}

func (self Instance) SetObjectId(value int) {
	class(self).SetObjectId(gd.Int(value))
}

//go:nosplit
func (self class) SetObjectId(id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EncodedObjectAsID.Bind_set_object_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetObjectId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EncodedObjectAsID.Bind_get_object_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsEncodedObjectAsID() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEncodedObjectAsID() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted      { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted   { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EncodedObjectAsID", func(ptr gd.Object) any {
		return [1]classdb.EncodedObjectAsID{*(*classdb.EncodedObjectAsID)(unsafe.Pointer(&ptr))}
	})
}
