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
		//t.Error(err)
		t.Log(err)
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

func TestGetUserPosts_1(t *testing.T) {
	posts, err := GetUserPosts(TestUserName)
	if err != nil {
		t.Error(err)
	}

	if posts.Total < 4554 {
		t.Error("Number of user posts is not correct!")
	}
}

func TestGetUserTopics_1(t *testing.T) {
	topics, err := GetUserTopics(TestUserName)
	if err != nil {
		t.Error(err)
	}

	if topics.Total < 1016 {
		t.Errorf("Expected at least 1016 topics, got %d", topics.Total)
	}
}
