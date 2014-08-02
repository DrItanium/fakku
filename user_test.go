package fakku

import (
	"testing"
)

const (
	TestUserName        = "jacob" // The man, the legend
	TestUserDisplayName = "Jacob"
	TestAchievementName = "FAKKU Gold"
)

func TestGetUserProfile_1(t *testing.T) {
	profile, err := GetUserProfile(TestUserName)
	if err != nil {
		t.Error(err)
	}

	if profile.Username != TestUserDisplayName {
		t.Error("Couldn't get Jacob's profile! Not good!")
	}
	t.Logf("Name: %s", profile.Username)
	t.Logf("Rank: %s", profile.Rank)
}

func TestGetUserFavorites_1(t *testing.T) {
	_, err := GetUserFavorites(TestUserName)
	// currently this isn't working....
	if err != nil {
		t.Error(err)
	}
	//TODO: put more elements here
}

func TestGetUserAchievements_1(t *testing.T) {
	achievements, err := GetUserAchievements(TestUserName)
	foundTargetAchievement := false
	if err != nil {
		t.Error(err)
	}
	for i := uint(0); i < achievements.Total; i++ {
		tmp := achievements.Achievements[i]
		t.Logf("Achievement: %s - %s", tmp.Name, tmp.Description)
		if tmp.Name == TestAchievementName {
			foundTargetAchievement = true
		}
	}
	if !foundTargetAchievement {
		t.Errorf("Didn't find achievement: %s!", TestAchievementName)
	}
}
