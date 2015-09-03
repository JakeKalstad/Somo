package controllers

import (
	"github.com/revel/revel"
	"Somo/app/Data"
)

type Track struct {
    Base
	*revel.Controller
}
func (c Track) Delete() revel.Result {
    var q data.TrackQuery
    c.Decode(&q)
    data.DeleteTrack(q)
    return c.RenderJson(q)
}
func (c Track) Save() revel.Result {
	var m data.TrackModel
    c.Decode(&m)
    return c.RenderJson(data.NewTrack(m))
}

func (c Track) Index() revel.Result {
	var q data.TrackQuery
    c.Decode(&q)
    return c.RenderJson(data.GetTracks(q))
}

func (c Track) Artist() revel.Result {
    var q data.TrackQuery
    c.Decode(&q)
    return c.RenderJson(data.GetTracksByArtist(q))
}