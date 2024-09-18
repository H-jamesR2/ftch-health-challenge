// domain/domain.go

package domain

import (
	"fmt"
	"net/url"
)

// DomainStats keeps track of availability statistics for each domain
type DomainStats struct {
	Up    int
	Total int
}

// GetDomain extracts the domain from a given URL
func GetDomain(endpointURL string) (string, error) {
	u, err := url.Parse(endpointURL)
	if err != nil {
		return "", fmt.Errorf("GetDomain: failed to parse URL '%s': %w", endpointURL, err)
	}
	return u.Hostname(), nil
}

// LogAvailability calculates and prints the availability percentage for each domain
func LogAvailability(domainStats map[string]*DomainStats) {
	fmt.Printf("- Availability Percentages:\n")
	for domain, stats := range domainStats {
		percentage := 0
		if stats.Total > 0 {
			percentage = (stats.Up * 100) / stats.Total
		}
		fmt.Printf("    %s has %d%% availability percentage\n", domain, percentage)
	}
	fmt.Printf("\n\n")
}
