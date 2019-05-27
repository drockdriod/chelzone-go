package controllers


import (
	"github.com/gin-gonic/gin"
    "github.com/drockdriod/chelzone-go/models"
    "net/http"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/drockdriod/chelzone-go/utils"
    "log"
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

func (controller TeamController) GetByTeamSlug(c *gin.Context) {
	slug := c.Param("team-slug")
	name := utils.CapitalizeEachWord(slug, "-")

	team, err := teamModel.GetByFilter([]bson.D{
		{{"$match", bson.M{"name": name }}},
		{{"$lookup",bson.M{
			"from": "teamstats",
			"localField": "id",
			"foreignField": "team.id",
			"as": "stats",
		}}},
		{{"$unwind", "$stats"}},
		{{"$lookup",bson.M{
			"from": "players",
			"localField": "id",
			"foreignField": "team.id",
			"as": "roster",
		}}},
	})

	log.Println(team)

	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})

	return
}