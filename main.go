package main

import (
	"encoding/json"
	"fmt"
	"github.com/JuanValero25/golangBack/repository"
	"net/http"
	"strings"
)

var (
	mockRepository = repository.MockRepository{}
)

func readAllTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(mockRepository.GetAllTransaction())

}

func transactionWriterHandler(w http.ResponseWriter, r *http.Request) {
	var transaction repository.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		fmt.Print("error with json decoding : ", err)
	}
	mockRepository.PostTransaction(&transaction)
}

func transactionReaderHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/transaction/")
	fmt.Print(id)
	transaction, err := mockRepository.GetTransactionById(id)
	if err == nil {

	} else {
		_ = json.NewEncoder(w).Encode(transaction)
	}

}

func handleRequests() {
	http.HandleFunc("/all", readAllTransaction)
	http.HandleFunc("/transaction", transactionWriterHandler)
	http.HandleFunc("/transaction/", transactionReaderHandler)
	_ = http.ListenAndServe(":10000", nil)
}

func main() {
	handleRequests()
}
