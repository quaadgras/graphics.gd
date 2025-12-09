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

package atomicfile

import (
	"io"
	"os"
)

type nopAtomic struct {
	*os.File
	doClose bool
}

func (a nopAtomic) GetFile() *os.File {
	return a.File
}

func (a nopAtomic) Commit() error {
	if a.doClose {
		return a.Close()
	}
	return nil
}

func isSpecial(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		if !stat.Mode().IsRegular() {
			return true
		}
	}
	return false
}

// Pick the best strategy for writing to the given path. Pipes and devices will
// be written to directly, otherwise write-rename.
func WriteAny(path string) (AtomicFile, error) {
	if path == "-" {
		return nopAtomic{os.Stdout, false}, nil
	}
	if isSpecial(path) {
		f, err := os.Create(path)
		return nopAtomic{f, true}, err
	}
	return New(path)
}

// If src and dest are the same, use src for reading and writing. If they are
// different, make a copy and open the destination as an atomicfile, after
// which src will be closed.
func WriteInPlace(src *os.File, dest string) (AtomicFile, error) {
	if src.Name() == dest {
		return nopAtomic{src, false}, nil
	}
	outfile, err := New(dest)
	if err != nil {
		return nil, err
	}
	if _, err := src.Seek(0, 0); err != nil {
		return nil, err
	}
	if _, err := io.Copy(outfile, src); err != nil {
		return nil, err
	}
	if _, err := outfile.Seek(0, 0); err != nil {
		return nil, err
	}
	err = src.Close()
	return outfile, err
}

// Write bytes to a file, using write-rename when appropriate
func WriteFile(path string, data []byte) error {
	f, err := WriteAny(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return err
	}
	return f.Commit()
}
