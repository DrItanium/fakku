package fakku

import (
	"testing"
)

func TestGetUserProfile_1(t *testing.T) {
	profile, err := GetUserProfile("jacob")
	if err != nil {
		t.Error(err)
	}

	if profile.Username != "Jacob" {
		t.Error("Couldn't get Jacob's profile! Not good!")
	}
	t.Logf("Name: %s", profile.Username)
	t.Logf("Rank: %s", profile.Rank)
}
