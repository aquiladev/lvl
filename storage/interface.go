package storage

type KeyValue struct {
	K []byte
	V []byte
}

type KeyValueStorage interface {
	Put([]byte, []byte) error
	PutBatch([]KeyValue) error
	Get([]byte) ([]byte, error)
}
