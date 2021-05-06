package controller

import (
	"github.com/hashicorp/go-hclog"
	"github.com/jaskeerat789/go-postgres-webserver/model"
)

type PersonController struct {
	Log        hclog.Logger
	DBInstance *model.DB
}

func NewPersonController(db *model.DB, l hclog.Logger) *PersonController {
	return &PersonController{
		Log:        l,
		DBInstance: db,
	}
}

func (pc *PersonController) FetchAllPerson() model.People {
	pc.Log.Debug("Fetch all person controller")

	var people model.People
	pc.DBInstance.DB.Find(&people)
	return people
}
