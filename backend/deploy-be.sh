#!/usr/bin/env bash
docker rmi canobbioe/reelo-be

docker build -t canobbioe/reelo-be .

docker push canobbioe/reelo-be
