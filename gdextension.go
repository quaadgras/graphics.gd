//go:generate go run ./internal/tool/generate
//go:generate go run ./internal/tool/generate/v2
//go:generate go fmt ./...
package gdunsafe

import (
	"math/rand"
	"time"

	"graphics.gd/internal/gdextension"
)

type Int = int64
type Variant = gdextension.Variant
type CallError = gdextension.CallError

// just a placeholder for functions that don't need to be implemented
// as they are already available in the Go standard library.

func randomize() { //gd:randomize
	rand.Seed(time.Now().UnixNano())
}

func seed(s int) { //gd:seed
	rand.Seed(int64(s))
}

func rand_from_seed(seed int) *rand.Rand { //gd:rand_from_seed
	return rand.New(rand.NewSource(int64(seed)))
}

func weakref(v any) any { return v } //gd:weakref
