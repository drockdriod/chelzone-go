package twitter

import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/controllers"
)

var twitterController = new(controllers.TwitterController)

func ServeRoutes(parentPath *gin.RouterGroup) *gin.RouterGroup {
	r := parentPath.Group("/twitter")

	r.GET("/tweets/*skip", twitterController.GetTweets)


	return r
}