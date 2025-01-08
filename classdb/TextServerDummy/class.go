// Package TextServerDummy provides methods for working with TextServerDummy object instances.
package TextServerDummy

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/TextServerExtension"
import "graphics.gd/classdb/TextServer"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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
type Instance [1]gdclass.TextServerDummy

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTextServerDummy() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TextServerDummy

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextServerDummy"))
	casted := Instance{*(*gdclass.TextServerDummy)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(TextServerExtension.Advanced(self.AsTextServerExtension()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(TextServerExtension.Instance(self.AsTextServerExtension()), name)
	}
}
func init() {
	gdclass.Register("TextServerDummy", func(ptr gd.Object) any {
		return [1]gdclass.TextServerDummy{*(*gdclass.TextServerDummy)(unsafe.Pointer(&ptr))}
	})
}
