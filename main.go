package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Details struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var allData []Details

// json, _ := ioutil.ReadAll(r.Body)
// fmt.Println("Got params: ", string(json))
func loginHandler(w http.ResponseWriter, r *http.Request) {
	formData := &Details{}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(formData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(*formData)
	allData = append(allData, *formData)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(formData); err != nil {
		log.Printf("Can't encode %v - %s", w, err)
	}
	fmt.Println(allData)
}
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/login", loginHandler).Methods("POST")
	fmt.Println("server is running")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
