package controllers


import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/models"
    playerReqs "github.com/drockdriod/chelzone-go/requests/players"
    "net/http"
    "github.com/mongodb/mongo-go-driver/bson"
    "github.com/drockdriod/chelzone-go/utils"
    "log"
)



type PlayerController struct{}

var playerModel = new(models.Player)

func (controller PlayerController) GetByPlayerSlug(c *gin.Context) {
	slug := c.Param("player-slug")
	name := utils.CapitalizeEachWord(slug, "-")

	player, err := playerModel.GetOne(bson.M{
		"person.fullName": name,
	})

	p := playerReqs.GetPlayerStatsById(player.Person.Id)

	log.Println(p)

	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"player": player,
		"stats": p,
	})

	return
}