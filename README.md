## Intro

A way to search for multiple permissions among different GCP roles. This tool performs ORed searches to match GCP roles for all the permissions you request in the search input box. Partial permission names can be used and are matched agains all the available permissions.
this tool is developed in typescript with svelte and sveltstrap on the front end. The backend is currently a proto quickly written in python 3 with flask

## Pre req

* npm with node js (v.12.22 is used here)
* python 3.9 with venv

## Setup

### setup front end
 
```
cd app
npm install
```

### setup backend

```
cd server
./bin/install -r requirements
```

## Run develop env

### start front end

```
cd app && npm run dev
```
This will run the frontend compiling the required static files into a server folder (under server/static/build)

### start the server

```
cd server/src && ./run-server.sh
```

This will run the flask devel server which will monitor even for changed static files (frontend) reloading the serving env after any modification either from frontend or from backend.

the server is accessible on localhost:5001

## Preseeding

In the current state the backend is a proto, you need to populate a pseudo DB by running the `role_db_builder.sh` script from the root of the repo. If you skip this step the backend wont run. To run the preseeding script you need to autenticate yourself into the GCP as the script makes use of the gcloud command.

```
mkdir ./server/src/roles
./role_db_builder.sh
```

this will take a bit of time depending on your machine and network latency. still you can interrupt the process and have a devel env even if not all the roles are downloaded from GCP. Just mind that killing the script can lead to some corrupted json files in the local pseudo DB, just remove them and the app will run.