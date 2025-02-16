package v1

import (
	"testing"
)

func TestValidateUser(t *testing.T) {
	_, err := NexusClient.ValidateUser()
	if err != nil {
		t.Fatal(err)
	}

	//util.PrettyPrint(user)
}

func TestGetEndorsements(t *testing.T) {
	_, err := NexusClient.GetEndorsements()
	if err != nil {
		t.Fatal(err)
	}

	//util.PrettyPrint(endorsements)
}

func TestGetTrackedMods(t *testing.T) {
	_, err := NexusClient.GetTrackedMods()
	if err != nil {
		t.Fatal(err)
	}

	//util.PrettyPrint(mods)
}
