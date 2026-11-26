#!/bin/bash
# gdsrc.sh graphics.gd-dir out.tar.gz: the source that the playground's go
# command builds graphics.gd programs from, as it is laid out in the page
# under /gd: graphics.gd and the packages of the other modules that its
# web build imports, each replaced by its directory in the program's go.mod.
set -e
gd=$(cd "$1" && pwd); out=$2
tmp=$(mktemp -d); trap 'rm -rf "$tmp"' EXIT
export GOOS=js GOARCH=wasm GOFLAGS=-mod=mod
mkdir -p "$tmp/graphics.gd"
(cd "$gd" && find . -path ./cmd -prune -o -path ./.git -prune -o -path ./.claude -prune -o -path ./releases -prune -o -path ./gd -prune -o -path ./playground -prune \
	-o -type d -name testdata -prune -o -type f \( -name '*.go' -o -name '*.s' -o -name '*.S' -o -name '*.h' -o -name '*.c' -o -name '*.cc' -o -name '*.cpp' -o -name '*.m' -o -name '*.mm' \) ! -name '*_test.go' -print \
	| tar -cf - -T - | tar -xf - -C "$tmp/graphics.gd")
# The other modules: only the packages that are imported.
(cd "$gd" && go list -deps -f '{{if .Module}}{{if ne .Module.Path "graphics.gd"}}{{.Module.Path}} {{.Module.Version}} {{.Module.Dir}} {{.Dir}}{{end}}{{end}}' \
	./classdb/... ./variant/... ./startup ./shaders/... ./nativeui/...) | sort -u > "$tmp/deps"
: > "$tmp/requires"
while read -r mod version moddir dir; do
	rel=${dir#"$moddir"}
	mkdir -p "$tmp/$mod$rel"
	find "$dir" -maxdepth 1 -type f \( -name '*.go' -o -name '*.s' -o -name '*.h' \) ! -name '*_test.go' -exec cp {} "$tmp/$mod$rel/" \;
	install -m 644 "$moddir/go.mod" "$tmp/$mod/go.mod"
	echo "$mod $version" >> "$tmp/requires"
done < "$tmp/deps"
sort -u "$tmp/requires" > "$tmp/requires.u"
# graphics.gd's go.mod, with only what its web build needs.
{
	echo "module graphics.gd"; echo; grep '^go ' "$gd/go.mod"; echo
	while read -r mod version; do echo "require $mod $version"; done < "$tmp/requires.u"
} > "$tmp/graphics.gd/go.mod"
# The go.mod of a program (the page writes it to /work/go.mod).
{
	echo "module play"; echo; grep '^go ' "$gd/go.mod"; echo
	echo "require graphics.gd v0.0.0"
	while read -r mod version; do echo "require $mod $version // indirect"; done < "$tmp/requires.u"
	echo; echo "replace graphics.gd => /gd/graphics.gd"
	while read -r mod version; do echo "replace $mod => /gd/$mod"; done < "$tmp/requires.u"
} > "$tmp/play.mod"
chmod -R u+w "$tmp"
rm "$tmp/deps" "$tmp/requires" "$tmp/requires.u"
(cd "$tmp" && find . -type f | sed 's|^\./||' | LC_ALL=C sort | tar --format=ustar --owner=0 --group=0 --mtime=@0 -cf - -T - | gzip -9n > "$out")
