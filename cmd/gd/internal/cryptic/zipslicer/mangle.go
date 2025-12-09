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

package zipslicer

import (
	"bytes"
	"time"

	"graphics.gd/cmd/gd/internal/cryptic/binpatch"
)

type Mangler struct {
	outz          *Directory
	patch         *binpatch.PatchSet
	newcontents   bytes.Buffer
	indir, insize int64
}

type MangleFunc func(*MangleFile) error

// Walk all the files in the directory in-order, invoking a callback that can
// decide whether to keep or discard each one. Returns a Mangler that can be
// used to add more files and eventually produce a binary patch against the
// original zip.
func (d *Directory) Mangle(callback MangleFunc) (*Mangler, error) {
	m := &Mangler{
		outz:   new(Directory),
		patch:  binpatch.New(),
		indir:  d.DirLoc,
		insize: d.Size,
	}
	for _, f := range d.File {
		mf := &MangleFile{File: *f, m: m}
		if err := callback(mf); err != nil {
			return nil, err
		}
		if mf.deleted {
			size, err := mf.GetTotalSize()
			if err != nil {
				return nil, err
			}
			m.patch.Add(int64(mf.Offset), size, nil)
		} else {
			if _, err := m.outz.AddFile(&mf.File); err != nil {
				return nil, err
			}
		}
	}
	return m, nil
}

// Add a new file to a zip mangler
func (m *Mangler) NewFile(name string, contents []byte) error {
	deflate := len(contents) != 0
	_, err := m.outz.NewFile(name, nil, contents, &m.newcontents, time.Now(), deflate, true)
	return err
}

// Create a binary patchset out of the operations performed in this mangler
func (m *Mangler) MakePatch(forceZip64 bool) (*binpatch.PatchSet, error) {
	w := &m.newcontents
	if err := m.outz.WriteDirectory(w, w, forceZip64); err != nil {
		return nil, err
	}
	m.patch.Add(m.indir, m.insize-m.indir, m.newcontents.Bytes())
	return m.patch, nil
}

type MangleFile struct {
	File
	m       *Mangler
	deleted bool
}

// Mark this file for deletion
func (f *MangleFile) Delete() {
	f.deleted = true
}
