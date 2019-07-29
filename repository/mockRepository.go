package repository

import (
	"github.com/google/uuid"
	"math/big"
	"sync"
	"time"
)

type Transaction struct {
	ID            string    `json:"id"`
	Type          string    `json:"type"`
	Amount        big.Rat   `json:"amount"`
	EffectiveDate time.Time `json:"effectiveDate"`
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
	sumAllTransaction = big.NewRat(0, 0)
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
	sumAllTransaction.Add(sumAllTransaction, &transaction.Amount)
	return nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func IsInValidTransaction(transactionInsert *Transaction) bool {
	if transactionInsert.Amount.Num().Int64() == 0 {
		return true
	}

	if len(transactionInsert.Type) > 1 {
		return true
	}
	return false
}
