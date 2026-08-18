# ld64.lld as WebAssembly

Build `ld64.lld` (LLVM's Mach-O linker) as a `wasm32-wasip1` module so
it runs under wazero, in-process, with **zero process spawns** — the
linker tier of the on-device iOS toolchain. A fresh module instance per
link gives the same re-entrancy and crash isolation that the wasm Go
toolchain gets (and that lld's own globals otherwise fight, see
`safeLldMain`).

**Verified on the real workload 2026-08-18.** lld.wasm under wazero
linked the full graphics.gd iOS game — MoltenVK + libgodot + libgo +
the faux-SDK `.tbd` frameworks, gd's exact `ld64` argument list — and
produced a **124,744,648-byte arm64 Mach-O, byte-identical to the
native `ld64.lld` output except the 16-byte `LC_UUID`** (and the ad-hoc
code-signature page hashes that re-hash from it; codesign replaces that
anyway). Load commands verified on macOS: `LC_BUILD_VERSION` platform
iOS / minos 15.0, `LC_DYLD_CHAINED_FIXUPS` present. ~14–17s per link on
the desktop (wazero compiling runtime); the iOS interpreter is slower.

## Why the patches

LLVM/lld assume a POSIX host. wasi (as wazero implements it) has no
threads, no signals, no processes, no sockets, and no passwd/rlimit/
umask APIs — but the linker needs **none** of those at runtime on this
path, so every patch replaces an unreachable OS dependency with a
trivial equivalent. Two buckets:

- `llvm-wasi.patch` — `#ifdef __wasi__` guards across LLVM Support's
  Unix `.inc` files (Signals, Program, Process, Path, Watchdog,
  ProgramStack), a `__wasi__` branch in `ADT/bit.h`, a socket-stream
  compile-out, an rpath-link cmake guard, and lld's `unlinkAsync` made
  synchronous. All behaviour-preserving on a single-instance,
  single-threaded guest.
- `shim-include/` — single-threaded C++ standard headers wasi's
  no-threads libc++ omits: `<mutex>`, `<shared_mutex>`,
  `<condition_variable>`, `<thread>`, `<future>` (async runs inline),
  and a `<setjmp.h>` whose `longjmp` traps to the host (crash recovery
  is dead code when the host re-instantiates per link).
- `wasi-compat.c` — provides `__cxa_thread_atexit` and `posix_madvise`,
  which wasi-libc omits; both no-op safely here.

## Build

```sh
export LLDWASM=$PWD/lldwasm
mkdir -p $LLDWASM && cd $LLDWASM
# wasi-sdk 25 (sysroot + wasm builtins; its own clang is glibc-linked
# so we drive the host clang 21 against its sysroot):
curl -L https://github.com/WebAssembly/wasi-sdk/releases/download/wasi-sdk-25/wasi-sdk-25.0-x86_64-linux.tar.gz | tar xz
git clone --depth 1 --branch release/21.x https://github.com/llvm/llvm-project
( cd llvm-project && git apply < .../llvm-wasi.patch )
# shim resource dir maps clang 21 headers + the sdk's wasm builtins:
mkdir -p resource/lib/wasm32-unknown-wasip1
ln -s $(clang -print-resource-dir)/include resource/include
ln -s wasi-sdk-25.0-x86_64-linux/lib/clang/19/lib/wasip1/libclang_rt.builtins-wasm32.a \
      resource/lib/wasm32-unknown-wasip1/libclang_rt.builtins.a
cp -r .../shim-include .../wasi-toolchain.cmake .../wasi-compat.c .../build.sh .
clang --target=wasm32-wasip1 --sysroot=wasi-sdk-*/share/wasi-sysroot -resource-dir=resource -Os -c wasi-compat.c -o wasi-compat.o
sh build.sh        # -> llvm-project/build-wasm/bin/lld  (the wasm module)
```

Run it with any wasip1 host (wazero, wasmtime) as `ld64.lld`:
`wasmrun build-wasm/bin/lld ld64.lld -arch arm64 ...`.

## Scope / next

This is the heavy link (against libgodot's C++ archive) that internal
linking can't do — it completes the buildkit story: Go compile+link and
now the Mach-O link all run as wasm under wazero. Remaining to wire it
into `buildkit.Runner`: mount the project + SDK dirs, feed gd's
`ld64` args, and hand the output to codesign/SideStore. clang.wasm (for
the future cgo tier) builds from the same tree with the same shims —
`-DLLVM_ENABLE_PROJECTS='lld;clang'`.
