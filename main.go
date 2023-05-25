package main

import (
	"KMF/cmd"
	"log"
)

func main() {
	// Load configuration
	config, err := cmd.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Start application
	cmd.Start(config)
}
