package Expression

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
An expression can be made of any arithmetic operation, built-in math function call, method call of a passed instance, or built-in type construction call.
An example expression text using the built-in math functions could be [code]sqrt(pow(3, 2) + pow(4, 2))[/code].
In the following example we use a [LineEdit] node to write our expression and show the result.
[codeblocks]
[gdscript]
var expression = Expression.new()

func _ready():

	$LineEdit.text_submitted.connect(self._on_text_submitted)

func _on_text_submitted(command):

	var error = expression.parse(command)
	if error != OK:
	    print(expression.get_error_text())
	    return
	var result = expression.execute()
	if not expression.has_execute_failed():
	    $LineEdit.text = str(result)

[/gdscript]
[csharp]
private Expression _expression = new Expression();

public override void _Ready()

	{
	    GetNode<LineEdit>("LineEdit").TextSubmitted += OnTextEntered;
	}

private void OnTextEntered(string command)

	{
	    Error error = _expression.Parse(command);
	    if (error != Error.Ok)
	    {
	        GD.Print(_expression.GetErrorText());
	        return;
	    }
	    Variant result = _expression.Execute();
	    if (!_expression.HasExecuteFailed())
	    {
	        GetNode<LineEdit>("LineEdit").Text = result.ToString();
	    }
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]classdb.Expression

/*
Parses the expression and returns an [enum Error] code.
You can optionally specify names of variables that may appear in the expression with [param input_names], so that you can bind them when it gets executed.
*/
func (self Instance) Parse(expression string) gd.Error {
	return gd.Error(class(self).Parse(gd.NewString(expression), gd.NewPackedStringSlice(([1][]string{}[0]))))
}

/*
Executes the expression that was previously parsed by [method parse] and returns the result. Before you use the returned object, you should check if the method failed by calling [method has_execute_failed].
If you defined input variables in [method parse], you can specify their values in the inputs array, in the same order.
*/
func (self Instance) Execute() gd.Variant {
	return gd.Variant(class(self).Execute(([1]gd.Array{}[0]), ([1]gd.Object{}[0]), true, false))
}

/*
Returns [code]true[/code] if [method execute] has failed.
*/
func (self Instance) HasExecuteFailed() bool {
	return bool(class(self).HasExecuteFailed())
}

/*
Returns the error text if [method parse] or [method execute] has failed.
*/
func (self Instance) GetErrorText() string {
	return string(class(self).GetErrorText().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Expression

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Expression"))
	return Instance{classdb.Expression(object)}
}

/*
Parses the expression and returns an [enum Error] code.
You can optionally specify names of variables that may appear in the expression with [param input_names], so that you can bind them when it gets executed.
*/
//go:nosplit
func (self class) Parse(expression gd.String, input_names gd.PackedStringArray) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(expression))
	callframe.Arg(frame, pointers.Get(input_names))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Expression.Bind_parse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Executes the expression that was previously parsed by [method parse] and returns the result. Before you use the returned object, you should check if the method failed by calling [method has_execute_failed].
If you defined input variables in [method parse], you can specify their values in the inputs array, in the same order.
*/
//go:nosplit
func (self class) Execute(inputs gd.Array, base_instance gd.Object, show_error bool, const_calls_only bool) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(inputs))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(base_instance))
	callframe.Arg(frame, show_error)
	callframe.Arg(frame, const_calls_only)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Expression.Bind_execute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [method execute] has failed.
*/
//go:nosplit
func (self class) HasExecuteFailed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Expression.Bind_has_execute_failed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the error text if [method parse] or [method execute] has failed.
*/
//go:nosplit
func (self class) GetErrorText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Expression.Bind_get_error_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsExpression() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsExpression() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("Expression", func(ptr gd.Object) any { return classdb.Expression(ptr) })
}
