package main

import (
	//	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go/jsonapi"
	"io/ioutil"
	"log"
	"net/http"
)

func addPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErr(w, err, "Error while parsing body", 403)
		return
	}
	var post Post
	err = jsonapi.Unmarshal(body, &post)
	if err != nil {
		sendErr(w, err, "Error while unmarshaling post", 403)
		return
	}
	log.Printf("%v", post)
	accepted(w)
}
