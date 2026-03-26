package artifacts

import (
	"database/sql"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

type Url struct {
	Title string
	Url   string
	Time  int64
}

func History(src string) (urls []Url, err error) {
	var query string

	switch strings.ToLower(filepath.Base(src)) {
	case "history":
		query = `
			SELECT u.url, u.title, (t.visit_time-11644473600000000)
			FROM urls AS u, visits AS t
			WHERE u.id = t.url
		;`
	case "places.sqlite":
		query = `
			SELECT u.url, COALESCE(u.title, ''), t.visit_date
			FROM moz_places AS u, moz_historyvisits AS t
			WHERE u.id = t.place_id
		;`
	}

	db, err := sql.Open("sqlite", src)

	if err != nil {
		return
	}

	defer func() { _ = db.Close() }()

	rows, err := db.Query(query)

	if err != nil {
		return
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {
		u := Url{}

		err = rows.Scan(&u.Url, &u.Title, &u.Time)

		if err != nil {
			return
		}

		urls = append(urls, u)
	}

	err = rows.Err()

	return
}
