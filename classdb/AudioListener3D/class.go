// Package AudioListener3D provides methods for working with AudioListener3D object instances.
package AudioListener3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Transform3D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
Once added to the scene tree and enabled using [method make_current], this node will override the location sounds are heard from. This can be used to listen from a location different from the [Camera3D].
*/
type Instance [1]gdclass.AudioListener3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioListener3D() Instance
}

/*
Enables the listener. This will override the current camera's listener.
*/
func (self Instance) MakeCurrent() {
	class(self).MakeCurrent()
}

/*
Disables the listener to use the current camera's listener instead.
*/
func (self Instance) ClearCurrent() {
	class(self).ClearCurrent()
}

/*
Returns [code]true[/code] if the listener was made current using [method make_current], [code]false[/code] otherwise.
[b]Note:[/b] There may be more than one AudioListener3D marked as "current" in the scene tree, but only the one that was made current last will be used.
*/
func (self Instance) IsCurrent() bool {
	return bool(class(self).IsCurrent())
}

/*
Returns the listener's global orthonormalized [Transform3D].
*/
func (self Instance) GetListenerTransform() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetListenerTransform())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioListener3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioListener3D"))
	casted := Instance{*(*gdclass.AudioListener3D)(unsafe.Pointer(&object))}
	return casted
}

/*
Enables the listener. This will override the current camera's listener.
*/
//go:nosplit
func (self class) MakeCurrent() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioListener3D.Bind_make_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disables the listener to use the current camera's listener instead.
*/
//go:nosplit
func (self class) ClearCurrent() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioListener3D.Bind_clear_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the listener was made current using [method make_current], [code]false[/code] otherwise.
[b]Note:[/b] There may be more than one AudioListener3D marked as "current" in the scene tree, but only the one that was made current last will be used.
*/
//go:nosplit
func (self class) IsCurrent() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioListener3D.Bind_is_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the listener's global orthonormalized [Transform3D].
*/
//go:nosplit
func (self class) GetListenerTransform() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioListener3D.Bind_get_listener_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioListener3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioListener3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced      { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance   { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced          { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance       { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("AudioListener3D", func(ptr gd.Object) any {
		return [1]gdclass.AudioListener3D{*(*gdclass.AudioListener3D)(unsafe.Pointer(&ptr))}
	})
}
