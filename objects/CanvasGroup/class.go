package CanvasGroup

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Child [CanvasItem] nodes of a [CanvasGroup] are drawn as a single object. It allows to e.g. draw overlapping translucent 2D nodes without blending (set [member CanvasItem.self_modulate] property of [CanvasGroup] to achieve this effect).
[b]Note:[/b] The [CanvasGroup] uses a custom shader to read from the backbuffer to draw its children. Assigning a [Material] to the [CanvasGroup] overrides the builtin shader. To duplicate the behavior of the builtin shader in a custom [Shader] use the following:
[codeblock]
shader_type canvas_item;
render_mode unshaded;

uniform sampler2D screen_texture : hint_screen_texture, repeat_disable, filter_nearest;

	void fragment() {
	    vec4 c = textureLod(screen_texture, SCREEN_UV, 0.0);

	    if (c.a > 0.0001) {
	        c.rgb /= c.a;
	    }

	    COLOR *= c;
	}

[/codeblock]
[b]Note:[/b] Since [CanvasGroup] and [member CanvasItem.clip_children] both utilize the backbuffer, children of a [CanvasGroup] who have their [member CanvasItem.clip_children] set to anything other than [constant CanvasItem.CLIP_CHILDREN_DISABLED] will not function correctly.
*/
type Instance [1]classdb.CanvasGroup

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CanvasGroup

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CanvasGroup"))
	return Instance{classdb.CanvasGroup(object)}
}

func (self Instance) FitMargin() float64 {
	return float64(float64(class(self).GetFitMargin()))
}

func (self Instance) SetFitMargin(value float64) {
	class(self).SetFitMargin(gd.Float(value))
}

func (self Instance) ClearMargin() float64 {
	return float64(float64(class(self).GetClearMargin()))
}

func (self Instance) SetClearMargin(value float64) {
	class(self).SetClearMargin(gd.Float(value))
}

func (self Instance) UseMipmaps() bool {
	return bool(class(self).IsUsingMipmaps())
}

func (self Instance) SetUseMipmaps(value bool) {
	class(self).SetUseMipmaps(value)
}

//go:nosplit
func (self class) SetFitMargin(fit_margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, fit_margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasGroup.Bind_set_fit_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFitMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasGroup.Bind_get_fit_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClearMargin(clear_margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, clear_margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasGroup.Bind_set_clear_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetClearMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasGroup.Bind_get_clear_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseMipmaps(use_mipmaps bool) {
	var frame = callframe.New()
	callframe.Arg(frame, use_mipmaps)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasGroup.Bind_set_use_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingMipmaps() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasGroup.Bind_is_using_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCanvasGroup() Advanced      { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCanvasGroup() Instance   { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("CanvasGroup", func(ptr gd.Object) any { return classdb.CanvasGroup(ptr) })
}
