// Package Performance provides methods for working with Performance object instances.
package Performance

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Array"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class provides access to a number of different monitors related to performance, such as memory usage, draw calls, and FPS. These are the same as the values displayed in the [b]Monitor[/b] tab in the editor's [b]Debugger[/b] panel. By using the [method get_monitor] method of this class, you can access this data from your code.
You can add custom monitors using the [method add_custom_monitor] method. Custom monitors are available in [b]Monitor[/b] tab in the editor's [b]Debugger[/b] panel together with built-in monitors.
[b]Note:[/b] Some of the built-in monitors are only available in debug mode and will always return [code]0[/code] when used in a project exported in release mode.
[b]Note:[/b] Some of the built-in monitors are not updated in real-time for performance reasons, so there may be a delay of up to 1 second between changes.
[b]Note:[/b] Custom monitors do not support negative values. Negative values are clamped to 0.
*/
var self [1]gdclass.Performance
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.Performance)
	self = *(*[1]gdclass.Performance)(unsafe.Pointer(&obj))
}

/*
Returns the value of one of the available built-in monitors. You should provide one of the [enum Monitor] constants as the argument, like this:
[codeblocks]
[gdscript]
print(Performance.get_monitor(Performance.TIME_FPS)) # Prints the FPS to the console.
[/gdscript]
[csharp]
GD.Print(Performance.GetMonitor(Performance.Monitor.TimeFps)); // Prints the FPS to the console.
[/csharp]
[/codeblocks]
See [method get_custom_monitor] to query custom performance monitors' values.
*/
func GetMonitor(monitor gdclass.PerformanceMonitor) Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetMonitor(monitor)))
}

/*
Adds a custom monitor with the name [param id]. You can specify the category of the monitor using slash delimiters in [param id] (for example: [code]"Game/NumberOfNPCs"[/code]). If there is more than one slash delimiter, then the default category is used. The default category is [code]"Custom"[/code]. Prints an error if given [param id] is already present.
[codeblocks]
[gdscript]
func _ready():

	var monitor_value = Callable(self, "get_monitor_value")

	# Adds monitor with name "MyName" to category "MyCategory".
	Performance.add_custom_monitor("MyCategory/MyMonitor", monitor_value)

	# Adds monitor with name "MyName" to category "Custom".
	# Note: "MyCategory/MyMonitor" and "MyMonitor" have same name but different IDs, so the code is valid.
	Performance.add_custom_monitor("MyMonitor", monitor_value)

	# Adds monitor with name "MyName" to category "Custom".
	# Note: "MyMonitor" and "Custom/MyMonitor" have same name and same category but different IDs, so the code is valid.
	Performance.add_custom_monitor("Custom/MyMonitor", monitor_value)

	# Adds monitor with name "MyCategoryOne/MyCategoryTwo/MyMonitor" to category "Custom".
	Performance.add_custom_monitor("MyCategoryOne/MyCategoryTwo/MyMonitor", monitor_value)

func get_monitor_value():

	return randi() % 25

[/gdscript]
[csharp]
public override void _Ready()

	{
	    var monitorValue = new Callable(this, MethodName.GetMonitorValue);

	    // Adds monitor with name "MyName" to category "MyCategory".
	    Performance.AddCustomMonitor("MyCategory/MyMonitor", monitorValue);
	    // Adds monitor with name "MyName" to category "Custom".
	    // Note: "MyCategory/MyMonitor" and "MyMonitor" have same name but different ids so the code is valid.
	    Performance.AddCustomMonitor("MyMonitor", monitorValue);

	    // Adds monitor with name "MyName" to category "Custom".
	    // Note: "MyMonitor" and "Custom/MyMonitor" have same name and same category but different ids so the code is valid.
	    Performance.AddCustomMonitor("Custom/MyMonitor", monitorValue);

	    // Adds monitor with name "MyCategoryOne/MyCategoryTwo/MyMonitor" to category "Custom".
	    Performance.AddCustomMonitor("MyCategoryOne/MyCategoryTwo/MyMonitor", monitorValue);
	}

public int GetMonitorValue()

	{
	    return GD.Randi() % 25;
	}

[/csharp]
[/codeblocks]
The debugger calls the callable to get the value of custom monitor. The callable must return a zero or positive integer or floating-point number.
Callables are called with arguments supplied in argument array.
*/
func AddCustomMonitor(id string, callable Callable.Any) {
	once.Do(singleton)
	class(self).AddCustomMonitor(gd.NewStringName(id), gd.NewCallable(callable), [1]Array.Any{}[0])
}

