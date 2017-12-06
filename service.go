package main

type QueueService interface {
	Put(string) (string, error)
	Pop(string) string
}

type queueService struct{}

func (queueService) Put(s string) (string, error) {
	return "", nil
}

func (queueService) Pop(s string) string {
	return ""
}
