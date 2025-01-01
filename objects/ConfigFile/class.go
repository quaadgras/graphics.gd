package ConfigFile

import "unsafe"
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
This helper class can be used to store [Variant] values on the filesystem using INI-style formatting. The stored values are identified by a section and a key:
[codeblock lang=text]
[section]
some_key=42
string_example="Hello World3D!"
a_vector=Vector3(1, 0, 2)
[/codeblock]
The stored data can be saved to or parsed from a file, though ConfigFile objects can also be used directly without accessing the filesystem.
The following example shows how to create a simple [ConfigFile] and save it on disc:
[codeblocks]
[gdscript]
# Create new ConfigFile object.
var config = ConfigFile.new()

# Store some values.
config.set_value("Player1", "player_name", "Steve")
config.set_value("Player1", "best_score", 10)
config.set_value("Player2", "player_name", "V3geta")
config.set_value("Player2", "best_score", 9001)

# Save it to a file (overwrite if already exists).
config.save("user://scores.cfg")
[/gdscript]
[csharp]
// Create new ConfigFile object.
var config = new ConfigFile();

// Store some values.
config.SetValue("Player1", "player_name", "Steve");
config.SetValue("Player1", "best_score", 10);
config.SetValue("Player2", "player_name", "V3geta");
config.SetValue("Player2", "best_score", 9001);

// Save it to a file (overwrite if already exists).
config.Save("user://scores.cfg");
[/csharp]
[/codeblocks]
This example shows how the above file could be loaded:
[codeblocks]
[gdscript]
var score_data = {}
var config = ConfigFile.new()

# Load data from a file.
var err = config.load("user://scores.cfg")

# If the file didn't load, ignore it.
if err != OK:

	return

# Iterate over all sections.
for player in config.get_sections():

	# Fetch the data for each section.
	var player_name = config.get_value(player, "player_name")
	var player_score = config.get_value(player, "best_score")
	score_data[player_name] = player_score

[/gdscript]
[csharp]
var score_data = new Godot.Collections.Dictionary();
var config = new ConfigFile();

// Load data from a file.
Error err = config.Load("user://scores.cfg");

// If the file didn't load, ignore it.
if (err != Error.Ok)

	{
	    return;
	}

// Iterate over all sections.
foreach (String player in config.GetSections())

	{
	    // Fetch the data for each section.
	    var player_name = (String)config.GetValue(player, "player_name");
	    var player_score = (int)config.GetValue(player, "best_score");
	    score_data[player_name] = player_score;
	}

[/csharp]
[/codeblocks]
Any operation that mutates the ConfigFile such as [method set_value], [method clear], or [method erase_section], only changes what is loaded in memory. If you want to write the change to a file, you have to save the changes with [method save], [method save_encrypted], or [method save_encrypted_pass].
Keep in mind that section and property names can't contain spaces. Anything after a space will be ignored on save and on load.
ConfigFiles can also contain manually written comment lines starting with a semicolon ([code];[/code]). Those lines will be ignored when parsing the file. Note that comments will be lost when saving the ConfigFile. This can still be useful for dedicated server configuration files, which are typically never overwritten without explicit user action.
[b]Note:[/b] The file extension given to a ConfigFile does not have any impact on its formatting or behavior. By convention, the [code].cfg[/code] extension is used here, but any other extension such as [code].ini[/code] is also valid. Since neither [code].cfg[/code] nor [code].ini[/code] are standardized, Godot's ConfigFile formatting may differ from files written by other programs.
*/
type Instance [1]classdb.ConfigFile
type Any interface {
	gd.IsClass
	AsConfigFile() Instance
}

/*
Assigns a value to the specified key of the specified section. If either the section or the key do not exist, they are created. Passing a [code]null[/code] value deletes the specified key if it exists, and deletes the section if it ends up empty once the key has been removed.
*/
func (self Instance) SetValue(section string, key string, value any) {
	class(self).SetValue(gd.NewString(section), gd.NewString(key), gd.NewVariant(value))
}

/*
Returns the current value for the specified section and key. If either the section or the key do not exist, the method returns the fallback [param default] value. If [param default] is not specified or set to [code]null[/code], an error is also raised.
*/
func (self Instance) GetValue(section string, key string) any {
	return any(class(self).GetValue(gd.NewString(section), gd.NewString(key), gd.NewVariant(gd.NewVariant(([1]any{}[0])))).Interface())
}

/*
Returns [code]true[/code] if the specified section exists.
*/
func (self Instance) HasSection(section string) bool {
	return bool(class(self).HasSection(gd.NewString(section)))
}

/*
Returns [code]true[/code] if the specified section-key pair exists.
*/
func (self Instance) HasSectionKey(section string, key string) bool {
	return bool(class(self).HasSectionKey(gd.NewString(section), gd.NewString(key)))
}

/*
Returns an array of all defined section identifiers.
*/
func (self Instance) GetSections() []string {
	return []string(class(self).GetSections().Strings())
}

/*
Returns an array of all defined key identifiers in the specified section. Raises an error and returns an empty array if the section does not exist.
*/
func (self Instance) GetSectionKeys(section string) []string {
	return []string(class(self).GetSectionKeys(gd.NewString(section)).Strings())
}

