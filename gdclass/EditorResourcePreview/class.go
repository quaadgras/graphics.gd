package EditorResourcePreview

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This node is used to generate previews for resources or files.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_resource_previewer].

*/
type Go [1]classdb.EditorResourcePreview

/*
Queue a resource file located at [param path] for preview. Once the preview is ready, the [param receiver]'s [param receiver_func] will be called. The [param receiver_func] must take the following four arguments: [String] path, [Texture2D] preview, [Texture2D] thumbnail_preview, [Variant] userdata. [param userdata] can be anything, and will be returned when [param receiver_func] is called.
[b]Note:[/b] If it was not possible to create the preview the [param receiver_func] will still be called, but the preview will be null.
*/
func (self Go) QueueResourcePreview(path string, receiver gd.Object, receiver_func string, userdata gd.Variant) {
	class(self).QueueResourcePreview(gd.NewString(path), receiver, gd.NewStringName(receiver_func), userdata)
}

/*
Queue the [param resource] being edited for preview. Once the preview is ready, the [param receiver]'s [param receiver_func] will be called. The [param receiver_func] must take the following four arguments: [String] path, [Texture2D] preview, [Texture2D] thumbnail_preview, [Variant] userdata. [param userdata] can be anything, and will be returned when [param receiver_func] is called.
[b]Note:[/b] If it was not possible to create the preview the [param receiver_func] will still be called, but the preview will be null.
*/
func (self Go) QueueEditedResourcePreview(resource gdclass.Resource, receiver gd.Object, receiver_func string, userdata gd.Variant) {
	class(self).QueueEditedResourcePreview(resource, receiver, gd.NewStringName(receiver_func), userdata)
}

/*
Create an own, custom preview generator.
*/
func (self Go) AddPreviewGenerator(generator gdclass.EditorResourcePreviewGenerator) {
	class(self).AddPreviewGenerator(generator)
}

/*
Removes a custom preview generator.
*/
func (self Go) RemovePreviewGenerator(generator gdclass.EditorResourcePreviewGenerator) {
	class(self).RemovePreviewGenerator(generator)
}

/*
Check if the resource changed, if so, it will be invalidated and the corresponding signal emitted.
*/
func (self Go) CheckForInvalidation(path string) {
	class(self).CheckForInvalidation(gd.NewString(path))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorResourcePreview
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorResourcePreview"))
	return Go{classdb.EditorResourcePreview(object)}
}

/*
Queue a resource file located at [param path] for preview. Once the preview is ready, the [param receiver]'s [param receiver_func] will be called. The [param receiver_func] must take the following four arguments: [String] path, [Texture2D] preview, [Texture2D] thumbnail_preview, [Variant] userdata. [param userdata] can be anything, and will be returned when [param receiver_func] is called.
[b]Note:[/b] If it was not possible to create the preview the [param receiver_func] will still be called, but the preview will be null.
*/
//go:nosplit
func (self class) QueueResourcePreview(path gd.String, receiver gd.Object, receiver_func gd.StringName, userdata gd.Variant)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(receiver))
	callframe.Arg(frame, discreet.Get(receiver_func))
	callframe.Arg(frame, discreet.Get(userdata))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePreview.Bind_queue_resource_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Queue the [param resource] being edited for preview. Once the preview is ready, the [param receiver]'s [param receiver_func] will be called. The [param receiver_func] must take the following four arguments: [String] path, [Texture2D] preview, [Texture2D] thumbnail_preview, [Variant] userdata. [param userdata] can be anything, and will be returned when [param receiver_func] is called.
[b]Note:[/b] If it was not possible to create the preview the [param receiver_func] will still be called, but the preview will be null.
*/
//go:nosplit
func (self class) QueueEditedResourcePreview(resource gdclass.Resource, receiver gd.Object, receiver_func gd.StringName, userdata gd.Variant)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(resource[0])[0])
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(receiver))
	callframe.Arg(frame, discreet.Get(receiver_func))
	callframe.Arg(frame, discreet.Get(userdata))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePreview.Bind_queue_edited_resource_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Create an own, custom preview generator.
*/
//go:nosplit
func (self class) AddPreviewGenerator(generator gdclass.EditorResourcePreviewGenerator)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(generator[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePreview.Bind_add_preview_generator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a custom preview generator.
*/
//go:nosplit
func (self class) RemovePreviewGenerator(generator gdclass.EditorResourcePreviewGenerator)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(generator[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePreview.Bind_remove_preview_generator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Check if the resource changed, if so, it will be invalidated and the corresponding signal emitted.
*/
//go:nosplit
func (self class) CheckForInvalidation(path gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePreview.Bind_check_for_invalidation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnPreviewInvalidated(cb func(path string)) {
	self[0].AsObject().Connect(gd.NewStringName("preview_invalidated"), gd.NewCallable(cb), 0)
}


func (self class) AsEditorResourcePreview() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorResourcePreview() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("EditorResourcePreview", func(ptr gd.Object) any { return classdb.EditorResourcePreview(ptr) })}
