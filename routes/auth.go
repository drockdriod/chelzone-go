package routes

import (
	"github.com/gin-gonic/gin"
    "io/ioutil"
    "github.com/drockdriod/chelzone-go/db"

)
// Binding from JSON
type Registration struct {
	Name     string `form:"name" json:"user" binding:"required"`
	Email 	 string `form: "email" json:"email" xml;"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`

}

func Auth(){
	r := gin.Default()
	
	// Be sure to use struts here to define a schema in which the JSON would conform to
	
	r.POST("/register", func(c *gin.Context) {
		// Perform get request on NHL API to get standings

		name := c.Param("name")
		email := c.Param("email")
		password := c.Param("password")

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


		c.JSON(200, gin.H{
			"message": "Your account is registered",
		})

	}

}