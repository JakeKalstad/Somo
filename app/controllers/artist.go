package controllers

import (
	"github.com/revel/revel"
	"Somo/app/Data"
)

type Artist struct {
    Base
	*revel.Controller
}


func (c Artist) Upload() revel.Result {
    return c.Render()
} 

func (c Artist) Save() revel.Result {
	var m data.ArtistModel
    c.Decode(&m)
    return c.RenderJson(data.NewArtist(m))
}

func (c Artist) Index() revel.Result {
    var q data.ArtistQuery
    c.Decode(&q)
    return c.RenderJson(data.GetArtists(q))
}
