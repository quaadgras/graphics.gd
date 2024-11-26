package XRServer

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Transform3D"
import "grow.graphics/gd/variant/Dictionary"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
The AR/VR server is the heart of our Advanced and Virtual Reality solution and handles all the processing.
*/
var self objects.XRServer
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.XRServer)
	self = *(*objects.XRServer)(unsafe.Pointer(&obj))
}

/*
Returns the reference frame transform. Mostly used internally and exposed for GDExtension build interfaces.
*/
func GetReferenceFrame() Transform3D.BasisOrigin {
	once.Do(singleton)
	return Transform3D.BasisOrigin(class(self).GetReferenceFrame())
}

/*
Clears the reference frame that was set by previous calls to [method center_on_hmd].
*/
func ClearReferenceFrame() {
	once.Do(singleton)
	class(self).ClearReferenceFrame()
}

/*
This is an important function to understand correctly. AR and VR platforms all handle positioning slightly differently.
For platforms that do not offer spatial tracking, our origin point (0, 0, 0) is the location of our HMD, but you have little control over the direction the player is facing in the real world.
For platforms that do offer spatial tracking, our origin point depends very much on the system. For OpenVR, our origin point is usually the center of the tracking space, on the ground. For other platforms, it's often the location of the tracking camera.
This method allows you to center your tracker on the location of the HMD. It will take the current location of the HMD and use that to adjust all your tracking data; in essence, realigning the real world to your player's current position in the game world.
For this method to produce usable results, tracking information must be available. This often takes a few frames after starting your game.
You should call this method after a few seconds have passed. For example, when the user requests a realignment of the display holding a designated button on a controller for a short period of time, or when implementing a teleport mechanism.
*/
func CenterOnHmd(rotation_mode classdb.XRServerRotationMode, keep_height bool) {
	once.Do(singleton)
	class(self).CenterOnHmd(rotation_mode, keep_height)
}

/*
Returns the primary interface's transformation.
*/
func GetHmdTransform() Transform3D.BasisOrigin {
	once.Do(singleton)
	return Transform3D.BasisOrigin(class(self).GetHmdTransform())
}

/*
Registers an [XRInterface] object.
*/
func AddInterface(intf objects.XRInterface) {
	once.Do(singleton)
	class(self).AddInterface(intf)
}

/*
Returns the number of interfaces currently registered with the AR/VR server. If your project supports multiple AR/VR platforms, you can look through the available interface, and either present the user with a selection or simply try to initialize each interface and use the first one that returns [code]true[/code].
*/
func GetInterfaceCount() int {
	once.Do(singleton)
	return int(int(class(self).GetInterfaceCount()))
}

/*
Removes this [param interface].
*/
func RemoveInterface(intf objects.XRInterface) {
	once.Do(singleton)
	class(self).RemoveInterface(intf)
}

/*
Returns the interface registered at the given [param idx] index in the list of interfaces.
*/
func GetInterface(idx int) objects.XRInterface {
	once.Do(singleton)
	return objects.XRInterface(class(self).GetInterface(gd.Int(idx)))
}

/*
Returns a list of available interfaces the ID and name of each interface.
*/
func GetInterfaces() gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).GetInterfaces())
}

/*
Finds an interface by its [param name]. For example, if your project uses capabilities of an AR/VR platform, you can find the interface for that platform by name and initialize it.
*/
func FindInterface(name string) objects.XRInterface {
	once.Do(singleton)
	return objects.XRInterface(class(self).FindInterface(gd.NewString(name)))
}

/*
Registers a new [XRTracker] that tracks a physical object.
*/
func AddTracker(tracker objects.XRTracker) {
	once.Do(singleton)
	class(self).AddTracker(tracker)
}

/*
Removes this [param tracker].
*/
func RemoveTracker(tracker objects.XRTracker) {
	once.Do(singleton)
	class(self).RemoveTracker(tracker)
}

/*
Returns a dictionary of trackers for [param tracker_types].
*/
func GetTrackers(tracker_types int) Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetTrackers(gd.Int(tracker_types)))
}

