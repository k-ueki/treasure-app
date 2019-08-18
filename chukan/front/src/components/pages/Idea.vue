<template>
	<div class="">
		<h2>Idea</h2> <div v-show="data" style="border:1px solid black;">
			<div>title - {{ ideaTitle }}</div>
			<div>body - {{ ideaBody }}</div>
			<div>tags - {{ tags }}</div>
		</div>

		<h2>Comments</h2>
		<div v-for="comment in comments" style="border:1px solid black;">
			<div>comment body - {{ comment.body }}</div>
			<button>X</button>
		</div>
		<button v-show="!postArea"v-on:click="openCommentArea">New Comment</button>

		<div v-show="postArea">
			<textarea v-model="newComment"></textarea>
			<button v-on:click="newPost()">投稿</button>
		</div>
	</div> 
</template>
<script>
import axios from 'axios'

const baseUrl = "http://localhost:1991/ideas"

export default{
	name:'Ideas',
	pops:{
		// Title:String,
		// Body:String,
	},
	hook:{
	},
	data(){
		return{
			ideaTitle:"",	
			ideaBody:"",
			data:false,
			comments:[],
			tags:[],

			newComment:"",

			postArea:false,
		}
	},
	created(){
		console.log(this.$route.params.id);
		// axios.get(baseUrl+"/"+this.$route.params.id)
		axios.get(baseUrl+"/1")
			.then(resp => {
				console.log(resp);
				this.ideaTitle = resp.data.title;
				this.ideaBody = resp.data.body;
				this.comments = resp.data.comments;
				this.tags = resp.data.tags;
				console.log(this.comments)
				this.data = true;
			})
	},
	methods:{
		newPost(){
			console.log("OK");
		},
		openCommentArea(){
			this.postArea=true;
		}
	}
}

</script>
<style>
</style>
