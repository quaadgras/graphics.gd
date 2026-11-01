package shaders

import (
	"fmt"
	"reflect"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"graphics.gd/shaders/internal/gpu"
)

/*
Loops on the GPU are written as Go for loops over over.Int. The loop body is
ordinary Go code that assigns to variables declared before the loop, so the
values that are carried from one iteration to the next (the state of the loop)
are not declared anywhere. They are found by running the pipeline function more
than once, with the loop body running a different number of times:

  - zero times, to see the value of each variable before the loop.
  - once, along each way that the conditions evaluated by the body can go (each
    path), to see what one iteration does, and which paths break.
  - twice, for each pair of paths, to see which of the values made by the first
    iteration are read by the second. These are the state of the loop.

Every expression created by the body is a gpu.Cell, numbered by the order it
was created in, so the same expression can be recognised in each run. Every
other expression is numbered by the order it was created in by the pipeline
function, so that two variables that start with equal values are not confused.

A use is a place where a value is read: a child of an expression that was
created by the loop body, or a child of an expression that was created after the
loop. The uses that see the same value as each other in every run read the same
variable. Once the state is known, the body of each path is rebuilt to read the
variables of a GLSL loop instead, and the paths are merged with if statements,
like the branches of a pipeline function. Finally, the pipeline function is run
once more, with the cells that hold the values of the variables after the loop
replaced by reads of the GLSL loop's variables.
*/

func init() {
	gpu.Iterate = iterate
}

// maxLoopRuns limits how many times a pipeline function is run to analyse its loops.
const maxLoopRuns = 4096

// loopDriver runs a pipeline function until all of its loops are analysed.
type loopDriver struct {
	instances []*loopInstance
	byKey     map[loopKey]*loopInstance
	current   int // the loop being analysed, the loops before it are done.
	runs      int

	// for the current run:
	occur       map[string]int
	traces      map[*gpu.Trace]int
	active      *loopInstance // the loop that ran an experiment in this run.
	activeTrace *gpu.Trace
	iters       []*iterRecord
}

type loopKey struct {
	site  string
	count int // the number of times the site was reached before, in this run.
}

var (
	driver    *loopDriver
	iterating *iteration // the iteration being traced.
)

// runWithLoops runs fn until all of the loops it reaches are analysed, and
// returns what the final run returns. reset is called before each run.
func runWithLoops(fn func() []gpu.Evaluator, reset func()) []gpu.Evaluator {
	d := &loopDriver{byKey: make(map[loopKey]*loopInstance)}
	saved := driver
	driver = d
	defer func() { driver = saved }()
	for {
		if d.runs++; d.runs > maxLoopRuns {
			panic(fmt.Sprintf("shaders: more than %d runs needed to analyse the loops of a pipeline function", maxLoopRuns))
		}
		reset()
		d.occur = make(map[string]int)
		d.traces = make(map[*gpu.Trace]int)
		d.active, d.activeTrace, d.iters = nil, nil, nil
		gpu.ResetStamps()
		values := fn()
		if d.end(values) {
			return values
		}
	}
}

// iterate is called by the over package.
func iterate(count gpu.Evaluator, yield func(gpu.Int) bool) {
	d := driver
	if d == nil {
		panic("shaders: over.Int can only be used by a pipeline function, while the shader is being compiled")
	}
	if iterating != nil || gpu.Tracing() {
		panic("shaders: a loop cannot be nested inside of another loop")
	}
	site := callSite()
	key := loopKey{site, d.occur[site]}
	d.occur[site]++
	inst := d.byKey[key]
	if inst == nil {
		inst = newLoopInstance(len(d.instances), count)
		d.instances = append(d.instances, inst)
		d.byKey[key] = inst
	}
	trace := new(gpu.Trace)
	d.traces[trace] = inst.n
	switch {
	case inst.n < d.current:
		d.final(inst, trace, yield)
	case inst.n == d.current:
		d.active, d.activeTrace = inst, trace
		d.iters = d.iterations(trace, inst.queue[0].scripts, inst.index, yield)
	}
	// loops after the one being analysed do not run, until it is done.
}

// callSite returns the file and line of the for statement that ranges over the
// loop, which is the first caller outside of the shaders packages.
func callSite() string {
	pc := make([]uintptr, 16)
	frames := runtime.CallersFrames(pc[:runtime.Callers(2, pc)])
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, "graphics.gd/shaders.") && !strings.HasPrefix(frame.Function, "graphics.gd/shaders/over.") && !strings.HasPrefix(frame.Function, "graphics.gd/shaders/internal/") {
			return fmt.Sprintf("%s:%d", frame.File, frame.Line)
		}
		if !more {
			return "?"
		}
	}
}

