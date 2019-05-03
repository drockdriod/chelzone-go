<template>
    <div>
        <v-parallax :src="rink">
            <v-container grid-list-md>
                <v-layout text-xs-center wrap>
                    <v-flex mb-4>
                        <h1 class="display-2 font-weight-bold mb-3">
                            Welcome to Chelzone
                        </h1>
                        <p class="subheading font-weight-regular">
                            Get up to date stats, standings, in game action and more!
                        </p>
                    </v-flex>

                    <v-flex xs12 sm12 md6 lg6 grid-list-md>
                        <LeaderStats v-bind:leaders="leaders" />
                    </v-flex>
                    <v-flex xs12 sm12 md12 lg12 grid-list-md>
                        <Divisions v-bind:teams="teamsByDivision" />
                    </v-flex>
                </v-layout>
            </v-container>
        </v-parallax>
        <v-container>
            <SocialMosaic />
        </v-container>
    </div>
</template>
<script>
    import { mapGetters, mapActions } from 'vuex'

    import Divisions from './Divisions'
    import SocialMosaic from './SocialMosaic'
    import LeaderStats from './LeaderStats'
    
    import rink from "../assets/images/rink.jpg"

    export default {
        components: { Divisions, SocialMosaic, LeaderStats },
        computed: mapGetters(['teams', 'teamsByDivision', 'leaders']),
        created() {
            this.getTeams();
            this.getLeaders();
        },
        data () {
            return {
                publicPath: process.env.BASE_URL,
                rink
            }
        },
        methods: {
            ...mapActions([
                'getTeams',
                'getLeaders'
            ])
        }
    }
</script>
<style>
    .v-parallax{
        height: auto !important
    }
</style>