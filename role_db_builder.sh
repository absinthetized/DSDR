#!/bin/bash

for role in $(gcloud iam roles list | grep "^name:" | cut -f 2 -d : | xargs); do gcloud iam roles describe $role --format=json > server/src/$role; done
