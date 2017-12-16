package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"
)

func TestPutGet(t *testing.T) {
	db, err := leveldb.OpenFile("D:/lvl-test", nil)
	assert.Equal(t, nil, err)

	storage := NewLevelDbStorage(db)
	err = storage.Put([]byte("testKey"), []byte("testValue"))
	assert.Equal(t, nil, err)

	v, err := storage.Get([]byte("testKey"))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("testValue"), v)
}

func TestPutBatchGet(t *testing.T) {
	db, err := leveldb.OpenFile("D:/lvl-batch-test", nil)
	assert.Equal(t, nil, err)

	storage := NewLevelDbStorage(db)
	list := make([]KeyValue, 2)
	list[0] = KeyValue{k: []byte("testK1"), v: []byte("testV1")}
	list[1] = KeyValue{k: []byte("testK2"), v: []byte("testV2")}

	err = storage.PutBatch(list)
	assert.Equal(t, nil, err)

	v, err := storage.Get([]byte("testK1"))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("testV1"), v)

	v, err = storage.Get([]byte("testK2"))
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("testV2"), v)
}
