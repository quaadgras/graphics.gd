/*
[gdscript]
val uri = "content://com.android..." # URI of the selected file or folder.
val persist = true # Set to false to release the persistable permission.
var android_runtime = Engine.get_singleton("AndroidRuntime")
android_runtime.updatePersistableUriPermission(uri, persist)
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Engine"
	"graphics.gd/variant/Object"
)

func DisplayServer_FileDialogShow() {
	var uri = "content://com.android..." // URI of the selected file or folder.
	var persist = true                   // Set to false to release the persistable permission.
	var android_runtime = Engine.GetSingleton("AndroidRuntime")
	Object.Call(android_runtime, "updatePersistableUriPermission", uri, persist)
}
