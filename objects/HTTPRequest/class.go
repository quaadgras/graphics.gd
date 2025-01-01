package HTTPRequest

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A node with the ability to send HTTP requests. Uses [HTTPClient] internally.
Can be used to make HTTP requests, i.e. download or upload files or web content via HTTP.
[b]Warning:[/b] See the notes and warnings on [HTTPClient] for limitations, especially regarding TLS security.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
[b]Example of contacting a REST API and printing one of its returned fields:[/b]
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
[b]Example of loading and displaying an image using HTTPRequest:[/b]
[codeblocks]
[gdscript]
func _ready():

	# Create an HTTP request node and connect its completion signal.
	var http_request = HTTPRequest.new()
	add_child(http_request)
	http_request.request_completed.connect(self._http_request_completed)

	# Perform the HTTP request. The URL below returns a PNG image as of writing.
	var error = http_request.request("https://via.placeholder.com/512")
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
	    Error error = httpRequest.Request("https://via.placeholder.com/512");
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
[b]Gzipped response bodies[/b]: HTTPRequest will automatically handle decompression of response bodies. A [code]Accept-Encoding[/code] header will be automatically added to each of your requests, unless one is already specified. Any response with a [code]Content-Encoding: gzip[/code] header will automatically be decompressed and delivered to you as uncompressed bytes.
*/
type Instance [1]classdb.HTTPRequest
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
func (self Instance) Request(url string) error {
	return error(class(self).Request(gd.NewString(url), gd.NewPackedStringSlice([1][]string{}[0]), 0, gd.NewString("")))
}

/*
Creates request on the underlying [HTTPClient] using a raw array of bytes for the request body. If there is no configuration errors, it tries to connect using [method HTTPClient.connect_to_host] and passes parameters onto [method HTTPClient.request].
Returns [constant OK] if request is successfully created. (Does not imply that the server has responded), [constant ERR_UNCONFIGURED] if not in the tree, [constant ERR_BUSY] if still processing previous request, [constant ERR_INVALID_PARAMETER] if given string is not a valid URL format, or [constant ERR_CANT_CONNECT] if not using thread and the [HTTPClient] cannot connect to host.
*/
func (self Instance) RequestRaw(url string) error {
	return error(class(self).RequestRaw(gd.NewString(url), gd.NewPackedStringSlice([1][]string{}[0]), 0, gd.NewPackedByteSlice([1][]byte{}[0])))
}

/*
Cancels the current request.
*/
func (self Instance) CancelRequest() {
	class(self).CancelRequest()
}

/*
Sets the [TLSOptions] to be used when connecting to an HTTPS server. See [method TLSOptions.client].
*/
func (self Instance) SetTlsOptions(client_options objects.TLSOptions) {
	class(self).SetTlsOptions(client_options)
}

/*
Returns the current status of the underlying [HTTPClient]. See [enum HTTPClient.Status].
*/
func (self Instance) GetHttpClientStatus() classdb.HTTPClientStatus {
	return classdb.HTTPClientStatus(class(self).GetHttpClientStatus())
}

/*
Returns the number of bytes this HTTPRequest downloaded.
*/
func (self Instance) GetDownloadedBytes() int {
	return int(int(class(self).GetDownloadedBytes()))
}

/*
Returns the response body length.
[b]Note:[/b] Some Web servers may not send a body length. In this case, the value returned will be [code]-1[/code]. If using chunked transfer encoding, the body length will also be [code]-1[/code].
*/
func (self Instance) GetBodySize() int {
	return int(int(class(self).GetBodySize()))
}

