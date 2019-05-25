package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
    "context"
    "github.com/drockdriod/chelzone-go/db"
	_ "github.com/joho/godotenv/autoload"
    "log"
    "github.com/drockdriod/chelzone-go/requests/crontasks"
    teamsRoute "github.com/drockdriod/chelzone-go/routes/teams"
    playersRoute "github.com/drockdriod/chelzone-go/routes/players"
    twitterRoute "github.com/drockdriod/chelzone-go/routes/twitter"
    youtubeRoute "github.com/drockdriod/chelzone-go/routes/youtube"
    "github.com/drockdriod/chelzone-go/requests/twitter"
    "github.com/drockdriod/chelzone-go/requests/youtube"
	// "github.com/drockdriod/chelzone-go/requests/tor"
)


func main() {
	ctx := context.Background()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT"}
	config.AllowHeaders = []string{"Accept", "access-control-allow-origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}
	r.Use(cors.New(config))
	
	_, err := db.Connect(ctx)

	if err != nil { 
		log.Fatal("error")
		log.Fatal(err) 
	}

	twitterClient := twitter.Connect()
	log.Println("TWIITERR")
	log.Println(twitterClient)

	youtubeClient := youtube.Connect()
	log.Println("YOUTUBE")
	log.Println(youtubeClient)

	// torClient, err := tor.Connect(ctx)
	// if err != nil {
	// 	log.Panicf("Unable to start Tor: %v", err)
	// }
	// log.Println("TOR")
	// log.Println(torClient)


	go crontasks.Init()
	// go crontasks.TestIp(torClient)

	root := r.Group("/api")
	{
		teamsRoute.ServeRoutes(root) 
		playersRoute.ServeRoutes(root)
		twitterRoute.ServeRoutes(root)
		youtubeRoute.ServeRoutes(root)
	}

	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}