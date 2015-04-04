package fakku

import (
	"fmt"
	"net/url"
	"time"
)

type userAchievementsApiFunction struct {
	userApiFunction
}

func (c userAchievementsApiFunction) Construct() string {
	return fmt.Sprintf("%s/achievements", c.userApiFunction.Construct())
}

func GetUserAchievements(user string) (UserAchievementList, error) {
	var c userAchievements
	url := userAchievementsApiFunction{
		userApiFunction: userApiFunction{Name: user},
	}
	if err := apiCall(url, &c); err != nil {
		return nil, err
	} else {
		if c.Total != uint(len(c.Achievements)) {
			return nil, fmt.Errorf("Achievement mismatch detected! Expected %d but got %d instead", c.Total, len(c.Achievements))
		} else {
			return c.Achievements, nil
		}
	}
}

type UserAchievementList []UserAchievement
type userAchievements struct {
	Achievements UserAchievementList `json:"achievements"`
	Total        uint                `json:"total"`
}
type UserAchievement struct {
	Name        string `json:"achievement_name"`
	Description string `json:"achievement_description"`
	RawIcon     string `json:"achievement_icon"`
	Class       string `json:"achievement_class"`
	RawDate     int64  `json:"achievement_date"`
}

func (this *UserAchievement) Date() time.Time {
	return time.Unix(this.RawDate, 0)
}
func (this *UserAchievement) IconUrl() (*url.URL, error) {
	return url.Parse(this.RawIcon)
}

func (this *UserAchievement) String() string {
	return fmt.Sprintf("%s - %s", this.Name, this.Description)
}

//func (this *UserProfile) Achievements() (
