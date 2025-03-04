// Package HTTPRequest provides methods for working with HTTPRequest object instances.
package HTTPRequest

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A node with the ability to send HTTP requests. Uses [HTTPClient] internally.
Can be used to make HTTP requests, i.e. download or upload files or web content via HTTP.
[b]Warning:[/b] See the notes and warnings on [HTTPClient] for limitations, especially regarding TLS security.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
[b]Example:[/b] Contact a REST API and print one of its returned fields:
[codeblocks]
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
[/codeblocks]
[b]Example:[/b] Load an image using [HTTPRequest] and display it:
[codeblocks]
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
[/codeblocks]
[b]Note:[/b] [HTTPRequest] nodes will automatically handle decompression of response bodies. A [code]Accept-Encoding[/code] header will be automatically added to each of your requests, unless one is already specified. Any response with a [code]Content-Encoding: gzip[/code] header will automatically be decompressed and delivered to you as uncompressed bytes.
*/
type Instance [1]gdclass.HTTPRequest

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsHTTPRequest() Instance
}

/*
Creates request on the underlying [HTTPClient]. If there is no configuration errors, it tries to connect using [method HTTPClient.connect_to_host] and passes parameters onto [method HTTPClient.request].
Returns [constant OK] if request is successfully created. (Does not imply that the server has responded), [constant ERR_UNCONFIGURED] if not in the tree, [constant ERR_BUSY] if still processing previous request, [constant ERR_INVALID_PARAMETER] if given string is not a valid URL format, or [constant ERR_CANT_CONNECT] if not using thread and the [HTTPClient] cannot connect to host.
[b]Note:[/b] When [param method] is [constant HTTPClient.METHOD_GET], the payload sent via [param request_data] might be ignored by the server or even cause the server to reject the request (check [url=https://datatracker.ietf.org/doc/html/rfc7231#section-4.3.1]RFC 7231 section 4.3.1[/url] for more details). As a workaround, you can send data as a query string in the URL (see [method String.uri_encode] for an example).
[b]Note:[/b] It's recommended to use transport encryption (TLS) and to avoid sending sensitive information (such as login credentials) in HTTP GET URL parameters. Consider using HTTP POST requests or HTTP headers for such information instead.
*/
func (self Instance) Request(url string) error { //gd:HTTPRequest.request
	return error(gd.ToError(class(self).Request(String.New(url), Packed.MakeStrings([1][]string{}[0]...), 0, String.New(""))))
}

/*
Creates request on the underlying [HTTPClient] using a raw array of bytes for the request body. If there is no configuration errors, it tries to connect using [method HTTPClient.connect_to_host] and passes parameters onto [method HTTPClient.request].
Returns [constant OK] if request is successfully created. (Does not imply that the server has responded), [constant ERR_UNCONFIGURED] if not in the tree, [constant ERR_BUSY] if still processing previous request, [constant ERR_INVALID_PARAMETER] if given string is not a valid URL format, or [constant ERR_CANT_CONNECT] if not using thread and the [HTTPClient] cannot connect to host.
*/
func (self Instance) RequestRaw(url string) error { //gd:HTTPRequest.request_raw
	return error(gd.ToError(class(self).RequestRaw(String.New(url), Packed.MakeStrings([1][]string{}[0]...), 0, Packed.Bytes(Packed.New([1][]byte{}[0]...)))))
}

/*
Cancels the current request.
*/
func (self Instance) CancelRequest() { //gd:HTTPRequest.cancel_request
	class(self).CancelRequest()
}

/*
Sets the [TLSOptions] to be used when connecting to an HTTPS server. See [method TLSOptions.client].
*/
func (self Instance) SetTlsOptions(client_options [1]gdclass.TLSOptions) { //gd:HTTPRequest.set_tls_options
	class(self).SetTlsOptions(client_options)
}

/*
Returns the current status of the underlying [HTTPClient]. See [enum HTTPClient.Status].
*/
func (self Instance) GetHttpClientStatus() gdclass.HTTPClientStatus { //gd:HTTPRequest.get_http_client_status
	return gdclass.HTTPClientStatus(class(self).GetHttpClientStatus())
}

/*
Returns the number of bytes this HTTPRequest downloaded.
*/
func (self Instance) GetDownloadedBytes() int { //gd:HTTPRequest.get_downloaded_bytes
	return int(int(class(self).GetDownloadedBytes()))
}

/*
Returns the response body length.
[b]Note:[/b] Some Web servers may not send a body length. In this case, the value returned will be [code]-1[/code]. If using chunked transfer encoding, the body length will also be [code]-1[/code].
*/
func (self Instance) GetBodySize() int { //gd:HTTPRequest.get_body_size
	return int(int(class(self).GetBodySize()))
}

