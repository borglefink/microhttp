// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Hostname        string
	Port            int
	DefaultDocument string
}

var configFilename = ""

// ------------------------------------------
// init
// ------------------------------------------
func init() {
	// Set the config file name to [thisexecutablefilename].config
	configFilename = strings.Replace(filepath.Base(os.Args[0]), ".exe", ".config", 1)
}

// ------------------------------------------
// save
// ------------------------------------------
func save(c Config) {
	var jsonstring, err = json.MarshalIndent(&c, "", "  ")
	if err != nil {
		fmt.Printf("json.Marshal(c), %s %v\n", string(jsonstring), err)
	}

	err = ioutil.WriteFile(configFilename, jsonstring, 0666)
	if err != nil {
		fmt.Printf("ioutil.WriteFile, %v\n", err)
	}
}

// ------------------------------------------
// Load
// ------------------------------------------
func Load() Config {
	var c Config

	// Read whole the file
	var jsonstring, err = ioutil.ReadFile(configFilename)
	if err != nil {
		var c = Config{"localhost", 80, "index.html"}
		save(c)
		return c
	}

	err = json.Unmarshal(jsonstring, &c)

	if err != nil {
		fmt.Printf("json.Unmarshal, %v\n", err)
	}

	return c
}
