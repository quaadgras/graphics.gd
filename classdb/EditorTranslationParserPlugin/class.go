// Package EditorTranslationParserPlugin provides methods for working with EditorTranslationParserPlugin object instances.
package EditorTranslationParserPlugin

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[EditorTranslationParserPlugin] is invoked when a file is being parsed to extract strings that require translation. To define the parsing and string extraction logic, override the [method _parse_file] method in script.
Add the extracted strings to argument [code]msgids[/code] or [code]msgids_context_plural[/code] if context or plural is used.
When adding to [code]msgids_context_plural[/code], you must add the data using the format [code]["A", "B", "C"][/code], where [code]A[/code] represents the extracted string, [code]B[/code] represents the context, and [code]C[/code] represents the plural version of the extracted string. If you want to add only context but not plural, put [code]""[/code] for the plural slot. The idea is the same if you only want to add plural but not context. See the code below for concrete examples.
The extracted strings will be written into a POT file selected by user under "POT Generation" in "Localization" tab in "Project Settings" menu.
Below shows an example of a custom parser that extracts strings from a CSV file to write into a POT.
[codeblocks]
[gdscript]
@tool
extends EditorTranslationParserPlugin

func _parse_file(path, msgids, msgids_context_plural):

	var file = FileAccess.open(path, FileAccess.READ)
	var text = file.get_as_text()
	var split_strs = text.split(",", false)
	for s in split_strs:
	    msgids.append(s)
	    #print("Extracted string: " + s)

func _get_recognized_extensions():

	return ["csv"]

[/gdscript]
[csharp]
using Godot;

[Tool]
public partial class CustomParser : EditorTranslationParserPlugin

	{
	    public override void _ParseFile(string path, Godot.Collections.Array<string> msgids, Godot.Collections.Array<Godot.Collections.Array> msgidsContextPlural)
	    {
	        using var file = FileAccess.Open(path, FileAccess.ModeFlags.Read);
	        string text = file.GetAsText();
	        string[] splitStrs = text.Split(",", allowEmpty: false);
	        foreach (string s in splitStrs)
	        {
	            msgids.Add(s);
	            //GD.Print($"Extracted string: {s}");
	        }
	    }

	    public override string[] _GetRecognizedExtensions()
	    {
	        return new string[] { "csv" };
	    }
	}

[/csharp]
[/codeblocks]
To add a translatable string associated with context or plural, add it to [code]msgids_context_plural[/code]:
[codeblocks]
[gdscript]
# This will add a message with msgid "Test 1", msgctxt "context", and msgid_plural "test 1 plurals".
msgids_context_plural.append(["Test 1", "context", "test 1 plurals"])
# This will add a message with msgid "A test without context" and msgid_plural "plurals".
msgids_context_plural.append(["A test without context", "", "plurals"])
# This will add a message with msgid "Only with context" and msgctxt "a friendly context".
msgids_context_plural.append(["Only with context", "a friendly context", ""])
[/gdscript]
[csharp]
// This will add a message with msgid "Test 1", msgctxt "context", and msgid_plural "test 1 plurals".
msgidsContextPlural.Add(new Godot.Collections.Array{"Test 1", "context", "test 1 Plurals"});
// This will add a message with msgid "A test without context" and msgid_plural "plurals".
msgidsContextPlural.Add(new Godot.Collections.Array{"A test without context", "", "plurals"});
// This will add a message with msgid "Only with context" and msgctxt "a friendly context".
msgidsContextPlural.Add(new Godot.Collections.Array{"Only with context", "a friendly context", ""});
[/csharp]
[/codeblocks]
[b]Note:[/b] If you override parsing logic for standard script types (GDScript, C#, etc.), it would be better to load the [code]path[/code] argument using [method ResourceLoader.load]. This is because built-in scripts are loaded as [Resource] type, not [FileAccess] type.
For example:
[codeblocks]
[gdscript]
func _parse_file(path, msgids, msgids_context_plural):

	var res = ResourceLoader.load(path, "Script")
	var text = res.source_code
	# Parsing logic.

func _get_recognized_extensions():

	return ["gd"]

[/gdscript]
[csharp]
public override void _ParseFile(string path, Godot.Collections.Array<string> msgids, Godot.Collections.Array<Godot.Collections.Array> msgidsContextPlural)

	{
	    var res = ResourceLoader.Load<Script>(path, "Script");
	    string text = res.SourceCode;
	    // Parsing logic.
	}

public override string[] _GetRecognizedExtensions()

	{
	    return new string[] { "gd" };
	}

[/csharp]
[/codeblocks]
To use [EditorTranslationParserPlugin], register it using the [method EditorPlugin.add_translation_parser_plugin] method first.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorTranslationParserPlugin)
*/
type Instance [1]gdclass.EditorTranslationParserPlugin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorTranslationParserPlugin() Instance
}
type Interface interface {
	//Override this method to define a custom parsing logic to extract the translatable strings.
	ParseFile(path string, msgids gd.Array, msgids_context_plural gd.Array)
	//Gets the list of file extensions to associate with this parser, e.g. [code]["csv"][/code].
	GetRecognizedExtensions() []string
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) ParseFile(path string, msgids gd.Array, msgids_context_plural gd.Array) {
	return
}
func (self Implementation) GetRecognizedExtensions() (_ []string) { return }

/*
Override this method to define a custom parsing logic to extract the translatable strings.
*/
func (Instance) _parse_file(impl func(ptr unsafe.Pointer, path string, msgids gd.Array, msgids_context_plural gd.Array)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(path)
		var msgids = pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(msgids)
		var msgids_context_plural = pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(msgids_context_plural)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, path.String(), msgids, msgids_context_plural)
	}
}

/*
Gets the list of file extensions to associate with this parser, e.g. [code]["csv"][/code].
*/
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorTranslationParserPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorTranslationParserPlugin"))
	casted := Instance{*(*gdclass.EditorTranslationParserPlugin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Override this method to define a custom parsing logic to extract the translatable strings.
*/
func (class) _parse_file(impl func(ptr unsafe.Pointer, path gd.String, msgids gd.Array, msgids_context_plural gd.Array)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		var msgids = pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		var msgids_context_plural = pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, path, msgids, msgids_context_plural)
	}
}

/*
Gets the list of file extensions to associate with this parser, e.g. [code]["csv"][/code].
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsEditorTranslationParserPlugin() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorTranslationParserPlugin() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_parse_file":
		return reflect.ValueOf(self._parse_file)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_parse_file":
		return reflect.ValueOf(self._parse_file)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorTranslationParserPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorTranslationParserPlugin{*(*gdclass.EditorTranslationParserPlugin)(unsafe.Pointer(&ptr))}
	})
}
