package router

import (
	"speQ/generated/api"
	"speQ/generated/model"
	"speQ/service"
)

type Router struct {
	queries  *model.Queries
	services *service.Services
}

var (
	_ api.ServerInterface = &Router{}
)

func NewRouter(queries *model.Queries, services *service.Services) *Router {
	return &Router{queries: queries, services: services}
}
