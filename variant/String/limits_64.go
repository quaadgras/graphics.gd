//go:build amd64 || arm64 || (wasm && !(gd && wasm.gc))

package String

const maxSafeInt = 9007199254740991
