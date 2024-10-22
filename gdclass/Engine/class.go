package Engine

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [Engine] singleton allows you to query and modify the project's run-time parameters, such as frames per second, time scale, and others. It also stores information about the current build of Godot, such as the current version.

*/
var self gdclass.Engine
var once sync.Once
func singleton() {
	gc := gd.Static
	obj := gc.API.Object.GetSingleton(gc, gc.API.Singletons.Engine)
	self = *(*gdclass.Engine)(unsafe.Pointer(&obj))
}

/*
Returns the fraction through the current physics tick we are at the time of rendering the frame. This can be used to implement fixed timestep interpolation.
*/
func GetPhysicsInterpolationFraction() float64 {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return float64(float64(class(self).GetPhysicsInterpolationFraction()))
}

/*
Returns the total number of frames drawn since the engine started.
[b]Note:[/b] On headless platforms, or if rendering is disabled with [code]--disable-render-loop[/code] via command line, this method always returns [code]0[/code]. See also [method get_process_frames].
*/
func GetFramesDrawn() int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).GetFramesDrawn()))
}

/*
Returns the average frames rendered every second (FPS), also known as the framerate.
*/
func GetFramesPerSecond() float64 {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return float64(float64(class(self).GetFramesPerSecond()))
}

/*
Returns the total number of frames passed since the engine started. This number is increased every [b]physics frame[/b]. See also [method get_process_frames].
This method can be used to run expensive logic less often without relying on a [Timer]:
[codeblocks]
[gdscript]
func _physics_process(_delta):
    if Engine.get_physics_frames() % 2 == 0:
        pass # Run expensive logic only once every 2 physics frames here.
[/gdscript]
[csharp]
public override void _PhysicsProcess(double delta)
{
    base._PhysicsProcess(delta);

    if (Engine.GetPhysicsFrames() % 2 == 0)
    {
        // Run expensive logic only once every 2 physics frames here.
    }
}
[/csharp]
[/codeblocks]
*/
func GetPhysicsFrames() int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).GetPhysicsFrames()))
}

/*
Returns the total number of frames passed since the engine started. This number is increased every [b]process frame[/b], regardless of whether the render loop is enabled. See also [method get_frames_drawn] and [method get_physics_frames].
This method can be used to run expensive logic less often without relying on a [Timer]:
[codeblocks]
[gdscript]
func _process(_delta):
    if Engine.get_process_frames() % 5 == 0:
        pass # Run expensive logic only once every 5 process (render) frames here.
[/gdscript]
[csharp]
public override void _Process(double delta)
{
    base._Process(delta);

    if (Engine.GetProcessFrames() % 5 == 0)
    {
        // Run expensive logic only once every 5 process (render) frames here.
    }
}
[/csharp]
[/codeblocks]
*/
func GetProcessFrames() int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).GetProcessFrames()))
}

/*
Returns the instance of the [MainLoop]. This is usually the main [SceneTree] and is the same as [method Node.get_tree].
[b]Note:[/b] The type instantiated as the main loop can changed with [member ProjectSettings.application/run/main_loop_type].
*/
func GetMainLoop() gdclass.MainLoop {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.MainLoop(class(self).GetMainLoop(gc))
}

