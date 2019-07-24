package main

import (
	"awesomeProject/repository"
	"fmt"
	"net/http"
)

var( r  repository.MockRepository)

func homePage(w http.ResponseWriter, r *http.Request){
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
