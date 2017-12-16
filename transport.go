package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/aquiladev/lvl/storage"
)

type KeyValue struct {
	K string
	V string
}

func makePutEndpoint(svc StorageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(putRequest)
		err := svc.Put([]byte(req.K), []byte(req.V))
		if err != nil {
			return putResponse{err.Error()}, nil
		}
		return putResponse{""}, nil
	}
}

func makePutBatchEndpoint(svc StorageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(putBatchRequest)

		pairs := make([]storage.KeyValue, len(req.P))
		for i, j := range req.P {
			pairs[i] = storage.KeyValue{
				V: []byte(j.V),
				K: []byte(j.K),
			}
		}

		err := svc.PutBatch(pairs)
		if err != nil {
			return putBatchResponse{err.Error()}, nil
		}
		return putBatchResponse{""}, nil
	}
}

func makeGetEndpoint(svc StorageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		v, err := svc.Get([]byte(req.K))
		if err != nil {
			return getResponse{"", err.Error()}, nil
		}
		return getResponse{string(v), ""}, nil
	}
}

func decodePutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request putRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodePutBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request putBatchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type putRequest struct {
	K string `json:"k"`
	V string `json:"v"`
}

type putResponse struct {
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type putBatchRequest struct {
	P []KeyValue `json:"p"`
}

type putBatchResponse struct {
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type getRequest struct {
	K string `json:"k"`
}

type getResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