/*
Returns the current engine version information as a [Dictionary] containing the following entries:
- [code]major[/code] - Major version number as an int;
- [code]minor[/code] - Minor version number as an int;
- [code]patch[/code] - Patch version number as an int;
- [code]hex[/code] - Full version encoded as a hexadecimal int with one byte (2 hex digits) per number (see example below);
- [code]status[/code] - Status (such as "beta", "rc1", "rc2", "stable", etc.) as a String;
- [code]build[/code] - Build name (e.g. "custom_build") as a String;
- [code]hash[/code] - Full Git commit hash as a String;
- [code]timestamp[/code] - Holds the Git commit date UNIX timestamp in seconds as an int, or [code]0[/code] if unavailable;
- [code]string[/code] - [code]major[/code], [code]minor[/code], [code]patch[/code], [code]status[/code], and [code]build[/code] in a single String.
The [code]hex[/code] value is encoded as follows, from left to right: one byte for the major, one byte for the minor, one byte for the patch version. For example, "3.1.12" would be [code]0x03010C[/code].
[b]Note:[/b] The [code]hex[/code] value is still an [int] internally, and printing it will give you its decimal representation, which is not particularly meaningful. Use hexadecimal literals for quick version comparisons from code:
[codeblocks]
[gdscript]
if Engine.get_version_info().hex >= 0x040100:
    pass # Do things specific to version 4.1 or later.
else:
    pass # Do things specific to versions before 4.1.
[/gdscript]
[csharp]
if ((int)Engine.GetVersionInfo()["hex"] >= 0x040100)
{
    // Do things specific to version 4.1 or later.
}
else
{
    // Do things specific to versions before 4.1.
}
[/csharp]
[/codeblocks]
*/
func GetVersionInfo() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Dictionary(class(self).GetVersionInfo(gc))
}

/*
Returns the engine author information as a [Dictionary], where each entry is an [Array] of strings with the names of notable contributors to the Godot Engine: [code]lead_developers[/code], [code]founders[/code], [code]project_managers[/code], and [code]developers[/code].
*/
func GetAuthorInfo() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Dictionary(class(self).GetAuthorInfo(gc))
}

/*
Returns an [Array] of dictionaries with copyright information for every component of Godot's source code.
Every [Dictionary] contains a [code]name[/code] identifier, and a [code]parts[/code] array of dictionaries. It describes the component in detail with the following entries:
- [code]files[/code] - [Array] of file paths from the source code affected by this component;
- [code]copyright[/code] - [Array] of owners of this component;
- [code]license[/code] - The license applied to this component (such as "[url=https://en.wikipedia.org/wiki/MIT_License#Ambiguity_and_variants]Expat[/url]" or "[url=https://creativecommons.org/licenses/by/4.0/]CC-BY-4.0[/url]").
*/
func GetCopyrightInfo() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.ArrayOf[gd.Dictionary](class(self).GetCopyrightInfo(gc))
}

/*
Returns a [Dictionary] of categorized donor names. Each entry is an [Array] of strings:
{[code]platinum_sponsors[/code], [code]gold_sponsors[/code], [code]silver_sponsors[/code], [code]bronze_sponsors[/code], [code]mini_sponsors[/code], [code]gold_donors[/code], [code]silver_donors[/code], [code]bronze_donors[/code]}
*/
func GetDonorInfo() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Dictionary(class(self).GetDonorInfo(gc))
}

/*
Returns a [Dictionary] of licenses used by Godot and included third party components. Each entry is a license name (such as "[url=https://en.wikipedia.org/wiki/MIT_License#Ambiguity_and_variants]Expat[/url]") and its associated text.
*/
func GetLicenseInfo() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Dictionary(class(self).GetLicenseInfo(gc))
}

/*
Returns the full Godot license text.
*/
func GetLicenseText() string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetLicenseText(gc).String())
}

/*
Returns the name of the CPU architecture the Godot binary was built for. Possible return values include [code]"x86_64"[/code], [code]"x86_32"[/code], [code]"arm64"[/code], [code]"arm32"[/code], [code]"rv64"[/code], [code]"riscv"[/code], [code]"ppc64"[/code], [code]"ppc"[/code], [code]"wasm64"[/code], and [code]"wasm32"[/code].
To detect whether the current build is 64-bit, you can use the fact that all 64-bit architecture names contain [code]64[/code] in their name:
[codeblocks]
[gdscript]
if "64" in Engine.get_architecture_name():
    print("Running a 64-bit build of Godot.")
else:
    print("Running a 32-bit build of Godot.")
[/gdscript]
[csharp]
if (Engine.GetArchitectureName().Contains("64"))
    GD.Print("Running a 64-bit build of Godot.");
else
    GD.Print("Running a 32-bit build of Godot.");
[/csharp]
[/codeblocks]
[b]Note:[/b] This method does [i]not[/i] return the name of the system's CPU architecture (like [method OS.get_processor_name]). For example, when running an [code]x86_32[/code] Godot binary on an [code]x86_64[/code] system, the returned value will still be [code]"x86_32"[/code].
*/
func GetArchitectureName() string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetArchitectureName(gc).String())
}

