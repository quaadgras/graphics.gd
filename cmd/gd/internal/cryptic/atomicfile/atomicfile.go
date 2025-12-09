//
// Copyright (c) SAS Institute Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Implement atomic write-rename file pattern. Instead of opening the named
// file it creates a temporary file next to it, then on Commit() renames it. If
// the file is Close()d  before Commit() then it is unlinked instead.
package atomicfile

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// File-like interface used by several functions in this package. Some of them
// may open a file or stdio directly without atomic semantics, in which case
// Commit() is an alias for Close()
type AtomicFile interface {
	io.Reader
	io.ReaderAt
	io.Writer
	io.WriterAt
	io.Seeker
	Truncate(size int64) error

	// Close and unlink the underlying file object, discarding the contents.
	// No-op if Close() or Commit() was already called.
	io.Closer
	// Get the underlying *File object
	GetFile() *os.File
	// Complete the write-rename pattern and close the file
	Commit() error
}

type atomicFile struct {
	*os.File
	name string
}

// Open a temporary file for reading and writing which will ultimately be
// renamed to the given name when Commit() is called.
func New(name string) (AtomicFile, error) {
	tempfile, err := ioutil.TempFile(filepath.Dir(name), filepath.Base(name)+".tmp")
	if err != nil {
		return nil, err
	}
	f := &atomicFile{tempfile, name}
	runtime.SetFinalizer(f, (*atomicFile).Close)
	return f, nil
}

func (f *atomicFile) GetFile() *os.File {
	return f.File
}

func (f *atomicFile) Close() error {
	if f.File == nil {
		return nil
	}
	f.File.Close()
	os.Remove(f.File.Name())
	f.File = nil
	runtime.SetFinalizer(f, nil)
	return nil
}

func (f *atomicFile) Commit() error {
	if f.File == nil {
		return errors.New("file is closed")
	}
	_ = f.File.Chmod(0644)
	if err := f.File.Close(); err != nil {
		return err
	}
	// rename can't overwrite on windows
	if err := os.Remove(f.name); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.Rename(f.File.Name(), f.name); err != nil {
		return err
	}
	f.File = nil
	runtime.SetFinalizer(f, nil)
	return nil
}
