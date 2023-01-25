package service

import (
	"fmt"
	"testing"
)

func TestGenerateLink(t *testing.T) {
	testCases := []struct {
		url           string
		expectedError error
	}{
		{"http://ya.com", nil},
		{"http://ozon.ru", nil},
		{"http://golang.org", nil},
		{"http://youtube.com", nil},
		{"", nil},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("URL: %s", tc.url), func(t *testing.T) {
			_, err := GenerateLink(tc.url)
			if err != tc.expectedError {
				t.Errorf("Expected error %v but got %v", tc.expectedError, err)
			}
		})
	}
}
