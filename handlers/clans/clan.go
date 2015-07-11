package clans

import (
	"encoding/json"
	"fmt"
	"github.com/cryptojuice/cocwarz/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type ClanHandler struct {
	session *mgo.Session
}

func NewClanHandler(s *mgo.Session) *ClanHandler {
	return &ClanHandler{s}
}

func (ch ClanHandler) Create(w http.ResponseWriter, r *http.Request) {
	c := models.Clan{}
	json.NewDecoder(r.Body).Decode(&c)
	c.Id = bson.NewObjectId()
	ch.session.DB("cocwarz").C("clans").Insert(c)

	cj, _ := json.Marshal(c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", cj)
}
