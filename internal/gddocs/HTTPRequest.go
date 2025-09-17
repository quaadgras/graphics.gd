/*
[gdscript]
func _ready():
	# Create an HTTP request node and connect its completion signal.
	var http_request = HTTPRequest.new()
	add_child(http_request)
	http_request.request_completed.connect(self._http_request_completed)

	# Perform a GET request. The URL below returns JSON as of writing.
	var error = http_request.request("https://httpbin.org/get")
	if error != OK:
		push_error("An error occurred in the HTTP request.")

	# Perform a POST request. The URL below returns JSON as of writing.
	# Note: Don't make simultaneous requests using a single HTTPRequest node.
	# The snippet below is provided for reference only.
	var body = JSON.new().stringify({"name": "Godette"})
	error = http_request.request("https://httpbin.org/post", [], HTTPClient.METHOD_POST, body)
	if error != OK:
		push_error("An error occurred in the HTTP request.")

# Called when the HTTP request is completed.
func _http_request_completed(result, response_code, headers, body):
	var json = JSON.new()
	json.parse(body.get_string_from_utf8())
	var response = json.get_data()

	# Will print the user agent string used by the HTTPRequest node (as recognized by httpbin.org).
	print(response.headers["User-Agent"])
[/gdscript]
[csharp]
public override void _Ready()
{
	// Create an HTTP request node and connect its completion signal.
	var httpRequest = new HttpRequest();
	AddChild(httpRequest);
	httpRequest.RequestCompleted += HttpRequestCompleted;

	// Perform a GET request. The URL below returns JSON as of writing.
	Error error = httpRequest.Request("https://httpbin.org/get");
	if (error != Error.Ok)
	{
		GD.PushError("An error occurred in the HTTP request.");
	}

	// Perform a POST request. The URL below returns JSON as of writing.
	// Note: Don't make simultaneous requests using a single HTTPRequest node.
	// The snippet below is provided for reference only.
	string body = new Json().Stringify(new Godot.Collections.Dictionary
	{
		{ "name", "Godette" }
	});
	error = httpRequest.Request("https://httpbin.org/post", null, HttpClient.Method.Post, body);
	if (error != Error.Ok)
	{
		GD.PushError("An error occurred in the HTTP request.");
	}
}

// Called when the HTTP request is completed.
private void HttpRequestCompleted(long result, long responseCode, string[] headers, byte[] body)
{
	var json = new Json();
	json.Parse(body.GetStringFromUtf8());
	var response = json.GetData().AsGodotDictionary();

	// Will print the user agent string used by the HTTPRequest node (as recognized by httpbin.org).
	GD.Print((response["headers"].AsGodotDictionary())["User-Agent"]);
}
[/csharp]
*/

package main

import (
	"encoding/json"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/HTTPClient"
	"graphics.gd/classdb/HTTPRequest"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Signal"
)

type ExampleHTTP struct {
	Node.Extension[ExampleHTTP]
}

func (n *ExampleHTTP) Ready() {
	// Create an HTTP request node and connect its completion signal.
	var httpRequest = HTTPRequest.New()
	n.AsNode().AddChild(httpRequest.AsNode())
	httpRequest.OnRequestCompleted(func(result HTTPRequest.Result, response_code int, headers []string, body []byte) {
		var Response struct {
			Headers map[string]string
		}
		json.Unmarshal(body, &Response)
		// Will print the user agent string used by the HTTPRequest node (as recognized by httpbin.org).
		println(Response.Headers["User-Agent"])

		// Perform a POST request. The URL below returns JSON as of writing.
		body, _ = json.Marshal(map[string]string{"name": "Godette"})
		var err = httpRequest.MoreArgs().Request("https://httpbin.org/post", nil, HTTPClient.MethodPost, string(body))
		if err != nil {
			Engine.Raise(err)
		}
	}, Signal.OneShot)
	// Perform a GET request. The URL below returns JSON as of writing.
	var err = httpRequest.MoreArgs().Request("https://httpbin.org/get", nil, HTTPClient.MethodGet, "")
	if err != nil {
		Engine.Raise(err)
	}
	// Note: Don't make simultaneous requests using a single HTTPRequest node.
}
