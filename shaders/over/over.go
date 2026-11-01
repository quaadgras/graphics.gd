/*
Package over provides loops that run on the GPU, for use with a Go for statement.

	pos := start
	for i := range over.Int(steps) {
		pos = vec3.Add(pos, step)
		if float.Gt(density(pos), 0.5).Evaluate() {
			break
		}
	}

The body of the loop is written like any other Go loop: values assigned to
variables declared before the loop are carried over to the next iteration,
continue (inside of a branch on a condition evaluated with [bool.X.Evaluate])
skips the rest of the iteration, and break stops the loop. Values computed before
the loop are computed once, outside of it.

To find out which values are carried over from one iteration to the next, the
pipeline function is run more than once, with the loop body running a few
times, so the body must do the same thing each time it runs with the same
values. A loop cannot (yet) be nested inside of another one.

The functions in this package can only be called by pipeline functions, while
the shader is being compiled.
*/
package over

import (
	"iter"

	"graphics.gd/shaders/int"
	"graphics.gd/shaders/internal/gpu"
)

// Int loops n times on the GPU, where n can be a value that is only known while
// the shader is running. Each iteration receives its index, counting up from zero.
func Int[N gpu.AnyInt](n N) iter.Seq[int.X] {
	count := gpu.NewInt(n)
	return func(yield func(int.X) bool) {
		if gpu.Iterate == nil {
			panic("shaders: over.Int can only be used while a shader is being compiled")
		}
		gpu.Iterate(count, func(i gpu.Int) bool { return yield(int.X(i)) })
	}
}
