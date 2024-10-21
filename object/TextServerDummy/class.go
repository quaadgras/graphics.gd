package TextServerDummy

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/TextServerExtension"
import "grow.graphics/gd/object/TextServer"

var _ unsafe.Pointer
var _ object.Engine
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
type Simple [1]classdb.TextServerDummy
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TextServerDummy
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsTextServerDummy() Expert { return self[0].AsTextServerDummy() }


//go:nosplit
func (self Simple) AsTextServerDummy() Simple { return self[0].AsTextServerDummy() }


//go:nosplit
func (self class) AsTextServerExtension() TextServerExtension.Expert { return self[0].AsTextServerExtension() }


//go:nosplit
func (self Simple) AsTextServerExtension() TextServerExtension.Simple { return self[0].AsTextServerExtension() }


//go:nosplit
func (self class) AsTextServer() TextServer.Expert { return self[0].AsTextServer() }


//go:nosplit
func (self Simple) AsTextServer() TextServer.Simple { return self[0].AsTextServer() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TextServerDummy", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
