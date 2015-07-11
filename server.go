package main

import (
	"gopkg.in/mgo.v2"
	"net/http"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:17017")

	if err != nil {
		panic(err)
	}

	return s
}

func main() {
	router := NewRouter()
	http.ListenAndServe(":8080", router)
}
