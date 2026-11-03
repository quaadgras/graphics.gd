//go:build !unix

package builder

// lockBuilds does nothing here, see the unix one.
func lockBuilds(dir string) (unlock func(), err error) {
	return func() {}, nil
}
