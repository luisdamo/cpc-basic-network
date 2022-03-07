#!/bin/bash

# if cli container already exists run:
# docker stop cli
# docker rm cli

echo "starting client container to interact with the network"
docker-compose up -d cli
docker exec -it cli bash