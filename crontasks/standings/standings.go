package standings

import (
	"fmt"
	"net/http"
	"encoding/json"
    "io/ioutil"
    "os"
    "github.com/drockdriod/chelzone-go/models"
)

func GetStandings() (models.Standings) {
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/standings")
		
	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Declare a map variable for the JSON data being retrieved
	var data models.Standings

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }

    // Get JSON from response body and assign to data variable
	json.Unmarshal(body, &data)

	return data
}