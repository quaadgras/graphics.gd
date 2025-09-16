/*
extends Node

var _my_js_callback = JavaScriptBridge.create_callback(myCallback) # This reference must be kept
var console = JavaScriptBridge.get_interface("console")

func _init():
	var buf = JavaScriptBridge.create_object("ArrayBuffer", 10) # new ArrayBuffer(10)
	print(buf) # Prints [JavaScriptObject:OBJECT_ID]
	var uint8arr = JavaScriptBridge.create_object("Uint8Array", buf) # new Uint8Array(buf)
	uint8arr[1] = 255
	prints(uint8arr[1], uint8arr.byteLength) # Prints "255 10"

	# Prints "Uint8Array(10) [ 0, 255, 0, 0, 0, 0, 0, 0, 0, 0 ]" in the browser's console.
	console.log(uint8arr)

	# Equivalent of JavaScriptBridge: Array.from(uint8arr).forEach(myCallback)
	JavaScriptBridge.get_interface("Array").from(uint8arr).forEach(_my_js_callback)

func myCallback(args):
	# Will be called with the parameters passed to the "forEach" callback
	# [0, 0, [JavaScriptObject:1173]]
	# [255, 1, [JavaScriptObject:1173]]
	# ...
	# [0, 9, [JavaScriptObject:1180]]
	print(args)
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/JavaScriptBridge"
	"graphics.gd/classdb/JavaScriptObject"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Object"
)

type MyJavaScriptObjects struct {
	Node.Extension[MyJavaScriptObjects]

	my_js_callback JavaScriptObject.Instance
	console        Object.Instance
}

func (n MyJavaScriptObjects) Init() {
	n.my_js_callback = JavaScriptBridge.CreateCallback(func(args []any) any {
		fmt.Println(args)
		return nil
	})

	var buf = JavaScriptBridge.CreateObject("ArrayBuffer", 10)
	fmt.Println(buf)
	var uint8arr = JavaScriptBridge.CreateObject("Uint8Array", buf).(Object.Instance)
	Object.SetIndex(uint8arr, 1, 255)
	fmt.Println(Object.Index(uint8arr, 1), Object.Get(uint8arr, "byteLength"))

	Object.Call(n.console, "log", uint8arr)

	Object.Call(Object.Call(JavaScriptBridge.GetInterface("Array"), "from").(Object.Instance), "forEach", n.my_js_callback)
}
