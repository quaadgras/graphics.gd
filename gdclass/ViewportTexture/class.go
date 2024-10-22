package ViewportTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture2D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A [ViewportTexture] provides the content of a [Viewport] as a dynamic [Texture2D]. This can be used to combine the rendering of [Control], [Node2D] and [Node3D] nodes. For example, you can use this texture to display a 3D scene inside a [TextureRect], or a 2D overlay in a [Sprite3D].
To get a [ViewportTexture] in code, use the [method Viewport.get_texture] method on the target viewport.
[b]Note:[/b] A [ViewportTexture] is always local to its scene (see [member Resource.resource_local_to_scene]). If the scene root is not ready, it may return incorrect data (see [signal Node.ready]).
[b]Note:[/b] Instantiating scenes containing a high-resolution [ViewportTexture] may cause noticeable stutter.

*/
type Go [1]classdb.ViewportTexture
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ViewportTexture
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ViewportTexture"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ViewportPath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetViewportPathInScene(gc).String())
}

func (self Go) SetViewportPath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetViewportPathInScene(gc.String(value).NodePath(gc))
}

//go:nosplit
func (self class) SetViewportPathInScene(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ViewportTexture.Bind_set_viewport_path_in_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetViewportPathInScene(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ViewportTexture.Bind_get_viewport_path_in_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsViewportTexture() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsViewportTexture() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.GD { return *((*Texture2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture2D() Texture2D.Go { return *((*Texture2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {classdb.Register("ViewportTexture", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
