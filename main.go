package main

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"
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

		err = util.CreateNewRequestTable(db)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		ru, _ := url.Parse("https://httpbin.org/anything")
		r := util.Request{
			Method:      "POST",
			Body:        "{}",
			ContentType: "application/json",
			URL:         *ru,
			Header:      make(http.Header),
		}

		resp, err := r.DoRequest()
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}
		log.Printf(resp.ContentType)
		log.Printf(resp.Body)

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
			os.Exit(3)
		}
		err = r.Insert(tx)
		if err != nil {
			log.Fatal(err)
			os.Exit(4)
		}
		tx.Commit()
	}
}
