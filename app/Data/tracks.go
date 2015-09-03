package data

import (
	"gopkg.in/mgo.v2/bson"
)
type TrackModel struct {
	entity
	Id    bson.ObjectId 	`json:"id" bson:"_id"`
 	Name   string        	`json:"name" bson:"name"`
 	Description   string     `json:"description" bson:"description"`
  	ImgUrl   string        	`json:"img_url" bson:"img_url`
  	Handle   string        	`json:"handle" bson:"handle`  
  	Filename string 		`json:"filename" bson:"filename"`
  	Tags []string			`json:"tags" bson:"tags"`
  	UID string 				`json:"uid" bson:"uid"`
}

type TrackQuery struct {
	Id string `json:"id" bson:"Id"` 
}

func NewTrack (trk TrackModel) TrackModel {
	trk.Id = bson.NewObjectId()
	create("track", &trk)
	return trk
}

func DeleteTrack(q TrackQuery) {
	delete(bson.M{"artist.tracks._id": bson.ObjectIdHex(q.Id)})
}

func GetTracksByArtist (q TrackQuery) []TrackModel {
	var results []TrackModel
	query("track", bson.M{"_id": bson.ObjectIdHex(q.Id)}, &results)
	return results
}

func GetTracks (q TrackQuery) []TrackModel {
	var results []TrackModel
	if (q.Id != "") {
		query("track", bson.M{"entity._id": bson.ObjectIdHex(q.Id)}, &results)
	} else {
		query("track", nil, &results)
	}
	return results
}