/*
Returns the positional tracker with the given [param tracker_name].
*/
func GetTracker(tracker_name string) objects.XRTracker {
	once.Do(singleton)
	return objects.XRTracker(class(self).GetTracker(gd.NewStringName(tracker_name)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.XRServer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

func WorldScale() Float.X {
	return Float.X(Float.X(class(self).GetWorldScale()))
}

func SetWorldScale(value Float.X) {
	class(self).SetWorldScale(gd.Float(value))
}

func WorldOrigin() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetWorldOrigin())
}

func SetWorldOrigin(value Transform3D.BasisOrigin) {
	class(self).SetWorldOrigin(gd.Transform3D(value))
}

func PrimaryInterface() objects.XRInterface {
	return objects.XRInterface(class(self).GetPrimaryInterface())
}

func SetPrimaryInterface(value objects.XRInterface) {
	class(self).SetPrimaryInterface(value)
}

//go:nosplit
func (self class) GetWorldScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_world_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWorldScale(scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_set_world_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWorldOrigin() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_world_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWorldOrigin(world_origin gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, world_origin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_set_world_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the reference frame transform. Mostly used internally and exposed for GDExtension build interfaces.
*/
//go:nosplit
func (self class) GetReferenceFrame() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_reference_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears the reference frame that was set by previous calls to [method center_on_hmd].
*/
//go:nosplit
func (self class) ClearReferenceFrame() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_clear_reference_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) CenterOnHmd(rotation_mode classdb.XRServerRotationMode, keep_height bool) {
	var frame = callframe.New()
	callframe.Arg(frame, rotation_mode)
	callframe.Arg(frame, keep_height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_center_on_hmd, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the primary interface's transformation.
*/
//go:nosplit
func (self class) GetHmdTransform() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_hmd_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Registers an [XRInterface] object.
*/
//go:nosplit
func (self class) AddInterface(intf objects.XRInterface) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(intf[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_add_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of interfaces currently registered with the AR/VR server. If your project supports multiple AR/VR platforms, you can look through the available interface, and either present the user with a selection or simply try to initialize each interface and use the first one that returns [code]true[/code].
*/
//go:nosplit
func (self class) GetInterfaceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_interface_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes this [param interface].
*/
//go:nosplit
func (self class) RemoveInterface(intf objects.XRInterface) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(intf[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_remove_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the interface registered at the given [param idx] index in the list of interfaces.
*/
//go:nosplit
func (self class) GetInterface(idx gd.Int) objects.XRInterface {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.XRInterface{classdb.XRInterface(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns a list of available interfaces the ID and name of each interface.
*/
//go:nosplit
func (self class) GetInterfaces() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_interfaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Finds an interface by its [param name]. For example, if your project uses capabilities of an AR/VR platform, you can find the interface for that platform by name and initialize it.
*/
//go:nosplit
func (self class) FindInterface(name gd.String) objects.XRInterface {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_find_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.XRInterface{classdb.XRInterface(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Registers a new [XRTracker] that tracks a physical object.
*/
//go:nosplit
func (self class) AddTracker(tracker objects.XRTracker) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tracker[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_add_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes this [param tracker].
*/
//go:nosplit
func (self class) RemoveTracker(tracker objects.XRTracker) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tracker[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_remove_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a dictionary of trackers for [param tracker_types].
*/
//go:nosplit
func (self class) GetTrackers(tracker_types gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, tracker_types)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_trackers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the positional tracker with the given [param tracker_name].
*/
//go:nosplit
func (self class) GetTracker(tracker_name gd.StringName) objects.XRTracker {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tracker_name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.XRTracker{classdb.XRTracker(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetPrimaryInterface() objects.XRInterface {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_get_primary_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.XRInterface{classdb.XRInterface(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPrimaryInterface(intf objects.XRInterface) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(intf[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRServer.Bind_set_primary_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func OnReferenceFrameChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("reference_frame_changed"), gd.NewCallable(cb), 0)
}

func OnInterfaceAdded(cb func(interface_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("interface_added"), gd.NewCallable(cb), 0)
}

func OnInterfaceRemoved(cb func(interface_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("interface_removed"), gd.NewCallable(cb), 0)
}

func OnTrackerAdded(cb func(tracker_name string, atype int)) {
	self[0].AsObject().Connect(gd.NewStringName("tracker_added"), gd.NewCallable(cb), 0)
}

func OnTrackerUpdated(cb func(tracker_name string, atype int)) {
	self[0].AsObject().Connect(gd.NewStringName("tracker_updated"), gd.NewCallable(cb), 0)
}

func OnTrackerRemoved(cb func(tracker_name string, atype int)) {
	self[0].AsObject().Connect(gd.NewStringName("tracker_removed"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() { classdb.Register("XRServer", func(ptr gd.Object) any { return classdb.XRServer(ptr) }) }

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