/*
Returns [code]true[/code] if the engine is inside the fixed physics process step of the main loop.
[codeblock]
func _enter_tree():
    # Depending on when the node is added to the tree,
    # prints either "true" or "false".
    print(Engine.is_in_physics_frame())

func _process(delta):
    print(Engine.is_in_physics_frame()) # Prints false

func _physics_process(delta):
    print(Engine.is_in_physics_frame()) # Prints true
[/codeblock]
*/
func IsInPhysicsFrame() bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).IsInPhysicsFrame())
}

/*
Returns [code]true[/code] if a singleton with the given [param name] exists in the global scope. See also [method get_singleton].
[codeblocks]
[gdscript]
print(Engine.has_singleton("OS"))          # Prints true
print(Engine.has_singleton("Engine"))      # Prints true
print(Engine.has_singleton("AudioServer")) # Prints true
print(Engine.has_singleton("Unknown"))     # Prints false
[/gdscript]
[csharp]
GD.Print(Engine.HasSingleton("OS"));          // Prints true
GD.Print(Engine.HasSingleton("Engine"));      // Prints true
GD.Print(Engine.HasSingleton("AudioServer")); // Prints true
GD.Print(Engine.HasSingleton("Unknown"));     // Prints false
[/csharp]
[/codeblocks]
[b]Note:[/b] Global singletons are not the same as autoloaded nodes, which are configurable in the project settings.
*/
func HasSingleton(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).HasSingleton(gc.StringName(name)))
}

/*
Returns the global singleton with the given [param name], or [code]null[/code] if it does not exist. Often used for plugins. See also [method has_singleton] and [method get_singleton_list].
[b]Note:[/b] Global singletons are not the same as autoloaded nodes, which are configurable in the project settings.
*/
func GetSingleton(name string) gd.Object {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Object(class(self).GetSingleton(gc, gc.StringName(name)))
}

/*
Registers the given [Object] [param instance] as a singleton, available globally under [param name]. Useful for plugins.
*/
func RegisterSingleton(name string, instance gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).RegisterSingleton(gc.StringName(name), instance)
}

/*
Removes the singleton registered under [param name]. The singleton object is [i]not[/i] freed. Only works with user-defined singletons registered with [method register_singleton].
*/
func UnregisterSingleton(name string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).UnregisterSingleton(gc.StringName(name))
}

/*
Returns a list of names of all available global singletons. See also [method get_singleton].
*/
func GetSingletonList() []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetSingletonList(gc).Strings(gc))
}

/*
Registers a [ScriptLanguage] instance to be available with [code]ScriptServer[/code].
Returns:
- [constant OK] on success;
- [constant ERR_UNAVAILABLE] if [code]ScriptServer[/code] has reached the limit and cannot register any new language;
- [constant ERR_ALREADY_EXISTS] if [code]ScriptServer[/code] already contains a language with similar extension/name/type.
*/
func RegisterScriptLanguage(language gdclass.ScriptLanguage) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Error(class(self).RegisterScriptLanguage(language))
}

/*
Unregisters the [ScriptLanguage] instance from [code]ScriptServer[/code].
Returns:
- [constant OK] on success;
- [constant ERR_DOES_NOT_EXIST] if the language is not registered in [code]ScriptServer[/code].
*/
func UnregisterScriptLanguage(language gdclass.ScriptLanguage) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.Error(class(self).UnregisterScriptLanguage(language))
}

/*
Returns the number of available script languages. Use with [method get_script_language].
*/
func GetScriptLanguageCount() int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).GetScriptLanguageCount()))
}

