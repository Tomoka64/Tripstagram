package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/search"
)

//

// type Image struct {
// 	width  int
// 	height int
// }
type NewPostId struct {
	CreatedID string
}

func PostProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(r)
	model := &NewPostId{}

	if r.Method == "POST" {
		title := r.FormValue("Title")
		location := r.FormValue("Location")
		// Image := r.FormValue("Image")
		review := r.FormValue("Review")
		rating := r.FormValue("Rating")

		index, err := search.Open("post")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		newpost := &NewPost{
			Title:    title,
			Location: location,
			Review:   review,
			Rating:   rating,
		}
		id, err := index.Put(ctx, "", newpost)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		model.CreatedID = id
	}

	err := tpl.ExecuteTemplate(w, "done", model)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
