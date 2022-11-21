package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type TestVersion struct {
	Version string
}

var toggle bool = false

func main() {
	http.HandleFunc("/apiVersion", ApiVersion)
	http.HandleFunc("/health", Health)
	http.HandleFunc("/toggle", Toggle)
	go ToggleReady()
	http.ListenAndServe(":8080", nil)
}

func ApiVersion(rw http.ResponseWriter, req *http.Request) {

	httpMethod := req.Method
	env := os.Getenv("server-name")

	log.Printf("Request method: %s | Server name: %s", httpMethod, env)

	rw.Header().Set("Content-Type", "application/json")

	versionResponse := TestVersion{
		Version: env,
	}

	json.NewEncoder(rw).Encode(versionResponse)
}

func Toggle(rw http.ResponseWriter, req *http.Request) {
	toggle = !toggle
	rw.WriteHeader(http.StatusOK)
}
func Health(rw http.ResponseWriter, req *http.Request) {
	if toggle == true {
		rw.WriteHeader(http.StatusOK)
		return
	}
	rw.WriteHeader(http.StatusInternalServerError)

}

func ToggleReady() {
	time.Sleep(time.Second * 10)
	fmt.Println("Server ready! ")
	toggle = true
}
