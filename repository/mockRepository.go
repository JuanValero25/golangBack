package repository

import (
	"math/rand"
	"sync"
	"time"
)

type Transaction struct {
	ID            string    `json:"id"`
	Type          string    `json:"type"`
	Amount        int       `json:"amount"`
	EffectiveDate time.Time `json:"effectiveDate"`
}

var (
	m              sync.RWMutex
	mapTransaction = make(map[string]*Transaction)
	values         []*Transaction
)

type MockRepository struct {
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
func (repository *MockRepository) reader(trasanctionId *string) Transaction {
	m.RLock()
	sleep()
	transactionFound := mapTransaction[*trasanctionId]
	m.RUnlock()
	return *transactionFound
}

func (repository *MockRepository) readerAll() *[]*Transaction {
	return &values
}

func (repository *MockRepository) writer(transactionInsert *Transaction) {
	m.Lock()
	sleep()
	mapTransaction[transactionInsert.ID] = transactionInsert
	values = append(values, transactionInsert)
	m.Unlock()
}
