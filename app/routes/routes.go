// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tBase struct {}
var Base tBase



type tTrack struct {}
var Track tTrack


func (_ tTrack) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Track.Delete", args).Url
}

func (_ tTrack) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Track.Save", args).Url
}

func (_ tTrack) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Track.Index", args).Url
}

func (_ tTrack) Artist(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Track.Artist", args).Url
}


type tArtist struct {}
var Artist tArtist


func (_ tArtist) Upload(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Artist.Upload", args).Url
}

func (_ tArtist) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Artist.Save", args).Url
}

func (_ tArtist) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Artist.Index", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Save", args).Url
}

func (_ tUser) DeAuth(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.DeAuth", args).Url
}

func (_ tUser) Auth(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Auth", args).Url
}

func (_ tUser) Account(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Account", args).Url
}

func (_ tUser) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Index", args).Url
}

func (_ tUser) Get(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Get", args).Url
}

func (_ tUser) CreateArtist(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.CreateArtist", args).Url
}

func (_ tUser) Signup(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Signup", args).Url
}

func (_ tUser) TSignUp(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.TSignUp", args).Url
}


type tLabel struct {}
var Label tLabel


func (_ tLabel) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Label.Save", args).Url
}

func (_ tLabel) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Label.Index", args).Url
}


type tAddress struct {}
var Address tAddress


func (_ tAddress) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Address.Save", args).Url
}

func (_ tAddress) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Address.Index", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tAudio struct {}
var Audio tAudio


func (_ tAudio) Upload(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Audio.Upload", args).Url
}

func (_ tAudio) Stream(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Audio.Stream", args).Url
}

func (_ tAudio) HandleUpload(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Audio.HandleUpload", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


