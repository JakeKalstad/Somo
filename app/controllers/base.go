package controllers

import (
	"github.com/revel/revel"
    "encoding/json"
	"Somo/app/Data"
    "fmt"
)

type Base struct {
	*revel.Controller
}

func (c Base) Decode(q interface{}) {
    decoder := json.NewDecoder(c.Request.Body)
    decoder.Decode(q)
}

func init() {
    revel.InterceptFunc(setuser, revel.BEFORE, &Base{})
}
    
func setuser(c *revel.Controller) revel.Result {
    fmt.Printf("SETTING USER")
    var user *data.UserModel
    if _, ok := c.Session["UID"]; ok {
        uid := c.Session["UID"]
        fmt.Printf(uid)
        user = data.GetUser(uid)
    }
    if (user != nil) {
        c.RenderArgs["user"] = user
    }
    return nil
}

func (c Base) connected() *data.UserModel {
    return c.RenderArgs["user"].(*data.UserModel)
}