//go:build reloads && !wasip1

package startup

// The reloads host: built with -tags reloads, this process owns the
// engine natively but registers no classes of its own — the project is
// also compiled to GOOS=wasip1 wasm and instantiated as a "guest" module
// (see startup_reloads_wasip1.go) that performs every registration
// through the generated "gd" bridge (reloads_wazero.go). When the
// project's source changes, the host rebuilds the guest, asks the
// running one to exit its yield loop, unregisters the classes the old
// module registered and instantiates the new build — while the engine
// (and the editor) keeps running, so new classes become available
// without a restart.
//
// Threading model: everything guest-related happens on the engine main
// thread. The guest's _start runs inside InstantiateModule here; each
// engine iteration runs inside the guest's yield host-call; engine
// callbacks forward back into the guest only as nested calls. The
// watcher goroutine touches nothing but atomics and the build cache.

import (
	"context"
	"fmt"
	"hash/fnv"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
	"graphics.gd/classdb/Startup"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

var reloadsExitCode atomic.Uint32 // reloadsKeepRunning/reloadsShutdown/reloadsSwap

const (
	reloadsKeepRunning = 0
	reloadsShutdown    = 1
	reloadsSwap        = 2
)

var reloadsCompiled atomic.Pointer[wazero.CompiledModule]

var (
	reloadsClassesMu sync.Mutex
	reloadsClasses   []string
)

var reloadsInitLevels []gdextension.InitializationLevel

var reloadsNeedsStart bool

var reloadsRuntime wazero.Runtime

var reloadsAdoptFn api.Function

// reloadsCShared is true when the extension was loaded into a godot
// process (c-shared): the engine drives the main loop and the guest
// session runs inside the main coroutine (see call_main_in_steps in
// startup_cgo.go), parking via pause_main between events instead of
// driving engine iterations from the yield import.
var reloadsCShared bool

var reloadsParked atomic.Bool

var reloadsAwaitingBuild atomic.Bool

// reloadsPark yields control from the main coroutine back to the
// engine; the reloads frame chain resumes it when there is something
// for the guest session to do (swap, shutdown, or a fresh build).
func reloadsPark() {
	if pause_main == nil {
		panic("graphics.gd: the reloads host requires the main coroutine (gd test builds must not use -tags reloads)")
	}
	reloadsParked.Store(true)
	pause_main(false)
	reloadsParked.Store(false)
}

// reloadsWakeWanted reports whether a parked guest session has
// something to do.
func reloadsWakeWanted() bool {
	if reloadsExitCode.Load() != reloadsKeepRunning {
		return true
	}
	return reloadsAwaitingBuild.Load() && reloadsCompiled.Load() != nil
}

func init() {
	reloadsSession = reloadsRun
}

// reloadsBase gives the reloads host access to the engine handle shared
// by the library startup variants (shared and static): any startup value
// embedding engineAsLibrary picks this method up. The GDExtension
// (.so-inside-godot) startup mode does not, and is not supported as a
// reloads host.
func (e *engineAsLibrary) reloadsBase() *engineAsLibrary { return e }

func reloadsEngine() *engineAsLibrary {
	base, ok := startup.(interface{ reloadsBase() *engineAsLibrary })
	if !ok {
		panic("graphics.gd: the reloads host requires the engine-as-library startup mode (run the built binary directly, not as a GDExtension)")
	}
	return base.reloadsBase()
}

// reloadsProjectDir locates the Go module being reloaded: the host is
// conventionally started from the project's graphics/ directory (the
// Godot project), with the Go module in its parent.
func reloadsProjectDir() string {
	wd, err := os.Getwd()
	if err != nil {
		return "."
	}
	if filepath.Base(wd) == "graphics" {
		return filepath.Dir(wd)
	}
	return wd
}

func reloadsBuildGuest() error {
	dir := reloadsProjectDir()
	out := filepath.Join(dir, "graphics", "library.wasm")
	cmd := exec.Command("go", "build", "-o", out, ".")
	cmd.Env = append(os.Environ(), "GOOS=wasip1", "GOARCH=wasm", "CGO_ENABLED=0")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	data, err := os.ReadFile(out)
	if err != nil {
		return err
	}
	compiled, err := reloadsRuntime.CompileModule(reloadsCtx, data)
	if err != nil {
		return err
	}
	reloadsCompiled.Store(&compiled)
	return nil
}

// reloadsSourceStamp hashes the name, size and mtime of every .go file
// (plus go.mod/go.sum) under the project directory.
func reloadsSourceStamp(dir string) uint64 {
	sum := fnv.New64a()
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			name := d.Name()
			if name == "graphics" || strings.HasPrefix(name, ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") && d.Name() != "go.mod" && d.Name() != "go.sum" {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return nil
		}
		fmt.Fprintf(sum, "%s %d %d\n", path, info.Size(), info.ModTime().UnixNano())
		return nil
	})
	return sum.Sum64()
}

