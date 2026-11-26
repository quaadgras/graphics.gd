// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// wasm_gc_js.js is the part of the host that a GOOS=js program has and a
// wasip1 one does not: the imports of module gojs, which package
// syscall/js calls (src/syscall/js/js_gc.go), and the object it knows
// the host by. wasm_gc.js uses it for a module that imports gojs.
//
// A JavaScript value is passed to and from these functions as itself
// (an externref): Go holds it in an object of its own, and the engine
// frees it when Go no longer has that object. A value that Go passes is
// a reference and a number: the reference if it is not null, else the
// number, or null if that is NaN (Go passes NaN itself as a reference,
// and undefined, true and false too: the values predef gives). A value
// that the host returns is returned as a reference if it is a string,
// a symbol, an object or a function, and as null otherwise; its type
// (and, if it is an exception that the call threw, 8 more) and, for a
// number or a boolean, its value, are written to two float64s of Go's,
// whose address predef is given.

// gojs returns the imports of module gojs and the object that syscall/js
// knows as jsGo. rt has:
//   memory:  the module's memory, once it is instantiated;
//   global:  what js.Global() is;
//   call(event): call a function made by js.FuncOf, as event says
//            ({id, this, args}), and return its result;
//   busy(on): the runtime says whether a call must wait (see jsbusy in
//            runtime/wasmgc/jscall.go).
export function gojs(rt) {
  const encoder = new TextEncoder(), decoder = new TextDecoder();
  let retp = 0, retv = null;
  // (A view is made again only when memory has grown.)
  const out = () => (retv !== null && retv.buffer === rt.memory.buffer ? retv : (retv = new Float64Array(rt.memory.buffer, retp, 2)));
  const ret = (x, threw) => {
    const r = out();
    let t, n = 0, ext = null;
    switch (typeof x) {
      case 'undefined': t = 0; break;
      case 'boolean': t = 2; n = x ? 1 : 0; break;
      case 'number': t = 3; n = x; break;
      case 'string': t = 4; ext = x; break;
      case 'symbol': t = 5; ext = x; break;
      case 'function': t = 7; ext = x; break;
      default:
        if (x === null) t = 1;
        else { t = 6; ext = x; }
    }
    r[0] = threw ? t + 8 : t;
    r[1] = n;
    return ext;
  };
  const val = (e, n) => (e !== null ? e : n === n ? n : null);
  // (The address of no bytes may be anything.)
  // (A short ASCII string, as most names of properties are, is made
  // without the decoder, which takes longer than the rest of a call.)
  let u8 = null;
  const str = (p, n) => {
    if (n === 0) return '';
    if (n <= 32) {
      if (u8 === null || u8.buffer !== rt.memory.buffer) u8 = new Uint8Array(rt.memory.buffer);
      let s = '';
      for (let i = p; i < p + n; i++) {
        const c = u8[i];
        if (c >= 0x80) return decoder.decode(u8.subarray(p, p + n));
        s += String.fromCharCode(c);
      }
      return s;
    }
    return decoder.decode(new Uint8Array(rt.memory.buffer, p, n));
  };
  const bytes = (p, n) => (n === 0 ? new Uint8Array(0) : new Uint8Array(rt.memory.buffer, p, n));

  // The arguments of a call that are not numbers, in order (valueArg):
  // a call takes as many from the end as it has NaNs among its numbers.
  const pending = [];
  const args = (p, n) => {
    const a = new Array(n);
    if (n === 0) return a;
    const f = new Float64Array(rt.memory.buffer, p, n);
    let k = 0;
    for (let i = 0; i < n; i++) if (f[i] !== f[i]) k++;
    let j = pending.length - k;
    for (let i = 0; i < n; i++) a[i] = f[i] === f[i] ? f[i] : pending[j++];
    pending.length -= k;
    return a;
  };

  let pendingString = null;

  const go = {
    _pendingEvent: null,
    // syscall/js.FuncOf: a JavaScript function that calls the Go
    // function numbered id.
    _makeFuncWrapper(id) {
      return function (...a) {
        return rt.call({ id, this: this, args: a });
      };
    },
    // (wasm_exec.js keeps the values Go has in _values, which a test of
    // syscall/js counts. Here the engine keeps them: there are none.)
    _values: [],
  };
  const predefs = [undefined, true, false, NaN];

  const imports = {
    'syscall/js.predef': (n, p) => {
      retp = p;
      return n < predefs.length ? predefs[n] : ret(n === 4 ? rt.global : go);
    },
    'syscall/js.stringVal': (p, n) => ret(str(p, n)),
    'syscall/js.valueGet': (v, p, n) => ret(Reflect.get(v, str(p, n))),
    'syscall/js.valueSet': (v, p, n, xe, xn) => { Reflect.set(v, str(p, n), val(xe, xn)); },
    'syscall/js.valueDelete': (v, p, n) => { Reflect.deleteProperty(v, str(p, n)); },
    'syscall/js.valueIndex': (v, i) => ret(Reflect.get(v, i)),
    'syscall/js.valueSetIndex': (v, i, xe, xn) => { Reflect.set(v, i, val(xe, xn)); },
    'syscall/js.valueArg': (x) => { pending.push(x); },
    'syscall/js.valueLength': (v) => v.length,
    'syscall/js.valueCall': (v, mp, mn, ap, an) => {
      const a = args(ap, an);
      try {
        const m = Reflect.get(v, str(mp, mn));
        return ret(Reflect.apply(m, v, a));
      } catch (err) {
        return ret(err, true);
      }
    },
    'syscall/js.valueInvoke': (ve, vn, ap, an) => {
      const a = args(ap, an);
      try {
        return ret(Reflect.apply(val(ve, vn), undefined, a));
      } catch (err) {
        return ret(err, true);
      }
    },
    'syscall/js.valueNew': (ve, vn, ap, an) => {
      const a = args(ap, an);
      try {
        return ret(Reflect.construct(val(ve, vn), a));
      } catch (err) {
        return ret(err, true);
      }
    },
    'syscall/js.valuePrepareString': (e, n) => {
      pendingString = encoder.encode(String(val(e, n)));
      return pendingString.length;
    },
    'syscall/js.valueLoadString': (p, n) => {
      bytes(p, n).set(pendingString.subarray(0, n));
      pendingString = null;
    },
    'syscall/js.valueInstanceOf': (ve, vn, te, tn) => val(ve, vn) instanceof val(te, tn),
    'syscall/js.valueEqual': (v, w) => v === w,
    'syscall/js.copyBytesToGo': (p, n, c, src) => {
      if (!(src instanceof Uint8Array || src instanceof Uint8ClampedArray)) return -1;
      const k = Math.min(n, src.length);
      bytes(p, k).set(src.subarray(0, k));
      return k;
    },
    'syscall/js.copyBytesToJS': (dst, p, n) => {
      if (!(dst instanceof Uint8Array || dst instanceof Uint8ClampedArray)) return -1;
      const k = Math.min(n, dst.length);
      dst.set(bytes(p, k));
      return k;
    },
    'runtime.busy': (on) => rt.busy(on !== 0),
    // (crypto/internal/sysrand, at most 65536 bytes.)
    'runtime.getRandomData': (p, n) => { crypto.getRandomValues(bytes(p, n)); },
  };
  return { imports, go };
}
