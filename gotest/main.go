package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Welcome to our website")
	//fmt.Fprint(w, "Welcome to our website")
	return
}

func main() {
	http.HandleFunc("/", welcome)

	http.ListenAndServe(":3000", nil)
}
