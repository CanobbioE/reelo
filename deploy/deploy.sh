#!/usr/bin/env bash
docker-compose rm -f -s -v reelo-be
docker-compose rm -f -s -v reelo-fe

docker-compose pull

docker-compose up -d reelo-be
docker-compose up -d reelo-fe

