package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	finalURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := strings.TrimSuffix(strings.ToLower(finalURL.Host+finalURL.Path), "/")

	return fullPath, nil
}
