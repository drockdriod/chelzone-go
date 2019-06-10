import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';
import Team from './components/Team';
import Player from './components/Player';

Vue.use(VueRouter);

export default new VueRouter({
    mode: 'history',
    routes: [
        { path: '/', component: Home, props: true, name: "home" },
        { path: '/teams/:slug', name:"team", component: Team, props: true },
        { path: '/players/:slug', name:"player", component: Player },
    ],
    scrollBehavior (to, from, savedPosition) {
	    if (to.hash) {
	    	return {
		    	selector: to.hash
		    }
	    }
	    else if (savedPosition) {
		    return savedPosition
		} else {
		    return { x: 0, y: 0 }
		}

	}
});