// iteration is a run of a loop body that is being traced.
type iteration struct {
	rec     *iterRecord
	scratch *canonicalizer // numbers conditions, to evaluate a repeated condition the same way.
	decided map[int]bool
}

func (it *iteration) decide(cond gpu.Evaluator) bool {
	id := it.scratch.of(cond)
	if decided, ok := it.decided[id]; ok {
		return decided
	}
	k := len(it.rec.decisions)
	decided := true
	if k < len(it.rec.script) {
		decided = it.rec.script[k]
	}
	it.rec.decisions = append(it.rec.decisions, decided)
	it.rec.conds = append(it.rec.conds, cond)
	it.decided[id] = decided
	return decided
}

// iterRecord is what one iteration of a loop body did.
type iterRecord struct {
	trace     *gpu.Trace
	iteration int
	script    []bool // which way to go for the first conditions.
	decisions []bool
	conds     []gpu.Evaluator
	broke     bool
	cells     []*gpu.Cell
	leaves    map[string]leaf
}

// leaf is the value seen at a use.
type leaf struct {
	key   string
	value gpu.Evaluator
}

// iterations runs the body once for each script, until it breaks.
func (d *loopDriver) iterations(trace *gpu.Trace, scripts [][]bool, index gpu.Int, yield func(gpu.Int) bool) []*iterRecord {
	var recs []*iterRecord
	for _, script := range scripts {
		rec := &iterRecord{trace: trace, iteration: len(trace.Cells) + 1, script: script}
		recs = append(recs, rec)
		ok := func() bool {
			iterating = &iteration{rec: rec, scratch: newCanonicalizer(), decided: make(map[int]bool)}
			gpu.StartIteration(trace)
			defer func() {
				iterating = nil
				gpu.StopIteration()
			}()
			return yield(index)
		}()
		rec.cells = trace.Cells[rec.iteration-1]
		if len(rec.decisions) < len(script) || !slices.Equal(rec.decisions[:len(script)], script) {
			panic("shaders: a loop body must evaluate the same conditions each time it runs")
		}
		if !ok {
			rec.broke = true
			break
		}
	}
	return recs
}

// final runs the body of a loop that has been analysed, and replaces the values
// that the body assigned to variables with reads of the variables of the loop.
func (d *loopDriver) final(inst *loopInstance, trace *gpu.Trace, yield func(gpu.Int) bool) {
	var scripts [][]bool
	for _, p := range inst.finalPaths {
		scripts = append(scripts, p.decisions)
	}
	recs := d.iterations(trace, scripts, inst.index, yield)
	if len(recs) != len(scripts) {
		panic("shaders: a loop body must do the same thing each time it runs")
	}
	gpu.Within(inst.block.Created, func() {
		for _, r := range inst.rebinds {
			trace.Cells[r.iteration-1][r.seq].Target = gpu.NewUnstamped(r.class.state)
		}
	})
}

// end records the run that just finished, and reports whether it was the final run.
func (d *loopDriver) end(values []gpu.Evaluator) bool {
	inst := d.active
	if inst == nil {
		return d.current >= len(d.instances)
	}
	exp := inst.queue[0]
	inst.queue = inst.queue[1:]
	for _, rec := range d.iters {
		d.leaves(rec)
	}
	run := &loopRun{iters: d.iters, obs: d.observe(values)}
	inst.record(exp, run)
	if len(inst.queue) == 0 && inst.analyse() {
		d.current++
	}
	return false
}

// keyOf returns a key that is the same for the same value in each run.
func (d *loopDriver) keyOf(e gpu.Evaluator) string {
	if info, ok := gpu.InfoOf(e); ok {
		switch {
		case info.Cell != nil:
			c := info.Cell
			if c.Trace == d.activeTrace {
				return fmt.Sprintf("c%d:%d", c.Iteration, c.Seq)
			}
			return fmt.Sprintf("C%d:%d:%d", d.traces[c.Trace], c.Iteration, c.Seq)
		case info.Stamp != 0:
			return "o" + strconv.Itoa(info.Stamp)
		case info.Identity != nil:
			return fmt.Sprintf("p%p", info.Identity)
		}
	}
	// a literal, or a vector or matrix made out of components.
	var b strings.Builder
	b.WriteString("k" + glslName(reflect.TypeOf(e)) + "[")
	d.components(reflect.ValueOf(e), &b)
	return b.String() + "]"
}

func (d *loopDriver) components(value reflect.Value, b *strings.Builder) {
	eachLeaf(value, func(ev gpu.Evaluator) {
		b.WriteString(d.keyOf(ev))
		b.WriteByte(',')
	}, func(v reflect.Value) {
		fmt.Fprint(b, v.Interface())
		b.WriteByte(',')
	})
}

