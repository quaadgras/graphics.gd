package ScriptEditorBase

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VBoxContainer"
import "grow.graphics/gd/object/BoxContainer"
import "grow.graphics/gd/object/Container"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base editor for editing scripts in the [ScriptEditor]. This does not include documentation items.

*/
type Simple [1]classdb.ScriptEditorBase
func (self Simple) GetBaseEditor() [1]classdb.Control {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Control(Expert(self).GetBaseEditor(gc))
}
func (self Simple) AddSyntaxHighlighter(highlighter [1]classdb.EditorSyntaxHighlighter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddSyntaxHighlighter(highlighter)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ScriptEditorBase
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the underlying [Control] used for editing scripts. For text scripts, this is a [CodeEdit].
*/
//go:nosplit
func (self class) GetBaseEditor(ctx gd.Lifetime) object.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditorBase.Bind_get_base_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Adds a [EditorSyntaxHighlighter] to the open script.
*/
//go:nosplit
func (self class) AddSyntaxHighlighter(highlighter object.EditorSyntaxHighlighter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(highlighter[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditorBase.Bind_add_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsScriptEditorBase() Expert { return self[0].AsScriptEditorBase() }


//go:nosplit
func (self Simple) AsScriptEditorBase() Simple { return self[0].AsScriptEditorBase() }


//go:nosplit
func (self class) AsVBoxContainer() VBoxContainer.Expert { return self[0].AsVBoxContainer() }


//go:nosplit
func (self Simple) AsVBoxContainer() VBoxContainer.Simple { return self[0].AsVBoxContainer() }


//go:nosplit
func (self class) AsBoxContainer() BoxContainer.Expert { return self[0].AsBoxContainer() }


//go:nosplit
func (self Simple) AsBoxContainer() BoxContainer.Simple { return self[0].AsBoxContainer() }


//go:nosplit
func (self class) AsContainer() Container.Expert { return self[0].AsContainer() }


//go:nosplit
func (self Simple) AsContainer() Container.Simple { return self[0].AsContainer() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


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
func init() {classdb.Register("ScriptEditorBase", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
