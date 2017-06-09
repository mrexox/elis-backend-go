package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go/jsonapi"
	"log"
	"net/http"
)

/* GET handler for /api/posts request.
 * Returns all posts that a database has
 */

func posts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := db()
	if err != nil {
		log.Println("Database was not properly opened")
		sendErr(w, err, "Database was not properly opened", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query(`
SELECT p.id, p.name, p.content, p.permalink, p.visible, p.created_at, IFNULL(likes.likes, 0), p.cover
FROM post p LEFT OUTER JOIN (SELECT post_id, COUNT(ip) AS  likes
                      FROM liker
                      GROUP BY post_id) likes ON p.id = likes.post_id
ORDER BY created_at;
	`)
	if err != nil {
		log.Println("Error in statement")
		sendErr(w, err, "Error in query blog", http.StatusInternalServerError)
		return
	}

	var posts []Post
	for rows.Next() {
		var post Post
		// getting a post
		var coverID sql.NullInt64
		err = rows.Scan(&post.ID, &post.Name, &post.Content, &post.Permalink,
			&post.Visible, &post.CreatedAt, &post.Likes, &coverID)
		if err != nil {
			log.Println("Error while scanning at home handler")
			sendErr(w, err, "Error while scanning blog", http.StatusInternalServerError)
			return
		}
		// Adding a cover if neded
		if coverID.Valid {
			row := db.QueryRow(`
SELECT id, url
FROM image
WHERE id = ?
`, coverID)
			var image Image
			err := row.Scan(&image.ID, &image.Url)
			if err != nil {
				sendErr(w, err, "Error while trying to get an image for a post", http.StatusInternalServerError)
				return
			}
			post.Cover = image
		}

		// Adding tags
		tagRows, err := db.Query(`
SELECT t.id, t.name 
FROM post_tag pt INNER JOIN tag t ON pt.tag_id = t.id
WHERE pt.post_id = ?;
		`, post.ID)
		if err != nil {
			log.Println("Error while scanning at home handler")
			sendErr(w, err, "Error while scanning blog", http.StatusInternalServerError)
			return
		}
		var tagsIDs []string
		var tags []Tag
		for tagRows.Next() {
			var tag Tag
			tagRows.Scan(&tag.ID, &tag.Name)
			tags = append(tags, tag)
			tagsIDs = append(tagsIDs, tag.ID)
		}
		post.TagsIDs = tagsIDs
		post.Tags = tags
		// Append to []Post
		posts = append(posts, post)
	}
	db.Close()
	jsn, err := jsonapi.Marshal(posts)
	send(w, jsn)
}
