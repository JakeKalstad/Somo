package data

import (
	"gopkg.in/mgo.v2/bson"
)

type LabelModel struct {
	entity
	Id    bson.ObjectId 	`json:"id" bson:"_id"`
	Name   		string       	`json:"name" bson:"name"`
  	ImgUrl   	string        	`json:"img_url" bson:"img_url`
  	Handle   	string     		`json:"handle" bson:"handle`
  	Blurb 		string			`json:"blurb" bson:"blurb`
}

type LabelQuery struct {
	Id string `bson:"Id"` 
}

func NewLabel (ady LabelModel) LabelModel {
	ady.Id = bson.NewObjectId()
	create("label", &ady)
	return ady
}

func GetLabels (q LabelQuery) []LabelModel {
	var results []LabelModel
	if (q.Id != "") {
		query("label", bson.M{"entity._id": bson.ObjectIdHex(q.Id)}, &results)
	} else {
		query("label", nil, &results)
	}
	return results
}