
export const teams = (state) =>{
	return state.teams
}

export const teamsByDivision = (state) => {
	let t = {}

	state.teams.map(team => {
		if(t[team.division.name] === undefined){
			t[team.division.name] = {
				name: team.division.name,
				teams: []
			}
		}
		t[team.division.name].teams.push(team)
	})

	return Object.values(t)
}