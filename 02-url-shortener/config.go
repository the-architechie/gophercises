package main

import "net/http"

type Router struct {
	Routes   map[string]string
	Fallback http.HandlerFunc
}

type pathURL struct {
	Path string `json:"path" yaml:"path"`
	URL  string `json:"url" yaml:"url"`
}

func buildRouter(entries []pathURL) Router {
	routes := make(map[string]string, len(entries))
	for _, e := range entries {
		routes[e.Path] = e.URL
	}
	return Router{Routes: routes}
}