/*
Sets the proxy server for HTTP requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
func (self Instance) SetHttpProxy(host string, port int) { //gd:HTTPRequest.set_http_proxy
	class(self).SetHttpProxy(String.New(host), int64(port))
}

/*
Sets the proxy server for HTTPS requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
func (self Instance) SetHttpsProxy(host string, port int) { //gd:HTTPRequest.set_https_proxy
	class(self).SetHttpsProxy(String.New(host), int64(port))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.HTTPRequest

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HTTPRequest"))
	casted := Instance{*(*gdclass.HTTPRequest)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) DownloadFile() string {
	return string(class(self).GetDownloadFile().String())
}

func (self Instance) SetDownloadFile(value string) {
	class(self).SetDownloadFile(String.New(value))
}

func (self Instance) DownloadChunkSize() int {
	return int(int(class(self).GetDownloadChunkSize()))
}

func (self Instance) SetDownloadChunkSize(value int) {
	class(self).SetDownloadChunkSize(int64(value))
}

func (self Instance) UseThreads() bool {
	return bool(class(self).IsUsingThreads())
}

func (self Instance) SetUseThreads(value bool) {
	class(self).SetUseThreads(value)
}

func (self Instance) AcceptGzip() bool {
	return bool(class(self).IsAcceptingGzip())
}

func (self Instance) SetAcceptGzip(value bool) {
	class(self).SetAcceptGzip(value)
}

func (self Instance) BodySizeLimit() int {
	return int(int(class(self).GetBodySizeLimit()))
}

func (self Instance) SetBodySizeLimit(value int) {
	class(self).SetBodySizeLimit(int64(value))
}

func (self Instance) MaxRedirects() int {
	return int(int(class(self).GetMaxRedirects()))
}

func (self Instance) SetMaxRedirects(value int) {
	class(self).SetMaxRedirects(int64(value))
}

func (self Instance) Timeout() Float.X {
	return Float.X(Float.X(class(self).GetTimeout()))
}

func (self Instance) SetTimeout(value Float.X) {
	class(self).SetTimeout(float64(value))
}

/*
Creates request on the underlying [HTTPClient]. If there is no configuration errors, it tries to connect using [method HTTPClient.connect_to_host] and passes parameters onto [method HTTPClient.request].
Returns [constant OK] if request is successfully created. (Does not imply that the server has responded), [constant ERR_UNCONFIGURED] if not in the tree, [constant ERR_BUSY] if still processing previous request, [constant ERR_INVALID_PARAMETER] if given string is not a valid URL format, or [constant ERR_CANT_CONNECT] if not using thread and the [HTTPClient] cannot connect to host.
[b]Note:[/b] When [param method] is [constant HTTPClient.METHOD_GET], the payload sent via [param request_data] might be ignored by the server or even cause the server to reject the request (check [url=https://datatracker.ietf.org/doc/html/rfc7231#section-4.3.1]RFC 7231 section 4.3.1[/url] for more details). As a workaround, you can send data as a query string in the URL (see [method String.uri_encode] for an example).
[b]Note:[/b] It's recommended to use transport encryption (TLS) and to avoid sending sensitive information (such as login credentials) in HTTP GET URL parameters. Consider using HTTP POST requests or HTTP headers for such information instead.
*/
//go:nosplit
func (self class) Request(url String.Readable, custom_headers Packed.Strings, method gdclass.HTTPClientMethod, request_data String.Readable) Error.Code { //gd:HTTPRequest.request
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(url)))
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(custom_headers)))
	callframe.Arg(frame, method)
	callframe.Arg(frame, pointers.Get(gd.InternalString(request_data)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_request, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates request on the underlying [HTTPClient] using a raw array of bytes for the request body. If there is no configuration errors, it tries to connect using [method HTTPClient.connect_to_host] and passes parameters onto [method HTTPClient.request].
Returns [constant OK] if request is successfully created. (Does not imply that the server has responded), [constant ERR_UNCONFIGURED] if not in the tree, [constant ERR_BUSY] if still processing previous request, [constant ERR_INVALID_PARAMETER] if given string is not a valid URL format, or [constant ERR_CANT_CONNECT] if not using thread and the [HTTPClient] cannot connect to host.
*/
//go:nosplit
func (self class) RequestRaw(url String.Readable, custom_headers Packed.Strings, method gdclass.HTTPClientMethod, request_data_raw Packed.Bytes) Error.Code { //gd:HTTPRequest.request_raw
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(url)))
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(custom_headers)))
	callframe.Arg(frame, method)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](request_data_raw))))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_request_raw, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Cancels the current request.
