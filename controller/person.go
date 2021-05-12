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

func (pc *PersonController) FetchPerson(id int) model.Person {
	pc.Log.Debug("Fetch person by id")

	var person model.Person
	var books []model.Book
	pc.DBInstance.DB.First(&person, id)
	pc.DBInstance.DB.Where("person_id=?", id).Find(&books)
	person.Books = books
	return person
}

func (pc *PersonController) SaveNewPerson(p model.Person) error {
	result := pc.DBInstance.DB.Create(&p)
	return result.Error
}
