package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"hash/fnv"
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

// type Person struct {
// 	UserId string `json:"UserId,omitempty" bson:"UserId,omitempty"`
// 	Name   string `json:"name,omitempty" bson:"name,omitempty"`
// 	Email  string `json:"email,omitempty" bson:"email,omitempty"`
// 	Mobile string `json:"mobile,omitempty" bson:"mobile,omitempty"`
// }

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

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

	person.UserId = hash(person.Name + person.Email + person.Mobile)

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
	var person []Person
	//_ = json.NewDecoder(request.Body).Decode(&person)
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := (params["id"])
	fmt.Println(id)
	err = c.Find(bson.M{"UserId": id}).All(&person)
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode(person)

	//------------------LOCALHOST CODE--------------------

	// response.Header().Set("content-type", "application/json")
	// params := mux.Vars(request)
	// id, _ := primitive.ObjectIDFromHex(params["id"])
	// var person Person
	// collection := client.Database("jay").Collection("people")
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// json.NewEncoder(response).Encode(person)
}

func RemovePersonEndpoint(response http.ResponseWriter, request *http.Request) {

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
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := (params["id"])
	fmt.Println(id)
	err = c.Remove(bson.M{"UserId": id})
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode("Deleted Record")

	//------------------LOCALHOST CODE--------------------

	// response.Header().Set("content-type", "application/json")
	// params := mux.Vars(request)
	// id, _ := primitive.ObjectIDFromHex(params["id"])
	// var person Person

	// collection := client.Database("jay").Collection("people")
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// filter := Person{ID: id}
	// result, err := collection.DeleteOne(ctx, filter)
	// fmt.Println(result.DeletedCount)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// json.NewEncoder(response).Encode(person)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {

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
	var person []Person

	response.Header().Set("content-type", "application/json")
	err = c.Find(bson.M{}).All(&person)
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode(person)

	//------------------LOCALHOST CODE--------------------

	// response.Header().Set("content-type", "application/json")
	// var people []Person
	// collection := client.Database("jay").Collection("people")
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// cursor, err := collection.Find(ctx, bson.M{})
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	var person Person
	// 	cursor.Decode(&person)
	// 	people = append(people, person)
	// }
	// if err := cursor.Err(); err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// json.NewEncoder(response).Encode(people)
}

func sendmail(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Sending mail")

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := (params["id"])
	var person []Person

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
	response.Header().Set("content-type", "application/json")
	fmt.Println(id)
	err = c.Find(bson.M{"UserId": id}).All(&person)
	if err != nil {
		fmt.Println("Can not Send")
	}
	fmt.Println("Sending........")
	to2 := person[0].Email
	from := "j.pathak.6732@gmail.com"

	// use we are sending email to
	to := to2

	// server we are authorized to send email through
	host := "smtp.gmail.com"

	// Create the authentication for the SendMail()
	// using PlainText, but other authentication methods are encouraged
	auth := smtp.PlainAuth("", from, "John6732", host)

	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style
	message := `To: "Some User"
	From: "Other User"
	Subject: Yayy Ticket Booked..!!

	This is the message we are sending. That's it!
	`

	if err := smtp.SendMail(host+":25", auth, from, []string{to}, []byte(message)); err != nil {
		fmt.Println("Error SendMail: ", err)
		os.Exit(1)
	}
	fmt.Println("Email Sent!")
}

func sendsms(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Sending sms")
	AccountSid := "AC6e3ee3aa9f0e15dc727f715cf9d05838"
	AuthToken := "01ae941c8c7f80e1a1f07a7a874d6d03"
	From := "+12015146384"

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := (params["id"])
	var person []Person

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
	response.Header().Set("content-type", "application/json")
	fmt.Println(id)
	err = c.Find(bson.M{"UserId": id}).All(&person)
	if err != nil {
		fmt.Println("Can not insert")
	}
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	to2 := person[0].Mobile
	To := to2
	// Initialize twilio Client
	c2 := twilio.NewClient(AccountSid, AuthToken, nil)

	// Send Message
	params2 := twilio.MessageParams{
		Body: "Hello Go!",
	}

	s, resp, err := c2.Messages.Send(From, To, params2)
	log.Println("Send:", s)
	log.Println("Response:", resp)
	log.Println("Err:", err)

}
