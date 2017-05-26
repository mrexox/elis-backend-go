package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go/jsonapi"
	"log"
	"net/http"
)

const (
	pass = "El15Typ3"
)

// DRY
func db() (*sql.DB, error) {
	return sql.Open("mysql", "elis:"+pass+"@/elis_test?charset=utf8")
}

func send(w http.ResponseWriter, jsn []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsn)
}

// /api/home Home
func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := db()
	if err != nil {
		log.Println("Database was not properly opened")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.Query(`
SELECT p.id, p.name, p.content, p.permalink, p.created_at, likes.likes
FROM post p LEFT OUTER JOIN (SELECT post_id, COUNT(ip) AS  likes
                      FROM liker
                      GROUP BY post_id) likes ON p.id = likes.post_id
ORDER BY created_at
LIMIT 4;
	`)
	db.Close()
	if err != nil {
		log.Println("Error in statement")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Name, &post.Content, &post.Permalink, &post.CreatedAt, &post.Likes)
		posts = append(posts, post)
	}
	jsn, err := jsonapi.Marshal(posts)
	send(w, jsn)
}

// /api/blog Blog
func blog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	home(w, r, p)
}

// /api/blog/:permalink
func blogPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	permalink := p.ByName("permalink")
	db, err := db()
	// Get post of given permalink
	row := db.QueryRow(`
SELECT p.id, p.name, p.content, p.permalink, p.created_at, likes.likes
FROM post p LEFT OUTER JOIN (SELECT post_id, COUNT(ip) AS  likes
                      FROM liker
                      GROUP BY post_id) likes ON p.id = likes.post_id
WHERE permalink = ?;
	`, permalink)
	var post Post
	err = row.Scan(&post.ID, &post.Name, &post.Content, &post.Permalink, &post.CreatedAt, &post.Likes)

	// Get tags of this post
	rows, err := db.Query(`
			SELECT t.id, t.name
			FROM post_tag pt INNER JOIN tag t ON pt.tag_id = t.id
			WHERE pt.post_id = ?;
			`, post.ID)
	db.Close()

	if err != nil {
		log.Printf("Error in query while selecting tags of post id:%s", post.ID)
		log.Println(rows.Err().Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Managing relations
	var tags []Tag
	var tagsIDs []string
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			log.Println("Error while iterating through tags")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tags = append(tags, tag)
		tagsIDs = append(tagsIDs, tag.ID)
	}
	post.Tags = tags
	post.TagsIDs = tagsIDs
	jsn, err := jsonapi.Marshal(post)
	if err != nil {
		log.Println("Error while marshalling")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	send(w, jsn)
}