// hasComponentValues reports whether e is a vector or matrix literal made out
// of other values (as opposed to a scalar literal).
func hasComponentValues(e gpu.Evaluator) bool {
	found := false
	eachComponent(reflect.ValueOf(e), func(gpu.Evaluator) { found = true })
	return found
}

// isComposite reports whether e is a vector or matrix literal made out of
// other values, whose components are uses of their own.
func isComposite(e gpu.Evaluator) bool {
	return !isExpression(e) && hasComponentValues(e)
}

// isExpression reports whether e is an expression, rather than a literal.
func isExpression(e gpu.Evaluator) bool {
	info, ok := gpu.InfoOf(e)
	return ok && (info.Cell != nil || info.Identity != nil)
}

// leaves records the value seen by each use inside of an iteration: each child
// of the expressions it created, that it did not create itself, and each of
// the conditions that it evaluated.
func (d *loopDriver) leaves(rec *iterRecord) {
	rec.leaves = make(map[string]leaf)
	var visit func(pos string, e gpu.Evaluator)
	visit = func(pos string, e gpu.Evaluator) {
		if e == nil {
			return
		}
		if info, ok := gpu.InfoOf(e); ok && info.Cell != nil && info.Cell.Trace == rec.trace && info.Cell.Iteration == rec.iteration {
			return // made by this iteration.
		}
		rec.leaves[pos] = leaf{d.keyOf(e), e}
		if isComposite(e) {
			i := 0
			eachComponent(reflect.ValueOf(e), func(c gpu.Evaluator) {
				visit(pos+"."+strconv.Itoa(i), c)
				i++
			})
		}
	}
	for _, c := range rec.cells {
		i := 0
		eachChild(resolve(c.Target), func(child gpu.Evaluator) {
			visit(strconv.Itoa(c.Seq)+"/"+strconv.Itoa(i), child)
			i++
		})
	}
	for k, cond := range rec.conds {
		visit("cond"+strconv.Itoa(k), cond)
	}
}

// observe records the value seen by each use in the outputs of a run, and in
// the conditions evaluated by the pipeline function.
func (d *loopDriver) observe(values []gpu.Evaluator) map[string]leaf {
	obs := make(map[string]leaf)
	seen := make(map[*byte]bool)
	var visit func(pos string, e gpu.Evaluator)
	visit = func(pos string, e gpu.Evaluator) {
		if e == nil {
			return
		}
		key := d.keyOf(e)
		obs[pos] = leaf{key, e}
		info, ok := gpu.InfoOf(e)
		switch {
		case ok && info.Cell != nil:
			// a value made by a loop body.
		case ok && info.Identity != nil:
			if seen[info.Identity] {
				return
			}
			seen[info.Identity] = true
			i := 0
			eachChild(resolve(e), func(child gpu.Evaluator) {
				visit(key+"/"+strconv.Itoa(i), child)
				i++
			})
		default:
			i := 0
			eachComponent(reflect.ValueOf(e), func(c gpu.Evaluator) {
				visit(pos+"."+strconv.Itoa(i), c)
				i++
			})
		}
	}
	for i, v := range values {
		visit("out"+strconv.Itoa(i), v)
	}
	if len(explorers) > 0 {
		for i, cond := range explorers[len(explorers)-1].conds {
			visit("cond"+strconv.Itoa(i), cond)
		}
	}
	return obs
}

type experimentKind int

const (
	zeroIterations experimentKind = iota
	onePath                       // discover a path.
	twoPaths                      // see which values the second path reads from the first.
	threePaths                    // see what the second path does to a value that is a copy.
)

type experiment struct {
	kind    experimentKind
	scripts [][]bool
	paths   []*bodyPath
}

type loopRun struct {
	iters []*iterRecord
	obs   map[string]leaf
}

// bodyPath is one way that the conditions evaluated by a loop body can go.
type bodyPath struct {
	decisions []bool
	broke     bool
	run       *loopRun // the run where only this path ran.
}

type pathPair [2]*bodyPath
type pathTriple [3]*bodyPath

// loopInstance is a loop in a pipeline function.
type loopInstance struct {
	n     int
	block *gpu.Block
	body  *gpu.Scope
	index gpu.Int

	queue   []experiment
	paths   []*bodyPath
	known   map[string]*bodyPath
	zero    *loopRun
	pairs   map[pathPair]*loopRun
	triples map[pathTriple]*loopRun
	stage   experimentKind

	finalPaths []*bodyPath
	rebinds    []rebind
}

type rebind struct {
	iteration, seq int
	class          *loopClass
}

