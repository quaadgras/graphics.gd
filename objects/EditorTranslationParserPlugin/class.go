package EditorTranslationParserPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

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

	// EditorTranslationParserPlugin methods that can be overridden by a [Class] that extends it.
	type EditorTranslationParserPlugin interface {
		//Override this method to define a custom parsing logic to extract the translatable strings.
		ParseFile(path string, msgids gd.Array, msgids_context_plural gd.Array)
		//Gets the list of file extensions to associate with this parser, e.g. [code]["csv"][/code].
		GetRecognizedExtensions() []string
	}
*/
type Instance [1]classdb.EditorTranslationParserPlugin

/*
Override this method to define a custom parsing logic to extract the translatable strings.
*/
func (Instance) _parse_file(impl func(ptr unsafe.Pointer, path string, msgids gd.Array, msgids_context_plural gd.Array)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		var msgids = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(msgids)
		var msgids_context_plural = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 2))
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
type class [1]classdb.EditorTranslationParserPlugin

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorTranslationParserPlugin"))
	return Instance{classdb.EditorTranslationParserPlugin(object)}
}

/*
Override this method to define a custom parsing logic to extract the translatable strings.
*/
func (class) _parse_file(impl func(ptr unsafe.Pointer, path gd.String, msgids gd.Array, msgids_context_plural gd.Array)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var msgids = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var msgids_context_plural = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 2))
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_parse_file":
		return reflect.ValueOf(self._parse_file)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_parse_file":
		return reflect.ValueOf(self._parse_file)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EditorTranslationParserPlugin", func(ptr gd.Object) any { return classdb.EditorTranslationParserPlugin(ptr) })
}
