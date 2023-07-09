package service

import (
	"database/sql"
	"speQ/generated/model"
)

type Services struct {
	db      *sql.DB
	queries *model.Queries
}

func NewService(db *sql.DB, queries *model.Queries) *Services {
	return &Services{db: db, queries: queries}
}
