// Package EngineDebugger provides methods for working with EngineDebugger object instances.
package EngineDebugger

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any

/*
[EngineDebugger] handles the communication between the editor and the running game. It is active in the running game. Messages can be sent/received through it. It also manages the profilers.
*/
var self [1]gdclass.EngineDebugger
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.EngineDebugger)
	self = *(*[1]gdclass.EngineDebugger)(unsafe.Pointer(&obj))
}

/*
Returns [code]true[/code] if the debugger is active otherwise [code]false[/code].
*/
func IsActive() bool { //gd:EngineDebugger.is_active
	once.Do(singleton)
	return bool(class(self).IsActive())
}

/*
Registers a profiler with the given [param name]. See [EngineProfiler] for more information.
*/
func RegisterProfiler(name string, profiler [1]gdclass.EngineProfiler) { //gd:EngineDebugger.register_profiler
	once.Do(singleton)
	class(self).RegisterProfiler(gd.NewStringName(name), profiler)
}

/*
Unregisters a profiler with given [param name].
*/
func UnregisterProfiler(name string) { //gd:EngineDebugger.unregister_profiler
	once.Do(singleton)
	class(self).UnregisterProfiler(gd.NewStringName(name))
}

/*
Returns [code]true[/code] if a profiler with the given name is present and active otherwise [code]false[/code].
*/
func IsProfiling(name string) bool { //gd:EngineDebugger.is_profiling
	once.Do(singleton)
	return bool(class(self).IsProfiling(gd.NewStringName(name)))
}

/*
Returns [code]true[/code] if a profiler with the given name is present otherwise [code]false[/code].
*/
func HasProfiler(name string) bool { //gd:EngineDebugger.has_profiler
	once.Do(singleton)
	return bool(class(self).HasProfiler(gd.NewStringName(name)))
}

/*
Calls the [code]add[/code] callable of the profiler with given [param name] and [param data].
*/
func ProfilerAddFrameData(name string, data []any) { //gd:EngineDebugger.profiler_add_frame_data
	once.Do(singleton)
	class(self).ProfilerAddFrameData(gd.NewStringName(name), gd.EngineArrayFromSlice(data))
}

/*
Calls the [code]toggle[/code] callable of the profiler with given [param name] and [param arguments]. Enables/Disables the same profiler depending on [param enable] argument.
*/
func ProfilerEnable(name string, enable bool) { //gd:EngineDebugger.profiler_enable
	once.Do(singleton)
	class(self).ProfilerEnable(gd.NewStringName(name), enable, Array.Nil)
}

/*
Registers a message capture with given [param name]. If [param name] is "my_message" then messages starting with "my_message:" will be called with the given callable.
Callable must accept a message string and a data array as argument. If the message and data are valid then callable must return [code]true[/code] otherwise [code]false[/code].
*/
func RegisterMessageCapture(name string, callable func(message string, data []any)) { //gd:EngineDebugger.register_message_capture
	once.Do(singleton)
	class(self).RegisterMessageCapture(gd.NewStringName(name), Callable.New(callable))
}

/*
Unregisters the message capture with given [param name].
*/
func UnregisterMessageCapture(name string) { //gd:EngineDebugger.unregister_message_capture
	once.Do(singleton)
	class(self).UnregisterMessageCapture(gd.NewStringName(name))
}

/*
Returns [code]true[/code] if a capture with the given name is present otherwise [code]false[/code].
*/
func HasCapture(name string) bool { //gd:EngineDebugger.has_capture
	once.Do(singleton)
	return bool(class(self).HasCapture(gd.NewStringName(name)))
}

/*
Forces a processing loop of debugger events. The purpose of this method is just processing events every now and then when the script might get too busy, so that bugs like infinite loops can be caught
*/
func LinePoll() { //gd:EngineDebugger.line_poll
	once.Do(singleton)
	class(self).LinePoll()
}

/*
Sends a message with given [param message] and [param data] array.
*/
func SendMessage(message string, data []any) { //gd:EngineDebugger.send_message
	once.Do(singleton)
	class(self).SendMessage(gd.NewString(message), gd.EngineArrayFromSlice(data))
}

/*
Starts a debug break in script execution, optionally specifying whether the program can continue based on [param can_continue] and whether the break was due to a breakpoint.
*/
func Debug() { //gd:EngineDebugger.debug
	once.Do(singleton)
	class(self).Debug(true, false)
}

