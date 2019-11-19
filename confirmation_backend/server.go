package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"net/http"
	"net/smtp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/subosito/twilio"
)

var client *mongo.Client
var mongodb_server = "mongodb+srv://jay:@movies-upn2q.mongodb.net/test?retryWrites=true&w=majority"
var mongodb_server_1 = "52.37.128.85:27017"
var mongodb_database = "movies"
var mongodb_collection = "people"
var mongodb_collection1 = "submissions"
var mongodb_username = "jay"
var mongodb_password = "jay"

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Inside get assignemts function")

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: []string{"movies-shard-00-00-upn2q.mongodb.net:27017",
			"movies-shard-00-01-upn2q.mongodb.net:27017",
			"movies-shard-00-02-upn2q.mongodb.net:27017"},
		Database: "admin",
		Username: "jay",
		Password: "jay",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongodb_database).C(mongodb_collection)
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	err = c.Insert(person)
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode("OK")

	//------------------LOCALHOST CODE--------------------

	// var assignments_array []assignment
	// err = c.Find(bson.M{}).All(&assignments_array)
	// fmt.Println("Assignments", assignments_array)
	//formatter.JSON(w, http.StatusOK, assignments_array)

	// response.Header().Set("content-type", "application/json")
	// var person Person
	// _ = json.NewDecoder(request.Body).Decode(&person)
	// fmt.Println("Name", person.Name)
	// collection := client.Database("movies").Collection("people")
	// fmt.Println("Name", person.Name)
	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// result, _ := collection.InsertOne(ctx, person)
	// json.NewEncoder(response).Encode(result)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {


	//------------------LOCALHOST CODE--------------------

	 response.Header().Set("content-type", "application/json")
	 params := mux.Vars(request)
	 id, _ := primitive.ObjectIDFromHex(params["id"])
	 var person Person
	 collection := client.Database("jay").Collection("people")
	 ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	 err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	 if err != nil {
	 	response.WriteHeader(http.StatusInternalServerError)
	 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	 	return
	 }
	 json.NewEncoder(response).Encode(person)
}
