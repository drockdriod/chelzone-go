<template>
    <v-container grid-list-md>
        <v-layout align-end justify-end row fill-height row>
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
            player: {}
        }
    },
    async beforeRouteEnter(to, from, next) {
        const slug = to.params.slug
        const result = await ApiClient.perform('get', `/players/player/${slug}`)
        next(vm => vm.setPlayer(result.player))
    },
    async beforeRouteUpdate(to, from, next) {
        this.player = {}
        const slug = to.params.slug
        const result = await ApiClient.perform('get', `/players/player/${slug}`)

        this.setPlayer(result.player)
        next()
    },
    methods: {
        convertBinary: (binary) => {
            return "data:image/png;base64," + binary
        },
        setPlayer(player) {
            this.player = player
        }
    }
}
</script>