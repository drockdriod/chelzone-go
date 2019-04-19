package controllers


import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/models"
    "net/http"
)

type TeamController struct{}

var teamModel = new(models.Team)

func (controller TeamController) GetAll(c *gin.Context) {
	teams := teamModel.GetAll()

	c.JSON(http.StatusOK, gin.H{
		"teams": teams,
	})
	return
}