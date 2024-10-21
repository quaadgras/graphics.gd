package JSONRPC

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[url=https://www.jsonrpc.org/]JSON-RPC[/url] is a standard which wraps a method call in a [JSON] object. The object has a particular structure and identifies which method is called, the parameters to that function, and carries an ID to keep track of responses. This class implements that standard on top of [Dictionary]; you will have to convert between a [Dictionary] and [JSON] with other functions.

*/
type Simple [1]classdb.JSONRPC
func (self Simple) SetScope(scope string, target gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScope(gc.String(scope), target)
}
func (self Simple) ProcessAction(action gd.Variant, recurse bool) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).ProcessAction(gc, action, recurse))
}
func (self Simple) ProcessString(action string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).ProcessString(gc, gc.String(action)).String())
}
func (self Simple) MakeRequest(method string, params gd.Variant, id gd.Variant) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).MakeRequest(gc, gc.String(method), params, id))
}
func (self Simple) MakeResponse(result gd.Variant, id gd.Variant) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).MakeResponse(gc, result, id))
}
func (self Simple) MakeNotification(method string, params gd.Variant) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).MakeNotification(gc, gc.String(method), params))
}
func (self Simple) MakeResponseError(code int, message string, id gd.Variant) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).MakeResponseError(gc, gd.Int(code), gc.String(message), id))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.JSONRPC
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetScope(scope gd.String, target gd.Object)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(scope))
	callframe.Arg(frame, mmm.End(target.AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSONRPC.Bind_set_scope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Given a Dictionary which takes the form of a JSON-RPC request: unpack the request and run it. Methods are resolved by looking at the field called "method" and looking for an equivalently named function in the JSONRPC object. If one is found that method is called.
To add new supported methods extend the JSONRPC class and call [method process_action] on your subclass.
[param action]: The action to be run, as a Dictionary in the form of a JSON-RPC request or notification.
*/
//go:nosplit
func (self class) ProcessAction(ctx gd.Lifetime, action gd.Variant, recurse bool) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	callframe.Arg(frame, recurse)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSONRPC.Bind_process_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) ProcessString(ctx gd.Lifetime, action gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSONRPC.Bind_process_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
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
func (self class) MakeRequest(ctx gd.Lifetime, method gd.String, params gd.Variant, id gd.Variant) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(method))
	callframe.Arg(frame, mmm.Get(params))
	callframe.Arg(frame, mmm.Get(id))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSONRPC.Bind_make_request, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
When a server has received and processed a request, it is expected to send a response. If you did not want a response then you need to have sent a Notification instead.
- [param result]: The return value of the function which was called.
- [param id]: The ID of the request this response is targeted to.
*/
//go:nosplit
func (self class) MakeResponse(ctx gd.Lifetime, result gd.Variant, id gd.Variant) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(result))
	callframe.Arg(frame, mmm.Get(id))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSONRPC.Bind_make_response, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a dictionary in the form of a JSON-RPC notification. Notifications are one-shot messages which do not expect a response.
- [param method]: Name of the method being called.
- [param params]: An array or dictionary of parameters being passed to the method.
*/
//go:nosplit
func (self class) MakeNotification(ctx gd.Lifetime, method gd.String, params gd.Variant) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(method))
	callframe.Arg(frame, mmm.Get(params))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSONRPC.Bind_make_notification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
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
func (self class) MakeResponseError(ctx gd.Lifetime, code gd.Int, message gd.String, id gd.Variant) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, code)
	callframe.Arg(frame, mmm.Get(message))
	callframe.Arg(frame, mmm.Get(id))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.JSONRPC.Bind_make_response_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsJSONRPC() Expert { return self[0].AsJSONRPC() }


//go:nosplit
func (self Simple) AsJSONRPC() Simple { return self[0].AsJSONRPC() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("JSONRPC", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ErrorCode = classdb.JSONRPCErrorCode

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
