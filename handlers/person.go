package handelers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/jaskeerat789/go-postgres-webserver/controller"
	"github.com/jaskeerat789/go-postgres-webserver/model"
)

type PersonHandler struct {
	Log hclog.Logger
	pc  *controller.PersonController
}

func NewPersonHandler(l hclog.Logger, db *model.DB) *PersonHandler {
	pc := controller.NewPersonController(db, l)
	return &PersonHandler{
		Log: l,
		pc:  pc,
	}
}

func (ph *PersonHandler) GetPeople(rw http.ResponseWriter, r *http.Request) {
	ph.Log.Debug("Get all people handler")
	rw.Header().Add("Content-Type", "application/json")
	people := ph.pc.FetchAllPerson()
	people.ToJSON(rw)
}

func (ph *PersonHandler) GetPerson(rw http.ResponseWriter, r *http.Request) {
	ph.Log.Debug("Get a person by id handler")
	rw.Header().Add("Content-Type", "application/json")

	id, err := getPersonId(r)
	if err != nil {
		ph.Log.Error("Failed to get valid Person Id", err)
		http.Error(rw, "Need a valid person id", http.StatusInternalServerError)
		return
	}
	person := ph.pc.FetchPerson(id)
	person.ToJSON(rw)

}

func (ph *PersonHandler) CreatePerson(rw http.ResponseWriter, r *http.Request) {
	ph.Log.Debug("Create person handler")
	person := model.NewPerson(r)
	err := ph.pc.SaveNewPerson(person)
	if err != nil {
		ph.Log.Error("Failed to save person", err)
		http.Error(rw, "Failed to save new person", http.StatusInternalServerError)
	}
	person.ToJSON(rw)
}

func getPersonId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}
