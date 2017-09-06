package main

import (
	"github.com/gorilla/mux"
	"net/http"
	// "log"
	"encoding/json"

	"fmt"
)

type dogs struct {
	ID    string `json:"id"`
	Color string `json:"color"`
	Breed string `json:"breed"`
}

func readparameter(w http.ResponseWriter, req *http.Request) {
	var dog dogs
	parameter := mux.Vars(req)
	dog.ID = parameter["id"]
	json.NewEncoder(w).Encode(dog.ID)
}
func Readmultiplequeryparameters(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello", req.URL.RawQuery)
	var dog dogs
	f := req.URL.Query()
	dog.Color = f["color"][0]
	dog.Breed = f["breed"][0]
	fmt.Fprintln(w, "breed=", dog.Breed)
	fmt.Fprintln(w, "color=", dog.Color)

}
func Storedatafrombody(w http.ResponseWriter, r *http.Request) {
	var dog dogs
	f := r.Body
	json.NewDecoder(f).Decode(&dog)
	fmt.Fprintln(w, "breed=", dog.Breed)
	fmt.Fprintln(w, "color=", dog.Color)
}
func Readfromheader(w http.ResponseWriter, r *http.Request) {
	var dog dogs
	byte := r.Header.Get("Breed")
	byte1 := r.Header.Get("Color")
	dog.Color = string(byte1)
	dog.Breed = string(byte)
	fmt.Fprintln(w, "Breed=", dog.Breed)
	fmt.Fprintln(w, "color=", dog.Color)
}

func main() {
	fmt.Println("welcome")
	//router := mux.NewRouter()
	//router.HandleFunc("/dogs/{id}",readparameter).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8081",router))
	http.HandleFunc("/dog", Readmultiplequeryparameters) //read query parameters with multiple values
	http.HandleFunc("/dogwrite", Storedatafrombody) //read the data from body (json data) and store it in a structure variables
	http.HandleFunc("/dogheader", Readfromheader) //read value in a header and print it.
	http.ListenAndServe(":8080", nil)

}