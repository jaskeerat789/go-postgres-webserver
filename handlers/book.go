package handelers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/jaskeerat789/go-postgres-webserver/model"
)

type BookHandler struct {
	Log        hclog.Logger
	DBInstance *model.DB
}

func NewBookHandler(l hclog.Logger, db *model.DB) *BookHandler {
	return &BookHandler{
		Log:        l,
		DBInstance: db,
	}
}
