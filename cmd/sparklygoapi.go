package main

import (
    "fmt"    
    "encoding/json"
    "github.com/sirupsen/logrus"
    "net/http"
    "os"
    "github.com/neo4j/neo4j-go-driver/neo4j"  
    "github.com/labstack/echo"  
)

//Declaration of package variables
var (
    err      error
    driver   neo4j.Driver
    session  neo4j.Session
    result   neo4j.Result
    movie interface{}
)

var log = logrus.New()
var file_log os.File

//Constants for connectivity (This shouldn't be here)
const SERVERDB = "bolt://localhost:7687"
const USER = "neo4j"
const PASS = "secret"

//TODO: parse the JSON receives in the request's body into this Struct.
type MovieStruct struct{
    id int `json:"id"`
    title string `json:"title"`
    released string `json:"released"`
    tagline string `json:"tagline"`
    roles string `json:"roles"`
}

func init() {

    // open a file
    file_log, err := os.OpenFile("./log/sparkly_api.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
        fmt.Printf("error opening file: %v", err)
    }

    // Log as JSON instead of the default ASCII formatter.
    logrus.SetFormatter(&logrus.JSONFormatter{})

    // Output to stderr instead of stdout, could also be a file.
    logrus.SetOutput(file_log)

    // Only log the warning severity or above.
    logrus.SetLevel(logrus.DebugLevel)

}


//API implementation with the ECHO library.
func main() {
    e := echo.New()

    // don't forget to close the log file.
    defer file_log.Close()

    //logrus.SetOutput(os.Stdout) //uncomment this to send the logging to the standar output (the console).

    logrus.WithFields(logrus.Fields{"App":"sparkly","function":"main"}).Info("Initiating the REST API.")


    // This is just an example for Golang of adding a new  node. By no means is  complete - TODO: Extract a method/function
    e.POST("/movies", func(c echo.Context) error {

        title:=""

        json_map := make(map[string]interface{})

        err := json.NewDecoder(c.Request().Body).Decode(&json_map)
        if err != nil {
            return err
        } else {
            //json_map has the JSON Payload decoded into a map - just for an example and lake of time, we work with the movie's title only.
            title = json_map["title"].(string)            
        }
        

        logrus.WithFields(logrus.Fields{"App":"sparkly","function":"main"}).Info("The movie title: " + title)

        result, err := addNewMovie(title)    //we just send the name - still at baby steps in golang...


        defer driver.Close() 
        defer session.Close()

        if err != nil {

            logrus.WithFields(logrus.Fields{"App":"sparkly","function":"main"}).Error(err)
            return c.String(http.StatusInternalServerError, "")
        }

        return c.String(http.StatusOK, result)
        
    })


    e.Logger.Fatal(e.Start(":8093"))

}

//This is the function that will interact with Neo4J - Here occurs the transaction.
func addNewMovie(newName string) (string, error) {

    driver, session, err = neo4jConnect(SERVERDB,USER,PASS)

    //Here we perform the proper transaction execution.
    movie, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

        result, err = transaction.Run(
            "CREATE (a:Movie {name: $name}) RETURN a.name",
            map[string]interface{}{"name": newName}) //cypher create command here...

        //Error handling here...
        if err != nil {
            logrus.WithFields(logrus.Fields{"App":"sparkly","function":"addNewMovie"}).Error(err)
            return nil, err
        }

        //Consuming the return (we putter that in the create as well, just like a message)
        if result.Next() {
            return result.Record().GetByIndex(0), nil
        }

        return nil, result.Err()
    }) // transaction ends here...

    if err != nil {
        logrus.WithFields(logrus.Fields{"App":"sparkly","function":"addNewMovie"}).Error(err)
        return "", err
    }

    return movie.(string), nil
}


//Just the equivalent to the "catch" section in a regular java Try-catch block.
func errorHandler(err error){

    if err!=nil {
        logrus.WithFields(logrus.Fields{"App":"sparkly","function":"errorHanlder"}).Error(err)
        fmt.Println(err)
        os.Exit(1)
    }

}

// Here we start the database connection...
func neo4jConnect(uri, username, password string)(neo4j.Driver, neo4j.Session, error){

    //We start the driver here, we need to pass it the neo4j url, usernama and password to connect with.
    driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))

    //Any error handler
    if err != nil {
        logrus.WithFields(logrus.Fields{"App":"sparkly","function":"neo4jConnect"}).Error(err)
        return nil,nil, err
    }

    //Connection start up...
    session, err = driver.Session(neo4j.AccessModeWrite)
    if err != nil {
        logrus.WithFields(logrus.Fields{"App":"sparkly","function":"neo4jConnect"}).Error(err)
        return nil, nil, err
    }

    return driver, session, nil
}


