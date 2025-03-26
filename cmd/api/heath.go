package main

import "net/http"

func (app *application) heathCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
