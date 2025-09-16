/*
[gdscript]
var fields = { "username": "user", "password": "pass" }
var query_string = http_client.query_string_from_dict(fields)
# Returns "username=user&password=pass"
[/gdscript]
[csharp]
var fields = new Godot.Collections.Dictionary { { "username", "user" }, { "password", "pass" } };
string queryString = httpClient.QueryStringFromDict(fields);
// Returns "username=user&password=pass"
[/csharp]
*/

package main

func HTTPClient_QueryStringFromDict() {
	var fields = map[string]string{"username": "user", "password": "pass"}
	var queryString = http_client.QueryStringFromDict(fields)
	// Returns "username=user&password=pass"
	_ = queryString
}