/*
Sets the proxy server for HTTP requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
func (self Instance) SetHttpProxy(host string, port int) {
	class(self).SetHttpProxy(gd.NewString(host), gd.Int(port))
}

/*
Sets the proxy server for HTTPS requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
func (self Instance) SetHttpsProxy(host string, port int) {
	class(self).SetHttpsProxy(gd.NewString(host), gd.Int(port))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.HTTPRequest

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HTTPRequest"))
	return Instance{classdb.HTTPRequest(object)}
}

func (self Instance) DownloadFile() string {
	return string(class(self).GetDownloadFile().String())
}

func (self Instance) SetDownloadFile(value string) {
	class(self).SetDownloadFile(gd.NewString(value))
}

func (self Instance) DownloadChunkSize() int {
	return int(int(class(self).GetDownloadChunkSize()))
}

func (self Instance) SetDownloadChunkSize(value int) {
	class(self).SetDownloadChunkSize(gd.Int(value))
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
	class(self).SetBodySizeLimit(gd.Int(value))
}

func (self Instance) MaxRedirects() int {
	return int(int(class(self).GetMaxRedirects()))
}

func (self Instance) SetMaxRedirects(value int) {
	class(self).SetMaxRedirects(gd.Int(value))
}

func (self Instance) Timeout() Float.X {
	return Float.X(Float.X(class(self).GetTimeout()))
}

func (self Instance) SetTimeout(value Float.X) {
	class(self).SetTimeout(gd.Float(value))
}

/*
Creates request on the underlying [HTTPClient]. If there is no configuration errors, it tries to connect using [method HTTPClient.connect_to_host] and passes parameters onto [method HTTPClient.request].
Returns [constant OK] if request is successfully created. (Does not imply that the server has responded), [constant ERR_UNCONFIGURED] if not in the tree, [constant ERR_BUSY] if still processing previous request, [constant ERR_INVALID_PARAMETER] if given string is not a valid URL format, or [constant ERR_CANT_CONNECT] if not using thread and the [HTTPClient] cannot connect to host.
[b]Note:[/b] When [param method] is [constant HTTPClient.METHOD_GET], the payload sent via [param request_data] might be ignored by the server or even cause the server to reject the request (check [url=https://datatracker.ietf.org/doc/html/rfc7231#section-4.3.1]RFC 7231 section 4.3.1[/url] for more details). As a workaround, you can send data as a query string in the URL (see [method String.uri_encode] for an example).
[b]Note:[/b] It's recommended to use transport encryption (TLS) and to avoid sending sensitive information (such as login credentials) in HTTP GET URL parameters. Consider using HTTP POST requests or HTTP headers for such information instead.
*/
//go:nosplit
func (self class) Request(url gd.String, custom_headers gd.PackedStringArray, method classdb.HTTPClientMethod, request_data gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(url))
	callframe.Arg(frame, pointers.Get(custom_headers))
	callframe.Arg(frame, method)
	callframe.Arg(frame, pointers.Get(request_data))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_request, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates request on the underlying [HTTPClient] using a raw array of bytes for the request body. If there is no configuration errors, it tries to connect using [method HTTPClient.connect_to_host] and passes parameters onto [method HTTPClient.request].
Returns [constant OK] if request is successfully created. (Does not imply that the server has responded), [constant ERR_UNCONFIGURED] if not in the tree, [constant ERR_BUSY] if still processing previous request, [constant ERR_INVALID_PARAMETER] if given string is not a valid URL format, or [constant ERR_CANT_CONNECT] if not using thread and the [HTTPClient] cannot connect to host.
*/
//go:nosplit
func (self class) RequestRaw(url gd.String, custom_headers gd.PackedStringArray, method classdb.HTTPClientMethod, request_data_raw gd.PackedByteArray) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(url))
	callframe.Arg(frame, pointers.Get(custom_headers))
	callframe.Arg(frame, method)
	callframe.Arg(frame, pointers.Get(request_data_raw))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_request_raw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Cancels the current request.
*/
//go:nosplit
func (self class) CancelRequest() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_cancel_request, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the [TLSOptions] to be used when connecting to an HTTPS server. See [method TLSOptions.client].
*/
//go:nosplit
func (self class) SetTlsOptions(client_options objects.TLSOptions) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(client_options[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_tls_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the current status of the underlying [HTTPClient]. See [enum HTTPClient.Status].
*/
//go:nosplit
func (self class) GetHttpClientStatus() classdb.HTTPClientStatus {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.HTTPClientStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_http_client_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseThreads(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_use_threads, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingThreads() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_is_using_threads, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAcceptGzip(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_accept_gzip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAcceptingGzip() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_is_accepting_gzip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBodySizeLimit(bytes gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_body_size_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBodySizeLimit() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_body_size_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxRedirects(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_max_redirects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxRedirects() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_max_redirects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDownloadFile(path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_download_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDownloadFile() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_download_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the number of bytes this HTTPRequest downloaded.
*/
//go:nosplit
func (self class) GetDownloadedBytes() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_downloaded_bytes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the response body length.
[b]Note:[/b] Some Web servers may not send a body length. In this case, the value returned will be [code]-1[/code]. If using chunked transfer encoding, the body length will also be [code]-1[/code].
*/
//go:nosplit
func (self class) GetBodySize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_body_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTimeout(timeout gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_timeout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTimeout() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_timeout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDownloadChunkSize(chunk_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, chunk_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_download_chunk_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDownloadChunkSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_get_download_chunk_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the proxy server for HTTP requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
//go:nosplit
func (self class) SetHttpProxy(host gd.String, port gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(host))
	callframe.Arg(frame, port)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_http_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the proxy server for HTTPS requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
//go:nosplit
func (self class) SetHttpsProxy(host gd.String, port gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(host))
	callframe.Arg(frame, port)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HTTPRequest.Bind_set_https_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnRequestCompleted(cb func(result int, response_code int, headers []string, body []byte)) {
	self[0].AsObject().Connect(gd.NewStringName("request_completed"), gd.NewCallable(cb), 0)
}

func (self class) AsHTTPRequest() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHTTPRequest() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced      { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance   { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {
	classdb.Register("HTTPRequest", func(ptr gd.Object) any { return [1]classdb.HTTPRequest{classdb.HTTPRequest(ptr)} })
}

type Result = classdb.HTTPRequestResult

const (
	/*Request successful.*/
	ResultSuccess                 Result = 0
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
	ResultBodyDecompressFailed  Result = 8
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

type Error int

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
