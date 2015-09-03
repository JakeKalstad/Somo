package data

import (
	"gopkg.in/mgo.v2/bson"
    "golang.org/x/crypto/bcrypt"
    "Somo/app/helpers"
    "fmt"
    )

type UserModel struct {
	entity
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Label LabelModel 	`json:"label" bson:"label,omitempty"`
	Artist ArtistModel 	`json:"artist" bson:"artist,omitempty"`
	Email string 		`json:"email" bson:"email"`
	Password string 	`json:"password" bson:"password"`
	Login string 		`json:"login" bson:"login"`
	Token string 		`json:"token" bson:"token"`
	UID string			`json:"uid"`
}

type UserResult struct {
	Authorized bool		`json:"authorized" bson:"authorized"`		
	Errors []Error		`json:"errors" bson:"errors"`
	User *UserModel		`json:"user" bson:"user"`
}

type UserQuery struct {
	Id string `bson:"Id"` 
}
 
func (m *UserModel) SetPassword () {
    password := []byte(m.Password)
    hashedPassword, _ := bcrypt.GenerateFromPassword(password, 10)
    m.Password = string(hashedPassword)
}

func (m *UserModel) Authorize () UserResult {
	var results []UserModel
	fmt.Printf(m.Password)
	fmt.Printf("/\\")
	query("user", bson.M{"login": m.Login}, &results)
	if (len(results)) == 0 {
		error := Error { Msg: "User not found", Type: "login" }
		user := UserResult { Authorized: false, Errors: []Error { error } }
		return user
	}
	userRecord := results[0]
	compResult := bcrypt.CompareHashAndPassword([]byte(userRecord.Password), []byte(m.Password)) == nil
	if (!compResult) {
		error := Error { Msg: "Password does not match", Type: "password" }
		user := UserResult { Authorized: false, Errors: []Error { error } }
		return user
	}
	token := helpers.GenerateHandler()
	update("user", userRecord.Id, bson.M{"token": token})
	userRecord.Token = token
	userRes := UserResult { Authorized: true, User: &userRecord }
	return userRes
}

func (user *UserModel) Update() {
	if (user.Artist.Id == "" && user.Artist.Moniker != "") {
        user.Artist.Id = GetId()
    	upsert("user", user.Id, user)
    } 
	if (user.Label.Id == "" && user.Label.Name != "") {
		user.Label.Id = GetId()
    }
	if (len(user.Artist.Tracks) != 0) {
		for i := 0; i < len(user.Artist.Tracks); i++ {
		    user.Artist.Tracks[i].Id = GetId()
		    user.Artist.Tracks[i].UID = user.Artist.Tracks[i].Id.Hex()
		}
	}
    upsert("user", user.Id, user)
}

func (user *UserModel) Deauthorize() {
	update("user", user.Id, bson.M{"token": ""})
}

func NewUser (user UserModel) UserModel {
	user.Id = bson.NewObjectId()
	user.UID = user.Id.Hex()
	create("user", &user)
	return user
}

func GetUsers (q UserQuery) []UserModel {
	var results []UserModel
	if (q.Id != "") {
		query("user", bson.M{"_id": bson.ObjectIdHex(q.Id), "entity.deleted" : bson.M{ "$exists" : false }}, &results)
	} else {
		query("user", nil, &results)
	}
	return results
}

func GetUser (id string) *UserModel {
	var results []UserModel
	if (id == "" || id =="-1") {
		return &UserModel { UID: "-1" }
	}
	query("user", bson.M{"_id": bson.ObjectIdHex(id)}, &results)
	if (len(results) > 0) {
		results[0].UID = results[0].Id.Hex()
		return &results[0]
	} else {
		return nil
	}
}