package tooling

import (
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"os"

	"runtime.link/api/xray"
)

func Download(dest, unzip, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", xray.New(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", xray.New(err)
	}
	var body = resp.Body
	if unzip != "" {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", xray.New(err)
		}
		archive, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
		if err != nil {
			return "", xray.New(err)
		}
		inZip, err := archive.Open(unzip)
		if err != nil {
			return "", xray.New(err)
		}
		defer inZip.Close()
		body = inZip
	}
	//executable
	file, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return "", xray.New(err)
	}
	defer file.Close()
	if _, err = io.Copy(file, body); err != nil {
		return "", xray.New(err)
	}
	return dest, nil
}
