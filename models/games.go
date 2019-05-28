package models

import (
	"github.com/drockdriod/chelzone-go/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"fmt"
	"log"
)

type Game struct {
	GamePk int `json:"gamePk" bson:"gamePk"`
	Link string `json:"link" bson:"link"`
	GameType string `json:"gameType" bson:"gameType"` 
	Season string `json:"season" bson:"season"`
	GameDate time.Time `json:"gameDate" bson:"gameDate"`
	Teams GameTeams `json:"teams" bson:"teams"`
}

type GameTeams struct {
	Away GameTeam `json:"away" bson:"away"`
	Home GameTeam `json:"home" bson:"home"`
}

type GameTeam struct {
	Team Team `json:"team" bson:"team"`
}

// type GameHighlights struct{
// 	Type string `json:"type" bson:"type"`
// 	TimeAbsolute string `json:"timeAbsolute" bson:"timeAbsolute"`
// 	Title string `json:"title" bson:"title"`
// 	Description string `json:"description" bson:"description"`
// 	TeamId string `json:"teamId" bson:"teamId"`
// 	PlayerId string `json:"playerId" bson:"playerId"`
// 	PeriodTime string `json:"periodTime" bson:"periodTime"`
// 	Period string `json:"period" bson:"period"`
// 	Image string `bson:"image"`
// }

type GameMilestone struct{
	GamePk int `bson:"gamePk"`
	Type string `json:"type" bson:"type"`
	TimeAbsolute string `json:"timeAbsolute" bson:"timeAbsolute"`
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	TeamId string `json:"teamId" bson:"teamId"`
	PlayerId string `json:"playerId" bson:"playerId"`
	PeriodTime string `json:"periodTime" bson:"periodTime"`
	Period string `json:"period" bson:"period"`
	Highlight Highlight `json:"highlight" bson:"highlight"`
}


type Highlight struct{
	Type string `json:"type" bson:"type"`
	Id string `json:"id" bson:"id"`
	Date string `json:"date" bson:"date"`
	Title string `json:"title" bson:"title"`
	Blurb string `json:"blurb" bson:"blurb"`
	Description string `json:"description" bson:"description"`
	Duration string `json:"duration" bson:"duration"`
	AuthFlow bool `json:"authFlow" bson:"authFlow"`
	MediaPlaybackId string `json:"mediaPlaybackId" bson:"mediaPlaybackId"`
	Image ScoreboardItemImage `json:"image,omitempty" bson:"image"`
	Playbacks []Playback `json:"playbacks" bson:"playbacks"`
}

type Playback struct{
	Name string `json:"name" bson:"name"`
	Width string `json:"width" bson:"width"`
	Height string `json:"height" bson:"height"`
	Url string `json:"url" bson:"url"`
}

type ScoreboardItemImage struct{
	Cuts map[string]interface{} `json:"cuts"`
}

func GetGamesByDate(chosenDate time.Time) []interface{} {
	startDate := time.Date(chosenDate.Year(), chosenDate.Month(), chosenDate.Day(),0,0,0,0, time.UTC)
	endDate := startDate.Add(time.Hour * 30)

	log.Println(fmt.Sprintf("ISODate(%s)",startDate.Format(time.RFC3339)))
	log.Println(endDate.Format(time.RFC3339))

	query := []bson.D{
		{{
			"$match", bson.M{
				"gameDate": bson.M{
					"$gte": startDate,
					"$lt":  endDate,
				},
			},
		}},
	}
	items, err := db.FindByAggregate("games", query)

	log.Println(items)

	if(err != nil){
		log.Fatal(err.Error())
	}

	return items

}

func (g GameMilestone) GetAllMilestones() []interface{}{
	query := []bson.D{
		{{"$addFields",bson.M{
			"timeAbsoluteDate": bson.M{
				"$dateFromString": bson.M{
					"dateString": "$timeAbsolute",
				},
			},
		}}},
		{{
			"$sort", bson.M{
				"timeAbsoluteDate": -1,
			},
		}},
	}
	items, err := db.FindByAggregate("gamemilestones", query)

	if(err != nil){
		log.Fatal(err.Error())
	}

	return items
}