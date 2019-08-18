import Vue from 'vue'
import VueRouter from 'vue-router'

import Auth from './components/pages/Auth.vue'
import Idea from './components/pages/Idea.vue'
import IdeasPreview from './components/pages/IdeasPreview.vue'
import HelloWorld from './components/HelloWorld.vue'

Vue.use(VueRouter)

const router = new VueRouter({
	mode:'history',
	routes:[
		{
			path:'/',
			componets:Auth
		},
		{
			path:'/ideas/:idea_id',
			component:Idea
		},
		{
			path:'/ideas',
			component:IdeasPreview
		}
	]
})

export default router
