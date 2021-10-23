## Intro

A way to search for multiple permissions among different GCP roles. This tool performs ORed searches to match GCP roles for all the permissions you request in the search input box. Partial permission names can be used and are matched agains all the available permissions.
this tool is developed in typescript with svelte and sveltstrap on the front end. The backend is currently a proto quickly written in go with gin

## Pre req

* npm with node js (v.12.22 is used here)
* python 1.17.2+

## Setup

### setup front end
 
```
cd app
npm install
```

### setup backend

```
cd server
go mod init
```

## Run develop env

### start front end

```
cd app && npm run dev
```
This will run the frontend compiling the required static files into a server folder (under server/static/build)

### start the server

```
cd server && go build -tags=jsoniter -o dsdr && ./dsdr
```

This will run the gin server. No hot reload of server or frontend yet, sorry.
The server is accessible on localhost:8080

## Preseeding

In the current state the backend is a proto, you need to populate a pseudo DB by running the `role_db_builder.sh` script from the root of the repo. If you skip this step the backend wont run. To run the preseeding script you need to autenticate yourself into the GCP as the script makes use of the gcloud command.

```
mkdir ./server/src/roles
./role_db_builder.sh
```

this will take a bit of time depending on your machine and network latency. still you can interrupt the process and have a devel env even if not all the roles are downloaded from GCP. Just mind that killing the script can lead to some corrupted json files in the local pseudo DB, just remove them and the app will run.