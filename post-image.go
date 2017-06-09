package main

import (
	//	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	//"github.com/manyminds/api2go/jsonapi"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	savePath = "/home/ian/data/pictures/"
)

/* addImage responce for accepting file uploading (images)
 * FIXME no other files should be allowed
 * FIXME accepting json files should be allowed
 */
func addImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseMultipartForm(32 << 24) // Size in bytes to be accepted 511 mb
	file, handler, err := r.FormFile("uploadimage")
	if err != nil {
		sendErr(w, err, "Error while uploading an image", 403)
		return
	}
	defer file.Close()
	t := time.Now()
	filename := handler.Filename + t.Format("20060102150405")
	f, err := os.OpenFile(savePath+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		sendErr(w, err, "Error while saving an image", 403)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	db, err := db()
	if err != nil {
		sendErr(w, err, "Error while connecting to database", 403)
		return
	}
	stmt, err := db.Prepare("INSERT INTO image (url) VALUES (?)")
	if err != nil {
		sendErr(w, err, "Error while making an sql statement", 403)
		return
	}
	_, err = stmt.Exec(filename)
	if err != nil {
		sendErr(w, err, "Error while storing an image", 403)
		return
	}
	accepted(w)
}
