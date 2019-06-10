import ApiClient from '../ApiClient'

export const getGameMilestones = async({commit, dispatch}) => {

	const results = await ApiClient.perform('get', '/games/milestones')

	commit('setGameMilestones', results.gameMilestones)
}