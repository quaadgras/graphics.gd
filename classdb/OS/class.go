// Package OS provides methods for working with OS object instances.
package OS

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
The [OS] class wraps the most common functionalities for communicating with the host operating system, such as the video driver, delays, environment variables, execution of binaries, command line, etc.
[b]Note:[/b] In Godot 4, [OS] functions related to window management, clipboard, and TTS were moved to the [DisplayServer] singleton (and the [Window] class). Functions related to time were removed and are only available in the [Time] class.
*/
var self [1]gdclass.OS
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.OS)
	self = *(*[1]gdclass.OS)(unsafe.Pointer(&obj))
}

/*
Generates a [PackedByteArray] of cryptographically secure random bytes with given [param size].
[b]Note:[/b] Generating large quantities of bytes using this method can result in locking and entropy of lower quality on most platforms. Using [method Crypto.generate_random_bytes] is preferred in most cases.
*/
func GetEntropy(size int) []byte {
	once.Do(singleton)
	return []byte(class(self).GetEntropy(gd.Int(size)).Bytes())
}

/*
Returns the list of certification authorities trusted by the operating system as a string of concatenated certificates in PEM format.
*/
func GetSystemCaCertificates() string {
	once.Do(singleton)
	return string(class(self).GetSystemCaCertificates().String())
}

/*
Returns an array of connected MIDI device names, if they exist. Returns an empty array if the system MIDI driver has not previously been initialized with [method open_midi_inputs]. See also [method close_midi_inputs].
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
*/
func GetConnectedMidiInputs() []string {
	once.Do(singleton)
	return []string(class(self).GetConnectedMidiInputs().Strings())
}

/*
Initializes the singleton for the system MIDI driver, allowing Godot to receive [InputEventMIDI]. See also [method get_connected_midi_inputs] and [method close_midi_inputs].
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
*/
func OpenMidiInputs() {
	once.Do(singleton)
	class(self).OpenMidiInputs()
}

/*
Shuts down the system MIDI driver. Godot will no longer receive [InputEventMIDI]. See also [method open_midi_inputs] and [method get_connected_midi_inputs].
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
*/
func CloseMidiInputs() {
	once.Do(singleton)
	class(self).CloseMidiInputs()
}

/*
Displays a modal dialog box using the host platform's implementation. The engine execution is blocked until the dialog is closed.
*/
func Alert(text string) {
	once.Do(singleton)
	class(self).Alert(gd.NewString(text), gd.NewString("Alert!"))
}

/*
Crashes the engine (or the editor if called within a [code]@tool[/code] script). See also [method kill].
[b]Note:[/b] This method should [i]only[/i] be used for testing the system's crash handler, not for any other purpose. For general error reporting, use (in order of preference) [method @GDScript.assert], [method @GlobalScope.push_error], or [method alert].
*/
func Crash(message string) {
	once.Do(singleton)
	class(self).Crash(gd.NewString(message))
}

/*
Returns the number of [i]logical[/i] CPU cores available on the host machine. On CPUs with HyperThreading enabled, this number will be greater than the number of [i]physical[/i] CPU cores.
*/
func GetProcessorCount() int {
	once.Do(singleton)
	return int(int(class(self).GetProcessorCount()))
}

/*
Returns the full name of the CPU model on the host machine (e.g. [code]"Intel(R) Core(TM) i7-6700K CPU @ 4.00GHz"[/code]).
[b]Note:[/b] This method is only implemented on Windows, macOS, Linux and iOS. On Android and Web, [method get_processor_name] returns an empty string.
*/
func GetProcessorName() string {
	once.Do(singleton)
	return string(class(self).GetProcessorName().String())
}

/*
Returns the list of font family names available.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS and Windows.
*/
func GetSystemFonts() []string {
	once.Do(singleton)
	return []string(class(self).GetSystemFonts().Strings())
}

/*
Returns the path to the system font file with [param font_name] and style. Returns an empty string if no matching fonts found.
The following aliases can be used to request default fonts: "sans-serif", "serif", "monospace", "cursive", and "fantasy".
[b]Note:[/b] Returned font might have different style if the requested style is not available.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS and Windows.
*/
func GetSystemFontPath(font_name string) string {
	once.Do(singleton)
	return string(class(self).GetSystemFontPath(gd.NewString(font_name), gd.Int(400), gd.Int(100), false).String())
}

/*
Returns an array of the system substitute font file paths, which are similar to the font with [param font_name] and style for the specified text, locale, and script. Returns an empty array if no matching fonts found.
The following aliases can be used to request default fonts: "sans-serif", "serif", "monospace", "cursive", and "fantasy".
[b]Note:[/b] Depending on OS, it's not guaranteed that any of the returned fonts will be suitable for rendering specified text. Fonts should be loaded and checked in the order they are returned, and the first suitable one used.
[b]Note:[/b] Returned fonts might have different style if the requested style is not available or belong to a different font family.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS and Windows.
*/
func GetSystemFontPathForText(font_name string, text string) []string {
	once.Do(singleton)
	return []string(class(self).GetSystemFontPathForText(gd.NewString(font_name), gd.NewString(text), gd.NewString(""), gd.NewString(""), gd.Int(400), gd.Int(100), false).Strings())
}

/*
Returns the file path to the current engine executable.
[b]Note:[/b] On macOS, if you want to launch another instance of Godot, always use [method create_instance] instead of relying on the executable path.
*/
func GetExecutablePath() string {
	once.Do(singleton)
	return string(class(self).GetExecutablePath().String())
}

/*
Reads a user input string from the standard input (usually the terminal). This operation is [i]blocking[/i], which causes the window to freeze if [method read_string_from_stdin] is called on the main thread. The thread calling [method read_string_from_stdin] will block until the program receives a line break in standard input (usually by the user pressing [kbd]Enter[/kbd]).
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
[b]Note:[/b] On exported Windows builds, run the console wrapper executable to access the terminal. Otherwise, the standard input will not work correctly. If you need a single executable with console support, use a custom build compiled with the [code]windows_subsystem=console[/code] flag.
*/
func ReadStringFromStdin() string {
	once.Do(singleton)
	return string(class(self).ReadStringFromStdin().String())
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
func Execute(path string, arguments []string) int {
	once.Do(singleton)
	return int(int(class(self).Execute(gd.NewString(path), gd.NewPackedStringSlice(arguments), [1]Array.Any{}[0], false, false)))
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
func ExecuteWithPipe(path string, arguments []string) Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).ExecuteWithPipe(gd.NewString(path), gd.NewPackedStringSlice(arguments)))
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
func CreateProcess(path string, arguments []string) int {
	once.Do(singleton)
	return int(int(class(self).CreateProcess(gd.NewString(path), gd.NewPackedStringSlice(arguments), false)))
}

/*
Creates a new instance of Godot that runs independently. The [param arguments] are used in the given order and separated by a space.
If the process is successfully created, this method returns the new process' ID, which you can use to monitor the process (and potentially terminate it with [method kill]). If the process cannot be created, this method returns [code]-1[/code].
See [method create_process] if you wish to run a different process.
[b]Note:[/b] This method is implemented on Android, Linux, macOS and Windows.
*/
func CreateInstance(arguments []string) int {
	once.Do(singleton)
	return int(int(class(self).CreateInstance(gd.NewPackedStringSlice(arguments))))
}

