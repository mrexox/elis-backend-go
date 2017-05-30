package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go/jsonapi"
	"log"
	"net/http"
)

/* GET for /api/portfolio
 * Returns the pictures (urls of them) that are thought to be portfolio
 * pictures
 */

func portfolio(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := db()
	if err != nil {
		log.Println("Error while opening database")
		sendErr(w, err, "Error while opening database", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query(`
SELECT im.id, im.url
FROM portfolio_image pi INNER JOIN image im ON pi.id = im.id;
	`)
	db.Close()
	if err != nil {
		log.Println("Error while query in portfolio")
		sendErr(w, err, "Error while query in portfolio", http.StatusInternalServerError)
		return
	}
	var portfolioImages []Image
	for rows.Next() {
		var portfolioImage Image
		err := rows.Scan(&portfolioImage.ID, &portfolioImage.Url)
		if err != nil {
			log.Println("Error while reading rows into go struct")
			sendErr(w, err, "Error while reading rows into go struct", http.StatusInternalServerError)
			return
		}
		portfolioImages = append(portfolioImages, portfolioImage)
	}
	jsn, err := jsonapi.Marshal(portfolioImages)

	if err != nil {
		log.Println("Error while marshalling portfolio")
		sendErr(w, err, "Error while marshalling portfolio", http.StatusInternalServerError)
		return
	}
	send(w, jsn)
}
