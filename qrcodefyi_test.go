package qrcodefyi_test

import (
	"testing"

	qrcodefyi "github.com/fyipedia/qrcodefyi-go"
)

func TestNewClient(t *testing.T) {
	c := qrcodefyi.NewClient()
	if c.BaseURL != qrcodefyi.DefaultBaseURL {
		t.Errorf("expected %s, got %s", qrcodefyi.DefaultBaseURL, c.BaseURL)
	}
	if c.HTTPClient == nil {
		t.Error("expected non-nil HTTPClient")
	}
}

func TestSearch(t *testing.T) {
	c := qrcodefyi.NewClient()
	result, err := c.Search("micro")
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}
	if result.Total == 0 {
		t.Error("expected results, got 0")
	}
}
