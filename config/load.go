package config

import (
	"sync"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"log"
	"fmt"
)

var (
	Filename = "config.yml"
	ENVName  = "APP_CONFIG"
	c        *cfg
	mu       sync.Mutex
)

func Get() *cfg {
	if c == nil {
		mu.Lock()
		defer mu.Unlock()
		// Double check after mutex unlocking if config not set
		if c != nil {
			return c
		}
		// Do to load form ENV
		var filename string
		if envFilename := os.Getenv(ENVName); envFilename != "" {
			filename = envFilename
		} else {
			// Default
			curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				log.Fatalf("Can't get config path. Error: %s", err)
			}
			filename = fmt.Sprintf("%s/%s", curDir, Filename)
		}
		err := load(filename)
		if err != nil {
			log.Fatalf("Config file didn't load. Error: %s", err)
		}

	}
	return c
}

func load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	conf := cfg{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	c = &conf
	return nil
}

// After Reset(), config will be reloaded
func Reset() {
	mu.Lock()
	defer mu.Unlock()
	c = nil
}