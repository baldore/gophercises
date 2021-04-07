package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/baldore/urlshort/urlshort"
)

func main() {
	var filepath string

	flag.StringVar(&filepath, "yml", "", "Yaml string to parse.")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml, err := getOptionsFromYAMLFile(filepath)
	if err != nil {
		errorString := fmt.Sprintf("Error: file %q not found", filepath)
		panic(errorString)
	}

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func getOptionsFromYAMLFile(filepath string) ([]byte, error) {
	file, err := ioutil.ReadFile(filepath)

	return file, err
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
