package models

import "gopkg.in/mgo.v2/bson"

type Attacker struct {
	Id           bson.ObjectId
	AttackerName string
	stars        int
}
