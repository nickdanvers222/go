package main

import(
	"fmt"
	"net/http"
)

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New users hit")
}