/*
Starts a debug break in script execution, optionally specifying whether the program can continue based on [param can_continue] and whether the break was due to a breakpoint.
*/
func ScriptDebug(language [1]gdclass.ScriptLanguage) { //gd:EngineDebugger.script_debug
	once.Do(singleton)
	class(self).ScriptDebug(language, true, false)
}

/*
Sets the current debugging lines that remain.
*/
func SetLinesLeft(lines int) { //gd:EngineDebugger.set_lines_left
	once.Do(singleton)
	class(self).SetLinesLeft(gd.Int(lines))
}

/*
Returns the number of lines that remain.
*/
func GetLinesLeft() int { //gd:EngineDebugger.get_lines_left
	once.Do(singleton)
	return int(int(class(self).GetLinesLeft()))
}

/*
Sets the current debugging depth.
*/
func SetDepth(depth int) { //gd:EngineDebugger.set_depth
	once.Do(singleton)
	class(self).SetDepth(gd.Int(depth))
}

/*
Returns the current debug depth.
*/
func GetDepth() int { //gd:EngineDebugger.get_depth
	once.Do(singleton)
	return int(int(class(self).GetDepth()))
}

/*
Returns [code]true[/code] if the given [param source] and [param line] represent an existing breakpoint.
*/
func IsBreakpoint(line int, source string) bool { //gd:EngineDebugger.is_breakpoint
	once.Do(singleton)
	return bool(class(self).IsBreakpoint(gd.Int(line), gd.NewStringName(source)))
}

/*
Returns [code]true[/code] if the debugger is skipping breakpoints otherwise [code]false[/code].
*/
func IsSkippingBreakpoints() bool { //gd:EngineDebugger.is_skipping_breakpoints
	once.Do(singleton)
	return bool(class(self).IsSkippingBreakpoints())
}

/*
Inserts a new breakpoint with the given [param source] and [param line].
*/
func InsertBreakpoint(line int, source string) { //gd:EngineDebugger.insert_breakpoint
	once.Do(singleton)
	class(self).InsertBreakpoint(gd.Int(line), gd.NewStringName(source))
}

/*
Removes a breakpoint with the given [param source] and [param line].
*/
func RemoveBreakpoint(line int, source string) { //gd:EngineDebugger.remove_breakpoint
	once.Do(singleton)
	class(self).RemoveBreakpoint(gd.Int(line), gd.NewStringName(source))
}

