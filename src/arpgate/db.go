package arpgate

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetConf(key string) string {
	db, err := sql.Open("sqlite3", "arpgate.db")
	if err != nil {
		return "ErROR - NO DB"
	}

	rows, err := db.Query("select val from config WHERE key='" + key + "'")
	if err != nil {
		log.Fatal(err)
	}
	var val string
	for rows.Next() {
		rows.Scan(&val)
	}
	return val
}

func TestDBExists() bool {

	db, err := sql.Open("sqlite3", "arpgate.db")
	if err != nil {
		return false
	}

	defer db.Close()

	rows, err := db.Query("select id, key, val, version from config")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var key string
		var val string
		var version int
		rows.Scan(&id, &key, &val, &version)
		fmt.Println(id, key, val, version)
	}
	rows.Close()

	return true
}
