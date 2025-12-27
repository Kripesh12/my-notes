package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("ok"))
	if err != nil {
		return
	}
}

func (app *application) greetHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("welcome to todo api"))
	if err != nil {
		return
	}
}