func newLoopInstance(n int, count gpu.Evaluator) *loopInstance {
	created := gpu.CurrentScope()
	body := &gpu.Scope{Parent: created, Loop: true}
	block := &gpu.Block{Kind: gpu.BlockFor, Created: created, Count: count, Bodies: []gpu.Body{{Scope: body}}}
	block.Index = &gpu.Local{Block: block, Type: reflect.TypeFor[gpu.Int](), Param: true}
	return &loopInstance{
		n:       n,
		block:   block,
		body:    body,
		index:   gpu.NewIntExpression(gpu.NewUnstamped(block.Index)),
		queue:   []experiment{{kind: zeroIterations}, {kind: onePath, scripts: [][]bool{nil}}},
		known:   make(map[string]*bodyPath),
		pairs:   make(map[pathPair]*loopRun),
		triples: make(map[pathTriple]*loopRun),
	}
}

func pathKey(decisions []bool, broke bool) string {
	var b strings.Builder
	for _, d := range decisions {
		if d {
			b.WriteByte('T')
		} else {
			b.WriteByte('F')
		}
	}
	if broke {
		b.WriteByte('!')
	}
	return b.String()
}

func (inst *loopInstance) record(exp experiment, run *loopRun) {
	switch exp.kind {
	case zeroIterations:
		inst.zero = run
	case onePath:
		rec := run.iters[0]
		key := pathKey(rec.decisions, rec.broke)
		if inst.known[key] == nil {
			p := &bodyPath{decisions: rec.decisions, broke: rec.broke, run: run}
			inst.known[key] = p
			inst.paths = append(inst.paths, p)
		}
		for i := len(exp.scripts[0]); i < len(rec.decisions); i++ {
			script := append(slices.Clone(rec.decisions[:i]), !rec.decisions[i])
			inst.queue = append(inst.queue, experiment{kind: onePath, scripts: [][]bool{script}})
		}
	case twoPaths:
		inst.pairs[pathPair(exp.paths)] = run
	case threePaths:
		inst.triples[pathTriple(exp.paths)] = run
	}
}

func (inst *loopInstance) continuing() []*bodyPath {
	var paths []*bodyPath
	for _, p := range inst.paths {
		if !p.broke {
			paths = append(paths, p)
		}
	}
	return paths
}

// analyse is called when the queue of experiments is empty, it either queues
// more experiments, or finishes the analysis and reports true.
func (inst *loopInstance) analyse() bool {
	if inst.stage == onePath || inst.stage == zeroIterations {
		inst.stage = twoPaths
		for _, x := range inst.continuing() {
			for _, p := range inst.paths {
				inst.queue = append(inst.queue, experiment{kind: twoPaths, scripts: [][]bool{x.decisions, p.decisions}, paths: []*bodyPath{x, p}})
			}
		}
		if len(inst.queue) > 0 {
			return false
		}
	}
	a := inst.classify()
	if needed := a.transfers(); len(needed) > 0 {
		inst.stage = threePaths
		for _, t := range needed {
			inst.queue = append(inst.queue, experiment{kind: threePaths, scripts: [][]bool{t[0].decisions, t[1].decisions, t[2].decisions}, paths: t[:]})
		}
		return false
	}
	a.build()
	a.plan()
	return true
}

// loopClass is a variable that a loop changes.
type loopClass struct {
	start    leaf                 // its value before the loop.
	after    map[*bodyPath]string // its value after one iteration along each path.
	observed []string             // uses after the loop.
	rebound  bool                 // used after the loop, through one of its components.
	uses     map[*bodyPath][]string
	typ      reflect.Type

	transfer map[*bodyPath]transfer // what each path does to it.

	state, param *gpu.Local
}

type transferKind int

const (
	keep     transferKind = iota // unchanged.
	compute                      // set to a value computed by the body.
	copyOf                       // set to the value of another variable at the start of the iteration.
	constant                     // set to a value computed before the loop.
)

type transfer struct {
	kind  transferKind
	seq   int        // compute: the cell that holds the value.
	other *loopClass // copyOf
	value gpu.Evaluator
}

type loopAnalysis struct {
	inst      *loopInstance
	classes   []*loopClass
	leafClass map[*bodyPath]map[string]*loopClass
}