/*
Kill (terminate) the process identified by the given process ID ([param pid]), such as the ID returned by [method execute] in non-blocking mode. See also [method crash].
[b]Note:[/b] This method can also be used to kill processes that were not spawned by the engine.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS and Windows.
*/
func Kill(pid int) error {
	once.Do(singleton)
	return error(class(self).Kill(gd.Int(pid)))
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
func ShellOpen(uri string) error {
	once.Do(singleton)
	return error(class(self).ShellOpen(gd.NewString(uri)))
}

/*
Requests the OS to open the file manager, navigate to the given [param file_or_dir_path] and select the target file or folder.
If [param open_folder] is [code]true[/code] and [param file_or_dir_path] is a valid directory path, the OS will open the file manager and navigate to the target folder without selecting anything.
Use [method ProjectSettings.globalize_path] to convert a [code]res://[/code] or [code]user://[/code] project path into a system path to use with this method.
[b]Note:[/b] This method is currently only implemented on Windows and macOS. On other platforms, it will fallback to [method shell_open] with a directory path of [param file_or_dir_path] prefixed with [code]file://[/code].
*/
func ShellShowInFileManager(file_or_dir_path string) error {
	once.Do(singleton)
	return error(class(self).ShellShowInFileManager(gd.NewString(file_or_dir_path), true))
}

/*
Returns [code]true[/code] if the child process ID ([param pid]) is still running or [code]false[/code] if it has terminated. [param pid] must be a valid ID generated from [method create_process].
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS, and Windows.
*/
func IsProcessRunning(pid int) bool {
	once.Do(singleton)
	return bool(class(self).IsProcessRunning(gd.Int(pid)))
}

/*
Returns the exit code of a spawned process once it has finished running (see [method is_process_running]).
Returns [code]-1[/code] if the [param pid] is not a PID of a spawned child process, the process is still running, or the method is not implemented for the current platform.
[b]Note:[/b] Returns [code]-1[/code] if the [param pid] is a macOS bundled app process.
[b]Note:[/b] This method is implemented on Android, Linux, macOS and Windows.
*/
func GetProcessExitCode(pid int) int {
	once.Do(singleton)
	return int(int(class(self).GetProcessExitCode(gd.Int(pid))))
}

/*
Returns the number used by the host machine to uniquely identify this application.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS, and Windows.
*/
func GetProcessId() int {
	once.Do(singleton)
	return int(int(class(self).GetProcessId()))
}

/*
Returns [code]true[/code] if the environment variable with the name [param variable] exists.
[b]Note:[/b] Double-check the casing of [param variable]. Environment variable names are case-sensitive on all platforms except Windows.
*/
func HasEnvironment(variable string) bool {
	once.Do(singleton)
	return bool(class(self).HasEnvironment(gd.NewString(variable)))
}

/*
Returns the value of the given environment variable, or an empty string if [param variable] doesn't exist.
[b]Note:[/b] Double-check the casing of [param variable]. Environment variable names are case-sensitive on all platforms except Windows.
[b]Note:[/b] On macOS, applications do not have access to shell environment variables.
*/
func GetEnvironment(variable string) string {
	once.Do(singleton)
	return string(class(self).GetEnvironment(gd.NewString(variable)).String())
}

/*
Sets the value of the environment variable [param variable] to [param value]. The environment variable will be set for the Godot process and any process executed with [method execute] after running [method set_environment]. The environment variable will [i]not[/i] persist to processes run after the Godot process was terminated.
[b]Note:[/b] Environment variable names are case-sensitive on all platforms except Windows. The [param variable] name cannot be empty or include the [code]=[/code] character. On Windows, there is a 32767 characters limit for the combined length of [param variable], [param value], and the [code]=[/code] and null terminator characters that will be registered in the environment block.
*/
func SetEnvironment(variable string, value string) {
	once.Do(singleton)
	class(self).SetEnvironment(gd.NewString(variable), gd.NewString(value))
}

/*
Removes the given environment variable from the current environment, if it exists. The [param variable] name cannot be empty or include the [code]=[/code] character. The environment variable will be removed for the Godot process and any process executed with [method execute] after running [method unset_environment]. The removal of the environment variable will [i]not[/i] persist to processes run after the Godot process was terminated.
[b]Note:[/b] Environment variable names are case-sensitive on all platforms except Windows.
*/
func UnsetEnvironment(variable string) {
	once.Do(singleton)
	class(self).UnsetEnvironment(gd.NewString(variable))
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
func GetName() string {
	once.Do(singleton)
	return string(class(self).GetName().String())
}

/*
Returns the name of the distribution for Linux and BSD platforms (e.g. "Ubuntu", "Manjaro", "OpenBSD", etc.).
Returns the same value as [method get_name] for stock Android ROMs, but attempts to return the custom ROM name for popular Android derivatives such as "LineageOS".
Returns the same value as [method get_name] for other platforms.
[b]Note:[/b] This method is not supported on the Web platform. It returns an empty string.
*/
func GetDistributionName() string {
	once.Do(singleton)
	return string(class(self).GetDistributionName().String())
}

/*
Returns the exact production and build version of the operating system. This is different from the branded version used in marketing. This helps to distinguish between different releases of operating systems, including minor versions, and insider and custom builds.
- For Windows, the major and minor version are returned, as well as the build number. For example, the returned string may look like [code]10.0.9926[/code] for a build of Windows 10, and it may look like [code]6.1.7601[/code] for a build of Windows 7 SP1.
- For rolling distributions, such as Arch Linux, an empty string is returned.
- For macOS and iOS, the major and minor version are returned, as well as the patch number.
- For Android, the SDK version and the incremental build number are returned. If it's a custom ROM, it attempts to return its version instead.
[b]Note:[/b] This method is not supported on the Web platform. It returns an empty string.
*/
func GetVersion() string {
	once.Do(singleton)
	return string(class(self).GetVersion().String())
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
func GetCmdlineArgs() []string {
	once.Do(singleton)
	return []string(class(self).GetCmdlineArgs().Strings())
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
func GetCmdlineUserArgs() []string {
	once.Do(singleton)
	return []string(class(self).GetCmdlineUserArgs().Strings())
}

/*
Returns the video adapter driver name and version for the user's currently active graphics card, as a [PackedStringArray]. See also [method RenderingServer.get_video_adapter_api_version].
The first element holds the driver name, such as [code]nvidia[/code], [code]amdgpu[/code], etc.
The second element holds the driver version. For example, on the [code]nvidia[/code] driver on a Linux/BSD platform, the version is in the format [code]510.85.02[/code]. For Windows, the driver's format is [code]31.0.15.1659[/code].
[b]Note:[/b] This method is only supported on Linux/BSD and Windows when not running in headless mode. On other platforms, it returns an empty array.
*/
func GetVideoAdapterDriverInfo() []string {
	once.Do(singleton)
	return []string(class(self).GetVideoAdapterDriverInfo().Strings())
}

/*
If [param restart] is [code]true[/code], restarts the project automatically when it is exited with [method SceneTree.quit] or [constant Node.NOTIFICATION_WM_CLOSE_REQUEST]. Command-line [param arguments] can be supplied. To restart the project with the same command line arguments as originally used to run the project, pass [method get_cmdline_args] as the value for [param arguments].
This method can be used to apply setting changes that require a restart. See also [method is_restart_on_exit_set] and [method get_restart_on_exit_arguments].
[b]Note:[/b] This method is only effective on desktop platforms, and only when the project isn't started from the editor. It will have no effect on mobile and Web platforms, or when the project is started from the editor.
[b]Note:[/b] If the project process crashes or is [i]killed[/i] by the user (by sending [code]SIGKILL[/code] instead of the usual [code]SIGTERM[/code]), the project won't restart automatically.
*/
func SetRestartOnExit(restart bool) {
	once.Do(singleton)
	class(self).SetRestartOnExit(restart, gd.NewPackedStringSlice([1][]string{}[0]))
}

/*
Returns [code]true[/code] if the project will automatically restart when it exits for any reason, [code]false[/code] otherwise. See also [method set_restart_on_exit] and [method get_restart_on_exit_arguments].
*/
func IsRestartOnExitSet() bool {
	once.Do(singleton)
	return bool(class(self).IsRestartOnExitSet())
}

/*
Returns the list of command line arguments that will be used when the project automatically restarts using [method set_restart_on_exit]. See also [method is_restart_on_exit_set].
*/
func GetRestartOnExitArguments() []string {
	once.Do(singleton)
	return []string(class(self).GetRestartOnExitArguments().Strings())
}

/*
Delays execution of the current thread by [param usec] microseconds. [param usec] must be greater than or equal to [code]0[/code]. Otherwise, [method delay_usec] does nothing and prints an error message.
[b]Note:[/b] [method delay_usec] is a [i]blocking[/i] way to delay code execution. To delay code execution in a non-blocking way, you may use [method SceneTree.create_timer]. Awaiting with a [SceneTreeTimer] delays the execution of code placed below the [code]await[/code] without affecting the rest of the project (or editor, for [EditorPlugin]s and [EditorScript]s).
[b]Note:[/b] When [method delay_usec] is called on the main thread, it will freeze the project and will prevent it from redrawing and registering input until the delay has passed. When using [method delay_usec] as part of an [EditorPlugin] or [EditorScript], it will freeze the editor but won't freeze the project if it is currently running (since the project is an independent child process).
*/
func DelayUsec(usec int) {
	once.Do(singleton)
	class(self).DelayUsec(gd.Int(usec))
}

/*
Delays execution of the current thread by [param msec] milliseconds. [param msec] must be greater than or equal to [code]0[/code]. Otherwise, [method delay_msec] does nothing and prints an error message.
[b]Note:[/b] [method delay_msec] is a [i]blocking[/i] way to delay code execution. To delay code execution in a non-blocking way, you may use [method SceneTree.create_timer]. Awaiting with [SceneTreeTimer] delays the execution of code placed below the [code]await[/code] without affecting the rest of the project (or editor, for [EditorPlugin]s and [EditorScript]s).
[b]Note:[/b] When [method delay_msec] is called on the main thread, it will freeze the project and will prevent it from redrawing and registering input until the delay has passed. When using [method delay_msec] as part of an [EditorPlugin] or [EditorScript], it will freeze the editor but won't freeze the project if it is currently running (since the project is an independent child process).
*/
func DelayMsec(msec int) {
	once.Do(singleton)
	class(self).DelayMsec(gd.Int(msec))
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
func GetLocale() string {
	once.Do(singleton)
	return string(class(self).GetLocale().String())
}

/*
Returns the host OS locale's 2 or 3-letter [url=https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes]language code[/url] as a string which should be consistent on all platforms. This is equivalent to extracting the [code]language[/code] part of the [method get_locale] string.
This can be used to narrow down fully specified locale strings to only the "common" language code, when you don't need the additional information about country code or variants. For example, for a French Canadian user with [code]fr_CA[/code] locale, this would return [code]fr[/code].
*/
func GetLocaleLanguage() string {
	once.Do(singleton)
	return string(class(self).GetLocaleLanguage().String())
}

/*
Returns the model name of the current device.
[b]Note:[/b] This method is implemented on Android and iOS. Returns [code]"GenericDevice"[/code] on unsupported platforms.
*/
func GetModelName() string {
	once.Do(singleton)
	return string(class(self).GetModelName().String())
}

/*
Returns [code]true[/code] if the [code]user://[/code] file system is persistent, that is, its state is the same after a player quits and starts the game again. Relevant to the Web platform, where this persistence may be unavailable.
*/
func IsUserfsPersistent() bool {
	once.Do(singleton)
	return bool(class(self).IsUserfsPersistent())
}

/*
Returns [code]true[/code] if the engine was executed with the [code]--verbose[/code] or [code]-v[/code] command line argument, or if [member ProjectSettings.debug/settings/stdout/verbose_stdout] is [code]true[/code]. See also [method @GlobalScope.print_verbose].
*/
func IsStdoutVerbose() bool {
	once.Do(singleton)
	return bool(class(self).IsStdoutVerbose())
}

/*
Returns [code]true[/code] if the Godot binary used to run the project is a [i]debug[/i] export template, or when running in the editor.
Returns [code]false[/code] if the Godot binary used to run the project is a [i]release[/i] export template.
[b]Note:[/b] To check whether the Godot binary used to run the project is an export template (debug or release), use [code]OS.has_feature("template")[/code] instead.
*/
func IsDebugBuild() bool {
	once.Do(singleton)
	return bool(class(self).IsDebugBuild())
}

/*
Returns the amount of static memory being used by the program in bytes. Only works in debug builds.
*/
func GetStaticMemoryUsage() int {
	once.Do(singleton)
	return int(int(class(self).GetStaticMemoryUsage()))
}

/*
Returns the maximum amount of static memory used. Only works in debug builds.
*/
func GetStaticMemoryPeakUsage() int {
	once.Do(singleton)
	return int(int(class(self).GetStaticMemoryPeakUsage()))
}

/*
Returns a [Dictionary] containing information about the current memory with the following entries:
- [code]"physical"[/code] - total amount of usable physical memory in bytes. This value can be slightly less than the actual physical memory amount, since it does not include memory reserved by the kernel and devices.
- [code]"free"[/code] - amount of physical memory, that can be immediately allocated without disk access or other costly operations, in bytes. The process might be able to allocate more physical memory, but this action will require moving inactive pages to disk, which can be expensive.
- [code]"available"[/code] - amount of memory that can be allocated without extending the swap file(s), in bytes. This value includes both physical memory and swap.
- [code]"stack"[/code] - size of the current thread stack in bytes.
[b]Note:[/b] Each entry's value may be [code]-1[/code] if it is unknown.
*/
func GetMemoryInfo() Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetMemoryInfo())
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
func MoveToTrash(path string) error {
	once.Do(singleton)
	return error(class(self).MoveToTrash(gd.NewString(path)))
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
func GetUserDataDir() string {
	once.Do(singleton)
	return string(class(self).GetUserDataDir().String())
}

/*
Returns the path to commonly used folders across different platforms, as defined by [param dir]. See the [enum SystemDir] constants for available locations.
[b]Note:[/b] This method is implemented on Android, Linux, macOS and Windows.
[b]Note:[/b] Shared storage is implemented on Android and allows to differentiate between app specific and shared directories, if [param shared_storage] is [code]true[/code]. Shared directories have additional restrictions on Android.
*/
func GetSystemDir(dir gdclass.OSSystemDir) string {
	once.Do(singleton)
	return string(class(self).GetSystemDir(dir, true).String())
}

/*
Returns the [i]global[/i] user configuration directory according to the operating system's standards.
On the Linux/BSD platform, this path can be overridden by setting the [code]XDG_CONFIG_HOME[/code] environment variable before starting the project. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] in the documentation for more information. See also [method get_cache_dir] and [method get_data_dir].
Not to be confused with [method get_user_data_dir], which returns the [i]project-specific[/i] user data path.
*/
func GetConfigDir() string {
	once.Do(singleton)
	return string(class(self).GetConfigDir().String())
}

/*
Returns the [i]global[/i] user data directory according to the operating system's standards.
On the Linux/BSD platform, this path can be overridden by setting the [code]XDG_DATA_HOME[/code] environment variable before starting the project. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] in the documentation for more information. See also [method get_cache_dir] and [method get_config_dir].
Not to be confused with [method get_user_data_dir], which returns the [i]project-specific[/i] user data path.
*/
func GetDataDir() string {
	once.Do(singleton)
	return string(class(self).GetDataDir().String())
}

/*
Returns the [i]global[/i] cache data directory according to the operating system's standards.
On the Linux/BSD platform, this path can be overridden by setting the [code]XDG_CACHE_HOME[/code] environment variable before starting the project. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] in the documentation for more information. See also [method get_config_dir] and [method get_data_dir].
Not to be confused with [method get_user_data_dir], which returns the [i]project-specific[/i] user data path.
*/
func GetCacheDir() string {
	once.Do(singleton)
	return string(class(self).GetCacheDir().String())
}

/*
Returns a string that is unique to the device.
[b]Note:[/b] This string may change without notice if the user reinstalls their operating system, upgrades it, or modifies their hardware. This means it should generally not be used to encrypt persistent data, as the data saved before an unexpected ID change would become inaccessible. The returned string may also be falsified using external programs, so do not rely on the string returned by this method for security purposes.
[b]Note:[/b] On Web, returns an empty string and generates an error, as this method cannot be implemented for security reasons.
*/
func GetUniqueId() string {
	once.Do(singleton)
	return string(class(self).GetUniqueId().String())
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
func GetKeycodeString(code Key) string {
	once.Do(singleton)
	return string(class(self).GetKeycodeString(code).String())
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
func IsKeycodeUnicode(code int) bool {
	once.Do(singleton)
	return bool(class(self).IsKeycodeUnicode(gd.Int(code)))
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
func FindKeycodeFromString(s string) Key {
	once.Do(singleton)
	return Key(class(self).FindKeycodeFromString(gd.NewString(s)))
}

/*
If [param enabled] is [code]true[/code], when opening a file for writing, a temporary file is used in its place. When closed, it is automatically applied to the target file.
This can useful when files may be opened by other applications, such as antiviruses, text editors, or even the Godot editor itself.
*/
func SetUseFileAccessSaveAndSwap(enabled bool) {
	once.Do(singleton)
	class(self).SetUseFileAccessSaveAndSwap(enabled)
}

/*
Assigns the given name to the current thread. Returns [constant ERR_UNAVAILABLE] if unavailable on the current platform.
*/
func SetThreadName(name string) error {
	once.Do(singleton)
	return error(class(self).SetThreadName(gd.NewString(name)))
}

/*
Returns the ID of the current thread. This can be used in logs to ease debugging of multi-threaded applications.
[b]Note:[/b] Thread IDs are not deterministic and may be reused across application restarts.
*/
func GetThreadCallerId() int {
	once.Do(singleton)
	return int(int(class(self).GetThreadCallerId()))
}

/*
Returns the ID of the main thread. See [method get_thread_caller_id].
[b]Note:[/b] Thread IDs are not deterministic and may be reused across application restarts.
*/
func GetMainThreadId() int {
	once.Do(singleton)
	return int(int(class(self).GetMainThreadId()))
}

/*
Returns [code]true[/code] if the feature for the given feature tag is supported in the currently running instance, depending on the platform, build, etc. Can be used to check whether you're currently running a debug build, on a certain platform or arch, etc. Refer to the [url=$DOCS_URL/tutorials/export/feature_tags.html]Feature Tags[/url] documentation for more details.
[b]Note:[/b] Tag names are case-sensitive.
[b]Note:[/b] On the Web platform, one of the following additional tags is defined to indicate the host platform: [code]web_android[/code], [code]web_ios[/code], [code]web_linuxbsd[/code], [code]web_macos[/code], or [code]web_windows[/code].
*/
func HasFeature(tag_name string) bool {
	once.Do(singleton)
	return bool(class(self).HasFeature(gd.NewString(tag_name)))
}

/*
Returns [code]true[/code] if the application is running in the sandbox.
[b]Note:[/b] This method is only implemented on macOS and Linux.
*/
func IsSandboxed() bool {
	once.Do(singleton)
	return bool(class(self).IsSandboxed())
}

/*
Requests permission from the OS for the given [param name]. Returns [code]true[/code] if the permission has been successfully granted.
[b]Note:[/b] This method is currently only implemented on Android, to specifically request permission for [code]"RECORD_AUDIO"[/code] by [code]AudioDriverOpenSL[/code].
*/
func RequestPermission(name string) bool {
	once.Do(singleton)
	return bool(class(self).RequestPermission(gd.NewString(name)))
}

/*
Requests [i]dangerous[/i] permissions from the OS. Returns [code]true[/code] if permissions have been successfully granted.
[b]Note:[/b] This method is only implemented on Android. Normal permissions are automatically granted at install time in Android applications.
*/
func RequestPermissions() bool {
	once.Do(singleton)
	return bool(class(self).RequestPermissions())
}

/*
On Android devices: Returns the list of dangerous permissions that have been granted.
On macOS: Returns the list of user selected folders accessible to the application (sandboxed applications only). Use the native file dialog to request folder access permission.
*/
func GetGrantedPermissions() []string {
	once.Do(singleton)
	return []string(class(self).GetGrantedPermissions().Strings())
}

/*
On macOS (sandboxed applications only), this function clears list of user selected folders accessible to the application.
*/
func RevokeGrantedPermissions() {
	once.Do(singleton)
	class(self).RevokeGrantedPermissions()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.OS

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

func LowProcessorUsageMode() bool {
	return bool(class(self).IsInLowProcessorUsageMode())
}

func SetLowProcessorUsageMode(value bool) {
	class(self).SetLowProcessorUsageMode(value)
}

func LowProcessorUsageModeSleepUsec() int {
	return int(int(class(self).GetLowProcessorUsageModeSleepUsec()))
}

func SetLowProcessorUsageModeSleepUsec(value int) {
	class(self).SetLowProcessorUsageModeSleepUsec(gd.Int(value))
}

func DeltaSmoothing() bool {
	return bool(class(self).IsDeltaSmoothingEnabled())
}

func SetDeltaSmoothing(value bool) {
	class(self).SetDeltaSmoothing(value)
}

/*
Generates a [PackedByteArray] of cryptographically secure random bytes with given [param size].
[b]Note:[/b] Generating large quantities of bytes using this method can result in locking and entropy of lower quality on most platforms. Using [method Crypto.generate_random_bytes] is preferred in most cases.
*/
//go:nosplit
func (self class) GetEntropy(size gd.Int) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_entropy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the list of certification authorities trusted by the operating system as a string of concatenated certificates in PEM format.
*/
//go:nosplit
func (self class) GetSystemCaCertificates() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_system_ca_certificates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of connected MIDI device names, if they exist. Returns an empty array if the system MIDI driver has not previously been initialized with [method open_midi_inputs]. See also [method close_midi_inputs].
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
*/
//go:nosplit
func (self class) GetConnectedMidiInputs() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_connected_midi_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Initializes the singleton for the system MIDI driver, allowing Godot to receive [InputEventMIDI]. See also [method get_connected_midi_inputs] and [method close_midi_inputs].
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
*/
//go:nosplit
func (self class) OpenMidiInputs() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_open_midi_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Shuts down the system MIDI driver. Godot will no longer receive [InputEventMIDI]. See also [method open_midi_inputs] and [method get_connected_midi_inputs].
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
*/
//go:nosplit
func (self class) CloseMidiInputs() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_close_midi_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Displays a modal dialog box using the host platform's implementation. The engine execution is blocked until the dialog is closed.
*/
//go:nosplit
func (self class) Alert(text gd.String, title gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, pointers.Get(title))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_alert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Crashes the engine (or the editor if called within a [code]@tool[/code] script). See also [method kill].
[b]Note:[/b] This method should [i]only[/i] be used for testing the system's crash handler, not for any other purpose. For general error reporting, use (in order of preference) [method @GDScript.assert], [method @GlobalScope.push_error], or [method alert].
*/
//go:nosplit
func (self class) Crash(message gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(message))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_crash, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetLowProcessorUsageMode(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_set_low_processor_usage_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsInLowProcessorUsageMode() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_is_in_low_processor_usage_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLowProcessorUsageModeSleepUsec(usec gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, usec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_set_low_processor_usage_mode_sleep_usec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLowProcessorUsageModeSleepUsec() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_low_processor_usage_mode_sleep_usec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeltaSmoothing(delta_smoothing_enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, delta_smoothing_enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_set_delta_smoothing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDeltaSmoothingEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_is_delta_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of [i]logical[/i] CPU cores available on the host machine. On CPUs with HyperThreading enabled, this number will be greater than the number of [i]physical[/i] CPU cores.
*/
//go:nosplit
func (self class) GetProcessorCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_processor_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the full name of the CPU model on the host machine (e.g. [code]"Intel(R) Core(TM) i7-6700K CPU @ 4.00GHz"[/code]).
[b]Note:[/b] This method is only implemented on Windows, macOS, Linux and iOS. On Android and Web, [method get_processor_name] returns an empty string.
*/
//go:nosplit
func (self class) GetProcessorName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_processor_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the list of font family names available.
[b]Note:[/b] This method is implemented on Android, iOS, Linux, macOS and Windows.
*/
//go:nosplit
func (self class) GetSystemFonts() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_system_fonts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
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
func (self class) GetSystemFontPath(font_name gd.String, weight gd.Int, stretch gd.Int, italic bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font_name))
	callframe.Arg(frame, weight)
	callframe.Arg(frame, stretch)
	callframe.Arg(frame, italic)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_system_font_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) GetSystemFontPathForText(font_name gd.String, text gd.String, locale gd.String, script gd.String, weight gd.Int, stretch gd.Int, italic bool) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font_name))
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, pointers.Get(locale))
	callframe.Arg(frame, pointers.Get(script))
	callframe.Arg(frame, weight)
	callframe.Arg(frame, stretch)
	callframe.Arg(frame, italic)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_system_font_path_for_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the file path to the current engine executable.
[b]Note:[/b] On macOS, if you want to launch another instance of Godot, always use [method create_instance] instead of relying on the executable path.
*/
//go:nosplit
func (self class) GetExecutablePath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_executable_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Reads a user input string from the standard input (usually the terminal). This operation is [i]blocking[/i], which causes the window to freeze if [method read_string_from_stdin] is called on the main thread. The thread calling [method read_string_from_stdin] will block until the program receives a line break in standard input (usually by the user pressing [kbd]Enter[/kbd]).
[b]Note:[/b] This method is implemented on Linux, macOS and Windows.
[b]Note:[/b] On exported Windows builds, run the console wrapper executable to access the terminal. Otherwise, the standard input will not work correctly. If you need a single executable with console support, use a custom build compiled with the [code]windows_subsystem=console[/code] flag.
*/
//go:nosplit
func (self class) ReadStringFromStdin() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_read_string_from_stdin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(arguments))
	callframe.Arg(frame, pointers.Get(output))
	callframe.Arg(frame, read_stderr)
	callframe.Arg(frame, open_console)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_execute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) ExecuteWithPipe(path gd.String, arguments gd.PackedStringArray) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(arguments))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_execute_with_pipe, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(arguments))
	callframe.Arg(frame, open_console)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_create_process, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(arguments))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_create_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) Kill(pid gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, pid)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_kill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) ShellOpen(uri gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(uri))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_shell_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) ShellShowInFileManager(file_or_dir_path gd.String, open_folder bool) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(file_or_dir_path))
	callframe.Arg(frame, open_folder)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_shell_show_in_file_manager, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_is_process_running, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_process_exit_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_process_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(variable))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_has_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetEnvironment(variable gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(variable))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the value of the environment variable [param variable] to [param value]. The environment variable will be set for the Godot process and any process executed with [method execute] after running [method set_environment]. The environment variable will [i]not[/i] persist to processes run after the Godot process was terminated.
