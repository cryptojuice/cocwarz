package models

import "gopkg.in/mgo.v2/bson"

type Target struct {
	Id       bson.ObjectId
	Name     string
	Called   bool
	Attacker []Attacker
}
