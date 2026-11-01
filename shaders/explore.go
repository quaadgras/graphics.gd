package shaders

import (
	"fmt"
	"reflect"

	dsl "graphics.gd/shaders/internal/gpu"
)

// maxRuns limits how many times a pipeline function (or loop body) is run to
// find each way that its branches can go.
const maxRuns = 4096

// explorer records the conditions evaluated during one run of a pipeline
// function, and decides which way they go.
type explorer struct {
	script  []bool // which way to go for the first conditions.
	conds   []dsl.Evaluator
	decided map[int]bool   // by condition number.
	scratch *canonicalizer // numbers conditions while running, see runWithLoops.
}

var explorers []*explorer

func init() {
	dsl.Decide = decide
}

func decide(cond dsl.Evaluator) bool {
	if len(explorers) == 0 {
		panic("shaders: Evaluate can only be called while a shader is being compiled")
	}
	if iterating != nil {
		return iterating.decide(cond)
	}
	// A condition that was already evaluated in this run goes the same way again.
	ex := explorers[len(explorers)-1]
	id := ex.scratch.of(cond)
	if decided, ok := ex.decided[id]; ok {
		return decided
	}
	k := len(ex.conds)
	ex.conds = append(ex.conds, cond)
	decided := true
	if k < len(ex.script) {
		decided = ex.script[k]
	}
	ex.decided[id] = decided
	return decided
}

// explore runs fn once for each way that the conditions it evaluates can go,
// and merges what each run returns into values that are computed by if
// statements wherever the runs differ. name names the i'th value, for errors.
func explore(fn func() []dsl.Evaluator, name func(i int) string) []dsl.Evaluator {
	runs := 0
	return exploreFrom(fn, name, nil, nil, &runs)
}

func exploreFrom(fn func() []dsl.Evaluator, name func(i int) string, script []bool, expect []int, runs *int) []dsl.Evaluator {
	if *runs++; *runs > maxRuns {
		panic(fmt.Sprintf("shaders: more than %d combinations of conditions to Evaluate, use bool.Select to choose between values that are cheap to compute", maxRuns))
	}
	ex := &explorer{script: script}
	explorers = append(explorers, ex)
	values := func() []dsl.Evaluator {
		defer func() { explorers = explorers[:len(explorers)-1] }()
		// Expressions made by a loop body can change while a pipeline function
		// runs (see loop.go), so conditions are numbered with a scratch
		// canonicalizer until the run is over.
		return runWithLoops(fn, func() {
			ex.conds = nil
			ex.decided = make(map[int]bool)
			ex.scratch = newCanonicalizer()
		})
	}()
	if len(ex.conds) < len(script) {
		panic("shaders: a pipeline function must evaluate the same conditions each time it runs")
	}
	ids := make([]int, len(ex.conds))
	for i, cond := range ex.conds {
		ids[i] = canon.of(cond)
		if i < len(expect) && ids[i] != expect[i] {
			panic("shaders: a pipeline function must evaluate the same conditions each time it runs")
		}
	}
	// This run took the true branch of every condition after the script, work
	// back from the last one, merging in the runs where it goes the other way.
	merged := values
	for k := len(ex.conds) - 1; k >= len(script); k-- {
		other := append(append([]bool(nil), script...), make([]bool, k+1-len(script))...)
		for i := len(script); i < k; i++ {
			other[i] = true
		}
		otherwise := exploreFrom(fn, name, other, ids[:k+1], runs)
		merged = merge(ex.conds[k], merged, otherwise, name)
	}
	return merged
}

// merger builds an if statement that computes the parts of two sets of values
// that differ, so that the parts they share are only computed once.
type merger struct {
	block *dsl.Block
	holes map[[2]int]dsl.Evaluator
	ifs   map[[2]*dsl.Block]*dsl.Block
}

