package controllers


import (
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/drockdriod/chelzone-go/models"
	"net/http"
)

type TwitterController struct{}

var twitterModel = new(models.Twitter)

func (controller TwitterController) GetTweets(c *gin.Context){
	tweets := twitterModel.GetTweets(bson.M{
		"in_reply_to_user_id_str":"",
	})


	c.JSON(http.StatusOK, gin.H{
		"tweets": tweets,
	})
}