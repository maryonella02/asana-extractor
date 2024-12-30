package main

import (
	"asana-extractor/cmd/config"
	"asana-extractor/internal"
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

	asanaClient := internal.NewClient(cfg)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			err = asanaClient.ExportUsersToFile()
			if err != nil {
				fmt.Errorf("error: %w", err)
			}

			err = asanaClient.ExportProjectsToFile()
			if err != nil {
				fmt.Errorf("error: %w", err)
			}
		}
	}()

	<-sigChan
	log.Println("Shutting down...")
}
