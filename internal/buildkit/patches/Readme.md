# ios-internal-linking.patch

Toolchain patches (against go1.26.5) that let the wasm-hosted toolchain
produce a **runnable** ios/arm64 executable with internal linking — no
external linker, no C toolchain, no Mac in the build path.

**Verified on hardware 2026-08-18**: a hello-world compiled by
`compile.wasm` and linked by `link.wasm` (both running under wazero,
zero process spawns), dev-signed and installed on an iPhone 8
(iOS 16.7), ran to completion and wrote its proof file:

    hello from an iPhone running a binary compiled and linked by a Go toolchain inside wazero
    built with zero process spawns for ios/arm64

Three changes:

1. `internal/platform/supported.go` — stop forcing external linking
   for ios/arm64, and add ios/arm64 to `InternalLinkPIESupported` (iOS
   defaults to PIE). Stock Go's refusal is a policy gate, not a
   technical wall.
2. `cmd/link/internal/ld/macho.go` — emit `LC_BUILD_VERSION`
   (platform iOS, minos 15.0) for internally linked iOS binaries; the
   stock code only emits it for macOS, and iOS refuses binaries with no
   declared platform.
3. `runtime/rt0_ios_arm64.s` — give `_rt0_arm64_ios` a real body
   (`JMP runtime·rt0_go`, same as darwin: dyld passes argc/argv in
   R0/R1). Stock Go deliberately plants `UNDEF` there ("ios/arm64 only
   supports external linking"), which is exactly the SIGILL you'll see
   if this patch is missing.

## Recipe

```sh
cp -a $(go env GOROOT) goroot-ios && (cd goroot-ios && patch -p1 < ios-internal-linking.patch)
export GOROOT=$PWD/goroot-ios GOTOOLCHAIN=local
go build -o patched-go cmd/go                                # its platform checks are baked in
GOOS=wasip1 GOARCH=wasm go build -o compile.wasm cmd/compile
GOOS=wasip1 GOARCH=wasm go build -o link.wasm cmd/link
# export data for the target (this builds std for ios with the patched rules):
GOOS=ios GOARCH=arm64 CGO_ENABLED=0 ./patched-go list -deps -export \
    -f '{{if .Export}}packagefile {{.ImportPath}}={{.Export}}{{end}}' . > importcfg
# then run compile.wasm / link.wasm via buildkit with guest env
# GOOS=ios GOARCH=arm64, link flags: -buildmode=pie -linkmode=internal
```

Sign the result into an .app with a development certificate and a
wildcard team profile; iOS installs and runs it. The linker's own
ad-hoc signature is replaced by codesign, so no extra tooling is
needed beyond what SideStore or a dev cert already provides.

Scope: pure-Go binaries. Linking against C archives (libgodot) still
needs a real Mach-O external linker — that's the lld tier.