*/
//go:nosplit
func (self class) CancelRequest() { //gd:HTTPRequest.cancel_request
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_cancel_request, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [TLSOptions] to be used when connecting to an HTTPS server. See [method TLSOptions.client].
*/
//go:nosplit
func (self class) SetTlsOptions(client_options [1]gdclass.TLSOptions) { //gd:HTTPRequest.set_tls_options
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(client_options[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_tls_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current status of the underlying [HTTPClient]. See [enum HTTPClient.Status].
*/
//go:nosplit
func (self class) GetHttpClientStatus() gdclass.HTTPClientStatus { //gd:HTTPRequest.get_http_client_status
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.HTTPClientStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_http_client_status, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseThreads(enable bool) { //gd:HTTPRequest.set_use_threads
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_use_threads, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingThreads() bool { //gd:HTTPRequest.is_using_threads
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_is_using_threads, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAcceptGzip(enable bool) { //gd:HTTPRequest.set_accept_gzip
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_accept_gzip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAcceptingGzip() bool { //gd:HTTPRequest.is_accepting_gzip
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_is_accepting_gzip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBodySizeLimit(bytes int64) { //gd:HTTPRequest.set_body_size_limit
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_body_size_limit, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBodySizeLimit() int64 { //gd:HTTPRequest.get_body_size_limit
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_body_size_limit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxRedirects(amount int64) { //gd:HTTPRequest.set_max_redirects
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_max_redirects, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxRedirects() int64 { //gd:HTTPRequest.get_max_redirects
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_max_redirects, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDownloadFile(path String.Readable) { //gd:HTTPRequest.set_download_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_download_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDownloadFile() String.Readable { //gd:HTTPRequest.get_download_file
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_download_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the number of bytes this HTTPRequest downloaded.
*/
//go:nosplit
func (self class) GetDownloadedBytes() int64 { //gd:HTTPRequest.get_downloaded_bytes
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_downloaded_bytes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the response body length.
[b]Note:[/b] Some Web servers may not send a body length. In this case, the value returned will be [code]-1[/code]. If using chunked transfer encoding, the body length will also be [code]-1[/code].
*/
//go:nosplit
func (self class) GetBodySize() int64 { //gd:HTTPRequest.get_body_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_body_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTimeout(timeout float64) { //gd:HTTPRequest.set_timeout
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_timeout, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTimeout() float64 { //gd:HTTPRequest.get_timeout
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_timeout, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDownloadChunkSize(chunk_size int64) { //gd:HTTPRequest.set_download_chunk_size
	var frame = callframe.New()
	callframe.Arg(frame, chunk_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_download_chunk_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDownloadChunkSize() int64 { //gd:HTTPRequest.get_download_chunk_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_download_chunk_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the proxy server for HTTP requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
//go:nosplit
func (self class) SetHttpProxy(host String.Readable, port int64) { //gd:HTTPRequest.set_http_proxy
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(host)))
	callframe.Arg(frame, port)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_http_proxy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the proxy server for HTTPS requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
//go:nosplit
func (self class) SetHttpsProxy(host String.Readable, port int64) { //gd:HTTPRequest.set_https_proxy
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(host)))
	callframe.Arg(frame, port)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_https_proxy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnRequestCompleted(cb func(result int, response_code int, headers []string, body []byte)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("request_completed"), gd.NewCallable(cb), 0)
}

func (self class) AsHTTPRequest() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHTTPRequest() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced      { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance   { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("HTTPRequest", func(ptr gd.Object) any { return [1]gdclass.HTTPRequest{*(*gdclass.HTTPRequest)(unsafe.Pointer(&ptr))} })
}

type Result = gdclass.HTTPRequestResult //gd:HTTPRequest.Result

const (
	/*Request successful.*/
	ResultSuccess Result = 0
	/*Request failed due to a mismatch between the expected and actual chunked body size during transfer. Possible causes include network errors, server misconfiguration, or issues with chunked encoding.*/
	ResultChunkedBodySizeMismatch Result = 1
	/*Request failed while connecting.*/
	ResultCantConnect Result = 2
	/*Request failed while resolving.*/
	ResultCantResolve Result = 3
	/*Request failed due to connection (read/write) error.*/
	ResultConnectionError Result = 4
	/*Request failed on TLS handshake.*/
	ResultTlsHandshakeError Result = 5
	/*Request does not have a response (yet).*/
	ResultNoResponse Result = 6
	/*Request exceeded its maximum size limit, see [member body_size_limit].*/
	ResultBodySizeLimitExceeded Result = 7
	/*Request failed due to an error while decompressing the response body. Possible causes include unsupported or incorrect compression format, corrupted data, or incomplete transfer.*/
	ResultBodyDecompressFailed Result = 8
	/*Request failed (currently unused).*/
	ResultRequestFailed Result = 9
	/*HTTPRequest couldn't open the download file.*/
	ResultDownloadFileCantOpen Result = 10
	/*HTTPRequest couldn't write to the download file.*/
	ResultDownloadFileWriteError Result = 11
	/*Request reached its maximum redirect limit, see [member max_redirects].*/
	ResultRedirectLimitReached Result = 12
	/*Request failed due to a timeout. If you expect requests to take a long time, try increasing the value of [member timeout] or setting it to [code]0.0[/code] to remove the timeout completely.*/
	ResultTimeout Result = 13
)
