package gd

import (
	"grow.graphics/gd/internal/mmm"
)

var Global API
var GC = Lifetime{
	Lifetime: mmm.NewThreadSafeLifetime(),
	API:      &Global,
}
var Static = Lifetime{
	Lifetime: mmm.NewThreadSafeLifetime(),
	API:      &Global,
}

// GarbageCollector TODO, idea is to have a thread-safe lifetime that
// collects anything older than n-1 frames
func GarbageCollector() Lifetime {
	return GC
}