/*
Clears all breakpoints.
*/
func ClearBreakpoints() { //gd:EngineDebugger.clear_breakpoints
	once.Do(singleton)
	class(self).ClearBreakpoints()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.EngineDebugger

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns [code]true[/code] if the debugger is active otherwise [code]false[/code].
*/
//go:nosplit
func (self class) IsActive() bool { //gd:EngineDebugger.is_active
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Registers a profiler with the given [param name]. See [EngineProfiler] for more information.
*/
//go:nosplit
func (self class) RegisterProfiler(name gd.StringName, profiler [1]gdclass.EngineProfiler) { //gd:EngineDebugger.register_profiler
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(profiler[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_register_profiler, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unregisters a profiler with given [param name].
*/
//go:nosplit
func (self class) UnregisterProfiler(name gd.StringName) { //gd:EngineDebugger.unregister_profiler
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_unregister_profiler, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if a profiler with the given name is present and active otherwise [code]false[/code].
*/
//go:nosplit
func (self class) IsProfiling(name gd.StringName) bool { //gd:EngineDebugger.is_profiling
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_is_profiling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if a profiler with the given name is present otherwise [code]false[/code].
*/
//go:nosplit
func (self class) HasProfiler(name gd.StringName) bool { //gd:EngineDebugger.has_profiler
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_has_profiler, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Calls the [code]add[/code] callable of the profiler with given [param name] and [param data].
*/
//go:nosplit
func (self class) ProfilerAddFrameData(name gd.StringName, data Array.Any) { //gd:EngineDebugger.profiler_add_frame_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(data)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_profiler_add_frame_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Calls the [code]toggle[/code] callable of the profiler with given [param name] and [param arguments]. Enables/Disables the same profiler depending on [param enable] argument.
*/
//go:nosplit
func (self class) ProfilerEnable(name gd.StringName, enable bool, arguments Array.Any) { //gd:EngineDebugger.profiler_enable
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, enable)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(arguments)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_profiler_enable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers a message capture with given [param name]. If [param name] is "my_message" then messages starting with "my_message:" will be called with the given callable.
Callable must accept a message string and a data array as argument. If the message and data are valid then callable must return [code]true[/code] otherwise [code]false[/code].
*/
//go:nosplit
func (self class) RegisterMessageCapture(name gd.StringName, callable Callable.Function) { //gd:EngineDebugger.register_message_capture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_register_message_capture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unregisters the message capture with given [param name].
*/
//go:nosplit
func (self class) UnregisterMessageCapture(name gd.StringName) { //gd:EngineDebugger.unregister_message_capture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_unregister_message_capture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if a capture with the given name is present otherwise [code]false[/code].
*/
//go:nosplit
func (self class) HasCapture(name gd.StringName) bool { //gd:EngineDebugger.has_capture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_has_capture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Forces a processing loop of debugger events. The purpose of this method is just processing events every now and then when the script might get too busy, so that bugs like infinite loops can be caught
*/
//go:nosplit
func (self class) LinePoll() { //gd:EngineDebugger.line_poll
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_line_poll, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sends a message with given [param message] and [param data] array.
*/
//go:nosplit
func (self class) SendMessage(message gd.String, data Array.Any) { //gd:EngineDebugger.send_message
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(message))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(data)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_send_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Starts a debug break in script execution, optionally specifying whether the program can continue based on [param can_continue] and whether the break was due to a breakpoint.
*/
//go:nosplit
func (self class) Debug(can_continue bool, is_error_breakpoint bool) { //gd:EngineDebugger.debug
	var frame = callframe.New()
	callframe.Arg(frame, can_continue)
	callframe.Arg(frame, is_error_breakpoint)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_debug, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Starts a debug break in script execution, optionally specifying whether the program can continue based on [param can_continue] and whether the break was due to a breakpoint.
*/
//go:nosplit
func (self class) ScriptDebug(language [1]gdclass.ScriptLanguage, can_continue bool, is_error_breakpoint bool) { //gd:EngineDebugger.script_debug
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language[0])[0])
	callframe.Arg(frame, can_continue)
	callframe.Arg(frame, is_error_breakpoint)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_script_debug, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the current debugging lines that remain.
*/
//go:nosplit
func (self class) SetLinesLeft(lines gd.Int) { //gd:EngineDebugger.set_lines_left
	var frame = callframe.New()
	callframe.Arg(frame, lines)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_set_lines_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of lines that remain.
*/
//go:nosplit
func (self class) GetLinesLeft() gd.Int { //gd:EngineDebugger.get_lines_left
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_get_lines_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the current debugging depth.
*/
//go:nosplit
func (self class) SetDepth(depth gd.Int) { //gd:EngineDebugger.set_depth
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current debug depth.
*/
//go:nosplit
func (self class) GetDepth() gd.Int { //gd:EngineDebugger.get_depth
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [param source] and [param line] represent an existing breakpoint.
*/
//go:nosplit
func (self class) IsBreakpoint(line gd.Int, source gd.StringName) bool { //gd:EngineDebugger.is_breakpoint
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, pointers.Get(source))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_is_breakpoint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the debugger is skipping breakpoints otherwise [code]false[/code].
*/
//go:nosplit
func (self class) IsSkippingBreakpoints() bool { //gd:EngineDebugger.is_skipping_breakpoints
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_is_skipping_breakpoints, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Inserts a new breakpoint with the given [param source] and [param line].
*/
//go:nosplit
func (self class) InsertBreakpoint(line gd.Int, source gd.StringName) { //gd:EngineDebugger.insert_breakpoint
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, pointers.Get(source))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_insert_breakpoint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a breakpoint with the given [param source] and [param line].
*/
//go:nosplit
func (self class) RemoveBreakpoint(line gd.Int, source gd.StringName) { //gd:EngineDebugger.remove_breakpoint
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, pointers.Get(source))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_remove_breakpoint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears all breakpoints.
*/
//go:nosplit
func (self class) ClearBreakpoints() { //gd:EngineDebugger.clear_breakpoints
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EngineDebugger.Bind_clear_breakpoints, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("EngineDebugger", func(ptr gd.Object) any {
		return [1]gdclass.EngineDebugger{*(*gdclass.EngineDebugger)(unsafe.Pointer(&ptr))}
	})
}
