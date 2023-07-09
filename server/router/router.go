package router

import (
	"speQ/generated/api"
	"speQ/generated/model"
)

type Router struct {
	queries *model.Queries
}

var (
	_ api.ServerInterface = &Router{}
)

func NewRouter(queries *model.Queries) *Router {
	return &Router{queries: queries}
}
