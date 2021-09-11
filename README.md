## Intro

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
This will run the front end compiling the required static files into the server folder

### start the server

```
cd src && ./run-server.sh
```

This will run the flask devel server which will monitor even for changed static files (front end) reloading the hot reloading serving env after any modification either from front end or from backend

## Preseeding

In the current state the backend is a proto, you need to populate a pseudo DB by running the `role_db_builder.sh` script from the root of the repo. If you skip this step the backend wont run. To run the preseeding script you need to autenticate yourself into the GCP as the script uses the gcloud command.

```
mkdir ./server/src/roles
./role_db_builder.sh
```

this will take a bit of time depending on your machine and network latency