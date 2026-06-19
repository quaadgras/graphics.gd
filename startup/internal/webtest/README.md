# Web (js/wasm) test harness

Runs the `internal` test suite inside a real browser against the Godot **web
export**, headlessly, and reports pass/fail to CI.

## How it works

1. `gd test` (with `GOOS=js`) builds the `go test` binary to `library.wasm`,
   downloads the Godot web export template, and runs the Godot **Web** export
   preset — producing a normal Godot web build (`index.html`, `index.wasm`,
   `index.pck`) plus the Go `library.wasm` and `wasm_exec.js`. It then serves
   the result on `:$PORT` with the cross-origin-isolation headers Godot needs.
2. `run_browser_test.mjs` launches headless Chromium, loads the page over the
   DevTools protocol, and streams the page console. Go's stdout is routed to
   `console.log` by `wasm_exec.js`, so the normal `go test` output (`=== RUN`,
   `--- PASS`/`--- FAIL`, `PASS`/`FAIL`) shows up there. A `PASS`/`ok` line ->
   exit 0; `FAIL`/`--- FAIL`/`exit code: N`/panic/uncaught exception -> exit 1;
   timeout -> exit 2.

`run_web_tests.sh` is the single entrypoint that wires these together (used by
`.github/workflows/web.yml`).

## Files

| file                   | purpose                                                        |
| ---------------------- | -------------------------------------------------------------- |
| `run_web_tests.sh`     | entrypoint: export + serve via `gd`, then drive the browser    |
| `run_browser_test.mjs` | headless-Chromium DevTools driver; exit code reflects the test |
| `serve.mjs`            | static server with COOP/COEP (for serving a prebuilt export)   |
| `Dockerfile`           | glibc node+chromium image (browser-run half only)              |
| `Dockerfile.full`      | glibc Go+node+chromium image (whole flow)                      |

## Requirements

- Go 1.26, Node **22+** (global `WebSocket`), and a Chrome/Chromium binary.
- A **glibc** host. Chromium's software WebGL2 (SwiftShader) is glibc-only, so
  the browser half does not run on a musl host — use the containers below.
- Godot web needs WebGL2; headless Chromium gets it from SwiftShader via
  `--enable-unsafe-swiftshader` (see the flag list in `run_browser_test.mjs`).

## Reproduce the CI locally with podman/docker

Mirror a GitHub `ubuntu-latest` runner and run the whole flow from scratch:

```sh
cd startup/internal/webtest
podman build -t gdwebtestfull -f Dockerfile.full .

# from the repo root:
podman run --rm --userns=keep-id \
  -v "$PWD":/repo \
  -v "$(go env GOMODCACHE)":/gomod -v "$(go env GOCACHE)":/gocache \
  -v gdcache:/cache:U \
  -e HOME=/cache -e REPO=/repo -e GDPATH=/cache/gd -e GOTOOLCHAIN=auto \
  -e GOMODCACHE=/gomod -e GOCACHE=/gocache -e GOFLAGS=-mod=mod \
  -e CHROMIUM=/usr/bin/chromium \
  -w /repo gdwebtestfull \
  bash /repo/startup/internal/webtest/run_web_tests.sh -v
```

Notes:
- `GDPATH`/`GOTOOLCHAIN=auto` are only needed inside the `golang` image, which
  ships `GOTOOLCHAIN=local` and has no real passwd home for `gd`'s downloader.
  A real CI runner needs neither.
- The `gdcache` volume caches the downloaded Godot (~144 MB) between runs.
- To exercise just the browser half against an export you already built on the
  host, use `Dockerfile` + `--network=host`, serving the export with
  `serve.mjs`.
