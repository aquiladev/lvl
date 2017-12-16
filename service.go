package main

import (
	"github.com/aquiladev/lvl/storage"
)

type StorageService interface {
	Put([]byte, []byte) error
	PutBatch([]storage.KeyValue) error
	Get([]byte) ([]byte, error)
}

type storageService struct {
	storage storage.KeyValueStorage
}

func newStorageService(storage storage.KeyValueStorage) *storageService {
	return &storageService{
		storage: storage,
	}
}

func (s *storageService) Put(k, v []byte) error {
	return s.storage.Put(k, v)
}

func (s *storageService) PutBatch(pairs []storage.KeyValue) error {
	return s.storage.PutBatch(pairs)
}

func (s *storageService) Get(k []byte) ([]byte, error) {
	return s.storage.Get(k)
}
