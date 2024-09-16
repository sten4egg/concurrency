package internal

import (
	"errors"
	"sync"
)

var errNotFound = errors.New("key not found")

type Engine struct {
	mu      sync.Mutex
	storage map[string]string
}

func NewEngine() *Engine {
	return &Engine{
		mu:      sync.Mutex{},
		storage: make(map[string]string),
	}
}

func (e *Engine) Get(key string) (string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	val, ok := e.storage[key]
	if !ok {
		return "", errNotFound
	}
	return val, nil
}

func (e *Engine) Set(key, value string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.storage[key] = value
}

func (e *Engine) Del(key string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.storage, key)
}
