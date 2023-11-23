package models

import (
	"github.com/upper/db/v4"
)

type Models struct {
	Users UsersModel
}

func New(db db.Session) Models {
	models := Models {
		Users: UsersModel{db: db},
	}

	return models
}