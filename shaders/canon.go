package shaders

import (
	"fmt"
	"reflect"
	"strings"

	dsl "graphics.gd/shaders/internal/gpu"
)

// canonicalizer numbers expressions by their structure, so that the same
// computation gets the same number, even when it was built more than once (as
// happens when a pipeline function is run once for each way its branches go).
type canonicalizer struct {
	byIdentity map[*byte]int
	bySig      map[string]int
	reps       []dsl.Evaluator // the first expression seen with each number.
	loops      []*dsl.Scope    // the outermost loop that each number was created inside of.
	blocks     map[*dsl.Block]string
	blockReps  map[string]*dsl.Block
}

func newCanonicalizer() *canonicalizer {
	return &canonicalizer{
		byIdentity: make(map[*byte]int),
		bySig:      make(map[string]int),
		blocks:     make(map[*dsl.Block]string),
		blockReps:  make(map[string]*dsl.Block),
	}
}

// canon is reset for each shader that is compiled.
var canon = newCanonicalizer()

var untyped = reflect.TypeFor[dsl.Expression]()

// of returns the number of e.
func (c *canonicalizer) of(e dsl.Evaluator) int {
	var id *byte
	if hi, ok := e.(hasIdentity); ok {
		if id = hi.Identity(); id != nil {
			if n, ok := c.byIdentity[id]; ok {
				return n
			}
		}
	}
	sig := c.signature(e)
	loop := dsl.ScopeOf(e).InnermostLoop()
	n, ok := c.bySig[sig]
	if !ok {
		n = len(c.reps)
		c.bySig[sig] = n
		c.reps = append(c.reps, e)
		c.loops = append(c.loops, loop)
	} else {
		if reflect.TypeOf(c.reps[n]) == untyped && reflect.TypeOf(e) != untyped {
			c.reps[n] = e // prefer a typed representative, so that it can be stored in a variable.
		}
		// The same computation created outside of a loop can be computed
		// outside of it.
		c.loops[n] = commonLoop(c.loops[n], loop)
	}
	if id != nil {
		c.byIdentity[id] = n
	}
	return n
}

func (c *canonicalizer) ref(e dsl.Evaluator) string {
	if e == nil {
		return "-"
	}
	return fmt.Sprint(c.of(e))
}

func resolve(e dsl.Evaluator) dsl.Evaluator {
	if r := dsl.Evaluate(e); r != nil {
		return r
	}
	return e
}

func (c *canonicalizer) signature(e dsl.Evaluator) string {
	switch v := resolve(e).(type) {
	case dsl.Identifier:
		return "i" + string(v)
	case dsl.Operation:
		return fmt.Sprintf("o%s(%s,%s)", v.Op, c.ref(v.A), c.ref(v.B))
	case dsl.Select:
		return fmt.Sprintf("s%s(%s)", v.Field, c.ref(v.Value))
	case dsl.Index:
		return fmt.Sprintf("x%d(%s)", v.Index, c.ref(v.Value))
	case dsl.Ternary:
		return fmt.Sprintf("t(%s,%s,%s)", c.ref(v.If), c.ref(v.A), c.ref(v.B))
	case dsl.FunctionCall:
		var b strings.Builder
		fmt.Fprintf(&b, "f%s(", v.Name)
		for _, arg := range v.Args {
			b.WriteString(c.ref(arg))
			b.WriteByte(',')
		}
		return b.String() + ")"
	case *dsl.Local:
		if v.Block.Kind == dsl.BlockIf {
			// The results of an if are numbered by what it computes, so that the
			// same branch reached in more than one run is the same variable.
			return fmt.Sprintf("r%d|%s", v.Slot, c.block(v.Block))
		}
		return fmt.Sprintf("l%p", v)
	default:
		var b strings.Builder
		b.WriteString("k" + glslName(reflect.TypeOf(v)) + "{")
		c.fields(reflect.ValueOf(v), &b)
		return b.String() + "}"
	}
}

// fields writes the components of a literal value.
func (c *canonicalizer) fields(value reflect.Value, b *strings.Builder) {
	eachLeaf(value, func(ev dsl.Evaluator) {
		b.WriteString(c.ref(ev))
		b.WriteByte(',')
	}, func(v reflect.Value) {
		fmt.Fprint(b, v.Interface())
		b.WriteByte(',')
	})
}

// block returns the signature of an if block.
func (c *canonicalizer) block(blk *dsl.Block) string {
	if sig, ok := c.blocks[blk]; ok {
		return sig
	}
	var b strings.Builder
	fmt.Fprintf(&b, "if(%s)", c.ref(blk.Cond))
	for _, body := range blk.Bodies {
		b.WriteString("{")
		for _, next := range body.Next {
			b.WriteString(c.ref(next))
			b.WriteByte(',')
		}
		b.WriteString("}")
	}
	sig := b.String()
	c.blocks[blk] = sig
	if _, ok := c.blockReps[sig]; !ok {
		c.blockReps[sig] = blk
	}
	return sig
}

// rep returns the block that represents all of the blocks that compute the
// same thing as blk.
func (c *canonicalizer) rep(blk *dsl.Block) *dsl.Block {
	if blk.Kind != dsl.BlockIf {
		return blk
	}
	return c.blockReps[c.block(blk)]
}

// list returns the numbers of values, joined into a string.
func (c *canonicalizer) list(values []dsl.Evaluator) string {
	var b strings.Builder
	for _, v := range values {
		b.WriteString(c.ref(v))
		b.WriteByte(',')
	}
	return b.String()
}

// glslName returns the GLSL type of a Go type, so that literals with different
// Go types that are the same on the GPU are numbered the same.
func glslName(t reflect.Type) (name string) {
	defer func() {
		if recover() != nil {
			name = t.String()
		}
	}()
	return glslTypeFor(t)
}

// commonLoop returns the innermost loop body that contains both of the loop
// bodies a and b (nil for neither). It must be called before the parents of
// the loop bodies are moved to where the loops are placed.
func commonLoop(a, b *dsl.Scope) *dsl.Scope {
	return commonScope(a, b).InnermostLoop()
}
