package main

import (
	"encoding/json"
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

	rw.Header().Set("Content-Type", "application/json")

	env := os.Getenv("env_test")
	versionResponse := TestVersion{
		Version: env,
	}

	json.NewEncoder(rw).Encode(versionResponse)
}
