package models

import (
	"github.com/drockdriod/chelzone-go/db"
    "github.com/mongodb/mongo-go-driver/bson"
	"log"
)

type Team struct {
	Id int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Stats *FullTeamStats `json:"stats,omitempty" bson:"stats,omitempty"`
	Roster *[]Player `json:"roster,omitempty" bson:"roster,omitempty"`
}

func (t Team) GetAll() ([]bson.M){
	results := db.GetItems("teams", bson.M{})

	log.Println(results)

	return results
}

func (t Team) GetByFilter(filter interface{}) (*Team, error) {
	team := new(Team)

	results, err := db.FindByAggregate("teams", filter)

	var body, _ = bson.Marshal(results[0])

	bson.Unmarshal(body, &team)

	return team, err
}