package repo

import (
	"KMF/entities"
	"sync"
)

type RequestRepository struct {
	mutex    sync.Mutex
	requests map[string]*entities.Request
}

func NewRequestRepository() *RequestRepository {
	return &RequestRepository{
		requests: make(map[string]*entities.Request),
	}
}

func (rr *RequestRepository) StoreRequest(request *entities.Request, requestID *string) {
	rr.mutex.Lock()
	defer rr.mutex.Unlock()

	rr.requests[*requestID] = request
}
