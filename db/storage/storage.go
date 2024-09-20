package storage

import "concurrency/db/storage/internal"

type Storage interface {
	Get(key string) (string, error)
	Set(key, value string)
	Del(key string)
}

func NewStorage() Storage {
	return internal.NewEngine()
}