/*
Returns an instance of a [ScriptLanguage] with the given [param index].
*/
func GetScriptLanguage(index int) gdclass.ScriptLanguage {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.ScriptLanguage(class(self).GetScriptLanguage(gc, gd.Int(index)))
}

/*
Returns [code]true[/code] if the script is currently running inside the editor, otherwise returns [code]false[/code]. This is useful for [code]@tool[/code] scripts to conditionally draw editor helpers, or prevent accidentally running "game" code that would affect the scene state while in the editor:
[codeblocks]
[gdscript]
if Engine.is_editor_hint():
    draw_gizmos()
else:
    simulate_physics()
[/gdscript]
[csharp]
if (Engine.IsEditorHint())
    DrawGizmos();
else
    SimulatePhysics();
[/csharp]
[/codeblocks]
See [url=$DOCS_URL/tutorials/plugins/running_code_in_the_editor.html]Running code in the editor[/url] in the documentation for more information.
[b]Note:[/b] To detect whether the script is running on an editor [i]build[/i] (such as when pressing [kbd]F5[/kbd]), use [method OS.has_feature] with the [code]"editor"[/code] argument instead. [code]OS.has_feature("editor")[/code] evaluate to [code]true[/code] both when the script is running in the editor and when running the project from the editor, but returns [code]false[/code] when run from an exported project.
*/
func IsEditorHint() bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).IsEditorHint())
}

/*
Returns the path to the [MovieWriter]'s output file, or an empty string if the engine wasn't started in Movie Maker mode. The default path can be changed in [member ProjectSettings.editor/movie_writer/movie_file].
*/
func GetWriteMoviePath() string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetWriteMoviePath(gc).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.Engine
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func PrintErrorMessages() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPrintingErrorMessages())
}

func SetPrintErrorMessages(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPrintErrorMessages(value)
}

func PhysicsTicksPerSecond() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPhysicsTicksPerSecond()))
}

func SetPhysicsTicksPerSecond(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPhysicsTicksPerSecond(gd.Int(value))
}

func MaxPhysicsStepsPerFrame() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxPhysicsStepsPerFrame()))
}

func SetMaxPhysicsStepsPerFrame(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxPhysicsStepsPerFrame(gd.Int(value))
}

func MaxFps() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxFps()))
}

func SetMaxFps(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxFps(gd.Int(value))
}

func TimeScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTimeScale()))
}

func SetTimeScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTimeScale(gd.Float(value))
}

func PhysicsJitterFix() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPhysicsJitterFix()))
}

func SetPhysicsJitterFix(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPhysicsJitterFix(gd.Float(value))
}

