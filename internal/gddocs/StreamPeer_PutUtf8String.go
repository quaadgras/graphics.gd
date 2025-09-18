/*
[gdscript]
put_data("Hello world".to_utf8_buffer())
[/gdscript]
[csharp]
PutData("Hello World".ToUtf8Buffer());
[/csharp]
*/

package main

func StreamPeer_PutUtf8String() {
	streamPeer.PutData([]byte("Hello World"))
}
