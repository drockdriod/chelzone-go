package teams

import (
	"fmt"
	"net/http"
	"encoding/json"
    "io/ioutil"
    "os"
    "log"
    "github.com/drockdriod/chelzone-go/models"
)

type Teams struct{
	TeamsList []models.Team `json:"teams"`
}

type Stats struct{
	StatsList []Stat `json:"stats"`
}

type Stat struct{
	Type map[string]string `json:"-"`
	Splits []Split `json:"splits"`
}

type Split struct{
	Stat models.TeamStats `json:"stat"`
	Team models.Team `json:"team"`
}

func GetTeams() (Teams) {
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/teams/")
		
	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Declare a map variable for the JSON data being retrieved
	var data Teams

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    // Get JSON from response body and assign to data variable
	json.Unmarshal(body, &data)

	return data
}


func GetStatsByTeam(team models.Team) models.TeamStats {
	resp, err := http.Get(fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/teams/%v/stats", team.Id))
		
	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Declare a map variable for the JSON data being retrieved
	var data Stats

	body1, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }


	body2, err := json.Marshal(team)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    // Get JSON from response body and assign to data variable
	json.Unmarshal(body1, &data)

	var moddedData = data.StatsList[0].Splits[0].Stat

	json.Unmarshal(body2, &moddedData.TeamRef)

	log.Println("team:")
	log.Println(moddedData)

	return moddedData
}