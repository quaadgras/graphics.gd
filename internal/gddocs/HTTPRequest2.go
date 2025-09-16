/*
[gdscript]
func _ready():
	# Create an HTTP request node and connect its completion signal.
	var http_request = HTTPRequest.new()
	add_child(http_request)
	http_request.request_completed.connect(self._http_request_completed)

	# Perform the HTTP request. The URL below returns a PNG image as of writing.
	var error = http_request.request("https://placehold.co/512")
	if error != OK:
		push_error("An error occurred in the HTTP request.")

# Called when the HTTP request is completed.
func _http_request_completed(result, response_code, headers, body):
	if result != HTTPRequest.RESULT_SUCCESS:
		push_error("Image couldn't be downloaded. Try a different image.")

	var image = Image.new()
	var error = image.load_png_from_buffer(body)
	if error != OK:
		push_error("Couldn't load the image.")

	var texture = ImageTexture.create_from_image(image)

	# Display the image in a TextureRect node.
	var texture_rect = TextureRect.new()
	add_child(texture_rect)
	texture_rect.texture = texture
[/gdscript]
[csharp]
public override void _Ready()
{
	// Create an HTTP request node and connect its completion signal.
	var httpRequest = new HttpRequest();
	AddChild(httpRequest);
	httpRequest.RequestCompleted += HttpRequestCompleted;

	// Perform the HTTP request. The URL below returns a PNG image as of writing.
	Error error = httpRequest.Request("https://placehold.co/512");
	if (error != Error.Ok)
	{
		GD.PushError("An error occurred in the HTTP request.");
	}
}

// Called when the HTTP request is completed.
private void HttpRequestCompleted(long result, long responseCode, string[] headers, byte[] body)
{
	if (result != (long)HttpRequest.Result.Success)
	{
		GD.PushError("Image couldn't be downloaded. Try a different image.");
	}
	var image = new Image();
	Error error = image.LoadPngFromBuffer(body);
	if (error != Error.Ok)
	{
		GD.PushError("Couldn't load the image.");
	}

	var texture = ImageTexture.CreateFromImage(image);

	// Display the image in a TextureRect node.
	var textureRect = new TextureRect();
	AddChild(textureRect);
	textureRect.Texture = texture;
}
[/csharp]
*/

package main

import (
	"errors"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/HTTPRequest"
	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/TextureRect"
	"graphics.gd/variant/Signal"
)

type ExampleDownloadImage struct {
	Node.Extension[ExampleDownloadImage]
}

func (n ExampleDownloadImage) Ready() {
	var http_request = HTTPRequest.New()
	n.AsNode().AddChild(http_request.AsNode())
	http_request.AsHTTPRequest().OnRequestCompleted(func(result HTTPRequest.Result, response_code int, headers []string, body []byte) {
		if result != HTTPRequest.ResultSuccess {
			Engine.Raise(errors.New("Image couldn't be downloaded. Try a different image."))
		}
		var image = Image.New()
		var err = image.LoadPngFromBuffer(body)
		if err != nil {
			Engine.Raise(errors.New("Couldn't load the image."))
		}
		var texture = ImageTexture.CreateFromImage(image)

		// Display the image in a TextureRect node.
		var texture_rect = TextureRect.New()
		texture_rect.AsTextureRect().SetTexture(texture.AsTexture2D())
		n.AsNode().AddChild(texture_rect.AsNode())
	}, Signal.OneShot)
	// Perform the HTTP request. The URL below returns a PNG image as of writing.
	var error = http_request.Request("https://placehold.co/512")
	if error != nil {
		panic("An error occurred in the HTTP request.")
	}
}