func reloadsWatch() {
	dir := reloadsProjectDir()
	last := reloadsSourceStamp(dir)
	for {
		time.Sleep(500 * time.Millisecond)
		if reloadsExitCode.Load() == reloadsShutdown {
			return
		}
		stamp := reloadsSourceStamp(dir)
		if stamp == last {
			continue
		}
		last = stamp
		fmt.Fprintln(os.Stderr, "graphics.gd: source change detected, rebuilding...")
		if err := reloadsBuildGuest(); err != nil {
			fmt.Fprintln(os.Stderr, "graphics.gd: reload build failed, keeping the running build:", err)
			continue
		}
		reloadsExitCode.CompareAndSwap(reloadsKeepRunning, reloadsSwap)
	}
}

// State that survives module swaps. Godot's unregister_extension_class
// frees method binds with no live-instance protection, so classes are
// never unregistered on a swap: the engine keeps the registration from
// the session that first introduced each class, duplicate registrations
// from later sessions are skipped, and their fresh ExtensionClassIDs
// are remapped so engine callbacks land in the current module. Anything
// referring to per-module state that died with a swapped-out module
// (extension instances, wrapper bindings, callables) is answered
// inertly instead of being forwarded (see reloadsGuardStale).
type reloadsInstanceInfo struct {
	obj   uint64 // engine object the instance is bound to
	class string // engine-registered class name
	epoch uint32 // module epoch that owns the binding
}

type reloadsMethodKey struct {
	class  string
	method string
	hash   uint32
}

type reloadsTokenInfo struct {
	key   reloadsMethodKey
	kind  uint8 // reloadsTokenRegistered / reloadsTokenVirtual / reloadsTokenCaller
	epoch uint32
}

const (
	reloadsTokenRegistered = iota
	reloadsTokenVirtual
	reloadsTokenCaller
)

var reloadsEpoch atomic.Uint32 // current module epoch, bumped on swap

var (
	reloadsEngineClasses    = map[string]gdextension.ExtensionClassID{} // name -> id the engine was registered with
	reloadsEngineClassNames = map[gdextension.ExtensionClassID]string{} // reverse of the above
	reloadsClassRemap       = map[gdextension.ExtensionClassID]gdextension.ExtensionClassID{}
	reloadsDeduped          = map[string]bool{}                // classes whose re-registration was skipped this session
	reloadsInstances        = map[uint64]reloadsInstanceInfo{} // extension instances by id
	reloadsBindings         = map[uint64]bool{}                // live wrapper bindings of the current module
	reloadsCallables        = map[uint64]bool{}                // live callable functions of the current module
	reloadsTokens           = map[uint64]reloadsTokenInfo{}    // method fn tokens / virtual callData by value
	reloadsCurrentTokens    = map[reloadsMethodKey]uint64{}    // latest token for a method (current epoch)
	reloadsListPushes       = map[uint64][]reloadsListPush{}   // method-list handle -> pushed methods
)

type reloadsListPush struct {
	name  string
	token uint64
}

// reloadsClassToken resolves the current module's class token for an
// engine-registered class name.
func reloadsClassToken(name string) (gdextension.ExtensionClassID, bool) {
	engineID, ok := reloadsEngineClasses[name]
	if !ok {
		return 0, false
	}
	if mapped, ok := reloadsClassRemap[engineID]; ok {
		return mapped, true
	}
	return engineID, true
}

