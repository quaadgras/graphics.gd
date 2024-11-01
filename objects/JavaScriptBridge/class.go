package JavaScriptBridge

import "unsafe"
import "sync"
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
The JavaScriptBridge singleton is implemented only in the Web export. It's used to access the browser's JavaScript context. This allows interaction with embedding pages or calling third-party JavaScript APIs.
[b]Note:[/b] This singleton can be disabled at build-time to improve security. By default, the JavaScriptBridge singleton is enabled. Official export templates also have the JavaScriptBridge singleton enabled. See [url=$DOCS_URL/contributing/development/compiling/compiling_for_web.html]Compiling for the Web[/url] in the documentation for more information.
*/
var self objects.JavaScriptBridge
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.JavaScriptBridge)
	self = *(*objects.JavaScriptBridge)(unsafe.Pointer(&obj))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.JavaScriptBridge

func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Execute the string [param code] as JavaScript code within the browser window. This is a call to the actual global JavaScript function [code skip-lint]eval()[/code].
If [param use_global_execution_context] is [code]true[/code], the code will be evaluated in the global execution context. Otherwise, it is evaluated in the execution context of a function within the engine's runtime environment.
*/

/*
Returns an interface to a JavaScript object that can be used by scripts. The [param interface] must be a valid property of the JavaScript [code]window[/code]. The callback must accept a single [Array] argument, which will contain the JavaScript [code]arguments[/code]. See [JavaScriptObject] for usage.
*/

/*
Creates a reference to a [Callable] that can be used as a callback by JavaScript. The reference must be kept until the callback happens, or it won't be called at all. See [JavaScriptObject] for usage.
*/

/*
Prompts the user to download a file containing the specified [param buffer]. The file will have the given [param name] and [param mime] type.
[b]Note:[/b] The browser may override the [url=https://en.wikipedia.org/wiki/Media_type]MIME type[/url] provided based on the file [param name]'s extension.
[b]Note:[/b] Browsers might block the download if [method download_buffer] is not being called from a user interaction (e.g. button click).
[b]Note:[/b] Browsers might ask the user for permission or block the download if multiple download requests are made in a quick succession.
*/

/*
Returns [code]true[/code] if a new version of the progressive web app is waiting to be activated.
[b]Note:[/b] Only relevant when exported as a Progressive Web App.
*/

/*
Performs the live update of the progressive web app. Forcing the new version to be installed and the page to be reloaded.
[b]Note:[/b] Your application will be [b]reloaded in all browser tabs[/b].
[b]Note:[/b] Only relevant when exported as a Progressive Web App and [method pwa_needs_update] returns [code]true[/code].
*/

/*
Force synchronization of the persistent file system (when enabled).
[b]Note:[/b] This is only useful for modules or extensions that can't use [FileAccess] to write files.
*/

func OnPwaUpdateAvailable(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("pwa_update_available"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("JavaScriptBridge", func(ptr gd.Object) any { return classdb.JavaScriptBridge(ptr) })
}
