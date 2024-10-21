package gd

var Global API


// GarbageCollector TODO, idea is to have a thread-safe lifetime that
// collects anything older than n-1 frames
func GarbageCollector() Lifetime {
	return NewLifetime(&Global)
}
