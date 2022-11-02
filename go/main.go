package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type TestVersion struct {
	Version string
}

func main() {
	http.HandleFunc("/api", apiVersion)
	http.ListenAndServe(":8080", nil)
}

func apiVersion(rw http.ResponseWriter, req *http.Request) {

	httpMethod := req.Method
	env := os.Getenv("server-name")

	log.Printf("Request method: %s | Server name: %s", httpMethod, env)

	rw.Header().Set("Content-Type", "application/json")

	versionResponse := TestVersion{
		Version: env,
	}

	json.NewEncoder(rw).Encode(versionResponse)
}
