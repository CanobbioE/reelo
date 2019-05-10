#!/usr/bin/env bash
docker-compose rm -f -s -v

docker-compose pull

docker-compose up -d

