#!/bin/bash
# build.sh [dir]: make the playground in dir (default dist/): the page and
# its scripts, and in dir/files the large files, which the page fetches by
# the URLs of dir/manifest.json (a hash of each is in its URL):
#
#	go.wasm.gz      compiler.gd's go command for GOOS=js
#	goroot.tar.gz   what it reads of its GOROOT
#	gd.tar.gz       graphics.gd's source and what it imports of other modules (source.sh)
#	cache.tar.gz    a build cache of what the examples import, made by the
#	                same go command under node (prebuild.mjs), as the page's makes it
#	engine.wasm.gz  the engine: graphics.gd's web template
#
# Needs: a checkout of graphics.gd's release branch (GD), compiler.gd built
# (GDROOT, its src/make.bash run), node 24, and gd with the stock Go
# toolchain in PATH, which exports the engine and the empty project it
# opens (or ENGINE=dir, a web export made earlier).
# CACHEDIR= a directory keeps the build cache between runs (JOBS= how many
# node processes fill it at once, 3 GB each).
# FILES=https://host/path is where the large files will be served from, if
# not beside the page.
set -e
here=$(cd "$(dirname "$0")" && pwd)
gd=$(cd "${GD:?set GD to a checkout of graphics.gd (its release branch)}" && pwd)
root=${GDROOT:?set GDROOT to a built checkout of compiler.gd}
site=${1:-$here/dist}
# (Built beside the site and put in its place when whole: a site that is
# being served is not left half made for the minutes the cache takes.)
out=$site.new
files=$out/files
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
fork() { GOROOT=$root GOTOOLCHAIN=local GOFLAGS= PATH=$root/bin:$PATH "$@"; }

rm -rf "$out"
mkdir -p "$files"
cp "$here"/{index.html,run.html,worker.mjs,fetch.mjs} "$out/"
cp -r "$here/examples" "$out/examples"
cp "$root"/lib/wasm/{wasm_gc.js,wasm_gc_js.js,wasm_gc_fs.js,wasm_gc_browser.js} "$root/doc/gd/site/playground/files.mjs" "$out/"

echo "engine" >&2
engine=$ENGINE
if [ -z "$engine" ]; then
	# An export of an empty project: gd downloads the web template that
	# matches it, and writes the project's pack.
	mkdir "$tmp/project"
	printf 'module project\n\ngo 1.27\n\nrequire graphics.gd v0.0.0\n\nreplace graphics.gd => %s\n' "$gd" > "$tmp/project/go.mod"
	printf 'package main\n\nimport "graphics.gd/startup"\n\nfunc main() {\n\tstartup.LoadingScene()\n\tstartup.Scene()\n}\n' > "$tmp/project/main.go"
	(cd "$tmp/project" && GOFLAGS=-mod=mod go mod tidy >&2 && GOFLAGS=-mod=mod GOOS=js gd build >&2)
	engine=$tmp/project/releases/js/wasm
fi
cp "$engine/index.js" "$out/godot.js"
cp "$engine/index.audio.worklet.js" "$out/godot.audio.worklet.js"
cp "$engine/index.audio.position.worklet.js" "$out/godot.audio.position.worklet.js"
cp "$engine/index.pck" "$out/play.pck"
gzip -9n -c "$engine/index.wasm" > "$files/engine.wasm.gz"
NODE_OPTIONS= node "$here/patch.mjs" "$out"

echo "go.wasm" >&2
(cd "$root/src" && fork env GOOS=js GOARCH=wasm go build -ldflags='-s -w' -o "$tmp/go.wasm" ./cmd/go)
gzip -9n -c "$tmp/go.wasm" > "$files/go.wasm.gz"

echo "goroot.tar.gz" >&2
(cd "$root" && {
	echo go.env
	[ -f VERSION ] && echo VERSION
	find pkg/include -type f
	find src -path src/cmd -prune -o -type d -name testdata -prune -o -type f \
		\( -name '*.go' -o -name '*.s' -o -name '*.h' -o -name go.mod -o -name go.sum -o -name modules.txt \) \
		! -name '*_test.go' -print
} | tar --format=ustar --owner=0 --group=0 --mtime=@0 -cf - -T - | gzip -9n) > "$files/goroot.tar.gz"

echo "gd.tar.gz" >&2
fork "$here/source.sh" "$gd" "$files/gd.tar.gz"

echo "cache.tar.gz (prebuild: minutes)" >&2
# What the examples import, of graphics.gd and of std, compiled by the go
# command above under node, into CACHEDIR if set (what is there already is
# not compiled again: the names are hashes of what went in; the go command
# and its environment are in all of them, so each has a directory of its own
# there; and so has each state of graphics.gd's source, or what was compiled
# of the one before would stay in the cache that the page downloads).
cachedir=${CACHEDIR:-$tmp}/$(cat "$tmp/go.wasm" "$out/files.mjs" "$files/gd.tar.gz" | sha256sum | cut -c1-12)
# (The examples are listed as packages of a module of their own, which
# has graphics.gd from the checkout.)
mkdir "$tmp/examples"
cp -r "$here/examples/." "$tmp/examples/"
tar -xzOf "$files/gd.tar.gz" play.mod | sed "s|=> /gd/|=> $tmp/src/|" > "$tmp/examples/go.mod"
mkdir "$tmp/src" && tar -xzf "$files/gd.tar.gz" -C "$tmp/src"
(cd "$tmp/examples" && fork env GOOS=js GOARCH=wasm GOFLAGS=-mod=mod GOPROXY=off go list -deps -f '{{.ImportPath}} {{join .Imports " "}}' ./...) | grep -v '^play/' > "$tmp/deps"
cp "$here/prebuild.mjs" "$out/prebuild.mjs"
NODE_OPTIONS= node --stack-size=4000 --max-old-space-size=8000 "$out/prebuild.mjs" "$tmp/go.wasm" "$files/goroot.tar.gz" "$files/gd.tar.gz" -cachedir "$cachedir" -deps "$tmp/deps" ${JOBS:+-jobs $JOBS}
rm "$out/prebuild.mjs"
(cd "$cachedir" && find . -type f ! -name '*.tmp*' | sed 's|^\./||' | LC_ALL=C sort | tar --format=ustar --owner=0 --group=0 -cf - -T - | gzip -6n) > "$files/cache.tar.gz"

# The manifest: where each large file is, with a hash of it.
base=${FILES:-files}
url() { echo "$base/$1?v=$(sha256sum "$files/$1" | cut -c1-12)"; }
cat > "$out/manifest.json" <<JSON
{"go": "$(url go.wasm.gz)", "goroot": "$(url goroot.tar.gz)", "gd": "$(url gd.tar.gz)", "cache": "$(url cache.tar.gz)", "engine": "$(url engine.wasm.gz)", "pck": "play.pck", "version": "$(cd "$gd" && git rev-parse --short HEAD) / $(cd "$root" && git rev-parse --short HEAD)"}
JSON
# The engine has threads: the page must be cross-origin isolated. (These
# are the headers as Cloudflare Pages reads them; serve.mjs sends the same.)
cat > "$out/_headers" <<HEADERS
/*
  Cross-Origin-Opener-Policy: same-origin
  Cross-Origin-Embedder-Policy: require-corp
HEADERS
rm -rf "$site.old"
[ -e "$site" ] && mv "$site" "$site.old"
mv "$out" "$site"
rm -rf "$site.old"
ls -l "$site" "$site/files" >&2
