package gpu

import "reflect"

// Scope is a block of statements on the GPU, such as the body of a loop or a
// branch of an if. Expressions remember the scope they were created in, so that
// a value computed outside of a loop body (in Go) is computed once, outside of
// the loop (on the GPU).
type Scope struct {
	Parent *Scope
	Loop   bool // the scope runs once for each iteration of a loop.
}

// current is the scope that new expressions are created in. Shaders are
// compiled on one goroutine at a time.
var current *Scope

// CurrentScope returns the scope that new expressions are created in.
func CurrentScope() *Scope { return current }

// Within calls fn with new expressions created in scope.
func Within(scope *Scope, fn func()) {
	saved := current
	current = scope
	defer func() { current = saved }()
	fn()
}

// ScopeOf returns the scope that e was created in.
func ScopeOf(e Evaluator) *Scope {
	if s, ok := e.(interface{ getScope() *Scope }); ok {
		return s.getScope()
	}
	return nil
}

// InnermostLoop returns the body of the innermost loop that s is inside of, or
// nil if s is not inside of a loop.
func (s *Scope) InnermostLoop() *Scope {
	for ; s != nil; s = s.Parent {
		if s.Loop {
			return s
		}
	}
	return nil
}

// Contains reports whether other is s or nested inside of s.
func (s *Scope) Contains(other *Scope) bool {
	for ; other != nil; other = other.Parent {
		if other == s {
			return true
		}
	}
	return s == nil
}

type BlockKind int

const (
	BlockIf    BlockKind = iota // Bodies[0] when Cond, else Bodies[1].
	BlockFor                    // Bodies[0] Count times, counting with Index.
	BlockWhile                  // Bodies[0] while Cond (evaluated in Bodies[0].Scope).
)

// Block is a statement with nested statements, a loop or an if. The values it
// produces are read through its State.
type Block struct {
	Kind    BlockKind
	Created *Scope // the scope that was current when the block was created.

	Cond  Evaluator
	Count Evaluator
	Index *Local // the loop counter, as read inside of the body.

	// State holds the variables that the block assigns, as read after the
	// block. A loop initializes them with Init, its body reads them through
	// Params and each iteration assigns the body's Next to them. Each branch of
	// an if assigns its Next to them.
	State  []*Local
	Params []*Local
	Init   []Evaluator
	Bodies []Body
}

// Body is a nested block of statements. The parent of its Scope is decided
// when the shader is compiled, as it depends on where the block's values are
// used.
type Body struct {
	Scope *Scope
	Next  []Evaluator
	Break Evaluator // for a loop, whether to stop after this iteration.
}

// Local is a variable declared by a [Block].
type Local struct {
	Block *Block
	Type  reflect.Type
	Slot  int
	Param bool // read inside of the body of the block, rather than after it.
}

func (l *Local) evaluate() Evaluator { return l }

// Decide is called by Evaluate on a GPU condition, to find out which way the
// condition goes in the current run of a pipeline function.
var Decide func(cond Evaluator) bool

// IsLiteral reports whether e is a constant written in Go.
func IsLiteral(e Evaluator) bool {
	if ex, ok := e.(interface{ isLiteral() bool }); ok {
		return ex.isLiteral()
	}
	return false
}

func (e Expression) isLiteral() bool { return e.literal }

// Cell is an expression created inside of a loop body while it is being
// traced. Its Target can be replaced after the loop, so that values the loop
// body assigned to Go variables can be made to read the variables of the loop.
type Cell struct {
	Target    Evaluator
	Trace     *Trace
	Iteration int // counting from 1.
	Seq       int // the order in which it was created, within its iteration.
}

func (c *Cell) evaluate() Evaluator { return c.Target.evaluate() }

// Trace records the expressions created by each iteration of a loop body.
type Trace struct {
	Iteration int
	Cells     [][]*Cell
}

var tracing *Trace

// StartIteration traces the expressions created from now on, as part of the
// next iteration of t.
func StartIteration(t *Trace) {
	t.Cells = append(t.Cells, nil)
	t.Iteration = len(t.Cells)
	tracing = t
}

// StopIteration stops tracing.
func StopIteration() { tracing = nil }

// Tracing reports whether a loop body is being traced.
func Tracing() bool { return tracing != nil }

var stamps int

// ResetStamps starts numbering expressions from the beginning again.
func ResetStamps() { stamps = 0 }

func literal() Expression {
	if tracing != nil {
		return Expression{literal: true}
	}
	stamps++
	return Expression{literal: true, stamp: stamps}
}

// Info describes how an expression was created.
type Info struct {
	Cell     *Cell
	Stamp    int
	Identity *byte
	Literal  bool
}

// InfoOf returns how e was created, if it is an expression.
func InfoOf(e Evaluator) (Info, bool) {
	if x, ok := e.(interface{ info() Info }); ok {
		return x.info(), true
	}
	return Info{}, false
}

// Iterate is called by the over package, to run a loop body on the GPU.
// It calls yield with the index of the iteration.
var Iterate func(count Evaluator, yield func(index Int) bool)
