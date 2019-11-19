package main

type Person struct {
	UserId uint32 `json:"UserId,omitempty" bson:"UserId,omitempty"`
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Email  string `json:"email,omitempty" bson:"email,omitempty"`
	Mobile string `json:"mobile,omitempty" bson:"mobile,omitempty"`
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}
