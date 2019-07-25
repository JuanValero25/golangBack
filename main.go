package main

import (
	"fmt"
	"net/http"
	"time"
)

//var( r  repository.MockRepository)

type Transaction struct {
	ID            string    `json:"id"`
	Type          string    `json:"type"`
	Amount        int       `json:"amount"`
	EffectiveDate time.Time `json:"effectiveDate"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":10000", nil)
	//log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
