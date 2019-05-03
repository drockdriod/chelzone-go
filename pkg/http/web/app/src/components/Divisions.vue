<template>
    <v-layout row wrap>
        <v-toolbar class="mb-2" color="indigo darken-5" dark flat>
            <v-toolbar-title>Divisions</v-toolbar-title>
        </v-toolbar>
        <v-flex xs12 sm6 md4 lg3 v-for="(item,index) in teams" v-bind:key="index">
            <v-card>
                <v-card-title class="subheading font-weight-bold">{{ item.name }}</v-card-title>
                <v-divider></v-divider>
                <v-list dense>
                    <v-list-tile v-for="(item,index2) in item.teams" v-bind:key="index2">
                        <v-list-tile-content>
                            <v-avatar color="white">
                                    <img v-bind:src="convertBinary(item.logo.Data)" alt="avatar">
                            </v-avatar>
                        </v-list-tile-content>
                        <v-list-tile-content class="align-end">{{ item.name }}</v-list-tile-content>
                    </v-list-tile>
                </v-list>
            </v-card>
        </v-flex>
    </v-layout>
</template>
<script>
export default {
    props: ["teams"],
    methods: {
        convertBinary: (binary) => {
            return "data:image/svg+xml;base64,"+binary
        }
    }
}
</script>