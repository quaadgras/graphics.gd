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
	"archive/zip"
	"bytes"
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"io"
	"net/http"
	"path"
	"strings"

	"graphics.gd/cmd/gd/internal/cryptic/pkcs7"
	"graphics.gd/cmd/gd/internal/cryptic/pkcs9"
	"graphics.gd/cmd/gd/internal/cryptic/x509tools"
)

var errNoDigests = errors.New("no recognized digests found")

type JarSignature struct {
	pkcs9.TimestampedSignature
	SignatureHeader http.Header
	Hash            crypto.Hash
}

func Verify(inz *zip.Reader, skipDigests bool) ([]*JarSignature, error) {
	var manifest []byte
	sigfiles := make(map[string][]byte)
	sigblobs := make(map[string][]byte)
	for _, f := range inz.File {
		dir, name := path.Split(strings.ToUpper(f.Name))
		if dir != "META-INF/" || name == "" {
			continue
		}
		i := strings.LastIndex(name, ".")
		if i < 0 {
			continue
		}
		base, ext := name[:i], name[i:]
		r2, err := f.Open()
		if err != nil {
			return nil, err
		}
		contents, err := io.ReadAll(r2)
		if err != nil {
			return nil, err
		}
		if err := r2.Close(); err != nil {
			return nil, err
		}
		if name == "MANIFEST.MF" {
			manifest = contents
		} else if ext == ".SF" {
			sigfiles[base] = contents
		} else if ext == ".RSA" || ext == ".DSA" || ext == ".EC" || strings.HasPrefix(name, "SIG-") {
			sigblobs[base] = contents
		}
	}
	if manifest == nil {
		return nil, errors.New("JAR contains no META-INF/MANIFEST.MF")
	} else if len(sigfiles) == 0 {
		return nil, errors.New("JAR is not signed (no META-INF/*.SF files found)")
	}
	sigs := make([]*JarSignature, 0, len(sigfiles))
	for base, sigfile := range sigfiles {
		pkcs := sigblobs[base]
		if pkcs == nil {
			return nil, fmt.Errorf("JAR contains sigfile META-INF/%s.SF with no matching signature", base)
		}
		psd, err := pkcs7.Unmarshal(pkcs)
		if err != nil {
			return nil, err
		}
		sig, err := psd.Content.Verify(sigfile, false)
		if err != nil {
			return nil, err
		}
		ts, err := pkcs9.VerifyOptionalTimestamp(sig)
		if err != nil {
			return nil, err
		}
		ts.Raw = pkcs
		hdr, err := verifySigFile(sigfile, manifest)
		if err != nil {
			return nil, err
		}
		hash, _ := x509tools.PkixDigestToHash(ts.SignerInfo.DigestAlgorithm)
		sigs = append(sigs, &JarSignature{
			TimestampedSignature: ts,
			Hash:                 hash,
			SignatureHeader:      hdr,
		})
	}
	if !skipDigests {
		if err := verifyManifest(inz, manifest); err != nil {
			return nil, err
		}
	}
	return sigs, nil
}

// Verify all digests in MANIFEST.MF
func verifyManifest(inz *zip.Reader, manifest []byte) error {
	parsed, err := ParseManifest(manifest)
	if err != nil {
		return err
	}
	zipfiles := make(map[string]*zip.File, len(inz.File))
	for _, fh := range inz.File {
		zipfiles[fh.Name] = fh
	}
	for filename, keys := range parsed.Files {
		if keys.Get("Magic") != "" {
			continue
		}
		fh := zipfiles[filename]
		if fh == nil {
			return fmt.Errorf("file %s is in manifest but not JAR", filename)
		}
		r, err := fh.Open()
		if err != nil {
			return err
		}
		if err := hashFile(keys, r, ""); err != nil {
			return fmt.Errorf("file \"%s\" in MANIFEST.MF: %w", filename, err)
		}
		if err := r.Close(); err != nil {
			return err
		}
	}
	return nil
}

type digester struct {
	key, value string
	hash       hash.Hash
}

// Verify any hashes present in a single manifest section against the given content.
func hashFile(keys http.Header, content io.Reader, suffix string) error {
	digesters := make([]digester, 0)
	suffix = "-Digest" + suffix
	for key, value := range keys {
		if !strings.HasSuffix(key, suffix) {
			continue
		}
		hashName := strings.ToUpper(key[:len(key)-len(suffix)])
		hash := x509tools.HashByName(hashName)
		if !hash.Available() {
			return fmt.Errorf("unknown digest key in manifest: %s", hashName)
		}
		digesters = append(digesters, digester{key, value[0], hash.New()})
	}
	if len(digesters) == 0 {
		return errNoDigests
	}
	buf := make([]byte, 32*1024)
	for {
		n, err := content.Read(buf)
		if n > 0 {
			for _, digester := range digesters {
				digester.hash.Write(buf[:n])
			}
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	for _, digester := range digesters {
		calculated := base64.StdEncoding.EncodeToString(digester.hash.Sum(nil))
		if calculated != digester.value {
			return fmt.Errorf("%s mismatch: manifest %s != calculated %s", digester.key, digester.value, calculated)
		}
	}
	return nil
}

func verifySigFile(sigfile, manifest []byte) (http.Header, error) {
	sfParsed, err := ParseManifest(sigfile)
	if err != nil {
		return nil, err
	}
	if err := hashFile(sfParsed.Main, bytes.NewReader(manifest), "-Manifest"); err != nil {
		if err != errNoDigests {
			return nil, fmt.Errorf("manifest signature: %w", err)
		}
		// fall through and verify all the section digests
	} else {
		// if the whole-file digest passed then skip the sections
		return sfParsed.Main, nil
	}
	sections, malformed := splitManifest(manifest)
	if malformed {
		return nil, ErrManifestLineEndings
	}
	sectionMap := make(map[string][]byte, len(sections)-1)
	for i, section := range sections {
		if i == 0 {
			if err := hashFile(sfParsed.Main, bytes.NewReader(sections[0]), "-Manifest-Main-Attributes"); err != nil {
				return nil, fmt.Errorf("manifest main attributes signature: %w", err)
			}
		} else {
			hdr, err := parseSection(section)
			if err != nil {
				return nil, err
			}
			name := hdr.Get("Name")
			if name == "" {
				return nil, errors.New("manifest has section with no \"Name\" attribute")
			}
			sectionMap[name] = section
		}
	}
	for name, keys := range sfParsed.Files {
		section := sectionMap[name]
		if section == nil {
			return nil, fmt.Errorf("manifest is missing signed section \"%s\"", name)
		}
		if err := hashFile(keys, bytes.NewReader(section), ""); err != nil {
			return nil, fmt.Errorf("manifest signature over section \"%s\": %w", name, err)
		}
	}
	return sfParsed.Main, nil
}
