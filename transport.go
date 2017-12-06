package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makePutEndpoint(svc QueueService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(putRequest)
		v, err := svc.Put(req.S)
		if err != nil {
			return putResponse{v, err.Error()}, nil
		}
		return putResponse{v, ""}, nil
	}
}

func makePopEndpoint(svc QueueService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(popRequest)
		v := svc.Pop(req.S)
		return popResponse{v}, nil
	}
}

func decodePutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request putRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodePopRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request popRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type putRequest struct {
	S string `json:"s"`
}

type putResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type popRequest struct {
	S string `json:"s"`
}

type popResponse struct {
	V string `json:"v"`
}
