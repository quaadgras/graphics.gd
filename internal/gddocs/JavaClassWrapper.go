/*
var LocalDateTime = JavaClassWrapper.wrap("java.time.LocalDateTime")
var DateTimeFormatter = JavaClassWrapper.wrap("java.time.format.DateTimeFormatter")

var datetime = LocalDateTime.now()
var formatter = DateTimeFormatter.ofPattern("dd-MM-yyyy HH:mm:ss")

print(datetime.format(formatter))
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/JavaClassWrapper"
	"graphics.gd/variant/Object"
)

func ExampleJavaClassWrapper() {
	var LocalDateTime = JavaClassWrapper.Wrap("java.time.LocalDateTime")
	var DateTimeFormatter = JavaClassWrapper.Wrap("java.time.format.DateTimeFormatter")
	var datetime = Object.Call(LocalDateTime, "now")
	var formatter = Object.Call(DateTimeFormatter, "ofPattern", "dd-MM-yyyy HH:mm:ss")
	fmt.Println(Object.Call(datetime.(Object.Instance), "format", formatter))
}
