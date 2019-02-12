package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
    "encoding/json"
    "io/ioutil"
)

func server() {
	r := gin.Default()
	r.GET("/standings", func(c *gin.Context) {
		// Perform get request on NHL API to get standings
		resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/standings")
		
		if err != nil {
			fmt.Printf("Error: %s",err)
			os.Exit(1)
		} else {

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

			fmt.Printf("Body: %s",data)

			c.JSON(200, gin.H{
				"standings": data,
			})

		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}