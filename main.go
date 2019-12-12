package main

import (
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
)

func main () {
	for _, k := range os.Args[1:] {
		log.Println("k = " + k)
		db, err := sql.Open("sqlite3", k)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer db.Close()

		_, err = db.Exec("CREATE TABLE IF NOT EXISTS FOO (id integer not null primary key, name text); DELETE FROM FOO;")
		if err != nil {
			log.Fatal(err)
		}

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare("INSERT INTO FOO (ID, NAME) VALUES (?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		for i := 0; i < 1000; i++ {
			_, err := stmt.Exec(i, fmt.Sprintf("%03d-tarou", i))
			if err != nil {
				log.Fatal(i)
			}
		}
		tx.Commit()
	}
}
