//go:build amd64 || arm64

// Package callerpc provides fast caller PC capture via assembly.
package callerpc

// Callerpc returns the PC of the caller's caller by walking the
// frame pointer chain. This is used to record which Go call site
// enqueued a ring buffer entry, for crash diagnostics.
func Callerpc() uintptr