// reloadsTrackRegistrations wraps the native registration entry points:
// class registrations are recorded (and deduplicated across sessions),
// extension instance setups and callable creations are recorded so
// stale callbacks can be recognized after a swap.
func reloadsTrackRegistrations() {
	register := gdextension.Host.ClassDB.Register.Class
	gdextension.Host.ClassDB.Register.Class = func(class, parent_class gdextension.StringName, id gdextension.ExtensionClassID, virtual, abstract, exposed, runtime bool, icon_path gdextension.String) {
		name := reloadsStringNameToString(class)
		reloadsClassesMu.Lock()
		engineID, known := reloadsEngineClasses[name]
		if known {
			// The engine already has this class from an earlier session;
			// keep that registration and route its callbacks to the
			// current module's class token. Its methods, properties and
			// signals must not re-register either (see below).
			reloadsClassRemap[engineID] = id
			reloadsDeduped[name] = true
			reloadsClassesMu.Unlock()
			return
		}
		reloadsEngineClasses[name] = id
		reloadsEngineClassNames[id] = name
		reloadsClasses = append(reloadsClasses, name)
		reloadsClassesMu.Unlock()
		register(class, parent_class, id, virtual, abstract, exposed, runtime, icon_path)
	}
	removal := gdextension.Host.ClassDB.Register.Removal
	gdextension.Host.ClassDB.Register.Removal = func(class gdextension.StringName) {
		if reloadsExitCode.Load() == reloadsSwap {
			// A module on its way out must not unregister classes the
			// engine may still have live instances of (Godot frees the
			// method binds unconditionally).
			return
		}
		name := reloadsStringNameToString(class)
		reloadsClassesMu.Lock()
		if i := slices.Index(reloadsClasses, name); i >= 0 {
			reloadsClasses = slices.Delete(reloadsClasses, i, i+1)
		}
		delete(reloadsEngineClasses, name)
		reloadsClassesMu.Unlock()
		removal(class)
	}
	// Drop method/constant/property/signal registrations for classes
	// whose (identical-name) registration was deduplicated this session:
	// the engine already has them from the original registration.
	registers := reflect.ValueOf(&gdextension.Host.ClassDB.Register).Elem()
	for i := range registers.NumField() {
		switch registers.Type().Field(i).Name {
		case "Class", "Removal":
			continue
		}
		f := registers.Field(i)
		if f.Kind() != reflect.Func || f.IsNil() || f.Type().NumIn() == 0 || f.Type().In(0) != reflect.TypeFor[gdextension.StringName]() {
			continue
		}
		prev := f.Interface()
		ftype := f.Type()
		f.Set(reflect.MakeFunc(ftype, func(args []reflect.Value) []reflect.Value {
			class := args[0].Interface().(gdextension.StringName)
			reloadsClassesMu.Lock()
			skip := reloadsDeduped[reloadsStringNameToString(class)]
			reloadsClassesMu.Unlock()
			if skip {
				out := make([]reflect.Value, ftype.NumOut())
				for k := range out {
					out[k] = reflect.Zero(ftype.Out(k))
				}
				return out
			}
			return reflect.ValueOf(prev).Call(args)
		}))
	}
	// Record every method registered through a method list so that
	// engine method binds cached from earlier sessions can be remapped
	// to the current module's tokens (see reloadsMapToken).
	push := gdextension.Host.ClassDB.MethodList.Push
	gdextension.Host.ClassDB.MethodList.Push = func(info gdextension.MethodList, name gdextension.StringName, call gdextension.FunctionID, method_flags gdextension.MethodFlags, return_value_info gdextension.PropertyList, arguments_info gdextension.PropertyList, count int, default_arguments gdextension.CallAccepts[gdextension.Variant]) {
		reloadsListPushes[uint64(info)] = append(reloadsListPushes[uint64(info)], reloadsListPush{
			name:  reloadsStringNameToString(name),
			token: uint64(call),
		})
		push(info, name, call, method_flags, return_value_info, arguments_info, count, default_arguments)
	}
	registerMethods := gdextension.Host.ClassDB.Register.Methods
	gdextension.Host.ClassDB.Register.Methods = func(class gdextension.StringName, list gdextension.MethodList) {
		className := reloadsStringNameToString(class)
		epoch := reloadsEpoch.Load()
		for _, m := range reloadsListPushes[uint64(list)] {
			key := reloadsMethodKey{class: className, method: m.name}
			reloadsCurrentTokens[key] = m.token
			reloadsTokens[m.token] = reloadsTokenInfo{key: key, kind: reloadsTokenRegistered, epoch: epoch}
		}
		delete(reloadsListPushes, uint64(list))
		registerMethods(class, list)
	}
	setup := gdextension.Host.Objects.Extension.Setup
	gdextension.Host.Objects.Extension.Setup = func(obj gdextension.Object, name gdextension.StringName, id gdextension.ExtensionInstanceID) {
		reloadsInstances[uint64(id)] = reloadsInstanceInfo{
			obj:   uint64(obj),
			class: reloadsStringNameToString(name),
			epoch: reloadsEpoch.Load(),
		}
		setup(obj, name, id)
	}
	callable := gdextension.Host.Callables.Create
	gdextension.Host.Callables.Create = func(id gdextension.FunctionID, object gdextension.ObjectID, result gdextension.CallReturns[gdextension.Callable]) {
		reloadsCallables[uint64(id)] = true
		callable(id, object, result)
	}
}

func reloadsStringNameToString(name gdextension.StringName) string {
	return pointers.Raw[gd.StringName](name).String()
}

// reloadsMarkSessionStale invalidates everything owned by the module
// that just exited: its extension instances become adoption candidates
// (see reloadsAdoptInstances), and callbacks referring to bindings or
// callables it owned are answered inertly from now on. Classes stay
// registered (see reloadsTrackRegistrations).
func reloadsMarkSessionStale() {
	reloadsEpoch.Add(1)
	clear(reloadsBindings)
	clear(reloadsCallables)
	reloadsClassesMu.Lock()
	clear(reloadsDeduped)
	reloadsClassesMu.Unlock()
}

