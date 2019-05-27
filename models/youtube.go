package models

import(
	"github.com/drockdriod/chelzone-go/db"
 //    "go.mongodb.org/mongo-driver/bson"
  	"google.golang.org/api/youtube/v3"
)

type Youtube struct{
	Activities []*youtube.Activity
	SearchResults []*youtube.SearchResult `json:"items" bson"items"`
}
func (model Youtube) GetYoutubeItems(filter interface{}) ([]interface{}, error){
	// twitter := new(Twitter)
	results, err := db.FindByAggregate("youtubeitems", filter)

	return results, err
}