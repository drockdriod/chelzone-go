import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import colors from 'vuetify/es5/util/colors'

import 'vuetify/src/stylus/app.styl'
import '@fortawesome/fontawesome-free/css/all.css'

import '../assets/css/index.css'


Vue.use(Vuetify, {
	iconfont: 'md',
  	icons: {
  		'twitter': 'fab fa-twitter',
  		'youtube': 'fab fa-youtube'
  	},
	theme: {
		primary: colors.indigo.base,
		secondary: colors.lime.base,
		accent: colors.amber.base,
		error: colors.red.base,
		warning: colors.orange.base,
		info: colors.cyan.base,
		success: colors.green.base
	}
})
