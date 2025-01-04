// Package JavaClass provides methods for working with JavaClass object instances.
package JavaClass

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Represents an object from the Java Native Interface. It is returned from [method JavaClassWrapper.wrap].
[b]Note:[/b] This class only works on Android. For any other build, this class does nothing.
[b]Note:[/b] This class is not to be confused with [JavaScriptObject].
*/
type Instance [1]gdclass.JavaClass
type Any interface {
	gd.IsClass
	AsJavaClass() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.JavaClass

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("JavaClass"))
	return Instance{*(*gdclass.JavaClass)(unsafe.Pointer(&object))}
}

func (self class) AsJavaClass() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJavaClass() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

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
	gdclass.Register("JavaClass", func(ptr gd.Object) any { return [1]gdclass.JavaClass{*(*gdclass.JavaClass)(unsafe.Pointer(&ptr))} })
}
