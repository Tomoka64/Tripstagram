package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

//

// type Image struct {
// 	width  int
// 	height int
// }

func PostProcess(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)
	if err != nil {
		log.Infof(ctx, "Attempt to post tweet from logged out user")
		http.Error(res, "You must be logged in", http.StatusForbidden)
		return
	} else {
		var user User
		json.Unmarshal(memItem.Value, &user)
		log.Infof(ctx, user.UserName)

		title := req.FormValue("Title")
		location := req.FormValue("Location")
		// Image := r.FormValue("Image")
		review := req.FormValue("Review")
		rating := req.FormValue("Rating")

		post := NewPost{
			Title:    title,
			Location: location,
			Review:   review,
			Rating:   rating,
			Time:     time.Now(),
			UserName: user.UserName,
		}
		err = putNewPost(req, &user, &post)
		if err != nil {
			log.Errorf(ctx, "error adding todo: %v", err)
			http.Error(res, err.Error(), 500)
			return
		}

	}
	time.Sleep(time.Millisecond * 500)
	http.Redirect(res, req, "/", 302)
}

func putNewPost(req *http.Request, user *User, post *NewPost) error {
	ctx := appengine.NewContext(req)
	userKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	key := datastore.NewIncompleteKey(ctx, "Post", userKey)
	_, err := datastore.Put(ctx, key, post)
	return err
}
