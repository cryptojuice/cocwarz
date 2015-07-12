package wars

import (
	"encoding/json"
	"fmt"
	"github.com/cryptojuice/cocwarz/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type WarHandler struct {
	session *mgo.Session
}

func NewWarHandler(s *mgo.Session) *WarHandler {
	return &WarHandler{s}
}

func (wh WarHandler) Create(w http.ResponseWriter, r *http.Request) {
	newWar := models.War{}
	json.NewDecoder(r.Body).Decode(newWar)

	newWar.Id = bson.NewObjectId()

	wh.session.DB("cocwarz").C("wars").Insert(newWar)
	wj, _ := json.Marshal(newWar)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", wj)
}

func (wh WarHandler) Index(w http.ResponseWriter, r *http.Request) {
}

func (wh WarHandler) Show(w http.ResponseWriter, r *http.Request) {
}

func (wh WarHandler) Update(w http.ResponseWriter, r *http.Request) {
}

func (wh WarHandler) Destroy(w http.ResponseWriter, r *http.Request) {
}
