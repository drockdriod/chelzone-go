package models

type SocialMediaUser struct {
	ScreenName string `bson:"screenName" json:"-"`
	Site string `bson:"site" json:"-"`
}