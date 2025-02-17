package v1

import "testing"

func TestGetColourSchemes(t *testing.T) {
	colourSchemes, err := NexusClient.GetColourSchemes()
	if err != nil {
		t.Fatal("failed to get colour schemes\n", err)
	}

	if colourSchemes == nil {
		t.Fatal("failed to get colour schemes. response was nil but no error was provided")
	}
}
