package models

import (
	"github.com/drockdriod/chelzone-go/db"
    "github.com/mongodb/mongo-go-driver/bson"
	"log"
)

type Team struct {
	Id int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Stats FullTeamStats `json:"stats,omitempty" bson:"stats,omitempty"`
}

func (t Team) GetAll() (*[]Team){
	teams := new([]Team)

	results := db.GetItems("teams", bson.M{})

	log.Println(results)

	var body, _ = bson.Marshal(results)

	bson.Unmarshal(body, &teams)

	return teams
}

func (t Team) GetByFilter(filter interface{}) (interface{}, error) {
	team := new(Team)

	results, err := db.FindByAggregate("teams", filter)

	var body, _ = bson.Marshal(results[0])

	bson.Unmarshal(body, &team)

	log.Println("results")
	log.Println(results)

	return team, err
}