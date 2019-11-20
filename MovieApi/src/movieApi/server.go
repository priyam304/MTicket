
package main

import(
	
	"context"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

)


var mongoURI = "mongodb://localhost:27017"
var client *mongo.Client


//Function to GET a particular movie in a database
func GetMovieEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	name, _ := params["name"]
	fmt.Printf("Name : %q" , name )
	var movie Movie
	collection := client.Database("Movie").Collection("Movies")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Movie{Name: name}).Decode(&movie)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(movie)
}


//Function to POST a particular movie in a database
func CreateMovieEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	collection := client.Database("Movie").Collection("Movies")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, movie)

	if result == nil{
		json.NewEncoder(response).Encode("{'message':'Already Exists'}")
		} else{

			json.NewEncoder(response).Encode(result)
		}	
}


//Function to Delete a particular movie in a database
func DeleteMovieEndpoint(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	name, _ := params["name"]
	filter := bson.D{{"name",name}}
	collection := client.Database("Movie").Collection("Movies")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
    log.Fatal(err)
	}
	//fmt.Printf("Deleted %v documents in the Movies collection\n", deleteResult.DeletedCount)

	json.NewEncoder(response).Encode(deleteResult)
}




//Function to Update a particular movie in a database 
func UpdateMovieEndpoint(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")

	var movie Movie
	params := mux.Vars(request)
	name, _ := params["name"]
	_ = json.NewDecoder(request.Body).Decode(&movie)

	// fmt.Printf("newName : %q" , movie.Name )
	// fmt.Printf("newRating : %q" , movie.Rating )

	filter := bson.D{{"name",name}}
	update := bson.D{
	{ "$set",  bson.D{
		{"name", movie.Name } ,{"rating", movie.Rating} ,
	}},
	}
	collection := client.Database("Movie").Collection("Movies")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
    log.Fatal(err)
	}

	//fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
 	json.NewEncoder(response).Encode(updateResult)
}




//Function to get all movies from the movies database
func GetAllMoviesEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var movies []Movie
	collection := client.Database("Movie").Collection("Movies")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var movie Movie
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(movies)
}





