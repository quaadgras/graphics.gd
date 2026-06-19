/*
[gdscript]
# This will add a message with msgid "Test 1", msgctxt "context", msgid_plural "test 1 plurals", comment "test 1 comment", and source line "7".
ret.append(PackedStringArray(["Test 1", "context", "test 1 plurals", "test 1 comment", "7"]))
# This will add a message with msgid "A test without context" and msgid_plural "plurals".
ret.append(PackedStringArray(["A test without context", "", "plurals"]))
# This will add a message with msgid "Only with context" and msgctxt "a friendly context".
ret.append(PackedStringArray(["Only with context", "a friendly context"]))
[/gdscript]
[csharp]
// This will add a message with msgid "Test 1", msgctxt "context", msgid_plural "test 1 plurals", comment "test 1 comment", and source line "7".
ret.Add(["Test 1", "context", "test 1 plurals", "test 1 comment", "7"]);
// This will add a message with msgid "A test without context" and msgid_plural "plurals".
ret.Add(["A test without context", "", "plurals"]);
// This will add a message with msgid "Only with context" and msgctxt "a friendly context".
ret.Add(["Only with context", "a friendly context"]);
[/csharp]
*/

package main

func ExampleEditorTranslationParserParse(ret *[][]string) {
	// Adds msgid "Test 1", msgctxt "context", msgid_plural "test 1 plurals", comment "test 1 comment", source line "7".
	*ret = append(*ret, []string{"Test 1", "context", "test 1 plurals", "test 1 comment", "7"})
	// Adds msgid "A test without context" and msgid_plural "plurals".
	*ret = append(*ret, []string{"A test without context", "", "plurals"})
	// Adds msgid "Only with context" and msgctxt "a friendly context".
	*ret = append(*ret, []string{"Only with context", "a friendly context"})
}
