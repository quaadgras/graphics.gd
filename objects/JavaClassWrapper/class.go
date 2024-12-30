package JavaClassWrapper

import "unsafe"
import "sync"
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
var _ = pointers.Root

/*
The JavaClassWrapper singleton provides a way for the Godot application to send and receive data through the [url=https://developer.android.com/training/articles/perf-jni]Java Native Interface[/url] (JNI).
[b]Note:[/b] This singleton is only available in Android builds.
*/
var self objects.JavaClassWrapper
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.JavaClassWrapper)
	self = *(*objects.JavaClassWrapper)(unsafe.Pointer(&obj))
}

/*
Wraps a class defined in Java, and returns it as a [JavaClass] [Object] type that Godot can interact with.
[b]Note:[/b] This method only works on Android. On every other platform, this method does nothing and returns an empty [JavaClass].
*/
func Wrap(name string) objects.JavaClass {
	once.Do(singleton)
	return objects.JavaClass(class(self).Wrap(gd.NewString(name)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.JavaClassWrapper

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Wraps a class defined in Java, and returns it as a [JavaClass] [Object] type that Godot can interact with.
[b]Note:[/b] This method only works on Android. On every other platform, this method does nothing and returns an empty [JavaClass].
*/
//go:nosplit
func (self class) Wrap(name gd.String) objects.JavaClass {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaClassWrapper.Bind_wrap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.JavaClass{classdb.JavaClass(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("JavaClassWrapper", func(ptr gd.Object) any { return classdb.JavaClassWrapper(ptr) })
}
