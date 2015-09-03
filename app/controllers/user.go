package controllers

import (
	"github.com/revel/revel"
	"Somo/app/Data"
    "fmt"
    "encoding/json"
)

type User struct {
    Base
	*revel.Controller
}

func encode(user interface {}) string {
    b, err := json.Marshal(user)
    if err != nil {
        return ""
    }
    return string(b)
}
func (c User) Save() revel.Result {
	var m data.UserModel
    c.Decode(&m)
    if (m.UID == "") {   
        m.SetPassword()
        user := data.NewUser(m)
        c.Session["UID"] = user.Id.Hex()
        fmt.Printf(c.Session["UID"])
        fmt.Printf("SESSSSSION /\\/\\")
        c.Session["user"] = encode(user)
        return c.RenderJson(user)
    } else {
        m.Update()
        return c.RenderJson(m)
    }
}

func (c User) DeAuth() revel.Result {
    user := c.connected()
    if (user != nil) {
        user.Deauthorize()
    } 
    c.Session["UID"] = "-1"
    success := true
    return c.RenderJson(success)
}

func (c User) Auth() revel.Result {
    var m data.UserModel
    c.Decode(&m)
    result := m.Authorize()
    if (result.Authorized == true) {
        c.Session["UID"] = result.User.Id.Hex()
    } else {
        c.Session["UID"] = "-1"
    }
    return c.RenderJson(result)
}

func (c User) Account() revel.Result {
    var user *data.UserModel
    c.Decode(&user)
    return c.Render(user)
}

func (c User) Index() revel.Result {
    var q data.UserQuery
    c.Decode(&q)
    return c.RenderJson(data.GetUsers(q))
}

type Get struct {
    Id string `json:"id"`
}

func (c User) Get() revel.Result {
    var m Get
    var user *data.UserModel
    c.Decode(&m)
    user = data.GetUser(m.Id)
    if (user == nil) {
        return c.RenderJson(map[string]interface{}{
            "id":  "",
        });
    }

    return c.RenderJson(user)
}

func (c User) CreateArtist() revel.Result {
    return c.Render();
}

func (c User) Signup() revel.Result {
    return c.Render()
}

// HTML TEMPLATE

func (c User) TSignUp () revel.Result {
    return c.RenderTemplate("Templates/sign_up.html")
}