package domain

import (
	"net/url"
	"fmt"
)

type DomainStats struct {
	Up    int
	Total int
}

func GetDomain(endpointURL string) (string, error) {
	u, err := url.Parse(endpointURL)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}

func LogAvailability(domainStats map[string]*DomainStats) {
	for domain, stats := range domainStats {
		percentage := 0
		if stats.Total > 0 {
			percentage = (stats.Up * 100) / stats.Total
		}
		fmt.Printf("%s has %d%% availability percentage\n", domain, percentage)
	}
}
