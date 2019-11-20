package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Show struct {
    ShowID string `json:"ShowID,omitempty"    bson:"ShowID,omitempty"`
    TheatreID  string `json:"TheatreID,omitempty"    bson:"TheatreID,omitempty"`
    MovieID string `json:"MovieID,omitempty"    bson:"MovieID,omitempty"`
    //Users [5]string
}