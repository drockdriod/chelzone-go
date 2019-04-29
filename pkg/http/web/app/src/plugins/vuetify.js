import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import 'vuetify/src/stylus/app.styl'
import colors from 'vuetify/es5/util/colors'

Vue.use(Vuetify, {
	iconfont: 'md',
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