[b]Note:[/b] Environment variable names are case-sensitive on all platforms except Windows. The [param variable] name cannot be empty or include the [code]=[/code] character. On Windows, there is a 32767 characters limit for the combined length of [param variable], [param value], and the [code]=[/code] and null terminator characters that will be registered in the environment block.
*/
//go:nosplit
func (self class) SetEnvironment(variable gd.String, value gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(variable))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the given environment variable from the current environment, if it exists. The [param variable] name cannot be empty or include the [code]=[/code] character. The environment variable will be removed for the Godot process and any process executed with [method execute] after running [method unset_environment]. The removal of the environment variable will [i]not[/i] persist to processes run after the Godot process was terminated.
[b]Note:[/b] Environment variable names are case-sensitive on all platforms except Windows.
*/
//go:nosplit
func (self class) UnsetEnvironment(variable gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(variable))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_unset_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) GetDistributionName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_distribution_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) GetVersion() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) GetCmdlineArgs() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_cmdline_args, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
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
func (self class) GetCmdlineUserArgs() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_cmdline_user_args, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
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
func (self class) GetVideoAdapterDriverInfo() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_video_adapter_driver_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
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
func (self class) SetRestartOnExit(restart bool, arguments gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, restart)
	callframe.Arg(frame, pointers.Get(arguments))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_set_restart_on_exit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the project will automatically restart when it exits for any reason, [code]false[/code] otherwise. See also [method set_restart_on_exit] and [method get_restart_on_exit_arguments].
