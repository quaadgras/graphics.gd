// The 'gd' command is designed as a drop-in replacement of the 'go' command when working
// with Godot-based projects. It will automatically download and install the supported
// version of Godot and ensure all go commands behave as expected (go build, go run, etc).
//
// Running 'gd' without any command line arguments will launch the Godot editor so that
// users can focus on designing the visual aspects of their project.
//
// The 'gd' command assumes that the Go module lives at the root of the project, an empty
// godot project will be created under a 'graphics' directory, this is where the user can
// edit the graphical representation of their project and manage their assets.
package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"log"
)

func main() {
	if err := wrap(); err != nil {
		log.Fatal(err)
	}
}

func useGodot() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	info, err := os.Stat(gopath + "/bin/godot-4.2.1")
	if os.IsNotExist(err) {
		switch runtime.GOOS {
		case "linux":
			fmt.Println("gd: downloading Godot v4.2.1 stable for linux")
			resp, err := http.Get("https://github.com/godotengine/godot-builds/releases/download/4.2.1-stable/Godot_v4.2.1-stable_linux.x86_64.zip")
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				return "", err
			}
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return "", err
			}
			archive, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
			if err != nil {
				return "", err
			}
			inZip, err := archive.Open("Godot_v4.2.1-stable_linux.x86_64")
			if err != nil {
				return "", err
			}
			defer inZip.Close()
			file, err := os.Create(gopath + "/bin/godot-4.2.1")
			if err != nil {
				return "", err
			}
			defer file.Close()
			if _, err = io.Copy(file, inZip); err != nil {
				return "", err
			}

		default:
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	if info.Mode()&0111 == 0 {
		if err := os.Chmod(gopath+"/bin/godot-4.2.1", 0755); err != nil {
			return "", err
		}
	}
	return gopath + "/bin/godot-4.2.1", nil
}

func wrap() error {
	if runtime.GOARCH != "amd64" && runtime.GOARCH != "arm64" {
		return errors.New("gd requires an amd64, or an arm64 system")
	}
	godot, err := useGodot()
	if err != nil {
		return fmt.Errorf("gd requires Godot to be installed as a binary at $GOPATH/bin/godot-4.2.1: %w", err)
	}
	if err := os.MkdirAll("./graphics", 0755); err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if _, err := os.Stat("./graphics/project.godot"); os.IsNotExist(err) {
		if err := os.WriteFile("./graphics/project.godot", []byte(`; Engine configuration file.
; It's best edited using the editor UI and not directly,
; since the parameters that go here are not all obvious.
;
; Format:
;   [section] ; section goes between []
;   param=value ; assign values to parameters

config_version=5

[application]

config/name="`+filepath.Base(wd)+`"`), 0644); err != nil {
			return err
		}
	}
	if len(os.Args) == 1 {
		godot := exec.Command(godot, "-e")
		godot.Dir = "./graphics"
		godot.Stderr = os.Stderr
		godot.Stdout = os.Stdout
		godot.Stdin = os.Stdin
		return godot.Run()
	}
	return errors.New("not implemented")
}