/*
Removes the custom monitor with given [param id]. Prints an error if the given [param id] is already absent.
*/
func RemoveCustomMonitor(id string) {
	once.Do(singleton)
	class(self).RemoveCustomMonitor(gd.NewStringName(id))
}

/*
Returns [code]true[/code] if custom monitor with the given [param id] is present, [code]false[/code] otherwise.
*/
func HasCustomMonitor(id string) bool {
	once.Do(singleton)
	return bool(class(self).HasCustomMonitor(gd.NewStringName(id)))
}

/*
Returns the value of custom monitor with given [param id]. The callable is called to get the value of custom monitor. See also [method has_custom_monitor]. Prints an error if the given [param id] is absent.
*/
func GetCustomMonitor(id string) any {
	once.Do(singleton)
	return any(class(self).GetCustomMonitor(gd.NewStringName(id)).Interface())
}

/*
Returns the last tick in which custom monitor was added/removed (in microseconds since the engine started). This is set to [method Time.get_ticks_usec] when the monitor is updated.
*/
func GetMonitorModificationTime() int {
	once.Do(singleton)
	return int(int(class(self).GetMonitorModificationTime()))
}

/*
Returns the names of active custom monitors in an [Array].
*/
func GetCustomMonitorNames() gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).GetCustomMonitorNames())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.Performance

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns the value of one of the available built-in monitors. You should provide one of the [enum Monitor] constants as the argument, like this:
[codeblocks]
[gdscript]
print(Performance.get_monitor(Performance.TIME_FPS)) # Prints the FPS to the console.
[/gdscript]
[csharp]
GD.Print(Performance.GetMonitor(Performance.Monitor.TimeFps)); // Prints the FPS to the console.
[/csharp]
[/codeblocks]
See [method get_custom_monitor] to query custom performance monitors' values.
*/
//go:nosplit
func (self class) GetMonitor(monitor gdclass.PerformanceMonitor) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, monitor)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Performance.Bind_get_monitor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a custom monitor with the name [param id]. You can specify the category of the monitor using slash delimiters in [param id] (for example: [code]"Game/NumberOfNPCs"[/code]). If there is more than one slash delimiter, then the default category is used. The default category is [code]"Custom"[/code]. Prints an error if given [param id] is already present.
[codeblocks]
[gdscript]
func _ready():
    var monitor_value = Callable(self, "get_monitor_value")

    # Adds monitor with name "MyName" to category "MyCategory".
    Performance.add_custom_monitor("MyCategory/MyMonitor", monitor_value)

    # Adds monitor with name "MyName" to category "Custom".
    # Note: "MyCategory/MyMonitor" and "MyMonitor" have same name but different IDs, so the code is valid.
    Performance.add_custom_monitor("MyMonitor", monitor_value)

    # Adds monitor with name "MyName" to category "Custom".
    # Note: "MyMonitor" and "Custom/MyMonitor" have same name and same category but different IDs, so the code is valid.
    Performance.add_custom_monitor("Custom/MyMonitor", monitor_value)

    # Adds monitor with name "MyCategoryOne/MyCategoryTwo/MyMonitor" to category "Custom".
    Performance.add_custom_monitor("MyCategoryOne/MyCategoryTwo/MyMonitor", monitor_value)

func get_monitor_value():
    return randi() % 25
