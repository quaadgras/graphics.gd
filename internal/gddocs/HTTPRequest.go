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
	var body = JSON.stringify({"name": "Godette"})
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
	string body = Json.Stringify(new Godot.Collections.Dictionary
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
	"graphics.gd/classdb/HTTPClient"
	"graphics.gd/classdb/HTTPRequest"
	"graphics.gd/classdb/JSON"
	"graphics.gd/classdb/Node"
)

type httpRequestExample struct {
	Node.Extension[httpRequestExample]
}

func (n httpRequestExample) Ready() {
	// Create an HTTP request node and connect its completion signal.
	var httpRequest = HTTPRequest.New()
	n.AsNode().AddChild(httpRequest.AsNode())
	httpRequest.OnRequestCompleted(n.httpRequestCompleted)

	// Perform a GET request. The URL below returns JSON as of writing.
	if err := httpRequest.Request("https://httpbin.org/get"); err != nil {
		// An error occurred in the HTTP request.
	}

	// Perform a POST request. The URL below returns JSON as of writing.
	// Note: Don't make simultaneous requests using a single HTTPRequest node.
	var body = JSON.Stringify(map[string]any{"name": "Godette"}, "", false)
	if err := httpRequest.MoreArgs().Request("https://httpbin.org/post", nil, HTTPClient.MethodPost, body); err != nil {
		// An error occurred in the HTTP request.
	}
}

// Called when the HTTP request is completed.
func (n httpRequestExample) httpRequestCompleted(result HTTPRequest.Result, responseCode int, headers []string, body []byte) {
	var json = JSON.New()
	json.Parse(string(body))
	var response = json.Data()
	_ = response
	// Will print the user agent string used by the HTTPRequest node (as recognized by httpbin.org).
	// fmt.Println(response["headers"]["User-Agent"])
}
