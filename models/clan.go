package models

import "gopkg.in/mgo.v2/bson"

type Clan struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	ClanId string        `json:"clan_id" bson:"clan_id"`
}
