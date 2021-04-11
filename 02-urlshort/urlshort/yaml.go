package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type option struct {
	Path string
	URL  string
}

type Options []option

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	options := Options{}

	err := yaml.Unmarshal(yml, &options)
	if err != nil {
		return fallback.ServeHTTP, err
	}

	urlMap := map[string]string{}
	for _, option := range options {
		urlMap[option.Path] = option.URL
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if url, ok := urlMap[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}

	return handler, nil
}