[/gdscript]
[csharp]
public override void _Ready()
{
    var monitorValue = new Callable(this, MethodName.GetMonitorValue);

    // Adds monitor with name "MyName" to category "MyCategory".
    Performance.AddCustomMonitor("MyCategory/MyMonitor", monitorValue);
    // Adds monitor with name "MyName" to category "Custom".
    // Note: "MyCategory/MyMonitor" and "MyMonitor" have same name but different ids so the code is valid.
    Performance.AddCustomMonitor("MyMonitor", monitorValue);

    // Adds monitor with name "MyName" to category "Custom".
    // Note: "MyMonitor" and "Custom/MyMonitor" have same name and same category but different ids so the code is valid.
    Performance.AddCustomMonitor("Custom/MyMonitor", monitorValue);

    // Adds monitor with name "MyCategoryOne/MyCategoryTwo/MyMonitor" to category "Custom".
    Performance.AddCustomMonitor("MyCategoryOne/MyCategoryTwo/MyMonitor", monitorValue);
}

public int GetMonitorValue()
{
    return GD.Randi() % 25;
}
[/csharp]
[/codeblocks]
The debugger calls the callable to get the value of custom monitor. The callable must return a zero or positive integer or floating-point number.
Callables are called with arguments supplied in argument array.
*/
//go:nosplit
func (self class) AddCustomMonitor(id gd.StringName, callable gd.Callable, arguments gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(id))
	callframe.Arg(frame, pointers.Get(callable))
	callframe.Arg(frame, pointers.Get(arguments))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Performance.Bind_add_custom_monitor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the custom monitor with given [param id]. Prints an error if the given [param id] is already absent.
*/
//go:nosplit
func (self class) RemoveCustomMonitor(id gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(id))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Performance.Bind_remove_custom_monitor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if custom monitor with the given [param id] is present, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) HasCustomMonitor(id gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(id))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Performance.Bind_has_custom_monitor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the value of custom monitor with given [param id]. The callable is called to get the value of custom monitor. See also [method has_custom_monitor]. Prints an error if the given [param id] is absent.
*/
//go:nosplit
func (self class) GetCustomMonitor(id gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(id))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Performance.Bind_get_custom_monitor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the last tick in which custom monitor was added/removed (in microseconds since the engine started). This is set to [method Time.get_ticks_usec] when the monitor is updated.
*/
//go:nosplit
func (self class) GetMonitorModificationTime() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Performance.Bind_get_monitor_modification_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the names of active custom monitors in an [Array].
*/
//go:nosplit
func (self class) GetCustomMonitorNames() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Performance.Bind_get_custom_monitor_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("Performance", func(ptr gd.Object) any { return [1]gdclass.Performance{*(*gdclass.Performance)(unsafe.Pointer(&ptr))} })
}

type Monitor = gdclass.PerformanceMonitor

