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

func getPlayerById(id string){
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/people/%s/stats?stats=statsSingleSeason", id)
		
	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Declare a map variable for the JSON data being retrieved
	var data map[string]interface{}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    // Get JSON from response body and assign to data variable
	json.Unmarshal(body, &data)

	return data
}