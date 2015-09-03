package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"fmt"
	"time"
)

type entity struct {
	Created time.Time `bson:"created"`
	Modified time.Time `bson:"modified",omitempty`
	Deleted time.Time `bson:"deleted",omitempty`
}

func GetId() bson.ObjectId {
	return bson.NewObjectId()
}
type Db struct {
	Con *mgo.Session 
}

func InitDb() *Db {
	db := new (Db) 
  	sess, err := mgo.Dial("127.0.0.1:27017")
  	if err != nil {
	    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    	os.Exit(1)
  	}
  	sess.SetSafe(&mgo.Safe{})
  	db.Con = sess
  	return db
}

func update(col string, id bson.ObjectId, update_cmd bson.M) {
	db := InitDb()
	db.Con.DB("SOMO").C(col).Update(bson.M{"_id": id}, bson.M{"$set": update_cmd})
  	defer db.Con.Close()
}

func updateM(col string, selector bson.M, update_cmd bson.M) {
	db := InitDb()
	db.Con.DB("SOMO").C(col).Update(selector, bson.M{"$set": update_cmd})
  	defer db.Con.Close()
}

func upsert(col string, id bson.ObjectId, model interface{}) {
    	db := InitDb()
		db.Con.DB("SOMO").C(col).Upsert(bson.M{"_id": id}, model)
	  	defer db.Con.Close()
}

func create(col string, model interface{}) {
	var db = InitDb()
	err := db.Con.DB("SOMO").C(col).Insert(model)
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
	}
  	defer db.Con.Close()
}

func query(col string, query bson.M, results interface{}) {
	db := InitDb()
	if (query != nil) {
		fromDate := time.Date(2014, time.November, 4, 0, 0, 0, 0, time.UTC)
		query["entity.deleted"] = bson.M{ "$lt" : fromDate }
		err := db.Con.DB("SOMO").C(col).Find(query).All(results)
		if err != nil {
			fmt.Printf("Error searching documents: %v\n", err)
		}
	} else {
		err := db.Con.DB("SOMO").C(col).Find(bson.M{"entity.deleted" : bson.M{ "$exists" : false }}).All(results)
		if err != nil {
			fmt.Printf("Error searching documents: %v\n", err)
		}
	}
}

func delete(query bson.M) {
	var m TrackModel
	db := InitDb()
	change := mgo.Change{
        Update: bson.M{".entity.deleted": time.Now()},
        ReturnNew: true,
	}
	_, err := db.Con.DB("SOMO").C("user").Find(query).Apply(change, &m) 
	fmt.Printf("Can't del document: %v\n", err)
}