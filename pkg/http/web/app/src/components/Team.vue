<template>
	<v-layout row wrap>
	  
		<v-toolbar
	    	color="info"
	    	flat
	    	prominent
	  	>

		    <v-toolbar-title>
		    	<v-avatar
	                :tile="true"
	                :size="130">
	                <img v-if="team.logo" v-bind:src="convertFromBinary('image/svg+xml',team.logo)" v-bind:alt="team.name">
	            </v-avatar>
		    	{{team.name}}
		    </v-toolbar-title>

		    <v-spacer></v-spacer>

		    <v-btn icon>
		      <v-icon>mdi-export</v-icon>
		    </v-btn>
	  	</v-toolbar>
	    
	    <v-container grid-list-md>
			<v-data-table
			    :headers="headers"
			    :items="team.roster"
    			:loading="true"
    			:rows-per-page-items="rowsPerPage"
    			>
   				<v-progress-linear slot="progress" color="blue" indeterminate v-show="loading"></v-progress-linear>
    			<template slot="items" slot-scope="props">
          			<tr @click="vm.$router.push({ name: 'player', params: { slug: props.item.person.fullName.split(' ').join('-') }})">
				        <td>
				        	<PlayerAvatar :binary="props.item.badgeImage" />
				        </td>
				        <td>{{ props.item.person.fullName }}</td>
				        <td>{{ props.item.position.name }}</td>
				        <td>{{ props.item.position.type }}</td>
				        <td>{{ props.item.jerseyNumber }}</td>
				    </tr>
			    </template>
			</v-data-table>
		</v-container>
	</v-layout>
</template>
<script>
    import ApiClient from '../ApiClient'
    import { convertFromBinary } from '../utils/common'
    import PlayerAvatar from './PlayerAvatar'

    export default {
        components: { PlayerAvatar },
    	data () {
		    return {
		      	team: {},
		      	headers: [
		      		{text: '', value: 'badgeImage'},
		      		{text: 'Name', value: 'person.fullName'},
		      		{text: 'Position Name', value: 'position.name'},
		      		{text: 'Position Type', value: 'position.type'},
		      		{text: 'Jersey Number', value: 'jerseyNumber'},
		      	],
		      	loading: true,
		      	rowsPerPage: [10,25,{"text":"All","value":-1}],
		      	vm: this
		    }
	  	},
	  	async beforeRouteEnter(to, from, next) {
        	const slug = to.params.slug
        	const result = await ApiClient.perform('get',`/teams/${slug}`)
        	next(vm => vm.setTeam(result.team))
	  		
	  	},
        async beforeRouteUpdate (to, from, next) {
            this.team = {}
            const slug = to.params.slug
            const result = await ApiClient.perform('get',`/teams/${slug}`)
            
            this.setTeam(result.team)
            next()
        },
        methods: {
            setTeam(team){
            	this.team = team
            	this.loading = false
            },
            convertFromBinary
        }
    }
</script>