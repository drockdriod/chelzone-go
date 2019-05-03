package twitter
import (
    "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
    "github.com/drockdriod/chelzone-go/db"
	"fmt"
	"os"
	"log"
	"encoding/json"
	mongoOptions "github.com/mongodb/mongo-go-driver/mongo/options"
)

var client *twitter.Client

func Connect() *twitter.Client{

	apiKey := os.Getenv("TWITTER_API_KEY")
	apiKeySecret := os.Getenv("TWITTER_API_KEY_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	
	configOauth := oauth1.NewConfig(apiKey, apiKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := configOauth.Client(oauth1.NoContext, token)

	// Twitter client
	client = twitter.NewClient(httpClient)

	// verifyParams := &twitter.AccountVerifyParams{
	// 	SkipStatus:   twitter.Bool(true),
	// 	IncludeEmail: twitter.Bool(true),
	// }
	// user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	// fmt.Printf("User's ACCOUNT:\n%+v\n", user)


	return client
}

func GetUsers(){

	// userLookupParams := &twitter.UserLookupParams{ScreenName: []string{"TSNBobMckenzie"}}
	// users, _, _ := client.Users.Lookup(userLookupParams)

	searchTweetParams := &twitter.SearchTweetParams{
		Query:     "from:TSNBobMckenzie",
		Count:     3,
		SinceID: 201905021800,
	}
	search, _, _ := client.Search.Tweets(searchTweetParams)

	options := mongoOptions.InsertMany()
	options.SetOrdered(false)

	var items []interface{}
	body, _ := json.Marshal(search.Statuses)
	json.Unmarshal(body, &items)

	log.Println(items)

	db.InsertMany("tweets",items,options)
	fmt.Printf("SEARCH TWEETS:\n%+v\n", search)
	fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)

}