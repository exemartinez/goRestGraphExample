# goRestGraphExample
A simple implementation of a REST API on GOLang with access to Neo4j. This is just a simpler PoC made over what we did in https://github.com/exemartinez/SimpleMovieGraphAPI.

We added a few scripts in Docker and Docker compose, about "how" we believe it might be resolved. However, those scripts still need work over them.

This example just adds a new movie, with just its title.

We have been using Echo and Logrus, just as an example of library usage. 

##Assumptions
* There is a localhost Neo4J solution on 7687 with the credentials as: neo4j/secret
* You got all the drivers for Neo4J installed, as is explained in https://neo4j.com/developer/go/

