package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           QueueService
}

func (mw instrumentingMiddleware) Put(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "put", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Put(s)
	return
}

func (mw instrumentingMiddleware) Pop(s string) (n string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "pop", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	n = mw.next.Pop(s)
	return
}