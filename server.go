package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
    "encoding/json"
    "context"
    "io/ioutil"
    "github.com/drockdriod/chelzone-go/db"
    "github.com/mongodb/mongo-go-driver/bson"
    "log"
)

func main() {
	ctx := context.Background()


	r := gin.Default()

	// Be sure to use struts here to define a schema in which the JSON would conform to
	
	r.GET("/standings", func(c *gin.Context) {
		// Perform get request on NHL API to get standings
		client, err := db.Connect(ctx)

		filter := bson.M{}

		cur, err := client.Collection("teams").Find(ctx, filter)


	   	if err != nil { log.Fatal(err) }
		
		defer cur.Close(ctx)
		
		for cur.Next(ctx) {
			elem := &bson.D{}
		   	if err := cur.Decode(elem); err != nil {
				log.Fatal(err)
			}

		   	
		   	// if err != nil { log.Fatal(err) }
	   		fmt.Printf("%s", elem)
		   // do something with elem....
		}
		
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
	r.Run() // listen and serve on 0.0.0.0:8080
}