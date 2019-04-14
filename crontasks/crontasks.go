package crontasks

import (
	"github.com/robfig/cron"
	"github.com/drockdriod/chelzone-go/crontasks/standings"
	"github.com/drockdriod/chelzone-go/crontasks/teams"
    "github.com/mongodb/mongo-go-driver/bson"
	mongoOptions "github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/drockdriod/chelzone-go/db"
	"log"

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

func CollectTeams(){
	results := teams.GetTeams()

	log.Println(results)

	options := mongoOptions.FindOneAndReplace()
	options.SetUpsert(true)

	for _, r := range results.TeamsList {
		log.Println(r)

		result := db.FindOneAndReplace("teams", bson.M{
			"id": r.Id,
		}, r, options)
		
		log.Println(result)
	}

}

func Init(){
	c := cron.New()

	c.AddFunc("0 30 22-23 * 10,11,12,1,2,3,4 *", func() { 
		CollectStandings()
	})

	c.AddFunc("0 30 2 1 9 ?", func() { 
		CollectTeams()
	})

	c.Start()
}