package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"
	"graphics.gd/cmd/gd/internal/tooling"
	"runtime.link/api/xray"
)

// templateSpec describes which files, out of the monolithic export-templates
// .tpz, a given target GOOS needs — plus a marker file whose presence in the
// templates directory means that platform is already installed.
type templateSpec struct {
	prefixes []string // template-file basename prefixes to extract
	marker   string   // basename whose presence means "already installed"
}

var templateSpecs = map[string]templateSpec{
	"android": {[]string{"android_"}, "android_release.apk"},
	"windows": {[]string{"windows_"}, "windows_release_x86_64.exe"},
	"linux":   {[]string{"linux_"}, "linux_release.x86_64"},
	"darwin":  {[]string{"macos."}, "macos.zip"},
	"ios":     {[]string{"ios."}, "ios.zip"},
}

// templateCommon files are needed by every platform's export.
var templateCommon = []string{"version.txt", "icudt_godot.dat"}

func exportTemplatesDir(version string) (string, bool) {
	switch runtime.GOOS {
	case "linux":
		return filepath.Join(os.Getenv("HOME"), ".local", "share", "godot", "export_templates", version+".stable"), true
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "Godot", "export_templates", version+".stable"), true
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Godot", "export_templates", version+".stable"), true
	}
	return "", false
}

// AssertExportTemplates makes sure the Godot export templates needed for the
// given target platform (goos) are installed. When the platform is known it
// pulls only that platform's files out of the remote .tpz via HTTP range
// requests instead of downloading the whole (~1.3GB) archive, falling back to a
// full download if range extraction is unavailable.
func AssertExportTemplates(version string, goos string) error {
	location, ok := exportTemplatesDir(version)
	if !ok {
		return nil
	}
	// js/web uses graphics.gd's own web template (handled by the browser
	// builder), not the stock Godot web templates.
	if goos == "js" || goos == "web" {
		return nil
	}
	url := "https://github.com/godotengine/godot/releases/download/" + version + "-stable/Godot_v" + version + "-stable_export_templates.tpz"
	if spec, selective := templateSpecs[goos]; selective {
		if _, err := os.Stat(filepath.Join(location, spec.marker)); err == nil {
			return nil // already installed for this platform
		}
		if err := extractTemplatesForPlatform(url, location, spec); err != nil {
			fmt.Printf("gd: per-platform template fetch failed (%v); falling back to full download\n", err)
			return fullDownloadTemplates(version, url, location)
		}
		return nil
	}
	// Unknown platform: keep the original whole-archive behaviour.
	if _, err := os.Stat(location); err == nil {
		return nil
	}
	return fullDownloadTemplates(version, url, location)
}

func wantTemplateFile(base string, spec templateSpec) bool {
	for _, c := range templateCommon {
		if base == c {
			return true
		}
	}
	for _, p := range spec.prefixes {
		if strings.HasPrefix(base, p) {
			return true
		}
	}
	return false
}

func extractTemplatesForPlatform(url, location string, spec templateSpec) error {
	size, err := remoteSize(url)
	if err != nil {
		return xray.New(err)
	}
	// archive/zip reads each entry sequentially; the block buffer turns the many
	// small ReadAts into a few large HTTP range requests.
	ra := &blockReaderAt{ra: httpReaderAt{url: url}, size: size, block: 8 << 20}
	zr, err := zip.NewReader(ra, size)
	if err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll(location, 0o755); err != nil {
		return xray.New(err)
	}
	var extracted int
	for _, f := range zr.File {
		if strings.HasSuffix(f.Name, "/") {
			continue
		}
		base := path.Base(f.Name)
		if !wantTemplateFile(base, spec) {
			continue
		}
		if err := extractZipEntry(f, filepath.Join(location, base)); err != nil {
			return xray.New(err)
		}
		extracted++
	}
	if extracted == 0 {
		return fmt.Errorf("no template files matched %v in %s", spec.prefixes, path.Base(url))
	}
	return nil
}

func extractZipEntry(f *zip.File, dest string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	out, err := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o755)
	if err != nil {
		return err
	}
	defer out.Close()
	bar := progressbar.DefaultBytes(int64(f.UncompressedSize64), "gd: "+path.Base(dest))
	if _, err := io.Copy(io.MultiWriter(out, bar), rc); err != nil {
		return err
	}
	return nil
}

