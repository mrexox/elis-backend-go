package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go/jsonapi"
	"log"
	"net/http"
)

/* GET for /api/blog/:permalink
* Returns the post of given :permalink
 */

func blogPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	permalink := p.ByName("permalink")
	db, err := db()
	row := db.QueryRow(`
SELECT p.id, p.name, p.content, p.permalink, p.created_at, IFNULL(likes.likes, 0)
FROM post p LEFT OUTER JOIN (SELECT post_id, COUNT(ip) AS  likes
                      FROM liker
                      GROUP BY post_id) likes ON p.id = likes.post_id
WHERE permalink = ?;
	`, permalink)
	var post Post
	err = row.Scan(&post.ID, &post.Name, &post.Content, &post.Permalink, &post.CreatedAt, &post.Likes)
	if err == sql.ErrNoRows {
		sendErr(w, err, "404", http.StatusNotFound)
		return
	} else if err != nil {
		sendErr(w, err, "Error while scanning row in blogPost", http.StatusInternalServerError)
		return
	}

	// Get tags of this post
	rows, err := db.Query(`
			SELECT t.id, t.name
			FROM post_tag pt INNER JOIN tag t ON pt.tag_id = t.id
			WHERE pt.post_id = ?;
			`, post.ID)
	db.Close()

	if err != nil {
		log.Printf("Error in query while selecting tags of post id:%s", post.ID)
		sendErr(w, err, "Error in query", http.StatusInternalServerError)
		return
	}

	// Managing relations
	var tags []Tag
	var tagsIDs []string
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			log.Println("Error while iterating through tags")
			sendErr(w, err, "Error while iterating through tags", http.StatusInternalServerError)
		}
		tags = append(tags, tag)
		tagsIDs = append(tagsIDs, tag.ID)
	}
	post.Tags = tags
	post.TagsIDs = tagsIDs
	jsn, err := jsonapi.Marshal(post)
	if err != nil {
		log.Println("Error while marshalling")
		sendErr(w, err, "Error while marshalling", http.StatusInternalServerError)
		return
	}
	send(w, jsn)
}
