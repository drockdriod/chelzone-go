import ApiClient from '../ApiClient'

export const getLeaders = async({commit}) => {
	const results = await ApiClient.perform('get', '/players/leaders')

	console.log(results)

	commit('setLeaders', results.leaders)
}