// httpReaderAt is an io.ReaderAt backed by HTTP range requests, so archive/zip
// can read a remote zip's central directory and entries without downloading the
// whole file. http.DefaultClient follows the GitHub release redirect, preserving
// the Range header, so each call lands on the signed asset URL.
type httpReaderAt struct{ url string }

func (h httpReaderAt) ReadAt(p []byte, off int64) (int, error) {
	req, err := http.NewRequest("GET", h.url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", off, off+int64(len(p))-1))
	req.Header.Set("User-Agent", "graphics.gd/cmd/gd")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("range GET %s: %s", h.url, resp.Status)
	}
	return io.ReadFull(resp.Body, p)
}

// blockReaderAt caches one large, block-aligned chunk of an underlying ReaderAt,
// so the sequential reads archive/zip performs hit the cache instead of issuing
// a request per read.
type blockReaderAt struct {
	ra     io.ReaderAt
	size   int64
	block  int64
	bufOff int64
	buf    []byte
}

func (b *blockReaderAt) ReadAt(p []byte, off int64) (int, error) {
	got := 0
	for got < len(p) {
		cur := off + int64(got)
		if cur >= b.size {
			return got, io.EOF
		}
		if b.buf == nil || cur < b.bufOff || cur >= b.bufOff+int64(len(b.buf)) {
			start := (cur / b.block) * b.block
			n := b.block
			if start+n > b.size {
				n = b.size - start
			}
			tmp := make([]byte, n)
			m, err := b.ra.ReadAt(tmp, start)
			if m == 0 {
				return got, err
			}
			b.bufOff, b.buf = start, tmp[:m]
		}
		got += copy(p[got:], b.buf[cur-b.bufOff:])
	}
	return got, nil
}

func remoteSize(url string) (int64, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Range", "bytes=0-0")
	req.Header.Set("User-Agent", "graphics.gd/cmd/gd")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return resp.ContentLength, nil // server ignored Range
	}
	if resp.StatusCode != http.StatusPartialContent {
		return 0, fmt.Errorf("range GET %s: %s", url, resp.Status)
	}
	_, total, ok := strings.Cut(resp.Header.Get("Content-Range"), "/")
	if !ok {
		return 0, fmt.Errorf("missing Content-Range total")
	}
	return strconv.ParseInt(strings.TrimSpace(total), 10, 64)
}

// fullDownloadTemplates downloads and extracts the entire .tpz (all platforms).
// Used as a fallback when per-platform range extraction is unavailable, and for
// platforms without a known file set.
func fullDownloadTemplates(version, url, location string) error {
	var dest = filepath.Join(filepath.Dir(location), version+".stable.download")
	if err := func() error {
		out, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return xray.New(err)
		}
		defer out.Close()
		stat, err := out.Stat()
		if err != nil {
			return xray.New(err)
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return xray.New(err)
		}
		if stat.Size() > 0 {
			req.Header.Set("Range", fmt.Sprintf("bytes=%d-", stat.Size()))
		}
		req.Header.Set("User-Agent", "graphics.gd/cmd/gd")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return xray.New(err)
		}
		defer resp.Body.Close()
		switch resp.StatusCode {
		case 200:
		case 206:
			if _, err := out.Seek(stat.Size(), io.SeekStart); err != nil {
				return xray.New(err)
			}
		case 416:
			contentRange := resp.Header.Get("Content-Range")
			if contentRange != fmt.Sprintf("bytes */%d", stat.Size()) {
				return fmt.Errorf("unable to resume download of 'export templates' (required for 'gd build'), please delete %v and try again\nGET %s HTTP status: %v", dest, url, resp.StatusCode)
			}
		default:
			return fmt.Errorf(
				"unable to download 'export templates' (required for 'gd build') and not found in %s, please install them, ie. %v\nGET %s HTTP status: %v",
				filepath.Dir(location), "https://godotengine.org/download/linux/", url, resp.StatusCode,
			)
		}
		if resp.StatusCode != 416 {
			bar := progressbar.DefaultBytes(
				resp.ContentLength,
				"gd: downloading export templates",
			)
			if _, err := io.Copy(io.MultiWriter(out, bar), resp.Body); err != nil {
				return xray.New(err)
			}
		}
		if err := tooling.ExtractArchive(dest, location, "zip", "", true); err != nil {
			return xray.New(err)
		}
		return nil
	}(); err != nil {
		return xray.New(err)
	}
	if err := os.Remove(dest); err != nil {
		return xray.New(err)
	}
	return nil
}
