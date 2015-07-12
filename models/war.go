package models

import "gopkg.in/mgo.v2/bson"

type War struct {
	Id            bson.ObjectId `json:"id" bson: "_id"`
	ClanId        string        `json:"clan_id" bson: "clan_id"`
	ClanName      string        `json:"clan_name" bson: "clan_name"`
	EnemyClanName string        `json:"enemy_clan_name" bson: "enemy_clan_name"`
	Size          int           `json:"size" bson: "size"`
	Target        []Target
}
