package utils

import (
	"log"
	"net/url"
)

// ValidateURL validates whether the provided URL is allowed
func ValidateURL(check string, baseURL *url.URL) (bool, *url.URL) {
	p, err := url.Parse(check)
	if err != nil {
		log.Printf("Cannot parse URL: %s", check)
		return false, p
	}

	p.Fragment = ""

	if p.Host == baseURL.Host {
		return true, p
	} else if p.Host == "" {
		p.Scheme = baseURL.Scheme
		p.Host = baseURL.Host
		return true, p
	} else {
		return false, p
	}
}
