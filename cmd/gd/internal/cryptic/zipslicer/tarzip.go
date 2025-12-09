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
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	TarMemberCD  = "zipdir.bin"
	TarMemberZip = "contents.zip"
)

// Make a tar archive with two members:
// - the central directory of the zip file
// - the complete zip file
// This lets us process the zip in one pass, which normally isn't possible with
// the directory at the end.
func ZipToTar(r *os.File, w io.Writer) error {
	size, err := r.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	dirLoc, err := FindDirectory(r, size)
	if err != nil {
		return err
	}
	tw := tar.NewWriter(w)
	if _, err := r.Seek(dirLoc, 0); err != nil {
		return err
	}
	if err := tarAddStream(tw, r, TarMemberCD, size-dirLoc); err != nil {
		return err
	}
	if _, err := r.Seek(0, 0); err != nil {
		return err
	}
	if err := tarAddStream(tw, r, TarMemberZip, size); err != nil {
		return err
	}
	return tw.Close()
}

func tarAddStream(tw *tar.Writer, r io.Reader, name string, size int64) error {
	hdr := &tar.Header{Name: name, Mode: 0644, Size: size}
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	if _, err := io.CopyN(tw, r, size); err != nil {
		return err
	}
	return nil
}

// Read a tar stream produced by ZipToTar and return the zip directory. Files
// must be read from the zip in order or an error will be raised.
func ReadZipTar(r io.Reader) (*Directory, error) {
	tr := tar.NewReader(r)
	hdr, err := tr.Next()
	if err != nil {
		return nil, fmt.Errorf("error reading tar: %w", err)
	} else if hdr.Name != TarMemberCD {
		return nil, errors.New("invalid tarzip")
	}
	zipdir, err := ioutil.ReadAll(tr)
	if err != nil {
		return nil, fmt.Errorf("error reading tar: %w", err)
	}
	hdr, err = tr.Next()
	if err != nil {
		return nil, err
	} else if hdr.Name != TarMemberZip {
		return nil, errors.New("invalid tarzip")
	}
	zr := &zipTarReader{tr: tr}
	return ReadStream(zr, hdr.Size, zipdir)
}

type zipTarReader struct {
	tr *tar.Reader
}

func (z *zipTarReader) Read(d []byte) (int, error) {
	if z.tr == nil {
		return 0, io.EOF
	}
	n, err := z.tr.Read(d)
	if err == io.EOF {
		_, err2 := z.tr.Next()
		if err2 == nil {
			err = errors.New("invalid tarzip")
		} else if err2 != io.EOF {
			err = err2
		}
		z.tr = nil
	}
	return n, err
}
