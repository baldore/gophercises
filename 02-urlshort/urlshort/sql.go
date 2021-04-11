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

	rows, err := db.Query("select url_path, url from urls")
	if err != nil {
		panic(err)
	}

	got := map[string]string{}

	for rows.Next() {
		var urlPath, path string
		err = rows.Scan(&urlPath, &path)

		if err != nil {
			panic(err)
		}

		got[urlPath] = path
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if url, ok := got[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}

	return handler
}
