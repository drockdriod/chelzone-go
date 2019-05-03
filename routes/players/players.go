package players

import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/controllers"
)

var playerController = new(controllers.PlayerController)

func ServeRoutes(parentPath *gin.RouterGroup) *gin.RouterGroup {
	r := parentPath.Group("/players")

	r.GET("/leaders", playerController.GetLeaders)

	r.GET("/r/:player-slug", playerController.GetByPlayerSlug)


	return r
}