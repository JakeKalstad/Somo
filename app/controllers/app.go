package controllers

import "github.com/revel/revel"

type App struct {
	Base
	*revel.Controller
}

func (c App) Index() revel.Result {
    UID := c.Session["UID"]
	return c.Render(UID)
}
