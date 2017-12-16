package main

import (
	"fmt"
	"time"

	"github.com/aquiladev/lvl/storage"
	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           StorageService
}

func (mw instrumentingMiddleware) Put(k, v []byte) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "put", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = mw.next.Put(k, v)
	return
}

func (mw instrumentingMiddleware) PutBatch(pairs []storage.KeyValue) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "put_batch", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = mw.next.PutBatch(pairs)
	return
}

func (mw instrumentingMiddleware) Get(k []byte) (v []byte, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	v, err = mw.next.Get(k)
	return
}