func (inst *loopInstance) classify() *loopAnalysis {
	a := &loopAnalysis{inst: inst, leafClass: make(map[*bodyPath]map[string]*loopClass)}
	continuing := inst.continuing()
	var breaking []*bodyPath
	for _, p := range inst.paths {
		if p.broke {
			breaking = append(breaking, p)
		}
	}
	keysOf := func(paths []*bodyPath, key func(*bodyPath) string) string {
		var b strings.Builder
		for _, p := range paths {
			b.WriteString(key(p))
			b.WriteByte('|')
		}
		return b.String()
	}
	// Uses after the loop, which are wherever a value made by the loop is seen.
	positions := make(map[string]bool)
	runs := []*loopRun{}
	for _, p := range inst.paths {
		runs = append(runs, p.run)
	}
	for _, run := range inst.pairs {
		runs = append(runs, run)
	}
	for _, run := range runs {
		for pos, l := range run.obs {
			if strings.HasPrefix(l.key, "c") {
				positions[pos] = true
			}
		}
	}
	byFull := make(map[string]*loopClass)
	byPrefix := make(map[string]*loopClass)
	for _, pos := range slices.Sorted(mapsKeys(positions)) {
		start, ok := inst.zero.obs[pos]
		if !ok {
			panic("shaders: the code after a loop must do the same thing however many times the loop runs")
		}
		after := make(map[*bodyPath]string)
		changes := false
		for _, p := range inst.paths {
			l, ok := p.run.obs[pos]
			if !ok {
				panic("shaders: the code after a loop must do the same thing however many times the loop runs")
			}
			after[p] = l.key
			changes = changes || l.key != start.key
		}
		if !changes {
			continue
		}
		prefix := start.key + "#" + keysOf(continuing, func(p *bodyPath) string { return after[p] })
		full := prefix + "#" + keysOf(breaking, func(p *bodyPath) string { return after[p] })
		c := byFull[full]
		if c == nil {
			c = &loopClass{start: start, after: after, uses: make(map[*bodyPath][]string)}
			byFull[full] = c
			a.classes = append(a.classes, c)
			if byPrefix[prefix] == nil {
				byPrefix[prefix] = c
			}
		}
		c.observed = append(c.observed, pos)
	}
	// Uses inside of the loop body.
	for _, p := range inst.paths {
		a.leafClass[p] = make(map[string]*loopClass)
		rec := p.run.iters[0]
		for _, pos := range slices.Sorted(mapsKeys(rec.leaves)) {
			// A vector literal is a use of its own when one of the runs sees
			// a value made by the loop in its place, otherwise each of its
			// components is.
			start := rec.leaves[pos]
			after := make(map[*bodyPath]string)
			changes, missing, composite := false, false, isComposite(start.value)
			for _, x := range continuing {
				l, ok := inst.pairs[pathPair{x, p}].iters[1].leaves[pos]
				if !ok {
					missing = true // part of a value that is a use of its own.
					break
				}
				after[x] = l.key
				changes = changes || l.key != start.key
				composite = composite && isComposite(l.value)
			}
			if missing || composite || !changes {
				continue
			}
			prefix := start.key + "#" + keysOf(continuing, func(x *bodyPath) string { return after[x] })
			c := byPrefix[prefix]
			if c == nil {
				c = &loopClass{start: start, after: after, uses: make(map[*bodyPath][]string)}
				byPrefix[prefix] = c
				a.classes = append(a.classes, c)
			}
			c.uses[p] = append(c.uses[p], pos)
			a.leafClass[p][pos] = c
		}
	}
	a.components()
	for _, c := range a.classes {
		c.typ = reflect.TypeOf(c.start.value)
		if c.typ == untyped || !isGPUValue(c.typ) {
			panic(fmt.Sprintf("shaders: cannot carry a value of type %v from one iteration of a loop to the next", c.typ))
		}
	}
	return a
}

func mapsKeys[V any](m map[string]V) func(func(string) bool) {
	return func(yield func(string) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}
}

// components replaces each variable used after the loop that is a component of
// another variable (as the X of a vector) with that other variable, so that it
// is not carried by the loop separately.
func (a *loopAnalysis) components() {
	keep := a.classes[:0]
	for _, c := range a.classes {
		if len(c.observed) == 0 {
			keep = append(keep, c)
			continue
		}
		var parent *loopClass
		parentAfter := make(map[*bodyPath]string)
		ok := true
		for _, p := range a.inst.paths {
			after := c.after[p]
			rest, isCell := strings.CutPrefix(after, "c1:")
			if !isCell {
				ok = ok && after == c.start.key
				continue
			}
			seq, _ := strconv.Atoi(rest)
			sel, isSelect := resolve(p.run.iters[0].cells[seq].Target).(gpu.Select)
			info, _ := gpu.InfoOf(sel.Value)
			if !isSelect || info.Cell == nil || info.Cell.Iteration != 1 || info.Cell.Trace != p.run.iters[0].trace {
				ok = false
				break
			}
			key := fmt.Sprintf("c1:%d", info.Cell.Seq)
			parentAfter[p] = key
			if !p.broke {
				for _, u := range a.classes {
					if u != c && u.after[p] == key && (parent == nil || parent == u) {
						parent = u
					}
				}
				ok = ok && parent != nil && parent.after[p] == key
			}
		}
		if !ok || parent == nil {
			keep = append(keep, c)
			continue
		}
		parent.rebound = true
		for p, key := range parentAfter {
			if p.broke {
				parent.after[p] = key
			}
		}
	}
	a.classes = keep
}

