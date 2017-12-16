package main

import (
	"time"

	"github.com/aquiladev/lvl/storage"
	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   StorageService
}

func (mw loggingMiddleware) Put(k, v []byte) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "put",
			"key", k,
			"value", v,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.Put(k, v)
	return
}

func (mw loggingMiddleware) PutBatch(pairs []storage.KeyValue) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "put_batch",
			"pairs_count", len(pairs),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.next.PutBatch(pairs)
	return
}

func (mw loggingMiddleware) Get(k []byte) (v []byte, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "get",
			"key", k,
			"value", v,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	v, err = mw.next.Get(k)
	return
}
