package main

import (
	"fmt"
	"github.com/JuanValero25/golangBack/repository"
	"net/http"
)

var (
	repositori = repository.MockRepository{}
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func readAllTransaction(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, repositori)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":10000", nil)
	//log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
