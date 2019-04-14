package models

type TeamStats struct {
	TeamRef Team `json:"team" bson:"team"`
	GoalsAgainst int `json:"goalsAgainst" bson:"goalsAgainst"`
	GoalsScored int `json:"goalsScored" bson:"goalsScored"`
	Points int `json:"points" bson:"points"`
	DivisionRank string `json:"divisionRank" bson:"divisionRank"`
	ConferenceRank string `json:"conferenceRank" bson:"conferenceRank"`
	LeagueRank string `json:"leagueRank" bson:"leagueRank"`
	WildCardRank string `json:"wildCardRank" bson:"wildCardRank"`
	Row int `json:"row" bson:"row"`
	GamesPlayed int `json:"gamesPlayed" bson:"gamesPlayed"`
}