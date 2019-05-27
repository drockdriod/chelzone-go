package models
// import "go.mongodb.org/mongo-driver/bson"

type TeamStats struct {
	TeamRef Team `json:"team" bson:"team,omitempty"`
	GoalsAgainst int `json:"goalsAgainst,omitempty" bson:"goalsAgainst,omitempty"`
	GoalsScored int `json:"goalsScored,omitempty" bson:"goalsScored,omitempty"`
	Points int `json:"points,omitempty" bson:"points,omitempty"`
	DivisionRank string `json:"divisionRank,omitempty" bson:"divisionRank,omitempty"`
	ConferenceRank string `json:"conferenceRank,omitempty" bson:"conferenceRank,omitempty"`
	LeagueRank string `json:"leagueRank,omitempty" bson:"leagueRank,omitempty"`
	WildCardRank string `json:"wildCardRank,omitempty" bson:"wildCardRank,omitempty"`
	Row int `json:"row,omitempty" bson:"row,omitempty"`
	GamesPlayed int `json:"gamesPlayed,omitempty" bson:"gamesPlayed,omitempty"`
	PtPctg string  `json:"ptPctg,omitempty" bson:"ptPctg,omitempty"`
	GoalsPerGame float32  `json:"goalsPerGame,omitempty" bson:"goalsPerGame,omitempty"`
	GoalsAgainstPerGame float32  `json:"goalsAgainstPerGame,omitempty" bson:"goalsAgainstPerGame,omitempty"`
	EvGGARatio float32  `json:"evGGARatio,omitempty" bson:"evGGARatio,omitempty"`
	PowerPlayPercentage string  `json:"powerPlayPercentage,omitempty" bson:"powerPlayPercentage,omitempty"`
	PowerPlayGoals int  `json:"powerPlayGoals,omitempty" bson:"powerPlayGoals,omitempty"`
	PowerPlayGoalsAgainst int  `json:"powerPlayGoalsAgainst,omitempty" bson:"powerPlayGoalsAgainst,omitempty"`
	PowerPlayOpportunities int  `json:"powerPlayOpportunities,omitempty" bson:"powerPlayOpportunities,omitempty"`
	PenaltyKillPercentage string  `json:"penaltyKillPercentage,omitempty" bson:"penaltyKillPercentage,omitempty"`
	ShotsPerGame float32  `json:"shotsPerGame,omitempty" bson:"shotsPerGame,omitempty"`
	ShotsAllowed float32  `json:"shotsAllowed,omitempty" bson:"shotsAllowed,omitempty"`
	WinScoreFirst float32  `json:"winScoreFirst,omitempty" bson:"winScoreFirst,omitempty"`
	WinOppScoreFirst float32  `json:"winOppScoreFirst,omitempty" bson:"winOppScoreFirst,omitempty"`
	WinLeadFirstPer float32  `json:"winLeadFirstPer,omitempty" bson:"winLeadFirstPer,omitempty"`
	WinLeadSecondPer float32  `json:"winLeadSecondPer,omitempty" bson:"winLeadSecondPer,omitempty"`
	WinOutshootOpp float32  `json:"winOutshootOpp,omitempty" bson:"winOutshootOpp,omitempty"`
	WinOutshotByOpp float32  `json:"winOutshotByOpp,omitempty" bson:"winOutshotByOpp,omitempty"`
	FaceOffsTaken int  `json:"faceOffsTaken,omitempty" bson:"faceOffsTaken,omitempty"`
	FaceOffsWon int  `json:"faceOffsWon,omitempty" bson:"faceOffsWon,omitempty"`
	FaceOffsLost int  `json:"faceOffsLost,omitempty" bson:"faceOffsLost,omitempty"`
	FaceOffWinPercentage string  `json:"faceOffWinPercentage,omitempty" bson:"faceOffWinPercentage,omitempty"`
	ShootingPctg float32  `json:"shootingPctg,omitempty" bson:"shootingPctg,omitempty"`
	SavePctg float32  `json:"savePctg,omitempty" bson:"savePctg,omitempty"`
	// LeagueRecord map[string]int `json:"leagueRecord,omitempty" json:"leagueRecord,omitempty` {
	// 	wins `json:"wins,omitempty" json:"wins,omitempty`
	// 	ot `json:"ot,omitempty" json:"ot,omitempty`
	// 	losses `json:"losses,omitempty" json:"losses,omitempty`
	// }
}