package models

import (
	"github.com/drockdriod/chelzone-go/db"
	"encoding/json"
    "github.com/mongodb/mongo-go-driver/bson"
	"log"
)

type Team struct {
	Id int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

func (t Team) GetAll() (*[]Team){
	teams := new([]Team)
	
	results := db.GetItems("teams", bson.M{})

	log.Println(results)

	var body, _ = json.Marshal(results)

	json.Unmarshal(body, &teams)

	return teams
}