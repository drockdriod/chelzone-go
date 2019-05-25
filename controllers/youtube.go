package controllers


import (
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/drockdriod/chelzone-go/models"
	"net/http"
	"strings"
	"strconv"
)


type YoutubeController struct{}

var youtubeModel = new(models.Youtube)

func (controller YoutubeController) GetYoutubeItems(c *gin.Context){
	skipParam := c.Param("skip")
	skip := "0"

	if(skipParam != "/"){
		skip = strings.Replace(skipParam, "/", "", 1)
	}
	skipInt, _ := strconv.Atoi(skip)

	youtubeitems, err := youtubeModel.GetYoutubeItems([]bson.D{
		{{"$sort", bson.M{
			"pubishedAt": -1,
		}}},
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
		"youtubeitems": youtubeitems,
	})
}