#!/bin/bash

go clean

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ap0001_mongo_engine cmd/main.go

docker build -t vickeyshrestha/ap0001_mongo_engine:$1 .

echo "Process finished"