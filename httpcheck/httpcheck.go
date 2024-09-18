package httpcheck

import (
	"ftch-health-challenge/config"
	"log"
	"net/http"
	"time"
)

func CheckEndpoint(endpoint config.Endpoint, resultChan chan<- string) {
	client := &http.Client{
		Timeout: time.Second * 1, // Timeout of 1 second
	}

	req, err := http.NewRequest(endpoint.Method, endpoint.URL, nil)
	if err != nil {
		log.Println(err)
		resultChan <- endpoint.URL + "|DOWN"
		return
	}

	// Add headers to the request if they exist
	for key, value := range endpoint.Headers {
		req.Header.Add(key, value)
	}

	start := time.Now()
	resp, err := client.Do(req)
	latency := time.Since(start).Milliseconds()

	// Check for success criteria: 2xx status code and < 500ms latency
	if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 || latency > 500 {
		resultChan <- endpoint.URL + "|DOWN"
	} else {
		resultChan <- endpoint.URL + "|UP"
	}
}
