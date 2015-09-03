package data

import (
	"gopkg.in/mgo.v2/bson"
)

type ArtistModel struct {
	entity
	Id    bson.ObjectId 	`json:"id" bson:"_id"`
 	Moniker   	string      `json:"moniker" bson:"moniker"`
  	ImgUrl   	string      `json:"img_url" bson:"img_url`
  	Blurb 		string		`json:"blurb" bson:"blurb` 
  	Tracks 		[]TrackModel		`json:"tracks" bson:"tracks`
}

type ArtistQuery struct {
	Id string `bson:"Id"` 
}

func NewArtist (art ArtistModel) ArtistModel {
	art.Id = bson.NewObjectId()
	return art
}

func GetArtists (q ArtistQuery) []ArtistModel {
	var results []ArtistModel
	if (q.Id != "") {
		query("artist", bson.M{"entity._id": bson.ObjectIdHex(q.Id)}, &results)
	} else {
		query("artist", nil, &results)
	}
	return results
}