// valueAfter returns the value of c after one iteration along p, as seen by one
// of its uses.
func (a *loopAnalysis) valueAfter(c *loopClass, p *bodyPath) (leaf, bool) {
	if len(c.observed) > 0 {
		l, ok := p.run.obs[c.observed[0]]
		return l, ok
	}
	for y, positions := range c.uses {
		if run := a.inst.pairs[pathPair{p, y}]; run != nil && len(run.iters) > 1 {
			l, ok := run.iters[1].leaves[positions[0]]
			return l, ok
		}
	}
	return leaf{}, false
}

// transfers works out what each path does to each variable. It returns the
// runs needed to tell apart the variables that a path sets to a copy of a value
// from before the iteration.
func (a *loopAnalysis) transfers() (needed []pathTriple) {
	continuing := a.inst.continuing()
	for _, c := range a.classes {
		c.transfer = make(map[*bodyPath]transfer)
		for _, p := range a.inst.paths {
			if p.broke && len(c.observed) == 0 && !c.rebound {
				c.transfer[p] = transfer{kind: keep} // never seen again.
				continue
			}
			var after string
			if p.broke {
				after = c.after[p]
			} else {
				after = c.after[p]
			}
			if rest, ok := strings.CutPrefix(after, "c1:"); ok {
				seq, _ := strconv.Atoi(rest)
				c.transfer[p] = transfer{kind: compute, seq: seq}
				continue
			}
			value, ok := a.valueAfter(c, p)
			if !ok {
				panic("shaders: a loop body must do the same thing each time it runs")
			}
			var candidates []transfer
			if after == c.start.key {
				candidates = append(candidates, transfer{kind: keep})
			}
			for _, u := range a.classes {
				if u != c && u.start.key == after {
					candidates = append(candidates, transfer{kind: copyOf, other: u})
				}
			}
			candidates = append(candidates, transfer{kind: constant, value: value.value})
			if len(candidates) == 1 {
				c.transfer[p] = candidates[0]
				continue
			}
			// Tell the candidates apart, by what the path does after another path.
			var chosen *transfer
		next:
			for i := range candidates {
				for _, x := range continuing {
					seen, ok := a.valueAfterTwo(c, x, p)
					if !ok {
						for z := range c.uses {
							needed = append(needed, pathTriple{x, p, z})
							break
						}
						continue next
					}
					var predicted string
					switch t := candidates[i]; t.kind {
					case keep:
						predicted = c.after[x]
					case copyOf:
						predicted = t.other.after[x]
					case constant:
						predicted = after
					}
					if predicted != seen {
						continue next
					}
				}
				chosen = &candidates[i]
				break
			}
			if chosen == nil {
				if len(needed) > 0 {
					continue
				}
				panic("shaders: cannot work out what a loop body does to one of the variables it changes")
			}
			c.transfer[p] = *chosen
		}
	}
	slices.SortFunc(needed, func(a, b pathTriple) int { return 0 })
	needed = slices.CompactFunc(needed, func(a, b pathTriple) bool { return a == b })
	var missing []pathTriple
	for _, t := range needed {
		if a.inst.triples[t] == nil {
			missing = append(missing, t)
		}
	}
	if len(needed) > 0 && len(missing) == 0 {
		panic("shaders: cannot work out what a loop body does to one of the variables it changes")
	}
	return missing
}

// valueAfterTwo returns the key of the value of c after an iteration along x,
// then one along p.
func (a *loopAnalysis) valueAfterTwo(c *loopClass, x, p *bodyPath) (string, bool) {
	if len(c.observed) > 0 {
		if run := a.inst.pairs[pathPair{x, p}]; run != nil {
			l, ok := run.obs[c.observed[0]]
			return l.key, ok
		}
		return "", false
	}
	for z, positions := range c.uses {
		if run := a.inst.triples[pathTriple{x, p, z}]; run != nil && len(run.iters) > 2 {
			l, ok := run.iters[2].leaves[positions[0]]
			return l.key, ok
		}
	}
	return "", false
}

