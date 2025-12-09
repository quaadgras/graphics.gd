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

package signjar

import (
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"graphics.gd/cmd/gd/internal/cryptic/x509tools"
	"graphics.gd/cmd/gd/internal/cryptic/zipslicer"
)

// found in the "extra" field of JAR files, not strictly required but it makes
// `file` output actually say JAR
var jarMagic = []byte{0xfe, 0xca, 0, 0}

type JarDigest struct {
	Digests  map[string]string
	Manifest []byte
	Hash     crypto.Hash
	inz      *zipslicer.Directory
}

func DigestJarStream(r io.Reader, hash crypto.Hash) (*JarDigest, error) {
	inz, err := zipslicer.ReadZipTar(r)
	if err != nil {
		return nil, err
	}
	return updateManifest(inz, hash)
}

// Digest all of the files in the JAR
func digestFiles(jar *zipslicer.Directory, hash crypto.Hash) (*JarDigest, error) {
	jd := &JarDigest{
		Hash:    hash,
		Digests: make(map[string]string),
		inz:     jar,
	}
	for _, f := range jar.File {
		if f.Name == manifestName {
			r, err := f.Open()
			if err != nil {
				return nil, fmt.Errorf("failed to read JAR manifest: %w", err)
			}
			jd.Manifest, err = io.ReadAll(r)
			if err != nil {
				return nil, fmt.Errorf("failed to read JAR manifest: %w", err)
			}
		} else if strings.HasSuffix(f.Name, "/") || !keepFile(f.Name) {
			// not hashing
		} else {
			r, err := f.Open()
			if err != nil {
				return nil, err
			}
			d := hash.New()
			if _, err := io.Copy(d, r); err != nil {
				return nil, fmt.Errorf("failed to digest JAR file %s: %w", f.Name, err)
			}
			if err := r.Close(); err != nil {
				return nil, fmt.Errorf("failed to digest JAR file %s: %w", f.Name, err)
			}
			jd.Digests[f.Name] = base64.StdEncoding.EncodeToString(d.Sum(nil))
		}
		// Ensure we get a copy of the zip metadata even if the file isn't
		// digested, because if we're reading from a stream we can't go back
		// and get it later.
		if _, err := f.GetDataDescriptor(); err != nil {
			return nil, fmt.Errorf("failed to read JAR manifest: %w", err)
		}
	}
	return jd, nil
}

// Check JAR contents against its manifest and adds digests if necessary
func updateManifest(jar *zipslicer.Directory, hash crypto.Hash) (*JarDigest, error) {
	jd, err := digestFiles(jar, hash)
	if err != nil {
		return nil, err
	}
	files, malformed, err := parseManifest(jd.Manifest)
	if err != nil {
		return nil, err
	}

	hashName := x509tools.HashNames[hash]
	if hashName == "" {
		return nil, errors.New("unsupported hash type")
	}
	hashName += "-Digest"
	changed := false
	for name, calculated := range jd.Digests {
		// if the manifest has a matching digest, check it. otherwise add to the manifest.
		attrs := files.Files[name]
		if attrs == nil {
			// file is not mentioned in the manifest at all
			files.Files[name] = http.Header{
				"Name":   []string{name},
				hashName: []string{calculated},
			}
			files.Order = append(files.Order, name)
			changed = true
		} else if attrs.Get("Magic") != "" {
			// magic means a special digester is required. hopefully it's already been digested.
		} else if existing := attrs.Get(hashName); existing != "" {
			// manifest has a digest already, check it
			if existing != calculated {
				return nil, fmt.Errorf("%s mismatch for JAR file %s: manifest %s != calculated %s", hashName, name, existing, calculated)
			}
		} else {
			// file in manifest but no matching digest
			attrs.Set(hashName, calculated)
			changed = true
		}
	}
	// Add empty digests for any dir entries with metadata
	for name, attrs := range files.Files {
		if strings.HasSuffix(name, "/") && attrs.Get(hashName) == "" {
			d := hash.New()
			calculated := base64.StdEncoding.EncodeToString(d.Sum(nil))
			attrs.Set(hashName, calculated)
			changed = true
		}
	}
	if files.Main == nil {
		files.Main = make(http.Header)
	}
	// Ensure default headers are present
	if files.Main.Get("Manifest-Version") == "" {
		files.Main.Set("Manifest-Version", "1.0")
		changed = true
	}
	if files.Main.Get("Created-By") == "" {
		files.Main.Set("Created-By", "Signflinger")
		changed = true
	}
	if files.Main.Get("Built-By") == "" {
		files.Main.Set("Built-By", "Signflinger")
		changed = true
	}
	if changed || malformed {
		jd.Manifest = files.Dump()
	}
	return jd, nil
}
