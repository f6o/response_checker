package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"fmt"
	"github.com/f6o/response_checker/util"
)

func main() {
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
		dbutil := util.DBUtil{Tx: tx}
		for i := 0; i < 1000; i++ {
			dbutil.InsertToFoo(i, fmt.Sprintf("%03d-tarou", i))
		}
		tx.Commit()
	}
}
