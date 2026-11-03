//go:build gdstub

package startup

// gd runs a Godot project with no Go code of its own with a Go program that
// is this package and nothing else, built with the gdstub tag: there are no
// Go packages to generate for such a project (which is not a Go module) and
// nothing for gd to build when the editor plays it.
const isStub = true
