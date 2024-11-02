package TextServerDummy

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/TextServerExtension"
import "grow.graphics/gd/objects/TextServer"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.TextServerDummy

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TextServerDummy

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextServerDummy"))
	return Instance{classdb.TextServerDummy(object)}
}

func (self class) AsTextServerDummy() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextServerDummy() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTextServerExtension() TextServerExtension.Advanced {
	return *((*TextServerExtension.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextServerExtension() TextServerExtension.Instance {
	return *((*TextServerExtension.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTextServer() TextServer.Advanced {
	return *((*TextServer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextServer() TextServer.Instance {
	return *((*TextServer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTextServerExtension(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTextServerExtension(), name)
	}
}
func init() {
	classdb.Register("TextServerDummy", func(ptr gd.Object) any { return classdb.TextServerDummy(ptr) })
}
