package main

import (
	"log"
	"net/http"
	"os"
)

func init() {
	requiredEnvs := []string{}

	for _, env := range requiredEnvs {
		if len(os.Getenv(env)) == 0 {
			log.Fatalf("`%s` environment variable not set", env)
		}
	}
}

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":80", router))
}
