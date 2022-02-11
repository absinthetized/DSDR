// Init Svelte

import App from './App.svelte';
const app = new App({
	target: document.body,
});

// Initialize Firebase

import { initializeApp } from "firebase/app";
import firebaseConfig from './firebase.config';
const fapp = initializeApp(firebaseConfig);

// Install sign-in/sign-out event listener

import { getAuth, onAuthStateChanged } from "firebase/auth";
import currentUser from "./models/user"

const auth = getAuth();
onAuthStateChanged(auth, (userInfo) => {
  if (userInfo) {
    currentUser.info = userInfo
    console.log('current user is: ' + currentUser.info.email)
  
} else {
    console.log('no user currently logged in')
	currentUser.info = null
  }
});

export default app;