/*
class PrintProxy:
	func println(content: String) -> void:
		print(content)

var print_proxy = PrintProxy.new()
var printer_object = JavaClassWrapper.create_proxy(print_proxy, ["android.util.Printer"])
printer_object.println("Hello Godot World!")
*/

package main

import (
	"graphics.gd/classdb/JavaClassWrapper"
	"graphics.gd/variant/Object"
)

// printProxy stands in for a Godot object implementing the println(String) method.
func ExampleJavaClassWrapperCreateProxy(printProxy Object.Instance) {
	var printerObject = JavaClassWrapper.CreateProxy(printProxy, []string{"android.util.Printer"})
	Object.Call(printerObject, "println", "Hello Godot World!") // dynamic Java call.
	_ = printerObject
}
