package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

type BastileResponse struct {
	Target string
	Release string
}

func BastilleStartEndpoint(response http.ResponseWriter, request *http.Request) {
	out, err := exec.Command("ls", "-l").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func BastilleStopEndpoint(response http.ResponseWriter, request *http.Request) {
	// bastille stop TARGET
}

func BastilleRestartEndpoint(response http.ResponseWriter, request *http.Request) {
	// bastille restart TARGET
}
func main() {
	fmt.Println("Starting the Bastille API...")

	router := mux.NewRouter()
	router.HandleFunc("/v1/bastille/start/{target}", BastilleStartEndpoint).Methods("GET")
	router.HandleFunc("/v1/bastille/stop/{target}", BastilleStopEndpoint).Methods("GET")
	router.HandleFunc("/v1/bastille/restart/{target}", BastilleRestartEndpoint).Methods("GET")
	fmt.Println("Listening on port 12345")
	http.ListenAndServe(":12345", router)
}
