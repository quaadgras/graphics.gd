// Package JSONRPC provides methods for working with JSONRPC object instances.
package JSONRPC

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"

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

/*
[url=https://www.jsonrpc.org/]JSON-RPC[/url] is a standard which wraps a method call in a [JSON] object. The object has a particular structure and identifies which method is called, the parameters to that function, and carries an ID to keep track of responses. This class implements that standard on top of [Dictionary]; you will have to convert between a [Dictionary] and [JSON] with other functions.
*/
type Instance [1]gdclass.JSONRPC

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsJSONRPC() Instance
}

func (self Instance) SetScope(scope string, target Object.Instance) { //gd:JSONRPC.set_scope
	class(self).SetScope(gd.NewString(scope), target)
}

/*
Given a Dictionary which takes the form of a JSON-RPC request: unpack the request and run it. Methods are resolved by looking at the field called "method" and looking for an equivalently named function in the JSONRPC object. If one is found that method is called.
To add new supported methods extend the JSONRPC class and call [method process_action] on your subclass.
[param action]: The action to be run, as a Dictionary in the form of a JSON-RPC request or notification.
*/
func (self Instance) ProcessAction(action any) any { //gd:JSONRPC.process_action
	return any(class(self).ProcessAction(gd.NewVariant(action), false).Interface())
}
func (self Instance) ProcessString(action string) string { //gd:JSONRPC.process_string
	return string(class(self).ProcessString(gd.NewString(action)).String())
}

/*
Returns a dictionary in the form of a JSON-RPC request. Requests are sent to a server with the expectation of a response. The ID field is used for the server to specify which exact request it is responding to.
- [param method]: Name of the method being called.
- [param params]: An array or dictionary of parameters being passed to the method.
- [param id]: Uniquely identifies this request. The server is expected to send a response with the same ID.
*/
func (self Instance) MakeRequest(method string, params any, id any) Request { //gd:JSONRPC.make_request
	return Request(gd.DictionaryAs[Request](class(self).MakeRequest(gd.NewString(method), gd.NewVariant(params), gd.NewVariant(id))))
}

/*
When a server has received and processed a request, it is expected to send a response. If you did not want a response then you need to have sent a Notification instead.
- [param result]: The return value of the function which was called.
- [param id]: The ID of the request this response is targeted to.
*/
func (self Instance) MakeResponse(result any, id any) Response { //gd:JSONRPC.make_response
	return Response(gd.DictionaryAs[Response](class(self).MakeResponse(gd.NewVariant(result), gd.NewVariant(id))))
}

/*
Returns a dictionary in the form of a JSON-RPC notification. Notifications are one-shot messages which do not expect a response.
- [param method]: Name of the method being called.
- [param params]: An array or dictionary of parameters being passed to the method.
*/
func (self Instance) MakeNotification(method string, params any) Notification { //gd:JSONRPC.make_notification
	return Notification(gd.DictionaryAs[Notification](class(self).MakeNotification(gd.NewString(method), gd.NewVariant(params))))
}

