package cmd

import "log"

func main() {

	_, err := LoadConfig()
	if err != nil {

		log.Fatalf("could not load configs")
	}

	// use later cfg
}
