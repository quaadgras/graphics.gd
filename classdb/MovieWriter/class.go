// Package MovieWriter provides methods for working with MovieWriter object instances.
package MovieWriter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Vector2i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Godot can record videos with non-real-time simulation. Like the [code]--fixed-fps[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url], this forces the reported [code]delta[/code] in [method Node._process] functions to be identical across frames, regardless of how long it actually took to render the frame. This can be used to record high-quality videos with perfect frame pacing regardless of your hardware's capabilities.
Godot has 2 built-in [MovieWriter]s:
- AVI container with MJPEG for video and uncompressed audio ([code].avi[/code] file extension). Lossy compression, medium file sizes, fast encoding. The lossy compression quality can be adjusted by changing [member ProjectSettings.editor/movie_writer/mjpeg_quality]. The resulting file can be viewed in most video players, but it must be converted to another format for viewing on the web or by Godot with [VideoStreamPlayer]. MJPEG does not support transparency. AVI output is currently limited to a file of 4 GB in size at most.
- PNG image sequence for video and WAV for audio ([code].png[/code] file extension). Lossless compression, large file sizes, slow encoding. Designed to be encoded to a video file with another tool such as [url=https://ffmpeg.org/]FFmpeg[/url] after recording. Transparency is currently not supported, even if the root viewport is set to be transparent.
If you need to encode to a different format or pipe a stream through third-party software, you can extend the [MovieWriter] class to create your own movie writers. This should typically be done using GDExtension for performance reasons.
[b]Editor usage:[/b] A default movie file path can be specified in [member ProjectSettings.editor/movie_writer/movie_file]. Alternatively, for running single scenes, a [code]movie_file[/code] metadata can be added to the root node, specifying the path to a movie file that will be used when recording that scene. Once a path is set, click the video reel icon in the top-right corner of the editor to enable Movie Maker mode, then run any scene as usual. The engine will start recording as soon as the splash screen is finished, and it will only stop recording when the engine quits. Click the video reel icon again to disable Movie Maker mode. Note that toggling Movie Maker mode does not affect project instances that are already running.
[b]Note:[/b] MovieWriter is available for use in both the editor and exported projects, but it is [i]not[/i] designed for use by end users to record videos while playing. Players wishing to record gameplay videos should install tools such as [url=https://obsproject.com/]OBS Studio[/url] or [url=https://www.maartenbaert.be/simplescreenrecorder/]SimpleScreenRecorder[/url] instead.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=MovieWriter)
*/
type Instance [1]gdclass.MovieWriter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMovieWriter() Instance
}
type Interface interface {
	//Called when the audio sample rate used for recording the audio is requested by the engine. The value returned must be specified in Hz. Defaults to 48000 Hz if [method _get_audio_mix_rate] is not overridden.
	GetAudioMixRate() int
	//Called when the audio speaker mode used for recording the audio is requested by the engine. This can affect the number of output channels in the resulting audio file/stream. Defaults to [constant AudioServer.SPEAKER_MODE_STEREO] if [method _get_audio_speaker_mode] is not overridden.
	GetAudioSpeakerMode() gdclass.AudioServerSpeakerMode
	//Called when the engine determines whether this [MovieWriter] is able to handle the file at [param path]. Must return [code]true[/code] if this [MovieWriter] is able to handle the given file path, [code]false[/code] otherwise. Typically, [method _handles_file] is overridden as follows to allow the user to record a file at any path with a given file extension:
	//[codeblock]
	//func _handles_file(path):
	//    # Allows specifying an output file with a `.mkv` file extension (case-insensitive),
	//    # either in the Project Settings or with the `--write-movie <path>` command line argument.
	//    return path.get_extension().to_lower() == "mkv"
	//[/codeblock]
	HandlesFile(path string) bool
	//Called once before the engine starts writing video and audio data. [param movie_size] is the width and height of the video to save. [param fps] is the number of frames per second specified in the project settings or using the [code]--fixed-fps <fps>[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url].
	WriteBegin(movie_size Vector2i.XY, fps int, base_path string) error
	//Called at the end of every rendered frame. The [param frame_image] and [param audio_frame_block] function arguments should be written to.
	WriteFrame(frame_image [1]gdclass.Image, audio_frame_block unsafe.Pointer) error
	//Called when the engine finishes writing. This occurs when the engine quits by pressing the window manager's close button, or when [method SceneTree.quit] is called.
	//[b]Note:[/b] Pressing [kbd]Ctrl + C[/kbd] on the terminal running the editor/project does [i]not[/i] result in [method _write_end] being called.
	WriteEnd()
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) GetAudioMixRate() (_ int)                                { return }
func (self Implementation) GetAudioSpeakerMode() (_ gdclass.AudioServerSpeakerMode) { return }
func (self Implementation) HandlesFile(path string) (_ bool)                        { return }
func (self Implementation) WriteBegin(movie_size Vector2i.XY, fps int, base_path string) (_ error) {
	return
}
func (self Implementation) WriteFrame(frame_image [1]gdclass.Image, audio_frame_block unsafe.Pointer) (_ error) {
	return
}
func (self Implementation) WriteEnd() { return }

/*
Called when the audio sample rate used for recording the audio is requested by the engine. The value returned must be specified in Hz. Defaults to 48000 Hz if [method _get_audio_mix_rate] is not overridden.
*/
func (Instance) _get_audio_mix_rate(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the audio speaker mode used for recording the audio is requested by the engine. This can affect the number of output channels in the resulting audio file/stream. Defaults to [constant AudioServer.SPEAKER_MODE_STEREO] if [method _get_audio_speaker_mode] is not overridden.
*/
func (Instance) _get_audio_speaker_mode(impl func(ptr unsafe.Pointer) gdclass.AudioServerSpeakerMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the engine determines whether this [MovieWriter] is able to handle the file at [param path]. Must return [code]true[/code] if this [MovieWriter] is able to handle the given file path, [code]false[/code] otherwise. Typically, [method _handles_file] is overridden as follows to allow the user to record a file at any path with a given file extension:
[codeblock]
func _handles_file(path):

	# Allows specifying an output file with a `.mkv` file extension (case-insensitive),
	# either in the Project Settings or with the `--write-movie <path>` command line argument.
	return path.get_extension().to_lower() == "mkv"

[/codeblock]
*/
func (Instance) _handles_file(impl func(ptr unsafe.Pointer, path string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called once before the engine starts writing video and audio data. [param movie_size] is the width and height of the video to save. [param fps] is the number of frames per second specified in the project settings or using the [code]--fixed-fps <fps>[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url].
*/
func (Instance) _write_begin(impl func(ptr unsafe.Pointer, movie_size Vector2i.XY, fps int, base_path string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var movie_size = gd.UnsafeGet[gd.Vector2i](p_args, 0)
		var fps = gd.UnsafeGet[gd.Int](p_args, 1)
		var base_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(base_path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, movie_size, int(fps), base_path.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called at the end of every rendered frame. The [param frame_image] and [param audio_frame_block] function arguments should be written to.
*/
func (Instance) _write_frame(impl func(ptr unsafe.Pointer, frame_image [1]gdclass.Image, audio_frame_block unsafe.Pointer) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var frame_image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(frame_image[0])
		var audio_frame_block = gd.UnsafeGet[unsafe.Pointer](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, frame_image, audio_frame_block)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the engine finishes writing. This occurs when the engine quits by pressing the window manager's close button, or when [method SceneTree.quit] is called.
[b]Note:[/b] Pressing [kbd]Ctrl + C[/kbd] on the terminal running the editor/project does [i]not[/i] result in [method _write_end] being called.
*/
func (Instance) _write_end(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Adds a writer to be usable by the engine. The supported file extensions can be set by overriding [method _handles_file].
[b]Note:[/b] [method add_writer] must be called early enough in the engine initialization to work, as movie writing is designed to start at the same time as the rest of the engine.
*/
func AddWriter(writer [1]gdclass.MovieWriter) {
	self := Instance{}
	class(self).AddWriter(writer)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MovieWriter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MovieWriter"))
	casted := Instance{*(*gdclass.MovieWriter)(unsafe.Pointer(&object))}
	return casted
}

/*
Called when the audio sample rate used for recording the audio is requested by the engine. The value returned must be specified in Hz. Defaults to 48000 Hz if [method _get_audio_mix_rate] is not overridden.
*/
func (class) _get_audio_mix_rate(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the audio speaker mode used for recording the audio is requested by the engine. This can affect the number of output channels in the resulting audio file/stream. Defaults to [constant AudioServer.SPEAKER_MODE_STEREO] if [method _get_audio_speaker_mode] is not overridden.
*/
func (class) _get_audio_speaker_mode(impl func(ptr unsafe.Pointer) gdclass.AudioServerSpeakerMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the engine determines whether this [MovieWriter] is able to handle the file at [param path]. Must return [code]true[/code] if this [MovieWriter] is able to handle the given file path, [code]false[/code] otherwise. Typically, [method _handles_file] is overridden as follows to allow the user to record a file at any path with a given file extension:
[codeblock]
func _handles_file(path):

	# Allows specifying an output file with a `.mkv` file extension (case-insensitive),
	# either in the Project Settings or with the `--write-movie <path>` command line argument.
	return path.get_extension().to_lower() == "mkv"

[/codeblock]
*/
func (class) _handles_file(impl func(ptr unsafe.Pointer, path gd.String) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called once before the engine starts writing video and audio data. [param movie_size] is the width and height of the video to save. [param fps] is the number of frames per second specified in the project settings or using the [code]--fixed-fps <fps>[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url].
*/
func (class) _write_begin(impl func(ptr unsafe.Pointer, movie_size gd.Vector2i, fps gd.Int, base_path gd.String) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var movie_size = gd.UnsafeGet[gd.Vector2i](p_args, 0)
		var fps = gd.UnsafeGet[gd.Int](p_args, 1)
		var base_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, movie_size, fps, base_path)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called at the end of every rendered frame. The [param frame_image] and [param audio_frame_block] function arguments should be written to.
*/
func (class) _write_frame(impl func(ptr unsafe.Pointer, frame_image [1]gdclass.Image, audio_frame_block unsafe.Pointer) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var frame_image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(frame_image[0])
		var audio_frame_block = gd.UnsafeGet[unsafe.Pointer](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, frame_image, audio_frame_block)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the engine finishes writing. This occurs when the engine quits by pressing the window manager's close button, or when [method SceneTree.quit] is called.
[b]Note:[/b] Pressing [kbd]Ctrl + C[/kbd] on the terminal running the editor/project does [i]not[/i] result in [method _write_end] being called.
*/
func (class) _write_end(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Adds a writer to be usable by the engine. The supported file extensions can be set by overriding [method _handles_file].
[b]Note:[/b] [method add_writer] must be called early enough in the engine initialization to work, as movie writing is designed to start at the same time as the rest of the engine.
*/
//go:nosplit
func (self class) AddWriter(writer [1]gdclass.MovieWriter) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(writer[0].AsObject()[0]))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MovieWriter.Bind_add_writer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsMovieWriter() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMovieWriter() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_audio_mix_rate":
		return reflect.ValueOf(self._get_audio_mix_rate)
	case "_get_audio_speaker_mode":
		return reflect.ValueOf(self._get_audio_speaker_mode)
	case "_handles_file":
		return reflect.ValueOf(self._handles_file)
	case "_write_begin":
		return reflect.ValueOf(self._write_begin)
	case "_write_frame":
		return reflect.ValueOf(self._write_frame)
	case "_write_end":
		return reflect.ValueOf(self._write_end)
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_audio_mix_rate":
		return reflect.ValueOf(self._get_audio_mix_rate)
	case "_get_audio_speaker_mode":
		return reflect.ValueOf(self._get_audio_speaker_mode)
	case "_handles_file":
		return reflect.ValueOf(self._handles_file)
	case "_write_begin":
		return reflect.ValueOf(self._write_begin)
	case "_write_frame":
		return reflect.ValueOf(self._write_frame)
	case "_write_end":
		return reflect.ValueOf(self._write_end)
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("MovieWriter", func(ptr gd.Object) any { return [1]gdclass.MovieWriter{*(*gdclass.MovieWriter)(unsafe.Pointer(&ptr))} })
}

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
