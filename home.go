package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func Yourpage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	posts, err := getPosts(req, nil)
	if err != nil {
		log.Errorf(ctx, "error getting postss: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	memItem, err := getSession(req)
	var sd SessionData
	if err == nil {
		// logged in
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
	}
	sd.NewPosts = posts
	tpl.ExecuteTemplate(res, "yourpage.html", &sd)
}

func getPosts(req *http.Request, user *User) ([]NewPost, error) {
	ctx := appengine.NewContext(req)

	var posts []NewPost
	q := datastore.NewQuery("Post")

	if user != nil {
		// show tweets of a specific user
		userKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
		q = q.Ancestor(userKey)
	}

	q = q.Order("-Time").Limit(20)
	_, err := q.GetAll(ctx, &posts)
	return posts, err
}
