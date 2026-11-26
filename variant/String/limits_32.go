//go:build arm || i386 || (gd && wasm.gc)

package String

import "math"

const maxSafeInt = math.MaxInt32
