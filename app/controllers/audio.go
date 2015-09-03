package controllers

import (
	"github.com/revel/revel"
	"Somo/app/helpers"
	"Somo/app/Data"
	"os"
	"fmt"
	)

type Audio struct {
	Base
	*revel.Controller
}
func (c *Audio) Upload() revel.Result {
	return c.Render()
}

type FileInfo struct {
	ContentType string
	Filename    string
	RealFormat  string `json:",omitempty"`
	Resolution  string `json:",omitempty"`
	Size        int
	Status      string `json:",omitempty"`
	Handler 	string `json:"handler"`
}

type TrackUrl struct {
	Name string
}
func (c *Audio) Stream() revel.Result {
	var user data.UserModel
    c.Decode(&user)
	return c.Render(user)
}

func (c *Audio) HandleUpload() revel.Result {
	var files [][]byte

	c.Params.Bind(&files, "file")
	filesInfo := make([]FileInfo, len(files))
	for i, _ := range files {
		handler := helpers.GenerateHandler()
		filesInfo[i] = FileInfo{
			ContentType: c.Params.Files["file[]"][i].Header.Get("Content-Type"),
			Filename:    c.Params.Files["file[]"][i].Filename,
			Size:        len(files[i]),
			Handler: 	 handler,
		}
		var name = revel.BasePath + "/public/audio/" + handler + filesInfo[i].Filename[:len(filesInfo[i].Filename)-4] + filesInfo[i].Filename[len(filesInfo[i].Filename)-4:len(filesInfo[i].Filename)]
	    fmt.Printf(name);
		fo, err := os.Create(name)
	    if err != nil {
	        panic(err)
	    }
	    fo.Write(files[i])

	    defer func() {
	        if err := fo.Close(); err != nil {
				panic(err)
	        }
	    }()
	}
	return c.RenderJson(map[string]interface{}{
		"Count":  len(files),
		"Files":  filesInfo,
		"Status": "Successfully uploaded",
	})
}
