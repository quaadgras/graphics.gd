package CameraServer

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
The [CameraServer] keeps track of different cameras accessible in Godot. These are external cameras such as webcams or the cameras on your phone.
It is notably used to provide AR modules with a video feed from the camera.
[b]Note:[/b] This class is currently only implemented on macOS and iOS. To get a [CameraFeed] on iOS, the camera plugin from [url=https://github.com/godotengine/godot-ios-plugins]godot-ios-plugins[/url] is required. On other platforms, no [CameraFeed]s will be available.
*/
var self objects.CameraServer
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.CameraServer)
	self = *(*objects.CameraServer)(unsafe.Pointer(&obj))
}

/*
Returns the [CameraFeed] corresponding to the camera with the given [param index].
*/
func GetFeed(index int) objects.CameraFeed {
	once.Do(singleton)
	return objects.CameraFeed(class(self).GetFeed(gd.Int(index)))
}

/*
Returns the number of [CameraFeed]s registered.
*/
func GetFeedCount() int {
	once.Do(singleton)
	return int(int(class(self).GetFeedCount()))
}

/*
Returns an array of [CameraFeed]s.
*/
func Feeds() gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).Feeds())
}

/*
Adds the camera [param feed] to the camera server.
*/
func AddFeed(feed objects.CameraFeed) {
	once.Do(singleton)
	class(self).AddFeed(feed)
}

/*
Removes the specified camera [param feed].
*/
func RemoveFeed(feed objects.CameraFeed) {
	once.Do(singleton)
	class(self).RemoveFeed(feed)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.CameraServer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns the [CameraFeed] corresponding to the camera with the given [param index].
*/
//go:nosplit
func (self class) GetFeed(index gd.Int) objects.CameraFeed {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraServer.Bind_get_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.CameraFeed{classdb.CameraFeed(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the number of [CameraFeed]s registered.
*/
//go:nosplit
func (self class) GetFeedCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraServer.Bind_get_feed_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of [CameraFeed]s.
*/
//go:nosplit
func (self class) Feeds() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraServer.Bind_feeds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds the camera [param feed] to the camera server.
*/
//go:nosplit
func (self class) AddFeed(feed objects.CameraFeed) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(feed[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraServer.Bind_add_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the specified camera [param feed].
*/
//go:nosplit
func (self class) RemoveFeed(feed objects.CameraFeed) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(feed[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraServer.Bind_remove_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func OnCameraFeedAdded(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("camera_feed_added"), gd.NewCallable(cb), 0)
}

func OnCameraFeedRemoved(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("camera_feed_removed"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("CameraServer", func(ptr gd.Object) any { return classdb.CameraServer(ptr) })
}

type FeedImage = classdb.CameraServerFeedImage

const (
	/*The RGBA camera image.*/
	FeedRgbaImage FeedImage = 0
	/*The [url=https://en.wikipedia.org/wiki/YCbCr]YCbCr[/url] camera image.*/
	FeedYcbcrImage FeedImage = 0
	/*The Y component camera image.*/
	FeedYImage FeedImage = 0
	/*The CbCr component camera image.*/
	FeedCbcrImage FeedImage = 1
)
