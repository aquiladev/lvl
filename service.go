package main

import (
	"github.com/aquiladev/lvl/storage"
	"github.com/shirou/gopsutil/disk"
	"github.com/pkg/errors"
)

type StorageService interface {
	Put([]byte, []byte) error
	PutBatch([]storage.KeyValue) error
	Get([]byte) ([]byte, error)
}

type storageService struct {
	storage      storage.KeyValueStorage
	dataDir      string
	maxDiskUsage float64
}

func newStorageService(storage storage.KeyValueStorage, dataDir string, maxDiskUsage float64) *storageService {
	return &storageService{
		storage:      storage,
		dataDir:      dataDir,
		maxDiskUsage: maxDiskUsage,
	}
}

func (s *storageService) Put(k, v []byte) error {
	if err := s.checkDiskUsage(); err != nil {
		return err
	}
	return s.storage.Put(k, v)
}

func (s *storageService) PutBatch(pairs []storage.KeyValue) error {
	if err := s.checkDiskUsage(); err != nil {
		return err
	}
	return s.storage.PutBatch(pairs)
}

func (s *storageService) Get(k []byte) ([]byte, error) {
	return s.storage.Get(k)
}

func (s *storageService) checkDiskUsage() error {
	usage, err := disk.Usage(s.dataDir)
	if err != nil {
		return err
	}

	if usage.UsedPercent >= s.maxDiskUsage {
		return errors.New("reached max disk usage")
	}

	return nil
}
