package repository

import (
	"math/rand"
	"sync"
	"time"
)

var (
	m sync.RWMutex
)

type MockRepository struct {
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
func (repository *MockRepository) reader() {
	m.RLock()
	sleep()
	m.RUnlock()
}
func (repository *MockRepository) writer() {
	m.Lock()
	sleep()
	m.Unlock()
}
