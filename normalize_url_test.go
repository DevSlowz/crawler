package main

import (
	"strings"
	"testing"
)

// Needs to accept URL returns normalized URL
//
//
//

func TestNormalizeURL(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name          string
		inputURL      string
		want          string
		errorContains string
	}

	testCases := []testCase{
		{name: "Remove Scheme", inputURL: "https://blog.boot.dev/path", want: "blog.boot.dev/path"},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			want:     "blog.boot.dev/path",
		},
		{
			name:     "lowercase capital letters",
			inputURL: "https://BLOG.boot.dev/PATH",
			want:     "blog.boot.dev/path",
		},
		{
			name:     "remove scheme and capitals and trailing slash",
			inputURL: "http://BLOG.boot.dev/path/",
			want:     "blog.boot.dev/path",
		},
		{
			name:          "handle invalid URL",
			inputURL:      `:\\invalidURL`,
			want:          "",
			errorContains: "couldn't parse URL",
		},
	}

	// t.Run lets the test case run in parallel
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.want {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.want, actual)
			}
		})
	}
}
