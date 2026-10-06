package main

import (
	"net/http"
)

func (rt Router) MapHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := r.URL.Path
		longUrl, ok := rt.Routes[route]
		if ok {
			http.Redirect(w, r, longUrl, http.StatusFound)
			return
		}

		errorHandler(w, r)
	}
}
