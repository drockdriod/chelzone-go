package players

import (
	"fmt"
	"net/http"
	"encoding/json"
    "io/ioutil"
    "os"
    "log"
    "github.com/drockdriod/chelzone-go/models"
    "github.com/drockdriod/chelzone-go/db"
    "github.com/mongodb/mongo-go-driver/bson"
	mongoOptions "github.com/mongodb/mongo-go-driver/mongo/options"
)

type Stats struct{
	Copyright string `json:"-"`
	StatsList []Stat `json:"stats"`
}

type Stat struct{
	Type map[string]string `json:"-"`
	Splits []Split `json:"splits"`
}

type Split struct{
	season string `json:"-"`
	PlayerStats models.PlayerStats `json:"stat"`
}

type StatsLeaders struct{
	GoalieStats []MeasurementItem `json:"goalie"` 
	PlayerStats []MeasurementItem `json:"skater"` 
}

type MeasurementItem struct{
	Measure string `json:"measure"`
	Leaders []models.LeaderStats `json:"leaders"`
}

func MarshalAndInsertLeaders(record *MeasurementItem){
	var player models.PlayerLeader
	player.Measure = record.Measure

	options := mongoOptions.FindOneAndReplace()
	options.SetUpsert(true)

	for _, p := range record.Leaders{
		player.PersonRef.Id = p.PlayerId
		player.PersonRef.FullName = p.FullName
		player.TeamRef.Id = p.TeamId
		player.TeamRef.Name = p.TeamFullName

		player.Stats = p

		db.FindOneAndReplace("playerleaders", bson.M{
			"stats.rankingIndex": player.Stats.RankingIndex,
			"measure": player.Measure,
		}, player, options)
	}
}

func GetPlayerStatsById(id int) (models.PlayerStats){
	uri := fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/people/%d/stats?stats=statsSingleSeason", id)
	log.Println(uri)
	resp, err := http.Get(uri)
		
	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Declare a map variable for the JSON data being retrieved
	var data Stats

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    // Get JSON from response body and assign to data variable
	json.Unmarshal(body, &data)

	log.Println(data)

	return data.StatsList[0].Splits[0].PlayerStats
}

func GetPlayerImage(player models.Player) []byte{
	resp, err := http.Get(fmt.Sprintf("https://nhl.bamcontent.com/images/headshots/current/168x168/%d.jpg", player.Person.Id))
		
	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    return body
}

func GetStatsLeaders() (StatsLeaders){
	resp, err := http.Get("http://www.nhl.com/stats/rest/leaders?season=current&gameType=2")

	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Declare a map variable for the JSON data being retrieved
	var data StatsLeaders

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    // Get JSON from response body and assign to data variable
	json.Unmarshal(body, &data)

	return data
}

func GetPlayersByTeam(team models.Team) ([]models.Player){
	resp, err := http.Get(fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/teams/%v/roster", team.Id))
		
	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Declare a map variable for the JSON data being retrieved
	var data models.Roster

	body1, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    // Get JSON from response body and assign to data variable
	json.Unmarshal(body1, &data)




	return data.PlayerList
}