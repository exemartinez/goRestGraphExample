#!/bin/bash
#To get the neo4j drivers - this script should be run from the go project's folder.
go get github.com/neo4j/neo4j-go-driver/neo4j

#To build the docker image and run it right away.
docker build -t sparklygoapi .
docker run -p 8093:8093 sparklygoapi