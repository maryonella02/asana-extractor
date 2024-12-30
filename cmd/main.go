package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	_, err := LoadConfig()
	if err != nil {

		log.Fatalf("could not load configs")
	}

	// use later cfg

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down...")
}
