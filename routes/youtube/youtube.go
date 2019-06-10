package youtube

import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/controllers"
)

var youtubeController = new(controllers.YoutubeController)

func ServeRoutes(parentPath *gin.RouterGroup) *gin.RouterGroup {
	r := parentPath.Group("/youtube")

	r.GET("/items/*skip", youtubeController.GetYoutubeItems)


	return r
}