// reloadsInstanceLive reports whether the extension instance id belongs
// to the current module.
func reloadsInstanceLive(id uint64) bool {
	info, ok := reloadsInstances[id]
	return ok && info.epoch == reloadsEpoch.Load()
}

func reloadsBindingLive(id uint64) bool  { return reloadsBindings[id] }
func reloadsCallableLive(id uint64) bool { return reloadsCallables[id] }

// reloadsAdoptInstances hands engine objects that outlived the previous
// module over to the current one: each surviving instance is re-bound
// to a fresh Go instance (zero-valued state) of the same class in the
// new module, so live objects keep dispatching into live code.
func reloadsAdoptInstances() {
	if reloadsAdoptFn == nil {
		return
	}
	current := reloadsEpoch.Load()
	type pending struct {
		id   uint64
		info reloadsInstanceInfo
	}
	var stale []pending
	for id, info := range reloadsInstances {
		if info.epoch != current {
			stale = append(stale, pending{id, info})
		}
	}
	for _, p := range stale {
		token, ok := reloadsClassToken(p.info.class)
		if !ok {
			continue // class no longer exists; the object stays inert
		}
		stack := []uint64{uint64(token), p.info.obj}
		if err := reloadsAdoptFn.CallWithStack(reloadsCtx, stack); err != nil {
			fmt.Fprintln(os.Stderr, "graphics.gd: failed to adopt instance of", p.info.class+":", err)
			continue
		}
		delete(reloadsInstances, p.id)
		if newID := stack[0]; newID != 0 {
			reloadsInstances[newID] = reloadsInstanceInfo{obj: p.info.obj, class: p.info.class, epoch: current}
		}
	}
	if len(stale) > 0 {
		fmt.Fprintln(os.Stderr, "graphics.gd: handed", len(stale), "instances over to the new build")
	}
}

// reloadsMapToken translates a method function token (or virtual call
// data pointer) minted by an earlier module into the current module's
// equivalent. Engine method binds cache these from the session that
// registered the class, so calls arriving after a swap must re-resolve.
func reloadsMapToken(token uint64) uint64 {
	info, ok := reloadsTokens[token]
	if !ok || info.epoch == reloadsEpoch.Load() {
		return token
	}
	if current, ok := reloadsCurrentTokens[info.key]; ok {
		if currentInfo, ok := reloadsTokens[current]; ok && currentInfo.epoch == reloadsEpoch.Load() {
			return current
		}
	}
	// Re-resolve through the class callbacks (records the new token).
	engineID, ok := reloadsEngineClasses[info.key.class]
	if !ok {
		return token
	}
	name := gd.NewStringName(info.key.method)
	defer name.Free()
	switch info.kind {
	case reloadsTokenVirtual:
		if fn := gdextension.On.Extension.Class.Method(engineID, pointers.Get(name), info.key.hash); fn != 0 {
			return uint64(fn)
		}
	case reloadsTokenCaller:
		if data := gdextension.On.Extension.Class.Caller(engineID, pointers.Get(name), info.key.hash); data != 0 {
			return uint64(data)
		}
	}
	return token
}

// reloadsGuardStale wraps every callback in the given group whose
// arguments carry an id of the given type: if the id is not owned by
// the live module, the callback returns zero values instead of being
// forwarded. Must run after reloadsInstallGuestCallbacks.
func reloadsGuardStale(group any, idType reflect.Type, live func(uint64) bool) {
	v := reflect.ValueOf(group).Elem()
	for i := range v.NumField() {
		f := v.Field(i)
		if f.Kind() != reflect.Func || f.IsNil() {
			continue
		}
		var idArgs []int
		for j := range f.Type().NumIn() {
			if f.Type().In(j) == idType {
				idArgs = append(idArgs, j)
			}
		}
		if len(idArgs) == 0 {
			continue
		}
		prev := f.Interface()
		ftype := f.Type()
		f.Set(reflect.MakeFunc(ftype, func(args []reflect.Value) []reflect.Value {
			for _, j := range idArgs {
				if !live(args[j].Uint()) {
					out := make([]reflect.Value, ftype.NumOut())
					for k := range out {
						out[k] = reflect.Zero(ftype.Out(k))
					}
					return out
				}
			}
			return reflect.ValueOf(prev).Call(args)
		}))
	}
}

