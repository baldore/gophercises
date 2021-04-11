package urlshort

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func SQLHandler(fallback http.Handler) http.HandlerFunc {
	fmt.Println("yepppp")

	connStr := "user=postgres password=asdfasdf dbname=urlshort sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select url from urls")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", rows)

	// _, err := db.QueryContext(ctx, "SELECT url, path FROM urls")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprintf("#v\n", rows)
	// return func(w http.ResponseWriter, r *http.Request) {
	// 	if url, ok := pathsToUrls[r.URL.Path]; ok {
	// 		http.Redirect(w, r, url, http.StatusMovedPermanently)
	// 	} else {
	// 		fallback.ServeHTTP(w, r)
	// 	}
	// }

	return fallback.ServeHTTP
}
