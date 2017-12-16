package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDbStorage struct {
	db *leveldb.DB
}

func NewLevelDbStorage(db *leveldb.DB) *LevelDbStorage {
	return &LevelDbStorage{
		db: db,
	}
}

func (s *LevelDbStorage) Put(k, v []byte) error {
	return s.db.Put(k, v, nil)
}

func (s *LevelDbStorage) PutBatch(pairs []KeyValue) error {
	batch := new(leveldb.Batch)
	for _, p := range pairs {
		batch.Put(p.K, p.V)
	}

	return s.db.Write(batch, nil)
}

func (s *LevelDbStorage) Get(k []byte) ([]byte, error) {
	return s.db.Get(k, nil)
}
