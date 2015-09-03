package data

type Error struct {
	Msg		string		`json:"msg" bson:"msg"`
	Type 	string		`json:"type" bson:"type"`
}