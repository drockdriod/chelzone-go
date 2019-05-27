package models

import(
	"github.com/drockdriod/chelzone-go/db"
    "go.mongodb.org/mongo-driver/bson"
)

type Person struct{
	Id int `json:"id" bson:"id"`
	FullName string `json:"fullName" bson:"fullName"`
	Link string `json:"link,omitempty" bson:"link,omitempty"`
}

type position struct {
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
	Abbreviation string `json:"abbreviation" bson:"abbreviation"`
}

type Player struct{
	TeamRef Team `json:"team,omitempty" bson:"team"`
	Person Person `json:"person" bson:"person"`
	JerseyNumber string `json:"jerseyNumber" bson:"jerseyNumber"`
	Position position `json:"position" bson:"position"`
	PlayerStats *PlayerStats `json:"stats,omitempty" bson:"-"`
	BadgeImage []byte `json:"badgeImage,omitempty" bson:"badgeImage,omitempty"`
}

type PlayerStats struct{
	TimeOnIce string `json:"timeOnIce" bson:"timeOnIce"`
	Assists int `json: "assists" bson:"assists"`
	Goals int `json: "goals" bson:"goals"`
	Pim int `json: "pim" bson:"pim"`
	Shots int `json: "shots" bson:"shots"`
	Games int `json: "games" bson:"games"`
	Hits int `json: "hits" bson:"hits"`
	PowerPlayGoals int `json: "powerPlayGoals" bson:"powerPlayGoals"`
	PowerPlayPoints int `json: "powerPlayPoints" bson:"powerPlayPoints"`
	PowerPlayTimeOnIce string `json:"powerPlayTimeOnIce" bson:"powerPlayTimeOnIce"`
	EvenTimeOnIce string `json:"evenTimeOnIce" bson:"evenTimeOnIce"`
	PenaltyMinutes string `json:"penaltyMinutes" bson:"penaltyMinutes"`
	FaceOffPct int `json: "faceOffPct" bson:"faceOffPct"`
	ShotPct float32 `json:"shotPct" bson:"shotPct"`
	GameWinningGoals int `json: "gameWinningGoals" bson:"gameWinningGoals"`
	OverTimeGoals int `json: "overTimeGoals" bson:"overTimeGoals"`
	ShortHandedGoals int `json: "shortHandedGoals" bson:"shortHandedGoals"`
	ShortHandedPoints int `json: "shortHandedPoints" bson:"shortHandedPoints"`
	ShortHandedTimeOnIce string `json:"shortHandedTimeOnIce" bson:"shortHandedTimeOnIce"`
	Blocked int `json: "blocked" bson:"blocked"`
	PlusMinus int `json: "plusMinus" bson:"plusMinus"`
	Points int `json: "points" bson:"points"`
	Shifts int `json: "shifts" bson:"shifts"`
	TimeOnIcePerGame string `json:"timeOnIcePerGame" bson:"timeOnIcePerGame"`
	EvenTimeOnIcePerGame string `json:"evenTimeOnIcePerGame" bson:"evenTimeOnIcePerGame"`
	ShortHandedTimeOnIcePerGame string `json:"shortHandedTimeOnIcePerGame" bson:"shortHandedTimeOnIcePerGame"`
	PowerPlayTimeOnIcePerGame string `json:"powerPlayTimeOnIcePerGame" bson:"powerPlayTimeOnIcePerGame"`
}

type PlayerLeader struct{
	Measure string `json:"measure" bson:"measure"`
	PersonRef Person `json:"person" bson:"person"`
	TeamRef Team `json:"team" bson:"team"`
	Stats LeaderStats `json:"stats" bson:"stats"`
}

type PlayerLeaderGroup struct{
	Id string `bson:"-"`
	Player PlayerLeader `bson:"player"`
}


type LeaderStats struct{
	PlayerId int `json:"playerId,omitempty"`
	FullName string `json:"fullName,omitempty"`
	TeamId int `json:"teamId,omitempty"`
	TeamFullName string `json:"teamFullName,omitempty"`
	RankingIndex int `json:"listIndex" bson:"rankingIndex"`
	Value float32 `json:"value" bson:"value"`
}

type Roster struct{
	PlayerList []Player `json:"roster"`
}

func (p Player) GetOne(filter interface{}) (*Player, error) {
	player := new(Player)

	results, err := db.FindOne("players", filter)

	var body, _ = bson.Marshal(results)

	bson.Unmarshal(body, &player)

	return player, err
}

func (p Player) GetLeaders() ([]interface{}, error) {

	query := []bson.D{
		{{"$sort", bson.M{"stats.rankingIndex": 1, "measure": 1 } }},
		{{"$group", bson.M{
			"_id": "$measure", 
			"player": bson.M{ 
				"$first": "$$ROOT",
			},
		} }},
        {{
            "$project", bson.M{ "_id": 0, },
        }},
		{{"$lookup", bson.M{
			"from": "teams", 
			"localField": "player.team.id", 
			"foreignField": "id", 
			"as": "player.team",
		} }},
		{{"$unwind", "$player.team"}},
		{{"$lookup", bson.M{
			"from": "players", 
			"localField": "player.person.id", 
			"foreignField": "person.id", 
			"as": "player.details",
		} }},
		{{"$unwind", "$player.details"}},
        {{
            "$project", bson.M{ 
            	"player": 1, 
            	"order": bson.M{
            		"$eq": []string{"$player.details.position.code", "G"},
        		},
            },
        }},
		{{"$sort", bson.M{"order": 1 } }},
	}

	results, err := db.FindByAggregate("playerleaders",query)


	return results, err

}