*/
//go:nosplit
func (self class) IsRestartOnExitSet() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_is_restart_on_exit_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of command line arguments that will be used when the project automatically restarts using [method set_restart_on_exit]. See also [method is_restart_on_exit_set].
*/
//go:nosplit
func (self class) GetRestartOnExitArguments() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_restart_on_exit_arguments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Delays execution of the current thread by [param usec] microseconds. [param usec] must be greater than or equal to [code]0[/code]. Otherwise, [method delay_usec] does nothing and prints an error message.
[b]Note:[/b] [method delay_usec] is a [i]blocking[/i] way to delay code execution. To delay code execution in a non-blocking way, you may use [method SceneTree.create_timer]. Awaiting with a [SceneTreeTimer] delays the execution of code placed below the [code]await[/code] without affecting the rest of the project (or editor, for [EditorPlugin]s and [EditorScript]s).
[b]Note:[/b] When [method delay_usec] is called on the main thread, it will freeze the project and will prevent it from redrawing and registering input until the delay has passed. When using [method delay_usec] as part of an [EditorPlugin] or [EditorScript], it will freeze the editor but won't freeze the project if it is currently running (since the project is an independent child process).
*/
//go:nosplit
func (self class) DelayUsec(usec gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, usec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_delay_usec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Delays execution of the current thread by [param msec] milliseconds. [param msec] must be greater than or equal to [code]0[/code]. Otherwise, [method delay_msec] does nothing and prints an error message.
[b]Note:[/b] [method delay_msec] is a [i]blocking[/i] way to delay code execution. To delay code execution in a non-blocking way, you may use [method SceneTree.create_timer]. Awaiting with [SceneTreeTimer] delays the execution of code placed below the [code]await[/code] without affecting the rest of the project (or editor, for [EditorPlugin]s and [EditorScript]s).
[b]Note:[/b] When [method delay_msec] is called on the main thread, it will freeze the project and will prevent it from redrawing and registering input until the delay has passed. When using [method delay_msec] as part of an [EditorPlugin] or [EditorScript], it will freeze the editor but won't freeze the project if it is currently running (since the project is an independent child process).
*/
//go:nosplit
func (self class) DelayMsec(msec gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, msec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_delay_msec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetLocale() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the host OS locale's 2 or 3-letter [url=https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes]language code[/url] as a string which should be consistent on all platforms. This is equivalent to extracting the [code]language[/code] part of the [method get_locale] string.
This can be used to narrow down fully specified locale strings to only the "common" language code, when you don't need the additional information about country code or variants. For example, for a French Canadian user with [code]fr_CA[/code] locale, this would return [code]fr[/code].
*/
//go:nosplit
func (self class) GetLocaleLanguage() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_locale_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the model name of the current device.
[b]Note:[/b] This method is implemented on Android and iOS. Returns [code]"GenericDevice"[/code] on unsupported platforms.
*/
//go:nosplit
func (self class) GetModelName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_model_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [code]user://[/code] file system is persistent, that is, its state is the same after a player quits and starts the game again. Relevant to the Web platform, where this persistence may be unavailable.
*/
//go:nosplit
func (self class) IsUserfsPersistent() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_is_userfs_persistent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the engine was executed with the [code]--verbose[/code] or [code]-v[/code] command line argument, or if [member ProjectSettings.debug/settings/stdout/verbose_stdout] is [code]true[/code]. See also [method @GlobalScope.print_verbose].
*/
//go:nosplit
func (self class) IsStdoutVerbose() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_is_stdout_verbose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_is_debug_build, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the amount of static memory being used by the program in bytes. Only works in debug builds.
*/
//go:nosplit
func (self class) GetStaticMemoryUsage() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_static_memory_usage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the maximum amount of static memory used. Only works in debug builds.
*/
//go:nosplit
func (self class) GetStaticMemoryPeakUsage() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_static_memory_peak_usage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetMemoryInfo() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_memory_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
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
func (self class) MoveToTrash(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_move_to_trash, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetUserDataDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_user_data_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the path to commonly used folders across different platforms, as defined by [param dir]. See the [enum SystemDir] constants for available locations.
[b]Note:[/b] This method is implemented on Android, Linux, macOS and Windows.
[b]Note:[/b] Shared storage is implemented on Android and allows to differentiate between app specific and shared directories, if [param shared_storage] is [code]true[/code]. Shared directories have additional restrictions on Android.
*/
//go:nosplit
func (self class) GetSystemDir(dir gdclass.OSSystemDir, shared_storage bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, dir)
	callframe.Arg(frame, shared_storage)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_system_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [i]global[/i] user configuration directory according to the operating system's standards.
On the Linux/BSD platform, this path can be overridden by setting the [code]XDG_CONFIG_HOME[/code] environment variable before starting the project. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] in the documentation for more information. See also [method get_cache_dir] and [method get_data_dir].
Not to be confused with [method get_user_data_dir], which returns the [i]project-specific[/i] user data path.
*/
//go:nosplit
func (self class) GetConfigDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_config_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [i]global[/i] user data directory according to the operating system's standards.
On the Linux/BSD platform, this path can be overridden by setting the [code]XDG_DATA_HOME[/code] environment variable before starting the project. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] in the documentation for more information. See also [method get_cache_dir] and [method get_config_dir].
Not to be confused with [method get_user_data_dir], which returns the [i]project-specific[/i] user data path.
*/
//go:nosplit
func (self class) GetDataDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_data_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [i]global[/i] cache data directory according to the operating system's standards.
On the Linux/BSD platform, this path can be overridden by setting the [code]XDG_CACHE_HOME[/code] environment variable before starting the project. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] in the documentation for more information. See also [method get_config_dir] and [method get_data_dir].
Not to be confused with [method get_user_data_dir], which returns the [i]project-specific[/i] user data path.
*/
//go:nosplit
func (self class) GetCacheDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_cache_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a string that is unique to the device.
[b]Note:[/b] This string may change without notice if the user reinstalls their operating system, upgrades it, or modifies their hardware. This means it should generally not be used to encrypt persistent data, as the data saved before an unexpected ID change would become inaccessible. The returned string may also be falsified using external programs, so do not rely on the string returned by this method for security purposes.
[b]Note:[/b] On Web, returns an empty string and generates an error, as this method cannot be implemented for security reasons.
*/
//go:nosplit
func (self class) GetUniqueId() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) GetKeycodeString(code Key) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, code)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_keycode_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
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
	var frame = callframe.New()
	callframe.Arg(frame, code)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_is_keycode_unicode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) FindKeycodeFromString(s gd.String) Key {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(s))
	var r_ret = callframe.Ret[Key](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_find_keycode_from_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], when opening a file for writing, a temporary file is used in its place. When closed, it is automatically applied to the target file.
