package fakku

import (
	"fmt"
)

type UserApiFunction struct {
	Name string
}

func (c UserApiFunction) ConstructApiFunction() string {
	return fmt.Sprintf("%s/users/%s", ApiHeader, c.Name)
}

func GetUserProfile(name string) (*UserProfile, error) {
	var c User
	url := UserApiFunction{Name: name}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		//cheat a little bit :D
		return &(*c.Profile), nil
	}
}

type User struct {
	Profile *UserProfile `json:"user"`
}

// DAMN, I can't just have Go convert these to bools for me.
// I'll need to do the conversion manually
type UserProfile struct {
	Username          string `json:"user_username"`
	Url               string `json:"user_url"`
	Rank              string `json:"user_rank"`
	Avatar            string `json:"user_avatar"`
	AvatarWidth       uint   `json:"user_avatar_width"`
	AvatarHeight      uint   `json:"user_avatar_height"`
	RegistrationDate  uint   `json:"user_registration_date"`
	LastVisit         uint   `json:"user_last_visit"`
	Subscribed        uint   `json:"user_subscribed"`
	Timezone          int    `json:"user_timezone"`
	Posts             uint   `json:"user_posts"`
	Topics            uint   `json:"user_topics"`
	Comments          uint   `json:"user_comments"`
	Signature         string `json:"user_signature"`
	ForumReputation   int    `json:"user_forum_reputation"`
	CommentReputation int    `json:"user_comment_reputation"`
	Gold              uint   `json:"user_gold"`
	Online            uint   `json:"user_online"`
}

type UserFavoritesApiFunction struct {
	UserApiFunction
	SupportsPagination
}

func (c UserFavoritesApiFunction) ConstructApiFunction() string {
	base := fmt.Sprintf("%s/favorites", c.UserApiFunction.ConstructApiFunction())
	return PaginateString(base, c.Page)
}

func GetUserFavoritesPage(user string, page uint) (*UserFavorites, error) {
	var c UserFavorites
	url := UserFavoritesApiFunction{
		UserApiFunction:    UserApiFunction{Name: user},
		SupportsPagination: SupportsPagination{Page: page},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

func GetUserFavorites(user string) (*UserFavorites, error) {
	return GetUserFavoritesPage(user, 0)
}

type UserFavorites struct {
	Favorites []*Content `json:"favorites"`
	Total     uint       `json:"total"`
	Pages     uint       `json:"pages"`
}

type UserAchievementsApiFunction struct {
	UserApiFunction
}

func (c UserAchievementsApiFunction) ConstructApiFunction() string {
	return fmt.Sprintf("%s/achievements", c.UserApiFunction.ConstructApiFunction())
}

func GetUserAchievements(user string) (*UserAchievements, error) {
	var c UserAchievements
	url := UserAchievementsApiFunction{
		UserApiFunction: UserApiFunction{Name: user},
	}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type UserAchievements struct {
	Achievements []*UserAchievement `json:"achievements"`
	Total        uint               `json:"total"`
}
type UserAchievement struct {
	Name        string `json:"achievement_name"`
	Description string `json:"achievement_description"`
	Icon        string `json:"achievement_icon"`
	Class       string `json:"achievement_class"`
	Date        uint   `json:"achievement_date"`
}
