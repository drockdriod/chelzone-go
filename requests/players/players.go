package players

import (
	"fmt"
	"net/http"
	"encoding/json"
    "io/ioutil"
    "os"
    "log"
    "github.com/drockdriod/chelzone-go/models"
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