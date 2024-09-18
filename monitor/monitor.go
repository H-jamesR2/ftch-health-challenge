package monitor

import (
	"ftch-health-challenge/config"
	"ftch-health-challenge/domain"
	"ftch-health-challenge/httpcheck"
	"log"
	"sync"
	"time"
)

func MonitorEndpoints(endpoints []config.Endpoint) {
	domainStats := make(map[string]*domain.DomainStats)
	for {
		resultChan := make(chan string, len(endpoints))
		var wg sync.WaitGroup

		// Check each endpoint concurrently
		for _, endpoint := range endpoints {
			wg.Add(1)
			go func(ep config.Endpoint) {
				defer wg.Done()
				httpcheck.CheckEndpoint(ep, resultChan)
			}(endpoint)
		}

		// Wait for all checks to complete
		go func() {
			wg.Wait()
			close(resultChan)
		}()

		// Update stats based on results
		for result := range resultChan {
			splitResult := splitResult(result)
			domainName, err := domain.GetDomain(splitResult[0])
			if err != nil {
				log.Println(err)
				continue
			}

			if domainStats[domainName] == nil {
				domainStats[domainName] = &domain.DomainStats{}
			}
			domainStats[domainName].Total++

			if splitResult[1] == "UP" {
				domainStats[domainName].Up++
			}
		}

		// Log the availability stats after each round
		domain.LogAvailability(domainStats)

		// Sleep for 15 seconds before next round
		time.Sleep(15 * time.Second)
	}
}

// Helper function to split result into URL and status (e.g., "URL|UP" -> [URL, UP])
func splitResult(result string) [2]string {
	var res [2]string
	for i, v := range result {
		if string(v) == "|" {
			res[0] = result[:i]
			res[1] = result[i+1:]
			break
		}
	}
	return res
}
