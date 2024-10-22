package TextServerDummy

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/TextServerExtension"
import "grow.graphics/gd/gdclass/TextServer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A dummy [TextServer] interface that doesn't do anything. Useful for freeing up memory when rendering text is not needed, as text servers are resource-intensive. It can also be used for performance comparisons in complex GUIs to check the impact of text rendering.
A dummy text server is always available at the start of a project. Here's how to access it:
[codeblock]
var dummy_text_server = TextServerManager.find_interface("Dummy")
if dummy_text_server != null:
    TextServerManager.set_primary_interface(dummy_text_server)
    # If the other text servers are unneeded, they can be removed:
    for i in TextServerManager.get_interface_count():
        var text_server = TextServerManager.get_interface(i)
        if text_server != dummy_text_server:
            TextServerManager.remove_interface(text_server)
[/codeblock]
The command line argument [code]--text-driver Dummy[/code] (case-sensitive) can be used to force the "Dummy" [TextServer] on any project.

*/
type Go [1]classdb.TextServerDummy
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextServerDummy
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TextServerDummy"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsTextServerDummy() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextServerDummy() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTextServerExtension() TextServerExtension.GD { return *((*TextServerExtension.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextServerExtension() TextServerExtension.Go { return *((*TextServerExtension.Go)(unsafe.Pointer(&self))) }
func (self class) AsTextServer() TextServer.GD { return *((*TextServer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextServer() TextServer.Go { return *((*TextServer.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTextServerExtension(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTextServerExtension(), name)
	}
}
func init() {classdb.Register("TextServerDummy", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
