package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	port = ":7070"
)

/*

get /api/posts - all posts
get /api/posts/:permalink - post with :permalink
get /api/home - 4 recent posts
get /api/portfolio - portfolio images
get /api/tags - get list of tags ADMIN
get /api/images - get all images ADMIN
get /api/messages - get all messages ADMIN

post /api/posts - create new post ADMIN
post /api/images - add new image ADMIN
post /api/mesages - send new message


*/

func main() {
	router := httprouter.New()
	router.GET("/api/home", posts)
	router.GET("/api/posts", posts)
	router.POST("/api/posts", addPost) // needs realization!
	router.POST("/api/images", addImage)
	router.GET("/api/posts/:permalink", blogPost)
	router.GET("/api/portfolio", portfolio)

	//router.POST("/api/login", login)
	// Admin area
	//router.POST("/api/posts", posts)
	//router.GET("/api/images", images)
	//router.GET("/api/messages", messages)
	//router.GET("/api/tags", tags)

	router.NotFound = http.FileServer(http.Dir("/home/ian/data"))
	http.ListenAndServe(port, router)
}
