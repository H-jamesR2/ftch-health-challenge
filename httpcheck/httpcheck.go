package httpcheck

import (
	"ftch-health-challenge/config"
	"ftch-health-challenge/util"
	//"log"
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// CheckEndpoint checks the status of a single endpoint
func CheckEndpoint(endpoint config.Endpoint, resultChan chan<- string) {
	client := &http.Client{
		Timeout: time.Second * 1, // Timeout of 1 second
	}

	var reqBody *bytes.Reader
	if endpoint.Body != "" {
		reqBody = bytes.NewReader([]byte(endpoint.Body))
	} else {
		reqBody = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(endpoint.Method, endpoint.URL, reqBody)
	if err != nil {
		util.LogError("CheckEndpoint: failed to create request", err)
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
	if err != nil {
		util.LogError("CheckEndpoint: request failed", err)
		resultChan <- endpoint.URL + "|DOWN"
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 || latency > 500 {
		util.LogInfo(fmt.Sprintf("CheckEndpoint: endpoint '%s' is DOWN (Status: %d, Latency: %dms)", endpoint.Name, resp.StatusCode, latency))
		resultChan <- endpoint.URL + "|DOWN"
	} else {
		util.LogInfo(fmt.Sprintf("CheckEndpoint: endpoint '%s' is UP (Status: %d, Latency: %dms)", endpoint.Name, resp.StatusCode, latency))
		resultChan <- endpoint.URL + "|UP"
	}
}