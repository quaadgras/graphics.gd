/*
var Intent = JavaClassWrapper.wrap("android.content.Intent")
var intent = Intent.Intent()
*/

package main

import (
	"graphics.gd/classdb/JavaClassWrapper"
	"graphics.gd/variant/Object"
)

func ExampleJavaClassWrapperWrap() {
	var Intent = JavaClassWrapper.Wrap("android.content.Intent")
	var intent = Object.Call(Intent, "Intent")
	_ = intent
}
