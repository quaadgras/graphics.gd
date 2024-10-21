package EditorResourcePreview

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This node is used to generate previews for resources or files.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_resource_previewer].

*/
type Simple [1]classdb.EditorResourcePreview
func (self Simple) QueueResourcePreview(path string, receiver gd.Object, receiver_func string, userdata gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).QueueResourcePreview(gc.String(path), receiver, gc.StringName(receiver_func), userdata)
}
func (self Simple) QueueEditedResourcePreview(resource [1]classdb.Resource, receiver gd.Object, receiver_func string, userdata gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).QueueEditedResourcePreview(resource, receiver, gc.StringName(receiver_func), userdata)
}
func (self Simple) AddPreviewGenerator(generator [1]classdb.EditorResourcePreviewGenerator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddPreviewGenerator(generator)
}
func (self Simple) RemovePreviewGenerator(generator [1]classdb.EditorResourcePreviewGenerator) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemovePreviewGenerator(generator)
}
func (self Simple) CheckForInvalidation(path string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CheckForInvalidation(gc.String(path))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorResourcePreview
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Queue a resource file located at [param path] for preview. Once the preview is ready, the [param receiver]'s [param receiver_func] will be called. The [param receiver_func] must take the following four arguments: [String] path, [Texture2D] preview, [Texture2D] thumbnail_preview, [Variant] userdata. [param userdata] can be anything, and will be returned when [param receiver_func] is called.
[b]Note:[/b] If it was not possible to create the preview the [param receiver_func] will still be called, but the preview will be null.
*/
//go:nosplit
func (self class) QueueResourcePreview(path gd.String, receiver gd.Object, receiver_func gd.StringName, userdata gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.End(receiver.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(receiver_func))
	callframe.Arg(frame, mmm.Get(userdata))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePreview.Bind_queue_resource_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Queue the [param resource] being edited for preview. Once the preview is ready, the [param receiver]'s [param receiver_func] will be called. The [param receiver_func] must take the following four arguments: [String] path, [Texture2D] preview, [Texture2D] thumbnail_preview, [Variant] userdata. [param userdata] can be anything, and will be returned when [param receiver_func] is called.
[b]Note:[/b] If it was not possible to create the preview the [param receiver_func] will still be called, but the preview will be null.
*/
//go:nosplit
func (self class) QueueEditedResourcePreview(resource object.Resource, receiver gd.Object, receiver_func gd.StringName, userdata gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(resource[0].AsPointer())[0])
	callframe.Arg(frame, mmm.End(receiver.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(receiver_func))
	callframe.Arg(frame, mmm.Get(userdata))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePreview.Bind_queue_edited_resource_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Create an own, custom preview generator.
*/
//go:nosplit
func (self class) AddPreviewGenerator(generator object.EditorResourcePreviewGenerator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(generator[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePreview.Bind_add_preview_generator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a custom preview generator.
*/
//go:nosplit
func (self class) RemovePreviewGenerator(generator object.EditorResourcePreviewGenerator)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(generator[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePreview.Bind_remove_preview_generator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Check if the resource changed, if so, it will be invalidated and the corresponding signal emitted.
*/
//go:nosplit
func (self class) CheckForInvalidation(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorResourcePreview.Bind_check_for_invalidation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsEditorResourcePreview() Expert { return self[0].AsEditorResourcePreview() }


//go:nosplit
func (self Simple) AsEditorResourcePreview() Simple { return self[0].AsEditorResourcePreview() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorResourcePreview", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