// reloadsInstallSwapGuards adapts the generated guest forwarders for
// module swaps: class ids remap to the live module, and callbacks for
// dead-module state answer inertly.
func reloadsInstallSwapGuards() {
	// Class callbacks: remap the engine's (first-session) class id to
	// the current module's token.
	class := &gdextension.On.Extension.Class
	createClass := class.Create
	class.Create = func(id gdextension.ExtensionClassID, notify bool) gdextension.Object {
		if mapped, ok := reloadsClassRemap[id]; ok {
			id = mapped
		}
		return createClass(id, notify)
	}
	method := class.Method
	class.Method = func(id gdextension.ExtensionClassID, name gdextension.StringName, hash uint32) gdextension.FunctionID {
		engineID := id
		if mapped, ok := reloadsClassRemap[id]; ok {
			id = mapped
		}
		fn := method(id, name, hash)
		if fn != 0 {
			if className, ok := reloadsEngineClassNames[engineID]; ok {
				key := reloadsMethodKey{class: className, method: reloadsStringNameToString(name), hash: hash}
				reloadsCurrentTokens[key] = uint64(fn)
				reloadsTokens[uint64(fn)] = reloadsTokenInfo{key: key, kind: reloadsTokenVirtual, epoch: reloadsEpoch.Load()}
			}
		}
		return fn
	}
	caller := class.Caller
	class.Caller = func(id gdextension.ExtensionClassID, name gdextension.StringName, hash uint32) uintptr {
		engineID := id
		if mapped, ok := reloadsClassRemap[id]; ok {
			id = mapped
		}
		data := caller(id, name, hash)
		if data != 0 {
			if className, ok := reloadsEngineClassNames[engineID]; ok {
				key := reloadsMethodKey{class: className, method: reloadsStringNameToString(name), hash: hash}
				reloadsCurrentTokens[key] = uint64(data)
				reloadsTokens[uint64(data)] = reloadsTokenInfo{key: key, kind: reloadsTokenCaller, epoch: reloadsEpoch.Load()}
			}
		}
		return data
	}
	// Track wrapper bindings as they are created, then guard everything
	// by liveness of the ids involved.
	binding := &gdextension.On.Extension.Binding
	created := binding.Created
	binding.Created = func(instance gdextension.ExtensionInstanceID) gdextension.ExtensionBindingID {
		id := created(instance)
		reloadsBindings[uint64(id)] = true
		return id
	}
	reloadsGuardStale(binding, reflect.TypeFor[gdextension.ExtensionBindingID](), reloadsBindingLive)
	reloadsGuardStale(&gdextension.On.Extension.Instance, reflect.TypeFor[gdextension.ExtensionInstanceID](), reloadsInstanceLive)
	reloadsGuardStale(&gdextension.On.Extension.Script, reflect.TypeFor[gdextension.ExtensionInstanceID](), reloadsInstanceLive)
	reloadsGuardStale(&gdextension.On.Callables, reflect.TypeFor[gdextension.FunctionID](), reloadsCallableLive)
	// Method binds and virtual call data the engine cached in earlier
	// sessions carry tokens from a discarded module: remap them to the
	// current module before forwarding.
	checked := gdextension.On.Extension.Instance.CheckedCall
	gdextension.On.Extension.Instance.CheckedCall = func(instance gdextension.ExtensionInstanceID, fn gdextension.FunctionID, result gdextension.Returns[any], args gdextension.Accepts[any]) {
		checked(instance, gdextension.FunctionID(reloadsMapToken(uint64(fn))), result, args)
	}
	variantCall := gdextension.On.Extension.Instance.VariantCall
	gdextension.On.Extension.Instance.VariantCall = func(instance gdextension.ExtensionInstanceID, fn gdextension.FunctionID, result gdextension.Returns[gdextension.Variant], args gdextension.Accepts[gdextension.Variant]) {
		variantCall(instance, gdextension.FunctionID(reloadsMapToken(uint64(fn))), result, args)
	}
	dynamicCall := gdextension.On.Extension.Instance.DynamicCall
	gdextension.On.Extension.Instance.DynamicCall = func(instance gdextension.ExtensionInstanceID, fn gdextension.FunctionID, result gdextension.Returns[gdextension.Variant], arg_count int, args gdextension.Accepts[gdextension.Variant], err gdextension.Returns[gdextension.CallError]) {
		dynamicCall(instance, gdextension.FunctionID(reloadsMapToken(uint64(fn))), result, arg_count, args, err)
	}
	called := gdextension.On.Extension.Instance.Called
	gdextension.On.Extension.Instance.Called = func(instance gdextension.ExtensionInstanceID, callData gdextension.Pointer, result gdextension.Returns[any], args gdextension.Accepts[any]) {
		called(instance, gdextension.Pointer(reloadsMapToken(uint64(callData))), result, args)
	}
	// Forget instances and callables when the engine releases them.
	free := gdextension.On.Extension.Instance.Free
	gdextension.On.Extension.Instance.Free = func(id gdextension.ExtensionInstanceID) {
		free(id)
		delete(reloadsInstances, uint64(id))
	}
	freeCallable := gdextension.On.Callables.Free
	gdextension.On.Callables.Free = func(id gdextension.FunctionID) {
		freeCallable(id)
		delete(reloadsCallables, uint64(id))
	}
}

