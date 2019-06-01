import ApiClient from '../ApiClient'

export const getTeams = async({commit, dispatch}) => {

	const results = await ApiClient.perform('get', '/teams/')

	commit('setTeams', results.teams)
}


export const getGamesByTeamId = async({commit}, teamId) => {
	const results = await ApiClient.perform('get', `/teams/${teamId}/games`)

	commit('setTeamGames', results.games)
}

export const getGameMilestonesByTeamId = async({commit}, teamId) => {
	const results = await ApiClient.perform('get', `/teams/${teamId}/gamemilestones`)

	commit('setTeamMilestones', results.gameMilestones)
}