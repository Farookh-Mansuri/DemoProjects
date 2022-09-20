package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	Name     string
	LastName string
	Age      uint8
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloHandler).Methods("Get")
	r.HandleFunc("/Hi", HiHandler).Methods("Get")
	r.HandleFunc("/show", showHandler).Methods("Get")
	r.HandleFunc("/", sendResponse).Methods("Get")
	log.Println("Listening...")

	http.ListenAndServe("localhost:9000", r)
}
func HelloHandler(resp http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(resp, "hello I am here")
	resp.WriteHeader(http.StatusOK)
}
func HiHandler(resp http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(resp, "Hi I am here")
	resp.WriteHeader(http.StatusOK)
}
func showHandler(resp http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(resp, "Show I am here")
	resp.WriteHeader(http.StatusOK)
}
func sendResponse(response http.ResponseWriter, request *http.Request) {
	person := person{Name: "Shashank", LastName: "Tiwari", Age: 30}

	jsonResponse, jsonError := json.Marshal(person)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	fmt.Println(string(jsonResponse))

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
