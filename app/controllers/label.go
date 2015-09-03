package controllers

import (
	"github.com/revel/revel"
	"Somo/app/Data"
)

type Label struct {
    Base
	*revel.Controller
}

func (c Label) Save() revel.Result {
	var m data.LabelModel
    c.Decode(&m)
    return c.RenderJson(data.NewLabel(m))
}

func (c Label) Index() revel.Result {
    var q data.LabelQuery
    c.Decode(&q)
    return c.RenderJson(data.GetLabels(q))
}
