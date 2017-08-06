package tpl

import (
	"github.com/flosch/pongo2"
	"fmt"
	"net/http"
	"github.com/alehano/gobootstrap/config"
)

type D map[string]interface{} // Syntax sugar alias
var store = map[string]*pongo2.Template{}
var storePaths = map[string]string{}
var defaultData = map[string]interface{}{}

func Register(name, path string) {
	if _, ok := store[name]; ok {
		panic(fmt.Sprintf("Template %q already exists", name))
		return
	}
	if config.Get().Debug {
		storePaths[name] = path
	} else {
		store[name] = pongo2.Must(pongo2.FromFile(path))
	}
}

func RegisterMulti(baseDir string, names map[string]string) {
	for name, path := range names {
		Register(name, baseDir+path)
	}
}

func RegisterDefaultData(data map[string]interface{}) {
	for key, value := range data {
		if _, exists := defaultData[key]; exists {
			panic(fmt.Sprintf("Data key %q already exists", key))
		} else {
			defaultData[key] = value
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	// Add Context
	data["context"] = r.Context().Value
	// Add default data
	for key, value := range defaultData {
		if _, dataExists := data[key]; !dataExists {
			data[key] = value
		}
	}

	var t *pongo2.Template
	exists := false

	// Rebuild on the fly in Debug mode
	if config.Get().Debug {
		path := ""
		path, exists = storePaths[name]
		if exists {
			t = pongo2.Must(pongo2.FromFile(path))
		}
	} else {
		t, exists = store[name]
	}

	if exists {
		err := t.ExecuteWriter(pongo2.Context(data), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, fmt.Sprintf("Template %q not exists", name), http.StatusInternalServerError)
	}
}