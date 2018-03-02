package main

import (
	"time"
)

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type SessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
	NewPosts  []NewPost
}

type NewPost struct {
	Title    string
	Location string
	// Image    multipart.File
	Review   string
	Rating   string
	Time     time.Time
	UserName string
}

//
// type Image interface {
// 	image.Image
// 	Set(x, y int, c color.Color)
// }
