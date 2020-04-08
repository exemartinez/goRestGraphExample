# goRestGraphExample
A simple implementation of a REST API on GOLang with access to Neo4j. This is just a simpler PoC made over what we did in https://github.com/exemartinez/SimpleMovieGraphAPI.

We added a few scripts in Docker and Docker compose, about "how" we believe it might be resolved. However, those scripts still need work over them.

This example just adds a new movie, with just its title.

We have been using Echo and Logrus, just as an example of library usage. 

## Assumptions
* There is a localhost Neo4J solution on 7687 with the credentials as: neo4j/secret.
* You got golang installed and your $GOPATH variables appropiately configured.
* You got all the drivers for Neo4J installed, as is explained in https://neo4j.com/developer/go/
* Dockerfile & Docker-Compose: the docker file that you might need to run the go binary or for a later upload. These are just samples and need still some work to make it work appropiately.
* sparklygoapi.go: here goes the simple standalone app, implemented with juest ONE REST service of the ones implemented in the Java version. This is done as a mode of example of what can be achieved in Golang in terms of simplicity. The complete explanation of this file and design can be found in the [pdf](https://github.com/exemartinez/SimpleMovieGraphAPI/blob/master/Project%20Proposal%20-%20Solution%20Architecture.pdf), section "Final Words".

## Run
In order to run the example, you'll need:
1. Placed on $GOPATH/src/github.com/exemartinez/goRestGraphExample. Execute the following command:
    ```
     go run cmd/sparklygoapi.go
     ```
This will run a local server that will hit the Neo4J pre-configured database.

2. Use the postman collection to test the call to the golang service. This will just run ONE call that adds a movie by its title.


