package BackBufferCopy

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node2D"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Node for back-buffering the currently-displayed screen. The region defined in the [BackBufferCopy] node is buffered with the content of the screen it covers, or the entire screen according to the [member copy_mode]. It can be accessed in shader scripts using the screen texture (i.e. a uniform sampler with [code]hint_screen_texture[/code]).
[b]Note:[/b] Since this node inherits from [Node2D] (and not [Control]), anchors and margins won't apply to child [Control]-derived nodes. This can be problematic when resizing the window. To avoid this, add [Control]-derived nodes as [i]siblings[/i] to the [BackBufferCopy] node instead of adding them as children.

*/
type Simple [1]classdb.BackBufferCopy
func (self Simple) SetRect(rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRect(rect)
}
func (self Simple) GetRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetRect())
}
func (self Simple) SetCopyMode(copy_mode classdb.BackBufferCopyCopyMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCopyMode(copy_mode)
}
func (self Simple) GetCopyMode() classdb.BackBufferCopyCopyMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BackBufferCopyCopyMode(Expert(self).GetCopyMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.BackBufferCopy
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetRect(rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BackBufferCopy.Bind_set_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BackBufferCopy.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCopyMode(copy_mode classdb.BackBufferCopyCopyMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, copy_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BackBufferCopy.Bind_set_copy_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCopyMode() classdb.BackBufferCopyCopyMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BackBufferCopyCopyMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BackBufferCopy.Bind_get_copy_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsBackBufferCopy() Expert { return self[0].AsBackBufferCopy() }


//go:nosplit
func (self Simple) AsBackBufferCopy() Simple { return self[0].AsBackBufferCopy() }


//go:nosplit
func (self class) AsNode2D() Node2D.Expert { return self[0].AsNode2D() }


//go:nosplit
func (self Simple) AsNode2D() Node2D.Simple { return self[0].AsNode2D() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

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
func init() {classdb.Register("BackBufferCopy", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type CopyMode = classdb.BackBufferCopyCopyMode

const (
/*Disables the buffering mode. This means the [BackBufferCopy] node will directly use the portion of screen it covers.*/
	CopyModeDisabled CopyMode = 0
/*[BackBufferCopy] buffers a rectangular region.*/
	CopyModeRect CopyMode = 1
/*[BackBufferCopy] buffers the entire screen.*/
	CopyModeViewport CopyMode = 2
)
