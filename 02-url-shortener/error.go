package main

import "net/http"

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/gophercises", 302)
}