// reloadsForwardCallbacks chains the host's own engine/main-loop
// callbacks (which keep the host runtime linked and pumping) with
// forwards into the live guest module.
func reloadsForwardCallbacks() {
	if !reloadsCShared && reloadsEngine().Library != (Startup.Instance{}) {
		// Static startup path: the core/servers initialization levels
		// fired during process initialization, before this chain
		// existed. The scene (and editor) levels fire later, inside
		// startup.Start(), and are recorded (and forwarded) live.
		reloadsInitLevels = []gdextension.InitializationLevel{0, 1}
	}
	engineInit := gdextension.On.Engine.Init
	gdextension.On.Engine.Init = func(level gdextension.InitializationLevel) {
		reloadsInitLevels = append(reloadsInitLevels, level)
		engineInit(level)
		// Forward live to the guest: the scene level must reach it only
		// once the engine has its own scene classes registered, which is
		// exactly when this callback fires.
		if g := reloadsGuest.Load(); g != nil {
			reloadsCall(g.on_engine_init, []uint64{uint64(level)})
		}
	}
	engineExit := gdextension.On.Engine.Exit
	gdextension.On.Engine.Exit = func(level gdextension.InitializationLevel) {
		if reloadsCShared && level == 2 && reloadsExitCode.Load() != reloadsShutdown {
			// The engine is shutting down: unwind the guest session (its
			// cleanups need the engine alive) before the host's own exit
			// handling runs.
			reloadsExitCode.Store(reloadsShutdown)
			if reloadsParked.Load() {
				resume_main()
			}
		}
		engineExit(level)
	}
	firstFrame := gdextension.On.MainLoop.FirstFrame
	gdextension.On.MainLoop.FirstFrame = func() {
		firstFrame()
		if g := reloadsGuest.Load(); g != nil {
			reloadsCall(g.on_first_frame, nil)
		}
	}
	everyFrame := gdextension.On.MainLoop.EveryFrame
	gdextension.On.MainLoop.EveryFrame = func() {
		everyFrame()
		if g := reloadsGuest.Load(); g != nil {
			reloadsCall(g.on_every_frame, nil)
		}
		if reloadsCShared && reloadsParked.Load() && reloadsWakeWanted() {
			resume_main()
		}
	}
	finalFrame := gdextension.On.MainLoop.FinalFrame
	gdextension.On.MainLoop.FinalFrame = func() {
		finalFrame()
		if g := reloadsGuest.Load(); g != nil {
			reloadsCall(g.on_final_frame, nil)
		}
	}
}

func reloadsInstantiateControl() error {
	_, err := reloadsRuntime.NewHostModuleBuilder("reloads").
		NewFunctionBuilder().WithGoModuleFunction(api.GoModuleFunc(reloadsHostReady), nil, nil).Export("ready").
		NewFunctionBuilder().WithGoModuleFunction(api.GoModuleFunc(reloadsHostYield), nil, []api.ValueType{api.ValueTypeI32}).Export("yield").
		Instantiate(reloadsCtx)
	return err
}

func reloadsHostReady(_ context.Context, m api.Module, stack []uint64) {
	g := reloadsBindGuest(m)
	reloadsGuest.Store(g)
	// Replay the initialization levels the engine has fired so far, so
	// the guest links itself exactly like a natively-started extension
	// would. On the first static-path session only core/servers have
	// fired; the scene level reaches the guest live, via the Engine.Init
	// chain, when startup.Start() below completes engine setup. On later
	// sessions (and the shared path) all levels replay here and the
	// guest's queued class registrations flush at the scene level.
	for _, level := range reloadsInitLevels {
		reloadsCall(g.on_engine_init, []uint64{uint64(level)})
	}
	reloadsAdoptFn = m.ExportedFunction("reloads_adopt")
	if reloadsNeedsStart {
		reloadsNeedsStart = false
		// Static startup path: complete engine scene setup now that the
		// first guest has registered its classes (the project's main
		// loop type may be one of them).
		startup.Start()
	}
	// Hand engine objects that survived a module swap over to this
	// module, now that its classes are registered.
	reloadsAdoptInstances()
	reloadsClassesMu.Lock()
	fmt.Fprintln(os.Stderr, "graphics.gd: serving classes:", reloadsClasses)
	reloadsClassesMu.Unlock()
}

func reloadsHostYield(_ context.Context, m api.Module, stack []uint64) {
	if code := reloadsExitCode.Load(); code != reloadsKeepRunning {
		stack[0] = uint64(code)
		return
	}
	if reloadsCShared {
		// The engine drives its own loop: park the main coroutine until
		// the frame chain wakes us with something to do.
		reloadsPark()
		stack[0] = uint64(reloadsExitCode.Load())
		return
	}
	if reloadsEngine().Library.Iteration() {
		reloadsExitCode.Store(reloadsShutdown)
		stack[0] = reloadsShutdown
		return
	}
	stack[0] = reloadsKeepRunning
}

