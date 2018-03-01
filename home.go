package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/search"
)

type HomeModel struct {
	Posts []NewPost
	Sd    SessionData
}

func Yourpage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	memItem, err := getSession(req)
	model := HomeModel{}
	if err != nil {
		// not logged in
		tpl.ExecuteTemplate(res, "yourpage.html", SessionData{})
	} else {
		// logged in
		var sd SessionData
		json.Unmarshal(memItem.Value, &sd)
		sd.LoggedIn = true
		model.Sd = sd

		ctx := appengine.NewContext(req)
		index, err := search.Open("post")
		if err != nil {
			panic(err)
		}
		iterator := index.Search(ctx, ".", nil)
		for {
			var post NewPost
			_, err := iterator.Next(&post)
			if err == search.Done {
				break
			} else if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			model.Posts = append(model.Posts, post)
		}

		err = tpl.ExecuteTemplate(res, "yourpage.html", model.Sd)

		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	}
}
