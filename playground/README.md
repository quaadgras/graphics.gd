# graphics.gd playground

A page where a graphics.gd program is edited, compiled in the browser and
run there on the engine. It is built on the playground of
[compiler.gd](https://compiler.gd/playground): the `go` command of that
fork of Go, built for `GOOS=js`, runs in a worker and builds the program;
the program then runs in the page beside graphics.gd's web build of the
engine.

## The pieces

- `index.html`: the page (editor, the frame the program runs in, output).
- `worker.mjs`: the go command. Its files are in memory: `GOROOT` at
  `/goroot`, graphics.gd's source at `/gd`, the program at `/work`, `$HOME`
  (with the build cache) at `/home`, kept in IndexedDB between visits.
- `run.html`: one run. The page makes a new frame for every run and sends
  it the engine's module (compiled once) and the program.
- `fetch.mjs`: the large files, kept in Cache Storage.
- `examples/`: the programs in the page's menu. What they import is what
  the build cache has precompiled.
- `source.sh`: graphics.gd's source as the page lays it out, with the
  `go.mod` of a program (`/gd/play.mod`): graphics.gd and the two modules
  its web build imports are `replace`d by directories, so no module is
  downloaded (`GOPROXY=off`).
- `prebuild.mjs`: the same go command under node, with the same paths, to
  make the build cache (the paths are part of the cache's keys) or to build
  a program as the page would.
- `patch.mjs`: the changes to compiler.gd's loader and to the engine's that
  let the two meet (see its comments).
- `build.sh`: all of it, into a directory; `serve.mjs` serves that on localhost.

## Building and trying it

	GD=/path/to/graphics.gd GDROOT=/path/to/compiler.gd playground/build.sh /tmp/playground
	node playground/serve.mjs /tmp/playground

compiler.gd must be built (`src/make.bash`), and `gd` with the stock Go
toolchain must be in `PATH`: it exports the empty project that the engine
opens and fetches the web template. The slow step is the build cache
(every package is compiled by a WebAssembly compiler under node, several
processes at once): `CACHEDIR=` a directory keeps it between builds, and
only what changed is compiled again.

To build one program as the page would, without a browser:

	node --stack-size=4000 prebuild.mjs go.wasm goroot.tar.gz gd.tar.gz -cachedir cache -main main.go -o main.wasm

(from a directory that has the loader's scripts, as `dist/` has).

## What it needs of the target

The fork's `GOOS=js` is WebAssembly GC with 4-byte `int`, `uintptr` and
pointers, no assembly, and pointers that cannot be made of integers.
graphics.gd's files for it have the build constraint `gd && wasm.gc`; the ones
they replace are for Go's own linear-memory port, which `gd` still builds
the web with.

The engine is graphics.gd's ordinary web template. It has threads, so the
page must be cross-origin isolated (`_headers`, for Cloudflare Pages), and
whatever serves the large files from another origin must allow that
(CORS, or `Cross-Origin-Resource-Policy: cross-origin`).

## Publishing

The site is this branch: Cloudflare Pages builds it on a push, and serves
`public/playground` at graphics.gd/playground (`public/_headers` isolates
it). What `build.sh` makes goes to two places:

- `dir/files`, the large files (more than Pages takes: 25 MiB a file), to
  the R2 bucket `graphics-gd`, which is release.graphics.gd, under
  `playground/`. Build with `FILES=https://release.graphics.gd/playground`
  so that the manifest names them there. The bucket's CORS rule lets
  https://graphics.gd fetch them, and no other origin.

	for f in /tmp/playground/files/*; do npx wrangler@4 r2 object put graphics-gd/playground/$(basename $f) --file $f --content-type application/gzip --remote; done

- everything else in `dir`, to `public/playground` (not `files`, `_headers`):

	rsync -a --delete --exclude files --exclude _headers /tmp/playground/ public/playground/

Then commit and push. Upload the files first: the manifest that the page
reads names each by a hash of it.