func merge(cond dsl.Evaluator, then, otherwise []dsl.Evaluator, name func(i int) string) []dsl.Evaluator {
	m := &merger{
		block: &dsl.Block{
			Kind:    dsl.BlockIf,
			Created: dsl.CurrentScope(),
			Cond:    cond,
			Bodies:  []dsl.Body{{Scope: &dsl.Scope{}}, {Scope: &dsl.Scope{}}},
		},
		holes: make(map[[2]int]dsl.Evaluator),
		ifs:   make(map[[2]*dsl.Block]*dsl.Block),
	}
	merged := make([]dsl.Evaluator, len(then))
	for i := range then {
		if (then[i] == nil) != (otherwise[i] == nil) {
			panic(fmt.Sprintf("shaders: %s is only set on one side of a branch, set it on both", name(i)))
		}
		value, ok := m.generalize(then[i], otherwise[i])
		if !ok {
			panic(fmt.Sprintf("shaders: %s has a different type on each side of a branch", name(i)))
		}
		merged[i] = value
	}
	return merged
}

// generalize returns an expression with the structure that a and b share, with
// the parts where they differ read from the results of the if statement.
func (m *merger) generalize(a, b dsl.Evaluator) (dsl.Evaluator, bool) {
	if a == nil || b == nil {
		return a, a == nil && b == nil
	}
	if canon.of(a) == canon.of(b) {
		return a, true
	}
	ra, rb := resolve(a), resolve(b)
	var shaped dsl.Evaluator
	switch x := ra.(type) {
	case dsl.Operation:
		if y, ok := rb.(dsl.Operation); ok && x.Op == y.Op {
			A, okA := m.generalize(x.A, y.A)
			B, okB := m.generalize(x.B, y.B)
			if okA && okB {
				shaped = dsl.Operation{A: A, B: B, Op: x.Op}
			}
		}
	case dsl.Select:
		if y, ok := rb.(dsl.Select); ok && x.Field == y.Field {
			if v, ok := m.generalize(x.Value, y.Value); ok {
				shaped = dsl.Select{Value: v, Field: x.Field}
			}
		}
	case dsl.Index:
		if y, ok := rb.(dsl.Index); ok && x.Index == y.Index {
			if v, ok := m.generalize(x.Value, y.Value); ok {
				shaped = dsl.Index{Value: v, Index: x.Index}
			}
		}
	case dsl.Ternary:
		if y, ok := rb.(dsl.Ternary); ok {
			c, ok1 := m.generalize(x.If, y.If)
			A, ok2 := m.generalize(x.A, y.A)
			B, ok3 := m.generalize(x.B, y.B)
			if ok1 && ok2 && ok3 {
				shaped = dsl.Ternary{If: c, A: A, B: B}
			}
		}
	case dsl.FunctionCall:
		if y, ok := rb.(dsl.FunctionCall); ok && x.Name == y.Name && len(x.Args) == len(y.Args) {
			args := make([]dsl.Evaluator, len(x.Args))
			all := true
			for i := range args {
				var ok bool
				if args[i], ok = m.generalize(x.Args[i], y.Args[i]); !ok {
					all = false
					break
				}
			}
			if all {
				shaped = dsl.FunctionCall{Name: x.Name, Args: args}
			}
		}
	case *dsl.Local:
		// An if that each side reached with the same condition becomes one if,
		// after this one, so that it is not repeated inside of both branches.
		if y, ok := rb.(*dsl.Local); ok && x.Slot == y.Slot {
			if blk := m.generalizeIf(x.Block, y.Block); blk != nil {
				ptr := reflect.New(reflect.TypeOf(a))
				bind(ptr, blk.State[x.Slot])
				return ptr.Elem().Interface().(dsl.Evaluator), true
			}
		}
	case dsl.Identifier:
	default:
		// a literal, or a vector or matrix made out of components.
		if reflect.TypeOf(ra) == reflect.TypeOf(rb) {
			value := reflect.New(reflect.TypeOf(ra)).Elem()
			if m.components(value, reflect.ValueOf(ra), reflect.ValueOf(rb)) {
				return value.Interface().(dsl.Evaluator), true
			}
		}
	}
	if shaped != nil {
		return rebuild(a, shaped), true
	}
	return m.hole(a, b)
}

