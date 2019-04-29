<template>
    <v-app>
        <v-toolbar app>
            <v-toolbar-title class="headline text-uppercase">
                <span>Chelzone</span>
                <span class="font-weight-light">Stats and Such</span>
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

export default {
    name: 'App',
    data() {
        return {
            //
        }
    }
}
</script>