// reloadsPrepareRuntime brings up the wasm runtime and builds the first
// guest, leaving reloadsRuntime closed and nil if anything fails. It
// touches nothing engine-side, so a failure can still be recovered from
// by running the project natively.
func reloadsPrepareRuntime() error {
	runtimeConfig := wazero.NewRuntimeConfig()
	if dir, err := os.UserCacheDir(); err == nil {
		// Persist compiled wasm across runs so session startup is fast.
		if cache, err := wazero.NewCompilationCacheWithDir(filepath.Join(dir, "graphics.gd", "reloads")); err == nil {
			runtimeConfig = runtimeConfig.WithCompilationCache(cache)
		}
	}
	reloadsRuntime = wazero.NewRuntimeWithConfig(reloadsCtx, runtimeConfig)
	prepare := func() error {
		if _, err := wasi_snapshot_preview1.Instantiate(reloadsCtx, reloadsRuntime); err != nil {
			return fmt.Errorf("failed to instantiate wasi: %w", err)
		}
		if err := reloadsInstantiateHostModule(reloadsCtx, reloadsRuntime); err != nil {
			return fmt.Errorf("failed to instantiate gd bridge: %w", err)
		}
		if err := reloadsInstantiateControl(); err != nil {
			return fmt.Errorf("failed to instantiate reloads control: %w", err)
		}
		fmt.Fprintln(os.Stderr, "graphics.gd: building wasm guest...")
		return reloadsBuildGuest()
	}
	if err := prepare(); err != nil {
		reloadsRuntime.Close(reloadsCtx)
		reloadsRuntime = nil
		return err
	}
	return nil
}

// reloadsFallback gives up on hot reloading and finishes the run the way
// a build without the reloads tag would. Nothing engine-side has been
// touched at this point, and the host binary is the project itself, so
// every class the project defines is compiled in and the editor (or the
// game) stays completely usable — only live swapping is lost.
func reloadsFallback(reason string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "graphics.gd: hot reloading is off (%s):\n%v\n", reason, err)
	} else {
		fmt.Fprintf(os.Stderr, "graphics.gd: hot reloading is off (%s)\n", reason)
	}
	reloadsSession = nil // Scene takes the ordinary path from here on.
	Scene()
}

