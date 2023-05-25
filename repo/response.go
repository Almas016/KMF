package repo

import (
	"KMF/entities"
	"sync"
)

type ResponseRepository struct {
	mutex     sync.Mutex
	responses map[string]*entities.Response
}

func NewResponseRepository() *ResponseRepository {
	return &ResponseRepository{
		responses: make(map[string]*entities.Response),
	}
}

func (rr *ResponseRepository) StoreResponse(response *entities.Response) {
	rr.mutex.Lock()
	defer rr.mutex.Unlock()

	// Store the response
	rr.responses[response.ID] = response
}
