#!/bin/sh

docker build -t anselbrandt/go-service:latest .
docker push anselbrandt/go-service:latest
ssh root@anselbrandt.dev << HERE
docker pull anselbrandt/go-service:latest
docker tag anselbrandt/go-service:latest dokku/go-service
dokku tags:deploy go-service
docker system prune -a
y
HERE

exit 0