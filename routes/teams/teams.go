package teams

import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/controllers"
)

var teamController = new(controllers.TeamController)

func ServeRoutes(parentPath *gin.RouterGroup) *gin.RouterGroup {
	r := parentPath.Group("/teams")

	r.GET("/", teamController.GetAll)

	r.GET("/:team-slug", teamController.GetByTeamSlug)

	r.GET("/:team-id/games", teamController.GetTeamGames)

	return r
}