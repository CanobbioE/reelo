#!/usr/bin/env bash
docker rmi canobbioe/reelo-fe

docker build -t canobbioe/reelo-fe .

docker push canobbioe/reelo-fe
