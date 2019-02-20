package models

type TeamStats struct {
	teamRef team
	goalsAgainst int
	goalsScored int
	points int
	divisionRank string
	conferenceRank string
	leagueRank string
	wildCardRank string
	row int
	gamesPlayed int
}