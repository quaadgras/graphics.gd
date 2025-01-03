package ViewportTexture

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Texture2D"
import "graphics.gd/objects/Texture"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Path"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A [ViewportTexture] provides the content of a [Viewport] as a dynamic [Texture2D]. This can be used to combine the rendering of [Control], [Node2D] and [Node3D] nodes. For example, you can use this texture to display a 3D scene inside a [TextureRect], or a 2D overlay in a [Sprite3D].
To get a [ViewportTexture] in code, use the [method Viewport.get_texture] method on the target viewport.
[b]Note:[/b] A [ViewportTexture] is always local to its scene (see [member Resource.resource_local_to_scene]). If the scene root is not ready, it may return incorrect data (see [signal Node.ready]).
[b]Note:[/b] Instantiating scenes containing a high-resolution [ViewportTexture] may cause noticeable stutter.
*/
type Instance [1]classdb.ViewportTexture
type Any interface {
	gd.IsClass
	AsViewportTexture() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ViewportTexture

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ViewportTexture"))
	return Instance{*(*classdb.ViewportTexture)(unsafe.Pointer(&object))}
}

func (self Instance) ViewportPath() Path.String {
	return Path.String(class(self).GetViewportPathInScene().String())
}

func (self Instance) SetViewportPath(value Path.String) {
	class(self).SetViewportPathInScene(gd.NewString(string(value)).NodePath())
}

//go:nosplit
func (self class) SetViewportPathInScene(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ViewportTexture.Bind_set_viewport_path_in_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetViewportPathInScene() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ViewportTexture.Bind_get_viewport_path_in_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsViewportTexture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsViewportTexture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {
	classdb.Register("ViewportTexture", func(ptr gd.Object) any {
		return [1]classdb.ViewportTexture{*(*classdb.ViewportTexture)(unsafe.Pointer(&ptr))}
	})
}
