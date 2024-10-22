package XMLParser

import "unsafe"
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
Provides a low-level interface for creating parsers for [url=https://en.wikipedia.org/wiki/XML]XML[/url] files. This class can serve as base to make custom XML parsers.
To parse XML, you must open a file with the [method open] method or a buffer with the [method open_buffer] method. Then, the [method read] method must be called to parse the next nodes. Most of the methods take into consideration the currently parsed node.
Here is an example of using [XMLParser] to parse an SVG file (which is based on XML), printing each element and its attributes as a dictionary:
[codeblocks]
[gdscript]
var parser = XMLParser.new()
parser.open("path/to/file.svg")
while parser.read() != ERR_FILE_EOF:
    if parser.get_node_type() == XMLParser.NODE_ELEMENT:
        var node_name = parser.get_node_name()
        var attributes_dict = {}
        for idx in range(parser.get_attribute_count()):
            attributes_dict[parser.get_attribute_name(idx)] = parser.get_attribute_value(idx)
        print("The ", node_name, " element has the following attributes: ", attributes_dict)
[/gdscript]
[csharp]
var parser = new XmlParser();
parser.Open("path/to/file.svg");
while (parser.Read() != Error.FileEof)
{
    if (parser.GetNodeType() == XmlParser.NodeType.Element)
    {
        var nodeName = parser.GetNodeName();
        var attributesDict = new Godot.Collections.Dictionary();
        for (int idx = 0; idx < parser.GetAttributeCount(); idx++)
        {
            attributesDict[parser.GetAttributeName(idx)] = parser.GetAttributeValue(idx);
        }
        GD.Print($"The {nodeName} element has the following attributes: {attributesDict}");
    }
}
[/csharp]
[/codeblocks]

*/
type Go [1]classdb.XMLParser

/*
Parses the next node in the file. This method returns an error code.
*/
func (self Go) Read() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Read())
}

/*
Returns the type of the current node. Compare with [enum NodeType] constants.
*/
func (self Go) GetNodeType() classdb.XMLParserNodeType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XMLParserNodeType(class(self).GetNodeType())
}

/*
Returns the name of a node. This method will raise an error if the currently parsed node is a text node.
[b]Note:[/b] The content of a [constant NODE_CDATA] node and the comment string of a [constant NODE_COMMENT] node are also considered names.
*/
func (self Go) GetNodeName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNodeName(gc).String())
}

/*
Returns the contents of a text node. This method will raise an error if the current parsed node is of any other type.
*/
func (self Go) GetNodeData() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNodeData(gc).String())
}

/*
Returns the byte offset of the currently parsed node since the beginning of the file or buffer. This is usually equivalent to the number of characters before the read position.
*/
func (self Go) GetNodeOffset() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetNodeOffset()))
}

/*
Returns the number of attributes in the currently parsed element.
[b]Note:[/b] If this method is used while the currently parsed node is not [constant NODE_ELEMENT] or [constant NODE_ELEMENT_END], this count will not be updated and will still reflect the last element.
*/
func (self Go) GetAttributeCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetAttributeCount()))
}

/*
Returns the name of an attribute of the currently parsed element, specified by the [param idx] index.
*/
func (self Go) GetAttributeName(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetAttributeName(gc, gd.Int(idx)).String())
}

/*
Returns the value of an attribute of the currently parsed element, specified by the [param idx] index.
*/
func (self Go) GetAttributeValue(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetAttributeValue(gc, gd.Int(idx)).String())
}

/*
Returns [code]true[/code] if the currently parsed element has an attribute with the [param name].
*/
func (self Go) HasAttribute(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasAttribute(gc.String(name)))
}

/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will raise an error if the element has no such attribute.
*/
func (self Go) GetNamedAttributeValue(name string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNamedAttributeValue(gc, gc.String(name)).String())
}

/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will return an empty string if the element has no such attribute.
*/
func (self Go) GetNamedAttributeValueSafe(name string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNamedAttributeValueSafe(gc, gc.String(name)).String())
}

/*
Returns [code]true[/code] if the currently parsed element is empty, e.g. [code]<element />[/code].
*/
func (self Go) IsEmpty() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsEmpty())
}

/*
Returns the current line in the parsed file, counting from 0.
*/
func (self Go) GetCurrentLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCurrentLine()))
}

/*
Skips the current section. If the currently parsed node contains more inner nodes, they will be ignored and the cursor will go to the closing of the current element.
*/
func (self Go) SkipSection() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SkipSection()
}

