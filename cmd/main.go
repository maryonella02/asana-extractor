package main

import (
	"asana-extractor/cmd/config"
	"asana-extractor/internal"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {

		log.Fatalf("could not load configs")
	}

	// use later cfg

	asanaClient := internal.NewClient(cfg)

	err = asanaClient.ExportUsersToFile()
	if err != nil {
		fmt.Errorf("error: %w", err)
		return
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down...")
}
