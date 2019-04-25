package players

import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/controllers"
)

var playerController = new(controllers.PlayerController)

func ServeRoutes(parentPath *gin.RouterGroup) *gin.RouterGroup {
	r := parentPath.Group("/players")

	r.GET("/:player-slug", playerController.GetByPlayerSlug)

	return r
}