package builder

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/tooling"
)

// linkStatic has zig link out from args (which must not name the output),
// unless out is already what those args link: a stamp kept beside it records
// what it was linked from, and linking an editor is most of what a build
// with little or nothing changed would take. archive (the Go code, small
// and rewritten by every build) is compared by content, every other file
// among the args (the engine, gigabytes of it) by size and modification
// time. The link is written beside out and renamed over it, so nobody
// ever executes half of one and a process running the old one keeps it.
//
// Callers hold [lockBuilds] on the directory.
func linkStatic(out, archive string, args ...string) error {
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return err
	}
	var from strings.Builder
	fmt.Fprintf(&from, "%s %s\n", zig, strings.Join(args, " "))
	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil || !info.Mode().IsRegular() {
			continue
		}
		if arg == archive {
			sum, err := hashFile(arg)
			if err != nil {
				return err
			}
			fmt.Fprintf(&from, "%s sha256 %s\n", arg, sum)
			continue
		}
		fmt.Fprintf(&from, "%s size %d modified %d\n", arg, info.Size(), info.ModTime().UnixNano())
	}
	// the stamp also names the file it is for, so that an editor put there by
	// anything else (a gd from before there were stamps) is linked again.
	stamped := func() string {
		info, err := os.Stat(out)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%s%s size %d modified %d\n", from.String(), out, info.Size(), info.ModTime().UnixNano())
	}
	stamp := out + ".stamp"
	if last, err := os.ReadFile(stamp); err == nil && string(last) == stamped() {
		return nil
	}
	tmp := out + ".tmp"
	defer os.Remove(tmp)
	if err := tooling.Zig.Exec(append(append([]string{}, args...), "-o", tmp)...); err != nil {
		return err
	}
	// never a stamp beside an editor that it does not describe.
	if err := os.Remove(stamp); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.Rename(tmp, out); err != nil {
		return err
	}
	if err := os.WriteFile(stamp+".tmp", []byte(stamped()), 0o644); err != nil {
		return err
	}
	return os.Rename(stamp+".tmp", stamp)
}

// replaceIfChanged moves built over name unless they hold the same bytes,
// in which case name (and its modification time) is left as it is.
func replaceIfChanged(name, built string) error {
	if was, err := hashFile(name); err == nil {
		now, err := hashFile(built)
		if err != nil {
			return err
		}
		if was == now {
			return os.Remove(built)
		}
	}
	return os.Rename(built, name)
}

func hashFile(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// held keeps the files opened by [holdExecutable] open for the life of gd.
var held []*os.File

// holdExecutable returns a path that goes on running the executable that is
// at path now, whatever is done to that name afterwards: another gd may link
// a new one over it as soon as this one lets go of [lockBuilds], and the
// project has to be run by the build that was made for it. On linux that is
// the file held open, reached through /proc (a child process starts with the
// same descriptors, so the path means the same to it), which costs nothing
// and leaves nothing behind: the old executable goes when the last process
// running or holding it does. Elsewhere it is the path itself.
func holdExecutable(path string) string {
	if runtime.GOOS != "linux" && runtime.GOOS != "android" {
		return path
	}
	file, err := os.Open(path)
	if err != nil {
		return path
	}
	through := fmt.Sprintf("/proc/self/fd/%d", file.Fd())
	if _, err := os.Stat(through); err != nil {
		file.Close()
		return path
	}
	held = append(held, file)
	return through
}
