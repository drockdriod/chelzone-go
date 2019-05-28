package players

import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/controllers"
)

var gameController = new(controllers.GameController)

func ServeRoutes(parentPath *gin.RouterGroup) *gin.RouterGroup {
	r := parentPath.Group("/games")

	r.GET("/milestones", gameController.GetAllMilestones)


	return r
}