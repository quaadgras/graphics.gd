package OS

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
The [OS] class wraps the most common functionalities for communicating with the host operating system, such as the video driver, delays, environment variables, execution of binaries, command line, etc.
[b]Note:[/b] In Godot 4, [OS] functions related to window management, clipboard, and TTS were moved to the [DisplayServer] singleton (and the [Window] class). Functions related to time were removed and are only available in the [Time] class.

*/
type Simple [1]classdb.OS
func (self Simple) GetEntropy(size int) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetEntropy(gc, gd.Int(size)).Bytes())
}
func (self Simple) GetSystemCaCertificates() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSystemCaCertificates(gc).String())
}
func (self Simple) GetConnectedMidiInputs() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetConnectedMidiInputs(gc))
}
func (self Simple) OpenMidiInputs() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).OpenMidiInputs()
}
func (self Simple) CloseMidiInputs() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CloseMidiInputs()
}
func (self Simple) Alert(text string, title string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Alert(gc.String(text), gc.String(title))
}
func (self Simple) Crash(message string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Crash(gc.String(message))
}
func (self Simple) SetLowProcessorUsageMode(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLowProcessorUsageMode(enable)
}
func (self Simple) IsInLowProcessorUsageMode() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInLowProcessorUsageMode())
}
func (self Simple) SetLowProcessorUsageModeSleepUsec(usec int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLowProcessorUsageModeSleepUsec(gd.Int(usec))
}
func (self Simple) GetLowProcessorUsageModeSleepUsec() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLowProcessorUsageModeSleepUsec()))
}
func (self Simple) SetDeltaSmoothing(delta_smoothing_enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDeltaSmoothing(delta_smoothing_enabled)
}
func (self Simple) IsDeltaSmoothingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDeltaSmoothingEnabled())
}
func (self Simple) GetProcessorCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetProcessorCount()))
}
func (self Simple) GetProcessorName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetProcessorName(gc).String())
}
func (self Simple) GetSystemFonts() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetSystemFonts(gc))
}
func (self Simple) GetSystemFontPath(font_name string, weight int, stretch int, italic bool) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSystemFontPath(gc, gc.String(font_name), gd.Int(weight), gd.Int(stretch), italic).String())
}
func (self Simple) GetSystemFontPathForText(font_name string, text string, locale string, script string, weight int, stretch int, italic bool) gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetSystemFontPathForText(gc, gc.String(font_name), gc.String(text), gc.String(locale), gc.String(script), gd.Int(weight), gd.Int(stretch), italic))
}
func (self Simple) GetExecutablePath() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetExecutablePath(gc).String())
}
func (self Simple) ReadStringFromStdin() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).ReadStringFromStdin(gc).String())
}
func (self Simple) Execute(path string, arguments gd.PackedStringArray, output gd.Array, read_stderr bool, open_console bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).Execute(gc.String(path), arguments, output, read_stderr, open_console)))
}
func (self Simple) ExecuteWithPipe(path string, arguments gd.PackedStringArray) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).ExecuteWithPipe(gc, gc.String(path), arguments))
}
func (self Simple) CreateProcess(path string, arguments gd.PackedStringArray, open_console bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).CreateProcess(gc.String(path), arguments, open_console)))
}
func (self Simple) CreateInstance(arguments gd.PackedStringArray) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).CreateInstance(arguments)))
}
func (self Simple) Kill(pid int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Kill(gd.Int(pid)))
}
func (self Simple) ShellOpen(uri string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ShellOpen(gc.String(uri)))
}
func (self Simple) ShellShowInFileManager(file_or_dir_path string, open_folder bool) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).ShellShowInFileManager(gc.String(file_or_dir_path), open_folder))
}
func (self Simple) IsProcessRunning(pid int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsProcessRunning(gd.Int(pid)))
}
func (self Simple) GetProcessExitCode(pid int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetProcessExitCode(gd.Int(pid))))
}
func (self Simple) GetProcessId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetProcessId()))
}
func (self Simple) HasEnvironment(variable string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasEnvironment(gc.String(variable)))
}
func (self Simple) GetEnvironment(variable string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetEnvironment(gc, gc.String(variable)).String())
}
func (self Simple) SetEnvironment(variable string, value string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnvironment(gc.String(variable), gc.String(value))
}
func (self Simple) UnsetEnvironment(variable string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnsetEnvironment(gc.String(variable))
}
func (self Simple) GetName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetName(gc).String())
}
func (self Simple) GetDistributionName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetDistributionName(gc).String())
}
func (self Simple) GetVersion() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetVersion(gc).String())
}
func (self Simple) GetCmdlineArgs() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetCmdlineArgs(gc))
}
func (self Simple) GetCmdlineUserArgs() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetCmdlineUserArgs(gc))
}
func (self Simple) GetVideoAdapterDriverInfo() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetVideoAdapterDriverInfo(gc))
}
func (self Simple) SetRestartOnExit(restart bool, arguments gd.PackedStringArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRestartOnExit(restart, arguments)
}
func (self Simple) IsRestartOnExitSet() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRestartOnExitSet())
}
func (self Simple) GetRestartOnExitArguments() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetRestartOnExitArguments(gc))
}
func (self Simple) DelayUsec(usec int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DelayUsec(gd.Int(usec))
}
func (self Simple) DelayMsec(msec int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DelayMsec(gd.Int(msec))
}
func (self Simple) GetLocale() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLocale(gc).String())
}
func (self Simple) GetLocaleLanguage() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLocaleLanguage(gc).String())
}
func (self Simple) GetModelName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetModelName(gc).String())
}
func (self Simple) IsUserfsPersistent() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUserfsPersistent())
}
func (self Simple) IsStdoutVerbose() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsStdoutVerbose())
}
func (self Simple) IsDebugBuild() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDebugBuild())
}
func (self Simple) GetStaticMemoryUsage() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetStaticMemoryUsage()))
}
func (self Simple) GetStaticMemoryPeakUsage() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetStaticMemoryPeakUsage()))
}
func (self Simple) GetMemoryInfo() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetMemoryInfo(gc))
}
func (self Simple) MoveToTrash(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).MoveToTrash(gc.String(path)))
}
func (self Simple) GetUserDataDir() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetUserDataDir(gc).String())
}
func (self Simple) GetSystemDir(dir classdb.OSSystemDir, shared_storage bool) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSystemDir(gc, dir, shared_storage).String())
}
func (self Simple) GetConfigDir() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetConfigDir(gc).String())
}
func (self Simple) GetDataDir() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetDataDir(gc).String())
}
func (self Simple) GetCacheDir() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCacheDir(gc).String())
}
func (self Simple) GetUniqueId() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetUniqueId(gc).String())
}
func (self Simple) GetKeycodeString(code gd.Key) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetKeycodeString(gc, code).String())
}
func (self Simple) IsKeycodeUnicode(code int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsKeycodeUnicode(gd.Int(code)))
}
func (self Simple) FindKeycodeFromString(s string) gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).FindKeycodeFromString(gc.String(s)))
}
func (self Simple) SetUseFileAccessSaveAndSwap(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseFileAccessSaveAndSwap(enabled)
}
func (self Simple) SetThreadName(name string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).SetThreadName(gc.String(name)))
}
func (self Simple) GetThreadCallerId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetThreadCallerId()))
}
func (self Simple) GetMainThreadId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMainThreadId()))
}
func (self Simple) HasFeature(tag_name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasFeature(gc.String(tag_name)))
}
func (self Simple) IsSandboxed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSandboxed())
}
func (self Simple) RequestPermission(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).RequestPermission(gc.String(name)))
}
func (self Simple) RequestPermissions() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).RequestPermissions())
}
func (self Simple) GetGrantedPermissions() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetGrantedPermissions(gc))
}
func (self Simple) RevokeGrantedPermissions() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RevokeGrantedPermissions()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OS
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Generates a [PackedByteArray] of cryptographically secure random bytes with given [param size].
[b]Note:[/b] Generating large quantities of bytes using this method can result in locking and entropy of lower quality on most platforms. Using [method Crypto.generate_random_bytes] is preferred in most cases.
*/
//go:nosplit
func (self class) GetEntropy(ctx gd.Lifetime, size gd.Int) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_entropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the list of certification authorities trusted by the operating system as a string of concatenated certificates in PEM format.
*/
//go:nosplit
func (self class) GetSystemCaCertificates(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_system_ca_certificates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of connected MIDI device names, if they exist. Returns an empty array if the system MIDI driver has not previously been initialized with [method open_midi_inputs]. See also [method close_midi_inputs].
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
*/
//go:nosplit
func (self class) GetConnectedMidiInputs(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_connected_midi_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Initializes the singleton for the system MIDI driver, allowing Godot to receive [InputEventMIDI]. See also [method get_connected_midi_inputs] and [method close_midi_inputs].
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
*/
//go:nosplit
func (self class) OpenMidiInputs()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_open_midi_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Shuts down the system MIDI driver. Godot will no longer receive [InputEventMIDI]. See also [method open_midi_inputs] and [method get_connected_midi_inputs].
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
*/
//go:nosplit
func (self class) CloseMidiInputs()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_close_midi_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Displays a modal dialog box using the host platform's implementation. The engine execution is blocked until the dialog is closed.
*/
//go:nosplit
func (self class) Alert(text gd.String, title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_alert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Crashes the engine (or the editor if called within a [code]@tool[/code] script). See also [method kill].
[b]Note:[/b] This method should [i]only[/i] be used for testing the system's crash handler, not for any other purpose. For general error reporting, use (in order of preference) [method @GDScript.assert], [method @GlobalScope.push_error], or [method alert].
*/
//go:nosplit
func (self class) Crash(message gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(message))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_crash, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetLowProcessorUsageMode(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_set_low_processor_usage_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsInLowProcessorUsageMode() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_is_in_low_processor_usage_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLowProcessorUsageModeSleepUsec(usec gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, usec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_set_low_processor_usage_mode_sleep_usec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLowProcessorUsageModeSleepUsec() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_low_processor_usage_mode_sleep_usec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeltaSmoothing(delta_smoothing_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delta_smoothing_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_set_delta_smoothing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDeltaSmoothingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_is_delta_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of [i]logical[/i] CPU cores available on the host machine. On CPUs with HyperThreading enabled, this number will be greater than the number of [i]physical[/i] CPU cores.
*/
//go:nosplit
func (self class) GetProcessorCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_processor_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the full name of the CPU model on the host machine (e.g. [code]"Intel(R) Core(TM) i7-6700K CPU @ 4.00GHz"[/code]).
[b]Note:[/b] This method is only implemented on Windows, macOS, Linux and iOS. On Android and Web, [method get_processor_name] returns an empty string.
*/
//go:nosplit
func (self class) GetProcessorName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_processor_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the list of font family names available.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS and Windows.
*/
//go:nosplit
func (self class) GetSystemFonts(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_system_fonts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the path to the system font file with [param font_name] and style. Returns an empty string if no matching fonts found.
The following aliases can be used to request default fonts: "sans-serif", "serif", "monospace", "cursive", and "fantasy".
[b]Note:[/b] Returned font might have different style if the requested style is not available.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS and Windows.
*/
//go:nosplit
func (self class) GetSystemFontPath(ctx gd.Lifetime, font_name gd.String, weight gd.Int, stretch gd.Int, italic bool) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font_name))
	callframe.Arg(frame, weight)
	callframe.Arg(frame, stretch)
	callframe.Arg(frame, italic)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_system_font_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of the system substitute font file paths, which are similar to the font with [param font_name] and style for the specified text, locale, and script. Returns an empty array if no matching fonts found.
The following aliases can be used to request default fonts: "sans-serif", "serif", "monospace", "cursive", and "fantasy".
[b]Note:[/b] Depending on OS, it's not guaranteed that any of the returned fonts will be suitable for rendering specified text. Fonts should be loaded and checked in the order they are returned, and the first suitable one used.
[b]Note:[/b] Returned fonts might have different style if the requested style is not available or belong to a different font family.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS and Windows.
*/
//go:nosplit
func (self class) GetSystemFontPathForText(ctx gd.Lifetime, font_name gd.String, text gd.String, locale gd.String, script gd.String, weight gd.Int, stretch gd.Int, italic bool) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font_name))
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, mmm.Get(locale))
	callframe.Arg(frame, mmm.Get(script))
	callframe.Arg(frame, weight)
	callframe.Arg(frame, stretch)
	callframe.Arg(frame, italic)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_system_font_path_for_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the file path to the current engine executable.
[b]Note:[/b] On macOS, if you want to launch another instance of Godot, always use [method create_instance] instead of relying on the executable path.
*/
//go:nosplit
func (self class) GetExecutablePath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_executable_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Reads a user input string from the standard input (usually the terminal). This operation is [i]blocking[/i], which causes the window to freeze if [method read_string_from_stdin] is called on the main thread. The thread calling [method read_string_from_stdin] will block until the program receives a line break in standard input (usually by the user pressing [kbd]Enter[/kbd]).
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
[b]Note:[/b] On exported Windows builds, run the console wrapper executable to access the terminal. Otherwise, the standard input will not work correctly. If you need a single executable with console support, use a custom build compiled with the [code]windows_subsystem=console[/code] flag.
*/
//go:nosplit
func (self class) ReadStringFromStdin(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_read_string_from_stdin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Executes the given process in a [i]blocking[/i] way. The file specified in [param path] must exist and be executable. The system path resolution will be used. The [param arguments] are used in the given order, separated by spaces, and wrapped in quotes.
If an [param output] array is provided, the complete shell output of the process is appended to [param output] as a single [String] element. If [param read_stderr] is [code]true[/code], the output to the standard error stream is also appended to the array.
On Windows, if [param open_console] is [code]true[/code] and the process is a console app, a new terminal window is opened.
This method returns the exit code of the command, or [code]-1[/code] if the process fails to execute.
[b]Note:[/b] The main thread will be blocked until the executed command terminates. Use [Thread] to create a separate thread that will not block the main thread, or use [method create_process] to create a completely independent process.
For example, to retrieve a list of the working directory's contents:
[codeblocks]
[gdscript]
var output = []
var exit_code = OS.execute("ls", ["-l", "/tmp"], output)
[/gdscript]
[csharp]
var output = new Godot.Collections.Array();
int exitCode = OS.Execute("ls", new string[] {"-l", "/tmp"}, output);
[/csharp]
[/codeblocks]
If you wish to access a shell built-in or execute a composite command, a platform-specific shell can be invoked. For example:
[codeblocks]
[gdscript]
var output = []
OS.execute("CMD.exe", ["/C", "cd %TEMP% && dir"], output)
[/gdscript]
[csharp]
var output = new Godot.Collections.Array();
OS.Execute("CMD.exe", new string[] {"/C", "cd %TEMP% && dir"}, output);
[/csharp]
[/codeblocks]
[b]Note:[/b] This method is implemented on Android, Linux, macOS, and Windows.
[b]Note:[/b] To execute a Windows command interpreter built-in command, specify [code]cmd.exe[/code] in [param path], [code]/c[/code] as the first argument, and the desired command as the second argument.
[b]Note:[/b] To execute a PowerShell built-in command, specify [code]powershell.exe[/code] in [param path], [code]-Command[/code] as the first argument, and the desired command as the second argument.
[b]Note:[/b] To execute a Unix shell built-in command, specify shell executable name in [param path], [code]-c[/code] as the first argument, and the desired command as the second argument.
[b]Note:[/b] On macOS, sandboxed applications are limited to run only embedded helper executables, specified during export.
[b]Note:[/b] On Android, system commands such as [code]dumpsys[/code] can only be run on a rooted device.
*/
//go:nosplit
func (self class) Execute(path gd.String, arguments gd.PackedStringArray, output gd.Array, read_stderr bool, open_console bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(arguments))
	callframe.Arg(frame, mmm.Get(output))
	callframe.Arg(frame, read_stderr)
	callframe.Arg(frame, open_console)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_execute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new process that runs independently of Godot with redirected IO. It will not terminate when Godot terminates. The path specified in [param path] must exist and be an executable file or macOS [code].app[/code] bundle. The path is resolved based on the current platform. The [param arguments] are used in the given order and separated by a space.
If the process cannot be created, this method returns an empty [Dictionary]. Otherwise, this method returns a [Dictionary] with the following keys:
- [code]"stdio"[/code] - [FileAccess] to access the process stdin and stdout pipes (read/write).
- [code]"stderr"[/code] - [FileAccess] to access the process stderr pipe (read only).
- [code]"pid"[/code] - Process ID as an [int], which you can use to monitor the process (and potentially terminate it with [method kill]).
[b]Note:[/b] This method is implemented on Android, Linux, macOS, and Windows.
[b]Note:[/b] To execute a Windows command interpreter built-in command, specify [code]cmd.exe[/code] in [param path], [code]/c[/code] as the first argument, and the desired command as the second argument.
[b]Note:[/b] To execute a PowerShell built-in command, specify [code]powershell.exe[/code] in [param path], [code]-Command[/code] as the first argument, and the desired command as the second argument.
[b]Note:[/b] To execute a Unix shell built-in command, specify shell executable name in [param path], [code]-c[/code] as the first argument, and the desired command as the second argument.
[b]Note:[/b] On macOS, sandboxed applications are limited to run only embedded helper executables, specified during export or system .app bundle, system .app bundles will ignore arguments.
*/
//go:nosplit
func (self class) ExecuteWithPipe(ctx gd.Lifetime, path gd.String, arguments gd.PackedStringArray) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(arguments))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_execute_with_pipe, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates a new process that runs independently of Godot. It will not terminate when Godot terminates. The path specified in [param path] must exist and be an executable file or macOS [code].app[/code] bundle. The path is resolved based on the current platform. The [param arguments] are used in the given order and separated by a space.
On Windows, if [param open_console] is [code]true[/code] and the process is a console app, a new terminal window will be opened.
If the process is successfully created, this method returns its process ID, which you can use to monitor the process (and potentially terminate it with [method kill]). Otherwise, this method returns [code]-1[/code].
For example, running another instance of the project:
[codeblocks]
[gdscript]
var pid = OS.create_process(OS.get_executable_path(), [])
[/gdscript]
[csharp]
var pid = OS.CreateProcess(OS.GetExecutablePath(), new string[] {});
[/csharp]
[/codeblocks]
See [method execute] if you wish to run an external command and retrieve the results.
[b]Note:[/b] This method is implemented on Android, Linux, macOS, and Windows.
[b]Note:[/b] On macOS, sandboxed applications are limited to run only embedded helper executables, specified during export or system .app bundle, system .app bundles will ignore arguments.
*/
//go:nosplit
func (self class) CreateProcess(path gd.String, arguments gd.PackedStringArray, open_console bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(arguments))
	callframe.Arg(frame, open_console)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_create_process, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new instance of Godot that runs independently. The [param arguments] are used in the given order and separated by a space.
If the process is successfully created, this method returns the new process' ID, which you can use to monitor the process (and potentially terminate it with [method kill]). If the process cannot be created, this method returns [code]-1[/code].
See [method create_process] if you wish to run a different process.
[b]Note:[/b] This method is implemented on Android, Linux, macOS and Windows.
*/
//go:nosplit
func (self class) CreateInstance(arguments gd.PackedStringArray) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(arguments))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_create_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Kill (terminate) the process identified by the given process ID ([param pid]), such as the ID returned by [method execute] in non-blocking mode. See also [method crash].
[b]Note:[/b] This method can also be used to kill processes that were not spawned by the engine.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS and Windows.
*/
//go:nosplit
func (self class) Kill(pid gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pid)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_kill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Requests the OS to open a resource identified by [param uri] with the most appropriate program. For example:
- [code]OS.shell_open("C:\\Users\name\Downloads")[/code] on Windows opens the file explorer at the user's Downloads folder.
- [code]OS.shell_open("https://godotengine.org")[/code] opens the default web browser on the official Godot website.
- [code]OS.shell_open("mailto:example@example.com")[/code] opens the default email client with the "To" field set to [code]example@example.com[/code]. See [url=https://datatracker.ietf.org/doc/html/rfc2368]RFC 2368 - The [code]mailto[/code] URL scheme[/url] for a list of fields that can be added.
Use [method ProjectSettings.globalize_path] to convert a [code]res://[/code] or [code]user://[/code] project path into a system path for use with this method.
[b]Note:[/b] Use [method String.uri_encode] to encode characters within URLs in a URL-safe, portable way. This is especially required for line breaks. Otherwise, [method shell_open] may not work correctly in a project exported to the Web platform.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux, macOS and Windows.
*/
//go:nosplit
func (self class) ShellOpen(uri gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(uri))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_shell_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Requests the OS to open the file manager, navigate to the given [param file_or_dir_path] and select the target file or folder.
If [param open_folder] is [code]true[/code] and [param file_or_dir_path] is a valid directory path, the OS will open the file manager and navigate to the target folder without selecting anything.
Use [method ProjectSettings.globalize_path] to convert a [code]res://[/code] or [code]user://[/code] project path into a system path to use with this method.
[b]Note:[/b] This method is currently only implemented on Windows and macOS. On other platforms, it will fallback to [method shell_open] with a directory path of [param file_or_dir_path] prefixed with [code]file://[/code].
*/
//go:nosplit
func (self class) ShellShowInFileManager(file_or_dir_path gd.String, open_folder bool) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file_or_dir_path))
	callframe.Arg(frame, open_folder)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_shell_show_in_file_manager, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the child process ID ([param pid]) is still running or [code]false[/code] if it has terminated. [param pid] must be a valid ID generated from [method create_process].
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS, and Windows.
*/
//go:nosplit
func (self class) IsProcessRunning(pid gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_is_process_running, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the exit code of a spawned process once it has finished running (see [method is_process_running]).
Returns [code]-1[/code] if the [param pid] is not a PID of a spawned child process, the process is still running, or the method is not implemented for the current platform.
[b]Note:[/b] Returns [code]-1[/code] if the [param pid] is a macOS bundled app process.
[b]Note:[/b] This method is implemented on Android, Linux, macOS and Windows.
*/
//go:nosplit
func (self class) GetProcessExitCode(pid gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_process_exit_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number used by the host machine to uniquely identify this application.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS, and Windows.
*/
//go:nosplit
func (self class) GetProcessId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_process_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the environment variable with the name [param variable] exists.
[b]Note:[/b] Double-check the casing of [param variable]. Environment variable names are case-sensitive on all platforms except Windows.
*/
//go:nosplit
func (self class) HasEnvironment(variable gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(variable))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_has_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the value of the given environment variable, or an empty string if [param variable] doesn't exist.
[b]Note:[/b] Double-check the casing of [param variable]. Environment variable names are case-sensitive on all platforms except Windows.
[b]Note:[/b] On macOS, applications do not have access to shell environment variables.
*/
//go:nosplit
func (self class) GetEnvironment(ctx gd.Lifetime, variable gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(variable))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the value of the environment variable [param variable] to [param value]. The environment variable will be set for the Godot process and any process executed with [method execute] after running [method set_environment]. The environment variable will [i]not[/i] persist to processes run after the Godot process was terminated.
[b]Note:[/b] Environment variable names are case-sensitive on all platforms except Windows. The [param variable] name cannot be empty or include the [code]=[/code] character. On Windows, there is a 32767 characters limit for the combined length of [param variable], [param value], and the [code]=[/code] and null terminator characters that will be registered in the environment block.
*/
//go:nosplit
func (self class) SetEnvironment(variable gd.String, value gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(variable))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given environment variable from the current environment, if it exists. The [param variable] name cannot be empty or include the [code]=[/code] character. The environment variable will be removed for the Godot process and any process executed with [method execute] after running [method unset_environment]. The removal of the environment variable will [i]not[/i] persist to processes run after the Godot process was terminated.
[b]Note:[/b] Environment variable names are case-sensitive on all platforms except Windows.
*/
//go:nosplit
func (self class) UnsetEnvironment(variable gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(variable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_unset_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the host platform.
- On Windows, this is [code]"Windows"[/code].
- On macOS, this is [code]"macOS"[/code].
- On Linux-based operating systems, this is [code]"Linux"[/code].
- On BSD-based operating systems, this is [code]"FreeBSD"[/code], [code]"NetBSD"[/code], [code]"OpenBSD"[/code], or [code]"BSD"[/code] as a fallback.
- On Android, this is [code]"Android"[/code].
- On iOS, this is [code]"iOS"[/code].
- On Web, this is [code]"Web"[/code].
[b]Note:[/b] Custom builds of the engine may support additional platforms, such as consoles, possibly returning other names.
[codeblocks]
[gdscript]
match OS.get_name():
    "Windows":
        print("Welcome to Windows!")
    "macOS":
        print("Welcome to macOS!")
    "Linux", "FreeBSD", "NetBSD", "OpenBSD", "BSD":
        print("Welcome to Linux/BSD!")
    "Android":
        print("Welcome to Android!")
    "iOS":
        print("Welcome to iOS!")
    "Web":
        print("Welcome to the Web!")
[/gdscript]
[csharp]
switch (OS.GetName())
{
    case "Windows":
        GD.Print("Welcome to Windows");
        break;
    case "macOS":
        GD.Print("Welcome to macOS!");
        break;
    case "Linux":
    case "FreeBSD":
    case "NetBSD":
    case "OpenBSD":
    case "BSD":
        GD.Print("Welcome to Linux/BSD!");
        break;
    case "Android":
        GD.Print("Welcome to Android!");
        break;
    case "iOS":
        GD.Print("Welcome to iOS!");
        break;
    case "Web":
        GD.Print("Welcome to the Web!");
        break;
}
[/csharp]
[/codeblocks]
[b]Note:[/b] On Web platforms, it is still possible to determine the host platform's OS with feature tags. See [method has_feature].
*/
//go:nosplit
func (self class) GetName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the distribution for Linux and BSD platforms (e.g. "Ubuntu", "Manjaro", "OpenBSD", etc.).
Returns the same value as [method get_name] for stock Android ROMs, but attempts to return the custom ROM name for popular Android derivatives such as "LineageOS".
Returns the same value as [method get_name] for other platforms.
[b]Note:[/b] This method is not supported on the Web platform. It returns an empty string.
*/
//go:nosplit
func (self class) GetDistributionName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_distribution_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the exact production and build version of the operating system. This is different from the branded version used in marketing. This helps to distinguish between different releases of operating systems, including minor versions, and insider and custom builds.
- For Windows, the major and minor version are returned, as well as the build number. For example, the returned string may look like [code]10.0.9926[/code] for a build of Windows 10, and it may look like [code]6.1.7601[/code] for a build of Windows 7 SP1.
- For rolling distributions, such as Arch Linux, an empty string is returned.
- For macOS and iOS, the major and minor version are returned, as well as the patch number.
- For Android, the SDK version and the incremental build number are returned. If it's a custom ROM, it attempts to return its version instead.
[b]Note:[/b] This method is not supported on the Web platform. It returns an empty string.
*/
//go:nosplit
func (self class) GetVersion(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the command-line arguments passed to the engine.
Command-line arguments can be written in any form, including both [code]--key value[/code] and [code]--key=value[/code] forms so they can be properly parsed, as long as custom command-line arguments do not conflict with engine arguments.
You can also incorporate environment variables using the [method get_environment] method.
You can set [member ProjectSettings.editor/run/main_run_args] to define command-line arguments to be passed by the editor when running the project.
Here's a minimal example on how to parse command-line arguments into a [Dictionary] using the [code]--key=value[/code] form for arguments:
[codeblocks]
[gdscript]
var arguments = {}
for argument in OS.get_cmdline_args():
    if argument.contains("="):
        var key_value = argument.split("=")
        arguments[key_value[0].trim_prefix("--")] = key_value[1]
    else:
        # Options without an argument will be present in the dictionary,
        # with the value set to an empty string.
        arguments[argument.trim_prefix("--")] = ""
[/gdscript]
[csharp]
var arguments = new Dictionary<string, string>();
foreach (var argument in OS.GetCmdlineArgs())
{
    if (argument.Contains('='))
    {
        string[] keyValue = argument.Split("=");
        arguments[keyValue[0].TrimPrefix("--")] = keyValue[1];
    }
    else
    {
        // Options without an argument will be present in the dictionary,
        // with the value set to an empty string.
        arguments[argument.TrimPrefix("--")] = "";
    }
}
[/csharp]
[/codeblocks]
[b]Note:[/b] Passing custom user arguments directly is not recommended, as the engine may discard or modify them. Instead, pass the standard UNIX double dash ([code]--[/code]) and then the custom arguments, which the engine will ignore by design. These can be read via [method get_cmdline_user_args].
*/
//go:nosplit
func (self class) GetCmdlineArgs(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_cmdline_args, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the command-line user arguments passed to the engine. User arguments are ignored by the engine and reserved for the user. They are passed after the double dash [code]--[/code] argument. [code]++[/code] may be used when [code]--[/code] is intercepted by another program (such as [code]startx[/code]).
[codeblock]
# Godot has been executed with the following command:
# godot --fullscreen -- --level=2 --hardcore

OS.get_cmdline_args()      # Returns ["--fullscreen", "--level=2", "--hardcore"]
OS.get_cmdline_user_args() # Returns ["--level=2", "--hardcore"]
[/codeblock]
To get all passed arguments, use [method get_cmdline_args].
*/
//go:nosplit
func (self class) GetCmdlineUserArgs(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_cmdline_user_args, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the video adapter driver name and version for the user's currently active graphics card, as a [PackedStringArray]. See also [method RenderingServer.get_video_adapter_api_version].
The first element holds the driver name, such as [code]nvidia[/code], [code]amdgpu[/code], etc.
The second element holds the driver version. For example, on the [code]nvidia[/code] driver on a Linux/BSD platform, the version is in the format [code]510.85.02[/code]. For Windows, the driver's format is [code]31.0.15.1659[/code].
[b]Note:[/b] This method is only supported on Linux/BSD and Windows when not running in headless mode. On other platforms, it returns an empty array.
*/
//go:nosplit
func (self class) GetVideoAdapterDriverInfo(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_video_adapter_driver_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
If [param restart] is [code]true[/code], restarts the project automatically when it is exited with [method SceneTree.quit] or [constant Node.NOTIFICATION_WM_CLOSE_REQUEST]. Command-line [param arguments] can be supplied. To restart the project with the same command line arguments as originally used to run the project, pass [method get_cmdline_args] as the value for [param arguments].
This method can be used to apply setting changes that require a restart. See also [method is_restart_on_exit_set] and [method get_restart_on_exit_arguments].
[b]Note:[/b] This method is only effective on desktop platforms, and only when the project isn't started from the editor. It will have no effect on mobile and Web platforms, or when the project is started from the editor.
[b]Note:[/b] If the project process crashes or is [i]killed[/i] by the user (by sending [code]SIGKILL[/code] instead of the usual [code]SIGTERM[/code]), the project won't restart automatically.
*/
//go:nosplit
func (self class) SetRestartOnExit(restart bool, arguments gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, restart)
	callframe.Arg(frame, mmm.Get(arguments))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_set_restart_on_exit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the project will automatically restart when it exits for any reason, [code]false[/code] otherwise. See also [method set_restart_on_exit] and [method get_restart_on_exit_arguments].
*/
//go:nosplit
func (self class) IsRestartOnExitSet() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_is_restart_on_exit_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of command line arguments that will be used when the project automatically restarts using [method set_restart_on_exit]. See also [method is_restart_on_exit_set].
*/
//go:nosplit
func (self class) GetRestartOnExitArguments(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_restart_on_exit_arguments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Delays execution of the current thread by [param usec] microseconds. [param usec] must be greater than or equal to [code]0[/code]. Otherwise, [method delay_usec] does nothing and prints an error message.
[b]Note:[/b] [method delay_usec] is a [i]blocking[/i] way to delay code execution. To delay code execution in a non-blocking way, you may use [method SceneTree.create_timer]. Awaiting with a [SceneTreeTimer] delays the execution of code placed below the [code]await[/code] without affecting the rest of the project (or editor, for [EditorPlugin]s and [EditorScript]s).
[b]Note:[/b] When [method delay_usec] is called on the main thread, it will freeze the project and will prevent it from redrawing and registering input until the delay has passed. When using [method delay_usec] as part of an [EditorPlugin] or [EditorScript], it will freeze the editor but won't freeze the project if it is currently running (since the project is an independent child process).
*/
//go:nosplit
func (self class) DelayUsec(usec gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, usec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_delay_usec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Delays execution of the current thread by [param msec] milliseconds. [param msec] must be greater than or equal to [code]0[/code]. Otherwise, [method delay_msec] does nothing and prints an error message.
[b]Note:[/b] [method delay_msec] is a [i]blocking[/i] way to delay code execution. To delay code execution in a non-blocking way, you may use [method SceneTree.create_timer]. Awaiting with [SceneTreeTimer] delays the execution of code placed below the [code]await[/code] without affecting the rest of the project (or editor, for [EditorPlugin]s and [EditorScript]s).
[b]Note:[/b] When [method delay_msec] is called on the main thread, it will freeze the project and will prevent it from redrawing and registering input until the delay has passed. When using [method delay_msec] as part of an [EditorPlugin] or [EditorScript], it will freeze the editor but won't freeze the project if it is currently running (since the project is an independent child process).
*/
//go:nosplit
func (self class) DelayMsec(msec gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_delay_msec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the host OS locale as a [String] of the form [code]language_Script_COUNTRY_VARIANT@extra[/code]. Every substring after [code]language[/code] is optional and may not exist.
- [code]language[/code] - 2 or 3-letter [url=https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes]language code[/url], in lower case.
- [code skip-lint]Script[/code] - 4-letter [url=https://en.wikipedia.org/wiki/ISO_15924]script code[/url], in title case.
- [code]COUNTRY[/code] - 2 or 3-letter [url=https://en.wikipedia.org/wiki/ISO_3166-1]country code[/url], in upper case.
- [code]VARIANT[/code] - language variant, region and sort order. The variant can have any number of underscored keywords.
- [code]extra[/code] - semicolon separated list of additional key words. This may include currency, calendar, sort order and numbering system information.
If you want only the language code and not the fully specified locale from the OS, you can use [method get_locale_language].
*/
//go:nosplit
func (self class) GetLocale(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the host OS locale's 2 or 3-letter [url=https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes]language code[/url] as a string which should be consistent on all platforms. This is equivalent to extracting the [code]language[/code] part of the [method get_locale] string.
This can be used to narrow down fully specified locale strings to only the "common" language code, when you don't need the additional information about country code or variants. For example, for a French Canadian user with [code]fr_CA[/code] locale, this would return [code]fr[/code].
*/
//go:nosplit
func (self class) GetLocaleLanguage(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_locale_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the model name of the current device.
[b]Note:[/b] This method is implemented on Android and iOS. Returns [code]"GenericDevice"[/code] on unsupported platforms.
*/
//go:nosplit
func (self class) GetModelName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_model_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the [code]user://[/code] file system is persistent, that is, its state is the same after a player quits and starts the game again. Relevant to the Web platform, where this persistence may be unavailable.
*/
//go:nosplit
func (self class) IsUserfsPersistent() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_is_userfs_persistent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the engine was executed with the [code]--verbose[/code] or [code]-v[/code] command line argument, or if [member ProjectSettings.debug/settings/stdout/verbose_stdout] is [code]true[/code]. See also [method @GlobalScope.print_verbose].
*/
//go:nosplit
func (self class) IsStdoutVerbose() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_is_stdout_verbose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the Godot binary used to run the project is a [i]debug[/i] export template, or when running in the editor.
Returns [code]false[/code] if the Godot binary used to run the project is a [i]release[/i] export template.
[b]Note:[/b] To check whether the Godot binary used to run the project is an export template (debug or release), use [code]OS.has_feature("template")[/code] instead.
*/
//go:nosplit
func (self class) IsDebugBuild() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_is_debug_build, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the amount of static memory being used by the program in bytes. Only works in debug builds.
*/
//go:nosplit
func (self class) GetStaticMemoryUsage() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_static_memory_usage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the maximum amount of static memory used. Only works in debug builds.
*/
//go:nosplit
func (self class) GetStaticMemoryPeakUsage() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_static_memory_peak_usage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [Dictionary] containing information about the current memory with the following entries:
- [code]"physical"[/code] - total amount of usable physical memory in bytes. This value can be slightly less than the actual physical memory amount, since it does not include memory reserved by the kernel and devices.
- [code]"free"[/code] - amount of physical memory, that can be immediately allocated without disk access or other costly operations, in bytes. The process might be able to allocate more physical memory, but this action will require moving inactive pages to disk, which can be expensive.
- [code]"available"[/code] - amount of memory that can be allocated without extending the swap file(s), in bytes. This value includes both physical memory and swap.
- [code]"stack"[/code] - size of the current thread stack in bytes.
[b]Note:[/b] Each entry's value may be [code]-1[/code] if it is unknown.
*/
//go:nosplit
func (self class) GetMemoryInfo(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_memory_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Moves the file or directory at the given [param path] to the system's recycle bin. See also [method DirAccess.remove].
The method takes only global paths, so you may need to use [method ProjectSettings.globalize_path]. Do not use it for files in [code]res://[/code] as it will not work in exported projects.
Returns [constant FAILED] if the file or directory cannot be found, or the system does not support this method.
[codeblocks]
[gdscript]
var file_to_remove = "user://slot1.save"
OS.move_to_trash(ProjectSettings.globalize_path(file_to_remove))
[/gdscript]
[csharp]
var fileToRemove = "user://slot1.save";
OS.MoveToTrash(ProjectSettings.GlobalizePath(fileToRemove));
[/csharp]
[/codeblocks]
[b]Note:[/b] This method is implemented on Android, Linux, macOS and Windows.
[b]Note:[/b] If the user has disabled the recycle bin on their system, the file will be permanently deleted instead.
*/
//go:nosplit
func (self class) MoveToTrash(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_move_to_trash, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the absolute directory path where user data is written (the [code]user://[/code] directory in Godot). The path depends on the project name and [member ProjectSettings.application/config/use_custom_user_dir].
- On Windows, this is [code]%AppData%\Godot\app_userdata\[project_name][/code], or [code]%AppData%\[custom_name][/code] if [code]use_custom_user_dir[/code] is set. [code]%AppData%[/code] expands to [code]%UserProfile%\AppData\Roaming[/code].
- On macOS, this is [code]~/Library/Application Support/Godot/app_userdata/[project_name][/code], or [code]~/Library/Application Support/[custom_name][/code] if [code]use_custom_user_dir[/code] is set.
- On Linux and BSD, this is [code]~/.local/share/godot/app_userdata/[project_name][/code], or [code]~/.local/share/[custom_name][/code] if [code]use_custom_user_dir[/code] is set.
- On Android and iOS, this is a sandboxed directory in either internal or external storage, depending on the user's configuration.
- On Web, this is a virtual directory managed by the browser.
If the project name is empty, [code][project_name][/code] falls back to [code][unnamed project][/code].
Not to be confused with [method get_data_dir], which returns the [i]global[/i] (non-project-specific) user home directory.
*/
//go:nosplit
func (self class) GetUserDataDir(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_user_data_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the path to commonly used folders across different platforms, as defined by [param dir]. See the [enum SystemDir] constants for available locations.
[b]Note:[/b] This method is implemented on Android, Linux, macOS and Windows.
[b]Note:[/b] Shared storage is implemented on Android and allows to differentiate between app specific and shared directories, if [param shared_storage] is [code]true[/code]. Shared directories have additional restrictions on Android.
*/
//go:nosplit
func (self class) GetSystemDir(ctx gd.Lifetime, dir classdb.OSSystemDir, shared_storage bool) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, dir)
	callframe.Arg(frame, shared_storage)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_system_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [i]global[/i] user configuration directory according to the operating system's standards.
On the Linux/BSD platform, this path can be overridden by setting the [code]XDG_CONFIG_HOME[/code] environment variable before starting the project. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] in the documentation for more information. See also [method get_cache_dir] and [method get_data_dir].
Not to be confused with [method get_user_data_dir], which returns the [i]project-specific[/i] user data path.
*/
//go:nosplit
func (self class) GetConfigDir(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_config_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [i]global[/i] user data directory according to the operating system's standards.
On the Linux/BSD platform, this path can be overridden by setting the [code]XDG_DATA_HOME[/code] environment variable before starting the project. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] in the documentation for more information. See also [method get_cache_dir] and [method get_config_dir].
Not to be confused with [method get_user_data_dir], which returns the [i]project-specific[/i] user data path.
*/
//go:nosplit
func (self class) GetDataDir(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_data_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [i]global[/i] cache data directory according to the operating system's standards.
On the Linux/BSD platform, this path can be overridden by setting the [code]XDG_CACHE_HOME[/code] environment variable before starting the project. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] in the documentation for more information. See also [method get_config_dir] and [method get_data_dir].
Not to be confused with [method get_user_data_dir], which returns the [i]project-specific[/i] user data path.
*/
//go:nosplit
func (self class) GetCacheDir(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_cache_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a string that is unique to the device.
[b]Note:[/b] This string may change without notice if the user reinstalls their operating system, upgrades it, or modifies their hardware. This means it should generally not be used to encrypt persistent data, as the data saved before an unexpected ID change would become inaccessible. The returned string may also be falsified using external programs, so do not rely on the string returned by this method for security purposes.
[b]Note:[/b] On Web, returns an empty string and generates an error, as this method cannot be implemented for security reasons.
*/
//go:nosplit
func (self class) GetUniqueId(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the given keycode as a [String].
[codeblocks]
[gdscript]
print(OS.get_keycode_string(KEY_C))                    # Prints "C"
print(OS.get_keycode_string(KEY_ESCAPE))               # Prints "Escape"
print(OS.get_keycode_string(KEY_MASK_SHIFT | KEY_TAB)) # Prints "Shift+Tab"
[/gdscript]
[csharp]
GD.Print(OS.GetKeycodeString(Key.C));                                    // Prints "C"
GD.Print(OS.GetKeycodeString(Key.Escape));                               // Prints "Escape"
GD.Print(OS.GetKeycodeString((Key)KeyModifierMask.MaskShift | Key.Tab)); // Prints "Shift+Tab"
[/csharp]
[/codeblocks]
See also [method find_keycode_from_string], [member InputEventKey.keycode], and [method InputEventKey.get_keycode_with_modifiers].
*/
//go:nosplit
func (self class) GetKeycodeString(ctx gd.Lifetime, code gd.Key) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, code)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_keycode_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the input keycode corresponds to a Unicode character. For a list of codes, see the [enum Key] constants.
[codeblocks]
[gdscript]
print(OS.is_keycode_unicode(KEY_G))      # Prints true
print(OS.is_keycode_unicode(KEY_KP_4))   # Prints true
print(OS.is_keycode_unicode(KEY_TAB))    # Prints false
print(OS.is_keycode_unicode(KEY_ESCAPE)) # Prints false
[/gdscript]
[csharp]
GD.Print(OS.IsKeycodeUnicode((long)Key.G));      // Prints true
GD.Print(OS.IsKeycodeUnicode((long)Key.Kp4));    // Prints true
GD.Print(OS.IsKeycodeUnicode((long)Key.Tab));    // Prints false
GD.Print(OS.IsKeycodeUnicode((long)Key.Escape)); // Prints false
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) IsKeycodeUnicode(code gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, code)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_is_keycode_unicode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Finds the keycode for the given string. The returned values are equivalent to the [enum Key] constants.
[codeblocks]
[gdscript]
print(OS.find_keycode_from_string("C"))         # Prints 67 (KEY_C)
print(OS.find_keycode_from_string("Escape"))    # Prints 4194305 (KEY_ESCAPE)
print(OS.find_keycode_from_string("Shift+Tab")) # Prints 37748738 (KEY_MASK_SHIFT | KEY_TAB)
print(OS.find_keycode_from_string("Unknown"))   # Prints 0 (KEY_NONE)
[/gdscript]
[csharp]
GD.Print(OS.FindKeycodeFromString("C"));         // Prints C (Key.C)
GD.Print(OS.FindKeycodeFromString("Escape"));    // Prints Escape (Key.Escape)
GD.Print(OS.FindKeycodeFromString("Shift+Tab")); // Prints 37748738 (KeyModifierMask.MaskShift | Key.Tab)
GD.Print(OS.FindKeycodeFromString("Unknown"));   // Prints None (Key.None)
[/csharp]
[/codeblocks]
See also [method get_keycode_string].
*/
//go:nosplit
func (self class) FindKeycodeFromString(s gd.String) gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_find_keycode_from_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enabled] is [code]true[/code], when opening a file for writing, a temporary file is used in its place. When closed, it is automatically applied to the target file.
This can useful when files may be opened by other applications, such as antiviruses, text editors, or even the Godot editor itself.
*/
//go:nosplit
func (self class) SetUseFileAccessSaveAndSwap(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_set_use_file_access_save_and_swap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Assigns the given name to the current thread. Returns [constant ERR_UNAVAILABLE] if unavailable on the current platform.
*/
//go:nosplit
func (self class) SetThreadName(name gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_set_thread_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ID of the current thread. This can be used in logs to ease debugging of multi-threaded applications.
[b]Note:[/b] Thread IDs are not deterministic and may be reused across application restarts.
*/
//go:nosplit
func (self class) GetThreadCallerId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_thread_caller_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ID of the main thread. See [method get_thread_caller_id].
[b]Note:[/b] Thread IDs are not deterministic and may be reused across application restarts.
*/
//go:nosplit
func (self class) GetMainThreadId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_main_thread_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the feature for the given feature tag is supported in the currently running instance, depending on the platform, build, etc. Can be used to check whether you're currently running a debug build, on a certain platform or arch, etc. Refer to the [url=$DOCS_URL/tutorials/export/feature_tags.html]Feature Tags[/url] documentation for more details.
[b]Note:[/b] Tag names are case-sensitive.
[b]Note:[/b] On the Web platform, one of the following additional tags is defined to indicate the host platform: [code]web_android[/code], [code]web_ios[/code], [code]web_linuxbsd[/code], [code]web_macos[/code], or [code]web_windows[/code].
*/
//go:nosplit
func (self class) HasFeature(tag_name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tag_name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the application is running in the sandbox.
[b]Note:[/b] This method is only implemented on macOS and Linux.
*/
//go:nosplit
func (self class) IsSandboxed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_is_sandboxed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Requests permission from the OS for the given [param name]. Returns [code]true[/code] if the permission has been successfully granted.
[b]Note:[/b] This method is currently only implemented on Android, to specifically request permission for [code]"RECORD_AUDIO"[/code] by [code]AudioDriverOpenSL[/code].
*/
//go:nosplit
func (self class) RequestPermission(name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_request_permission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Requests [i]dangerous[/i] permissions from the OS. Returns [code]true[/code] if permissions have been successfully granted.
[b]Note:[/b] This method is only implemented on Android. Normal permissions are automatically granted at install time in Android applications.
*/
//go:nosplit
func (self class) RequestPermissions() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_request_permissions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
On Android devices: Returns the list of dangerous permissions that have been granted.
On macOS: Returns the list of user selected folders accessible to the application (sandboxed applications only). Use the native file dialog to request folder access permission.
*/
//go:nosplit
func (self class) GetGrantedPermissions(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_get_granted_permissions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
On macOS (sandboxed applications only), this function clears list of user selected folders accessible to the application.
*/
//go:nosplit
func (self class) RevokeGrantedPermissions()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OS.Bind_revoke_granted_permissions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsOS() Expert { return self[0].AsOS() }


//go:nosplit
func (self Simple) AsOS() Simple { return self[0].AsOS() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("OS", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type RenderingDriver = classdb.OSRenderingDriver

const (
/*The Vulkan rendering driver. It requires Vulkan 1.0 support and automatically uses features from Vulkan 1.1 and 1.2 if available.*/
	RenderingDriverVulkan RenderingDriver = 0
/*The OpenGL 3 rendering driver. It uses OpenGL 3.3 Core Profile on desktop platforms, OpenGL ES 3.0 on mobile devices, and WebGL 2.0 on Web.*/
	RenderingDriverOpengl3 RenderingDriver = 1
/*The Direct3D 12 rendering driver.*/
	RenderingDriverD3d12 RenderingDriver = 2
)
type SystemDir = classdb.OSSystemDir

const (
/*Refers to the Desktop directory path.*/
	SystemDirDesktop SystemDir = 0
/*Refers to the DCIM (Digital Camera Images) directory path.*/
	SystemDirDcim SystemDir = 1
/*Refers to the Documents directory path.*/
	SystemDirDocuments SystemDir = 2
/*Refers to the Downloads directory path.*/
	SystemDirDownloads SystemDir = 3
/*Refers to the Movies (or Videos) directory path.*/
	SystemDirMovies SystemDir = 4
/*Refers to the Music directory path.*/
	SystemDirMusic SystemDir = 5
/*Refers to the Pictures directory path.*/
	SystemDirPictures SystemDir = 6
/*Refers to the Ringtones directory path.*/
	SystemDirRingtones SystemDir = 7
)