const (
	/*The number of frames rendered in the last second. This metric is only updated once per second, even if queried more often. [i]Higher is better.[/i]*/
	TimeFps Monitor = 0
	/*Time it took to complete one frame, in seconds. [i]Lower is better.[/i]*/
	TimeProcess Monitor = 1
	/*Time it took to complete one physics frame, in seconds. [i]Lower is better.[/i]*/
	TimePhysicsProcess Monitor = 2
	/*Time it took to complete one navigation step, in seconds. This includes navigation map updates as well as agent avoidance calculations. [i]Lower is better.[/i]*/
	TimeNavigationProcess Monitor = 3
	/*Static memory currently used, in bytes. Not available in release builds. [i]Lower is better.[/i]*/
	MemoryStatic Monitor = 4
	/*Available static memory. Not available in release builds. [i]Lower is better.[/i]*/
	MemoryStaticMax Monitor = 5
	/*Largest amount of memory the message queue buffer has used, in bytes. The message queue is used for deferred functions calls and notifications. [i]Lower is better.[/i]*/
	MemoryMessageBufferMax Monitor = 6
	/*Number of objects currently instantiated (including nodes). [i]Lower is better.[/i]*/
	ObjectCount Monitor = 7
	/*Number of resources currently used. [i]Lower is better.[/i]*/
	ObjectResourceCount Monitor = 8
	/*Number of nodes currently instantiated in the scene tree. This also includes the root node. [i]Lower is better.[/i]*/
	ObjectNodeCount Monitor = 9
	/*Number of orphan nodes, i.e. nodes which are not parented to a node of the scene tree. [i]Lower is better.[/i]*/
	ObjectOrphanNodeCount Monitor = 10
	/*The total number of objects in the last rendered frame. This metric doesn't include culled objects (either via hiding nodes, frustum culling or occlusion culling). [i]Lower is better.[/i]*/
	RenderTotalObjectsInFrame Monitor = 11
	/*The total number of vertices or indices rendered in the last rendered frame. This metric doesn't include primitives from culled objects (either via hiding nodes, frustum culling or occlusion culling). Due to the depth prepass and shadow passes, the number of primitives is always higher than the actual number of vertices in the scene (typically double or triple the original vertex count). [i]Lower is better.[/i]*/
	RenderTotalPrimitivesInFrame Monitor = 12
	/*The total number of draw calls performed in the last rendered frame. This metric doesn't include culled objects (either via hiding nodes, frustum culling or occlusion culling), since they do not result in draw calls. [i]Lower is better.[/i]*/
	RenderTotalDrawCallsInFrame Monitor = 13
	/*The amount of video memory used (texture and vertex memory combined, in bytes). Since this metric also includes miscellaneous allocations, this value is always greater than the sum of [constant RENDER_TEXTURE_MEM_USED] and [constant RENDER_BUFFER_MEM_USED]. [i]Lower is better.[/i]*/
	RenderVideoMemUsed Monitor = 14
	/*The amount of texture memory used (in bytes). [i]Lower is better.[/i]*/
	RenderTextureMemUsed Monitor = 15
	/*The amount of render buffer memory used (in bytes). [i]Lower is better.[/i]*/
	RenderBufferMemUsed Monitor = 16
	/*Number of active [RigidBody2D] nodes in the game. [i]Lower is better.[/i]*/
	Physics2dActiveObjects Monitor = 17
	/*Number of collision pairs in the 2D physics engine. [i]Lower is better.[/i]*/
	Physics2dCollisionPairs Monitor = 18
	/*Number of islands in the 2D physics engine. [i]Lower is better.[/i]*/
	Physics2dIslandCount Monitor = 19
	/*Number of active [RigidBody3D] and [VehicleBody3D] nodes in the game. [i]Lower is better.[/i]*/
	Physics3dActiveObjects Monitor = 20
	/*Number of collision pairs in the 3D physics engine. [i]Lower is better.[/i]*/
	Physics3dCollisionPairs Monitor = 21
	/*Number of islands in the 3D physics engine. [i]Lower is better.[/i]*/
	Physics3dIslandCount Monitor = 22
	/*Output latency of the [AudioServer]. Equivalent to calling [method AudioServer.get_output_latency], it is not recommended to call this every frame.*/
	AudioOutputLatency Monitor = 23
	/*Number of active navigation maps in the [NavigationServer3D]. This also includes the two empty default navigation maps created by World2D and World3D.*/
	NavigationActiveMaps Monitor = 24
	/*Number of active navigation regions in the [NavigationServer3D].*/
	NavigationRegionCount Monitor = 25
	/*Number of active navigation agents processing avoidance in the [NavigationServer3D].*/
	NavigationAgentCount Monitor = 26
	/*Number of active navigation links in the [NavigationServer3D].*/
	NavigationLinkCount Monitor = 27
	/*Number of navigation mesh polygons in the [NavigationServer3D].*/
	NavigationPolygonCount Monitor = 28
	/*Number of navigation mesh polygon edges in the [NavigationServer3D].*/
	NavigationEdgeCount Monitor = 29
	/*Number of navigation mesh polygon edges that were merged due to edge key overlap in the [NavigationServer3D].*/
	NavigationEdgeMergeCount Monitor = 30
	/*Number of polygon edges that are considered connected by edge proximity [NavigationServer3D].*/
	NavigationEdgeConnectionCount Monitor = 31
	/*Number of navigation mesh polygon edges that could not be merged in the [NavigationServer3D]. The edges still may be connected by edge proximity or with links.*/
	NavigationEdgeFreeCount Monitor = 32
	/*Represents the size of the [enum Monitor] enum.*/
	MonitorMax Monitor = 33
)
