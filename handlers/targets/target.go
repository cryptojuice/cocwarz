package targets

import (
	// "encoding/json"
	// "fmt"
	// "github.com/cryptojuice/cocwarz/models"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"net/http"
)

type TargetHandler struct {
	session *mgo.Session
}

func NewTargetHandler(s *mgo.Session) *TargetHandler {
	return &TargetHandler{s}
}

func (th TargetHandler) Index(w http.ResponseWriter, r *http.Request) {
}

func (th TargetHandler) Create(w http.ResponseWriter, r *http.Request) {
}

func (th TargetHandler) Update(w http.ResponseWriter, r *http.Request) {
}
