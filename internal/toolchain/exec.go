//go:build !ios

package toolchain

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Shell runs a shell command in the project directory and returns its
// combined output. Output is capped; long-running commands time out.
func (t *Toolchain) Shell(command string, timeout time.Duration) (string, error) {
	if timeout <= 0 {
		timeout = 5 * time.Minute
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	cmd.Dir = t.Project
	out, err := cmd.CombinedOutput()
	const cap = 100 << 10
	if len(out) > cap {
		out = append([]byte("(output truncated)\n"), out[len(out)-cap:]...)
	}
	if ctx.Err() == context.DeadlineExceeded {
		err = fmt.Errorf("timed out after %v", timeout)
	}
	return string(out), err
}

// GD runs a gd verb (build, run, test) against the project, streaming
// output to t.Out. An empty goos builds for the host's default target.
func (t *Toolchain) GD(verb string, goos string) error {
	gd, err := exec.LookPath("gd")
	if err != nil {
		return fmt.Errorf("the gd command is not on PATH: go install graphics.gd/cmd/gd@release")
	}
	cmd := exec.Command(gd, verb)
	cmd.Dir = t.Project
	cmd.Env = os.Environ()
	if goos != "" {
		cmd.Env = append(cmd.Env, "GOOS="+goos)
	}
	cmd.Stdout = t.Out
	cmd.Stderr = t.Out
	fmt.Fprintf(t.Out, "$ %s\n", strings.TrimSpace(strings.TrimPrefix("GOOS="+goos+" gd "+verb, "GOOS= ")))
	return cmd.Run()
}
