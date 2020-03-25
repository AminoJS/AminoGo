package test

import (
	"github.com/AminoJS/AminoGo/routes"
	"testing"
)

func TestMainEndPoint(t *testing.T) {
	expectedEndpointValue := "https://service.narvii.com"
	if routes.ENDPOINT != expectedEndpointValue {
		t.Errorf("Expecting main endpoint to be %s, but got %v", expectedEndpointValue, routes.ENDPOINT)
	}
}
