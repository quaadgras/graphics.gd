/*
In order to make graphics.gd goroutine-safe:

 1. New objects created on non-main threads should allocate a Go pointer and use a runtime cleanup for GC.
 2. Objects created on the main thread can continue to work as they do now.
 3. Objects borrowed from Godot on another thread, should always include an ObjectID which must be checked
    whenever the object is used.

The gd command should include a compile-time race-detector 'gd run -race' that ensures data-race-free Go
rules are followed:

 1. global variables may not be mutated after initialization.
 2. unsyncronised reference types may not be passed via channels, nor captured by newly created goroutines.
 3. all goroutines must immediately defer a call that recovers.
*/
package threadsafe
