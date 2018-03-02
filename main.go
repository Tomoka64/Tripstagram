package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//
// var tpl = template.Must(
// 	template.New("").
// 		Funcs(template.FuncMap{}).
// 		ParseGlob("templates/*.gohtml"),
// )

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/form/login", Login)
	r.GET("/form/signup", Signup)
	r.POST("/api/checkusername", checkUserName)
	r.POST("/api/createuser", createUser)
	r.POST("/api/login", loginProcess)
	r.GET("/api/logout", logout)
	r.GET("/yourpage", Yourpage)
	r.GET("/post", PostPage)
	r.POST("/post-process", PostProcess)
	r.GET("/post/done", Done)
	// r.GET("/delete", Delete)
	// r.GET("/loginWithFB", handleFacebookLogin)
	// r.GET("/login/facebook", facebookk)
	// r.POST("/oauth2callback", handleFacebookCallback)
	// r.GET("/logout", logout)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	tpl = template.Must(template.ParseGlob("template/*"))
}

func PostPage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	memItem, err := getSession(req)
	ctx := appengine.NewContext(req)
	if err != nil {
		// not logged in
		log.Infof(ctx, "Attempt to post from logged out user")
		http.Redirect(res, req, "form/login", http.StatusSeeOther)
		return
	} else {
		// logged in
		var sd SessionData
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
		tpl.ExecuteTemplate(res, "new.html", sd)
	}
}

func Done(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "done.html")
}
func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "home.html")
}
func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "login.html")
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "signup.html")
}
