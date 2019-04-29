<template>
    <v-container>
        <v-layout text-xs-center wrap>
            <v-flex mb-4>
                <h1 class="display-2 font-weight-bold mb-3">
                    Welcome to Vuetify
                </h1>
                <p class="subheading font-weight-regular">
                    For help and collaboration with other Vuetify developers,
                    <br>please join our online
                    <a href="https://community.vuetifyjs.com" target="_blank">Discord Community</a>
                </p>
            </v-flex>
            <Divisions v-bind:teams="teamsByDivision" />
        </v-layout>
    </v-container>
</template>
<script>
    import { mapGetters, mapActions } from 'vuex'
    import Divisions from './Divisions'
    export default {
        components: { Divisions },
        computed: mapGetters(['teams', 'teamsByDivision']),
        created() {
            this.getTeams();
        },
        methods: {
            ...mapActions([
                'getTeams',
            ])
        }
    }
</script>
<style>
</style>