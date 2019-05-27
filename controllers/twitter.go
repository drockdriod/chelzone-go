package controllers


import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/drockdriod/chelzone-go/models"
	"net/http"
	// "log"
	"strings"
	"strconv"
)

const QUERY_LIMIT = 50

type TwitterController struct{}

var twitterModel = new(models.Twitter)

func (controller TwitterController) GetTweets(c *gin.Context){
	skipParam := c.Param("skip")
	skip := "0"

	if(skipParam != "/"){
		skip = strings.Replace(skipParam, "/", "", 1)
	}
	skipInt, _ := strconv.Atoi(skip)

	tweets, err := twitterModel.GetTweets([]bson.D{
		{{"$match", bson.M{
			"in_reply_to_user_id_str":"",
		}}},
		// {{"$sort", bson.M{
		// 	"created_at": -1,
		// }}},
		{{"$skip", skipInt}},
		{{"$limit", QUERY_LIMIT}},
	})

	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"tweets": tweets,
	})
}