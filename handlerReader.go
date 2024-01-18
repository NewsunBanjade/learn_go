package main

import "net/http"

func handlerReader(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
