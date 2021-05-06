package handelers

import (
	"net/http"

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
