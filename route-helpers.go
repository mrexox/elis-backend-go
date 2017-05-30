package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

var pass = os.Getenv("ELIS_MYSQL")

// DRY
func db() (*sql.DB, error) {
	return sql.Open("mysql", "elis:"+pass+"@/elis_test?charset=utf8")
}

func send(w http.ResponseWriter, jsn []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsn)
}

// For sendErr

func sendErr(w http.ResponseWriter, err error, msg string, status int) {
	http.Error(w, err.Error(), status)
	w.Write([]byte(msg))
}