This can useful when files may be opened by other applications, such as antiviruses, text editors, or even the Godot editor itself.
*/
//go:nosplit
func (self class) SetUseFileAccessSaveAndSwap(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_set_use_file_access_save_and_swap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Assigns the given name to the current thread. Returns [constant ERR_UNAVAILABLE] if unavailable on the current platform.
*/
//go:nosplit
func (self class) SetThreadName(name gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_set_thread_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_thread_caller_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_main_thread_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tag_name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_is_sandboxed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_request_permission, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_request_permissions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
On Android devices: Returns the list of dangerous permissions that have been granted.
On macOS: Returns the list of user selected folders accessible to the application (sandboxed applications only). Use the native file dialog to request folder access permission.
*/
//go:nosplit
func (self class) GetGrantedPermissions() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_get_granted_permissions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
On macOS (sandboxed applications only), this function clears list of user selected folders accessible to the application.
*/
//go:nosplit
func (self class) RevokeGrantedPermissions() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OS.Bind_revoke_granted_permissions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("OS", func(ptr gd.Object) any { return [1]gdclass.OS{*(*gdclass.OS)(unsafe.Pointer(&ptr))} })
}

type RenderingDriver = gdclass.OSRenderingDriver

const (
	/*The Vulkan rendering driver. It requires Vulkan 1.0 support and automatically uses features from Vulkan 1.1 and 1.2 if available.*/
	RenderingDriverVulkan RenderingDriver = 0
	/*The OpenGL 3 rendering driver. It uses OpenGL 3.3 Core Profile on desktop platforms, OpenGL ES 3.0 on mobile devices, and WebGL 2.0 on Web.*/
	RenderingDriverOpengl3 RenderingDriver = 1
	/*The Direct3D 12 rendering driver.*/
	RenderingDriverD3d12 RenderingDriver = 2
)

type SystemDir = gdclass.OSSystemDir

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

type Key int

const (
	/*Enum value which doesn't correspond to any key. This is used to initialize [enum Key] properties with a generic state.*/
	KeyNone Key = 0
	/*Keycodes with this bit applied are non-printable.*/
	KeySpecial Key = 4194304
	/*Escape key.*/
	KeyEscape Key = 4194305
	/*Tab key.*/
	KeyTab Key = 4194306
	/*Shift + Tab key.*/
	KeyBacktab Key = 4194307
	/*Backspace key.*/
	KeyBackspace Key = 4194308
	/*Return key (on the main keyboard).*/
	KeyEnter Key = 4194309
	/*Enter key on the numeric keypad.*/
	KeyKpEnter Key = 4194310
	/*Insert key.*/
	KeyInsert Key = 4194311
	/*Delete key.*/
	KeyDelete Key = 4194312
	/*Pause key.*/
	KeyPause Key = 4194313
	/*Print Screen key.*/
	KeyPrint Key = 4194314
	/*System Request key.*/
	KeySysreq Key = 4194315
	/*Clear key.*/
	KeyClear Key = 4194316
	/*Home key.*/
	KeyHome Key = 4194317
	/*End key.*/
	KeyEnd Key = 4194318
	/*Left arrow key.*/
	KeyLeft Key = 4194319
	/*Up arrow key.*/
	KeyUp Key = 4194320
	/*Right arrow key.*/
	KeyRight Key = 4194321
	/*Down arrow key.*/
	KeyDown Key = 4194322
	/*Page Up key.*/
	KeyPageup Key = 4194323
	/*Page Down key.*/
	KeyPagedown Key = 4194324
	/*Shift key.*/
	KeyShift Key = 4194325
	/*Control key.*/
	KeyCtrl Key = 4194326
	/*Meta key.*/
	KeyMeta Key = 4194327
	/*Alt key.*/
	KeyAlt Key = 4194328
	/*Caps Lock key.*/
	KeyCapslock Key = 4194329
	/*Num Lock key.*/
	KeyNumlock Key = 4194330
	/*Scroll Lock key.*/
	KeyScrolllock Key = 4194331
	/*F1 key.*/
	KeyF1 Key = 4194332
	/*F2 key.*/
	KeyF2 Key = 4194333
	/*F3 key.*/
	KeyF3 Key = 4194334
	/*F4 key.*/
	KeyF4 Key = 4194335
	/*F5 key.*/
	KeyF5 Key = 4194336
	/*F6 key.*/
	KeyF6 Key = 4194337
	/*F7 key.*/
	KeyF7 Key = 4194338
	/*F8 key.*/
	KeyF8 Key = 4194339
	/*F9 key.*/
	KeyF9 Key = 4194340
	/*F10 key.*/
	KeyF10 Key = 4194341
	/*F11 key.*/
	KeyF11 Key = 4194342
	/*F12 key.*/
	KeyF12 Key = 4194343
	/*F13 key.*/
	KeyF13 Key = 4194344
	/*F14 key.*/
	KeyF14 Key = 4194345
	/*F15 key.*/
	KeyF15 Key = 4194346
	/*F16 key.*/
	KeyF16 Key = 4194347
	/*F17 key.*/
	KeyF17 Key = 4194348
	/*F18 key.*/
	KeyF18 Key = 4194349
	/*F19 key.*/
	KeyF19 Key = 4194350
	/*F20 key.*/
	KeyF20 Key = 4194351
	/*F21 key.*/
	KeyF21 Key = 4194352
	/*F22 key.*/
	KeyF22 Key = 4194353
	/*F23 key.*/
	KeyF23 Key = 4194354
	/*F24 key.*/
	KeyF24 Key = 4194355
	/*F25 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF25 Key = 4194356
	/*F26 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF26 Key = 4194357
	/*F27 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF27 Key = 4194358
	/*F28 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF28 Key = 4194359
	/*F29 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF29 Key = 4194360
	/*F30 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF30 Key = 4194361
	/*F31 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF31 Key = 4194362
	/*F32 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF32 Key = 4194363
	/*F33 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF33 Key = 4194364
	/*F34 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF34 Key = 4194365
	/*F35 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF35 Key = 4194366
	/*Multiply (*) key on the numeric keypad.*/
	KeyKpMultiply Key = 4194433
	/*Divide (/) key on the numeric keypad.*/
	KeyKpDivide Key = 4194434
	/*Subtract (-) key on the numeric keypad.*/
	KeyKpSubtract Key = 4194435
	/*Period (.) key on the numeric keypad.*/
	KeyKpPeriod Key = 4194436
	/*Add (+) key on the numeric keypad.*/
	KeyKpAdd Key = 4194437
	/*Number 0 on the numeric keypad.*/
	KeyKp0 Key = 4194438
	/*Number 1 on the numeric keypad.*/
	KeyKp1 Key = 4194439
	/*Number 2 on the numeric keypad.*/
	KeyKp2 Key = 4194440
	/*Number 3 on the numeric keypad.*/
	KeyKp3 Key = 4194441
	/*Number 4 on the numeric keypad.*/
	KeyKp4 Key = 4194442
	/*Number 5 on the numeric keypad.*/
	KeyKp5 Key = 4194443
	/*Number 6 on the numeric keypad.*/
	KeyKp6 Key = 4194444
	/*Number 7 on the numeric keypad.*/
	KeyKp7 Key = 4194445
	/*Number 8 on the numeric keypad.*/
	KeyKp8 Key = 4194446
	/*Number 9 on the numeric keypad.*/
	KeyKp9 Key = 4194447
	/*Context menu key.*/
	KeyMenu Key = 4194370
	/*Hyper key. (On Linux/X11 only).*/
	KeyHyper Key = 4194371
	/*Help key.*/
	KeyHelp Key = 4194373
	/*Media back key. Not to be confused with the Back button on an Android device.*/
	KeyBack Key = 4194376
	/*Media forward key.*/
	KeyForward Key = 4194377
	/*Media stop key.*/
	KeyStop Key = 4194378
	/*Media refresh key.*/
	KeyRefresh Key = 4194379
	/*Volume down key.*/
	KeyVolumedown Key = 4194380
	/*Mute volume key.*/
	KeyVolumemute Key = 4194381
	/*Volume up key.*/
	KeyVolumeup Key = 4194382
	/*Media play key.*/
	KeyMediaplay Key = 4194388
	/*Media stop key.*/
	KeyMediastop Key = 4194389
	/*Previous song key.*/
	KeyMediaprevious Key = 4194390
	/*Next song key.*/
	KeyMedianext Key = 4194391
	/*Media record key.*/
	KeyMediarecord Key = 4194392
	/*Home page key.*/
	KeyHomepage Key = 4194393
	/*Favorites key.*/
	KeyFavorites Key = 4194394
	/*Search key.*/
	KeySearch Key = 4194395
	/*Standby key.*/
	KeyStandby Key = 4194396
	/*Open URL / Launch Browser key.*/
	KeyOpenurl Key = 4194397
	/*Launch Mail key.*/
	KeyLaunchmail Key = 4194398
	/*Launch Media key.*/
	KeyLaunchmedia Key = 4194399
	/*Launch Shortcut 0 key.*/
	KeyLaunch0 Key = 4194400
	/*Launch Shortcut 1 key.*/
	KeyLaunch1 Key = 4194401
	/*Launch Shortcut 2 key.*/
	KeyLaunch2 Key = 4194402
	/*Launch Shortcut 3 key.*/
	KeyLaunch3 Key = 4194403
	/*Launch Shortcut 4 key.*/
	KeyLaunch4 Key = 4194404
	/*Launch Shortcut 5 key.*/
	KeyLaunch5 Key = 4194405
	/*Launch Shortcut 6 key.*/
	KeyLaunch6 Key = 4194406
	/*Launch Shortcut 7 key.*/
	KeyLaunch7 Key = 4194407
	/*Launch Shortcut 8 key.*/
	KeyLaunch8 Key = 4194408
	/*Launch Shortcut 9 key.*/
	KeyLaunch9 Key = 4194409
	/*Launch Shortcut A key.*/
	KeyLauncha Key = 4194410
	/*Launch Shortcut B key.*/
	KeyLaunchb Key = 4194411
	/*Launch Shortcut C key.*/
	KeyLaunchc Key = 4194412
	/*Launch Shortcut D key.*/
	KeyLaunchd Key = 4194413
	/*Launch Shortcut E key.*/
	KeyLaunche Key = 4194414
	/*Launch Shortcut F key.*/
	KeyLaunchf Key = 4194415
	/*"Globe" key on Mac / iPad keyboard.*/
	KeyGlobe Key = 4194416
	/*"On-screen keyboard" key on iPad keyboard.*/
	KeyKeyboard Key = 4194417
	/*英数 key on Mac keyboard.*/
	KeyJisEisu Key = 4194418
	/*かな key on Mac keyboard.*/
	KeyJisKana Key = 4194419
	/*Unknown key.*/
	KeyUnknown Key = 8388607
	/*Space key.*/
	KeySpace Key = 32
	/*! key.*/
	KeyExclam Key = 33
	/*" key.*/
	KeyQuotedbl Key = 34
	/*# key.*/
	KeyNumbersign Key = 35
	/*$ key.*/
	KeyDollar Key = 36
	/*% key.*/
	KeyPercent Key = 37
	/*& key.*/
	KeyAmpersand Key = 38
	/*' key.*/
	KeyApostrophe Key = 39
	/*( key.*/
	KeyParenleft Key = 40
	/*) key.*/
	KeyParenright Key = 41
	/** key.*/
	KeyAsterisk Key = 42
	/*+ key.*/
	KeyPlus Key = 43
	/*, key.*/
	KeyComma Key = 44
	/*- key.*/
	KeyMinus Key = 45
	/*. key.*/
	KeyPeriod Key = 46
	/*/ key.*/
	KeySlash Key = 47
	/*Number 0 key.*/
	Key0 Key = 48
	/*Number 1 key.*/
	Key1 Key = 49
	/*Number 2 key.*/
	Key2 Key = 50
	/*Number 3 key.*/
	Key3 Key = 51
	/*Number 4 key.*/
	Key4 Key = 52
	/*Number 5 key.*/
	Key5 Key = 53
	/*Number 6 key.*/
	Key6 Key = 54
	/*Number 7 key.*/
	Key7 Key = 55
	/*Number 8 key.*/
	Key8 Key = 56
	/*Number 9 key.*/
	Key9 Key = 57
	/*: key.*/
	KeyColon Key = 58
	/*; key.*/
	KeySemicolon Key = 59
	/*< key.*/
	KeyLess Key = 60
	/*= key.*/
	KeyEqual Key = 61
	/*> key.*/
	KeyGreater Key = 62
	/*? key.*/
	KeyQuestion Key = 63
	/*@ key.*/
	KeyAt Key = 64
	/*A key.*/
	KeyA Key = 65
	/*B key.*/
	KeyB Key = 66
	/*C key.*/
	KeyC Key = 67
	/*D key.*/
	KeyD Key = 68
	/*E key.*/
	KeyE Key = 69
	/*F key.*/
	KeyF Key = 70
	/*G key.*/
	KeyG Key = 71
	/*H key.*/
	KeyH Key = 72
	/*I key.*/
	KeyI Key = 73
	/*J key.*/
	KeyJ Key = 74
	/*K key.*/
	KeyK Key = 75
	/*L key.*/
	KeyL Key = 76
	/*M key.*/
	KeyM Key = 77
	/*N key.*/
	KeyN Key = 78
	/*O key.*/
	KeyO Key = 79
	/*P key.*/
	KeyP Key = 80
	/*Q key.*/
	KeyQ Key = 81
	/*R key.*/
	KeyR Key = 82
	/*S key.*/
	KeyS Key = 83
	/*T key.*/
	KeyT Key = 84
	/*U key.*/
	KeyU Key = 85
	/*V key.*/
	KeyV Key = 86
	/*W key.*/
	KeyW Key = 87
	/*X key.*/
	KeyX Key = 88
	/*Y key.*/
	KeyY Key = 89
	/*Z key.*/
	KeyZ Key = 90
	/*[ key.*/
	KeyBracketleft Key = 91
	/*\ key.*/
	KeyBackslash Key = 92
	/*] key.*/
	KeyBracketright Key = 93
	/*^ key.*/
	KeyAsciicircum Key = 94
	/*_ key.*/
	KeyUnderscore Key = 95
	/*` key.*/
	KeyQuoteleft Key = 96
	/*{ key.*/
	KeyBraceleft Key = 123
	/*| key.*/
	KeyBar Key = 124
	/*} key.*/
	KeyBraceright Key = 125
	/*~ key.*/
	KeyAsciitilde Key = 126
	/*¥ key.*/
	KeyYen Key = 165
	/*§ key.*/
	KeySection Key = 167
)
