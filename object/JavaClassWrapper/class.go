package JavaClassWrapper

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The JavaClassWrapper singleton provides a way for the Godot application to send and receive data through the [url=https://developer.android.com/training/articles/perf-jni]Java Native Interface[/url] (JNI).
[b]Note:[/b] This singleton is only available in Android builds.

*/
type Simple [1]classdb.JavaClassWrapper
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.JavaClassWrapper
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Wraps a class defined in Java, and returns it as a [JavaClass] [Object] type that Godot can interact with.
[b]Note:[/b] This method only works on Android. On every other platform, this method does nothing and returns an empty [JavaClass].
*/


//go:nosplit
func (self class) AsJavaClassWrapper() Expert { return self[0].AsJavaClassWrapper() }


//go:nosplit
func (self Simple) AsJavaClassWrapper() Simple { return self[0].AsJavaClassWrapper() }

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
func init() {classdb.Register("JavaClassWrapper", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
