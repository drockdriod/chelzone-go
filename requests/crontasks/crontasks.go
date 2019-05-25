package crontasks

import (
	"github.com/robfig/cron"
	"github.com/drockdriod/chelzone-go/requests/standings"
	"github.com/drockdriod/chelzone-go/requests/teams"
	"github.com/drockdriod/chelzone-go/requests/players"
	"github.com/drockdriod/chelzone-go/requests/twitter"
	"github.com/drockdriod/chelzone-go/requests/youtube"
    "github.com/mongodb/mongo-go-driver/bson"
	mongoOptions "github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/drockdriod/chelzone-go/db"
    "github.com/drockdriod/chelzone-go/models"
	"log"
	"net/http"
	"fmt"
	"os"
    "io/ioutil"

)

func CollectStandings(){
	results := standings.GetStandings()

	records := results.RecordList

	options := mongoOptions.FindOneAndReplace()
	options.SetUpsert(true)

	for _, record := range records {
			  // log.Println(record.TeamRecords )
	  	for _, teamStats := range record.TeamRecords {
			log.Println(teamStats.TeamRef)

			result := db.FindOneAndReplace("standings", bson.M{
				"team.id": teamStats.TeamRef.Id,
			}, teamStats, options)
			
			log.Println(result)

		}
	}


}

func CollectTopLeaders(){
	
	statsLeaders := players.GetStatsLeaders()

	goalieStats := statsLeaders.GoalieStats
	playerStats := statsLeaders.PlayerStats


	for _, grecord := range goalieStats {
		
		players.MarshalAndInsertLeaders(&grecord)

	}


	for _, precord := range playerStats {
		
		players.MarshalAndInsertLeaders(&precord)

		
	}
}


func CollectTeams(){
	results := teams.GetTeams()


	options := mongoOptions.FindOneAndReplace()
	options.SetUpsert(true)

	for _, r := range results.TeamsList {

		logo := teams.GetLogoByTeam(r)

		r.Logo = logo

		result := db.FindOneAndReplace("teams", bson.M{
			"id": r.Id,
		}, r, options)
		
		log.Println(result)
	}

}

func CollectTeamStats(){

	var team models.Team
	options := mongoOptions.FindOneAndReplace()
	options.SetUpsert(true)

	results := db.GetItems("teams", bson.M{})

	for _, v := range results {
		
		var body, err = bson.Marshal(v)

		if(err != nil){
			log.Fatal(err)
		}

		bson.Unmarshal(body, &team)

		stats := teams.GetStatsByTeam(team)
		
		_ = db.FindOneAndReplace("teamstats", bson.M{
			"team.id": team.Id,
		}, stats, options)
	}
}

func CollectTeamPlayers(){

	var team models.Team
	options := mongoOptions.FindOneAndReplace()
	options.SetUpsert(true)

	results := db.GetItems("teams", bson.M{})

	for _, v := range results {
		
		var body, err = bson.Marshal(v)

		if(err != nil){
			log.Fatal(err)
		}

		bson.Unmarshal(body, &team)

		results2 := players.GetPlayersByTeam(team)

		
		for _, player := range results2 {
			bson.Unmarshal(body, &player.TeamRef)

			_ = db.FindOneAndReplace("players", bson.M{
				"person.id": player.Person.Id,
			}, player, options)
		}
	}
}

func CollectPlayerImages(){
	var player models.Player

	results := db.GetItems("players", bson.M{})


	for _, v := range results {
		var body, err = bson.Marshal(v)

		if(err != nil){
			log.Fatal(err)
		}

		bson.Unmarshal(body, &player)

		img := players.GetPlayerImage(player)


		player.BadgeImage = img

		_, err = db.UpdateObj(
			"players", 
			bson.M{"person.id": player.Person.Id, }, 
			bson.D{
				{"$set", bson.M{ 
					"badgeImage": player.BadgeImage,
				}},
			},
		)

		if(err != nil){
			log.Fatal(err.Error())
		}

	}
}

func TestIp(t *http.Client){
	resp, err := t.Get("https://www.whatismyip.com/")

	if err != nil {
		fmt.Printf("Error: %s",err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	responseString := string(responseData)
	log.Println(responseString)
}

func CollectTweetsFromFollowers(){
	users := db.GetItems("socialmediausers", bson.M{"site":"twitter"})

	twitter.GetTweetsFromUsers(users)
}

func CollectYoutubeSearchResults(){
	users := db.GetItems("socialmediausers", bson.M{"site":"youtube"})
	youtube.GetVideosFromUsers(users)
}

func Init(){
	c := cron.New()

	c.AddFunc("0 30 22-23 * 10,11,12,1,2,3,4 *", func() { 
		go CollectStandings()
		go CollectTopLeaders()
	})

	c.AddFunc("0 30 2 1 9 ?", func() { 
		go CollectTeams()
	})

	c.AddFunc("@every 0h30m", func() {
		go CollectTweetsFromFollowers()
	})

	c.AddFunc("@every 1h", func() {
		go CollectYoutubeSearchResults()
	})


	c.AddFunc("0 30 2 1 3,7,10 ?", func() { 
		go CollectTeamPlayers()
		go CollectPlayerImages()
	})

	c.AddFunc("0 30 2 * 10,11,12,1,2,3,4 *", func() { 
		go CollectTeamStats()
	})

	c.Start()
}