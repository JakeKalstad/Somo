package controllers

import (
	"github.com/revel/revel"
	"Somo/app/Data"
)

type Address struct {
    Base
	*revel.Controller
}

func (c Address) Save() revel.Result {
	var m data.AddressModel
    c.Decode(&m)
    return c.RenderJson(data.NewAddress(m))
}

func (c Address) Index() revel.Result {
    var q data.AddressQuery
    c.Decode(&q)
    return c.RenderJson(data.GetAddresses(q))
}
