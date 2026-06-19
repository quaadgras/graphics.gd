/*
[gdscript]
var fields = { "single": 123, "not_valued": null, "multiple": [22, 33, 44] }
var query_string = http_client.query_string_from_dict(fields)
# Returns "single=123&not_valued&multiple=22&multiple=33&multiple=44"
[/gdscript]
[csharp]
var fields = new Godot.Collections.Dictionary
{
	{ "single", 123 },
	{ "notValued", default },
	{ "multiple", new Godot.Collections.Array { 22, 33, 44 } },
};
string queryString = httpClient.QueryStringFromDict(fields);
// Returns "single=123&not_valued&multiple=22&multiple=33&multiple=44"
[/csharp]
*/

package main

import "graphics.gd/classdb/HTTPClient"

func ExampleQueryStringFromDict(httpClient HTTPClient.Instance) {
	var fields = map[string]string{"single": "123", "not_valued": "", "multiple": "22"}
	var queryString = httpClient.QueryStringFromDict(fields)
	// Returns "single=123&not_valued&multiple=22&multiple=33&multiple=44"
	_ = queryString
}
