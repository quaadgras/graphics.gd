// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// wasm_gc_browser.js runs a GOWASM=gc program (GOOS=wasip1 or js) in a
// browser:
//
//	import { runGo } from './wasm_gc_browser.js';
//	const code = await runGo('prog.wasm', { args: ['prog', 'x'], stdout: (line) => ... });
//
// The browser needs WebAssembly GC and JSPI. The system interface is the
// standard output and error (by line), the clocks, random numbers, the
// arguments, the environment, and files as an Emscripten program has
// them (wasm_gc_fs.js). The options for files:
//
//   - fs: where the files are, an object with the interface of
//     Emscripten's FS. The FS of an Emscripten program on the page
//     (Module.FS) gives the Go program that program's files. The
//     default is memFS() of wasm_gc_fs.js: files in memory, lost when
//     the page is. null is no files (the calls fail with ENOSYS).
//   - persist: directories (of the default fs) kept in IndexedDB, where
//     Emscripten's IDBFS keeps a directory of the same name: read
//     before the program starts, written as it changes them and when it
//     exits. (With an FS of Emscripten's, that is Emscripten's to do:
//     FS.mount(IDBFS, ...) and FS.syncfs.)
//   - files: files to start with, their contents (bytes or a string)
//     by path.
//
// The working directory is the environment's PWD, or /. options.wasi, a
// function of the program's memory (a function that returns its
// buffer), gives calls of wasi_snapshot_preview1 that replace these
// (wasm_gc_browser_exec.mjs, which runs tests in a browser, has the
// files of the machine it is on).
//
// Unlike under node, the engine cannot be told to collect garbage
// (unless the page has a way, which it gives as options.collect):
// runtime.GC only yields (weak pointers and cleanups go when the browser
// collects), and the host reserves memory up front (options.reserve, 64
// MB by default), since memory that grows on a goroutine's stack can
// make V8 collect there, which it does not survive.
//
// A GOOS=js program (one that imports gojs) runs where runGo is called:
// js.Global() is that global (options.global), and the program can use
// the page (document...) if that is a page's main thread. Its files are
// those of options.fs too, through the globals fs and process that
// package syscall uses (node's interface: nodeFS in wasm_gc_fs.js; and
// path, for path.resolve),
// which runGo sets on the global if it has none of its own. A program
// that waits for JavaScript to call it (js.FuncOf) does not end:
// runGo's promise is settled only when the program exits.
//
// A program that imports functions of its own (//go:wasmimport) is given
// them as options.imports, an object of modules of functions as an
// import object is, which is added to what the runtime imports. And
// options.started, if set, is called with the instance when it is made,
// before the program runs: for a page that calls what the program
// exports (//go:wasmexport) or needs its memory (exports.mem).

import { moduleImports, runWasmGC } from './wasm_gc.js';
import { memFS, nodeFS, wasiFS } from './wasm_gc_fs.js';

const ENOSYS = 52, EBADF = 8, ESUCCESS = 0;

// runGo runs the program at source (a URL, a Response, bytes or a
// WebAssembly.Module) and returns a promise of its exit code.
// ours marks the globals a run of runGo set (see jsGlobals).
const ours = Symbol('runGo');