// generalizeIf returns an if block that computes what a computes on one side
// of the merged branch, and what b computes on the other side, or nil if they
// do not share the same condition and structure.
func (m *merger) generalizeIf(a, b *dsl.Block) *dsl.Block {
	if a.Kind != dsl.BlockIf || b.Kind != dsl.BlockIf || len(a.State) != len(b.State) {
		return nil
	}
	key := [2]*dsl.Block{a, b}
	if blk, ok := m.ifs[key]; ok {
		return blk
	}
	if canon.of(a.Cond) != canon.of(b.Cond) {
		return nil
	}
	cond := a.Cond
	blk := &dsl.Block{
		Kind:    dsl.BlockIf,
		Created: dsl.CurrentScope(),
		Cond:    cond,
		Bodies:  []dsl.Body{{Scope: &dsl.Scope{}}, {Scope: &dsl.Scope{}}},
	}
	for i := range a.Bodies {
		for slot := range a.State {
			next, ok := m.generalize(a.Bodies[i].Next[slot], b.Bodies[i].Next[slot])
			if !ok {
				return nil
			}
			blk.Bodies[i].Next = append(blk.Bodies[i].Next, next)
		}
	}
	for slot, l := range a.State {
		blk.State = append(blk.State, &dsl.Local{Block: blk, Type: l.Type, Slot: slot})
	}
	m.ifs[key] = blk
	return blk
}

// components generalizes each component of a and b into value.
func (m *merger) components(value, a, b reflect.Value) bool {
	switch a.Kind() {
	case reflect.Struct:
		for i := range a.NumField() {
			if !a.Type().Field(i).IsExported() {
				continue
			}
			if !m.component(value.Field(i), a.Field(i), b.Field(i)) {
				return false
			}
		}
		return true
	case reflect.Array:
		for i := range a.Len() {
			if !m.component(value.Index(i), a.Index(i), b.Index(i)) {
				return false
			}
		}
		return true
	default:
		value.Set(a)
		return a.Interface() == b.Interface()
	}
}

func (m *merger) component(value, a, b reflect.Value) bool {
	ea, ok := a.Interface().(dsl.Evaluator)
	if !ok {
		return m.components(value, a, b)
	}
	v, ok := m.generalize(ea, b.Interface().(dsl.Evaluator))
	if !ok || reflect.TypeOf(v) != value.Type() {
		return false
	}
	value.Set(reflect.ValueOf(v))
	return true
}

// hole adds a result to the if statement, which is a on one side and b on the
// other.
func (m *merger) hole(a, b dsl.Evaluator) (dsl.Evaluator, bool) {
	// Choosing between true and false is the condition itself.
	if x, ok := boolLiteral(a); ok {
		if y, ok := boolLiteral(b); ok && x != y {
			if x {
				return m.block.Cond, true
			}
			return dsl.NewBoolExpression(dsl.Op(nil, "!", m.block.Cond)), true
		}
	}
	rtype := reflect.TypeOf(a)
	if rtype == untyped || !isGPUValue(rtype) || !sameGLSLType(rtype, reflect.TypeOf(b)) {
		return nil, false
	}
	key := [2]int{canon.of(a), canon.of(b)}
	if v, ok := m.holes[key]; ok {
		return v, true
	}
	l := &dsl.Local{Block: m.block, Type: rtype, Slot: len(m.block.State)}
	m.block.State = append(m.block.State, l)
	m.block.Bodies[0].Next = append(m.block.Bodies[0].Next, a)
	m.block.Bodies[1].Next = append(m.block.Bodies[1].Next, b)
	ptr := reflect.New(rtype)
	bind(ptr, l)
	v := ptr.Elem().Interface().(dsl.Evaluator)
	m.holes[key] = v
	return v, true
}

func sameGLSLType(a, b reflect.Type) bool {
	return glslName(a) == glslName(b)
}

// rebuild returns a new expression of the same type as like, that computes e.
func rebuild(like dsl.Evaluator, e dsl.Evaluator) dsl.Evaluator {
	rtype := reflect.TypeOf(like)
	if rtype == untyped {
		return dsl.New(e)
	}
	ptr := reflect.New(rtype)
	dsl.Set(ptr.Interface().(dsl.Pointer), dsl.New(e))
	return ptr.Elem().Interface().(dsl.Evaluator)
}

// boolLiteral returns the value of a boolean written in Go.
func boolLiteral(e dsl.Evaluator) (bool, bool) {
	if !dsl.IsLiteral(e) {
		return false, false
	}
	rvalue := reflect.ValueOf(e)
	if !rvalue.Type().ConvertibleTo(reflect.TypeFor[dsl.Bool]()) {
		return false, false
	}
	return rvalue.Convert(reflect.TypeFor[dsl.Bool]()).Interface().(dsl.Bool).X, true
}
