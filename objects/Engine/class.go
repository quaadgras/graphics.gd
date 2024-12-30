package Engine

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Dictionary"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
The [Engine] singleton allows you to query and modify the project's run-time parameters, such as frames per second, time scale, and others. It also stores information about the current build of Godot, such as the current version.
*/
var self objects.Engine
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.Engine)
	self = *(*objects.Engine)(unsafe.Pointer(&obj))
}

/*
Returns the fraction through the current physics tick we are at the time of rendering the frame. This can be used to implement fixed timestep interpolation.
*/
func GetPhysicsInterpolationFraction() Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetPhysicsInterpolationFraction()))
}

/*
Returns the total number of frames drawn since the engine started.
[b]Note:[/b] On headless platforms, or if rendering is disabled with [code]--disable-render-loop[/code] via command line, this method always returns [code]0[/code]. See also [method get_process_frames].
*/
func GetFramesDrawn() int {
	once.Do(singleton)
	return int(int(class(self).GetFramesDrawn()))
}

/*
Returns the average frames rendered every second (FPS), also known as the framerate.
*/
func GetFramesPerSecond() Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetFramesPerSecond()))
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
	once.Do(singleton)
	return int(int(class(self).GetProcessFrames()))
}

/*
Returns the instance of the [MainLoop]. This is usually the main [SceneTree] and is the same as [method Node.get_tree].
[b]Note:[/b] The type instantiated as the main loop can changed with [member ProjectSettings.application/run/main_loop_type].
*/
func GetMainLoop() objects.MainLoop {
	once.Do(singleton)
	return objects.MainLoop(class(self).GetMainLoop())
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
func GetVersionInfo() Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetVersionInfo())
}

/*
Returns the engine author information as a [Dictionary], where each entry is an [Array] of strings with the names of notable contributors to the Godot Engine: [code]lead_developers[/code], [code]founders[/code], [code]project_managers[/code], and [code]developers[/code].
*/
func GetAuthorInfo() Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetAuthorInfo())
}

/*
Returns an [Array] of dictionaries with copyright information for every component of Godot's source code.
Every [Dictionary] contains a [code]name[/code] identifier, and a [code]parts[/code] array of dictionaries. It describes the component in detail with the following entries:
- [code]files[/code] - [Array] of file paths from the source code affected by this component;
- [code]copyright[/code] - [Array] of owners of this component;
- [code]license[/code] - The license applied to this component (such as "[url=https://en.wikipedia.org/wiki/MIT_License#Ambiguity_and_variants]Expat[/url]" or "[url=https://creativecommons.org/licenses/by/4.0/]CC-BY-4.0[/url]").
*/
func GetCopyrightInfo() gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).GetCopyrightInfo())
}

/*
Returns a [Dictionary] of categorized donor names. Each entry is an [Array] of strings:
{[code]platinum_sponsors[/code], [code]gold_sponsors[/code], [code]silver_sponsors[/code], [code]bronze_sponsors[/code], [code]mini_sponsors[/code], [code]gold_donors[/code], [code]silver_donors[/code], [code]bronze_donors[/code]}
*/
func GetDonorInfo() Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetDonorInfo())
}

/*
Returns a [Dictionary] of licenses used by Godot and included third party components. Each entry is a license name (such as "[url=https://en.wikipedia.org/wiki/MIT_License#Ambiguity_and_variants]Expat[/url]") and its associated text.
*/
func GetLicenseInfo() Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetLicenseInfo())
}

/*
Returns the full Godot license text.
*/
func GetLicenseText() string {
	once.Do(singleton)
	return string(class(self).GetLicenseText().String())
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
	once.Do(singleton)
	return string(class(self).GetArchitectureName().String())
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
	once.Do(singleton)
	return bool(class(self).HasSingleton(gd.NewStringName(name)))
}

/*
Returns the global singleton with the given [param name], or [code]null[/code] if it does not exist. Often used for plugins. See also [method has_singleton] and [method get_singleton_list].
[b]Note:[/b] Global singletons are not the same as autoloaded nodes, which are configurable in the project settings.
*/
func GetSingleton(name string) gd.Object {
	once.Do(singleton)
	return gd.Object(class(self).GetSingleton(gd.NewStringName(name)))
}

/*
Registers the given [Object] [param instance] as a singleton, available globally under [param name]. Useful for plugins.
*/
func RegisterSingleton(name string, instance gd.Object) {
	once.Do(singleton)
	class(self).RegisterSingleton(gd.NewStringName(name), instance)
}