// build rebuilds the body of each path to read the variables of the loop, and
// merges the paths into the body of the loop.
func (a *loopAnalysis) build() {
	inst := a.inst
	block := inst.block
	gpu.Within(inst.body, func() {
		params := make(map[*loopClass]gpu.Evaluator)
		for i, c := range a.classes {
			c.state = &gpu.Local{Block: block, Type: c.typ, Slot: i}
			c.param = &gpu.Local{Block: block, Type: c.typ, Slot: i, Param: true}
			block.State = append(block.State, c.state)
			block.Params = append(block.Params, c.param)
			block.Init = append(block.Init, c.start.value)
			ptr := reflect.New(c.typ)
			bind(ptr, c.param)
			params[c] = ptr.Elem().Interface().(gpu.Evaluator)
		}
		breaks := false
		values := make(map[*bodyPath][]gpu.Evaluator)
		conds := make(map[*bodyPath][]gpu.Evaluator)
		for _, p := range inst.paths {
			r := &rebuilder{rec: p.run.iters[0], classes: a.leafClass[p], params: params, built: make(map[int]gpu.Evaluator)}
			var next []gpu.Evaluator
			for _, c := range a.classes {
				switch t := c.transfer[p]; t.kind {
				case keep:
					next = append(next, params[c])
				case copyOf:
					next = append(next, params[t.other])
				case constant:
					next = append(next, t.value)
				case compute:
					cell := r.rec.cells[t.seq]
					next = append(next, r.cell(cell, c.typ))
				}
			}
			next = append(next, gpu.NewBool(p.broke))
			breaks = breaks || p.broke
			values[p] = next
			for k, cond := range r.rec.conds {
				conds[p] = append(conds[p], r.value("cond"+strconv.Itoa(k), cond))
			}
		}
		merged := mergePaths(inst.paths, 0, values, conds)
		block.Bodies[0].Next = merged[:len(merged)-1]
		if breaks {
			block.Bodies[0].Break = merged[len(merged)-1]
		}
	})
}

func mergePaths(paths []*bodyPath, depth int, values, conds map[*bodyPath][]gpu.Evaluator) []gpu.Evaluator {
	if len(paths) == 1 {
		return values[paths[0]]
	}
	var then, otherwise []*bodyPath
	for _, p := range paths {
		if p.decisions[depth] {
			then = append(then, p)
		} else {
			otherwise = append(otherwise, p)
		}
	}
	if len(then) == 0 {
		return mergePaths(otherwise, depth+1, values, conds)
	}
	if len(otherwise) == 0 {
		return mergePaths(then, depth+1, values, conds)
	}
	return merge(conds[then[0]][depth], mergePaths(then, depth+1, values, conds), mergePaths(otherwise, depth+1, values, conds), func(int) string { return "a variable of the loop" })
}

// rebuilder rebuilds the expressions of one iteration, with each use of a
// variable of the loop replaced by a read of its parameter.
type rebuilder struct {
	rec     *iterRecord
	classes map[string]*loopClass
	params  map[*loopClass]gpu.Evaluator
	built   map[int]gpu.Evaluator
}

func (r *rebuilder) cell(c *gpu.Cell, like reflect.Type) gpu.Evaluator {
	node, ok := r.built[c.Seq]
	if !ok {
		i := 0
		node = mapChildren(resolve(c.Target), func(child gpu.Evaluator) gpu.Evaluator {
			v := r.value(strconv.Itoa(c.Seq)+"/"+strconv.Itoa(i), child)
			i++
			return v
		})
		r.built[c.Seq] = node
	}
	if like == nil || like == untyped {
		return gpu.New(node)
	}
	ptr := reflect.New(like)
	gpu.Set(ptr.Interface().(gpu.Pointer), gpu.New(node))
	return ptr.Elem().Interface().(gpu.Evaluator)
}

func (r *rebuilder) value(pos string, v gpu.Evaluator) gpu.Evaluator {
	if v == nil {
		return nil
	}
	if info, ok := gpu.InfoOf(v); ok && info.Cell != nil && info.Cell.Trace == r.rec.trace && info.Cell.Iteration == r.rec.iteration {
		return r.cell(info.Cell, reflect.TypeOf(v))
	}
	if c := r.classes[pos]; c != nil {
		return r.params[c]
	}
	if isComposite(v) {
		value := reflect.New(reflect.TypeOf(v)).Elem()
		value.Set(reflect.ValueOf(v))
		i := 0
		setComponents(value, func(c gpu.Evaluator) gpu.Evaluator {
			out := r.value(pos+"."+strconv.Itoa(i), c)
			i++
			return out
		})
		return value.Interface().(gpu.Evaluator)
	}
	if c := r.classes[pos]; c != nil {
		return r.params[c]
	}
	return v
}

