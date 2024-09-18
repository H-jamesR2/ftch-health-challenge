package main

import (
	"fmt"
	"ftch-health-challenge/config"
	"ftch-health-challenge/monitor"
	"log"
	"os"
)

func main() {
	// Configure the logger to include date, time, and short file name
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Ensure filepath argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <config_file.yaml>")
		return
	}
	configFile := os.Args[1]

	// Load the YAML configuration
	endpoints, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("main: failed to load configuration: %v", err)
	}

	// Default to GET method if not provided
	for i := range endpoints {
		if endpoints[i].Method == "" {
			endpoints[i].Method = "GET"
		}
	}

	// Start monitoring the endpoints
	monitor.MonitorEndpoints(endpoints)
}
