// config/config.go

package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Endpoint defines the structure for each HTTP endpoint from the YAML configuration
type Endpoint struct {
	Name    string            `yaml:"name"`
	URL     string            `yaml:"url"`
	Method  string            `yaml:"method"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}

// LoadConfig loads and parses the YAML configuration file with validation.
// It returns a slice of valid endpoints and reports any invalid configurations.
func LoadConfig(filepath string) ([]Endpoint, error) {
	var allEndpoints []Endpoint
	var validEndpoints []Endpoint
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("LoadConfig: failed to read file %s: %w", filepath, err)
	}
	err = yaml.Unmarshal(data, &allEndpoints)
	if err != nil {
		return nil, fmt.Errorf("LoadConfig: failed to parse YAML: %w", err)
	}

	// Validate each endpoint
	for i, endpoint := range allEndpoints {
		if endpoint.Name == "" {
			log.Printf("LoadConfig: skipping endpoint at index %d due to missing 'name'\n", i)
			continue
		}
		if endpoint.URL == "" {
			log.Printf("LoadConfig: skipping endpoint '%s' due to missing 'url'\n", endpoint.Name)
			continue
		}
		if endpoint.Method != "" {
			validMethods := map[string]bool{
				"GET":     true,
				"POST":    true,
				"PUT":     true,
				"DELETE":  true,
				"PATCH":   true,
				"HEAD":    true,
				"OPTIONS": true,
			}
			method := endpoint.Method
			if _, exists := validMethods[method]; !exists {
				log.Printf("LoadConfig: skipping endpoint '%s' due to invalid HTTP method '%s'\n", endpoint.Name, method)
				continue
			}
		}
		// Optionally, validate that 'body' is valid JSON if present
		if endpoint.Body != "" {
			if !isValidJSON(endpoint.Body) {
				log.Printf("LoadConfig: skipping endpoint '%s' due to invalid JSON in 'body'\n", endpoint.Name)
				continue
			}
		}
		// If all validations pass, add to validEndpoints
		validEndpoints = append(validEndpoints, endpoint)
	}

	if len(validEndpoints) == 0 {
		return nil, errors.New("LoadConfig: no valid endpoints found in configuration")
	}

	return validEndpoints, nil
}

// isValidJSON checks if a string is valid JSON
func isValidJSON(s string) bool {
	var js interface{}
	return yaml.Unmarshal([]byte(s), &js) == nil
}
