/*
var cb = func (content: String) -> void:
	print(content)
var callback = JavaClassWrapper.create_sam_callback("android.util.Printer", cb)
callback.println("Hello Godot World!")
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/JavaClassWrapper"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Object"
)

func ExampleJavaClassWrapperCreateSamCallback() {
	cb := func(content string) { fmt.Println(content) }
	var callback = JavaClassWrapper.CreateSamCallback("android.util.Printer", Callable.New(cb))
	Object.Call(callback, "println", "Hello Godot World!") // dynamic Java call.
	_ = callback
}
