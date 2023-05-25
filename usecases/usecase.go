package usecases

import (
	"KMF/entities"
	"KMF/repo"
	"github.com/google/uuid"
)

type UseCase struct {
	requestRepository  *repo.RequestRepository
	responseRepository *repo.ResponseRepository
}

func NewUseCase(requestRepository *repo.RequestRepository, responseRepository *repo.ResponseRepository) *UseCase {
	return &UseCase{
		requestRepository:  requestRepository,
		responseRepository: responseRepository,
	}
}

func (uc *UseCase) Process(request *entities.Request) *entities.Response {
	requestID := uuid.New().String()

	// Perform necessary processing and interaction with external services
	// ...

	uc.requestRepository.StoreRequest(request, &requestID)

	response := &entities.Response{
		ID:      requestID,
		Status:  200,
		Headers: make(map[string]string),
		Length:  0, // Update with actual length from the response
	}

	uc.responseRepository.StoreResponse(response)

	return response
}
