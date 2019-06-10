package games

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
	"log"
	"strings"
	"github.com/drockdriod/chelzone-go/models"
)

type Schedule struct {
	Dates []ScheduleDate `json:"dates"`
}

type ScheduleDate struct {
	Date string `json:"date"`
	Games []models.Game `json:"games"`
}

type GameContent struct{
	Media GameMedia `json:"media"`
}

type GameMedia struct{
	Milestones GameMilestones `json:"milestones"`
}

type GameMilestones struct{
	Title string `json:"title"`
	StreamStart string `json:"streamStart"`
	Items []models.GameMilestone `json:"items"`
}

// type ImageCuts struct{
// 	Cut ImageCut `json:"1136x640"`
// }

// type ImageCut struct{
// 	Src string `json:"src"`
// }

func GetSchedule(criteria ...string) Schedule{
	var strBuilder strings.Builder	
	param := ""

	if(len(criteria) > 0 && criteria[0] != ""){
		strBuilder.WriteString(fmt.Sprintf("?date=%s",criteria[0]))
	}

	log.Println("criteria")
	log.Println(criteria)

	if(len(criteria) > 1 && criteria[1] != ""){
		strBuilder.WriteString(fmt.Sprintf("&teamId=%s",criteria[1]))
	}

	if(len(criteria) > 0){
		param = strBuilder.String()
	}

	resp, err := http.Get(fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/schedule%s",param))

	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Declare a map variable for the JSON data being retrieved
	var data Schedule

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    // Get JSON from response body and assign to data variable
	json.Unmarshal(body, &data)

	return data
}

func GetGameContent(gamePK int) GameContent {
	resp, err := http.Get(fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/game/%d/content",gamePK))

	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Declare a map variable for the JSON data being retrieved
	var data GameContent

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    // Get JSON from response body and assign to data variable
	err = json.Unmarshal(body, &data)
	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
	return data
}