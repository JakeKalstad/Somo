package data

import (
	"gopkg.in/mgo.v2/bson"
)

type AddressModel struct {
	entity
	Id    bson.ObjectId 	`json:"id" bson:"_id"`
	Street string	`json:"street" bson:"street"`
	Other string	`json:"other" bson:"other"`
	State string	`json:"state" bson:"state"`
	Country string	`json:"country" bson:"country"`
	Zipcode string	`json:"zipcode" bson:"zipcode"`
}

type AddressQuery struct {
	Id string `bson:"Id"` 
}

func NewAddress (ady AddressModel) AddressModel {
	ady.Id = bson.NewObjectId()
	create("address", &ady)
	return ady
}

func GetAddresses (q AddressQuery) []AddressModel {
	var results []AddressModel
	if (q.Id != "") {
		query("address", bson.M{"entity._id": bson.ObjectIdHex(q.Id)}, &results)
	} else {
		query("address", nil, &results)
	}
	return results
}