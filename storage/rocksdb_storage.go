package storage

//import "github.com/tecbot/gorocksdb"
//
//type RocksDbStorage struct {
//	db *gorocksdb.DB
//	wo *gorocksdb.WriteOptions
//}
//
//func NewRocksDbStorage(db *gorocksdb.DB) *RocksDbStorage {
//	return &RocksDbStorage{
//		db: db,
//		wo: gorocksdb.NewDefaultWriteOptions(),
//	}
//}
//
//func (s *RocksDbStorage) Put(k []byte, v []byte) error {
//	return s.db.Put(s.wo, k, v)
//}
//
//func (s *RocksDbStorage) PutBatch(pairs []KeyValue) error {
//	wo := gorocksdb.NewDefaultWriteOptions()
//	wb := gorocksdb.NewWriteBatch()
//	defer wb.Clear()
//
//	for _, p := range pairs {
//		wb.Put(p.k, p.v)
//	}
//
//	return s.db.Write(wo, wb)
//}
