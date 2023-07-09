package service

import "speQ/generated/model"

type Services struct {
	queries *model.Queries
}

func NewService(queries *model.Queries) *Services {
	return &Services{queries: queries}
}
