package gd

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"grow.graphics/gd/internal/mmm"
)

var Global API

var allocsRing [4]atomic.Int64
var globalLifetimeRing = [4]Lifetime{
	{
		Lifetime: mmm.NewThreadSafeLifetime().WithAllocationCounter(&allocsRing[0]),
		API:      &Global,
	},
	{
		Lifetime: mmm.NewThreadSafeLifetime().WithAllocationCounter(&allocsRing[1]),
		API:      &Global,
	},
	{
		Lifetime: mmm.NewThreadSafeLifetime().WithAllocationCounter(&allocsRing[2]),
		API:      &Global,
	},
	{
		Lifetime: mmm.NewThreadSafeLifetime().WithAllocationCounter(&allocsRing[3]),
		API:      &Global,
	},
}
var currentGlobalLifetime atomic.Int32

var allocs atomic.Int64
var lastGC time.Time
var lastGCns atomic.Int64

var DEBUG = os.Getenv("DEBUG") != ""

func init() {
	ticker := time.NewTicker(time.Second / 60)
	go func() {
		for range ticker.C {
			index := currentGlobalLifetime.Load()
			next := (index + 1) % int32(len(allocsRing))
			prev := (index + int32(len(allocsRing)-1)) % int32(len(allocsRing))
			allocs := allocsRing[index].Load()
			if allocs > 0 && currentGlobalLifetime.CompareAndSwap(index, next) {
				globalLifetimeRing[prev].End()
				if DEBUG {
					fmt.Printf("GC cleaned up %d allocs \n", allocs)
				}
				allocsRing[prev].Store(0)
				globalLifetimeRing[prev].Lifetime = mmm.NewThreadSafeLifetime().WithAllocationCounter(&allocsRing[prev])
				lastGC = time.Now()
				lastGCns.Store(lastGC.UnixNano())
			}
		}
	}()
}

// GarbageCollector TODO, idea is to have a thread-safe lifetime that
// collects anything older than n-1 frames
func GarbageCollector() Lifetime {
	index := currentGlobalLifetime.Load()
	return globalLifetimeRing[index%int32(len(globalLifetimeRing))]
}
