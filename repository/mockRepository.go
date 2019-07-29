package repository

import (
	"fmt"
	"github.com/google/uuid"
	"math/big"
	"sync"
	"time"
)

type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

type Transaction struct {
	ID            string          `json:"id"`
	Type          TransactionType `json:"type"`
	Amount        big.Float       `json:"amount"`
	EffectiveDate time.Time       `json:"effectiveDate"`
}

type TrasactionError struct {
	Message string `json:"Message"`
}

func (t *TrasactionError) Error() string {
	return "Error Message : " + t.Message
}

var (
	mutex             sync.RWMutex
	mockedDBMap       = make(map[string]*Transaction)
	allTransaction    []*Transaction
	sumAllTransaction big.Float
)

type MockRepository struct {
}

func (c *MockRepository) GetTransactionById(trasanctionId string) (*Transaction, error) {
	if !IsValidUUID(trasanctionId) {
		return nil, &TrasactionError{"id is not valid"}
	}
	mutex.Lock()
	defer mutex.Unlock()
	transactionFound := mockedDBMap[trasanctionId]
	return transactionFound, nil
}

func (c *MockRepository) GetAllTransaction() *[]*Transaction {
	mutex.Lock()
	defer mutex.Unlock()
	return &allTransaction
}

func (c *MockRepository) PostTransaction(transaction *Transaction) error {
	if IsInValidTransaction(transaction) {
		return &TrasactionError{Message: "invalid transaction"}
	}
	mutex.Lock()
	defer mutex.Unlock()
	transaction.EffectiveDate = time.Now()
	transaction.ID = uuid.New().String()
	mockedDBMap[transaction.ID] = transaction
	allTransaction = append(allTransaction, transaction)
	SumTransaction(transaction)
	return nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func SumTransaction(transaction *Transaction) {

	switch transaction.Type {
	case Debit:
		sumAllTransaction = *sumAllTransaction.Sub(&sumAllTransaction, &transaction.Amount)

	case Credit:
		sumAllTransaction = *sumAllTransaction.Add(&sumAllTransaction, &transaction.Amount)
	}

}

func IsInValidTransaction(transactionInsert *Transaction) bool {
	amountValue , _  := transactionInsert.Amount.Uint64()
	fmt.Print("the precion amount is ", amountValue)
	if amountValue == 0 {
		return true
	}

	if transactionInsert.Type != Credit || transactionInsert.Type != Debit {
		return true
	}
	return false
}
