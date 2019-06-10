import ApiClient from '../ApiClient'

export const getLeaders = async({commit, dispatch}) => {
	const results = await ApiClient.perform('get', '/players/leaders')

	commit('setLeaders', results.leaders)
}