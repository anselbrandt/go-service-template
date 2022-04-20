#!/bin/sh

docker build -t anselbrandt/go-service:latest .
docker push anselbrandt/go-service:latest
ssh root@anselbrandt.dev << HERE
dokku apps:create go-service
dokku domains:clear-global
dokku domains:set go-service anselbrandt.dev
dokku proxy:ports-set go-service http:80:8080
dokku certs:add go-service < cert-key.tar
docker pull anselbrandt/go-service:latest
docker tag anselbrandt/go-service:latest dokku/go-service
dokku tags:deploy go-service
docker system prune -a
y
HERE

exit 0
