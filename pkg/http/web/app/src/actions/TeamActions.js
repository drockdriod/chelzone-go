import ApiClient from '../ApiClient'

export const getTeams = async({commit, dispatch}) => {

	const results = await ApiClient.perform('get', '/teams/')

	commit('setTeams', results.teams)
}