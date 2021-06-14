package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Details struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var allData []Details

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var formData Details
	_ = json.NewDecoder(r.Body).Decode(&formData)
	allData = append(allData, formData)
	json.NewEncoder(w).Encode(&formData)
	fmt.Println("Got THE form: ", formData)
	fmt.Println("Got ALL forms: ", allData)
	fmt.Fprintf(os.Stdout, "Employee is %#v\n", formData)

}
func main() {
	// myRouter := mux.NewRouter()
	// myRouter.HandleFunc("/login", loginHandler).Methods("POST")
	http.HandleFunc("/login", loginHandler)

	fmt.Println("server is running")
	// log.Fatal(http.ListenAndServe(":10000", myRouter))
	log.Fatal(http.ListenAndServe(":10000", nil))

}
