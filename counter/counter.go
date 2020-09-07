package counter

import (
	"database/sql"
	"time"
)

// Simple counter service based on sqlite3.

// Sqlite3 schema
// $ sqlite> create table views (path varchar(50), timestamp integer);
// $ sqlite> insert into views values ("/api/v1/search/programming-in-p5js", 1599458576);
// $ sqlite> select * from views;
// $ /api/v1/search/programming-in-p5js|1599458576

// GetStats returns page view for a URL.
func GetStats(db sql.DB, link string) int64 {
	rows, err := db.Query("SELECT COUNT(*) FROM views WHERE path = ?", link)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var cnt int64
		if err := rows.Scan(&cnt); err != nil {
			panic(err)
		}
		return cnt
	}
	return 0
}

// GetTotalViews returns total number of views for the website.
func GetTotalViews(db sql.DB) int64 {
	rows, err := db.Query("SELECT COUNT(*) FROM views")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var cnt int64
		if err := rows.Scan(&cnt); err != nil {
			panic(err)
		}
		return cnt
	}
	return 0
}

// UpdateDB writes <url, timestamp> into sqlite3 database.
func UpdateDB(db sql.DB, link string) {
	statement, _ := db.Prepare("insert into views values (?, ?)")
	unixTime := time.Now().Unix()
	_, err := statement.Exec(link, unixTime)
	if err != nil {
		panic(err)
	}
}
