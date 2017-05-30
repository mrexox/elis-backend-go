package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go/jsonapi"
	"log"
	"net/http"
)

/* GET handler for /api/posts request.
 * Returns all posts that a database has
 * FIXME: add tags to posts
 */

func posts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := db()
	if err != nil {
		log.Println("Database was not properly opened")
		sendErr(w, err, "Database was not properly opened", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query(`
SELECT p.id, p.name, p.content, p.permalink, p.created_at, IFNULL(likes.likes, 0)
FROM post p LEFT OUTER JOIN (SELECT post_id, COUNT(ip) AS  likes
                      FROM liker
                      GROUP BY post_id) likes ON p.id = likes.post_id
ORDER BY created_at;
	`)
	db.Close()
	if err != nil {
		log.Println("Error in statement")
		sendErr(w, err, "Error in query blog", http.StatusInternalServerError)
		return
	}
	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Name, &post.Content, &post.Permalink, &post.CreatedAt, &post.Likes)
		if err != nil {
			log.Println("Error while scanning at home handler")
			sendErr(w, err, "Error while scanning blog", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	jsn, err := jsonapi.Marshal(posts)
	send(w, jsn)
}
