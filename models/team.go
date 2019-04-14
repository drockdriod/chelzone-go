package models

type Team struct {
	Id int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}