func reloadsRun() {
	if startup == nil {
		startup = new(engineAsSharedLibrary)
	}
	_, isLibrary := startup.(interface{ reloadsBase() *engineAsLibrary })
	reloadsCShared = !isLibrary
	if engineStarted {
		// The project called startup.LoadingScene() before Scene(), which
		// is its documented right: the engine is already all the way up.
		// The host has therefore missed the scene-level initialization it
		// has to link each guest against, and has already registered its
		// own classes under the names the guest would claim. There is no
		// way to install hot reloading after the fact, so run the host
		// build, which is the whole project anyway.
		reloadsFallback("startup.LoadingScene starts the engine before the reloads host can take it over", nil)
		return
	}
	// The host binary is the project itself, compiled with the reloads
	// tag, so its sources are known to build for this platform: a guest
	// build that fails does so for wasip1/wasm specifically (a cgo
	// dependency, a build constraint) and no later save will clear it.
	// Probe it before touching the engine, so that a project which
	// cannot target wasm runs the host's classes instead of opening an
	// editor with none of them.
	if err := reloadsPrepareRuntime(); err != nil {
		reloadsFallback("the project does not build for wasip1/wasm", err)
		return
	}
	defer reloadsRuntime.Close(reloadsCtx)
	if reloadsCShared {
		// Frame-driven fastcb residency permanently locks the engine
		// thread on the first frame, which changes the OS-thread-lock
		// state between the main coroutine's creation and later resumes
		// — a fatal mismatch for runtime.coroswitch. The reloads host
		// gains nothing from residency (all project code runs in the
		// wasm guest), so keep the stock callback path. Runs before the
		// first frame, so fastcbFrame never engages.
		fastcbFrameDriving = false
	}
	loadingSceneWasCalled = true
	reloadsForwardCallbacks()
	reloadsTrackRegistrations()
	reloadsInstallGuestCallbacks()
	reloadsInstallSwapGuards()
	if os.Getenv("GD_RELOADS_TRACE") != "" {
		create := gdextension.On.Extension.Class.Create
		gdextension.On.Extension.Class.Create = func(id gdextension.ExtensionClassID, notify bool) gdextension.Object {
			obj := create(id, notify)
			fmt.Fprintf(os.Stderr, "graphics.gd: class create id=%x -> obj=%x\n", uintptr(id), uintptr(obj))
			return obj
		}
		var lookups int
		lookup := gdextension.Host.Objects.Method.Lookup
		gdextension.Host.Objects.Method.Lookup = func(name, method gdextension.StringName, hash int64) gdextension.MethodForClass {
			bind := lookup(name, method, hash)
			if lookups < 40 || bind == 0 && lookups < 400 {
				lookups++
				fmt.Fprintf(os.Stderr, "graphics.gd: method_lookup(%q, %q, %d) -> %x\n",
					reloadsStringNameToString(name), reloadsStringNameToString(method), hash, uintptr(bind))
			}
			return bind
		}
		make_ := gdextension.Host.Objects.Make
		gdextension.Host.Objects.Make = func(name gdextension.StringName) gdextension.Object {
			obj := make_(name)
			fmt.Fprintf(os.Stderr, "graphics.gd: object_make(%q) -> %x\n", reloadsStringNameToString(name), uintptr(obj))
			return obj
		}
	}

	switch {
	case reloadsCShared:
		// The extension was loaded into a godot process: we are running
		// as the main coroutine, on the engine main thread, inside the
		// scene-level initialization callback — so levels 0..2 have
		// already fired natively and the engine is mid-initialization.
		// Sessions replay them into each guest; later levels (editor)
		// reach the guest live via the Engine.Init chain.
		reloadsInitLevels = []gdextension.InitializationLevel{0, 1, 2}
	case reloadsEngine().Library == (Startup.Instance{}):
		// Shared library path: the engine does not exist until Start(),
		// so it runs up front and the guest registers as soon as the
		// engine is live.
		startup.Start()
	default:
		// Static startup path: the engine instance is created during
		// process initialization and startup.Start() completes scene
		// setup — which resolves the project's main loop type. The
		// guest's classes (including GoMainLoop) must be registered
		// first, so Start() is deferred into reloadsHostReady, which
		// runs right after the first guest's registrations flush.
		reloadsNeedsStart = true
	}

	go reloadsWatch()

	for session := 0; ; session++ {
		reloadsExitCode.Store(reloadsKeepRunning)
		if reloadsCompiled.Load() == nil {
			// No build has succeeded yet. If engine scene setup is still
			// pending (static path), it waits for the first successful
			// build — the project's main loop class comes from the
			// guest. Otherwise the engine runs host-driven so the editor
			// stays usable while the project is broken. Either way the
			// watcher keeps rebuilding on source changes.
			if reloadsCShared {
				// The engine drives itself; park the coroutine until a
				// build succeeds or the engine shuts down.
				fmt.Fprintln(os.Stderr, "graphics.gd: waiting for the project to compile...")
				reloadsAwaitingBuild.Store(true)
				for reloadsCompiled.Load() == nil && reloadsExitCode.Load() != reloadsShutdown {
					reloadsPark()
				}
				reloadsAwaitingBuild.Store(false)
				if reloadsExitCode.Load() == reloadsShutdown {
					break
				}
				continue
			}
			if reloadsNeedsStart {
				// The editor can open without the project's Go classes
				// (they hot-add once the project compiles), but a game
				// run cannot: its main loop class comes from the guest.
				if slices.Contains(os.Args, "-e") || slices.Contains(os.Args, "--editor") {
					reloadsNeedsStart = false
					startup.Start()
				} else {
					fmt.Fprintln(os.Stderr, "graphics.gd: waiting for the project to compile...")
					for reloadsCompiled.Load() == nil {
						time.Sleep(100 * time.Millisecond)
					}
					continue
				}
			}
			for reloadsCompiled.Load() == nil {
				if reloadsEngine().Library.Iteration() {
					reloadsExitCode.Store(reloadsShutdown)
					break
				}
			}
			if reloadsExitCode.Load() == reloadsShutdown {
				break
			}
			continue
		}
		compiled := *reloadsCompiled.Load()
		config := wazero.NewModuleConfig().
			WithName(fmt.Sprintf("guest-%d", session)).
			WithStdout(os.Stdout).
			WithStderr(os.Stderr).
			WithStdin(os.Stdin).
			WithArgs("library.wasm").
			WithSysWalltime().
			WithSysNanotime().
			WithFSConfig(wazero.NewFSConfig().WithDirMount(reloadsProjectDir(), "/"))
		for _, env := range os.Environ() {
			if k, v, ok := strings.Cut(env, "="); ok {
				config = config.WithEnv(k, v)
			}
		}
		mod, err := reloadsRuntime.InstantiateModule(reloadsCtx, compiled, config)
		if err != nil {
			if exit, ok := err.(*sys.ExitError); !ok || exit.ExitCode() != 0 {
				fmt.Fprintln(os.Stderr, "graphics.gd: wasm guest failed:", err)
			}
		}
		reloadsGuest.Store(nil)
		if mod != nil {
			mod.Close(reloadsCtx)
		}
		if reloadsExitCode.Load() != reloadsSwap {
			break
		}
		reloadsMarkSessionStale()
		fmt.Fprintln(os.Stderr, "graphics.gd: reloaded")
	}

	if !reloadsCShared {
		if lib := reloadsEngine(); lib.destroy != nil {
			lib.destroy()
		}
	}
}
