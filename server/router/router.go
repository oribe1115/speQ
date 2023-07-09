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
