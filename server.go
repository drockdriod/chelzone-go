package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
	"fmt"
	"os"
    "encoding/json"
    "context"
    "io/ioutil"
    "github.com/drockdriod/chelzone-go/db"
	_ "github.com/joho/godotenv/autoload"
    "github.com/mongodb/mongo-go-driver/bson"
    "log"
    "github.com/drockdriod/chelzone-go/requests/crontasks"
    teamsRoute "github.com/drockdriod/chelzone-go/routes/teams"
    playersRoute "github.com/drockdriod/chelzone-go/routes/players"
    "github.com/drockdriod/chelzone-go/requests/twitter"
)

type element map[string]interface{}

func main() {
	ctx := context.Background()


	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT"}
	config.AllowHeaders = []string{"Accept", "access-control-allow-origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}
	r.Use(cors.New(config))
	
	client, err := db.Connect(ctx)

	if err != nil { 
		log.Fatal("error")
		log.Fatal(err) 
	}

	twitterClient := twitter.Connect()
	log.Println("TWIITERR")
	log.Println(twitterClient)
	// twitter.GetUsers()


	go crontasks.Init()

	root := r.Group("/api")
	{
		teamsRoute.ServeRoutes(root) 
		playersRoute.ServeRoutes(root)
	}

	// Be sure to use struts here to define a schema in which the JSON would conform to
	
	r.GET("/standings", func(c *gin.Context) {
		// Perform get request on NHL API to get standings

		filter := bson.M{}

		cur, err := client.Collection("teams").Find(ctx, filter)

		teams := make(map[string]string)

	   	if err != nil { log.Fatal(err) }
		
		defer cur.Close(ctx)
		
		for cur.Next(ctx) {
			elem := &element{}
		   	if err := cur.Decode(elem); err != nil {
				log.Fatal(err)
			}

			teams["lol"] = fmt.Sprint(elem)
		 
		   	// if err != nil { log.Fatal(err) }
		   // do something with elem....
		}
	   	
	   	fmt.Printf("%s", teams)
		
		if err := cur.Err(); err != nil {
			fmt.Printf("%s", err)
            os.Exit(1)
		}

		resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/standings")
		
		if err != nil {
			fmt.Printf("Error: %s",err)
			os.Exit(1)
		}

		defer resp.Body.Close()

		// Declare a map variable for the JSON data being retrieved
		data := &map[string]interface{}{}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }

        // Get JSON from response body and assign to data variable
		json.Unmarshal(body, data)

		// fmt.Printf("Body: %s",data)

		c.JSON(200, gin.H{
			"standings": data,
		})

		
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}