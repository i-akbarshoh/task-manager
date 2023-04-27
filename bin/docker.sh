#!/bin/bash

# This script is used to login to docker hub
# It is used by the CI/CD pipeline to login to docker hub
# and push the docker image to docker hub

# Login to docker hub
docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD

# Build the docker image
docker build -t $DOCKER_USERNAME/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG .

# Push the docker image to docker hub
docker push $DOCKER_USERNAME/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG

# Logout from docker hub
docker logout

# Exit with status code 0
exit 0
