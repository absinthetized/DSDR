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
import type { User } from "firebase/auth";

let currentUser: User

const auth = getAuth();
onAuthStateChanged(auth, (user) => {
  if (user) {
    // User is signed in, see docs for a list of available properties
    // https://firebase.google.com/docs/reference/js/firebase.User
    const uid = user.uid;
    // ...

    currentUser = user
    console.log('current user is: ' + currentUser.email)
  } else {
    // User is signed out
    // ...
    console.log('no user currently logged in')
  }
});

export default app;