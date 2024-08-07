package v1

import (
	"testing"

	"github.com/the-egg-corp/gonexus/util"
)

func TestValidateUser(t *testing.T) {
	user, err := NexusClient.ValidateUser()
	if err != nil {
		t.Fatal(err)
	}

	util.PrettyPrint(user)
}
