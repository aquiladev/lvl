package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   QueueService
}

func (mw loggingMiddleware) Put(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "put",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Put(s)
	return
}

func (mw loggingMiddleware) Pop(s string) (n string) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "pop",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Pop(s)
	return
}