/*
Removes the singleton registered under [param name]. The singleton object is [i]not[/i] freed. Only works with user-defined singletons registered with [method register_singleton].
*/
func UnregisterSingleton(name string) {
	once.Do(singleton)
	class(self).UnregisterSingleton(gd.NewStringName(name))
}

/*
Returns a list of names of all available global singletons. See also [method get_singleton].
*/
func GetSingletonList() []string {
	once.Do(singleton)
	return []string(class(self).GetSingletonList().Strings())
}

/*
Registers a [ScriptLanguage] instance to be available with [code]ScriptServer[/code].
Returns:
- [constant OK] on success;
- [constant ERR_UNAVAILABLE] if [code]ScriptServer[/code] has reached the limit and cannot register any new language;
- [constant ERR_ALREADY_EXISTS] if [code]ScriptServer[/code] already contains a language with similar extension/name/type.
*/
func RegisterScriptLanguage(language objects.ScriptLanguage) error {
	once.Do(singleton)
	return error(class(self).RegisterScriptLanguage(language))
}

/*
Unregisters the [ScriptLanguage] instance from [code]ScriptServer[/code].
Returns:
- [constant OK] on success;
- [constant ERR_DOES_NOT_EXIST] if the language is not registered in [code]ScriptServer[/code].
*/
func UnregisterScriptLanguage(language objects.ScriptLanguage) error {
	once.Do(singleton)
	return error(class(self).UnregisterScriptLanguage(language))
}

/*
Returns the number of available script languages. Use with [method get_script_language].
*/
func GetScriptLanguageCount() int {
	once.Do(singleton)
	return int(int(class(self).GetScriptLanguageCount()))
}

/*
Returns an instance of a [ScriptLanguage] with the given [param index].
*/
func GetScriptLanguage(index int) objects.ScriptLanguage {
	once.Do(singleton)
	return objects.ScriptLanguage(class(self).GetScriptLanguage(gd.Int(index)))
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
	once.Do(singleton)
	return bool(class(self).IsEditorHint())
}

