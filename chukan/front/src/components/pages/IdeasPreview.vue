<template>
	<div class="all_ideas">
		<h2>All Ideas</h2>
		<div v-for="idea in ideas" v-on:click="selectIdea(idea.id)" style="border:1px solid black;">
			<br/>
			<!-- <div>idea_id &#45; {{ idea.id }}</div> -->
			<div>title - {{ idea.title }}</div>
			<div>body - {{ idea.body }}</div>
			<br/>
		</div>
		<button v-on:click="openIdeaArea()"style="font-size:12px;">+</button>

		<div v-show="postArea">
			<textarea v-model="title" placeholder="Title"></textarea>
			<textarea v-model="body" placeholder="Body"></textarea>
			<button v-on:click="newIdeaPost()">Post</button>
		</div>
	</div> 
</template>
<script>
import axios from 'axios'

import router from '../../router.js'

const baseUrl = "http://localhost:1991/ideas"

export default{
	name:'IdeasPreview',
	data(){
		return{
			ideas:[],
			postArea:false,
			title:"",
			body:"",
		}
	},
	created(){
		axios.get(baseUrl)
			.then(resp => {
				console.log(resp);
				this.ideas = resp.data;
			}).catch(err=>{
				throw new Error(err)
			})
	},
	methods:{
		selectIdea(idea_id){
			console.log(idea_id)
			router.push({ path: '/ideas/'+idea_id })
		},
		openIdeaArea(){
			this.postArea = true;
		},
		newIdeaPost(){
			var datas = {
				'title':this.title,
				'body':this.body,
				// <!-- 'tag_ids':' -->
			}
			axios.post(baseUrl,{

			},data)
				.then(resp => {

				})
		}
	}
}

</script>
<style>
</style>

