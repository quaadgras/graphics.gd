/*
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
*/

package main

import "graphics.gd/classdb/XMLParser"

func ExampleXMLParser() {
	var parser = XMLParser.New()
	parser.Open("path/to/file.svg")
	for parser.Read() != nil {
		if parser.GetNodeType() == XMLParser.NodeElement {
			var nodeName = parser.GetNodeName()
			var attributesDict = make(map[string]string)
			for idx := 0; idx < parser.GetAttributeCount(); idx++ {
				attributesDict[parser.GetAttributeName(idx)] = parser.GetAttributeValue(idx)
			}
			print("The ", nodeName, " element has the following attributes: ", attributesDict)
		}
	}
}
