package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"github.com/unrolled/render"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "12345"
	}

	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//clientOptions := options.Client().ApplyURI("mongodb://cmpe281:cmpe281@3.89.47.220:27017")
	clientOptions := options.Client().ApplyURI("mongodb+srv://jay:jay@movies-upn2q.mongodb.net/test?retryWrites=true&w=majority")
	fmt.Println("Client Options set...")
	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Mongo Connected...")
	router := mux.NewRouter()

	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")

	http.ListenAndServe(":"+port, router)
}
