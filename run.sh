#!/usr/bin/env bash

docker build . -t golang-demo
docker run -i -t -p 8080:9205 golang-demo