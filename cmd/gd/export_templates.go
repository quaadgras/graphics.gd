package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/schollz/progressbar/v3"
	"graphics.gd/cmd/gd/internal/tooling"
	"runtime.link/api/xray"
)

func AssertExportTemplates(version string) error {
	var location string
	switch runtime.GOOS {
	case "linux":
		location = filepath.Join(os.Getenv("HOME"), ".local", "share", "godot", "export_templates", version+".stable")
	case "windows":
		location = filepath.Join(os.Getenv("APPDATA"), "Godot", "export_templates", version+".stable")
	case "darwin":
		location = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Godot", "export_templates", version+".stable")
	default:
		return nil
	}
	var url = "https://github.com/godotengine/godot/releases/download/" + version + "-stable/Godot_v" + version + "-stable_export_templates.tpz"
	if _, err := os.Stat(location); err == nil {
		return nil
	}
	var dest = filepath.Join(filepath.Dir(location), version+".stable.download")
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
	if err := os.Remove(dest); err != nil {
		return xray.New(err)
	}
	return nil
}
