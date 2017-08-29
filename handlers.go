package main

import (
	"fmt"
	"net/http"
)

//Index plain handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

//Your Handlers here
