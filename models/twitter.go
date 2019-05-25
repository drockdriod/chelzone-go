package models

import(
	"github.com/drockdriod/chelzone-go/db"
    // "github.com/mongodb/mongo-go-driver/bson"
    twitterGo "github.com/dghubble/go-twitter/twitter"
)

type Twitter struct{
	Tweets []twitterGo.Tweet `json:"tweets" bson"tweets"`
}

func (model Twitter) GetTweets(filter interface{}) ([]interface{}, error){
	// twitter := new(Twitter)
	results, err := db.FindByAggregate("tweets", filter)

	return results, err
}