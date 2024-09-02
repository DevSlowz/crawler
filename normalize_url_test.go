package main

import (
	"testing"
)

// Needs to accept URL returns normalized URL
//
//
//

func TestNormalizeURL(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		inputURL string
		want     string
	}

	testCases := []testCase{
		{name: "Remove Scheme", inputURL: "https://blog.boot.dev/path", want: "blog.boot.dev/path"},
	}

	for _, tc := range testCases {
		got := normalizeURL(tc.inputURL)

		if tc.want != got {
			t.Errorf("normalizeURL(%s): want %s, got %s", tc.inputURL, tc.want, got)
		}
	}
}
