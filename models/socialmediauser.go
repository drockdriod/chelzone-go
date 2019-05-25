package models

type SocialMediaUser struct {
	Identity string `bson:"identity" json:"-"`
	Site string `bson:"site" json:"-"`
}