export async function runGo(source, options = {}) {
  const args = options.args || ['prog'];
  const env = Object.entries(options.env || {}).map(([k, v]) => `${k}=${v}`);
  const out = { 1: options.stdout || ((t) => console.log(t)), 2: options.stderr || ((t) => console.error(t)) };
  let module;
  if (source instanceof WebAssembly.Module) {
    module = source;
  } else if (typeof source === 'string' || source instanceof URL || source instanceof Response) {
    module = await WebAssembly.compileStreaming(source instanceof Response ? source : fetch(source));
  } else {
    module = await WebAssembly.compile(source);
  }

  const fs = options.fs === undefined ? memFS() : options.fs;
  if (options.persist && !(fs && fs.persist)) throw new Error('runGo: persist is for the default fs');
  for (const dir of options.persist || []) await fs.persist(dir);
  for (const [path, data] of Object.entries(options.files || {})) {
    fs.mkdirTree(path.slice(0, path.lastIndexOf('/')));
    fs.writeFile(path, data);
  }

  let instance, resolveExit;
  const exited = new Promise((r) => { resolveExit = r; });
  const mem = () => instance.exports.mem.buffer;
  const view = () => new DataView(mem());
  const bytes = (p, n) => new Uint8Array(mem(), p, n);
  const encoder = new TextEncoder(), decoder = new TextDecoder();
  const partial = { 1: '', 2: '' };
  const write = (fd, data) => {
    const text = partial[fd] + decoder.decode(data, { stream: true });
    const lines = text.split('\n');
    partial[fd] = lines.pop();
    for (const l of lines) out[fd](l);
  };
  const strings = (list, ptrs, buf) => {
    let p = buf;
    list.forEach((s, i) => {
      const b = encoder.encode(s + '\0');
      view().setUint32(ptrs + 4 * i, p, true);
      bytes(p, b.length).set(b);
      p += b.length;
    });
    return ESUCCESS;
  };
  const sizes = (list, np, sizep) => {
    view().setUint32(np, list.length, true);
    view().setUint32(sizep, list.reduce((n, s) => n + encoder.encode(s).length + 1, 0), true);
    return ESUCCESS;
  };
  const start = performance.now();
  const wasi = {
    args_get: (ptrs, buf) => strings(args, ptrs, buf),
    args_sizes_get: (np, sizep) => sizes(args, np, sizep),
    environ_get: (ptrs, buf) => strings(env, ptrs, buf),
    environ_sizes_get: (np, sizep) => sizes(env, np, sizep),
    clock_res_get: (id, resp) => { view().setBigUint64(resp, 1000n, true); return ESUCCESS; },
    clock_time_get: (id, precision, timep) => {
      const ns = id === 0 ? BigInt(Date.now()) * 1000000n : BigInt(Math.round((performance.now() - start) * 1e6));
      view().setBigUint64(timep, ns, true);
      return ESUCCESS;
    },
    random_get: (p, n) => {
      for (let i = 0; i < n; i += 65536) crypto.getRandomValues(bytes(p + i, Math.min(65536, n - i)));
      return ESUCCESS;
    },
    fd_write: (fd, iovs, n, nwrittenp) => {
      if (fd !== 1 && fd !== 2) return EBADF;
      let total = 0;
      for (let i = 0; i < n; i++) {
        const p = view().getUint32(iovs + 8 * i, true), l = view().getUint32(iovs + 8 * i + 4, true);
        write(fd, bytes(p, l).slice());
        total += l;
      }
      view().setUint32(nwrittenp, total, true);
      return ESUCCESS;
    },
    fd_read: (fd, iovs, n, nreadp) => {
      if (fd !== 0) return EBADF;
      view().setUint32(nreadp, 0, true); // (end of input)
      return ESUCCESS;
    },
    fd_fdstat_get: (fd, statp) => {
      if (fd > 2) return EBADF;
      view().setUint8(statp, 2); // character device
      view().setUint16(statp + 2, 0, true);
      view().setBigUint64(statp + 8, 0xffffffffn, true);
      view().setBigUint64(statp + 16, 0xffffffffn, true);
      return ESUCCESS;
    },
    fd_fdstat_set_flags: () => ESUCCESS,
    fd_prestat_get: () => EBADF, // no directories
    fd_prestat_dir_name: () => EBADF,
    fd_close: (fd) => (fd <= 2 ? ESUCCESS : EBADF),
    sched_yield: () => ESUCCESS,
    proc_exit: (code) => goExit(code),
  };
  // The files: all but the descriptors of the standard input and output.
  if (fs) {
    for (const [name, f] of Object.entries(wasiFS(fs, mem))) {
      const std = wasi[name];
      wasi[name] = std ? (fd, ...a) => (fd <= 2 ? std(fd, ...a) : f(fd, ...a)) : f;
    }
  }
  if (options.wasi) Object.assign(wasi, options.wasi(mem), { proc_exit: wasi.proc_exit });
  const global = options.global ?? globalThis;
  const jsGlobals = () => {
    // GOOS=js: node's fs and process, for the same files.
    const pwd = env.find((e) => e.startsWith('PWD='));
    const none = new Proxy({}, { get: () => () => { throw { errno: ENOSYS }; } });
    const node = nodeFS(fs || none, write, pwd ? pwd.slice(4) : '/');
    // Each run's own, replacing an earlier run's (whose fs writes into
    // that run's output, and keeps a last line without a newline there),
    // but not the engine's (node's process).
    for (const k of ['fs', 'process', 'path']) {
      if (global[k] !== undefined && !global[k][ours]) continue;
      node[k][ours] = true;
      global[k] = node[k];
    }
  };
  // (If the engine cannot list the imports, the program is known to be
  // GOOS=js only once it is instantiated: see moduleImports.)
  const imported = moduleImports(module);
  if (imported && imported.some((i) => i.module === 'gojs')) jsGlobals();
  const imports = new Proxy(wasi, { get: (t, name) => t[name] || (() => ENOSYS) });

  function goExit(code) {
    for (const fd of [1, 2]) if (partial[fd]) { out[fd](partial[fd]); partial[fd] = ''; }
    resolveExit(code);
    // Unwind whatever called it: the program is over.
    throw Object.assign(new Error('exit ' + code), { goExit: code });
  }

  await runWasmGC(module, {
    wasi: { imports, start: (inst) => { instance = inst; if (options.started) options.started(inst); if (!imported && 'callback' in inst.exports) jsGlobals(); } },
    write: (fd, text) => out[fd](text),
    exit: goExit,
    collect: options.collect || null,
    nanotime: () => BigInt(Math.round((performance.now() - start) * 1e6)),
    reserve: options.reserve ?? 64 << 20,
    trace: !!options.trace,
    poison: !!options.poison, // (for testing: see wasm_gc.js)
    maxStacks: options.maxStacks || 0,
    global,
    imports: options.imports,
  });
  const code = await exited;
  if (fs && options.persist) await fs.sync();
  return code;
}
