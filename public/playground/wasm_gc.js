// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// wasm_gc.js runs a GOWASM=gc program: the part of the host that does
// not depend on where it runs. wasm_gc_node.mjs (node) and
// wasm_gc_browser.js (browsers) give it the system interface (WASI) and
// what the engine allows them.
//
// The program needs WebAssembly GC and JSPI (WebAssembly.Suspending and
// WebAssembly.promising): each goroutine runs on a stack of its own.

import { gojs } from './wasm_gc_js.js';

// runWasmGC instantiates module and runs the program. host has:
//   wasi:      the imports of wasi_snapshot_preview1, and start(instance)
//              to give it the memory (or null);
//   write(fd, text): a line of output from the runtime's print (fd 2);
//   exit(code): the program ended;
//   collect(major): collect garbage now, if the engine allows it (else
//              null): runtime.GC, and after memory grows (V8 aborts if
//              it collects for the memory later, on a goroutine's stack);
//   nanotime(): a monotonic clock, as a BigInt of nanoseconds;
//   trace:     print the stack when the program fails;
//   poison, maxStacks: for testing (see below).
// For a GOOS=js program (one that imports module gojs: see
// wasm_gc_js.js), also:
//   global:    what js.Global() is (globalThis by default);
//   idle(deadlock): no goroutine can run and no timer is set; the
//              program waits for JavaScript to call it. deadlock()
//              ends it if it is still so (node calls it when its event
//              loop is empty).
// It returns a promise that is settled when the program has started.
export async function runWasmGC(module, host) {
  // Whether this is a GOOS=js program; null if it is not known until it
  // is instantiated (see moduleImports).
  const imported = moduleImports(module);
  let isJS = imported ? imported.some((i) => i.module === 'gojs') : null;
  let exited = false;
  let line = [];
  const flush = () => { host.write(2, new TextDecoder().decode(new Uint8Array(line))); line = []; };
  let instance;
  const objids = new WeakMap();
  let lastid = 0;
  // Without a way to make the engine collect (a browser's page has
  // none), it is told what the program allocates: an engine counts the
  // buffers that are allocated, and collects when they are many, and
  // knows nothing of what is in linear memory. When a collection is due
  // for what was allocated there, as much is allocated in buffers and
  // dropped: the engine collects when it would have for that much of
  // its own, and what it collects is freed in callbacks, then or at a
  // later time that the program waits. For runtime.GC, which asks for a
  // collection now, 128 MB of them, and a moment for the callbacks; not
  // more than five times a second (a program may call it in a loop).
  let freed = 0; // the blocks of linear memory freed
  let pressed = -Infinity;
  const pressure = (major, allocated) => {
    if (!major) return undefined;
    if (allocated > 0) {
      for (let left = allocated; left > 0; left -= 16 << 20) new ArrayBuffer(Math.min(left, 16 << 20));
      return undefined;
    }
    if (performance.now() - pressed < 200) return undefined;
    pressed = performance.now();
    for (let i = 0; i < 8; i++) new ArrayBuffer(16 << 20);
    const before = freed;
    return new Promise((done) => {
      let n = 0;
      const wait = () => (freed !== before || ++n > 10 ? done() : setTimeout(wait, 5));
      setTimeout(wait, 0);
    });
  };
  const collect = host.collect || pressure;

  // A block of linear memory is freed when the object that owns it is
  // collected. (Freeing a block may queue its cleanups: see
  // runtime/wasmgc/linear.go.)
  let free, slotsof, mkptr; // (instance.exports is slow to read: these are kept)
  const registry = new FinalizationRegistry((addr) => {
    freed++;
    if (host.poison) {
      // For testing: what is read from a block after it is freed is
      // seen to be wrong (an address in it is out of memory), not what
      // was there. (The header of a block holds its size: see
      // runtime/wasmgc/linear.go.)
      const mem = instance.exports.mem.buffer;
      const size = 2 ** new Uint32Array(mem, addr - 8, 1)[0] - 8;
      new Uint8Array(mem, addr, size).fill(0xA5);
    }
    free(BigInt(addr));
    wakeIfIdle();
  });

  // Goroutines (see runtime/wasmgc/proc.go). Each runs on a stack of its
  // own: one that blocks calls park, which suspends it until wake is
  // called.
  const parked = new Map();
  let idle = false;
  // Each time the program goes idle is numbered: a timeout set in one
  // must not wake (or declare deadlock in) a later one, nor run the
  // scheduler while a goroutine runs.
  let idleEpoch = 0;
  const wakeIfIdle = () => {
    if (idle) {
      idle = false;
      schedule();
    }
  };

  // The cleanups of runtime.AddCleanup: the runtime is told when the
  // object is collected, and a goroutine runs the cleanup.
  const cleanupTokens = new Map();
  const cleanups = new FinalizationRegistry((id) => {
    cleanupTokens.delete(id);
    instance.exports.cleanup(id);
    wakeIfIdle();
  });

  // Memory grows on the host's stack (see runtime/wasmgc/linear.go), and
  // its buffer is made here: the first look at it, from a goroutine's
  // stack, could abort.
  function growsync(pages) {
    const r = instance.exports.mem.grow(pages);
    instance.exports.mem.buffer.byteLength;
    return r;
  }
  function growmem(pages) {
    const r = instance.exports.mem.grow(pages);
    instance.exports.mem.buffer.byteLength;
    if (pages >= 16) collect(false);
    return r;
  }

  // A goroutine that ran out of the engine's stack (the engine's stacks
  // do not grow) fails as Go's do: a fatal error, which deferred calls
  // do not see. id is the goroutine's.
  const isOverflow = (err) =>
    (err instanceof RangeError && /call stack/i.test(err.message)) || // V8, JavaScriptCore
    (err && err.name === 'InternalError' && /recursion/i.test(err.message)); // SpiderMonkey
  const fail = (err, id = 1) => {
    if (err && err.goExit !== undefined) return; // (the program exited)
    if (line.length) flush();
    if (err && /(continuation|suspendable) stacks/.test(String(err.message))) {
      // (A stack that grows takes another of the engine's: see hop.)
      host.write(2, 'fatal error: the WebAssembly engine has no more stacks (too many goroutines are blocked, or their stacks have grown)\n');
      host.exit(2);
      return;
    }
    if (isOverflow(err)) {
      const out = ["runtime: goroutine stack exceeds the WebAssembly engine's stack size", 'fatal error: stack overflow', '', `goroutine ${id} [running]:`];
      // (As many frames as the engine keeps in an error's stack.)
      for (const l of String(err.stack).split('\n')) {
        if (l.includes('wasm-function')) out.push(l.trim().replace(/^at /, '').split(/[ @(]/)[0] + '(...)');
      }
      out.push('...additional frames elided...');
      host.write(2, out.join('\n'));
      host.exit(2);
      return;
    }
    // (Firefox's stack does not begin with the message, as V8's does.)
    let text = String(err && err.stack || err);
    if (err && err.stack && !text.includes(String(err.message))) text = `${err}\n${text}`;
    host.write(2, text);
    if (host.trace) host.write(2, '\npanicline ' + instance.exports.panicline?.value + '\n');
    host.exit(2);
  };
  let gorun;
  function schedule() {
    if (deferred.length > 0 && !exited) {
      // A call that came while the runtime was busy: now nothing runs
      // (and callGo schedules again after it).
      idle = true;
      deferred.shift()();
      return;
    }
    const next = instance.exports.sched();
    if (next < 0n) {
      // Nothing can run until a timer fires, if one is set, or a
      // cleanup is due (see cleanups).
      idle = true;
      const epoch = ++idleEpoch;
      const still = () => idle && epoch === idleEpoch;
      const wait = instance.exports.timerwait();
      if (wait < 0n && isJS) {
        // (Or until JavaScript calls into Go.)
        if (host.idle) {
          host.idle(() => {
            if (!still() || exited) return;
            try { instance.exports.deadlock(); } catch (err) { fail(err); return; }
            idle = false;
            schedule();
          });
        }
      } else if (wait < 0n) {
        // Before deadlock is declared, what was collected may have
        // cleanups to run.
        if (host.collect) host.collect(true);
        setTimeout(() => {
          if (!still()) return;
          // (It returns if goroutines only wait for a stack of the
          // engine's, to try again.)
          try { instance.exports.deadlock(); } catch (err) { fail(err); return; }
          idle = false;
          schedule();
        }, 10);
      } else {
        // (A longer delay than 2^31-1 ms would fire at once.)
        setTimeout(() => { if (still()) { idle = false; schedule(); } }, Math.min(Number(wait / 1000000n), 0x7fffffff));
      }
      return;
    }
    idle = false;
    const id = Number(next >> 1n);
    if (next & 1n) {
      queueMicrotask(() => {
        let run;
        try { run = gorun(BigInt(id)); } catch (err) { run = Promise.reject(err); }
        run.then(schedule, (err) => {
          // If the goroutine never ran, the engine had no stack for it
          // (Firefox has some 2,000): it waits until one is free.
          if (!(err && err.goExit !== undefined) && instance.exports.unstart(BigInt(id))) schedule();
          else fail(err, id);
        });
      });
    } else {
      const wake = parked.get(id);
      parked.delete(id);
      wake();
    }
  }

  // Calls from JavaScript into Go (GOOS=js: the functions js.FuncOf
  // makes; see runtime/wasmgc/jscall.go). The function runs on a
  // goroutine of its own, started here on a stack of its own: if it
  // ends without blocking, the call returns its result; if it blocks,
  // the call returns a promise of the result, and the goroutine goes on
  // when it is woken, as any other. While it runs, whatever goroutine
  // called into JavaScript (if one did) waits in that call. Until the
  // call returns, the goroutine of the call is suspended only when it
  // blocks, which does not run the scheduler (park): an engine
  // suspends a stack at every call of a WebAssembly.Suspending
  // function, whatever it returns, and the runtime does without the
  // others then (see jssync in runtime/wasmgc/jscall.go).
  let syncDepth = 0;
  // While the runtime is busy (see jsbusy in runtime/wasmgc/jscall.go),
  // a call waits in deferred until the scheduler next runs.
  let busy = false;
  const deferred = [];
  function callGo(event) {
    if (exited) throw new Error('Go program has already exited');
    if (busy) {
      return new Promise((done, failed) => deferred.push(() => {
        try {
          Promise.resolve(callGo(event)).then(done, failed);
        } catch (err) {
          failed(err);
        }
      }));
    }
    const wasIdle = idle;
    idle = false;
    js.go._pendingEvent = event;
    const id = instance.exports.callback();
    let done;
    syncDepth++;
    try { done = gorun(id); } catch (err) { done = Promise.reject(err); }
    syncDepth--;
    if (!exited) instance.exports.callbackdone();
    const finished = 'result' in event;
    if (finished) done.catch((err) => fail(err, Number(id)));
    else done.then(schedule, (err) => fail(err, Number(id)));
    if (wasIdle && !exited) schedule();
    return finished ? event.result : done.then(() => event.result);
  }

  // The names of the functions, from the module's name section, for
  // runtime.FuncForPC.
  const funcNames = new Map();
  // (Without moduleImports, the functions named first: the names are
  // of the module's own functions, which come after the imported ones.)
  let nimports = imported ? imported.filter((i) => i.kind === 'function').length : -1;
  for (const sec of WebAssembly.Module.customSections(module, 'name')) {
    const b = new Uint8Array(sec);
    let i = 0;
    const leb = () => { let r = 0, s = 0, c; do { c = b[i++]; r |= (c & 0x7f) << s; s += 7; } while (c & 0x80); return r >>> 0; };
    while (i < b.length) {
      const id = b[i++], len = leb(), end = i + len;
      if (id === 1) {
        for (let n = leb(); n > 0; n--) {
          const idx = leb(), l = leb();
          if (nimports < 0) nimports = idx;
          funcNames.set(idx, new TextDecoder().decode(b.subarray(i, i + l)));
          i += l;
        }
      }
      i = end;
    }
  }

  // The files and lines of the code, from the section "go.lines" (which
  // the linker leaves out with -s), read when first asked: by function,
  // where each line starts as an offset in the module.
  let lineTable;
  const lineInfo = () => {
    if (lineTable) return lineTable;
    lineTable = { files: [], names: [], funcs: new Map() };
    for (const sec of WebAssembly.Module.customSections(module, 'go.lines')) {
      const b = new Uint8Array(sec);
      let i = 0;
      const leb = () => { let r = 0, s = 0, c; do { c = b[i++]; r += (c & 0x7f) * 2 ** s; s += 7; } while (c & 0x80); return r; };
      const sleb = () => { let r = 0, s = 0, c; do { c = b[i++]; r += (c & 0x7f) * 2 ** s; s += 7; } while (c & 0x80); return c & 0x40 ? r - 2 ** s : r; };
      const dec = new TextDecoder();
      for (const list of [lineTable.files, lineTable.names]) {
        for (let n = leb(); n > 0; n--) {
          const l = leb();
          list.push(dec.decode(b.subarray(i, i + l)));
          i += l;
        }
      }
      // A function's frames: several at one offset where calls were
      // inlined, the innermost first, each but the last with the name
      // of the function inlined.
      for (let n = leb(); n > 0; n--) {
        const idx = leb(), start = leb(), count = leb();
        const f = { start, offs: new Uint32Array(count), files: new Uint32Array(count), lines: new Int32Array(count), names: new Uint32Array(count) };
        let off = start, line = 0;
        for (let k = 0; k < count; k++) {
          off += leb();
          f.offs[k] = off;
          f.files[k] = leb();
          line += sleb();
          f.lines[k] = line;
          f.names[k] = leb();
        }
        lineTable.funcs.set(idx, f);
      }
    }
    return lineTable;
  };
  // lineFrames returns the frames of function idx for the code at off,
  // an offset in the module: the first and the last index in its
  // table, or [-1, -1].
  const lineFrames = (idx, off) => {
    const f = lineInfo().funcs.get(idx);
    if (!f) return [-1, -1];
    let lo = 0, hi = f.offs.length;
    while (lo < hi) {
      const mid = (lo + hi) >> 1;
      if (f.offs[mid] <= off) lo = mid + 1; else hi = mid;
    }
    let first = lo - 1;
    while (first > 0 && f.offs[first - 1] === f.offs[first]) first--;
    return [first, lo - 1];
  };

  // panicFrames returns the frames of function idx nearest to those
  // given whose line (the innermost's) is the one the program noted
  // for its panic: those given if there are none.
  const panicFrames = (idx, first, last) => {
    const line = instance.exports.panicline?.value;
    const f = lineTable.funcs.get(idx);
    if (!line || f.lines[first] === line) return [first, last];
    let best = -1;
    for (let e = 0; e < f.offs.length; e++) {
      if (e > 0 && f.offs[e - 1] === f.offs[e]) continue; // (not the innermost)
      if (f.lines[e] === line && (best < 0 || Math.abs(e - first) < Math.abs(best - first))) best = e;
    }
    if (best < 0) return [first, last];
    let end = best;
    while (end + 1 < f.offs.length && f.offs[end + 1] === f.offs[best]) end++;
    return [best, end];
  };

  const hops = new Map(); // the functions that hop called, by their index in the table
  const weakids = new WeakMap(), weakrefs = new Map();
  let lastweak = 0;

  let heapStart = 0;
  const imports = {
    gd: {
      heapstart: () => BigInt(heapStart),
      putc: (c) => { if (c === 10) flush(); else line.push(c); },
      nanotime: () => host.nanotime(),
      walltime: () => BigInt(Date.now()) * 1000000n,
      funcname: (pc, addr, n) => {
        const idx = Number(pc >> 16n) - 0x1000 + nimports;
        let name = funcNames.get(idx) || '';
        // (A frame of a call that was inlined: the function inlined.)
        const e = (Number(pc & 0xffffn) >> 1) - 1;
        if (e >= 0 && lineTable) {
          const f = lineTable.funcs.get(idx);
          if (f && e < f.names.length && f.names[e] > 0) name = lineTable.names[f.names[e] - 1];
        }
        const bytes = new TextEncoder().encode(name).subarray(0, n);
        new Uint8Array(instance.exports.mem.buffer, Number(addr), bytes.length).set(bytes);
        return bytes.length;
      },
      callers: (addr, n) => {
        const limit = Error.stackTraceLimit;
        Error.stackTraceLimit = n + 64;
        const stack = new Error().stack;
        Error.stackTraceLimit = limit;
        const pcs = new BigUint64Array(instance.exports.mem.buffer, Number(addr), n);
        let k = 0;
        // A pc is the function (its index << 16, as a func value is)
        // plus, in its low bits, 1 and twice the frame's place in the
        // function's table of lines (see pcline), if it has one: a
        // frame of the engine's is several where calls were inlined.
        // The frame that called one of the runtime's functions that
        // panic for the compiler (the innermost such) is at the line
        // the program noted (see wasm.GCGlobalPanicLine): the engine's
        // position for it may not be exact.
        let panicked = false, noted = false;
        // The record of the innermost function that defers calls, in
        // linear memory (see GCGlobalDefer in cmd/internal/obj/wasm):
        // such a function notes where each call it makes is, as the
        // engine may not give it (its code is in a try block).
        let rec = instance.exports.deferframe ? instance.exports.deferframe.value >>> 0 : 0;
        const mem = new DataView(instance.exports.mem.buffer);
        for (const m of stack.matchAll(/wasm-function\[(\d+)\](?::0x([0-9a-f]+))?/g)) {
          if (k === n) break;
          const idx = Number(m[1]);
          let noted_at = -1;
          if (rec && mem.getBigUint64(rec + 8, true) === BigInt(idx - nimports + 0x1000) << 16n) {
            noted_at = mem.getInt32(rec + 4, true);
            rec = mem.getUint32(rec, true);
          }
          // (Not the runtime's call of a function on a new segment of
          // stack, which is not a frame of Go's; nor the wrappers the
          // compiler makes for defer and go statements, which Go's
          // tracebacks leave out.)
          if (/^runtime\.morestackgc$|\.(deferwrap|gowrap)\d+$/.test(funcNames.get(idx) || '')) continue;
          const fix = panicked && !noted;
          panicked = !noted && /^runtime\.(panic|goPanic)/.test(funcNames.get(idx) || '');
          noted ||= fix;
          // (The engine gives where the call is, not where it returns to.)
          let [first, last] = m[2] === undefined ? [-1, -1] : lineFrames(idx, parseInt(m[2], 16));
          if (noted_at >= 0) {
            const f = lineInfo().funcs.get(idx);
            if (f) [first, last] = lineFrames(idx, f.start + noted_at);
          }
          if (fix && first >= 0) [first, last] = panicFrames(idx, first, last);
          const fn = first < 0 ? null : lineTable.funcs.get(idx);
          for (let e = first; e <= last && k < n; e++) {
            // (As Go's traceback: not the wrappers the compiler makes
            // for methods, which are of no file.)
            if (fn && lineTable.files[fn.files[e]] === '<autogenerated>') continue;
            pcs[k++] = (BigInt(idx - nimports + 0x1000) << 16n) + 1n + BigInt(2 * Math.min(e + 1, 0x7fff));
          }
        }
        return k;
      },
      // pcline writes the name of the file of the code at pc at addr and
      // returns its length << 32 and the line: 0 if they are not known.
      pcline: (pc, addr, n) => {
        const f = lineInfo().funcs.get(Number(pc >> 16n) - 0x1000 + nimports);
        const e = (Number(pc & 0xffffn) >> 1) - 1;
        if (!f || e < 0 || e >= f.offs.length) return 0n;
        const bytes = new TextEncoder().encode(lineTable.files[f.files[e]]).subarray(0, n);
        new Uint8Array(instance.exports.mem.buffer, Number(addr), bytes.length).set(bytes);
        return (BigInt(bytes.length) << 32n) | BigInt(f.lines[e]);
      },
      grow: new WebAssembly.Suspending((pages) => new Promise((done) => queueMicrotask(() => done(growmem(pages))))),
      // (On a goroutine's stack during a call from JavaScript that has
      // not returned (see callGo), the engine is not asked to collect.)
      growsync: (pages) => (syncDepth > 0 ? growsync(pages) : growmem(pages)),
      cleanup: (obj, id) => {
        if (slotsof) obj = slotsof(obj);
        const token = {};
        cleanupTokens.set(id, token);
        cleanups.register(obj, id, token);
      },
      uncleanup: (id) => {
        const token = cleanupTokens.get(id);
        if (token) {
          cleanups.unregister(token);
          cleanupTokens.delete(id);
        }
      },
      // runtime.GC: a collection (if the engine allows), from a task of
      // its own, on the host's stack; then the finalization callbacks may
      // run.
      gc: new WebAssembly.Suspending((allocated) => new Promise((done) => setTimeout(() => Promise.resolve(collect(true, allocated)).then(() => setTimeout(done, 0)), 0))),
      weakmake: (o) => {
        // (With GOWASM=gc,split, weak on the array of the block: see
        // register.)
        if (slotsof) o = slotsof(o);
        let id = weakids.get(o);
        if (id === undefined) {
          weakids.set(o, id = ++lastweak);
          weakrefs.set(id, new WeakRef(o));
        }
        return id;
      },
      weakget: (id) => {
        const o = weakrefs.get(id)?.deref();
        if (o === undefined) {
          weakrefs.delete(id);
          return null;
        }
        return o;
      },
      // With GOWASM=gc,split: the pointer at addr into the block whose
      // array weakmake gave id, off bytes from its first element, or
      // null if the block was collected.
      weakgetat: (id, addr, off, lim) => {
        const o = weakrefs.get(id)?.deref();
        if (o === undefined) {
          weakrefs.delete(id);
          return null;
        }
        return mkptr(addr, o, off, lim);
      },
      objid: (o) => {
        let id = objids.get(o);
        if (id === undefined) objids.set(o, id = ++lastid);
        return id;
      },
      // (With GOWASM=gc,split a pointer is made anew for each place it
      // points to: the block lives as long as its array of pointers,
      // which the module gives for a pointer as "slotsof".)
      register: (p, addr) => registry.register(slotsof ? slotsof(p) : p, addr),
      // A function whose frame does not fit what is left of its
      // goroutine's stack calls itself here (see GCHop in
      // cmd/internal/obj/wasm): with its arguments, whatever they are,
      // and then itself, as a func value. It is called on a new stack
      // of the engine's, which do not grow, and what it returns (or
      // throws) is returned to the first stack, suspended meanwhile.
      // (From a microtask: the host's stack unwinds first, or it would
      // hold a frame for every stack started this way.)
      hop: new WebAssembly.Suspending((...args) => {
        const index = Number(args.pop() >> 16n);
        let f = hops.get(index);
        if (!f) hops.set(index, f = WebAssembly.promising(instance.exports.table.get(index)));
        return new Promise((done, failed) => queueMicrotask(() => f(...args).then(done, failed)));
      }),
      // (The next goroutine is chosen from a microtask, on the host's
      // stack.)
      park: new WebAssembly.Suspending((id) => new Promise((wake) => { parked.set(id, wake); if (syncDepth === 0) queueMicrotask(schedule); })),
      exit: (code) => {
        exited = true;
        if (line.length) flush();
        if (code !== 0 && host.trace) host.write(2, new Error('exit ' + code).stack);
        host.exit(code);
      },
    },
    wasi_snapshot_preview1: host.wasi.imports,
  };
  // What the program imports of its own (//go:wasmimport of another
  // module than these, or more functions of one of them): the host's.
  for (const [m, fns] of Object.entries(host.imports || {})) imports[m] = Object.assign(imports[m] || {}, fns);
  let js = null;
  if (isJS !== false) {
    js = gojs({ get memory() { return instance.exports.mem; }, global: host.global ?? globalThis, call: callGo, busy: (on) => { busy = on; } });
    imports.gojs = js.imports;
    // (For the tests of syscall/js, as wasm_exec.js has them.)
    imports._gotest = {
      add: (a, b) => a + b,
      callExport: new WebAssembly.Suspending((a, b) => {
        const exports = instance.exports;
        return WebAssembly.promising(exports.testExport0)().then(() => WebAssembly.promising(exports.testExport)(a, b));
      }),
    };
  }
  instance = await WebAssembly.instantiate(module, imports);
  if (isJS === null) isJS = 'callback' in instance.exports;
  heapStart = instance.exports.mem.buffer.byteLength;
  ({ free, slotsof, mkptr } = instance.exports);
  host.wasi.start(instance);
  if (host.reserve) {
    // Memory reserved up front, from the host's stack: growing it later,
    // on a goroutine's stack, can abort an engine that cannot be told to
    // collect (see growmem).
    const have = instance.exports.mem.buffer.byteLength;
    if (host.reserve > have) {
      instance.exports.mem.grow(Math.ceil((host.reserve - have) / 65536));
      instance.exports.mem.buffer.byteLength;
    }
  }
  gorun = WebAssembly.promising(instance.exports.gorun);
  if (host.maxStacks) {
    // For testing: an engine that has only so many stacks (as Firefox
    // has some 2,000) refuses to start a goroutine beyond them.
    const run = gorun;
    let live = 1;
    gorun = (id) => {
      if (live >= host.maxStacks) return Promise.reject(new Error('too many continuation stacks'));
      live++;
      return run(id).finally(() => { live--; });
    };
  }
  WebAssembly.promising(instance.exports.run)().catch(fail);
  return instance;
}

// moduleImports returns WebAssembly.Module.imports(module), or null if
// the engine cannot say: Safari 27 throws for a module whose types use
// WebAssembly GC references ("unable to produce import descriptors"),
// as every GOWASM=gc program's do. Then a GOOS=js program is one whose
// instance exports callback (the js entry point of wasm_gc_js.js), and
// it is given the imports of both kinds of program.
export function moduleImports(module) {
  try {
    return WebAssembly.Module.imports(module);
  } catch {
    return null;
  }
}
