package api

import "testing"

func TestGetURLPathWith(t *testing.T) {
	path := "/api/v1/upload/"
	result := getURLPath(path)

	if result != "/api/v1/upload" {
		t.Errorf("getURLPath returned %s", result)
	}
}

func TestGetURLPathWithEmptyPath(t *testing.T) {
	path := "/api/v1/upload"
	result := getURLPath(path)

	if result != "/api/v1/upload" {
		t.Errorf("getURLPath returned %s", result)
	}
}