/*
Creates a response which indicates a previous reply has failed in some way.
- [param code]: The error code corresponding to what kind of error this is. See the [enum ErrorCode] constants.
- [param message]: A custom message about this error.
- [param id]: The request this error is a response to.
*/
func (self Instance) MakeResponseError(code int, message string) ResponseError { //gd:JSONRPC.make_response_error
	return ResponseError(gd.DictionaryAs[ResponseError](class(self).MakeResponseError(gd.Int(code), gd.NewString(message), gd.NewVariant(gd.NewVariant(([1]any{}[0]))))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.JSONRPC

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("JSONRPC"))
	casted := Instance{*(*gdclass.JSONRPC)(unsafe.Pointer(&object))}
	return casted
}

//go:nosplit
func (self class) SetScope(scope gd.String, target [1]gd.Object) { //gd:JSONRPC.set_scope
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scope))
	callframe.Arg(frame, pointers.Get(target[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_set_scope, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Given a Dictionary which takes the form of a JSON-RPC request: unpack the request and run it. Methods are resolved by looking at the field called "method" and looking for an equivalently named function in the JSONRPC object. If one is found that method is called.
To add new supported methods extend the JSONRPC class and call [method process_action] on your subclass.
[param action]: The action to be run, as a Dictionary in the form of a JSON-RPC request or notification.
*/
//go:nosplit
func (self class) ProcessAction(action gd.Variant, recurse bool) gd.Variant { //gd:JSONRPC.process_action
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, recurse)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_process_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) ProcessString(action gd.String) gd.String { //gd:JSONRPC.process_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_process_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a dictionary in the form of a JSON-RPC request. Requests are sent to a server with the expectation of a response. The ID field is used for the server to specify which exact request it is responding to.
- [param method]: Name of the method being called.
- [param params]: An array or dictionary of parameters being passed to the method.
- [param id]: Uniquely identifies this request. The server is expected to send a response with the same ID.
*/
//go:nosplit
func (self class) MakeRequest(method gd.String, params gd.Variant, id gd.Variant) Dictionary.Any { //gd:JSONRPC.make_request
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(method))
	callframe.Arg(frame, pointers.Get(params))
	callframe.Arg(frame, pointers.Get(id))
	var r_ret = callframe.Ret[Dictionary.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_make_request, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
When a server has received and processed a request, it is expected to send a response. If you did not want a response then you need to have sent a Notification instead.
- [param result]: The return value of the function which was called.
- [param id]: The ID of the request this response is targeted to.
*/
//go:nosplit
func (self class) MakeResponse(result gd.Variant, id gd.Variant) Dictionary.Any { //gd:JSONRPC.make_response
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(result))
	callframe.Arg(frame, pointers.Get(id))
	var r_ret = callframe.Ret[Dictionary.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_make_response, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a dictionary in the form of a JSON-RPC notification. Notifications are one-shot messages which do not expect a response.
- [param method]: Name of the method being called.
- [param params]: An array or dictionary of parameters being passed to the method.
*/
//go:nosplit
func (self class) MakeNotification(method gd.String, params gd.Variant) Dictionary.Any { //gd:JSONRPC.make_notification
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(method))
	callframe.Arg(frame, pointers.Get(params))
	var r_ret = callframe.Ret[Dictionary.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_make_notification, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a response which indicates a previous reply has failed in some way.
- [param code]: The error code corresponding to what kind of error this is. See the [enum ErrorCode] constants.
- [param message]: A custom message about this error.
- [param id]: The request this error is a response to.
*/
//go:nosplit
func (self class) MakeResponseError(code gd.Int, message gd.String, id gd.Variant) Dictionary.Any { //gd:JSONRPC.make_response_error
	var frame = callframe.New()
	callframe.Arg(frame, code)
	callframe.Arg(frame, pointers.Get(message))
	callframe.Arg(frame, pointers.Get(id))
	var r_ret = callframe.Ret[Dictionary.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_make_response_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsJSONRPC() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJSONRPC() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("JSONRPC", func(ptr gd.Object) any { return [1]gdclass.JSONRPC{*(*gdclass.JSONRPC)(unsafe.Pointer(&ptr))} })
}

type ErrorCode = gdclass.JSONRPCErrorCode //gd:JSONRPC.ErrorCode

const (
	/*The request could not be parsed as it was not valid by JSON standard ([method JSON.parse] failed).*/
	ParseError ErrorCode = -32700
	/*A method call was requested but the request's format is not valid.*/
	InvalidRequest ErrorCode = -32600
	/*A method call was requested but no function of that name existed in the JSONRPC subclass.*/
	MethodNotFound ErrorCode = -32601
	/*A method call was requested but the given method parameters are not valid. Not used by the built-in JSONRPC.*/
	InvalidParams ErrorCode = -32602
	/*An internal error occurred while processing the request. Not used by the built-in JSONRPC.*/
	InternalError ErrorCode = -32603
)

type Request struct {
	Method string      `gd:"method"`
	Params interface{} `gd:"params"`
	ID     string      `gd:"id"`
}
type Response struct {
	Result interface{} `gd:"result"`
	ID     string      `gd:"id"`
}
type Notification struct {
	Method string      `gd:"method"`
	Params interface{} `gd:"params"`
}
type ResponseError struct {
	Code    int    `gd:"code"`
	Message string `gd:"message"`
	ID      int    `gd:"id"`
}
