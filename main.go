package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/f6o/response_checker/util"
	_ "github.com/mattn/go-sqlite3"
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

		err = util.CreateRequestTable(db)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
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