/*
Returns the path to the [MovieWriter]'s output file, or an empty string if the engine wasn't started in Movie Maker mode. The default path can be changed in [member ProjectSettings.editor/movie_writer/movie_file].
*/
func GetWriteMoviePath() string {
	once.Do(singleton)
	return string(class(self).GetWriteMoviePath().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.Engine

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

func PrintErrorMessages() bool {
	return bool(class(self).IsPrintingErrorMessages())
}

func SetPrintErrorMessages(value bool) {
	class(self).SetPrintErrorMessages(value)
}

func PhysicsTicksPerSecond() int {
	return int(int(class(self).GetPhysicsTicksPerSecond()))
}

func SetPhysicsTicksPerSecond(value int) {
	class(self).SetPhysicsTicksPerSecond(gd.Int(value))
}

func MaxPhysicsStepsPerFrame() int {
	return int(int(class(self).GetMaxPhysicsStepsPerFrame()))
}

func SetMaxPhysicsStepsPerFrame(value int) {
	class(self).SetMaxPhysicsStepsPerFrame(gd.Int(value))
}

func MaxFps() int {
	return int(int(class(self).GetMaxFps()))
}

func SetMaxFps(value int) {
	class(self).SetMaxFps(gd.Int(value))
}

func TimeScale() Float.X {
	return Float.X(Float.X(class(self).GetTimeScale()))
}

func SetTimeScale(value Float.X) {
	class(self).SetTimeScale(gd.Float(value))
}

func PhysicsJitterFix() Float.X {
	return Float.X(Float.X(class(self).GetPhysicsJitterFix()))
}

func SetPhysicsJitterFix(value Float.X) {
	class(self).SetPhysicsJitterFix(gd.Float(value))
}

//go:nosplit
func (self class) SetPhysicsTicksPerSecond(physics_ticks_per_second gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, physics_ticks_per_second)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_set_physics_ticks_per_second, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsTicksPerSecond() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_physics_ticks_per_second, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxPhysicsStepsPerFrame(max_physics_steps gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_physics_steps)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_set_max_physics_steps_per_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxPhysicsStepsPerFrame() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_max_physics_steps_per_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsJitterFix(physics_jitter_fix gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, physics_jitter_fix)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_set_physics_jitter_fix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsJitterFix() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_physics_jitter_fix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the fraction through the current physics tick we are at the time of rendering the frame. This can be used to implement fixed timestep interpolation.
*/
//go:nosplit
func (self class) GetPhysicsInterpolationFraction() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_physics_interpolation_fraction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxFps(max_fps gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_fps)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_set_max_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxFps() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_max_fps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTimeScale(time_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_set_time_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTimeScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_time_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_frames_drawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the average frames rendered every second (FPS), also known as the framerate.
*/
//go:nosplit
func (self class) GetFramesPerSecond() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_frames_per_second, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_physics_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_process_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the instance of the [MainLoop]. This is usually the main [SceneTree] and is the same as [method Node.get_tree].
[b]Note:[/b] The type instantiated as the main loop can changed with [member ProjectSettings.application/run/main_loop_type].
*/
//go:nosplit
func (self class) GetMainLoop() objects.MainLoop {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_main_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.MainLoop{classdb.MainLoop(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
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
func (self class) GetVersionInfo() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_version_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the engine author information as a [Dictionary], where each entry is an [Array] of strings with the names of notable contributors to the Godot Engine: [code]lead_developers[/code], [code]founders[/code], [code]project_managers[/code], and [code]developers[/code].
*/
//go:nosplit
func (self class) GetAuthorInfo() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_author_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
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
func (self class) GetCopyrightInfo() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_copyright_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a [Dictionary] of categorized donor names. Each entry is an [Array] of strings:
{[code]platinum_sponsors[/code], [code]gold_sponsors[/code], [code]silver_sponsors[/code], [code]bronze_sponsors[/code], [code]mini_sponsors[/code], [code]gold_donors[/code], [code]silver_donors[/code], [code]bronze_donors[/code]}
*/
//go:nosplit
func (self class) GetDonorInfo() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_donor_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a [Dictionary] of licenses used by Godot and included third party components. Each entry is a license name (such as "[url=https://en.wikipedia.org/wiki/MIT_License#Ambiguity_and_variants]Expat[/url]") and its associated text.
*/
//go:nosplit
func (self class) GetLicenseInfo() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_license_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the full Godot license text.
*/
//go:nosplit
func (self class) GetLicenseText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_license_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) GetArchitectureName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_architecture_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_is_in_physics_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_has_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the global singleton with the given [param name], or [code]null[/code] if it does not exist. Often used for plugins. See also [method has_singleton] and [method get_singleton_list].
[b]Note:[/b] Global singletons are not the same as autoloaded nodes, which are configurable in the project settings.
*/
//go:nosplit
func (self class) GetSingleton(name gd.StringName) gd.Object {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo(r_ret.Get())
	frame.Free()
	return ret
}

/*
Registers the given [Object] [param instance] as a singleton, available globally under [param name]. Useful for plugins.
*/
//go:nosplit
func (self class) RegisterSingleton(name gd.StringName, instance gd.Object) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(instance))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_register_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the singleton registered under [param name]. The singleton object is [i]not[/i] freed. Only works with user-defined singletons registered with [method register_singleton].
*/
//go:nosplit
func (self class) UnregisterSingleton(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_unregister_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of names of all available global singletons. See also [method get_singleton].
*/
//go:nosplit
func (self class) GetSingletonList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_singleton_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
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
func (self class) RegisterScriptLanguage(language objects.ScriptLanguage) error {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(language[0])))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_register_script_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) UnregisterScriptLanguage(language objects.ScriptLanguage) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language[0])[0])
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_unregister_script_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of available script languages. Use with [method get_script_language].
*/
//go:nosplit
func (self class) GetScriptLanguageCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_script_language_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an instance of a [ScriptLanguage] with the given [param index].
*/
//go:nosplit
func (self class) GetScriptLanguage(index gd.Int) objects.ScriptLanguage {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_script_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.ScriptLanguage{classdb.ScriptLanguage(gd.PointerMustAssertInstanceID(r_ret.Get()))}
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_is_editor_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the path to the [MovieWriter]'s output file, or an empty string if the engine wasn't started in Movie Maker mode. The default path can be changed in [member ProjectSettings.editor/movie_writer/movie_file].
*/
//go:nosplit
func (self class) GetWriteMoviePath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_get_write_movie_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPrintErrorMessages(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_set_print_error_messages, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsPrintingErrorMessages() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Engine.Bind_is_printing_error_messages, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() { classdb.Register("Engine", func(ptr gd.Object) any { return classdb.Engine(ptr) }) }

type Error int

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