//go:nosplit
func (self class) SetPhysicsTicksPerSecond(physics_ticks_per_second gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, physics_ticks_per_second)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_set_physics_ticks_per_second, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsTicksPerSecond() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_physics_ticks_per_second, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxPhysicsStepsPerFrame(max_physics_steps gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_physics_steps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_set_max_physics_steps_per_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxPhysicsStepsPerFrame() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_max_physics_steps_per_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPhysicsJitterFix(physics_jitter_fix gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, physics_jitter_fix)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_set_physics_jitter_fix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsJitterFix() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_physics_jitter_fix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the fraction through the current physics tick we are at the time of rendering the frame. This can be used to implement fixed timestep interpolation.
*/
//go:nosplit
func (self class) GetPhysicsInterpolationFraction() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_physics_interpolation_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxFps(max_fps gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_fps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_set_max_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxFps() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_max_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTimeScale(time_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, time_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_set_time_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTimeScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_time_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the total number of frames drawn since the engine started.
[b]Note:[/b] On headless platforms, or if rendering is disabled with [code]--disable-render-loop[/code] via command line, this method always returns [code]0[/code]. See also [method get_process_frames].
*/
//go:nosplit
func (self class) GetFramesDrawn() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_frames_drawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the average frames rendered every second (FPS), also known as the framerate.
*/
//go:nosplit
func (self class) GetFramesPerSecond() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_frames_per_second, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the total number of frames passed since the engine started. This number is increased every [b]physics frame[/b]. See also [method get_process_frames].
This method can be used to run expensive logic less often without relying on a [Timer]:
[codeblocks]
[gdscript]
func _physics_process(_delta):
    if Engine.get_physics_frames() % 2 == 0:
        pass # Run expensive logic only once every 2 physics frames here.
[/gdscript]
[csharp]
public override void _PhysicsProcess(double delta)
{
    base._PhysicsProcess(delta);

    if (Engine.GetPhysicsFrames() % 2 == 0)
    {
        // Run expensive logic only once every 2 physics frames here.
    }
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetPhysicsFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_physics_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the total number of frames passed since the engine started. This number is increased every [b]process frame[/b], regardless of whether the render loop is enabled. See also [method get_frames_drawn] and [method get_physics_frames].
This method can be used to run expensive logic less often without relying on a [Timer]:
[codeblocks]
[gdscript]
func _process(_delta):
    if Engine.get_process_frames() % 5 == 0:
        pass # Run expensive logic only once every 5 process (render) frames here.
[/gdscript]
[csharp]
public override void _Process(double delta)
{
    base._Process(delta);

    if (Engine.GetProcessFrames() % 5 == 0)
    {
        // Run expensive logic only once every 5 process (render) frames here.
    }
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetProcessFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_process_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the instance of the [MainLoop]. This is usually the main [SceneTree] and is the same as [method Node.get_tree].
[b]Note:[/b] The type instantiated as the main loop can changed with [member ProjectSettings.application/run/main_loop_type].
*/
//go:nosplit
func (self class) GetMainLoop(ctx gd.Lifetime) gdclass.MainLoop {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_main_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.MainLoop
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the current engine version information as a [Dictionary] containing the following entries:
- [code]major[/code] - Major version number as an int;
- [code]minor[/code] - Minor version number as an int;
- [code]patch[/code] - Patch version number as an int;
- [code]hex[/code] - Full version encoded as a hexadecimal int with one byte (2 hex digits) per number (see example below);
- [code]status[/code] - Status (such as "beta", "rc1", "rc2", "stable", etc.) as a String;
- [code]build[/code] - Build name (e.g. "custom_build") as a String;
- [code]hash[/code] - Full Git commit hash as a String;
- [code]timestamp[/code] - Holds the Git commit date UNIX timestamp in seconds as an int, or [code]0[/code] if unavailable;
- [code]string[/code] - [code]major[/code], [code]minor[/code], [code]patch[/code], [code]status[/code], and [code]build[/code] in a single String.
The [code]hex[/code] value is encoded as follows, from left to right: one byte for the major, one byte for the minor, one byte for the patch version. For example, "3.1.12" would be [code]0x03010C[/code].
[b]Note:[/b] The [code]hex[/code] value is still an [int] internally, and printing it will give you its decimal representation, which is not particularly meaningful. Use hexadecimal literals for quick version comparisons from code:
[codeblocks]
[gdscript]
if Engine.get_version_info().hex >= 0x040100:
    pass # Do things specific to version 4.1 or later.
else:
    pass # Do things specific to versions before 4.1.
[/gdscript]
[csharp]
if ((int)Engine.GetVersionInfo()["hex"] >= 0x040100)
{
    // Do things specific to version 4.1 or later.
}
else
{
    // Do things specific to versions before 4.1.
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetVersionInfo(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_version_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the engine author information as a [Dictionary], where each entry is an [Array] of strings with the names of notable contributors to the Godot Engine: [code]lead_developers[/code], [code]founders[/code], [code]project_managers[/code], and [code]developers[/code].
*/
//go:nosplit
func (self class) GetAuthorInfo(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_author_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an [Array] of dictionaries with copyright information for every component of Godot's source code.
Every [Dictionary] contains a [code]name[/code] identifier, and a [code]parts[/code] array of dictionaries. It describes the component in detail with the following entries:
- [code]files[/code] - [Array] of file paths from the source code affected by this component;
- [code]copyright[/code] - [Array] of owners of this component;
- [code]license[/code] - The license applied to this component (such as "[url=https://en.wikipedia.org/wiki/MIT_License#Ambiguity_and_variants]Expat[/url]" or "[url=https://creativecommons.org/licenses/by/4.0/]CC-BY-4.0[/url]").
*/
//go:nosplit
func (self class) GetCopyrightInfo(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_copyright_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Returns a [Dictionary] of categorized donor names. Each entry is an [Array] of strings:
{[code]platinum_sponsors[/code], [code]gold_sponsors[/code], [code]silver_sponsors[/code], [code]bronze_sponsors[/code], [code]mini_sponsors[/code], [code]gold_donors[/code], [code]silver_donors[/code], [code]bronze_donors[/code]}
*/
//go:nosplit
func (self class) GetDonorInfo(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_donor_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [Dictionary] of licenses used by Godot and included third party components. Each entry is a license name (such as "[url=https://en.wikipedia.org/wiki/MIT_License#Ambiguity_and_variants]Expat[/url]") and its associated text.
*/
//go:nosplit
func (self class) GetLicenseInfo(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_license_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the full Godot license text.
*/
//go:nosplit
func (self class) GetLicenseText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_license_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the CPU architecture the Godot binary was built for. Possible return values include [code]"x86_64"[/code], [code]"x86_32"[/code], [code]"arm64"[/code], [code]"arm32"[/code], [code]"rv64"[/code], [code]"riscv"[/code], [code]"ppc64"[/code], [code]"ppc"[/code], [code]"wasm64"[/code], and [code]"wasm32"[/code].
To detect whether the current build is 64-bit, you can use the fact that all 64-bit architecture names contain [code]64[/code] in their name:
[codeblocks]
[gdscript]
if "64" in Engine.get_architecture_name():
    print("Running a 64-bit build of Godot.")
else:
    print("Running a 32-bit build of Godot.")
[/gdscript]
[csharp]
if (Engine.GetArchitectureName().Contains("64"))
    GD.Print("Running a 64-bit build of Godot.");
else
    GD.Print("Running a 32-bit build of Godot.");
[/csharp]
[/codeblocks]
[b]Note:[/b] This method does [i]not[/i] return the name of the system's CPU architecture (like [method OS.get_processor_name]). For example, when running an [code]x86_32[/code] Godot binary on an [code]x86_64[/code] system, the returned value will still be [code]"x86_32"[/code].
*/
//go:nosplit
func (self class) GetArchitectureName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_architecture_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the engine is inside the fixed physics process step of the main loop.
[codeblock]
func _enter_tree():
    # Depending on when the node is added to the tree,
    # prints either "true" or "false".
    print(Engine.is_in_physics_frame())

func _process(delta):
    print(Engine.is_in_physics_frame()) # Prints false

func _physics_process(delta):
    print(Engine.is_in_physics_frame()) # Prints true
[/codeblock]
*/
//go:nosplit
func (self class) IsInPhysicsFrame() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_is_in_physics_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a singleton with the given [param name] exists in the global scope. See also [method get_singleton].
[codeblocks]
[gdscript]
print(Engine.has_singleton("OS"))          # Prints true
print(Engine.has_singleton("Engine"))      # Prints true
print(Engine.has_singleton("AudioServer")) # Prints true
print(Engine.has_singleton("Unknown"))     # Prints false
[/gdscript]
[csharp]
GD.Print(Engine.HasSingleton("OS"));          // Prints true
GD.Print(Engine.HasSingleton("Engine"));      // Prints true
GD.Print(Engine.HasSingleton("AudioServer")); // Prints true
GD.Print(Engine.HasSingleton("Unknown"));     // Prints false
[/csharp]
[/codeblocks]
[b]Note:[/b] Global singletons are not the same as autoloaded nodes, which are configurable in the project settings.
*/
//go:nosplit
func (self class) HasSingleton(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_has_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the global singleton with the given [param name], or [code]null[/code] if it does not exist. Often used for plugins. See also [method has_singleton] and [method get_singleton_list].
[b]Note:[/b] Global singletons are not the same as autoloaded nodes, which are configurable in the project settings.
*/
//go:nosplit
func (self class) GetSingleton(ctx gd.Lifetime, name gd.StringName) gd.Object {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gd.Object
	ret.SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Registers the given [Object] [param instance] as a singleton, available globally under [param name]. Useful for plugins.
*/
//go:nosplit
func (self class) RegisterSingleton(name gd.StringName, instance gd.Object)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.End(instance.AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_register_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the singleton registered under [param name]. The singleton object is [i]not[/i] freed. Only works with user-defined singletons registered with [method register_singleton].
*/
//go:nosplit
func (self class) UnregisterSingleton(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_unregister_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of names of all available global singletons. See also [method get_singleton].
*/
//go:nosplit
func (self class) GetSingletonList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_singleton_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Registers a [ScriptLanguage] instance to be available with [code]ScriptServer[/code].
Returns:
- [constant OK] on success;
- [constant ERR_UNAVAILABLE] if [code]ScriptServer[/code] has reached the limit and cannot register any new language;
- [constant ERR_ALREADY_EXISTS] if [code]ScriptServer[/code] already contains a language with similar extension/name/type.
*/
//go:nosplit
func (self class) RegisterScriptLanguage(language gdclass.ScriptLanguage) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(language[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_register_script_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Unregisters the [ScriptLanguage] instance from [code]ScriptServer[/code].
Returns:
- [constant OK] on success;
- [constant ERR_DOES_NOT_EXIST] if the language is not registered in [code]ScriptServer[/code].
*/
//go:nosplit
func (self class) UnregisterScriptLanguage(language gdclass.ScriptLanguage) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_unregister_script_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of available script languages. Use with [method get_script_language].
*/
//go:nosplit
func (self class) GetScriptLanguageCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_script_language_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an instance of a [ScriptLanguage] with the given [param index].
*/
//go:nosplit
func (self class) GetScriptLanguage(ctx gd.Lifetime, index gd.Int) gdclass.ScriptLanguage {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_script_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.ScriptLanguage
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the script is currently running inside the editor, otherwise returns [code]false[/code]. This is useful for [code]@tool[/code] scripts to conditionally draw editor helpers, or prevent accidentally running "game" code that would affect the scene state while in the editor:
[codeblocks]
[gdscript]
if Engine.is_editor_hint():
    draw_gizmos()
else:
    simulate_physics()
[/gdscript]
[csharp]
if (Engine.IsEditorHint())
    DrawGizmos();
else
    SimulatePhysics();
[/csharp]
[/codeblocks]
See [url=$DOCS_URL/tutorials/plugins/running_code_in_the_editor.html]Running code in the editor[/url] in the documentation for more information.
[b]Note:[/b] To detect whether the script is running on an editor [i]build[/i] (such as when pressing [kbd]F5[/kbd]), use [method OS.has_feature] with the [code]"editor"[/code] argument instead. [code]OS.has_feature("editor")[/code] evaluate to [code]true[/code] both when the script is running in the editor and when running the project from the editor, but returns [code]false[/code] when run from an exported project.
*/
//go:nosplit
func (self class) IsEditorHint() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_is_editor_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the path to the [MovieWriter]'s output file, or an empty string if the engine wasn't started in Movie Maker mode. The default path can be changed in [member ProjectSettings.editor/movie_writer/movie_file].
*/
//go:nosplit
func (self class) GetWriteMoviePath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_get_write_movie_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPrintErrorMessages(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_set_print_error_messages, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPrintingErrorMessages() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Engine.Bind_is_printing_error_messages, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("Engine", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
