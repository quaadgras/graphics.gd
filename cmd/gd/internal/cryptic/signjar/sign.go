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
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"errors"
	"path"
	"strings"
	"time"

	"graphics.gd/cmd/gd/internal/cryptic/binpatch"
	"graphics.gd/cmd/gd/internal/cryptic/certloader"
	"graphics.gd/cmd/gd/internal/cryptic/pkcs7"
	"graphics.gd/cmd/gd/internal/cryptic/pkcs9"
	"graphics.gd/cmd/gd/internal/cryptic/zipslicer"
)

func (jd *JarDigest) Sign(ctx context.Context, cert *certloader.Certificate, alias string, sectionsOnly, inlineSignature, apkV2 bool) (*binpatch.PatchSet, *pkcs9.TimestampedSignature, error) {
	// Create sigfile from the manifest
	sf, err := DigestManifest(jd.Manifest, jd.Hash, sectionsOnly, apkV2)
	if err != nil {
		return nil, nil, err
	}
	// Sign sigfile
	sig := pkcs7.NewBuilder(cert.Signer(), cert.Chain(), jd.Hash)
	if err := sig.SetContentData(sf); err != nil {
		return nil, nil, err
	}
	psd, err := sig.Sign()
	if err != nil {
		return nil, nil, err
	}
	ts, err := pkcs9.TimestampAndMarshal(ctx, psd, cert.Timestamper, false)
	if err != nil {
		return nil, nil, err
	}
	// Rebuild zip with updated manifest, sigfile, and signature
	psig := ts.Raw
	if !inlineSignature {
		if _, err := psd.Detach(); err != nil {
			return nil, nil, err
		}
		psig, err = asn1.Marshal(*psd)
		if err != nil {
			return nil, nil, err
		}
	}
	patch, err := jd.insertSignature(cert.Leaf, alias, sf, psig)
	if err != nil {
		return nil, nil, err
	}
	return patch, ts, nil
}

func (jd *JarDigest) insertSignature(cert *x509.Certificate, alias string, sf, sig []byte) (*binpatch.PatchSet, error) {
	signame, pkcsname := sigNames(cert.PublicKey, alias)
	deflate := jd.shouldDeflate()
	// Add new files to beginning of zip
	outz := new(zipslicer.Directory)
	var zipcon bytes.Buffer
	mtime := time.Now()
	if _, err := outz.NewFile(manifestName, jarMagic, jd.Manifest, &zipcon, mtime, deflate, false); err != nil {
		return nil, err
	}
	if _, err := outz.NewFile(metaInf+signame, nil, sf, &zipcon, mtime, deflate, false); err != nil {
		return nil, err
	}
	if _, err := outz.NewFile(metaInf+pkcsname, nil, sig, &zipcon, mtime, deflate, false); err != nil {
		return nil, err
	}
	// Patch out old files
	patch := binpatch.New()
	patch.Add(0, 0, zipcon.Bytes())
	for _, f := range jd.inz.File {
		if keepFile(f.Name) {
			// Add existing file to the new zip directory. Its offset will be changed.
			if _, err := outz.AddFile(f); err != nil {
				return nil, err
			}
		} else {
			// remove this region from the old zip
			size, err := f.GetTotalSize()
			if err != nil {
				return nil, err
			}
			if size > 0xffffffff {
				return nil, errors.New("signature file too big")
			}
			patch.Add(int64(f.Offset), size, nil)
		}
	}
	zipdir := new(bytes.Buffer)
	if err := outz.WriteDirectory(zipdir, zipdir, false); err != nil {
		return nil, err
	}
	patch.Add(jd.inz.DirLoc, jd.inz.Size-jd.inz.DirLoc, zipdir.Bytes())
	return patch, nil
}

// name for the signature file is based on the key type
func sigNames(pubkey crypto.PublicKey, alias string) (signame, pkcsname string) {
	signame = strings.ToUpper(alias) + ".SF"
	pkcsname = strings.ToUpper(alias)
	switch pubkey.(type) {
	case *rsa.PublicKey:
		pkcsname += ".RSA"
	case *ecdsa.PublicKey:
		pkcsname += ".EC"
	default:
		signame = "SIG-" + signame
		pkcsname = "SIG-" + pkcsname + ".SIG"
	}
	return
}

// deflate if the original manifest was deflated
func (jd *JarDigest) shouldDeflate() bool {
	for _, f := range jd.inz.File {
		if f.Name == manifestName {
			return f.Method != zip.Store
		}
	}
	return false
}

func keepFile(name string) bool {
	if name == metaInf {
		// META-INF/ itself gets updated
		return false
	}
	if path.Dir(name)+"/" != metaInf {
		// everything not an immediate child of META-INF/ is kept
		return true
	}
	name = path.Base(name)
	switch {
	case strings.HasPrefix(name, "SIG-"):
		// delete all old signatures
		return false
	case name == "MANIFEST.MF":
		// replace the manifest
		return false
	}
	switch path.Ext(name) {
	case ".SF":
		// delete all old signatures
		return false
	case ".RSA", ".DSA", ".EC", ".SIG":
		return false
	default:
		// all other META-INF/ files are kept
		return true
	}
}
