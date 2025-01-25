// Package JavaClassWrapper provides methods for working with JavaClassWrapper object instances.
package JavaClassWrapper

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
The JavaClassWrapper singleton provides a way for the Godot application to send and receive data through the [url=https://developer.android.com/training/articles/perf-jni]Java Native Interface[/url] (JNI).
[b]Note:[/b] This singleton is only available in Android builds.
*/
var self [1]gdclass.JavaClassWrapper
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.JavaClassWrapper)
	self = *(*[1]gdclass.JavaClassWrapper)(unsafe.Pointer(&obj))
}

/*
Wraps a class defined in Java, and returns it as a [JavaClass] [Object] type that Godot can interact with.
[b]Note:[/b] This method only works on Android. On every other platform, this method does nothing and returns an empty [JavaClass].
*/
func Wrap(name string) [1]gdclass.JavaClass {
	once.Do(singleton)
	return [1]gdclass.JavaClass(class(self).Wrap(gd.NewString(name)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.JavaClassWrapper

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Wraps a class defined in Java, and returns it as a [JavaClass] [Object] type that Godot can interact with.
[b]Note:[/b] This method only works on Android. On every other platform, this method does nothing and returns an empty [JavaClass].
*/
//go:nosplit
func (self class) Wrap(name gd.String) [1]gdclass.JavaClass {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaClassWrapper.Bind_wrap, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.JavaClass{gd.PointerWithOwnershipTransferredToGo[gdclass.JavaClass](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("JavaClassWrapper", func(ptr gd.Object) any {
		return [1]gdclass.JavaClassWrapper{*(*gdclass.JavaClassWrapper)(unsafe.Pointer(&ptr))}
	})
}
