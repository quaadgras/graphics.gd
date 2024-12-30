package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"graphics.gd/internal/gdjson"
)

// LoadSpecification, either from a local file or by downloading
// it from the Godot Github repository.
func LoadSpecification() (*gdjson.Specification, error) {
	file, err := os.Open("./extension_api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return nil, err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		file, err = os.Create("./extension_api.json")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return nil, err
		}
		file.Seek(0, 0)
		resp.Body.Close()
	}
	var spec gdjson.Specification
	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		return nil, err
	}
	return &spec, nil
}
