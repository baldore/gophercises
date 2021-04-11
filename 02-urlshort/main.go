package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/baldore/urlshort/urlshort"
)

func main() {
	var (
		ymlFilepath  string
		jsonFilepath string
	)

	flag.StringVar(&ymlFilepath, "y", "", "Yaml file to parse.")
	flag.StringVar(&jsonFilepath, "j", "", "Json file to parse.")
	flag.Parse()

	mux := defaultMux()

	pathsToUrls, err := getOptionsFromJSONFile(jsonFilepath)
	if err != nil {
		errorString := fmt.Sprintf("Error: json file %q not found.", ymlFilepath)
		panic(errorString)
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml, err := getOptionsFromYAMLFile(ymlFilepath)
	if err != nil {
		errorString := fmt.Sprintf("Error: yml file %q not found.", ymlFilepath)
		panic(errorString)
	}

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	sqlHandler := urlshort.SQLHandler(yamlHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", sqlHandler)
}

func getOptionsFromJSONFile(filepath string) (map[string]string, error) {
	if filepath == "" {
		return map[string]string{}, nil
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	urlsMap := map[string]string{}
	fileReader := bufio.NewReader(file)
	decoder := json.NewDecoder(fileReader)

	if err := decoder.Decode(&urlsMap); err != nil {
		return nil, err
	}

	return urlsMap, nil
}

func getOptionsFromYAMLFile(filepath string) ([]byte, error) {
	if filepath == "" {
		return []byte{}, nil
	}

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
