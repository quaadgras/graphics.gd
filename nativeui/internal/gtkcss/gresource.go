package gtkcss

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

// Themes like Ubuntu's Yaru keep their stylesheets and assets in a compiled
// GResource bundle, gtk-3.0/gtk.gresource, which GTK registers before it
// loads gtk.css; their gtk.css is then a single
// @import url("resource:///...") into it.

const resourceScheme = "resource://"

var resources sync.Map // resource path ("/com/ubuntu/...") -> []byte

// LoadResources registers the files of the GResource bundle at path, so that
// resource:// URLs resolve to them.
func LoadResources(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	files, err := parseGResource(data)
	if err != nil {
		return err
	}
	for name, content := range files {
		resources.Store(name, content)
	}
	return nil
}

// readFile reads a file on disk or a registered resource:// one.
func readFile(name string) ([]byte, error) {
	if rest, ok := strings.CutPrefix(name, resourceScheme); ok {
		if v, ok := resources.Load(path.Clean("/" + rest)); ok {
			return v.([]byte), nil
		}
		return nil, os.ErrNotExist
	}
	return os.ReadFile(name)
}

// resolve is ref relative to the directory base, either of which may be a
// resource:// URL.
func resolve(base, ref string) string {
	if strings.HasPrefix(ref, resourceScheme) || filepath.IsAbs(ref) {
		return ref
	}
	if rest, ok := strings.CutPrefix(base, resourceScheme); ok {
		return resourceScheme + path.Join("/"+rest, ref)
	}
	return filepath.Join(base, ref)
}

// dir is the directory of a file on disk or a resource:// URL.
func dir(name string) string {
	if rest, ok := strings.CutPrefix(name, resourceScheme); ok {
		return resourceScheme + path.Dir("/"+rest)
	}
	return filepath.Dir(name)
}

var errGResource = errors.New("gtkcss: not a GResource bundle")

// parseGResource reads a GResource bundle: a GVDB hash table whose keys are
// paths ("/com/ubuntu/themes/Yaru/3.0/gtk.css", each made of its parent's
// key and its own) and whose values are (uuay) variants: the size, the
// flags (1: zlib compressed) and the data.
func parseGResource(data []byte) (map[string][]byte, error) {
	le := binary.LittleEndian
	if len(data) < 24 || string(data[:8]) != "GVariant" {
		return nil, errGResource
	}
	slice := func(start, end uint32) ([]byte, bool) {
		if start > end || int(end) > len(data) {
			return nil, false
		}
		return data[start:end], true
	}
	table, ok := slice(le.Uint32(data[16:]), le.Uint32(data[20:]))
	if !ok || len(table) < 8 {
		return nil, errGResource
	}
	nBloom := le.Uint32(table[0:]) & (1<<27 - 1)
	nBuckets := le.Uint32(table[4:])
	itemsAt := 8 + 4*(uint64(nBloom)+uint64(nBuckets))
	if itemsAt > uint64(len(table)) {
		return nil, errGResource
	}
	const itemSize = 24
	items := table[itemsAt:]
	n := len(items) / itemSize
	type item struct {
		parent uint32
		key    string
		kind   byte
		value  []byte
	}
	all := make([]item, n)
	for i := range all {
		it := items[i*itemSize:]
		key, ok := slice(le.Uint32(it[8:]), le.Uint32(it[8:])+uint32(le.Uint16(it[12:])))
		if !ok {
			return nil, errGResource
		}
		value, _ := slice(le.Uint32(it[16:]), le.Uint32(it[20:]))
		all[i] = item{parent: le.Uint32(it[4:]), key: string(key), kind: it[14], value: value}
	}
	fullKey := func(i int) string {
		key := ""
		for depth := 0; depth <= n; depth++ {
			key = all[i].key + key
			p := all[i].parent
			if p == 0xffffffff || int(p) >= n {
				return key
			}
			i = int(p)
		}
		return key // a cycle: give up on the prefix
	}
	files := map[string][]byte{}
	for i, it := range all {
		if it.kind != 'v' || strings.HasSuffix(it.key, "/") {
			continue
		}
		// A variant: the (uuay) value, a nul, then its type string.
		v := it.value
		j := bytes.LastIndexByte(v, 0)
		if j < 8 || string(v[j+1:]) != "(uuay)" {
			continue
		}
		size, flags, content := le.Uint32(v[0:]), le.Uint32(v[4:]), v[8:j]
		if flags&1 != 0 {
			r, err := zlib.NewReader(bytes.NewReader(content))
			if err != nil {
				continue
			}
			content, err = io.ReadAll(io.LimitReader(r, int64(size)))
			if err != nil {
				continue
			}
		} else if uint32(len(content)) > size {
			content = content[:size] // the nul GResource appends
		}
		files[fullKey(i)] = content
	}
	return files, nil
}