// setComponents replaces each component of a vector or matrix literal, in the
// same order as eachComponent visits them.
func setComponents(value reflect.Value, fn func(gpu.Evaluator) gpu.Evaluator) {
	switch value.Kind() {
	case reflect.Struct:
		for i := range value.NumField() {
			if !value.Type().Field(i).IsExported() {
				continue
			}
			field := value.Field(i)
			if ev, ok := field.Interface().(gpu.Evaluator); ok {
				out := reflect.ValueOf(fn(ev))
				if out.Type().ConvertibleTo(field.Type()) {
					field.Set(out.Convert(field.Type()))
				}
			} else {
				setComponents(field, fn)
			}
		}
	case reflect.Array:
		for i := range value.Len() {
			elem := value.Index(i)
			if ev, ok := elem.Interface().(gpu.Evaluator); ok {
				out := reflect.ValueOf(fn(ev))
				if out.Type().ConvertibleTo(elem.Type()) {
					elem.Set(out.Convert(elem.Type()))
				}
			} else {
				setComponents(elem, fn)
			}
		}
	}
}

// mapChildren returns resolved with each of its children replaced, visiting
// them in the same order as eachChild.
func mapChildren(resolved gpu.Evaluator, fn func(gpu.Evaluator) gpu.Evaluator) gpu.Evaluator {
	switch v := resolved.(type) {
	case gpu.Operation:
		if v.A != nil {
			v.A = fn(v.A)
		}
		v.B = fn(v.B)
		return v
	case gpu.Select:
		v.Value = fn(v.Value)
		return v
	case gpu.Index:
		v.Value = fn(v.Value)
		return v
	case gpu.FunctionCall:
		args := make([]gpu.Evaluator, len(v.Args))
		for i, arg := range v.Args {
			args[i] = fn(arg)
		}
		v.Args = args
		return v
	case gpu.Ternary:
		v.If = fn(v.If)
		v.A = fn(v.A)
		v.B = fn(v.B)
		return v
	default:
		return resolved
	}
}

// plan chooses which paths to run, in the final run, so that each variable
// used after the loop holds a value made by the loop body that no other
// variable holds, which can then be replaced by a read of the loop's variable.
func (a *loopAnalysis) plan() {
	var observed []*loopClass
	for _, c := range a.classes {
		if len(c.observed) > 0 || c.rebound {
			observed = append(observed, c)
		}
	}
	if len(observed) == 0 {
		return
	}
	continuing := a.inst.continuing()
	var candidates [][]*bodyPath
	var extend func(seq []*bodyPath, depth int)
	extend = func(seq []*bodyPath, depth int) {
		if len(seq) > 0 {
			candidates = append(candidates, slices.Clone(seq))
		}
		if depth == 0 || len(candidates) > 100000 {
			return
		}
		for _, p := range a.inst.paths {
			if len(seq) > 0 && seq[len(seq)-1].broke {
				return
			}
			extend(append(seq, p), depth-1)
		}
	}
	extend(nil, min(len(observed)+1, 4))
	_ = continuing
	slices.SortStableFunc(candidates, func(x, y []*bodyPath) int { return len(x) - len(y) })
	for _, seq := range candidates {
		if rebinds, ok := a.simulate(seq, observed); ok {
			a.inst.finalPaths = seq
			a.inst.rebinds = rebinds
			return
		}
	}
	panic("shaders: cannot work out the value of every variable after a loop; a variable that is only set right before a break needs to be set before every break")
}

// simulate follows what the paths in seq do to each variable, and returns
// which value made by the loop body each variable in observed holds after
// them, unless two of them hold the same value, or one holds a value made
// before the loop.
func (a *loopAnalysis) simulate(seq []*bodyPath, observed []*loopClass) ([]rebind, bool) {
	holds := make(map[*loopClass]rebind)
	for j, p := range seq {
		if p.broke && j != len(seq)-1 {
			return nil, false
		}
		next := make(map[*loopClass]rebind)
		for _, c := range a.classes {
			switch t := c.transfer[p]; t.kind {
			case compute:
				next[c] = rebind{iteration: j + 1, seq: t.seq}
			case keep:
				if r, ok := holds[c]; ok {
					next[c] = r
				}
			case copyOf:
				if r, ok := holds[t.other]; ok {
					next[c] = r
				}
			}
		}
		holds = next
	}
	seen := make(map[[2]int]bool)
	var rebinds []rebind
	for _, c := range observed {
		r, ok := holds[c]
		if !ok || seen[[2]int{r.iteration, r.seq}] {
			return nil, false
		}
		seen[[2]int{r.iteration, r.seq}] = true
		r.class = c
		rebinds = append(rebinds, r)
	}
	return rebinds, true
}
