package main

import (
	"net/http"
	"os"

	"github.com/aquiladev/lvl/storage"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "storage",
		Subsystem: "storage_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "storage",
		Subsystem: "storage_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	db, err := leveldb.OpenFile("d:/lvl/data", nil)
	if err != nil {
		panic(err)
	}

	var svc StorageService
	svc = newStorageService(storage.NewLevelDbStorage(db))
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, svc}

	putHandler := httptransport.NewServer(
		makePutEndpoint(svc),
		decodePutRequest,
		encodeResponse,
	)

	putBatchHandler := httptransport.NewServer(
		makePutBatchEndpoint(svc),
		decodePutBatchRequest,
		encodeResponse,
	)

	getHandler := httptransport.NewServer(
		makeGetEndpoint(svc),
		decodeGetRequest,
		encodeResponse,
	)

	http.Handle("/put", putHandler)
	http.Handle("/put-batch", putBatchHandler)
	http.Handle("/get", getHandler)

	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
