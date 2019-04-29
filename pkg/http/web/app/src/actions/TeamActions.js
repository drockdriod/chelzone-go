import ApiClient from '../ApiClient'

export const getTeams = async({commit}) => {
	const results = await ApiClient.perform('get', '/teams/')

	console.log(results)

	commit('setTeams', results.teams)
}