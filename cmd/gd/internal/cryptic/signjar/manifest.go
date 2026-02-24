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
	"bytes"
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"graphics.gd/cmd/gd/internal/cryptic/x509tools"
)

// See https://docs.oracle.com/javase/8/docs/technotes/guides/jar/jar.html#JAR_Manifest

const (
	metaInf      = "META-INF/"
	manifestName = metaInf + "MANIFEST.MF"
)

var ErrManifestLineEndings = errors.New("manifest has incorrect line ending sequence")

type FilesMap struct {
	Main  http.Header
	Order []string
	Files map[string]http.Header
}

func ParseManifest(manifest []byte) (files *FilesMap, err error) {
	files, malformed, err := parseManifest(manifest)
	if err != nil {
		return nil, err
	} else if malformed {
		return nil, ErrManifestLineEndings
	}
	return files, nil
}

func parseManifest(manifest []byte) (files *FilesMap, malformed bool, err error) {
	sections, malformed := splitManifest(manifest)
	if len(sections) == 0 {
		return &FilesMap{
			Order: make([]string, 0, 0),
			Files: make(map[string]http.Header, 0),
		}, false, nil
	}
	files = &FilesMap{
		Order: make([]string, 0, len(sections)-1),
		Files: make(map[string]http.Header, len(sections)-1),
	}
	for i, section := range sections {
		if i > 0 && len(section) == 0 {
			continue
		}
		hdr, err := parseSection(section)
		if err != nil {
			return nil, false, err
		}
		if i == 0 {
			files.Main = hdr
		} else {
			name := hdr.Get("Name")
			if name == "" {
				return nil, false, errors.New("manifest has section with no \"Name\" attribute")
			}
			files.Order = append(files.Order, name)
			files.Files[name] = hdr
		}
	}
	return files, malformed, nil
}

func (m *FilesMap) Dump() []byte {
	var out bytes.Buffer
	writeSection(&out, m.Main, "Manifest-Version")
	for _, name := range m.Order {
		section := m.Files[name]
		if section != nil {
			writeSection(&out, section, "Name")
		}
	}
	return out.Bytes()
}

func splitManifest(manifest []byte) ([][]byte, bool) {
	var malformed bool
	sections := make([][]byte, 0)
	for len(manifest) != 0 {
		i1 := bytes.Index(manifest, []byte("\r\n\r\n"))
		i2 := bytes.Index(manifest, []byte("\n\n"))
		var idx int
		switch {
		case i1 >= 0:
			idx = i1 + 4
		case i2 >= 0:
			idx = i2 + 2
		default:
			// If there is not a proper 2x line ending,
			// then it's technically not valid but we can sign it anyway
			// as long as it gets rewritten with correct endings.
			idx = len(manifest)
			malformed = true
		}
		section := manifest[:idx]
		manifest = manifest[idx:]
		if len(bytes.TrimSpace(section)) == 0 {
			// Excessive line endings have created an empty section
			malformed = true
			continue
		}
		sections = append(sections, section)
	}
	return sections, malformed
}

func parseSection(section []byte) (http.Header, error) {
	section = bytes.ReplaceAll(section, []byte("\r\n"), []byte{'\n'})
	section = bytes.ReplaceAll(section, []byte("\n "), []byte{})
	keys := bytes.Split(section, []byte{'\n'})
	hdr := make(http.Header)
	for _, line := range keys {
		if len(line) == 0 {
			continue
		}
		idx := bytes.IndexRune(line, ':')
		if idx < 0 {
			return nil, errors.New("jar manifest is malformed")
		}
		key := strings.TrimSpace(string(line[:idx]))
		value := strings.TrimSpace(string(line[idx+1:]))
		hdr.Set(key, value)
	}
	return hdr, nil
}

func hashSection(hash crypto.Hash, section []byte) string {
	d := hash.New()
	d.Write(section)
	return base64.StdEncoding.EncodeToString(d.Sum(nil))
}

// Transform a MANIFEST.MF into a *.SF by digesting each section with the
// specified hash
func DigestManifest(manifest []byte, hash crypto.Hash, sectionsOnly, apkV2 bool) ([]byte, error) {
	sections, malformed := splitManifest(manifest)
	if malformed {
		return nil, ErrManifestLineEndings
	}
	hashName := x509tools.HashNames[hash]
	if hashName == "" {
		return nil, errors.New("unsupported hash type")
	}
	var output bytes.Buffer
	writeAttribute(&output, "Signature-Version", "1.0")
	//writeAttribute(&output, hashName+"-Digest-Manifest-Main-Attributes", hashSection(hash, sections[0]))
	if !sectionsOnly {
		writeAttribute(&output, hashName+"-Digest-Manifest", hashSection(hash, manifest))
	}
	writeAttribute(&output, "Created-By", "Signflinger")
	if apkV2 {
		writeAttribute(&output, "X-Android-APK-Signed", "2")
	}
	output.WriteString("\r\n")
	for _, section := range sections {
		hdr, err := parseSection(section)
		if err != nil {
			return nil, err
		}
		name := hdr.Get("Name")
		if name == "" {
			continue // main section
		}
		writeAttribute(&output, "Name", name)
		writeAttribute(&output, hashName+"-Digest", hashSection(hash, section))
		output.WriteString("\r\n")
	}
	return output.Bytes(), nil
}

const maxLineLength = 70

// Write a key-value pair, wrapping long lines as necessary
func writeAttribute(out *bytes.Buffer, key, value string) {
	line := fmt.Appendf(nil, "%s: %s", key, value)
	for i := 0; i < len(line); {
		goal := maxLineLength
		if i != 0 {
			out.Write([]byte{' '})
			goal--
		}
		j := min(i+goal, len(line))
		out.Write(line[i:j])
		out.Write([]byte("\r\n"))
		i = j
	}
}

func writeSection(out *bytes.Buffer, hdr http.Header, first string) {
	value := hdr.Get(first)
	if value != "" {
		writeAttribute(out, first, value)
	}
	keys := make([]string, 0, len(hdr))
	for key := range hdr {
		if key == first {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		for _, value := range hdr[key] {
			writeAttribute(out, key, value)
		}
	}
	out.Write([]byte("\r\n"))
}
