package twitter
import (
    "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
    "github.com/drockdriod/chelzone-go/db"
    "github.com/drockdriod/chelzone-go/models"
	"fmt"
	"os"
	"time"
    "strconv"
	"log"
	"encoding/json"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
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

func GetFollowers() (*twitter.Followers){
	followers, _, _ := client.Followers.List(&twitter.FollowerListParams{})

	fmt.Printf("SEARCH Followers:\n%+v\n", followers)
	return followers
}

func GetTweetsFromUsers(users []bson.M){
	var user models.SocialMediaUser
	// userLookupParams := &twitter.UserLookupParams{ScreenName: []string{"TSNBobMckenzie"}}
	// users, _, _ := client.Users.Lookup(userLookupParams)

    now := time.Now()
    dateStart := now.Add(-30*time.Minute)

	sinceIdStr := dateStart.Format("200601021504")

    sinceId, err := strconv.ParseInt(sinceIdStr,10,64)
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }

    log.Println(sinceId)

	for _, item := range users {

		body, _ := bson.Marshal(item)

		bson.Unmarshal(body, &user)

		query := fmt.Sprintf("from: %s",user.Identity)

		searchTweetParams := &twitter.SearchTweetParams{
			Query:     query,
			Count:     2,
			SinceID: sinceId,
			TweetMode: "extended",
		}

		search, _, _ := client.Search.Tweets(searchTweetParams)

		options := mongoOptions.InsertMany()
		options.SetOrdered(false)

		var items []interface{}
		body2, _ := json.Marshal(search.Statuses)
		json.Unmarshal(body2, &items)

		log.Println(items)


		db.InsertMany("tweets",items,options)
		fmt.Printf("SEARCH TWEETS:\n%+v\n", search)
		fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)
	}
}