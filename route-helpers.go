package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

// Database password
var pass = os.Getenv("ELIS_MYSQL")

// Connecting to database. Needed for every query
func db() (*sql.DB, error) {
	return sql.Open("mysql", "elis:"+pass+"@/elis_test?charset=utf8")
}

// Fills headers and sends data (json or not)
func send(w http.ResponseWriter, jsn []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Access-Control-Allow-Origin", "*") // FIXME for my server requests only
	w.WriteHeader(http.StatusOK)
	w.Write(jsn)
}

// For methods that don't need to send anything back (e.g. post methods)
func accepted(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*") // FIXME for my server requests only
	w.WriteHeader(http.StatusAccepted)
}

// For sending errors not in JSON format FIXME should be in JSON
func sendErr(w http.ResponseWriter, err error, msg string, status int) {
	http.Error(w, err.Error(), status)
	w.Write([]byte(msg))
}
