<template>
	<div class="auth">
		<button v-if="stateUser==null" v-on:click="login()">Pleade login</button>
		<div>aoiep</div>
	</div>
</template>
<script>
import firebase from "firebase/app"
import "firebase/auth"

const firebaseConfig = {
  apiKey: process.env.FIREBASE_APIKEY,
  authDomain: process.env.FIREBASE_AUTHDOMAIN,
  databaseURL: process.env.FIREBASE_DATABASEURL,
  projectId: process.env.FIREBASE_PROJECTID,
  messagingSenderId: process.env.FIREBASE_MESSAGINGSENDERID,
  appId: process.env.FIREBASE_APPID
};

const githubProvider = new firebase.auth.GithubAuthProvider();

export default {
  name: 'Auth',
	data(){
	},
	methods:{
		login:function(){
  		firebase.initializeApp(firebaseConfig);
			firebase.auth().signInWithPopup(githubProvider);
		},
		logout:function(){
			firebase.auth().signOut();
		}
	}

}


</script>
<style>
</style>
