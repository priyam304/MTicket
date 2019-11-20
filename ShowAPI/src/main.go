package main

import (
    "context"
    "fmt"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


func main() {

    //mongoURI := "mongodb://cmpe281:cmpe281@10.0.1.244:27017"

    fmt.Println("Starting the application")
    ctx,_:= context.WithTimeout(context.Background(),10*time.Second)

    //clientOptions:= options.Client().ApplyURI()
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client,_=mongo.Connect(ctx,clientOptions)
    router:=mux.NewRouter()
    router.HandleFunc("/show", CreateShowEndpoint).Methods("POST")
    router.HandleFunc("/show/{id}", GetShowEndpoint).Methods("GET")
    router.HandleFunc("/shows", GetAllShowsEndpoint).Methods("GET")
    router.HandleFunc("/show/{id}", DeleteShowEndpoint).Methods("DELETE")
    router.HandleFunc("/show/{id}", UpdateShowEndpoint).Methods("PUT")
    http.ListenAndServe(":12345",router)
    //"mongodb://localhost:27017"
}