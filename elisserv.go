package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	port = ":7070"
)

func main() {
	router := httprouter.New()
	router.GET("/api/home", posts)
	router.GET("/api/posts", posts)
	//	router.POST("/api/posts", addPost) // needs realization!
	router.GET("/api/blog/:permalink", blogPost)
	router.GET("/api/portfolio", portfolio)
	//	router.GET("/pictures/:picture_name", getPicture)
	//router.POST("/api/login", login)
	//router.POST("/api/contact-me", contactMe)
	// Admin area
	//router.POST("/api/posts", posts)
	//router.GET("/api/images", images)
	//router.GET("/api/messages", messages)
	//router.GET("/api/tags", tags)

	http.ListenAndServe(port, router)
}
