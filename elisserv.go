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
	router.GET("/api/home", home)
	//router.GET("/api/login", login)
	router.GET("/api/blog", blog)
	router.GET("/api/blog/:permalink", blogPost)
	//router.GET("/api/portfolio", portfolio)
	//router.GET("/api/contact-me", contactMe)
	//router.GET("/api/search/:tag", search)
	// Admin area
	//router.GET("/api/posts", posts)
	//router.GET("/api/images", images)
	//router.GET("/api/messages", messages)
	//router.GET("/tags", tags)
	http.ListenAndServe(port, router)
}