/*
Deletes the specified section along with all the key-value pairs inside. Raises an error if the section does not exist.
*/
func (self Instance) EraseSection(section string) {
	class(self).EraseSection(gd.NewString(section))
}

/*
Deletes the specified key in a section. Raises an error if either the section or the key do not exist.
*/
func (self Instance) EraseSectionKey(section string, key string) {
	class(self).EraseSectionKey(gd.NewString(section), gd.NewString(key))
}

/*
Loads the config file specified as a parameter. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) Load(path string) error {
	return error(class(self).Load(gd.NewString(path)))
}

/*
Parses the passed string as the contents of a config file. The string is parsed and loaded in the ConfigFile object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) Parse(data string) error {
	return error(class(self).Parse(gd.NewString(data)))
}

/*
Saves the contents of the [ConfigFile] object to the file specified as a parameter. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) Save(path string) error {
	return error(class(self).Save(gd.NewString(path)))
}

/*
Obtain the text version of this config file (the same text that would be written to a file).
*/
func (self Instance) EncodeToText() string {
	return string(class(self).EncodeToText().String())
}

/*
Loads the encrypted config file specified as a parameter, using the provided [param key] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) LoadEncrypted(path string, key []byte) error {
	return error(class(self).LoadEncrypted(gd.NewString(path), gd.NewPackedByteSlice(key)))
}

/*
Loads the encrypted config file specified as a parameter, using the provided [param password] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) LoadEncryptedPass(path string, password string) error {
	return error(class(self).LoadEncryptedPass(gd.NewString(path), gd.NewString(password)))
}

/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param key] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) SaveEncrypted(path string, key []byte) error {
	return error(class(self).SaveEncrypted(gd.NewString(path), gd.NewPackedByteSlice(key)))
}

/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param password] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
func (self Instance) SaveEncryptedPass(path string, password string) error {
	return error(class(self).SaveEncryptedPass(gd.NewString(path), gd.NewString(password)))
}

/*
Removes the entire contents of the config.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ConfigFile

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ConfigFile"))
	return Instance{classdb.ConfigFile(object)}
}

/*
Assigns a value to the specified key of the specified section. If either the section or the key do not exist, they are created. Passing a [code]null[/code] value deletes the specified key if it exists, and deletes the section if it ends up empty once the key has been removed.
*/
//go:nosplit
func (self class) SetValue(section gd.String, key gd.String, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(section))
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_set_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the current value for the specified section and key. If either the section or the key do not exist, the method returns the fallback [param default] value. If [param default] is not specified or set to [code]null[/code], an error is also raised.
*/
//go:nosplit
func (self class) GetValue(section gd.String, key gd.String, def gd.Variant) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(section))
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, pointers.Get(def))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_get_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified section exists.
*/
//go:nosplit
func (self class) HasSection(section gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(section))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_has_section, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified section-key pair exists.
*/
//go:nosplit
func (self class) HasSectionKey(section gd.String, key gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(section))
	callframe.Arg(frame, pointers.Get(key))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_has_section_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of all defined section identifiers.
*/
//go:nosplit
func (self class) GetSections() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_get_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of all defined key identifiers in the specified section. Raises an error and returns an empty array if the section does not exist.
*/
//go:nosplit
func (self class) GetSectionKeys(section gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(section))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_get_section_keys, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Deletes the specified section along with all the key-value pairs inside. Raises an error if the section does not exist.
*/
//go:nosplit
func (self class) EraseSection(section gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(section))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_erase_section, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Deletes the specified key in a section. Raises an error if either the section or the key do not exist.
*/
//go:nosplit
func (self class) EraseSectionKey(section gd.String, key gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(section))
	callframe.Arg(frame, pointers.Get(key))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_erase_section_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Loads the config file specified as a parameter. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) Load(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_load, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Parses the passed string as the contents of a config file. The string is parsed and loaded in the ConfigFile object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) Parse(data gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_parse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the contents of the [ConfigFile] object to the file specified as a parameter. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) Save(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_save, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Obtain the text version of this config file (the same text that would be written to a file).
*/
//go:nosplit
func (self class) EncodeToText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_encode_to_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Loads the encrypted config file specified as a parameter, using the provided [param key] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) LoadEncrypted(path gd.String, key gd.PackedByteArray) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(key))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_load_encrypted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads the encrypted config file specified as a parameter, using the provided [param password] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) LoadEncryptedPass(path gd.String, password gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(password))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_load_encrypted_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param key] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) SaveEncrypted(path gd.String, key gd.PackedByteArray) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(key))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_save_encrypted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [param password] to encrypt it. The output file uses an INI-style structure.
Returns [constant OK] on success, or one of the other [enum Error] values if the operation failed.
*/
//go:nosplit
func (self class) SaveEncryptedPass(path gd.String, password gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(password))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_save_encrypted_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the entire contents of the config.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConfigFile.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsConfigFile() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsConfigFile() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("ConfigFile", func(ptr gd.Object) any { return [1]classdb.ConfigFile{classdb.ConfigFile(ptr)} })
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
