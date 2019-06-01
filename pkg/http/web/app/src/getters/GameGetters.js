import { TOP_MILESTONE_LIMIT } from '../utils/common'

export const gameMilestones = (state) =>{
	return state.gameMilestones
}

export const topGameMilestones = (state) => {
	const milestones = state.gameMilestones.slice(0,TOP_MILESTONE_LIMIT)
	return [...milestones.map(m => (
		
		{
			...m,
			videoUrl: m.highlight.playbacks[m.highlight.playbacks.length - 1].url,
			imageUrl: m.highlight.image.cuts['768x432'].src
		}
	))]
}