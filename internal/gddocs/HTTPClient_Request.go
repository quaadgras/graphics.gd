/*
[gdscript]
var fields = {"username" : "user", "password" : "pass"}
var query_string = http_client.query_string_from_dict(fields)
var headers = ["Content-Type: application/x-www-form-urlencoded", "Content-Length: " + str(query_string.length())]
var result = http_client.request(http_client.METHOD_POST, "/index.php", headers, query_string)
[/gdscript]
[csharp]
var fields = new Godot.Collections.Dictionary { { "username", "user" }, { "password", "pass" } };
string queryString = new HttpClient().QueryStringFromDict(fields);
string[] headers = ["Content-Type: application/x-www-form-urlencoded", $"Content-Length: {queryString.Length}"];
var result = new HttpClient().Request(HttpClient.Method.Post, "index.php", headers, queryString);
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/HTTPClient"
)

var http_client HTTPClient.Instance

func HTTPClient_Request() {
	var fields = map[string]string{"username": "user", "password": "pass"}
	var queryString = http_client.QueryStringFromDict(fields)
	var headers = []string{"Content-Type: application/x-www-form-urlencoded", "Content-Length: " + fmt.Sprint(len(queryString))}
	var result = HTTPClient.Expanded(http_client).Request(HTTPClient.MethodPost, "/index.php", headers, queryString)
	_ = result
}
