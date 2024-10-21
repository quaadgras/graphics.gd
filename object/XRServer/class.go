package XRServer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The AR/VR server is the heart of our Advanced and Virtual Reality solution and handles all the processing.

*/
type Simple [1]classdb.XRServer
func (self Simple) GetWorldScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetWorldScale()))
}
func (self Simple) SetWorldScale(scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWorldScale(gd.Float(scale))
}
func (self Simple) GetWorldOrigin() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetWorldOrigin())
}
func (self Simple) SetWorldOrigin(world_origin gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWorldOrigin(world_origin)
}
func (self Simple) GetReferenceFrame() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetReferenceFrame())
}
func (self Simple) ClearReferenceFrame() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearReferenceFrame()
}
func (self Simple) CenterOnHmd(rotation_mode classdb.XRServerRotationMode, keep_height bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CenterOnHmd(rotation_mode, keep_height)
}
func (self Simple) GetHmdTransform() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetHmdTransform())
}
func (self Simple) AddInterface(intf [1]classdb.XRInterface) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddInterface(intf)
}
func (self Simple) GetInterfaceCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInterfaceCount()))
}
func (self Simple) RemoveInterface(intf [1]classdb.XRInterface) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveInterface(intf)
}
func (self Simple) GetInterface(idx int) [1]classdb.XRInterface {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.XRInterface(Expert(self).GetInterface(gc, gd.Int(idx)))
}
func (self Simple) GetInterfaces() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetInterfaces(gc))
}
func (self Simple) FindInterface(name string) [1]classdb.XRInterface {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.XRInterface(Expert(self).FindInterface(gc, gc.String(name)))
}
func (self Simple) AddTracker(tracker [1]classdb.XRTracker) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddTracker(tracker)
}
func (self Simple) RemoveTracker(tracker [1]classdb.XRTracker) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveTracker(tracker)
}
func (self Simple) GetTrackers(tracker_types int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetTrackers(gc, gd.Int(tracker_types)))
}
func (self Simple) GetTracker(tracker_name string) [1]classdb.XRTracker {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.XRTracker(Expert(self).GetTracker(gc, gc.StringName(tracker_name)))
}
func (self Simple) GetPrimaryInterface() [1]classdb.XRInterface {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.XRInterface(Expert(self).GetPrimaryInterface(gc))
}
func (self Simple) SetPrimaryInterface(intf [1]classdb.XRInterface) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPrimaryInterface(intf)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.XRServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) GetWorldScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_world_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWorldScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_set_world_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWorldOrigin() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_world_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWorldOrigin(world_origin gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, world_origin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_set_world_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the reference frame transform. Mostly used internally and exposed for GDExtension build interfaces.
*/
//go:nosplit
func (self class) GetReferenceFrame() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_reference_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears the reference frame that was set by previous calls to [method center_on_hmd].
*/
//go:nosplit
func (self class) ClearReferenceFrame()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_clear_reference_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This is an important function to understand correctly. AR and VR platforms all handle positioning slightly differently.
For platforms that do not offer spatial tracking, our origin point (0, 0, 0) is the location of our HMD, but you have little control over the direction the player is facing in the real world.
For platforms that do offer spatial tracking, our origin point depends very much on the system. For OpenVR, our origin point is usually the center of the tracking space, on the ground. For other platforms, it's often the location of the tracking camera.
This method allows you to center your tracker on the location of the HMD. It will take the current location of the HMD and use that to adjust all your tracking data; in essence, realigning the real world to your player's current position in the game world.
For this method to produce usable results, tracking information must be available. This often takes a few frames after starting your game.
You should call this method after a few seconds have passed. For example, when the user requests a realignment of the display holding a designated button on a controller for a short period of time, or when implementing a teleport mechanism.
*/
//go:nosplit
func (self class) CenterOnHmd(rotation_mode classdb.XRServerRotationMode, keep_height bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rotation_mode)
	callframe.Arg(frame, keep_height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_center_on_hmd, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the primary interface's transformation.
*/
//go:nosplit
func (self class) GetHmdTransform() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_hmd_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Registers an [XRInterface] object.
*/
//go:nosplit
func (self class) AddInterface(intf object.XRInterface)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(intf[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_add_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of interfaces currently registered with the AR/VR server. If your project supports multiple AR/VR platforms, you can look through the available interface, and either present the user with a selection or simply try to initialize each interface and use the first one that returns [code]true[/code].
*/
//go:nosplit
func (self class) GetInterfaceCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_interface_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes this [param interface].
*/
//go:nosplit
func (self class) RemoveInterface(intf object.XRInterface)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(intf[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_remove_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the interface registered at the given [param idx] index in the list of interfaces.
*/
//go:nosplit
func (self class) GetInterface(ctx gd.Lifetime, idx gd.Int) object.XRInterface {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.XRInterface
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a list of available interfaces the ID and name of each interface.
*/
//go:nosplit
func (self class) GetInterfaces(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_interfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Finds an interface by its [param name]. For example, if your project uses capabilities of an AR/VR platform, you can find the interface for that platform by name and initialize it.
*/
//go:nosplit
func (self class) FindInterface(ctx gd.Lifetime, name gd.String) object.XRInterface {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_find_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.XRInterface
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Registers a new [XRTracker] that tracks a physical object.
*/
//go:nosplit
func (self class) AddTracker(tracker object.XRTracker)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tracker[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_add_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes this [param tracker].
*/
//go:nosplit
func (self class) RemoveTracker(tracker object.XRTracker)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tracker[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_remove_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a dictionary of trackers for [param tracker_types].
*/
//go:nosplit
func (self class) GetTrackers(ctx gd.Lifetime, tracker_types gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, tracker_types)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_trackers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the positional tracker with the given [param tracker_name].
*/
//go:nosplit
func (self class) GetTracker(ctx gd.Lifetime, tracker_name gd.StringName) object.XRTracker {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tracker_name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.XRTracker
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetPrimaryInterface(ctx gd.Lifetime) object.XRInterface {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_get_primary_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.XRInterface
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPrimaryInterface(intf object.XRInterface)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(intf[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRServer.Bind_set_primary_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsXRServer() Expert { return self[0].AsXRServer() }


//go:nosplit
func (self Simple) AsXRServer() Simple { return self[0].AsXRServer() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("XRServer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type TrackerType = classdb.XRServerTrackerType

const (
/*The tracker tracks the location of the players head. This is usually a location centered between the players eyes. Note that for handheld AR devices this can be the current location of the device.*/
	TrackerHead TrackerType = 1
/*The tracker tracks the location of a controller.*/
	TrackerController TrackerType = 2
/*The tracker tracks the location of a base station.*/
	TrackerBasestation TrackerType = 4
/*The tracker tracks the location and size of an AR anchor.*/
	TrackerAnchor TrackerType = 8
/*The tracker tracks the location and joints of a hand.*/
	TrackerHand TrackerType = 16
/*The tracker tracks the location and joints of a body.*/
	TrackerBody TrackerType = 32
/*The tracker tracks the expressions of a face.*/
	TrackerFace TrackerType = 64
/*Used internally to filter trackers of any known type.*/
	TrackerAnyKnown TrackerType = 127
/*Used internally if we haven't set the tracker type yet.*/
	TrackerUnknown TrackerType = 128
/*Used internally to select all trackers.*/
	TrackerAny TrackerType = 255
)
type RotationMode = classdb.XRServerRotationMode

const (
/*Fully reset the orientation of the HMD. Regardless of what direction the user is looking to in the real world. The user will look dead ahead in the virtual world.*/
	ResetFullRotation RotationMode = 0
/*Resets the orientation but keeps the tilt of the device. So if we're looking down, we keep looking down but heading will be reset.*/
	ResetButKeepTilt RotationMode = 1
/*Does not reset the orientation of the HMD, only the position of the player gets centered.*/
	DontResetRotation RotationMode = 2
)
