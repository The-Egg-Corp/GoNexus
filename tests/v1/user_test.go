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
