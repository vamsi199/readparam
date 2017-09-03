package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"


)



type dogs struct {
	ID string `json:"id"`
	Color string `json:"color"`
}

func readparameter(w http.ResponseWriter,req *http.Request)  {
	var dog dogs
	parameter := mux.Vars(req)
	dog.ID = parameter["id"]
	json.NewEncoder(w).Encode(dog.ID)
}
func Qweryread(w http.ResponseWriter,req *http.Request){
	var dog dogs
	f := req.URL.Query()
	dog.Color = f.Get("color")
	json.NewEncoder(w).Encode(dog.Color)

}





func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dogs/{id}",readparameter).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081",router))
	http.HandleFunc("/breed",Qweryread)
	http.ListenAndServe(":8082",nil)




}
