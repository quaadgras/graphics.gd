/*
[gdscript]
var code_preview = TextEdit.new()
var highlighter = GDScriptSyntaxHighlighter.new()
code_preview.syntax_highlighter = highlighter
[/gdscript]
[csharp]
var codePreview = new TextEdit();
var highlighter = new GDScriptSyntaxHighlighter();
codePreview.SyntaxHighlighter = highlighter;
[/csharp]
*/
package main

import (
	"graphics.gd/classdb/GDScriptSyntaxHighlighter"
	"graphics.gd/classdb/TextEdit"
)

func ExampleScriptSyntaxHighlighter() {
	var code_preview = TextEdit.New()
	var highlighter = GDScriptSyntaxHighlighter.New()
	TextEdit.Advanced(code_preview).SetSyntaxHighlighter(highlighter.AsSyntaxHighlighter())
}
