package controllers

import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/models"
    "net/http"
)

type GameController struct{}

var gameMilestoneModel = new(models.GameMilestone)

func (controller GameController) GetAllMilestones(c *gin.Context) {
	gameMilestones := gameMilestoneModel.GetAllMilestones()

	c.JSON(http.StatusOK, gin.H{
		"gameMilestones": gameMilestones,
	})
	return
}
