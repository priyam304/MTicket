package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Movie represents the Movie structure which has fields defined below
type Movie struct{

	Id     primitive.ObjectID `json:"_id,omitempty"    bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty"   bson:"name,omitempty"`
	Rating string             `json:"rating,omitempty" bson:"rating,omitempty"`
	

}

//Movies is an array of type Movie
type Movies []Movie
