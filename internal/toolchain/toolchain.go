// Package toolchain abstracts how the harness compiles, runs and ships
// the project it is working on.
//
// On hosts with process spawning (desktop, Termux) the exec tier shells
// out to the gd command, which already knows how to build, export,
// install and serve every target — including GOOS=ios with a SideStore
// install URL. On iOS itself there is no exec; the ios tier will run
// the toolchain in-process (compiler.gd + embedded lld) and is stubbed
// until then.
package toolchain

import "io"

// Toolchain builds and runs the project rooted at Project, streaming
// human-readable progress to Out.
type Toolchain struct {
	Project string
	Out     io.Writer
}

func New(project string, out io.Writer) *Toolchain {
	return &Toolchain{Project: project, Out: out}
}