/*
Moves the buffer cursor to a certain offset (since the beginning) and reads the next node there. This method returns an error code.
*/
func (self Go) SeekTo(position int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).SeekTo(gd.Int(position)))
}

/*
Opens an XML [param file] for parsing. This method returns an error code.
*/
func (self Go) Open(file string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Open(gc.String(file)))
}

/*
Opens an XML raw [param buffer] for parsing. This method returns an error code.
*/
func (self Go) OpenBuffer(buffer []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).OpenBuffer(gc.PackedByteSlice(buffer)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XMLParser
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("XMLParser"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Parses the next node in the file. This method returns an error code.
*/
//go:nosplit
func (self class) Read() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_read, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the type of the current node. Compare with [enum NodeType] constants.
*/
//go:nosplit
func (self class) GetNodeType() classdb.XMLParserNodeType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XMLParserNodeType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_node_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of a node. This method will raise an error if the currently parsed node is a text node.
[b]Note:[/b] The content of a [constant NODE_CDATA] node and the comment string of a [constant NODE_COMMENT] node are also considered names.
*/
//go:nosplit
func (self class) GetNodeName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_node_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the contents of a text node. This method will raise an error if the current parsed node is of any other type.
*/
//go:nosplit
func (self class) GetNodeData(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_node_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the byte offset of the currently parsed node since the beginning of the file or buffer. This is usually equivalent to the number of characters before the read position.
*/
//go:nosplit
func (self class) GetNodeOffset() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_node_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of attributes in the currently parsed element.
[b]Note:[/b] If this method is used while the currently parsed node is not [constant NODE_ELEMENT] or [constant NODE_ELEMENT_END], this count will not be updated and will still reflect the last element.
*/
//go:nosplit
func (self class) GetAttributeCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_attribute_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of an attribute of the currently parsed element, specified by the [param idx] index.
*/
//go:nosplit
func (self class) GetAttributeName(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_attribute_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the value of an attribute of the currently parsed element, specified by the [param idx] index.
*/
//go:nosplit
func (self class) GetAttributeValue(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_attribute_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the currently parsed element has an attribute with the [param name].
*/
//go:nosplit
func (self class) HasAttribute(name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_has_attribute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will raise an error if the element has no such attribute.
*/
//go:nosplit
func (self class) GetNamedAttributeValue(ctx gd.Lifetime, name gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_named_attribute_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the value of an attribute of the currently parsed element, specified by its [param name]. This method will return an empty string if the element has no such attribute.
*/
//go:nosplit
func (self class) GetNamedAttributeValueSafe(ctx gd.Lifetime, name gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_named_attribute_value_safe, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the currently parsed element is empty, e.g. [code]<element />[/code].
*/
//go:nosplit
func (self class) IsEmpty() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_is_empty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current line in the parsed file, counting from 0.
*/
//go:nosplit
func (self class) GetCurrentLine() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_get_current_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Skips the current section. If the currently parsed node contains more inner nodes, they will be ignored and the cursor will go to the closing of the current element.
*/
//go:nosplit
func (self class) SkipSection()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_skip_section, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the buffer cursor to a certain offset (since the beginning) and reads the next node there. This method returns an error code.
*/
//go:nosplit
func (self class) SeekTo(position gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Opens an XML [param file] for parsing. This method returns an error code.
*/
//go:nosplit
func (self class) Open(file gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Opens an XML raw [param buffer] for parsing. This method returns an error code.
*/
//go:nosplit
func (self class) OpenBuffer(buffer gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XMLParser.Bind_open_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXMLParser() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXMLParser() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("XMLParser", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type NodeType = classdb.XMLParserNodeType

const (
/*There's no node (no file or buffer opened).*/
	NodeNone NodeType = 0
/*An element node type, also known as a tag, e.g. [code]<title>[/code].*/
	NodeElement NodeType = 1
/*An end of element node type, e.g. [code]</title>[/code].*/
	NodeElementEnd NodeType = 2
/*A text node type, i.e. text that is not inside an element. This includes whitespace.*/
	NodeText NodeType = 3
/*A comment node type, e.g. [code]<!--A comment-->[/code].*/
	NodeComment NodeType = 4
/*A node type for CDATA (Character Data) sections, e.g. [code]<![CDATA[CDATA section]]>[/code].*/
	NodeCdata NodeType = 5
/*An unknown node type.*/
	NodeUnknown NodeType = 6
)
