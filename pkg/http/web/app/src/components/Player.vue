<template>
    <v-container grid-list-md>
        <v-layout align-end justify-end row fill-height>
            <v-flex md9 lg9>
                <v-data-table
                    :hide-actions="true"
                    class="elevation-1"
                    :items="[player.stats]"
                    :headers="statsHeaders"
                >
                    <template v-slot:items="props">
                        <td class="text-xs-right">{{ props.item.season}}</td>
                        <td class="text-xs-right">{{ props.item.games}}</td>
                        <td class="text-xs-right">{{ props.item.goals}}</td>
                        <td class="text-xs-right">{{ props.item.assists}}</td>
                        <td class="text-xs-right">{{ props.item.shots}}</td>
                    </template>
                </v-data-table>
            </v-flex>
            <v-flex md3 lg3>
                <v-card>
                    <div class="text-xs-center">
                        <PlayerAvatar :binary="player.badgeImage" />
                    </div>
                    <!-- <v-card-title class="subheading font-weight-bold">{{ item.name }}</v-card-title>
                    <v-divider></v-divider> -->
                    <v-list dense>
                        <v-list-tile>
                            <v-list-tile-content>
                                <strong>Name</strong>
                            </v-list-tile-content>
                            <v-list-tile-content class="align-end">
                                {{player.person && player.person.fullName}}
                            </v-list-tile-content>
                        </v-list-tile>
                        <v-list-tile>
                            <v-list-tile-content>
                                <strong>Position Name</strong>
                            </v-list-tile-content>
                            <v-list-tile-content class="align-end">
                                {{player.position && player.position.name}}
                            </v-list-tile-content>
                        </v-list-tile>
                        <v-list-tile>
                            <v-list-tile-content>
                                <strong>Position Type</strong>
                            </v-list-tile-content>
                            <v-list-tile-content class="align-end">
                                {{player.position && player.position.type}}
                            </v-list-tile-content>
                        </v-list-tile>
                        <v-list-tile>
                            <v-list-tile-content>
                                <strong>Team</strong>
                            </v-list-tile-content>
                            <v-list-tile-content class="align-end">
                                {{player.team && player.team.name}}
                            </v-list-tile-content>
                        </v-list-tile>
                    </v-list>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>
<script>
import ApiClient from '../ApiClient'
import PlayerAvatar from './PlayerAvatar'

export default {
    components: { PlayerAvatar },
    data() {
        return {
            player: {},
            seasonStats: [],
            statsHeaders: []
        }
    },
    async beforeRouteEnter(to, from, next) {
        const slug = to.params.slug
        const result = await ApiClient.perform('get', `/players/player/${slug}`)
        console.log(result.player)
        
        next(vm => {
            vm.setPlayer(result.player)
            vm.setStatsHeaders(result.player.stats)
        })
    },
    async beforeRouteUpdate(to, from, next) {
        this.player = {}
        const slug = to.params.slug
        const result = await ApiClient.perform('get', `/players/player/${slug}`)

        this.setPlayer(result.player)
        this.setStatsHeaders(result.player.stats)
        console.log(result.player)
        next()
    },
    methods: {
        convertBinary: (binary) => {
            return "data:image/png;base64," + binary
        },
        setPlayer(player) {
            this.player = player
        },
        setStatsHeaders(stats) {
            this.statsHeaders = this.player.stats && Object.keys(this.player.stats).map(s => {
                return s === Object.keys(this.player.stats)[0] ? { text: "Season", value: this.player.stats && this.player.stats.season } : { text: s, value: s}
            })
        },
        customStatsSort(a,b){
            return ["season", "goals", "assists"].includes(a) ? 1 : -1
        }
    }
}
</script>