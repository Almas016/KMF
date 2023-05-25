package cmd

import (
	"KMF/handlers"
	"KMF/repo"
	"KMF/usecases"
	"log"
	"net/http"
)

// Start application
func Start(config *Config) {
	requestRepository := repo.NewRequestRepository()
	responseRepository := repo.NewResponseRepository()
	usecase := usecases.NewUseCase(requestRepository, responseRepository)
	handler := handlers.NewHandler(usecase)

	http.HandleFunc("/", handler.Handle)

	log.Println("Start server on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
