package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Details struct {
	Email    string
	Password string
}

var allData []Details

func processPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processPostHandler is running")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	var formData Details
	json.Unmarshal(reqBody, &formData)
	allData = append(allData, formData)
	json.NewEncoder(w).Encode(formData)
	json.NewEncoder(w).Encode(allData)
	fmt.Println(formData)
	fmt.Println(allData)
}
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/processpost", processPostHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
