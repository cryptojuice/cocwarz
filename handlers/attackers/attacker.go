package attackers

import (
	// "encoding/json"
	// "fmt"
	// "github.com/cryptojuice/cocwarz/models"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"net/http"
)

type AttackerHandler struct {
	session *mgo.Session
}

func NewAttackerHandler(s *mgo.Session) *AttackerHandler {
	return &AttackerHandler{s}
}

func (ah AttackerHandler) Index(w http.ResponseWriter, r *http.Request) {
}

func (ah AttackerHandler) Create(w http.ResponseWriter, r *http.Request) {
}

func (ah AttackerHandler) Update(w http.ResponseWriter, r *http.Request) {
}

func (ah AttackerHandler) Destroy(w http.ResponseWriter, r *http.Request) {
}
