<template>
    <v-app>
        <v-toolbar app dark>
            <v-toolbar-title class="headline text-uppercase" @click="navigate('home')">
                <span>Chelzone</span>
                <span class="font-weight-light"> Hockey First</span>
            </v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn flat href="https://github.com/vuetifyjs/vuetify/releases/latest" target="_blank">
                <span class="mr-2">Latest Release</span>
            </v-btn>
        </v-toolbar>

        <v-content>
            <router-view></router-view>
        </v-content>
    </v-app>
</template>
<script>

import { mapGetters, mapActions } from 'vuex'
export default {
    name: 'App',
    data() {
        return {
            //
        }
    },
    created() {
        this.setLoading(true)
        Promise.all([
            this.$store.dispatch('getSocialMedia',0),
            this.getLeaders()
        ]).then(() => {
            this.getTeams()
            this.getGameMilestones()
            this.setLoading(false)
        })
    },
    methods: {
        navigate(name, params){
            this.$router.push({ name, params})
        },
        ...mapActions([
            'getTeams',
            'getLeaders',
            'getSocialMedia',
            'setLoading',
            'getGameMilestones'
        ])
    }
}
</script>
<style>
    .headline {
        cursor: pointer;
    }
</style>